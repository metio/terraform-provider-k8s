/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kafka_strimzi_io_v1beta2

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ datasource.DataSource = &KafkaStrimziIoKafkaNodePoolV1Beta2Manifest{}
)

func NewKafkaStrimziIoKafkaNodePoolV1Beta2Manifest() datasource.DataSource {
	return &KafkaStrimziIoKafkaNodePoolV1Beta2Manifest{}
}

type KafkaStrimziIoKafkaNodePoolV1Beta2Manifest struct{}

type KafkaStrimziIoKafkaNodePoolV1Beta2ManifestData struct {
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
		JvmOptions *struct {
			_XX                  *map[string]string `tfsdk:"xx" json:"-XX,omitempty"`
			_Xms                 *string            `tfsdk:"xms" json:"-Xms,omitempty"`
			_Xmx                 *string            `tfsdk:"xmx" json:"-Xmx,omitempty"`
			GcLoggingEnabled     *bool              `tfsdk:"gc_logging_enabled" json:"gcLoggingEnabled,omitempty"`
			JavaSystemProperties *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"java_system_properties" json:"javaSystemProperties,omitempty"`
		} `tfsdk:"jvm_options" json:"jvmOptions,omitempty"`
		Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		Roles   *[]string `tfsdk:"roles" json:"roles,omitempty"`
		Storage *struct {
			Class         *string `tfsdk:"class" json:"class,omitempty"`
			DeleteClaim   *bool   `tfsdk:"delete_claim" json:"deleteClaim,omitempty"`
			Id            *int64  `tfsdk:"id" json:"id,omitempty"`
			KraftMetadata *string `tfsdk:"kraft_metadata" json:"kraftMetadata,omitempty"`
			Overrides     *[]struct {
				Broker *int64  `tfsdk:"broker" json:"broker,omitempty"`
				Class  *string `tfsdk:"class" json:"class,omitempty"`
			} `tfsdk:"overrides" json:"overrides,omitempty"`
			Selector  *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
			Size      *string            `tfsdk:"size" json:"size,omitempty"`
			SizeLimit *string            `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
			Type      *string            `tfsdk:"type" json:"type,omitempty"`
			Volumes   *[]struct {
				Class         *string `tfsdk:"class" json:"class,omitempty"`
				DeleteClaim   *bool   `tfsdk:"delete_claim" json:"deleteClaim,omitempty"`
				Id            *int64  `tfsdk:"id" json:"id,omitempty"`
				KraftMetadata *string `tfsdk:"kraft_metadata" json:"kraftMetadata,omitempty"`
				Overrides     *[]struct {
					Broker *int64  `tfsdk:"broker" json:"broker,omitempty"`
					Class  *string `tfsdk:"class" json:"class,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				Selector  *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
				Size      *string            `tfsdk:"size" json:"size,omitempty"`
				SizeLimit *string            `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				Type      *string            `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		Template *struct {
			InitContainer *struct {
				Env *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				SecurityContext *struct {
					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
					AppArmorProfile          *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
						Type             *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"app_armor_profile" json:"appArmorProfile,omitempty"`
					Capabilities *struct {
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
				VolumeMounts *[]struct {
					MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
					Name              *string `tfsdk:"name" json:"name,omitempty"`
					ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
					SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
				} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			} `tfsdk:"init_container" json:"initContainer,omitempty"`
			KafkaContainer *struct {
				Env *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				SecurityContext *struct {
					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
					AppArmorProfile          *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
						Type             *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"app_armor_profile" json:"appArmorProfile,omitempty"`
					Capabilities *struct {
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
				VolumeMounts *[]struct {
					MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
					Name              *string `tfsdk:"name" json:"name,omitempty"`
					ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
					SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
				} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			} `tfsdk:"kafka_container" json:"kafkaContainer,omitempty"`
			PerPodIngress *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
			} `tfsdk:"per_pod_ingress" json:"perPodIngress,omitempty"`
			PerPodRoute *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
			} `tfsdk:"per_pod_route" json:"perPodRoute,omitempty"`
			PerPodService *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
			} `tfsdk:"per_pod_service" json:"perPodService,omitempty"`
			PersistentVolumeClaim *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
			} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
			Pod *struct {
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
				EnableServiceLinks *bool `tfsdk:"enable_service_links" json:"enableServiceLinks,omitempty"`
				HostAliases        *[]struct {
					Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
					Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
				} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
				ImagePullSecrets *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
				SchedulerName     *string `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
				SecurityContext   *struct {
					AppArmorProfile *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
						Type             *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"app_armor_profile" json:"appArmorProfile,omitempty"`
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
				TerminationGracePeriodSeconds *int64  `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
				TmpDirSizeLimit               *string `tfsdk:"tmp_dir_size_limit" json:"tmpDirSizeLimit,omitempty"`
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
					EmptyDir *struct {
						Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
						SizeLimit *struct {
							Amount *string `tfsdk:"amount" json:"amount,omitempty"`
							Format *string `tfsdk:"format" json:"format,omitempty"`
						} `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
					} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
					Name                  *string `tfsdk:"name" json:"name,omitempty"`
					PersistentVolumeClaim *struct {
						ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
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
				} `tfsdk:"volumes" json:"volumes,omitempty"`
			} `tfsdk:"pod" json:"pod,omitempty"`
			PodSet *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
			} `tfsdk:"pod_set" json:"podSet,omitempty"`
		} `tfsdk:"template" json:"template,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KafkaStrimziIoKafkaNodePoolV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_strimzi_io_kafka_node_pool_v1beta2_manifest"
}

func (r *KafkaStrimziIoKafkaNodePoolV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "The specification of the KafkaNodePool.",
				MarkdownDescription: "The specification of the KafkaNodePool.",
				Attributes: map[string]schema.Attribute{
					"jvm_options": schema.SingleNestedAttribute{
						Description:         "JVM Options for pods.",
						MarkdownDescription: "JVM Options for pods.",
						Attributes: map[string]schema.Attribute{
							"xx": schema.MapAttribute{
								Description:         "A map of -XX options to the JVM.",
								MarkdownDescription: "A map of -XX options to the JVM.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"xms": schema.StringAttribute{
								Description:         "-Xms option to to the JVM.",
								MarkdownDescription: "-Xms option to to the JVM.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
								},
							},

							"xmx": schema.StringAttribute{
								Description:         "-Xmx option to to the JVM.",
								MarkdownDescription: "-Xmx option to to the JVM.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
								},
							},

							"gc_logging_enabled": schema.BoolAttribute{
								Description:         "Specifies whether the Garbage Collection logging is enabled. The default is false.",
								MarkdownDescription: "Specifies whether the Garbage Collection logging is enabled. The default is false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"java_system_properties": schema.ListNestedAttribute{
								Description:         "A map of additional system properties which will be passed using the '-D' option to the JVM.",
								MarkdownDescription: "A map of additional system properties which will be passed using the '-D' option to the JVM.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "The system property name.",
											MarkdownDescription: "The system property name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "The system property value.",
											MarkdownDescription: "The system property value.",
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

					"replicas": schema.Int64Attribute{
						Description:         "The number of pods in the pool.",
						MarkdownDescription: "The number of pods in the pool.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "CPU and memory resources to reserve.",
						MarkdownDescription: "CPU and memory resources to reserve.",
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
											Optional:            true,
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

					"roles": schema.ListAttribute{
						Description:         "The roles that the nodes in this pool will have when KRaft mode is enabled. Supported values are 'broker' and 'controller'. This field is required. When KRaft mode is disabled, the only allowed value if 'broker'.",
						MarkdownDescription: "The roles that the nodes in this pool will have when KRaft mode is enabled. Supported values are 'broker' and 'controller'. This field is required. When KRaft mode is disabled, the only allowed value if 'broker'.",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage configuration (disk). Cannot be updated.",
						MarkdownDescription: "Storage configuration (disk). Cannot be updated.",
						Attributes: map[string]schema.Attribute{
							"class": schema.StringAttribute{
								Description:         "The storage class to use for dynamic volume allocation.",
								MarkdownDescription: "The storage class to use for dynamic volume allocation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delete_claim": schema.BoolAttribute{
								Description:         "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",
								MarkdownDescription: "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"id": schema.Int64Attribute{
								Description:         "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",
								MarkdownDescription: "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"kraft_metadata": schema.StringAttribute{
								Description:         "Specifies whether this volume should be used for storing KRaft metadata. This property is optional. When set, the only currently supported value is 'shared'. At most one volume can have this property set.",
								MarkdownDescription: "Specifies whether this volume should be used for storing KRaft metadata. This property is optional. When set, the only currently supported value is 'shared'. At most one volume can have this property set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("shared"),
								},
							},

							"overrides": schema.ListNestedAttribute{
								Description:         "Overrides for individual brokers. The 'overrides' field allows you to specify a different configuration for different brokers.",
								MarkdownDescription: "Overrides for individual brokers. The 'overrides' field allows you to specify a different configuration for different brokers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"broker": schema.Int64Attribute{
											Description:         "Id of the kafka broker (broker identifier).",
											MarkdownDescription: "Id of the kafka broker (broker identifier).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"class": schema.StringAttribute{
											Description:         "The storage class to use for dynamic volume allocation for this broker.",
											MarkdownDescription: "The storage class to use for dynamic volume allocation for this broker.",
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

							"selector": schema.MapAttribute{
								Description:         "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",
								MarkdownDescription: "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "When 'type=persistent-claim', defines the size of the persistent volume claim, such as 100Gi. Mandatory when 'type=persistent-claim'.",
								MarkdownDescription: "When 'type=persistent-claim', defines the size of the persistent volume claim, such as 100Gi. Mandatory when 'type=persistent-claim'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size_limit": schema.StringAttribute{
								Description:         "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",
								MarkdownDescription: "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
								},
							},

							"type": schema.StringAttribute{
								Description:         "Storage type, must be either 'ephemeral', 'persistent-claim', or 'jbod'.",
								MarkdownDescription: "Storage type, must be either 'ephemeral', 'persistent-claim', or 'jbod'.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ephemeral", "persistent-claim", "jbod"),
								},
							},

							"volumes": schema.ListNestedAttribute{
								Description:         "List of volumes as Storage objects representing the JBOD disks array.",
								MarkdownDescription: "List of volumes as Storage objects representing the JBOD disks array.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"class": schema.StringAttribute{
											Description:         "The storage class to use for dynamic volume allocation.",
											MarkdownDescription: "The storage class to use for dynamic volume allocation.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"delete_claim": schema.BoolAttribute{
											Description:         "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",
											MarkdownDescription: "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"id": schema.Int64Attribute{
											Description:         "Storage identification number. Mandatory for storage volumes defined with a 'jbod' storage type configuration.",
											MarkdownDescription: "Storage identification number. Mandatory for storage volumes defined with a 'jbod' storage type configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"kraft_metadata": schema.StringAttribute{
											Description:         "Specifies whether this volume should be used for storing KRaft metadata. This property is optional. When set, the only currently supported value is 'shared'. At most one volume can have this property set.",
											MarkdownDescription: "Specifies whether this volume should be used for storing KRaft metadata. This property is optional. When set, the only currently supported value is 'shared'. At most one volume can have this property set.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("shared"),
											},
										},

										"overrides": schema.ListNestedAttribute{
											Description:         "Overrides for individual brokers. The 'overrides' field allows you to specify a different configuration for different brokers.",
											MarkdownDescription: "Overrides for individual brokers. The 'overrides' field allows you to specify a different configuration for different brokers.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"broker": schema.Int64Attribute{
														Description:         "Id of the kafka broker (broker identifier).",
														MarkdownDescription: "Id of the kafka broker (broker identifier).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"class": schema.StringAttribute{
														Description:         "The storage class to use for dynamic volume allocation for this broker.",
														MarkdownDescription: "The storage class to use for dynamic volume allocation for this broker.",
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

										"selector": schema.MapAttribute{
											Description:         "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",
											MarkdownDescription: "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size": schema.StringAttribute{
											Description:         "When 'type=persistent-claim', defines the size of the persistent volume claim, such as 100Gi. Mandatory when 'type=persistent-claim'.",
											MarkdownDescription: "When 'type=persistent-claim', defines the size of the persistent volume claim, such as 100Gi. Mandatory when 'type=persistent-claim'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size_limit": schema.StringAttribute{
											Description:         "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",
											MarkdownDescription: "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
											},
										},

										"type": schema.StringAttribute{
											Description:         "Storage type, must be either 'ephemeral' or 'persistent-claim'.",
											MarkdownDescription: "Storage type, must be either 'ephemeral' or 'persistent-claim'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("ephemeral", "persistent-claim"),
											},
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

					"template": schema.SingleNestedAttribute{
						Description:         "Template for pool resources. The template allows users to specify how the resources belonging to this pool are generated.",
						MarkdownDescription: "Template for pool resources. The template allows users to specify how the resources belonging to this pool are generated.",
						Attributes: map[string]schema.Attribute{
							"init_container": schema.SingleNestedAttribute{
								Description:         "Template for the Kafka init container.",
								MarkdownDescription: "Template for the Kafka init container.",
								Attributes: map[string]schema.Attribute{
									"env": schema.ListNestedAttribute{
										Description:         "Environment variables which should be applied to the container.",
										MarkdownDescription: "Environment variables which should be applied to the container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The environment variable key.",
													MarkdownDescription: "The environment variable key.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "The environment variable value.",
													MarkdownDescription: "The environment variable value.",
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

									"security_context": schema.SingleNestedAttribute{
										Description:         "Security context for the container.",
										MarkdownDescription: "Security context for the container.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"app_armor_profile": schema.SingleNestedAttribute{
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
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
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
														Required:            false,
														Optional:            true,
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

									"volume_mounts": schema.ListNestedAttribute{
										Description:         "Additional volume mounts which should be applied to the container.",
										MarkdownDescription: "Additional volume mounts which should be applied to the container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"mount_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
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

												"recursive_read_only": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"kafka_container": schema.SingleNestedAttribute{
								Description:         "Template for the Kafka broker container.",
								MarkdownDescription: "Template for the Kafka broker container.",
								Attributes: map[string]schema.Attribute{
									"env": schema.ListNestedAttribute{
										Description:         "Environment variables which should be applied to the container.",
										MarkdownDescription: "Environment variables which should be applied to the container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The environment variable key.",
													MarkdownDescription: "The environment variable key.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "The environment variable value.",
													MarkdownDescription: "The environment variable value.",
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

									"security_context": schema.SingleNestedAttribute{
										Description:         "Security context for the container.",
										MarkdownDescription: "Security context for the container.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"app_armor_profile": schema.SingleNestedAttribute{
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
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
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
														Required:            false,
														Optional:            true,
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

									"volume_mounts": schema.ListNestedAttribute{
										Description:         "Additional volume mounts which should be applied to the container.",
										MarkdownDescription: "Additional volume mounts which should be applied to the container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"mount_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
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

												"recursive_read_only": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"per_pod_ingress": schema.SingleNestedAttribute{
								Description:         "Template for Kafka per-pod 'Ingress' used for access from outside of Kubernetes.",
								MarkdownDescription: "Template for Kafka per-pod 'Ingress' used for access from outside of Kubernetes.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations added to the Kubernetes resource.",
												MarkdownDescription: "Annotations added to the Kubernetes resource.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"per_pod_route": schema.SingleNestedAttribute{
								Description:         "Template for Kafka per-pod 'Routes' used for access from outside of OpenShift.",
								MarkdownDescription: "Template for Kafka per-pod 'Routes' used for access from outside of OpenShift.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations added to the Kubernetes resource.",
												MarkdownDescription: "Annotations added to the Kubernetes resource.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"per_pod_service": schema.SingleNestedAttribute{
								Description:         "Template for Kafka per-pod 'Services' used for access from outside of Kubernetes.",
								MarkdownDescription: "Template for Kafka per-pod 'Services' used for access from outside of Kubernetes.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations added to the Kubernetes resource.",
												MarkdownDescription: "Annotations added to the Kubernetes resource.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"persistent_volume_claim": schema.SingleNestedAttribute{
								Description:         "Template for all Kafka 'PersistentVolumeClaims'.",
								MarkdownDescription: "Template for all Kafka 'PersistentVolumeClaims'.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations added to the Kubernetes resource.",
												MarkdownDescription: "Annotations added to the Kubernetes resource.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod": schema.SingleNestedAttribute{
								Description:         "Template for Kafka 'Pods'.",
								MarkdownDescription: "Template for Kafka 'Pods'.",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "The pod's affinity rules.",
										MarkdownDescription: "The pod's affinity rules.",
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
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"weight": schema.Int64Attribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"weight": schema.Int64Attribute{
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
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"weight": schema.Int64Attribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable_service_links": schema.BoolAttribute{
										Description:         "Indicates whether information about services should be injected into Pod's environment variables.",
										MarkdownDescription: "Indicates whether information about services should be injected into Pod's environment variables.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_aliases": schema.ListNestedAttribute{
										Description:         "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",
										MarkdownDescription: "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",
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

									"image_pull_secrets": schema.ListNestedAttribute{
										Description:         "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
										MarkdownDescription: "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
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
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations added to the Kubernetes resource.",
												MarkdownDescription: "Annotations added to the Kubernetes resource.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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

									"priority_class_name": schema.StringAttribute{
										Description:         "The name of the priority class used to assign priority to the pods. ",
										MarkdownDescription: "The name of the priority class used to assign priority to the pods. ",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheduler_name": schema.StringAttribute{
										Description:         "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
										MarkdownDescription: "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"security_context": schema.SingleNestedAttribute{
										Description:         "Configures pod-level security attributes and common container settings.",
										MarkdownDescription: "Configures pod-level security attributes and common container settings.",
										Attributes: map[string]schema.Attribute{
											"app_armor_profile": schema.SingleNestedAttribute{
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
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

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
														Required:            false,
														Optional:            true,
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

									"termination_grace_period_seconds": schema.Int64Attribute{
										Description:         "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
										MarkdownDescription: "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"tmp_dir_size_limit": schema.StringAttribute{
										Description:         "Defines the total amount of pod memory allocated for the temporary 'EmptyDir' volume '/tmp'. Specify the allocation in memory units, for example, '100Mi' for 100 mebibytes. Default value is '5Mi'. The '/tmp' volume is backed by pod memory, not disk storage, so avoid setting a high value as it consumes pod memory resources.",
										MarkdownDescription: "Defines the total amount of pod memory allocated for the temporary 'EmptyDir' volume '/tmp'. Specify the allocation in memory units, for example, '100Mi' for 100 mebibytes. Default value is '5Mi'. The '/tmp' volume is backed by pod memory, not disk storage, so avoid setting a high value as it consumes pod memory resources.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
										},
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "The pod's tolerations.",
										MarkdownDescription: "The pod's tolerations.",
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
										Description:         "The pod's topology spread constraints.",
										MarkdownDescription: "The pod's topology spread constraints.",
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
													Required:            false,
													Optional:            true,
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
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"when_unsatisfiable": schema.StringAttribute{
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

									"volumes": schema.ListNestedAttribute{
										Description:         "Additional volumes that can be mounted to the pod.",
										MarkdownDescription: "Additional volumes that can be mounted to the pod.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap to use to populate the volume.",
													MarkdownDescription: "ConfigMap to use to populate the volume.",
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
																		Required:            false,
																		Optional:            true,
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

												"empty_dir": schema.SingleNestedAttribute{
													Description:         "EmptyDir to use to populate the volume.",
													MarkdownDescription: "EmptyDir to use to populate the volume.",
													Attributes: map[string]schema.Attribute{
														"medium": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"size_limit": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"amount": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"format": schema.StringAttribute{
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

												"name": schema.StringAttribute{
													Description:         "Name to use for the volume. Required.",
													MarkdownDescription: "Name to use for the volume. Required.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"persistent_volume_claim": schema.SingleNestedAttribute{
													Description:         "PersistentVolumeClaim object to use to populate the volume.",
													MarkdownDescription: "PersistentVolumeClaim object to use to populate the volume.",
													Attributes: map[string]schema.Attribute{
														"claim_name": schema.StringAttribute{
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

												"secret": schema.SingleNestedAttribute{
													Description:         "Secret to use populate the volume.",
													MarkdownDescription: "Secret to use populate the volume.",
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
																		Required:            false,
																		Optional:            true,
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

							"pod_set": schema.SingleNestedAttribute{
								Description:         "Template for Kafka 'StrimziPodSet' resource.",
								MarkdownDescription: "Template for Kafka 'StrimziPodSet' resource.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations added to the Kubernetes resource.",
												MarkdownDescription: "Annotations added to the Kubernetes resource.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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
		},
	}
}

func (r *KafkaStrimziIoKafkaNodePoolV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_strimzi_io_kafka_node_pool_v1beta2_manifest")

	var model KafkaStrimziIoKafkaNodePoolV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kafka.strimzi.io/v1beta2")
	model.Kind = pointer.String("KafkaNodePool")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
