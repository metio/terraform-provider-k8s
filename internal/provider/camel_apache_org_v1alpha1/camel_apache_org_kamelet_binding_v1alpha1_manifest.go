/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1alpha1

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
	_ datasource.DataSource = &CamelApacheOrgKameletBindingV1Alpha1Manifest{}
)

func NewCamelApacheOrgKameletBindingV1Alpha1Manifest() datasource.DataSource {
	return &CamelApacheOrgKameletBindingV1Alpha1Manifest{}
}

type CamelApacheOrgKameletBindingV1Alpha1Manifest struct{}

type CamelApacheOrgKameletBindingV1Alpha1ManifestData struct {
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
		ErrorHandler *map[string]string `tfsdk:"error_handler" json:"errorHandler,omitempty"`
		Integration  *struct {
			Configuration *[]struct {
				Type  *string `tfsdk:"type" json:"type,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"configuration" json:"configuration,omitempty"`
			Dependencies   *[]string            `tfsdk:"dependencies" json:"dependencies,omitempty"`
			Flows          *[]map[string]string `tfsdk:"flows" json:"flows,omitempty"`
			IntegrationKit *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"integration_kit" json:"integrationKit,omitempty"`
			Profile            *string   `tfsdk:"profile" json:"profile,omitempty"`
			Replicas           *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
			Repositories       *[]string `tfsdk:"repositories" json:"repositories,omitempty"`
			ServiceAccountName *string   `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
			Sources            *[]struct {
				Compression    *bool     `tfsdk:"compression" json:"compression,omitempty"`
				Content        *string   `tfsdk:"content" json:"content,omitempty"`
				ContentKey     *string   `tfsdk:"content_key" json:"contentKey,omitempty"`
				ContentRef     *string   `tfsdk:"content_ref" json:"contentRef,omitempty"`
				ContentType    *string   `tfsdk:"content_type" json:"contentType,omitempty"`
				From_kamelet   *bool     `tfsdk:"from_kamelet" json:"from-kamelet,omitempty"`
				Interceptors   *[]string `tfsdk:"interceptors" json:"interceptors,omitempty"`
				Language       *string   `tfsdk:"language" json:"language,omitempty"`
				Loader         *string   `tfsdk:"loader" json:"loader,omitempty"`
				Name           *string   `tfsdk:"name" json:"name,omitempty"`
				Path           *string   `tfsdk:"path" json:"path,omitempty"`
				Property_names *[]string `tfsdk:"property_names" json:"property-names,omitempty"`
				RawContent     *string   `tfsdk:"raw_content" json:"rawContent,omitempty"`
				Type           *string   `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"sources" json:"sources,omitempty"`
			Template *struct {
				Spec *struct {
					ActiveDeadlineSeconds        *int64 `tfsdk:"active_deadline_seconds" json:"activeDeadlineSeconds,omitempty"`
					AutomountServiceAccountToken *bool  `tfsdk:"automount_service_account_token" json:"automountServiceAccountToken,omitempty"`
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
					DnsPolicy           *string `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
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
					NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					RestartPolicy   *string            `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
					SecurityContext *struct {
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
					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
					TopologySpreadConstraints     *[]struct {
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
			Traits *struct {
				Threescale *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				} `tfsdk:"threescale" json:"3scale,omitempty"`
				Addons   *map[string]string `tfsdk:"addons" json:"addons,omitempty"`
				Affinity *struct {
					Configuration         *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled               *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					NodeAffinityLabels    *[]string          `tfsdk:"node_affinity_labels" json:"nodeAffinityLabels,omitempty"`
					PodAffinity           *bool              `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					PodAffinityLabels     *[]string          `tfsdk:"pod_affinity_labels" json:"podAffinityLabels,omitempty"`
					PodAntiAffinity       *bool              `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					PodAntiAffinityLabels *[]string          `tfsdk:"pod_anti_affinity_labels" json:"podAntiAffinityLabels,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				Builder *struct {
					Annotations           *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					BaseImage             *string            `tfsdk:"base_image" json:"baseImage,omitempty"`
					Configuration         *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled               *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					IncrementalImageBuild *bool              `tfsdk:"incremental_image_build" json:"incrementalImageBuild,omitempty"`
					LimitCPU              *string            `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
					LimitMemory           *string            `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
					MavenProfiles         *[]string          `tfsdk:"maven_profiles" json:"mavenProfiles,omitempty"`
					NodeSelector          *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					OrderStrategy         *string            `tfsdk:"order_strategy" json:"orderStrategy,omitempty"`
					Platforms             *[]string          `tfsdk:"platforms" json:"platforms,omitempty"`
					Properties            *[]string          `tfsdk:"properties" json:"properties,omitempty"`
					RequestCPU            *string            `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
					RequestMemory         *string            `tfsdk:"request_memory" json:"requestMemory,omitempty"`
					Strategy              *string            `tfsdk:"strategy" json:"strategy,omitempty"`
					Tasks                 *[]string          `tfsdk:"tasks" json:"tasks,omitempty"`
					TasksFilter           *string            `tfsdk:"tasks_filter" json:"tasksFilter,omitempty"`
					TasksLimitCPU         *[]string          `tfsdk:"tasks_limit_cpu" json:"tasksLimitCPU,omitempty"`
					TasksLimitMemory      *[]string          `tfsdk:"tasks_limit_memory" json:"tasksLimitMemory,omitempty"`
					TasksRequestCPU       *[]string          `tfsdk:"tasks_request_cpu" json:"tasksRequestCPU,omitempty"`
					TasksRequestMemory    *[]string          `tfsdk:"tasks_request_memory" json:"tasksRequestMemory,omitempty"`
					Verbose               *bool              `tfsdk:"verbose" json:"verbose,omitempty"`
				} `tfsdk:"builder" json:"builder,omitempty"`
				Camel *struct {
					Configuration  *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Properties     *[]string          `tfsdk:"properties" json:"properties,omitempty"`
					RuntimeVersion *string            `tfsdk:"runtime_version" json:"runtimeVersion,omitempty"`
				} `tfsdk:"camel" json:"camel,omitempty"`
				Container *struct {
					AllowPrivilegeEscalation *bool              `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
					Auto                     *bool              `tfsdk:"auto" json:"auto,omitempty"`
					CapabilitiesAdd          *[]string          `tfsdk:"capabilities_add" json:"capabilitiesAdd,omitempty"`
					CapabilitiesDrop         *[]string          `tfsdk:"capabilities_drop" json:"capabilitiesDrop,omitempty"`
					Configuration            *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled                  *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Expose                   *bool              `tfsdk:"expose" json:"expose,omitempty"`
					Image                    *string            `tfsdk:"image" json:"image,omitempty"`
					ImagePullPolicy          *string            `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
					LimitCPU                 *string            `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
					LimitMemory              *string            `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
					Name                     *string            `tfsdk:"name" json:"name,omitempty"`
					Port                     *int64             `tfsdk:"port" json:"port,omitempty"`
					PortName                 *string            `tfsdk:"port_name" json:"portName,omitempty"`
					RequestCPU               *string            `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
					RequestMemory            *string            `tfsdk:"request_memory" json:"requestMemory,omitempty"`
					RunAsNonRoot             *bool              `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
					RunAsUser                *int64             `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					SeccompProfileType       *string            `tfsdk:"seccomp_profile_type" json:"seccompProfileType,omitempty"`
					ServicePort              *int64             `tfsdk:"service_port" json:"servicePort,omitempty"`
					ServicePortName          *string            `tfsdk:"service_port_name" json:"servicePortName,omitempty"`
				} `tfsdk:"container" json:"container,omitempty"`
				Cron *struct {
					ActiveDeadlineSeconds   *int64             `tfsdk:"active_deadline_seconds" json:"activeDeadlineSeconds,omitempty"`
					Auto                    *bool              `tfsdk:"auto" json:"auto,omitempty"`
					BackoffLimit            *int64             `tfsdk:"backoff_limit" json:"backoffLimit,omitempty"`
					Components              *string            `tfsdk:"components" json:"components,omitempty"`
					ConcurrencyPolicy       *string            `tfsdk:"concurrency_policy" json:"concurrencyPolicy,omitempty"`
					Configuration           *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled                 *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Fallback                *bool              `tfsdk:"fallback" json:"fallback,omitempty"`
					Schedule                *string            `tfsdk:"schedule" json:"schedule,omitempty"`
					StartingDeadlineSeconds *int64             `tfsdk:"starting_deadline_seconds" json:"startingDeadlineSeconds,omitempty"`
					TimeZone                *string            `tfsdk:"time_zone" json:"timeZone,omitempty"`
				} `tfsdk:"cron" json:"cron,omitempty"`
				Dependencies *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"dependencies" json:"dependencies,omitempty"`
				Deployer *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Kind          *string            `tfsdk:"kind" json:"kind,omitempty"`
					UseSSA        *bool              `tfsdk:"use_ssa" json:"useSSA,omitempty"`
				} `tfsdk:"deployer" json:"deployer,omitempty"`
				Deployment *struct {
					Configuration               *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled                     *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					ProgressDeadlineSeconds     *int64             `tfsdk:"progress_deadline_seconds" json:"progressDeadlineSeconds,omitempty"`
					RollingUpdateMaxSurge       *string            `tfsdk:"rolling_update_max_surge" json:"rollingUpdateMaxSurge,omitempty"`
					RollingUpdateMaxUnavailable *string            `tfsdk:"rolling_update_max_unavailable" json:"rollingUpdateMaxUnavailable,omitempty"`
					Strategy                    *string            `tfsdk:"strategy" json:"strategy,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
				Environment *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					ContainerMeta *bool              `tfsdk:"container_meta" json:"containerMeta,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					HttpProxy     *bool              `tfsdk:"http_proxy" json:"httpProxy,omitempty"`
					Vars          *[]string          `tfsdk:"vars" json:"vars,omitempty"`
				} `tfsdk:"environment" json:"environment,omitempty"`
				Error_handler *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Ref           *string            `tfsdk:"ref" json:"ref,omitempty"`
				} `tfsdk:"error_handler" json:"error-handler,omitempty"`
				Gc *struct {
					Configuration  *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					DiscoveryCache *string            `tfsdk:"discovery_cache" json:"discoveryCache,omitempty"`
					Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"gc" json:"gc,omitempty"`
				Health *struct {
					Configuration             *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled                   *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					LivenessFailureThreshold  *int64             `tfsdk:"liveness_failure_threshold" json:"livenessFailureThreshold,omitempty"`
					LivenessInitialDelay      *int64             `tfsdk:"liveness_initial_delay" json:"livenessInitialDelay,omitempty"`
					LivenessPeriod            *int64             `tfsdk:"liveness_period" json:"livenessPeriod,omitempty"`
					LivenessProbe             *string            `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
					LivenessProbeEnabled      *bool              `tfsdk:"liveness_probe_enabled" json:"livenessProbeEnabled,omitempty"`
					LivenessScheme            *string            `tfsdk:"liveness_scheme" json:"livenessScheme,omitempty"`
					LivenessSuccessThreshold  *int64             `tfsdk:"liveness_success_threshold" json:"livenessSuccessThreshold,omitempty"`
					LivenessTimeout           *int64             `tfsdk:"liveness_timeout" json:"livenessTimeout,omitempty"`
					ReadinessFailureThreshold *int64             `tfsdk:"readiness_failure_threshold" json:"readinessFailureThreshold,omitempty"`
					ReadinessInitialDelay     *int64             `tfsdk:"readiness_initial_delay" json:"readinessInitialDelay,omitempty"`
					ReadinessPeriod           *int64             `tfsdk:"readiness_period" json:"readinessPeriod,omitempty"`
					ReadinessProbe            *string            `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
					ReadinessProbeEnabled     *bool              `tfsdk:"readiness_probe_enabled" json:"readinessProbeEnabled,omitempty"`
					ReadinessScheme           *string            `tfsdk:"readiness_scheme" json:"readinessScheme,omitempty"`
					ReadinessSuccessThreshold *int64             `tfsdk:"readiness_success_threshold" json:"readinessSuccessThreshold,omitempty"`
					ReadinessTimeout          *int64             `tfsdk:"readiness_timeout" json:"readinessTimeout,omitempty"`
					StartupFailureThreshold   *int64             `tfsdk:"startup_failure_threshold" json:"startupFailureThreshold,omitempty"`
					StartupInitialDelay       *int64             `tfsdk:"startup_initial_delay" json:"startupInitialDelay,omitempty"`
					StartupPeriod             *int64             `tfsdk:"startup_period" json:"startupPeriod,omitempty"`
					StartupProbe              *string            `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
					StartupProbeEnabled       *bool              `tfsdk:"startup_probe_enabled" json:"startupProbeEnabled,omitempty"`
					StartupScheme             *string            `tfsdk:"startup_scheme" json:"startupScheme,omitempty"`
					StartupSuccessThreshold   *int64             `tfsdk:"startup_success_threshold" json:"startupSuccessThreshold,omitempty"`
					StartupTimeout            *int64             `tfsdk:"startup_timeout" json:"startupTimeout,omitempty"`
				} `tfsdk:"health" json:"health,omitempty"`
				Ingress *struct {
					Annotations      *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Auto             *bool              `tfsdk:"auto" json:"auto,omitempty"`
					Configuration    *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled          *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Host             *string            `tfsdk:"host" json:"host,omitempty"`
					IngressClassName *string            `tfsdk:"ingress_class_name" json:"ingressClassName,omitempty"`
					Path             *string            `tfsdk:"path" json:"path,omitempty"`
					PathType         *string            `tfsdk:"path_type" json:"pathType,omitempty"`
					TlsHosts         *[]string          `tfsdk:"tls_hosts" json:"tlsHosts,omitempty"`
					TlsSecretName    *string            `tfsdk:"tls_secret_name" json:"tlsSecretName,omitempty"`
				} `tfsdk:"ingress" json:"ingress,omitempty"`
				Istio *struct {
					Allow         *string            `tfsdk:"allow" json:"allow,omitempty"`
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Inject        *bool              `tfsdk:"inject" json:"inject,omitempty"`
				} `tfsdk:"istio" json:"istio,omitempty"`
				Jolokia *struct {
					CACert                     *string            `tfsdk:"ca_cert" json:"CACert,omitempty"`
					ClientPrincipal            *[]string          `tfsdk:"client_principal" json:"clientPrincipal,omitempty"`
					Configuration              *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					DiscoveryEnabled           *bool              `tfsdk:"discovery_enabled" json:"discoveryEnabled,omitempty"`
					Enabled                    *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					ExtendedClientCheck        *bool              `tfsdk:"extended_client_check" json:"extendedClientCheck,omitempty"`
					Host                       *string            `tfsdk:"host" json:"host,omitempty"`
					Options                    *[]string          `tfsdk:"options" json:"options,omitempty"`
					Password                   *string            `tfsdk:"password" json:"password,omitempty"`
					Port                       *int64             `tfsdk:"port" json:"port,omitempty"`
					Protocol                   *string            `tfsdk:"protocol" json:"protocol,omitempty"`
					UseSSLClientAuthentication *bool              `tfsdk:"use_ssl_client_authentication" json:"useSSLClientAuthentication,omitempty"`
					User                       *string            `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"jolokia" json:"jolokia,omitempty"`
				Jvm *struct {
					Classpath     *string            `tfsdk:"classpath" json:"classpath,omitempty"`
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Debug         *bool              `tfsdk:"debug" json:"debug,omitempty"`
					DebugAddress  *string            `tfsdk:"debug_address" json:"debugAddress,omitempty"`
					DebugSuspend  *bool              `tfsdk:"debug_suspend" json:"debugSuspend,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Jar           *string            `tfsdk:"jar" json:"jar,omitempty"`
					Options       *[]string          `tfsdk:"options" json:"options,omitempty"`
					PrintCommand  *bool              `tfsdk:"print_command" json:"printCommand,omitempty"`
				} `tfsdk:"jvm" json:"jvm,omitempty"`
				Kamelets *struct {
					Auto          *bool              `tfsdk:"auto" json:"auto,omitempty"`
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					List          *string            `tfsdk:"list" json:"list,omitempty"`
					MountPoint    *string            `tfsdk:"mount_point" json:"mountPoint,omitempty"`
				} `tfsdk:"kamelets" json:"kamelets,omitempty"`
				Keda *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				} `tfsdk:"keda" json:"keda,omitempty"`
				Knative *struct {
					Auto                 *bool              `tfsdk:"auto" json:"auto,omitempty"`
					ChannelSinks         *[]string          `tfsdk:"channel_sinks" json:"channelSinks,omitempty"`
					ChannelSources       *[]string          `tfsdk:"channel_sources" json:"channelSources,omitempty"`
					Config               *string            `tfsdk:"config" json:"config,omitempty"`
					Configuration        *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled              *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					EndpointSinks        *[]string          `tfsdk:"endpoint_sinks" json:"endpointSinks,omitempty"`
					EndpointSources      *[]string          `tfsdk:"endpoint_sources" json:"endpointSources,omitempty"`
					EventSinks           *[]string          `tfsdk:"event_sinks" json:"eventSinks,omitempty"`
					EventSources         *[]string          `tfsdk:"event_sources" json:"eventSources,omitempty"`
					FilterEventType      *bool              `tfsdk:"filter_event_type" json:"filterEventType,omitempty"`
					FilterSourceChannels *bool              `tfsdk:"filter_source_channels" json:"filterSourceChannels,omitempty"`
					Filters              *[]string          `tfsdk:"filters" json:"filters,omitempty"`
					NamespaceLabel       *bool              `tfsdk:"namespace_label" json:"namespaceLabel,omitempty"`
					SinkBinding          *bool              `tfsdk:"sink_binding" json:"sinkBinding,omitempty"`
				} `tfsdk:"knative" json:"knative,omitempty"`
				Knative_service *struct {
					Annotations       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Auto              *bool              `tfsdk:"auto" json:"auto,omitempty"`
					AutoscalingMetric *string            `tfsdk:"autoscaling_metric" json:"autoscalingMetric,omitempty"`
					AutoscalingTarget *int64             `tfsdk:"autoscaling_target" json:"autoscalingTarget,omitempty"`
					Class             *string            `tfsdk:"class" json:"class,omitempty"`
					Configuration     *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled           *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					MaxScale          *int64             `tfsdk:"max_scale" json:"maxScale,omitempty"`
					MinScale          *int64             `tfsdk:"min_scale" json:"minScale,omitempty"`
					RolloutDuration   *string            `tfsdk:"rollout_duration" json:"rolloutDuration,omitempty"`
					TimeoutSeconds    *int64             `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
					Visibility        *string            `tfsdk:"visibility" json:"visibility,omitempty"`
				} `tfsdk:"knative_service" json:"knative-service,omitempty"`
				Logging *struct {
					Color           *bool              `tfsdk:"color" json:"color,omitempty"`
					Configuration   *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled         *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Format          *string            `tfsdk:"format" json:"format,omitempty"`
					Json            *bool              `tfsdk:"json" json:"json,omitempty"`
					JsonPrettyPrint *bool              `tfsdk:"json_pretty_print" json:"jsonPrettyPrint,omitempty"`
					Level           *string            `tfsdk:"level" json:"level,omitempty"`
				} `tfsdk:"logging" json:"logging,omitempty"`
				Master *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				} `tfsdk:"master" json:"master,omitempty"`
				Mount *struct {
					Configs                          *[]string          `tfsdk:"configs" json:"configs,omitempty"`
					Configuration                    *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					EmptyDirs                        *[]string          `tfsdk:"empty_dirs" json:"emptyDirs,omitempty"`
					Enabled                          *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					HotReload                        *bool              `tfsdk:"hot_reload" json:"hotReload,omitempty"`
					Resources                        *[]string          `tfsdk:"resources" json:"resources,omitempty"`
					ScanKameletsImplicitLabelSecrets *bool              `tfsdk:"scan_kamelets_implicit_label_secrets" json:"scanKameletsImplicitLabelSecrets,omitempty"`
					Volumes                          *[]string          `tfsdk:"volumes" json:"volumes,omitempty"`
				} `tfsdk:"mount" json:"mount,omitempty"`
				Openapi *struct {
					Configmaps    *[]string          `tfsdk:"configmaps" json:"configmaps,omitempty"`
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"openapi" json:"openapi,omitempty"`
				Owner *struct {
					Configuration     *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled           *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					TargetAnnotations *[]string          `tfsdk:"target_annotations" json:"targetAnnotations,omitempty"`
					TargetLabels      *[]string          `tfsdk:"target_labels" json:"targetLabels,omitempty"`
				} `tfsdk:"owner" json:"owner,omitempty"`
				Pdb *struct {
					Configuration  *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					MaxUnavailable *string            `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					MinAvailable   *string            `tfsdk:"min_available" json:"minAvailable,omitempty"`
				} `tfsdk:"pdb" json:"pdb,omitempty"`
				Platform *struct {
					Auto          *bool              `tfsdk:"auto" json:"auto,omitempty"`
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					CreateDefault *bool              `tfsdk:"create_default" json:"createDefault,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Global        *bool              `tfsdk:"global" json:"global,omitempty"`
				} `tfsdk:"platform" json:"platform,omitempty"`
				Pod *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"pod" json:"pod,omitempty"`
				Prometheus *struct {
					Configuration    *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled          *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					PodMonitor       *bool              `tfsdk:"pod_monitor" json:"podMonitor,omitempty"`
					PodMonitorLabels *[]string          `tfsdk:"pod_monitor_labels" json:"podMonitorLabels,omitempty"`
				} `tfsdk:"prometheus" json:"prometheus,omitempty"`
				Pull_secret *struct {
					Auto                  *bool              `tfsdk:"auto" json:"auto,omitempty"`
					Configuration         *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled               *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					ImagePullerDelegation *bool              `tfsdk:"image_puller_delegation" json:"imagePullerDelegation,omitempty"`
					SecretName            *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
				} `tfsdk:"pull_secret" json:"pull-secret,omitempty"`
				Quarkus *struct {
					BuildMode          *[]string          `tfsdk:"build_mode" json:"buildMode,omitempty"`
					Configuration      *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled            *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					NativeBaseImage    *string            `tfsdk:"native_base_image" json:"nativeBaseImage,omitempty"`
					NativeBuilderImage *string            `tfsdk:"native_builder_image" json:"nativeBuilderImage,omitempty"`
					PackageTypes       *[]string          `tfsdk:"package_types" json:"packageTypes,omitempty"`
				} `tfsdk:"quarkus" json:"quarkus,omitempty"`
				Registry *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"registry" json:"registry,omitempty"`
				Route *struct {
					Annotations                       *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Configuration                     *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled                           *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Host                              *string            `tfsdk:"host" json:"host,omitempty"`
					TlsCACertificate                  *string            `tfsdk:"tls_ca_certificate" json:"tlsCACertificate,omitempty"`
					TlsCACertificateSecret            *string            `tfsdk:"tls_ca_certificate_secret" json:"tlsCACertificateSecret,omitempty"`
					TlsCertificate                    *string            `tfsdk:"tls_certificate" json:"tlsCertificate,omitempty"`
					TlsCertificateSecret              *string            `tfsdk:"tls_certificate_secret" json:"tlsCertificateSecret,omitempty"`
					TlsDestinationCACertificate       *string            `tfsdk:"tls_destination_ca_certificate" json:"tlsDestinationCACertificate,omitempty"`
					TlsDestinationCACertificateSecret *string            `tfsdk:"tls_destination_ca_certificate_secret" json:"tlsDestinationCACertificateSecret,omitempty"`
					TlsInsecureEdgeTerminationPolicy  *string            `tfsdk:"tls_insecure_edge_termination_policy" json:"tlsInsecureEdgeTerminationPolicy,omitempty"`
					TlsKey                            *string            `tfsdk:"tls_key" json:"tlsKey,omitempty"`
					TlsKeySecret                      *string            `tfsdk:"tls_key_secret" json:"tlsKeySecret,omitempty"`
					TlsTermination                    *string            `tfsdk:"tls_termination" json:"tlsTermination,omitempty"`
				} `tfsdk:"route" json:"route,omitempty"`
				Security_context *struct {
					Configuration      *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled            *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					RunAsNonRoot       *bool              `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
					RunAsUser          *int64             `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
					SeccompProfileType *string            `tfsdk:"seccomp_profile_type" json:"seccompProfileType,omitempty"`
				} `tfsdk:"security_context" json:"security-context,omitempty"`
				Service *struct {
					Auto          *bool              `tfsdk:"auto" json:"auto,omitempty"`
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					NodePort      *bool              `tfsdk:"node_port" json:"nodePort,omitempty"`
					Type          *string            `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"service" json:"service,omitempty"`
				Service_binding *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Services      *[]string          `tfsdk:"services" json:"services,omitempty"`
				} `tfsdk:"service_binding" json:"service-binding,omitempty"`
				Strimzi *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				} `tfsdk:"strimzi" json:"strimzi,omitempty"`
				Telemetry *struct {
					Auto                 *bool              `tfsdk:"auto" json:"auto,omitempty"`
					Configuration        *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled              *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Endpoint             *string            `tfsdk:"endpoint" json:"endpoint,omitempty"`
					Sampler              *string            `tfsdk:"sampler" json:"sampler,omitempty"`
					Sampler_parent_based *bool              `tfsdk:"sampler_parent_based" json:"sampler-parent-based,omitempty"`
					Sampler_ratio        *string            `tfsdk:"sampler_ratio" json:"sampler-ratio,omitempty"`
					ServiceName          *string            `tfsdk:"service_name" json:"serviceName,omitempty"`
				} `tfsdk:"telemetry" json:"telemetry,omitempty"`
				Toleration *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
					Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
					Taints        *[]string          `tfsdk:"taints" json:"taints,omitempty"`
				} `tfsdk:"toleration" json:"toleration,omitempty"`
				Tracing *struct {
					Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				} `tfsdk:"tracing" json:"tracing,omitempty"`
			} `tfsdk:"traits" json:"traits,omitempty"`
		} `tfsdk:"integration" json:"integration,omitempty"`
		Replicas           *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		Sink               *struct {
			DataTypes *struct {
				Format *string `tfsdk:"format" json:"format,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"data_types" json:"dataTypes,omitempty"`
			Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
			Ref        *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"ref" json:"ref,omitempty"`
			Types *struct {
				MediaType *string `tfsdk:"media_type" json:"mediaType,omitempty"`
				Schema    *struct {
					Dollarschema *string            `tfsdk:"dollarschema" json:"$schema,omitempty"`
					Description  *string            `tfsdk:"description" json:"description,omitempty"`
					Example      *map[string]string `tfsdk:"example" json:"example,omitempty"`
					ExternalDocs *struct {
						Description *string `tfsdk:"description" json:"description,omitempty"`
						Url         *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"external_docs" json:"externalDocs,omitempty"`
					Id         *string `tfsdk:"id" json:"id,omitempty"`
					Properties *struct {
						Default          *map[string]string `tfsdk:"default" json:"default,omitempty"`
						Deprecated       *bool              `tfsdk:"deprecated" json:"deprecated,omitempty"`
						Description      *string            `tfsdk:"description" json:"description,omitempty"`
						Enum             *[]string          `tfsdk:"enum" json:"enum,omitempty"`
						Example          *map[string]string `tfsdk:"example" json:"example,omitempty"`
						ExclusiveMaximum *bool              `tfsdk:"exclusive_maximum" json:"exclusiveMaximum,omitempty"`
						ExclusiveMinimum *bool              `tfsdk:"exclusive_minimum" json:"exclusiveMinimum,omitempty"`
						Format           *string            `tfsdk:"format" json:"format,omitempty"`
						Id               *string            `tfsdk:"id" json:"id,omitempty"`
						MaxItems         *int64             `tfsdk:"max_items" json:"maxItems,omitempty"`
						MaxLength        *int64             `tfsdk:"max_length" json:"maxLength,omitempty"`
						MaxProperties    *int64             `tfsdk:"max_properties" json:"maxProperties,omitempty"`
						Maximum          *string            `tfsdk:"maximum" json:"maximum,omitempty"`
						MinItems         *int64             `tfsdk:"min_items" json:"minItems,omitempty"`
						MinLength        *int64             `tfsdk:"min_length" json:"minLength,omitempty"`
						MinProperties    *int64             `tfsdk:"min_properties" json:"minProperties,omitempty"`
						Minimum          *string            `tfsdk:"minimum" json:"minimum,omitempty"`
						MultipleOf       *string            `tfsdk:"multiple_of" json:"multipleOf,omitempty"`
						Nullable         *bool              `tfsdk:"nullable" json:"nullable,omitempty"`
						Pattern          *string            `tfsdk:"pattern" json:"pattern,omitempty"`
						Title            *string            `tfsdk:"title" json:"title,omitempty"`
						Type             *string            `tfsdk:"type" json:"type,omitempty"`
						UniqueItems      *bool              `tfsdk:"unique_items" json:"uniqueItems,omitempty"`
						X_descriptors    *[]string          `tfsdk:"x_descriptors" json:"x-descriptors,omitempty"`
					} `tfsdk:"properties" json:"properties,omitempty"`
					Required *[]string `tfsdk:"required" json:"required,omitempty"`
					Title    *string   `tfsdk:"title" json:"title,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"schema" json:"schema,omitempty"`
			} `tfsdk:"types" json:"types,omitempty"`
			Uri *string `tfsdk:"uri" json:"uri,omitempty"`
		} `tfsdk:"sink" json:"sink,omitempty"`
		Source *struct {
			DataTypes *struct {
				Format *string `tfsdk:"format" json:"format,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"data_types" json:"dataTypes,omitempty"`
			Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
			Ref        *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"ref" json:"ref,omitempty"`
			Types *struct {
				MediaType *string `tfsdk:"media_type" json:"mediaType,omitempty"`
				Schema    *struct {
					Dollarschema *string            `tfsdk:"dollarschema" json:"$schema,omitempty"`
					Description  *string            `tfsdk:"description" json:"description,omitempty"`
					Example      *map[string]string `tfsdk:"example" json:"example,omitempty"`
					ExternalDocs *struct {
						Description *string `tfsdk:"description" json:"description,omitempty"`
						Url         *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"external_docs" json:"externalDocs,omitempty"`
					Id         *string `tfsdk:"id" json:"id,omitempty"`
					Properties *struct {
						Default          *map[string]string `tfsdk:"default" json:"default,omitempty"`
						Deprecated       *bool              `tfsdk:"deprecated" json:"deprecated,omitempty"`
						Description      *string            `tfsdk:"description" json:"description,omitempty"`
						Enum             *[]string          `tfsdk:"enum" json:"enum,omitempty"`
						Example          *map[string]string `tfsdk:"example" json:"example,omitempty"`
						ExclusiveMaximum *bool              `tfsdk:"exclusive_maximum" json:"exclusiveMaximum,omitempty"`
						ExclusiveMinimum *bool              `tfsdk:"exclusive_minimum" json:"exclusiveMinimum,omitempty"`
						Format           *string            `tfsdk:"format" json:"format,omitempty"`
						Id               *string            `tfsdk:"id" json:"id,omitempty"`
						MaxItems         *int64             `tfsdk:"max_items" json:"maxItems,omitempty"`
						MaxLength        *int64             `tfsdk:"max_length" json:"maxLength,omitempty"`
						MaxProperties    *int64             `tfsdk:"max_properties" json:"maxProperties,omitempty"`
						Maximum          *string            `tfsdk:"maximum" json:"maximum,omitempty"`
						MinItems         *int64             `tfsdk:"min_items" json:"minItems,omitempty"`
						MinLength        *int64             `tfsdk:"min_length" json:"minLength,omitempty"`
						MinProperties    *int64             `tfsdk:"min_properties" json:"minProperties,omitempty"`
						Minimum          *string            `tfsdk:"minimum" json:"minimum,omitempty"`
						MultipleOf       *string            `tfsdk:"multiple_of" json:"multipleOf,omitempty"`
						Nullable         *bool              `tfsdk:"nullable" json:"nullable,omitempty"`
						Pattern          *string            `tfsdk:"pattern" json:"pattern,omitempty"`
						Title            *string            `tfsdk:"title" json:"title,omitempty"`
						Type             *string            `tfsdk:"type" json:"type,omitempty"`
						UniqueItems      *bool              `tfsdk:"unique_items" json:"uniqueItems,omitempty"`
						X_descriptors    *[]string          `tfsdk:"x_descriptors" json:"x-descriptors,omitempty"`
					} `tfsdk:"properties" json:"properties,omitempty"`
					Required *[]string `tfsdk:"required" json:"required,omitempty"`
					Title    *string   `tfsdk:"title" json:"title,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"schema" json:"schema,omitempty"`
			} `tfsdk:"types" json:"types,omitempty"`
			Uri *string `tfsdk:"uri" json:"uri,omitempty"`
		} `tfsdk:"source" json:"source,omitempty"`
		Steps *[]struct {
			DataTypes *struct {
				Format *string `tfsdk:"format" json:"format,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"data_types" json:"dataTypes,omitempty"`
			Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
			Ref        *struct {
				ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
				FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
				Name            *string `tfsdk:"name" json:"name,omitempty"`
				Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
				Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
			} `tfsdk:"ref" json:"ref,omitempty"`
			Types *struct {
				MediaType *string `tfsdk:"media_type" json:"mediaType,omitempty"`
				Schema    *struct {
					Dollarschema *string            `tfsdk:"dollarschema" json:"$schema,omitempty"`
					Description  *string            `tfsdk:"description" json:"description,omitempty"`
					Example      *map[string]string `tfsdk:"example" json:"example,omitempty"`
					ExternalDocs *struct {
						Description *string `tfsdk:"description" json:"description,omitempty"`
						Url         *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"external_docs" json:"externalDocs,omitempty"`
					Id         *string `tfsdk:"id" json:"id,omitempty"`
					Properties *struct {
						Default          *map[string]string `tfsdk:"default" json:"default,omitempty"`
						Deprecated       *bool              `tfsdk:"deprecated" json:"deprecated,omitempty"`
						Description      *string            `tfsdk:"description" json:"description,omitempty"`
						Enum             *[]string          `tfsdk:"enum" json:"enum,omitempty"`
						Example          *map[string]string `tfsdk:"example" json:"example,omitempty"`
						ExclusiveMaximum *bool              `tfsdk:"exclusive_maximum" json:"exclusiveMaximum,omitempty"`
						ExclusiveMinimum *bool              `tfsdk:"exclusive_minimum" json:"exclusiveMinimum,omitempty"`
						Format           *string            `tfsdk:"format" json:"format,omitempty"`
						Id               *string            `tfsdk:"id" json:"id,omitempty"`
						MaxItems         *int64             `tfsdk:"max_items" json:"maxItems,omitempty"`
						MaxLength        *int64             `tfsdk:"max_length" json:"maxLength,omitempty"`
						MaxProperties    *int64             `tfsdk:"max_properties" json:"maxProperties,omitempty"`
						Maximum          *string            `tfsdk:"maximum" json:"maximum,omitempty"`
						MinItems         *int64             `tfsdk:"min_items" json:"minItems,omitempty"`
						MinLength        *int64             `tfsdk:"min_length" json:"minLength,omitempty"`
						MinProperties    *int64             `tfsdk:"min_properties" json:"minProperties,omitempty"`
						Minimum          *string            `tfsdk:"minimum" json:"minimum,omitempty"`
						MultipleOf       *string            `tfsdk:"multiple_of" json:"multipleOf,omitempty"`
						Nullable         *bool              `tfsdk:"nullable" json:"nullable,omitempty"`
						Pattern          *string            `tfsdk:"pattern" json:"pattern,omitempty"`
						Title            *string            `tfsdk:"title" json:"title,omitempty"`
						Type             *string            `tfsdk:"type" json:"type,omitempty"`
						UniqueItems      *bool              `tfsdk:"unique_items" json:"uniqueItems,omitempty"`
						X_descriptors    *[]string          `tfsdk:"x_descriptors" json:"x-descriptors,omitempty"`
					} `tfsdk:"properties" json:"properties,omitempty"`
					Required *[]string `tfsdk:"required" json:"required,omitempty"`
					Title    *string   `tfsdk:"title" json:"title,omitempty"`
					Type     *string   `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"schema" json:"schema,omitempty"`
			} `tfsdk:"types" json:"types,omitempty"`
			Uri *string `tfsdk:"uri" json:"uri,omitempty"`
		} `tfsdk:"steps" json:"steps,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CamelApacheOrgKameletBindingV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_kamelet_binding_v1alpha1_manifest"
}

func (r *CamelApacheOrgKameletBindingV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "KameletBinding is the Schema for the kamelets binding API.",
		MarkdownDescription: "KameletBinding is the Schema for the kamelets binding API.",
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
				Description:         "the specification of a KameletBinding",
				MarkdownDescription: "the specification of a KameletBinding",
				Attributes: map[string]schema.Attribute{
					"error_handler": schema.MapAttribute{
						Description:         "ErrorHandler is an optional handler called upon an error occurring in the integration",
						MarkdownDescription: "ErrorHandler is an optional handler called upon an error occurring in the integration",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"integration": schema.SingleNestedAttribute{
						Description:         "Integration is an optional integration used to specify custom parameters",
						MarkdownDescription: "Integration is an optional integration used to specify custom parameters",
						Attributes: map[string]schema.Attribute{
							"configuration": schema.ListNestedAttribute{
								Description:         "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes",
								MarkdownDescription: "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											Description:         "represents the type of configuration, ie: property, configmap, secret, ...",
											MarkdownDescription: "represents the type of configuration, ie: property, configmap, secret, ...",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "the value to assign to the configuration (syntax may vary depending on the 'Type')",
											MarkdownDescription: "the value to assign to the configuration (syntax may vary depending on the 'Type')",
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

							"dependencies": schema.ListAttribute{
								Description:         "the list of Camel or Maven dependencies required by the Integration",
								MarkdownDescription: "the list of Camel or Maven dependencies required by the Integration",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"flows": schema.ListAttribute{
								Description:         "a source in YAML DSL language which contain the routes to run",
								MarkdownDescription: "a source in YAML DSL language which contain the routes to run",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"integration_kit": schema.SingleNestedAttribute{
								Description:         "the reference of the 'IntegrationKit' which is used for this Integration",
								MarkdownDescription: "the reference of the 'IntegrationKit' which is used for this Integration",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"field_path": schema.StringAttribute{
										Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"profile": schema.StringAttribute{
								Description:         "the profile needed to run this Integration",
								MarkdownDescription: "the profile needed to run this Integration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "the number of 'Pods' needed for the running Integration",
								MarkdownDescription: "the number of 'Pods' needed for the running Integration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repositories": schema.ListAttribute{
								Description:         "additional Maven repositories to be used",
								MarkdownDescription: "additional Maven repositories to be used",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account_name": schema.StringAttribute{
								Description:         "custom SA to use for the Integration",
								MarkdownDescription: "custom SA to use for the Integration",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sources": schema.ListNestedAttribute{
								Description:         "the sources which contain the Camel routes to run",
								MarkdownDescription: "the sources which contain the Camel routes to run",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"compression": schema.BoolAttribute{
											Description:         "if the content is compressed (base64 encrypted)",
											MarkdownDescription: "if the content is compressed (base64 encrypted)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"content": schema.StringAttribute{
											Description:         "the source code (plain text)",
											MarkdownDescription: "the source code (plain text)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"content_key": schema.StringAttribute{
											Description:         "the confimap key holding the source content",
											MarkdownDescription: "the confimap key holding the source content",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"content_ref": schema.StringAttribute{
											Description:         "the confimap reference holding the source content",
											MarkdownDescription: "the confimap reference holding the source content",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"content_type": schema.StringAttribute{
											Description:         "the content type (tipically text or binary)",
											MarkdownDescription: "the content type (tipically text or binary)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"from_kamelet": schema.BoolAttribute{
											Description:         "True if the spec is generated from a Kamelet",
											MarkdownDescription: "True if the spec is generated from a Kamelet",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"interceptors": schema.ListAttribute{
											Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
											MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"language": schema.StringAttribute{
											Description:         "specify which is the language (Camel DSL) used to interpret this source code",
											MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"loader": schema.StringAttribute{
											Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
											MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "the name of the specification",
											MarkdownDescription: "the name of the specification",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "the path where the file is stored",
											MarkdownDescription: "the path where the file is stored",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"property_names": schema.ListAttribute{
											Description:         "List of property names defined in the source (e.g. if type is 'template')",
											MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"raw_content": schema.StringAttribute{
											Description:         "the source code (binary)",
											MarkdownDescription: "the source code (binary)",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												validators.Base64Validator(),
											},
										},

										"type": schema.StringAttribute{
											Description:         "Type defines the kind of source described by this object",
											MarkdownDescription: "Type defines the kind of source described by this object",
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

							"template": schema.SingleNestedAttribute{
								Description:         "Pod template customization",
								MarkdownDescription: "Pod template customization",
								Attributes: map[string]schema.Attribute{
									"spec": schema.SingleNestedAttribute{
										Description:         "the specification",
										MarkdownDescription: "the specification",
										Attributes: map[string]schema.Attribute{
											"active_deadline_seconds": schema.Int64Attribute{
												Description:         "ActiveDeadlineSeconds",
												MarkdownDescription: "ActiveDeadlineSeconds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"automount_service_account_token": schema.BoolAttribute{
												Description:         "AutomountServiceAccountToken",
												MarkdownDescription: "AutomountServiceAccountToken",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"containers": schema.ListNestedAttribute{
												Description:         "Containers",
												MarkdownDescription: "Containers",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"args": schema.ListAttribute{
															Description:         "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															MarkdownDescription: "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"command": schema.ListAttribute{
															Description:         "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															MarkdownDescription: "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"env": schema.ListNestedAttribute{
															Description:         "List of environment variables to set in the container. Cannot be updated.",
															MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",
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
																		Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																		MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "Specify whether the ConfigMap or its key must be defined",
																						MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
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
																				Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				Attributes: map[string]schema.Attribute{
																					"api_version": schema.StringAttribute{
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"field_path": schema.StringAttribute{
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",
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
																				Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				Attributes: map[string]schema.Attribute{
																					"container_name": schema.StringAttribute{
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"divisor": schema.StringAttribute{
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"resource": schema.StringAttribute{
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",
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
																				Description:         "Selects a key of a secret in the pod's namespace",
																				MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key of the secret to select from. Must be a valid secret key.",
																						MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

														"image": schema.StringAttribute{
															Description:         "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
															MarkdownDescription: "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_policy": schema.StringAttribute{
															Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
															MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
															Required:            false,
															Optional:            true,
															Computed:            false,
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
																					Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																					MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																			Description:         "HTTPGet specifies the http request to perform.",
																			MarkdownDescription: "HTTPGet specifies the http request to perform.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																					MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "The header field value",
																								MarkdownDescription: "The header field value",
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
																					Description:         "Path to access on the HTTP server.",
																					MarkdownDescription: "Path to access on the HTTP server.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"scheme": schema.StringAttribute{
																					Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																					MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																			Description:         "Sleep represents the duration that the container should sleep before being terminated.",
																			MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
																			Attributes: map[string]schema.Attribute{
																				"seconds": schema.Int64Attribute{
																					Description:         "Seconds is the number of seconds to sleep.",
																					MarkdownDescription: "Seconds is the number of seconds to sleep.",
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
																			Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																					MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
																	MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "Exec specifies the action to take.",
																			MarkdownDescription: "Exec specifies the action to take.",
																			Attributes: map[string]schema.Attribute{
																				"command": schema.ListAttribute{
																					Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																					MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																			Description:         "HTTPGet specifies the http request to perform.",
																			MarkdownDescription: "HTTPGet specifies the http request to perform.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																					MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "The header field value",
																								MarkdownDescription: "The header field value",
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
																					Description:         "Path to access on the HTTP server.",
																					MarkdownDescription: "Path to access on the HTTP server.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"scheme": schema.StringAttribute{
																					Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																					MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																			Description:         "Sleep represents the duration that the container should sleep before being terminated.",
																			MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
																			Attributes: map[string]schema.Attribute{
																				"seconds": schema.Int64Attribute{
																					Description:         "Seconds is the number of seconds to sleep.",
																					MarkdownDescription: "Seconds is the number of seconds to sleep.",
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
																			Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																					MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
															Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																	Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "GRPC specifies an action involving a GRPC port.",
																	MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
																	Attributes: map[string]schema.Attribute{
																		"port": schema.Int64Attribute{
																			Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"service": schema.StringAttribute{
																			Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
																			MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
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
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																	Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"period_seconds": schema.Int64Attribute{
																	Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"success_threshold": schema.Int64Attribute{
																	Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "TCPSocket specifies an action involving a TCP port.",
																	MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
															Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
															MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"ports": schema.ListNestedAttribute{
															Description:         "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
															MarkdownDescription: "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"container_port": schema.Int64Attribute{
																		Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
																		MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"host_ip": schema.StringAttribute{
																		Description:         "What host IP to bind the external port to.",
																		MarkdownDescription: "What host IP to bind the external port to.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"host_port": schema.Int64Attribute{
																		Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
																		MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
																		MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"protocol": schema.StringAttribute{
																		Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
																		MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
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
															Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																	Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "GRPC specifies an action involving a GRPC port.",
																	MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
																	Attributes: map[string]schema.Attribute{
																		"port": schema.Int64Attribute{
																			Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"service": schema.StringAttribute{
																			Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
																			MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
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
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																	Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"period_seconds": schema.Int64Attribute{
																	Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"success_threshold": schema.Int64Attribute{
																	Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "TCPSocket specifies an action involving a TCP port.",
																	MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
															Description:         "Resources resize policy for the container.",
															MarkdownDescription: "Resources resize policy for the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"resource_name": schema.StringAttribute{
																		Description:         "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
																		MarkdownDescription: "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"restart_policy": schema.StringAttribute{
																		Description:         "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
																		MarkdownDescription: "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
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
															Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

														"restart_policy": schema.StringAttribute{
															Description:         "RestartPolicy defines the restart behavior of individual containers in a pod. This field may only be set for init containers, and the only allowed value is 'Always'. For non-init containers or when this field is not specified, the restart behavior is defined by the Pod's restart policy and the container type. Setting the RestartPolicy as 'Always' for the init container will have the following effect: this init container will be continually restarted on exit until all regular containers have terminated. Once all regular containers have completed, all init containers with restartPolicy 'Always' will be shut down. This lifecycle differs from normal init containers and is often referred to as a 'sidecar' container. Although this init container still starts in the init container sequence, it does not wait for the container to complete before proceeding to the next init container. Instead, the next init container starts immediately after this init container is started, or after any startupProbe has successfully completed.",
															MarkdownDescription: "RestartPolicy defines the restart behavior of individual containers in a pod. This field may only be set for init containers, and the only allowed value is 'Always'. For non-init containers or when this field is not specified, the restart behavior is defined by the Pod's restart policy and the container type. Setting the RestartPolicy as 'Always' for the init container will have the following effect: this init container will be continually restarted on exit until all regular containers have terminated. Once all regular containers have completed, all init containers with restartPolicy 'Always' will be shut down. This lifecycle differs from normal init containers and is often referred to as a 'sidecar' container. Although this init container still starts in the init container sequence, it does not wait for the container to complete before proceeding to the next init container. Instead, the next init container starts immediately after this init container is started, or after any startupProbe has successfully completed.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_context": schema.SingleNestedAttribute{
															Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
															MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
															Attributes: map[string]schema.Attribute{
																"allow_privilege_escalation": schema.BoolAttribute{
																	Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
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
																	Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"proc_mount": schema.StringAttribute{
																	Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only_root_filesystem": schema.BoolAttribute{
																	Description:         "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_group": schema.Int64Attribute{
																	Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_non_root": schema.BoolAttribute{
																	Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																	MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_user": schema.Int64Attribute{
																	Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"se_linux_options": schema.SingleNestedAttribute{
																	Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
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
																	Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
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

																"windows_options": schema.SingleNestedAttribute{
																	Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
																	MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
																	Attributes: map[string]schema.Attribute{
																		"gmsa_credential_spec": schema.StringAttribute{
																			Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																			MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"gmsa_credential_spec_name": schema.StringAttribute{
																			Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																			MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"host_process": schema.BoolAttribute{
																			Description:         "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
																			MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"run_as_user_name": schema.StringAttribute{
																			Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																			MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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
															Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																	Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "GRPC specifies an action involving a GRPC port.",
																	MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
																	Attributes: map[string]schema.Attribute{
																		"port": schema.Int64Attribute{
																			Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"service": schema.StringAttribute{
																			Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
																			MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
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
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																	Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"period_seconds": schema.Int64Attribute{
																	Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"success_threshold": schema.Int64Attribute{
																	Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "TCPSocket specifies an action involving a TCP port.",
																	MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
															Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
															MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"stdin_once": schema.BoolAttribute{
															Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
															MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"termination_message_path": schema.StringAttribute{
															Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
															MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"termination_message_policy": schema.StringAttribute{
															Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
															MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tty": schema.BoolAttribute{
															Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
															MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"volume_devices": schema.ListNestedAttribute{
															Description:         "volumeDevices is the list of block devices to be used by the container.",
															MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"device_path": schema.StringAttribute{
																		Description:         "devicePath is the path inside of the container that the device will be mapped to.",
																		MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "name must match the name of a persistentVolumeClaim in the pod",
																		MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",
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
															Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
															MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",
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
																		Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
																		MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
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

														"working_dir": schema.StringAttribute{
															Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
															MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"dns_policy": schema.StringAttribute{
												Description:         "DNSPolicy",
												MarkdownDescription: "DNSPolicy",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ephemeral_containers": schema.ListNestedAttribute{
												Description:         "EphemeralContainers",
												MarkdownDescription: "EphemeralContainers",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"args": schema.ListAttribute{
															Description:         "Arguments to the entrypoint. The image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															MarkdownDescription: "Arguments to the entrypoint. The image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"command": schema.ListAttribute{
															Description:         "Entrypoint array. Not executed within a shell. The image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															MarkdownDescription: "Entrypoint array. Not executed within a shell. The image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"env": schema.ListNestedAttribute{
															Description:         "List of environment variables to set in the container. Cannot be updated.",
															MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",
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
																		Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																		MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "Specify whether the ConfigMap or its key must be defined",
																						MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
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
																				Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				Attributes: map[string]schema.Attribute{
																					"api_version": schema.StringAttribute{
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"field_path": schema.StringAttribute{
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",
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
																				Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				Attributes: map[string]schema.Attribute{
																					"container_name": schema.StringAttribute{
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"divisor": schema.StringAttribute{
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"resource": schema.StringAttribute{
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",
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
																				Description:         "Selects a key of a secret in the pod's namespace",
																				MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key of the secret to select from. Must be a valid secret key.",
																						MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

														"image": schema.StringAttribute{
															Description:         "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images",
															MarkdownDescription: "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_policy": schema.StringAttribute{
															Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
															MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"lifecycle": schema.SingleNestedAttribute{
															Description:         "Lifecycle is not allowed for ephemeral containers.",
															MarkdownDescription: "Lifecycle is not allowed for ephemeral containers.",
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
																					Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																					MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																			Description:         "HTTPGet specifies the http request to perform.",
																			MarkdownDescription: "HTTPGet specifies the http request to perform.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																					MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "The header field value",
																								MarkdownDescription: "The header field value",
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
																					Description:         "Path to access on the HTTP server.",
																					MarkdownDescription: "Path to access on the HTTP server.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"scheme": schema.StringAttribute{
																					Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																					MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																			Description:         "Sleep represents the duration that the container should sleep before being terminated.",
																			MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
																			Attributes: map[string]schema.Attribute{
																				"seconds": schema.Int64Attribute{
																					Description:         "Seconds is the number of seconds to sleep.",
																					MarkdownDescription: "Seconds is the number of seconds to sleep.",
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
																			Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																					MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
																	MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "Exec specifies the action to take.",
																			MarkdownDescription: "Exec specifies the action to take.",
																			Attributes: map[string]schema.Attribute{
																				"command": schema.ListAttribute{
																					Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																					MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																			Description:         "HTTPGet specifies the http request to perform.",
																			MarkdownDescription: "HTTPGet specifies the http request to perform.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																					MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "The header field value",
																								MarkdownDescription: "The header field value",
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
																					Description:         "Path to access on the HTTP server.",
																					MarkdownDescription: "Path to access on the HTTP server.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"scheme": schema.StringAttribute{
																					Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																					MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																			Description:         "Sleep represents the duration that the container should sleep before being terminated.",
																			MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
																			Attributes: map[string]schema.Attribute{
																				"seconds": schema.Int64Attribute{
																					Description:         "Seconds is the number of seconds to sleep.",
																					MarkdownDescription: "Seconds is the number of seconds to sleep.",
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
																			Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																					MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
															Description:         "Probes are not allowed for ephemeral containers.",
															MarkdownDescription: "Probes are not allowed for ephemeral containers.",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																	Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "GRPC specifies an action involving a GRPC port.",
																	MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
																	Attributes: map[string]schema.Attribute{
																		"port": schema.Int64Attribute{
																			Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"service": schema.StringAttribute{
																			Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
																			MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
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
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																	Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"period_seconds": schema.Int64Attribute{
																	Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"success_threshold": schema.Int64Attribute{
																	Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "TCPSocket specifies an action involving a TCP port.",
																	MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
															Description:         "Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.",
															MarkdownDescription: "Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"ports": schema.ListNestedAttribute{
															Description:         "Ports are not allowed for ephemeral containers.",
															MarkdownDescription: "Ports are not allowed for ephemeral containers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"container_port": schema.Int64Attribute{
																		Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
																		MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"host_ip": schema.StringAttribute{
																		Description:         "What host IP to bind the external port to.",
																		MarkdownDescription: "What host IP to bind the external port to.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"host_port": schema.Int64Attribute{
																		Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
																		MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
																		MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"protocol": schema.StringAttribute{
																		Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
																		MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
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
															Description:         "Probes are not allowed for ephemeral containers.",
															MarkdownDescription: "Probes are not allowed for ephemeral containers.",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																	Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "GRPC specifies an action involving a GRPC port.",
																	MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
																	Attributes: map[string]schema.Attribute{
																		"port": schema.Int64Attribute{
																			Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"service": schema.StringAttribute{
																			Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
																			MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
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
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																	Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"period_seconds": schema.Int64Attribute{
																	Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"success_threshold": schema.Int64Attribute{
																	Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "TCPSocket specifies an action involving a TCP port.",
																	MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
															Description:         "Resources resize policy for the container.",
															MarkdownDescription: "Resources resize policy for the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"resource_name": schema.StringAttribute{
																		Description:         "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
																		MarkdownDescription: "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"restart_policy": schema.StringAttribute{
																		Description:         "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
																		MarkdownDescription: "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
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
															Description:         "Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
															MarkdownDescription: "Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
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

														"restart_policy": schema.StringAttribute{
															Description:         "Restart policy for the container to manage the restart behavior of each container within a pod. This may only be set for init containers. You cannot set this field on ephemeral containers.",
															MarkdownDescription: "Restart policy for the container to manage the restart behavior of each container within a pod. This may only be set for init containers. You cannot set this field on ephemeral containers.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_context": schema.SingleNestedAttribute{
															Description:         "Optional: SecurityContext defines the security options the ephemeral container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",
															MarkdownDescription: "Optional: SecurityContext defines the security options the ephemeral container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",
															Attributes: map[string]schema.Attribute{
																"allow_privilege_escalation": schema.BoolAttribute{
																	Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
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
																	Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"proc_mount": schema.StringAttribute{
																	Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only_root_filesystem": schema.BoolAttribute{
																	Description:         "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_group": schema.Int64Attribute{
																	Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_non_root": schema.BoolAttribute{
																	Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																	MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_user": schema.Int64Attribute{
																	Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"se_linux_options": schema.SingleNestedAttribute{
																	Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
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
																	Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
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

																"windows_options": schema.SingleNestedAttribute{
																	Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
																	MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
																	Attributes: map[string]schema.Attribute{
																		"gmsa_credential_spec": schema.StringAttribute{
																			Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																			MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"gmsa_credential_spec_name": schema.StringAttribute{
																			Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																			MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"host_process": schema.BoolAttribute{
																			Description:         "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
																			MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"run_as_user_name": schema.StringAttribute{
																			Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																			MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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
															Description:         "Probes are not allowed for ephemeral containers.",
															MarkdownDescription: "Probes are not allowed for ephemeral containers.",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																	Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "GRPC specifies an action involving a GRPC port.",
																	MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
																	Attributes: map[string]schema.Attribute{
																		"port": schema.Int64Attribute{
																			Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"service": schema.StringAttribute{
																			Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
																			MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
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
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																	Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"period_seconds": schema.Int64Attribute{
																	Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"success_threshold": schema.Int64Attribute{
																	Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "TCPSocket specifies an action involving a TCP port.",
																	MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
															Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
															MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"stdin_once": schema.BoolAttribute{
															Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
															MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target_container_name": schema.StringAttribute{
															Description:         "If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container uses the namespaces configured in the Pod spec. The container runtime must implement support for this feature. If the runtime does not support namespace targeting then the result of setting this field is undefined.",
															MarkdownDescription: "If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container uses the namespaces configured in the Pod spec. The container runtime must implement support for this feature. If the runtime does not support namespace targeting then the result of setting this field is undefined.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"termination_message_path": schema.StringAttribute{
															Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
															MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"termination_message_policy": schema.StringAttribute{
															Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
															MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tty": schema.BoolAttribute{
															Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
															MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"volume_devices": schema.ListNestedAttribute{
															Description:         "volumeDevices is the list of block devices to be used by the container.",
															MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"device_path": schema.StringAttribute{
																		Description:         "devicePath is the path inside of the container that the device will be mapped to.",
																		MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "name must match the name of a persistentVolumeClaim in the pod",
																		MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",
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
															Description:         "Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. Cannot be updated.",
															MarkdownDescription: "Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. Cannot be updated.",
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
																		Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
																		MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
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

														"working_dir": schema.StringAttribute{
															Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
															MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
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
												Description:         "InitContainers",
												MarkdownDescription: "InitContainers",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"args": schema.ListAttribute{
															Description:         "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															MarkdownDescription: "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"command": schema.ListAttribute{
															Description:         "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															MarkdownDescription: "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"env": schema.ListNestedAttribute{
															Description:         "List of environment variables to set in the container. Cannot be updated.",
															MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",
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
																		Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																		MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "Specify whether the ConfigMap or its key must be defined",
																						MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
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
																				Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																				Attributes: map[string]schema.Attribute{
																					"api_version": schema.StringAttribute{
																						Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"field_path": schema.StringAttribute{
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",
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
																				Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				Attributes: map[string]schema.Attribute{
																					"container_name": schema.StringAttribute{
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"divisor": schema.StringAttribute{
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"resource": schema.StringAttribute{
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",
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
																				Description:         "Selects a key of a secret in the pod's namespace",
																				MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "The key of the secret to select from. Must be a valid secret key.",
																						MarkdownDescription: "The key of the secret to select from. Must be a valid secret key.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"name": schema.StringAttribute{
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

														"image": schema.StringAttribute{
															Description:         "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
															MarkdownDescription: "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image_pull_policy": schema.StringAttribute{
															Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
															MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
															Required:            false,
															Optional:            true,
															Computed:            false,
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
																					Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																					MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																			Description:         "HTTPGet specifies the http request to perform.",
																			MarkdownDescription: "HTTPGet specifies the http request to perform.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																					MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "The header field value",
																								MarkdownDescription: "The header field value",
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
																					Description:         "Path to access on the HTTP server.",
																					MarkdownDescription: "Path to access on the HTTP server.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"scheme": schema.StringAttribute{
																					Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																					MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																			Description:         "Sleep represents the duration that the container should sleep before being terminated.",
																			MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
																			Attributes: map[string]schema.Attribute{
																				"seconds": schema.Int64Attribute{
																					Description:         "Seconds is the number of seconds to sleep.",
																					MarkdownDescription: "Seconds is the number of seconds to sleep.",
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
																			Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																					MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
																	MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
																	Attributes: map[string]schema.Attribute{
																		"exec": schema.SingleNestedAttribute{
																			Description:         "Exec specifies the action to take.",
																			MarkdownDescription: "Exec specifies the action to take.",
																			Attributes: map[string]schema.Attribute{
																				"command": schema.ListAttribute{
																					Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																					MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																			Description:         "HTTPGet specifies the http request to perform.",
																			MarkdownDescription: "HTTPGet specifies the http request to perform.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"http_headers": schema.ListNestedAttribute{
																					Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																					MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"value": schema.StringAttribute{
																								Description:         "The header field value",
																								MarkdownDescription: "The header field value",
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
																					Description:         "Path to access on the HTTP server.",
																					MarkdownDescription: "Path to access on the HTTP server.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"scheme": schema.StringAttribute{
																					Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																					MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																			Description:         "Sleep represents the duration that the container should sleep before being terminated.",
																			MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
																			Attributes: map[string]schema.Attribute{
																				"seconds": schema.Int64Attribute{
																					Description:         "Seconds is the number of seconds to sleep.",
																					MarkdownDescription: "Seconds is the number of seconds to sleep.",
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
																			Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
																			Attributes: map[string]schema.Attribute{
																				"host": schema.StringAttribute{
																					Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																					MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"port": schema.StringAttribute{
																					Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																					MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
															Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																	Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "GRPC specifies an action involving a GRPC port.",
																	MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
																	Attributes: map[string]schema.Attribute{
																		"port": schema.Int64Attribute{
																			Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"service": schema.StringAttribute{
																			Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
																			MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
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
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																	Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"period_seconds": schema.Int64Attribute{
																	Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"success_threshold": schema.Int64Attribute{
																	Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "TCPSocket specifies an action involving a TCP port.",
																	MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
															Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
															MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"ports": schema.ListNestedAttribute{
															Description:         "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
															MarkdownDescription: "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"container_port": schema.Int64Attribute{
																		Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
																		MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"host_ip": schema.StringAttribute{
																		Description:         "What host IP to bind the external port to.",
																		MarkdownDescription: "What host IP to bind the external port to.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"host_port": schema.Int64Attribute{
																		Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
																		MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
																		MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"protocol": schema.StringAttribute{
																		Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
																		MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
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
															Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																	Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "GRPC specifies an action involving a GRPC port.",
																	MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
																	Attributes: map[string]schema.Attribute{
																		"port": schema.Int64Attribute{
																			Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"service": schema.StringAttribute{
																			Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
																			MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
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
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																	Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"period_seconds": schema.Int64Attribute{
																	Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"success_threshold": schema.Int64Attribute{
																	Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "TCPSocket specifies an action involving a TCP port.",
																	MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
															Description:         "Resources resize policy for the container.",
															MarkdownDescription: "Resources resize policy for the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"resource_name": schema.StringAttribute{
																		Description:         "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
																		MarkdownDescription: "Name of the resource to which this resource resize policy applies. Supported values: cpu, memory.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"restart_policy": schema.StringAttribute{
																		Description:         "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
																		MarkdownDescription: "Restart policy to apply when specified resource is resized. If not specified, it defaults to NotRequired.",
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
															Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

														"restart_policy": schema.StringAttribute{
															Description:         "RestartPolicy defines the restart behavior of individual containers in a pod. This field may only be set for init containers, and the only allowed value is 'Always'. For non-init containers or when this field is not specified, the restart behavior is defined by the Pod's restart policy and the container type. Setting the RestartPolicy as 'Always' for the init container will have the following effect: this init container will be continually restarted on exit until all regular containers have terminated. Once all regular containers have completed, all init containers with restartPolicy 'Always' will be shut down. This lifecycle differs from normal init containers and is often referred to as a 'sidecar' container. Although this init container still starts in the init container sequence, it does not wait for the container to complete before proceeding to the next init container. Instead, the next init container starts immediately after this init container is started, or after any startupProbe has successfully completed.",
															MarkdownDescription: "RestartPolicy defines the restart behavior of individual containers in a pod. This field may only be set for init containers, and the only allowed value is 'Always'. For non-init containers or when this field is not specified, the restart behavior is defined by the Pod's restart policy and the container type. Setting the RestartPolicy as 'Always' for the init container will have the following effect: this init container will be continually restarted on exit until all regular containers have terminated. Once all regular containers have completed, all init containers with restartPolicy 'Always' will be shut down. This lifecycle differs from normal init containers and is often referred to as a 'sidecar' container. Although this init container still starts in the init container sequence, it does not wait for the container to complete before proceeding to the next init container. Instead, the next init container starts immediately after this init container is started, or after any startupProbe has successfully completed.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"security_context": schema.SingleNestedAttribute{
															Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
															MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
															Attributes: map[string]schema.Attribute{
																"allow_privilege_escalation": schema.BoolAttribute{
																	Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
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
																	Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"proc_mount": schema.StringAttribute{
																	Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only_root_filesystem": schema.BoolAttribute{
																	Description:         "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_group": schema.Int64Attribute{
																	Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_non_root": schema.BoolAttribute{
																	Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																	MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"run_as_user": schema.Int64Attribute{
																	Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"se_linux_options": schema.SingleNestedAttribute{
																	Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is windows.",
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
																	Description:         "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
																	MarkdownDescription: "The seccomp options to use by this container. If seccomp options are provided at both the pod & container level, the container options override the pod options. Note that this field cannot be set when spec.os.name is windows.",
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

																"windows_options": schema.SingleNestedAttribute{
																	Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
																	MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
																	Attributes: map[string]schema.Attribute{
																		"gmsa_credential_spec": schema.StringAttribute{
																			Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																			MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"gmsa_credential_spec_name": schema.StringAttribute{
																			Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																			MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"host_process": schema.BoolAttribute{
																			Description:         "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
																			MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"run_as_user_name": schema.StringAttribute{
																			Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																			MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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
															Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "Exec specifies the action to take.",
																	MarkdownDescription: "Exec specifies the action to take.",
																	Attributes: map[string]schema.Attribute{
																		"command": schema.ListAttribute{
																			Description:         "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																			MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
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
																	Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"grpc": schema.SingleNestedAttribute{
																	Description:         "GRPC specifies an action involving a GRPC port.",
																	MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
																	Attributes: map[string]schema.Attribute{
																		"port": schema.Int64Attribute{
																			Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"service": schema.StringAttribute{
																			Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
																			MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md). If this is not specified, the default behavior is defined by gRPC.",
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
																	Description:         "HTTPGet specifies the http request to perform.",
																	MarkdownDescription: "HTTPGet specifies the http request to perform.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																			MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						MarkdownDescription: "The header field name. This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
																						Description:         "The header field value",
																						MarkdownDescription: "The header field value",
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
																			Description:         "Path to access on the HTTP server.",
																			MarkdownDescription: "Path to access on the HTTP server.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"scheme": schema.StringAttribute{
																			Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																			MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",
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
																	Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"period_seconds": schema.Int64Attribute{
																	Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"success_threshold": schema.Int64Attribute{
																	Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "TCPSocket specifies an action involving a TCP port.",
																	MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																			MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"port": schema.StringAttribute{
																			Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																			MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
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
																	Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"timeout_seconds": schema.Int64Attribute{
																	Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
																	MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
															Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
															MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"stdin_once": schema.BoolAttribute{
															Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
															MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"termination_message_path": schema.StringAttribute{
															Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
															MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"termination_message_policy": schema.StringAttribute{
															Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
															MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"tty": schema.BoolAttribute{
															Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
															MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"volume_devices": schema.ListNestedAttribute{
															Description:         "volumeDevices is the list of block devices to be used by the container.",
															MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"device_path": schema.StringAttribute{
																		Description:         "devicePath is the path inside of the container that the device will be mapped to.",
																		MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "name must match the name of a persistentVolumeClaim in the pod",
																		MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",
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
															Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
															MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",
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
																		Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
																		MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
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

														"working_dir": schema.StringAttribute{
															Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
															MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
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
												Description:         "NodeSelector",
												MarkdownDescription: "NodeSelector",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"restart_policy": schema.StringAttribute{
												Description:         "RestartPolicy",
												MarkdownDescription: "RestartPolicy",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"security_context": schema.SingleNestedAttribute{
												Description:         "PodSecurityContext",
												MarkdownDescription: "PodSecurityContext",
												Attributes: map[string]schema.Attribute{
													"fs_group": schema.Int64Attribute{
														Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod: 1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw---- If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"fs_group_change_policy": schema.StringAttribute{
														Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_group": schema.Int64Attribute{
														Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_non_root": schema.BoolAttribute{
														Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_user": schema.Int64Attribute{
														Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"se_linux_options": schema.SingleNestedAttribute{
														Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container. May also be set in SecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
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
														Description:         "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "The seccomp options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
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
														Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sysctls": schema.ListNestedAttribute{
														Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of a property to set",
																	MarkdownDescription: "Name of a property to set",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "Value of a property to set",
																	MarkdownDescription: "Value of a property to set",
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
														Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
														MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
														Attributes: map[string]schema.Attribute{
															"gmsa_credential_spec": schema.StringAttribute{
																Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"gmsa_credential_spec_name": schema.StringAttribute{
																Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"host_process": schema.BoolAttribute{
																Description:         "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
																MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container. All of a Pod's containers must have the same effective HostProcess value (it is not allowed to have a mix of HostProcess containers and non-HostProcess containers). In addition, if HostProcess is true then HostNetwork must also be set to true.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"run_as_user_name": schema.StringAttribute{
																Description:         "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
																MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process. Defaults to the user specified in image metadata if unspecified. May also be set in PodSecurityContext. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
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

											"termination_grace_period_seconds": schema.Int64Attribute{
												Description:         "TerminationGracePeriodSeconds",
												MarkdownDescription: "TerminationGracePeriodSeconds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"topology_spread_constraints": schema.ListNestedAttribute{
												Description:         "TopologySpreadConstraints",
												MarkdownDescription: "TopologySpreadConstraints",
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
															Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector. This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
															MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector. This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_skew": schema.Int64Attribute{
															Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | | P P | P P | P | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
															MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | | P P | P P | P | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"min_domains": schema.Int64Attribute{
															Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule. For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | | P P | P P | P P | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew. This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
															MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule. For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | | P P | P P | P P | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew. This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_affinity_policy": schema.StringAttribute{
															Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations. If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
															MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations. If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_taints_policy": schema.StringAttribute{
															Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included. If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
															MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included. If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
															MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"when_unsatisfiable": schema.StringAttribute{
															Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P | P | P | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
															MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P | P | P | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
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
												Description:         "Volumes",
												MarkdownDescription: "Volumes",
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
																	Optional:            true,
																	Computed:            false,
																},

																"partition": schema.Int64Attribute{
																	Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																	MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_id": schema.StringAttribute{
																	Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
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
															Description:         "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
															MarkdownDescription: "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
															Attributes: map[string]schema.Attribute{
																"caching_mode": schema.StringAttribute{
																	Description:         "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																	MarkdownDescription: "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"disk_name": schema.StringAttribute{
																	Description:         "diskName is the Name of the data disk in the blob storage",
																	MarkdownDescription: "diskName is the Name of the data disk in the blob storage",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"disk_uri": schema.StringAttribute{
																	Description:         "diskURI is the URI of data disk in the blob storage",
																	MarkdownDescription: "diskURI is the URI of data disk in the blob storage",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"fs_type": schema.StringAttribute{
																	Description:         "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "kind expected values are Shared: multiple blob disks per storage account Dedicated: single blob disk per storage account Managed: azure managed data disk (only in managed availability set). defaults to shared",
																	MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account Dedicated: single blob disk per storage account Managed: azure managed data disk (only in managed availability set). defaults to shared",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
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
															Description:         "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
															MarkdownDescription: "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
															Attributes: map[string]schema.Attribute{
																"read_only": schema.BoolAttribute{
																	Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_name": schema.StringAttribute{
																	Description:         "secretName is the name of secret that contains Azure Storage Account Name and Key",
																	MarkdownDescription: "secretName is the name of secret that contains Azure Storage Account Name and Key",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"share_name": schema.StringAttribute{
																	Description:         "shareName is the azure share Name",
																	MarkdownDescription: "shareName is the azure share Name",
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
															Description:         "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
															MarkdownDescription: "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
															Attributes: map[string]schema.Attribute{
																"monitors": schema.ListAttribute{
																	Description:         "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																	MarkdownDescription: "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_file": schema.StringAttribute{
																	Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Description:         "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
															Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																	MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Description:         "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	MarkdownDescription: "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
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
															Description:         "configMap represents a configMap that should populate this volume",
															MarkdownDescription: "configMap represents a configMap that should populate this volume",
															Attributes: map[string]schema.Attribute{
																"default_mode": schema.Int64Attribute{
																	Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"items": schema.ListNestedAttribute{
																	Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																	MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the key to project.",
																				MarkdownDescription: "key is the key to project.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"mode": schema.Int64Attribute{
																				Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.StringAttribute{
																				Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"optional": schema.BoolAttribute{
																	Description:         "optional specify whether the ConfigMap or its keys must be defined",
																	MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
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
															Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
															MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
															Attributes: map[string]schema.Attribute{
																"driver": schema.StringAttribute{
																	Description:         "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																	MarkdownDescription: "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"fs_type": schema.StringAttribute{
																	Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																	MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"node_publish_secret_ref": schema.SingleNestedAttribute{
																	Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																	MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Description:         "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																	MarkdownDescription: "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_attributes": schema.MapAttribute{
																	Description:         "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
																	MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
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
															Description:         "downwardAPI represents downward API about the pod that should populate this volume",
															MarkdownDescription: "downwardAPI represents downward API about the pod that should populate this volume",
															Attributes: map[string]schema.Attribute{
																"default_mode": schema.Int64Attribute{
																	Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
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
																						Optional:            true,
																						Computed:            false,
																					},

																					"field_path": schema.StringAttribute{
																						Description:         "Path of the field to select in the specified API version.",
																						MarkdownDescription: "Path of the field to select in the specified API version.",
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
																				Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.StringAttribute{
																				Description:         "Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																				MarkdownDescription: "Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"resource_field_ref": schema.SingleNestedAttribute{
																				Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																				Attributes: map[string]schema.Attribute{
																					"container_name": schema.StringAttribute{
																						Description:         "Container name: required for volumes, optional for env vars",
																						MarkdownDescription: "Container name: required for volumes, optional for env vars",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"divisor": schema.StringAttribute{
																						Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																						MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"resource": schema.StringAttribute{
																						Description:         "Required: resource to select",
																						MarkdownDescription: "Required: resource to select",
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
															Description:         "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
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

														"ephemeral": schema.SingleNestedAttribute{
															Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed. Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim). Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod. Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information. A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
															MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed. Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim). Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod. Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information. A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
															Attributes: map[string]schema.Attribute{
																"volume_claim_template": schema.SingleNestedAttribute{
																	Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod. The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long). An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster. This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created. Required, must not be nil.",
																	MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod. The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long). An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster. This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created. Required, must not be nil.",
																	Attributes: map[string]schema.Attribute{
																		"metadata": schema.MapAttribute{
																			Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																			MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
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
																					Optional:            true,
																					Computed:            false,
																				},

																				"data_source": schema.SingleNestedAttribute{
																					Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																					MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
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
																					Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																					MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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

																						"namespace": schema.StringAttribute{
																							Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																							MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
																					Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																					MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
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
																					Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																					MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_attributes_class_name": schema.StringAttribute{
																					Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim. If specified, the CSI driver will create or update the volume with the attributes defined in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName, it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass will be applied to the claim but it's not allowed to reset this field to empty string once it is set. If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass will be set by the persistentvolume controller if it exists. If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource exists. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass (Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
																					MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim. If specified, the CSI driver will create or update the volume with the attributes defined in the corresponding VolumeAttributesClass. This has a different purpose than storageClassName, it can be changed after the claim is created. An empty string value means that no VolumeAttributesClass will be applied to the claim but it's not allowed to reset this field to empty string once it is set. If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClass will be set by the persistentvolume controller if it exists. If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will be set to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resource exists. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#volumeattributesclass (Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_mode": schema.StringAttribute{
																					Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																					MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_name": schema.StringAttribute{
																					Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
																					MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
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
															Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
															MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"lun": schema.Int64Attribute{
																	Description:         "lun is Optional: FC target lun number",
																	MarkdownDescription: "lun is Optional: FC target lun number",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"target_ww_ns": schema.ListAttribute{
																	Description:         "targetWWNs is Optional: FC target worldwide names (WWNs)",
																	MarkdownDescription: "targetWWNs is Optional: FC target worldwide names (WWNs)",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"wwids": schema.ListAttribute{
																	Description:         "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
																	MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
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
															Description:         "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
															MarkdownDescription: "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
															Attributes: map[string]schema.Attribute{
																"driver": schema.StringAttribute{
																	Description:         "driver is the name of the driver to use for this volume.",
																	MarkdownDescription: "driver is the name of the driver to use for this volume.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"options": schema.MapAttribute{
																	Description:         "options is Optional: this field holds extra command options if any.",
																	MarkdownDescription: "options is Optional: this field holds extra command options if any.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																	MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
															Description:         "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
															MarkdownDescription: "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
															Attributes: map[string]schema.Attribute{
																"dataset_name": schema.StringAttribute{
																	Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																	MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"dataset_uuid": schema.StringAttribute{
																	Description:         "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																	MarkdownDescription: "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
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
															Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																	MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"partition": schema.Int64Attribute{
																	Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"pd_name": schema.StringAttribute{
																	Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
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
															Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
															MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
															Attributes: map[string]schema.Attribute{
																"directory": schema.StringAttribute{
																	Description:         "directory is the target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																	MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'. If '.' is supplied, the volume directory will be the git repository. Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"repository": schema.StringAttribute{
																	Description:         "repository is the URL",
																	MarkdownDescription: "repository is the URL",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"revision": schema.StringAttribute{
																	Description:         "revision is the commit hash for the specified revision.",
																	MarkdownDescription: "revision is the commit hash for the specified revision.",
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
															Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
															MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
															Attributes: map[string]schema.Attribute{
																"endpoints": schema.StringAttribute{
																	Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	MarkdownDescription: "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
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
															Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
															MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
															Attributes: map[string]schema.Attribute{
																"path": schema.StringAttribute{
																	Description:         "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																	MarkdownDescription: "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																	MarkdownDescription: "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
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
															Description:         "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
															MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
															Attributes: map[string]schema.Attribute{
																"chap_auth_discovery": schema.BoolAttribute{
																	Description:         "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																	MarkdownDescription: "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"chap_auth_session": schema.BoolAttribute{
																	Description:         "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																	MarkdownDescription: "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																	MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"initiator_name": schema.StringAttribute{
																	Description:         "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																	MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"iqn": schema.StringAttribute{
																	Description:         "iqn is the target iSCSI Qualified Name.",
																	MarkdownDescription: "iqn is the target iSCSI Qualified Name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"iscsi_interface": schema.StringAttribute{
																	Description:         "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																	MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"lun": schema.Int64Attribute{
																	Description:         "lun represents iSCSI Target Lun number.",
																	MarkdownDescription: "lun represents iSCSI Target Lun number.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"portals": schema.ListAttribute{
																	Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																	MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																	MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																	MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																	MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
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
															Description:         "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"nfs": schema.SingleNestedAttribute{
															Description:         "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															Attributes: map[string]schema.Attribute{
																"path": schema.StringAttribute{
																	Description:         "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	MarkdownDescription: "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"server": schema.StringAttribute{
																	Description:         "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	MarkdownDescription: "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
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
															Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
															MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
															Attributes: map[string]schema.Attribute{
																"claim_name": schema.StringAttribute{
																	Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																	MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
																	MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
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
															Description:         "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
															MarkdownDescription: "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"pd_id": schema.StringAttribute{
																	Description:         "pdID is the ID that identifies Photon Controller persistent disk",
																	MarkdownDescription: "pdID is the ID that identifies Photon Controller persistent disk",
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
															Description:         "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
															MarkdownDescription: "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_id": schema.StringAttribute{
																	Description:         "volumeID uniquely identifies a Portworx volume",
																	MarkdownDescription: "volumeID uniquely identifies a Portworx volume",
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
															Description:         "projected items for all in one resources secrets, configmaps, and downward API",
															MarkdownDescription: "projected items for all in one resources secrets, configmaps, and downward API",
															Attributes: map[string]schema.Attribute{
																"default_mode": schema.Int64Attribute{
																	Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"sources": schema.ListNestedAttribute{
																	Description:         "sources is the list of volume projections",
																	MarkdownDescription: "sources is the list of volume projections",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"cluster_trust_bundle": schema.SingleNestedAttribute{
																				Description:         "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' field of ClusterTrustBundle objects in an auto-updating file. Alpha, gated by the ClusterTrustBundleProjection feature gate. ClusterTrustBundle objects can either be selected by name, or by the combination of signer name and a label selector. Kubelet performs aggressive normalization of the PEM contents written into the pod filesystem. Esoteric PEM features such as inter-block comments and block headers are stripped. Certificates are deduplicated. The ordering of certificates within the file is arbitrary, and Kubelet may change the order over time.",
																				MarkdownDescription: "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' field of ClusterTrustBundle objects in an auto-updating file. Alpha, gated by the ClusterTrustBundleProjection feature gate. ClusterTrustBundle objects can either be selected by name, or by the combination of signer name and a label selector. Kubelet performs aggressive normalization of the PEM contents written into the pod filesystem. Esoteric PEM features such as inter-block comments and block headers are stripped. Certificates are deduplicated. The ordering of certificates within the file is arbitrary, and Kubelet may change the order over time.",
																				Attributes: map[string]schema.Attribute{
																					"label_selector": schema.SingleNestedAttribute{
																						Description:         "Select all ClusterTrustBundles that match this label selector. Only has effect if signerName is set. Mutually-exclusive with name. If unset, interpreted as 'match nothing'. If set but empty, interpreted as 'match everything'.",
																						MarkdownDescription: "Select all ClusterTrustBundles that match this label selector. Only has effect if signerName is set. Mutually-exclusive with name. If unset, interpreted as 'match nothing'. If set but empty, interpreted as 'match everything'.",
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

																					"name": schema.StringAttribute{
																						Description:         "Select a single ClusterTrustBundle by object name. Mutually-exclusive with signerName and labelSelector.",
																						MarkdownDescription: "Select a single ClusterTrustBundle by object name. Mutually-exclusive with signerName and labelSelector.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "If true, don't block pod startup if the referenced ClusterTrustBundle(s) aren't available. If using name, then the named ClusterTrustBundle is allowed not to exist. If using signerName, then the combination of signerName and labelSelector is allowed to match zero ClusterTrustBundles.",
																						MarkdownDescription: "If true, don't block pod startup if the referenced ClusterTrustBundle(s) aren't available. If using name, then the named ClusterTrustBundle is allowed not to exist. If using signerName, then the combination of signerName and labelSelector is allowed to match zero ClusterTrustBundles.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "Relative path from the volume root to write the bundle.",
																						MarkdownDescription: "Relative path from the volume root to write the bundle.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"signer_name": schema.StringAttribute{
																						Description:         "Select all ClusterTrustBundles that match this signer name. Mutually-exclusive with name. The contents of all selected ClusterTrustBundles will be unified and deduplicated.",
																						MarkdownDescription: "Select all ClusterTrustBundles that match this signer name. Mutually-exclusive with name. The contents of all selected ClusterTrustBundles will be unified and deduplicated.",
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
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"mode": schema.Int64Attribute{
																									Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"path": schema.StringAttribute{
																									Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																									MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "optional specify whether the ConfigMap or its keys must be defined",
																						MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
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
																											Optional:            true,
																											Computed:            false,
																										},

																										"field_path": schema.StringAttribute{
																											Description:         "Path of the field to select in the specified API version.",
																											MarkdownDescription: "Path of the field to select in the specified API version.",
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
																									Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"path": schema.StringAttribute{
																									Description:         "Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																									MarkdownDescription: "Required: Path is the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"resource_field_ref": schema.SingleNestedAttribute{
																									Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																									MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																									Attributes: map[string]schema.Attribute{
																										"container_name": schema.StringAttribute{
																											Description:         "Container name: required for volumes, optional for env vars",
																											MarkdownDescription: "Container name: required for volumes, optional for env vars",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"divisor": schema.StringAttribute{
																											Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																											MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"resource": schema.StringAttribute{
																											Description:         "Required: resource to select",
																											MarkdownDescription: "Required: resource to select",
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
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"mode": schema.Int64Attribute{
																									Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"path": schema.StringAttribute{
																									Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																									MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"optional": schema.BoolAttribute{
																						Description:         "optional field specify whether the Secret or its key must be defined",
																						MarkdownDescription: "optional field specify whether the Secret or its key must be defined",
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
																				Description:         "serviceAccountToken is information about the serviceAccountToken data to project",
																				MarkdownDescription: "serviceAccountToken is information about the serviceAccountToken data to project",
																				Attributes: map[string]schema.Attribute{
																					"audience": schema.StringAttribute{
																						Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																						MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"expiration_seconds": schema.Int64Attribute{
																						Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																						MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "path is the path relative to the mount point of the file to project the token into.",
																						MarkdownDescription: "path is the path relative to the mount point of the file to project the token into.",
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
															Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
															MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
															Attributes: map[string]schema.Attribute{
																"group": schema.StringAttribute{
																	Description:         "group to map volume access to Default is no group",
																	MarkdownDescription: "group to map volume access to Default is no group",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																	MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"registry": schema.StringAttribute{
																	Description:         "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																	MarkdownDescription: "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"tenant": schema.StringAttribute{
																	Description:         "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																	MarkdownDescription: "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"user": schema.StringAttribute{
																	Description:         "user to map volume access to Defaults to serivceaccount user",
																	MarkdownDescription: "user to map volume access to Defaults to serivceaccount user",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume": schema.StringAttribute{
																	Description:         "volume is a string that references an already created Quobyte volume by name.",
																	MarkdownDescription: "volume is a string that references an already created Quobyte volume by name.",
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
															Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
															MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																	MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"keyring": schema.StringAttribute{
																	Description:         "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"monitors": schema.ListAttribute{
																	Description:         "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"pool": schema.StringAttribute{
																	Description:         "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Description:         "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
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
															Description:         "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
															MarkdownDescription: "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"gateway": schema.StringAttribute{
																	Description:         "gateway is the host address of the ScaleIO API Gateway.",
																	MarkdownDescription: "gateway is the host address of the ScaleIO API Gateway.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"protection_domain": schema.StringAttribute{
																	Description:         "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																	MarkdownDescription: "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																	MarkdownDescription: "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Description:         "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																	MarkdownDescription: "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"storage_mode": schema.StringAttribute{
																	Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																	MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"storage_pool": schema.StringAttribute{
																	Description:         "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																	MarkdownDescription: "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"system": schema.StringAttribute{
																	Description:         "system is the name of the storage system as configured in ScaleIO.",
																	MarkdownDescription: "system is the name of the storage system as configured in ScaleIO.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"volume_name": schema.StringAttribute{
																	Description:         "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
																	MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
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
															Description:         "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
															MarkdownDescription: "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
															Attributes: map[string]schema.Attribute{
																"default_mode": schema.Int64Attribute{
																	Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"items": schema.ListNestedAttribute{
																	Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																	MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the key to project.",
																				MarkdownDescription: "key is the key to project.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"mode": schema.Int64Attribute{
																				Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.StringAttribute{
																				Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																	Description:         "optional field specify whether the Secret or its keys must be defined",
																	MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_name": schema.StringAttribute{
																	Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																	MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
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
															Description:         "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
															MarkdownDescription: "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef specifies the secret to use for obtaining the StorageOS API credentials. If not specified, default values will be attempted.",
																	MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS API credentials. If not specified, default values will be attempted.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Description:         "volumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.",
																	MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume. Volume names are only unique within a namespace.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_namespace": schema.StringAttribute{
																	Description:         "volumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																	MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS. If no namespace is specified then the Pod's namespace will be used. This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
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
															Description:         "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
															MarkdownDescription: "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"storage_policy_id": schema.StringAttribute{
																	Description:         "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																	MarkdownDescription: "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"storage_policy_name": schema.StringAttribute{
																	Description:         "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																	MarkdownDescription: "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_path": schema.StringAttribute{
																	Description:         "volumePath is the path that identifies vSphere volume vmdk",
																	MarkdownDescription: "volumePath is the path that identifies vSphere volume vmdk",
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

							"traits": schema.SingleNestedAttribute{
								Description:         "the traits needed to run this Integration",
								MarkdownDescription: "the traits needed to run this Integration",
								Attributes: map[string]schema.Attribute{
									"threescale": schema.SingleNestedAttribute{
										Description:         "Deprecated: for backward compatibility.",
										MarkdownDescription: "Deprecated: for backward compatibility.",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "TraitConfiguration parameters configuration",
												MarkdownDescription: "TraitConfiguration parameters configuration",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"addons": schema.MapAttribute{
										Description:         "The extension point with addon traits",
										MarkdownDescription: "The extension point with addon traits",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"affinity": schema.SingleNestedAttribute{
										Description:         "The configuration of Affinity trait",
										MarkdownDescription: "The configuration of Affinity trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_affinity_labels": schema.ListAttribute{
												Description:         "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",
												MarkdownDescription: "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_affinity": schema.BoolAttribute{
												Description:         "Always co-locates multiple replicas of the integration in the same node (default 'false').",
												MarkdownDescription: "Always co-locates multiple replicas of the integration in the same node (default 'false').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_affinity_labels": schema.ListAttribute{
												Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",
												MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_anti_affinity": schema.BoolAttribute{
												Description:         "Never co-locates multiple replicas of the integration in the same node (default 'false').",
												MarkdownDescription: "Never co-locates multiple replicas of the integration in the same node (default 'false').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_anti_affinity_labels": schema.ListAttribute{
												Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",
												MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",
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

									"builder": schema.SingleNestedAttribute{
										Description:         "The configuration of Builder trait",
										MarkdownDescription: "The configuration of Builder trait",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "When using 'pod' strategy, annotation to use for the builder pod.",
												MarkdownDescription: "When using 'pod' strategy, annotation to use for the builder pod.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"base_image": schema.StringAttribute{
												Description:         "Specify a base image. In order to have the application working properly it must be a container image which has a Java JDK installed and ready to use on path (ie '/usr/bin/java').",
												MarkdownDescription: "Specify a base image. In order to have the application working properly it must be a container image which has a Java JDK installed and ready to use on path (ie '/usr/bin/java').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"incremental_image_build": schema.BoolAttribute{
												Description:         "Use the incremental image build option, to reuse existing containers (default 'true')",
												MarkdownDescription: "Use the incremental image build option, to reuse existing containers (default 'true')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"limit_cpu": schema.StringAttribute{
												Description:         "When using 'pod' strategy, the maximum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
												MarkdownDescription: "When using 'pod' strategy, the maximum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"limit_memory": schema.StringAttribute{
												Description:         "When using 'pod' strategy, the maximum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
												MarkdownDescription: "When using 'pod' strategy, the maximum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"maven_profiles": schema.ListAttribute{
												Description:         "A list of references pointing to configmaps/secrets that contains a maven profile. This configmap/secret is a resource of the IntegrationKit created, therefore it needs to be present in the namespace where the operator is going to create the IntegrationKit. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
												MarkdownDescription: "A list of references pointing to configmaps/secrets that contains a maven profile. This configmap/secret is a resource of the IntegrationKit created, therefore it needs to be present in the namespace where the operator is going to create the IntegrationKit. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_selector": schema.MapAttribute{
												Description:         "Defines a set of nodes the builder pod is eligible to be scheduled on, based on labels on the node.",
												MarkdownDescription: "Defines a set of nodes the builder pod is eligible to be scheduled on, based on labels on the node.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"order_strategy": schema.StringAttribute{
												Description:         "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default is the platform default)",
												MarkdownDescription: "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default is the platform default)",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("dependencies", "fifo", "sequential"),
												},
											},

											"platforms": schema.ListAttribute{
												Description:         "The list of manifest platforms to use to build a container image (default 'linux/amd64').",
												MarkdownDescription: "The list of manifest platforms to use to build a container image (default 'linux/amd64').",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"properties": schema.ListAttribute{
												Description:         "A list of properties to be provided to the build task",
												MarkdownDescription: "A list of properties to be provided to the build task",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"request_cpu": schema.StringAttribute{
												Description:         "When using 'pod' strategy, the minimum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
												MarkdownDescription: "When using 'pod' strategy, the minimum amount of CPU required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"request_memory": schema.StringAttribute{
												Description:         "When using 'pod' strategy, the minimum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
												MarkdownDescription: "When using 'pod' strategy, the minimum amount of memory required by the pod builder. Deprecated: use TasksRequestCPU instead with task name 'builder'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strategy": schema.StringAttribute{
												Description:         "The strategy to use, either 'pod' or 'routine' (default 'routine')",
												MarkdownDescription: "The strategy to use, either 'pod' or 'routine' (default 'routine')",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("pod", "routine"),
												},
											},

											"tasks": schema.ListAttribute{
												Description:         "A list of tasks to be executed (available only when using 'pod' strategy) with format '<name>;<container-image>;<container-command>'.",
												MarkdownDescription: "A list of tasks to be executed (available only when using 'pod' strategy) with format '<name>;<container-image>;<container-command>'.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tasks_filter": schema.StringAttribute{
												Description:         "A list of tasks sorted by the order of execution in a csv format, ie, '<taskName1>,<taskName2>,...'. Mind that you must include also the operator tasks ('builder', 'quarkus-native', 'package', 'jib', 's2i') if you need to execute them. Useful only with 'pod' strategy.",
												MarkdownDescription: "A list of tasks sorted by the order of execution in a csv format, ie, '<taskName1>,<taskName2>,...'. Mind that you must include also the operator tasks ('builder', 'quarkus-native', 'package', 'jib', 's2i') if you need to execute them. Useful only with 'pod' strategy.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tasks_limit_cpu": schema.ListAttribute{
												Description:         "A list of limit cpu configuration for the specific task with format '<task-name>:<limit-cpu-conf>'.",
												MarkdownDescription: "A list of limit cpu configuration for the specific task with format '<task-name>:<limit-cpu-conf>'.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tasks_limit_memory": schema.ListAttribute{
												Description:         "A list of limit memory configuration for the specific task with format '<task-name>:<limit-memory-conf>'.",
												MarkdownDescription: "A list of limit memory configuration for the specific task with format '<task-name>:<limit-memory-conf>'.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tasks_request_cpu": schema.ListAttribute{
												Description:         "A list of request cpu configuration for the specific task with format '<task-name>:<request-cpu-conf>'.",
												MarkdownDescription: "A list of request cpu configuration for the specific task with format '<task-name>:<request-cpu-conf>'.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tasks_request_memory": schema.ListAttribute{
												Description:         "A list of request memory configuration for the specific task with format '<task-name>:<request-memory-conf>'.",
												MarkdownDescription: "A list of request memory configuration for the specific task with format '<task-name>:<request-memory-conf>'.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"verbose": schema.BoolAttribute{
												Description:         "Enable verbose logging on build components that support it (e.g. Kaniko build pod). Deprecated no longer in use",
												MarkdownDescription: "Enable verbose logging on build components that support it (e.g. Kaniko build pod). Deprecated no longer in use",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"camel": schema.SingleNestedAttribute{
										Description:         "The configuration of Camel trait",
										MarkdownDescription: "The configuration of Camel trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"properties": schema.ListAttribute{
												Description:         "A list of properties to be provided to the Integration runtime",
												MarkdownDescription: "A list of properties to be provided to the Integration runtime",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"runtime_version": schema.StringAttribute{
												Description:         "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform. You can use a fixed version (for example '3.2.3') or a semantic version (for example '3.x') which will try to resolve to the best matching Catalog existing on the cluster.",
												MarkdownDescription: "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform. You can use a fixed version (for example '3.2.3') or a semantic version (for example '3.x') which will try to resolve to the best matching Catalog existing on the cluster.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"container": schema.SingleNestedAttribute{
										Description:         "The configuration of Container trait",
										MarkdownDescription: "The configuration of Container trait",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "Security Context AllowPrivilegeEscalation configuration (default false).",
												MarkdownDescription: "Security Context AllowPrivilegeEscalation configuration (default false).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"auto": schema.BoolAttribute{
												Description:         "To automatically enable the trait",
												MarkdownDescription: "To automatically enable the trait",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"capabilities_add": schema.ListAttribute{
												Description:         "Security Context Capabilities Add configuration (default none).",
												MarkdownDescription: "Security Context Capabilities Add configuration (default none).",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"capabilities_drop": schema.ListAttribute{
												Description:         "Security Context Capabilities Drop configuration (default ALL).",
												MarkdownDescription: "Security Context Capabilities Drop configuration (default ALL).",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"expose": schema.BoolAttribute{
												Description:         "Can be used to enable/disable exposure via kubernetes Service.",
												MarkdownDescription: "Can be used to enable/disable exposure via kubernetes Service.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image": schema.StringAttribute{
												Description:         "The main container image to use for the Integration. When using this parameter the operator will create a synthetic IntegrationKit which won't be able to execute traits requiring CamelCatalog. If the container image you're using is coming from an IntegrationKit, use instead Integration '.spec.integrationKit' parameter. If you're moving the Integration across environments, you will also need to create an 'external' IntegrationKit.",
												MarkdownDescription: "The main container image to use for the Integration. When using this parameter the operator will create a synthetic IntegrationKit which won't be able to execute traits requiring CamelCatalog. If the container image you're using is coming from an IntegrationKit, use instead Integration '.spec.integrationKit' parameter. If you're moving the Integration across environments, you will also need to create an 'external' IntegrationKit.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_pull_policy": schema.StringAttribute{
												Description:         "The pull policy: Always|Never|IfNotPresent",
												MarkdownDescription: "The pull policy: Always|Never|IfNotPresent",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
												},
											},

											"limit_cpu": schema.StringAttribute{
												Description:         "The maximum amount of CPU to be provided (default 500 millicores).",
												MarkdownDescription: "The maximum amount of CPU to be provided (default 500 millicores).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"limit_memory": schema.StringAttribute{
												Description:         "The maximum amount of memory to be provided (default 512 Mi).",
												MarkdownDescription: "The maximum amount of memory to be provided (default 512 Mi).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "The main container name. It's named 'integration' by default.",
												MarkdownDescription: "The main container name. It's named 'integration' by default.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "To configure a different port exposed by the container (default '8080').",
												MarkdownDescription: "To configure a different port exposed by the container (default '8080').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port_name": schema.StringAttribute{
												Description:         "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",
												MarkdownDescription: "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"request_cpu": schema.StringAttribute{
												Description:         "The minimum amount of CPU required (default 125 millicores).",
												MarkdownDescription: "The minimum amount of CPU required (default 125 millicores).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"request_memory": schema.StringAttribute{
												Description:         "The minimum amount of memory required (default 128 Mi).",
												MarkdownDescription: "The minimum amount of memory required (default 128 Mi).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "Security Context RunAsNonRoot configuration (default false).",
												MarkdownDescription: "Security Context RunAsNonRoot configuration (default false).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "Security Context RunAsUser configuration (default none): this value is automatically retrieved in Openshift clusters when not explicitly set.",
												MarkdownDescription: "Security Context RunAsUser configuration (default none): this value is automatically retrieved in Openshift clusters when not explicitly set.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"seccomp_profile_type": schema.StringAttribute{
												Description:         "Security Context SeccompProfileType configuration (default RuntimeDefault).",
												MarkdownDescription: "Security Context SeccompProfileType configuration (default RuntimeDefault).",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Unconfined", "RuntimeDefault"),
												},
											},

											"service_port": schema.Int64Attribute{
												Description:         "To configure under which service port the container port is to be exposed (default '80').",
												MarkdownDescription: "To configure under which service port the container port is to be exposed (default '80').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_port_name": schema.StringAttribute{
												Description:         "To configure under which service port name the container port is to be exposed (default 'http').",
												MarkdownDescription: "To configure under which service port name the container port is to be exposed (default 'http').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cron": schema.SingleNestedAttribute{
										Description:         "The configuration of Cron trait",
										MarkdownDescription: "The configuration of Cron trait",
										Attributes: map[string]schema.Attribute{
											"active_deadline_seconds": schema.Int64Attribute{
												Description:         "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",
												MarkdownDescription: "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"auto": schema.BoolAttribute{
												Description:         "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer). It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",
												MarkdownDescription: "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer). It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"backoff_limit": schema.Int64Attribute{
												Description:         "Specifies the number of retries before marking the job failed. It defaults to 2.",
												MarkdownDescription: "Specifies the number of retries before marking the job failed. It defaults to 2.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"components": schema.StringAttribute{
												Description:         "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library). Supported components are currently: 'cron', 'timer' and 'quartz'.",
												MarkdownDescription: "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library). Supported components are currently: 'cron', 'timer' and 'quartz'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"concurrency_policy": schema.StringAttribute{
												Description:         "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",
												MarkdownDescription: "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Allow", "Forbid", "Replace"),
												},
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"fallback": schema.BoolAttribute{
												Description:         "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",
												MarkdownDescription: "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"schedule": schema.StringAttribute{
												Description:         "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",
												MarkdownDescription: "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"starting_deadline_seconds": schema.Int64Attribute{
												Description:         "Optional deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones.",
												MarkdownDescription: "Optional deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"time_zone": schema.StringAttribute{
												Description:         "The timezone that the CronJob will run on",
												MarkdownDescription: "The timezone that the CronJob will run on",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"dependencies": schema.SingleNestedAttribute{
										Description:         "The configuration of Dependencies trait",
										MarkdownDescription: "The configuration of Dependencies trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployer": schema.SingleNestedAttribute{
										Description:         "The configuration of Deployer trait",
										MarkdownDescription: "The configuration of Deployer trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",
												MarkdownDescription: "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("deployment", "cron-job", "knative-service"),
												},
											},

											"use_ssa": schema.BoolAttribute{
												Description:         "Deprecated: won't be able to enforce client side update in the future. Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",
												MarkdownDescription: "Deprecated: won't be able to enforce client side update in the future. Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment": schema.SingleNestedAttribute{
										Description:         "The configuration of Deployment trait",
										MarkdownDescription: "The configuration of Deployment trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"progress_deadline_seconds": schema.Int64Attribute{
												Description:         "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to '60s'.",
												MarkdownDescription: "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to '60s'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"rolling_update_max_surge": schema.StringAttribute{
												Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to '25%'.",
												MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to '25%'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"rolling_update_max_unavailable": schema.StringAttribute{
												Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to '25%'.",
												MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to '25%'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"strategy": schema.StringAttribute{
												Description:         "The deployment strategy to use to replace existing pods with new ones.",
												MarkdownDescription: "The deployment strategy to use to replace existing pods with new ones.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Recreate", "RollingUpdate"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"environment": schema.SingleNestedAttribute{
										Description:         "The configuration of Environment trait",
										MarkdownDescription: "The configuration of Environment trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"container_meta": schema.BoolAttribute{
												Description:         "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",
												MarkdownDescription: "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_proxy": schema.BoolAttribute{
												Description:         "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",
												MarkdownDescription: "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"vars": schema.ListAttribute{
												Description:         "A list of environment variables to be added to the integration container. The syntax is either VAR=VALUE or VAR=[configmap|secret]:name/key, where name represents the resource name, and key represents the resource key to be mapped as and environment variable. These take precedence over any previously defined environment variables.",
												MarkdownDescription: "A list of environment variables to be added to the integration container. The syntax is either VAR=VALUE or VAR=[configmap|secret]:name/key, where name represents the resource name, and key represents the resource key to be mapped as and environment variable. These take precedence over any previously defined environment variables.",
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

									"error_handler": schema.SingleNestedAttribute{
										Description:         "The configuration of Error Handler trait",
										MarkdownDescription: "The configuration of Error Handler trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ref": schema.StringAttribute{
												Description:         "The error handler ref name provided or found in application properties",
												MarkdownDescription: "The error handler ref name provided or found in application properties",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"gc": schema.SingleNestedAttribute{
										Description:         "The configuration of GC trait",
										MarkdownDescription: "The configuration of GC trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"discovery_cache": schema.StringAttribute{
												Description:         "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",
												MarkdownDescription: "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("disabled", "disk", "memory"),
												},
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"health": schema.SingleNestedAttribute{
										Description:         "The configuration of Health trait",
										MarkdownDescription: "The configuration of Health trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"liveness_failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",
												MarkdownDescription: "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"liveness_initial_delay": schema.Int64Attribute{
												Description:         "Number of seconds after the container has started before the liveness probe is initiated.",
												MarkdownDescription: "Number of seconds after the container has started before the liveness probe is initiated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"liveness_period": schema.Int64Attribute{
												Description:         "How often to perform the liveness probe.",
												MarkdownDescription: "How often to perform the liveness probe.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"liveness_probe": schema.StringAttribute{
												Description:         "The liveness probe path to use (default provided by the Catalog runtime used).",
												MarkdownDescription: "The liveness probe path to use (default provided by the Catalog runtime used).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"liveness_probe_enabled": schema.BoolAttribute{
												Description:         "Configures the liveness probe for the integration container (default 'false').",
												MarkdownDescription: "Configures the liveness probe for the integration container (default 'false').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"liveness_scheme": schema.StringAttribute{
												Description:         "Scheme to use when connecting to the liveness probe (default 'HTTP').",
												MarkdownDescription: "Scheme to use when connecting to the liveness probe (default 'HTTP').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"liveness_success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",
												MarkdownDescription: "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"liveness_timeout": schema.Int64Attribute{
												Description:         "Number of seconds after which the liveness probe times out.",
												MarkdownDescription: "Number of seconds after which the liveness probe times out.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"readiness_failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",
												MarkdownDescription: "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"readiness_initial_delay": schema.Int64Attribute{
												Description:         "Number of seconds after the container has started before the readiness probe is initiated.",
												MarkdownDescription: "Number of seconds after the container has started before the readiness probe is initiated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"readiness_period": schema.Int64Attribute{
												Description:         "How often to perform the readiness probe.",
												MarkdownDescription: "How often to perform the readiness probe.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"readiness_probe": schema.StringAttribute{
												Description:         "The readiness probe path to use (default provided by the Catalog runtime used).",
												MarkdownDescription: "The readiness probe path to use (default provided by the Catalog runtime used).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"readiness_probe_enabled": schema.BoolAttribute{
												Description:         "Configures the readiness probe for the integration container (default 'true').",
												MarkdownDescription: "Configures the readiness probe for the integration container (default 'true').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"readiness_scheme": schema.StringAttribute{
												Description:         "Scheme to use when connecting to the readiness probe (default 'HTTP').",
												MarkdownDescription: "Scheme to use when connecting to the readiness probe (default 'HTTP').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"readiness_success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",
												MarkdownDescription: "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"readiness_timeout": schema.Int64Attribute{
												Description:         "Number of seconds after which the readiness probe times out.",
												MarkdownDescription: "Number of seconds after which the readiness probe times out.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"startup_failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the startup probe to be considered failed after having succeeded.",
												MarkdownDescription: "Minimum consecutive failures for the startup probe to be considered failed after having succeeded.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"startup_initial_delay": schema.Int64Attribute{
												Description:         "Number of seconds after the container has started before the startup probe is initiated.",
												MarkdownDescription: "Number of seconds after the container has started before the startup probe is initiated.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"startup_period": schema.Int64Attribute{
												Description:         "How often to perform the startup probe.",
												MarkdownDescription: "How often to perform the startup probe.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"startup_probe": schema.StringAttribute{
												Description:         "The startup probe path to use (default provided by the Catalog runtime used).",
												MarkdownDescription: "The startup probe path to use (default provided by the Catalog runtime used).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"startup_probe_enabled": schema.BoolAttribute{
												Description:         "Configures the startup probe for the integration container (default 'false').",
												MarkdownDescription: "Configures the startup probe for the integration container (default 'false').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"startup_scheme": schema.StringAttribute{
												Description:         "Scheme to use when connecting to the startup probe (default 'HTTP').",
												MarkdownDescription: "Scheme to use when connecting to the startup probe (default 'HTTP').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"startup_success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the startup probe to be considered successful after having failed.",
												MarkdownDescription: "Minimum consecutive successes for the startup probe to be considered successful after having failed.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"startup_timeout": schema.Int64Attribute{
												Description:         "Number of seconds after which the startup probe times out.",
												MarkdownDescription: "Number of seconds after which the startup probe times out.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"ingress": schema.SingleNestedAttribute{
										Description:         "The configuration of Ingress trait",
										MarkdownDescription: "The configuration of Ingress trait",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "The annotations added to the ingress. This can be used to set controller specific annotations, e.g., when using the NGINX Ingress controller: See https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md",
												MarkdownDescription: "The annotations added to the ingress. This can be used to set controller specific annotations, e.g., when using the NGINX Ingress controller: See https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"auto": schema.BoolAttribute{
												Description:         "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",
												MarkdownDescription: "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host": schema.StringAttribute{
												Description:         "To configure the host exposed by the ingress.",
												MarkdownDescription: "To configure the host exposed by the ingress.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ingress_class_name": schema.StringAttribute{
												Description:         "The Ingress class name as defined by the Ingress spec See https://kubernetes.io/docs/concepts/services-networking/ingress/",
												MarkdownDescription: "The Ingress class name as defined by the Ingress spec See https://kubernetes.io/docs/concepts/services-networking/ingress/",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"path": schema.StringAttribute{
												Description:         "To configure the path exposed by the ingress (default '/').",
												MarkdownDescription: "To configure the path exposed by the ingress (default '/').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"path_type": schema.StringAttribute{
												Description:         "To configure the path type exposed by the ingress. One of 'Exact', 'Prefix', 'ImplementationSpecific' (default to 'Prefix').",
												MarkdownDescription: "To configure the path type exposed by the ingress. One of 'Exact', 'Prefix', 'ImplementationSpecific' (default to 'Prefix').",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Exact", "Prefix", "ImplementationSpecific"),
												},
											},

											"tls_hosts": schema.ListAttribute{
												Description:         "To configure tls hosts",
												MarkdownDescription: "To configure tls hosts",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_secret_name": schema.StringAttribute{
												Description:         "To configure tls secret name",
												MarkdownDescription: "To configure tls secret name",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"istio": schema.SingleNestedAttribute{
										Description:         "The configuration of Istio trait",
										MarkdownDescription: "The configuration of Istio trait",
										Attributes: map[string]schema.Attribute{
											"allow": schema.StringAttribute{
												Description:         "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",
												MarkdownDescription: "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"inject": schema.BoolAttribute{
												Description:         "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",
												MarkdownDescription: "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"jolokia": schema.SingleNestedAttribute{
										Description:         "The configuration of Jolokia trait",
										MarkdownDescription: "The configuration of Jolokia trait",
										Attributes: map[string]schema.Attribute{
											"ca_cert": schema.StringAttribute{
												Description:         "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",
												MarkdownDescription: "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"client_principal": schema.ListAttribute{
												Description:         "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",
												MarkdownDescription: "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"discovery_enabled": schema.BoolAttribute{
												Description:         "Listen for multicast requests (default 'false')",
												MarkdownDescription: "Listen for multicast requests (default 'false')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"extended_client_check": schema.BoolAttribute{
												Description:         "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",
												MarkdownDescription: "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host": schema.StringAttribute{
												Description:         "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",
												MarkdownDescription: "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"options": schema.ListAttribute{
												Description:         "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",
												MarkdownDescription: "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"password": schema.StringAttribute{
												Description:         "The password used for authentication, applicable when the 'user' option is set.",
												MarkdownDescription: "The password used for authentication, applicable when the 'user' option is set.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "The Jolokia endpoint port (default '8778').",
												MarkdownDescription: "The Jolokia endpoint port (default '8778').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"protocol": schema.StringAttribute{
												Description:         "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",
												MarkdownDescription: "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"use_ssl_client_authentication": schema.BoolAttribute{
												Description:         "Whether client certificates should be used for authentication (default 'true' for OpenShift).",
												MarkdownDescription: "Whether client certificates should be used for authentication (default 'true' for OpenShift).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"user": schema.StringAttribute{
												Description:         "The user to be used for authentication",
												MarkdownDescription: "The user to be used for authentication",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"jvm": schema.SingleNestedAttribute{
										Description:         "The configuration of JVM trait",
										MarkdownDescription: "The configuration of JVM trait",
										Attributes: map[string]schema.Attribute{
											"classpath": schema.StringAttribute{
												Description:         "Additional JVM classpath (use 'Linux' classpath separator)",
												MarkdownDescription: "Additional JVM classpath (use 'Linux' classpath separator)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"debug": schema.BoolAttribute{
												Description:         "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",
												MarkdownDescription: "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"debug_address": schema.StringAttribute{
												Description:         "Transport address at which to listen for the newly launched JVM (default '*:5005')",
												MarkdownDescription: "Transport address at which to listen for the newly launched JVM (default '*:5005')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"debug_suspend": schema.BoolAttribute{
												Description:         "Suspends the target JVM immediately before the main class is loaded",
												MarkdownDescription: "Suspends the target JVM immediately before the main class is loaded",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"jar": schema.StringAttribute{
												Description:         "The Jar dependency which will run the application. Leave it empty for managed Integrations.",
												MarkdownDescription: "The Jar dependency which will run the application. Leave it empty for managed Integrations.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"options": schema.ListAttribute{
												Description:         "A list of JVM options",
												MarkdownDescription: "A list of JVM options",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"print_command": schema.BoolAttribute{
												Description:         "Prints the command used the start the JVM in the container logs (default 'true') Deprecated: no longer in use.",
												MarkdownDescription: "Prints the command used the start the JVM in the container logs (default 'true') Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"kamelets": schema.SingleNestedAttribute{
										Description:         "The configuration of Kamelets trait",
										MarkdownDescription: "The configuration of Kamelets trait",
										Attributes: map[string]schema.Attribute{
											"auto": schema.BoolAttribute{
												Description:         "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",
												MarkdownDescription: "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"list": schema.StringAttribute{
												Description:         "Comma separated list of Kamelet names to load into the current integration",
												MarkdownDescription: "Comma separated list of Kamelet names to load into the current integration",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"mount_point": schema.StringAttribute{
												Description:         "The directory where the application mounts and reads Kamelet spec (default '/etc/camel/kamelets')",
												MarkdownDescription: "The directory where the application mounts and reads Kamelet spec (default '/etc/camel/kamelets')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"keda": schema.SingleNestedAttribute{
										Description:         "Deprecated: for backward compatibility.",
										MarkdownDescription: "Deprecated: for backward compatibility.",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "TraitConfiguration parameters configuration",
												MarkdownDescription: "TraitConfiguration parameters configuration",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"knative": schema.SingleNestedAttribute{
										Description:         "The configuration of Knative trait",
										MarkdownDescription: "The configuration of Knative trait",
										Attributes: map[string]schema.Attribute{
											"auto": schema.BoolAttribute{
												Description:         "Enable automatic discovery of all trait properties.",
												MarkdownDescription: "Enable automatic discovery of all trait properties.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"channel_sinks": schema.ListAttribute{
												Description:         "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",
												MarkdownDescription: "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"channel_sources": schema.ListAttribute{
												Description:         "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",
												MarkdownDescription: "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"config": schema.StringAttribute{
												Description:         "Can be used to inject a Knative complete configuration in JSON format.",
												MarkdownDescription: "Can be used to inject a Knative complete configuration in JSON format.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"endpoint_sinks": schema.ListAttribute{
												Description:         "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",
												MarkdownDescription: "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"endpoint_sources": schema.ListAttribute{
												Description:         "List of channels used as source of integration routes.",
												MarkdownDescription: "List of channels used as source of integration routes.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"event_sinks": schema.ListAttribute{
												Description:         "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",
												MarkdownDescription: "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"event_sources": schema.ListAttribute{
												Description:         "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",
												MarkdownDescription: "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"filter_event_type": schema.BoolAttribute{
												Description:         "Enables the default filtering for the Knative trigger using the event type If this is true, the created Knative trigger uses the event type as a filter on the event stream when no other filter criteria is given. (default: true)",
												MarkdownDescription: "Enables the default filtering for the Knative trigger using the event type If this is true, the created Knative trigger uses the event type as a filter on the event stream when no other filter criteria is given. (default: true)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"filter_source_channels": schema.BoolAttribute{
												Description:         "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",
												MarkdownDescription: "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"filters": schema.ListAttribute{
												Description:         "Sets filter attributes on the event stream (such as event type, source, subject and so on). A list of key-value pairs that represent filter attributes and its values. The syntax is KEY=VALUE, e.g., 'source='my.source''. Filter attributes get set on the Knative trigger that is being created as part of this integration.",
												MarkdownDescription: "Sets filter attributes on the event stream (such as event type, source, subject and so on). A list of key-value pairs that represent filter attributes and its values. The syntax is KEY=VALUE, e.g., 'source='my.source''. Filter attributes get set on the Knative trigger that is being created as part of this integration.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace_label": schema.BoolAttribute{
												Description:         "Enables the camel-k-operator to set the 'bindings.knative.dev/include=true' label to the namespace As Knative requires this label to perform injection of K_SINK URL into the service. If this is false, the integration pod may start and fail, read the SinkBinding Knative documentation. (default: true)",
												MarkdownDescription: "Enables the camel-k-operator to set the 'bindings.knative.dev/include=true' label to the namespace As Knative requires this label to perform injection of K_SINK URL into the service. If this is false, the integration pod may start and fail, read the SinkBinding Knative documentation. (default: true)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sink_binding": schema.BoolAttribute{
												Description:         "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",
												MarkdownDescription: "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"knative_service": schema.SingleNestedAttribute{
										Description:         "The configuration of Knative Service trait",
										MarkdownDescription: "The configuration of Knative Service trait",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "The annotations added to route. This can be used to set knative service specific annotations CLI usage example: -t 'knative-service.annotations.'haproxy.router.openshift.io/balance'=true'",
												MarkdownDescription: "The annotations added to route. This can be used to set knative service specific annotations CLI usage example: -t 'knative-service.annotations.'haproxy.router.openshift.io/balance'=true'",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"auto": schema.BoolAttribute{
												Description:         "Automatically deploy the integration as Knative service when all conditions hold: * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",
												MarkdownDescription: "Automatically deploy the integration as Knative service when all conditions hold: * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"autoscaling_metric": schema.StringAttribute{
												Description:         "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling). Refer to the Knative documentation for more information.",
												MarkdownDescription: "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling). Refer to the Knative documentation for more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"autoscaling_target": schema.Int64Attribute{
												Description:         "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod. Refer to the Knative documentation for more information.",
												MarkdownDescription: "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod. Refer to the Knative documentation for more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"class": schema.StringAttribute{
												Description:         "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling). Refer to the Knative documentation for more information.",
												MarkdownDescription: "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling). Refer to the Knative documentation for more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("kpa.autoscaling.knative.dev", "hpa.autoscaling.knative.dev"),
												},
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_scale": schema.Int64Attribute{
												Description:         "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation. Refer to the Knative documentation for more information.",
												MarkdownDescription: "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation. Refer to the Knative documentation for more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"min_scale": schema.Int64Attribute{
												Description:         "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time. Refer to the Knative documentation for more information.",
												MarkdownDescription: "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time. Refer to the Knative documentation for more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"rollout_duration": schema.StringAttribute{
												Description:         "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",
												MarkdownDescription: "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The maximum duration in seconds that the request instance is allowed to respond to a request. This field propagates to the integration pod's terminationGracePeriodSeconds Refer to the Knative documentation for more information.",
												MarkdownDescription: "The maximum duration in seconds that the request instance is allowed to respond to a request. This field propagates to the integration pod's terminationGracePeriodSeconds Refer to the Knative documentation for more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"visibility": schema.StringAttribute{
												Description:         "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service. Refer to the Knative documentation for more information.",
												MarkdownDescription: "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service. Refer to the Knative documentation for more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("cluster-local"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"logging": schema.SingleNestedAttribute{
										Description:         "The configuration of Logging trait",
										MarkdownDescription: "The configuration of Logging trait",
										Attributes: map[string]schema.Attribute{
											"color": schema.BoolAttribute{
												Description:         "Colorize the log output",
												MarkdownDescription: "Colorize the log output",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"format": schema.StringAttribute{
												Description:         "Logs message format",
												MarkdownDescription: "Logs message format",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"json": schema.BoolAttribute{
												Description:         "Output the logs in JSON",
												MarkdownDescription: "Output the logs in JSON",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"json_pretty_print": schema.BoolAttribute{
												Description:         "Enable 'pretty printing' of the JSON logs",
												MarkdownDescription: "Enable 'pretty printing' of the JSON logs",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"level": schema.StringAttribute{
												Description:         "Adjust the logging level (defaults to 'INFO')",
												MarkdownDescription: "Adjust the logging level (defaults to 'INFO')",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("FATAL", "WARN", "INFO", "DEBUG", "TRACE"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"master": schema.SingleNestedAttribute{
										Description:         "Deprecated: for backward compatibility.",
										MarkdownDescription: "Deprecated: for backward compatibility.",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "TraitConfiguration parameters configuration",
												MarkdownDescription: "TraitConfiguration parameters configuration",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"mount": schema.SingleNestedAttribute{
										Description:         "The configuration of Mount trait",
										MarkdownDescription: "The configuration of Mount trait",
										Attributes: map[string]schema.Attribute{
											"configs": schema.ListAttribute{
												Description:         "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",
												MarkdownDescription: "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"empty_dirs": schema.ListAttribute{
												Description:         "A list of EmptyDir volumes to be mounted. Syntax: [name:/container/path]",
												MarkdownDescription: "A list of EmptyDir volumes to be mounted. Syntax: [name:/container/path]",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"hot_reload": schema.BoolAttribute{
												Description:         "Enable 'hot reload' when a secret/configmap mounted is edited (default 'false'). The configmap/secret must be marked with 'camel.apache.org/integration' label to be taken in account. The resource will be watched for any kind change, also for changes in metadata.",
												MarkdownDescription: "Enable 'hot reload' when a secret/configmap mounted is edited (default 'false'). The configmap/secret must be marked with 'camel.apache.org/integration' label to be taken in account. The resource will be watched for any kind change, also for changes in metadata.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.ListAttribute{
												Description:         "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",
												MarkdownDescription: "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"scan_kamelets_implicit_label_secrets": schema.BoolAttribute{
												Description:         "Deprecated: no longer available since version 2.5.",
												MarkdownDescription: "Deprecated: no longer available since version 2.5.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"volumes": schema.ListAttribute{
												Description:         "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",
												MarkdownDescription: "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",
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

									"openapi": schema.SingleNestedAttribute{
										Description:         "The configuration of OpenAPI trait",
										MarkdownDescription: "The configuration of OpenAPI trait",
										Attributes: map[string]schema.Attribute{
											"configmaps": schema.ListAttribute{
												Description:         "The configmaps holding the spec of the OpenAPI (compatible with > 3.0 spec only).",
												MarkdownDescription: "The configmaps holding the spec of the OpenAPI (compatible with > 3.0 spec only).",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"owner": schema.SingleNestedAttribute{
										Description:         "The configuration of Owner trait",
										MarkdownDescription: "The configuration of Owner trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_annotations": schema.ListAttribute{
												Description:         "The set of annotations to be transferred",
												MarkdownDescription: "The set of annotations to be transferred",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_labels": schema.ListAttribute{
												Description:         "The set of labels to be transferred",
												MarkdownDescription: "The set of labels to be transferred",
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

									"pdb": schema.SingleNestedAttribute{
										Description:         "The configuration of PDB trait",
										MarkdownDescription: "The configuration of PDB trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_unavailable": schema.StringAttribute{
												Description:         "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",
												MarkdownDescription: "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"min_available": schema.StringAttribute{
												Description:         "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",
												MarkdownDescription: "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"platform": schema.SingleNestedAttribute{
										Description:         "The configuration of Platform trait",
										MarkdownDescription: "The configuration of Platform trait",
										Attributes: map[string]schema.Attribute{
											"auto": schema.BoolAttribute{
												Description:         "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift or when a registry address is set). Deprecated: Platform is auto generated by the operator install procedure - maintained for backward compatibility",
												MarkdownDescription: "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift or when a registry address is set). Deprecated: Platform is auto generated by the operator install procedure - maintained for backward compatibility",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"create_default": schema.BoolAttribute{
												Description:         "To create a default (empty) platform when the platform is missing. Deprecated: Platform is auto generated by the operator install procedure - maintained for backward compatibility",
												MarkdownDescription: "To create a default (empty) platform when the platform is missing. Deprecated: Platform is auto generated by the operator install procedure - maintained for backward compatibility",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"global": schema.BoolAttribute{
												Description:         "Indicates if the platform should be created globally in the case of global operator (default true). Deprecated: Platform is auto generated by the operator install procedure - maintained for backward compatibility",
												MarkdownDescription: "Indicates if the platform should be created globally in the case of global operator (default true). Deprecated: Platform is auto generated by the operator install procedure - maintained for backward compatibility",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod": schema.SingleNestedAttribute{
										Description:         "The configuration of Pod trait",
										MarkdownDescription: "The configuration of Pod trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"prometheus": schema.SingleNestedAttribute{
										Description:         "The configuration of Prometheus trait",
										MarkdownDescription: "The configuration of Prometheus trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_monitor": schema.BoolAttribute{
												Description:         "Whether a 'PodMonitor' resource is created (default 'true').",
												MarkdownDescription: "Whether a 'PodMonitor' resource is created (default 'true').",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_monitor_labels": schema.ListAttribute{
												Description:         "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",
												MarkdownDescription: "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",
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

									"pull_secret": schema.SingleNestedAttribute{
										Description:         "The configuration of Pull Secret trait",
										MarkdownDescription: "The configuration of Pull Secret trait",
										Attributes: map[string]schema.Attribute{
											"auto": schema.BoolAttribute{
												Description:         "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",
												MarkdownDescription: "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_puller_delegation": schema.BoolAttribute{
												Description:         "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",
												MarkdownDescription: "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret_name": schema.StringAttribute{
												Description:         "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",
												MarkdownDescription: "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"quarkus": schema.SingleNestedAttribute{
										Description:         "The configuration of Quarkus trait",
										MarkdownDescription: "The configuration of Quarkus trait",
										Attributes: map[string]schema.Attribute{
											"build_mode": schema.ListAttribute{
												Description:         "The Quarkus mode to run: either 'jvm' or 'native' (default 'jvm'). In case both 'jvm' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'jvm' one once ready.",
												MarkdownDescription: "The Quarkus mode to run: either 'jvm' or 'native' (default 'jvm'). In case both 'jvm' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'jvm' one once ready.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"native_base_image": schema.StringAttribute{
												Description:         "The base image to use when running a native build (default 'quay.io/quarkus/quarkus-micro-image:2.0')",
												MarkdownDescription: "The base image to use when running a native build (default 'quay.io/quarkus/quarkus-micro-image:2.0')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"native_builder_image": schema.StringAttribute{
												Description:         "The image containing the tooling required for a native build (by default it will use the one provided in the runtime catalog)",
												MarkdownDescription: "The image containing the tooling required for a native build (by default it will use the one provided in the runtime catalog)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"package_types": schema.ListAttribute{
												Description:         "The Quarkus package types, 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the native kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists. Deprecated: use 'build-mode' instead.",
												MarkdownDescription: "The Quarkus package types, 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the native kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists. Deprecated: use 'build-mode' instead.",
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

									"registry": schema.SingleNestedAttribute{
										Description:         "The configuration of Registry trait (support removed since version 2.5.0). Deprecated: use jvm trait or read documentation.",
										MarkdownDescription: "The configuration of Registry trait (support removed since version 2.5.0). Deprecated: use jvm trait or read documentation.",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
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
										Description:         "The configuration of Route trait",
										MarkdownDescription: "The configuration of Route trait",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "The annotations added to route. This can be used to set route specific annotations For annotations options see https://docs.openshift.com/container-platform/3.11/architecture/networking/routes.html#route-specific-annotations CLI usage example: -t 'route.annotations.'haproxy.router.openshift.io/balance'=true'",
												MarkdownDescription: "The annotations added to route. This can be used to set route specific annotations For annotations options see https://docs.openshift.com/container-platform/3.11/architecture/networking/routes.html#route-specific-annotations CLI usage example: -t 'route.annotations.'haproxy.router.openshift.io/balance'=true'",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host": schema.StringAttribute{
												Description:         "To configure the host exposed by the route.",
												MarkdownDescription: "To configure the host exposed by the route.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_ca_certificate": schema.StringAttribute{
												Description:         "The TLS CA certificate contents. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "The TLS CA certificate contents. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_ca_certificate_secret": schema.StringAttribute{
												Description:         "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_certificate": schema.StringAttribute{
												Description:         "The TLS certificate contents. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "The TLS certificate contents. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_certificate_secret": schema.StringAttribute{
												Description:         "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_destination_ca_certificate": schema.StringAttribute{
												Description:         "The destination CA certificate provides the contents of the ca certificate of the final destination. When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "The destination CA certificate provides the contents of the ca certificate of the final destination. When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_destination_ca_certificate_secret": schema.StringAttribute{
												Description:         "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_insecure_edge_termination_policy": schema.StringAttribute{
												Description:         "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("None", "Allow", "Redirect"),
												},
											},

											"tls_key": schema.StringAttribute{
												Description:         "The TLS certificate key contents. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "The TLS certificate key contents. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_key_secret": schema.StringAttribute{
												Description:         "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls_termination": schema.StringAttribute{
												Description:         "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'. Refer to the OpenShift route documentation for additional information.",
												MarkdownDescription: "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'. Refer to the OpenShift route documentation for additional information.",
												Required:            false,
												Optional:            true,
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

									"security_context": schema.SingleNestedAttribute{
										Description:         "The configuration of Security Context trait",
										MarkdownDescription: "The configuration of Security Context trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Deprecated: no longer in use.",
												MarkdownDescription: "Deprecated: no longer in use.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "Security Context RunAsNonRoot configuration (default false).",
												MarkdownDescription: "Security Context RunAsNonRoot configuration (default false).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "Security Context RunAsUser configuration (default none): this value is automatically retrieved in Openshift clusters when not explicitly set.",
												MarkdownDescription: "Security Context RunAsUser configuration (default none): this value is automatically retrieved in Openshift clusters when not explicitly set.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"seccomp_profile_type": schema.StringAttribute{
												Description:         "Security Context SeccompProfileType configuration (default RuntimeDefault).",
												MarkdownDescription: "Security Context SeccompProfileType configuration (default RuntimeDefault).",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Unconfined", "RuntimeDefault"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"service": schema.SingleNestedAttribute{
										Description:         "The configuration of Service trait",
										MarkdownDescription: "The configuration of Service trait",
										Attributes: map[string]schema.Attribute{
											"auto": schema.BoolAttribute{
												Description:         "To automatically detect from the code if a Service needs to be created.",
												MarkdownDescription: "To automatically detect from the code if a Service needs to be created.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_port": schema.BoolAttribute{
												Description:         "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",
												MarkdownDescription: "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",
												MarkdownDescription: "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",
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

									"service_binding": schema.SingleNestedAttribute{
										Description:         "The configuration of Service Binding trait",
										MarkdownDescription: "The configuration of Service Binding trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"services": schema.ListAttribute{
												Description:         "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",
												MarkdownDescription: "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",
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

									"strimzi": schema.SingleNestedAttribute{
										Description:         "Deprecated: for backward compatibility.",
										MarkdownDescription: "Deprecated: for backward compatibility.",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "TraitConfiguration parameters configuration",
												MarkdownDescription: "TraitConfiguration parameters configuration",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"telemetry": schema.SingleNestedAttribute{
										Description:         "The configuration of Telemetry trait",
										MarkdownDescription: "The configuration of Telemetry trait",
										Attributes: map[string]schema.Attribute{
											"auto": schema.BoolAttribute{
												Description:         "Enables automatic configuration of the trait, including automatic discovery of the telemetry endpoint.",
												MarkdownDescription: "Enables automatic configuration of the trait, including automatic discovery of the telemetry endpoint.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"endpoint": schema.StringAttribute{
												Description:         "The target endpoint of the Telemetry service (automatically discovered by default)",
												MarkdownDescription: "The target endpoint of the Telemetry service (automatically discovered by default)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sampler": schema.StringAttribute{
												Description:         "The sampler of the telemetry used for tracing (default 'on')",
												MarkdownDescription: "The sampler of the telemetry used for tracing (default 'on')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sampler_parent_based": schema.BoolAttribute{
												Description:         "The sampler of the telemetry used for tracing is parent based (default 'true')",
												MarkdownDescription: "The sampler of the telemetry used for tracing is parent based (default 'true')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"sampler_ratio": schema.StringAttribute{
												Description:         "The sampler ratio of the telemetry used for tracing",
												MarkdownDescription: "The sampler ratio of the telemetry used for tracing",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_name": schema.StringAttribute{
												Description:         "The name of the service that publishes telemetry data (defaults to the integration name)",
												MarkdownDescription: "The name of the service that publishes telemetry data (defaults to the integration name)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"toleration": schema.SingleNestedAttribute{
										Description:         "The configuration of Toleration trait",
										MarkdownDescription: "The configuration of Toleration trait",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Can be used to enable or disable a trait. All traits share this common property.",
												MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"taints": schema.ListAttribute{
												Description:         "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",
												MarkdownDescription: "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",
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

									"tracing": schema.SingleNestedAttribute{
										Description:         "Deprecated: for backward compatibility.",
										MarkdownDescription: "Deprecated: for backward compatibility.",
										Attributes: map[string]schema.Attribute{
											"configuration": schema.MapAttribute{
												Description:         "TraitConfiguration parameters configuration",
												MarkdownDescription: "TraitConfiguration parameters configuration",
												ElementType:         types.StringType,
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

					"replicas": schema.Int64Attribute{
						Description:         "Replicas is the number of desired replicas for the binding",
						MarkdownDescription: "Replicas is the number of desired replicas for the binding",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "Custom SA to use for the binding",
						MarkdownDescription: "Custom SA to use for the binding",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sink": schema.SingleNestedAttribute{
						Description:         "Sink is the destination of the integration defined by this binding",
						MarkdownDescription: "Sink is the destination of the integration defined by this binding",
						Attributes: map[string]schema.Attribute{
							"data_types": schema.SingleNestedAttribute{
								Description:         "DataTypes defines the data type of the data produced/consumed by the endpoint and references a given data type specification.",
								MarkdownDescription: "DataTypes defines the data type of the data produced/consumed by the endpoint and references a given data type specification.",
								Attributes: map[string]schema.Attribute{
									"format": schema.StringAttribute{
										Description:         "the data type format name",
										MarkdownDescription: "the data type format name",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "the data type component scheme",
										MarkdownDescription: "the data type component scheme",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"properties": schema.MapAttribute{
								Description:         "Properties are a key value representation of endpoint properties",
								MarkdownDescription: "Properties are a key value representation of endpoint properties",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ref": schema.SingleNestedAttribute{
								Description:         "Ref can be used to declare a Kubernetes resource as source/sink endpoint",
								MarkdownDescription: "Ref can be used to declare a Kubernetes resource as source/sink endpoint",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"field_path": schema.StringAttribute{
										Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"types": schema.SingleNestedAttribute{
								Description:         "Types defines the data type of the data produced/consumed by the endpoint and references a given data type specification. Deprecated: In favor of using DataTypes",
								MarkdownDescription: "Types defines the data type of the data produced/consumed by the endpoint and references a given data type specification. Deprecated: In favor of using DataTypes",
								Attributes: map[string]schema.Attribute{
									"media_type": schema.StringAttribute{
										Description:         "media type as expected for HTTP media types (ie, application/json)",
										MarkdownDescription: "media type as expected for HTTP media types (ie, application/json)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"schema": schema.SingleNestedAttribute{
										Description:         "the expected schema for the event",
										MarkdownDescription: "the expected schema for the event",
										Attributes: map[string]schema.Attribute{
											"dollarschema": schema.StringAttribute{
												Description:         "JSONSchemaURL represents a schema url.",
												MarkdownDescription: "JSONSchemaURL represents a schema url.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"description": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"example": schema.MapAttribute{
												Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
												MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_docs": schema.SingleNestedAttribute{
												Description:         "ExternalDocumentation allows referencing an external resource for extended documentation.",
												MarkdownDescription: "ExternalDocumentation allows referencing an external resource for extended documentation.",
												Attributes: map[string]schema.Attribute{
													"description": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
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

											"id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"properties": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"default": schema.MapAttribute{
														Description:         "default is a default value for undefined object fields.",
														MarkdownDescription: "default is a default value for undefined object fields.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"deprecated": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"description": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enum": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"example": schema.MapAttribute{
														Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
														MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"exclusive_maximum": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"exclusive_minimum": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"format": schema.StringAttribute{
														Description:         "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated: - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
														MarkdownDescription: "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated: - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_items": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_length": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_properties": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"maximum": schema.StringAttribute{
														Description:         "A Number represents a JSON number literal.",
														MarkdownDescription: "A Number represents a JSON number literal.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_items": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_length": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_properties": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"minimum": schema.StringAttribute{
														Description:         "A Number represents a JSON number literal.",
														MarkdownDescription: "A Number represents a JSON number literal.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"multiple_of": schema.StringAttribute{
														Description:         "A Number represents a JSON number literal.",
														MarkdownDescription: "A Number represents a JSON number literal.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"nullable": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"title": schema.StringAttribute{
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

													"unique_items": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"x_descriptors": schema.ListAttribute{
														Description:         "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
														MarkdownDescription: "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
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

											"required": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"title": schema.StringAttribute{
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

							"uri": schema.StringAttribute{
								Description:         "URI can be used to specify the (Camel) endpoint explicitly",
								MarkdownDescription: "URI can be used to specify the (Camel) endpoint explicitly",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"source": schema.SingleNestedAttribute{
						Description:         "Source is the starting point of the integration defined by this binding",
						MarkdownDescription: "Source is the starting point of the integration defined by this binding",
						Attributes: map[string]schema.Attribute{
							"data_types": schema.SingleNestedAttribute{
								Description:         "DataTypes defines the data type of the data produced/consumed by the endpoint and references a given data type specification.",
								MarkdownDescription: "DataTypes defines the data type of the data produced/consumed by the endpoint and references a given data type specification.",
								Attributes: map[string]schema.Attribute{
									"format": schema.StringAttribute{
										Description:         "the data type format name",
										MarkdownDescription: "the data type format name",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "the data type component scheme",
										MarkdownDescription: "the data type component scheme",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"properties": schema.MapAttribute{
								Description:         "Properties are a key value representation of endpoint properties",
								MarkdownDescription: "Properties are a key value representation of endpoint properties",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ref": schema.SingleNestedAttribute{
								Description:         "Ref can be used to declare a Kubernetes resource as source/sink endpoint",
								MarkdownDescription: "Ref can be used to declare a Kubernetes resource as source/sink endpoint",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"field_path": schema.StringAttribute{
										Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"types": schema.SingleNestedAttribute{
								Description:         "Types defines the data type of the data produced/consumed by the endpoint and references a given data type specification. Deprecated: In favor of using DataTypes",
								MarkdownDescription: "Types defines the data type of the data produced/consumed by the endpoint and references a given data type specification. Deprecated: In favor of using DataTypes",
								Attributes: map[string]schema.Attribute{
									"media_type": schema.StringAttribute{
										Description:         "media type as expected for HTTP media types (ie, application/json)",
										MarkdownDescription: "media type as expected for HTTP media types (ie, application/json)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"schema": schema.SingleNestedAttribute{
										Description:         "the expected schema for the event",
										MarkdownDescription: "the expected schema for the event",
										Attributes: map[string]schema.Attribute{
											"dollarschema": schema.StringAttribute{
												Description:         "JSONSchemaURL represents a schema url.",
												MarkdownDescription: "JSONSchemaURL represents a schema url.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"description": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"example": schema.MapAttribute{
												Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
												MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"external_docs": schema.SingleNestedAttribute{
												Description:         "ExternalDocumentation allows referencing an external resource for extended documentation.",
												MarkdownDescription: "ExternalDocumentation allows referencing an external resource for extended documentation.",
												Attributes: map[string]schema.Attribute{
													"description": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"url": schema.StringAttribute{
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

											"id": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"properties": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"default": schema.MapAttribute{
														Description:         "default is a default value for undefined object fields.",
														MarkdownDescription: "default is a default value for undefined object fields.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"deprecated": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"description": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"enum": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"example": schema.MapAttribute{
														Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
														MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"exclusive_maximum": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"exclusive_minimum": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"format": schema.StringAttribute{
														Description:         "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated: - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
														MarkdownDescription: "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated: - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"id": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_items": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_length": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_properties": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"maximum": schema.StringAttribute{
														Description:         "A Number represents a JSON number literal.",
														MarkdownDescription: "A Number represents a JSON number literal.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_items": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_length": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"min_properties": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"minimum": schema.StringAttribute{
														Description:         "A Number represents a JSON number literal.",
														MarkdownDescription: "A Number represents a JSON number literal.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"multiple_of": schema.StringAttribute{
														Description:         "A Number represents a JSON number literal.",
														MarkdownDescription: "A Number represents a JSON number literal.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"nullable": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pattern": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"title": schema.StringAttribute{
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

													"unique_items": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"x_descriptors": schema.ListAttribute{
														Description:         "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
														MarkdownDescription: "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
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

											"required": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"title": schema.StringAttribute{
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

							"uri": schema.StringAttribute{
								Description:         "URI can be used to specify the (Camel) endpoint explicitly",
								MarkdownDescription: "URI can be used to specify the (Camel) endpoint explicitly",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"steps": schema.ListNestedAttribute{
						Description:         "Steps contains an optional list of intermediate steps that are executed between the Source and the Sink",
						MarkdownDescription: "Steps contains an optional list of intermediate steps that are executed between the Source and the Sink",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"data_types": schema.SingleNestedAttribute{
									Description:         "DataTypes defines the data type of the data produced/consumed by the endpoint and references a given data type specification.",
									MarkdownDescription: "DataTypes defines the data type of the data produced/consumed by the endpoint and references a given data type specification.",
									Attributes: map[string]schema.Attribute{
										"format": schema.StringAttribute{
											Description:         "the data type format name",
											MarkdownDescription: "the data type format name",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scheme": schema.StringAttribute{
											Description:         "the data type component scheme",
											MarkdownDescription: "the data type component scheme",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"properties": schema.MapAttribute{
									Description:         "Properties are a key value representation of endpoint properties",
									MarkdownDescription: "Properties are a key value representation of endpoint properties",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"ref": schema.SingleNestedAttribute{
									Description:         "Ref can be used to declare a Kubernetes resource as source/sink endpoint",
									MarkdownDescription: "Ref can be used to declare a Kubernetes resource as source/sink endpoint",
									Attributes: map[string]schema.Attribute{
										"api_version": schema.StringAttribute{
											Description:         "API version of the referent.",
											MarkdownDescription: "API version of the referent.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"field_path": schema.StringAttribute{
											Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
											MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"types": schema.SingleNestedAttribute{
									Description:         "Types defines the data type of the data produced/consumed by the endpoint and references a given data type specification. Deprecated: In favor of using DataTypes",
									MarkdownDescription: "Types defines the data type of the data produced/consumed by the endpoint and references a given data type specification. Deprecated: In favor of using DataTypes",
									Attributes: map[string]schema.Attribute{
										"media_type": schema.StringAttribute{
											Description:         "media type as expected for HTTP media types (ie, application/json)",
											MarkdownDescription: "media type as expected for HTTP media types (ie, application/json)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"schema": schema.SingleNestedAttribute{
											Description:         "the expected schema for the event",
											MarkdownDescription: "the expected schema for the event",
											Attributes: map[string]schema.Attribute{
												"dollarschema": schema.StringAttribute{
													Description:         "JSONSchemaURL represents a schema url.",
													MarkdownDescription: "JSONSchemaURL represents a schema url.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"description": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"example": schema.MapAttribute{
													Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
													MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"external_docs": schema.SingleNestedAttribute{
													Description:         "ExternalDocumentation allows referencing an external resource for extended documentation.",
													MarkdownDescription: "ExternalDocumentation allows referencing an external resource for extended documentation.",
													Attributes: map[string]schema.Attribute{
														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"url": schema.StringAttribute{
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

												"id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"properties": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"default": schema.MapAttribute{
															Description:         "default is a default value for undefined object fields.",
															MarkdownDescription: "default is a default value for undefined object fields.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"deprecated": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"enum": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"example": schema.MapAttribute{
															Description:         "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
															MarkdownDescription: "JSON represents any valid JSON value. These types are supported: bool, int64, float64, string, []interface{}, map[string]interface{} and nil.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"exclusive_maximum": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"exclusive_minimum": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"format": schema.StringAttribute{
															Description:         "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated: - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
															MarkdownDescription: "format is an OpenAPI v3 format string. Unknown formats are ignored. The following formats are validated: - bsonobjectid: a bson object ID, i.e. a 24 characters hex string - uri: an URI as parsed by Golang net/url.ParseRequestURI - email: an email address as parsed by Golang net/mail.ParseAddress - hostname: a valid representation for an Internet host name, as defined by RFC 1034, section 3.1 [RFC1034]. - ipv4: an IPv4 IP as parsed by Golang net.ParseIP - ipv6: an IPv6 IP as parsed by Golang net.ParseIP - cidr: a CIDR as parsed by Golang net.ParseCIDR - mac: a MAC address as parsed by Golang net.ParseMAC - uuid: an UUID that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid3: an UUID3 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?3[0-9a-f]{3}-?[0-9a-f]{4}-?[0-9a-f]{12}$ - uuid4: an UUID4 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?4[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - uuid5: an UUID5 that allows uppercase defined by the regex (?i)^[0-9a-f]{8}-?[0-9a-f]{4}-?5[0-9a-f]{3}-?[89ab][0-9a-f]{3}-?[0-9a-f]{12}$ - isbn: an ISBN10 or ISBN13 number string like '0321751043' or '978-0321751041' - isbn10: an ISBN10 number string like '0321751043' - isbn13: an ISBN13 number string like '978-0321751041' - creditcard: a credit card number defined by the regex ^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35d{3})d{11})$ with any non digit characters mixed in - ssn: a U.S. social security number following the regex ^d{3}[- ]?d{2}[- ]?d{4}$ - hexcolor: an hexadecimal color code like '#FFFFFF' following the regex ^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$ - rgbcolor: an RGB color code like rgb like 'rgb(255,255,255)' - byte: base64 encoded binary data - password: any kind of string - date: a date string like '2006-01-02' as defined by full-date in RFC3339 - duration: a duration string like '22 ns' as parsed by Golang time.ParseDuration or compatible with Scala duration format - datetime: a date time string like '2014-12-15T19:30:20.000Z' as defined by date-time in RFC3339.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_items": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_length": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"max_properties": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"maximum": schema.StringAttribute{
															Description:         "A Number represents a JSON number literal.",
															MarkdownDescription: "A Number represents a JSON number literal.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"min_items": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"min_length": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"min_properties": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"minimum": schema.StringAttribute{
															Description:         "A Number represents a JSON number literal.",
															MarkdownDescription: "A Number represents a JSON number literal.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"multiple_of": schema.StringAttribute{
															Description:         "A Number represents a JSON number literal.",
															MarkdownDescription: "A Number represents a JSON number literal.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"nullable": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pattern": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"title": schema.StringAttribute{
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

														"unique_items": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"x_descriptors": schema.ListAttribute{
															Description:         "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
															MarkdownDescription: "XDescriptors is a list of extended properties that trigger a custom behavior in external systems",
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

												"required": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"title": schema.StringAttribute{
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

								"uri": schema.StringAttribute{
									Description:         "URI can be used to specify the (Camel) endpoint explicitly",
									MarkdownDescription: "URI can be used to specify the (Camel) endpoint explicitly",
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

func (r *CamelApacheOrgKameletBindingV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_camel_apache_org_kamelet_binding_v1alpha1_manifest")

	var model CamelApacheOrgKameletBindingV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("camel.apache.org/v1alpha1")
	model.Kind = pointer.String("KameletBinding")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
