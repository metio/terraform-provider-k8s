/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type KafkaStrimziIoKafkaV1Beta2Resource struct{}

var (
	_ resource.Resource = (*KafkaStrimziIoKafkaV1Beta2Resource)(nil)
)

type KafkaStrimziIoKafkaV1Beta2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KafkaStrimziIoKafkaV1Beta2GoModel struct {
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
		ClientsCa *struct {
			CertificateExpirationPolicy *string `tfsdk:"certificate_expiration_policy" yaml:"certificateExpirationPolicy,omitempty"`

			GenerateCertificateAuthority *bool `tfsdk:"generate_certificate_authority" yaml:"generateCertificateAuthority,omitempty"`

			GenerateSecretOwnerReference *bool `tfsdk:"generate_secret_owner_reference" yaml:"generateSecretOwnerReference,omitempty"`

			RenewalDays *int64 `tfsdk:"renewal_days" yaml:"renewalDays,omitempty"`

			ValidityDays *int64 `tfsdk:"validity_days" yaml:"validityDays,omitempty"`
		} `tfsdk:"clients_ca" yaml:"clientsCa,omitempty"`

		ClusterCa *struct {
			CertificateExpirationPolicy *string `tfsdk:"certificate_expiration_policy" yaml:"certificateExpirationPolicy,omitempty"`

			GenerateCertificateAuthority *bool `tfsdk:"generate_certificate_authority" yaml:"generateCertificateAuthority,omitempty"`

			GenerateSecretOwnerReference *bool `tfsdk:"generate_secret_owner_reference" yaml:"generateSecretOwnerReference,omitempty"`

			RenewalDays *int64 `tfsdk:"renewal_days" yaml:"renewalDays,omitempty"`

			ValidityDays *int64 `tfsdk:"validity_days" yaml:"validityDays,omitempty"`
		} `tfsdk:"cluster_ca" yaml:"clusterCa,omitempty"`

		CruiseControl *struct {
			BrokerCapacity *struct {
				Cpu *string `tfsdk:"cpu" yaml:"cpu,omitempty"`

				CpuUtilization *int64 `tfsdk:"cpu_utilization" yaml:"cpuUtilization,omitempty"`

				Disk *string `tfsdk:"disk" yaml:"disk,omitempty"`

				InboundNetwork *string `tfsdk:"inbound_network" yaml:"inboundNetwork,omitempty"`

				OutboundNetwork *string `tfsdk:"outbound_network" yaml:"outboundNetwork,omitempty"`

				Overrides *[]struct {
					Brokers *[]string `tfsdk:"brokers" yaml:"brokers,omitempty"`

					Cpu *string `tfsdk:"cpu" yaml:"cpu,omitempty"`

					InboundNetwork *string `tfsdk:"inbound_network" yaml:"inboundNetwork,omitempty"`

					OutboundNetwork *string `tfsdk:"outbound_network" yaml:"outboundNetwork,omitempty"`
				} `tfsdk:"overrides" yaml:"overrides,omitempty"`
			} `tfsdk:"broker_capacity" yaml:"brokerCapacity,omitempty"`

			Config utilities.Dynamic `tfsdk:"config" yaml:"config,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			JvmOptions *struct {
				_XX utilities.Dynamic `tfsdk:"xx" yaml:"-XX,omitempty"`

				_Xms *string `tfsdk:"xms" yaml:"-Xms,omitempty"`

				_Xmx *string `tfsdk:"xmx" yaml:"-Xmx,omitempty"`

				GcLoggingEnabled *bool `tfsdk:"gc_logging_enabled" yaml:"gcLoggingEnabled,omitempty"`

				JavaSystemProperties *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"java_system_properties" yaml:"javaSystemProperties,omitempty"`
			} `tfsdk:"jvm_options" yaml:"jvmOptions,omitempty"`

			LivenessProbe *struct {
				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

			Logging *struct {
				Loggers utilities.Dynamic `tfsdk:"loggers" yaml:"loggers,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"logging" yaml:"logging,omitempty"`

			MetricsConfig *struct {
				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"metrics_config" yaml:"metricsConfig,omitempty"`

			ReadinessProbe *struct {
				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

			Resources *struct {
				Limits utilities.Dynamic `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests utilities.Dynamic `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Template *struct {
				ApiService *struct {
					IpFamilies *[]string `tfsdk:"ip_families" yaml:"ipFamilies,omitempty"`

					IpFamilyPolicy *string `tfsdk:"ip_family_policy" yaml:"ipFamilyPolicy,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"api_service" yaml:"apiService,omitempty"`

				CruiseControlContainer *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"cruise_control_container" yaml:"cruiseControlContainer,omitempty"`

				Deployment *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"deployment" yaml:"deployment,omitempty"`

				Pod *struct {
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" yaml:"affinity,omitempty"`

					EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

					HostAliases *[]struct {
						Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

						Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
					} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

					ImagePullSecrets *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

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

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TmpDirSizeLimit *string `tfsdk:"tmp_dir_size_limit" yaml:"tmpDirSizeLimit,omitempty"`

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

							MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						MatchLabelKeys *[]string `tfsdk:"match_label_keys" yaml:"matchLabelKeys,omitempty"`

						MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

						MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

						NodeAffinityPolicy *string `tfsdk:"node_affinity_policy" yaml:"nodeAffinityPolicy,omitempty"`

						NodeTaintsPolicy *string `tfsdk:"node_taints_policy" yaml:"nodeTaintsPolicy,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

						WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
					} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
				} `tfsdk:"pod" yaml:"pod,omitempty"`

				PodDisruptionBudget *struct {
					MaxUnavailable *int64 `tfsdk:"max_unavailable" yaml:"maxUnavailable,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"pod_disruption_budget" yaml:"podDisruptionBudget,omitempty"`

				ServiceAccount *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

				TlsSidecarContainer *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"tls_sidecar_container" yaml:"tlsSidecarContainer,omitempty"`
			} `tfsdk:"template" yaml:"template,omitempty"`

			TlsSidecar *struct {
				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				LivenessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				ReadinessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

				Resources *struct {
					Limits utilities.Dynamic `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests utilities.Dynamic `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`
			} `tfsdk:"tls_sidecar" yaml:"tlsSidecar,omitempty"`
		} `tfsdk:"cruise_control" yaml:"cruiseControl,omitempty"`

		EntityOperator *struct {
			Template *struct {
				Deployment *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"deployment" yaml:"deployment,omitempty"`

				Pod *struct {
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" yaml:"affinity,omitempty"`

					EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

					HostAliases *[]struct {
						Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

						Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
					} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

					ImagePullSecrets *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

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

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TmpDirSizeLimit *string `tfsdk:"tmp_dir_size_limit" yaml:"tmpDirSizeLimit,omitempty"`

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

							MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						MatchLabelKeys *[]string `tfsdk:"match_label_keys" yaml:"matchLabelKeys,omitempty"`

						MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

						MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

						NodeAffinityPolicy *string `tfsdk:"node_affinity_policy" yaml:"nodeAffinityPolicy,omitempty"`

						NodeTaintsPolicy *string `tfsdk:"node_taints_policy" yaml:"nodeTaintsPolicy,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

						WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
					} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
				} `tfsdk:"pod" yaml:"pod,omitempty"`

				ServiceAccount *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

				TlsSidecarContainer *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"tls_sidecar_container" yaml:"tlsSidecarContainer,omitempty"`

				TopicOperatorContainer *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"topic_operator_container" yaml:"topicOperatorContainer,omitempty"`

				UserOperatorContainer *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"user_operator_container" yaml:"userOperatorContainer,omitempty"`
			} `tfsdk:"template" yaml:"template,omitempty"`

			TlsSidecar *struct {
				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				LivenessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

				ReadinessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

				Resources *struct {
					Limits utilities.Dynamic `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests utilities.Dynamic `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`
			} `tfsdk:"tls_sidecar" yaml:"tlsSidecar,omitempty"`

			TopicOperator *struct {
				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				JvmOptions *struct {
					_XX utilities.Dynamic `tfsdk:"xx" yaml:"-XX,omitempty"`

					_Xms *string `tfsdk:"xms" yaml:"-Xms,omitempty"`

					_Xmx *string `tfsdk:"xmx" yaml:"-Xmx,omitempty"`

					GcLoggingEnabled *bool `tfsdk:"gc_logging_enabled" yaml:"gcLoggingEnabled,omitempty"`

					JavaSystemProperties *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"java_system_properties" yaml:"javaSystemProperties,omitempty"`
				} `tfsdk:"jvm_options" yaml:"jvmOptions,omitempty"`

				LivenessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				Logging *struct {
					Loggers utilities.Dynamic `tfsdk:"loggers" yaml:"loggers,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"logging" yaml:"logging,omitempty"`

				ReadinessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

				ReconciliationIntervalSeconds *int64 `tfsdk:"reconciliation_interval_seconds" yaml:"reconciliationIntervalSeconds,omitempty"`

				Resources *struct {
					Limits utilities.Dynamic `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests utilities.Dynamic `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				StartupProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

				TopicMetadataMaxAttempts *int64 `tfsdk:"topic_metadata_max_attempts" yaml:"topicMetadataMaxAttempts,omitempty"`

				WatchedNamespace *string `tfsdk:"watched_namespace" yaml:"watchedNamespace,omitempty"`

				ZookeeperSessionTimeoutSeconds *int64 `tfsdk:"zookeeper_session_timeout_seconds" yaml:"zookeeperSessionTimeoutSeconds,omitempty"`
			} `tfsdk:"topic_operator" yaml:"topicOperator,omitempty"`

			UserOperator *struct {
				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				JvmOptions *struct {
					_XX utilities.Dynamic `tfsdk:"xx" yaml:"-XX,omitempty"`

					_Xms *string `tfsdk:"xms" yaml:"-Xms,omitempty"`

					_Xmx *string `tfsdk:"xmx" yaml:"-Xmx,omitempty"`

					GcLoggingEnabled *bool `tfsdk:"gc_logging_enabled" yaml:"gcLoggingEnabled,omitempty"`

					JavaSystemProperties *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"java_system_properties" yaml:"javaSystemProperties,omitempty"`
				} `tfsdk:"jvm_options" yaml:"jvmOptions,omitempty"`

				LivenessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				Logging *struct {
					Loggers utilities.Dynamic `tfsdk:"loggers" yaml:"loggers,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
					} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
				} `tfsdk:"logging" yaml:"logging,omitempty"`

				ReadinessProbe *struct {
					FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

				ReconciliationIntervalSeconds *int64 `tfsdk:"reconciliation_interval_seconds" yaml:"reconciliationIntervalSeconds,omitempty"`

				Resources *struct {
					Limits utilities.Dynamic `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests utilities.Dynamic `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				SecretPrefix *string `tfsdk:"secret_prefix" yaml:"secretPrefix,omitempty"`

				WatchedNamespace *string `tfsdk:"watched_namespace" yaml:"watchedNamespace,omitempty"`

				ZookeeperSessionTimeoutSeconds *int64 `tfsdk:"zookeeper_session_timeout_seconds" yaml:"zookeeperSessionTimeoutSeconds,omitempty"`
			} `tfsdk:"user_operator" yaml:"userOperator,omitempty"`
		} `tfsdk:"entity_operator" yaml:"entityOperator,omitempty"`

		JmxTrans *struct {
			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			KafkaQueries *[]struct {
				Attributes *[]string `tfsdk:"attributes" yaml:"attributes,omitempty"`

				Outputs *[]string `tfsdk:"outputs" yaml:"outputs,omitempty"`

				TargetMBean *string `tfsdk:"target_m_bean" yaml:"targetMBean,omitempty"`
			} `tfsdk:"kafka_queries" yaml:"kafkaQueries,omitempty"`

			LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

			OutputDefinitions *[]struct {
				FlushDelayInSeconds *int64 `tfsdk:"flush_delay_in_seconds" yaml:"flushDelayInSeconds,omitempty"`

				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				OutputType *string `tfsdk:"output_type" yaml:"outputType,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				TypeNames *[]string `tfsdk:"type_names" yaml:"typeNames,omitempty"`
			} `tfsdk:"output_definitions" yaml:"outputDefinitions,omitempty"`

			Resources *struct {
				Limits utilities.Dynamic `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests utilities.Dynamic `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Template *struct {
				Container *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"container" yaml:"container,omitempty"`

				Deployment *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"deployment" yaml:"deployment,omitempty"`

				Pod *struct {
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" yaml:"affinity,omitempty"`

					EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

					HostAliases *[]struct {
						Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

						Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
					} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

					ImagePullSecrets *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

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

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TmpDirSizeLimit *string `tfsdk:"tmp_dir_size_limit" yaml:"tmpDirSizeLimit,omitempty"`

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

							MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						MatchLabelKeys *[]string `tfsdk:"match_label_keys" yaml:"matchLabelKeys,omitempty"`

						MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

						MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

						NodeAffinityPolicy *string `tfsdk:"node_affinity_policy" yaml:"nodeAffinityPolicy,omitempty"`

						NodeTaintsPolicy *string `tfsdk:"node_taints_policy" yaml:"nodeTaintsPolicy,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

						WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
					} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
				} `tfsdk:"pod" yaml:"pod,omitempty"`

				ServiceAccount *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`
			} `tfsdk:"template" yaml:"template,omitempty"`
		} `tfsdk:"jmx_trans" yaml:"jmxTrans,omitempty"`

		Kafka *struct {
			Authorization *struct {
				AllowOnError *bool `tfsdk:"allow_on_error" yaml:"allowOnError,omitempty"`

				AuthorizerClass *string `tfsdk:"authorizer_class" yaml:"authorizerClass,omitempty"`

				ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

				ConnectTimeoutSeconds *int64 `tfsdk:"connect_timeout_seconds" yaml:"connectTimeoutSeconds,omitempty"`

				DelegateToKafkaAcls *bool `tfsdk:"delegate_to_kafka_acls" yaml:"delegateToKafkaAcls,omitempty"`

				DisableTlsHostnameVerification *bool `tfsdk:"disable_tls_hostname_verification" yaml:"disableTlsHostnameVerification,omitempty"`

				EnableMetrics *bool `tfsdk:"enable_metrics" yaml:"enableMetrics,omitempty"`

				ExpireAfterMs *int64 `tfsdk:"expire_after_ms" yaml:"expireAfterMs,omitempty"`

				GrantsRefreshPeriodSeconds *int64 `tfsdk:"grants_refresh_period_seconds" yaml:"grantsRefreshPeriodSeconds,omitempty"`

				GrantsRefreshPoolSize *int64 `tfsdk:"grants_refresh_pool_size" yaml:"grantsRefreshPoolSize,omitempty"`

				InitialCacheCapacity *int64 `tfsdk:"initial_cache_capacity" yaml:"initialCacheCapacity,omitempty"`

				MaximumCacheSize *int64 `tfsdk:"maximum_cache_size" yaml:"maximumCacheSize,omitempty"`

				ReadTimeoutSeconds *int64 `tfsdk:"read_timeout_seconds" yaml:"readTimeoutSeconds,omitempty"`

				SuperUsers *[]string `tfsdk:"super_users" yaml:"superUsers,omitempty"`

				SupportsAdminApi *bool `tfsdk:"supports_admin_api" yaml:"supportsAdminApi,omitempty"`

				TlsTrustedCertificates *[]struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"tls_trusted_certificates" yaml:"tlsTrustedCertificates,omitempty"`

				TokenEndpointUri *string `tfsdk:"token_endpoint_uri" yaml:"tokenEndpointUri,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"authorization" yaml:"authorization,omitempty"`

			BrokerRackInitImage *string `tfsdk:"broker_rack_init_image" yaml:"brokerRackInitImage,omitempty"`

			Config utilities.Dynamic `tfsdk:"config" yaml:"config,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			JmxOptions *struct {
				Authentication *struct {
					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"authentication" yaml:"authentication,omitempty"`
			} `tfsdk:"jmx_options" yaml:"jmxOptions,omitempty"`

			JvmOptions *struct {
				_XX utilities.Dynamic `tfsdk:"xx" yaml:"-XX,omitempty"`

				_Xms *string `tfsdk:"xms" yaml:"-Xms,omitempty"`

				_Xmx *string `tfsdk:"xmx" yaml:"-Xmx,omitempty"`

				GcLoggingEnabled *bool `tfsdk:"gc_logging_enabled" yaml:"gcLoggingEnabled,omitempty"`

				JavaSystemProperties *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"java_system_properties" yaml:"javaSystemProperties,omitempty"`
			} `tfsdk:"jvm_options" yaml:"jvmOptions,omitempty"`

			Listeners *[]struct {
				Authentication *struct {
					AccessTokenIsJwt *bool `tfsdk:"access_token_is_jwt" yaml:"accessTokenIsJwt,omitempty"`

					CheckAccessTokenType *bool `tfsdk:"check_access_token_type" yaml:"checkAccessTokenType,omitempty"`

					CheckAudience *bool `tfsdk:"check_audience" yaml:"checkAudience,omitempty"`

					CheckIssuer *bool `tfsdk:"check_issuer" yaml:"checkIssuer,omitempty"`

					ClientAudience *string `tfsdk:"client_audience" yaml:"clientAudience,omitempty"`

					ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`

					ClientScope *string `tfsdk:"client_scope" yaml:"clientScope,omitempty"`

					ClientSecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

					ConnectTimeoutSeconds *int64 `tfsdk:"connect_timeout_seconds" yaml:"connectTimeoutSeconds,omitempty"`

					CustomClaimCheck *string `tfsdk:"custom_claim_check" yaml:"customClaimCheck,omitempty"`

					DisableTlsHostnameVerification *bool `tfsdk:"disable_tls_hostname_verification" yaml:"disableTlsHostnameVerification,omitempty"`

					EnableECDSA *bool `tfsdk:"enable_ecdsa" yaml:"enableECDSA,omitempty"`

					EnableMetrics *bool `tfsdk:"enable_metrics" yaml:"enableMetrics,omitempty"`

					EnableOauthBearer *bool `tfsdk:"enable_oauth_bearer" yaml:"enableOauthBearer,omitempty"`

					EnablePlain *bool `tfsdk:"enable_plain" yaml:"enablePlain,omitempty"`

					FailFast *bool `tfsdk:"fail_fast" yaml:"failFast,omitempty"`

					FallbackUserNameClaim *string `tfsdk:"fallback_user_name_claim" yaml:"fallbackUserNameClaim,omitempty"`

					FallbackUserNamePrefix *string `tfsdk:"fallback_user_name_prefix" yaml:"fallbackUserNamePrefix,omitempty"`

					GroupsClaim *string `tfsdk:"groups_claim" yaml:"groupsClaim,omitempty"`

					GroupsClaimDelimiter *string `tfsdk:"groups_claim_delimiter" yaml:"groupsClaimDelimiter,omitempty"`

					IntrospectionEndpointUri *string `tfsdk:"introspection_endpoint_uri" yaml:"introspectionEndpointUri,omitempty"`

					JwksEndpointUri *string `tfsdk:"jwks_endpoint_uri" yaml:"jwksEndpointUri,omitempty"`

					JwksExpirySeconds *int64 `tfsdk:"jwks_expiry_seconds" yaml:"jwksExpirySeconds,omitempty"`

					JwksIgnoreKeyUse *bool `tfsdk:"jwks_ignore_key_use" yaml:"jwksIgnoreKeyUse,omitempty"`

					JwksMinRefreshPauseSeconds *int64 `tfsdk:"jwks_min_refresh_pause_seconds" yaml:"jwksMinRefreshPauseSeconds,omitempty"`

					JwksRefreshSeconds *int64 `tfsdk:"jwks_refresh_seconds" yaml:"jwksRefreshSeconds,omitempty"`

					ListenerConfig utilities.Dynamic `tfsdk:"listener_config" yaml:"listenerConfig,omitempty"`

					MaxSecondsWithoutReauthentication *int64 `tfsdk:"max_seconds_without_reauthentication" yaml:"maxSecondsWithoutReauthentication,omitempty"`

					ReadTimeoutSeconds *int64 `tfsdk:"read_timeout_seconds" yaml:"readTimeoutSeconds,omitempty"`

					Sasl *bool `tfsdk:"sasl" yaml:"sasl,omitempty"`

					Secrets *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"secrets" yaml:"secrets,omitempty"`

					TlsTrustedCertificates *[]struct {
						Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"tls_trusted_certificates" yaml:"tlsTrustedCertificates,omitempty"`

					TokenEndpointUri *string `tfsdk:"token_endpoint_uri" yaml:"tokenEndpointUri,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					UserInfoEndpointUri *string `tfsdk:"user_info_endpoint_uri" yaml:"userInfoEndpointUri,omitempty"`

					UserNameClaim *string `tfsdk:"user_name_claim" yaml:"userNameClaim,omitempty"`

					ValidIssuerUri *string `tfsdk:"valid_issuer_uri" yaml:"validIssuerUri,omitempty"`

					ValidTokenType *string `tfsdk:"valid_token_type" yaml:"validTokenType,omitempty"`
				} `tfsdk:"authentication" yaml:"authentication,omitempty"`

				Configuration *struct {
					Bootstrap *struct {
						AlternativeNames *[]string `tfsdk:"alternative_names" yaml:"alternativeNames,omitempty"`

						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`

						LoadBalancerIP *string `tfsdk:"load_balancer_ip" yaml:"loadBalancerIP,omitempty"`

						NodePort *int64 `tfsdk:"node_port" yaml:"nodePort,omitempty"`
					} `tfsdk:"bootstrap" yaml:"bootstrap,omitempty"`

					BrokerCertChainAndKey *struct {
						Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"broker_cert_chain_and_key" yaml:"brokerCertChainAndKey,omitempty"`

					Brokers *[]struct {
						AdvertisedHost *string `tfsdk:"advertised_host" yaml:"advertisedHost,omitempty"`

						AdvertisedPort *int64 `tfsdk:"advertised_port" yaml:"advertisedPort,omitempty"`

						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Broker *int64 `tfsdk:"broker" yaml:"broker,omitempty"`

						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`

						LoadBalancerIP *string `tfsdk:"load_balancer_ip" yaml:"loadBalancerIP,omitempty"`

						NodePort *int64 `tfsdk:"node_port" yaml:"nodePort,omitempty"`
					} `tfsdk:"brokers" yaml:"brokers,omitempty"`

					Class *string `tfsdk:"class" yaml:"class,omitempty"`

					CreateBootstrapService *bool `tfsdk:"create_bootstrap_service" yaml:"createBootstrapService,omitempty"`

					ExternalTrafficPolicy *string `tfsdk:"external_traffic_policy" yaml:"externalTrafficPolicy,omitempty"`

					Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

					IpFamilies *[]string `tfsdk:"ip_families" yaml:"ipFamilies,omitempty"`

					IpFamilyPolicy *string `tfsdk:"ip_family_policy" yaml:"ipFamilyPolicy,omitempty"`

					LoadBalancerSourceRanges *[]string `tfsdk:"load_balancer_source_ranges" yaml:"loadBalancerSourceRanges,omitempty"`

					MaxConnectionCreationRate *int64 `tfsdk:"max_connection_creation_rate" yaml:"maxConnectionCreationRate,omitempty"`

					MaxConnections *int64 `tfsdk:"max_connections" yaml:"maxConnections,omitempty"`

					PreferredNodePortAddressType *string `tfsdk:"preferred_node_port_address_type" yaml:"preferredNodePortAddressType,omitempty"`

					UseServiceDnsDomain *bool `tfsdk:"use_service_dns_domain" yaml:"useServiceDnsDomain,omitempty"`
				} `tfsdk:"configuration" yaml:"configuration,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				NetworkPolicyPeers *[]struct {
					IpBlock *struct {
						Cidr *string `tfsdk:"cidr" yaml:"cidr,omitempty"`

						Except *[]string `tfsdk:"except" yaml:"except,omitempty"`
					} `tfsdk:"ip_block" yaml:"ipBlock,omitempty"`

					NamespaceSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

					PodSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"pod_selector" yaml:"podSelector,omitempty"`
				} `tfsdk:"network_policy_peers" yaml:"networkPolicyPeers,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				Tls *bool `tfsdk:"tls" yaml:"tls,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"listeners" yaml:"listeners,omitempty"`

			LivenessProbe *struct {
				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

			Logging *struct {
				Loggers utilities.Dynamic `tfsdk:"loggers" yaml:"loggers,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"logging" yaml:"logging,omitempty"`

			MetricsConfig *struct {
				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"metrics_config" yaml:"metricsConfig,omitempty"`

			Rack *struct {
				TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
			} `tfsdk:"rack" yaml:"rack,omitempty"`

			ReadinessProbe *struct {
				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			Resources *struct {
				Limits utilities.Dynamic `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests utilities.Dynamic `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Storage *struct {
				Class *string `tfsdk:"class" yaml:"class,omitempty"`

				DeleteClaim *bool `tfsdk:"delete_claim" yaml:"deleteClaim,omitempty"`

				Id *int64 `tfsdk:"id" yaml:"id,omitempty"`

				Overrides *[]struct {
					Broker *int64 `tfsdk:"broker" yaml:"broker,omitempty"`

					Class *string `tfsdk:"class" yaml:"class,omitempty"`
				} `tfsdk:"overrides" yaml:"overrides,omitempty"`

				Selector utilities.Dynamic `tfsdk:"selector" yaml:"selector,omitempty"`

				Size *string `tfsdk:"size" yaml:"size,omitempty"`

				SizeLimit *string `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				Volumes *[]struct {
					Class *string `tfsdk:"class" yaml:"class,omitempty"`

					DeleteClaim *bool `tfsdk:"delete_claim" yaml:"deleteClaim,omitempty"`

					Id *int64 `tfsdk:"id" yaml:"id,omitempty"`

					Overrides *[]struct {
						Broker *int64 `tfsdk:"broker" yaml:"broker,omitempty"`

						Class *string `tfsdk:"class" yaml:"class,omitempty"`
					} `tfsdk:"overrides" yaml:"overrides,omitempty"`

					Selector utilities.Dynamic `tfsdk:"selector" yaml:"selector,omitempty"`

					Size *string `tfsdk:"size" yaml:"size,omitempty"`

					SizeLimit *string `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"volumes" yaml:"volumes,omitempty"`
			} `tfsdk:"storage" yaml:"storage,omitempty"`

			Template *struct {
				BootstrapService *struct {
					IpFamilies *[]string `tfsdk:"ip_families" yaml:"ipFamilies,omitempty"`

					IpFamilyPolicy *string `tfsdk:"ip_family_policy" yaml:"ipFamilyPolicy,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"bootstrap_service" yaml:"bootstrapService,omitempty"`

				BrokersService *struct {
					IpFamilies *[]string `tfsdk:"ip_families" yaml:"ipFamilies,omitempty"`

					IpFamilyPolicy *string `tfsdk:"ip_family_policy" yaml:"ipFamilyPolicy,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"brokers_service" yaml:"brokersService,omitempty"`

				ClusterCaCert *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"cluster_ca_cert" yaml:"clusterCaCert,omitempty"`

				ClusterRoleBinding *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"cluster_role_binding" yaml:"clusterRoleBinding,omitempty"`

				ExternalBootstrapIngress *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"external_bootstrap_ingress" yaml:"externalBootstrapIngress,omitempty"`

				ExternalBootstrapRoute *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"external_bootstrap_route" yaml:"externalBootstrapRoute,omitempty"`

				ExternalBootstrapService *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"external_bootstrap_service" yaml:"externalBootstrapService,omitempty"`

				InitContainer *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"init_container" yaml:"initContainer,omitempty"`

				JmxSecret *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"jmx_secret" yaml:"jmxSecret,omitempty"`

				KafkaContainer *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"kafka_container" yaml:"kafkaContainer,omitempty"`

				PerPodIngress *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"per_pod_ingress" yaml:"perPodIngress,omitempty"`

				PerPodRoute *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"per_pod_route" yaml:"perPodRoute,omitempty"`

				PerPodService *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"per_pod_service" yaml:"perPodService,omitempty"`

				PersistentVolumeClaim *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`

				Pod *struct {
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" yaml:"affinity,omitempty"`

					EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

					HostAliases *[]struct {
						Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

						Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
					} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

					ImagePullSecrets *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

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

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TmpDirSizeLimit *string `tfsdk:"tmp_dir_size_limit" yaml:"tmpDirSizeLimit,omitempty"`

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

							MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						MatchLabelKeys *[]string `tfsdk:"match_label_keys" yaml:"matchLabelKeys,omitempty"`

						MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

						MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

						NodeAffinityPolicy *string `tfsdk:"node_affinity_policy" yaml:"nodeAffinityPolicy,omitempty"`

						NodeTaintsPolicy *string `tfsdk:"node_taints_policy" yaml:"nodeTaintsPolicy,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

						WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
					} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
				} `tfsdk:"pod" yaml:"pod,omitempty"`

				PodDisruptionBudget *struct {
					MaxUnavailable *int64 `tfsdk:"max_unavailable" yaml:"maxUnavailable,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"pod_disruption_budget" yaml:"podDisruptionBudget,omitempty"`

				PodSet *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"pod_set" yaml:"podSet,omitempty"`

				ServiceAccount *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

				Statefulset *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					PodManagementPolicy *string `tfsdk:"pod_management_policy" yaml:"podManagementPolicy,omitempty"`
				} `tfsdk:"statefulset" yaml:"statefulset,omitempty"`
			} `tfsdk:"template" yaml:"template,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"kafka" yaml:"kafka,omitempty"`

		KafkaExporter *struct {
			EnableSaramaLogging *bool `tfsdk:"enable_sarama_logging" yaml:"enableSaramaLogging,omitempty"`

			GroupRegex *string `tfsdk:"group_regex" yaml:"groupRegex,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			LivenessProbe *struct {
				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

			Logging *string `tfsdk:"logging" yaml:"logging,omitempty"`

			ReadinessProbe *struct {
				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

			Resources *struct {
				Limits utilities.Dynamic `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests utilities.Dynamic `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Template *struct {
				Container *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"container" yaml:"container,omitempty"`

				Deployment *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"deployment" yaml:"deployment,omitempty"`

				Pod *struct {
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" yaml:"affinity,omitempty"`

					EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

					HostAliases *[]struct {
						Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

						Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
					} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

					ImagePullSecrets *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

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

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TmpDirSizeLimit *string `tfsdk:"tmp_dir_size_limit" yaml:"tmpDirSizeLimit,omitempty"`

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

							MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						MatchLabelKeys *[]string `tfsdk:"match_label_keys" yaml:"matchLabelKeys,omitempty"`

						MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

						MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

						NodeAffinityPolicy *string `tfsdk:"node_affinity_policy" yaml:"nodeAffinityPolicy,omitempty"`

						NodeTaintsPolicy *string `tfsdk:"node_taints_policy" yaml:"nodeTaintsPolicy,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

						WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
					} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
				} `tfsdk:"pod" yaml:"pod,omitempty"`

				Service *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"service" yaml:"service,omitempty"`

				ServiceAccount *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`
			} `tfsdk:"template" yaml:"template,omitempty"`

			TopicRegex *string `tfsdk:"topic_regex" yaml:"topicRegex,omitempty"`
		} `tfsdk:"kafka_exporter" yaml:"kafkaExporter,omitempty"`

		MaintenanceTimeWindows *[]string `tfsdk:"maintenance_time_windows" yaml:"maintenanceTimeWindows,omitempty"`

		Zookeeper *struct {
			Config utilities.Dynamic `tfsdk:"config" yaml:"config,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			JmxOptions *struct {
				Authentication *struct {
					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"authentication" yaml:"authentication,omitempty"`
			} `tfsdk:"jmx_options" yaml:"jmxOptions,omitempty"`

			JvmOptions *struct {
				_XX utilities.Dynamic `tfsdk:"xx" yaml:"-XX,omitempty"`

				_Xms *string `tfsdk:"xms" yaml:"-Xms,omitempty"`

				_Xmx *string `tfsdk:"xmx" yaml:"-Xmx,omitempty"`

				GcLoggingEnabled *bool `tfsdk:"gc_logging_enabled" yaml:"gcLoggingEnabled,omitempty"`

				JavaSystemProperties *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"java_system_properties" yaml:"javaSystemProperties,omitempty"`
			} `tfsdk:"jvm_options" yaml:"jvmOptions,omitempty"`

			LivenessProbe *struct {
				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

			Logging *struct {
				Loggers utilities.Dynamic `tfsdk:"loggers" yaml:"loggers,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"logging" yaml:"logging,omitempty"`

			MetricsConfig *struct {
				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
			} `tfsdk:"metrics_config" yaml:"metricsConfig,omitempty"`

			ReadinessProbe *struct {
				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			Resources *struct {
				Limits utilities.Dynamic `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests utilities.Dynamic `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Storage *struct {
				Class *string `tfsdk:"class" yaml:"class,omitempty"`

				DeleteClaim *bool `tfsdk:"delete_claim" yaml:"deleteClaim,omitempty"`

				Id *int64 `tfsdk:"id" yaml:"id,omitempty"`

				Overrides *[]struct {
					Broker *int64 `tfsdk:"broker" yaml:"broker,omitempty"`

					Class *string `tfsdk:"class" yaml:"class,omitempty"`
				} `tfsdk:"overrides" yaml:"overrides,omitempty"`

				Selector utilities.Dynamic `tfsdk:"selector" yaml:"selector,omitempty"`

				Size *string `tfsdk:"size" yaml:"size,omitempty"`

				SizeLimit *string `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"storage" yaml:"storage,omitempty"`

			Template *struct {
				ClientService *struct {
					IpFamilies *[]string `tfsdk:"ip_families" yaml:"ipFamilies,omitempty"`

					IpFamilyPolicy *string `tfsdk:"ip_family_policy" yaml:"ipFamilyPolicy,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"client_service" yaml:"clientService,omitempty"`

				JmxSecret *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"jmx_secret" yaml:"jmxSecret,omitempty"`

				NodesService *struct {
					IpFamilies *[]string `tfsdk:"ip_families" yaml:"ipFamilies,omitempty"`

					IpFamilyPolicy *string `tfsdk:"ip_family_policy" yaml:"ipFamilyPolicy,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"nodes_service" yaml:"nodesService,omitempty"`

				PersistentVolumeClaim *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"persistent_volume_claim" yaml:"persistentVolumeClaim,omitempty"`

				Pod *struct {
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
									} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

									NamespaceSelector *struct {
										MatchExpressions *[]struct {
											Key *string `tfsdk:"key" yaml:"key,omitempty"`

											Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

											Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
										} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

										MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
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

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" yaml:"affinity,omitempty"`

					EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

					HostAliases *[]struct {
						Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

						Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
					} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

					ImagePullSecrets *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

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

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TmpDirSizeLimit *string `tfsdk:"tmp_dir_size_limit" yaml:"tmpDirSizeLimit,omitempty"`

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

							MatchLabels utilities.Dynamic `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						MatchLabelKeys *[]string `tfsdk:"match_label_keys" yaml:"matchLabelKeys,omitempty"`

						MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

						MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

						NodeAffinityPolicy *string `tfsdk:"node_affinity_policy" yaml:"nodeAffinityPolicy,omitempty"`

						NodeTaintsPolicy *string `tfsdk:"node_taints_policy" yaml:"nodeTaintsPolicy,omitempty"`

						TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

						WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
					} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
				} `tfsdk:"pod" yaml:"pod,omitempty"`

				PodDisruptionBudget *struct {
					MaxUnavailable *int64 `tfsdk:"max_unavailable" yaml:"maxUnavailable,omitempty"`

					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"pod_disruption_budget" yaml:"podDisruptionBudget,omitempty"`

				PodSet *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"pod_set" yaml:"podSet,omitempty"`

				ServiceAccount *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`
				} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

				Statefulset *struct {
					Metadata *struct {
						Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

						Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					PodManagementPolicy *string `tfsdk:"pod_management_policy" yaml:"podManagementPolicy,omitempty"`
				} `tfsdk:"statefulset" yaml:"statefulset,omitempty"`

				ZookeeperContainer *struct {
					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

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
				} `tfsdk:"zookeeper_container" yaml:"zookeeperContainer,omitempty"`
			} `tfsdk:"template" yaml:"template,omitempty"`
		} `tfsdk:"zookeeper" yaml:"zookeeper,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKafkaStrimziIoKafkaV1Beta2Resource() resource.Resource {
	return &KafkaStrimziIoKafkaV1Beta2Resource{}
}

func (r *KafkaStrimziIoKafkaV1Beta2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kafka_strimzi_io_kafka_v1beta2"
}

func (r *KafkaStrimziIoKafkaV1Beta2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "The specification of the Kafka and ZooKeeper clusters, and Topic Operator.",
				MarkdownDescription: "The specification of the Kafka and ZooKeeper clusters, and Topic Operator.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"clients_ca": {
						Description:         "Configuration of the clients certificate authority.",
						MarkdownDescription: "Configuration of the clients certificate authority.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"certificate_expiration_policy": {
								Description:         "How should CA certificate expiration be handled when 'generateCertificateAuthority=true'. The default is for a new CA certificate to be generated reusing the existing private key.",
								MarkdownDescription: "How should CA certificate expiration be handled when 'generateCertificateAuthority=true'. The default is for a new CA certificate to be generated reusing the existing private key.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("renew-certificate", "replace-key"),
								},
							},

							"generate_certificate_authority": {
								Description:         "If true then Certificate Authority certificates will be generated automatically. Otherwise the user will need to provide a Secret with the CA certificate. Default is true.",
								MarkdownDescription: "If true then Certificate Authority certificates will be generated automatically. Otherwise the user will need to provide a Secret with the CA certificate. Default is true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"generate_secret_owner_reference": {
								Description:         "If 'true', the Cluster and Client CA Secrets are configured with the 'ownerReference' set to the 'Kafka' resource. If the 'Kafka' resource is deleted when 'true', the CA Secrets are also deleted. If 'false', the 'ownerReference' is disabled. If the 'Kafka' resource is deleted when 'false', the CA Secrets are retained and available for reuse. Default is 'true'.",
								MarkdownDescription: "If 'true', the Cluster and Client CA Secrets are configured with the 'ownerReference' set to the 'Kafka' resource. If the 'Kafka' resource is deleted when 'true', the CA Secrets are also deleted. If 'false', the 'ownerReference' is disabled. If the 'Kafka' resource is deleted when 'false', the CA Secrets are retained and available for reuse. Default is 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"renewal_days": {
								Description:         "The number of days in the certificate renewal period. This is the number of days before the a certificate expires during which renewal actions may be performed. When 'generateCertificateAuthority' is true, this will cause the generation of a new certificate. When 'generateCertificateAuthority' is true, this will cause extra logging at WARN level about the pending certificate expiry. Default is 30.",
								MarkdownDescription: "The number of days in the certificate renewal period. This is the number of days before the a certificate expires during which renewal actions may be performed. When 'generateCertificateAuthority' is true, this will cause the generation of a new certificate. When 'generateCertificateAuthority' is true, this will cause extra logging at WARN level about the pending certificate expiry. Default is 30.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"validity_days": {
								Description:         "The number of days generated certificates should be valid for. The default is 365.",
								MarkdownDescription: "The number of days generated certificates should be valid for. The default is 365.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_ca": {
						Description:         "Configuration of the cluster certificate authority.",
						MarkdownDescription: "Configuration of the cluster certificate authority.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"certificate_expiration_policy": {
								Description:         "How should CA certificate expiration be handled when 'generateCertificateAuthority=true'. The default is for a new CA certificate to be generated reusing the existing private key.",
								MarkdownDescription: "How should CA certificate expiration be handled when 'generateCertificateAuthority=true'. The default is for a new CA certificate to be generated reusing the existing private key.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("renew-certificate", "replace-key"),
								},
							},

							"generate_certificate_authority": {
								Description:         "If true then Certificate Authority certificates will be generated automatically. Otherwise the user will need to provide a Secret with the CA certificate. Default is true.",
								MarkdownDescription: "If true then Certificate Authority certificates will be generated automatically. Otherwise the user will need to provide a Secret with the CA certificate. Default is true.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"generate_secret_owner_reference": {
								Description:         "If 'true', the Cluster and Client CA Secrets are configured with the 'ownerReference' set to the 'Kafka' resource. If the 'Kafka' resource is deleted when 'true', the CA Secrets are also deleted. If 'false', the 'ownerReference' is disabled. If the 'Kafka' resource is deleted when 'false', the CA Secrets are retained and available for reuse. Default is 'true'.",
								MarkdownDescription: "If 'true', the Cluster and Client CA Secrets are configured with the 'ownerReference' set to the 'Kafka' resource. If the 'Kafka' resource is deleted when 'true', the CA Secrets are also deleted. If 'false', the 'ownerReference' is disabled. If the 'Kafka' resource is deleted when 'false', the CA Secrets are retained and available for reuse. Default is 'true'.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"renewal_days": {
								Description:         "The number of days in the certificate renewal period. This is the number of days before the a certificate expires during which renewal actions may be performed. When 'generateCertificateAuthority' is true, this will cause the generation of a new certificate. When 'generateCertificateAuthority' is true, this will cause extra logging at WARN level about the pending certificate expiry. Default is 30.",
								MarkdownDescription: "The number of days in the certificate renewal period. This is the number of days before the a certificate expires during which renewal actions may be performed. When 'generateCertificateAuthority' is true, this will cause the generation of a new certificate. When 'generateCertificateAuthority' is true, this will cause extra logging at WARN level about the pending certificate expiry. Default is 30.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"validity_days": {
								Description:         "The number of days generated certificates should be valid for. The default is 365.",
								MarkdownDescription: "The number of days generated certificates should be valid for. The default is 365.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cruise_control": {
						Description:         "Configuration for Cruise Control deployment. Deploys a Cruise Control instance when specified.",
						MarkdownDescription: "Configuration for Cruise Control deployment. Deploys a Cruise Control instance when specified.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"broker_capacity": {
								Description:         "The Cruise Control 'brokerCapacity' configuration.",
								MarkdownDescription: "The Cruise Control 'brokerCapacity' configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cpu": {
										Description:         "Broker capacity for CPU resource in cores or millicores. For example, 1, 1.500, 1500m. For more information on valid CPU resource units see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.",
										MarkdownDescription: "Broker capacity for CPU resource in cores or millicores. For example, 1, 1.500, 1500m. For more information on valid CPU resource units see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([.][0-9]{0,3}|[m]?)$`), ""),
										},
									},

									"cpu_utilization": {
										Description:         "Broker capacity for CPU resource utilization as a percentage (0 - 100).",
										MarkdownDescription: "Broker capacity for CPU resource utilization as a percentage (0 - 100).",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),

											int64validator.AtMost(100),
										},
									},

									"disk": {
										Description:         "Broker capacity for disk in bytes. Use a number value with either standard Kubernetes byte units (K, M, G, or T), their bibyte (power of two) equivalents (Ki, Mi, Gi, or Ti), or a byte value with or without E notation. For example, 100000M, 100000Mi, 104857600000, or 1e+11.",
										MarkdownDescription: "Broker capacity for disk in bytes. Use a number value with either standard Kubernetes byte units (K, M, G, or T), their bibyte (power of two) equivalents (Ki, Mi, Gi, or Ti), or a byte value with or without E notation. For example, 100000M, 100000Mi, 104857600000, or 1e+11.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([.][0-9]*)?([KMGTPE]i?|e[0-9]+)?$`), ""),
										},
									},

									"inbound_network": {
										Description:         "Broker capacity for inbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
										MarkdownDescription: "Broker capacity for inbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([KMG]i?)?B/s$`), ""),
										},
									},

									"outbound_network": {
										Description:         "Broker capacity for outbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
										MarkdownDescription: "Broker capacity for outbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([KMG]i?)?B/s$`), ""),
										},
									},

									"overrides": {
										Description:         "Overrides for individual brokers. The 'overrides' property lets you specify a different capacity configuration for different brokers.",
										MarkdownDescription: "Overrides for individual brokers. The 'overrides' property lets you specify a different capacity configuration for different brokers.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"brokers": {
												Description:         "List of Kafka brokers (broker identifiers).",
												MarkdownDescription: "List of Kafka brokers (broker identifiers).",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"cpu": {
												Description:         "Broker capacity for CPU resource in cores or millicores. For example, 1, 1.500, 1500m. For more information on valid CPU resource units see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.",
												MarkdownDescription: "Broker capacity for CPU resource in cores or millicores. For example, 1, 1.500, 1500m. For more information on valid CPU resource units see https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/#meaning-of-cpu.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([.][0-9]{0,3}|[m]?)$`), ""),
												},
											},

											"inbound_network": {
												Description:         "Broker capacity for inbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
												MarkdownDescription: "Broker capacity for inbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([KMG]i?)?B/s$`), ""),
												},
											},

											"outbound_network": {
												Description:         "Broker capacity for outbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",
												MarkdownDescription: "Broker capacity for outbound network throughput in bytes per second. Use an integer value with standard Kubernetes byte units (K, M, G) or their bibyte (power of two) equivalents (Ki, Mi, Gi) per second. For example, 10000KiB/s.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+([KMG]i?)?B/s$`), ""),
												},
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

							"config": {
								Description:         "The Cruise Control configuration. For a full list of configuration options refer to https://github.com/linkedin/cruise-control/wiki/Configurations. Note that properties with the following prefixes cannot be set: bootstrap.servers, client.id, zookeeper., network., security., failed.brokers.zk.path,webserver.http., webserver.api.urlprefix, webserver.session.path, webserver.accesslog., two.step., request.reason.required,metric.reporter.sampler.bootstrap.servers, capacity.config.file, self.healing., ssl., kafka.broker.failure.detection.enable, topic.config.provider.class (with the exception of: ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols, webserver.http.cors.enabled, webserver.http.cors.origin, webserver.http.cors.exposeheaders, webserver.security.enable, webserver.ssl.enable).",
								MarkdownDescription: "The Cruise Control configuration. For a full list of configuration options refer to https://github.com/linkedin/cruise-control/wiki/Configurations. Note that properties with the following prefixes cannot be set: bootstrap.servers, client.id, zookeeper., network., security., failed.brokers.zk.path,webserver.http., webserver.api.urlprefix, webserver.session.path, webserver.accesslog., two.step., request.reason.required,metric.reporter.sampler.bootstrap.servers, capacity.config.file, self.healing., ssl., kafka.broker.failure.detection.enable, topic.config.provider.class (with the exception of: ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols, webserver.http.cors.enabled, webserver.http.cors.origin, webserver.http.cors.exposeheaders, webserver.security.enable, webserver.ssl.enable).",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "The docker image for the pods.",
								MarkdownDescription: "The docker image for the pods.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jvm_options": {
								Description:         "JVM Options for the Cruise Control container.",
								MarkdownDescription: "JVM Options for the Cruise Control container.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"xx": {
										Description:         "A map of -XX options to the JVM.",
										MarkdownDescription: "A map of -XX options to the JVM.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"xms": {
										Description:         "-Xms option to to the JVM.",
										MarkdownDescription: "-Xms option to to the JVM.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
										},
									},

									"xmx": {
										Description:         "-Xmx option to to the JVM.",
										MarkdownDescription: "-Xmx option to to the JVM.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
										},
									},

									"gc_logging_enabled": {
										Description:         "Specifies whether the Garbage Collection logging is enabled. The default is false.",
										MarkdownDescription: "Specifies whether the Garbage Collection logging is enabled. The default is false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"java_system_properties": {
										Description:         "A map of additional system properties which will be passed using the '-D' option to the JVM.",
										MarkdownDescription: "A map of additional system properties which will be passed using the '-D' option to the JVM.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "The system property name.",
												MarkdownDescription: "The system property name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "The system property value.",
												MarkdownDescription: "The system property value.",

												Type: types.StringType,

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
								Description:         "Pod liveness checking for the Cruise Control container.",
								MarkdownDescription: "Pod liveness checking for the Cruise Control container.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": {
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": {
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": {
								Description:         "Logging configuration (Log4j 2) for Cruise Control.",
								MarkdownDescription: "Logging configuration (Log4j 2) for Cruise Control.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"loggers": {
										Description:         "A Map from logger name to logger level.",
										MarkdownDescription: "A Map from logger name to logger level.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Logging type, must be either 'inline' or 'external'.",
										MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("inline", "external"),
										},
									},

									"value_from": {
										Description:         "'ConfigMap' entry where the logging configuration is stored. ",
										MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
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

							"metrics_config": {
								Description:         "Metrics configuration.",
								MarkdownDescription: "Metrics configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"type": {
										Description:         "Metrics type. Only 'jmxPrometheusExporter' supported currently.",
										MarkdownDescription: "Metrics type. Only 'jmxPrometheusExporter' supported currently.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("jmxPrometheusExporter"),
										},
									},

									"value_from": {
										Description:         "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",
										MarkdownDescription: "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"readiness_probe": {
								Description:         "Pod readiness checking for the Cruise Control container.",
								MarkdownDescription: "Pod readiness checking for the Cruise Control container.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": {
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": {
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "CPU and memory resources to reserve for the Cruise Control container.",
								MarkdownDescription: "CPU and memory resources to reserve for the Cruise Control container.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requests": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

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
								Description:         "Template to specify how Cruise Control resources, 'Deployments' and 'Pods', are generated.",
								MarkdownDescription: "Template to specify how Cruise Control resources, 'Deployments' and 'Pods', are generated.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"api_service": {
										Description:         "Template for Cruise Control API 'Service'.",
										MarkdownDescription: "Template for Cruise Control API 'Service'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ip_families": {
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_family_policy": {
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
											},

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"cruise_control_container": {
										Description:         "Template for the Cruise Control container.",
										MarkdownDescription: "Template for the Cruise Control container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment": {
										Description:         "Template for Cruise Control 'Deployment'.",
										MarkdownDescription: "Template for Cruise Control 'Deployment'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"pod": {
										Description:         "Template for Cruise Control 'Pods'.",
										MarkdownDescription: "Template for Cruise Control 'Pods'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"affinity": {
												Description:         "The pod's affinity rules.",
												MarkdownDescription: "The pod's affinity rules.",

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

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

											"enable_service_links": {
												Description:         "Indicates whether information about services should be injected into Pod's environment variables.",
												MarkdownDescription: "Indicates whether information about services should be injected into Pod's environment variables.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_aliases": {
												Description:         "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",
												MarkdownDescription: "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",

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

											"image_pull_secrets": {
												Description:         "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
												MarkdownDescription: "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",

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

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scheduler_name": {
												Description:         "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
												MarkdownDescription: "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": {
												Description:         "Configures pod-level security attributes and common container settings.",
												MarkdownDescription: "Configures pod-level security attributes and common container settings.",

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

											"termination_grace_period_seconds": {
												Description:         "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
												MarkdownDescription: "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"tmp_dir_size_limit": {
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
												},
											},

											"tolerations": {
												Description:         "The pod's tolerations.",
												MarkdownDescription: "The pod's tolerations.",

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
												Description:         "The pod's topology spread constraints.",
												MarkdownDescription: "The pod's topology spread constraints.",

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

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_label_keys": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_skew": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"min_domains": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_affinity_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_taints_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"when_unsatisfiable": {
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

									"pod_disruption_budget": {
										Description:         "Template for Cruise Control 'PodDisruptionBudget'.",
										MarkdownDescription: "Template for Cruise Control 'PodDisruptionBudget'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max_unavailable": {
												Description:         "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
												MarkdownDescription: "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"metadata": {
												Description:         "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
												MarkdownDescription: "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
										Description:         "Template for the Cruise Control service account.",
										MarkdownDescription: "Template for the Cruise Control service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"tls_sidecar_container": {
										Description:         "Template for the Cruise Control TLS sidecar container.",
										MarkdownDescription: "Template for the Cruise Control TLS sidecar container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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

							"tls_sidecar": {
								Description:         "TLS sidecar configuration.",
								MarkdownDescription: "TLS sidecar configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image": {
										Description:         "The docker image for the container.",
										MarkdownDescription: "The docker image for the container.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_probe": {
										Description:         "Pod liveness checking.",
										MarkdownDescription: "Pod liveness checking.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"failure_threshold": {
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": {
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": {
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_level": {
										Description:         "The log level for the TLS sidecar. Default value is 'notice'.",
										MarkdownDescription: "The log level for the TLS sidecar. Default value is 'notice'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("emerg", "alert", "crit", "err", "warning", "notice", "info", "debug"),
										},
									},

									"readiness_probe": {
										Description:         "Pod readiness checking.",
										MarkdownDescription: "Pod readiness checking.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"failure_threshold": {
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": {
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": {
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "CPU and memory resources to reserve.",
										MarkdownDescription: "CPU and memory resources to reserve.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

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

					"entity_operator": {
						Description:         "Configuration of the Entity Operator.",
						MarkdownDescription: "Configuration of the Entity Operator.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"template": {
								Description:         "Template for Entity Operator resources. The template allows users to specify how a 'Deployment' and 'Pod' is generated.",
								MarkdownDescription: "Template for Entity Operator resources. The template allows users to specify how a 'Deployment' and 'Pod' is generated.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"deployment": {
										Description:         "Template for Entity Operator 'Deployment'.",
										MarkdownDescription: "Template for Entity Operator 'Deployment'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"pod": {
										Description:         "Template for Entity Operator 'Pods'.",
										MarkdownDescription: "Template for Entity Operator 'Pods'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"affinity": {
												Description:         "The pod's affinity rules.",
												MarkdownDescription: "The pod's affinity rules.",

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

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

											"enable_service_links": {
												Description:         "Indicates whether information about services should be injected into Pod's environment variables.",
												MarkdownDescription: "Indicates whether information about services should be injected into Pod's environment variables.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_aliases": {
												Description:         "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",
												MarkdownDescription: "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",

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

											"image_pull_secrets": {
												Description:         "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
												MarkdownDescription: "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",

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

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scheduler_name": {
												Description:         "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
												MarkdownDescription: "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": {
												Description:         "Configures pod-level security attributes and common container settings.",
												MarkdownDescription: "Configures pod-level security attributes and common container settings.",

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

											"termination_grace_period_seconds": {
												Description:         "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
												MarkdownDescription: "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"tmp_dir_size_limit": {
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
												},
											},

											"tolerations": {
												Description:         "The pod's tolerations.",
												MarkdownDescription: "The pod's tolerations.",

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
												Description:         "The pod's topology spread constraints.",
												MarkdownDescription: "The pod's topology spread constraints.",

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

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_label_keys": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_skew": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"min_domains": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_affinity_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_taints_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"when_unsatisfiable": {
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
										Description:         "Template for the Entity Operator service account.",
										MarkdownDescription: "Template for the Entity Operator service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"tls_sidecar_container": {
										Description:         "Template for the Entity Operator TLS sidecar container.",
										MarkdownDescription: "Template for the Entity Operator TLS sidecar container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"topic_operator_container": {
										Description:         "Template for the Entity Topic Operator container.",
										MarkdownDescription: "Template for the Entity Topic Operator container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"user_operator_container": {
										Description:         "Template for the Entity User Operator container.",
										MarkdownDescription: "Template for the Entity User Operator container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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

							"tls_sidecar": {
								Description:         "TLS sidecar configuration.",
								MarkdownDescription: "TLS sidecar configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image": {
										Description:         "The docker image for the container.",
										MarkdownDescription: "The docker image for the container.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_probe": {
										Description:         "Pod liveness checking.",
										MarkdownDescription: "Pod liveness checking.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"failure_threshold": {
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": {
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": {
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_level": {
										Description:         "The log level for the TLS sidecar. Default value is 'notice'.",
										MarkdownDescription: "The log level for the TLS sidecar. Default value is 'notice'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("emerg", "alert", "crit", "err", "warning", "notice", "info", "debug"),
										},
									},

									"readiness_probe": {
										Description:         "Pod readiness checking.",
										MarkdownDescription: "Pod readiness checking.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"failure_threshold": {
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": {
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": {
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "CPU and memory resources to reserve.",
										MarkdownDescription: "CPU and memory resources to reserve.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

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

							"topic_operator": {
								Description:         "Configuration of the Topic Operator.",
								MarkdownDescription: "Configuration of the Topic Operator.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image": {
										Description:         "The image to use for the Topic Operator.",
										MarkdownDescription: "The image to use for the Topic Operator.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jvm_options": {
										Description:         "JVM Options for pods.",
										MarkdownDescription: "JVM Options for pods.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"xx": {
												Description:         "A map of -XX options to the JVM.",
												MarkdownDescription: "A map of -XX options to the JVM.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"xms": {
												Description:         "-Xms option to to the JVM.",
												MarkdownDescription: "-Xms option to to the JVM.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
												},
											},

											"xmx": {
												Description:         "-Xmx option to to the JVM.",
												MarkdownDescription: "-Xmx option to to the JVM.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
												},
											},

											"gc_logging_enabled": {
												Description:         "Specifies whether the Garbage Collection logging is enabled. The default is false.",
												MarkdownDescription: "Specifies whether the Garbage Collection logging is enabled. The default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"java_system_properties": {
												Description:         "A map of additional system properties which will be passed using the '-D' option to the JVM.",
												MarkdownDescription: "A map of additional system properties which will be passed using the '-D' option to the JVM.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The system property name.",
														MarkdownDescription: "The system property name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The system property value.",
														MarkdownDescription: "The system property value.",

														Type: types.StringType,

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
										Description:         "Pod liveness checking.",
										MarkdownDescription: "Pod liveness checking.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"failure_threshold": {
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": {
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": {
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"logging": {
										Description:         "Logging configuration.",
										MarkdownDescription: "Logging configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"loggers": {
												Description:         "A Map from logger name to logger level.",
												MarkdownDescription: "A Map from logger name to logger level.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "Logging type, must be either 'inline' or 'external'.",
												MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("inline", "external"),
												},
											},

											"value_from": {
												Description:         "'ConfigMap' entry where the logging configuration is stored. ",
												MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_key_ref": {
														Description:         "Reference to the key in the ConfigMap containing the configuration.",
														MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
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

									"readiness_probe": {
										Description:         "Pod readiness checking.",
										MarkdownDescription: "Pod readiness checking.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"failure_threshold": {
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": {
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": {
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reconciliation_interval_seconds": {
										Description:         "Interval between periodic reconciliations.",
										MarkdownDescription: "Interval between periodic reconciliations.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"resources": {
										Description:         "CPU and memory resources to reserve.",
										MarkdownDescription: "CPU and memory resources to reserve.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

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
										Description:         "Pod startup checking.",
										MarkdownDescription: "Pod startup checking.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"failure_threshold": {
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": {
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": {
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"topic_metadata_max_attempts": {
										Description:         "The number of attempts at getting topic metadata.",
										MarkdownDescription: "The number of attempts at getting topic metadata.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"watched_namespace": {
										Description:         "The namespace the Topic Operator should watch.",
										MarkdownDescription: "The namespace the Topic Operator should watch.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"zookeeper_session_timeout_seconds": {
										Description:         "Timeout for the ZooKeeper session.",
										MarkdownDescription: "Timeout for the ZooKeeper session.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"user_operator": {
								Description:         "Configuration of the User Operator.",
								MarkdownDescription: "Configuration of the User Operator.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image": {
										Description:         "The image to use for the User Operator.",
										MarkdownDescription: "The image to use for the User Operator.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jvm_options": {
										Description:         "JVM Options for pods.",
										MarkdownDescription: "JVM Options for pods.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"xx": {
												Description:         "A map of -XX options to the JVM.",
												MarkdownDescription: "A map of -XX options to the JVM.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"xms": {
												Description:         "-Xms option to to the JVM.",
												MarkdownDescription: "-Xms option to to the JVM.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
												},
											},

											"xmx": {
												Description:         "-Xmx option to to the JVM.",
												MarkdownDescription: "-Xmx option to to the JVM.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
												},
											},

											"gc_logging_enabled": {
												Description:         "Specifies whether the Garbage Collection logging is enabled. The default is false.",
												MarkdownDescription: "Specifies whether the Garbage Collection logging is enabled. The default is false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"java_system_properties": {
												Description:         "A map of additional system properties which will be passed using the '-D' option to the JVM.",
												MarkdownDescription: "A map of additional system properties which will be passed using the '-D' option to the JVM.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The system property name.",
														MarkdownDescription: "The system property name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The system property value.",
														MarkdownDescription: "The system property value.",

														Type: types.StringType,

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
										Description:         "Pod liveness checking.",
										MarkdownDescription: "Pod liveness checking.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"failure_threshold": {
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": {
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": {
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"logging": {
										Description:         "Logging configuration.",
										MarkdownDescription: "Logging configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"loggers": {
												Description:         "A Map from logger name to logger level.",
												MarkdownDescription: "A Map from logger name to logger level.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "Logging type, must be either 'inline' or 'external'.",
												MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("inline", "external"),
												},
											},

											"value_from": {
												Description:         "'ConfigMap' entry where the logging configuration is stored. ",
												MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config_map_key_ref": {
														Description:         "Reference to the key in the ConfigMap containing the configuration.",
														MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
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

									"readiness_probe": {
										Description:         "Pod readiness checking.",
										MarkdownDescription: "Pod readiness checking.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"failure_threshold": {
												Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"initial_delay_seconds": {
												Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
												MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"timeout_seconds": {
												Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
												MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"reconciliation_interval_seconds": {
										Description:         "Interval between periodic reconciliations.",
										MarkdownDescription: "Interval between periodic reconciliations.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"resources": {
										Description:         "CPU and memory resources to reserve.",
										MarkdownDescription: "CPU and memory resources to reserve.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "",
												MarkdownDescription: "",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_prefix": {
										Description:         "The prefix that will be added to the KafkaUser name to be used as the Secret name.",
										MarkdownDescription: "The prefix that will be added to the KafkaUser name to be used as the Secret name.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"watched_namespace": {
										Description:         "The namespace the User Operator should watch.",
										MarkdownDescription: "The namespace the User Operator should watch.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"zookeeper_session_timeout_seconds": {
										Description:         "Timeout for the ZooKeeper session.",
										MarkdownDescription: "Timeout for the ZooKeeper session.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
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

					"jmx_trans": {
						Description:         "Configuration for JmxTrans. When the property is present a JmxTrans deployment is created for gathering JMX metrics from each Kafka broker. For more information see https://github.com/jmxtrans/jmxtrans[JmxTrans GitHub].",
						MarkdownDescription: "Configuration for JmxTrans. When the property is present a JmxTrans deployment is created for gathering JMX metrics from each Kafka broker. For more information see https://github.com/jmxtrans/jmxtrans[JmxTrans GitHub].",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"image": {
								Description:         "The image to use for the JmxTrans.",
								MarkdownDescription: "The image to use for the JmxTrans.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kafka_queries": {
								Description:         "Queries to send to the Kafka brokers to define what data should be read from each broker. For more information on these properties see, xref:type-JmxTransQueryTemplate-reference['JmxTransQueryTemplate' schema reference].",
								MarkdownDescription: "Queries to send to the Kafka brokers to define what data should be read from each broker. For more information on these properties see, xref:type-JmxTransQueryTemplate-reference['JmxTransQueryTemplate' schema reference].",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"attributes": {
										Description:         "Determine which attributes of the targeted MBean should be included.",
										MarkdownDescription: "Determine which attributes of the targeted MBean should be included.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"outputs": {
										Description:         "List of the names of output definitions specified in the spec.kafka.jmxTrans.outputDefinitions that have defined where JMX metrics are pushed to, and in which data format.",
										MarkdownDescription: "List of the names of output definitions specified in the spec.kafka.jmxTrans.outputDefinitions that have defined where JMX metrics are pushed to, and in which data format.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"target_m_bean": {
										Description:         "If using wildcards instead of a specific MBean then the data is gathered from multiple MBeans. Otherwise if specifying an MBean then data is gathered from that specified MBean.",
										MarkdownDescription: "If using wildcards instead of a specific MBean then the data is gathered from multiple MBeans. Otherwise if specifying an MBean then data is gathered from that specified MBean.",

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

							"log_level": {
								Description:         "Sets the logging level of the JmxTrans deployment.For more information see, https://github.com/jmxtrans/jmxtrans-agent/wiki/Troubleshooting[JmxTrans Logging Level].",
								MarkdownDescription: "Sets the logging level of the JmxTrans deployment.For more information see, https://github.com/jmxtrans/jmxtrans-agent/wiki/Troubleshooting[JmxTrans Logging Level].",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"output_definitions": {
								Description:         "Defines the output hosts that will be referenced later on. For more information on these properties see, xref:type-JmxTransOutputDefinitionTemplate-reference['JmxTransOutputDefinitionTemplate' schema reference].",
								MarkdownDescription: "Defines the output hosts that will be referenced later on. For more information on these properties see, xref:type-JmxTransOutputDefinitionTemplate-reference['JmxTransOutputDefinitionTemplate' schema reference].",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"flush_delay_in_seconds": {
										Description:         "How many seconds the JmxTrans waits before pushing a new set of data out.",
										MarkdownDescription: "How many seconds the JmxTrans waits before pushing a new set of data out.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host": {
										Description:         "The DNS/hostname of the remote host that the data is pushed to.",
										MarkdownDescription: "The DNS/hostname of the remote host that the data is pushed to.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Template for setting the name of the output definition. This is used to identify where to send the results of queries should be sent.",
										MarkdownDescription: "Template for setting the name of the output definition. This is used to identify where to send the results of queries should be sent.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"output_type": {
										Description:         "Template for setting the format of the data that will be pushed.For more information see https://github.com/jmxtrans/jmxtrans/wiki/OutputWriters[JmxTrans OutputWriters].",
										MarkdownDescription: "Template for setting the format of the data that will be pushed.For more information see https://github.com/jmxtrans/jmxtrans/wiki/OutputWriters[JmxTrans OutputWriters].",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"port": {
										Description:         "The port of the remote host that the data is pushed to.",
										MarkdownDescription: "The port of the remote host that the data is pushed to.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type_names": {
										Description:         "Template for filtering data to be included in response to a wildcard query. For more information see https://github.com/jmxtrans/jmxtrans/wiki/Queries[JmxTrans queries].",
										MarkdownDescription: "Template for filtering data to be included in response to a wildcard query. For more information see https://github.com/jmxtrans/jmxtrans/wiki/Queries[JmxTrans queries].",

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

							"resources": {
								Description:         "CPU and memory resources to reserve.",
								MarkdownDescription: "CPU and memory resources to reserve.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requests": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

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
								Description:         "Template for JmxTrans resources.",
								MarkdownDescription: "Template for JmxTrans resources.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"container": {
										Description:         "Template for JmxTrans container.",
										MarkdownDescription: "Template for JmxTrans container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment": {
										Description:         "Template for JmxTrans 'Deployment'.",
										MarkdownDescription: "Template for JmxTrans 'Deployment'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"pod": {
										Description:         "Template for JmxTrans 'Pods'.",
										MarkdownDescription: "Template for JmxTrans 'Pods'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"affinity": {
												Description:         "The pod's affinity rules.",
												MarkdownDescription: "The pod's affinity rules.",

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

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

											"enable_service_links": {
												Description:         "Indicates whether information about services should be injected into Pod's environment variables.",
												MarkdownDescription: "Indicates whether information about services should be injected into Pod's environment variables.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_aliases": {
												Description:         "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",
												MarkdownDescription: "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",

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

											"image_pull_secrets": {
												Description:         "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
												MarkdownDescription: "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",

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

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scheduler_name": {
												Description:         "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
												MarkdownDescription: "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": {
												Description:         "Configures pod-level security attributes and common container settings.",
												MarkdownDescription: "Configures pod-level security attributes and common container settings.",

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

											"termination_grace_period_seconds": {
												Description:         "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
												MarkdownDescription: "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"tmp_dir_size_limit": {
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
												},
											},

											"tolerations": {
												Description:         "The pod's tolerations.",
												MarkdownDescription: "The pod's tolerations.",

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
												Description:         "The pod's topology spread constraints.",
												MarkdownDescription: "The pod's topology spread constraints.",

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

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_label_keys": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_skew": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"min_domains": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_affinity_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_taints_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"when_unsatisfiable": {
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
										Description:         "Template for the JMX Trans service account.",
										MarkdownDescription: "Template for the JMX Trans service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

					"kafka": {
						Description:         "Configuration of the Kafka cluster.",
						MarkdownDescription: "Configuration of the Kafka cluster.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"authorization": {
								Description:         "Authorization configuration for Kafka brokers.",
								MarkdownDescription: "Authorization configuration for Kafka brokers.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"allow_on_error": {
										Description:         "Defines whether a Kafka client should be allowed or denied by default when the authorizer fails to query the Open Policy Agent, for example, when it is temporarily unavailable). Defaults to 'false' - all actions will be denied.",
										MarkdownDescription: "Defines whether a Kafka client should be allowed or denied by default when the authorizer fails to query the Open Policy Agent, for example, when it is temporarily unavailable). Defaults to 'false' - all actions will be denied.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorizer_class": {
										Description:         "Authorization implementation class, which must be available in classpath.",
										MarkdownDescription: "Authorization implementation class, which must be available in classpath.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_id": {
										Description:         "OAuth Client ID which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",
										MarkdownDescription: "OAuth Client ID which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"connect_timeout_seconds": {
										Description:         "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",
										MarkdownDescription: "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"delegate_to_kafka_acls": {
										Description:         "Whether authorization decision should be delegated to the 'Simple' authorizer if DENIED by Keycloak Authorization Services policies. Default value is 'false'.",
										MarkdownDescription: "Whether authorization decision should be delegated to the 'Simple' authorizer if DENIED by Keycloak Authorization Services policies. Default value is 'false'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_tls_hostname_verification": {
										Description:         "Enable or disable TLS hostname verification. Default value is 'false'.",
										MarkdownDescription: "Enable or disable TLS hostname verification. Default value is 'false'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enable_metrics": {
										Description:         "Enable or disable OAuth metrics. Default value is 'false'.",
										MarkdownDescription: "Enable or disable OAuth metrics. Default value is 'false'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"expire_after_ms": {
										Description:         "The expiration of the records kept in the local cache to avoid querying the Open Policy Agent for every request. Defines how often the cached authorization decisions are reloaded from the Open Policy Agent server. In milliseconds. Defaults to '3600000'.",
										MarkdownDescription: "The expiration of the records kept in the local cache to avoid querying the Open Policy Agent for every request. Defines how often the cached authorization decisions are reloaded from the Open Policy Agent server. In milliseconds. Defaults to '3600000'.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grants_refresh_period_seconds": {
										Description:         "The time between two consecutive grants refresh runs in seconds. The default value is 60.",
										MarkdownDescription: "The time between two consecutive grants refresh runs in seconds. The default value is 60.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"grants_refresh_pool_size": {
										Description:         "The number of threads to use to refresh grants for active sessions. The more threads, the more parallelism, so the sooner the job completes. However, using more threads places a heavier load on the authorization server. The default value is 5.",
										MarkdownDescription: "The number of threads to use to refresh grants for active sessions. The more threads, the more parallelism, so the sooner the job completes. However, using more threads places a heavier load on the authorization server. The default value is 5.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"initial_cache_capacity": {
										Description:         "Initial capacity of the local cache used by the authorizer to avoid querying the Open Policy Agent for every request Defaults to '5000'.",
										MarkdownDescription: "Initial capacity of the local cache used by the authorizer to avoid querying the Open Policy Agent for every request Defaults to '5000'.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"maximum_cache_size": {
										Description:         "Maximum capacity of the local cache used by the authorizer to avoid querying the Open Policy Agent for every request. Defaults to '50000'.",
										MarkdownDescription: "Maximum capacity of the local cache used by the authorizer to avoid querying the Open Policy Agent for every request. Defaults to '50000'.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"read_timeout_seconds": {
										Description:         "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",
										MarkdownDescription: "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"super_users": {
										Description:         "List of super users, which are user principals with unlimited access rights.",
										MarkdownDescription: "List of super users, which are user principals with unlimited access rights.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"supports_admin_api": {
										Description:         "Indicates whether the custom authorizer supports the APIs for managing ACLs using the Kafka Admin API. Defaults to 'false'.",
										MarkdownDescription: "Indicates whether the custom authorizer supports the APIs for managing ACLs using the Kafka Admin API. Defaults to 'false'.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_trusted_certificates": {
										Description:         "Trusted certificates for TLS connection to the OAuth server.",
										MarkdownDescription: "Trusted certificates for TLS connection to the OAuth server.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"certificate": {
												Description:         "The name of the file certificate in the Secret.",
												MarkdownDescription: "The name of the file certificate in the Secret.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"secret_name": {
												Description:         "The name of the Secret containing the certificate.",
												MarkdownDescription: "The name of the Secret containing the certificate.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"token_endpoint_uri": {
										Description:         "Authorization server token endpoint URI.",
										MarkdownDescription: "Authorization server token endpoint URI.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Authorization type. Currently, the supported types are 'simple', 'keycloak', 'opa' and 'custom'. 'simple' authorization type uses Kafka's 'kafka.security.authorizer.AclAuthorizer' class for authorization. 'keycloak' authorization type uses Keycloak Authorization Services for authorization. 'opa' authorization type uses Open Policy Agent based authorization.'custom' authorization type uses user-provided implementation for authorization.",
										MarkdownDescription: "Authorization type. Currently, the supported types are 'simple', 'keycloak', 'opa' and 'custom'. 'simple' authorization type uses Kafka's 'kafka.security.authorizer.AclAuthorizer' class for authorization. 'keycloak' authorization type uses Keycloak Authorization Services for authorization. 'opa' authorization type uses Open Policy Agent based authorization.'custom' authorization type uses user-provided implementation for authorization.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("simple", "opa", "keycloak", "custom"),
										},
									},

									"url": {
										Description:         "The URL used to connect to the Open Policy Agent server. The URL has to include the policy which will be queried by the authorizer. This option is required.",
										MarkdownDescription: "The URL used to connect to the Open Policy Agent server. The URL has to include the policy which will be queried by the authorizer. This option is required.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"broker_rack_init_image": {
								Description:         "The image of the init container used for initializing the 'broker.rack'.",
								MarkdownDescription: "The image of the init container used for initializing the 'broker.rack'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"config": {
								Description:         "Kafka broker config properties with the following prefixes cannot be set: listeners, advertised., broker., listener., host.name, port, inter.broker.listener.name, sasl., ssl., security., password., log.dir, zookeeper.connect, zookeeper.set.acl, zookeeper.ssl, zookeeper.clientCnxnSocket, authorizer., super.user, cruise.control.metrics.topic, cruise.control.metrics.reporter.bootstrap.servers,node.id, process.roles, controller. (with the exception of: zookeeper.connection.timeout.ms, ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols, sasl.server.max.receive.size,cruise.control.metrics.topic.num.partitions, cruise.control.metrics.topic.replication.factor, cruise.control.metrics.topic.retention.ms,cruise.control.metrics.topic.auto.create.retries, cruise.control.metrics.topic.auto.create.timeout.ms,cruise.control.metrics.topic.min.insync.replicas,controller.quorum.election.backoff.max.ms, controller.quorum.election.timeout.ms, controller.quorum.fetch.timeout.ms).",
								MarkdownDescription: "Kafka broker config properties with the following prefixes cannot be set: listeners, advertised., broker., listener., host.name, port, inter.broker.listener.name, sasl., ssl., security., password., log.dir, zookeeper.connect, zookeeper.set.acl, zookeeper.ssl, zookeeper.clientCnxnSocket, authorizer., super.user, cruise.control.metrics.topic, cruise.control.metrics.reporter.bootstrap.servers,node.id, process.roles, controller. (with the exception of: zookeeper.connection.timeout.ms, ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols, sasl.server.max.receive.size,cruise.control.metrics.topic.num.partitions, cruise.control.metrics.topic.replication.factor, cruise.control.metrics.topic.retention.ms,cruise.control.metrics.topic.auto.create.retries, cruise.control.metrics.topic.auto.create.timeout.ms,cruise.control.metrics.topic.min.insync.replicas,controller.quorum.election.backoff.max.ms, controller.quorum.election.timeout.ms, controller.quorum.fetch.timeout.ms).",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "The docker image for the pods. The default value depends on the configured 'Kafka.spec.kafka.version'.",
								MarkdownDescription: "The docker image for the pods. The default value depends on the configured 'Kafka.spec.kafka.version'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jmx_options": {
								Description:         "JMX Options for Kafka brokers.",
								MarkdownDescription: "JMX Options for Kafka brokers.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"authentication": {
										Description:         "Authentication configuration for connecting to the JMX port.",
										MarkdownDescription: "Authentication configuration for connecting to the JMX port.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"type": {
												Description:         "Authentication type. Currently the only supported types are 'password'.'password' type creates a username and protected port with no TLS.",
												MarkdownDescription: "Authentication type. Currently the only supported types are 'password'.'password' type creates a username and protected port with no TLS.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("password"),
												},
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

							"jvm_options": {
								Description:         "JVM Options for pods.",
								MarkdownDescription: "JVM Options for pods.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"xx": {
										Description:         "A map of -XX options to the JVM.",
										MarkdownDescription: "A map of -XX options to the JVM.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"xms": {
										Description:         "-Xms option to to the JVM.",
										MarkdownDescription: "-Xms option to to the JVM.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
										},
									},

									"xmx": {
										Description:         "-Xmx option to to the JVM.",
										MarkdownDescription: "-Xmx option to to the JVM.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
										},
									},

									"gc_logging_enabled": {
										Description:         "Specifies whether the Garbage Collection logging is enabled. The default is false.",
										MarkdownDescription: "Specifies whether the Garbage Collection logging is enabled. The default is false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"java_system_properties": {
										Description:         "A map of additional system properties which will be passed using the '-D' option to the JVM.",
										MarkdownDescription: "A map of additional system properties which will be passed using the '-D' option to the JVM.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "The system property name.",
												MarkdownDescription: "The system property name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "The system property value.",
												MarkdownDescription: "The system property value.",

												Type: types.StringType,

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

							"listeners": {
								Description:         "Configures listeners of Kafka brokers.",
								MarkdownDescription: "Configures listeners of Kafka brokers.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"authentication": {
										Description:         "Authentication configuration for this listener.",
										MarkdownDescription: "Authentication configuration for this listener.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"access_token_is_jwt": {
												Description:         "Configure whether the access token is treated as JWT. This must be set to 'false' if the authorization server returns opaque tokens. Defaults to 'true'.",
												MarkdownDescription: "Configure whether the access token is treated as JWT. This must be set to 'false' if the authorization server returns opaque tokens. Defaults to 'true'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"check_access_token_type": {
												Description:         "Configure whether the access token type check is performed or not. This should be set to 'false' if the authorization server does not include 'typ' claim in JWT token. Defaults to 'true'.",
												MarkdownDescription: "Configure whether the access token type check is performed or not. This should be set to 'false' if the authorization server does not include 'typ' claim in JWT token. Defaults to 'true'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"check_audience": {
												Description:         "Enable or disable audience checking. Audience checks identify the recipients of tokens. If audience checking is enabled, the OAuth Client ID also has to be configured using the 'clientId' property. The Kafka broker will reject tokens that do not have its 'clientId' in their 'aud' (audience) claim.Default value is 'false'.",
												MarkdownDescription: "Enable or disable audience checking. Audience checks identify the recipients of tokens. If audience checking is enabled, the OAuth Client ID also has to be configured using the 'clientId' property. The Kafka broker will reject tokens that do not have its 'clientId' in their 'aud' (audience) claim.Default value is 'false'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"check_issuer": {
												Description:         "Enable or disable issuer checking. By default issuer is checked using the value configured by 'validIssuerUri'. Default value is 'true'.",
												MarkdownDescription: "Enable or disable issuer checking. By default issuer is checked using the value configured by 'validIssuerUri'. Default value is 'true'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_audience": {
												Description:         "The audience to use when making requests to the authorization server's token endpoint. Used for inter-broker authentication and for configuring OAuth 2.0 over PLAIN using the 'clientId' and 'secret' method.",
												MarkdownDescription: "The audience to use when making requests to the authorization server's token endpoint. Used for inter-broker authentication and for configuring OAuth 2.0 over PLAIN using the 'clientId' and 'secret' method.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_id": {
												Description:         "OAuth Client ID which the Kafka broker can use to authenticate against the authorization server and use the introspect endpoint URI.",
												MarkdownDescription: "OAuth Client ID which the Kafka broker can use to authenticate against the authorization server and use the introspect endpoint URI.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_scope": {
												Description:         "The scope to use when making requests to the authorization server's token endpoint. Used for inter-broker authentication and for configuring OAuth 2.0 over PLAIN using the 'clientId' and 'secret' method.",
												MarkdownDescription: "The scope to use when making requests to the authorization server's token endpoint. Used for inter-broker authentication and for configuring OAuth 2.0 over PLAIN using the 'clientId' and 'secret' method.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"client_secret": {
												Description:         "Link to Kubernetes Secret containing the OAuth client secret which the Kafka broker can use to authenticate against the authorization server and use the introspect endpoint URI.",
												MarkdownDescription: "Link to Kubernetes Secret containing the OAuth client secret which the Kafka broker can use to authenticate against the authorization server and use the introspect endpoint URI.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key under which the secret value is stored in the Kubernetes Secret.",
														MarkdownDescription: "The key under which the secret value is stored in the Kubernetes Secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "The name of the Kubernetes Secret containing the secret value.",
														MarkdownDescription: "The name of the Kubernetes Secret containing the secret value.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"connect_timeout_seconds": {
												Description:         "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",
												MarkdownDescription: "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"custom_claim_check": {
												Description:         "JsonPath filter query to be applied to the JWT token or to the response of the introspection endpoint for additional token validation. Not set by default.",
												MarkdownDescription: "JsonPath filter query to be applied to the JWT token or to the response of the introspection endpoint for additional token validation. Not set by default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"disable_tls_hostname_verification": {
												Description:         "Enable or disable TLS hostname verification. Default value is 'false'.",
												MarkdownDescription: "Enable or disable TLS hostname verification. Default value is 'false'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_ecdsa": {
												Description:         "Enable or disable ECDSA support by installing BouncyCastle crypto provider. ECDSA support is always enabled. The BouncyCastle libraries are no longer packaged with Strimzi. Value is ignored.",
												MarkdownDescription: "Enable or disable ECDSA support by installing BouncyCastle crypto provider. ECDSA support is always enabled. The BouncyCastle libraries are no longer packaged with Strimzi. Value is ignored.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_metrics": {
												Description:         "Enable or disable OAuth metrics. Default value is 'false'.",
												MarkdownDescription: "Enable or disable OAuth metrics. Default value is 'false'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_oauth_bearer": {
												Description:         "Enable or disable OAuth authentication over SASL_OAUTHBEARER. Default value is 'true'.",
												MarkdownDescription: "Enable or disable OAuth authentication over SASL_OAUTHBEARER. Default value is 'true'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enable_plain": {
												Description:         "Enable or disable OAuth authentication over SASL_PLAIN. There is no re-authentication support when this mechanism is used. Default value is 'false'.",
												MarkdownDescription: "Enable or disable OAuth authentication over SASL_PLAIN. There is no re-authentication support when this mechanism is used. Default value is 'false'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"fail_fast": {
												Description:         "Enable or disable termination of Kafka broker processes due to potentially recoverable runtime errors during startup. Default value is 'true'.",
												MarkdownDescription: "Enable or disable termination of Kafka broker processes due to potentially recoverable runtime errors during startup. Default value is 'true'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"fallback_user_name_claim": {
												Description:         "The fallback username claim to be used for the user id if the claim specified by 'userNameClaim' is not present. This is useful when 'client_credentials' authentication only results in the client id being provided in another claim. It only takes effect if 'userNameClaim' is set.",
												MarkdownDescription: "The fallback username claim to be used for the user id if the claim specified by 'userNameClaim' is not present. This is useful when 'client_credentials' authentication only results in the client id being provided in another claim. It only takes effect if 'userNameClaim' is set.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"fallback_user_name_prefix": {
												Description:         "The prefix to use with the value of 'fallbackUserNameClaim' to construct the user id. This only takes effect if 'fallbackUserNameClaim' is true, and the value is present for the claim. Mapping usernames and client ids into the same user id space is useful in preventing name collisions.",
												MarkdownDescription: "The prefix to use with the value of 'fallbackUserNameClaim' to construct the user id. This only takes effect if 'fallbackUserNameClaim' is true, and the value is present for the claim. Mapping usernames and client ids into the same user id space is useful in preventing name collisions.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"groups_claim": {
												Description:         "JsonPath query used to extract groups for the user during authentication. Extracted groups can be used by a custom authorizer. By default no groups are extracted.",
												MarkdownDescription: "JsonPath query used to extract groups for the user during authentication. Extracted groups can be used by a custom authorizer. By default no groups are extracted.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"groups_claim_delimiter": {
												Description:         "A delimiter used to parse groups when they are extracted as a single String value rather than a JSON array. Default value is ',' (comma).",
												MarkdownDescription: "A delimiter used to parse groups when they are extracted as a single String value rather than a JSON array. Default value is ',' (comma).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"introspection_endpoint_uri": {
												Description:         "URI of the token introspection endpoint which can be used to validate opaque non-JWT tokens.",
												MarkdownDescription: "URI of the token introspection endpoint which can be used to validate opaque non-JWT tokens.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jwks_endpoint_uri": {
												Description:         "URI of the JWKS certificate endpoint, which can be used for local JWT validation.",
												MarkdownDescription: "URI of the JWKS certificate endpoint, which can be used for local JWT validation.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jwks_expiry_seconds": {
												Description:         "Configures how often are the JWKS certificates considered valid. The expiry interval has to be at least 60 seconds longer then the refresh interval specified in 'jwksRefreshSeconds'. Defaults to 360 seconds.",
												MarkdownDescription: "Configures how often are the JWKS certificates considered valid. The expiry interval has to be at least 60 seconds longer then the refresh interval specified in 'jwksRefreshSeconds'. Defaults to 360 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"jwks_ignore_key_use": {
												Description:         "Flag to ignore the 'use' attribute of 'key' declarations in a JWKS endpoint response. Default value is 'false'.",
												MarkdownDescription: "Flag to ignore the 'use' attribute of 'key' declarations in a JWKS endpoint response. Default value is 'false'.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"jwks_min_refresh_pause_seconds": {
												Description:         "The minimum pause between two consecutive refreshes. When an unknown signing key is encountered the refresh is scheduled immediately, but will always wait for this minimum pause. Defaults to 1 second.",
												MarkdownDescription: "The minimum pause between two consecutive refreshes. When an unknown signing key is encountered the refresh is scheduled immediately, but will always wait for this minimum pause. Defaults to 1 second.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"jwks_refresh_seconds": {
												Description:         "Configures how often are the JWKS certificates refreshed. The refresh interval has to be at least 60 seconds shorter then the expiry interval specified in 'jwksExpirySeconds'. Defaults to 300 seconds.",
												MarkdownDescription: "Configures how often are the JWKS certificates refreshed. The refresh interval has to be at least 60 seconds shorter then the expiry interval specified in 'jwksExpirySeconds'. Defaults to 300 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(1),
												},
											},

											"listener_config": {
												Description:         "Configuration to be used for a specific listener. All values are prefixed with listener.name._<listener_name>_.",
												MarkdownDescription: "Configuration to be used for a specific listener. All values are prefixed with listener.name._<listener_name>_.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_seconds_without_reauthentication": {
												Description:         "Maximum number of seconds the authenticated session remains valid without re-authentication. This enables Apache Kafka re-authentication feature, and causes sessions to expire when the access token expires. If the access token expires before max time or if max time is reached, the client has to re-authenticate, otherwise the server will drop the connection. Not set by default - the authenticated session does not expire when the access token expires. This option only applies to SASL_OAUTHBEARER authentication mechanism (when 'enableOauthBearer' is 'true').",
												MarkdownDescription: "Maximum number of seconds the authenticated session remains valid without re-authentication. This enables Apache Kafka re-authentication feature, and causes sessions to expire when the access token expires. If the access token expires before max time or if max time is reached, the client has to re-authenticate, otherwise the server will drop the connection. Not set by default - the authenticated session does not expire when the access token expires. This option only applies to SASL_OAUTHBEARER authentication mechanism (when 'enableOauthBearer' is 'true').",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_timeout_seconds": {
												Description:         "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",
												MarkdownDescription: "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sasl": {
												Description:         "Enable or disable SASL on this listener.",
												MarkdownDescription: "Enable or disable SASL on this listener.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secrets": {
												Description:         "Secrets to be mounted to /opt/kafka/custom-authn-secrets/custom-listener-_<listener_name>-<port>_/_<secret_name>_.",
												MarkdownDescription: "Secrets to be mounted to /opt/kafka/custom-authn-secrets/custom-listener-_<listener_name>-<port>_/_<secret_name>_.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
														Description:         "The key under which the secret value is stored in the Kubernetes Secret.",
														MarkdownDescription: "The key under which the secret value is stored in the Kubernetes Secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "The name of the Kubernetes Secret containing the secret value.",
														MarkdownDescription: "The name of the Kubernetes Secret containing the secret value.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls_trusted_certificates": {
												Description:         "Trusted certificates for TLS connection to the OAuth server.",
												MarkdownDescription: "Trusted certificates for TLS connection to the OAuth server.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"certificate": {
														Description:         "The name of the file certificate in the Secret.",
														MarkdownDescription: "The name of the file certificate in the Secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "The name of the Secret containing the certificate.",
														MarkdownDescription: "The name of the Secret containing the certificate.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"token_endpoint_uri": {
												Description:         "URI of the Token Endpoint to use with SASL_PLAIN mechanism when the client authenticates with 'clientId' and a 'secret'. If set, the client can authenticate over SASL_PLAIN by either setting 'username' to 'clientId', and setting 'password' to client 'secret', or by setting 'username' to account username, and 'password' to access token prefixed with '$accessToken:'. If this option is not set, the 'password' is always interpreted as an access token (without a prefix), and 'username' as the account username (a so called 'no-client-credentials' mode).",
												MarkdownDescription: "URI of the Token Endpoint to use with SASL_PLAIN mechanism when the client authenticates with 'clientId' and a 'secret'. If set, the client can authenticate over SASL_PLAIN by either setting 'username' to 'clientId', and setting 'password' to client 'secret', or by setting 'username' to account username, and 'password' to access token prefixed with '$accessToken:'. If this option is not set, the 'password' is always interpreted as an access token (without a prefix), and 'username' as the account username (a so called 'no-client-credentials' mode).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "Authentication type. 'oauth' type uses SASL OAUTHBEARER Authentication. 'scram-sha-512' type uses SASL SCRAM-SHA-512 Authentication. 'tls' type uses TLS Client Authentication. 'tls' type is supported only on TLS listeners.'custom' type allows for any authentication type to be used.",
												MarkdownDescription: "Authentication type. 'oauth' type uses SASL OAUTHBEARER Authentication. 'scram-sha-512' type uses SASL SCRAM-SHA-512 Authentication. 'tls' type uses TLS Client Authentication. 'tls' type is supported only on TLS listeners.'custom' type allows for any authentication type to be used.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("tls", "scram-sha-512", "oauth", "custom"),
												},
											},

											"user_info_endpoint_uri": {
												Description:         "URI of the User Info Endpoint to use as a fallback to obtaining the user id when the Introspection Endpoint does not return information that can be used for the user id. ",
												MarkdownDescription: "URI of the User Info Endpoint to use as a fallback to obtaining the user id when the Introspection Endpoint does not return information that can be used for the user id. ",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"user_name_claim": {
												Description:         "Name of the claim from the JWT authentication token, Introspection Endpoint response or User Info Endpoint response which will be used to extract the user id. Defaults to 'sub'.",
												MarkdownDescription: "Name of the claim from the JWT authentication token, Introspection Endpoint response or User Info Endpoint response which will be used to extract the user id. Defaults to 'sub'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"valid_issuer_uri": {
												Description:         "URI of the token issuer used for authentication.",
												MarkdownDescription: "URI of the token issuer used for authentication.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"valid_token_type": {
												Description:         "Valid value for the 'token_type' attribute returned by the Introspection Endpoint. No default value, and not checked by default.",
												MarkdownDescription: "Valid value for the 'token_type' attribute returned by the Introspection Endpoint. No default value, and not checked by default.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"configuration": {
										Description:         "Additional listener configuration.",
										MarkdownDescription: "Additional listener configuration.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"bootstrap": {
												Description:         "Bootstrap configuration.",
												MarkdownDescription: "Bootstrap configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"alternative_names": {
														Description:         "Additional alternative names for the bootstrap service. The alternative names will be added to the list of subject alternative names of the TLS certificates.",
														MarkdownDescription: "Additional alternative names for the bootstrap service. The alternative names will be added to the list of subject alternative names of the TLS certificates.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"annotations": {
														Description:         "Annotations that will be added to the 'Ingress', 'Route', or 'Service' resource. You can use this field to configure DNS providers such as External DNS. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",
														MarkdownDescription: "Annotations that will be added to the 'Ingress', 'Route', or 'Service' resource. You can use this field to configure DNS providers such as External DNS. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"host": {
														Description:         "The bootstrap host. This field will be used in the Ingress resource or in the Route resource to specify the desired hostname. This field can be used only with 'route' (optional) or 'ingress' (required) type listeners.",
														MarkdownDescription: "The bootstrap host. This field will be used in the Ingress resource or in the Route resource to specify the desired hostname. This field can be used only with 'route' (optional) or 'ingress' (required) type listeners.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels that will be added to the 'Ingress', 'Route', or 'Service' resource. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",
														MarkdownDescription: "Labels that will be added to the 'Ingress', 'Route', or 'Service' resource. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"load_balancer_ip": {
														Description:         "The loadbalancer is requested with the IP address specified in this field. This feature depends on whether the underlying cloud provider supports specifying the 'loadBalancerIP' when a load balancer is created. This field is ignored if the cloud provider does not support the feature.This field can be used only with 'loadbalancer' type listener.",
														MarkdownDescription: "The loadbalancer is requested with the IP address specified in this field. This feature depends on whether the underlying cloud provider supports specifying the 'loadBalancerIP' when a load balancer is created. This field is ignored if the cloud provider does not support the feature.This field can be used only with 'loadbalancer' type listener.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_port": {
														Description:         "Node port for the bootstrap service. This field can be used only with 'nodeport' type listener.",
														MarkdownDescription: "Node port for the bootstrap service. This field can be used only with 'nodeport' type listener.",

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

											"broker_cert_chain_and_key": {
												Description:         "Reference to the 'Secret' which holds the certificate and private key pair which will be used for this listener. The certificate can optionally contain the whole chain. This field can be used only with listeners with enabled TLS encryption.",
												MarkdownDescription: "Reference to the 'Secret' which holds the certificate and private key pair which will be used for this listener. The certificate can optionally contain the whole chain. This field can be used only with listeners with enabled TLS encryption.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"certificate": {
														Description:         "The name of the file certificate in the Secret.",
														MarkdownDescription: "The name of the file certificate in the Secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"key": {
														Description:         "The name of the private key in the Secret.",
														MarkdownDescription: "The name of the private key in the Secret.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"secret_name": {
														Description:         "The name of the Secret containing the certificate.",
														MarkdownDescription: "The name of the Secret containing the certificate.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"brokers": {
												Description:         "Per-broker configurations.",
												MarkdownDescription: "Per-broker configurations.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"advertised_host": {
														Description:         "The host name which will be used in the brokers' 'advertised.brokers'.",
														MarkdownDescription: "The host name which will be used in the brokers' 'advertised.brokers'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"advertised_port": {
														Description:         "The port number which will be used in the brokers' 'advertised.brokers'.",
														MarkdownDescription: "The port number which will be used in the brokers' 'advertised.brokers'.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"annotations": {
														Description:         "Annotations that will be added to the 'Ingress' or 'Service' resource. You can use this field to configure DNS providers such as External DNS. This field can be used only with 'loadbalancer', 'nodeport', or 'ingress' type listeners.",
														MarkdownDescription: "Annotations that will be added to the 'Ingress' or 'Service' resource. You can use this field to configure DNS providers such as External DNS. This field can be used only with 'loadbalancer', 'nodeport', or 'ingress' type listeners.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"broker": {
														Description:         "ID of the kafka broker (broker identifier). Broker IDs start from 0 and correspond to the number of broker replicas.",
														MarkdownDescription: "ID of the kafka broker (broker identifier). Broker IDs start from 0 and correspond to the number of broker replicas.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"host": {
														Description:         "The broker host. This field will be used in the Ingress resource or in the Route resource to specify the desired hostname. This field can be used only with 'route' (optional) or 'ingress' (required) type listeners.",
														MarkdownDescription: "The broker host. This field will be used in the Ingress resource or in the Route resource to specify the desired hostname. This field can be used only with 'route' (optional) or 'ingress' (required) type listeners.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels that will be added to the 'Ingress', 'Route', or 'Service' resource. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",
														MarkdownDescription: "Labels that will be added to the 'Ingress', 'Route', or 'Service' resource. This field can be used only with 'loadbalancer', 'nodeport', 'route', or 'ingress' type listeners.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"load_balancer_ip": {
														Description:         "The loadbalancer is requested with the IP address specified in this field. This feature depends on whether the underlying cloud provider supports specifying the 'loadBalancerIP' when a load balancer is created. This field is ignored if the cloud provider does not support the feature.This field can be used only with 'loadbalancer' type listener.",
														MarkdownDescription: "The loadbalancer is requested with the IP address specified in this field. This feature depends on whether the underlying cloud provider supports specifying the 'loadBalancerIP' when a load balancer is created. This field is ignored if the cloud provider does not support the feature.This field can be used only with 'loadbalancer' type listener.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_port": {
														Description:         "Node port for the per-broker service. This field can be used only with 'nodeport' type listener.",
														MarkdownDescription: "Node port for the per-broker service. This field can be used only with 'nodeport' type listener.",

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

											"class": {
												Description:         "Configures the 'Ingress' class that defines which 'Ingress' controller will be used. This field can be used only with 'ingress' type listener. If not specified, the default Ingress controller will be used.",
												MarkdownDescription: "Configures the 'Ingress' class that defines which 'Ingress' controller will be used. This field can be used only with 'ingress' type listener. If not specified, the default Ingress controller will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"create_bootstrap_service": {
												Description:         "Whether to create the bootstrap service or not. The bootstrap service is created by default (if not specified differently). This field can be used with the 'loadBalancer' type listener.",
												MarkdownDescription: "Whether to create the bootstrap service or not. The bootstrap service is created by default (if not specified differently). This field can be used with the 'loadBalancer' type listener.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"external_traffic_policy": {
												Description:         "Specifies whether the service routes external traffic to node-local or cluster-wide endpoints. 'Cluster' may cause a second hop to another node and obscures the client source IP. 'Local' avoids a second hop for LoadBalancer and Nodeport type services and preserves the client source IP (when supported by the infrastructure). If unspecified, Kubernetes will use 'Cluster' as the default.This field can be used only with 'loadbalancer' or 'nodeport' type listener.",
												MarkdownDescription: "Specifies whether the service routes external traffic to node-local or cluster-wide endpoints. 'Cluster' may cause a second hop to another node and obscures the client source IP. 'Local' avoids a second hop for LoadBalancer and Nodeport type services and preserves the client source IP (when supported by the infrastructure). If unspecified, Kubernetes will use 'Cluster' as the default.This field can be used only with 'loadbalancer' or 'nodeport' type listener.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("Local", "Cluster"),
												},
											},

											"finalizers": {
												Description:         "A list of finalizers which will be configured for the 'LoadBalancer' type Services created for this listener. If supported by the platform, the finalizer 'service.kubernetes.io/load-balancer-cleanup' to make sure that the external load balancer is deleted together with the service.For more information, see https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#garbage-collecting-load-balancers. This field can be used only with 'loadbalancer' type listeners.",
												MarkdownDescription: "A list of finalizers which will be configured for the 'LoadBalancer' type Services created for this listener. If supported by the platform, the finalizer 'service.kubernetes.io/load-balancer-cleanup' to make sure that the external load balancer is deleted together with the service.For more information, see https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/#garbage-collecting-load-balancers. This field can be used only with 'loadbalancer' type listeners.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_families": {
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_family_policy": {
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
											},

											"load_balancer_source_ranges": {
												Description:         "A list of CIDR ranges (for example '10.0.0.0/8' or '130.211.204.1/32') from which clients can connect to load balancer type listeners. If supported by the platform, traffic through the loadbalancer is restricted to the specified CIDR ranges. This field is applicable only for loadbalancer type services and is ignored if the cloud provider does not support the feature. This field can be used only with 'loadbalancer' type listener.",
												MarkdownDescription: "A list of CIDR ranges (for example '10.0.0.0/8' or '130.211.204.1/32') from which clients can connect to load balancer type listeners. If supported by the platform, traffic through the loadbalancer is restricted to the specified CIDR ranges. This field is applicable only for loadbalancer type services and is ignored if the cloud provider does not support the feature. This field can be used only with 'loadbalancer' type listener.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_connection_creation_rate": {
												Description:         "The maximum connection creation rate we allow in this listener at any time. New connections will be throttled if the limit is reached.",
												MarkdownDescription: "The maximum connection creation rate we allow in this listener at any time. New connections will be throttled if the limit is reached.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_connections": {
												Description:         "The maximum number of connections we allow for this listener in the broker at any time. New connections are blocked if the limit is reached.",
												MarkdownDescription: "The maximum number of connections we allow for this listener in the broker at any time. New connections are blocked if the limit is reached.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"preferred_node_port_address_type": {
												Description:         "Defines which address type should be used as the node address. Available types are: 'ExternalDNS', 'ExternalIP', 'InternalDNS', 'InternalIP' and 'Hostname'. By default, the addresses will be used in the following order (the first one found will be used):* 'ExternalDNS'* 'ExternalIP'* 'InternalDNS'* 'InternalIP'* 'Hostname'This field is used to select the preferred address type, which is checked first. If no address is found for this address type, the other types are checked in the default order. This field can only be used with 'nodeport' type listener.",
												MarkdownDescription: "Defines which address type should be used as the node address. Available types are: 'ExternalDNS', 'ExternalIP', 'InternalDNS', 'InternalIP' and 'Hostname'. By default, the addresses will be used in the following order (the first one found will be used):* 'ExternalDNS'* 'ExternalIP'* 'InternalDNS'* 'InternalIP'* 'Hostname'This field is used to select the preferred address type, which is checked first. If no address is found for this address type, the other types are checked in the default order. This field can only be used with 'nodeport' type listener.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ExternalIP", "ExternalDNS", "InternalIP", "InternalDNS", "Hostname"),
												},
											},

											"use_service_dns_domain": {
												Description:         "Configures whether the Kubernetes service DNS domain should be used or not. If set to 'true', the generated addresses will contain the service DNS domain suffix (by default '.cluster.local', can be configured using environment variable 'KUBERNETES_SERVICE_DNS_DOMAIN'). Defaults to 'false'.This field can be used only with 'internal' type listener.",
												MarkdownDescription: "Configures whether the Kubernetes service DNS domain should be used or not. If set to 'true', the generated addresses will contain the service DNS domain suffix (by default '.cluster.local', can be configured using environment variable 'KUBERNETES_SERVICE_DNS_DOMAIN'). Defaults to 'false'.This field can be used only with 'internal' type listener.",

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

									"name": {
										Description:         "Name of the listener. The name will be used to identify the listener and the related Kubernetes objects. The name has to be unique within given a Kafka cluster. The name can consist of lowercase characters and numbers and be up to 11 characters long.",
										MarkdownDescription: "Name of the listener. The name will be used to identify the listener and the related Kubernetes objects. The name has to be unique within given a Kafka cluster. The name can consist of lowercase characters and numbers and be up to 11 characters long.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]{1,11}$`), ""),
										},
									},

									"network_policy_peers": {
										Description:         "List of peers which should be able to connect to this listener. Peers in this list are combined using a logical OR operation. If this field is empty or missing, all connections will be allowed for this listener. If this field is present and contains at least one item, the listener only allows the traffic which matches at least one item in this list.",
										MarkdownDescription: "List of peers which should be able to connect to this listener. Peers in this list are combined using a logical OR operation. If this field is empty or missing, all connections will be allowed for this listener. If this field is present and contains at least one item, the listener only allows the traffic which matches at least one item in this list.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"ip_block": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cidr": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"except": {
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

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_selector": {
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

														Type: utilities.DynamicType{},

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

									"port": {
										Description:         "Port number used by the listener inside Kafka. The port number has to be unique within a given Kafka cluster. Allowed port numbers are 9092 and higher with the exception of ports 9404 and 9999, which are already used for Prometheus and JMX. Depending on the listener type, the port number might not be the same as the port number that connects Kafka clients.",
										MarkdownDescription: "Port number used by the listener inside Kafka. The port number has to be unique within a given Kafka cluster. Allowed port numbers are 9092 and higher with the exception of ports 9404 and 9999, which are already used for Prometheus and JMX. Depending on the listener type, the port number might not be the same as the port number that connects Kafka clients.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(9092),
										},
									},

									"tls": {
										Description:         "Enables TLS encryption on the listener. This is a required property.",
										MarkdownDescription: "Enables TLS encryption on the listener. This is a required property.",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"type": {
										Description:         "Type of the listener. Currently the supported types are 'internal', 'route', 'loadbalancer', 'nodeport' and 'ingress'. * 'internal' type exposes Kafka internally only within the Kubernetes cluster.* 'route' type uses OpenShift Routes to expose Kafka.* 'loadbalancer' type uses LoadBalancer type services to expose Kafka.* 'nodeport' type uses NodePort type services to expose Kafka.* 'ingress' type uses Kubernetes Nginx Ingress to expose Kafka with TLS passthrought.* 'cluster-ip' type uses ClusterIP service with per broker port number. Can be exposed over Nginx Ingress Controller with tcp port config.",
										MarkdownDescription: "Type of the listener. Currently the supported types are 'internal', 'route', 'loadbalancer', 'nodeport' and 'ingress'. * 'internal' type exposes Kafka internally only within the Kubernetes cluster.* 'route' type uses OpenShift Routes to expose Kafka.* 'loadbalancer' type uses LoadBalancer type services to expose Kafka.* 'nodeport' type uses NodePort type services to expose Kafka.* 'ingress' type uses Kubernetes Nginx Ingress to expose Kafka with TLS passthrought.* 'cluster-ip' type uses ClusterIP service with per broker port number. Can be exposed over Nginx Ingress Controller with tcp port config.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("internal", "route", "loadbalancer", "nodeport", "ingress", "cluster-ip"),
										},
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"liveness_probe": {
								Description:         "Pod liveness checking.",
								MarkdownDescription: "Pod liveness checking.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": {
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": {
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": {
								Description:         "Logging configuration for Kafka.",
								MarkdownDescription: "Logging configuration for Kafka.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"loggers": {
										Description:         "A Map from logger name to logger level.",
										MarkdownDescription: "A Map from logger name to logger level.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Logging type, must be either 'inline' or 'external'.",
										MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("inline", "external"),
										},
									},

									"value_from": {
										Description:         "'ConfigMap' entry where the logging configuration is stored. ",
										MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
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

							"metrics_config": {
								Description:         "Metrics configuration.",
								MarkdownDescription: "Metrics configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"type": {
										Description:         "Metrics type. Only 'jmxPrometheusExporter' supported currently.",
										MarkdownDescription: "Metrics type. Only 'jmxPrometheusExporter' supported currently.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("jmxPrometheusExporter"),
										},
									},

									"value_from": {
										Description:         "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",
										MarkdownDescription: "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rack": {
								Description:         "Configuration of the 'broker.rack' broker config.",
								MarkdownDescription: "Configuration of the 'broker.rack' broker config.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"topology_key": {
										Description:         "A key that matches labels assigned to the Kubernetes cluster nodes. The value of the label is used to set a broker's 'broker.rack' config, and the 'client.rack' config for Kafka Connect or MirrorMaker 2.0.",
										MarkdownDescription: "A key that matches labels assigned to the Kubernetes cluster nodes. The value of the label is used to set a broker's 'broker.rack' config, and the 'client.rack' config for Kafka Connect or MirrorMaker 2.0.",

										Type: types.StringType,

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
								Description:         "Pod readiness checking.",
								MarkdownDescription: "Pod readiness checking.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": {
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": {
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": {
								Description:         "The number of pods in the cluster.",
								MarkdownDescription: "The number of pods in the cluster.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"resources": {
								Description:         "CPU and memory resources to reserve.",
								MarkdownDescription: "CPU and memory resources to reserve.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requests": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage": {
								Description:         "Storage configuration (disk). Cannot be updated.",
								MarkdownDescription: "Storage configuration (disk). Cannot be updated.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"class": {
										Description:         "The storage class to use for dynamic volume allocation.",
										MarkdownDescription: "The storage class to use for dynamic volume allocation.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"delete_claim": {
										Description:         "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",
										MarkdownDescription: "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",
										MarkdownDescription: "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"overrides": {
										Description:         "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",
										MarkdownDescription: "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"broker": {
												Description:         "Id of the kafka broker (broker identifier).",
												MarkdownDescription: "Id of the kafka broker (broker identifier).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"class": {
												Description:         "The storage class to use for dynamic volume allocation for this broker.",
												MarkdownDescription: "The storage class to use for dynamic volume allocation for this broker.",

												Type: types.StringType,

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
										Description:         "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",
										MarkdownDescription: "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"size": {
										Description:         "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",
										MarkdownDescription: "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"size_limit": {
										Description:         "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",
										MarkdownDescription: "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
										},
									},

									"type": {
										Description:         "Storage type, must be either 'ephemeral', 'persistent-claim', or 'jbod'.",
										MarkdownDescription: "Storage type, must be either 'ephemeral', 'persistent-claim', or 'jbod'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("ephemeral", "persistent-claim", "jbod"),
										},
									},

									"volumes": {
										Description:         "List of volumes as Storage objects representing the JBOD disks array.",
										MarkdownDescription: "List of volumes as Storage objects representing the JBOD disks array.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"class": {
												Description:         "The storage class to use for dynamic volume allocation.",
												MarkdownDescription: "The storage class to use for dynamic volume allocation.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"delete_claim": {
												Description:         "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",
												MarkdownDescription: "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"id": {
												Description:         "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",
												MarkdownDescription: "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"overrides": {
												Description:         "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",
												MarkdownDescription: "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"broker": {
														Description:         "Id of the kafka broker (broker identifier).",
														MarkdownDescription: "Id of the kafka broker (broker identifier).",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"class": {
														Description:         "The storage class to use for dynamic volume allocation for this broker.",
														MarkdownDescription: "The storage class to use for dynamic volume allocation for this broker.",

														Type: types.StringType,

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
												Description:         "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",
												MarkdownDescription: "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",

												Type: utilities.DynamicType{},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size": {
												Description:         "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",
												MarkdownDescription: "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size_limit": {
												Description:         "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",
												MarkdownDescription: "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
												},
											},

											"type": {
												Description:         "Storage type, must be either 'ephemeral' or 'persistent-claim'.",
												MarkdownDescription: "Storage type, must be either 'ephemeral' or 'persistent-claim'.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("ephemeral", "persistent-claim"),
												},
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

							"template": {
								Description:         "Template for Kafka cluster resources. The template allows users to specify how the 'StatefulSet', 'Pods', and 'Services' are generated.",
								MarkdownDescription: "Template for Kafka cluster resources. The template allows users to specify how the 'StatefulSet', 'Pods', and 'Services' are generated.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"bootstrap_service": {
										Description:         "Template for Kafka bootstrap 'Service'.",
										MarkdownDescription: "Template for Kafka bootstrap 'Service'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ip_families": {
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_family_policy": {
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
											},

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"brokers_service": {
										Description:         "Template for Kafka broker 'Service'.",
										MarkdownDescription: "Template for Kafka broker 'Service'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ip_families": {
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_family_policy": {
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
											},

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"cluster_ca_cert": {
										Description:         "Template for Secret with Kafka Cluster certificate public key.",
										MarkdownDescription: "Template for Secret with Kafka Cluster certificate public key.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"cluster_role_binding": {
										Description:         "Template for the Kafka ClusterRoleBinding.",
										MarkdownDescription: "Template for the Kafka ClusterRoleBinding.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"external_bootstrap_ingress": {
										Description:         "Template for Kafka external bootstrap 'Ingress'.",
										MarkdownDescription: "Template for Kafka external bootstrap 'Ingress'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"external_bootstrap_route": {
										Description:         "Template for Kafka external bootstrap 'Route'.",
										MarkdownDescription: "Template for Kafka external bootstrap 'Route'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"external_bootstrap_service": {
										Description:         "Template for Kafka external bootstrap 'Service'.",
										MarkdownDescription: "Template for Kafka external bootstrap 'Service'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"init_container": {
										Description:         "Template for the Kafka init container.",
										MarkdownDescription: "Template for the Kafka init container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"jmx_secret": {
										Description:         "Template for Secret of the Kafka Cluster JMX authentication.",
										MarkdownDescription: "Template for Secret of the Kafka Cluster JMX authentication.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"kafka_container": {
										Description:         "Template for the Kafka broker container.",
										MarkdownDescription: "Template for the Kafka broker container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"per_pod_ingress": {
										Description:         "Template for Kafka per-pod 'Ingress' used for access from outside of Kubernetes.",
										MarkdownDescription: "Template for Kafka per-pod 'Ingress' used for access from outside of Kubernetes.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"per_pod_route": {
										Description:         "Template for Kafka per-pod 'Routes' used for access from outside of OpenShift.",
										MarkdownDescription: "Template for Kafka per-pod 'Routes' used for access from outside of OpenShift.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"per_pod_service": {
										Description:         "Template for Kafka per-pod 'Services' used for access from outside of Kubernetes.",
										MarkdownDescription: "Template for Kafka per-pod 'Services' used for access from outside of Kubernetes.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"persistent_volume_claim": {
										Description:         "Template for all Kafka 'PersistentVolumeClaims'.",
										MarkdownDescription: "Template for all Kafka 'PersistentVolumeClaims'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"pod": {
										Description:         "Template for Kafka 'Pods'.",
										MarkdownDescription: "Template for Kafka 'Pods'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"affinity": {
												Description:         "The pod's affinity rules.",
												MarkdownDescription: "The pod's affinity rules.",

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

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

											"enable_service_links": {
												Description:         "Indicates whether information about services should be injected into Pod's environment variables.",
												MarkdownDescription: "Indicates whether information about services should be injected into Pod's environment variables.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_aliases": {
												Description:         "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",
												MarkdownDescription: "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",

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

											"image_pull_secrets": {
												Description:         "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
												MarkdownDescription: "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",

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

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scheduler_name": {
												Description:         "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
												MarkdownDescription: "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": {
												Description:         "Configures pod-level security attributes and common container settings.",
												MarkdownDescription: "Configures pod-level security attributes and common container settings.",

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

											"termination_grace_period_seconds": {
												Description:         "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
												MarkdownDescription: "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"tmp_dir_size_limit": {
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
												},
											},

											"tolerations": {
												Description:         "The pod's tolerations.",
												MarkdownDescription: "The pod's tolerations.",

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
												Description:         "The pod's topology spread constraints.",
												MarkdownDescription: "The pod's topology spread constraints.",

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

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_label_keys": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_skew": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"min_domains": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_affinity_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_taints_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"when_unsatisfiable": {
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

									"pod_disruption_budget": {
										Description:         "Template for Kafka 'PodDisruptionBudget'.",
										MarkdownDescription: "Template for Kafka 'PodDisruptionBudget'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max_unavailable": {
												Description:         "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
												MarkdownDescription: "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"metadata": {
												Description:         "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
												MarkdownDescription: "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"pod_set": {
										Description:         "Template for Kafka 'StrimziPodSet' resource.",
										MarkdownDescription: "Template for Kafka 'StrimziPodSet' resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
										Description:         "Template for the Kafka service account.",
										MarkdownDescription: "Template for the Kafka service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"statefulset": {
										Description:         "Template for Kafka 'StatefulSet'.",
										MarkdownDescription: "Template for Kafka 'StatefulSet'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_management_policy": {
												Description:         "PodManagementPolicy which will be used for this StatefulSet. Valid values are 'Parallel' and 'OrderedReady'. Defaults to 'Parallel'.",
												MarkdownDescription: "PodManagementPolicy which will be used for this StatefulSet. Valid values are 'Parallel' and 'OrderedReady'. Defaults to 'Parallel'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("OrderedReady", "Parallel"),
												},
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

							"version": {
								Description:         "The kafka broker version. Defaults to {DefaultKafkaVersion}. Consult the user documentation to understand the process required to upgrade or downgrade the version.",
								MarkdownDescription: "The kafka broker version. Defaults to {DefaultKafkaVersion}. Consult the user documentation to understand the process required to upgrade or downgrade the version.",

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

					"kafka_exporter": {
						Description:         "Configuration of the Kafka Exporter. Kafka Exporter can provide additional metrics, for example lag of consumer group at topic/partition.",
						MarkdownDescription: "Configuration of the Kafka Exporter. Kafka Exporter can provide additional metrics, for example lag of consumer group at topic/partition.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enable_sarama_logging": {
								Description:         "Enable Sarama logging, a Go client library used by the Kafka Exporter.",
								MarkdownDescription: "Enable Sarama logging, a Go client library used by the Kafka Exporter.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"group_regex": {
								Description:         "Regular expression to specify which consumer groups to collect. Default value is '.*'.",
								MarkdownDescription: "Regular expression to specify which consumer groups to collect. Default value is '.*'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "The docker image for the pods.",
								MarkdownDescription: "The docker image for the pods.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"liveness_probe": {
								Description:         "Pod liveness check.",
								MarkdownDescription: "Pod liveness check.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": {
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": {
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": {
								Description:         "Only log messages with the given severity or above. Valid levels: ['info', 'debug', 'trace']. Default log level is 'info'.",
								MarkdownDescription: "Only log messages with the given severity or above. Valid levels: ['info', 'debug', 'trace']. Default log level is 'info'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"readiness_probe": {
								Description:         "Pod readiness check.",
								MarkdownDescription: "Pod readiness check.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": {
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": {
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "CPU and memory resources to reserve.",
								MarkdownDescription: "CPU and memory resources to reserve.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requests": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

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
								Description:         "Customization of deployment templates and pods.",
								MarkdownDescription: "Customization of deployment templates and pods.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"container": {
										Description:         "Template for the Kafka Exporter container.",
										MarkdownDescription: "Template for the Kafka Exporter container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"deployment": {
										Description:         "Template for Kafka Exporter 'Deployment'.",
										MarkdownDescription: "Template for Kafka Exporter 'Deployment'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"pod": {
										Description:         "Template for Kafka Exporter 'Pods'.",
										MarkdownDescription: "Template for Kafka Exporter 'Pods'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"affinity": {
												Description:         "The pod's affinity rules.",
												MarkdownDescription: "The pod's affinity rules.",

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

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

											"enable_service_links": {
												Description:         "Indicates whether information about services should be injected into Pod's environment variables.",
												MarkdownDescription: "Indicates whether information about services should be injected into Pod's environment variables.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_aliases": {
												Description:         "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",
												MarkdownDescription: "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",

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

											"image_pull_secrets": {
												Description:         "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
												MarkdownDescription: "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",

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

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scheduler_name": {
												Description:         "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
												MarkdownDescription: "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": {
												Description:         "Configures pod-level security attributes and common container settings.",
												MarkdownDescription: "Configures pod-level security attributes and common container settings.",

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

											"termination_grace_period_seconds": {
												Description:         "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
												MarkdownDescription: "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"tmp_dir_size_limit": {
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
												},
											},

											"tolerations": {
												Description:         "The pod's tolerations.",
												MarkdownDescription: "The pod's tolerations.",

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
												Description:         "The pod's topology spread constraints.",
												MarkdownDescription: "The pod's topology spread constraints.",

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

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_label_keys": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_skew": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"min_domains": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_affinity_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_taints_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"when_unsatisfiable": {
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

									"service": {
										Description:         "Template for Kafka Exporter 'Service'.",
										MarkdownDescription: "Template for Kafka Exporter 'Service'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
										Description:         "Template for the Kafka Exporter service account.",
										MarkdownDescription: "Template for the Kafka Exporter service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

							"topic_regex": {
								Description:         "Regular expression to specify which topics to collect. Default value is '.*'.",
								MarkdownDescription: "Regular expression to specify which topics to collect. Default value is '.*'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"maintenance_time_windows": {
						Description:         "A list of time windows for maintenance tasks (that is, certificates renewal). Each time window is defined by a cron expression.",
						MarkdownDescription: "A list of time windows for maintenance tasks (that is, certificates renewal). Each time window is defined by a cron expression.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"zookeeper": {
						Description:         "Configuration of the ZooKeeper cluster.",
						MarkdownDescription: "Configuration of the ZooKeeper cluster.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"config": {
								Description:         "The ZooKeeper broker config. Properties with the following prefixes cannot be set: server., dataDir, dataLogDir, clientPort, authProvider, quorum.auth, requireClientAuthScheme, snapshot.trust.empty, standaloneEnabled, reconfigEnabled, 4lw.commands.whitelist, secureClientPort, ssl., serverCnxnFactory, sslQuorum (with the exception of: ssl.protocol, ssl.quorum.protocol, ssl.enabledProtocols, ssl.quorum.enabledProtocols, ssl.ciphersuites, ssl.quorum.ciphersuites, ssl.hostnameVerification, ssl.quorum.hostnameVerification).",
								MarkdownDescription: "The ZooKeeper broker config. Properties with the following prefixes cannot be set: server., dataDir, dataLogDir, clientPort, authProvider, quorum.auth, requireClientAuthScheme, snapshot.trust.empty, standaloneEnabled, reconfigEnabled, 4lw.commands.whitelist, secureClientPort, ssl., serverCnxnFactory, sslQuorum (with the exception of: ssl.protocol, ssl.quorum.protocol, ssl.enabledProtocols, ssl.quorum.enabledProtocols, ssl.ciphersuites, ssl.quorum.ciphersuites, ssl.hostnameVerification, ssl.quorum.hostnameVerification).",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "The docker image for the pods.",
								MarkdownDescription: "The docker image for the pods.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"jmx_options": {
								Description:         "JMX Options for Zookeeper nodes.",
								MarkdownDescription: "JMX Options for Zookeeper nodes.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"authentication": {
										Description:         "Authentication configuration for connecting to the JMX port.",
										MarkdownDescription: "Authentication configuration for connecting to the JMX port.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"type": {
												Description:         "Authentication type. Currently the only supported types are 'password'.'password' type creates a username and protected port with no TLS.",
												MarkdownDescription: "Authentication type. Currently the only supported types are 'password'.'password' type creates a username and protected port with no TLS.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("password"),
												},
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

							"jvm_options": {
								Description:         "JVM Options for pods.",
								MarkdownDescription: "JVM Options for pods.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"xx": {
										Description:         "A map of -XX options to the JVM.",
										MarkdownDescription: "A map of -XX options to the JVM.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"xms": {
										Description:         "-Xms option to to the JVM.",
										MarkdownDescription: "-Xms option to to the JVM.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
										},
									},

									"xmx": {
										Description:         "-Xmx option to to the JVM.",
										MarkdownDescription: "-Xmx option to to the JVM.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+[mMgG]?$`), ""),
										},
									},

									"gc_logging_enabled": {
										Description:         "Specifies whether the Garbage Collection logging is enabled. The default is false.",
										MarkdownDescription: "Specifies whether the Garbage Collection logging is enabled. The default is false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"java_system_properties": {
										Description:         "A map of additional system properties which will be passed using the '-D' option to the JVM.",
										MarkdownDescription: "A map of additional system properties which will be passed using the '-D' option to the JVM.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"name": {
												Description:         "The system property name.",
												MarkdownDescription: "The system property name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "The system property value.",
												MarkdownDescription: "The system property value.",

												Type: types.StringType,

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
								Description:         "Pod liveness checking.",
								MarkdownDescription: "Pod liveness checking.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": {
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": {
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logging": {
								Description:         "Logging configuration for ZooKeeper.",
								MarkdownDescription: "Logging configuration for ZooKeeper.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"loggers": {
										Description:         "A Map from logger name to logger level.",
										MarkdownDescription: "A Map from logger name to logger level.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Logging type, must be either 'inline' or 'external'.",
										MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("inline", "external"),
										},
									},

									"value_from": {
										Description:         "'ConfigMap' entry where the logging configuration is stored. ",
										MarkdownDescription: "'ConfigMap' entry where the logging configuration is stored. ",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
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

							"metrics_config": {
								Description:         "Metrics configuration.",
								MarkdownDescription: "Metrics configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"type": {
										Description:         "Metrics type. Only 'jmxPrometheusExporter' supported currently.",
										MarkdownDescription: "Metrics type. Only 'jmxPrometheusExporter' supported currently.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("jmxPrometheusExporter"),
										},
									},

									"value_from": {
										Description:         "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",
										MarkdownDescription: "ConfigMap entry where the Prometheus JMX Exporter configuration is stored. For details of the structure of this configuration, see the {JMXExporter}.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Reference to the key in the ConfigMap containing the configuration.",
												MarkdownDescription: "Reference to the key in the ConfigMap containing the configuration.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"key": {
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"readiness_probe": {
								Description:         "Pod readiness checking.",
								MarkdownDescription: "Pod readiness checking.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"initial_delay_seconds": {
										Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
										MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},

									"timeout_seconds": {
										Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
										MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": {
								Description:         "The number of pods in the cluster.",
								MarkdownDescription: "The number of pods in the cluster.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"resources": {
								Description:         "CPU and memory resources to reserve.",
								MarkdownDescription: "CPU and memory resources to reserve.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"limits": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"requests": {
										Description:         "",
										MarkdownDescription: "",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage": {
								Description:         "Storage configuration (disk). Cannot be updated.",
								MarkdownDescription: "Storage configuration (disk). Cannot be updated.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"class": {
										Description:         "The storage class to use for dynamic volume allocation.",
										MarkdownDescription: "The storage class to use for dynamic volume allocation.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"delete_claim": {
										Description:         "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",
										MarkdownDescription: "Specifies if the persistent volume claim has to be deleted when the cluster is un-deployed.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"id": {
										Description:         "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",
										MarkdownDescription: "Storage identification number. It is mandatory only for storage volumes defined in a storage of type 'jbod'.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(0),
										},
									},

									"overrides": {
										Description:         "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",
										MarkdownDescription: "Overrides for individual brokers. The 'overrides' field allows to specify a different configuration for different brokers.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"broker": {
												Description:         "Id of the kafka broker (broker identifier).",
												MarkdownDescription: "Id of the kafka broker (broker identifier).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"class": {
												Description:         "The storage class to use for dynamic volume allocation for this broker.",
												MarkdownDescription: "The storage class to use for dynamic volume allocation for this broker.",

												Type: types.StringType,

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
										Description:         "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",
										MarkdownDescription: "Specifies a specific persistent volume to use. It contains key:value pairs representing labels for selecting such a volume.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"size": {
										Description:         "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",
										MarkdownDescription: "When type=persistent-claim, defines the size of the persistent volume claim (i.e 1Gi). Mandatory when type=persistent-claim.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"size_limit": {
										Description:         "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",
										MarkdownDescription: "When type=ephemeral, defines the total amount of local storage required for this EmptyDir volume (for example 1Gi).",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
										},
									},

									"type": {
										Description:         "Storage type, must be either 'ephemeral' or 'persistent-claim'.",
										MarkdownDescription: "Storage type, must be either 'ephemeral' or 'persistent-claim'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("ephemeral", "persistent-claim"),
										},
									},
								}),

								Required: true,
								Optional: false,
								Computed: false,
							},

							"template": {
								Description:         "Template for ZooKeeper cluster resources. The template allows users to specify how the 'StatefulSet', 'Pods', and 'Services' are generated.",
								MarkdownDescription: "Template for ZooKeeper cluster resources. The template allows users to specify how the 'StatefulSet', 'Pods', and 'Services' are generated.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client_service": {
										Description:         "Template for ZooKeeper client 'Service'.",
										MarkdownDescription: "Template for ZooKeeper client 'Service'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ip_families": {
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_family_policy": {
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
											},

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"jmx_secret": {
										Description:         "Template for Secret of the Zookeeper Cluster JMX authentication.",
										MarkdownDescription: "Template for Secret of the Zookeeper Cluster JMX authentication.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"nodes_service": {
										Description:         "Template for ZooKeeper nodes 'Service'.",
										MarkdownDescription: "Template for ZooKeeper nodes 'Service'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ip_families": {
												Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting. Available on Kubernetes 1.20 and newer.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ip_family_policy": {
												Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",
												MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type. Available on Kubernetes 1.20 and newer.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("SingleStack", "PreferDualStack", "RequireDualStack"),
												},
											},

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"persistent_volume_claim": {
										Description:         "Template for all ZooKeeper 'PersistentVolumeClaims'.",
										MarkdownDescription: "Template for all ZooKeeper 'PersistentVolumeClaims'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"pod": {
										Description:         "Template for ZooKeeper 'Pods'.",
										MarkdownDescription: "Template for ZooKeeper 'Pods'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"affinity": {
												Description:         "The pod's affinity rules.",
												MarkdownDescription: "The pod's affinity rules.",

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

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																						Type: utilities.DynamicType{},

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

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},
																		}),

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"weight": {
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

																				Type: utilities.DynamicType{},

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

																				Type: utilities.DynamicType{},

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

											"enable_service_links": {
												Description:         "Indicates whether information about services should be injected into Pod's environment variables.",
												MarkdownDescription: "Indicates whether information about services should be injected into Pod's environment variables.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"host_aliases": {
												Description:         "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",
												MarkdownDescription: "The pod's HostAliases. HostAliases is an optional list of hosts and IPs that will be injected into the Pod's hosts file if specified.",

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

											"image_pull_secrets": {
												Description:         "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
												MarkdownDescription: "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",

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

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
												Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
												MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scheduler_name": {
												Description:         "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
												MarkdownDescription: "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"security_context": {
												Description:         "Configures pod-level security attributes and common container settings.",
												MarkdownDescription: "Configures pod-level security attributes and common container settings.",

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

											"termination_grace_period_seconds": {
												Description:         "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
												MarkdownDescription: "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"tmp_dir_size_limit": {
												Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
												MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9.]+)([eEinumkKMGTP]*[-+]?[0-9]*)$`), ""),
												},
											},

											"tolerations": {
												Description:         "The pod's tolerations.",
												MarkdownDescription: "The pod's tolerations.",

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
												Description:         "The pod's topology spread constraints.",
												MarkdownDescription: "The pod's topology spread constraints.",

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

																Type: utilities.DynamicType{},

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"match_label_keys": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"max_skew": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"min_domains": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_affinity_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_taints_policy": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"when_unsatisfiable": {
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

									"pod_disruption_budget": {
										Description:         "Template for ZooKeeper 'PodDisruptionBudget'.",
										MarkdownDescription: "Template for ZooKeeper 'PodDisruptionBudget'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max_unavailable": {
												Description:         "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
												MarkdownDescription: "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),
												},
											},

											"metadata": {
												Description:         "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
												MarkdownDescription: "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"pod_set": {
										Description:         "Template for ZooKeeper 'StrimziPodSet' resource.",
										MarkdownDescription: "Template for ZooKeeper 'StrimziPodSet' resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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
										Description:         "Template for the ZooKeeper service account.",
										MarkdownDescription: "Template for the ZooKeeper service account.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

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

									"statefulset": {
										Description:         "Template for ZooKeeper 'StatefulSet'.",
										MarkdownDescription: "Template for ZooKeeper 'StatefulSet'.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "Metadata applied to the resource.",
												MarkdownDescription: "Metadata applied to the resource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
														Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"labels": {
														Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
														MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

														Type: utilities.DynamicType{},

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pod_management_policy": {
												Description:         "PodManagementPolicy which will be used for this StatefulSet. Valid values are 'Parallel' and 'OrderedReady'. Defaults to 'Parallel'.",
												MarkdownDescription: "PodManagementPolicy which will be used for this StatefulSet. Valid values are 'Parallel' and 'OrderedReady'. Defaults to 'Parallel'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.OneOf("OrderedReady", "Parallel"),
												},
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"zookeeper_container": {
										Description:         "Template for the ZooKeeper container.",
										MarkdownDescription: "Template for the ZooKeeper container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"env": {
												Description:         "Environment variables which should be applied to the container.",
												MarkdownDescription: "Environment variables which should be applied to the container.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The environment variable key.",
														MarkdownDescription: "The environment variable key.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "The environment variable value.",
														MarkdownDescription: "The environment variable value.",

														Type: types.StringType,

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
												Description:         "Security context for the container.",
												MarkdownDescription: "Security context for the container.",

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
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *KafkaStrimziIoKafkaV1Beta2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kafka_strimzi_io_kafka_v1beta2")

	var state KafkaStrimziIoKafkaV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KafkaStrimziIoKafkaV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kafka.strimzi.io/v1beta2")
	goModel.Kind = utilities.Ptr("Kafka")

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

func (r *KafkaStrimziIoKafkaV1Beta2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_strimzi_io_kafka_v1beta2")
	// NO-OP: All data is already in Terraform state
}

func (r *KafkaStrimziIoKafkaV1Beta2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kafka_strimzi_io_kafka_v1beta2")

	var state KafkaStrimziIoKafkaV1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KafkaStrimziIoKafkaV1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kafka.strimzi.io/v1beta2")
	goModel.Kind = utilities.Ptr("Kafka")

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

func (r *KafkaStrimziIoKafkaV1Beta2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kafka_strimzi_io_kafka_v1beta2")
	// NO-OP: Terraform removes the state automatically for us
}
