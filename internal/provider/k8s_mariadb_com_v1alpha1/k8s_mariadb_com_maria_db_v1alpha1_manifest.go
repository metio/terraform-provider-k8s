/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package k8s_mariadb_com_v1alpha1

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
	_ datasource.DataSource = &K8SMariadbComMariaDbV1Alpha1Manifest{}
)

func NewK8SMariadbComMariaDbV1Alpha1Manifest() datasource.DataSource {
	return &K8SMariadbComMariaDbV1Alpha1Manifest{}
}

type K8SMariadbComMariaDbV1Alpha1Manifest struct{}

type K8SMariadbComMariaDbV1Alpha1ManifestData struct {
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
		Affinity *struct {
			AntiAffinityEnabled *bool `tfsdk:"anti_affinity_enabled" json:"antiAffinityEnabled,omitempty"`
			PodAntiAffinity     *struct {
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
						TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
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
					TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
				} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
			} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		Args          *[]string `tfsdk:"args" json:"args,omitempty"`
		BootstrapFrom *struct {
			BackupRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"backup_ref" json:"backupRef,omitempty"`
			RestoreJob *struct {
				Affinity *struct {
					AntiAffinityEnabled *bool `tfsdk:"anti_affinity_enabled" json:"antiAffinityEnabled,omitempty"`
					PodAntiAffinity     *struct {
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
								TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
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
							TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				Args     *[]string `tfsdk:"args" json:"args,omitempty"`
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Resources *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"restore_job" json:"restoreJob,omitempty"`
			S3 *struct {
				AccessKeyIdSecretKeyRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"access_key_id_secret_key_ref" json:"accessKeyIdSecretKeyRef,omitempty"`
				Bucket                      *string `tfsdk:"bucket" json:"bucket,omitempty"`
				Endpoint                    *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
				Prefix                      *string `tfsdk:"prefix" json:"prefix,omitempty"`
				Region                      *string `tfsdk:"region" json:"region,omitempty"`
				SecretAccessKeySecretKeyRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret_access_key_secret_key_ref" json:"secretAccessKeySecretKeyRef,omitempty"`
				SessionTokenSecretKeyRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"session_token_secret_key_ref" json:"sessionTokenSecretKeyRef,omitempty"`
				Tls *struct {
					CaSecretKeyRef *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"ca_secret_key_ref" json:"caSecretKeyRef,omitempty"`
					Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"s3" json:"s3,omitempty"`
			TargetRecoveryTime *string `tfsdk:"target_recovery_time" json:"targetRecoveryTime,omitempty"`
			Volume             *struct {
				Csi *struct {
					Driver               *string `tfsdk:"driver" json:"driver,omitempty"`
					FsType               *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					NodePublishSecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"node_publish_secret_ref" json:"nodePublishSecretRef,omitempty"`
					ReadOnly         *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
					VolumeAttributes *map[string]string `tfsdk:"volume_attributes" json:"volumeAttributes,omitempty"`
				} `tfsdk:"csi" json:"csi,omitempty"`
				EmptyDir *struct {
					Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
					SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
				} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
				Nfs *struct {
					Path     *string `tfsdk:"path" json:"path,omitempty"`
					ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					Server   *string `tfsdk:"server" json:"server,omitempty"`
				} `tfsdk:"nfs" json:"nfs,omitempty"`
				PersistentVolumeClaim *struct {
					ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
			} `tfsdk:"volume" json:"volume,omitempty"`
		} `tfsdk:"bootstrap_from" json:"bootstrapFrom,omitempty"`
		Command    *[]string `tfsdk:"command" json:"command,omitempty"`
		Connection *struct {
			HealthCheck *struct {
				Interval      *string `tfsdk:"interval" json:"interval,omitempty"`
				RetryInterval *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			Params         *map[string]string `tfsdk:"params" json:"params,omitempty"`
			Port           *int64             `tfsdk:"port" json:"port,omitempty"`
			SecretName     *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
			SecretTemplate *struct {
				DatabaseKey *string `tfsdk:"database_key" json:"databaseKey,omitempty"`
				Format      *string `tfsdk:"format" json:"format,omitempty"`
				HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
				Key         *string `tfsdk:"key" json:"key,omitempty"`
				Metadata    *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
				PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
				UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			} `tfsdk:"secret_template" json:"secretTemplate,omitempty"`
			ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
		} `tfsdk:"connection" json:"connection,omitempty"`
		Database *string `tfsdk:"database" json:"database,omitempty"`
		Env      *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
			ValueFrom *struct {
				ConfigMapKeyRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				FieldRef *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
				SecretKeyRef *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" json:"valueFrom,omitempty"`
		} `tfsdk:"env" json:"env,omitempty"`
		EnvFrom *[]struct {
			ConfigMapRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
			Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
			SecretRef *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"env_from" json:"envFrom,omitempty"`
		Galera *struct {
			Agent *struct {
				Args      *[]string `tfsdk:"args" json:"args,omitempty"`
				BasicAuth *struct {
					Enabled              *bool `tfsdk:"enabled" json:"enabled,omitempty"`
					PasswordSecretKeyRef *struct {
						Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"password_secret_key_ref" json:"passwordSecretKeyRef,omitempty"`
					Username *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"basic_auth" json:"basicAuth,omitempty"`
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
				Env     *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				EnvFrom *[]struct {
					ConfigMapRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
					Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
					SecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"env_from" json:"envFrom,omitempty"`
				GracefulShutdownTimeout *string `tfsdk:"graceful_shutdown_timeout" json:"gracefulShutdownTimeout,omitempty"`
				Image                   *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy         *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				KubernetesAuth          *struct {
					AuthDelegatorRoleName *string `tfsdk:"auth_delegator_role_name" json:"authDelegatorRoleName,omitempty"`
					Enabled               *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"kubernetes_auth" json:"kubernetesAuth,omitempty"`
				LivenessProbe *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					HttpGet          *struct {
						Host   *string `tfsdk:"host" json:"host,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *string `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
				Port           *int64 `tfsdk:"port" json:"port,omitempty"`
				ReadinessProbe *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					HttpGet          *struct {
						Host   *string `tfsdk:"host" json:"host,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *string `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
				Resources *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				SecurityContext *struct {
					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
					Capabilities             *struct {
						Add  *[]string `tfsdk:"add" json:"add,omitempty"`
						Drop *[]string `tfsdk:"drop" json:"drop,omitempty"`
					} `tfsdk:"capabilities" json:"capabilities,omitempty"`
					Privileged             *bool  `tfsdk:"privileged" json:"privileged,omitempty"`
					ReadOnlyRootFilesystem *bool  `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
					RunAsGroup             *int64 `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
					RunAsNonRoot           *bool  `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
					RunAsUser              *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
				} `tfsdk:"security_context" json:"securityContext,omitempty"`
				VolumeMounts *[]struct {
					MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SubPath   *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			} `tfsdk:"agent" json:"agent,omitempty"`
			AvailableWhenDonor *bool `tfsdk:"available_when_donor" json:"availableWhenDonor,omitempty"`
			Config             *struct {
				ReuseStorageVolume  *bool `tfsdk:"reuse_storage_volume" json:"reuseStorageVolume,omitempty"`
				VolumeClaimTemplate *struct {
					AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
					Metadata    *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
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
					StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			Enabled       *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			GaleraLibPath *string `tfsdk:"galera_lib_path" json:"galeraLibPath,omitempty"`
			InitContainer *struct {
				Args    *[]string `tfsdk:"args" json:"args,omitempty"`
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
				Env     *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Value     *string `tfsdk:"value" json:"value,omitempty"`
					ValueFrom *struct {
						ConfigMapKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
						SecretKeyRef *struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
					} `tfsdk:"value_from" json:"valueFrom,omitempty"`
				} `tfsdk:"env" json:"env,omitempty"`
				EnvFrom *[]struct {
					ConfigMapRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
					Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
					SecretRef *struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"env_from" json:"envFrom,omitempty"`
				Image           *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				LivenessProbe   *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					HttpGet          *struct {
						Host   *string `tfsdk:"host" json:"host,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *string `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
				ReadinessProbe *struct {
					Exec *struct {
						Command *[]string `tfsdk:"command" json:"command,omitempty"`
					} `tfsdk:"exec" json:"exec,omitempty"`
					FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
					HttpGet          *struct {
						Host   *string `tfsdk:"host" json:"host,omitempty"`
						Path   *string `tfsdk:"path" json:"path,omitempty"`
						Port   *string `tfsdk:"port" json:"port,omitempty"`
						Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
					} `tfsdk:"http_get" json:"httpGet,omitempty"`
					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
					PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
					SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
					TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
				Resources *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				SecurityContext *struct {
					AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
					Capabilities             *struct {
						Add  *[]string `tfsdk:"add" json:"add,omitempty"`
						Drop *[]string `tfsdk:"drop" json:"drop,omitempty"`
					} `tfsdk:"capabilities" json:"capabilities,omitempty"`
					Privileged             *bool  `tfsdk:"privileged" json:"privileged,omitempty"`
					ReadOnlyRootFilesystem *bool  `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
					RunAsGroup             *int64 `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
					RunAsNonRoot           *bool  `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
					RunAsUser              *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
				} `tfsdk:"security_context" json:"securityContext,omitempty"`
				VolumeMounts *[]struct {
					MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SubPath   *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
			} `tfsdk:"init_container" json:"initContainer,omitempty"`
			InitJob *struct {
				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				Resources *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
			} `tfsdk:"init_job" json:"initJob,omitempty"`
			Primary *struct {
				AutomaticFailover *bool  `tfsdk:"automatic_failover" json:"automaticFailover,omitempty"`
				PodIndex          *int64 `tfsdk:"pod_index" json:"podIndex,omitempty"`
			} `tfsdk:"primary" json:"primary,omitempty"`
			ProviderOptions *map[string]string `tfsdk:"provider_options" json:"providerOptions,omitempty"`
			Recovery        *struct {
				ClusterBootstrapTimeout    *string `tfsdk:"cluster_bootstrap_timeout" json:"clusterBootstrapTimeout,omitempty"`
				ClusterDownscaleTimeout    *string `tfsdk:"cluster_downscale_timeout" json:"clusterDownscaleTimeout,omitempty"`
				ClusterHealthyTimeout      *string `tfsdk:"cluster_healthy_timeout" json:"clusterHealthyTimeout,omitempty"`
				ClusterMonitorInterval     *string `tfsdk:"cluster_monitor_interval" json:"clusterMonitorInterval,omitempty"`
				ClusterUpscaleTimeout      *string `tfsdk:"cluster_upscale_timeout" json:"clusterUpscaleTimeout,omitempty"`
				Enabled                    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				ForceClusterBootstrapInPod *string `tfsdk:"force_cluster_bootstrap_in_pod" json:"forceClusterBootstrapInPod,omitempty"`
				Job                        *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					PodAffinity *bool `tfsdk:"pod_affinity" json:"podAffinity,omitempty"`
					Resources   *struct {
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
				} `tfsdk:"job" json:"job,omitempty"`
				MinClusterSize     *string `tfsdk:"min_cluster_size" json:"minClusterSize,omitempty"`
				PodRecoveryTimeout *string `tfsdk:"pod_recovery_timeout" json:"podRecoveryTimeout,omitempty"`
				PodSyncTimeout     *string `tfsdk:"pod_sync_timeout" json:"podSyncTimeout,omitempty"`
			} `tfsdk:"recovery" json:"recovery,omitempty"`
			ReplicaThreads *int64  `tfsdk:"replica_threads" json:"replicaThreads,omitempty"`
			Sst            *string `tfsdk:"sst" json:"sst,omitempty"`
		} `tfsdk:"galera" json:"galera,omitempty"`
		Image            *string `tfsdk:"image" json:"image,omitempty"`
		ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		InheritMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"inherit_metadata" json:"inheritMetadata,omitempty"`
		InitContainers *[]struct {
			Args            *[]string `tfsdk:"args" json:"args,omitempty"`
			Command         *[]string `tfsdk:"command" json:"command,omitempty"`
			Image           *string   `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string   `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			Resources       *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			VolumeMounts *[]struct {
				MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath   *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"init_containers" json:"initContainers,omitempty"`
		LivenessProbe *struct {
			Exec *struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
			} `tfsdk:"exec" json:"exec,omitempty"`
			FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			HttpGet          *struct {
				Host   *string `tfsdk:"host" json:"host,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Port   *string `tfsdk:"port" json:"port,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"http_get" json:"httpGet,omitempty"`
			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
			PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
			SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
			TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"liveness_probe" json:"livenessProbe,omitempty"`
		MaxScale *struct {
			Admin *struct {
				GuiEnabled *bool  `tfsdk:"gui_enabled" json:"guiEnabled,omitempty"`
				Port       *int64 `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"admin" json:"admin,omitempty"`
			Auth *struct {
				AdminPasswordSecretKeyRef *struct {
					Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"admin_password_secret_key_ref" json:"adminPasswordSecretKeyRef,omitempty"`
				AdminUsername              *string `tfsdk:"admin_username" json:"adminUsername,omitempty"`
				ClientMaxConnections       *int64  `tfsdk:"client_max_connections" json:"clientMaxConnections,omitempty"`
				ClientPasswordSecretKeyRef *struct {
					Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"client_password_secret_key_ref" json:"clientPasswordSecretKeyRef,omitempty"`
				ClientUsername              *string `tfsdk:"client_username" json:"clientUsername,omitempty"`
				DeleteDefaultAdmin          *bool   `tfsdk:"delete_default_admin" json:"deleteDefaultAdmin,omitempty"`
				Generate                    *bool   `tfsdk:"generate" json:"generate,omitempty"`
				MetricsPasswordSecretKeyRef *struct {
					Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"metrics_password_secret_key_ref" json:"metricsPasswordSecretKeyRef,omitempty"`
				MetricsUsername             *string `tfsdk:"metrics_username" json:"metricsUsername,omitempty"`
				MonitorMaxConnections       *int64  `tfsdk:"monitor_max_connections" json:"monitorMaxConnections,omitempty"`
				MonitorPasswordSecretKeyRef *struct {
					Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"monitor_password_secret_key_ref" json:"monitorPasswordSecretKeyRef,omitempty"`
				MonitorUsername            *string `tfsdk:"monitor_username" json:"monitorUsername,omitempty"`
				ServerMaxConnections       *int64  `tfsdk:"server_max_connections" json:"serverMaxConnections,omitempty"`
				ServerPasswordSecretKeyRef *struct {
					Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"server_password_secret_key_ref" json:"serverPasswordSecretKeyRef,omitempty"`
				ServerUsername           *string `tfsdk:"server_username" json:"serverUsername,omitempty"`
				SyncMaxConnections       *int64  `tfsdk:"sync_max_connections" json:"syncMaxConnections,omitempty"`
				SyncPasswordSecretKeyRef *struct {
					Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"sync_password_secret_key_ref" json:"syncPasswordSecretKeyRef,omitempty"`
				SyncUsername *string `tfsdk:"sync_username" json:"syncUsername,omitempty"`
			} `tfsdk:"auth" json:"auth,omitempty"`
			Config *struct {
				Params *map[string]string `tfsdk:"params" json:"params,omitempty"`
				Sync   *struct {
					Database *string `tfsdk:"database" json:"database,omitempty"`
					Interval *string `tfsdk:"interval" json:"interval,omitempty"`
					Timeout  *string `tfsdk:"timeout" json:"timeout,omitempty"`
				} `tfsdk:"sync" json:"sync,omitempty"`
				VolumeClaimTemplate *struct {
					AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
					Metadata    *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
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
					StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
			} `tfsdk:"config" json:"config,omitempty"`
			Connection *struct {
				HealthCheck *struct {
					Interval      *string `tfsdk:"interval" json:"interval,omitempty"`
					RetryInterval *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
				} `tfsdk:"health_check" json:"healthCheck,omitempty"`
				Params         *map[string]string `tfsdk:"params" json:"params,omitempty"`
				Port           *int64             `tfsdk:"port" json:"port,omitempty"`
				SecretName     *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
				SecretTemplate *struct {
					DatabaseKey *string `tfsdk:"database_key" json:"databaseKey,omitempty"`
					Format      *string `tfsdk:"format" json:"format,omitempty"`
					HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
					Key         *string `tfsdk:"key" json:"key,omitempty"`
					Metadata    *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
					PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
					UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
				} `tfsdk:"secret_template" json:"secretTemplate,omitempty"`
				ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
			} `tfsdk:"connection" json:"connection,omitempty"`
			Enabled              *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			GuiKubernetesService *struct {
				AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
				ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
				LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
				LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
				Metadata                      *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				SessionAffinity *string `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
				Type            *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"gui_kubernetes_service" json:"guiKubernetesService,omitempty"`
			Image             *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy   *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			KubernetesService *struct {
				AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
				ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
				LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
				LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
				Metadata                      *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				SessionAffinity *string `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
				Type            *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"kubernetes_service" json:"kubernetesService,omitempty"`
			Metrics *struct {
				Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				Exporter *struct {
					Affinity *struct {
						AntiAffinityEnabled *bool `tfsdk:"anti_affinity_enabled" json:"antiAffinityEnabled,omitempty"`
						PodAntiAffinity     *struct {
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
									TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
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
								TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
							} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
						} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					} `tfsdk:"affinity" json:"affinity,omitempty"`
					Image            *string `tfsdk:"image" json:"image,omitempty"`
					ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
					ImagePullSecrets *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
					NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					PodMetadata  *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
						Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
					PodSecurityContext *struct {
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
						SupplementalGroups       *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
						SupplementalGroupsPolicy *string   `tfsdk:"supplemental_groups_policy" json:"supplementalGroupsPolicy,omitempty"`
						Sysctls                  *[]struct {
							Name  *string `tfsdk:"name" json:"name,omitempty"`
							Value *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"sysctls" json:"sysctls,omitempty"`
						WindowsOptions *struct {
							GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
							GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
							HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
							RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
						} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
					} `tfsdk:"pod_security_context" json:"podSecurityContext,omitempty"`
					Port              *int64  `tfsdk:"port" json:"port,omitempty"`
					PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
					Resources         *struct {
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Tolerations *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"exporter" json:"exporter,omitempty"`
				ServiceMonitor *struct {
					Interval          *string `tfsdk:"interval" json:"interval,omitempty"`
					JobLabel          *string `tfsdk:"job_label" json:"jobLabel,omitempty"`
					PrometheusRelease *string `tfsdk:"prometheus_release" json:"prometheusRelease,omitempty"`
					ScrapeTimeout     *string `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
				} `tfsdk:"service_monitor" json:"serviceMonitor,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Monitor *struct {
				CooperativeMonitoring *string            `tfsdk:"cooperative_monitoring" json:"cooperativeMonitoring,omitempty"`
				Interval              *string            `tfsdk:"interval" json:"interval,omitempty"`
				Module                *string            `tfsdk:"module" json:"module,omitempty"`
				Name                  *string            `tfsdk:"name" json:"name,omitempty"`
				Params                *map[string]string `tfsdk:"params" json:"params,omitempty"`
				Suspend               *bool              `tfsdk:"suspend" json:"suspend,omitempty"`
			} `tfsdk:"monitor" json:"monitor,omitempty"`
			PodDisruptionBudget *struct {
				MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				MinAvailable   *string `tfsdk:"min_available" json:"minAvailable,omitempty"`
			} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
			Replicas        *int64  `tfsdk:"replicas" json:"replicas,omitempty"`
			RequeueInterval *string `tfsdk:"requeue_interval" json:"requeueInterval,omitempty"`
			Services        *[]struct {
				Listener *struct {
					Name     *string            `tfsdk:"name" json:"name,omitempty"`
					Params   *map[string]string `tfsdk:"params" json:"params,omitempty"`
					Port     *int64             `tfsdk:"port" json:"port,omitempty"`
					Protocol *string            `tfsdk:"protocol" json:"protocol,omitempty"`
					Suspend  *bool              `tfsdk:"suspend" json:"suspend,omitempty"`
				} `tfsdk:"listener" json:"listener,omitempty"`
				Name    *string            `tfsdk:"name" json:"name,omitempty"`
				Params  *map[string]string `tfsdk:"params" json:"params,omitempty"`
				Router  *string            `tfsdk:"router" json:"router,omitempty"`
				Suspend *bool              `tfsdk:"suspend" json:"suspend,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
			UpdateStrategy *struct {
				RollingUpdate *struct {
					MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					Partition      *int64  `tfsdk:"partition" json:"partition,omitempty"`
				} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
		} `tfsdk:"max_scale" json:"maxScale,omitempty"`
		MaxScaleRef *struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
		} `tfsdk:"max_scale_ref" json:"maxScaleRef,omitempty"`
		Metrics *struct {
			Enabled  *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Exporter *struct {
				Affinity *struct {
					AntiAffinityEnabled *bool `tfsdk:"anti_affinity_enabled" json:"antiAffinityEnabled,omitempty"`
					PodAntiAffinity     *struct {
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
								TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
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
							TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" json:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				Image            *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				ImagePullSecrets *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				PodMetadata  *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
				PodSecurityContext *struct {
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
					SupplementalGroups       *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
					SupplementalGroupsPolicy *string   `tfsdk:"supplemental_groups_policy" json:"supplementalGroupsPolicy,omitempty"`
					Sysctls                  *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"sysctls" json:"sysctls,omitempty"`
					WindowsOptions *struct {
						GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
						GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
						HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
						RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
				} `tfsdk:"pod_security_context" json:"podSecurityContext,omitempty"`
				Port              *int64  `tfsdk:"port" json:"port,omitempty"`
				PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
				Resources         *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Tolerations *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"exporter" json:"exporter,omitempty"`
			PasswordSecretKeyRef *struct {
				Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"password_secret_key_ref" json:"passwordSecretKeyRef,omitempty"`
			ServiceMonitor *struct {
				Interval          *string `tfsdk:"interval" json:"interval,omitempty"`
				JobLabel          *string `tfsdk:"job_label" json:"jobLabel,omitempty"`
				PrometheusRelease *string `tfsdk:"prometheus_release" json:"prometheusRelease,omitempty"`
				ScrapeTimeout     *string `tfsdk:"scrape_timeout" json:"scrapeTimeout,omitempty"`
			} `tfsdk:"service_monitor" json:"serviceMonitor,omitempty"`
			Username *string `tfsdk:"username" json:"username,omitempty"`
		} `tfsdk:"metrics" json:"metrics,omitempty"`
		MyCnf                *string `tfsdk:"my_cnf" json:"myCnf,omitempty"`
		MyCnfConfigMapKeyRef *struct {
			Key  *string `tfsdk:"key" json:"key,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"my_cnf_config_map_key_ref" json:"myCnfConfigMapKeyRef,omitempty"`
		NodeSelector             *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		PasswordHashSecretKeyRef *struct {
			Key  *string `tfsdk:"key" json:"key,omitempty"`
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"password_hash_secret_key_ref" json:"passwordHashSecretKeyRef,omitempty"`
		PasswordPlugin *struct {
			PluginArgSecretKeyRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"plugin_arg_secret_key_ref" json:"pluginArgSecretKeyRef,omitempty"`
			PluginNameSecretKeyRef *struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"plugin_name_secret_key_ref" json:"pluginNameSecretKeyRef,omitempty"`
		} `tfsdk:"password_plugin" json:"passwordPlugin,omitempty"`
		PasswordSecretKeyRef *struct {
			Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"password_secret_key_ref" json:"passwordSecretKeyRef,omitempty"`
		PodDisruptionBudget *struct {
			MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
			MinAvailable   *string `tfsdk:"min_available" json:"minAvailable,omitempty"`
		} `tfsdk:"pod_disruption_budget" json:"podDisruptionBudget,omitempty"`
		PodMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
		PodSecurityContext *struct {
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
			SupplementalGroups       *[]string `tfsdk:"supplemental_groups" json:"supplementalGroups,omitempty"`
			SupplementalGroupsPolicy *string   `tfsdk:"supplemental_groups_policy" json:"supplementalGroupsPolicy,omitempty"`
			Sysctls                  *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"sysctls" json:"sysctls,omitempty"`
			WindowsOptions *struct {
				GmsaCredentialSpec     *string `tfsdk:"gmsa_credential_spec" json:"gmsaCredentialSpec,omitempty"`
				GmsaCredentialSpecName *string `tfsdk:"gmsa_credential_spec_name" json:"gmsaCredentialSpecName,omitempty"`
				HostProcess            *bool   `tfsdk:"host_process" json:"hostProcess,omitempty"`
				RunAsUserName          *string `tfsdk:"run_as_user_name" json:"runAsUserName,omitempty"`
			} `tfsdk:"windows_options" json:"windowsOptions,omitempty"`
		} `tfsdk:"pod_security_context" json:"podSecurityContext,omitempty"`
		Port              *int64 `tfsdk:"port" json:"port,omitempty"`
		PrimaryConnection *struct {
			HealthCheck *struct {
				Interval      *string `tfsdk:"interval" json:"interval,omitempty"`
				RetryInterval *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			Params         *map[string]string `tfsdk:"params" json:"params,omitempty"`
			Port           *int64             `tfsdk:"port" json:"port,omitempty"`
			SecretName     *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
			SecretTemplate *struct {
				DatabaseKey *string `tfsdk:"database_key" json:"databaseKey,omitempty"`
				Format      *string `tfsdk:"format" json:"format,omitempty"`
				HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
				Key         *string `tfsdk:"key" json:"key,omitempty"`
				Metadata    *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
				PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
				UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			} `tfsdk:"secret_template" json:"secretTemplate,omitempty"`
			ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
		} `tfsdk:"primary_connection" json:"primaryConnection,omitempty"`
		PrimaryService *struct {
			AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
			ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
			LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
			Metadata                      *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			SessionAffinity *string `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
			Type            *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"primary_service" json:"primaryService,omitempty"`
		PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		ReadinessProbe    *struct {
			Exec *struct {
				Command *[]string `tfsdk:"command" json:"command,omitempty"`
			} `tfsdk:"exec" json:"exec,omitempty"`
			FailureThreshold *int64 `tfsdk:"failure_threshold" json:"failureThreshold,omitempty"`
			HttpGet          *struct {
				Host   *string `tfsdk:"host" json:"host,omitempty"`
				Path   *string `tfsdk:"path" json:"path,omitempty"`
				Port   *string `tfsdk:"port" json:"port,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
			} `tfsdk:"http_get" json:"httpGet,omitempty"`
			InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
			PeriodSeconds       *int64 `tfsdk:"period_seconds" json:"periodSeconds,omitempty"`
			SuccessThreshold    *int64 `tfsdk:"success_threshold" json:"successThreshold,omitempty"`
			TimeoutSeconds      *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
		} `tfsdk:"readiness_probe" json:"readinessProbe,omitempty"`
		Replicas                *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		ReplicasAllowEvenNumber *bool  `tfsdk:"replicas_allow_even_number" json:"replicasAllowEvenNumber,omitempty"`
		Replication             *struct {
			Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			Primary *struct {
				AutomaticFailover *bool  `tfsdk:"automatic_failover" json:"automaticFailover,omitempty"`
				PodIndex          *int64 `tfsdk:"pod_index" json:"podIndex,omitempty"`
			} `tfsdk:"primary" json:"primary,omitempty"`
			ProbesEnabled *bool `tfsdk:"probes_enabled" json:"probesEnabled,omitempty"`
			Replica       *struct {
				ConnectionRetries        *int64  `tfsdk:"connection_retries" json:"connectionRetries,omitempty"`
				ConnectionTimeout        *string `tfsdk:"connection_timeout" json:"connectionTimeout,omitempty"`
				Gtid                     *string `tfsdk:"gtid" json:"gtid,omitempty"`
				ReplPasswordSecretKeyRef *struct {
					Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"repl_password_secret_key_ref" json:"replPasswordSecretKeyRef,omitempty"`
				SyncTimeout *string `tfsdk:"sync_timeout" json:"syncTimeout,omitempty"`
				WaitPoint   *string `tfsdk:"wait_point" json:"waitPoint,omitempty"`
			} `tfsdk:"replica" json:"replica,omitempty"`
			SyncBinlog *bool `tfsdk:"sync_binlog" json:"syncBinlog,omitempty"`
		} `tfsdk:"replication" json:"replication,omitempty"`
		Resources *struct {
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		RootEmptyPassword        *bool `tfsdk:"root_empty_password" json:"rootEmptyPassword,omitempty"`
		RootPasswordSecretKeyRef *struct {
			Generate *bool   `tfsdk:"generate" json:"generate,omitempty"`
			Key      *string `tfsdk:"key" json:"key,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"root_password_secret_key_ref" json:"rootPasswordSecretKeyRef,omitempty"`
		SecondaryConnection *struct {
			HealthCheck *struct {
				Interval      *string `tfsdk:"interval" json:"interval,omitempty"`
				RetryInterval *string `tfsdk:"retry_interval" json:"retryInterval,omitempty"`
			} `tfsdk:"health_check" json:"healthCheck,omitempty"`
			Params         *map[string]string `tfsdk:"params" json:"params,omitempty"`
			Port           *int64             `tfsdk:"port" json:"port,omitempty"`
			SecretName     *string            `tfsdk:"secret_name" json:"secretName,omitempty"`
			SecretTemplate *struct {
				DatabaseKey *string `tfsdk:"database_key" json:"databaseKey,omitempty"`
				Format      *string `tfsdk:"format" json:"format,omitempty"`
				HostKey     *string `tfsdk:"host_key" json:"hostKey,omitempty"`
				Key         *string `tfsdk:"key" json:"key,omitempty"`
				Metadata    *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
				PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
				PortKey     *string `tfsdk:"port_key" json:"portKey,omitempty"`
				UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			} `tfsdk:"secret_template" json:"secretTemplate,omitempty"`
			ServiceName *string `tfsdk:"service_name" json:"serviceName,omitempty"`
		} `tfsdk:"secondary_connection" json:"secondaryConnection,omitempty"`
		SecondaryService *struct {
			AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
			ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
			LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
			Metadata                      *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			SessionAffinity *string `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
			Type            *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"secondary_service" json:"secondaryService,omitempty"`
		SecurityContext *struct {
			AllowPrivilegeEscalation *bool `tfsdk:"allow_privilege_escalation" json:"allowPrivilegeEscalation,omitempty"`
			Capabilities             *struct {
				Add  *[]string `tfsdk:"add" json:"add,omitempty"`
				Drop *[]string `tfsdk:"drop" json:"drop,omitempty"`
			} `tfsdk:"capabilities" json:"capabilities,omitempty"`
			Privileged             *bool  `tfsdk:"privileged" json:"privileged,omitempty"`
			ReadOnlyRootFilesystem *bool  `tfsdk:"read_only_root_filesystem" json:"readOnlyRootFilesystem,omitempty"`
			RunAsGroup             *int64 `tfsdk:"run_as_group" json:"runAsGroup,omitempty"`
			RunAsNonRoot           *bool  `tfsdk:"run_as_non_root" json:"runAsNonRoot,omitempty"`
			RunAsUser              *int64 `tfsdk:"run_as_user" json:"runAsUser,omitempty"`
		} `tfsdk:"security_context" json:"securityContext,omitempty"`
		Service *struct {
			AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
			ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
			LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
			LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
			Metadata                      *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			SessionAffinity *string `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
			Type            *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"service" json:"service,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		SidecarContainers  *[]struct {
			Args            *[]string `tfsdk:"args" json:"args,omitempty"`
			Command         *[]string `tfsdk:"command" json:"command,omitempty"`
			Image           *string   `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string   `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			Resources       *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			VolumeMounts *[]struct {
				MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath   *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"sidecar_containers" json:"sidecarContainers,omitempty"`
		Storage *struct {
			Ephemeral           *bool   `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
			ResizeInUseVolumes  *bool   `tfsdk:"resize_in_use_volumes" json:"resizeInUseVolumes,omitempty"`
			Size                *string `tfsdk:"size" json:"size,omitempty"`
			StorageClassName    *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			VolumeClaimTemplate *struct {
				AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
				Metadata    *struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				} `tfsdk:"metadata" json:"metadata,omitempty"`
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
				StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
			WaitForVolumeResize *bool `tfsdk:"wait_for_volume_resize" json:"waitForVolumeResize,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		Suspend     *bool   `tfsdk:"suspend" json:"suspend,omitempty"`
		TimeZone    *string `tfsdk:"time_zone" json:"timeZone,omitempty"`
		Tolerations *[]struct {
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
		UpdateStrategy *struct {
			AutoUpdateDataPlane *bool `tfsdk:"auto_update_data_plane" json:"autoUpdateDataPlane,omitempty"`
			RollingUpdate       *struct {
				MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
				Partition      *int64  `tfsdk:"partition" json:"partition,omitempty"`
			} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
			Type *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
		Username     *string `tfsdk:"username" json:"username,omitempty"`
		VolumeMounts *[]struct {
			MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
			SubPath   *string `tfsdk:"sub_path" json:"subPath,omitempty"`
		} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		Volumes *[]struct {
			Csi *struct {
				Driver               *string `tfsdk:"driver" json:"driver,omitempty"`
				FsType               *string `tfsdk:"fs_type" json:"fsType,omitempty"`
				NodePublishSecretRef *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"node_publish_secret_ref" json:"nodePublishSecretRef,omitempty"`
				ReadOnly         *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
				VolumeAttributes *map[string]string `tfsdk:"volume_attributes" json:"volumeAttributes,omitempty"`
			} `tfsdk:"csi" json:"csi,omitempty"`
			EmptyDir *struct {
				Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
				SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
			} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
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
		} `tfsdk:"volumes" json:"volumes,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *K8SMariadbComMariaDbV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_k8s_mariadb_com_maria_db_v1alpha1_manifest"
}

func (r *K8SMariadbComMariaDbV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "MariaDB is the Schema for the mariadbs API. It is used to define MariaDB clusters.",
		MarkdownDescription: "MariaDB is the Schema for the mariadbs API. It is used to define MariaDB clusters.",
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
				Description:         "MariaDBSpec defines the desired state of MariaDB",
				MarkdownDescription: "MariaDBSpec defines the desired state of MariaDB",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "Affinity to be used in the Pod.",
						MarkdownDescription: "Affinity to be used in the Pod.",
						Attributes: map[string]schema.Attribute{
							"anti_affinity_enabled": schema.BoolAttribute{
								Description:         "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
								MarkdownDescription: "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_anti_affinity": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
															MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

														"topology_key": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: true,
													Optional: false,
													Computed: false,
												},

												"weight": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
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

									"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
													MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

												"topology_key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

					"args": schema.ListAttribute{
						Description:         "Args to be used in the Container.",
						MarkdownDescription: "Args to be used in the Container.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bootstrap_from": schema.SingleNestedAttribute{
						Description:         "BootstrapFrom defines a source to bootstrap from.",
						MarkdownDescription: "BootstrapFrom defines a source to bootstrap from.",
						Attributes: map[string]schema.Attribute{
							"backup_ref": schema.SingleNestedAttribute{
								Description:         "BackupRef is a reference to a Backup object. It has priority over S3 and Volume.",
								MarkdownDescription: "BackupRef is a reference to a Backup object. It has priority over S3 and Volume.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

							"restore_job": schema.SingleNestedAttribute{
								Description:         "RestoreJob defines additional properties for the Job used to perform the Restore.",
								MarkdownDescription: "RestoreJob defines additional properties for the Job used to perform the Restore.",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "Affinity to be used in the Pod.",
										MarkdownDescription: "Affinity to be used in the Pod.",
										Attributes: map[string]schema.Attribute{
											"anti_affinity_enabled": schema.BoolAttribute{
												Description:         "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
												MarkdownDescription: "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_anti_affinity": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
												Attributes: map[string]schema.Attribute{
													"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"pod_affinity_term": schema.SingleNestedAttribute{
																	Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
																	MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
																			MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

																		"topology_key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"weight": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
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

													"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
																	MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

																"topology_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

									"args": schema.ListAttribute{
										Description:         "Args to be used in the Container.",
										MarkdownDescription: "Args to be used in the Container.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata defines additional metadata for the bootstrap Jobs.",
										MarkdownDescription: "Metadata defines additional metadata for the bootstrap Jobs.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"resources": schema.SingleNestedAttribute{
										Description:         "Resouces describes the compute resource requirements.",
										MarkdownDescription: "Resouces describes the compute resource requirements.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

							"s3": schema.SingleNestedAttribute{
								Description:         "S3 defines the configuration to restore backups from a S3 compatible storage. It has priority over Volume.",
								MarkdownDescription: "S3 defines the configuration to restore backups from a S3 compatible storage. It has priority over Volume.",
								Attributes: map[string]schema.Attribute{
									"access_key_id_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "AccessKeyIdSecretKeyRef is a reference to a Secret key containing the S3 access key id.",
										MarkdownDescription: "AccessKeyIdSecretKeyRef is a reference to a Secret key containing the S3 access key id.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"bucket": schema.StringAttribute{
										Description:         "Bucket is the name Name of the bucket to store backups.",
										MarkdownDescription: "Bucket is the name Name of the bucket to store backups.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"endpoint": schema.StringAttribute{
										Description:         "Endpoint is the S3 API endpoint without scheme.",
										MarkdownDescription: "Endpoint is the S3 API endpoint without scheme.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"prefix": schema.StringAttribute{
										Description:         "Prefix indicates a folder/subfolder in the bucket. For example: mariadb/ or mariadb/backups. A trailing slash '/' is added if not provided.",
										MarkdownDescription: "Prefix indicates a folder/subfolder in the bucket. For example: mariadb/ or mariadb/backups. A trailing slash '/' is added if not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"region": schema.StringAttribute{
										Description:         "Region is the S3 region name to use.",
										MarkdownDescription: "Region is the S3 region name to use.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_access_key_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "AccessKeyIdSecretKeyRef is a reference to a Secret key containing the S3 secret key.",
										MarkdownDescription: "AccessKeyIdSecretKeyRef is a reference to a Secret key containing the S3 secret key.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"session_token_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "SessionTokenSecretKeyRef is a reference to a Secret key containing the S3 session token.",
										MarkdownDescription: "SessionTokenSecretKeyRef is a reference to a Secret key containing the S3 session token.",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
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

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS provides the configuration required to establish TLS connections with S3.",
										MarkdownDescription: "TLS provides the configuration required to establish TLS connections with S3.",
										Attributes: map[string]schema.Attribute{
											"ca_secret_key_ref": schema.SingleNestedAttribute{
												Description:         "CASecretKeyRef is a reference to a Secret key containing a CA bundle in PEM format used to establish TLS connections with S3. By default, the system trust chain will be used, but you can use this field to add more CAs to the bundle.",
												MarkdownDescription: "CASecretKeyRef is a reference to a Secret key containing a CA bundle in PEM format used to establish TLS connections with S3. By default, the system trust chain will be used, but you can use this field to add more CAs to the bundle.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
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

											"enabled": schema.BoolAttribute{
												Description:         "Enabled is a flag to enable TLS.",
												MarkdownDescription: "Enabled is a flag to enable TLS.",
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

							"target_recovery_time": schema.StringAttribute{
								Description:         "TargetRecoveryTime is a RFC3339 (1970-01-01T00:00:00Z) date and time that defines the point in time recovery objective. It is used to determine the closest restoration source in time.",
								MarkdownDescription: "TargetRecoveryTime is a RFC3339 (1970-01-01T00:00:00Z) date and time that defines the point in time recovery objective. It is used to determine the closest restoration source in time.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									validators.DateTime64Validator(),
								},
							},

							"volume": schema.SingleNestedAttribute{
								Description:         "Volume is a Kubernetes Volume object that contains a backup.",
								MarkdownDescription: "Volume is a Kubernetes Volume object that contains a backup.",
								Attributes: map[string]schema.Attribute{
									"csi": schema.SingleNestedAttribute{
										Description:         "Represents a source location of a volume to mount, managed by an external CSI driver",
										MarkdownDescription: "Represents a source location of a volume to mount, managed by an external CSI driver",
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
														Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

									"empty_dir": schema.SingleNestedAttribute{
										Description:         "Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.",
										MarkdownDescription: "Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.",
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

									"nfs": schema.SingleNestedAttribute{
										Description:         "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",
										MarkdownDescription: "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",
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
										Description:         "PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).",
										MarkdownDescription: "PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).",
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

					"command": schema.ListAttribute{
						Description:         "Command to be used in the Container.",
						MarkdownDescription: "Command to be used in the Container.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"connection": schema.SingleNestedAttribute{
						Description:         "Connection defines a template to configure the general Connection object. This Connection provides the initial User access to the initial Database. It will make use of the Service to route network traffic to all Pods.",
						MarkdownDescription: "Connection defines a template to configure the general Connection object. This Connection provides the initial User access to the initial Database. It will make use of the Service to route network traffic to all Pods.",
						Attributes: map[string]schema.Attribute{
							"health_check": schema.SingleNestedAttribute{
								Description:         "HealthCheck to be used in the Connection.",
								MarkdownDescription: "HealthCheck to be used in the Connection.",
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Description:         "Interval used to perform health checks.",
										MarkdownDescription: "Interval used to perform health checks.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_interval": schema.StringAttribute{
										Description:         "RetryInterval is the interval used to perform health check retries.",
										MarkdownDescription: "RetryInterval is the interval used to perform health check retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"params": schema.MapAttribute{
								Description:         "Params to be used in the Connection.",
								MarkdownDescription: "Params to be used in the Connection.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
								MarkdownDescription: "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName to be used in the Connection.",
								MarkdownDescription: "SecretName to be used in the Connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_template": schema.SingleNestedAttribute{
								Description:         "SecretTemplate to be used in the Connection.",
								MarkdownDescription: "SecretTemplate to be used in the Connection.",
								Attributes: map[string]schema.Attribute{
									"database_key": schema.StringAttribute{
										Description:         "DatabaseKey to be used in the Secret.",
										MarkdownDescription: "DatabaseKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"format": schema.StringAttribute{
										Description:         "Format to be used in the Secret.",
										MarkdownDescription: "Format to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_key": schema.StringAttribute{
										Description:         "HostKey to be used in the Secret.",
										MarkdownDescription: "HostKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "Key to be used in the Secret.",
										MarkdownDescription: "Key to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata to be added to the Secret object.",
										MarkdownDescription: "Metadata to be added to the Secret object.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"password_key": schema.StringAttribute{
										Description:         "PasswordKey to be used in the Secret.",
										MarkdownDescription: "PasswordKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port_key": schema.StringAttribute{
										Description:         "PortKey to be used in the Secret.",
										MarkdownDescription: "PortKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username_key": schema.StringAttribute{
										Description:         "UsernameKey to be used in the Secret.",
										MarkdownDescription: "UsernameKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_name": schema.StringAttribute{
								Description:         "ServiceName to be used in the Connection.",
								MarkdownDescription: "ServiceName to be used in the Connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"database": schema.StringAttribute{
						Description:         "Database is the name of the initial Database.",
						MarkdownDescription: "Database is the name of the initial Database.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"env": schema.ListNestedAttribute{
						Description:         "Env represents the environment variables to be injected in a container.",
						MarkdownDescription: "Env represents the environment variables to be injected in a container.",
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
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value_from": schema.SingleNestedAttribute{
									Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.",
									MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.",
									Attributes: map[string]schema.Attribute{
										"config_map_key_ref": schema.SingleNestedAttribute{
											Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#configmapkeyselector-v1-core.",
											MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#configmapkeyselector-v1-core.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
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

										"field_ref": schema.SingleNestedAttribute{
											Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectfieldselector-v1-core.",
											MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectfieldselector-v1-core.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
											Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#secretkeyselector-v1-core.",
											MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#secretkeyselector-v1-core.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"env_from": schema.ListNestedAttribute{
						Description:         "EnvFrom represents the references (via ConfigMap and Secrets) to environment variables to be injected in the container.",
						MarkdownDescription: "EnvFrom represents the references (via ConfigMap and Secrets) to environment variables to be injected in the container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_map_ref": schema.SingleNestedAttribute{
									Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
									MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

								"prefix": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_ref": schema.SingleNestedAttribute{
									Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
									MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

					"galera": schema.SingleNestedAttribute{
						Description:         "Replication configures high availability via Galera.",
						MarkdownDescription: "Replication configures high availability via Galera.",
						Attributes: map[string]schema.Attribute{
							"agent": schema.SingleNestedAttribute{
								Description:         "GaleraAgent is a sidecar agent that co-operates with mariadb-operator.",
								MarkdownDescription: "GaleraAgent is a sidecar agent that co-operates with mariadb-operator.",
								Attributes: map[string]schema.Attribute{
									"args": schema.ListAttribute{
										Description:         "Args to be used in the Container.",
										MarkdownDescription: "Args to be used in the Container.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"basic_auth": schema.SingleNestedAttribute{
										Description:         "BasicAuth to be used by the agent container",
										MarkdownDescription: "BasicAuth to be used by the agent container",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Enabled is a flag to enable BasicAuth",
												MarkdownDescription: "Enabled is a flag to enable BasicAuth",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"password_secret_key_ref": schema.SingleNestedAttribute{
												Description:         "PasswordSecretKeyRef to be used for basic authentication",
												MarkdownDescription: "PasswordSecretKeyRef to be used for basic authentication",
												Attributes: map[string]schema.Attribute{
													"generate": schema.BoolAttribute{
														Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
														MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
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

											"username": schema.StringAttribute{
												Description:         "Username to be used for basic authentication",
												MarkdownDescription: "Username to be used for basic authentication",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"command": schema.ListAttribute{
										Description:         "Command to be used in the Container.",
										MarkdownDescription: "Command to be used in the Container.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env": schema.ListNestedAttribute{
										Description:         "Env represents the environment variables to be injected in a container.",
										MarkdownDescription: "Env represents the environment variables to be injected in a container.",
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
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value_from": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
															Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#configmapkeyselector-v1-core.",
															MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#configmapkeyselector-v1-core.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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

														"field_ref": schema.SingleNestedAttribute{
															Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectfieldselector-v1-core.",
															MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectfieldselector-v1-core.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#secretkeyselector-v1-core.",
															MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#secretkeyselector-v1-core.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"env_from": schema.ListNestedAttribute{
										Description:         "EnvFrom represents the references (via ConfigMap and Secrets) to environment variables to be injected in the container.",
										MarkdownDescription: "EnvFrom represents the references (via ConfigMap and Secrets) to environment variables to be injected in the container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"config_map_ref": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

									"graceful_shutdown_timeout": schema.StringAttribute{
										Description:         "GracefulShutdownTimeout is the time we give to the agent container in order to gracefully terminate in-flight requests.",
										MarkdownDescription: "GracefulShutdownTimeout is the time we give to the agent container in order to gracefully terminate in-flight requests.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image": schema.StringAttribute{
										Description:         "Image name to be used by the MariaDB instances. The supported format is '<image>:<tag>'.",
										MarkdownDescription: "Image name to be used by the MariaDB instances. The supported format is '<image>:<tag>'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
										MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
										},
									},

									"kubernetes_auth": schema.SingleNestedAttribute{
										Description:         "KubernetesAuth to be used by the agent container",
										MarkdownDescription: "KubernetesAuth to be used by the agent container",
										Attributes: map[string]schema.Attribute{
											"auth_delegator_role_name": schema.StringAttribute{
												Description:         "AuthDelegatorRoleName is the name of the ClusterRoleBinding that is associated with the 'system:auth-delegator' ClusterRole. It is necessary for creating TokenReview objects in order for the agent to validate the service account token.",
												MarkdownDescription: "AuthDelegatorRoleName is the name of the ClusterRoleBinding that is associated with the 'system:auth-delegator' ClusterRole. It is necessary for creating TokenReview objects in order for the agent to validate the service account token.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "Enabled is a flag to enable KubernetesAuth",
												MarkdownDescription: "Enabled is a flag to enable KubernetesAuth",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"liveness_probe": schema.SingleNestedAttribute{
										Description:         "LivenessProbe to be used in the Container.",
										MarkdownDescription: "LivenessProbe to be used in the Container.",
										Attributes: map[string]schema.Attribute{
											"exec": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
												Attributes: map[string]schema.Attribute{
													"command": schema.ListAttribute{
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

											"failure_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_get": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
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

													"port": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"scheme": schema.StringAttribute{
														Description:         "URIScheme identifies the scheme used for connection to a host for Get actions",
														MarkdownDescription: "URIScheme identifies the scheme used for connection to a host for Get actions",
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
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
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

									"port": schema.Int64Attribute{
										Description:         "Port where the agent will be listening for connections.",
										MarkdownDescription: "Port where the agent will be listening for connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"readiness_probe": schema.SingleNestedAttribute{
										Description:         "ReadinessProbe to be used in the Container.",
										MarkdownDescription: "ReadinessProbe to be used in the Container.",
										Attributes: map[string]schema.Attribute{
											"exec": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
												Attributes: map[string]schema.Attribute{
													"command": schema.ListAttribute{
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

											"failure_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_get": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
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

													"port": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"scheme": schema.StringAttribute{
														Description:         "URIScheme identifies the scheme used for connection to a host for Get actions",
														MarkdownDescription: "URIScheme identifies the scheme used for connection to a host for Get actions",
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
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
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

									"resources": schema.SingleNestedAttribute{
										Description:         "Resouces describes the compute resource requirements.",
										MarkdownDescription: "Resouces describes the compute resource requirements.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

									"security_context": schema.SingleNestedAttribute{
										Description:         "SecurityContext holds security configuration that will be applied to a container.",
										MarkdownDescription: "SecurityContext holds security configuration that will be applied to a container.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"capabilities": schema.SingleNestedAttribute{
												Description:         "Adds and removes POSIX capabilities from running containers.",
												MarkdownDescription: "Adds and removes POSIX capabilities from running containers.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_mounts": schema.ListNestedAttribute{
										Description:         "VolumeMounts to be used in the Container.",
										MarkdownDescription: "VolumeMounts to be used in the Container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"mount_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
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

							"available_when_donor": schema.BoolAttribute{
								Description:         "AvailableWhenDonor indicates whether a donor node should be responding to queries. It defaults to false.",
								MarkdownDescription: "AvailableWhenDonor indicates whether a donor node should be responding to queries. It defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"config": schema.SingleNestedAttribute{
								Description:         "GaleraConfig defines storage options for the Galera configuration files.",
								MarkdownDescription: "GaleraConfig defines storage options for the Galera configuration files.",
								Attributes: map[string]schema.Attribute{
									"reuse_storage_volume": schema.BoolAttribute{
										Description:         "ReuseStorageVolume indicates that storage volume used by MariaDB should be reused to store the Galera configuration files. It defaults to false, which implies that a dedicated volume for the Galera configuration files is provisioned.",
										MarkdownDescription: "ReuseStorageVolume indicates that storage volume used by MariaDB should be reused to store the Galera configuration files. It defaults to false, which implies that a dedicated volume for the Galera configuration files is provisioned.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_claim_template": schema.SingleNestedAttribute{
										Description:         "VolumeClaimTemplate is a template for the PVC that will contain the Galera configuration files shared between the InitContainer, Agent and MariaDB.",
										MarkdownDescription: "VolumeClaimTemplate is a template for the PVC that will contain the Galera configuration files shared between the InitContainer, Agent and MariaDB.",
										Attributes: map[string]schema.Attribute{
											"access_modes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata": schema.SingleNestedAttribute{
												Description:         "Metadata to be added to the PVC metadata.",
												MarkdownDescription: "Metadata to be added to the PVC metadata.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations to be added to children resources.",
														MarkdownDescription: "Annotations to be added to children resources.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Labels to be added to children resources.",
														MarkdownDescription: "Labels to be added to children resources.",
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

											"resources": schema.SingleNestedAttribute{
												Description:         "VolumeResourceRequirements describes the storage resource requirements for a volume.",
												MarkdownDescription: "VolumeResourceRequirements describes the storage resource requirements for a volume.",
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
												Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
												MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

							"enabled": schema.BoolAttribute{
								Description:         "Enabled is a flag to enable Galera.",
								MarkdownDescription: "Enabled is a flag to enable Galera.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"galera_lib_path": schema.StringAttribute{
								Description:         "GaleraLibPath is a path inside the MariaDB image to the wsrep provider plugin. It is defaulted if not provided. More info: https://galeracluster.com/library/documentation/mysql-wsrep-options.html#wsrep-provider.",
								MarkdownDescription: "GaleraLibPath is a path inside the MariaDB image to the wsrep provider plugin. It is defaulted if not provided. More info: https://galeracluster.com/library/documentation/mysql-wsrep-options.html#wsrep-provider.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"init_container": schema.SingleNestedAttribute{
								Description:         "InitContainer is an init container that runs in the MariaDB Pod and co-operates with mariadb-operator.",
								MarkdownDescription: "InitContainer is an init container that runs in the MariaDB Pod and co-operates with mariadb-operator.",
								Attributes: map[string]schema.Attribute{
									"args": schema.ListAttribute{
										Description:         "Args to be used in the Container.",
										MarkdownDescription: "Args to be used in the Container.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"command": schema.ListAttribute{
										Description:         "Command to be used in the Container.",
										MarkdownDescription: "Command to be used in the Container.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"env": schema.ListNestedAttribute{
										Description:         "Env represents the environment variables to be injected in a container.",
										MarkdownDescription: "Env represents the environment variables to be injected in a container.",
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
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value_from": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#envvarsource-v1-core.",
													Attributes: map[string]schema.Attribute{
														"config_map_key_ref": schema.SingleNestedAttribute{
															Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#configmapkeyselector-v1-core.",
															MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#configmapkeyselector-v1-core.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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

														"field_ref": schema.SingleNestedAttribute{
															Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectfieldselector-v1-core.",
															MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#objectfieldselector-v1-core.",
															Attributes: map[string]schema.Attribute{
																"api_version": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"field_path": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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
															Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#secretkeyselector-v1-core.",
															MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#secretkeyselector-v1-core.",
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"env_from": schema.ListNestedAttribute{
										Description:         "EnvFrom represents the references (via ConfigMap and Secrets) to environment variables to be injected in the container.",
										MarkdownDescription: "EnvFrom represents the references (via ConfigMap and Secrets) to environment variables to be injected in the container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"config_map_ref": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

												"prefix": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
													MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#localobjectreference-v1-core.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
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

									"image": schema.StringAttribute{
										Description:         "Image name to be used by the MariaDB instances. The supported format is '<image>:<tag>'.",
										MarkdownDescription: "Image name to be used by the MariaDB instances. The supported format is '<image>:<tag>'.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
										MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
										},
									},

									"liveness_probe": schema.SingleNestedAttribute{
										Description:         "LivenessProbe to be used in the Container.",
										MarkdownDescription: "LivenessProbe to be used in the Container.",
										Attributes: map[string]schema.Attribute{
											"exec": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
												Attributes: map[string]schema.Attribute{
													"command": schema.ListAttribute{
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

											"failure_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_get": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
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

													"port": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"scheme": schema.StringAttribute{
														Description:         "URIScheme identifies the scheme used for connection to a host for Get actions",
														MarkdownDescription: "URIScheme identifies the scheme used for connection to a host for Get actions",
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
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
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

									"readiness_probe": schema.SingleNestedAttribute{
										Description:         "ReadinessProbe to be used in the Container.",
										MarkdownDescription: "ReadinessProbe to be used in the Container.",
										Attributes: map[string]schema.Attribute{
											"exec": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
												Attributes: map[string]schema.Attribute{
													"command": schema.ListAttribute{
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

											"failure_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_get": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
												Attributes: map[string]schema.Attribute{
													"host": schema.StringAttribute{
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

													"port": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"scheme": schema.StringAttribute{
														Description:         "URIScheme identifies the scheme used for connection to a host for Get actions",
														MarkdownDescription: "URIScheme identifies the scheme used for connection to a host for Get actions",
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
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"period_seconds": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"success_threshold": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout_seconds": schema.Int64Attribute{
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

									"resources": schema.SingleNestedAttribute{
										Description:         "Resouces describes the compute resource requirements.",
										MarkdownDescription: "Resouces describes the compute resource requirements.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

									"security_context": schema.SingleNestedAttribute{
										Description:         "SecurityContext holds security configuration that will be applied to a container.",
										MarkdownDescription: "SecurityContext holds security configuration that will be applied to a container.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"capabilities": schema.SingleNestedAttribute{
												Description:         "Adds and removes POSIX capabilities from running containers.",
												MarkdownDescription: "Adds and removes POSIX capabilities from running containers.",
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_mounts": schema.ListNestedAttribute{
										Description:         "VolumeMounts to be used in the Container.",
										MarkdownDescription: "VolumeMounts to be used in the Container.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"mount_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
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

							"init_job": schema.SingleNestedAttribute{
								Description:         "InitJob defines a Job that co-operates with mariadb-operator by performing initialization tasks.",
								MarkdownDescription: "InitJob defines a Job that co-operates with mariadb-operator by performing initialization tasks.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata defines additional metadata for the Galera init Job.",
										MarkdownDescription: "Metadata defines additional metadata for the Galera init Job.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"resources": schema.SingleNestedAttribute{
										Description:         "Resouces describes the compute resource requirements.",
										MarkdownDescription: "Resouces describes the compute resource requirements.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

							"primary": schema.SingleNestedAttribute{
								Description:         "Primary is the Galera configuration for the primary node.",
								MarkdownDescription: "Primary is the Galera configuration for the primary node.",
								Attributes: map[string]schema.Attribute{
									"automatic_failover": schema.BoolAttribute{
										Description:         "AutomaticFailover indicates whether the operator should automatically update PodIndex to perform an automatic primary failover.",
										MarkdownDescription: "AutomaticFailover indicates whether the operator should automatically update PodIndex to perform an automatic primary failover.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_index": schema.Int64Attribute{
										Description:         "PodIndex is the StatefulSet index of the primary node. The user may change this field to perform a manual switchover.",
										MarkdownDescription: "PodIndex is the StatefulSet index of the primary node. The user may change this field to perform a manual switchover.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"provider_options": schema.MapAttribute{
								Description:         "ProviderOptions is map of Galera configuration parameters. More info: https://mariadb.com/kb/en/galera-cluster-system-variables/#wsrep_provider_options.",
								MarkdownDescription: "ProviderOptions is map of Galera configuration parameters. More info: https://mariadb.com/kb/en/galera-cluster-system-variables/#wsrep_provider_options.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"recovery": schema.SingleNestedAttribute{
								Description:         "GaleraRecovery is the recovery process performed by the operator whenever the Galera cluster is not healthy. More info: https://galeracluster.com/library/documentation/crash-recovery.html.",
								MarkdownDescription: "GaleraRecovery is the recovery process performed by the operator whenever the Galera cluster is not healthy. More info: https://galeracluster.com/library/documentation/crash-recovery.html.",
								Attributes: map[string]schema.Attribute{
									"cluster_bootstrap_timeout": schema.StringAttribute{
										Description:         "ClusterBootstrapTimeout is the time limit for bootstrapping a cluster. Once this timeout is reached, the Galera recovery state is reset and a new cluster bootstrap will be attempted.",
										MarkdownDescription: "ClusterBootstrapTimeout is the time limit for bootstrapping a cluster. Once this timeout is reached, the Galera recovery state is reset and a new cluster bootstrap will be attempted.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_downscale_timeout": schema.StringAttribute{
										Description:         "ClusterDownscaleTimeout represents the maximum duration for downscaling the cluster's StatefulSet during the recovery process.",
										MarkdownDescription: "ClusterDownscaleTimeout represents the maximum duration for downscaling the cluster's StatefulSet during the recovery process.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_healthy_timeout": schema.StringAttribute{
										Description:         "ClusterHealthyTimeout represents the duration at which a Galera cluster, that consistently failed health checks, is considered unhealthy, and consequently the Galera recovery process will be initiated by the operator.",
										MarkdownDescription: "ClusterHealthyTimeout represents the duration at which a Galera cluster, that consistently failed health checks, is considered unhealthy, and consequently the Galera recovery process will be initiated by the operator.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_monitor_interval": schema.StringAttribute{
										Description:         "ClusterMonitorInterval represents the interval used to monitor the Galera cluster health.",
										MarkdownDescription: "ClusterMonitorInterval represents the interval used to monitor the Galera cluster health.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cluster_upscale_timeout": schema.StringAttribute{
										Description:         "ClusterUpscaleTimeout represents the maximum duration for upscaling the cluster's StatefulSet during the recovery process.",
										MarkdownDescription: "ClusterUpscaleTimeout represents the maximum duration for upscaling the cluster's StatefulSet during the recovery process.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enabled is a flag to enable GaleraRecovery.",
										MarkdownDescription: "Enabled is a flag to enable GaleraRecovery.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"force_cluster_bootstrap_in_pod": schema.StringAttribute{
										Description:         "ForceClusterBootstrapInPod allows you to manually initiate the bootstrap process in a specific Pod. IMPORTANT: Use this option only in exceptional circumstances. Not selecting the Pod with the highest sequence number may result in data loss. IMPORTANT: Ensure you unset this field after completing the bootstrap to allow the operator to choose the appropriate Pod to bootstrap from in an event of cluster recovery.",
										MarkdownDescription: "ForceClusterBootstrapInPod allows you to manually initiate the bootstrap process in a specific Pod. IMPORTANT: Use this option only in exceptional circumstances. Not selecting the Pod with the highest sequence number may result in data loss. IMPORTANT: Ensure you unset this field after completing the bootstrap to allow the operator to choose the appropriate Pod to bootstrap from in an event of cluster recovery.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"job": schema.SingleNestedAttribute{
										Description:         "Job defines a Job that co-operates with mariadb-operator by performing the Galera cluster recovery .",
										MarkdownDescription: "Job defines a Job that co-operates with mariadb-operator by performing the Galera cluster recovery .",
										Attributes: map[string]schema.Attribute{
											"metadata": schema.SingleNestedAttribute{
												Description:         "Metadata defines additional metadata for the Galera recovery Jobs.",
												MarkdownDescription: "Metadata defines additional metadata for the Galera recovery Jobs.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations to be added to children resources.",
														MarkdownDescription: "Annotations to be added to children resources.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Labels to be added to children resources.",
														MarkdownDescription: "Labels to be added to children resources.",
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

											"pod_affinity": schema.BoolAttribute{
												Description:         "PodAffinity indicates whether the recovery Jobs should run in the same Node as the MariaDB Pods. It defaults to true.",
												MarkdownDescription: "PodAffinity indicates whether the recovery Jobs should run in the same Node as the MariaDB Pods. It defaults to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resouces describes the compute resource requirements.",
												MarkdownDescription: "Resouces describes the compute resource requirements.",
												Attributes: map[string]schema.Attribute{
													"limits": schema.MapAttribute{
														Description:         "ResourceList is a set of (resource name, quantity) pairs.",
														MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.MapAttribute{
														Description:         "ResourceList is a set of (resource name, quantity) pairs.",
														MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

									"min_cluster_size": schema.StringAttribute{
										Description:         "MinClusterSize is the minimum number of replicas to consider the cluster healthy. It can be either a number of replicas (1) or a percentage (50%). If Galera consistently reports less replicas than this value for the given 'ClusterHealthyTimeout' interval, a cluster recovery is iniated. It defaults to '1' replica, and it is highly recommendeded to keep this value at '1' in most cases. If set to more than one replica, the cluster recovery process may restart the healthy replicas as well.",
										MarkdownDescription: "MinClusterSize is the minimum number of replicas to consider the cluster healthy. It can be either a number of replicas (1) or a percentage (50%). If Galera consistently reports less replicas than this value for the given 'ClusterHealthyTimeout' interval, a cluster recovery is iniated. It defaults to '1' replica, and it is highly recommendeded to keep this value at '1' in most cases. If set to more than one replica, the cluster recovery process may restart the healthy replicas as well.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_recovery_timeout": schema.StringAttribute{
										Description:         "PodRecoveryTimeout is the time limit for recevorying the sequence of a Pod during the cluster recovery.",
										MarkdownDescription: "PodRecoveryTimeout is the time limit for recevorying the sequence of a Pod during the cluster recovery.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_sync_timeout": schema.StringAttribute{
										Description:         "PodSyncTimeout is the time limit for a Pod to join the cluster after having performed a cluster bootstrap during the cluster recovery.",
										MarkdownDescription: "PodSyncTimeout is the time limit for a Pod to join the cluster after having performed a cluster bootstrap during the cluster recovery.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replica_threads": schema.Int64Attribute{
								Description:         "ReplicaThreads is the number of replica threads used to apply Galera write sets in parallel. More info: https://mariadb.com/kb/en/galera-cluster-system-variables/#wsrep_slave_threads.",
								MarkdownDescription: "ReplicaThreads is the number of replica threads used to apply Galera write sets in parallel. More info: https://mariadb.com/kb/en/galera-cluster-system-variables/#wsrep_slave_threads.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sst": schema.StringAttribute{
								Description:         "SST is the Snapshot State Transfer used when new Pods join the cluster. More info: https://galeracluster.com/library/documentation/sst.html.",
								MarkdownDescription: "SST is the Snapshot State Transfer used when new Pods join the cluster. More info: https://galeracluster.com/library/documentation/sst.html.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("rsync", "mariabackup", "mysqldump"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": schema.StringAttribute{
						Description:         "Image name to be used by the MariaDB instances. The supported format is '<image>:<tag>'. Only MariaDB official images are supported.",
						MarkdownDescription: "Image name to be used by the MariaDB instances. The supported format is '<image>:<tag>'. Only MariaDB official images are supported.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
						MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
						},
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
						MarkdownDescription: "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
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

					"inherit_metadata": schema.SingleNestedAttribute{
						Description:         "InheritMetadata defines the metadata to be inherited by children resources.",
						MarkdownDescription: "InheritMetadata defines the metadata to be inherited by children resources.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations to be added to children resources.",
								MarkdownDescription: "Annotations to be added to children resources.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels to be added to children resources.",
								MarkdownDescription: "Labels to be added to children resources.",
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

					"init_containers": schema.ListNestedAttribute{
						Description:         "InitContainers to be used in the Pod.",
						MarkdownDescription: "InitContainers to be used in the Pod.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"args": schema.ListAttribute{
									Description:         "Args to be used in the Container.",
									MarkdownDescription: "Args to be used in the Container.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"command": schema.ListAttribute{
									Description:         "Command to be used in the Container.",
									MarkdownDescription: "Command to be used in the Container.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "Image name to be used by the container. The supported format is '<image>:<tag>'.",
									MarkdownDescription: "Image name to be used by the container. The supported format is '<image>:<tag>'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"image_pull_policy": schema.StringAttribute{
									Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
									MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
									},
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Resouces describes the compute resource requirements.",
									MarkdownDescription: "Resouces describes the compute resource requirements.",
									Attributes: map[string]schema.Attribute{
										"limits": schema.MapAttribute{
											Description:         "ResourceList is a set of (resource name, quantity) pairs.",
											MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"requests": schema.MapAttribute{
											Description:         "ResourceList is a set of (resource name, quantity) pairs.",
											MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

								"volume_mounts": schema.ListNestedAttribute{
									Description:         "VolumeMounts to be used in the Container.",
									MarkdownDescription: "VolumeMounts to be used in the Container.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"mount_path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
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

					"liveness_probe": schema.SingleNestedAttribute{
						Description:         "LivenessProbe to be used in the Container.",
						MarkdownDescription: "LivenessProbe to be used in the Container.",
						Attributes: map[string]schema.Attribute{
							"exec": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
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

							"failure_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_get": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
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

									"port": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "URIScheme identifies the scheme used for connection to a host for Get actions",
										MarkdownDescription: "URIScheme identifies the scheme used for connection to a host for Get actions",
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
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"period_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"success_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout_seconds": schema.Int64Attribute{
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

					"max_scale": schema.SingleNestedAttribute{
						Description:         "MaxScale is the MaxScale specification that defines the MaxScale resource to be used with the current MariaDB. When enabling this field, MaxScaleRef is automatically set.",
						MarkdownDescription: "MaxScale is the MaxScale specification that defines the MaxScale resource to be used with the current MariaDB. When enabling this field, MaxScaleRef is automatically set.",
						Attributes: map[string]schema.Attribute{
							"admin": schema.SingleNestedAttribute{
								Description:         "Admin configures the admin REST API and GUI.",
								MarkdownDescription: "Admin configures the admin REST API and GUI.",
								Attributes: map[string]schema.Attribute{
									"gui_enabled": schema.BoolAttribute{
										Description:         "GuiEnabled indicates whether the admin GUI should be enabled.",
										MarkdownDescription: "GuiEnabled indicates whether the admin GUI should be enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Port where the admin REST API and GUI will be exposed.",
										MarkdownDescription: "Port where the admin REST API and GUI will be exposed.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"auth": schema.SingleNestedAttribute{
								Description:         "Auth defines the credentials required for MaxScale to connect to MariaDB.",
								MarkdownDescription: "Auth defines the credentials required for MaxScale to connect to MariaDB.",
								Attributes: map[string]schema.Attribute{
									"admin_password_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "AdminPasswordSecretKeyRef is Secret key reference to the admin password to call the admin REST API. It is defaulted if not provided.",
										MarkdownDescription: "AdminPasswordSecretKeyRef is Secret key reference to the admin password to call the admin REST API. It is defaulted if not provided.",
										Attributes: map[string]schema.Attribute{
											"generate": schema.BoolAttribute{
												Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
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

									"admin_username": schema.StringAttribute{
										Description:         "AdminUsername is an admin username to call the admin REST API. It is defaulted if not provided.",
										MarkdownDescription: "AdminUsername is an admin username to call the admin REST API. It is defaulted if not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_max_connections": schema.Int64Attribute{
										Description:         "ClientMaxConnections defines the maximum number of connections that the client can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
										MarkdownDescription: "ClientMaxConnections defines the maximum number of connections that the client can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_password_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "ClientPasswordSecretKeyRef is Secret key reference to the password to connect to MaxScale. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										MarkdownDescription: "ClientPasswordSecretKeyRef is Secret key reference to the password to connect to MaxScale. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										Attributes: map[string]schema.Attribute{
											"generate": schema.BoolAttribute{
												Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
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

									"client_username": schema.StringAttribute{
										Description:         "ClientUsername is the user to connect to MaxScale. It is defaulted if not provided.",
										MarkdownDescription: "ClientUsername is the user to connect to MaxScale. It is defaulted if not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"delete_default_admin": schema.BoolAttribute{
										Description:         "DeleteDefaultAdmin determines whether the default admin user should be deleted after the initial configuration. If not provided, it defaults to true.",
										MarkdownDescription: "DeleteDefaultAdmin determines whether the default admin user should be deleted after the initial configuration. If not provided, it defaults to true.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"generate": schema.BoolAttribute{
										Description:         "Generate defies whether the operator should generate users and grants for MaxScale to work. It only supports MariaDBs specified via spec.mariaDbRef.",
										MarkdownDescription: "Generate defies whether the operator should generate users and grants for MaxScale to work. It only supports MariaDBs specified via spec.mariaDbRef.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metrics_password_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "MetricsPasswordSecretKeyRef is Secret key reference to the metrics password to call the admib REST API. It is defaulted if metrics are enabled. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										MarkdownDescription: "MetricsPasswordSecretKeyRef is Secret key reference to the metrics password to call the admib REST API. It is defaulted if metrics are enabled. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										Attributes: map[string]schema.Attribute{
											"generate": schema.BoolAttribute{
												Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
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

									"metrics_username": schema.StringAttribute{
										Description:         "MetricsUsername is an metrics username to call the REST API. It is defaulted if metrics are enabled.",
										MarkdownDescription: "MetricsUsername is an metrics username to call the REST API. It is defaulted if metrics are enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"monitor_max_connections": schema.Int64Attribute{
										Description:         "MonitorMaxConnections defines the maximum number of connections that the monitor can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
										MarkdownDescription: "MonitorMaxConnections defines the maximum number of connections that the monitor can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"monitor_password_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "MonitorPasswordSecretKeyRef is Secret key reference to the password used by MaxScale monitor to connect to MariaDB server. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										MarkdownDescription: "MonitorPasswordSecretKeyRef is Secret key reference to the password used by MaxScale monitor to connect to MariaDB server. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										Attributes: map[string]schema.Attribute{
											"generate": schema.BoolAttribute{
												Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
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

									"monitor_username": schema.StringAttribute{
										Description:         "MonitorUsername is the user used by MaxScale monitor to connect to MariaDB server. It is defaulted if not provided.",
										MarkdownDescription: "MonitorUsername is the user used by MaxScale monitor to connect to MariaDB server. It is defaulted if not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server_max_connections": schema.Int64Attribute{
										Description:         "ServerMaxConnections defines the maximum number of connections that the server can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
										MarkdownDescription: "ServerMaxConnections defines the maximum number of connections that the server can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"server_password_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "ServerPasswordSecretKeyRef is Secret key reference to the password used by MaxScale to connect to MariaDB server. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										MarkdownDescription: "ServerPasswordSecretKeyRef is Secret key reference to the password used by MaxScale to connect to MariaDB server. It is defaulted if not provided. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										Attributes: map[string]schema.Attribute{
											"generate": schema.BoolAttribute{
												Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
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

									"server_username": schema.StringAttribute{
										Description:         "ServerUsername is the user used by MaxScale to connect to MariaDB server. It is defaulted if not provided.",
										MarkdownDescription: "ServerUsername is the user used by MaxScale to connect to MariaDB server. It is defaulted if not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sync_max_connections": schema.Int64Attribute{
										Description:         "SyncMaxConnections defines the maximum number of connections that the sync can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
										MarkdownDescription: "SyncMaxConnections defines the maximum number of connections that the sync can establish. If HA is enabled, make sure to increase this value, as more MaxScale replicas implies more connections. It defaults to 30 times the number of MaxScale replicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sync_password_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "SyncPasswordSecretKeyRef is Secret key reference to the password used by MaxScale config to connect to MariaDB server. It is defaulted when HA is enabled. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										MarkdownDescription: "SyncPasswordSecretKeyRef is Secret key reference to the password used by MaxScale config to connect to MariaDB server. It is defaulted when HA is enabled. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
										Attributes: map[string]schema.Attribute{
											"generate": schema.BoolAttribute{
												Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
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

									"sync_username": schema.StringAttribute{
										Description:         "MonitoSyncUsernamerUsername is the user used by MaxScale config sync to connect to MariaDB server. It is defaulted when HA is enabled.",
										MarkdownDescription: "MonitoSyncUsernamerUsername is the user used by MaxScale config sync to connect to MariaDB server. It is defaulted when HA is enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"config": schema.SingleNestedAttribute{
								Description:         "Config defines the MaxScale configuration.",
								MarkdownDescription: "Config defines the MaxScale configuration.",
								Attributes: map[string]schema.Attribute{
									"params": schema.MapAttribute{
										Description:         "Params is a key value pair of parameters to be used in the MaxScale static configuration file. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#global-settings.",
										MarkdownDescription: "Params is a key value pair of parameters to be used in the MaxScale static configuration file. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#global-settings.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sync": schema.SingleNestedAttribute{
										Description:         "Sync defines how to replicate configuration across MaxScale replicas. It is defaulted when HA is enabled.",
										MarkdownDescription: "Sync defines how to replicate configuration across MaxScale replicas. It is defaulted when HA is enabled.",
										Attributes: map[string]schema.Attribute{
											"database": schema.StringAttribute{
												Description:         "Database is the MariaDB logical database where the 'maxscale_config' table will be created in order to persist and synchronize config changes. If not provided, it defaults to 'mysql'.",
												MarkdownDescription: "Database is the MariaDB logical database where the 'maxscale_config' table will be created in order to persist and synchronize config changes. If not provided, it defaults to 'mysql'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"interval": schema.StringAttribute{
												Description:         "Interval defines the config synchronization interval. It is defaulted if not provided.",
												MarkdownDescription: "Interval defines the config synchronization interval. It is defaulted if not provided.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"timeout": schema.StringAttribute{
												Description:         "Interval defines the config synchronization timeout. It is defaulted if not provided.",
												MarkdownDescription: "Interval defines the config synchronization timeout. It is defaulted if not provided.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_claim_template": schema.SingleNestedAttribute{
										Description:         "VolumeClaimTemplate provides a template to define the PVCs for storing MaxScale runtime configuration files. It is defaulted if not provided.",
										MarkdownDescription: "VolumeClaimTemplate provides a template to define the PVCs for storing MaxScale runtime configuration files. It is defaulted if not provided.",
										Attributes: map[string]schema.Attribute{
											"access_modes": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata": schema.SingleNestedAttribute{
												Description:         "Metadata to be added to the PVC metadata.",
												MarkdownDescription: "Metadata to be added to the PVC metadata.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations to be added to children resources.",
														MarkdownDescription: "Annotations to be added to children resources.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Labels to be added to children resources.",
														MarkdownDescription: "Labels to be added to children resources.",
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

											"resources": schema.SingleNestedAttribute{
												Description:         "VolumeResourceRequirements describes the storage resource requirements for a volume.",
												MarkdownDescription: "VolumeResourceRequirements describes the storage resource requirements for a volume.",
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
												Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
												MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

							"connection": schema.SingleNestedAttribute{
								Description:         "Connection provides a template to define the Connection for MaxScale.",
								MarkdownDescription: "Connection provides a template to define the Connection for MaxScale.",
								Attributes: map[string]schema.Attribute{
									"health_check": schema.SingleNestedAttribute{
										Description:         "HealthCheck to be used in the Connection.",
										MarkdownDescription: "HealthCheck to be used in the Connection.",
										Attributes: map[string]schema.Attribute{
											"interval": schema.StringAttribute{
												Description:         "Interval used to perform health checks.",
												MarkdownDescription: "Interval used to perform health checks.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retry_interval": schema.StringAttribute{
												Description:         "RetryInterval is the interval used to perform health check retries.",
												MarkdownDescription: "RetryInterval is the interval used to perform health check retries.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"params": schema.MapAttribute{
										Description:         "Params to be used in the Connection.",
										MarkdownDescription: "Params to be used in the Connection.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
										MarkdownDescription: "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_name": schema.StringAttribute{
										Description:         "SecretName to be used in the Connection.",
										MarkdownDescription: "SecretName to be used in the Connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret_template": schema.SingleNestedAttribute{
										Description:         "SecretTemplate to be used in the Connection.",
										MarkdownDescription: "SecretTemplate to be used in the Connection.",
										Attributes: map[string]schema.Attribute{
											"database_key": schema.StringAttribute{
												Description:         "DatabaseKey to be used in the Secret.",
												MarkdownDescription: "DatabaseKey to be used in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"format": schema.StringAttribute{
												Description:         "Format to be used in the Secret.",
												MarkdownDescription: "Format to be used in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host_key": schema.StringAttribute{
												Description:         "HostKey to be used in the Secret.",
												MarkdownDescription: "HostKey to be used in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "Key to be used in the Secret.",
												MarkdownDescription: "Key to be used in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"metadata": schema.SingleNestedAttribute{
												Description:         "Metadata to be added to the Secret object.",
												MarkdownDescription: "Metadata to be added to the Secret object.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations to be added to children resources.",
														MarkdownDescription: "Annotations to be added to children resources.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Labels to be added to children resources.",
														MarkdownDescription: "Labels to be added to children resources.",
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

											"password_key": schema.StringAttribute{
												Description:         "PasswordKey to be used in the Secret.",
												MarkdownDescription: "PasswordKey to be used in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port_key": schema.StringAttribute{
												Description:         "PortKey to be used in the Secret.",
												MarkdownDescription: "PortKey to be used in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"username_key": schema.StringAttribute{
												Description:         "UsernameKey to be used in the Secret.",
												MarkdownDescription: "UsernameKey to be used in the Secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"service_name": schema.StringAttribute{
										Description:         "ServiceName to be used in the Connection.",
										MarkdownDescription: "ServiceName to be used in the Connection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Enabled is a flag to enable a MaxScale instance to be used with the current MariaDB.",
								MarkdownDescription: "Enabled is a flag to enable a MaxScale instance to be used with the current MariaDB.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gui_kubernetes_service": schema.SingleNestedAttribute{
								Description:         "GuiKubernetesService define a template for a Kubernetes Service object to connect to MaxScale's GUI.",
								MarkdownDescription: "GuiKubernetesService define a template for a Kubernetes Service object to connect to MaxScale's GUI.",
								Attributes: map[string]schema.Attribute{
									"allocate_load_balancer_node_ports": schema.BoolAttribute{
										Description:         "AllocateLoadBalancerNodePorts Service field.",
										MarkdownDescription: "AllocateLoadBalancerNodePorts Service field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_traffic_policy": schema.StringAttribute{
										Description:         "ExternalTrafficPolicy Service field.",
										MarkdownDescription: "ExternalTrafficPolicy Service field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"load_balancer_ip": schema.StringAttribute{
										Description:         "LoadBalancerIP Service field.",
										MarkdownDescription: "LoadBalancerIP Service field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"load_balancer_source_ranges": schema.ListAttribute{
										Description:         "LoadBalancerSourceRanges Service field.",
										MarkdownDescription: "LoadBalancerSourceRanges Service field.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata to be added to the Service metadata.",
										MarkdownDescription: "Metadata to be added to the Service metadata.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"session_affinity": schema.StringAttribute{
										Description:         "SessionAffinity Service field.",
										MarkdownDescription: "SessionAffinity Service field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
										MarkdownDescription: "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
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

							"image": schema.StringAttribute{
								Description:         "Image name to be used by the MaxScale instances. The supported format is '<image>:<tag>'. Only MariaDB official images are supported.",
								MarkdownDescription: "Image name to be used by the MaxScale instances. The supported format is '<image>:<tag>'. Only MariaDB official images are supported.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
								MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
								},
							},

							"kubernetes_service": schema.SingleNestedAttribute{
								Description:         "KubernetesService defines a template for a Kubernetes Service object to connect to MaxScale.",
								MarkdownDescription: "KubernetesService defines a template for a Kubernetes Service object to connect to MaxScale.",
								Attributes: map[string]schema.Attribute{
									"allocate_load_balancer_node_ports": schema.BoolAttribute{
										Description:         "AllocateLoadBalancerNodePorts Service field.",
										MarkdownDescription: "AllocateLoadBalancerNodePorts Service field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_traffic_policy": schema.StringAttribute{
										Description:         "ExternalTrafficPolicy Service field.",
										MarkdownDescription: "ExternalTrafficPolicy Service field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"load_balancer_ip": schema.StringAttribute{
										Description:         "LoadBalancerIP Service field.",
										MarkdownDescription: "LoadBalancerIP Service field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"load_balancer_source_ranges": schema.ListAttribute{
										Description:         "LoadBalancerSourceRanges Service field.",
										MarkdownDescription: "LoadBalancerSourceRanges Service field.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata to be added to the Service metadata.",
										MarkdownDescription: "Metadata to be added to the Service metadata.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"session_affinity": schema.StringAttribute{
										Description:         "SessionAffinity Service field.",
										MarkdownDescription: "SessionAffinity Service field.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
										MarkdownDescription: "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
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

							"metrics": schema.SingleNestedAttribute{
								Description:         "Metrics configures metrics and how to scrape them.",
								MarkdownDescription: "Metrics configures metrics and how to scrape them.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enabled is a flag to enable Metrics",
										MarkdownDescription: "Enabled is a flag to enable Metrics",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"exporter": schema.SingleNestedAttribute{
										Description:         "Exporter defines the metrics exporter container.",
										MarkdownDescription: "Exporter defines the metrics exporter container.",
										Attributes: map[string]schema.Attribute{
											"affinity": schema.SingleNestedAttribute{
												Description:         "Affinity to be used in the Pod.",
												MarkdownDescription: "Affinity to be used in the Pod.",
												Attributes: map[string]schema.Attribute{
													"anti_affinity_enabled": schema.BoolAttribute{
														Description:         "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
														MarkdownDescription: "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"pod_anti_affinity": schema.SingleNestedAttribute{
														Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
														MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
														Attributes: map[string]schema.Attribute{
															"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"pod_affinity_term": schema.SingleNestedAttribute{
																			Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
																			MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
																			Attributes: map[string]schema.Attribute{
																				"label_selector": schema.SingleNestedAttribute{
																					Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
																					MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

																				"topology_key": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"weight": schema.Int64Attribute{
																			Description:         "",
																			MarkdownDescription: "",
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

															"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
																			MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

																		"topology_key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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

											"image": schema.StringAttribute{
												Description:         "Image name to be used as metrics exporter. The supported format is '<image>:<tag>'. Only mysqld-exporter >= v0.15.0 is supported: https://github.com/prometheus/mysqld_exporter",
												MarkdownDescription: "Image name to be used as metrics exporter. The supported format is '<image>:<tag>'. Only mysqld-exporter >= v0.15.0 is supported: https://github.com/prometheus/mysqld_exporter",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"image_pull_policy": schema.StringAttribute{
												Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
												MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
												},
											},

											"image_pull_secrets": schema.ListNestedAttribute{
												Description:         "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
												MarkdownDescription: "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
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

											"node_selector": schema.MapAttribute{
												Description:         "NodeSelector to be used in the Pod.",
												MarkdownDescription: "NodeSelector to be used in the Pod.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_metadata": schema.SingleNestedAttribute{
												Description:         "PodMetadata defines extra metadata for the Pod.",
												MarkdownDescription: "PodMetadata defines extra metadata for the Pod.",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Annotations to be added to children resources.",
														MarkdownDescription: "Annotations to be added to children resources.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Labels to be added to children resources.",
														MarkdownDescription: "Labels to be added to children resources.",
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

											"pod_security_context": schema.SingleNestedAttribute{
												Description:         "SecurityContext holds pod-level security attributes and common container settings.",
												MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.",
												Attributes: map[string]schema.Attribute{
													"app_armor_profile": schema.SingleNestedAttribute{
														Description:         "appArmorProfile is the AppArmor options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "appArmorProfile is the AppArmor options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
														Attributes: map[string]schema.Attribute{
															"localhost_profile": schema.StringAttribute{
																Description:         "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
																MarkdownDescription: "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"type": schema.StringAttribute{
																Description:         "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
																MarkdownDescription: "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

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
														Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID and fsGroup (if specified). If the SupplementalGroupsPolicy feature is enabled, the supplementalGroupsPolicy field determines whether these are in addition to or instead of any group memberships defined in the container image. If unspecified, no additional groups are added, though group memberships defined in the container image may still be used, depending on the supplementalGroupsPolicy field. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID and fsGroup (if specified). If the SupplementalGroupsPolicy feature is enabled, the supplementalGroupsPolicy field determines whether these are in addition to or instead of any group memberships defined in the container image. If unspecified, no additional groups are added, though group memberships defined in the container image may still be used, depending on the supplementalGroupsPolicy field. Note that this field cannot be set when spec.os.name is windows.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"supplemental_groups_policy": schema.StringAttribute{
														Description:         "Defines how supplemental groups of the first container processes are calculated. Valid values are 'Merge' and 'Strict'. If not specified, 'Merge' is used. (Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled and the container runtime must implement support for this feature. Note that this field cannot be set when spec.os.name is windows.",
														MarkdownDescription: "Defines how supplemental groups of the first container processes are calculated. Valid values are 'Merge' and 'Strict'. If not specified, 'Merge' is used. (Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled and the container runtime must implement support for this feature. Note that this field cannot be set when spec.os.name is windows.",
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

											"port": schema.Int64Attribute{
												Description:         "Port where the exporter will be listening for connections.",
												MarkdownDescription: "Port where the exporter will be listening for connections.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"priority_class_name": schema.StringAttribute{
												Description:         "PriorityClassName to be used in the Pod.",
												MarkdownDescription: "PriorityClassName to be used in the Pod.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Resouces describes the compute resource requirements.",
												MarkdownDescription: "Resouces describes the compute resource requirements.",
												Attributes: map[string]schema.Attribute{
													"limits": schema.MapAttribute{
														Description:         "ResourceList is a set of (resource name, quantity) pairs.",
														MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.MapAttribute{
														Description:         "ResourceList is a set of (resource name, quantity) pairs.",
														MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

											"tolerations": schema.ListNestedAttribute{
												Description:         "Tolerations to be used in the Pod.",
												MarkdownDescription: "Tolerations to be used in the Pod.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"effect": schema.StringAttribute{
															Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
															MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
															MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
															MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"toleration_seconds": schema.Int64Attribute{
															Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
															MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
															MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

									"service_monitor": schema.SingleNestedAttribute{
										Description:         "ServiceMonitor defines the ServiceMonior object.",
										MarkdownDescription: "ServiceMonitor defines the ServiceMonior object.",
										Attributes: map[string]schema.Attribute{
											"interval": schema.StringAttribute{
												Description:         "Interval for scraping metrics.",
												MarkdownDescription: "Interval for scraping metrics.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"job_label": schema.StringAttribute{
												Description:         "JobLabel to add to the ServiceMonitor object.",
												MarkdownDescription: "JobLabel to add to the ServiceMonitor object.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"prometheus_release": schema.StringAttribute{
												Description:         "PrometheusRelease is the release label to add to the ServiceMonitor object.",
												MarkdownDescription: "PrometheusRelease is the release label to add to the ServiceMonitor object.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"scrape_timeout": schema.StringAttribute{
												Description:         "ScrapeTimeout defines the timeout for scraping metrics.",
												MarkdownDescription: "ScrapeTimeout defines the timeout for scraping metrics.",
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

							"monitor": schema.SingleNestedAttribute{
								Description:         "Monitor monitors MariaDB server instances.",
								MarkdownDescription: "Monitor monitors MariaDB server instances.",
								Attributes: map[string]schema.Attribute{
									"cooperative_monitoring": schema.StringAttribute{
										Description:         "CooperativeMonitoring enables coordination between multiple MaxScale instances running monitors. It is defaulted when HA is enabled.",
										MarkdownDescription: "CooperativeMonitoring enables coordination between multiple MaxScale instances running monitors. It is defaulted when HA is enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("majority_of_all", "majority_of_running"),
										},
									},

									"interval": schema.StringAttribute{
										Description:         "Interval used to monitor MariaDB servers. It is defaulted if not provided.",
										MarkdownDescription: "Interval used to monitor MariaDB servers. It is defaulted if not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"module": schema.StringAttribute{
										Description:         "Module is the module to use to monitor MariaDB servers. It is mandatory when no MariaDB reference is provided.",
										MarkdownDescription: "Module is the module to use to monitor MariaDB servers. It is mandatory when no MariaDB reference is provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "Name is the identifier of the monitor. It is defaulted if not provided.",
										MarkdownDescription: "Name is the identifier of the monitor. It is defaulted if not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"params": schema.MapAttribute{
										Description:         "Params defines extra parameters to pass to the monitor. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-common-monitor-parameters/. Monitor specific parameter are also suported: https://mariadb.com/kb/en/mariadb-maxscale-2308-galera-monitor/#galera-monitor-optional-parameters. https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-monitor/#configuration.",
										MarkdownDescription: "Params defines extra parameters to pass to the monitor. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-common-monitor-parameters/. Monitor specific parameter are also suported: https://mariadb.com/kb/en/mariadb-maxscale-2308-galera-monitor/#galera-monitor-optional-parameters. https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-monitor/#configuration.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"suspend": schema.BoolAttribute{
										Description:         "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
										MarkdownDescription: "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_disruption_budget": schema.SingleNestedAttribute{
								Description:         "PodDisruptionBudget defines the budget for replica availability.",
								MarkdownDescription: "PodDisruptionBudget defines the budget for replica availability.",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "MaxUnavailable defines the number of maximum unavailable Pods.",
										MarkdownDescription: "MaxUnavailable defines the number of maximum unavailable Pods.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"min_available": schema.StringAttribute{
										Description:         "MinAvailable defines the number of minimum available Pods.",
										MarkdownDescription: "MinAvailable defines the number of minimum available Pods.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Replicas indicates the number of desired instances.",
								MarkdownDescription: "Replicas indicates the number of desired instances.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requeue_interval": schema.StringAttribute{
								Description:         "RequeueInterval is used to perform requeue reconciliations.",
								MarkdownDescription: "RequeueInterval is used to perform requeue reconciliations.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"services": schema.ListNestedAttribute{
								Description:         "Services define how the traffic is forwarded to the MariaDB servers.",
								MarkdownDescription: "Services define how the traffic is forwarded to the MariaDB servers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"listener": schema.SingleNestedAttribute{
											Description:         "MaxScaleListener defines how the MaxScale server will listen for connections.",
											MarkdownDescription: "MaxScaleListener defines how the MaxScale server will listen for connections.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name is the identifier of the listener. It is defaulted if not provided",
													MarkdownDescription: "Name is the identifier of the listener. It is defaulted if not provided",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"params": schema.MapAttribute{
													Description:         "Params defines extra parameters to pass to the listener. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#listener_1.",
													MarkdownDescription: "Params defines extra parameters to pass to the listener. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#listener_1.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"port": schema.Int64Attribute{
													Description:         "Port is the network port where the MaxScale server will listen.",
													MarkdownDescription: "Port is the network port where the MaxScale server will listen.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"protocol": schema.StringAttribute{
													Description:         "Protocol is the MaxScale protocol to use when communicating with the client. If not provided, it defaults to MariaDBProtocol.",
													MarkdownDescription: "Protocol is the MaxScale protocol to use when communicating with the client. If not provided, it defaults to MariaDBProtocol.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"suspend": schema.BoolAttribute{
													Description:         "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
													MarkdownDescription: "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: true,
											Optional: false,
											Computed: false,
										},

										"name": schema.StringAttribute{
											Description:         "Name is the identifier of the MaxScale service.",
											MarkdownDescription: "Name is the identifier of the MaxScale service.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"params": schema.MapAttribute{
											Description:         "Params defines extra parameters to pass to the service. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#service_1. Router specific parameter are also suported: https://mariadb.com/kb/en/mariadb-maxscale-2308-readwritesplit/#configuration. https://mariadb.com/kb/en/mariadb-maxscale-2308-readconnroute/#configuration.",
											MarkdownDescription: "Params defines extra parameters to pass to the service. Any parameter supported by MaxScale may be specified here. See reference: https://mariadb.com/kb/en/mariadb-maxscale-2308-mariadb-maxscale-configuration-guide/#service_1. Router specific parameter are also suported: https://mariadb.com/kb/en/mariadb-maxscale-2308-readwritesplit/#configuration. https://mariadb.com/kb/en/mariadb-maxscale-2308-readconnroute/#configuration.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"router": schema.StringAttribute{
											Description:         "Router is the type of router to use.",
											MarkdownDescription: "Router is the type of router to use.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("readwritesplit", "readconnroute"),
											},
										},

										"suspend": schema.BoolAttribute{
											Description:         "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
											MarkdownDescription: "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
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

							"update_strategy": schema.SingleNestedAttribute{
								Description:         "UpdateStrategy defines the update strategy for the StatefulSet object.",
								MarkdownDescription: "UpdateStrategy defines the update strategy for the StatefulSet object.",
								Attributes: map[string]schema.Attribute{
									"rolling_update": schema.SingleNestedAttribute{
										Description:         "RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.",
										MarkdownDescription: "RollingUpdate is used to communicate parameters when Type is RollingUpdateStatefulSetStrategyType.",
										Attributes: map[string]schema.Attribute{
											"max_unavailable": schema.StringAttribute{
												Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0. Defaults to 1. This field is alpha-level and is only honored by servers that enable the MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it will be counted towards MaxUnavailable.",
												MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0. Defaults to 1. This field is alpha-level and is only honored by servers that enable the MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it will be counted towards MaxUnavailable.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"partition": schema.Int64Attribute{
												Description:         "Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0.",
												MarkdownDescription: "Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.",
										MarkdownDescription: "Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.",
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

					"max_scale_ref": schema.SingleNestedAttribute{
						Description:         "MaxScaleRef is a reference to a MaxScale resource to be used with the current MariaDB. Providing this field implies delegating high availability tasks such as primary failover to MaxScale.",
						MarkdownDescription: "MaxScaleRef is a reference to a MaxScale resource to be used with the current MariaDB. Providing this field implies delegating high availability tasks such as primary failover to MaxScale.",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespace": schema.StringAttribute{
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

					"metrics": schema.SingleNestedAttribute{
						Description:         "Metrics configures metrics and how to scrape them.",
						MarkdownDescription: "Metrics configures metrics and how to scrape them.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled is a flag to enable Metrics",
								MarkdownDescription: "Enabled is a flag to enable Metrics",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"exporter": schema.SingleNestedAttribute{
								Description:         "Exporter defines the metrics exporter container.",
								MarkdownDescription: "Exporter defines the metrics exporter container.",
								Attributes: map[string]schema.Attribute{
									"affinity": schema.SingleNestedAttribute{
										Description:         "Affinity to be used in the Pod.",
										MarkdownDescription: "Affinity to be used in the Pod.",
										Attributes: map[string]schema.Attribute{
											"anti_affinity_enabled": schema.BoolAttribute{
												Description:         "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
												MarkdownDescription: "AntiAffinityEnabled configures PodAntiAffinity so each Pod is scheduled in a different Node, enabling HA. Make sure you have at least as many Nodes available as the replicas to not end up with unscheduled Pods.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"pod_anti_affinity": schema.SingleNestedAttribute{
												Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
												MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podantiaffinity-v1-core.",
												Attributes: map[string]schema.Attribute{
													"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"pod_affinity_term": schema.SingleNestedAttribute{
																	Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
																	MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#podaffinityterm-v1-core.",
																	Attributes: map[string]schema.Attribute{
																		"label_selector": schema.SingleNestedAttribute{
																			Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
																			MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

																		"topology_key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"weight": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
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

													"required_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
																	MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

																"topology_key": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

									"image": schema.StringAttribute{
										Description:         "Image name to be used as metrics exporter. The supported format is '<image>:<tag>'. Only mysqld-exporter >= v0.15.0 is supported: https://github.com/prometheus/mysqld_exporter",
										MarkdownDescription: "Image name to be used as metrics exporter. The supported format is '<image>:<tag>'. Only mysqld-exporter >= v0.15.0 is supported: https://github.com/prometheus/mysqld_exporter",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_policy": schema.StringAttribute{
										Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
										MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
										},
									},

									"image_pull_secrets": schema.ListNestedAttribute{
										Description:         "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
										MarkdownDescription: "ImagePullSecrets is the list of pull Secrets to be used to pull the image.",
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

									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector to be used in the Pod.",
										MarkdownDescription: "NodeSelector to be used in the Pod.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_metadata": schema.SingleNestedAttribute{
										Description:         "PodMetadata defines extra metadata for the Pod.",
										MarkdownDescription: "PodMetadata defines extra metadata for the Pod.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"pod_security_context": schema.SingleNestedAttribute{
										Description:         "SecurityContext holds pod-level security attributes and common container settings.",
										MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.",
										Attributes: map[string]schema.Attribute{
											"app_armor_profile": schema.SingleNestedAttribute{
												Description:         "appArmorProfile is the AppArmor options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "appArmorProfile is the AppArmor options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
														MarkdownDescription: "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
														MarkdownDescription: "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

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
												Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID and fsGroup (if specified). If the SupplementalGroupsPolicy feature is enabled, the supplementalGroupsPolicy field determines whether these are in addition to or instead of any group memberships defined in the container image. If unspecified, no additional groups are added, though group memberships defined in the container image may still be used, depending on the supplementalGroupsPolicy field. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID and fsGroup (if specified). If the SupplementalGroupsPolicy feature is enabled, the supplementalGroupsPolicy field determines whether these are in addition to or instead of any group memberships defined in the container image. If unspecified, no additional groups are added, though group memberships defined in the container image may still be used, depending on the supplementalGroupsPolicy field. Note that this field cannot be set when spec.os.name is windows.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"supplemental_groups_policy": schema.StringAttribute{
												Description:         "Defines how supplemental groups of the first container processes are calculated. Valid values are 'Merge' and 'Strict'. If not specified, 'Merge' is used. (Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled and the container runtime must implement support for this feature. Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Defines how supplemental groups of the first container processes are calculated. Valid values are 'Merge' and 'Strict'. If not specified, 'Merge' is used. (Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled and the container runtime must implement support for this feature. Note that this field cannot be set when spec.os.name is windows.",
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

									"port": schema.Int64Attribute{
										Description:         "Port where the exporter will be listening for connections.",
										MarkdownDescription: "Port where the exporter will be listening for connections.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"priority_class_name": schema.StringAttribute{
										Description:         "PriorityClassName to be used in the Pod.",
										MarkdownDescription: "PriorityClassName to be used in the Pod.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Resouces describes the compute resource requirements.",
										MarkdownDescription: "Resouces describes the compute resource requirements.",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "ResourceList is a set of (resource name, quantity) pairs.",
												MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations to be used in the Pod.",
										MarkdownDescription: "Tolerations to be used in the Pod.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
													MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
													MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"password_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "PasswordSecretKeyRef is a reference to the password of the monitoring user used by the exporter. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								MarkdownDescription: "PasswordSecretKeyRef is a reference to the password of the monitoring user used by the exporter. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
								Attributes: map[string]schema.Attribute{
									"generate": schema.BoolAttribute{
										Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

							"service_monitor": schema.SingleNestedAttribute{
								Description:         "ServiceMonitor defines the ServiceMonior object.",
								MarkdownDescription: "ServiceMonitor defines the ServiceMonior object.",
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Description:         "Interval for scraping metrics.",
										MarkdownDescription: "Interval for scraping metrics.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"job_label": schema.StringAttribute{
										Description:         "JobLabel to add to the ServiceMonitor object.",
										MarkdownDescription: "JobLabel to add to the ServiceMonitor object.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"prometheus_release": schema.StringAttribute{
										Description:         "PrometheusRelease is the release label to add to the ServiceMonitor object.",
										MarkdownDescription: "PrometheusRelease is the release label to add to the ServiceMonitor object.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scrape_timeout": schema.StringAttribute{
										Description:         "ScrapeTimeout defines the timeout for scraping metrics.",
										MarkdownDescription: "ScrapeTimeout defines the timeout for scraping metrics.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"username": schema.StringAttribute{
								Description:         "Username is the username of the monitoring user used by the exporter.",
								MarkdownDescription: "Username is the username of the monitoring user used by the exporter.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"my_cnf": schema.StringAttribute{
						Description:         "MyCnf allows to specify the my.cnf file mounted by Mariadb. Updating this field will trigger an update to the Mariadb resource.",
						MarkdownDescription: "MyCnf allows to specify the my.cnf file mounted by Mariadb. Updating this field will trigger an update to the Mariadb resource.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"my_cnf_config_map_key_ref": schema.SingleNestedAttribute{
						Description:         "MyCnfConfigMapKeyRef is a reference to the my.cnf config file provided via a ConfigMap. If not provided, it will be defaulted with a reference to a ConfigMap containing the MyCnf field. If the referred ConfigMap is labeled with 'k8s.mariadb.com/watch', an update to the Mariadb resource will be triggered when the ConfigMap is updated.",
						MarkdownDescription: "MyCnfConfigMapKeyRef is a reference to the my.cnf config file provided via a ConfigMap. If not provided, it will be defaulted with a reference to a ConfigMap containing the MyCnf field. If the referred ConfigMap is labeled with 'k8s.mariadb.com/watch', an update to the Mariadb resource will be triggered when the ConfigMap is updated.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
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

					"node_selector": schema.MapAttribute{
						Description:         "NodeSelector to be used in the Pod.",
						MarkdownDescription: "NodeSelector to be used in the Pod.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"password_hash_secret_key_ref": schema.SingleNestedAttribute{
						Description:         "PasswordHashSecretKeyRef is a reference to the password hash to be used by the initial User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password hash.",
						MarkdownDescription: "PasswordHashSecretKeyRef is a reference to the password hash to be used by the initial User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password hash.",
						Attributes: map[string]schema.Attribute{
							"key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
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

					"password_plugin": schema.SingleNestedAttribute{
						Description:         "PasswordPlugin is a reference to the password plugin and arguments to be used by the initial User.",
						MarkdownDescription: "PasswordPlugin is a reference to the password plugin and arguments to be used by the initial User.",
						Attributes: map[string]schema.Attribute{
							"plugin_arg_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "PluginArgSecretKeyRef is a reference to the arguments to be provided to the authentication plugin for the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the authentication plugin arguments.",
								MarkdownDescription: "PluginArgSecretKeyRef is a reference to the arguments to be provided to the authentication plugin for the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the authentication plugin arguments.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

							"plugin_name_secret_key_ref": schema.SingleNestedAttribute{
								Description:         "PluginNameSecretKeyRef is a reference to the authentication plugin to be used by the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the authentication plugin.",
								MarkdownDescription: "PluginNameSecretKeyRef is a reference to the authentication plugin to be used by the User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the authentication plugin.",
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"name": schema.StringAttribute{
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

					"password_secret_key_ref": schema.SingleNestedAttribute{
						Description:         "PasswordSecretKeyRef is a reference to a Secret that contains the password to be used by the initial User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
						MarkdownDescription: "PasswordSecretKeyRef is a reference to a Secret that contains the password to be used by the initial User. If the referred Secret is labeled with 'k8s.mariadb.com/watch', updates may be performed to the Secret in order to update the password.",
						Attributes: map[string]schema.Attribute{
							"generate": schema.BoolAttribute{
								Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
								MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
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

					"pod_disruption_budget": schema.SingleNestedAttribute{
						Description:         "PodDisruptionBudget defines the budget for replica availability.",
						MarkdownDescription: "PodDisruptionBudget defines the budget for replica availability.",
						Attributes: map[string]schema.Attribute{
							"max_unavailable": schema.StringAttribute{
								Description:         "MaxUnavailable defines the number of maximum unavailable Pods.",
								MarkdownDescription: "MaxUnavailable defines the number of maximum unavailable Pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_available": schema.StringAttribute{
								Description:         "MinAvailable defines the number of minimum available Pods.",
								MarkdownDescription: "MinAvailable defines the number of minimum available Pods.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_metadata": schema.SingleNestedAttribute{
						Description:         "PodMetadata defines extra metadata for the Pod.",
						MarkdownDescription: "PodMetadata defines extra metadata for the Pod.",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations to be added to children resources.",
								MarkdownDescription: "Annotations to be added to children resources.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels to be added to children resources.",
								MarkdownDescription: "Labels to be added to children resources.",
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

					"pod_security_context": schema.SingleNestedAttribute{
						Description:         "SecurityContext holds pod-level security attributes and common container settings.",
						MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.",
						Attributes: map[string]schema.Attribute{
							"app_armor_profile": schema.SingleNestedAttribute{
								Description:         "appArmorProfile is the AppArmor options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "appArmorProfile is the AppArmor options to use by the containers in this pod. Note that this field cannot be set when spec.os.name is windows.",
								Attributes: map[string]schema.Attribute{
									"localhost_profile": schema.StringAttribute{
										Description:         "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
										MarkdownDescription: "localhostProfile indicates a profile loaded on the node that should be used. The profile must be preconfigured on the node to work. Must match the loaded name of the profile. Must be set if and only if type is 'Localhost'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
										MarkdownDescription: "type indicates which kind of AppArmor profile will be applied. Valid options are: Localhost - a profile pre-loaded on the node. RuntimeDefault - the container runtime's default profile. Unconfined - no AppArmor enforcement.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

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
								Description:         "A list of groups applied to the first process run in each container, in addition to the container's primary GID and fsGroup (if specified). If the SupplementalGroupsPolicy feature is enabled, the supplementalGroupsPolicy field determines whether these are in addition to or instead of any group memberships defined in the container image. If unspecified, no additional groups are added, though group memberships defined in the container image may still be used, depending on the supplementalGroupsPolicy field. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "A list of groups applied to the first process run in each container, in addition to the container's primary GID and fsGroup (if specified). If the SupplementalGroupsPolicy feature is enabled, the supplementalGroupsPolicy field determines whether these are in addition to or instead of any group memberships defined in the container image. If unspecified, no additional groups are added, though group memberships defined in the container image may still be used, depending on the supplementalGroupsPolicy field. Note that this field cannot be set when spec.os.name is windows.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"supplemental_groups_policy": schema.StringAttribute{
								Description:         "Defines how supplemental groups of the first container processes are calculated. Valid values are 'Merge' and 'Strict'. If not specified, 'Merge' is used. (Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled and the container runtime must implement support for this feature. Note that this field cannot be set when spec.os.name is windows.",
								MarkdownDescription: "Defines how supplemental groups of the first container processes are calculated. Valid values are 'Merge' and 'Strict'. If not specified, 'Merge' is used. (Alpha) Using the field requires the SupplementalGroupsPolicy feature gate to be enabled and the container runtime must implement support for this feature. Note that this field cannot be set when spec.os.name is windows.",
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

					"port": schema.Int64Attribute{
						Description:         "Port where the instances will be listening for connections.",
						MarkdownDescription: "Port where the instances will be listening for connections.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"primary_connection": schema.SingleNestedAttribute{
						Description:         "PrimaryConnection defines a template to configure the primary Connection object. This Connection provides the initial User access to the initial Database. It will make use of the PrimaryService to route network traffic to the primary Pod.",
						MarkdownDescription: "PrimaryConnection defines a template to configure the primary Connection object. This Connection provides the initial User access to the initial Database. It will make use of the PrimaryService to route network traffic to the primary Pod.",
						Attributes: map[string]schema.Attribute{
							"health_check": schema.SingleNestedAttribute{
								Description:         "HealthCheck to be used in the Connection.",
								MarkdownDescription: "HealthCheck to be used in the Connection.",
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Description:         "Interval used to perform health checks.",
										MarkdownDescription: "Interval used to perform health checks.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_interval": schema.StringAttribute{
										Description:         "RetryInterval is the interval used to perform health check retries.",
										MarkdownDescription: "RetryInterval is the interval used to perform health check retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"params": schema.MapAttribute{
								Description:         "Params to be used in the Connection.",
								MarkdownDescription: "Params to be used in the Connection.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
								MarkdownDescription: "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName to be used in the Connection.",
								MarkdownDescription: "SecretName to be used in the Connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_template": schema.SingleNestedAttribute{
								Description:         "SecretTemplate to be used in the Connection.",
								MarkdownDescription: "SecretTemplate to be used in the Connection.",
								Attributes: map[string]schema.Attribute{
									"database_key": schema.StringAttribute{
										Description:         "DatabaseKey to be used in the Secret.",
										MarkdownDescription: "DatabaseKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"format": schema.StringAttribute{
										Description:         "Format to be used in the Secret.",
										MarkdownDescription: "Format to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_key": schema.StringAttribute{
										Description:         "HostKey to be used in the Secret.",
										MarkdownDescription: "HostKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "Key to be used in the Secret.",
										MarkdownDescription: "Key to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata to be added to the Secret object.",
										MarkdownDescription: "Metadata to be added to the Secret object.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"password_key": schema.StringAttribute{
										Description:         "PasswordKey to be used in the Secret.",
										MarkdownDescription: "PasswordKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port_key": schema.StringAttribute{
										Description:         "PortKey to be used in the Secret.",
										MarkdownDescription: "PortKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username_key": schema.StringAttribute{
										Description:         "UsernameKey to be used in the Secret.",
										MarkdownDescription: "UsernameKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_name": schema.StringAttribute{
								Description:         "ServiceName to be used in the Connection.",
								MarkdownDescription: "ServiceName to be used in the Connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"primary_service": schema.SingleNestedAttribute{
						Description:         "PrimaryService defines a template to configure the primary Service object. The network traffic of this Service will be routed to the primary Pod.",
						MarkdownDescription: "PrimaryService defines a template to configure the primary Service object. The network traffic of this Service will be routed to the primary Pod.",
						Attributes: map[string]schema.Attribute{
							"allocate_load_balancer_node_ports": schema.BoolAttribute{
								Description:         "AllocateLoadBalancerNodePorts Service field.",
								MarkdownDescription: "AllocateLoadBalancerNodePorts Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_traffic_policy": schema.StringAttribute{
								Description:         "ExternalTrafficPolicy Service field.",
								MarkdownDescription: "ExternalTrafficPolicy Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_ip": schema.StringAttribute{
								Description:         "LoadBalancerIP Service field.",
								MarkdownDescription: "LoadBalancerIP Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_source_ranges": schema.ListAttribute{
								Description:         "LoadBalancerSourceRanges Service field.",
								MarkdownDescription: "LoadBalancerSourceRanges Service field.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata to be added to the Service metadata.",
								MarkdownDescription: "Metadata to be added to the Service metadata.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations to be added to children resources.",
										MarkdownDescription: "Annotations to be added to children resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels to be added to children resources.",
										MarkdownDescription: "Labels to be added to children resources.",
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

							"session_affinity": schema.StringAttribute{
								Description:         "SessionAffinity Service field.",
								MarkdownDescription: "SessionAffinity Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
								MarkdownDescription: "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
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

					"priority_class_name": schema.StringAttribute{
						Description:         "PriorityClassName to be used in the Pod.",
						MarkdownDescription: "PriorityClassName to be used in the Pod.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"readiness_probe": schema.SingleNestedAttribute{
						Description:         "ReadinessProbe to be used in the Container.",
						MarkdownDescription: "ReadinessProbe to be used in the Container.",
						Attributes: map[string]schema.Attribute{
							"exec": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#execaction-v1-core.",
								Attributes: map[string]schema.Attribute{
									"command": schema.ListAttribute{
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

							"failure_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"http_get": schema.SingleNestedAttribute{
								Description:         "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
								MarkdownDescription: "Refer to the Kubernetes docs: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.31/#httpgetaction-v1-core.",
								Attributes: map[string]schema.Attribute{
									"host": schema.StringAttribute{
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

									"port": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "URIScheme identifies the scheme used for connection to a host for Get actions",
										MarkdownDescription: "URIScheme identifies the scheme used for connection to a host for Get actions",
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
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"period_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"success_threshold": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timeout_seconds": schema.Int64Attribute{
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

					"replicas": schema.Int64Attribute{
						Description:         "Replicas indicates the number of desired instances.",
						MarkdownDescription: "Replicas indicates the number of desired instances.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replicas_allow_even_number": schema.BoolAttribute{
						Description:         "disables the validation check for an odd number of replicas.",
						MarkdownDescription: "disables the validation check for an odd number of replicas.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replication": schema.SingleNestedAttribute{
						Description:         "Replication configures high availability via replication. This feature is still in alpha, use Galera if you are looking for a more production-ready HA.",
						MarkdownDescription: "Replication configures high availability via replication. This feature is still in alpha, use Galera if you are looking for a more production-ready HA.",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled is a flag to enable Replication.",
								MarkdownDescription: "Enabled is a flag to enable Replication.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"primary": schema.SingleNestedAttribute{
								Description:         "Primary is the replication configuration for the primary node.",
								MarkdownDescription: "Primary is the replication configuration for the primary node.",
								Attributes: map[string]schema.Attribute{
									"automatic_failover": schema.BoolAttribute{
										Description:         "AutomaticFailover indicates whether the operator should automatically update PodIndex to perform an automatic primary failover.",
										MarkdownDescription: "AutomaticFailover indicates whether the operator should automatically update PodIndex to perform an automatic primary failover.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_index": schema.Int64Attribute{
										Description:         "PodIndex is the StatefulSet index of the primary node. The user may change this field to perform a manual switchover.",
										MarkdownDescription: "PodIndex is the StatefulSet index of the primary node. The user may change this field to perform a manual switchover.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"probes_enabled": schema.BoolAttribute{
								Description:         "ProbesEnabled indicates to use replication specific liveness and readiness probes. This probes check that the primary can receive queries and that the replica has the replication thread running.",
								MarkdownDescription: "ProbesEnabled indicates to use replication specific liveness and readiness probes. This probes check that the primary can receive queries and that the replica has the replication thread running.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica": schema.SingleNestedAttribute{
								Description:         "ReplicaReplication is the replication configuration for the replica nodes.",
								MarkdownDescription: "ReplicaReplication is the replication configuration for the replica nodes.",
								Attributes: map[string]schema.Attribute{
									"connection_retries": schema.Int64Attribute{
										Description:         "ConnectionRetries to be used when the replica connects to the primary.",
										MarkdownDescription: "ConnectionRetries to be used when the replica connects to the primary.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"connection_timeout": schema.StringAttribute{
										Description:         "ConnectionTimeout to be used when the replica connects to the primary.",
										MarkdownDescription: "ConnectionTimeout to be used when the replica connects to the primary.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"gtid": schema.StringAttribute{
										Description:         "Gtid indicates which Global Transaction ID should be used when connecting a replica to the master. See: https://mariadb.com/kb/en/gtid/#using-current_pos-vs-slave_pos.",
										MarkdownDescription: "Gtid indicates which Global Transaction ID should be used when connecting a replica to the master. See: https://mariadb.com/kb/en/gtid/#using-current_pos-vs-slave_pos.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("CurrentPos", "SlavePos"),
										},
									},

									"repl_password_secret_key_ref": schema.SingleNestedAttribute{
										Description:         "ReplPasswordSecretKeyRef provides a reference to the Secret to use as password for the replication user.",
										MarkdownDescription: "ReplPasswordSecretKeyRef provides a reference to the Secret to use as password for the replication user.",
										Attributes: map[string]schema.Attribute{
											"generate": schema.BoolAttribute{
												Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
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

									"sync_timeout": schema.StringAttribute{
										Description:         "SyncTimeout defines the timeout for a replica to be synced with the primary when performing a primary switchover. If the timeout is reached, the replica GTID will be reset and the switchover will continue.",
										MarkdownDescription: "SyncTimeout defines the timeout for a replica to be synced with the primary when performing a primary switchover. If the timeout is reached, the replica GTID will be reset and the switchover will continue.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"wait_point": schema.StringAttribute{
										Description:         "WaitPoint defines whether the transaction should wait for ACK before committing to the storage engine. More info: https://mariadb.com/kb/en/semisynchronous-replication/#rpl_semi_sync_master_wait_point.",
										MarkdownDescription: "WaitPoint defines whether the transaction should wait for ACK before committing to the storage engine. More info: https://mariadb.com/kb/en/semisynchronous-replication/#rpl_semi_sync_master_wait_point.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("AfterSync", "AfterCommit"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"sync_binlog": schema.BoolAttribute{
								Description:         "SyncBinlog indicates whether the binary log should be synchronized to the disk after every event. It trades off performance for consistency. See: https://mariadb.com/kb/en/replication-and-binary-log-system-variables/#sync_binlog.",
								MarkdownDescription: "SyncBinlog indicates whether the binary log should be synchronized to the disk after every event. It trades off performance for consistency. See: https://mariadb.com/kb/en/replication-and-binary-log-system-variables/#sync_binlog.",
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
						Description:         "Resouces describes the compute resource requirements.",
						MarkdownDescription: "Resouces describes the compute resource requirements.",
						Attributes: map[string]schema.Attribute{
							"limits": schema.MapAttribute{
								Description:         "ResourceList is a set of (resource name, quantity) pairs.",
								MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "ResourceList is a set of (resource name, quantity) pairs.",
								MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

					"root_empty_password": schema.BoolAttribute{
						Description:         "RootEmptyPassword indicates if the root password should be empty. Don't use this feature in production, it is only intended for development and test environments.",
						MarkdownDescription: "RootEmptyPassword indicates if the root password should be empty. Don't use this feature in production, it is only intended for development and test environments.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"root_password_secret_key_ref": schema.SingleNestedAttribute{
						Description:         "RootPasswordSecretKeyRef is a reference to a Secret key containing the root password.",
						MarkdownDescription: "RootPasswordSecretKeyRef is a reference to a Secret key containing the root password.",
						Attributes: map[string]schema.Attribute{
							"generate": schema.BoolAttribute{
								Description:         "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
								MarkdownDescription: "Generate indicates whether the Secret should be generated if the Secret referenced is not present.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
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

					"secondary_connection": schema.SingleNestedAttribute{
						Description:         "SecondaryConnection defines a template to configure the secondary Connection object. This Connection provides the initial User access to the initial Database. It will make use of the SecondaryService to route network traffic to the secondary Pods.",
						MarkdownDescription: "SecondaryConnection defines a template to configure the secondary Connection object. This Connection provides the initial User access to the initial Database. It will make use of the SecondaryService to route network traffic to the secondary Pods.",
						Attributes: map[string]schema.Attribute{
							"health_check": schema.SingleNestedAttribute{
								Description:         "HealthCheck to be used in the Connection.",
								MarkdownDescription: "HealthCheck to be used in the Connection.",
								Attributes: map[string]schema.Attribute{
									"interval": schema.StringAttribute{
										Description:         "Interval used to perform health checks.",
										MarkdownDescription: "Interval used to perform health checks.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"retry_interval": schema.StringAttribute{
										Description:         "RetryInterval is the interval used to perform health check retries.",
										MarkdownDescription: "RetryInterval is the interval used to perform health check retries.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"params": schema.MapAttribute{
								Description:         "Params to be used in the Connection.",
								MarkdownDescription: "Params to be used in the Connection.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"port": schema.Int64Attribute{
								Description:         "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
								MarkdownDescription: "Port to connect to. If not provided, it defaults to the MariaDB port or to the first MaxScale listener.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "SecretName to be used in the Connection.",
								MarkdownDescription: "SecretName to be used in the Connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_template": schema.SingleNestedAttribute{
								Description:         "SecretTemplate to be used in the Connection.",
								MarkdownDescription: "SecretTemplate to be used in the Connection.",
								Attributes: map[string]schema.Attribute{
									"database_key": schema.StringAttribute{
										Description:         "DatabaseKey to be used in the Secret.",
										MarkdownDescription: "DatabaseKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"format": schema.StringAttribute{
										Description:         "Format to be used in the Secret.",
										MarkdownDescription: "Format to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_key": schema.StringAttribute{
										Description:         "HostKey to be used in the Secret.",
										MarkdownDescription: "HostKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"key": schema.StringAttribute{
										Description:         "Key to be used in the Secret.",
										MarkdownDescription: "Key to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata to be added to the Secret object.",
										MarkdownDescription: "Metadata to be added to the Secret object.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"password_key": schema.StringAttribute{
										Description:         "PasswordKey to be used in the Secret.",
										MarkdownDescription: "PasswordKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port_key": schema.StringAttribute{
										Description:         "PortKey to be used in the Secret.",
										MarkdownDescription: "PortKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username_key": schema.StringAttribute{
										Description:         "UsernameKey to be used in the Secret.",
										MarkdownDescription: "UsernameKey to be used in the Secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_name": schema.StringAttribute{
								Description:         "ServiceName to be used in the Connection.",
								MarkdownDescription: "ServiceName to be used in the Connection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"secondary_service": schema.SingleNestedAttribute{
						Description:         "SecondaryService defines a template to configure the secondary Service object. The network traffic of this Service will be routed to the secondary Pods.",
						MarkdownDescription: "SecondaryService defines a template to configure the secondary Service object. The network traffic of this Service will be routed to the secondary Pods.",
						Attributes: map[string]schema.Attribute{
							"allocate_load_balancer_node_ports": schema.BoolAttribute{
								Description:         "AllocateLoadBalancerNodePorts Service field.",
								MarkdownDescription: "AllocateLoadBalancerNodePorts Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_traffic_policy": schema.StringAttribute{
								Description:         "ExternalTrafficPolicy Service field.",
								MarkdownDescription: "ExternalTrafficPolicy Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_ip": schema.StringAttribute{
								Description:         "LoadBalancerIP Service field.",
								MarkdownDescription: "LoadBalancerIP Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_source_ranges": schema.ListAttribute{
								Description:         "LoadBalancerSourceRanges Service field.",
								MarkdownDescription: "LoadBalancerSourceRanges Service field.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata to be added to the Service metadata.",
								MarkdownDescription: "Metadata to be added to the Service metadata.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations to be added to children resources.",
										MarkdownDescription: "Annotations to be added to children resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels to be added to children resources.",
										MarkdownDescription: "Labels to be added to children resources.",
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

							"session_affinity": schema.StringAttribute{
								Description:         "SessionAffinity Service field.",
								MarkdownDescription: "SessionAffinity Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
								MarkdownDescription: "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
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

					"security_context": schema.SingleNestedAttribute{
						Description:         "SecurityContext holds security configuration that will be applied to a container.",
						MarkdownDescription: "SecurityContext holds security configuration that will be applied to a container.",
						Attributes: map[string]schema.Attribute{
							"allow_privilege_escalation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"capabilities": schema.SingleNestedAttribute{
								Description:         "Adds and removes POSIX capabilities from running containers.",
								MarkdownDescription: "Adds and removes POSIX capabilities from running containers.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service": schema.SingleNestedAttribute{
						Description:         "Service defines a template to configure the general Service object. The network traffic of this Service will be routed to all Pods.",
						MarkdownDescription: "Service defines a template to configure the general Service object. The network traffic of this Service will be routed to all Pods.",
						Attributes: map[string]schema.Attribute{
							"allocate_load_balancer_node_ports": schema.BoolAttribute{
								Description:         "AllocateLoadBalancerNodePorts Service field.",
								MarkdownDescription: "AllocateLoadBalancerNodePorts Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_traffic_policy": schema.StringAttribute{
								Description:         "ExternalTrafficPolicy Service field.",
								MarkdownDescription: "ExternalTrafficPolicy Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_ip": schema.StringAttribute{
								Description:         "LoadBalancerIP Service field.",
								MarkdownDescription: "LoadBalancerIP Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"load_balancer_source_ranges": schema.ListAttribute{
								Description:         "LoadBalancerSourceRanges Service field.",
								MarkdownDescription: "LoadBalancerSourceRanges Service field.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata to be added to the Service metadata.",
								MarkdownDescription: "Metadata to be added to the Service metadata.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations to be added to children resources.",
										MarkdownDescription: "Annotations to be added to children resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels to be added to children resources.",
										MarkdownDescription: "Labels to be added to children resources.",
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

							"session_affinity": schema.StringAttribute{
								Description:         "SessionAffinity Service field.",
								MarkdownDescription: "SessionAffinity Service field.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
								MarkdownDescription: "Type is the Service type. One of 'ClusterIP', 'NodePort' or 'LoadBalancer'. If not defined, it defaults to 'ClusterIP'.",
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

					"service_account_name": schema.StringAttribute{
						Description:         "ServiceAccountName is the name of the ServiceAccount to be used by the Pods.",
						MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to be used by the Pods.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sidecar_containers": schema.ListNestedAttribute{
						Description:         "SidecarContainers to be used in the Pod.",
						MarkdownDescription: "SidecarContainers to be used in the Pod.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"args": schema.ListAttribute{
									Description:         "Args to be used in the Container.",
									MarkdownDescription: "Args to be used in the Container.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"command": schema.ListAttribute{
									Description:         "Command to be used in the Container.",
									MarkdownDescription: "Command to be used in the Container.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image": schema.StringAttribute{
									Description:         "Image name to be used by the container. The supported format is '<image>:<tag>'.",
									MarkdownDescription: "Image name to be used by the container. The supported format is '<image>:<tag>'.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"image_pull_policy": schema.StringAttribute{
									Description:         "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
									MarkdownDescription: "ImagePullPolicy is the image pull policy. One of 'Always', 'Never' or 'IfNotPresent'. If not defined, it defaults to 'IfNotPresent'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Always", "Never", "IfNotPresent"),
									},
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Resouces describes the compute resource requirements.",
									MarkdownDescription: "Resouces describes the compute resource requirements.",
									Attributes: map[string]schema.Attribute{
										"limits": schema.MapAttribute{
											Description:         "ResourceList is a set of (resource name, quantity) pairs.",
											MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"requests": schema.MapAttribute{
											Description:         "ResourceList is a set of (resource name, quantity) pairs.",
											MarkdownDescription: "ResourceList is a set of (resource name, quantity) pairs.",
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

								"volume_mounts": schema.ListNestedAttribute{
									Description:         "VolumeMounts to be used in the Container.",
									MarkdownDescription: "VolumeMounts to be used in the Container.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"mount_path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
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

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage defines the storage options to be used for provisioning the PVCs mounted by MariaDB.",
						MarkdownDescription: "Storage defines the storage options to be used for provisioning the PVCs mounted by MariaDB.",
						Attributes: map[string]schema.Attribute{
							"ephemeral": schema.BoolAttribute{
								Description:         "Ephemeral indicates whether to use ephemeral storage in the PVCs. It is only compatible with non HA MariaDBs.",
								MarkdownDescription: "Ephemeral indicates whether to use ephemeral storage in the PVCs. It is only compatible with non HA MariaDBs.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resize_in_use_volumes": schema.BoolAttribute{
								Description:         "ResizeInUseVolumes indicates whether the PVCs can be resized. The 'StorageClassName' used should have 'allowVolumeExpansion' set to 'true' to allow resizing. It defaults to true.",
								MarkdownDescription: "ResizeInUseVolumes indicates whether the PVCs can be resized. The 'StorageClassName' used should have 'allowVolumeExpansion' set to 'true' to allow resizing. It defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "Size of the PVCs to be mounted by MariaDB. Required if not provided in 'VolumeClaimTemplate'. It superseeds the storage size specified in 'VolumeClaimTemplate'.",
								MarkdownDescription: "Size of the PVCs to be mounted by MariaDB. Required if not provided in 'VolumeClaimTemplate'. It superseeds the storage size specified in 'VolumeClaimTemplate'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_class_name": schema.StringAttribute{
								Description:         "StorageClassName to be used to provision the PVCS. It superseeds the 'StorageClassName' specified in 'VolumeClaimTemplate'. If not provided, the default 'StorageClass' configured in the cluster is used.",
								MarkdownDescription: "StorageClassName to be used to provision the PVCS. It superseeds the 'StorageClassName' specified in 'VolumeClaimTemplate'. If not provided, the default 'StorageClass' configured in the cluster is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_claim_template": schema.SingleNestedAttribute{
								Description:         "VolumeClaimTemplate provides a template to define the PVCs.",
								MarkdownDescription: "VolumeClaimTemplate provides a template to define the PVCs.",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"metadata": schema.SingleNestedAttribute{
										Description:         "Metadata to be added to the PVC metadata.",
										MarkdownDescription: "Metadata to be added to the PVC metadata.",
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Annotations to be added to children resources.",
												MarkdownDescription: "Annotations to be added to children resources.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Labels to be added to children resources.",
												MarkdownDescription: "Labels to be added to children resources.",
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

									"resources": schema.SingleNestedAttribute{
										Description:         "VolumeResourceRequirements describes the storage resource requirements for a volume.",
										MarkdownDescription: "VolumeResourceRequirements describes the storage resource requirements for a volume.",
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
										Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
										MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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

							"wait_for_volume_resize": schema.BoolAttribute{
								Description:         "WaitForVolumeResize indicates whether to wait for the PVCs to be resized before marking the MariaDB object as ready. This will block other operations such as cluster recovery while the resize is in progress. It defaults to true.",
								MarkdownDescription: "WaitForVolumeResize indicates whether to wait for the PVCs to be resized before marking the MariaDB object as ready. This will block other operations such as cluster recovery while the resize is in progress. It defaults to true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"suspend": schema.BoolAttribute{
						Description:         "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
						MarkdownDescription: "Suspend indicates whether the current resource should be suspended or not. This can be useful for maintenance, as disabling the reconciliation prevents the operator from interfering with user operations during maintenance activities.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"time_zone": schema.StringAttribute{
						Description:         "TimeZone sets the default timezone. If not provided, it defaults to SYSTEM and the timezone data is not loaded.",
						MarkdownDescription: "TimeZone sets the default timezone. If not provided, it defaults to SYSTEM and the timezone data is not loaded.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "Tolerations to be used in the Pod.",
						MarkdownDescription: "Tolerations to be used in the Pod.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"key": schema.StringAttribute{
									Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"operator": schema.StringAttribute{
									Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
									MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
						Description:         "TopologySpreadConstraints to be used in the Pod.",
						MarkdownDescription: "TopologySpreadConstraints to be used in the Pod.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"label_selector": schema.SingleNestedAttribute{
									Description:         "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
									MarkdownDescription: "A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.",
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
									Required:            true,
									Optional:            false,
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
									Description:         "NodeInclusionPolicy defines the type of node inclusion policy",
									MarkdownDescription: "NodeInclusionPolicy defines the type of node inclusion policy",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"node_taints_policy": schema.StringAttribute{
									Description:         "NodeInclusionPolicy defines the type of node inclusion policy",
									MarkdownDescription: "NodeInclusionPolicy defines the type of node inclusion policy",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"topology_key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"when_unsatisfiable": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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

					"update_strategy": schema.SingleNestedAttribute{
						Description:         "UpdateStrategy defines how a MariaDB resource is updated.",
						MarkdownDescription: "UpdateStrategy defines how a MariaDB resource is updated.",
						Attributes: map[string]schema.Attribute{
							"auto_update_data_plane": schema.BoolAttribute{
								Description:         "AutoUpdateDataPlane indicates whether the Galera data-plane version (agent and init containers) should be automatically updated based on the operator version. It defaults to false. Updating the operator will trigger updates on all the MariaDB instances that have this flag set to true. Thus, it is recommended to progressively set this flag after having updated the operator.",
								MarkdownDescription: "AutoUpdateDataPlane indicates whether the Galera data-plane version (agent and init containers) should be automatically updated based on the operator version. It defaults to false. Updating the operator will trigger updates on all the MariaDB instances that have this flag set to true. Thus, it is recommended to progressively set this flag after having updated the operator.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"rolling_update": schema.SingleNestedAttribute{
								Description:         "RollingUpdate defines parameters for the RollingUpdate type.",
								MarkdownDescription: "RollingUpdate defines parameters for the RollingUpdate type.",
								Attributes: map[string]schema.Attribute{
									"max_unavailable": schema.StringAttribute{
										Description:         "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0. Defaults to 1. This field is alpha-level and is only honored by servers that enable the MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it will be counted towards MaxUnavailable.",
										MarkdownDescription: "The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0. Defaults to 1. This field is alpha-level and is only honored by servers that enable the MaxUnavailableStatefulSet feature. The field applies to all pods in the range 0 to Replicas-1. That means if there is any unavailable pod in the range 0 to Replicas-1, it will be counted towards MaxUnavailable.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"partition": schema.Int64Attribute{
										Description:         "Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0.",
										MarkdownDescription: "Partition indicates the ordinal at which the StatefulSet should be partitioned for updates. During a rolling update, all pods from ordinal Replicas-1 to Partition are updated. All pods from ordinal Partition-1 to 0 remain untouched. This is helpful in being able to do a canary based deployment. The default value is 0.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"type": schema.StringAttribute{
								Description:         "Type defines the type of updates. One of 'ReplicasFirstPrimaryLast', 'RollingUpdate' or 'OnDelete'. If not defined, it defaults to 'ReplicasFirstPrimaryLast'.",
								MarkdownDescription: "Type defines the type of updates. One of 'ReplicasFirstPrimaryLast', 'RollingUpdate' or 'OnDelete'. If not defined, it defaults to 'ReplicasFirstPrimaryLast'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ReplicasFirstPrimaryLast", "RollingUpdate", "OnDelete", "Never"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"username": schema.StringAttribute{
						Description:         "Username is the initial username to be created by the operator once MariaDB is ready. It has all privileges on the initial database. The initial User will have ALL PRIVILEGES in the initial Database.",
						MarkdownDescription: "Username is the initial username to be created by the operator once MariaDB is ready. It has all privileges on the initial database. The initial User will have ALL PRIVILEGES in the initial Database.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_mounts": schema.ListNestedAttribute{
						Description:         "VolumeMounts to be used in the Container.",
						MarkdownDescription: "VolumeMounts to be used in the Container.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mount_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"volumes": schema.ListNestedAttribute{
						Description:         "Volumes to be used in the Pod.",
						MarkdownDescription: "Volumes to be used in the Pod.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"csi": schema.SingleNestedAttribute{
									Description:         "Represents a source location of a volume to mount, managed by an external CSI driver",
									MarkdownDescription: "Represents a source location of a volume to mount, managed by an external CSI driver",
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
													Description:         "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
													MarkdownDescription: "Name of the referent. This field is effectively required, but due to backwards compatibility is allowed to be empty. Instances of this type with an empty value here are almost certainly wrong. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
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

								"empty_dir": schema.SingleNestedAttribute{
									Description:         "Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.",
									MarkdownDescription: "Represents an empty directory for a pod. Empty directory volumes support ownership management and SELinux relabeling.",
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

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"nfs": schema.SingleNestedAttribute{
									Description:         "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",
									MarkdownDescription: "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",
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
									Description:         "PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).",
									MarkdownDescription: "PersistentVolumeClaimVolumeSource references the user's PVC in the same namespace. This volume finds the bound PV and mounts that volume for the pod. A PersistentVolumeClaimVolumeSource is, essentially, a wrapper around another type of volume that is owned by someone else (the system).",
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
	}
}

func (r *K8SMariadbComMariaDbV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_k8s_mariadb_com_maria_db_v1alpha1_manifest")

	var model K8SMariadbComMariaDbV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("k8s.mariadb.com/v1alpha1")
	model.Kind = pointer.String("MariaDB")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
