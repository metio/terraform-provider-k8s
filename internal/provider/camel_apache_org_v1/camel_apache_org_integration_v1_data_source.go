/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package camel_apache_org_v1

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
	_ datasource.DataSource              = &CamelApacheOrgIntegrationV1DataSource{}
	_ datasource.DataSourceWithConfigure = &CamelApacheOrgIntegrationV1DataSource{}
)

func NewCamelApacheOrgIntegrationV1DataSource() datasource.DataSource {
	return &CamelApacheOrgIntegrationV1DataSource{}
}

type CamelApacheOrgIntegrationV1DataSource struct {
	kubernetesClient dynamic.Interface
}

type CamelApacheOrgIntegrationV1DataSourceData struct {
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
				ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" json:"activeDeadlineSeconds,omitempty"`
				Containers            *[]struct {
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
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				LimitCPU      *string            `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
				LimitMemory   *string            `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
				MavenProfiles *[]string          `tfsdk:"maven_profiles" json:"mavenProfiles,omitempty"`
				OrderStrategy *string            `tfsdk:"order_strategy" json:"orderStrategy,omitempty"`
				Properties    *[]string          `tfsdk:"properties" json:"properties,omitempty"`
				RequestCPU    *string            `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
				RequestMemory *string            `tfsdk:"request_memory" json:"requestMemory,omitempty"`
				Strategy      *string            `tfsdk:"strategy" json:"strategy,omitempty"`
				Tasks         *[]string          `tfsdk:"tasks" json:"tasks,omitempty"`
				Verbose       *bool              `tfsdk:"verbose" json:"verbose,omitempty"`
			} `tfsdk:"builder" json:"builder,omitempty"`
			Camel *struct {
				Configuration  *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled        *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Properties     *[]string          `tfsdk:"properties" json:"properties,omitempty"`
				RuntimeVersion *string            `tfsdk:"runtime_version" json:"runtimeVersion,omitempty"`
			} `tfsdk:"camel" json:"camel,omitempty"`
			Container *struct {
				Auto            *bool              `tfsdk:"auto" json:"auto,omitempty"`
				Configuration   *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled         *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Expose          *bool              `tfsdk:"expose" json:"expose,omitempty"`
				Image           *string            `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy *string            `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				LimitCPU        *string            `tfsdk:"limit_cpu" json:"limitCPU,omitempty"`
				LimitMemory     *string            `tfsdk:"limit_memory" json:"limitMemory,omitempty"`
				Name            *string            `tfsdk:"name" json:"name,omitempty"`
				Port            *int64             `tfsdk:"port" json:"port,omitempty"`
				PortName        *string            `tfsdk:"port_name" json:"portName,omitempty"`
				RequestCPU      *string            `tfsdk:"request_cpu" json:"requestCPU,omitempty"`
				RequestMemory   *string            `tfsdk:"request_memory" json:"requestMemory,omitempty"`
				ServicePort     *int64             `tfsdk:"service_port" json:"servicePort,omitempty"`
				ServicePortName *string            `tfsdk:"service_port_name" json:"servicePortName,omitempty"`
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
				RollingUpdateMaxSurge       *int64             `tfsdk:"rolling_update_max_surge" json:"rollingUpdateMaxSurge,omitempty"`
				RollingUpdateMaxUnavailable *int64             `tfsdk:"rolling_update_max_unavailable" json:"rollingUpdateMaxUnavailable,omitempty"`
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
				LivenessProbeEnabled      *bool              `tfsdk:"liveness_probe_enabled" json:"livenessProbeEnabled,omitempty"`
				LivenessScheme            *string            `tfsdk:"liveness_scheme" json:"livenessScheme,omitempty"`
				LivenessSuccessThreshold  *int64             `tfsdk:"liveness_success_threshold" json:"livenessSuccessThreshold,omitempty"`
				LivenessTimeout           *int64             `tfsdk:"liveness_timeout" json:"livenessTimeout,omitempty"`
				ReadinessFailureThreshold *int64             `tfsdk:"readiness_failure_threshold" json:"readinessFailureThreshold,omitempty"`
				ReadinessInitialDelay     *int64             `tfsdk:"readiness_initial_delay" json:"readinessInitialDelay,omitempty"`
				ReadinessPeriod           *int64             `tfsdk:"readiness_period" json:"readinessPeriod,omitempty"`
				ReadinessProbeEnabled     *bool              `tfsdk:"readiness_probe_enabled" json:"readinessProbeEnabled,omitempty"`
				ReadinessScheme           *string            `tfsdk:"readiness_scheme" json:"readinessScheme,omitempty"`
				ReadinessSuccessThreshold *int64             `tfsdk:"readiness_success_threshold" json:"readinessSuccessThreshold,omitempty"`
				ReadinessTimeout          *int64             `tfsdk:"readiness_timeout" json:"readinessTimeout,omitempty"`
				StartupFailureThreshold   *int64             `tfsdk:"startup_failure_threshold" json:"startupFailureThreshold,omitempty"`
				StartupInitialDelay       *int64             `tfsdk:"startup_initial_delay" json:"startupInitialDelay,omitempty"`
				StartupPeriod             *int64             `tfsdk:"startup_period" json:"startupPeriod,omitempty"`
				StartupProbeEnabled       *bool              `tfsdk:"startup_probe_enabled" json:"startupProbeEnabled,omitempty"`
				StartupScheme             *string            `tfsdk:"startup_scheme" json:"startupScheme,omitempty"`
				StartupSuccessThreshold   *int64             `tfsdk:"startup_success_threshold" json:"startupSuccessThreshold,omitempty"`
				StartupTimeout            *int64             `tfsdk:"startup_timeout" json:"startupTimeout,omitempty"`
			} `tfsdk:"health" json:"health,omitempty"`
			Ingress *struct {
				Annotations   *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Auto          *bool              `tfsdk:"auto" json:"auto,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Host          *string            `tfsdk:"host" json:"host,omitempty"`
				Path          *string            `tfsdk:"path" json:"path,omitempty"`
				PathType      *string            `tfsdk:"path_type" json:"pathType,omitempty"`
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
				Options       *[]string          `tfsdk:"options" json:"options,omitempty"`
				PrintCommand  *bool              `tfsdk:"print_command" json:"printCommand,omitempty"`
			} `tfsdk:"jvm" json:"jvm,omitempty"`
			Kamelets *struct {
				Auto          *bool              `tfsdk:"auto" json:"auto,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				List          *string            `tfsdk:"list" json:"list,omitempty"`
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
				FilterSourceChannels *bool              `tfsdk:"filter_source_channels" json:"filterSourceChannels,omitempty"`
				NamespaceLabel       *bool              `tfsdk:"namespace_label" json:"namespaceLabel,omitempty"`
				SinkBinding          *bool              `tfsdk:"sink_binding" json:"sinkBinding,omitempty"`
			} `tfsdk:"knative" json:"knative,omitempty"`
			Knative_service *struct {
				Auto              *bool              `tfsdk:"auto" json:"auto,omitempty"`
				AutoscalingMetric *string            `tfsdk:"autoscaling_metric" json:"autoscalingMetric,omitempty"`
				AutoscalingTarget *int64             `tfsdk:"autoscaling_target" json:"autoscalingTarget,omitempty"`
				Class             *string            `tfsdk:"class" json:"class,omitempty"`
				Configuration     *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled           *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				MaxScale          *int64             `tfsdk:"max_scale" json:"maxScale,omitempty"`
				MinScale          *int64             `tfsdk:"min_scale" json:"minScale,omitempty"`
				RolloutDuration   *string            `tfsdk:"rollout_duration" json:"rolloutDuration,omitempty"`
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
				Configs       *[]string          `tfsdk:"configs" json:"configs,omitempty"`
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				HotReload     *bool              `tfsdk:"hot_reload" json:"hotReload,omitempty"`
				Resources     *[]string          `tfsdk:"resources" json:"resources,omitempty"`
				Volumes       *[]string          `tfsdk:"volumes" json:"volumes,omitempty"`
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
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				PackageTypes  *[]string          `tfsdk:"package_types" json:"packageTypes,omitempty"`
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
			Toleration *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
				Enabled       *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Taints        *[]string          `tfsdk:"taints" json:"taints,omitempty"`
			} `tfsdk:"toleration" json:"toleration,omitempty"`
			Tracing *struct {
				Configuration *map[string]string `tfsdk:"configuration" json:"configuration,omitempty"`
			} `tfsdk:"tracing" json:"tracing,omitempty"`
		} `tfsdk:"traits" json:"traits,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *CamelApacheOrgIntegrationV1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_camel_apache_org_integration_v1"
}

func (r *CamelApacheOrgIntegrationV1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Integration is the Schema for the integrations API.",
		MarkdownDescription: "Integration is the Schema for the integrations API.",
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
				Description:         "the desired Integration specification",
				MarkdownDescription: "the desired Integration specification",
				Attributes: map[string]schema.Attribute{
					"configuration": schema.ListNestedAttribute{
						Description:         "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes",
						MarkdownDescription: "Deprecated: Use camel trait (camel.properties) to manage properties Use mount trait (mount.configs) to manage configs Use mount trait (mount.resources) to manage resources Use mount trait (mount.volumes) to manage volumes",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"type": schema.StringAttribute{
									Description:         "represents the type of configuration, ie: property, configmap, secret, ...",
									MarkdownDescription: "represents the type of configuration, ie: property, configmap, secret, ...",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"value": schema.StringAttribute{
									Description:         "the value to assign to the configuration (syntax may vary depending on the 'Type')",
									MarkdownDescription: "the value to assign to the configuration (syntax may vary depending on the 'Type')",
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

					"dependencies": schema.ListAttribute{
						Description:         "the list of Camel or Maven dependencies required by the Integration",
						MarkdownDescription: "the list of Camel or Maven dependencies required by the Integration",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"flows": schema.ListAttribute{
						Description:         "a source in YAML DSL language which contain the routes to run",
						MarkdownDescription: "a source in YAML DSL language which contain the routes to run",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"integration_kit": schema.SingleNestedAttribute{
						Description:         "the reference of the 'IntegrationKit' which is used for this Integration",
						MarkdownDescription: "the reference of the 'IntegrationKit' which is used for this Integration",
						Attributes: map[string]schema.Attribute{
							"api_version": schema.StringAttribute{
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"field_path": schema.StringAttribute{
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object. TODO: this design is not final and this field is subject to change in the future.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"name": schema.StringAttribute{
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"namespace": schema.StringAttribute{
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resource_version": schema.StringAttribute{
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"uid": schema.StringAttribute{
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"profile": schema.StringAttribute{
						Description:         "the profile needed to run this Integration",
						MarkdownDescription: "the profile needed to run this Integration",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replicas": schema.Int64Attribute{
						Description:         "the number of 'Pods' needed for the running Integration",
						MarkdownDescription: "the number of 'Pods' needed for the running Integration",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"repositories": schema.ListAttribute{
						Description:         "additional Maven repositories to be used",
						MarkdownDescription: "additional Maven repositories to be used",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"service_account_name": schema.StringAttribute{
						Description:         "custom SA to use for the Integration",
						MarkdownDescription: "custom SA to use for the Integration",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
									Optional:            false,
									Computed:            true,
								},

								"content": schema.StringAttribute{
									Description:         "the source code (plain text)",
									MarkdownDescription: "the source code (plain text)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_key": schema.StringAttribute{
									Description:         "the confimap key holding the source content",
									MarkdownDescription: "the confimap key holding the source content",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_ref": schema.StringAttribute{
									Description:         "the confimap reference holding the source content",
									MarkdownDescription: "the confimap reference holding the source content",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"content_type": schema.StringAttribute{
									Description:         "the content type (tipically text or binary)",
									MarkdownDescription: "the content type (tipically text or binary)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"interceptors": schema.ListAttribute{
									Description:         "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
									MarkdownDescription: "Interceptors are optional identifiers the org.apache.camel.k.RoutesLoader uses to pre/post process sources",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"language": schema.StringAttribute{
									Description:         "specify which is the language (Camel DSL) used to interpret this source code",
									MarkdownDescription: "specify which is the language (Camel DSL) used to interpret this source code",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"loader": schema.StringAttribute{
									Description:         "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
									MarkdownDescription: "Loader is an optional id of the org.apache.camel.k.RoutesLoader that will interpret this source at runtime",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"name": schema.StringAttribute{
									Description:         "the name of the specification",
									MarkdownDescription: "the name of the specification",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"path": schema.StringAttribute{
									Description:         "the path where the file is stored",
									MarkdownDescription: "the path where the file is stored",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"property_names": schema.ListAttribute{
									Description:         "List of property names defined in the source (e.g. if type is 'template')",
									MarkdownDescription: "List of property names defined in the source (e.g. if type is 'template')",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"raw_content": schema.StringAttribute{
									Description:         "the source code (binary)",
									MarkdownDescription: "the source code (binary)",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"type": schema.StringAttribute{
									Description:         "Type defines the kind of source described by this object",
									MarkdownDescription: "Type defines the kind of source described by this object",
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
										Optional:            false,
										Computed:            true,
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
															Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
															MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
															Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

									"dns_policy": schema.StringAttribute{
										Description:         "DNSPolicy",
										MarkdownDescription: "DNSPolicy",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
													Optional:            false,
													Computed:            true,
												},

												"command": schema.ListAttribute{
													Description:         "Entrypoint array. Not executed within a shell. The image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
													MarkdownDescription: "Entrypoint array. Not executed within a shell. The image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
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
													Description:         "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images",
													MarkdownDescription: "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images",
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
													Description:         "Probes are not allowed for ephemeral containers.",
													MarkdownDescription: "Probes are not allowed for ephemeral containers.",
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
													Description:         "Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.",
													MarkdownDescription: "Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"ports": schema.ListNestedAttribute{
													Description:         "Ports are not allowed for ephemeral containers.",
													MarkdownDescription: "Ports are not allowed for ephemeral containers.",
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
													Description:         "Probes are not allowed for ephemeral containers.",
													MarkdownDescription: "Probes are not allowed for ephemeral containers.",
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
													Description:         "Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
													MarkdownDescription: "Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
													Attributes: map[string]schema.Attribute{
														"claims": schema.ListNestedAttribute{
															Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
															MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
															Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
													Description:         "Optional: SecurityContext defines the security options the ephemeral container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",
													MarkdownDescription: "Optional: SecurityContext defines the security options the ephemeral container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.",
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
													Description:         "Probes are not allowed for ephemeral containers.",
													MarkdownDescription: "Probes are not allowed for ephemeral containers.",
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

												"target_container_name": schema.StringAttribute{
													Description:         "If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container uses the namespaces configured in the Pod spec.  The container runtime must implement support for this feature. If the runtime does not support namespace targeting then the result of setting this field is undefined.",
													MarkdownDescription: "If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container uses the namespaces configured in the Pod spec.  The container runtime must implement support for this feature. If the runtime does not support namespace targeting then the result of setting this field is undefined.",
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
													Description:         "Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. Cannot be updated.",
													MarkdownDescription: "Pod volumes to mount into the container's filesystem. Subpath mounts are not allowed for ephemeral containers. Cannot be updated.",
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
															Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
															MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
															Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector",
										MarkdownDescription: "NodeSelector",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"restart_policy": schema.StringAttribute{
										Description:         "RestartPolicy",
										MarkdownDescription: "RestartPolicy",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"security_context": schema.SingleNestedAttribute{
										Description:         "PodSecurityContext",
										MarkdownDescription: "PodSecurityContext",
										Attributes: map[string]schema.Attribute{
											"fs_group": schema.Int64Attribute{
												Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume. Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"fs_group_change_policy": schema.StringAttribute{
												Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used. Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "Indicates that the container must run as a non-root user. If true, the Kubelet will validate the image at runtime to ensure that it does not run as UID 0 (root) and fail to start the container if it does. If unset or false, no such validation will be performed. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"se_linux_options": schema.SingleNestedAttribute{
												Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container. Note that this field cannot be set when spec.os.name is windows.",
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

											"supplemental_groups": schema.ListAttribute{
												Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID, the fsGroup (if specified), and group memberships defined in the container image for the uid of the container process. If unspecified, no additional groups are added to any container. Note that group memberships defined in the container image for the uid of the container process are still effective, even if they are not included in this list. Note that this field cannot be set when spec.os.name is windows.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sysctls": schema.ListNestedAttribute{
												Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch. Note that this field cannot be set when spec.os.name is windows.",
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
												Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
												MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence. Note that this field cannot be set when spec.os.name is linux.",
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

									"termination_grace_period_seconds": schema.Int64Attribute{
										Description:         "TerminationGracePeriodSeconds",
										MarkdownDescription: "TerminationGracePeriodSeconds",
										Required:            false,
										Optional:            false,
										Computed:            true,
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

												"match_label_keys": schema.ListAttribute{
													Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
													MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_skew": schema.Int64Attribute{
													Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
													MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min_domains": schema.Int64Attribute{
													Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
													MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_affinity_policy": schema.StringAttribute{
													Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
													MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_taints_policy": schema.StringAttribute{
													Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
													MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"topology_key": schema.StringAttribute{
													Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
													MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"when_unsatisfiable": schema.StringAttribute{
													Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location,   but giving higher precedence to topologies that would help reduce the   skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
													MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location,   but giving higher precedence to topologies that would help reduce the   skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
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
															Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
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
													Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
													MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
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
																			Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																			MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
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
																			Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. * While dataSource only allows local objects, dataSourceRef allows objects   in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																			MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. * While dataSource only allows local objects, dataSourceRef allows objects   in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
																					Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																					MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
																				"claims": schema.ListNestedAttribute{
																					Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																					MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
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
																					Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																					MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"addons": schema.MapAttribute{
								Description:         "The extension point with addon traits",
								MarkdownDescription: "The extension point with addon traits",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"node_affinity_labels": schema.ListAttribute{
										Description:         "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",
										MarkdownDescription: "Defines a set of nodes the integration pod(s) are eligible to be scheduled on, based on labels on the node.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_affinity": schema.BoolAttribute{
										Description:         "Always co-locates multiple replicas of the integration in the same node (default *false*).",
										MarkdownDescription: "Always co-locates multiple replicas of the integration in the same node (default *false*).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_affinity_labels": schema.ListAttribute{
										Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",
										MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should be co-located with.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_anti_affinity": schema.BoolAttribute{
										Description:         "Never co-locates multiple replicas of the integration in the same node (default *false*).",
										MarkdownDescription: "Never co-locates multiple replicas of the integration in the same node (default *false*).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_anti_affinity_labels": schema.ListAttribute{
										Description:         "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",
										MarkdownDescription: "Defines a set of pods (namely those matching the label selector, relative to the given namespace) that the integration pod(s) should not be co-located with.",
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

							"builder": schema.SingleNestedAttribute{
								Description:         "The configuration of Builder trait",
								MarkdownDescription: "The configuration of Builder trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_cpu": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the maximum amount of CPU required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the maximum amount of CPU required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_memory": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the maximum amount of memory required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the maximum amount of memory required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"maven_profiles": schema.ListAttribute{
										Description:         "A list of references pointing to configmaps/secrets that contains a maven profile. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
										MarkdownDescription: "A list of references pointing to configmaps/secrets that contains a maven profile. The content of the maven profile is expected to be a text containing a valid maven profile starting with '<profile>' and ending with '</profile>' that will be integrated as an inline profile in the POM. Syntax: [configmap|secret]:name[/key], where name represents the resource name, key optionally represents the resource key to be filtered (default key value = profile.xml).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"order_strategy": schema.StringAttribute{
										Description:         "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default sequential)",
										MarkdownDescription: "The build order strategy to use, either 'dependencies', 'fifo' or 'sequential' (default sequential)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"properties": schema.ListAttribute{
										Description:         "A list of properties to be provided to the build task",
										MarkdownDescription: "A list of properties to be provided to the build task",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_cpu": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the minimum amount of CPU required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the minimum amount of CPU required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_memory": schema.StringAttribute{
										Description:         "When using 'pod' strategy, the minimum amount of memory required by the pod builder.",
										MarkdownDescription: "When using 'pod' strategy, the minimum amount of memory required by the pod builder.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"strategy": schema.StringAttribute{
										Description:         "The strategy to use, either 'pod' or 'routine' (default routine)",
										MarkdownDescription: "The strategy to use, either 'pod' or 'routine' (default routine)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tasks": schema.ListAttribute{
										Description:         "A list of tasks to be executed (available only when using 'pod' strategy) with format <name>;<container-image>;<container-command>",
										MarkdownDescription: "A list of tasks to be executed (available only when using 'pod' strategy) with format <name>;<container-image>;<container-command>",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"verbose": schema.BoolAttribute{
										Description:         "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",
										MarkdownDescription: "Enable verbose logging on build components that support it (e.g. Kaniko build pod).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"properties": schema.ListAttribute{
										Description:         "A list of properties to be provided to the Integration runtime",
										MarkdownDescription: "A list of properties to be provided to the Integration runtime",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"runtime_version": schema.StringAttribute{
										Description:         "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",
										MarkdownDescription: "The camel-k-runtime version to use for the integration. It overrides the default version set in the Integration Platform.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"container": schema.SingleNestedAttribute{
								Description:         "The configuration of Container trait",
								MarkdownDescription: "The configuration of Container trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "To automatically enable the trait",
										MarkdownDescription: "To automatically enable the trait",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"expose": schema.BoolAttribute{
										Description:         "Can be used to enable/disable exposure via kubernetes Service.",
										MarkdownDescription: "Can be used to enable/disable exposure via kubernetes Service.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"image": schema.StringAttribute{
										Description:         "The main container image",
										MarkdownDescription: "The main container image",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "The pull policy: Always|Never|IfNotPresent",
										MarkdownDescription: "The pull policy: Always|Never|IfNotPresent",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_cpu": schema.StringAttribute{
										Description:         "The maximum amount of CPU required.",
										MarkdownDescription: "The maximum amount of CPU required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"limit_memory": schema.StringAttribute{
										Description:         "The maximum amount of memory required.",
										MarkdownDescription: "The maximum amount of memory required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"name": schema.StringAttribute{
										Description:         "The main container name. It's named 'integration' by default.",
										MarkdownDescription: "The main container name. It's named 'integration' by default.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.Int64Attribute{
										Description:         "To configure a different port exposed by the container (default '8080').",
										MarkdownDescription: "To configure a different port exposed by the container (default '8080').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port_name": schema.StringAttribute{
										Description:         "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",
										MarkdownDescription: "To configure a different port name for the port exposed by the container. It defaults to 'http' only when the 'expose' parameter is true.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_cpu": schema.StringAttribute{
										Description:         "The minimum amount of CPU required.",
										MarkdownDescription: "The minimum amount of CPU required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"request_memory": schema.StringAttribute{
										Description:         "The minimum amount of memory required.",
										MarkdownDescription: "The minimum amount of memory required.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"service_port": schema.Int64Attribute{
										Description:         "To configure under which service port the container port is to be exposed (default '80').",
										MarkdownDescription: "To configure under which service port the container port is to be exposed (default '80').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"service_port_name": schema.StringAttribute{
										Description:         "To configure under which service port name the container port is to be exposed (default 'http').",
										MarkdownDescription: "To configure under which service port name the container port is to be exposed (default 'http').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"cron": schema.SingleNestedAttribute{
								Description:         "The configuration of Cron trait",
								MarkdownDescription: "The configuration of Cron trait",
								Attributes: map[string]schema.Attribute{
									"active_deadline_seconds": schema.Int64Attribute{
										Description:         "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",
										MarkdownDescription: "Specifies the duration in seconds, relative to the start time, that the job may be continuously active before it is considered to be failed. It defaults to 60s.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"auto": schema.BoolAttribute{
										Description:         "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer).  It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",
										MarkdownDescription: "Automatically deploy the integration as CronJob when all routes are either starting from a periodic consumer (only 'cron', 'timer' and 'quartz' are supported) or a passive consumer (e.g. 'direct' is a passive consumer).  It's required that all periodic consumers have the same period, and it can be expressed as cron schedule (e.g. '1m' can be expressed as '0/1 * * * *', while '35m' or '50s' cannot).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"backoff_limit": schema.Int64Attribute{
										Description:         "Specifies the number of retries before marking the job failed. It defaults to 2.",
										MarkdownDescription: "Specifies the number of retries before marking the job failed. It defaults to 2.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"components": schema.StringAttribute{
										Description:         "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library).  Supported components are currently: 'cron', 'timer' and 'quartz'.",
										MarkdownDescription: "A comma separated list of the Camel components that need to be customized in order for them to work when the schedule is triggered externally by Kubernetes. A specific customizer is activated for each specified component. E.g. for the 'timer' component, the 'cron-timer' customizer is activated (it's present in the 'org.apache.camel.k:camel-k-cron' library).  Supported components are currently: 'cron', 'timer' and 'quartz'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"concurrency_policy": schema.StringAttribute{
										Description:         "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",
										MarkdownDescription: "Specifies how to treat concurrent executions of a Job. Valid values are: - 'Allow': allows CronJobs to run concurrently; - 'Forbid' (default): forbids concurrent runs, skipping next run if previous run hasn't finished yet; - 'Replace': cancels currently running job and replaces it with a new one",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"fallback": schema.BoolAttribute{
										Description:         "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",
										MarkdownDescription: "Use the default Camel implementation of the 'cron' endpoint ('quartz') instead of trying to materialize the integration as Kubernetes CronJob.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"schedule": schema.StringAttribute{
										Description:         "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",
										MarkdownDescription: "The CronJob schedule for the whole integration. If multiple routes are declared, they must have the same schedule for this mechanism to work correctly.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"starting_deadline_seconds": schema.Int64Attribute{
										Description:         "Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.",
										MarkdownDescription: "Optional deadline in seconds for starting the job if it misses scheduled time for any reason.  Missed jobs executions will be counted as failed ones.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"kind": schema.StringAttribute{
										Description:         "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",
										MarkdownDescription: "Allows to explicitly select the desired deployment kind between 'deployment', 'cron-job' or 'knative-service' when creating the resources for running the integration.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"use_ssa": schema.BoolAttribute{
										Description:         "Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",
										MarkdownDescription: "Use server-side apply to update the owned resources (default 'true'). Note that it automatically falls back to client-side patching, if SSA is not available, e.g., on old Kubernetes clusters.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"progress_deadline_seconds": schema.Int64Attribute{
										Description:         "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to 60s.",
										MarkdownDescription: "The maximum time in seconds for the deployment to make progress before it is considered to be failed. It defaults to 60s.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"rolling_update_max_surge": schema.Int64Attribute{
										Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%.",
										MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"rolling_update_max_unavailable": schema.Int64Attribute{
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"strategy": schema.StringAttribute{
										Description:         "The deployment strategy to use to replace existing pods with new ones.",
										MarkdownDescription: "The deployment strategy to use to replace existing pods with new ones.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"container_meta": schema.BoolAttribute{
										Description:         "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",
										MarkdownDescription: "Enables injection of 'NAMESPACE' and 'POD_NAME' environment variables (default 'true')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"http_proxy": schema.BoolAttribute{
										Description:         "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",
										MarkdownDescription: "Propagates the 'HTTP_PROXY', 'HTTPS_PROXY' and 'NO_PROXY' environment variables (default 'true')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"vars": schema.ListAttribute{
										Description:         "A list of environment variables to be added to the integration container. The syntax is KEY=VALUE, e.g., 'MY_VAR='my value''. These take precedence over the previously defined environment variables.",
										MarkdownDescription: "A list of environment variables to be added to the integration container. The syntax is KEY=VALUE, e.g., 'MY_VAR='my value''. These take precedence over the previously defined environment variables.",
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

							"error_handler": schema.SingleNestedAttribute{
								Description:         "The configuration of Error Handler trait",
								MarkdownDescription: "The configuration of Error Handler trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ref": schema.StringAttribute{
										Description:         "The error handler ref name provided or found in application properties",
										MarkdownDescription: "The error handler ref name provided or found in application properties",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"discovery_cache": schema.StringAttribute{
										Description:         "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",
										MarkdownDescription: "Discovery client cache to be used, either 'disabled', 'disk' or 'memory' (default 'memory'). Deprecated: to be removed from trait configuration.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the liveness probe to be considered failed after having succeeded.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_initial_delay": schema.Int64Attribute{
										Description:         "Number of seconds after the container has started before the liveness probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the liveness probe is initiated.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_period": schema.Int64Attribute{
										Description:         "How often to perform the liveness probe.",
										MarkdownDescription: "How often to perform the liveness probe.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_probe_enabled": schema.BoolAttribute{
										Description:         "Configures the liveness probe for the integration container (default 'false').",
										MarkdownDescription: "Configures the liveness probe for the integration container (default 'false').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_scheme": schema.StringAttribute{
										Description:         "Scheme to use when connecting to the liveness probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the liveness probe (default 'HTTP').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the liveness probe to be considered successful after having failed.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"liveness_timeout": schema.Int64Attribute{
										Description:         "Number of seconds after which the liveness probe times out.",
										MarkdownDescription: "Number of seconds after which the liveness probe times out.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the readiness probe to be considered failed after having succeeded.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_initial_delay": schema.Int64Attribute{
										Description:         "Number of seconds after the container has started before the readiness probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the readiness probe is initiated.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_period": schema.Int64Attribute{
										Description:         "How often to perform the readiness probe.",
										MarkdownDescription: "How often to perform the readiness probe.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_probe_enabled": schema.BoolAttribute{
										Description:         "Configures the readiness probe for the integration container (default 'true').",
										MarkdownDescription: "Configures the readiness probe for the integration container (default 'true').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_scheme": schema.StringAttribute{
										Description:         "Scheme to use when connecting to the readiness probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the readiness probe (default 'HTTP').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the readiness probe to be considered successful after having failed.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"readiness_timeout": schema.Int64Attribute{
										Description:         "Number of seconds after which the readiness probe times out.",
										MarkdownDescription: "Number of seconds after which the readiness probe times out.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the startup probe to be considered failed after having succeeded.",
										MarkdownDescription: "Minimum consecutive failures for the startup probe to be considered failed after having succeeded.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_initial_delay": schema.Int64Attribute{
										Description:         "Number of seconds after the container has started before the startup probe is initiated.",
										MarkdownDescription: "Number of seconds after the container has started before the startup probe is initiated.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_period": schema.Int64Attribute{
										Description:         "How often to perform the startup probe.",
										MarkdownDescription: "How often to perform the startup probe.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_probe_enabled": schema.BoolAttribute{
										Description:         "Configures the startup probe for the integration container (default 'false').",
										MarkdownDescription: "Configures the startup probe for the integration container (default 'false').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_scheme": schema.StringAttribute{
										Description:         "Scheme to use when connecting to the startup probe (default 'HTTP').",
										MarkdownDescription: "Scheme to use when connecting to the startup probe (default 'HTTP').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the startup probe to be considered successful after having failed.",
										MarkdownDescription: "Minimum consecutive successes for the startup probe to be considered successful after having failed.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"startup_timeout": schema.Int64Attribute{
										Description:         "Number of seconds after which the startup probe times out.",
										MarkdownDescription: "Number of seconds after which the startup probe times out.",
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
								Description:         "The configuration of Ingress trait",
								MarkdownDescription: "The configuration of Ingress trait",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "The annotations added to the ingress. This can be used to set controller specific annotations, e.g., when using the NGINX Ingress controller: See https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md",
										MarkdownDescription: "The annotations added to the ingress. This can be used to set controller specific annotations, e.g., when using the NGINX Ingress controller: See https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"auto": schema.BoolAttribute{
										Description:         "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",
										MarkdownDescription: "To automatically add an ingress whenever the integration uses an HTTP endpoint consumer.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"host": schema.StringAttribute{
										Description:         "To configure the host exposed by the ingress.",
										MarkdownDescription: "To configure the host exposed by the ingress.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"path": schema.StringAttribute{
										Description:         "To configure the path exposed by the ingress (default '/').",
										MarkdownDescription: "To configure the path exposed by the ingress (default '/').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"path_type": schema.StringAttribute{
										Description:         "To configure the path type exposed by the ingress. One of 'Exact', 'Prefix', 'ImplementationSpecific' (default to 'Prefix').",
										MarkdownDescription: "To configure the path type exposed by the ingress. One of 'Exact', 'Prefix', 'ImplementationSpecific' (default to 'Prefix').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"istio": schema.SingleNestedAttribute{
								Description:         "The configuration of Istio trait",
								MarkdownDescription: "The configuration of Istio trait",
								Attributes: map[string]schema.Attribute{
									"allow": schema.StringAttribute{
										Description:         "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",
										MarkdownDescription: "Configures a (comma-separated) list of CIDR subnets that should not be intercepted by the Istio proxy ('10.0.0.0/8,172.16.0.0/12,192.168.0.0/16' by default).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"inject": schema.BoolAttribute{
										Description:         "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",
										MarkdownDescription: "Forces the value for labels 'sidecar.istio.io/inject'. By default the label is set to 'true' on deployment and not set on Knative Service.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"jolokia": schema.SingleNestedAttribute{
								Description:         "The configuration of Jolokia trait",
								MarkdownDescription: "The configuration of Jolokia trait",
								Attributes: map[string]schema.Attribute{
									"ca_cert": schema.StringAttribute{
										Description:         "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",
										MarkdownDescription: "The PEM encoded CA certification file path, used to verify client certificates, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default '/var/run/secrets/kubernetes.io/serviceaccount/service-ca.crt' for OpenShift).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"client_principal": schema.ListAttribute{
										Description:         "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",
										MarkdownDescription: "The principal(s) which must be given in a client certificate to allow access to the Jolokia endpoint, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'clientPrincipal=cn=system:master-proxy', 'cn=hawtio-online.hawtio.svc' and 'cn=fuse-console.fuse.svc' for OpenShift).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"discovery_enabled": schema.BoolAttribute{
										Description:         "Listen for multicast requests (default 'false')",
										MarkdownDescription: "Listen for multicast requests (default 'false')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"extended_client_check": schema.BoolAttribute{
										Description:         "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",
										MarkdownDescription: "Mandate the client certificate contains a client flag in the extended key usage section, applicable when 'protocol' is 'https' and 'use-ssl-client-authentication' is 'true' (default 'true' for OpenShift).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"host": schema.StringAttribute{
										Description:         "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",
										MarkdownDescription: "The Host address to which the Jolokia agent should bind to. If ''*'' or ''0.0.0.0'' is given, the servers binds to every network interface (default ''*'').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"options": schema.ListAttribute{
										Description:         "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",
										MarkdownDescription: "A list of additional Jolokia options as defined in https://jolokia.org/reference/html/agents.html#agent-jvm-config[JVM agent configuration options]",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"password": schema.StringAttribute{
										Description:         "The password used for authentication, applicable when the 'user' option is set.",
										MarkdownDescription: "The password used for authentication, applicable when the 'user' option is set.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"port": schema.Int64Attribute{
										Description:         "The Jolokia endpoint port (default '8778').",
										MarkdownDescription: "The Jolokia endpoint port (default '8778').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"protocol": schema.StringAttribute{
										Description:         "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",
										MarkdownDescription: "The protocol to use, either 'http' or 'https' (default 'https' for OpenShift)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"use_ssl_client_authentication": schema.BoolAttribute{
										Description:         "Whether client certificates should be used for authentication (default 'true' for OpenShift).",
										MarkdownDescription: "Whether client certificates should be used for authentication (default 'true' for OpenShift).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"user": schema.StringAttribute{
										Description:         "The user to be used for authentication",
										MarkdownDescription: "The user to be used for authentication",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"jvm": schema.SingleNestedAttribute{
								Description:         "The configuration of JVM trait",
								MarkdownDescription: "The configuration of JVM trait",
								Attributes: map[string]schema.Attribute{
									"classpath": schema.StringAttribute{
										Description:         "Additional JVM classpath (use 'Linux' classpath separator)",
										MarkdownDescription: "Additional JVM classpath (use 'Linux' classpath separator)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"debug": schema.BoolAttribute{
										Description:         "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",
										MarkdownDescription: "Activates remote debugging, so that a debugger can be attached to the JVM, e.g., using port-forwarding",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"debug_address": schema.StringAttribute{
										Description:         "Transport address at which to listen for the newly launched JVM (default '*:5005')",
										MarkdownDescription: "Transport address at which to listen for the newly launched JVM (default '*:5005')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"debug_suspend": schema.BoolAttribute{
										Description:         "Suspends the target JVM immediately before the main class is loaded",
										MarkdownDescription: "Suspends the target JVM immediately before the main class is loaded",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"options": schema.ListAttribute{
										Description:         "A list of JVM options",
										MarkdownDescription: "A list of JVM options",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"print_command": schema.BoolAttribute{
										Description:         "Prints the command used the start the JVM in the container logs (default 'true')",
										MarkdownDescription: "Prints the command used the start the JVM in the container logs (default 'true')",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"kamelets": schema.SingleNestedAttribute{
								Description:         "The configuration of Kamelets trait",
								MarkdownDescription: "The configuration of Kamelets trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",
										MarkdownDescription: "Automatically inject all referenced Kamelets and their default configuration (enabled by default)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"list": schema.StringAttribute{
										Description:         "Comma separated list of Kamelet names to load into the current integration",
										MarkdownDescription: "Comma separated list of Kamelet names to load into the current integration",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"keda": schema.SingleNestedAttribute{
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",
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

							"knative": schema.SingleNestedAttribute{
								Description:         "The configuration of Knative trait",
								MarkdownDescription: "The configuration of Knative trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "Enable automatic discovery of all trait properties.",
										MarkdownDescription: "Enable automatic discovery of all trait properties.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"channel_sinks": schema.ListAttribute{
										Description:         "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",
										MarkdownDescription: "List of channels used as destination of integration routes. Can contain simple channel names or full Camel URIs.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"channel_sources": schema.ListAttribute{
										Description:         "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",
										MarkdownDescription: "List of channels used as source of integration routes. Can contain simple channel names or full Camel URIs.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"config": schema.StringAttribute{
										Description:         "Can be used to inject a Knative complete configuration in JSON format.",
										MarkdownDescription: "Can be used to inject a Knative complete configuration in JSON format.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"endpoint_sinks": schema.ListAttribute{
										Description:         "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",
										MarkdownDescription: "List of endpoints used as destination of integration routes. Can contain simple endpoint names or full Camel URIs.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"endpoint_sources": schema.ListAttribute{
										Description:         "List of channels used as source of integration routes.",
										MarkdownDescription: "List of channels used as source of integration routes.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"event_sinks": schema.ListAttribute{
										Description:         "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",
										MarkdownDescription: "List of event types that the integration will produce. Can contain simple event types or full Camel URIs (to use a specific broker).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"event_sources": schema.ListAttribute{
										Description:         "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",
										MarkdownDescription: "List of event types that the integration will be subscribed to. Can contain simple event types or full Camel URIs (to use a specific broker different from 'default').",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"filter_source_channels": schema.BoolAttribute{
										Description:         "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",
										MarkdownDescription: "Enables filtering on events based on the header 'ce-knativehistory'. Since this header has been removed in newer versions of Knative, filtering is disabled by default.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"namespace_label": schema.BoolAttribute{
										Description:         "Enables the camel-k-operator to set the 'bindings.knative.dev/include=true' label to the namespace As Knative requires this label to perform injection of K_SINK URL into the service. If this is false, the integration pod may start and fail, read the SinkBinding Knative documentation. (default: true)",
										MarkdownDescription: "Enables the camel-k-operator to set the 'bindings.knative.dev/include=true' label to the namespace As Knative requires this label to perform injection of K_SINK URL into the service. If this is false, the integration pod may start and fail, read the SinkBinding Knative documentation. (default: true)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"sink_binding": schema.BoolAttribute{
										Description:         "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",
										MarkdownDescription: "Allows binding the integration to a sink via a Knative SinkBinding resource. This can be used when the integration targets a single sink. It's enabled by default when the integration targets a single sink (except when the integration is owned by a Knative source).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"knative_service": schema.SingleNestedAttribute{
								Description:         "The configuration of Knative Service trait",
								MarkdownDescription: "The configuration of Knative Service trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "Automatically deploy the integration as Knative service when all conditions hold:  * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",
										MarkdownDescription: "Automatically deploy the integration as Knative service when all conditions hold:  * Integration is using the Knative profile * All routes are either starting from an HTTP based consumer or a passive consumer (e.g. 'direct' is a passive consumer)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"autoscaling_metric": schema.StringAttribute{
										Description:         "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling).  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Configures the Knative autoscaling metric property (e.g. to set 'concurrency' based or 'cpu' based autoscaling).  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"autoscaling_target": schema.Int64Attribute{
										Description:         "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Sets the allowed concurrency level or CPU percentage (depending on the autoscaling metric) for each Pod.  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"class": schema.StringAttribute{
										Description:         "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling).  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Configures the Knative autoscaling class property (e.g. to set 'hpa.autoscaling.knative.dev' or 'kpa.autoscaling.knative.dev' autoscaling).  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_scale": schema.Int64Attribute{
										Description:         "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "An upper bound for the number of Pods that can be running in parallel for the integration. Knative has its own cap value that depends on the installation.  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"min_scale": schema.Int64Attribute{
										Description:         "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "The minimum number of Pods that should be running at any time for the integration. It's **zero** by default, meaning that the integration is scaled down to zero when not used for a configured amount of time.  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"rollout_duration": schema.StringAttribute{
										Description:         "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",
										MarkdownDescription: "Enables to gradually shift traffic to the latest Revision and sets the rollout duration. It's disabled by default and must be expressed as a Golang 'time.Duration' string representation, rounded to a second precision.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"visibility": schema.StringAttribute{
										Description:         "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service.  Refer to the Knative documentation for more information.",
										MarkdownDescription: "Setting 'cluster-local', Knative service becomes a private service. Specifically, this option applies the 'networking.knative.dev/visibility' label to Knative service.  Refer to the Knative documentation for more information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"logging": schema.SingleNestedAttribute{
								Description:         "The configuration of Logging trait",
								MarkdownDescription: "The configuration of Logging trait",
								Attributes: map[string]schema.Attribute{
									"color": schema.BoolAttribute{
										Description:         "Colorize the log output",
										MarkdownDescription: "Colorize the log output",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"format": schema.StringAttribute{
										Description:         "Logs message format",
										MarkdownDescription: "Logs message format",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"json": schema.BoolAttribute{
										Description:         "Output the logs in JSON",
										MarkdownDescription: "Output the logs in JSON",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"json_pretty_print": schema.BoolAttribute{
										Description:         "Enable 'pretty printing' of the JSON logs",
										MarkdownDescription: "Enable 'pretty printing' of the JSON logs",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"level": schema.StringAttribute{
										Description:         "Adjust the logging level (defaults to INFO)",
										MarkdownDescription: "Adjust the logging level (defaults to INFO)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"master": schema.SingleNestedAttribute{
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",
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

							"mount": schema.SingleNestedAttribute{
								Description:         "The configuration of Mount trait",
								MarkdownDescription: "The configuration of Mount trait",
								Attributes: map[string]schema.Attribute{
									"configs": schema.ListAttribute{
										Description:         "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",
										MarkdownDescription: "A list of configuration pointing to configmap/secret. The configuration are expected to be UTF-8 resources as they are processed by runtime Camel Context and tried to be parsed as property files. They are also made available on the classpath in order to ease their usage directly from the Route. Syntax: [configmap|secret]:name[/key], where name represents the resource name and key optionally represents the resource key to be filtered",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"hot_reload": schema.BoolAttribute{
										Description:         "Enable 'hot reload' when a secret/configmap mounted is edited (default false)",
										MarkdownDescription: "Enable 'hot reload' when a secret/configmap mounted is edited (default false)",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"resources": schema.ListAttribute{
										Description:         "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",
										MarkdownDescription: "A list of resources (text or binary content) pointing to configmap/secret. The resources are expected to be any resource type (text or binary content). The destination path can be either a default location or any path specified by the user. Syntax: [configmap|secret]:name[/key][@path], where name represents the resource name, key optionally represents the resource key to be filtered and path represents the destination path",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"volumes": schema.ListAttribute{
										Description:         "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",
										MarkdownDescription: "A list of Persistent Volume Claims to be mounted. Syntax: [pvcname:/container/path]",
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

							"openapi": schema.SingleNestedAttribute{
								Description:         "The configuration of OpenAPI trait",
								MarkdownDescription: "The configuration of OpenAPI trait",
								Attributes: map[string]schema.Attribute{
									"configmaps": schema.ListAttribute{
										Description:         "The configmaps holding the spec of the OpenAPI",
										MarkdownDescription: "The configmaps holding the spec of the OpenAPI",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"target_annotations": schema.ListAttribute{
										Description:         "The set of annotations to be transferred",
										MarkdownDescription: "The set of annotations to be transferred",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"target_labels": schema.ListAttribute{
										Description:         "The set of labels to be transferred",
										MarkdownDescription: "The set of labels to be transferred",
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

							"pdb": schema.SingleNestedAttribute{
								Description:         "The configuration of PDB trait",
								MarkdownDescription: "The configuration of PDB trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_unavailable": schema.StringAttribute{
										Description:         "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",
										MarkdownDescription: "The number of pods for the Integration that can be unavailable after an eviction. It can be either an absolute number or a percentage (default '1' if 'min-available' is also not set). Only one of 'max-unavailable' and 'min-available' can be specified.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"min_available": schema.StringAttribute{
										Description:         "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",
										MarkdownDescription: "The number of pods for the Integration that must still be available after an eviction. It can be either an absolute number or a percentage. Only one of 'min-available' and 'max-unavailable' can be specified.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"platform": schema.SingleNestedAttribute{
								Description:         "The configuration of Platform trait",
								MarkdownDescription: "The configuration of Platform trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift only).",
										MarkdownDescription: "To automatically detect from the environment if a default platform can be created (it will be created on OpenShift only).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"create_default": schema.BoolAttribute{
										Description:         "To create a default (empty) platform when the platform is missing.",
										MarkdownDescription: "To create a default (empty) platform when the platform is missing.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"global": schema.BoolAttribute{
										Description:         "Indicates if the platform should be created globally in the case of global operator (default true).",
										MarkdownDescription: "Indicates if the platform should be created globally in the case of global operator (default true).",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_monitor": schema.BoolAttribute{
										Description:         "Whether a 'PodMonitor' resource is created (default 'true').",
										MarkdownDescription: "Whether a 'PodMonitor' resource is created (default 'true').",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"pod_monitor_labels": schema.ListAttribute{
										Description:         "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",
										MarkdownDescription: "The 'PodMonitor' resource labels, applicable when 'pod-monitor' is 'true'.",
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

							"pull_secret": schema.SingleNestedAttribute{
								Description:         "The configuration of Pull Secret trait",
								MarkdownDescription: "The configuration of Pull Secret trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",
										MarkdownDescription: "Automatically configures the platform registry secret on the pod if it is of type 'kubernetes.io/dockerconfigjson'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"image_puller_delegation": schema.BoolAttribute{
										Description:         "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",
										MarkdownDescription: "When using a global operator with a shared platform, this enables delegation of the 'system:image-puller' cluster role on the operator namespace to the integration service account.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_name": schema.StringAttribute{
										Description:         "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",
										MarkdownDescription: "The pull secret name to set on the Pod. If left empty this is automatically taken from the 'IntegrationPlatform' registry configuration.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"quarkus": schema.SingleNestedAttribute{
								Description:         "The configuration of Quarkus trait",
								MarkdownDescription: "The configuration of Quarkus trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"package_types": schema.ListAttribute{
										Description:         "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",
										MarkdownDescription: "The Quarkus package types, either 'fast-jar' or 'native' (default 'fast-jar'). In case both 'fast-jar' and 'native' are specified, two 'IntegrationKit' resources are created, with the 'native' kit having precedence over the 'fast-jar' one once ready. The order influences the resolution of the current kit for the integration. The kit corresponding to the first package type will be assigned to the integration in case no existing kit that matches the integration exists.",
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

							"registry": schema.SingleNestedAttribute{
								Description:         "The configuration of Registry trait",
								MarkdownDescription: "The configuration of Registry trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"host": schema.StringAttribute{
										Description:         "To configure the host exposed by the route.",
										MarkdownDescription: "To configure the host exposed by the route.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_ca_certificate": schema.StringAttribute{
										Description:         "The TLS CA certificate contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS CA certificate contents.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_ca_certificate_secret": schema.StringAttribute{
										Description:         "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_certificate": schema.StringAttribute{
										Description:         "The TLS certificate contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS certificate contents.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_certificate_secret": schema.StringAttribute{
										Description:         "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_destination_ca_certificate": schema.StringAttribute{
										Description:         "The destination CA certificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The destination CA certificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_destination_ca_certificate_secret": schema.StringAttribute{
										Description:         "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the destination CA certificate. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_insecure_edge_termination_policy": schema.StringAttribute{
										Description:         "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "To configure how to deal with insecure traffic, e.g. 'Allow', 'Disable' or 'Redirect' traffic.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_key": schema.StringAttribute{
										Description:         "The TLS certificate key contents.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS certificate key contents.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_key_secret": schema.StringAttribute{
										Description:         "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The secret name and key reference to the TLS certificate key. The format is 'secret-name[/key-name]', the value represents the secret name, if there is only one key in the secret it will be read, otherwise you can set a key name separated with a '/'.  Refer to the OpenShift route documentation for additional information.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tls_termination": schema.StringAttribute{
										Description:         "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'.  Refer to the OpenShift route documentation for additional information.",
										MarkdownDescription: "The TLS termination type, like 'edge', 'passthrough' or 'reencrypt'.  Refer to the OpenShift route documentation for additional information.",
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
								Description:         "The configuration of Service trait",
								MarkdownDescription: "The configuration of Service trait",
								Attributes: map[string]schema.Attribute{
									"auto": schema.BoolAttribute{
										Description:         "To automatically detect from the code if a Service needs to be created.",
										MarkdownDescription: "To automatically detect from the code if a Service needs to be created.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"node_port": schema.BoolAttribute{
										Description:         "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",
										MarkdownDescription: "Enable Service to be exposed as NodePort (default 'false'). Deprecated: Use service type instead.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"type": schema.StringAttribute{
										Description:         "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",
										MarkdownDescription: "The type of service to be used, either 'ClusterIP', 'NodePort' or 'LoadBalancer'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
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
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"services": schema.ListAttribute{
										Description:         "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",
										MarkdownDescription: "List of Services in the form [[apigroup/]version:]kind:[namespace/]name",
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

							"strimzi": schema.SingleNestedAttribute{
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",
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

							"toleration": schema.SingleNestedAttribute{
								Description:         "The configuration of Toleration trait",
								MarkdownDescription: "The configuration of Toleration trait",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										MarkdownDescription: "Legacy trait configuration parameters. Deprecated: for backward compatibility.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Can be used to enable or disable a trait. All traits share this common property.",
										MarkdownDescription: "Can be used to enable or disable a trait. All traits share this common property.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"taints": schema.ListAttribute{
										Description:         "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",
										MarkdownDescription: "The list of taints to tolerate, in the form 'Key[=Value]:Effect[:Seconds]'",
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

							"tracing": schema.SingleNestedAttribute{
								Description:         "Deprecated: for backward compatibility.",
								MarkdownDescription: "Deprecated: for backward compatibility.",
								Attributes: map[string]schema.Attribute{
									"configuration": schema.MapAttribute{
										Description:         "TraitConfiguration parameters configuration",
										MarkdownDescription: "TraitConfiguration parameters configuration",
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
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *CamelApacheOrgIntegrationV1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
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

func (r *CamelApacheOrgIntegrationV1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_camel_apache_org_integration_v1")

	var data CamelApacheOrgIntegrationV1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "camel.apache.org", Version: "v1", Resource: "integrations"}).
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

	var readResponse CamelApacheOrgIntegrationV1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("camel.apache.org/v1")
	data.Kind = pointer.String("Integration")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
