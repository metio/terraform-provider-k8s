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

type AppLightbendComAkkaClusterV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*AppLightbendComAkkaClusterV1Alpha1Resource)(nil)
)

type AppLightbendComAkkaClusterV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AppLightbendComAkkaClusterV1Alpha1GoModel struct {
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
		RevisionHistoryLimit *int64 `tfsdk:"revision_history_limit" yaml:"revisionHistoryLimit,omitempty"`

		Selector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"selector" yaml:"selector,omitempty"`

		Strategy *struct {
			RollingUpdate *struct {
				MaxUnavailable *string `tfsdk:"max_unavailable" yaml:"maxUnavailable,omitempty"`

				MaxSurge *string `tfsdk:"max_surge" yaml:"maxSurge,omitempty"`
			} `tfsdk:"rolling_update" yaml:"rollingUpdate,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"strategy" yaml:"strategy,omitempty"`

		Template *struct {
			Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

			Spec *struct {
				DnsConfig *struct {
					Nameservers *[]string `tfsdk:"nameservers" yaml:"nameservers,omitempty"`

					Options *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"options" yaml:"options,omitempty"`

					Searches *[]string `tfsdk:"searches" yaml:"searches,omitempty"`
				} `tfsdk:"dns_config" yaml:"dnsConfig,omitempty"`

				EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

				HostAliases *[]struct {
					Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

					Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
				} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

				PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

				Priority *int64 `tfsdk:"priority" yaml:"priority,omitempty"`

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

				TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

				Containers *[]struct {
					Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`

						ValueFrom *struct {
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

							ConfigMapKeyRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
						} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

					EnvFrom *[]struct {
						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						ConfigMapRef *struct {
							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

						Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`
					} `tfsdk:"env_from" yaml:"envFrom,omitempty"`

					Image *string `tfsdk:"image" yaml:"image,omitempty"`

					Ports *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

						HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

						HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`
					} `tfsdk:"ports" yaml:"ports,omitempty"`

					SecurityContext *struct {
						Capabilities *struct {
							Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

							Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
						} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

						Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

						ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

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

						ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

						RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

						RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

					ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

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
					} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

					Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

					VolumeDevices *[]struct {
						DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

					VolumeMounts *[]struct {
						SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

						SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`

						MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

						MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
					} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

					Resources *struct {
						Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

						Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

					TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

					Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

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
						} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
					} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

					LivenessProbe *struct {
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
					} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ReadinessProbe *struct {
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

						SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

						TcpSocket *struct {
							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

						TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
					} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

					StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

					WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
				} `tfsdk:"containers" yaml:"containers,omitempty"`

				Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

				Overhead *map[string]string `tfsdk:"overhead" yaml:"overhead,omitempty"`

				RestartPolicy *string `tfsdk:"restart_policy" yaml:"restartPolicy,omitempty"`

				RuntimeClassName *string `tfsdk:"runtime_class_name" yaml:"runtimeClassName,omitempty"`

				ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

				ShareProcessNamespace *bool `tfsdk:"share_process_namespace" yaml:"shareProcessNamespace,omitempty"`

				DnsPolicy *string `tfsdk:"dns_policy" yaml:"dnsPolicy,omitempty"`

				SchedulerName *string `tfsdk:"scheduler_name" yaml:"schedulerName,omitempty"`

				Volumes *[]struct {
					PersistentVolumeClaim *struct {
						ClaimName *string `tfsdk:"claim_name" yaml:"claimName,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
					} `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`

					Iscsi *struct {
						IscsiInterface *string `tfsdk:"iscsi_interface" yaml:"iscsiInterface,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						TargetPortal *string `tfsdk:"target_portal" yaml:"targetPortal,omitempty"`

						ChapAuthSession *bool `tfsdk:"chap_auth_session" yaml:"chapAuthSession,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Iqn *string `tfsdk:"iqn" yaml:"iqn,omitempty"`

						Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

						Portals *[]string `tfsdk:"portals" yaml:"portals,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						ChapAuthDiscovery *bool `tfsdk:"chap_auth_discovery" yaml:"chapAuthDiscovery,omitempty"`

						InitiatorName *string `tfsdk:"initiator_name" yaml:"initiatorName,omitempty"`
					} `tfsdk:"iscsi" yaml:"iscsi,omitempty"`

					Nfs *struct {
						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						Server *string `tfsdk:"server" yaml:"server,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"nfs" yaml:"nfs,omitempty"`

					DownwardAPI *struct {
						DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

						Items *[]struct {
							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

								Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

								Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

								FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

							Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`
						} `tfsdk:"items" yaml:"items,omitempty"`
					} `tfsdk:"downward_api" yaml:"downwardAPI,omitempty"`

					HostPath *struct {
						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"host_path" yaml:"hostPath,omitempty"`

					Projected *struct {
						DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

						Sources *[]struct {
							ConfigMap *struct {
								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

								Items *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`
								} `tfsdk:"items" yaml:"items,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"config_map" yaml:"configMap,omitempty"`

							DownwardAPI *struct {
								Items *[]struct {
									Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

									Path *string `tfsdk:"path" yaml:"path,omitempty"`

									ResourceFieldRef *struct {
										ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

										Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

										Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
									} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

									FieldRef *struct {
										ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

										FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
									} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`
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

					Csi *struct {
						Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						NodePublishSecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"node_publish_secret_ref" yaml:"nodePublishSecretRef,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						VolumeAttributes *map[string]string `tfsdk:"volume_attributes" yaml:"volumeAttributes,omitempty"`
					} `tfsdk:"csi" yaml:"csi,omitempty"`

					PhotonPersistentDisk *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						PdID *string `tfsdk:"pd_id" yaml:"pdID,omitempty"`
					} `tfsdk:"photon_persistent_disk" yaml:"photonPersistentDisk,omitempty"`

					PortworxVolume *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
					} `tfsdk:"portworx_volume" yaml:"portworxVolume,omitempty"`

					AzureDisk *struct {
						DiskName *string `tfsdk:"disk_name" yaml:"diskName,omitempty"`

						DiskURI *string `tfsdk:"disk_uri" yaml:"diskURI,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						CachingMode *string `tfsdk:"caching_mode" yaml:"cachingMode,omitempty"`
					} `tfsdk:"azure_disk" yaml:"azureDisk,omitempty"`

					Fc *struct {
						Wwids *[]string `tfsdk:"wwids" yaml:"wwids,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Lun *int64 `tfsdk:"lun" yaml:"lun,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						TargetWWNs *[]string `tfsdk:"target_ww_ns" yaml:"targetWWNs,omitempty"`
					} `tfsdk:"fc" yaml:"fc,omitempty"`

					Flocker *struct {
						DatasetUUID *string `tfsdk:"dataset_uuid" yaml:"datasetUUID,omitempty"`

						DatasetName *string `tfsdk:"dataset_name" yaml:"datasetName,omitempty"`
					} `tfsdk:"flocker" yaml:"flocker,omitempty"`

					Glusterfs *struct {
						Endpoints *string `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
					} `tfsdk:"glusterfs" yaml:"glusterfs,omitempty"`

					Quobyte *struct {
						Tenant *string `tfsdk:"tenant" yaml:"tenant,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`

						Volume *string `tfsdk:"volume" yaml:"volume,omitempty"`

						Group *string `tfsdk:"group" yaml:"group,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						Registry *string `tfsdk:"registry" yaml:"registry,omitempty"`
					} `tfsdk:"quobyte" yaml:"quobyte,omitempty"`

					ScaleIO *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						ProtectionDomain *string `tfsdk:"protection_domain" yaml:"protectionDomain,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						SslEnabled *bool `tfsdk:"ssl_enabled" yaml:"sslEnabled,omitempty"`

						StorageMode *string `tfsdk:"storage_mode" yaml:"storageMode,omitempty"`

						VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`

						Gateway *string `tfsdk:"gateway" yaml:"gateway,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						StoragePool *string `tfsdk:"storage_pool" yaml:"storagePool,omitempty"`

						System *string `tfsdk:"system" yaml:"system,omitempty"`
					} `tfsdk:"scale_io" yaml:"scaleIO,omitempty"`

					ConfigMap *struct {
						DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`

						Items *[]struct {
							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`
						} `tfsdk:"items" yaml:"items,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					GitRepo *struct {
						Directory *string `tfsdk:"directory" yaml:"directory,omitempty"`

						Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`

						Revision *string `tfsdk:"revision" yaml:"revision,omitempty"`
					} `tfsdk:"git_repo" yaml:"gitRepo,omitempty"`

					FlexVolume *struct {
						Options *map[string]string `tfsdk:"options" yaml:"options,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`
					} `tfsdk:"flex_volume" yaml:"flexVolume,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Rbd *struct {
						Keyring *string `tfsdk:"keyring" yaml:"keyring,omitempty"`

						Monitors *[]string `tfsdk:"monitors" yaml:"monitors,omitempty"`

						Pool *string `tfsdk:"pool" yaml:"pool,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SecretRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Image *string `tfsdk:"image" yaml:"image,omitempty"`
					} `tfsdk:"rbd" yaml:"rbd,omitempty"`

					VsphereVolume *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						StoragePolicyID *string `tfsdk:"storage_policy_id" yaml:"storagePolicyID,omitempty"`

						StoragePolicyName *string `tfsdk:"storage_policy_name" yaml:"storagePolicyName,omitempty"`

						VolumePath *string `tfsdk:"volume_path" yaml:"volumePath,omitempty"`
					} `tfsdk:"vsphere_volume" yaml:"vsphereVolume,omitempty"`

					AwsElasticBlockStore *struct {
						VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`

						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Partition *int64 `tfsdk:"partition" yaml:"partition,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
					} `tfsdk:"aws_elastic_block_store" yaml:"awsElasticBlockStore,omitempty"`

					EmptyDir *struct {
						Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

						SizeLimit *string `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
					} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

					GcePersistentDisk *struct {
						FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

						Partition *int64 `tfsdk:"partition" yaml:"partition,omitempty"`

						PdName *string `tfsdk:"pd_name" yaml:"pdName,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`
					} `tfsdk:"gce_persistent_disk" yaml:"gcePersistentDisk,omitempty"`

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
				} `tfsdk:"volumes" yaml:"volumes,omitempty"`

				ImagePullSecrets *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

				PreemptionPolicy *string `tfsdk:"preemption_policy" yaml:"preemptionPolicy,omitempty"`

				Subdomain *string `tfsdk:"subdomain" yaml:"subdomain,omitempty"`

				ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" yaml:"activeDeadlineSeconds,omitempty"`

				EphemeralContainers *[]struct {
					TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

					VolumeDevices *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`
					} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

					VolumeMounts *[]struct {
						MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

						MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

						SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
					} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Ports *[]struct {
						HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

						HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`
					} `tfsdk:"ports" yaml:"ports,omitempty"`

					SecurityContext *struct {
						AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

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

						Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

						ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

						ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

						RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

						SeLinuxOptions *struct {
							Type *string `tfsdk:"type" yaml:"type,omitempty"`

							User *string `tfsdk:"user" yaml:"user,omitempty"`

							Level *string `tfsdk:"level" yaml:"level,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`
						} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

					StartupProbe *struct {
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
							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

						TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`

						Exec *struct {
							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
						} `tfsdk:"exec" yaml:"exec,omitempty"`

						FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`
					} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

					Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

					StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

					TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

					EnvFrom *[]struct {
						ConfigMapRef *struct {
							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

						Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

						SecretRef *struct {
							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
					} `tfsdk:"env_from" yaml:"envFrom,omitempty"`

					ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

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

					Resources *struct {
						Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

						Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					Lifecycle *struct {
						PreStop *struct {
							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port *string `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

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
						} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`

						PostStart *struct {
							Exec *struct {
								Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
							} `tfsdk:"exec" yaml:"exec,omitempty"`

							HttpGet *struct {
								HttpHeaders *[]struct {
									Value *string `tfsdk:"value" yaml:"value,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`
								} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

								Path *string `tfsdk:"path" yaml:"path,omitempty"`

								Port *string `tfsdk:"port" yaml:"port,omitempty"`

								Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

								Host *string `tfsdk:"host" yaml:"host,omitempty"`
							} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

							TcpSocket *struct {
								Port *string `tfsdk:"port" yaml:"port,omitempty"`

								Host *string `tfsdk:"host" yaml:"host,omitempty"`
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
							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`

							Port *string `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`
					} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

					TargetContainerName *string `tfsdk:"target_container_name" yaml:"targetContainerName,omitempty"`

					Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

					Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`

						ValueFrom *struct {
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

					WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
				} `tfsdk:"ephemeral_containers" yaml:"ephemeralContainers,omitempty"`

				HostNetwork *bool `tfsdk:"host_network" yaml:"hostNetwork,omitempty"`

				HostPID *bool `tfsdk:"host_pid" yaml:"hostPID,omitempty"`

				NodeName *string `tfsdk:"node_name" yaml:"nodeName,omitempty"`

				NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

				ReadinessGates *[]struct {
					ConditionType *string `tfsdk:"condition_type" yaml:"conditionType,omitempty"`
				} `tfsdk:"readiness_gates" yaml:"readinessGates,omitempty"`

				AutomountServiceAccountToken *bool `tfsdk:"automount_service_account_token" yaml:"automountServiceAccountToken,omitempty"`

				SecurityContext *struct {
					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

					SeLinuxOptions *struct {
						Level *string `tfsdk:"level" yaml:"level,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`
					} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

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

					FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" yaml:"fsGroupChangePolicy,omitempty"`

					RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

					SupplementalGroups *[]string `tfsdk:"supplemental_groups" yaml:"supplementalGroups,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

				ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

				Affinity *struct {
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
							TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

							LabelSelector *struct {
								MatchExpressions *[]struct {
									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

									Key *string `tfsdk:"key" yaml:"key,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

							Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`

					NodeAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							Preference *struct {
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
									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`
								} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
							} `tfsdk:"node_selector_terms" yaml:"nodeSelectorTerms,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`

					PodAffinity *struct {
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

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

							Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`
				} `tfsdk:"affinity" yaml:"affinity,omitempty"`

				HostIPC *bool `tfsdk:"host_ipc" yaml:"hostIPC,omitempty"`

				InitContainers *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ReadinessProbe *struct {
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
					} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

					TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

					Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`

						ValueFrom *struct {
							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

								Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

								Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

							SecretKeyRef *struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

								Key *string `tfsdk:"key" yaml:"key,omitempty"`
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

					Image *string `tfsdk:"image" yaml:"image,omitempty"`

					ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

					Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

					StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

					Lifecycle *struct {
						PostStart *struct {
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
						} `tfsdk:"post_start" yaml:"postStart,omitempty"`

						PreStop *struct {
							TcpSocket *struct {
								Host *string `tfsdk:"host" yaml:"host,omitempty"`

								Port *string `tfsdk:"port" yaml:"port,omitempty"`
							} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

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
						} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
					} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

					StartupProbe *struct {
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

						InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`
					} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

					SecurityContext *struct {
						ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

						RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

						RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

						SeLinuxOptions *struct {
							Type *string `tfsdk:"type" yaml:"type,omitempty"`

							User *string `tfsdk:"user" yaml:"user,omitempty"`

							Level *string `tfsdk:"level" yaml:"level,omitempty"`

							Role *string `tfsdk:"role" yaml:"role,omitempty"`
						} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

						WindowsOptions *struct {
							GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

							RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`

							GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`
						} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

						AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

						Capabilities *struct {
							Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

							Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
						} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

						Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

						ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

						RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`
					} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

					VolumeMounts *[]struct {
						MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

						SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`

						MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`
					} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

					WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`

					Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

					EnvFrom *[]struct {
						ConfigMapRef *struct {
							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"config_map_ref" yaml:"configMapRef,omitempty"`

						Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

						SecretRef *struct {
							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`
						} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`
					} `tfsdk:"env_from" yaml:"envFrom,omitempty"`

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
							Port *string `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							HttpHeaders *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`
							} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`
					} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

					VolumeDevices *[]struct {
						DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

					Ports *[]struct {
						Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

						ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

						HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

						HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"ports" yaml:"ports,omitempty"`

					Resources *struct {
						Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

						Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`
				} `tfsdk:"init_containers" yaml:"initContainers,omitempty"`
			} `tfsdk:"spec" yaml:"spec,omitempty"`
		} `tfsdk:"template" yaml:"template,omitempty"`

		MinReadySeconds *int64 `tfsdk:"min_ready_seconds" yaml:"minReadySeconds,omitempty"`

		Paused *bool `tfsdk:"paused" yaml:"paused,omitempty"`

		ProgressDeadlineSeconds *int64 `tfsdk:"progress_deadline_seconds" yaml:"progressDeadlineSeconds,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAppLightbendComAkkaClusterV1Alpha1Resource() resource.Resource {
	return &AppLightbendComAkkaClusterV1Alpha1Resource{}
}

func (r *AppLightbendComAkkaClusterV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_lightbend_com_akka_cluster_v1alpha1"
}

func (r *AppLightbendComAkkaClusterV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "AkkaCluster is the Schema for the akkaclusters API",
		MarkdownDescription: "AkkaCluster is the Schema for the akkaclusters API",
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
				Description:         "DeploymentSpec is the specification of the desired behavior of the Deployment.",
				MarkdownDescription: "DeploymentSpec is the specification of the desired behavior of the Deployment.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"revision_history_limit": {
						Description:         "The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.",
						MarkdownDescription: "The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"selector": {
						Description:         "Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment. It must match the pod template's labels.",
						MarkdownDescription: "Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment. It must match the pod template's labels.",

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

						Required: true,
						Optional: false,
						Computed: false,
					},

					"strategy": {
						Description:         "The deployment strategy to use to replace existing pods with new ones.",
						MarkdownDescription: "The deployment strategy to use to replace existing pods with new ones.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"rolling_update": {
								Description:         "Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate. --- TODO: Update this to follow our convention for oneOf, whatever we decide it to be.",
								MarkdownDescription: "Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate. --- TODO: Update this to follow our convention for oneOf, whatever we decide it to be.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"max_unavailable": {
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding down. This can not be 0 if MaxSurge is 0. Defaults to 25%. Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_surge": {
										Description:         "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",
										MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. Defaults to 25%. Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that total number of pods running at any time during the update is at most 130% of desired pods.",

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

							"type": {
								Description:         "Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.",
								MarkdownDescription: "Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.",

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
						Description:         "Template describes the pods that will be created.",
						MarkdownDescription: "Template describes the pods that will be created.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"metadata": {
								Description:         "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spec": {
								Description:         "Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
								MarkdownDescription: "Specification of the desired behavior of the pod. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"dns_config": {
										Description:         "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",
										MarkdownDescription: "Specifies the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS configuration based on DNSPolicy.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"nameservers": {
												Description:         "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",
												MarkdownDescription: "A list of DNS name server IP addresses. This will be appended to the base nameservers generated from DNSPolicy. Duplicated nameservers will be removed.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"options": {
												Description:         "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",
												MarkdownDescription: "A list of DNS resolver options. This will be merged with the base options generated from DNSPolicy. Duplicated entries will be removed. Resolution options given in Options will override those that appear in the base DNSPolicy.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Required.",
														MarkdownDescription: "Required.",

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
												Description:         "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",
												MarkdownDescription: "A list of DNS search domains for host-name lookup. This will be appended to the base search paths generated from DNSPolicy. Duplicated search paths will be removed.",

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

									"enable_service_links": {
										Description:         "EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.",
										MarkdownDescription: "EnableServiceLinks indicates whether information about services should be injected into pod's environment variables, matching the syntax of Docker links. Optional: Defaults to true.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host_aliases": {
										Description:         "HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.",
										MarkdownDescription: "HostAliases is an optional list of hosts and IPs that will be injected into the pod's hosts file if specified. This is only valid for non-hostNetwork pods.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"hostnames": {
												Description:         "Hostnames for the above IP address.",
												MarkdownDescription: "Hostnames for the above IP address.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip": {
												Description:         "IP address of the host file entry.",
												MarkdownDescription: "IP address of the host file entry.",

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

									"priority_class_name": {
										Description:         "If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.",
										MarkdownDescription: "If specified, indicates the pod's priority. 'system-node-critical' and 'system-cluster-critical' are two special keywords which indicate the highest priorities with the former being the highest priority. Any other name must be defined by creating a PriorityClass object with that name. If not specified, the pod priority will be default or zero if there is no default.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"priority": {
										Description:         "The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.",
										MarkdownDescription: "The priority value. Various system components use this field to find the priority of the pod. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"topology_spread_constraints": {
										Description:         "TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. This field is only honored by clusters that enable the EvenPodsSpread feature. All topologySpreadConstraints are ANDed.",
										MarkdownDescription: "TopologySpreadConstraints describes how a group of pods ought to spread across topology domains. Scheduler will schedule pods in a way which abides by the constraints. This field is only honored by clusters that enable the EvenPodsSpread feature. All topologySpreadConstraints are ANDed.",

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
												Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. It's the maximum permitted difference between the number of matching pods in any two topology domains of a given topology type. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 1/1/0: | zone1 | zone2 | zone3 | |   P   |   P   |       | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 1/1/1; scheduling it onto zone1(zone2) would make the ActualSkew(2-0) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. It's a required field. Default value is 1 and 0 is not allowed.",
												MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. It's the maximum permitted difference between the number of matching pods in any two topology domains of a given topology type. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 1/1/0: | zone1 | zone2 | zone3 | |   P   |   P   |       | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 1/1/1; scheduling it onto zone1(zone2) would make the ActualSkew(2-0) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. It's a required field. Default value is 1 and 0 is not allowed.",

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
												Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it - ScheduleAnyway tells the scheduler to still schedule it It's considered as 'Unsatisfiable' if and only if placing incoming pod on any topology violates 'MaxSkew'. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
												MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it - ScheduleAnyway tells the scheduler to still schedule it It's considered as 'Unsatisfiable' if and only if placing incoming pod on any topology violates 'MaxSkew'. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",

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

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully. May be decreased in delete request. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. Defaults to 30 seconds.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"containers": {
										Description:         "List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.",
										MarkdownDescription: "List of containers belonging to the pod. Containers cannot currently be added or removed. There must be at least one container in a Pod. Cannot be updated.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"stdin": {
												Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
												MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",

												Type: types.BoolType,

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
														Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value_from": {
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"field_ref": {
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

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

																		Type: types.StringType,

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

													"config_map_ref": {
														Description:         "The ConfigMap to select from",
														MarkdownDescription: "The ConfigMap to select from",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"optional": {
																Description:         "Specify whether the ConfigMap must be defined",
																MarkdownDescription: "Specify whether the ConfigMap must be defined",

																Type: types.BoolType,

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

											"ports": {
												Description:         "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",
												MarkdownDescription: "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": {
												Description:         "Security options the pod should run with. More info: https://kubernetes.io/docs/concepts/policy/security-context/ More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
												MarkdownDescription: "Security options the pod should run with. More info: https://kubernetes.io/docs/concepts/policy/security-context/ More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"capabilities": {
														Description:         "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.",
														MarkdownDescription: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.",

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
														Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.",
														MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only_root_filesystem": {
														Description:         "Whether this container has a read-only root filesystem. Default is false.",
														MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"se_linux_options": {
														Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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

													"windows_options": {
														Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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

													"allow_privilege_escalation": {
														Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN",
														MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"proc_mount": {
														Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled.",
														MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_group": {
														Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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
												}),

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

											"startup_probe": {
												Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. This is a beta feature enabled by the StartupProbe feature flag. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. This is a beta feature enabled by the StartupProbe feature flag. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
														MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

																Type: types.StringType,

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
														Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": {
												Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"limits": {
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requests": {
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

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

											"args": {
												Description:         "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"command": {
												Description:         "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

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
																Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
																MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

																		Type: types.StringType,

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
																Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
																MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
														Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period. Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period. Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
																MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

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

																		Type: types.StringType,

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
																Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
																MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
												Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
														MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

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
														Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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

											"readiness_probe": {
												Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
														MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: types.StringType,

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
														Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

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

											"stdin_once": {
												Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
												MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",

												Type: types.BoolType,

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

									"hostname": {
										Description:         "Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.",
										MarkdownDescription: "Specifies the hostname of the Pod If not specified, the pod's hostname will be set to a system-defined value.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"overhead": {
										Description:         "Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.16, and is only honored by servers that enable the PodOverhead feature.",
										MarkdownDescription: "Overhead represents the resource overhead associated with running a pod for a given RuntimeClass. This field will be autopopulated at admission time by the RuntimeClass admission controller. If the RuntimeClass admission controller is enabled, overhead must not be set in Pod create requests. The RuntimeClass admission controller will reject Pod create requests which have the overhead already set. If RuntimeClass is configured and selected in the PodSpec, Overhead will be set to the value defined in the corresponding RuntimeClass, otherwise it will remain unset and treated as zero. More info: https://git.k8s.io/enhancements/keps/sig-node/20190226-pod-overhead.md This field is alpha-level as of Kubernetes v1.16, and is only honored by servers that enable the PodOverhead feature.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"restart_policy": {
										Description:         "Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy",
										MarkdownDescription: "Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"runtime_class_name": {
										Description:         "RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the 'legacy' RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is a beta feature as of Kubernetes v1.14.",
										MarkdownDescription: "RuntimeClassName refers to a RuntimeClass object in the node.k8s.io group, which should be used to run this pod.  If no RuntimeClass resource matches the named class, the pod will not be run. If unset or empty, the 'legacy' RuntimeClass will be used, which is an implicit class with an empty definition that uses the default runtime handler. More info: https://git.k8s.io/enhancements/keps/sig-node/runtime-class.md This is a beta feature as of Kubernetes v1.14.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_account": {
										Description:         "DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.",
										MarkdownDescription: "DeprecatedServiceAccount is a depreciated alias for ServiceAccountName. Deprecated: Use serviceAccountName instead.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"share_process_namespace": {
										Description:         "Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.",
										MarkdownDescription: "Share a single process namespace between all of the containers in a pod. When this is set containers will be able to view and signal processes from other containers in the same pod, and the first process in each container will not be assigned PID 1. HostPID and ShareProcessNamespace cannot both be set. Optional: Default to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"dns_policy": {
										Description:         "Set DNS policy for the pod. Defaults to 'ClusterFirst'. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.",
										MarkdownDescription: "Set DNS policy for the pod. Defaults to 'ClusterFirst'. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scheduler_name": {
										Description:         "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.",
										MarkdownDescription: "If specified, the pod will be dispatched by specified scheduler. If not specified, the pod will be dispatched by default scheduler.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volumes": {
										Description:         "List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes",
										MarkdownDescription: "List of volumes that can be mounted by containers belonging to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

											"iscsi": {
												Description:         "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
												MarkdownDescription: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"iscsi_interface": {
														Description:         "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
														MarkdownDescription: "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",

														Type: types.StringType,

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

													"target_portal": {
														Description:         "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
														MarkdownDescription: "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

														Type: types.StringType,

														Required: true,
														Optional: false,
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

													"iqn": {
														Description:         "Target iSCSI Qualified Name.",
														MarkdownDescription: "Target iSCSI Qualified Name.",

														Type: types.StringType,

														Required: true,
														Optional: false,
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

													"chap_auth_discovery": {
														Description:         "whether support iSCSI Discovery CHAP authentication",
														MarkdownDescription: "whether support iSCSI Discovery CHAP authentication",

														Type: types.BoolType,

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"nfs": {
												Description:         "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
												MarkdownDescription: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

													"path": {
														Description:         "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

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

											"downward_api": {
												Description:         "DownwardAPI represents downward API about the pod that should populate this volume",
												MarkdownDescription: "DownwardAPI represents downward API about the pod that should populate this volume",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "Items is a list of downward API volume file",
														MarkdownDescription: "Items is a list of downward API volume file",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

																		Type: types.StringType,

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
																Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

											"host_path": {
												Description:         "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
												MarkdownDescription: "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"type": {
														Description:         "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

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
														Description:         "Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

																	"optional": {
																		Description:         "Specify whether the ConfigMap or its keys must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",

																		Type: types.BoolType,

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
																				Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

																			"mode": {
																				Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

																						Type: types.StringType,

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
																				Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

														Required: true,
														Optional: false,
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
														Description:         "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
																Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

											"cephfs": {
												Description:         "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
												MarkdownDescription: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"user": {
														Description:         "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

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

											"csi": {
												Description:         "CSI (Container Storage Interface) represents storage that is handled by an external CSI driver (Alpha feature).",
												MarkdownDescription: "CSI (Container Storage Interface) represents storage that is handled by an external CSI driver (Alpha feature).",

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

											"azure_disk": {
												Description:         "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
												MarkdownDescription: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

													"caching_mode": {
														Description:         "Host Caching mode: None, Read Only, Read Write.",
														MarkdownDescription: "Host Caching mode: None, Read Only, Read Write.",

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
												Description:         "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
												MarkdownDescription: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"wwids": {
														Description:         "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
														MarkdownDescription: "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"flocker": {
												Description:         "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
												MarkdownDescription: "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dataset_uuid": {
														Description:         "UUID of the dataset. This is unique identifier of a Flocker dataset",
														MarkdownDescription: "UUID of the dataset. This is unique identifier of a Flocker dataset",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dataset_name": {
														Description:         "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
														MarkdownDescription: "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",

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

											"quobyte": {
												Description:         "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
												MarkdownDescription: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

													"protection_domain": {
														Description:         "The name of the ScaleIO Protection Domain for the configured storage.",
														MarkdownDescription: "The name of the ScaleIO Protection Domain for the configured storage.",

														Type: types.StringType,

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

													"volume_name": {
														Description:         "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
														MarkdownDescription: "The name of a volume already created in the ScaleIO system that is associated with this volume source.",

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

													"read_only": {
														Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

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
														Description:         "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"path": {
																Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"key": {
																Description:         "The key to project.",
																MarkdownDescription: "The key to project.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mode": {
																Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

											"flex_volume": {
												Description:         "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
												MarkdownDescription: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"rbd": {
												Description:         "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
												MarkdownDescription: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"aws_elastic_block_store": {
												Description:         "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
												MarkdownDescription: "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"volume_id": {
														Description:         "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_pull_secrets": {
										Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",
										MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any of the images used by this PodSpec. If specified, these secrets will be passed to individual puller implementations for them to use. For example, in the case of docker, only DockerConfig type secrets are honored. More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

									"preemption_policy": {
										Description:         "PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.",
										MarkdownDescription: "PreemptionPolicy is the Policy for preempting pods with lower priority. One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. This field is alpha-level and is only honored by servers that enable the NonPreemptingPriority feature.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"subdomain": {
										Description:         "If specified, the fully qualified Pod hostname will be '<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>'. If not specified, the pod will not have a domainname at all.",
										MarkdownDescription: "If specified, the fully qualified Pod hostname will be '<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>'. If not specified, the pod will not have a domainname at all.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"active_deadline_seconds": {
										Description:         "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",
										MarkdownDescription: "Optional duration in seconds the pod may be active on the node relative to StartTime before the system will actively try to mark it failed and kill associated containers. Value must be a positive integer.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ephemeral_containers": {
										Description:         "List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing pod to perform user-initiated actions such as debugging. This list cannot be specified when creating a pod, and it cannot be modified by updating the pod spec. In order to add an ephemeral container to an existing pod, use the pod's ephemeralcontainers subresource. This field is alpha-level and is only honored by servers that enable the EphemeralContainers feature.",
										MarkdownDescription: "List of ephemeral containers run in this pod. Ephemeral containers may be run in an existing pod to perform user-initiated actions such as debugging. This list cannot be specified when creating a pod, and it cannot be modified by updating the pod spec. In order to add an ephemeral container to an existing pod, use the pod's ephemeralcontainers subresource. This field is alpha-level and is only honored by servers that enable the EphemeralContainers feature.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"termination_message_policy": {
												Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
												MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_devices": {
												Description:         "volumeDevices is the list of block devices to be used by the container.",
												MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "name must match the name of a persistentVolumeClaim in the pod",
														MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"device_path": {
														Description:         "devicePath is the path inside of the container that the device will be mapped to.",
														MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",

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

													"container_port": {
														Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
														MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",

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
												Description:         "SecurityContext is not allowed for ephemeral containers.",
												MarkdownDescription: "SecurityContext is not allowed for ephemeral containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"allow_privilege_escalation": {
														Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN",
														MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"capabilities": {
														Description:         "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.",
														MarkdownDescription: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"drop": {
																Description:         "Removed capabilities",
																MarkdownDescription: "Removed capabilities",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"add": {
																Description:         "Added capabilities",
																MarkdownDescription: "Added capabilities",

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
														Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_user": {
														Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"windows_options": {
														Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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

													"privileged": {
														Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.",
														MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"proc_mount": {
														Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled.",
														MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only_root_filesystem": {
														Description:         "Whether this container has a read-only root filesystem. Default is false.",
														MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false.",

														Type: types.BoolType,

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

													"se_linux_options": {
														Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

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
														Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

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
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exec": {
														Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
														MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

															"optional": {
																Description:         "Specify whether the Secret must be defined",
																MarkdownDescription: "Specify whether the Secret must be defined",

																Type: types.BoolType,

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
												Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
												MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"liveness_probe": {
												Description:         "Probes are not allowed for ephemeral containers.",
												MarkdownDescription: "Probes are not allowed for ephemeral containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
														Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exec": {
														Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
														MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

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

											"resources": {
												Description:         "Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
												MarkdownDescription: "Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"limits": {
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requests": {
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

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

											"lifecycle": {
												Description:         "Lifecycle is not allowed for ephemeral containers.",
												MarkdownDescription: "Lifecycle is not allowed for ephemeral containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"pre_stop": {
														Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period. Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period. Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"tcp_socket": {
																Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
																MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
																Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
																MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

																		Type: types.StringType,

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

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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

													"post_start": {
														Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"exec": {
																Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
																MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

																	"http_headers": {
																		Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																		MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"value": {
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"name": {
																				Description:         "The header field name",
																				MarkdownDescription: "The header field name",

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

																		Type: types.StringType,

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

																	"host": {
																		Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																		MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
																Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
																MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"port": {
																		Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																		MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"host": {
																		Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																		MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

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

											"readiness_probe": {
												Description:         "Probes are not allowed for ephemeral containers.",
												MarkdownDescription: "Probes are not allowed for ephemeral containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
														Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exec": {
														Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
														MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"scheme": {
																Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
																MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

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

											"target_container_name": {
												Description:         "If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container is run in whatever namespaces are shared for the pod. Note that the container runtime must support this feature.",
												MarkdownDescription: "If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container is run in whatever namespaces are shared for the pod. Note that the container runtime must support this feature.",

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

											"args": {
												Description:         "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"command": {
												Description:         "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

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
														Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value_from": {
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

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

																		Type: types.StringType,

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

											"image": {
												Description:         "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images",
												MarkdownDescription: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images",

												Type: types.StringType,

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

									"host_network": {
										Description:         "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.",
										MarkdownDescription: "Host networking requested for this pod. Use the host's network namespace. If this option is set, the ports that will be used must be specified. Default to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host_pid": {
										Description:         "Use the host's pid namespace. Optional: Default to false.",
										MarkdownDescription: "Use the host's pid namespace. Optional: Default to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_name": {
										Description:         "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.",
										MarkdownDescription: "NodeName is a request to schedule this pod onto a specific node. If it is non-empty, the scheduler simply schedules this pod onto that node, assuming that it fits resource requirements.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_selector": {
										Description:         "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
										MarkdownDescription: "NodeSelector is a selector which must be true for the pod to fit on a node. Selector which must match a node's labels for the pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_gates": {
										Description:         "If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to 'True' More info: https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md",
										MarkdownDescription: "If specified, all readiness gates will be evaluated for pod readiness. A pod is ready when all its containers are ready AND all conditions specified in the readiness gates have status equal to 'True' More info: https://git.k8s.io/enhancements/keps/sig-network/0007-pod-ready%2B%2B.md",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"condition_type": {
												Description:         "ConditionType refers to a condition in the pod's condition list with matching type.",
												MarkdownDescription: "ConditionType refers to a condition in the pod's condition list with matching type.",

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

									"automount_service_account_token": {
										Description:         "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.",
										MarkdownDescription: "AutomountServiceAccountToken indicates whether a service account token should be automatically mounted.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"security_context": {
										Description:         "SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field.",
										MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings. Optional: Defaults to empty.  See type description for default values of each field.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"run_as_group": {
												Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",
												MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",

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

											"se_linux_options": {
												Description:         "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",
												MarkdownDescription: "The SELinux context to be applied to all containers. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",

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

											"sysctls": {
												Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch.",
												MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported sysctls (by the container runtime) might fail to launch.",

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
												Description:         "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options within a container's SecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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

											"fs_group": {
												Description:         "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume.",
												MarkdownDescription: "A special supplemental group that applies to all containers in a pod. Some volume types allow the Kubelet to change the ownership of that volume to be owned by the pod:  1. The owning GID will be the FSGroup 2. The setgid bit is set (new files created in the volume will be owned by FSGroup) 3. The permission bits are OR'd with rw-rw----  If unset, the Kubelet will not modify the ownership and permissions of any volume.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"fs_group_change_policy": {
												Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified defaults to 'Always'.",
												MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volume before being exposed inside Pod. This field will only apply to volume types which support fsGroup based ownership(and permissions). It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir. Valid values are 'OnRootMismatch' and 'Always'. If not specified defaults to 'Always'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_user": {
												Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",
												MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in SecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence for that container.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"supplemental_groups": {
												Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container.",
												MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID.  If unspecified, no groups will be added to any container.",

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

									"service_account_name": {
										Description:         "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",
										MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run this pod. More info: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "If specified, the pod's tolerations.",
										MarkdownDescription: "If specified, the pod's tolerations.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
												MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
												MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
												MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
												MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

									"affinity": {
										Description:         "If specified, the pod's scheduling constraints",
										MarkdownDescription: "If specified, the pod's scheduling constraints",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"pod_anti_affinity": {
												Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
												MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"preferred_during_scheduling_ignored_during_execution": {
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"pod_affinity_term": {
																Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"topology_key": {
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"label_selector": {
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",

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

																	"namespaces": {
																		Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
																		MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",

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
																Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

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
														Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
														MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"topology_key": {
																Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"label_selector": {
																Description:         "A label query over a set of resources, in this case pods.",
																MarkdownDescription: "A label query over a set of resources, in this case pods.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"match_expressions": {
																		Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

																			"key": {
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",

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

															"namespaces": {
																Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
																MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",

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

											"node_affinity": {
												Description:         "Describes node affinity scheduling rules for the pod.",
												MarkdownDescription: "Describes node affinity scheduling rules for the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"preferred_during_scheduling_ignored_during_execution": {
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"preference": {
																Description:         "A node selector term, associated with the corresponding weight.",
																MarkdownDescription: "A node selector term, associated with the corresponding weight.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"match_expressions": {
																		Description:         "A list of node selector requirements by node's labels.",
																		MarkdownDescription: "A list of node selector requirements by node's labels.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"operator": {
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"key": {
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",

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
																		Description:         "A list of node selector requirements by node's fields.",
																		MarkdownDescription: "A list of node selector requirements by node's fields.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"operator": {
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

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
																Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",

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
														Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
														MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"node_selector_terms": {
																Description:         "Required. A list of node selector terms. The terms are ORed.",
																MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"match_expressions": {
																		Description:         "A list of node selector requirements by node's labels.",
																		MarkdownDescription: "A list of node selector requirements by node's labels.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"operator": {
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

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
																		Description:         "A list of node selector requirements by node's fields.",
																		MarkdownDescription: "A list of node selector requirements by node's fields.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"values": {
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"key": {
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"operator": {
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

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
												Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
												MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"required_during_scheduling_ignored_during_execution": {
														Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
														MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"label_selector": {
																Description:         "A label query over a set of resources, in this case pods.",
																MarkdownDescription: "A label query over a set of resources, in this case pods.",

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

															"namespaces": {
																Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
																MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"topology_key": {
																Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

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

													"preferred_during_scheduling_ignored_during_execution": {
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"pod_affinity_term": {
																Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"label_selector": {
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",

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

																	"namespaces": {
																		Description:         "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",
																		MarkdownDescription: "namespaces specifies which namespaces the labelSelector applies to (matches against); null or empty list means 'this pod's namespace'",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"topology_key": {
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

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
																Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

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

									"host_ipc": {
										Description:         "Use the host's ipc namespace. Optional: Default to false.",
										MarkdownDescription: "Use the host's ipc namespace. Optional: Default to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"init_containers": {
										Description:         "List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",
										MarkdownDescription: "List of initialization containers belonging to the pod. Init containers are executed in order prior to containers being started. If any init container fails, the pod is considered to have failed and is handled according to its restartPolicy. The name for an init container or normal container must be unique among all containers. Init containers may not have Lifecycle actions, Readiness probes, Liveness probes, or Startup probes. The resourceRequirements of an init container are taken into account during scheduling by finding the highest request/limit for each resource type, and then using the max of of that value or the sum of the normal containers. Limits are applied to init containers in a similar fashion. Init containers cannot currently be added or removed. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
												MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"readiness_probe": {
												Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
														Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"host": {
																Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

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
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exec": {
														Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
														MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																Type: types.StringType,

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

															"host": {
																Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
												}),

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
														Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value_from": {
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																		Type: types.StringType,

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

																	"key": {
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

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
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"field_path": {
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"api_version": {
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",

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

											"command": {
												Description:         "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

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
																Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
																MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

																		Type: types.StringType,

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
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"tcp_socket": {
																Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
																MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
														Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period. Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
														MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The reason for termination is passed to the handler. The Pod's termination grace period countdown begins before the PreStop hooked is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period. Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"tcp_socket": {
																Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
																MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
																Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
																MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

																		Type: types.StringType,

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
												Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. This is a beta feature enabled by the StartupProbe feature flag. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
												MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. This is a beta feature enabled by the StartupProbe feature flag. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
														Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exec": {
														Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
														MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

																	"value": {
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"name": {
																		Description:         "The header field name",
																		MarkdownDescription: "The header field name",

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

																Type: types.StringType,

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": {
												Description:         "Security options the pod should run with. More info: https://kubernetes.io/docs/concepts/policy/security-context/ More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
												MarkdownDescription: "Security options the pod should run with. More info: https://kubernetes.io/docs/concepts/policy/security-context/ More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"read_only_root_filesystem": {
														Description:         "Whether this container has a read-only root filesystem. Default is false.",
														MarkdownDescription: "Whether this container has a read-only root filesystem. Default is false.",

														Type: types.BoolType,

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
														Description:         "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The UID to run the entrypoint of the container process. Defaults to user specified in image metadata if unspecified. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"se_linux_options": {
														Description:         "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The SELinux context to be applied to the container. If unspecified, the container runtime will allocate a random SELinux context for each container.  May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"windows_options": {
														Description:         "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The Windows specific settings applied to all containers. If unspecified, the options from the PodSecurityContext will be used. If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"gmsa_credential_spec_name": {
																Description:         "GMSACredentialSpecName is the name of the GMSA credential spec to use.",
																MarkdownDescription: "GMSACredentialSpecName is the name of the GMSA credential spec to use.",

																Type: types.StringType,

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

															"gmsa_credential_spec": {
																Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",
																MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of the GMSA credential spec named by the GMSACredentialSpecName field.",

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
														Description:         "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN",
														MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more privileges than its parent process. This bool directly controls if the no_new_privs flag will be set on the container process. AllowPrivilegeEscalation is true always when the container is: 1) run as Privileged 2) has CAP_SYS_ADMIN",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"capabilities": {
														Description:         "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.",
														MarkdownDescription: "The capabilities to add/drop when running containers. Defaults to the default set of capabilities granted by the container runtime.",

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
														Description:         "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.",
														MarkdownDescription: "Run container in privileged mode. Processes in privileged containers are essentially equivalent to root on the host. Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"proc_mount": {
														Description:         "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled.",
														MarkdownDescription: "procMount denotes the type of proc mount to use for the containers. The default is DefaultProcMount which uses the container runtime defaults for readonly paths and masked paths. This requires the ProcMountType feature flag to be enabled.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"run_as_group": {
														Description:         "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The GID to run the entrypoint of the container process. Uses runtime default if unset. May also be set in PodSecurityContext.  If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.",

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

											"volume_mounts": {
												Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
												MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

													"mount_path": {
														Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
														MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",

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

											"working_dir": {
												Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
												MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"args": {
												Description:         "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
												MarkdownDescription: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

												Type: types.ListType{ElemType: types.StringType},

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

															"optional": {
																Description:         "Specify whether the ConfigMap must be defined",
																MarkdownDescription: "Specify whether the ConfigMap must be defined",

																Type: types.BoolType,

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

															"optional": {
																Description:         "Specify whether the Secret must be defined",
																MarkdownDescription: "Specify whether the Secret must be defined",

																Type: types.BoolType,

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
														Description:         "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported TODO: implement a realistic TCP lifecycle hook",

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
														Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
														MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"exec": {
														Description:         "One and only one of the following should be specified. Exec specifies the action to take.",
														MarkdownDescription: "One and only one of the following should be specified. Exec specifies the action to take.",

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

													"http_get": {
														Description:         "HTTPGet specifies the http request to perform.",
														MarkdownDescription: "HTTPGet specifies the http request to perform.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"port": {
																Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
																MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

																Type: types.StringType,

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

											"ports": {
												Description:         "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",
												MarkdownDescription: "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"protocol": {
														Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
														MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": {
												Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
												MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"limits": {
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"requests": {
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container/",

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

											"termination_message_path": {
												Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
												MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",

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

						Required: true,
						Optional: false,
						Computed: false,
					},

					"min_ready_seconds": {
						Description:         "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",
						MarkdownDescription: "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"paused": {
						Description:         "Indicates that the deployment is paused.",
						MarkdownDescription: "Indicates that the deployment is paused.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"progress_deadline_seconds": {
						Description:         "The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.",
						MarkdownDescription: "The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": {
						Description:         "Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.",
						MarkdownDescription: "Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.",

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

func (r *AppLightbendComAkkaClusterV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_app_lightbend_com_akka_cluster_v1alpha1")

	var state AppLightbendComAkkaClusterV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppLightbendComAkkaClusterV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("app.lightbend.com/v1alpha1")
	goModel.Kind = utilities.Ptr("AkkaCluster")

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

func (r *AppLightbendComAkkaClusterV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_lightbend_com_akka_cluster_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *AppLightbendComAkkaClusterV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_app_lightbend_com_akka_cluster_v1alpha1")

	var state AppLightbendComAkkaClusterV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppLightbendComAkkaClusterV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("app.lightbend.com/v1alpha1")
	goModel.Kind = utilities.Ptr("AkkaCluster")

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

func (r *AppLightbendComAkkaClusterV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_app_lightbend_com_akka_cluster_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
