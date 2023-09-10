/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kafka_strimzi_io_v1beta2

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
	_ datasource.DataSource              = &KafkaStrimziIoKafkaBridgeV1Beta2DataSource{}
	_ datasource.DataSourceWithConfigure = &KafkaStrimziIoKafkaBridgeV1Beta2DataSource{}
)

func NewKafkaStrimziIoKafkaBridgeV1Beta2DataSource() datasource.DataSource {
	return &KafkaStrimziIoKafkaBridgeV1Beta2DataSource{}
}

type KafkaStrimziIoKafkaBridgeV1Beta2DataSource struct {
	kubernetesClient dynamic.Interface
}

type KafkaStrimziIoKafkaBridgeV1Beta2DataSourceData struct {
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
		AdminClient *struct {
			Config *map[string]string `tfsdk:"config" json:"config,omitempty"`
		} `tfsdk:"admin_client" json:"adminClient,omitempty"`
		Authentication *struct {
			AccessToken *struct {
				Key        *string `tfsdk:"key" json:"key,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"access_token" json:"accessToken,omitempty"`
			AccessTokenIsJwt  *bool   `tfsdk:"access_token_is_jwt" json:"accessTokenIsJwt,omitempty"`
			Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
			CertificateAndKey *struct {
				Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
				Key         *string `tfsdk:"key" json:"key,omitempty"`
				SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"certificate_and_key" json:"certificateAndKey,omitempty"`
			ClientId     *string `tfsdk:"client_id" json:"clientId,omitempty"`
			ClientSecret *struct {
				Key        *string `tfsdk:"key" json:"key,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"client_secret" json:"clientSecret,omitempty"`
			ConnectTimeoutSeconds          *int64 `tfsdk:"connect_timeout_seconds" json:"connectTimeoutSeconds,omitempty"`
			DisableTlsHostnameVerification *bool  `tfsdk:"disable_tls_hostname_verification" json:"disableTlsHostnameVerification,omitempty"`
			EnableMetrics                  *bool  `tfsdk:"enable_metrics" json:"enableMetrics,omitempty"`
			HttpRetries                    *int64 `tfsdk:"http_retries" json:"httpRetries,omitempty"`
			HttpRetryPauseMs               *int64 `tfsdk:"http_retry_pause_ms" json:"httpRetryPauseMs,omitempty"`
			MaxTokenExpirySeconds          *int64 `tfsdk:"max_token_expiry_seconds" json:"maxTokenExpirySeconds,omitempty"`
			PasswordSecret                 *struct {
				Password   *string `tfsdk:"password" json:"password,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"password_secret" json:"passwordSecret,omitempty"`
			ReadTimeoutSeconds *int64 `tfsdk:"read_timeout_seconds" json:"readTimeoutSeconds,omitempty"`
			RefreshToken       *struct {
				Key        *string `tfsdk:"key" json:"key,omitempty"`
				SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"refresh_token" json:"refreshToken,omitempty"`
			Scope                  *string `tfsdk:"scope" json:"scope,omitempty"`
			TlsTrustedCertificates *[]struct {
				Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
				SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"tls_trusted_certificates" json:"tlsTrustedCertificates,omitempty"`
			TokenEndpointUri *string `tfsdk:"token_endpoint_uri" json:"tokenEndpointUri,omitempty"`
			Type             *string `tfsdk:"type" json:"type,omitempty"`
			Username         *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"authentication" json:"authentication,omitempty"`
		BootstrapServers    *string `tfsdk:"bootstrap_servers" json:"bootstrapServers,omitempty"`
		ClientRackInitImage *string `tfsdk:"client_rack_init_image" json:"clientRackInitImage,omitempty"`
		Consumer            *struct {
			Config *map[string]string `tfsdk:"config" json:"config,omitempty"`
		} `tfsdk:"consumer" json:"consumer,omitempty"`
		EnableMetrics *bool `tfsdk:"enable_metrics" json:"enableMetrics,omitempty"`
		Http          *struct {
			Cors *struct {
				AllowedMethods *[]string `tfsdk:"allowed_methods" json:"allowedMethods,omitempty"`
				AllowedOrigins *[]string `tfsdk:"allowed_origins" json:"allowedOrigins,omitempty"`
			} `tfsdk:"cors" json:"cors,omitempty"`
			Port *int64 `tfsdk:"port" json:"port,omitempty"`
		} `tfsdk:"http" json:"http,omitempty"`
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
		Producer *struct {
			Config *map[string]string `tfsdk:"config" json:"config,omitempty"`
		} `tfsdk:"producer" json:"producer,omitempty"`
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
		Template *struct {
			ApiService *struct {
				IpFamilies     *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
				IpFamilyPolicy *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
				Metadata       *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
			} `tfsdk:"api_service" json:"apiService,omitempty"`
			BridgeContainer *struct {
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
			} `tfsdk:"bridge_container" json:"bridgeContainer,omitempty"`
			ClusterRoleBinding *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
			} `tfsdk:"cluster_role_binding" json:"clusterRoleBinding,omitempty"`
			Deployment *struct {
				DeploymentStrategy *string `tfsdk:"deployment_strategy" json:"deploymentStrategy,omitempty"`
				Metadata           *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
			} `tfsdk:"deployment" json:"deployment,omitempty"`
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
		} `tfsdk:"template" json:"template,omitempty"`
		Tls *struct {
			TrustedCertificates *[]struct {
				Certificate *string `tfsdk:"certificate" json:"certificate,omitempty"`
				SecretName  *string `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"trusted_certificates" json:"trustedCertificates,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		Tracing *struct {
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"tracing" json:"tracing,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KafkaStrimziIoKafkaBridgeV1Beta2DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kafka_strimzi_io_kafka_bridge_v1beta2"
}

func (r *KafkaStrimziIoKafkaBridgeV1Beta2DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "The specification of the Kafka Bridge.",
				MarkdownDescription: "The specification of the Kafka Bridge.",
				Attributes: map[string]schema.Attribute{
					"admin_client": schema.SingleNestedAttribute{
						Description:         "Kafka AdminClient related configuration.",
						MarkdownDescription: "Kafka AdminClient related configuration.",
						Attributes: map[string]schema.Attribute{
							"config": schema.MapAttribute{
								Description:         "The Kafka AdminClient configuration used for AdminClient instances created by the bridge.",
								MarkdownDescription: "The Kafka AdminClient configuration used for AdminClient instances created by the bridge.",
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

					"authentication": schema.SingleNestedAttribute{
						Description:         "Authentication configuration for connecting to the cluster.",
						MarkdownDescription: "Authentication configuration for connecting to the cluster.",
						Attributes: map[string]schema.Attribute{
							"access_token": schema.SingleNestedAttribute{
								Description:         "Link to Kubernetes Secret containing the access token which was obtained from the authorization server.",
								MarkdownDescription: "Link to Kubernetes Secret containing the access token which was obtained from the authorization server.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key under which the secret value is stored in the Kubernetes Secret.",
										MarkdownDescription: "The key under which the secret value is stored in the Kubernetes Secret.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_name": schema.StringAttribute{
										Description:         "The name of the Kubernetes Secret containing the secret value.",
										MarkdownDescription: "The name of the Kubernetes Secret containing the secret value.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"access_token_is_jwt": schema.BoolAttribute{
								Description:         "Configure whether access token should be treated as JWT. This should be set to 'false' if the authorization server returns opaque tokens. Defaults to 'true'.",
								MarkdownDescription: "Configure whether access token should be treated as JWT. This should be set to 'false' if the authorization server returns opaque tokens. Defaults to 'true'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"audience": schema.StringAttribute{
								Description:         "OAuth audience to use when authenticating against the authorization server. Some authorization servers require the audience to be explicitly set. The possible values depend on how the authorization server is configured. By default, 'audience' is not specified when performing the token endpoint request.",
								MarkdownDescription: "OAuth audience to use when authenticating against the authorization server. Some authorization servers require the audience to be explicitly set. The possible values depend on how the authorization server is configured. By default, 'audience' is not specified when performing the token endpoint request.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"certificate_and_key": schema.SingleNestedAttribute{
								Description:         "Reference to the 'Secret' which holds the certificate and private key pair.",
								MarkdownDescription: "Reference to the 'Secret' which holds the certificate and private key pair.",
								Attributes: map[string]schema.Attribute{
									"certificate": schema.StringAttribute{
										Description:         "The name of the file certificate in the Secret.",
										MarkdownDescription: "The name of the file certificate in the Secret.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"key": schema.StringAttribute{
										Description:         "The name of the private key in the Secret.",
										MarkdownDescription: "The name of the private key in the Secret.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_name": schema.StringAttribute{
										Description:         "The name of the Secret containing the certificate.",
										MarkdownDescription: "The name of the Secret containing the certificate.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"client_id": schema.StringAttribute{
								Description:         "OAuth Client ID which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",
								MarkdownDescription: "OAuth Client ID which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"client_secret": schema.SingleNestedAttribute{
								Description:         "Link to Kubernetes Secret containing the OAuth client secret which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",
								MarkdownDescription: "Link to Kubernetes Secret containing the OAuth client secret which the Kafka client can use to authenticate against the OAuth server and use the token endpoint URI.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key under which the secret value is stored in the Kubernetes Secret.",
										MarkdownDescription: "The key under which the secret value is stored in the Kubernetes Secret.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_name": schema.StringAttribute{
										Description:         "The name of the Kubernetes Secret containing the secret value.",
										MarkdownDescription: "The name of the Kubernetes Secret containing the secret value.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"connect_timeout_seconds": schema.Int64Attribute{
								Description:         "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",
								MarkdownDescription: "The connect timeout in seconds when connecting to authorization server. If not set, the effective connect timeout is 60 seconds.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"disable_tls_hostname_verification": schema.BoolAttribute{
								Description:         "Enable or disable TLS hostname verification. Default value is 'false'.",
								MarkdownDescription: "Enable or disable TLS hostname verification. Default value is 'false'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"enable_metrics": schema.BoolAttribute{
								Description:         "Enable or disable OAuth metrics. Default value is 'false'.",
								MarkdownDescription: "Enable or disable OAuth metrics. Default value is 'false'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"http_retries": schema.Int64Attribute{
								Description:         "The maximum number of retries to attempt if an initial HTTP request fails. If not set, the default is to not attempt any retries.",
								MarkdownDescription: "The maximum number of retries to attempt if an initial HTTP request fails. If not set, the default is to not attempt any retries.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"http_retry_pause_ms": schema.Int64Attribute{
								Description:         "The pause to take before retrying a failed HTTP request. If not set, the default is to not pause at all but to immediately repeat a request.",
								MarkdownDescription: "The pause to take before retrying a failed HTTP request. If not set, the default is to not pause at all but to immediately repeat a request.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"max_token_expiry_seconds": schema.Int64Attribute{
								Description:         "Set or limit time-to-live of the access tokens to the specified number of seconds. This should be set if the authorization server returns opaque tokens.",
								MarkdownDescription: "Set or limit time-to-live of the access tokens to the specified number of seconds. This should be set if the authorization server returns opaque tokens.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"password_secret": schema.SingleNestedAttribute{
								Description:         "Reference to the 'Secret' which holds the password.",
								MarkdownDescription: "Reference to the 'Secret' which holds the password.",
								Attributes: map[string]schema.Attribute{
									"password": schema.StringAttribute{
										Description:         "The name of the key in the Secret under which the password is stored.",
										MarkdownDescription: "The name of the key in the Secret under which the password is stored.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_name": schema.StringAttribute{
										Description:         "The name of the Secret containing the password.",
										MarkdownDescription: "The name of the Secret containing the password.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"read_timeout_seconds": schema.Int64Attribute{
								Description:         "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",
								MarkdownDescription: "The read timeout in seconds when connecting to authorization server. If not set, the effective read timeout is 60 seconds.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"refresh_token": schema.SingleNestedAttribute{
								Description:         "Link to Kubernetes Secret containing the refresh token which can be used to obtain access token from the authorization server.",
								MarkdownDescription: "Link to Kubernetes Secret containing the refresh token which can be used to obtain access token from the authorization server.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "The key under which the secret value is stored in the Kubernetes Secret.",
										MarkdownDescription: "The key under which the secret value is stored in the Kubernetes Secret.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"secret_name": schema.StringAttribute{
										Description:         "The name of the Kubernetes Secret containing the secret value.",
										MarkdownDescription: "The name of the Kubernetes Secret containing the secret value.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"scope": schema.StringAttribute{
								Description:         "OAuth scope to use when authenticating against the authorization server. Some authorization servers require this to be set. The possible values depend on how authorization server is configured. By default 'scope' is not specified when doing the token endpoint request.",
								MarkdownDescription: "OAuth scope to use when authenticating against the authorization server. Some authorization servers require this to be set. The possible values depend on how authorization server is configured. By default 'scope' is not specified when doing the token endpoint request.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tls_trusted_certificates": schema.ListNestedAttribute{
								Description:         "Trusted certificates for TLS connection to the OAuth server.",
								MarkdownDescription: "Trusted certificates for TLS connection to the OAuth server.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"certificate": schema.StringAttribute{
											Description:         "The name of the file certificate in the Secret.",
											MarkdownDescription: "The name of the file certificate in the Secret.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_name": schema.StringAttribute{
											Description:         "The name of the Secret containing the certificate.",
											MarkdownDescription: "The name of the Secret containing the certificate.",
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

							"token_endpoint_uri": schema.StringAttribute{
								Description:         "Authorization server token endpoint URI.",
								MarkdownDescription: "Authorization server token endpoint URI.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Authentication type. Currently the supported types are 'tls', 'scram-sha-256', 'scram-sha-512', 'plain', and 'oauth'. 'scram-sha-256' and 'scram-sha-512' types use SASL SCRAM-SHA-256 and SASL SCRAM-SHA-512 Authentication, respectively. 'plain' type uses SASL PLAIN Authentication. 'oauth' type uses SASL OAUTHBEARER Authentication. The 'tls' type uses TLS Client Authentication. The 'tls' type is supported only over TLS connections.",
								MarkdownDescription: "Authentication type. Currently the supported types are 'tls', 'scram-sha-256', 'scram-sha-512', 'plain', and 'oauth'. 'scram-sha-256' and 'scram-sha-512' types use SASL SCRAM-SHA-256 and SASL SCRAM-SHA-512 Authentication, respectively. 'plain' type uses SASL PLAIN Authentication. 'oauth' type uses SASL OAUTHBEARER Authentication. The 'tls' type uses TLS Client Authentication. The 'tls' type is supported only over TLS connections.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"username": schema.StringAttribute{
								Description:         "Username used for the authentication.",
								MarkdownDescription: "Username used for the authentication.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"bootstrap_servers": schema.StringAttribute{
						Description:         "A list of host:port pairs for establishing the initial connection to the Kafka cluster.",
						MarkdownDescription: "A list of host:port pairs for establishing the initial connection to the Kafka cluster.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"client_rack_init_image": schema.StringAttribute{
						Description:         "The image of the init container used for initializing the 'client.rack'.",
						MarkdownDescription: "The image of the init container used for initializing the 'client.rack'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"consumer": schema.SingleNestedAttribute{
						Description:         "Kafka consumer related configuration.",
						MarkdownDescription: "Kafka consumer related configuration.",
						Attributes: map[string]schema.Attribute{
							"config": schema.MapAttribute{
								Description:         "The Kafka consumer configuration used for consumer instances created by the bridge. Properties with the following prefixes cannot be set: ssl., bootstrap.servers, group.id, sasl., security. (with the exception of: ssl.endpoint.identification.algorithm, ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols).",
								MarkdownDescription: "The Kafka consumer configuration used for consumer instances created by the bridge. Properties with the following prefixes cannot be set: ssl., bootstrap.servers, group.id, sasl., security. (with the exception of: ssl.endpoint.identification.algorithm, ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols).",
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

					"enable_metrics": schema.BoolAttribute{
						Description:         "Enable the metrics for the Kafka Bridge. Default is false.",
						MarkdownDescription: "Enable the metrics for the Kafka Bridge. Default is false.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"http": schema.SingleNestedAttribute{
						Description:         "The HTTP related configuration.",
						MarkdownDescription: "The HTTP related configuration.",
						Attributes: map[string]schema.Attribute{
							"cors": schema.SingleNestedAttribute{
								Description:         "CORS configuration for the HTTP Bridge.",
								MarkdownDescription: "CORS configuration for the HTTP Bridge.",
								Attributes: map[string]schema.Attribute{
									"allowed_methods": schema.ListAttribute{
										Description:         "List of allowed HTTP methods.",
										MarkdownDescription: "List of allowed HTTP methods.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"allowed_origins": schema.ListAttribute{
										Description:         "List of allowed origins. Java regular expressions can be used.",
										MarkdownDescription: "List of allowed origins. Java regular expressions can be used.",
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

							"port": schema.Int64Attribute{
								Description:         "The port which is the server listening on.",
								MarkdownDescription: "The port which is the server listening on.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"image": schema.StringAttribute{
						Description:         "The docker image for the pods.",
						MarkdownDescription: "The docker image for the pods.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"jvm_options": schema.SingleNestedAttribute{
						Description:         "**Currently not supported** JVM Options for pods.",
						MarkdownDescription: "**Currently not supported** JVM Options for pods.",
						Attributes: map[string]schema.Attribute{
							"xx": schema.MapAttribute{
								Description:         "A map of -XX options to the JVM.",
								MarkdownDescription: "A map of -XX options to the JVM.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"xms": schema.StringAttribute{
								Description:         "-Xms option to to the JVM.",
								MarkdownDescription: "-Xms option to to the JVM.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"xmx": schema.StringAttribute{
								Description:         "-Xmx option to to the JVM.",
								MarkdownDescription: "-Xmx option to to the JVM.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"gc_logging_enabled": schema.BoolAttribute{
								Description:         "Specifies whether the Garbage Collection logging is enabled. The default is false.",
								MarkdownDescription: "Specifies whether the Garbage Collection logging is enabled. The default is false.",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
											Optional:            false,
											Computed:            true,
										},

										"value": schema.StringAttribute{
											Description:         "The system property value.",
											MarkdownDescription: "The system property value.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"liveness_probe": schema.SingleNestedAttribute{
						Description:         "Pod liveness checking.",
						MarkdownDescription: "Pod liveness checking.",
						Attributes: map[string]schema.Attribute{
							"failure_threshold": schema.Int64Attribute{
								Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"initial_delay_seconds": schema.Int64Attribute{
								Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
								MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
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
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"timeout_seconds": schema.Int64Attribute{
								Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
								MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
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
						Description:         "Logging configuration for Kafka Bridge.",
						MarkdownDescription: "Logging configuration for Kafka Bridge.",
						Attributes: map[string]schema.Attribute{
							"loggers": schema.MapAttribute{
								Description:         "A Map from logger name to logger level.",
								MarkdownDescription: "A Map from logger name to logger level.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "Logging type, must be either 'inline' or 'external'.",
								MarkdownDescription: "Logging type, must be either 'inline' or 'external'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"optional": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
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

					"producer": schema.SingleNestedAttribute{
						Description:         "Kafka producer related configuration.",
						MarkdownDescription: "Kafka producer related configuration.",
						Attributes: map[string]schema.Attribute{
							"config": schema.MapAttribute{
								Description:         "The Kafka producer configuration used for producer instances created by the bridge. Properties with the following prefixes cannot be set: ssl., bootstrap.servers, sasl., security. (with the exception of: ssl.endpoint.identification.algorithm, ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols).",
								MarkdownDescription: "The Kafka producer configuration used for producer instances created by the bridge. Properties with the following prefixes cannot be set: ssl., bootstrap.servers, sasl., security. (with the exception of: ssl.endpoint.identification.algorithm, ssl.cipher.suites, ssl.protocol, ssl.enabled.protocols).",
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

					"rack": schema.SingleNestedAttribute{
						Description:         "Configuration of the node label which will be used as the client.rack consumer configuration.",
						MarkdownDescription: "Configuration of the node label which will be used as the client.rack consumer configuration.",
						Attributes: map[string]schema.Attribute{
							"topology_key": schema.StringAttribute{
								Description:         "A key that matches labels assigned to the Kubernetes cluster nodes. The value of the label is used to set a broker's 'broker.rack' config, and the 'client.rack' config for Kafka Connect or MirrorMaker 2.",
								MarkdownDescription: "A key that matches labels assigned to the Kubernetes cluster nodes. The value of the label is used to set a broker's 'broker.rack' config, and the 'client.rack' config for Kafka Connect or MirrorMaker 2.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"readiness_probe": schema.SingleNestedAttribute{
						Description:         "Pod readiness checking.",
						MarkdownDescription: "Pod readiness checking.",
						Attributes: map[string]schema.Attribute{
							"failure_threshold": schema.Int64Attribute{
								Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"initial_delay_seconds": schema.Int64Attribute{
								Description:         "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
								MarkdownDescription: "The initial delay before first the health is first checked. Default to 15 seconds. Minimum value is 0.",
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
								Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
								MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"timeout_seconds": schema.Int64Attribute{
								Description:         "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
								MarkdownDescription: "The timeout for each attempted health check. Default to 5 seconds. Minimum value is 1.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"replicas": schema.Int64Attribute{
						Description:         "The number of pods in the 'Deployment'.  Defaults to '1'.",
						MarkdownDescription: "The number of pods in the 'Deployment'.  Defaults to '1'.",
						Required:            false,
						Optional:            false,
						Computed:            true,
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
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"requests": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
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

					"template": schema.SingleNestedAttribute{
						Description:         "Template for Kafka Bridge resources. The template allows users to specify how a 'Deployment' and 'Pod' is generated.",
						MarkdownDescription: "Template for Kafka Bridge resources. The template allows users to specify how a 'Deployment' and 'Pod' is generated.",
						Attributes: map[string]schema.Attribute{
							"api_service": schema.SingleNestedAttribute{
								Description:         "Template for Kafka Bridge API 'Service'.",
								MarkdownDescription: "Template for Kafka Bridge API 'Service'.",
								Attributes: map[string]schema.Attribute{
									"ip_families": schema.ListAttribute{
										Description:         "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
										MarkdownDescription: "Specifies the IP Families used by the service. Available options are 'IPv4' and 'IPv6'. If unspecified, Kubernetes will choose the default value based on the 'ipFamilyPolicy' setting.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"ip_family_policy": schema.StringAttribute{
										Description:         "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
										MarkdownDescription: "Specifies the IP Family Policy used by the service. Available options are 'SingleStack', 'PreferDualStack' and 'RequireDualStack'. 'SingleStack' is for a single IP family. 'PreferDualStack' is for two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. 'RequireDualStack' fails unless there are two IP families on dual-stack configured clusters. If unspecified, Kubernetes will choose the default value based on the service type.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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

							"bridge_container": schema.SingleNestedAttribute{
								Description:         "Template for the Kafka Bridge container.",
								MarkdownDescription: "Template for the Kafka Bridge container.",
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
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "The environment variable value.",
													MarkdownDescription: "The environment variable value.",
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

									"security_context": schema.SingleNestedAttribute{
										Description:         "Security context for the container.",
										MarkdownDescription: "Security context for the container.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
														Optional:            false,
														Computed:            true,
													},

													"drop": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"proc_mount": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"read_only_root_filesystem": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"se_linux_options": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"level": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"role": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"user": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"gmsa_credential_spec": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"gmsa_credential_spec_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"host_process": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"run_as_user_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

							"cluster_role_binding": schema.SingleNestedAttribute{
								Description:         "Template for the Kafka Bridge ClusterRoleBinding.",
								MarkdownDescription: "Template for the Kafka Bridge ClusterRoleBinding.",
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
												Optional:            false,
												Computed:            true,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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

							"deployment": schema.SingleNestedAttribute{
								Description:         "Template for Kafka Bridge 'Deployment'.",
								MarkdownDescription: "Template for Kafka Bridge 'Deployment'.",
								Attributes: map[string]schema.Attribute{
									"deployment_strategy": schema.StringAttribute{
										Description:         "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
										MarkdownDescription: "Pod replacement strategy for deployment configuration changes. Valid values are 'RollingUpdate' and 'Recreate'. Defaults to 'RollingUpdate'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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

							"init_container": schema.SingleNestedAttribute{
								Description:         "Template for the Kafka Bridge init container.",
								MarkdownDescription: "Template for the Kafka Bridge init container.",
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
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "The environment variable value.",
													MarkdownDescription: "The environment variable value.",
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

									"security_context": schema.SingleNestedAttribute{
										Description:         "Security context for the container.",
										MarkdownDescription: "Security context for the container.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
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
														Optional:            false,
														Computed:            true,
													},

													"drop": schema.ListAttribute{
														Description:         "",
														MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"proc_mount": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"read_only_root_filesystem": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"se_linux_options": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"level": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"role": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"user": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"gmsa_credential_spec": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"gmsa_credential_spec_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"host_process": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"run_as_user_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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

							"pod": schema.SingleNestedAttribute{
								Description:         "Template for Kafka Bridge 'Pods'.",
								MarkdownDescription: "Template for Kafka Bridge 'Pods'.",
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
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"match_fields": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"weight": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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

																		"match_fields": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"key": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																								Optional:            false,
																								Computed:            true,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"values": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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
																					Description:         "",
																					MarkdownDescription: "",
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
																								Optional:            false,
																								Computed:            true,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"values": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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
																					Description:         "",
																					MarkdownDescription: "",
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

																		"namespaces": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"weight": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
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
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
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

																"namespaces": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
												},
												Required: false,
												Optional: false,
												Computed: true,
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
																								Optional:            false,
																								Computed:            true,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"values": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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
																					Description:         "",
																					MarkdownDescription: "",
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
																								Optional:            false,
																								Computed:            true,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"values": schema.ListAttribute{
																								Description:         "",
																								MarkdownDescription: "",
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
																					Description:         "",
																					MarkdownDescription: "",
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

																		"namespaces": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"topology_key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},
																	},
																	Required: false,
																	Optional: false,
																	Computed: true,
																},

																"weight": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
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
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
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
																						Optional:            false,
																						Computed:            true,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            false,
																						Optional:            false,
																						Computed:            true,
																					},

																					"values": schema.ListAttribute{
																						Description:         "",
																						MarkdownDescription: "",
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
																			Description:         "",
																			MarkdownDescription: "",
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

																"namespaces": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

									"enable_service_links": schema.BoolAttribute{
										Description:         "Indicates whether information about services should be injected into Pod's environment variables.",
										MarkdownDescription: "Indicates whether information about services should be injected into Pod's environment variables.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
													Optional:            false,
													Computed:            true,
												},

												"ip": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"image_pull_secrets": schema.ListNestedAttribute{
										Description:         "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
										MarkdownDescription: "List of references to secrets in the same namespace to use for pulling any of the images used by this Pod. When the 'STRIMZI_IMAGE_PULL_SECRETS' environment variable in Cluster Operator and the 'imagePullSecrets' option are specified, only the 'imagePullSecrets' variable is used and the 'STRIMZI_IMAGE_PULL_SECRETS' variable is ignored.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata applied to the resource.",
										MarkdownDescription: "Metadata applied to the resource.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations added to the Kubernetes resource.",
												MarkdownDescription: "Annotations added to the Kubernetes resource.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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

									"priority_class_name": schema.StringAttribute{
										Description:         "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
										MarkdownDescription: "The name of the priority class used to assign priority to the pods. For more information about priority classes, see {K8sPriorityClass}.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"scheduler_name": schema.StringAttribute{
										Description:         "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
										MarkdownDescription: "The name of the scheduler used to dispatch this 'Pod'. If not specified, the default scheduler will be used.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"security_context": schema.SingleNestedAttribute{
										Description:         "Configures pod-level security attributes and common container settings.",
										MarkdownDescription: "Configures pod-level security attributes and common container settings.",
										Attributes: map[string]schema.Attribute{
											"fs_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"fs_group_change_policy": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"se_linux_options": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"level": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"role": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"user": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            false,
												Computed:            true,
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
															Optional:            false,
															Computed:            true,
														},

														"value": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"gmsa_credential_spec": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"gmsa_credential_spec_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"host_process": schema.BoolAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"run_as_user_name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
										Description:         "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
										MarkdownDescription: "The grace period is the duration in seconds after the processes running in the pod are sent a termination signal, and the time when the processes are forcibly halted with a kill signal. Set this value to longer than the expected cleanup time for your process. Value must be a non-negative integer. A zero value indicates delete immediately. You might need to increase the grace period for very large Kafka clusters, so that the Kafka brokers have enough time to transfer their work to another broker before they are terminated. Defaults to 30 seconds.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"tmp_dir_size_limit": schema.StringAttribute{
										Description:         "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
										MarkdownDescription: "Defines the total amount (for example '1Gi') of local storage required for temporary EmptyDir volume ('/tmp'). Default value is '5Mi'.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
													Optional:            false,
													Computed:            true,
												},

												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"operator": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"value": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
																		Optional:            false,
																		Computed:            true,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"values": schema.ListAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
															Description:         "",
															MarkdownDescription: "",
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
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"max_skew": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"min_domains": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_affinity_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"node_taints_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"topology_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"when_unsatisfiable": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"pod_disruption_budget": schema.SingleNestedAttribute{
								Description:         "Template for Kafka Bridge 'PodDisruptionBudget'.",
								MarkdownDescription: "Template for Kafka Bridge 'PodDisruptionBudget'.",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.Int64Attribute{
										Description:         "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
										MarkdownDescription: "Maximum number of unavailable pods to allow automatic Pod eviction. A Pod eviction is allowed when the 'maxUnavailable' number of pods or fewer are unavailable after the eviction. Setting this value to 0 prevents all voluntary evictions, so the pods must be evicted manually. Defaults to 1.",
										Required:            false,
										Optional:            false,
										Computed:            true,
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
												Optional:            false,
												Computed:            true,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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

							"service_account": schema.SingleNestedAttribute{
								Description:         "Template for the Kafka Bridge service account.",
								MarkdownDescription: "Template for the Kafka Bridge service account.",
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
												Optional:            false,
												Computed:            true,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels added to the Kubernetes resource.",
												MarkdownDescription: "Labels added to the Kubernetes resource.",
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

					"tls": schema.SingleNestedAttribute{
						Description:         "TLS configuration for connecting Kafka Bridge to the cluster.",
						MarkdownDescription: "TLS configuration for connecting Kafka Bridge to the cluster.",
						Attributes: map[string]schema.Attribute{
							"trusted_certificates": schema.ListNestedAttribute{
								Description:         "Trusted certificates for TLS connection.",
								MarkdownDescription: "Trusted certificates for TLS connection.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"certificate": schema.StringAttribute{
											Description:         "The name of the file certificate in the Secret.",
											MarkdownDescription: "The name of the file certificate in the Secret.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_name": schema.StringAttribute{
											Description:         "The name of the Secret containing the certificate.",
											MarkdownDescription: "The name of the Secret containing the certificate.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tracing": schema.SingleNestedAttribute{
						Description:         "The configuration of tracing in Kafka Bridge.",
						MarkdownDescription: "The configuration of tracing in Kafka Bridge.",
						Attributes: map[string]schema.Attribute{
							"type": schema.StringAttribute{
								Description:         "Type of the tracing used. Currently the only supported type is 'opentelemetry' for OpenTelemetry tracing. As of Strimzi 0.37.0, 'jaeger' type is not supported anymore and this option is ignored.",
								MarkdownDescription: "Type of the tracing used. Currently the only supported type is 'opentelemetry' for OpenTelemetry tracing. As of Strimzi 0.37.0, 'jaeger' type is not supported anymore and this option is ignored.",
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
	}
}

func (r *KafkaStrimziIoKafkaBridgeV1Beta2DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *KafkaStrimziIoKafkaBridgeV1Beta2DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_kafka_strimzi_io_kafka_bridge_v1beta2")

	var data KafkaStrimziIoKafkaBridgeV1Beta2DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "kafka.strimzi.io", Version: "v1beta2", Resource: "KafkaBridge"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse KafkaStrimziIoKafkaBridgeV1Beta2DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("kafka.strimzi.io/v1beta2")
	data.Kind = pointer.String("KafkaBridge")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
