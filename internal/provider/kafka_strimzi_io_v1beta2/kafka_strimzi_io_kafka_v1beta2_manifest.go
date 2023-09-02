/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kafka_strimzi_io_v1beta2

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KafkaStrimziIoKafkaV1Beta2Manifest{}
)

func NewKafkaStrimziIoKafkaV1Beta2Manifest() datasource.DataSource {
	return &KafkaStrimziIoKafkaV1Beta2Manifest{}
}

type KafkaStrimziIoKafkaV1Beta2Manifest struct{}

type KafkaStrimziIoKafkaV1Beta2ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		ClientsCa *struct {
			CertificateExpirationPolicy  *string `tfsdk:"certificate_expiration_policy" json:"certificateExpirationPolicy,omitempty"`
			GenerateCertificateAuthority *bool   `tfsdk:"generate_certificate_authority" json:"generateCertificateAuthority,omitempty"`
			GenerateSecretOwnerReference *bool   `tfsdk:"generate_secret_owner_reference" json:"generateSecretOwnerReference,omitempty"`
			RenewalDays                  *int64  `tfsdk:"renewal_days" json:"renewalDays,omitempty"`
			ValidityDays                 *int64  `tfsdk:"validity_days" json:"validityDays,omitempty"`
		} `tfsdk:"clients_ca" json:"clientsCa,omitempty"`
		ClusterCa *struct {
			CertificateExpirationPolicy  *string `tfsdk:"certificate_expiration_policy" json:"certificateExpirationPolicy,omitempty"`
			GenerateCertificateAuthority *bool   `tfsdk:"generate_certificate_authority" json:"generateCertificateAuthority,omitempty"`
			GenerateSecretOwnerReference *bool   `tfsdk:"generate_secret_owner_reference" json:"generateSecretOwnerReference,omitempty"`
			RenewalDays                  *int64  `tfsdk:"renewal_days" json:"renewalDays,omitempty"`
			ValidityDays                 *int64  `tfsdk:"validity_days" json:"validityDays,omitempty"`
		} `tfsdk:"cluster_ca" json:"clusterCa,omitempty"`
		CruiseControl *struct {
			BrokerCapacity *struct {
				Cpu             *string `tfsdk:"cpu" json:"cpu,omitempty"`
				CpuUtilization  *int64  `tfsdk:"cpu_utilization" json:"cpuUtilization,omitempty"`
				Disk            *string `tfsdk:"disk" json:"disk,omitempty"`
				InboundNetwork  *string `tfsdk:"inbound_network" json:"inboundNetwork,omitempty"`
				OutboundNetwork *string `tfsdk:"outbound_network" json:"outboundNetwork,omitempty"`
				Overrides       *[]struct {
					Brokers         *[]string `tfsdk:"brokers" json:"brokers,omitempty"`
					Cpu             *string   `tfsdk:"cpu" json:"cpu,omitempty"`
					InboundNetwork  *string   `tfsdk:"inbound_network" json:"inboundNetwork,omitempty"`
					OutboundNetwork *string   `tfsdk:"outbound_network" json:"outboundNetwork,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
			} `tfsdk:"broker_capacity" json:"brokerCapacity,omitempty"`
			Config     *map[string]string `tfsdk:"config" json:"config,omitempty"`
			Image      *string            `tfsdk:"image" json:"image,omitempty"`
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
			LivenessProbe *struct {
				FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
				PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
				SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
				TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			Logging *struct {
				Loggers   *map[string]string `tfsdk:"loggers" json:"loggers,omitempty"`
				Type      *string            `tfsdk:"type" json:"type,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"logging" json:"logging,omitempty"`
			MetricsConfig *struct {
				Type      *string `tfsdk:"type" json:"type,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"metrics_config" json:"metricsConfig,omitempty"`
			ReadinessProbe *struct {
				FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
				PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
				SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
				TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Template *struct {
				ApiService *struct {
					IpFamilies     *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
					IpFamilyPolicy *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"api_service" json:"apiService,omitempty"`
				CruiseControlContainer *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"cruise_control_container" json:"cruiseControlContainer,omitempty"`
				Deployment *struct {
					DeploymentStrategy *string `tfsdk:"deployment_strategy" json:"deploymentStrategy,omitempty"`
					Metadata           *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
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
				} `tfsdk:"pod" json:"pod,omitempty"`
				PodDisruptionBudget *struct {
					MaxUnavailable *int64 `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
				ServiceAccount *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				TlsSidecarContainer *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"tls_sidecar_container" json:"tlsSidecarContainer,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
			TlsSidecar *struct {
				Image         *string `tfsdk:"image" json:"image,omitempty"`
				LivenessProbe *struct {
					FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
				LogLevel       *string `tfsdk:"log_level" json:"logLevel,omitempty"`
				ReadinessProbe *struct {
					FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"tls_sidecar" json:"tlsSidecar,omitempty"`
		} `tfsdk:"cruise_control" json:"cruiseControl,omitempty"`
		EntityOperator *struct {
			Template *struct {
				Deployment *struct {
					DeploymentStrategy *string `tfsdk:"deployment_strategy" json:"deploymentStrategy,omitempty"`
					Metadata           *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
				EntityOperatorRole *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"entity_operator_role" json:"entityOperatorRole,omitempty"`
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
				} `tfsdk:"pod" json:"pod,omitempty"`
				ServiceAccount *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				TlsSidecarContainer *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"tls_sidecar_container" json:"tlsSidecarContainer,omitempty"`
				TopicOperatorContainer *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"topic_operator_container" json:"topicOperatorContainer,omitempty"`
				TopicOperatorRoleBinding *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"topic_operator_role_binding" json:"topicOperatorRoleBinding,omitempty"`
				UserOperatorContainer *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"user_operator_container" json:"userOperatorContainer,omitempty"`
				UserOperatorRoleBinding *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"user_operator_role_binding" json:"userOperatorRoleBinding,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
			TlsSidecar *struct {
				Image         *string `tfsdk:"image" json:"image,omitempty"`
				LivenessProbe *struct {
					FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
				LogLevel       *string `tfsdk:"log_level" json:"logLevel,omitempty"`
				ReadinessProbe *struct {
					FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
				Resources *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"tls_sidecar" json:"tlsSidecar,omitempty"`
			TopicOperator *struct {
				Image      *string `tfsdk:"image" json:"image,omitempty"`
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
				LivenessProbe *struct {
					FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
				Logging *struct {
					Loggers   *map[string]string `tfsdk:"loggers" json:"loggers,omitempty"`
					Type      *string            `tfsdk:"type" json:"type,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"logging" json:"logging,omitempty"`
				ReadinessProbe *struct {
					FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
				ReconciliationIntervalSeconds *int64 `tfsdk:"reconciliation_interval_seconds" json:"reconciliationIntervalSeconds,omitempty"`
				Resources                     *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				StartupProbe *struct {
					FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"startup_probe" json:"startupProbe,omitempty"`
				TopicMetadataMaxAttempts       *int64  `tfsdk:"topic_metadata_max_attempts" json:"topicMetadataMaxAttempts,omitempty"`
				WatchedNamespace               *string `tfsdk:"watched_namespace" json:"watchedNamespace,omitempty"`
				ZookeeperSessionTimeoutSeconds *int64  `tfsdk:"zookeeper_session_timeout_seconds" json:"zookeeperSessionTimeoutSeconds,omitempty"`
			} `tfsdk:"topic_operator" json:"topicOperator,omitempty"`
			UserOperator *struct {
				Image      *string `tfsdk:"image" json:"image,omitempty"`
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
				LivenessProbe *struct {
					FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
				Logging *struct {
					Loggers   *map[string]string `tfsdk:"loggers" json:"loggers,omitempty"`
					Type      *string            `tfsdk:"type" json:"type,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key      *string `tfsdk:"key" json:"key,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"logging" json:"logging,omitempty"`
				ReadinessProbe *struct {
					FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
				ReconciliationIntervalSeconds *int64 `tfsdk:"reconciliation_interval_seconds" json:"reconciliationIntervalSeconds,omitempty"`
				Resources                     *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				SecretPrefix                   *string `tfsdk:"secret_prefix" json:"secretPrefix,omitempty"`
				WatchedNamespace               *string `tfsdk:"watched_namespace" json:"watchedNamespace,omitempty"`
				ZookeeperSessionTimeoutSeconds *int64  `tfsdk:"zookeeper_session_timeout_seconds" json:"zookeeperSessionTimeoutSeconds,omitempty"`
			} `tfsdk:"user_operator" json:"userOperator,omitempty"`
		} `tfsdk:"entity_operator" json:"entityOperator,omitempty"`
		JmxTrans *struct {
			Image        *string `tfsdk:"image" json:"image,omitempty"`
			KafkaQueries *[]struct {
				Attributes  *[]string `tfsdk:"attributes" json:"attributes,omitempty"`
				Outputs     *[]string `tfsdk:"outputs" json:"outputs,omitempty"`
				TargetMBean *string   `tfsdk:"target_m_bean" json:"targetMBean,omitempty"`
			} `tfsdk:"kafka_queries" json:"kafkaQueries,omitempty"`
			LogLevel          *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			OutputDefinitions *[]struct {
				FlushDelayInSeconds *int64    `tfsdk:"flush_delay_in_seconds" json:"flushDelayInSeconds,omitempty"`
				Host                *string   `tfsdk:"host" json:"host,omitempty"`
				Name                *string   `tfsdk:"name" json:"name,omitempty"`
				OutputType          *string   `tfsdk:"output_type" json:"outputType,omitempty"`
				Port                *int64    `tfsdk:"port" json:"port,omitempty"`
				TypeNames           *[]string `tfsdk:"type_names" json:"typeNames,omitempty"`
			} `tfsdk:"output_definitions" json:"outputDefinitions,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Template *struct {
				Container *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"container" json:"container,omitempty"`
				Deployment *struct {
					DeploymentStrategy *string `tfsdk:"deployment_strategy" json:"deploymentStrategy,omitempty"`
					Metadata           *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
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
				} `tfsdk:"pod" json:"pod,omitempty"`
				ServiceAccount *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
		} `tfsdk:"jmx_trans" json:"jmxTrans,omitempty"`
		Kafka *struct {
			Authorization *struct {
				AllowOnError                   *bool     `tfsdk:"allow_on_error" json:"allowOnError,omitempty"`
				AuthorizerClass                *string   `tfsdk:"authorizer_class" json:"authorizerClass,omitempty"`
				ClientId                       *string   `tfsdk:"client_id" json:"clientId,omitempty"`
				ConnectTimeoutSeconds          *int64    `tfsdk:"connect_timeout_seconds" json:"connectTimeoutSeconds,omitempty"`
				DelegateToKafkaAcls            *bool     `tfsdk:"delegate_to_kafka_acls" json:"delegateToKafkaAcls,omitempty"`
				DisableTlsHostnameVerification *bool     `tfsdk:"disable_tls_hostname_verification" json:"disableTlsHostnameVerification,omitempty"`
				EnableMetrics                  *bool     `tfsdk:"enable_metrics" json:"enableMetrics,omitempty"`
				ExpireAfterMs                  *int64    `tfsdk:"expire_after_ms" json:"expireAfterMs,omitempty"`
				GrantsAlwaysLatest             *bool     `tfsdk:"grants_always_latest" json:"grantsAlwaysLatest,omitempty"`
				GrantsGcPeriodSeconds          *int64    `tfsdk:"grants_gc_period_seconds" json:"grantsGcPeriodSeconds,omitempty"`
				GrantsMaxIdleTimeSeconds       *int64    `tfsdk:"grants_max_idle_time_seconds" json:"grantsMaxIdleTimeSeconds,omitempty"`
				GrantsRefreshPeriodSeconds     *int64    `tfsdk:"grants_refresh_period_seconds" json:"grantsRefreshPeriodSeconds,omitempty"`
				GrantsRefreshPoolSize          *int64    `tfsdk:"grants_refresh_pool_size" json:"grantsRefreshPoolSize,omitempty"`
				HttpRetries                    *int64    `tfsdk:"http_retries" json:"httpRetries,omitempty"`
				InitialCacheCapacity           *int64    `tfsdk:"initial_cache_capacity" json:"initialCacheCapacity,omitempty"`
				MaximumCacheSize               *int64    `tfsdk:"maximum_cache_size" json:"maximumCacheSize,omitempty"`
				ReadTimeoutSeconds             *int64    `tfsdk:"read_timeout_seconds" json:"readTimeoutSeconds,omitempty"`
				SuperUsers                     *[]string `tfsdk:"super_users" json:"superUsers,omitempty"`
				SupportsAdminApi               *bool     `tfsdk:"supports_admin_api" json:"supportsAdminApi,omitempty"`
				TlsTrustedCertificates         *[]struct {
					Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
					SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				} `tfsdk:"tls_trusted_certificates" json:"tlsTrustedCertificates,omitempty"`
				TokenEndpointUri *string `tfsdk:"token_endpoint_uri" json:"tokenEndpointUri,omitempty"`
				Type             *string `tfsdk:"type" json:"type,omitempty"`
				Url              *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"authorization" json:"authorization,omitempty"`
			BrokerRackInitImage *string            `tfsdk:"broker_rack_init_image" json:"brokerRackInitImage,omitempty"`
			Config              *map[string]string `tfsdk:"config" json:"config,omitempty"`
			Image               *string            `tfsdk:"image" json:"image,omitempty"`
			JmxOptions          *struct {
				Authentication *struct {
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"authentication" json:"authentication,omitempty"`
			} `tfsdk:"jmx_options" json:"jmxOptions,omitempty"`
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
			Listeners *[]struct {
				Authentication *struct {
					AccessTokenIsJwt     *bool   `tfsdk:"access_token_is_jwt" json:"accessTokenIsJwt,omitempty"`
					CheckAccessTokenType *bool   `tfsdk:"check_access_token_type" json:"checkAccessTokenType,omitempty"`
					CheckAudience        *bool   `tfsdk:"check_audience" json:"checkAudience,omitempty"`
					CheckIssuer          *bool   `tfsdk:"check_issuer" json:"checkIssuer,omitempty"`
					ClientAudience       *string `tfsdk:"client_audience" json:"clientAudience,omitempty"`
					ClientId             *string `tfsdk:"client_id" json:"clientId,omitempty"`
					ClientScope          *string `tfsdk:"client_scope" json:"clientScope,omitempty"`
					ClientSecret         *struct {
						Key        *string `tfsdk:"key" json:"key,omitempty"`
						SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
					ConnectTimeoutSeconds             *int64             `tfsdk:"connect_timeout_seconds" json:"connectTimeoutSeconds,omitempty"`
					CustomClaimCheck                  *string            `tfsdk:"custom_claim_check" json:"customClaimCheck,omitempty"`
					DisableTlsHostnameVerification    *bool              `tfsdk:"disable_tls_hostname_verification" json:"disableTlsHostnameVerification,omitempty"`
					EnableECDSA                       *bool              `tfsdk:"enable_ecdsa" json:"enableECDSA,omitempty"`
					EnableMetrics                     *bool              `tfsdk:"enable_metrics" json:"enableMetrics,omitempty"`
					EnableOauthBearer                 *bool              `tfsdk:"enable_oauth_bearer" json:"enableOauthBearer,omitempty"`
					EnablePlain                       *bool              `tfsdk:"enable_plain" json:"enablePlain,omitempty"`
					FailFast                          *bool              `tfsdk:"fail_fast" json:"failFast,omitempty"`
					FallbackUserNameClaim             *string            `tfsdk:"fallback_user_name_claim" json:"fallbackUserNameClaim,omitempty"`
					FallbackUserNamePrefix            *string            `tfsdk:"fallback_user_name_prefix" json:"fallbackUserNamePrefix,omitempty"`
					GroupsClaim                       *string            `tfsdk:"groups_claim" json:"groupsClaim,omitempty"`
					GroupsClaimDelimiter              *string            `tfsdk:"groups_claim_delimiter" json:"groupsClaimDelimiter,omitempty"`
					HttpRetries                       *int64             `tfsdk:"http_retries" json:"httpRetries,omitempty"`
					HttpRetryPauseMs                  *int64             `tfsdk:"http_retry_pause_ms" json:"httpRetryPauseMs,omitempty"`
					IntrospectionEndpointUri          *string            `tfsdk:"introspection_endpoint_uri" json:"introspectionEndpointUri,omitempty"`
					JwksEndpointUri                   *string            `tfsdk:"jwks_endpoint_uri" json:"jwksEndpointUri,omitempty"`
					JwksExpirySeconds                 *int64             `tfsdk:"jwks_expiry_seconds" json:"jwksExpirySeconds,omitempty"`
					JwksIgnoreKeyUse                  *bool              `tfsdk:"jwks_ignore_key_use" json:"jwksIgnoreKeyUse,omitempty"`
					JwksMinRefreshPauseSeconds        *int64             `tfsdk:"jwks_min_refresh_pause_seconds" json:"jwksMinRefreshPauseSeconds,omitempty"`
					JwksRefreshSeconds                *int64             `tfsdk:"jwks_refresh_seconds" json:"jwksRefreshSeconds,omitempty"`
					ListenerConfig                    *map[string]string `tfsdk:"listener_config" json:"listenerConfig,omitempty"`
					MaxSecondsWithoutReauthentication *int64             `tfsdk:"max_seconds_without_reauthentication" json:"maxSecondsWithoutReauthentication,omitempty"`
					ReadTimeoutSeconds                *int64             `tfsdk:"read_timeout_seconds" json:"readTimeoutSeconds,omitempty"`
					Sasl                              *bool              `tfsdk:"sasl" json:"sasl,omitempty"`
					Secrets                           *[]struct {
						Key        *string `tfsdk:"key" json:"key,omitempty"`
						SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					} `tfsdk:"secrets" json:"secrets,omitempty"`
					TlsTrustedCertificates *[]struct {
						Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
						SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					} `tfsdk:"tls_trusted_certificates" json:"tlsTrustedCertificates,omitempty"`
					TokenEndpointUri    *string `tfsdk:"token_endpoint_uri" json:"tokenEndpointUri,omitempty"`
					Type                *string `tfsdk:"type" json:"type,omitempty"`
					UserInfoEndpointUri *string `tfsdk:"user_info_endpoint_uri" json:"userInfoEndpointUri,omitempty"`
					UserNameClaim       *string `tfsdk:"user_name_claim" json:"userNameClaim,omitempty"`
					ValidIssuerUri      *string `tfsdk:"valid_issuer_uri" json:"validIssuerUri,omitempty"`
					ValidTokenType      *string `tfsdk:"valid_token_type" json:"validTokenType,omitempty"`
				} `tfsdk:"authentication" json:"authentication,omitempty"`
				Configuration *struct {
					Bootstrap *struct {
						AlternativeNames *[]string          `tfsdk:"alternative_names" json:"alternativeNames,omitempty"`
						Annotations      *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Host             *string            `tfsdk:"host" json:"host,omitempty"`
						Labels           *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						LoadBalancerIP   *string            `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
						NodePort         *int64             `tfsdk:"node_port" json:"nodePort,omitempty"`
					} `tfsdk:"bootstrap" json:"bootstrap,omitempty"`
					BrokerCertChainAndKey *struct {
						Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
						Key         *string `tfsdk:"key" json:"key,omitempty"`
						SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					} `tfsdk:"broker_cert_chain_and_key" json:"brokerCertChainAndKey,omitempty"`
					Brokers *[]struct {
						AdvertisedHost *string            `tfsdk:"advertised_host" json:"advertisedHost,omitempty"`
						AdvertisedPort *int64             `tfsdk:"advertised_port" json:"advertisedPort,omitempty"`
						Annotations    *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Broker         *int64             `tfsdk:"broker" json:"broker,omitempty"`
						Host           *string            `tfsdk:"host" json:"host,omitempty"`
						Labels         *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						LoadBalancerIP *string            `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
						NodePort       *int64             `tfsdk:"node_port" json:"nodePort,omitempty"`
					} `tfsdk:"brokers" json:"brokers,omitempty"`
					Class                        *string   `tfsdk:"class" json:"class,omitempty"`
					CreateBootstrapService       *bool     `tfsdk:"create_bootstrap_service" json:"createBootstrapService,omitempty"`
					ExternalTrafficPolicy        *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
					Finalizers                   *[]string `tfsdk:"finalizers" json:"finalizers,omitempty"`
					IpFamilies                   *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
					IpFamilyPolicy               *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
					LoadBalancerSourceRanges     *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
					MaxConnectionCreationRate    *int64    `tfsdk:"max_connection_creation_rate" json:"maxConnectionCreationRate,omitempty"`
					MaxConnections               *int64    `tfsdk:"max_connections" json:"maxConnections,omitempty"`
					PreferredNodePortAddressType *string   `tfsdk:"preferred_node_port_address_type" json:"preferredNodePortAddressType,omitempty"`
					UseServiceDnsDomain          *bool     `tfsdk:"use_service_dns_domain" json:"useServiceDnsDomain,omitempty"`
				} `tfsdk:"configuration" json:"configuration,omitempty"`
				Name               *string `tfsdk:"name" json:"name,omitempty"`
				NetworkPolicyPeers *[]struct {
					IpBlock *struct {
						Cidr   *string   `tfsdk:"cidr" json:"cidr,omitempty"`
						Except *[]string `tfsdk:"except" json:"except,omitempty"`
					} `tfsdk:"ip_block" json:"ipBlock,omitempty"`
					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" json:"namespaceSelector,omitempty"`
					PodSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" json:"podSelector,omitempty"`
				} `tfsdk:"network_policy_peers" json:"networkPolicyPeers,omitempty"`
				Port *int64  `tfsdk:"port" json:"port,omitempty"`
				Tls  *bool   `tfsdk:"tls" json:"tls,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"listeners" json:"listeners,omitempty"`
			LivenessProbe *struct {
				FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
				PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
				SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
				TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			Logging *struct {
				Loggers   *map[string]string `tfsdk:"loggers" json:"loggers,omitempty"`
				Type      *string            `tfsdk:"type" json:"type,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"logging" json:"logging,omitempty"`
			MetricsConfig *struct {
				Type      *string `tfsdk:"type" json:"type,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"metrics_config" json:"metricsConfig,omitempty"`
			Rack *struct {
				TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
			} `tfsdk:"rack" json:"rack,omitempty"`
			ReadinessProbe *struct {
				FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
				PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
				SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
				TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Storage *struct {
				Class       *string `tfsdk:"class" json:"class,omitempty"`
				DeleteClaim *bool   `tfsdk:"delete_claim" json:"deleteClaim,omitempty"`
				Id          *int64  `tfsdk:"id" json:"id,omitempty"`
				Overrides   *[]struct {
					Broker *int64  `tfsdk:"broker" json:"broker,omitempty"`
					Class  *string `tfsdk:"class" json:"class,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				Selector  *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
				Size      *string            `tfsdk:"size" json:"size,omitempty"`
				SizeLimit *string            `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				Type      *string            `tfsdk:"type" json:"type,omitempty"`
				Volumes   *[]struct {
					Class       *string `tfsdk:"class" json:"class,omitempty"`
					DeleteClaim *bool   `tfsdk:"delete_claim" json:"deleteClaim,omitempty"`
					Id          *int64  `tfsdk:"id" json:"id,omitempty"`
					Overrides   *[]struct {
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
				BootstrapService *struct {
					IpFamilies     *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
					IpFamilyPolicy *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"bootstrap_service" json:"bootstrapService,omitempty"`
				BrokersService *struct {
					IpFamilies     *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
					IpFamilyPolicy *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"brokers_service" json:"brokersService,omitempty"`
				ClusterCaCert *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"cluster_ca_cert" json:"clusterCaCert,omitempty"`
				ClusterRoleBinding *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"cluster_role_binding" json:"clusterRoleBinding,omitempty"`
				ExternalBootstrapIngress *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"external_bootstrap_ingress" json:"externalBootstrapIngress,omitempty"`
				ExternalBootstrapRoute *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"external_bootstrap_route" json:"externalBootstrapRoute,omitempty"`
				ExternalBootstrapService *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"external_bootstrap_service" json:"externalBootstrapService,omitempty"`
				InitContainer *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"init_container" json:"initContainer,omitempty"`
				JmxSecret *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"jmx_secret" json:"jmxSecret,omitempty"`
				KafkaContainer *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"pod" json:"pod,omitempty"`
				PodDisruptionBudget *struct {
					MaxUnavailable *int64 `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
				PodSet *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"pod_set" json:"podSet,omitempty"`
				ServiceAccount *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				Statefulset *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					PodManagementPolicy *string `tfsdk:"pod_management_policy" json:"podManagementPolicy,omitempty"`
				} `tfsdk:"statefulset" json:"statefulset,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
			Version *string `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"kafka" json:"kafka,omitempty"`
		KafkaExporter *struct {
			EnableSaramaLogging *bool   `tfsdk:"enable_sarama_logging" json:"enableSaramaLogging,omitempty"`
			GroupExcludeRegex   *string `tfsdk:"group_exclude_regex" json:"groupExcludeRegex,omitempty"`
			GroupRegex          *string `tfsdk:"group_regex" json:"groupRegex,omitempty"`
			Image               *string `tfsdk:"image" json:"image,omitempty"`
			LivenessProbe       *struct {
				FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
				PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
				SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
				TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			Logging        *string `tfsdk:"logging" json:"logging,omitempty"`
			ReadinessProbe *struct {
				FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
				PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
				SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
				TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Template *struct {
				Container *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"container" json:"container,omitempty"`
				Deployment *struct {
					DeploymentStrategy *string `tfsdk:"deployment_strategy" json:"deploymentStrategy,omitempty"`
					Metadata           *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"deployment" json:"deployment,omitempty"`
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
				} `tfsdk:"pod" json:"pod,omitempty"`
				Service *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"service" json:"service,omitempty"`
				ServiceAccount *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
			TopicExcludeRegex *string `tfsdk:"topic_exclude_regex" json:"topicExcludeRegex,omitempty"`
			TopicRegex        *string `tfsdk:"topic_regex" json:"topicRegex,omitempty"`
		} `tfsdk:"kafka_exporter" json:"kafkaExporter,omitempty"`
		MaintenanceTimeWindows *[]string `tfsdk:"maintenance_time_windows" json:"maintenanceTimeWindows,omitempty"`
		Zookeeper              *struct {
			Config     *map[string]string `tfsdk:"config" json:"config,omitempty"`
			Image      *string            `tfsdk:"image" json:"image,omitempty"`
			JmxOptions *struct {
				Authentication *struct {
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"authentication" json:"authentication,omitempty"`
			} `tfsdk:"jmx_options" json:"jmxOptions,omitempty"`
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
			LivenessProbe *struct {
				FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
				PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
				SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
				TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
			Logging *struct {
				Loggers   *map[string]string `tfsdk:"loggers" json:"loggers,omitempty"`
				Type      *string            `tfsdk:"type" json:"type,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"logging" json:"logging,omitempty"`
			MetricsConfig *struct {
				Type      *string `tfsdk:"type" json:"type,omitempty"`
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" json:"valueFrom,omitempty"`
			} `tfsdk:"metrics_config" json:"metricsConfig,omitempty"`
			ReadinessProbe *struct {
				FailureThreshold    *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
				PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
				SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
				TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
			Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Storage *struct {
				Class       *string `tfsdk:"class" json:"class,omitempty"`
				DeleteClaim *bool   `tfsdk:"delete_claim" json:"deleteClaim,omitempty"`
				Id          *int64  `tfsdk:"id" json:"id,omitempty"`
				Overrides   *[]struct {
					Broker *int64  `tfsdk:"broker" json:"broker,omitempty"`
					Class  *string `tfsdk:"class" json:"class,omitempty"`
				} `tfsdk:"overrides" json:"overrides,omitempty"`
				Selector  *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
				Size      *string            `tfsdk:"size" json:"size,omitempty"`
				SizeLimit *string            `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				Type      *string            `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
			Template *struct {
				ClientService *struct {
					IpFamilies     *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
					IpFamilyPolicy *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"client_service" json:"clientService,omitempty"`
				JmxSecret *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"jmx_secret" json:"jmxSecret,omitempty"`
				NodesService *struct {
					IpFamilies     *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
					IpFamilyPolicy *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"nodes_service" json:"nodesService,omitempty"`
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
				} `tfsdk:"pod" json:"pod,omitempty"`
				PodDisruptionBudget *struct {
					MaxUnavailable *int64 `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					Metadata       *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
				PodSet *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"pod_set" json:"podSet,omitempty"`
				ServiceAccount *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
				} `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				Statefulset *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					PodManagementPolicy *string `tfsdk:"pod_management_policy" json:"podManagementPolicy,omitempty"`
				} `tfsdk:"statefulset" json:"statefulset,omitempty"`
				ZookeeperContainer *struct {
					Env *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"env" json:"env,omitempty"`
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
				} `tfsdk:"zookeeper_container" json:"zookeeperContainer,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
		} `tfsdk:"zookeeper" json:"zookeeper,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KafkaStrimziIoKafkaV1Beta2Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_strimzi_io_kafka_v1beta2_manifest"
}

func (r *KafkaStrimziIoKafkaV1Beta2Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
				Description:         "The specification of the Kafka and ZooKeeper clusters, and Topic Operator.",
				MarkdownDescription: "The specification of the Kafka and ZooKeeper clusters, and Topic Operator.",
				Attributes: map[string]schema.Attribute{
					"clients_ca": schema.SingleNestedAttribute{
						Description:         "Configuration of the clients certificate authority.",
						MarkdownDescription: "Configuration of the clients certificate authority.",
						Attributes: map[string]schema.Attribute{
							"certificate_expiration_policy": schema.StringAttribute{
								Description:         "How should CA certificate expiration be handled when 'generateCertificateAuthority=true'. The default is for a new CA certificate to be generated reusing the existing private key.",
								MarkdownDescription: "How should CA certificate expiration be handled when 'generateCertificateAuthority=true'. The default is for a new CA certificate to be generated reusing the existing private key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("renew-certificate", "replace-key"),
								},
							},

							"generate_certificate_authority": schema.BoolAttribute{
								Description:         "If true then Certificate Authority certificates will be generated automatically. Otherwise the user will need to provide a Secret with the CA certificate. Default is true.",
								MarkdownDescription: "If true then Certificate Authority certificates will be generated automatically. Otherwise the user will need to provide a Secret with the CA certificate. Default is true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"generate_secret_owner_reference": schema.BoolAttribute{
								Description:         "If 'true', the Cluster and Client CA Secrets are configured with the 'ownerReference' set to the 'Kafka' resource. If the 'Kafka' resource is deleted when 'true', the CA Secrets are also deleted. If 'false', the 'ownerReference' is disabled. If the 'Kafka' resource is deleted when 'false', the CA Secrets are retained and available for reuse. Default is 'true'.",
								MarkdownDescription: "If 'true', the Cluster and Client CA Secrets are configured with the 'ownerReference' set to the 'Kafka' resource. If the 'Kafka' resource is deleted when 'true', the CA Secrets are also deleted. If 'false', the 'ownerReference' is disabled. If the 'Kafka' resource is deleted when 'false', the CA Secrets are retained and available for reuse. Default is 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"renewal_days": schema.Int64Attribute{
								Description:         "The number of days in the certificate renewal period. This is the number of days before the a certificate expires during which renewal actions may be performed. When 'generateCertificateAuthority' is true, this will cause the generation of a new certificate. When 'generateCertificateAuthority' is true, this will cause extra logging at WARN level about the pending certificate expiry. Default is 30.",
								MarkdownDescription: "The number of days in the certificate renewal period. This is the number of days before the a certificate expires during which renewal actions may be performed. When 'generateCertificateAuthority' is true, this will cause the generation of a new certificate. When 'generateCertificateAuthority' is true, this will cause extra logging at WARN level about the pending certificate expiry. Default is 30.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"validity_days": schema.Int64Attribute{
								Description:         "The number of days generated certificates should be valid for. The default is 365.",
								MarkdownDescription: "The number of days generated certificates should be valid for. The default is 365.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_ca": schema.SingleNestedAttribute{
						Description:         "Configuration of the cluster certificate authority.",
						MarkdownDescription: "Configuration of the cluster certificate authority.",
						Attributes: map[string]schema.Attribute{
							"certificate_expiration_policy": schema.StringAttribute{
								Description:         "How should CA certificate expiration be handled when 'generateCertificateAuthority=true'. The default is for a new CA certificate to be generated reusing the existing private key.",
								MarkdownDescription: "How should CA certificate expiration be handled when 'generateCertificateAuthority=true'. The default is for a new CA certificate to be generated reusing the existing private key.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("renew-certificate", "replace-key"),
								},
							},

							"generate_certificate_authority": schema.BoolAttribute{
								Description:         "If true then Certificate Authority certificates will be generated automatically. Otherwise the user will need to provide a Secret with the CA certificate. Default is true.",
								MarkdownDescription: "If true then Certificate Authority certificates will be generated automatically. Otherwise the user will need to provide a Secret with the CA certificate. Default is true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"generate_secret_owner_reference": schema.BoolAttribute{
								Description:         "If 'true', the Cluster and Client CA Secrets are configured with the 'ownerReference' set to the 'Kafka' resource. If the 'Kafka' resource is deleted when 'true', the CA Secrets are also deleted. If 'false', the 'ownerReference' is disabled. If the 'Kafka' resource is deleted when 'false', the CA Secrets are retained and available for reuse. Default is 'true'.",
								MarkdownDescription: "If 'true', the Cluster and Client CA Secrets are configured with the 'ownerReference' set to the 'Kafka' resource. If the 'Kafka' resource is deleted when 'true', the CA Secrets are also deleted. If 'false', the 'ownerReference' is disabled. If the 'Kafka' resource is deleted when 'false', the CA Secrets are retained and available for reuse. Default is 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"renewal_days": schema.Int64Attribute{
								Description:         "The number of days in the certificate renewal period. This is the number of days before the a certificate expires during which renewal actions may be performed. When 'generateCertificateAuthority' is true, this will cause the generation of a new certificate. When 'generateCertificateAuthority' is true, this will cause extra logging at WARN level about the pending certificate expiry. Default is 30.",
								MarkdownDescription: "The number of days in the certificate renewal period. This is the number of days before the a certificate expires during which renewal actions may be performed. When 'generateCertificateAuthority' is true, this will cause the generation of a new certificate. When 'generateCertificateAuthority' is true, this will cause extra logging at WARN level about the pending certificate expiry. Default is 30.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"validity_days": schema.Int64Attribute{
								Description:         "The number of days generated certificates should be valid for. The default is 365.",
								MarkdownDescription: "The number of days generated certificates should be valid for. The default is 365.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cruise_control": schema.SingleNestedAttribute{
						Description:         "Configuration for Cruise Control deployment. Deploys a Cruise Control instance when specified.",
						MarkdownDescription: "Configuration for Cruise Control deployment. Deploys a Cruise Control instance when specified.",
						Attributes: map[string]schema.Attribute{
							"broker_capacity": schema.SingleNestedAttribute{
								Description:         "The Cruise Control 'brokerCapacity' configuration.",
								MarkdownDescription: "The Cruise Control 'brokerCapacity' configuration.",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "Broker capacity for CPU resource in cores or millicores. For example, 1, 1.500, 1500m. For more information on valid CPU resource units see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.",
										MarkdownDescription: "Broker capacity for CPU resource in cores or millicores. For example, 1, 1.500, 1500m. For more information on valid CPU resource units see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([.][0-9]{0,3}|[m]?)$`), ""),
										},
									},

									"cpu_utilization": schema.Int64Attribute{
										Description:         "Broker capacity for CPU resource utilization as a percentage (0 - 100).",
										MarkdownDescription: "Broker capacity for CPU resource utilization as a percentage (0 - 100).",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
											int64validator.AtMost(100),
										},
									},

									"disk": schema.StringAttribute{
										Description:         "Broker capacity for disk in bytes. Use a number value with either standard Kubernetes byte units (K, M, G, or T), their bibyte (power of two) equivalents (Ki, Mi, Gi, or Ti), or a byte value with or without E notation. For example, 100000M, 100000Mi, 104857600000, or 1e+11.",
										MarkdownDescription: "Broker capacity for disk in bytes. Use a number value with either standard Kubernetes byte units (K, M, G, or T), their bibyte (power of two) equivalents (Ki, Mi, Gi, or Ti), or a byte value with or without E notation. For example, 100000M, 100000Mi, 104857600000, or 1e+11.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([.][0-9]*)?([KMGTPE]i?|e[0-9]+)?$`), ""),
										},
									},

									"inbound_network": schema.StringAttribute{
										Description:         "Broker capacity for inbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
										MarkdownDescription: "Broker capacity for inbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([KMG]i?)?B/s$`), ""),
										},
									},

									"outbound_network": schema.StringAttribute{
										Description:         "Broker capacity for outbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
										MarkdownDescription: "Broker capacity for outbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([KMG]i?)?B/s$`), ""),
										},
									},

									"overrides": schema.ListNestedAttribute{
										Description:         "Overrides for individual brokers. The 'overrides' property lets you specify a different capacity configuration for different brokers.",
										MarkdownDescription: "Overrides for individual brokers. The 'overrides' property lets you specify a different capacity configuration for different brokers.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"brokers": schema.ListAttribute{
													Description:         "List of Kafka brokers (broker identifiers).",
													MarkdownDescription: "List of Kafka brokers (broker identifiers).",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"cpu": schema.StringAttribute{
													Description:         "Broker capacity for CPU resource in cores or millicores. For example, 1, 1.500, 1500m. For more information on valid CPU resource units see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.",
													MarkdownDescription: "Broker capacity for CPU resource in cores or millicores. For example, 1, 1.500, 1500m. For more information on valid CPU resource units see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([.][0-9]{0,3}|[m]?)$`), ""),
													},
												},

												"inbound_network": schema.StringAttribute{
													Description:         "Broker capacity for inbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
													MarkdownDescription: "Broker capacity for inbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([KMG]i?)?B/s$`), ""),
													},
												},

												"outbound_network": schema.StringAttribute{
													Description:         "Broker capacity for outbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
													MarkdownDescription: "Broker capacity for outbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([KMG]i?)?B/s$`), ""),
													},
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

							"config": schema.MapAttribute{
								Description:         "The Cruise Control configuration. For a full list of configuration options refer to https://github.com/linkedin/cruise-control/wiki/Configurations. Note that properties with the following prefixes cannot be set: bootstrap.servers, client.id, zookeeper., network., security., failed.brokers.zk.path,webserver.http., webserver.api.urlprefix, webserver.session.path, webserver.accesslog., two.step., request.reason.required,metric.reporter.sampler.bootstrap.servers, capacity.config.file, self.healing., ssl., kafka.broker.failure.detection.enable, topic.config.provider.class (with the exception of: ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols, webserver.http.cors.enabled, webserver.http.cors.origin, webserver.http.cors.exposeheaders, webserver.security.enable, webserver.ssl.enable).",
								MarkdownDescription: "The Cruise Control configuration. For a full list of configuration options refer to https://github.com/linkedin/cruise-control/wiki/Configurations. Note that properties with the following prefixes cannot be set: bootstrap.servers, client.id, zookeeper., network., security., failed.brokers.zk.path,webserver.http., webserver.api.urlprefix, webserver.session.path, webserver.accesslog., two.step., request.reason.required,metric.reporter.sampler.bootstrap.servers, capacity.config.file, self.healing., ssl., kafka.broker.failure.detection.enable, topic.config.provider.class (with the exception of: ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols, webserver.http.cors.enabled, webserver.http.cors.origin, webserver.http.cors.exposeheaders, webserver.security.enable, webserver.ssl.enable).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "The docker image for the pods.",
								MarkdownDescription: "The docker image for the pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jvm_options": schema.SingleNestedAttribute{
								Description:         "JVM Options for the Cruise Control container.",
								MarkdownDescription: "JVM Options for the Cruise Control container.",
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

							"liveness_probe": schema.SingleNestedAttribute{
								Description:         "Pod liveness checking for the Cruise Control container.",
								MarkdownDescription: "Pod liveness checking for the Cruise Control container.",
								Attributes: map[string]schema.Attribute{
									"failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": schema.Int64Attribute{
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"period_seconds": schema.Int64Attribute{
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": schema.SingleNestedAttribute{
								Description:         "Logging configuration (Log4j 2) for Cruise Control.",
								MarkdownDescription: "Logging configuration (Log4j 2) for Cruise Control.",
								Attributes: map[string]schema.Attribute{
									"loggers": schema.MapAttribute{
										Description:         "A Map from logger name to logger level.",
										MarkdownDescription: "A Map from logger name to logger level.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Logging type, must be either 'inline' or 'external'.",
										MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("inline", "external"),
										},
									},

									"value_from": schema.SingleNestedAttribute{
										Description:         "'ConfigMap' entry where the logging configuration is stored. ",
										MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics_config": schema.SingleNestedAttribute{
								Description:         "Metrics configuration.",
								MarkdownDescription: "Metrics configuration.",
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Description:         "Metrics type. Only 'jmxPrometheusExporter' supported currently.",
										MarkdownDescription: "Metrics type. Only 'jmxPrometheusExporter' supported currently.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("jmxPrometheusExporter"),
										},
									},

									"value_from": schema.SingleNestedAttribute{
										Description:         "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",
										MarkdownDescription: "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
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
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"readiness_probe": schema.SingleNestedAttribute{
								Description:         "Pod readiness checking for the Cruise Control container.",
								MarkdownDescription: "Pod readiness checking for the Cruise Control container.",
								Attributes: map[string]schema.Attribute{
									"failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": schema.Int64Attribute{
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"period_seconds": schema.Int64Attribute{
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "CPU and memory resources to reserve for the Cruise Control container.",
								MarkdownDescription: "CPU and memory resources to reserve for the Cruise Control container.",
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

							"template": schema.SingleNestedAttribute{
								Description:         "Template to specify how Cruise Control resources, 'Deployments' and 'Pods', are generated.",
								MarkdownDescription: "Template to specify how Cruise Control resources, 'Deployments' and 'Pods', are generated.",
								Attributes: map[string]schema.Attribute{
									"api_service": schema.SingleNestedAttribute{
										Description:         "Template for Cruise Control API 'Service'.",
										MarkdownDescription: "Template for Cruise Control API 'Service'.",
										Attributes: map[string]schema.Attribute{
											"ip_families": schema.ListAttribute{
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cruise_control_container": schema.SingleNestedAttribute{
										Description:         "Template for the Cruise Control container.",
										MarkdownDescription: "Template for the Cruise Control container.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment": schema.SingleNestedAttribute{
										Description:         "Template for Cruise Control 'Deployment'.",
										MarkdownDescription: "Template for Cruise Control 'Deployment'.",
										Attributes: map[string]schema.Attribute{
											"deployment_strategy": schema.StringAttribute{
												Description:         "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
												MarkdownDescription: "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("RollingUpdate", "Recreate"),
												},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod": schema.SingleNestedAttribute{
										Description:         "Template for Cruise Control 'Pods'.",
										MarkdownDescription: "Template for Cruise Control 'Pods'.",
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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
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
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_disruption_budget": schema.SingleNestedAttribute{
										Description:         "Template for Cruise Control 'PodDisruptionBudget'.",
										MarkdownDescription: "Template for Cruise Control 'PodDisruptionBudget'.",
										Attributes: map[string]schema.Attribute{
											"max_unavailable": schema.Int64Attribute{
												Description:         "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
												MarkdownDescription: "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"metadata": schema.SingleNestedAttribute{
												Description:         "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
												MarkdownDescription: "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
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

									"service_account": schema.SingleNestedAttribute{
										Description:         "Template for the Cruise Control service account.",
										MarkdownDescription: "Template for the Cruise Control service account.",
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

									"tls_sidecar_container": schema.SingleNestedAttribute{
										Description:         "Template for the Cruise Control TLS sidecar container.",
										MarkdownDescription: "Template for the Cruise Control TLS sidecar container.",
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

							"tls_sidecar": schema.SingleNestedAttribute{
								Description:         "TLS sidecar configuration.",
								MarkdownDescription: "TLS sidecar configuration.",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "The docker image for the container.",
										MarkdownDescription: "The docker image for the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"liveness_probe": schema.SingleNestedAttribute{
										Description:         "Pod liveness checking.",
										MarkdownDescription: "Pod liveness checking.",
										Attributes: map[string]schema.Attribute{
											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_level": schema.StringAttribute{
										Description:         "The log level for the TLS sidecar. Default value is 'notice'.",
										MarkdownDescription: "The log level for the TLS sidecar. Default value is 'notice'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("emerg", "alert", "crit", "err", "warning", "notice", "info", "debug"),
										},
									},

									"readiness_probe": schema.SingleNestedAttribute{
										Description:         "Pod readiness checking.",
										MarkdownDescription: "Pod readiness checking.",
										Attributes: map[string]schema.Attribute{
											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
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

					"entity_operator": schema.SingleNestedAttribute{
						Description:         "Configuration of the Entity Operator.",
						MarkdownDescription: "Configuration of the Entity Operator.",
						Attributes: map[string]schema.Attribute{
							"template": schema.SingleNestedAttribute{
								Description:         "Template for Entity Operator resources. The template allows users to specify how a 'Deployment' and 'Pod' is generated.",
								MarkdownDescription: "Template for Entity Operator resources. The template allows users to specify how a 'Deployment' and 'Pod' is generated.",
								Attributes: map[string]schema.Attribute{
									"deployment": schema.SingleNestedAttribute{
										Description:         "Template for Entity Operator 'Deployment'.",
										MarkdownDescription: "Template for Entity Operator 'Deployment'.",
										Attributes: map[string]schema.Attribute{
											"deployment_strategy": schema.StringAttribute{
												Description:         "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
												MarkdownDescription: "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("RollingUpdate", "Recreate"),
												},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"entity_operator_role": schema.SingleNestedAttribute{
										Description:         "Template for the Entity Operator Role.",
										MarkdownDescription: "Template for the Entity Operator Role.",
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
										Description:         "Template for Entity Operator 'Pods'.",
										MarkdownDescription: "Template for Entity Operator 'Pods'.",
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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
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
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_account": schema.SingleNestedAttribute{
										Description:         "Template for the Entity Operator service account.",
										MarkdownDescription: "Template for the Entity Operator service account.",
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

									"tls_sidecar_container": schema.SingleNestedAttribute{
										Description:         "Template for the Entity Operator TLS sidecar container.",
										MarkdownDescription: "Template for the Entity Operator TLS sidecar container.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"topic_operator_container": schema.SingleNestedAttribute{
										Description:         "Template for the Entity Topic Operator container.",
										MarkdownDescription: "Template for the Entity Topic Operator container.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"topic_operator_role_binding": schema.SingleNestedAttribute{
										Description:         "Template for the Entity Topic Operator RoleBinding.",
										MarkdownDescription: "Template for the Entity Topic Operator RoleBinding.",
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

									"user_operator_container": schema.SingleNestedAttribute{
										Description:         "Template for the Entity User Operator container.",
										MarkdownDescription: "Template for the Entity User Operator container.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_operator_role_binding": schema.SingleNestedAttribute{
										Description:         "Template for the Entity Topic Operator RoleBinding.",
										MarkdownDescription: "Template for the Entity Topic Operator RoleBinding.",
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

							"tls_sidecar": schema.SingleNestedAttribute{
								Description:         "TLS sidecar configuration.",
								MarkdownDescription: "TLS sidecar configuration.",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "The docker image for the container.",
										MarkdownDescription: "The docker image for the container.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"liveness_probe": schema.SingleNestedAttribute{
										Description:         "Pod liveness checking.",
										MarkdownDescription: "Pod liveness checking.",
										Attributes: map[string]schema.Attribute{
											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_level": schema.StringAttribute{
										Description:         "The log level for the TLS sidecar. Default value is 'notice'.",
										MarkdownDescription: "The log level for the TLS sidecar. Default value is 'notice'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("emerg", "alert", "crit", "err", "warning", "notice", "info", "debug"),
										},
									},

									"readiness_probe": schema.SingleNestedAttribute{
										Description:         "Pod readiness checking.",
										MarkdownDescription: "Pod readiness checking.",
										Attributes: map[string]schema.Attribute{
											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"topic_operator": schema.SingleNestedAttribute{
								Description:         "Configuration of the Topic Operator.",
								MarkdownDescription: "Configuration of the Topic Operator.",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "The image to use for the Topic Operator.",
										MarkdownDescription: "The image to use for the Topic Operator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

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

									"liveness_probe": schema.SingleNestedAttribute{
										Description:         "Pod liveness checking.",
										MarkdownDescription: "Pod liveness checking.",
										Attributes: map[string]schema.Attribute{
											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"logging": schema.SingleNestedAttribute{
										Description:         "Logging configuration.",
										MarkdownDescription: "Logging configuration.",
										Attributes: map[string]schema.Attribute{
											"loggers": schema.MapAttribute{
												Description:         "A Map from logger name to logger level.",
												MarkdownDescription: "A Map from logger name to logger level.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Logging type, must be either 'inline' or 'external'.",
												MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("inline", "external"),
												},
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "'ConfigMap' entry where the logging configuration is stored. ",
												MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",
												Attributes: map[string]schema.Attribute{
													"config_map_key_ref": schema.SingleNestedAttribute{
														Description:         "Reference to the key in the ConfigMap containing the configuration.",
														MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_probe": schema.SingleNestedAttribute{
										Description:         "Pod readiness checking.",
										MarkdownDescription: "Pod readiness checking.",
										Attributes: map[string]schema.Attribute{
											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"reconciliation_interval_seconds": schema.Int64Attribute{
										Description:         "Interval between periodic reconciliations.",
										MarkdownDescription: "Interval between periodic reconciliations.",
										Required:            false,
										Optional:            true,
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

									"startup_probe": schema.SingleNestedAttribute{
										Description:         "Pod startup checking.",
										MarkdownDescription: "Pod startup checking.",
										Attributes: map[string]schema.Attribute{
											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"topic_metadata_max_attempts": schema.Int64Attribute{
										Description:         "The number of attempts at getting topic metadata.",
										MarkdownDescription: "The number of attempts at getting topic metadata.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"watched_namespace": schema.StringAttribute{
										Description:         "The namespace the Topic Operator should watch.",
										MarkdownDescription: "The namespace the Topic Operator should watch.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"zookeeper_session_timeout_seconds": schema.Int64Attribute{
										Description:         "Timeout for the ZooKeeper session.",
										MarkdownDescription: "Timeout for the ZooKeeper session.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_operator": schema.SingleNestedAttribute{
								Description:         "Configuration of the User Operator.",
								MarkdownDescription: "Configuration of the User Operator.",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "The image to use for the User Operator.",
										MarkdownDescription: "The image to use for the User Operator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

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

									"liveness_probe": schema.SingleNestedAttribute{
										Description:         "Pod liveness checking.",
										MarkdownDescription: "Pod liveness checking.",
										Attributes: map[string]schema.Attribute{
											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"logging": schema.SingleNestedAttribute{
										Description:         "Logging configuration.",
										MarkdownDescription: "Logging configuration.",
										Attributes: map[string]schema.Attribute{
											"loggers": schema.MapAttribute{
												Description:         "A Map from logger name to logger level.",
												MarkdownDescription: "A Map from logger name to logger level.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "Logging type, must be either 'inline' or 'external'.",
												MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("inline", "external"),
												},
											},

											"value_from": schema.SingleNestedAttribute{
												Description:         "'ConfigMap' entry where the logging configuration is stored. ",
												MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",
												Attributes: map[string]schema.Attribute{
													"config_map_key_ref": schema.SingleNestedAttribute{
														Description:         "Reference to the key in the ConfigMap containing the configuration.",
														MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
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
										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_probe": schema.SingleNestedAttribute{
										Description:         "Pod readiness checking.",
										MarkdownDescription: "Pod readiness checking.",
										Attributes: map[string]schema.Attribute{
											"failure_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": schema.Int64Attribute{
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": schema.Int64Attribute{
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"reconciliation_interval_seconds": schema.Int64Attribute{
										Description:         "Interval between periodic reconciliations.",
										MarkdownDescription: "Interval between periodic reconciliations.",
										Required:            false,
										Optional:            true,
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

									"secret_prefix": schema.StringAttribute{
										Description:         "The prefix that will be added to the KafkaUser name to be used as the Secret name.",
										MarkdownDescription: "The prefix that will be added to the KafkaUser name to be used as the Secret name.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"watched_namespace": schema.StringAttribute{
										Description:         "The namespace the User Operator should watch.",
										MarkdownDescription: "The namespace the User Operator should watch.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"zookeeper_session_timeout_seconds": schema.Int64Attribute{
										Description:         "Timeout for the ZooKeeper session.",
										MarkdownDescription: "Timeout for the ZooKeeper session.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
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

					"jmx_trans": schema.SingleNestedAttribute{
						Description:         "As of Strimzi 0.35.0, JMXTrans is not supported anymore and this option is ignored.",
						MarkdownDescription: "As of Strimzi 0.35.0, JMXTrans is not supported anymore and this option is ignored.",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "The image to use for the JmxTrans.",
								MarkdownDescription: "The image to use for the JmxTrans.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kafka_queries": schema.ListNestedAttribute{
								Description:         "Queries to send to the Kafka brokers to define what data should be read from each broker. For more information on these properties see, xref:type-JmxTransQueryTemplate-reference['JmxTransQueryTemplate' schema reference].",
								MarkdownDescription: "Queries to send to the Kafka brokers to define what data should be read from each broker. For more information on these properties see, xref:type-JmxTransQueryTemplate-reference['JmxTransQueryTemplate' schema reference].",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"attributes": schema.ListAttribute{
											Description:         "Determine which attributes of the targeted MBean should be included.",
											MarkdownDescription: "Determine which attributes of the targeted MBean should be included.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"outputs": schema.ListAttribute{
											Description:         "List of the names of output definitions specified in the spec.kafka.jmxTrans.outputDefinitions that have defined where JMX metrics are pushed to, and in which data format.",
											MarkdownDescription: "List of the names of output definitions specified in the spec.kafka.jmxTrans.outputDefinitions that have defined where JMX metrics are pushed to, and in which data format.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"target_m_bean": schema.StringAttribute{
											Description:         "If using wildcards instead of a specific MBean then the data is gathered from multiple MBeans. Otherwise if specifying an MBean then data is gathered from that specified MBean.",
											MarkdownDescription: "If using wildcards instead of a specific MBean then the data is gathered from multiple MBeans. Otherwise if specifying an MBean then data is gathered from that specified MBean.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"log_level": schema.StringAttribute{
								Description:         "Sets the logging level of the JmxTrans deployment.For more information see, https://github.com/jmxtrans/jmxtrans-agent/wiki/Troubleshooting[JmxTrans Logging Level].",
								MarkdownDescription: "Sets the logging level of the JmxTrans deployment.For more information see, https://github.com/jmxtrans/jmxtrans-agent/wiki/Troubleshooting[JmxTrans Logging Level].",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"output_definitions": schema.ListNestedAttribute{
								Description:         "Defines the output hosts that will be referenced later on. For more information on these properties see, xref:type-JmxTransOutputDefinitionTemplate-reference['JmxTransOutputDefinitionTemplate' schema reference].",
								MarkdownDescription: "Defines the output hosts that will be referenced later on. For more information on these properties see, xref:type-JmxTransOutputDefinitionTemplate-reference['JmxTransOutputDefinitionTemplate' schema reference].",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"flush_delay_in_seconds": schema.Int64Attribute{
											Description:         "How many seconds the JmxTrans waits before pushing a new set of data out.",
											MarkdownDescription: "How many seconds the JmxTrans waits before pushing a new set of data out.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"host": schema.StringAttribute{
											Description:         "The DNS/hostname of the remote host that the data is pushed to.",
											MarkdownDescription: "The DNS/hostname of the remote host that the data is pushed to.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Template for setting the name of the output definition. This is used to identify where to send the results of queries should be sent.",
											MarkdownDescription: "Template for setting the name of the output definition. This is used to identify where to send the results of queries should be sent.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"output_type": schema.StringAttribute{
											Description:         "Template for setting the format of the data that will be pushed.For more information see https://github.com/jmxtrans/jmxtrans/wiki/OutputWriters[JmxTrans OutputWriters].",
											MarkdownDescription: "Template for setting the format of the data that will be pushed.For more information see https://github.com/jmxtrans/jmxtrans/wiki/OutputWriters[JmxTrans OutputWriters].",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "The port of the remote host that the data is pushed to.",
											MarkdownDescription: "The port of the remote host that the data is pushed to.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"type_names": schema.ListAttribute{
											Description:         "Template for filtering data to be included in response to a wildcard query. For more information see https://github.com/jmxtrans/jmxtrans/wiki/Queries[JmxTrans queries].",
											MarkdownDescription: "Template for filtering data to be included in response to a wildcard query. For more information see https://github.com/jmxtrans/jmxtrans/wiki/Queries[JmxTrans queries].",
											ElementType:         types.StringType,
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

							"template": schema.SingleNestedAttribute{
								Description:         "Template for JmxTrans resources.",
								MarkdownDescription: "Template for JmxTrans resources.",
								Attributes: map[string]schema.Attribute{
									"container": schema.SingleNestedAttribute{
										Description:         "Template for JmxTrans container.",
										MarkdownDescription: "Template for JmxTrans container.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment": schema.SingleNestedAttribute{
										Description:         "Template for JmxTrans 'Deployment'.",
										MarkdownDescription: "Template for JmxTrans 'Deployment'.",
										Attributes: map[string]schema.Attribute{
											"deployment_strategy": schema.StringAttribute{
												Description:         "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
												MarkdownDescription: "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("RollingUpdate", "Recreate"),
												},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod": schema.SingleNestedAttribute{
										Description:         "Template for JmxTrans 'Pods'.",
										MarkdownDescription: "Template for JmxTrans 'Pods'.",
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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
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
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_account": schema.SingleNestedAttribute{
										Description:         "Template for the JmxTrans service account.",
										MarkdownDescription: "Template for the JmxTrans service account.",
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

					"kafka": schema.SingleNestedAttribute{
						Description:         "Configuration of the Kafka cluster.",
						MarkdownDescription: "Configuration of the Kafka cluster.",
						Attributes: map[string]schema.Attribute{
							"authorization": schema.SingleNestedAttribute{
								Description:         "Authorization configuration for Kafka brokers.",
								MarkdownDescription: "Authorization configuration for Kafka brokers.",
								Attributes: map[string]schema.Attribute{
									"allow_on_error": schema.BoolAttribute{
										Description:         "Defines whether a Kafka client should be allowed or denied by default when the authorizer fails to query the Open Policy Agent, for example, when it is temporarily unavailable). Defaults to 'false' - all actions will be denied.",
										MarkdownDescription: "Defines whether a Kafka client should be allowed or denied by default when the authorizer fails to query the Open Policy Agent, for example, when it is temporarily unavailable). Defaults to 'false' - all actions will be denied.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"authorizer_class": schema.StringAttribute{
										Description:         "Authorization implementation class, which must be available in classpath.",
										MarkdownDescription: "Authorization implementation class, which must be available in classpath.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_id": schema.StringAttribute{
										Description:         "OAuth Client ID which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",
										MarkdownDescription: "OAuth Client ID which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connect_timeout_seconds": schema.Int64Attribute{
										Description:         "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",
										MarkdownDescription: "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"delegate_to_kafka_acls": schema.BoolAttribute{
										Description:         "Whether authorization decision should be delegated to the 'Simple' authorizer if DENIED by Keycloak Authorization Services policies. Default value is 'false'.",
										MarkdownDescription: "Whether authorization decision should be delegated to the 'Simple' authorizer if DENIED by Keycloak Authorization Services policies. Default value is 'false'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_tls_hostname_verification": schema.BoolAttribute{
										Description:         "Enable or disable TLS hostname verification. Default value is 'false'.",
										MarkdownDescription: "Enable or disable TLS hostname verification. Default value is 'false'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable_metrics": schema.BoolAttribute{
										Description:         "Enable or disable OAuth metrics. The default value is 'false'.",
										MarkdownDescription: "Enable or disable OAuth metrics. The default value is 'false'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"expire_after_ms": schema.Int64Attribute{
										Description:         "The expiration of the records kept in the local cache to avoid querying the Open Policy Agent for every request. Defines how often the cached authorization decisions are reloaded from the Open Policy Agent server. In milliseconds. Defaults to '3600000'.",
										MarkdownDescription: "The expiration of the records kept in the local cache to avoid querying the Open Policy Agent for every request. Defines how often the cached authorization decisions are reloaded from the Open Policy Agent server. In milliseconds. Defaults to '3600000'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"grants_always_latest": schema.BoolAttribute{
										Description:         "Controls whether the latest grants are fetched for a new session. When enabled, grants are retrieved from Keycloak and cached for the user. The default value is 'false'.",
										MarkdownDescription: "Controls whether the latest grants are fetched for a new session. When enabled, grants are retrieved from Keycloak and cached for the user. The default value is 'false'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"grants_gc_period_seconds": schema.Int64Attribute{
										Description:         "The time, in seconds, between consecutive runs of a job that cleans stale grants from the cache. The default value is 300.",
										MarkdownDescription: "The time, in seconds, between consecutive runs of a job that cleans stale grants from the cache. The default value is 300.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"grants_max_idle_time_seconds": schema.Int64Attribute{
										Description:         "The time, in seconds, after which an idle grant can be evicted from the cache. The default value is 300.",
										MarkdownDescription: "The time, in seconds, after which an idle grant can be evicted from the cache. The default value is 300.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"grants_refresh_period_seconds": schema.Int64Attribute{
										Description:         "The time between two consecutive grants refresh runs in seconds. The default value is 60.",
										MarkdownDescription: "The time between two consecutive grants refresh runs in seconds. The default value is 60.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"grants_refresh_pool_size": schema.Int64Attribute{
										Description:         "The number of threads to use to refresh grants for active sessions. The more threads, the more parallelism, so the sooner the job completes. However, using more threads places a heavier load on the authorization server. The default value is 5.",
										MarkdownDescription: "The number of threads to use to refresh grants for active sessions. The more threads, the more parallelism, so the sooner the job completes. However, using more threads places a heavier load on the authorization server. The default value is 5.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"http_retries": schema.Int64Attribute{
										Description:         "The maximum number of retries to attempt if an initial HTTP request fails. If not set, the default is to not attempt any retries.",
										MarkdownDescription: "The maximum number of retries to attempt if an initial HTTP request fails. If not set, the default is to not attempt any retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"initial_cache_capacity": schema.Int64Attribute{
										Description:         "Initial capacity of the local cache used by the authorizer to avoid querying the Open Policy Agent for every request Defaults to '5000'.",
										MarkdownDescription: "Initial capacity of the local cache used by the authorizer to avoid querying the Open Policy Agent for every request Defaults to '5000'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"maximum_cache_size": schema.Int64Attribute{
										Description:         "Maximum capacity of the local cache used by the authorizer to avoid querying the Open Policy Agent for every request. Defaults to '50000'.",
										MarkdownDescription: "Maximum capacity of the local cache used by the authorizer to avoid querying the Open Policy Agent for every request. Defaults to '50000'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"read_timeout_seconds": schema.Int64Attribute{
										Description:         "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",
										MarkdownDescription: "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"super_users": schema.ListAttribute{
										Description:         "List of super users, which are user principals with unlimited access rights.",
										MarkdownDescription: "List of super users, which are user principals with unlimited access rights.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"supports_admin_api": schema.BoolAttribute{
										Description:         "Indicates whether the custom authorizer supports the APIs for managing ACLs using the Kafka Admin API. Defaults to 'false'.",
										MarkdownDescription: "Indicates whether the custom authorizer supports the APIs for managing ACLs using the Kafka Admin API. Defaults to 'false'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls_trusted_certificates": schema.ListNestedAttribute{
										Description:         "Trusted certificates for TLS connection to the OAuth server.",
										MarkdownDescription: "Trusted certificates for TLS connection to the OAuth server.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"certificate": schema.StringAttribute{
													Description:         "The name of the file certificate in the Secret.",
													MarkdownDescription: "The name of the file certificate in the Secret.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "The name of the Secret containing the certificate.",
													MarkdownDescription: "The name of the Secret containing the certificate.",
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

									"token_endpoint_uri": schema.StringAttribute{
										Description:         "Authorization server token endpoint URI.",
										MarkdownDescription: "Authorization server token endpoint URI.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Authorization type. Currently, the supported types are 'simple', 'keycloak', 'opa' and 'custom'. 'simple' authorization type uses Kafka's 'kafka.security.authorizer.AclAuthorizer' class for authorization. 'keycloak' authorization type uses Keycloak Authorization Services for authorization. 'opa' authorization type uses Open Policy Agent based authorization.'custom' authorization type uses user-provided implementation for authorization.",
										MarkdownDescription: "Authorization type. Currently, the supported types are 'simple', 'keycloak', 'opa' and 'custom'. 'simple' authorization type uses Kafka's 'kafka.security.authorizer.AclAuthorizer' class for authorization. 'keycloak' authorization type uses Keycloak Authorization Services for authorization. 'opa' authorization type uses Open Policy Agent based authorization.'custom' authorization type uses user-provided implementation for authorization.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("simple", "opa", "keycloak", "custom"),
										},
									},

									"url": schema.StringAttribute{
										Description:         "The URL used to connect to the Open Policy Agent server. The URL has to include the policy which will be queried by the authorizer. This option is required.",
										MarkdownDescription: "The URL used to connect to the Open Policy Agent server. The URL has to include the policy which will be queried by the authorizer. This option is required.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"broker_rack_init_image": schema.StringAttribute{
								Description:         "The image of the init container used for initializing the 'broker.rack'.",
								MarkdownDescription: "The image of the init container used for initializing the 'broker.rack'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config": schema.MapAttribute{
								Description:         "Kafka broker config properties with the following prefixes cannot be set: listeners, advertised., broker., listener., host.name, port, inter.broker.listener.name, sasl., ssl., security., password., log.dir, zookeeper.connect, zookeeper.set.acl, zookeeper.ssl, zookeeper.clientCnxnSocket, authorizer., super.user, cruise.control.metrics.topic, cruise.control.metrics.reporter.bootstrap.servers,node.id, process.roles, controller. (with the exception of: zookeeper.connection.timeout.ms, sasl.server.max.receive.size,ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols, ssl.secure.random.implementation,cruise.control.metrics.topic.num.partitions, cruise.control.metrics.topic.replication.factor, cruise.control.metrics.topic.retention.ms,cruise.control.metrics.topic.auto.create.retries, cruise.control.metrics.topic.auto.create.timeout.ms,cruise.control.metrics.topic.min.insync.replicas,controller.quorum.election.backoff.max.ms, controller.quorum.election.timeout.ms, controller.quorum.fetch.timeout.ms).",
								MarkdownDescription: "Kafka broker config properties with the following prefixes cannot be set: listeners, advertised., broker., listener., host.name, port, inter.broker.listener.name, sasl., ssl., security., password., log.dir, zookeeper.connect, zookeeper.set.acl, zookeeper.ssl, zookeeper.clientCnxnSocket, authorizer., super.user, cruise.control.metrics.topic, cruise.control.metrics.reporter.bootstrap.servers,node.id, process.roles, controller. (with the exception of: zookeeper.connection.timeout.ms, sasl.server.max.receive.size,ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols, ssl.secure.random.implementation,cruise.control.metrics.topic.num.partitions, cruise.control.metrics.topic.replication.factor, cruise.control.metrics.topic.retention.ms,cruise.control.metrics.topic.auto.create.retries, cruise.control.metrics.topic.auto.create.timeout.ms,cruise.control.metrics.topic.min.insync.replicas,controller.quorum.election.backoff.max.ms, controller.quorum.election.timeout.ms, controller.quorum.fetch.timeout.ms).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "The docker image for the pods. The default value depends on the configured 'Kafka.spec.kafka.version'.",
								MarkdownDescription: "The docker image for the pods. The default value depends on the configured 'Kafka.spec.kafka.version'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jmx_options": schema.SingleNestedAttribute{
								Description:         "JMX Options for Kafka brokers.",
								MarkdownDescription: "JMX Options for Kafka brokers.",
								Attributes: map[string]schema.Attribute{
									"authentication": schema.SingleNestedAttribute{
										Description:         "Authentication configuration for connecting to the JMX port.",
										MarkdownDescription: "Authentication configuration for connecting to the JMX port.",
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												Description:         "Authentication type. Currently the only supported types are 'password'.'password' type creates a username and protected port with no TLS.",
												MarkdownDescription: "Authentication type. Currently the only supported types are 'password'.'password' type creates a username and protected port with no TLS.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("password"),
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

							"listeners": schema.ListNestedAttribute{
								Description:         "Configures listeners of Kafka brokers.",
								MarkdownDescription: "Configures listeners of Kafka brokers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"authentication": schema.SingleNestedAttribute{
											Description:         "Authentication configuration for this listener.",
											MarkdownDescription: "Authentication configuration for this listener.",
											Attributes: map[string]schema.Attribute{
												"access_token_is_jwt": schema.BoolAttribute{
													Description:         "Configure whether the access token is treated as JWT. This must be set to 'false' if the authorization server returns opaque tokens. Defaults to 'true'.",
													MarkdownDescription: "Configure whether the access token is treated as JWT. This must be set to 'false' if the authorization server returns opaque tokens. Defaults to 'true'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"check_access_token_type": schema.BoolAttribute{
													Description:         "Configure whether the access token type check is performed or not. This should be set to 'false' if the authorization server does not include 'typ' claim in JWT token. Defaults to 'true'.",
													MarkdownDescription: "Configure whether the access token type check is performed or not. This should be set to 'false' if the authorization server does not include 'typ' claim in JWT token. Defaults to 'true'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"check_audience": schema.BoolAttribute{
													Description:         "Enable or disable audience checking. Audience checks identify the recipients of tokens. If audience checking is enabled, the OAuth Client ID also has to be configured using the 'clientId' property. The Kafka broker will reject tokens that do not have its 'clientId' in their 'aud' (audience) claim.Default value is 'false'.",
													MarkdownDescription: "Enable or disable audience checking. Audience checks identify the recipients of tokens. If audience checking is enabled, the OAuth Client ID also has to be configured using the 'clientId' property. The Kafka broker will reject tokens that do not have its 'clientId' in their 'aud' (audience) claim.Default value is 'false'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"check_issuer": schema.BoolAttribute{
													Description:         "Enable or disable issuer checking. By default issuer is checked using the value configured by 'validIssuerUri'. Default value is 'true'.",
													MarkdownDescription: "Enable or disable issuer checking. By default issuer is checked using the value configured by 'validIssuerUri'. Default value is 'true'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_audience": schema.StringAttribute{
													Description:         "The audience to use when making requests to the authorization server's token endpoint. Used for inter-broker authentication and for configuring OAuth 2.0 over PLAIN using the 'clientId' and 'secret' method.",
													MarkdownDescription: "The audience to use when making requests to the authorization server's token endpoint. Used for inter-broker authentication and for configuring OAuth 2.0 over PLAIN using the 'clientId' and 'secret' method.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_id": schema.StringAttribute{
													Description:         "OAuth Client ID which the Kafka broker can use to authenticate against the authorization server and use the introspect endpoint URI.",
													MarkdownDescription: "OAuth Client ID which the Kafka broker can use to authenticate against the authorization server and use the introspect endpoint URI.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_scope": schema.StringAttribute{
													Description:         "The scope to use when making requests to the authorization server's token endpoint. Used for inter-broker authentication and for configuring OAuth 2.0 over PLAIN using the 'clientId' and 'secret' method.",
													MarkdownDescription: "The scope to use when making requests to the authorization server's token endpoint. Used for inter-broker authentication and for configuring OAuth 2.0 over PLAIN using the 'clientId' and 'secret' method.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"client_secret": schema.SingleNestedAttribute{
													Description:         "Link to Kubernetes Secret containing the OAuth client secret which the Kafka broker can use to authenticate against the authorization server and use the introspect endpoint URI.",
													MarkdownDescription: "Link to Kubernetes Secret containing the OAuth client secret which the Kafka broker can use to authenticate against the authorization server and use the introspect endpoint URI.",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key under which the secret value is stored in the Kubernetes Secret.",
															MarkdownDescription: "The key under which the secret value is stored in the Kubernetes Secret.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"secret_name": schema.StringAttribute{
															Description:         "The name of the Kubernetes Secret containing the secret value.",
															MarkdownDescription: "The name of the Kubernetes Secret containing the secret value.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"connect_timeout_seconds": schema.Int64Attribute{
													Description:         "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",
													MarkdownDescription: "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"custom_claim_check": schema.StringAttribute{
													Description:         "JsonPath filter query to be applied to the JWT token or to the response of the introspection endpoint for additional token validation. Not set by default.",
													MarkdownDescription: "JsonPath filter query to be applied to the JWT token or to the response of the introspection endpoint for additional token validation. Not set by default.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"disable_tls_hostname_verification": schema.BoolAttribute{
													Description:         "Enable or disable TLS hostname verification. Default value is 'false'.",
													MarkdownDescription: "Enable or disable TLS hostname verification. Default value is 'false'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"enable_ecdsa": schema.BoolAttribute{
													Description:         "Enable or disable ECDSA support by installing BouncyCastle crypto provider. ECDSA support is always enabled. The BouncyCastle libraries are no longer packaged with Strimzi. Value is ignored.",
													MarkdownDescription: "Enable or disable ECDSA support by installing BouncyCastle crypto provider. ECDSA support is always enabled. The BouncyCastle libraries are no longer packaged with Strimzi. Value is ignored.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"enable_metrics": schema.BoolAttribute{
													Description:         "Enable or disable OAuth metrics. Default value is 'false'.",
													MarkdownDescription: "Enable or disable OAuth metrics. Default value is 'false'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"enable_oauth_bearer": schema.BoolAttribute{
													Description:         "Enable or disable OAuth authentication over SASL_OAUTHBEARER. Default value is 'true'.",
													MarkdownDescription: "Enable or disable OAuth authentication over SASL_OAUTHBEARER. Default value is 'true'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"enable_plain": schema.BoolAttribute{
													Description:         "Enable or disable OAuth authentication over SASL_PLAIN. There is no re-authentication support when this mechanism is used. Default value is 'false'.",
													MarkdownDescription: "Enable or disable OAuth authentication over SASL_PLAIN. There is no re-authentication support when this mechanism is used. Default value is 'false'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"fail_fast": schema.BoolAttribute{
													Description:         "Enable or disable termination of Kafka broker processes due to potentially recoverable runtime errors during startup. Default value is 'true'.",
													MarkdownDescription: "Enable or disable termination of Kafka broker processes due to potentially recoverable runtime errors during startup. Default value is 'true'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"fallback_user_name_claim": schema.StringAttribute{
													Description:         "The fallback username claim to be used for the user id if the claim specified by 'userNameClaim' is not present. This is useful when 'client_credentials' authentication only results in the client id being provided in another claim. It only takes effect if 'userNameClaim' is set.",
													MarkdownDescription: "The fallback username claim to be used for the user id if the claim specified by 'userNameClaim' is not present. This is useful when 'client_credentials' authentication only results in the client id being provided in another claim. It only takes effect if 'userNameClaim' is set.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"fallback_user_name_prefix": schema.StringAttribute{
													Description:         "The prefix to use with the value of 'fallbackUserNameClaim' to construct the user id. This only takes effect if 'fallbackUserNameClaim' is true, and the value is present for the claim. Mapping usernames and client ids into the same user id space is useful in preventing name collisions.",
													MarkdownDescription: "The prefix to use with the value of 'fallbackUserNameClaim' to construct the user id. This only takes effect if 'fallbackUserNameClaim' is true, and the value is present for the claim. Mapping usernames and client ids into the same user id space is useful in preventing name collisions.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"groups_claim": schema.StringAttribute{
													Description:         "JsonPath query used to extract groups for the user during authentication. Extracted groups can be used by a custom authorizer. By default no groups are extracted.",
													MarkdownDescription: "JsonPath query used to extract groups for the user during authentication. Extracted groups can be used by a custom authorizer. By default no groups are extracted.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"groups_claim_delimiter": schema.StringAttribute{
													Description:         "A delimiter used to parse groups when they are extracted as a single String value rather than a JSON array. Default value is ',' (comma).",
													MarkdownDescription: "A delimiter used to parse groups when they are extracted as a single String value rather than a JSON array. Default value is ',' (comma).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http_retries": schema.Int64Attribute{
													Description:         "The maximum number of retries to attempt if an initial HTTP request fails. If not set, the default is to not attempt any retries.",
													MarkdownDescription: "The maximum number of retries to attempt if an initial HTTP request fails. If not set, the default is to not attempt any retries.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http_retry_pause_ms": schema.Int64Attribute{
													Description:         "The pause to take before retrying a failed HTTP request. If not set, the default is to not pause at all but to immediately repeat a request.",
													MarkdownDescription: "The pause to take before retrying a failed HTTP request. If not set, the default is to not pause at all but to immediately repeat a request.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"introspection_endpoint_uri": schema.StringAttribute{
													Description:         "URI of the token introspection endpoint which can be used to validate opaque non-JWT tokens.",
													MarkdownDescription: "URI of the token introspection endpoint which can be used to validate opaque non-JWT tokens.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"jwks_endpoint_uri": schema.StringAttribute{
													Description:         "URI of the JWKS certificate endpoint, which can be used for local JWT validation.",
													MarkdownDescription: "URI of the JWKS certificate endpoint, which can be used for local JWT validation.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"jwks_expiry_seconds": schema.Int64Attribute{
													Description:         "Configures how often are the JWKS certificates considered valid. The expiry interval has to be at least 60 seconds longer then the refresh interval specified in 'jwksRefreshSeconds'. Defaults to 360 seconds.",
													MarkdownDescription: "Configures how often are the JWKS certificates considered valid. The expiry interval has to be at least 60 seconds longer then the refresh interval specified in 'jwksRefreshSeconds'. Defaults to 360 seconds.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"jwks_ignore_key_use": schema.BoolAttribute{
													Description:         "Flag to ignore the 'use' attribute of 'key' declarations in a JWKS endpoint response. Default value is 'false'.",
													MarkdownDescription: "Flag to ignore the 'use' attribute of 'key' declarations in a JWKS endpoint response. Default value is 'false'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"jwks_min_refresh_pause_seconds": schema.Int64Attribute{
													Description:         "The minimum pause between two consecutive refreshes. When an unknown signing key is encountered the refresh is scheduled immediately, but will always wait for this minimum pause. Defaults to 1 second.",
													MarkdownDescription: "The minimum pause between two consecutive refreshes. When an unknown signing key is encountered the refresh is scheduled immediately, but will always wait for this minimum pause. Defaults to 1 second.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"jwks_refresh_seconds": schema.Int64Attribute{
													Description:         "Configures how often are the JWKS certificates refreshed. The refresh interval has to be at least 60 seconds shorter then the expiry interval specified in 'jwksExpirySeconds'. Defaults to 300 seconds.",
													MarkdownDescription: "Configures how often are the JWKS certificates refreshed. The refresh interval has to be at least 60 seconds shorter then the expiry interval specified in 'jwksExpirySeconds'. Defaults to 300 seconds.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"listener_config": schema.MapAttribute{
													Description:         "Configuration to be used for a specific listener. All values are prefixed with listener.name._<listener_name>_.",
													MarkdownDescription: "Configuration to be used for a specific listener. All values are prefixed with listener.name._<listener_name>_.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_seconds_without_reauthentication": schema.Int64Attribute{
													Description:         "Maximum number of seconds the authenticated session remains valid without re-authentication. This enables Apache Kafka re-authentication feature, and causes sessions to expire when the access token expires. If the access token expires before max time or if max time is reached, the client has to re-authenticate, otherwise the server will drop the connection. Not set by default - the authenticated session does not expire when the access token expires. This option only applies to SASL_OAUTHBEARER authentication mechanism (when 'enableOauthBearer' is 'true').",
													MarkdownDescription: "Maximum number of seconds the authenticated session remains valid without re-authentication. This enables Apache Kafka re-authentication feature, and causes sessions to expire when the access token expires. If the access token expires before max time or if max time is reached, the client has to re-authenticate, otherwise the server will drop the connection. Not set by default - the authenticated session does not expire when the access token expires. This option only applies to SASL_OAUTHBEARER authentication mechanism (when 'enableOauthBearer' is 'true').",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_timeout_seconds": schema.Int64Attribute{
													Description:         "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",
													MarkdownDescription: "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sasl": schema.BoolAttribute{
													Description:         "Enable or disable SASL on this listener.",
													MarkdownDescription: "Enable or disable SASL on this listener.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secrets": schema.ListNestedAttribute{
													Description:         "Secrets to be mounted to /opt/kafka/custom-authn-secrets/custom-listener-_<listener_name>-<port>_/_<secret_name>_.",
													MarkdownDescription: "Secrets to be mounted to /opt/kafka/custom-authn-secrets/custom-listener-_<listener_name>-<port>_/_<secret_name>_.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "The key under which the secret value is stored in the Kubernetes Secret.",
																MarkdownDescription: "The key under which the secret value is stored in the Kubernetes Secret.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret_name": schema.StringAttribute{
																Description:         "The name of the Kubernetes Secret containing the secret value.",
																MarkdownDescription: "The name of the Kubernetes Secret containing the secret value.",
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

												"tls_trusted_certificates": schema.ListNestedAttribute{
													Description:         "Trusted certificates for TLS connection to the OAuth server.",
													MarkdownDescription: "Trusted certificates for TLS connection to the OAuth server.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"certificate": schema.StringAttribute{
																Description:         "The name of the file certificate in the Secret.",
																MarkdownDescription: "The name of the file certificate in the Secret.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"secret_name": schema.StringAttribute{
																Description:         "The name of the Secret containing the certificate.",
																MarkdownDescription: "The name of the Secret containing the certificate.",
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

												"token_endpoint_uri": schema.StringAttribute{
													Description:         "URI of the Token Endpoint to use with SASL_PLAIN mechanism when the client authenticates with 'clientId' and a 'secret'. If set, the client can authenticate over SASL_PLAIN by either setting 'username' to 'clientId', and setting 'password' to client 'secret', or by setting 'username' to account username, and 'password' to access token prefixed with '$accessToken:'. If this option is not set, the 'password' is always interpreted as an access token (without a prefix), and 'username' as the account username (a so called 'no-client-credentials' mode).",
													MarkdownDescription: "URI of the Token Endpoint to use with SASL_PLAIN mechanism when the client authenticates with 'clientId' and a 'secret'. If set, the client can authenticate over SASL_PLAIN by either setting 'username' to 'clientId', and setting 'password' to client 'secret', or by setting 'username' to account username, and 'password' to access token prefixed with '$accessToken:'. If this option is not set, the 'password' is always interpreted as an access token (without a prefix), and 'username' as the account username (a so called 'no-client-credentials' mode).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "Authentication type. 'oauth' type uses SASL OAUTHBEARER Authentication. 'scram-sha-512' type uses SASL SCRAM-SHA-512 Authentication. 'tls' type uses TLS Client Authentication. 'tls' type is supported only on TLS listeners.'custom' type allows for any authentication type to be used.",
													MarkdownDescription: "Authentication type. 'oauth' type uses SASL OAUTHBEARER Authentication. 'scram-sha-512' type uses SASL SCRAM-SHA-512 Authentication. 'tls' type uses TLS Client Authentication. 'tls' type is supported only on TLS listeners.'custom' type allows for any authentication type to be used.",
													Required:            true,
													Optional:            false,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("tls", "scram-sha-512", "oauth", "custom"),
													},
												},

												"user_info_endpoint_uri": schema.StringAttribute{
													Description:         "URI of the User Info Endpoint to use as a fallback to obtaining the user id when the Introspection Endpoint does not return information that can be used for the user id. ",
													MarkdownDescription: "URI of the User Info Endpoint to use as a fallback to obtaining the user id when the Introspection Endpoint does not return information that can be used for the user id. ",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"user_name_claim": schema.StringAttribute{
													Description:         "Name of the claim from the JWT authentication token, Introspection Endpoint response or User Info Endpoint response which will be used to extract the user id. Defaults to 'sub'.",
													MarkdownDescription: "Name of the claim from the JWT authentication token, Introspection Endpoint response or User Info Endpoint response which will be used to extract the user id. Defaults to 'sub'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"valid_issuer_uri": schema.StringAttribute{
													Description:         "URI of the token issuer used for authentication.",
													MarkdownDescription: "URI of the token issuer used for authentication.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"valid_token_type": schema.StringAttribute{
													Description:         "Valid value for the 'token_type' attribute returned by the Introspection Endpoint. No default value, and not checked by default.",
													MarkdownDescription: "Valid value for the 'token_type' attribute returned by the Introspection Endpoint. No default value, and not checked by default.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"configuration": schema.SingleNestedAttribute{
											Description:         "Additional listener configuration.",
											MarkdownDescription: "Additional listener configuration.",
											Attributes: map[string]schema.Attribute{
												"bootstrap": schema.SingleNestedAttribute{
													Description:         "Bootstrap configuration.",
													MarkdownDescription: "Bootstrap configuration.",
													Attributes: map[string]schema.Attribute{
														"alternative_names": schema.ListAttribute{
															Description:         "Additional alternative names for the bootstrap service. The alternative names will be added to the list of subject alternative names of the TLS certificates.",
															MarkdownDescription: "Additional alternative names for the bootstrap service. The alternative names will be added to the list of subject alternative names of the TLS certificates.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"annotations": schema.MapAttribute{
															Description:         "Annotations that will be added to the 'Ingress', 'Route', or 'Service' resource. You can use this field to configure DNS providers such as External DNS. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",
															MarkdownDescription: "Annotations that will be added to the 'Ingress', 'Route', or 'Service' resource. You can use this field to configure DNS providers such as External DNS. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"host": schema.StringAttribute{
															Description:         "The bootstrap host. This field will be used in the Ingress resource or in the Route resource to specify the desired hostname. This field can be used only with 'route' (optional) or 'ingress' (required) type listeners.",
															MarkdownDescription: "The bootstrap host. This field will be used in the Ingress resource or in the Route resource to specify the desired hostname. This field can be used only with 'route' (optional) or 'ingress' (required) type listeners.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"labels": schema.MapAttribute{
															Description:         "Labels that will be added to the 'Ingress', 'Route', or 'Service' resource. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",
															MarkdownDescription: "Labels that will be added to the 'Ingress', 'Route', or 'Service' resource. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"load_balancer_ip": schema.StringAttribute{
															Description:         "The loadbalancer is requested with the IP address specified in this field. This feature depends on whether the underlying cloud provider supports specifying the 'loadBalancerIP' when a load balancer is created. This field is ignored if the cloud provider does not support the feature.This field can be used only with 'loadbalancer' type listener.",
															MarkdownDescription: "The loadbalancer is requested with the IP address specified in this field. This feature depends on whether the underlying cloud provider supports specifying the 'loadBalancerIP' when a load balancer is created. This field is ignored if the cloud provider does not support the feature.This field can be used only with 'loadbalancer' type listener.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_port": schema.Int64Attribute{
															Description:         "Node port for the bootstrap service. This field can be used only with 'nodeport' type listener.",
															MarkdownDescription: "Node port for the bootstrap service. This field can be used only with 'nodeport' type listener.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"broker_cert_chain_and_key": schema.SingleNestedAttribute{
													Description:         "Reference to the 'Secret' which holds the certificate and private key pair which will be used for this listener. The certificate can optionally contain the whole chain. This field can be used only with listeners with enabled TLS encryption.",
													MarkdownDescription: "Reference to the 'Secret' which holds the certificate and private key pair which will be used for this listener. The certificate can optionally contain the whole chain. This field can be used only with listeners with enabled TLS encryption.",
													Attributes: map[string]schema.Attribute{
														"certificate": schema.StringAttribute{
															Description:         "The name of the file certificate in the Secret.",
															MarkdownDescription: "The name of the file certificate in the Secret.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "The name of the private key in the Secret.",
															MarkdownDescription: "The name of the private key in the Secret.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"secret_name": schema.StringAttribute{
															Description:         "The name of the Secret containing the certificate.",
															MarkdownDescription: "The name of the Secret containing the certificate.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"brokers": schema.ListNestedAttribute{
													Description:         "Per-broker configurations.",
													MarkdownDescription: "Per-broker configurations.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"advertised_host": schema.StringAttribute{
																Description:         "The host name used in the brokers' 'advertised.listeners'.",
																MarkdownDescription: "The host name used in the brokers' 'advertised.listeners'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"advertised_port": schema.Int64Attribute{
																Description:         "The port number used in the brokers' 'advertised.listeners'.",
																MarkdownDescription: "The port number used in the brokers' 'advertised.listeners'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"annotations": schema.MapAttribute{
																Description:         "Annotations that will be added to the 'Ingress' or 'Service' resource. You can use this field to configure DNS providers such as External DNS. This field can be used only with 'loadbalancer', 'nodeport', or 'ingress' type listeners.",
																MarkdownDescription: "Annotations that will be added to the 'Ingress' or 'Service' resource. You can use this field to configure DNS providers such as External DNS. This field can be used only with 'loadbalancer', 'nodeport', or 'ingress' type listeners.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"broker": schema.Int64Attribute{
																Description:         "ID of the kafka broker (broker identifier). Broker IDs start from 0 and correspond to the number of broker replicas.",
																MarkdownDescription: "ID of the kafka broker (broker identifier). Broker IDs start from 0 and correspond to the number of broker replicas.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"host": schema.StringAttribute{
																Description:         "The broker host. This field will be used in the Ingress resource or in the Route resource to specify the desired hostname. This field can be used only with 'route' (optional) or 'ingress' (required) type listeners.",
																MarkdownDescription: "The broker host. This field will be used in the Ingress resource or in the Route resource to specify the desired hostname. This field can be used only with 'route' (optional) or 'ingress' (required) type listeners.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"labels": schema.MapAttribute{
																Description:         "Labels that will be added to the 'Ingress', 'Route', or 'Service' resource. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",
																MarkdownDescription: "Labels that will be added to the 'Ingress', 'Route', or 'Service' resource. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"load_balancer_ip": schema.StringAttribute{
																Description:         "The loadbalancer is requested with the IP address specified in this field. This feature depends on whether the underlying cloud provider supports specifying the 'loadBalancerIP' when a load balancer is created. This field is ignored if the cloud provider does not support the feature.This field can be used only with 'loadbalancer' type listener.",
																MarkdownDescription: "The loadbalancer is requested with the IP address specified in this field. This feature depends on whether the underlying cloud provider supports specifying the 'loadBalancerIP' when a load balancer is created. This field is ignored if the cloud provider does not support the feature.This field can be used only with 'loadbalancer' type listener.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"node_port": schema.Int64Attribute{
																Description:         "Node port for the per-broker service. This field can be used only with 'nodeport' type listener.",
																MarkdownDescription: "Node port for the per-broker service. This field can be used only with 'nodeport' type listener.",
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

												"class": schema.StringAttribute{
													Description:         "Configures a specific class for 'Ingress' and 'LoadBalancer' that defines which controller will be used. This field can only be used with 'ingress' and 'loadbalancer' type listeners. If not specified, the default controller is used. For an 'ingress' listener, set the 'ingressClassName' property in the 'Ingress' resources. For a 'loadbalancer' listener, set the 'loadBalancerClass' property  in the 'Service' resources.",
													MarkdownDescription: "Configures a specific class for 'Ingress' and 'LoadBalancer' that defines which controller will be used. This field can only be used with 'ingress' and 'loadbalancer' type listeners. If not specified, the default controller is used. For an 'ingress' listener, set the 'ingressClassName' property in the 'Ingress' resources. For a 'loadbalancer' listener, set the 'loadBalancerClass' property  in the 'Service' resources.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"create_bootstrap_service": schema.BoolAttribute{
													Description:         "Whether to create the bootstrap service or not. The bootstrap service is created by default (if not specified differently). This field can be used with the 'loadBalancer' type listener.",
													MarkdownDescription: "Whether to create the bootstrap service or not. The bootstrap service is created by default (if not specified differently). This field can be used with the 'loadBalancer' type listener.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"external_traffic_policy": schema.StringAttribute{
													Description:         "Specifies whether the service routes external traffic to node-local or cluster-wide endpoints. 'Cluster' may cause a second hop to another node and obscures the client source IP. 'Local' avoids a second hop for LoadBalancer and Nodeport type services and preserves the client source IP (when supported by the infrastructure). If unspecified, Kubernetes will use 'Cluster' as the default.This field can be used only with 'loadbalancer' or 'nodeport' type listener.",
													MarkdownDescription: "Specifies whether the service routes external traffic to node-local or cluster-wide endpoints. 'Cluster' may cause a second hop to another node and obscures the client source IP. 'Local' avoids a second hop for LoadBalancer and Nodeport type services and preserves the client source IP (when supported by the infrastructure). If unspecified, Kubernetes will use 'Cluster' as the default.This field can be used only with 'loadbalancer' or 'nodeport' type listener.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Local", "Cluster"),
													},
												},

												"finalizers": schema.ListAttribute{
													Description:         "A list of finalizers which will be configured for the 'LoadBalancer' type Services created for this listener. If supported by the platform, the finalizer 'service.kubernetes.io/load-balancer-cleanup' to make sure that the external load balancer is deleted together with the service.For more information, see https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#garbage-collecting-load-balancers. This field can be used only with 'loadbalancer' type listeners.",
													MarkdownDescription: "A list of finalizers which will be configured for the 'LoadBalancer' type Services created for this listener. If supported by the platform, the finalizer 'service.kubernetes.io/load-balancer-cleanup' to make sure that the external load balancer is deleted together with the service.For more information, see https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#garbage-collecting-load-balancers. This field can be used only with 'loadbalancer' type listeners.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_families": schema.ListAttribute{
													Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
													MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"ip_family_policy": schema.StringAttribute{
													Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
													MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
													},
												},

												"load_balancer_source_ranges": schema.ListAttribute{
													Description:         "A list of CIDR ranges (for example '10.0.0.0/8' or '130.211.204.1/32') from which clients can connect to load balancer type listeners. If supported by the platform, traffic through the loadbalancer is restricted to the specified CIDR ranges. This field is applicable only for loadbalancer type services and is ignored if the cloud provider does not support the feature. This field can be used only with 'loadbalancer' type listener.",
													MarkdownDescription: "A list of CIDR ranges (for example '10.0.0.0/8' or '130.211.204.1/32') from which clients can connect to load balancer type listeners. If supported by the platform, traffic through the loadbalancer is restricted to the specified CIDR ranges. This field is applicable only for loadbalancer type services and is ignored if the cloud provider does not support the feature. This field can be used only with 'loadbalancer' type listener.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_connection_creation_rate": schema.Int64Attribute{
													Description:         "The maximum connection creation rate we allow in this listener at any time. New connections will be throttled if the limit is reached.",
													MarkdownDescription: "The maximum connection creation rate we allow in this listener at any time. New connections will be throttled if the limit is reached.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"max_connections": schema.Int64Attribute{
													Description:         "The maximum number of connections we allow for this listener in the broker at any time. New connections are blocked if the limit is reached.",
													MarkdownDescription: "The maximum number of connections we allow for this listener in the broker at any time. New connections are blocked if the limit is reached.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"preferred_node_port_address_type": schema.StringAttribute{
													Description:         "Defines which address type should be used as the node address. Available types are: 'ExternalDNS', 'ExternalIP', 'InternalDNS', 'InternalIP' and 'Hostname'. By default, the addresses will be used in the following order (the first one found will be used):* 'ExternalDNS'* 'ExternalIP'* 'InternalDNS'* 'InternalIP'* 'Hostname'This field is used to select the preferred address type, which is checked first. If no address is found for this address type, the other types are checked in the default order. This field can only be used with 'nodeport' type listener.",
													MarkdownDescription: "Defines which address type should be used as the node address. Available types are: 'ExternalDNS', 'ExternalIP', 'InternalDNS', 'InternalIP' and 'Hostname'. By default, the addresses will be used in the following order (the first one found will be used):* 'ExternalDNS'* 'ExternalIP'* 'InternalDNS'* 'InternalIP'* 'Hostname'This field is used to select the preferred address type, which is checked first. If no address is found for this address type, the other types are checked in the default order. This field can only be used with 'nodeport' type listener.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("ExternalIP", "ExternalDNS", "InternalIP", "InternalDNS", "Hostname"),
													},
												},

												"use_service_dns_domain": schema.BoolAttribute{
													Description:         "Configures whether the Kubernetes service DNS domain should be used or not. If set to 'true', the generated addresses will contain the service DNS domain suffix (by default '.cluster.local', can be configured using environment variable 'KUBERNETES_SERVICE_DNS_DOMAIN'). Defaults to 'false'.This field can be used only with 'internal' and 'cluster-ip' type listeners.",
													MarkdownDescription: "Configures whether the Kubernetes service DNS domain should be used or not. If set to 'true', the generated addresses will contain the service DNS domain suffix (by default '.cluster.local', can be configured using environment variable 'KUBERNETES_SERVICE_DNS_DOMAIN'). Defaults to 'false'.This field can be used only with 'internal' and 'cluster-ip' type listeners.",
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
											Description:         "Name of the listener. The name will be used to identify the listener and the related Kubernetes objects. The name has to be unique within given a Kafka cluster. The name can consist of lowercase characters and numbers and be up to 11 characters long.",
											MarkdownDescription: "Name of the listener. The name will be used to identify the listener and the related Kubernetes objects. The name has to be unique within given a Kafka cluster. The name can consist of lowercase characters and numbers and be up to 11 characters long.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]{1,11}$`), ""),
											},
										},

										"network_policy_peers": schema.ListNestedAttribute{
											Description:         "List of peers which should be able to connect to this listener. Peers in this list are combined using a logical OR operation. If this field is empty or missing, all connections will be allowed for this listener. If this field is present and contains at least one item, the listener only allows the traffic which matches at least one item in this list.",
											MarkdownDescription: "List of peers which should be able to connect to this listener. Peers in this list are combined using a logical OR operation. If this field is empty or missing, all connections will be allowed for this listener. If this field is present and contains at least one item, the listener only allows the traffic which matches at least one item in this list.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"ip_block": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"cidr": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"except": schema.ListAttribute{
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

													"pod_selector": schema.SingleNestedAttribute{
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
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"port": schema.Int64Attribute{
											Description:         "Port number used by the listener inside Kafka. The port number has to be unique within a given Kafka cluster. Allowed port numbers are 9092 and higher with the exception of ports 9404 and 9999, which are already used for Prometheus and JMX. Depending on the listener type, the port number might not be the same as the port number that connects Kafka clients.",
											MarkdownDescription: "Port number used by the listener inside Kafka. The port number has to be unique within a given Kafka cluster. Allowed port numbers are 9092 and higher with the exception of ports 9404 and 9999, which are already used for Prometheus and JMX. Depending on the listener type, the port number might not be the same as the port number that connects Kafka clients.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(9092),
											},
										},

										"tls": schema.BoolAttribute{
											Description:         "Enables TLS encryption on the listener. This is a required property.",
											MarkdownDescription: "Enables TLS encryption on the listener. This is a required property.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"type": schema.StringAttribute{
											Description:         "Type of the listener. Currently the supported types are 'internal', 'route', 'loadbalancer', 'nodeport' and 'ingress'. * 'internal' type exposes Kafka internally only within the Kubernetes cluster.* 'route' type uses OpenShift Routes to expose Kafka.* 'loadbalancer' type uses LoadBalancer type services to expose Kafka.* 'nodeport' type uses NodePort type services to expose Kafka.* 'ingress' type uses Kubernetes Nginx Ingress to expose Kafka with TLS passthrough.* 'cluster-ip' type uses a per-broker 'ClusterIP' service.",
											MarkdownDescription: "Type of the listener. Currently the supported types are 'internal', 'route', 'loadbalancer', 'nodeport' and 'ingress'. * 'internal' type exposes Kafka internally only within the Kubernetes cluster.* 'route' type uses OpenShift Routes to expose Kafka.* 'loadbalancer' type uses LoadBalancer type services to expose Kafka.* 'nodeport' type uses NodePort type services to expose Kafka.* 'ingress' type uses Kubernetes Nginx Ingress to expose Kafka with TLS passthrough.* 'cluster-ip' type uses a per-broker 'ClusterIP' service.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("internal", "route", "loadbalancer", "nodeport", "ingress", "cluster-ip"),
											},
										},
									},
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"liveness_probe": schema.SingleNestedAttribute{
								Description:         "Pod liveness checking.",
								MarkdownDescription: "Pod liveness checking.",
								Attributes: map[string]schema.Attribute{
									"failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": schema.Int64Attribute{
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"period_seconds": schema.Int64Attribute{
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": schema.SingleNestedAttribute{
								Description:         "Logging configuration for Kafka.",
								MarkdownDescription: "Logging configuration for Kafka.",
								Attributes: map[string]schema.Attribute{
									"loggers": schema.MapAttribute{
										Description:         "A Map from logger name to logger level.",
										MarkdownDescription: "A Map from logger name to logger level.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Logging type, must be either 'inline' or 'external'.",
										MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("inline", "external"),
										},
									},

									"value_from": schema.SingleNestedAttribute{
										Description:         "'ConfigMap' entry where the logging configuration is stored. ",
										MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics_config": schema.SingleNestedAttribute{
								Description:         "Metrics configuration.",
								MarkdownDescription: "Metrics configuration.",
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Description:         "Metrics type. Only 'jmxPrometheusExporter' supported currently.",
										MarkdownDescription: "Metrics type. Only 'jmxPrometheusExporter' supported currently.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("jmxPrometheusExporter"),
										},
									},

									"value_from": schema.SingleNestedAttribute{
										Description:         "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",
										MarkdownDescription: "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
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
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rack": schema.SingleNestedAttribute{
								Description:         "Configuration of the 'broker.rack' broker config.",
								MarkdownDescription: "Configuration of the 'broker.rack' broker config.",
								Attributes: map[string]schema.Attribute{
									"topology_key": schema.StringAttribute{
										Description:         "A key that matches labels assigned to the Kubernetes cluster nodes. The value of the label is used to set a broker's 'broker.rack' config, and the 'client.rack' config for Kafka Connect or MirrorMaker 2.",
										MarkdownDescription: "A key that matches labels assigned to the Kubernetes cluster nodes. The value of the label is used to set a broker's 'broker.rack' config, and the 'client.rack' config for Kafka Connect or MirrorMaker 2.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"readiness_probe": schema.SingleNestedAttribute{
								Description:         "Pod readiness checking.",
								MarkdownDescription: "Pod readiness checking.",
								Attributes: map[string]schema.Attribute{
									"failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": schema.Int64Attribute{
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"period_seconds": schema.Int64Attribute{
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "The number of pods in the cluster.",
								MarkdownDescription: "The number of pods in the cluster.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
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

									"overrides": schema.ListNestedAttribute{
										Description:         "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",
										MarkdownDescription: "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",
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
										Description:         "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",
										MarkdownDescription: "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",
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
													Description:         "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",
													MarkdownDescription: "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(0),
													},
												},

												"overrides": schema.ListNestedAttribute{
													Description:         "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",
													MarkdownDescription: "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",
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
													Description:         "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",
													MarkdownDescription: "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",
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
								Description:         "Template for Kafka cluster resources. The template allows users to specify how the Kubernetes resources are generated.",
								MarkdownDescription: "Template for Kafka cluster resources. The template allows users to specify how the Kubernetes resources are generated.",
								Attributes: map[string]schema.Attribute{
									"bootstrap_service": schema.SingleNestedAttribute{
										Description:         "Template for Kafka bootstrap 'Service'.",
										MarkdownDescription: "Template for Kafka bootstrap 'Service'.",
										Attributes: map[string]schema.Attribute{
											"ip_families": schema.ListAttribute{
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"brokers_service": schema.SingleNestedAttribute{
										Description:         "Template for Kafka broker 'Service'.",
										MarkdownDescription: "Template for Kafka broker 'Service'.",
										Attributes: map[string]schema.Attribute{
											"ip_families": schema.ListAttribute{
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cluster_ca_cert": schema.SingleNestedAttribute{
										Description:         "Template for Secret with Kafka Cluster certificate public key.",
										MarkdownDescription: "Template for Secret with Kafka Cluster certificate public key.",
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

									"cluster_role_binding": schema.SingleNestedAttribute{
										Description:         "Template for the Kafka ClusterRoleBinding.",
										MarkdownDescription: "Template for the Kafka ClusterRoleBinding.",
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

									"external_bootstrap_ingress": schema.SingleNestedAttribute{
										Description:         "Template for Kafka external bootstrap 'Ingress'.",
										MarkdownDescription: "Template for Kafka external bootstrap 'Ingress'.",
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

									"external_bootstrap_route": schema.SingleNestedAttribute{
										Description:         "Template for Kafka external bootstrap 'Route'.",
										MarkdownDescription: "Template for Kafka external bootstrap 'Route'.",
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

									"external_bootstrap_service": schema.SingleNestedAttribute{
										Description:         "Template for Kafka external bootstrap 'Service'.",
										MarkdownDescription: "Template for Kafka external bootstrap 'Service'.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"jmx_secret": schema.SingleNestedAttribute{
										Description:         "Template for Secret of the Kafka Cluster JMX authentication.",
										MarkdownDescription: "Template for Secret of the Kafka Cluster JMX authentication.",
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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
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
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_disruption_budget": schema.SingleNestedAttribute{
										Description:         "Template for Kafka 'PodDisruptionBudget'.",
										MarkdownDescription: "Template for Kafka 'PodDisruptionBudget'.",
										Attributes: map[string]schema.Attribute{
											"max_unavailable": schema.Int64Attribute{
												Description:         "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
												MarkdownDescription: "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"metadata": schema.SingleNestedAttribute{
												Description:         "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
												MarkdownDescription: "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
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

									"service_account": schema.SingleNestedAttribute{
										Description:         "Template for the Kafka service account.",
										MarkdownDescription: "Template for the Kafka service account.",
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

									"statefulset": schema.SingleNestedAttribute{
										Description:         "Template for Kafka 'StatefulSet'.",
										MarkdownDescription: "Template for Kafka 'StatefulSet'.",
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

											"pod_management_policy": schema.StringAttribute{
												Description:         "PodManagementPolicy which will be used for this StatefulSet. Valid values are 'Parallel' and 'OrderedReady'. Defaults to 'Parallel'.",
												MarkdownDescription: "PodManagementPolicy which will be used for this StatefulSet. Valid values are 'Parallel' and 'OrderedReady'. Defaults to 'Parallel'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("OrderedReady", "Parallel"),
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

							"version": schema.StringAttribute{
								Description:         "The kafka broker version. Defaults to {DefaultKafkaVersion}. Consult the user documentation to understand the process required to upgrade or downgrade the version.",
								MarkdownDescription: "The kafka broker version. Defaults to {DefaultKafkaVersion}. Consult the user documentation to understand the process required to upgrade or downgrade the version.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"kafka_exporter": schema.SingleNestedAttribute{
						Description:         "Configuration of the Kafka Exporter. Kafka Exporter can provide additional metrics, for example lag of consumer group at topic/partition.",
						MarkdownDescription: "Configuration of the Kafka Exporter. Kafka Exporter can provide additional metrics, for example lag of consumer group at topic/partition.",
						Attributes: map[string]schema.Attribute{
							"enable_sarama_logging": schema.BoolAttribute{
								Description:         "Enable Sarama logging, a Go client library used by the Kafka Exporter.",
								MarkdownDescription: "Enable Sarama logging, a Go client library used by the Kafka Exporter.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_exclude_regex": schema.StringAttribute{
								Description:         "Regular expression to specify which consumer groups to exclude.",
								MarkdownDescription: "Regular expression to specify which consumer groups to exclude.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"group_regex": schema.StringAttribute{
								Description:         "Regular expression to specify which consumer groups to collect. Default value is '.*'.",
								MarkdownDescription: "Regular expression to specify which consumer groups to collect. Default value is '.*'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "The docker image for the pods.",
								MarkdownDescription: "The docker image for the pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"liveness_probe": schema.SingleNestedAttribute{
								Description:         "Pod liveness check.",
								MarkdownDescription: "Pod liveness check.",
								Attributes: map[string]schema.Attribute{
									"failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": schema.Int64Attribute{
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"period_seconds": schema.Int64Attribute{
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": schema.StringAttribute{
								Description:         "Only log messages with the given severity or above. Valid levels: ['info', 'debug', 'trace']. Default log level is 'info'.",
								MarkdownDescription: "Only log messages with the given severity or above. Valid levels: ['info', 'debug', 'trace']. Default log level is 'info'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"readiness_probe": schema.SingleNestedAttribute{
								Description:         "Pod readiness check.",
								MarkdownDescription: "Pod readiness check.",
								Attributes: map[string]schema.Attribute{
									"failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": schema.Int64Attribute{
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"period_seconds": schema.Int64Attribute{
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
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

							"template": schema.SingleNestedAttribute{
								Description:         "Customization of deployment templates and pods.",
								MarkdownDescription: "Customization of deployment templates and pods.",
								Attributes: map[string]schema.Attribute{
									"container": schema.SingleNestedAttribute{
										Description:         "Template for the Kafka Exporter container.",
										MarkdownDescription: "Template for the Kafka Exporter container.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment": schema.SingleNestedAttribute{
										Description:         "Template for Kafka Exporter 'Deployment'.",
										MarkdownDescription: "Template for Kafka Exporter 'Deployment'.",
										Attributes: map[string]schema.Attribute{
											"deployment_strategy": schema.StringAttribute{
												Description:         "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
												MarkdownDescription: "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("RollingUpdate", "Recreate"),
												},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod": schema.SingleNestedAttribute{
										Description:         "Template for Kafka Exporter 'Pods'.",
										MarkdownDescription: "Template for Kafka Exporter 'Pods'.",
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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
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
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"service": schema.SingleNestedAttribute{
										Description:         "Template for Kafka Exporter 'Service'.",
										MarkdownDescription: "Template for Kafka Exporter 'Service'.",
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

									"service_account": schema.SingleNestedAttribute{
										Description:         "Template for the Kafka Exporter service account.",
										MarkdownDescription: "Template for the Kafka Exporter service account.",
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

							"topic_exclude_regex": schema.StringAttribute{
								Description:         "Regular expression to specify which topics to exclude.",
								MarkdownDescription: "Regular expression to specify which topics to exclude.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"topic_regex": schema.StringAttribute{
								Description:         "Regular expression to specify which topics to collect. Default value is '.*'.",
								MarkdownDescription: "Regular expression to specify which topics to collect. Default value is '.*'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"maintenance_time_windows": schema.ListAttribute{
						Description:         "A list of time windows for maintenance tasks (that is, certificates renewal). Each time window is defined by a cron expression.",
						MarkdownDescription: "A list of time windows for maintenance tasks (that is, certificates renewal). Each time window is defined by a cron expression.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"zookeeper": schema.SingleNestedAttribute{
						Description:         "Configuration of the ZooKeeper cluster.",
						MarkdownDescription: "Configuration of the ZooKeeper cluster.",
						Attributes: map[string]schema.Attribute{
							"config": schema.MapAttribute{
								Description:         "The ZooKeeper broker config. Properties with the following prefixes cannot be set: server., dataDir, dataLogDir, clientPort, authProvider, quorum.auth, requireClientAuthScheme, snapshot.trust.empty, standaloneEnabled, reconfigEnabled, 4lw.commands.whitelist, secureClientPort, ssl., serverCnxnFactory, sslQuorum (with the exception of: ssl.protocol, ssl.quorum.protocol, ssl.enabledProtocols, ssl.quorum.enabledProtocols, ssl.ciphersuites, ssl.quorum.ciphersuites, ssl.hostnameVerification, ssl.quorum.hostnameVerification).",
								MarkdownDescription: "The ZooKeeper broker config. Properties with the following prefixes cannot be set: server., dataDir, dataLogDir, clientPort, authProvider, quorum.auth, requireClientAuthScheme, snapshot.trust.empty, standaloneEnabled, reconfigEnabled, 4lw.commands.whitelist, secureClientPort, ssl., serverCnxnFactory, sslQuorum (with the exception of: ssl.protocol, ssl.quorum.protocol, ssl.enabledProtocols, ssl.quorum.enabledProtocols, ssl.ciphersuites, ssl.quorum.ciphersuites, ssl.hostnameVerification, ssl.quorum.hostnameVerification).",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "The docker image for the pods.",
								MarkdownDescription: "The docker image for the pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"jmx_options": schema.SingleNestedAttribute{
								Description:         "JMX Options for Zookeeper nodes.",
								MarkdownDescription: "JMX Options for Zookeeper nodes.",
								Attributes: map[string]schema.Attribute{
									"authentication": schema.SingleNestedAttribute{
										Description:         "Authentication configuration for connecting to the JMX port.",
										MarkdownDescription: "Authentication configuration for connecting to the JMX port.",
										Attributes: map[string]schema.Attribute{
											"type": schema.StringAttribute{
												Description:         "Authentication type. Currently the only supported types are 'password'.'password' type creates a username and protected port with no TLS.",
												MarkdownDescription: "Authentication type. Currently the only supported types are 'password'.'password' type creates a username and protected port with no TLS.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("password"),
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

							"liveness_probe": schema.SingleNestedAttribute{
								Description:         "Pod liveness checking.",
								MarkdownDescription: "Pod liveness checking.",
								Attributes: map[string]schema.Attribute{
									"failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": schema.Int64Attribute{
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"period_seconds": schema.Int64Attribute{
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": schema.SingleNestedAttribute{
								Description:         "Logging configuration for ZooKeeper.",
								MarkdownDescription: "Logging configuration for ZooKeeper.",
								Attributes: map[string]schema.Attribute{
									"loggers": schema.MapAttribute{
										Description:         "A Map from logger name to logger level.",
										MarkdownDescription: "A Map from logger name to logger level.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Logging type, must be either 'inline' or 'external'.",
										MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("inline", "external"),
										},
									},

									"value_from": schema.SingleNestedAttribute{
										Description:         "'ConfigMap' entry where the logging configuration is stored. ",
										MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics_config": schema.SingleNestedAttribute{
								Description:         "Metrics configuration.",
								MarkdownDescription: "Metrics configuration.",
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Description:         "Metrics type. Only 'jmxPrometheusExporter' supported currently.",
										MarkdownDescription: "Metrics type. Only 'jmxPrometheusExporter' supported currently.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("jmxPrometheusExporter"),
										},
									},

									"value_from": schema.SingleNestedAttribute{
										Description:         "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",
										MarkdownDescription: "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",
										Attributes: map[string]schema.Attribute{
											"config_map_key_ref": schema.SingleNestedAttribute{
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
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
										Required: true,
										Optional: false,
										Computed: false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"readiness_probe": schema.SingleNestedAttribute{
								Description:         "Pod readiness checking.",
								MarkdownDescription: "Pod readiness checking.",
								Attributes: map[string]schema.Attribute{
									"failure_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": schema.Int64Attribute{
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"period_seconds": schema.Int64Attribute{
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"success_threshold": schema.Int64Attribute{
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": schema.Int64Attribute{
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "The number of pods in the cluster.",
								MarkdownDescription: "The number of pods in the cluster.",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
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

									"overrides": schema.ListNestedAttribute{
										Description:         "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",
										MarkdownDescription: "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",
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
										Description:         "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",
										MarkdownDescription: "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",
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
								Required: true,
								Optional: false,
								Computed: false,
							},

							"template": schema.SingleNestedAttribute{
								Description:         "Template for ZooKeeper cluster resources. The template allows users to specify how the Kubernetes resources are generated.",
								MarkdownDescription: "Template for ZooKeeper cluster resources. The template allows users to specify how the Kubernetes resources are generated.",
								Attributes: map[string]schema.Attribute{
									"client_service": schema.SingleNestedAttribute{
										Description:         "Template for ZooKeeper client 'Service'.",
										MarkdownDescription: "Template for ZooKeeper client 'Service'.",
										Attributes: map[string]schema.Attribute{
											"ip_families": schema.ListAttribute{
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"jmx_secret": schema.SingleNestedAttribute{
										Description:         "Template for Secret of the Zookeeper Cluster JMX authentication.",
										MarkdownDescription: "Template for Secret of the Zookeeper Cluster JMX authentication.",
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

									"nodes_service": schema.SingleNestedAttribute{
										Description:         "Template for ZooKeeper nodes 'Service'.",
										MarkdownDescription: "Template for ZooKeeper nodes 'Service'.",
										Attributes: map[string]schema.Attribute{
											"ip_families": schema.ListAttribute{
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"persistent_volume_claim": schema.SingleNestedAttribute{
										Description:         "Template for all ZooKeeper 'PersistentVolumeClaims'.",
										MarkdownDescription: "Template for all ZooKeeper 'PersistentVolumeClaims'.",
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
										Description:         "Template for ZooKeeper 'Pods'.",
										MarkdownDescription: "Template for ZooKeeper 'Pods'.",
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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
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
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_disruption_budget": schema.SingleNestedAttribute{
										Description:         "Template for ZooKeeper 'PodDisruptionBudget'.",
										MarkdownDescription: "Template for ZooKeeper 'PodDisruptionBudget'.",
										Attributes: map[string]schema.Attribute{
											"max_unavailable": schema.Int64Attribute{
												Description:         "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
												MarkdownDescription: "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"metadata": schema.SingleNestedAttribute{
												Description:         "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
												MarkdownDescription: "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
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

									"pod_set": schema.SingleNestedAttribute{
										Description:         "Template for ZooKeeper 'StrimziPodSet' resource.",
										MarkdownDescription: "Template for ZooKeeper 'StrimziPodSet' resource.",
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

									"service_account": schema.SingleNestedAttribute{
										Description:         "Template for the ZooKeeper service account.",
										MarkdownDescription: "Template for the ZooKeeper service account.",
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

									"statefulset": schema.SingleNestedAttribute{
										Description:         "Template for ZooKeeper 'StatefulSet'.",
										MarkdownDescription: "Template for ZooKeeper 'StatefulSet'.",
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

											"pod_management_policy": schema.StringAttribute{
												Description:         "PodManagementPolicy which will be used for this StatefulSet. Valid values are 'Parallel' and 'OrderedReady'. Defaults to 'Parallel'.",
												MarkdownDescription: "PodManagementPolicy which will be used for this StatefulSet. Valid values are 'Parallel' and 'OrderedReady'. Defaults to 'Parallel'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("OrderedReady", "Parallel"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"zookeeper_container": schema.SingleNestedAttribute{
										Description:         "Template for the ZooKeeper container.",
										MarkdownDescription: "Template for the ZooKeeper container.",
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
	}
}

func (r *KafkaStrimziIoKafkaV1Beta2Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_strimzi_io_kafka_v1beta2_manifest")

	var model KafkaStrimziIoKafkaV1Beta2ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("kafka.strimzi.io/v1beta2")
	model.Kind = pointer.String("Kafka")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
