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

type KafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource struct{}

var (
	_ resource.Resource = (*KafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource)(nil)
)

type KafkaStrimziIoKafkaMirrorMaker2V1Beta2TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KafkaStrimziIoKafkaMirrorMaker2V1Beta2GoModel struct {
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
		JvmOptions *struct {
			GcLoggingEnabled *bool `tfsdk:"gc_logging_enabled" yaml:"gcLoggingEnabled,omitempty"`

			JavaSystemProperties *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"java_system_properties" yaml:"javaSystemProperties,omitempty"`

			_XX *map[string]string `tfsdk:"__xx" yaml:"-XX,omitempty"`

			_Xms *string `tfsdk:"___xms" yaml:"-Xms,omitempty"`

			_Xmx *string `tfsdk:"___xmx" yaml:"-Xmx,omitempty"`
		} `tfsdk:"jvm_options" yaml:"jvmOptions,omitempty"`

		Logging *struct {
			Loggers *map[string]string `tfsdk:"loggers" yaml:"loggers,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`

			ValueFrom *struct {
				ConfigMapKeyRef *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`
			} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
		} `tfsdk:"logging" yaml:"logging,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`

		ClientRackInitImage *string `tfsdk:"client_rack_init_image" yaml:"clientRackInitImage,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

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

		Tracing *struct {
			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"tracing" yaml:"tracing,omitempty"`

		ExternalConfiguration *struct {
			Env *[]struct {
				ValueFrom *struct {
					ConfigMapKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

					SecretKeyRef *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
				} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"env" yaml:"env,omitempty"`

			Volumes *[]struct {
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

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Secret *struct {
					Items *[]struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Mode *int64 `tfsdk:"mode" yaml:"mode,omitempty"`

						Path *string `tfsdk:"path" yaml:"path,omitempty"`
					} `tfsdk:"items" yaml:"items,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

					DefaultMode *int64 `tfsdk:"default_mode" yaml:"defaultMode,omitempty"`
				} `tfsdk:"secret" yaml:"secret,omitempty"`
			} `tfsdk:"volumes" yaml:"volumes,omitempty"`
		} `tfsdk:"external_configuration" yaml:"externalConfiguration,omitempty"`

		Mirrors *[]struct {
			TopicsPattern *string `tfsdk:"topics_pattern" yaml:"topicsPattern,omitempty"`

			CheckpointConnector *struct {
				Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

				Pause *bool `tfsdk:"pause" yaml:"pause,omitempty"`

				TasksMax *int64 `tfsdk:"tasks_max" yaml:"tasksMax,omitempty"`
			} `tfsdk:"checkpoint_connector" yaml:"checkpointConnector,omitempty"`

			GroupsBlacklistPattern *string `tfsdk:"groups_blacklist_pattern" yaml:"groupsBlacklistPattern,omitempty"`

			GroupsExcludePattern *string `tfsdk:"groups_exclude_pattern" yaml:"groupsExcludePattern,omitempty"`

			SourceCluster *string `tfsdk:"source_cluster" yaml:"sourceCluster,omitempty"`

			TargetCluster *string `tfsdk:"target_cluster" yaml:"targetCluster,omitempty"`

			TopicsExcludePattern *string `tfsdk:"topics_exclude_pattern" yaml:"topicsExcludePattern,omitempty"`

			GroupsPattern *string `tfsdk:"groups_pattern" yaml:"groupsPattern,omitempty"`

			HeartbeatConnector *struct {
				Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

				Pause *bool `tfsdk:"pause" yaml:"pause,omitempty"`

				TasksMax *int64 `tfsdk:"tasks_max" yaml:"tasksMax,omitempty"`
			} `tfsdk:"heartbeat_connector" yaml:"heartbeatConnector,omitempty"`

			SourceConnector *struct {
				Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

				Pause *bool `tfsdk:"pause" yaml:"pause,omitempty"`

				TasksMax *int64 `tfsdk:"tasks_max" yaml:"tasksMax,omitempty"`
			} `tfsdk:"source_connector" yaml:"sourceConnector,omitempty"`

			TopicsBlacklistPattern *string `tfsdk:"topics_blacklist_pattern" yaml:"topicsBlacklistPattern,omitempty"`
		} `tfsdk:"mirrors" yaml:"mirrors,omitempty"`

		JmxOptions *struct {
			Authentication *struct {
				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"authentication" yaml:"authentication,omitempty"`
		} `tfsdk:"jmx_options" yaml:"jmxOptions,omitempty"`

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

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

		Resources *struct {
			Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

			Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
		} `tfsdk:"resources" yaml:"resources,omitempty"`

		Clusters *[]struct {
			Alias *string `tfsdk:"alias" yaml:"alias,omitempty"`

			Authentication *struct {
				ConnectTimeoutSeconds *int64 `tfsdk:"connect_timeout_seconds" yaml:"connectTimeoutSeconds,omitempty"`

				PasswordSecret *struct {
					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"password_secret" yaml:"passwordSecret,omitempty"`

				ReadTimeoutSeconds *int64 `tfsdk:"read_timeout_seconds" yaml:"readTimeoutSeconds,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`

				AccessToken *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"access_token" yaml:"accessToken,omitempty"`

				AccessTokenIsJwt *bool `tfsdk:"access_token_is_jwt" yaml:"accessTokenIsJwt,omitempty"`

				TokenEndpointUri *string `tfsdk:"token_endpoint_uri" yaml:"tokenEndpointUri,omitempty"`

				Username *string `tfsdk:"username" yaml:"username,omitempty"`

				MaxTokenExpirySeconds *int64 `tfsdk:"max_token_expiry_seconds" yaml:"maxTokenExpirySeconds,omitempty"`

				RefreshToken *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"refresh_token" yaml:"refreshToken,omitempty"`

				Scope *string `tfsdk:"scope" yaml:"scope,omitempty"`

				CertificateAndKey *struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"certificate_and_key" yaml:"certificateAndKey,omitempty"`

				ClientSecret *struct {
					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`
				} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

				DisableTlsHostnameVerification *bool `tfsdk:"disable_tls_hostname_verification" yaml:"disableTlsHostnameVerification,omitempty"`

				TlsTrustedCertificates *[]struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"tls_trusted_certificates" yaml:"tlsTrustedCertificates,omitempty"`

				Audience *string `tfsdk:"audience" yaml:"audience,omitempty"`

				ClientId *string `tfsdk:"client_id" yaml:"clientId,omitempty"`
			} `tfsdk:"authentication" yaml:"authentication,omitempty"`

			BootstrapServers *string `tfsdk:"bootstrap_servers" yaml:"bootstrapServers,omitempty"`

			Config *map[string]string `tfsdk:"config" yaml:"config,omitempty"`

			Tls *struct {
				TrustedCertificates *[]struct {
					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"trusted_certificates" yaml:"trustedCertificates,omitempty"`
			} `tfsdk:"tls" yaml:"tls,omitempty"`
		} `tfsdk:"clusters" yaml:"clusters,omitempty"`

		ConnectCluster *string `tfsdk:"connect_cluster" yaml:"connectCluster,omitempty"`

		LivenessProbe *struct {
			FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

			PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

			SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

			TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
		} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

		Template *struct {
			ConnectContainer *struct {
				Env *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"env" yaml:"env,omitempty"`

				SecurityContext *struct {
					SeccompProfile *struct {
						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`
					} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

					Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

					ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

					RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

					SeLinuxOptions *struct {
						User *string `tfsdk:"user" yaml:"user,omitempty"`

						Level *string `tfsdk:"level" yaml:"level,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

					Capabilities *struct {
						Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

						Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
					} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

					ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

					WindowsOptions *struct {
						HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`

						GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
			} `tfsdk:"connect_container" yaml:"connectContainer,omitempty"`

			Deployment *struct {
				DeploymentStrategy *string `tfsdk:"deployment_strategy" yaml:"deploymentStrategy,omitempty"`

				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`
			} `tfsdk:"deployment" yaml:"deployment,omitempty"`

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

					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					WindowsOptions *struct {
						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

						HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`

						GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

					Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

					ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

					ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

					RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

					SeLinuxOptions *struct {
						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`

						Level *string `tfsdk:"level" yaml:"level,omitempty"`
					} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

					SeccompProfile *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
			} `tfsdk:"init_container" yaml:"initContainer,omitempty"`

			Pod *struct {
				HostAliases *[]struct {
					Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

					Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
				} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

				ImagePullSecrets *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

				SchedulerName *string `tfsdk:"scheduler_name" yaml:"schedulerName,omitempty"`

				SecurityContext *struct {
					FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

					SeccompProfile *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

					SupplementalGroups *[]string `tfsdk:"supplemental_groups" yaml:"supplementalGroups,omitempty"`

					Sysctls *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"sysctls" yaml:"sysctls,omitempty"`

					FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" yaml:"fsGroupChangePolicy,omitempty"`

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

						HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

				TmpDirSizeLimit *string `tfsdk:"tmp_dir_size_limit" yaml:"tmpDirSizeLimit,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

				Affinity *struct {
					NodeAffinity *struct {
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

						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							Preference *struct {
								MatchFields *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`

								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
							} `tfsdk:"preference" yaml:"preference,omitempty"`

							Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`

					PodAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							PodAffinityTerm *struct {
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
								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`

								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

							NamespaceSelector *struct {
								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`

								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
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
									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

							NamespaceSelector *struct {
								MatchExpressions *[]struct {
									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

							Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

							TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" yaml:"affinity,omitempty"`

				EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`

				PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

				TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

				TopologySpreadConstraints *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

							Key *string `tfsdk:"key" yaml:"key,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

					MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

					MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

					TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

					WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
				} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
			} `tfsdk:"pod" yaml:"pod,omitempty"`

			BuildConfig *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`

				PullSecret *string `tfsdk:"pull_secret" yaml:"pullSecret,omitempty"`
			} `tfsdk:"build_config" yaml:"buildConfig,omitempty"`

			BuildPod *struct {
				Affinity *struct {
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
									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`

									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

							Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

						RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

									Key *string `tfsdk:"key" yaml:"key,omitempty"`
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
										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

										Key *string `tfsdk:"key" yaml:"key,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

							Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

						RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
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

							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`

					NodeAffinity *struct {
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
					} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`
				} `tfsdk:"affinity" yaml:"affinity,omitempty"`

				EnableServiceLinks *bool `tfsdk:"enable_service_links" yaml:"enableServiceLinks,omitempty"`

				HostAliases *[]struct {
					Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

					Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
				} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

				PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

				SchedulerName *string `tfsdk:"scheduler_name" yaml:"schedulerName,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

				ImagePullSecrets *[]struct {
					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`

				SecurityContext *struct {
					Sysctls *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"sysctls" yaml:"sysctls,omitempty"`

					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

					SeLinuxOptions *struct {
						Level *string `tfsdk:"level" yaml:"level,omitempty"`

						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`
					} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`

					SeccompProfile *struct {
						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`
					} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

					SupplementalGroups *[]string `tfsdk:"supplemental_groups" yaml:"supplementalGroups,omitempty"`

					WindowsOptions *struct {
						GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`

						HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

					FsGroup *int64 `tfsdk:"fs_group" yaml:"fsGroup,omitempty"`

					FsGroupChangePolicy *string `tfsdk:"fs_group_change_policy" yaml:"fsGroupChangePolicy,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

				TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

				TmpDirSizeLimit *string `tfsdk:"tmp_dir_size_limit" yaml:"tmpDirSizeLimit,omitempty"`

				TopologySpreadConstraints *[]struct {
					MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

					TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

					WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`

					LabelSelector *struct {
						MatchExpressions *[]struct {
							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`

							Key *string `tfsdk:"key" yaml:"key,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

					MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`
				} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
			} `tfsdk:"build_pod" yaml:"buildPod,omitempty"`

			BuildServiceAccount *struct {
				Metadata *struct {
					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`
			} `tfsdk:"build_service_account" yaml:"buildServiceAccount,omitempty"`

			ClusterRoleBinding *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`
			} `tfsdk:"cluster_role_binding" yaml:"clusterRoleBinding,omitempty"`

			ServiceAccount *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`
			} `tfsdk:"service_account" yaml:"serviceAccount,omitempty"`

			ApiService *struct {
				IpFamilies *[]string `tfsdk:"ip_families" yaml:"ipFamilies,omitempty"`

				IpFamilyPolicy *string `tfsdk:"ip_family_policy" yaml:"ipFamilyPolicy,omitempty"`

				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`
			} `tfsdk:"api_service" yaml:"apiService,omitempty"`

			BuildContainer *struct {
				Env *[]struct {
					Value *string `tfsdk:"value" yaml:"value,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"env" yaml:"env,omitempty"`

				SecurityContext *struct {
					Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

					ProcMount *string `tfsdk:"proc_mount" yaml:"procMount,omitempty"`

					RunAsUser *int64 `tfsdk:"run_as_user" yaml:"runAsUser,omitempty"`

					SeccompProfile *struct {
						LocalhostProfile *string `tfsdk:"localhost_profile" yaml:"localhostProfile,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`
					} `tfsdk:"seccomp_profile" yaml:"seccompProfile,omitempty"`

					WindowsOptions *struct {
						HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`

						GmsaCredentialSpec *string `tfsdk:"gmsa_credential_spec" yaml:"gmsaCredentialSpec,omitempty"`

						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" yaml:"gmsaCredentialSpecName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`

					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" yaml:"allowPrivilegeEscalation,omitempty"`

					Capabilities *struct {
						Add *[]string `tfsdk:"add" yaml:"add,omitempty"`

						Drop *[]string `tfsdk:"drop" yaml:"drop,omitempty"`
					} `tfsdk:"capabilities" yaml:"capabilities,omitempty"`

					ReadOnlyRootFilesystem *bool `tfsdk:"read_only_root_filesystem" yaml:"readOnlyRootFilesystem,omitempty"`

					RunAsGroup *int64 `tfsdk:"run_as_group" yaml:"runAsGroup,omitempty"`

					RunAsNonRoot *bool `tfsdk:"run_as_non_root" yaml:"runAsNonRoot,omitempty"`

					SeLinuxOptions *struct {
						Role *string `tfsdk:"role" yaml:"role,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						User *string `tfsdk:"user" yaml:"user,omitempty"`

						Level *string `tfsdk:"level" yaml:"level,omitempty"`
					} `tfsdk:"se_linux_options" yaml:"seLinuxOptions,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
			} `tfsdk:"build_container" yaml:"buildContainer,omitempty"`

			JmxSecret *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`
			} `tfsdk:"jmx_secret" yaml:"jmxSecret,omitempty"`

			PodDisruptionBudget *struct {
				MaxUnavailable *int64 `tfsdk:"max_unavailable" yaml:"maxUnavailable,omitempty"`

				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
				} `tfsdk:"metadata" yaml:"metadata,omitempty"`
			} `tfsdk:"pod_disruption_budget" yaml:"podDisruptionBudget,omitempty"`
		} `tfsdk:"template" yaml:"template,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource() resource.Resource {
	return &KafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource{}
}

func (r *KafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kafka_strimzi_io_kafka_mirror_maker2_v1beta2"
}

func (r *KafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "The specification of the Kafka MirrorMaker 2.0 cluster.",
				MarkdownDescription: "The specification of the Kafka MirrorMaker 2.0 cluster.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"jvm_options": {
						Description:         "JVM Options for pods.",
						MarkdownDescription: "JVM Options for pods.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

							"__xx": {
								Description:         "A map of -XX options to the JVM.",
								MarkdownDescription: "A map of -XX options to the JVM.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"___xms": {
								Description:         "-Xms option to to the JVM.",
								MarkdownDescription: "-Xms option to to the JVM.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"___xmx": {
								Description:         "-Xmx option to to the JVM.",
								MarkdownDescription: "-Xmx option to to the JVM.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"logging": {
						Description:         "Logging configuration for Kafka Connect.",
						MarkdownDescription: "Logging configuration for Kafka Connect.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"loggers": {
								Description:         "A Map from logger name to logger level.",
								MarkdownDescription: "A Map from logger name to logger level.",

								Type: types.MapType{ElemType: types.StringType},

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

					"version": {
						Description:         "The Kafka Connect version. Defaults to {DefaultKafkaVersion}. Consult the user documentation to understand the process required to upgrade or downgrade the version.",
						MarkdownDescription: "The Kafka Connect version. Defaults to {DefaultKafkaVersion}. Consult the user documentation to understand the process required to upgrade or downgrade the version.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"client_rack_init_image": {
						Description:         "The image of the init container used for initializing the 'client.rack'.",
						MarkdownDescription: "The image of the init container used for initializing the 'client.rack'.",

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

					"rack": {
						Description:         "Configuration of the node label which will be used as the 'client.rack' consumer configuration.",
						MarkdownDescription: "Configuration of the node label which will be used as the 'client.rack' consumer configuration.",

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
							},

							"initial_delay_seconds": {
								Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
								MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

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
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout_seconds": {
								Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
								MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

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

					"tracing": {
						Description:         "The configuration of tracing in Kafka Connect.",
						MarkdownDescription: "The configuration of tracing in Kafka Connect.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"type": {
								Description:         "Type of the tracing used. Currently the only supported type is 'jaeger' for Jaeger tracing. The Jaeger tracing is deprecated.",
								MarkdownDescription: "Type of the tracing used. Currently the only supported type is 'jaeger' for Jaeger tracing. The Jaeger tracing is deprecated.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_configuration": {
						Description:         "Pass data from Secrets or ConfigMaps to the Kafka Connect pods and use them to configure connectors.",
						MarkdownDescription: "Pass data from Secrets or ConfigMaps to the Kafka Connect pods and use them to configure connectors.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"env": {
								Description:         "Makes data from a Secret or ConfigMap available in the Kafka Connect pods as environment variables.",
								MarkdownDescription: "Makes data from a Secret or ConfigMap available in the Kafka Connect pods as environment variables.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"value_from": {
										Description:         "Value of the environment variable which will be passed to the Kafka Connect pods. It can be passed either as a reference to Secret or ConfigMap field. The field has to specify exactly one Secret or ConfigMap.",
										MarkdownDescription: "Value of the environment variable which will be passed to the Kafka Connect pods. It can be passed either as a reference to Secret or ConfigMap field. The field has to specify exactly one Secret or ConfigMap.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_key_ref": {
												Description:         "Reference to a key in a ConfigMap.",
												MarkdownDescription: "Reference to a key in a ConfigMap.",

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

											"secret_key_ref": {
												Description:         "Reference to a key in a Secret.",
												MarkdownDescription: "Reference to a key in a Secret.",

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

									"name": {
										Description:         "Name of the environment variable which will be passed to the Kafka Connect pods. The name of the environment variable cannot start with 'KAFKA_' or 'STRIMZI_'.",
										MarkdownDescription: "Name of the environment variable which will be passed to the Kafka Connect pods. The name of the environment variable cannot start with 'KAFKA_' or 'STRIMZI_'.",

										Type: types.StringType,

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
								Description:         "Makes data from a Secret or ConfigMap available in the Kafka Connect pods as volumes.",
								MarkdownDescription: "Makes data from a Secret or ConfigMap available in the Kafka Connect pods as volumes.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"config_map": {
										Description:         "Reference to a key in a ConfigMap. Exactly one Secret or ConfigMap has to be specified.",
										MarkdownDescription: "Reference to a key in a ConfigMap. Exactly one Secret or ConfigMap has to be specified.",

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

									"name": {
										Description:         "Name of the volume which will be added to the Kafka Connect pods.",
										MarkdownDescription: "Name of the volume which will be added to the Kafka Connect pods.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret": {
										Description:         "Reference to a key in a Secret. Exactly one Secret or ConfigMap has to be specified.",
										MarkdownDescription: "Reference to a key in a Secret. Exactly one Secret or ConfigMap has to be specified.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"items": {
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

														Required: false,
														Optional: true,
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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"mirrors": {
						Description:         "Configuration of the MirrorMaker 2.0 connectors.",
						MarkdownDescription: "Configuration of the MirrorMaker 2.0 connectors.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"topics_pattern": {
								Description:         "A regular expression matching the topics to be mirrored, for example, 'topic1|topic2|topic3'. Comma-separated lists are also supported.",
								MarkdownDescription: "A regular expression matching the topics to be mirrored, for example, 'topic1|topic2|topic3'. Comma-separated lists are also supported.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"checkpoint_connector": {
								Description:         "The specification of the Kafka MirrorMaker 2.0 checkpoint connector.",
								MarkdownDescription: "The specification of the Kafka MirrorMaker 2.0 checkpoint connector.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config": {
										Description:         "The Kafka Connector configuration. The following properties cannot be set: connector.class, tasks.max.",
										MarkdownDescription: "The Kafka Connector configuration. The following properties cannot be set: connector.class, tasks.max.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pause": {
										Description:         "Whether the connector should be paused. Defaults to false.",
										MarkdownDescription: "Whether the connector should be paused. Defaults to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tasks_max": {
										Description:         "The maximum number of tasks for the Kafka Connector.",
										MarkdownDescription: "The maximum number of tasks for the Kafka Connector.",

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

							"groups_blacklist_pattern": {
								Description:         "A regular expression matching the consumer groups to exclude from mirroring. Comma-separated lists are also supported.",
								MarkdownDescription: "A regular expression matching the consumer groups to exclude from mirroring. Comma-separated lists are also supported.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"groups_exclude_pattern": {
								Description:         "A regular expression matching the consumer groups to exclude from mirroring. Comma-separated lists are also supported.",
								MarkdownDescription: "A regular expression matching the consumer groups to exclude from mirroring. Comma-separated lists are also supported.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"source_cluster": {
								Description:         "The alias of the source cluster used by the Kafka MirrorMaker 2.0 connectors. The alias must match a cluster in the list at 'spec.clusters'.",
								MarkdownDescription: "The alias of the source cluster used by the Kafka MirrorMaker 2.0 connectors. The alias must match a cluster in the list at 'spec.clusters'.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"target_cluster": {
								Description:         "The alias of the target cluster used by the Kafka MirrorMaker 2.0 connectors. The alias must match a cluster in the list at 'spec.clusters'.",
								MarkdownDescription: "The alias of the target cluster used by the Kafka MirrorMaker 2.0 connectors. The alias must match a cluster in the list at 'spec.clusters'.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"topics_exclude_pattern": {
								Description:         "A regular expression matching the topics to exclude from mirroring. Comma-separated lists are also supported.",
								MarkdownDescription: "A regular expression matching the topics to exclude from mirroring. Comma-separated lists are also supported.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"groups_pattern": {
								Description:         "A regular expression matching the consumer groups to be mirrored. Comma-separated lists are also supported.",
								MarkdownDescription: "A regular expression matching the consumer groups to be mirrored. Comma-separated lists are also supported.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"heartbeat_connector": {
								Description:         "The specification of the Kafka MirrorMaker 2.0 heartbeat connector.",
								MarkdownDescription: "The specification of the Kafka MirrorMaker 2.0 heartbeat connector.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config": {
										Description:         "The Kafka Connector configuration. The following properties cannot be set: connector.class, tasks.max.",
										MarkdownDescription: "The Kafka Connector configuration. The following properties cannot be set: connector.class, tasks.max.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pause": {
										Description:         "Whether the connector should be paused. Defaults to false.",
										MarkdownDescription: "Whether the connector should be paused. Defaults to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tasks_max": {
										Description:         "The maximum number of tasks for the Kafka Connector.",
										MarkdownDescription: "The maximum number of tasks for the Kafka Connector.",

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

							"source_connector": {
								Description:         "The specification of the Kafka MirrorMaker 2.0 source connector.",
								MarkdownDescription: "The specification of the Kafka MirrorMaker 2.0 source connector.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config": {
										Description:         "The Kafka Connector configuration. The following properties cannot be set: connector.class, tasks.max.",
										MarkdownDescription: "The Kafka Connector configuration. The following properties cannot be set: connector.class, tasks.max.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pause": {
										Description:         "Whether the connector should be paused. Defaults to false.",
										MarkdownDescription: "Whether the connector should be paused. Defaults to false.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tasks_max": {
										Description:         "The maximum number of tasks for the Kafka Connector.",
										MarkdownDescription: "The maximum number of tasks for the Kafka Connector.",

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

							"topics_blacklist_pattern": {
								Description:         "A regular expression matching the topics to exclude from mirroring. Comma-separated lists are also supported.",
								MarkdownDescription: "A regular expression matching the topics to exclude from mirroring. Comma-separated lists are also supported.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"jmx_options": {
						Description:         "JMX Options.",
						MarkdownDescription: "JMX Options.",

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
									},
								}),

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

					"replicas": {
						Description:         "The number of pods in the Kafka Connect group.",
						MarkdownDescription: "The number of pods in the Kafka Connect group.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "The maximum limits for CPU and memory resources and the requested initial resources.",
						MarkdownDescription: "The maximum limits for CPU and memory resources and the requested initial resources.",

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

					"clusters": {
						Description:         "Kafka clusters for mirroring.",
						MarkdownDescription: "Kafka clusters for mirroring.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"alias": {
								Description:         "Alias used to reference the Kafka cluster.",
								MarkdownDescription: "Alias used to reference the Kafka cluster.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"authentication": {
								Description:         "Authentication configuration for connecting to the cluster.",
								MarkdownDescription: "Authentication configuration for connecting to the cluster.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"connect_timeout_seconds": {
										Description:         "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",
										MarkdownDescription: "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"password_secret": {
										Description:         "Reference to the 'Secret' which holds the password.",
										MarkdownDescription: "Reference to the 'Secret' which holds the password.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"password": {
												Description:         "The name of the key in the Secret under which the password is stored.",
												MarkdownDescription: "The name of the key in the Secret under which the password is stored.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"secret_name": {
												Description:         "The name of the Secret containing the password.",
												MarkdownDescription: "The name of the Secret containing the password.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

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

									"type": {
										Description:         "Authentication type. Currently the only supported types are 'tls', 'scram-sha-256', 'scram-sha-512', and 'plain'. 'scram-sha-256' and 'scram-sha-512' types use SASL SCRAM-SHA-256 and SASL SCRAM-SHA-512 Authentication, respectively. 'plain' type uses SASL PLAIN Authentication. 'oauth' type uses SASL OAUTHBEARER Authentication. The 'tls' type uses TLS Client Authentication. The 'tls' type is supported only over TLS connections.",
										MarkdownDescription: "Authentication type. Currently the only supported types are 'tls', 'scram-sha-256', 'scram-sha-512', and 'plain'. 'scram-sha-256' and 'scram-sha-512' types use SASL SCRAM-SHA-256 and SASL SCRAM-SHA-512 Authentication, respectively. 'plain' type uses SASL PLAIN Authentication. 'oauth' type uses SASL OAUTHBEARER Authentication. The 'tls' type uses TLS Client Authentication. The 'tls' type is supported only over TLS connections.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"access_token": {
										Description:         "Link to Kubernetes Secret containing the access token which was obtained from the authorization server.",
										MarkdownDescription: "Link to Kubernetes Secret containing the access token which was obtained from the authorization server.",

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

									"access_token_is_jwt": {
										Description:         "Configure whether access token should be treated as JWT. This should be set to 'false' if the authorization server returns opaque tokens. Defaults to 'true'.",
										MarkdownDescription: "Configure whether access token should be treated as JWT. This should be set to 'false' if the authorization server returns opaque tokens. Defaults to 'true'.",

										Type: types.BoolType,

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

									"username": {
										Description:         "Username used for the authentication.",
										MarkdownDescription: "Username used for the authentication.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_token_expiry_seconds": {
										Description:         "Set or limit time-to-live of the access tokens to the specified number of seconds. This should be set if the authorization server returns opaque tokens.",
										MarkdownDescription: "Set or limit time-to-live of the access tokens to the specified number of seconds. This should be set if the authorization server returns opaque tokens.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"refresh_token": {
										Description:         "Link to Kubernetes Secret containing the refresh token which can be used to obtain access token from the authorization server.",
										MarkdownDescription: "Link to Kubernetes Secret containing the refresh token which can be used to obtain access token from the authorization server.",

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

									"scope": {
										Description:         "OAuth scope to use when authenticating against the authorization server. Some authorization servers require this to be set. The possible values depend on how authorization server is configured. By default 'scope' is not specified when doing the token endpoint request.",
										MarkdownDescription: "OAuth scope to use when authenticating against the authorization server. Some authorization servers require this to be set. The possible values depend on how authorization server is configured. By default 'scope' is not specified when doing the token endpoint request.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"certificate_and_key": {
										Description:         "Reference to the 'Secret' which holds the certificate and private key pair.",
										MarkdownDescription: "Reference to the 'Secret' which holds the certificate and private key pair.",

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

									"client_secret": {
										Description:         "Link to Kubernetes Secret containing the OAuth client secret which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",
										MarkdownDescription: "Link to Kubernetes Secret containing the OAuth client secret which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"secret_name": {
												Description:         "The name of the Kubernetes Secret containing the secret value.",
												MarkdownDescription: "The name of the Kubernetes Secret containing the secret value.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"key": {
												Description:         "The key under which the secret value is stored in the Kubernetes Secret.",
												MarkdownDescription: "The key under which the secret value is stored in the Kubernetes Secret.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

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

									"audience": {
										Description:         "OAuth audience to use when authenticating against the authorization server. Some authorization servers require the audience to be explicitly set. The possible values depend on how the authorization server is configured. By default, 'audience' is not specified when performing the token endpoint request.",
										MarkdownDescription: "OAuth audience to use when authenticating against the authorization server. Some authorization servers require the audience to be explicitly set. The possible values depend on how the authorization server is configured. By default, 'audience' is not specified when performing the token endpoint request.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bootstrap_servers": {
								Description:         "A comma-separated list of 'host:port' pairs for establishing the connection to the Kafka cluster.",
								MarkdownDescription: "A comma-separated list of 'host:port' pairs for establishing the connection to the Kafka cluster.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"config": {
								Description:         "The MirrorMaker 2.0 cluster config. Properties with the following prefixes cannot be set: ssl., sasl., security., listeners, plugin.path, rest., bootstrap.servers, consumer.interceptor.classes, producer.interceptor.classes (with the exception of: ssl.endpoint.identification.algorithm, ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols).",
								MarkdownDescription: "The MirrorMaker 2.0 cluster config. Properties with the following prefixes cannot be set: ssl., sasl., security., listeners, plugin.path, rest., bootstrap.servers, consumer.interceptor.classes, producer.interceptor.classes (with the exception of: ssl.endpoint.identification.algorithm, ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols).",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls": {
								Description:         "TLS configuration for connecting MirrorMaker 2.0 connectors to a cluster.",
								MarkdownDescription: "TLS configuration for connecting MirrorMaker 2.0 connectors to a cluster.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"trusted_certificates": {
										Description:         "Trusted certificates for TLS connection.",
										MarkdownDescription: "Trusted certificates for TLS connection.",

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
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"connect_cluster": {
						Description:         "The cluster alias used for Kafka Connect. The alias must match a cluster in the list at 'spec.clusters'.",
						MarkdownDescription: "The cluster alias used for Kafka Connect. The alias must match a cluster in the list at 'spec.clusters'.",

						Type: types.StringType,

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
							},

							"initial_delay_seconds": {
								Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
								MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",

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
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout_seconds": {
								Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
								MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",

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
						Description:         "Template for Kafka Connect and Kafka Mirror Maker 2 resources. The template allows users to specify how the 'Deployment', 'Pods' and 'Service' are generated.",
						MarkdownDescription: "Template for Kafka Connect and Kafka Mirror Maker 2 resources. The template allows users to specify how the 'Deployment', 'Pods' and 'Service' are generated.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"connect_container": {
								Description:         "Template for the Kafka Connect container.",
								MarkdownDescription: "Template for the Kafka Connect container.",

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

											"seccomp_profile": {
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

													"localhost_profile": {
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

											"windows_options": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
										}),

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
								Description:         "Template for Kafka Connect 'Deployment'.",
								MarkdownDescription: "Template for Kafka Connect 'Deployment'.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"deployment_strategy": {
										Description:         "DeploymentStrategy which will be used for this Deployment. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
										MarkdownDescription: "DeploymentStrategy which will be used for this Deployment. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",

										Type: types.StringType,

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

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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

											"run_as_group": {
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

													"level": {
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
										}),

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
								Description:         "Template for Kafka Connect 'Pods'.",
								MarkdownDescription: "Template for Kafka Connect 'Pods'.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

											"fs_group_change_policy": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

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

									"tmp_dir_size_limit": {
										Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
										MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
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

									"affinity": {
										Description:         "The pod's affinity rules.",
										MarkdownDescription: "The pod's affinity rules.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"node_affinity": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

													"preferred_during_scheduling_ignored_during_execution": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"preference": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																		Required: false,
																		Optional: true,
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

																	"match_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

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
																}),

																Required: false,
																Optional: true,
																Computed: false,
															},

															"namespace_selector": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"match_labels": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.MapType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

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

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

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

									"metadata": {
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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

									"priority_class_name": {
										Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
										MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",

										Type: types.StringType,

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

							"build_config": {
								Description:         "Template for the Kafka Connect BuildConfig used to build new container images. The BuildConfig is used only on OpenShift.",
								MarkdownDescription: "Template for the Kafka Connect BuildConfig used to build new container images. The BuildConfig is used only on OpenShift.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
										MarkdownDescription: "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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

									"pull_secret": {
										Description:         "Container Registry Secret with the credentials for pulling the base image.",
										MarkdownDescription: "Container Registry Secret with the credentials for pulling the base image.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"build_pod": {
								Description:         "Template for Kafka Connect Build 'Pods'. The build pod is used only on Kubernetes.",
								MarkdownDescription: "Template for Kafka Connect Build 'Pods'. The build pod is used only on Kubernetes.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"affinity": {
										Description:         "The pod's affinity rules.",
										MarkdownDescription: "The pod's affinity rules.",

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

																			"match_labels": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

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

																Required: false,
																Optional: true,
																Computed: false,
															},
														}),

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

																Required: false,
																Optional: true,
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
												}),

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

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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
										Description:         "Configures pod-level security attributes and common container settings.",
										MarkdownDescription: "Configures pod-level security attributes and common container settings.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

													"type": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"localhost_profile": {
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

											"run_as_non_root": {
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

									"termination_grace_period_seconds": {
										Description:         "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
										MarkdownDescription: "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tmp_dir_size_limit": {
										Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
										MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"topology_spread_constraints": {
										Description:         "The pod's topology spread constraints.",
										MarkdownDescription: "The pod's topology spread constraints.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"min_domains": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

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

											"label_selector": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"match_expressions": {
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

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"build_service_account": {
								Description:         "Template for the Kafka Connect Build service account.",
								MarkdownDescription: "Template for the Kafka Connect Build service account.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"annotations": {
												Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_role_binding": {
								Description:         "Template for the Kafka Connect ClusterRoleBinding.",
								MarkdownDescription: "Template for the Kafka Connect ClusterRoleBinding.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_account": {
								Description:         "Template for the Kafka Connect service account.",
								MarkdownDescription: "Template for the Kafka Connect service account.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"api_service": {
								Description:         "Template for Kafka Connect API 'Service'.",
								MarkdownDescription: "Template for Kafka Connect API 'Service'.",

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
									},

									"metadata": {
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"build_container": {
								Description:         "Template for the Kafka Connect Build container. The build container is used only on Kubernetes.",
								MarkdownDescription: "Template for the Kafka Connect Build container. The build container is used only on Kubernetes.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"env": {
										Description:         "Environment variables which should be applied to the container.",
										MarkdownDescription: "Environment variables which should be applied to the container.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"value": {
												Description:         "The environment variable value.",
												MarkdownDescription: "The environment variable value.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "The environment variable key.",
												MarkdownDescription: "The environment variable key.",

												Type: types.StringType,

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

											"run_as_user": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.Int64Type,

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

											"se_linux_options": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

													"level": {
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
								Description:         "Template for Secret of the Kafka Connect Cluster JMX authentication.",
								MarkdownDescription: "Template for Secret of the Kafka Connect Cluster JMX authentication.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"metadata": {
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_disruption_budget": {
								Description:         "Template for Kafka Connect 'PodDisruptionBudget'.",
								MarkdownDescription: "Template for Kafka Connect 'PodDisruptionBudget'.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"max_unavailable": {
										Description:         "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
										MarkdownDescription: "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"metadata": {
										Description:         "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",
										MarkdownDescription: "Metadata to apply to the 'PodDisruptionBudgetTemplate' resource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Annotations added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",
												MarkdownDescription: "Labels added to the resource template. Can be applied to different resources such as 'StatefulSets', 'Deployments', 'Pods', and 'Services'.",

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

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

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

func (r *KafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kafka_strimzi_io_kafka_mirror_maker2_v1beta2")

	var state KafkaStrimziIoKafkaMirrorMaker2V1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KafkaStrimziIoKafkaMirrorMaker2V1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kafka.strimzi.io/v1beta2")
	goModel.Kind = utilities.Ptr("KafkaMirrorMaker2")

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

func (r *KafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kafka_strimzi_io_kafka_mirror_maker2_v1beta2")
	// NO-OP: All data is already in Terraform state
}

func (r *KafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kafka_strimzi_io_kafka_mirror_maker2_v1beta2")

	var state KafkaStrimziIoKafkaMirrorMaker2V1Beta2TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KafkaStrimziIoKafkaMirrorMaker2V1Beta2GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kafka.strimzi.io/v1beta2")
	goModel.Kind = utilities.Ptr("KafkaMirrorMaker2")

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

func (r *KafkaStrimziIoKafkaMirrorMaker2V1Beta2Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kafka_strimzi_io_kafka_mirror_maker2_v1beta2")
	// NO-OP: Terraform removes the state automatically for us
}
