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

type OpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*OpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource)(nil)
)

type OpentelemetryIoOpenTelemetryCollectorV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type OpentelemetryIoOpenTelemetryCollectorV1Alpha1GoModel struct {
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

		Args *map[string]string `tfsdk:"args" yaml:"args,omitempty"`

		Autoscaler *struct {
			Behavior *struct {
				ScaleDown *struct {
					Policies *[]struct {
						PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						Value *int64 `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"policies" yaml:"policies,omitempty"`

					SelectPolicy *string `tfsdk:"select_policy" yaml:"selectPolicy,omitempty"`

					StabilizationWindowSeconds *int64 `tfsdk:"stabilization_window_seconds" yaml:"stabilizationWindowSeconds,omitempty"`
				} `tfsdk:"scale_down" yaml:"scaleDown,omitempty"`

				ScaleUp *struct {
					Policies *[]struct {
						PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						Value *int64 `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"policies" yaml:"policies,omitempty"`

					SelectPolicy *string `tfsdk:"select_policy" yaml:"selectPolicy,omitempty"`

					StabilizationWindowSeconds *int64 `tfsdk:"stabilization_window_seconds" yaml:"stabilizationWindowSeconds,omitempty"`
				} `tfsdk:"scale_up" yaml:"scaleUp,omitempty"`
			} `tfsdk:"behavior" yaml:"behavior,omitempty"`

			TargetCPUUtilization *int64 `tfsdk:"target_cpu_utilization" yaml:"targetCPUUtilization,omitempty"`
		} `tfsdk:"autoscaler" yaml:"autoscaler,omitempty"`

		Config *string `tfsdk:"config" yaml:"config,omitempty"`

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

		HostNetwork *bool `tfsdk:"host_network" yaml:"hostNetwork,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

		Ingress *struct {
			Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Hostname *string `tfsdk:"hostname" yaml:"hostname,omitempty"`

			Tls *[]struct {
				Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"tls" yaml:"tls,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"ingress" yaml:"ingress,omitempty"`

		MaxReplicas *int64 `tfsdk:"max_replicas" yaml:"maxReplicas,omitempty"`

		MinReplicas *int64 `tfsdk:"min_replicas" yaml:"minReplicas,omitempty"`

		Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

		NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		PodAnnotations *map[string]string `tfsdk:"pod_annotations" yaml:"podAnnotations,omitempty"`

		PodSecurityContext *struct {
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
		} `tfsdk:"pod_security_context" yaml:"podSecurityContext,omitempty"`

		Ports *[]struct {
			AppProtocol *string `tfsdk:"app_protocol" yaml:"appProtocol,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			NodePort *int64 `tfsdk:"node_port" yaml:"nodePort,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`

			TargetPort utilities.IntOrString `tfsdk:"target_port" yaml:"targetPort,omitempty"`
		} `tfsdk:"ports" yaml:"ports,omitempty"`

		PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

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

		ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

		TargetAllocator *struct {
			AllocationStrategy *string `tfsdk:"allocation_strategy" yaml:"allocationStrategy,omitempty"`

			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			FilterStrategy *string `tfsdk:"filter_strategy" yaml:"filterStrategy,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			PrometheusCR *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"prometheus_cr" yaml:"prometheusCR,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			ServiceAccount *string `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`
		} `tfsdk:"target_allocator" yaml:"targetAllocator,omitempty"`

		Tolerations *[]struct {
			Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

			TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

		UpgradeStrategy *string `tfsdk:"upgrade_strategy" yaml:"upgradeStrategy,omitempty"`

		VolumeClaimTemplates *[]struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

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

			Status *struct {
				AccessModes *[]string `tfsdk:"access_modes" yaml:"accessModes,omitempty"`

				AllocatedResources *map[string]string `tfsdk:"allocated_resources" yaml:"allocatedResources,omitempty"`

				Capacity *map[string]string `tfsdk:"capacity" yaml:"capacity,omitempty"`

				Conditions *[]struct {
					LastProbeTime *string `tfsdk:"last_probe_time" yaml:"lastProbeTime,omitempty"`

					LastTransitionTime *string `tfsdk:"last_transition_time" yaml:"lastTransitionTime,omitempty"`

					Message *string `tfsdk:"message" yaml:"message,omitempty"`

					Reason *string `tfsdk:"reason" yaml:"reason,omitempty"`

					Status *string `tfsdk:"status" yaml:"status,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"conditions" yaml:"conditions,omitempty"`

				Phase *string `tfsdk:"phase" yaml:"phase,omitempty"`

				ResizeStatus *string `tfsdk:"resize_status" yaml:"resizeStatus,omitempty"`
			} `tfsdk:"status" yaml:"status,omitempty"`
		} `tfsdk:"volume_claim_templates" yaml:"volumeClaimTemplates,omitempty"`

		VolumeMounts *[]struct {
			MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

			MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

			SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

			SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
		} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

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
}

func NewOpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource() resource.Resource {
	return &OpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource{}
}

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_opentelemetry_io_open_telemetry_collector_v1alpha1"
}

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "OpenTelemetryCollector is the Schema for the opentelemetrycollectors API.",
		MarkdownDescription: "OpenTelemetryCollector is the Schema for the opentelemetrycollectors API.",
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
				Description:         "OpenTelemetryCollectorSpec defines the desired state of OpenTelemetryCollector.",
				MarkdownDescription: "OpenTelemetryCollectorSpec defines the desired state of OpenTelemetryCollector.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"affinity": {
						Description:         "If specified, indicates the pod's scheduling constraints",
						MarkdownDescription: "If specified, indicates the pod's scheduling constraints",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

													"namespace_selector": {
														Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
														MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

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
														Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
														MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

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

											"namespace_selector": {
												Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
												MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

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
												Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
												MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

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

													"namespace_selector": {
														Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
														MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

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
														Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
														MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

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

									"required_during_scheduling_ignored_during_execution": {
										Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

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

											"namespace_selector": {
												Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
												MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

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
												Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
												MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

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

					"args": {
						Description:         "Args is the set of arguments to pass to the OpenTelemetry Collector binary",
						MarkdownDescription: "Args is the set of arguments to pass to the OpenTelemetry Collector binary",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"autoscaler": {
						Description:         "Autoscaler specifies the pod autoscaling configuration to use for the OpenTelemetryCollector workload.",
						MarkdownDescription: "Autoscaler specifies the pod autoscaling configuration to use for the OpenTelemetryCollector workload.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"behavior": {
								Description:         "HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).",
								MarkdownDescription: "HorizontalPodAutoscalerBehavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"scale_down": {
										Description:         "scaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e., the highest recommendation for the last 300sec is used).",
										MarkdownDescription: "scaleDown is scaling policy for scaling Down. If not set, the default value is to allow to scale down to minReplicas pods, with a 300 second stabilization window (i.e., the highest recommendation for the last 300sec is used).",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"policies": {
												Description:         "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
												MarkdownDescription: "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"period_seconds": {
														Description:         "PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
														MarkdownDescription: "PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"type": {
														Description:         "Type is used to specify the scaling policy.",
														MarkdownDescription: "Type is used to specify the scaling policy.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Value contains the amount of change which is permitted by the policy. It must be greater than zero",
														MarkdownDescription: "Value contains the amount of change which is permitted by the policy. It must be greater than zero",

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

											"select_policy": {
												Description:         "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
												MarkdownDescription: "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stabilization_window_seconds": {
												Description:         "StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
												MarkdownDescription: "StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",

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

									"scale_up": {
										Description:         "scaleUp is scaling policy for scaling Up. If not set, the default value is the higher of: * increase no more than 4 pods per 60 seconds * double the number of pods per 60 seconds No stabilization is used.",
										MarkdownDescription: "scaleUp is scaling policy for scaling Up. If not set, the default value is the higher of: * increase no more than 4 pods per 60 seconds * double the number of pods per 60 seconds No stabilization is used.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"policies": {
												Description:         "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",
												MarkdownDescription: "policies is a list of potential scaling polices which can be used during scaling. At least one policy must be specified, otherwise the HPAScalingRules will be discarded as invalid",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"period_seconds": {
														Description:         "PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",
														MarkdownDescription: "PeriodSeconds specifies the window of time for which the policy should hold true. PeriodSeconds must be greater than zero and less than or equal to 1800 (30 min).",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"type": {
														Description:         "Type is used to specify the scaling policy.",
														MarkdownDescription: "Type is used to specify the scaling policy.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Value contains the amount of change which is permitted by the policy. It must be greater than zero",
														MarkdownDescription: "Value contains the amount of change which is permitted by the policy. It must be greater than zero",

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

											"select_policy": {
												Description:         "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",
												MarkdownDescription: "selectPolicy is used to specify which policy should be used. If not set, the default value Max is used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"stabilization_window_seconds": {
												Description:         "StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",
												MarkdownDescription: "StabilizationWindowSeconds is the number of seconds for which past recommendations should be considered while scaling up or scaling down. StabilizationWindowSeconds must be greater than or equal to zero and less than or equal to 3600 (one hour). If not set, use the default values: - For scale up: 0 (i.e. no stabilization is done). - For scale down: 300 (i.e. the stabilization window is 300 seconds long).",

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

							"target_cpu_utilization": {
								Description:         "TargetCPUUtilization sets the target average CPU used across all replicas. If average CPU exceeds this value, the HPA will scale up. Defaults to 90 percent.",
								MarkdownDescription: "TargetCPUUtilization sets the target average CPU used across all replicas. If average CPU exceeds this value, the HPA will scale up. Defaults to 90 percent.",

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

					"config": {
						Description:         "Config is the raw JSON to be used as the collector's configuration. Refer to the OpenTelemetry Collector documentation for details.",
						MarkdownDescription: "Config is the raw JSON to be used as the collector's configuration. Refer to the OpenTelemetry Collector documentation for details.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"env": {
						Description:         "ENV vars to set on the OpenTelemetry Collector's Pods. These can then in certain cases be consumed in the config file for the Collector.",
						MarkdownDescription: "ENV vars to set on the OpenTelemetry Collector's Pods. These can then in certain cases be consumed in the config file for the Collector.",

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
						Description:         "List of sources to populate environment variables on the OpenTelemetry Collector's Pods. These can then in certain cases be consumed in the config file for the Collector.",
						MarkdownDescription: "List of sources to populate environment variables on the OpenTelemetry Collector's Pods. These can then in certain cases be consumed in the config file for the Collector.",

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

					"host_network": {
						Description:         "HostNetwork indicates if the pod should run in the host networking namespace.",
						MarkdownDescription: "HostNetwork indicates if the pod should run in the host networking namespace.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "Image indicates the container image to use for the OpenTelemetry Collector.",
						MarkdownDescription: "Image indicates the container image to use for the OpenTelemetry Collector.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_policy": {
						Description:         "ImagePullPolicy indicates the pull policy to be used for retrieving the container image (Always, Never, IfNotPresent)",
						MarkdownDescription: "ImagePullPolicy indicates the pull policy to be used for retrieving the container image (Always, Never, IfNotPresent)",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ingress": {
						Description:         "Ingress is used to specify how OpenTelemetry Collector is exposed. This functionality is only available if one of the valid modes is set. Valid modes are: deployment, daemonset and statefulset.",
						MarkdownDescription: "Ingress is used to specify how OpenTelemetry Collector is exposed. This functionality is only available if one of the valid modes is set. Valid modes are: deployment, daemonset and statefulset.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"annotations": {
								Description:         "Annotations to add to ingress. e.g. 'cert-manager.io/cluster-issuer: 'letsencrypt''",
								MarkdownDescription: "Annotations to add to ingress. e.g. 'cert-manager.io/cluster-issuer: 'letsencrypt''",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hostname": {
								Description:         "Hostname by which the ingress proxy can be reached.",
								MarkdownDescription: "Hostname by which the ingress proxy can be reached.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": {
								Description:         "TLS configuration.",
								MarkdownDescription: "TLS configuration.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"hosts": {
										Description:         "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
										MarkdownDescription: "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
										MarkdownDescription: "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",

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
								Description:         "Type default value is: '' Supported types are: ingress",
								MarkdownDescription: "Type default value is: '' Supported types are: ingress",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("ingress"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"max_replicas": {
						Description:         "MaxReplicas sets an upper bound to the autoscaling feature. If MaxReplicas is set autoscaling is enabled.",
						MarkdownDescription: "MaxReplicas sets an upper bound to the autoscaling feature. If MaxReplicas is set autoscaling is enabled.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"min_replicas": {
						Description:         "MinReplicas sets a lower bound to the autoscaling feature.  Set this if your are using autoscaling. It must be at least 1",
						MarkdownDescription: "MinReplicas sets a lower bound to the autoscaling feature.  Set this if your are using autoscaling. It must be at least 1",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mode": {
						Description:         "Mode represents how the collector should be deployed (deployment, daemonset, statefulset or sidecar)",
						MarkdownDescription: "Mode represents how the collector should be deployed (deployment, daemonset, statefulset or sidecar)",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("daemonset", "deployment", "sidecar", "statefulset"),
						},
					},

					"node_selector": {
						Description:         "NodeSelector to schedule OpenTelemetry Collector pods. This is only relevant to daemonset, statefulset, and deployment mode",
						MarkdownDescription: "NodeSelector to schedule OpenTelemetry Collector pods. This is only relevant to daemonset, statefulset, and deployment mode",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_annotations": {
						Description:         "PodAnnotations is the set of annotations that will be attached to Collector and Target Allocator pods.",
						MarkdownDescription: "PodAnnotations is the set of annotations that will be attached to Collector and Target Allocator pods.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_security_context": {
						Description:         "PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.",
						MarkdownDescription: "PodSecurityContext holds pod-level security attributes and common container settings. Some fields are also present in container.securityContext.  Field values of container.securityContext take precedence over field values of PodSecurityContext.",

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

					"ports": {
						Description:         "Ports allows a set of ports to be exposed by the underlying v1.Service. By default, the operator will attempt to infer the required ports by parsing the .Spec.Config property but this property can be used to open additional ports that can't be inferred by the operator, like for custom receivers.",
						MarkdownDescription: "Ports allows a set of ports to be exposed by the underlying v1.Service. By default, the operator will attempt to infer the required ports by parsing the .Spec.Config property but this property can be used to open additional ports that can't be inferred by the operator, like for custom receivers.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"app_protocol": {
								Description:         "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.",
								MarkdownDescription: "The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
								MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_port": {
								Description:         "The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
								MarkdownDescription: "The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "The port that will be exposed by this service.",
								MarkdownDescription: "The port that will be exposed by this service.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"protocol": {
								Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
								MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target_port": {
								Description:         "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
								MarkdownDescription: "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",

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

					"priority_class_name": {
						Description:         "If specified, indicates the pod's priority. If not specified, the pod priority will be default or zero if there is no default.",
						MarkdownDescription: "If specified, indicates the pod's priority. If not specified, the pod priority will be default or zero if there is no default.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": {
						Description:         "Replicas is the number of pod instances for the underlying OpenTelemetry Collector. Set this if your are not using autoscaling",
						MarkdownDescription: "Replicas is the number of pod instances for the underlying OpenTelemetry Collector. Set this if your are not using autoscaling",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "Resources to set on the OpenTelemetry Collector pods.",
						MarkdownDescription: "Resources to set on the OpenTelemetry Collector pods.",

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
						Description:         "SecurityContext will be set as the container security context.",
						MarkdownDescription: "SecurityContext will be set as the container security context.",

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

					"service_account": {
						Description:         "ServiceAccount indicates the name of an existing service account to use with this instance.",
						MarkdownDescription: "ServiceAccount indicates the name of an existing service account to use with this instance.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"target_allocator": {
						Description:         "TargetAllocator indicates a value which determines whether to spawn a target allocation resource or not.",
						MarkdownDescription: "TargetAllocator indicates a value which determines whether to spawn a target allocation resource or not.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"allocation_strategy": {
								Description:         "AllocationStrategy determines which strategy the target allocator should use for allocation. The current options are least-weighted and consistent-hashing. The default option is least-weighted",
								MarkdownDescription: "AllocationStrategy determines which strategy the target allocator should use for allocation. The current options are least-weighted and consistent-hashing. The default option is least-weighted",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": {
								Description:         "Enabled indicates whether to use a target allocation mechanism for Prometheus targets or not.",
								MarkdownDescription: "Enabled indicates whether to use a target allocation mechanism for Prometheus targets or not.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"filter_strategy": {
								Description:         "FilterStrategy determines how to filter targets before allocating them among the collectors. The only current option is relabel-config (drops targets based on prom relabel_config). Filtering is disabled by default.",
								MarkdownDescription: "FilterStrategy determines how to filter targets before allocating them among the collectors. The only current option is relabel-config (drops targets based on prom relabel_config). Filtering is disabled by default.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "Image indicates the container image to use for the OpenTelemetry TargetAllocator.",
								MarkdownDescription: "Image indicates the container image to use for the OpenTelemetry TargetAllocator.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"prometheus_cr": {
								Description:         "PrometheusCR defines the configuration for the retrieval of PrometheusOperator CRDs ( servicemonitor.monitoring.coreos.com/v1 and podmonitor.monitoring.coreos.com/v1 )  retrieval. All CR instances which the ServiceAccount has access to will be retrieved. This includes other namespaces.",
								MarkdownDescription: "PrometheusCR defines the configuration for the retrieval of PrometheusOperator CRDs ( servicemonitor.monitoring.coreos.com/v1 and podmonitor.monitoring.coreos.com/v1 )  retrieval. All CR instances which the ServiceAccount has access to will be retrieved. This includes other namespaces.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Enabled indicates whether to use a PrometheusOperator custom resources as targets or not.",
										MarkdownDescription: "Enabled indicates whether to use a PrometheusOperator custom resources as targets or not.",

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

							"replicas": {
								Description:         "Replicas is the number of pod instances for the underlying TargetAllocator. This should only be set to a value other than 1 if a strategy that allows for high availability is chosen. Currently, the only allocation strategy that can be run in a high availability mode is consistent-hashing.",
								MarkdownDescription: "Replicas is the number of pod instances for the underlying TargetAllocator. This should only be set to a value other than 1 if a strategy that allows for high availability is chosen. Currently, the only allocation strategy that can be run in a high availability mode is consistent-hashing.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_account": {
								Description:         "ServiceAccount indicates the name of an existing service account to use with this instance.",
								MarkdownDescription: "ServiceAccount indicates the name of an existing service account to use with this instance.",

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

					"tolerations": {
						Description:         "Toleration to schedule OpenTelemetry Collector pods. This is only relevant to daemonset, statefulset, and deployment mode",
						MarkdownDescription: "Toleration to schedule OpenTelemetry Collector pods. This is only relevant to daemonset, statefulset, and deployment mode",

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

					"upgrade_strategy": {
						Description:         "UpgradeStrategy represents how the operator will handle upgrades to the CR when a newer version of the operator is deployed",
						MarkdownDescription: "UpgradeStrategy represents how the operator will handle upgrades to the CR when a newer version of the operator is deployed",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("automatic", "none"),
						},
					},

					"volume_claim_templates": {
						Description:         "VolumeClaimTemplates will provide stable storage using PersistentVolumes. Only available when the mode=statefulset.",
						MarkdownDescription: "VolumeClaimTemplates will provide stable storage using PersistentVolumes. Only available when the mode=statefulset.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
								MarkdownDescription: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metadata": {
								Description:         "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
								MarkdownDescription: "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",

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
								Description:         "spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
								MarkdownDescription: "spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"access_modes": {
										Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"data_source": {
										Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
										MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",

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
										Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
										MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",

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
										Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",

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
										Description:         "selector is a label query over volumes to consider for binding.",
										MarkdownDescription: "selector is a label query over volumes to consider for binding.",

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
										Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",

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
										Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
										MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",

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

							"status": {
								Description:         "status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
								MarkdownDescription: "status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"access_modes": {
										Description:         "accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										MarkdownDescription: "accessModes contains the actual access modes the volume backing the PVC has. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allocated_resources": {
										Description:         "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
										MarkdownDescription: "allocatedResources is the storage resource within AllocatedResources tracks the capacity allocated to a PVC. It may be larger than the actual capacity when a volume expansion operation is requested. For storage quota, the larger value from allocatedResources and PVC.spec.resources is used. If allocatedResources is not set, PVC.spec.resources alone is used for quota calculation. If a volume expansion capacity request is lowered, allocatedResources is only lowered if there are no expansion operations in progress and if the actual volume capacity is equal or lower than the requested capacity. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"capacity": {
										Description:         "capacity represents the actual resources of the underlying volume.",
										MarkdownDescription: "capacity represents the actual resources of the underlying volume.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"conditions": {
										Description:         "conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.",
										MarkdownDescription: "conditions is the current Condition of persistent volume claim. If underlying persistent volume is being resized then the Condition will be set to 'ResizeStarted'.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"last_probe_time": {
												Description:         "lastProbeTime is the time we probed the condition.",
												MarkdownDescription: "lastProbeTime is the time we probed the condition.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.DateTime64Validator(),
												},
											},

											"last_transition_time": {
												Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
												MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													validators.DateTime64Validator(),
												},
											},

											"message": {
												Description:         "message is the human-readable message indicating details about last transition.",
												MarkdownDescription: "message is the human-readable message indicating details about last transition.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"reason": {
												Description:         "reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'ResizeStarted' that means the underlying persistent volume is being resized.",
												MarkdownDescription: "reason is a unique, this should be a short, machine understandable string that gives the reason for condition's last transition. If it reports 'ResizeStarted' that means the underlying persistent volume is being resized.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"status": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",
												MarkdownDescription: "PersistentVolumeClaimConditionType is a valid value of PersistentVolumeClaimCondition.Type",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"phase": {
										Description:         "phase represents the current phase of PersistentVolumeClaim.",
										MarkdownDescription: "phase represents the current phase of PersistentVolumeClaim.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resize_status": {
										Description:         "resizeStatus stores status of resize operation. ResizeStatus is not set by default but when expansion is complete resizeStatus is set to empty string by resize controller or kubelet. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",
										MarkdownDescription: "resizeStatus stores status of resize operation. ResizeStatus is not set by default but when expansion is complete resizeStatus is set to empty string by resize controller or kubelet. This is an alpha field and requires enabling RecoverVolumeExpansionFailure feature.",

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

					"volume_mounts": {
						Description:         "VolumeMounts represents the mount points to use in the underlying collector deployment(s)",
						MarkdownDescription: "VolumeMounts represents the mount points to use in the underlying collector deployment(s)",

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

					"volumes": {
						Description:         "Volumes represents which volumes to use in the underlying collector deployment(s).",
						MarkdownDescription: "Volumes represents which volumes to use in the underlying collector deployment(s).",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"aws_elastic_block_store": {
								Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
								MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
										MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"partition": {
										Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
										MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
										MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_id": {
										Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
										MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

										Type: types.StringType,

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
								Description:         "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
								MarkdownDescription: "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"caching_mode": {
										Description:         "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
										MarkdownDescription: "cachingMode is the Host Caching mode: None, Read Only, Read Write.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disk_name": {
										Description:         "diskName is the Name of the data disk in the blob storage",
										MarkdownDescription: "diskName is the Name of the data disk in the blob storage",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"disk_uri": {
										Description:         "diskURI is the URI of data disk in the blob storage",
										MarkdownDescription: "diskURI is the URI of data disk in the blob storage",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"fs_type": {
										Description:         "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
										MarkdownDescription: "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
										MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
										MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

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
								Description:         "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
								MarkdownDescription: "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"read_only": {
										Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
										MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "secretName is the  name of secret that contains Azure Storage Account Name and Key",
										MarkdownDescription: "secretName is the  name of secret that contains Azure Storage Account Name and Key",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"share_name": {
										Description:         "shareName is the azure share Name",
										MarkdownDescription: "shareName is the azure share Name",

										Type: types.StringType,

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
								Description:         "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
								MarkdownDescription: "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"monitors": {
										Description:         "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
										MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"path": {
										Description:         "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
										MarkdownDescription: "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
										MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_file": {
										Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
										MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
										MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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
										Description:         "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
										MarkdownDescription: "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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
								Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
								MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
										MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
										MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
										MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",

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
										Description:         "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
										MarkdownDescription: "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

										Type: types.StringType,

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
								Description:         "configMap represents a configMap that should populate this volume",
								MarkdownDescription: "configMap represents a configMap that should populate this volume",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"default_mode": {
										Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
										MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"items": {
										Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
										MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the key to project.",
												MarkdownDescription: "key is the key to project.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"mode": {
												Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
												MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
												MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

												Type: types.StringType,

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
										Description:         "optional specify whether the ConfigMap or its keys must be defined",
										MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",

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
								Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
								MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"driver": {
										Description:         "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
										MarkdownDescription: "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"fs_type": {
										Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
										MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"node_publish_secret_ref": {
										Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
										MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",

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
										Description:         "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
										MarkdownDescription: "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_attributes": {
										Description:         "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
										MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",

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
								Description:         "downwardAPI represents downward API about the pod that should populate this volume",
								MarkdownDescription: "downwardAPI represents downward API about the pod that should populate this volume",

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
								Description:         "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
								MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"medium": {
										Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
										MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"size_limit": {
										Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
										MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",

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
								Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
								MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"volume_claim_template": {
										Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
										MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
												MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",

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
												Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
												MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_modes": {
														Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
														MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"data_source": {
														Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
														MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",

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
														Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
														MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",

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
														Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
														MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",

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
														Description:         "selector is a label query over volumes to consider for binding.",
														MarkdownDescription: "selector is a label query over volumes to consider for binding.",

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
														Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
														MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",

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
														Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
														MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",

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
								Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
								MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
										MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"lun": {
										Description:         "lun is Optional: FC target lun number",
										MarkdownDescription: "lun is Optional: FC target lun number",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
										MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_ww_ns": {
										Description:         "targetWWNs is Optional: FC target worldwide names (WWNs)",
										MarkdownDescription: "targetWWNs is Optional: FC target worldwide names (WWNs)",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"wwids": {
										Description:         "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
										MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",

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
								Description:         "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
								MarkdownDescription: "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"driver": {
										Description:         "driver is the name of the driver to use for this volume.",
										MarkdownDescription: "driver is the name of the driver to use for this volume.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"fs_type": {
										Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
										MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"options": {
										Description:         "options is Optional: this field holds extra command options if any.",
										MarkdownDescription: "options is Optional: this field holds extra command options if any.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
										MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
										MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",

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
								Description:         "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
								MarkdownDescription: "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"dataset_name": {
										Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
										MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"dataset_uuid": {
										Description:         "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
										MarkdownDescription: "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",

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
								Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
								MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
										MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"partition": {
										Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
										MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pd_name": {
										Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
										MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
										MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

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
								Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
								MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"directory": {
										Description:         "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
										MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"repository": {
										Description:         "repository is the URL",
										MarkdownDescription: "repository is the URL",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"revision": {
										Description:         "revision is the commit hash for the specified revision.",
										MarkdownDescription: "revision is the commit hash for the specified revision.",

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
								Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
								MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"endpoints": {
										Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
										MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"path": {
										Description:         "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
										MarkdownDescription: "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
										MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

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
								Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
								MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"path": {
										Description:         "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
										MarkdownDescription: "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"type": {
										Description:         "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
										MarkdownDescription: "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

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
								Description:         "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
								MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"chap_auth_discovery": {
										Description:         "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
										MarkdownDescription: "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"chap_auth_session": {
										Description:         "chapAuthSession defines whether support iSCSI Session CHAP authentication",
										MarkdownDescription: "chapAuthSession defines whether support iSCSI Session CHAP authentication",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"fs_type": {
										Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
										MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"initiator_name": {
										Description:         "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
										MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"iqn": {
										Description:         "iqn is the target iSCSI Qualified Name.",
										MarkdownDescription: "iqn is the target iSCSI Qualified Name.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"iscsi_interface": {
										Description:         "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
										MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"lun": {
										Description:         "lun represents iSCSI Target Lun number.",
										MarkdownDescription: "lun represents iSCSI Target Lun number.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"portals": {
										Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
										MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
										MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
										MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",

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
										Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
										MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

										Type: types.StringType,

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
								Description:         "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"nfs": {
								Description:         "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
								MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"path": {
										Description:         "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
										MarkdownDescription: "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
										MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"server": {
										Description:         "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
										MarkdownDescription: "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

										Type: types.StringType,

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
								Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
								MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"claim_name": {
										Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
										MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
										MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",

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
								Description:         "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
								MarkdownDescription: "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
										MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pd_id": {
										Description:         "pdID is the ID that identifies Photon Controller persistent disk",
										MarkdownDescription: "pdID is the ID that identifies Photon Controller persistent disk",

										Type: types.StringType,

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
								Description:         "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
								MarkdownDescription: "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
										MarkdownDescription: "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
										MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_id": {
										Description:         "volumeID uniquely identifies a Portworx volume",
										MarkdownDescription: "volumeID uniquely identifies a Portworx volume",

										Type: types.StringType,

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
								Description:         "projected items for all in one resources secrets, configmaps, and downward API",
								MarkdownDescription: "projected items for all in one resources secrets, configmaps, and downward API",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"default_mode": {
										Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
										MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sources": {
										Description:         "sources is the list of volume projections",
										MarkdownDescription: "sources is the list of volume projections",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "configMap information about the configMap data to project",
												MarkdownDescription: "configMap information about the configMap data to project",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"items": {
														Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mode": {
																Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

																Type: types.StringType,

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
														Description:         "optional specify whether the ConfigMap or its keys must be defined",
														MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",

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
												Description:         "downwardAPI information about the downwardAPI data to project",
												MarkdownDescription: "downwardAPI information about the downwardAPI data to project",

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
												Description:         "secret information about the secret data to project",
												MarkdownDescription: "secret information about the secret data to project",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"items": {
														Description:         "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mode": {
																Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

																Type: types.StringType,

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
														Description:         "optional field specify whether the Secret or its key must be defined",
														MarkdownDescription: "optional field specify whether the Secret or its key must be defined",

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
												Description:         "serviceAccountToken is information about the serviceAccountToken data to project",
												MarkdownDescription: "serviceAccountToken is information about the serviceAccountToken data to project",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"audience": {
														Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
														MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expiration_seconds": {
														Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
														MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "path is the path relative to the mount point of the file to project the token into.",
														MarkdownDescription: "path is the path relative to the mount point of the file to project the token into.",

														Type: types.StringType,

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
								Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
								MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"group": {
										Description:         "group to map volume access to Default is no group",
										MarkdownDescription: "group to map volume access to Default is no group",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
										MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"registry": {
										Description:         "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
										MarkdownDescription: "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"tenant": {
										Description:         "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
										MarkdownDescription: "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user": {
										Description:         "user to map volume access to Defaults to serivceaccount user",
										MarkdownDescription: "user to map volume access to Defaults to serivceaccount user",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume": {
										Description:         "volume is a string that references an already created Quobyte volume by name.",
										MarkdownDescription: "volume is a string that references an already created Quobyte volume by name.",

										Type: types.StringType,

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
								Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
								MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
										MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
										MarkdownDescription: "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"keyring": {
										Description:         "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
										MarkdownDescription: "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"monitors": {
										Description:         "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
										MarkdownDescription: "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"pool": {
										Description:         "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
										MarkdownDescription: "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
										MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
										MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

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
										Description:         "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
										MarkdownDescription: "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

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
								Description:         "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
								MarkdownDescription: "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
										MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"gateway": {
										Description:         "gateway is the host address of the ScaleIO API Gateway.",
										MarkdownDescription: "gateway is the host address of the ScaleIO API Gateway.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"protection_domain": {
										Description:         "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
										MarkdownDescription: "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
										MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
										MarkdownDescription: "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",

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
										Description:         "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
										MarkdownDescription: "sslEnabled Flag enable/disable SSL communication with Gateway, default false",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_mode": {
										Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
										MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_pool": {
										Description:         "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
										MarkdownDescription: "storagePool is the ScaleIO Storage Pool associated with the protection domain.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"system": {
										Description:         "system is the name of the storage system as configured in ScaleIO.",
										MarkdownDescription: "system is the name of the storage system as configured in ScaleIO.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"volume_name": {
										Description:         "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
										MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",

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
								Description:         "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
								MarkdownDescription: "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"default_mode": {
										Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
										MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"items": {
										Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
										MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"key": {
												Description:         "key is the key to project.",
												MarkdownDescription: "key is the key to project.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"mode": {
												Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
												MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
												MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

												Type: types.StringType,

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
										Description:         "optional field specify whether the Secret or its keys must be defined",
										MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
										MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

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
								Description:         "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
								MarkdownDescription: "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
										MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_only": {
										Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
										MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": {
										Description:         "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
										MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",

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
										Description:         "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
										MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_namespace": {
										Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
										MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",

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
								Description:         "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
								MarkdownDescription: "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_type": {
										Description:         "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
										MarkdownDescription: "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_policy_id": {
										Description:         "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
										MarkdownDescription: "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage_policy_name": {
										Description:         "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
										MarkdownDescription: "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_path": {
										Description:         "volumePath is the path that identifies vSphere volume vmdk",
										MarkdownDescription: "volumePath is the path that identifies vSphere volume vmdk",

										Type: types.StringType,

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

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_opentelemetry_io_open_telemetry_collector_v1alpha1")

	var state OpentelemetryIoOpenTelemetryCollectorV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OpentelemetryIoOpenTelemetryCollectorV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("opentelemetry.io/v1alpha1")
	goModel.Kind = utilities.Ptr("OpenTelemetryCollector")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_opentelemetry_io_open_telemetry_collector_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_opentelemetry_io_open_telemetry_collector_v1alpha1")

	var state OpentelemetryIoOpenTelemetryCollectorV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel OpentelemetryIoOpenTelemetryCollectorV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("opentelemetry.io/v1alpha1")
	goModel.Kind = utilities.Ptr("OpenTelemetryCollector")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *OpentelemetryIoOpenTelemetryCollectorV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_opentelemetry_io_open_telemetry_collector_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
