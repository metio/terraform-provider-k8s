/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package app_redislabs_com_v1alpha1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &AppRedislabsComRedisEnterpriseClusterV1Alpha1Manifest{}
)

func NewAppRedislabsComRedisEnterpriseClusterV1Alpha1Manifest() datasource.DataSource {
	return &AppRedislabsComRedisEnterpriseClusterV1Alpha1Manifest{}
}

type AppRedislabsComRedisEnterpriseClusterV1Alpha1Manifest struct{}

type AppRedislabsComRedisEnterpriseClusterV1Alpha1ManifestData struct {
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
		ActiveActive *struct {
			ApiIngressUrl      *string            `tfsdk:"api_ingress_url" json:"apiIngressUrl,omitempty"`
			DbIngressSuffix    *string            `tfsdk:"db_ingress_suffix" json:"dbIngressSuffix,omitempty"`
			IngressAnnotations *map[string]string `tfsdk:"ingress_annotations" json:"ingressAnnotations,omitempty"`
			Method             *string            `tfsdk:"method" json:"method,omitempty"`
		} `tfsdk:"active_active" json:"activeActive,omitempty"`
		AntiAffinityAdditionalTopologyKeys *[]string `tfsdk:"anti_affinity_additional_topology_keys" json:"antiAffinityAdditionalTopologyKeys,omitempty"`
		BootstrapperImageSpec              *struct {
			DigestHash      *string `tfsdk:"digest_hash" json:"digestHash,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			Repository      *string `tfsdk:"repository" json:"repository,omitempty"`
			VersionTag      *string `tfsdk:"version_tag" json:"versionTag,omitempty"`
		} `tfsdk:"bootstrapper_image_spec" json:"bootstrapperImageSpec,omitempty"`
		BootstrapperResources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"bootstrapper_resources" json:"bootstrapperResources,omitempty"`
		Certificates *struct {
			ApiCertificateSecretName             *string `tfsdk:"api_certificate_secret_name" json:"apiCertificateSecretName,omitempty"`
			CmCertificateSecretName              *string `tfsdk:"cm_certificate_secret_name" json:"cmCertificateSecretName,omitempty"`
			LdapClientCertificateSecretName      *string `tfsdk:"ldap_client_certificate_secret_name" json:"ldapClientCertificateSecretName,omitempty"`
			MetricsExporterCertificateSecretName *string `tfsdk:"metrics_exporter_certificate_secret_name" json:"metricsExporterCertificateSecretName,omitempty"`
			ProxyCertificateSecretName           *string `tfsdk:"proxy_certificate_secret_name" json:"proxyCertificateSecretName,omitempty"`
			SyncerCertificateSecretName          *string `tfsdk:"syncer_certificate_secret_name" json:"syncerCertificateSecretName,omitempty"`
		} `tfsdk:"certificates" json:"certificates,omitempty"`
		ClusterCredentialSecretName *string `tfsdk:"cluster_credential_secret_name" json:"clusterCredentialSecretName,omitempty"`
		ClusterCredentialSecretRole *string `tfsdk:"cluster_credential_secret_role" json:"clusterCredentialSecretRole,omitempty"`
		ClusterCredentialSecretType *string `tfsdk:"cluster_credential_secret_type" json:"clusterCredentialSecretType,omitempty"`
		ClusterRecovery             *bool   `tfsdk:"cluster_recovery" json:"clusterRecovery,omitempty"`
		ContainerTimezone           *struct {
			PropagateHost *map[string]string `tfsdk:"propagate_host" json:"propagateHost,omitempty"`
			TimezoneName  *string            `tfsdk:"timezone_name" json:"timezoneName,omitempty"`
		} `tfsdk:"container_timezone" json:"containerTimezone,omitempty"`
		CreateServiceAccount    *bool `tfsdk:"create_service_account" json:"createServiceAccount,omitempty"`
		DataInternodeEncryption *bool `tfsdk:"data_internode_encryption" json:"dataInternodeEncryption,omitempty"`
		EncryptPkeys            *bool `tfsdk:"encrypt_pkeys" json:"encryptPkeys,omitempty"`
		EnforceIPv4             *bool `tfsdk:"enforce_i_pv4" json:"enforceIPv4,omitempty"`
		ExtraEnvVars            *[]struct {
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
		} `tfsdk:"extra_env_vars" json:"extraEnvVars,omitempty"`
		ExtraLabels *map[string]string `tfsdk:"extra_labels" json:"extraLabels,omitempty"`
		HostAliases *[]struct {
			Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
			Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
		} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
		IngressOrRouteSpec *struct {
			ApiFqdnUrl         *string            `tfsdk:"api_fqdn_url" json:"apiFqdnUrl,omitempty"`
			DbFqdnSuffix       *string            `tfsdk:"db_fqdn_suffix" json:"dbFqdnSuffix,omitempty"`
			IngressAnnotations *map[string]string `tfsdk:"ingress_annotations" json:"ingressAnnotations,omitempty"`
			Method             *string            `tfsdk:"method" json:"method,omitempty"`
		} `tfsdk:"ingress_or_route_spec" json:"ingressOrRouteSpec,omitempty"`
		Ldap *struct {
			AuthenticationQuery *struct {
				Query *struct {
					Base   *string `tfsdk:"base" json:"base,omitempty"`
					Filter *string `tfsdk:"filter" json:"filter,omitempty"`
					Scope  *string `tfsdk:"scope" json:"scope,omitempty"`
				} `tfsdk:"query" json:"query,omitempty"`
				Template *string `tfsdk:"template" json:"template,omitempty"`
			} `tfsdk:"authentication_query" json:"authenticationQuery,omitempty"`
			AuthorizationQuery *struct {
				Attribute *string `tfsdk:"attribute" json:"attribute,omitempty"`
				Query     *struct {
					Base   *string `tfsdk:"base" json:"base,omitempty"`
					Filter *string `tfsdk:"filter" json:"filter,omitempty"`
					Scope  *string `tfsdk:"scope" json:"scope,omitempty"`
				} `tfsdk:"query" json:"query,omitempty"`
			} `tfsdk:"authorization_query" json:"authorizationQuery,omitempty"`
			BindCredentialsSecretName *string `tfsdk:"bind_credentials_secret_name" json:"bindCredentialsSecretName,omitempty"`
			CaCertificateSecretName   *string `tfsdk:"ca_certificate_secret_name" json:"caCertificateSecretName,omitempty"`
			CacheTTLSeconds           *int64  `tfsdk:"cache_ttl_seconds" json:"cacheTTLSeconds,omitempty"`
			EnabledForControlPlane    *bool   `tfsdk:"enabled_for_control_plane" json:"enabledForControlPlane,omitempty"`
			EnabledForDataPlane       *bool   `tfsdk:"enabled_for_data_plane" json:"enabledForDataPlane,omitempty"`
			Protocol                  *string `tfsdk:"protocol" json:"protocol,omitempty"`
			Servers                   *[]struct {
				Host *string `tfsdk:"host" json:"host,omitempty"`
				Port *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"servers" json:"servers,omitempty"`
		} `tfsdk:"ldap" json:"ldap,omitempty"`
		License           *string            `tfsdk:"license" json:"license,omitempty"`
		LicenseSecretName *string            `tfsdk:"license_secret_name" json:"licenseSecretName,omitempty"`
		NodeSelector      *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
		Nodes             *int64             `tfsdk:"nodes" json:"nodes,omitempty"`
		OcspConfiguration *struct {
			OcspFunctionality *bool  `tfsdk:"ocsp_functionality" json:"ocspFunctionality,omitempty"`
			QueryFrequency    *int64 `tfsdk:"query_frequency" json:"queryFrequency,omitempty"`
			RecoveryFrequency *int64 `tfsdk:"recovery_frequency" json:"recoveryFrequency,omitempty"`
			RecoveryMaxTries  *int64 `tfsdk:"recovery_max_tries" json:"recoveryMaxTries,omitempty"`
			ResponseTimeout   *int64 `tfsdk:"response_timeout" json:"responseTimeout,omitempty"`
		} `tfsdk:"ocsp_configuration" json:"ocspConfiguration,omitempty"`
		PersistentSpec *struct {
			Enabled          *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
			VolumeSize       *string `tfsdk:"volume_size" json:"volumeSize,omitempty"`
		} `tfsdk:"persistent_spec" json:"persistentSpec,omitempty"`
		PodAnnotations  *map[string]string `tfsdk:"pod_annotations" json:"podAnnotations,omitempty"`
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
		PodSecurityPolicyName *string `tfsdk:"pod_security_policy_name" json:"podSecurityPolicyName,omitempty"`
		PodStartingPolicy     *struct {
			Enabled                  *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
			StartingThresholdSeconds *int64 `tfsdk:"starting_threshold_seconds" json:"startingThresholdSeconds,omitempty"`
		} `tfsdk:"pod_starting_policy" json:"podStartingPolicy,omitempty"`
		PodTolerations *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"pod_tolerations" json:"podTolerations,omitempty"`
		PriorityClassName *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		PullSecrets       *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"pull_secrets" json:"pullSecrets,omitempty"`
		RackAwarenessNodeLabel                     *string `tfsdk:"rack_awareness_node_label" json:"rackAwarenessNodeLabel,omitempty"`
		RedisEnterpriseAdditionalPodSpecAttributes *struct {
			ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" json:"activeDeadlineSeconds,omitempty"`
			Affinity              *struct {
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
			AutomountServiceAccountToken *bool `tfsdk:"automount_service_account_token" json:"automountServiceAccountToken,omitempty"`
			DnsConfig                    *struct {
				Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				Options     *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
			} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
			DnsPolicy           *string `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
			EnableServiceLinks  *bool   `tfsdk:"enable_service_links" json:"enableServiceLinks,omitempty"`
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
			HostAliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
				Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
			} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
			HostIPC          *bool   `tfsdk:"host_ipc" json:"hostIPC,omitempty"`
			HostNetwork      *bool   `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			HostPID          *bool   `tfsdk:"host_pid" json:"hostPID,omitempty"`
			HostUsers        *bool   `tfsdk:"host_users" json:"hostUsers,omitempty"`
			Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
			ImagePullSecrets *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
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
			NodeName     *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			Os           *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"os" json:"os,omitempty"`
			Overhead          *map[string]string `tfsdk:"overhead" json:"overhead,omitempty"`
			PreemptionPolicy  *string            `tfsdk:"preemption_policy" json:"preemptionPolicy,omitempty"`
			Priority          *int64             `tfsdk:"priority" json:"priority,omitempty"`
			PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
			ReadinessGates    *[]struct {
				ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
			} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
			ResourceClaims *[]struct {
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Source *struct {
					ResourceClaimName         *string `tfsdk:"resource_claim_name" json:"resourceClaimName,omitempty"`
					ResourceClaimTemplateName *string `tfsdk:"resource_claim_template_name" json:"resourceClaimTemplateName,omitempty"`
				} `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"resource_claims" json:"resourceClaims,omitempty"`
			RestartPolicy    *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
			RuntimeClassName *string `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
			SchedulerName    *string `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
			SchedulingGates  *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"scheduling_gates" json:"schedulingGates,omitempty"`
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
			ServiceAccount                *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
			ServiceAccountName            *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
			SetHostnameAsFQDN             *bool   `tfsdk:"set_hostname_as_fqdn" json:"setHostnameAsFQDN,omitempty"`
			ShareProcessNamespace         *bool   `tfsdk:"share_process_namespace" json:"shareProcessNamespace,omitempty"`
			Subdomain                     *string `tfsdk:"subdomain" json:"subdomain,omitempty"`
			TerminationGracePeriodSeconds *int64  `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
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
		} `tfsdk:"redis_enterprise_additional_pod_spec_attributes" json:"redisEnterpriseAdditionalPodSpecAttributes,omitempty"`
		RedisEnterpriseImageSpec *struct {
			DigestHash      *string `tfsdk:"digest_hash" json:"digestHash,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			Repository      *string `tfsdk:"repository" json:"repository,omitempty"`
			VersionTag      *string `tfsdk:"version_tag" json:"versionTag,omitempty"`
		} `tfsdk:"redis_enterprise_image_spec" json:"redisEnterpriseImageSpec,omitempty"`
		RedisEnterpriseNodeResources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"redis_enterprise_node_resources" json:"redisEnterpriseNodeResources,omitempty"`
		RedisEnterprisePodAnnotations        *map[string]string `tfsdk:"redis_enterprise_pod_annotations" json:"redisEnterprisePodAnnotations,omitempty"`
		RedisEnterpriseServicesConfiguration *struct {
			CmServer *struct {
				OperatingMode *string `tfsdk:"operating_mode" json:"operatingMode,omitempty"`
			} `tfsdk:"cm_server" json:"cmServer,omitempty"`
			CrdbCoordinator *struct {
				OperatingMode *string `tfsdk:"operating_mode" json:"operatingMode,omitempty"`
			} `tfsdk:"crdb_coordinator" json:"crdbCoordinator,omitempty"`
			CrdbWorker *struct {
				OperatingMode *string `tfsdk:"operating_mode" json:"operatingMode,omitempty"`
			} `tfsdk:"crdb_worker" json:"crdbWorker,omitempty"`
			MdnsServer *struct {
				OperatingMode *string `tfsdk:"operating_mode" json:"operatingMode,omitempty"`
			} `tfsdk:"mdns_server" json:"mdnsServer,omitempty"`
			PdnsServer *struct {
				OperatingMode *string `tfsdk:"operating_mode" json:"operatingMode,omitempty"`
			} `tfsdk:"pdns_server" json:"pdnsServer,omitempty"`
			Saslauthd *struct {
				OperatingMode *string `tfsdk:"operating_mode" json:"operatingMode,omitempty"`
			} `tfsdk:"saslauthd" json:"saslauthd,omitempty"`
			StatsArchiver *struct {
				OperatingMode *string `tfsdk:"operating_mode" json:"operatingMode,omitempty"`
			} `tfsdk:"stats_archiver" json:"statsArchiver,omitempty"`
		} `tfsdk:"redis_enterprise_services_configuration" json:"redisEnterpriseServicesConfiguration,omitempty"`
		RedisEnterpriseServicesRiggerImageSpec *struct {
			DigestHash      *string `tfsdk:"digest_hash" json:"digestHash,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			Repository      *string `tfsdk:"repository" json:"repository,omitempty"`
			VersionTag      *string `tfsdk:"version_tag" json:"versionTag,omitempty"`
		} `tfsdk:"redis_enterprise_services_rigger_image_spec" json:"redisEnterpriseServicesRiggerImageSpec,omitempty"`
		RedisEnterpriseServicesRiggerResources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"redis_enterprise_services_rigger_resources" json:"redisEnterpriseServicesRiggerResources,omitempty"`
		RedisEnterpriseTerminationGracePeriodSeconds *int64 `tfsdk:"redis_enterprise_termination_grace_period_seconds" json:"redisEnterpriseTerminationGracePeriodSeconds,omitempty"`
		RedisEnterpriseVolumeMounts                  *[]struct {
			MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
			MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
			Name             *string `tfsdk:"name" json:"name,omitempty"`
			ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
			SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
		} `tfsdk:"redis_enterprise_volume_mounts" json:"redisEnterpriseVolumeMounts,omitempty"`
		RedisOnFlashSpec *struct {
			BigStoreDriver     *string `tfsdk:"big_store_driver" json:"bigStoreDriver,omitempty"`
			Enabled            *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			FlashDiskSize      *string `tfsdk:"flash_disk_size" json:"flashDiskSize,omitempty"`
			FlashStorageEngine *string `tfsdk:"flash_storage_engine" json:"flashStorageEngine,omitempty"`
			StorageClassName   *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
		} `tfsdk:"redis_on_flash_spec" json:"redisOnFlashSpec,omitempty"`
		RedisUpgradePolicy *string `tfsdk:"redis_upgrade_policy" json:"redisUpgradePolicy,omitempty"`
		ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		Services           *struct {
			ApiService *struct {
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"api_service" json:"apiService,omitempty"`
			ServicesAnnotations *map[string]string `tfsdk:"services_annotations" json:"servicesAnnotations,omitempty"`
		} `tfsdk:"services" json:"services,omitempty"`
		ServicesRiggerSpec *struct {
			DatabaseServiceType *string `tfsdk:"database_service_type" json:"databaseServiceType,omitempty"`
			ExtraEnvVars        *[]struct {
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
			} `tfsdk:"extra_env_vars" json:"extraEnvVars,omitempty"`
			PodAnnotations                            *map[string]string `tfsdk:"pod_annotations" json:"podAnnotations,omitempty"`
			ServiceNaming                             *string            `tfsdk:"service_naming" json:"serviceNaming,omitempty"`
			ServicesRiggerAdditionalPodSpecAttributes *struct {
				ActiveDeadlineSeconds *int64 `tfsdk:"active_deadline_seconds" json:"activeDeadlineSeconds,omitempty"`
				Affinity              *struct {
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
				AutomountServiceAccountToken *bool `tfsdk:"automount_service_account_token" json:"automountServiceAccountToken,omitempty"`
				DnsConfig                    *struct {
					Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
					Options     *[]struct {
						Name  *string `tfsdk:"name" json:"name,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"options" json:"options,omitempty"`
					Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
				} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
				DnsPolicy           *string `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
				EnableServiceLinks  *bool   `tfsdk:"enable_service_links" json:"enableServiceLinks,omitempty"`
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
				HostAliases *[]struct {
					Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
					Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
				} `tfsdk:"host_aliases" json:"hostAliases,omitempty"`
				HostIPC          *bool   `tfsdk:"host_ipc" json:"hostIPC,omitempty"`
				HostNetwork      *bool   `tfsdk:"host_network" json:"hostNetwork,omitempty"`
				HostPID          *bool   `tfsdk:"host_pid" json:"hostPID,omitempty"`
				HostUsers        *bool   `tfsdk:"host_users" json:"hostUsers,omitempty"`
				Hostname         *string `tfsdk:"hostname" json:"hostname,omitempty"`
				ImagePullSecrets *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
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
				NodeName     *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Os           *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"os" json:"os,omitempty"`
				Overhead          *map[string]string `tfsdk:"overhead" json:"overhead,omitempty"`
				PreemptionPolicy  *string            `tfsdk:"preemption_policy" json:"preemptionPolicy,omitempty"`
				Priority          *int64             `tfsdk:"priority" json:"priority,omitempty"`
				PriorityClassName *string            `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
				ReadinessGates    *[]struct {
					ConditionType *string `tfsdk:"condition_type" json:"conditionType,omitempty"`
				} `tfsdk:"readiness_gates" json:"readinessGates,omitempty"`
				ResourceClaims *[]struct {
					Name   *string `tfsdk:"name" json:"name,omitempty"`
					Source *struct {
						ResourceClaimName         *string `tfsdk:"resource_claim_name" json:"resourceClaimName,omitempty"`
						ResourceClaimTemplateName *string `tfsdk:"resource_claim_template_name" json:"resourceClaimTemplateName,omitempty"`
					} `tfsdk:"source" json:"source,omitempty"`
				} `tfsdk:"resource_claims" json:"resourceClaims,omitempty"`
				RestartPolicy    *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
				RuntimeClassName *string `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
				SchedulerName    *string `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
				SchedulingGates  *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"scheduling_gates" json:"schedulingGates,omitempty"`
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
				ServiceAccount                *string `tfsdk:"service_account" json:"serviceAccount,omitempty"`
				ServiceAccountName            *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
				SetHostnameAsFQDN             *bool   `tfsdk:"set_hostname_as_fqdn" json:"setHostnameAsFQDN,omitempty"`
				ShareProcessNamespace         *bool   `tfsdk:"share_process_namespace" json:"shareProcessNamespace,omitempty"`
				Subdomain                     *string `tfsdk:"subdomain" json:"subdomain,omitempty"`
				TerminationGracePeriodSeconds *int64  `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
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
			} `tfsdk:"services_rigger_additional_pod_spec_attributes" json:"servicesRiggerAdditionalPodSpecAttributes,omitempty"`
		} `tfsdk:"services_rigger_spec" json:"servicesRiggerSpec,omitempty"`
		SideContainersSpec *[]struct {
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
		} `tfsdk:"side_containers_spec" json:"sideContainersSpec,omitempty"`
		SlaveHA *struct {
			SlaveHAGracePeriod *int64 `tfsdk:"slave_ha_grace_period" json:"slaveHAGracePeriod,omitempty"`
		} `tfsdk:"slave_ha" json:"slaveHA,omitempty"`
		UiAnnotations *map[string]string `tfsdk:"ui_annotations" json:"uiAnnotations,omitempty"`
		UiServiceType *string            `tfsdk:"ui_service_type" json:"uiServiceType,omitempty"`
		UpgradeSpec   *struct {
			AutoUpgradeRedisEnterprise *bool `tfsdk:"auto_upgrade_redis_enterprise" json:"autoUpgradeRedisEnterprise,omitempty"`
		} `tfsdk:"upgrade_spec" json:"upgradeSpec,omitempty"`
		Username      *string `tfsdk:"username" json:"username,omitempty"`
		VaultCASecret *string `tfsdk:"vault_ca_secret" json:"vaultCASecret,omitempty"`
		Volumes       *[]struct {
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
}

func (r *AppRedislabsComRedisEnterpriseClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_app_redislabs_com_redis_enterprise_cluster_v1alpha1_manifest"
}

func (r *AppRedislabsComRedisEnterpriseClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"active_active": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"api_ingress_url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"db_ingress_suffix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"ingress_annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("openShiftRoute", "ingress", "istio"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"anti_affinity_additional_topology_keys": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"bootstrapper_image_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"digest_hash": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version_tag": schema.StringAttribute{
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

					"bootstrapper_resources": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

					"certificates": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"api_certificate_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cm_certificate_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ldap_client_certificate_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"metrics_exporter_certificate_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"proxy_certificate_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"syncer_certificate_secret_name": schema.StringAttribute{
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

					"cluster_credential_secret_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_credential_secret_role": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_credential_secret_type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("vault", "kubernetes"),
						},
					},

					"cluster_recovery": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"container_timezone": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"propagate_host": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timezone_name": schema.StringAttribute{
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

					"create_service_account": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"data_internode_encryption": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"encrypt_pkeys": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enforce_i_pv4": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"extra_env_vars": schema.ListNestedAttribute{
						Description:         "ADVANCED USAGE: use carefully. Add environment variables to RS StatefulSet's containers.",
						MarkdownDescription: "ADVANCED USAGE: use carefully. Add environment variables to RS StatefulSet's containers.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
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
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"config_map_key_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"field_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"resource_field_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"container_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"divisor": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource": schema.StringAttribute{
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
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"extra_labels": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_aliases": schema.ListNestedAttribute{
						Description:         "Adds hostAliases entries to the Redis Enterprise pods",
						MarkdownDescription: "Adds hostAliases entries to the Redis Enterprise pods",
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

					"ingress_or_route_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"api_fqdn_url": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"db_fqdn_suffix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"ingress_annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("openShiftRoute", "ingress", "istio"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ldap": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"authentication_query": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"query": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"base": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"filter": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"scope": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("BaseObject", "SingleLevel", "WholeSubtree"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"template": schema.StringAttribute{
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

							"authorization_query": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"attribute": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"base": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"filter": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"scope": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("BaseObject", "SingleLevel", "WholeSubtree"),
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

							"bind_credentials_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ca_certificate_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_ttl_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled_for_control_plane": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled_for_data_plane": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protocol": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("LDAP", "LDAPS", "STARTTLS"),
								},
							},

							"servers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"host": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"port": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"license": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"license_secret_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_selector": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nodes": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ocsp_configuration": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"ocsp_functionality": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"query_frequency": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"recovery_frequency": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"recovery_max_tries": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"response_timeout": schema.Int64Attribute{
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

					"persistent_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_class_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_size": schema.StringAttribute{
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

					"pod_annotations": schema.MapAttribute{
						Description:         "annotations for the service rigger and redis enterprise pods",
						MarkdownDescription: "annotations for the service rigger and redis enterprise pods",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
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

					"pod_security_policy_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_starting_policy": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"starting_threshold_seconds": schema.Int64Attribute{
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

					"pod_tolerations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
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

					"priority_class_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pull_secrets": schema.ListNestedAttribute{
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

					"rack_awareness_node_label": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redis_enterprise_additional_pod_spec_attributes": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"active_deadline_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"affinity": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
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
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
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

							"automount_service_account_token": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"dns_config": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"nameservers": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"options": schema.ListNestedAttribute{
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

									"searches": schema.ListAttribute{
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

							"dns_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_service_links": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ephemeral_containers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"args": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"command": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"env": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"config_map_key_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
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

															"field_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
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

															"resource_field_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"container_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"divisor": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"resource": schema.StringAttribute{
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
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"env_from": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
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

													"prefix": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"image": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image_pull_policy": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"lifecycle": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"post_start": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"grpc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
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

												"http_get": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
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

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"ports": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"host_ip": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host_port": schema.Int64Attribute{
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

													"protocol": schema.StringAttribute{
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

										"readiness_probe": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"grpc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
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

												"http_get": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
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
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"claims": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
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

										"security_context": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"startup_probe": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"grpc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
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

												"http_get": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
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

										"stdin": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"stdin_once": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"target_container_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"termination_message_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"termination_message_policy": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tty": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_devices": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"device_path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
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

										"volume_mounts": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"mount_path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
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

										"working_dir": schema.StringAttribute{
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

							"host_aliases": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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

							"host_ipc": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_network": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_pid": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_users": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"hostname": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_secrets": schema.ListNestedAttribute{
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

							"init_containers": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"args": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"command": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"env": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
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
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"config_map_key_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
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

															"field_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
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

															"resource_field_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"container_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"divisor": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"resource": schema.StringAttribute{
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
																Description:         "",
																MarkdownDescription: "",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"env_from": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
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

													"prefix": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"image": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image_pull_policy": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"lifecycle": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"post_start": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"grpc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
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

												"http_get": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
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

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"ports": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_port": schema.Int64Attribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"host_ip": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host_port": schema.Int64Attribute{
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

													"protocol": schema.StringAttribute{
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

										"readiness_probe": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"grpc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
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

												"http_get": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
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
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"claims": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
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

										"security_context": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"startup_probe": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"grpc": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
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

												"http_get": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"termination_grace_period_seconds": schema.Int64Attribute{
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

										"stdin": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"stdin_once": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"termination_message_path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"termination_message_policy": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tty": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_devices": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"device_path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
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

										"volume_mounts": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"mount_path": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
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

										"working_dir": schema.StringAttribute{
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

							"node_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"os": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
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

							"overhead": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"preemption_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_class_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"readiness_gates": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"condition_type": schema.StringAttribute{
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

							"resource_claims": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"source": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"resource_claim_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource_claim_template_name": schema.StringAttribute{
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

							"restart_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"runtime_class_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduling_gates": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

							"security_context": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"value": schema.StringAttribute{
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

							"service_account": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"set_hostname_as_fqdn": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"share_process_namespace": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"subdomain": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"termination_grace_period_seconds": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
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

							"volumes": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"aws_elastic_block_store": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"partition": schema.Int64Attribute{
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

												"volume_id": schema.StringAttribute{
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

										"azure_disk": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"caching_mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"disk_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"disk_uri": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"kind": schema.StringAttribute{
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

										"azure_file": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"read_only": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"share_name": schema.StringAttribute{
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

										"cephfs": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"monitors": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"path": schema.StringAttribute{
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

												"secret_file": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"cinder": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
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

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"volume_id": schema.StringAttribute{
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

										"config_map": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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
																Required:            true,
																Optional:            false,
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

										"csi": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"driver": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_publish_secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"read_only": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_attributes": schema.MapAttribute{
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

										"downward_api": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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
															"field_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
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
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"resource_field_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"container_name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"divisor": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"resource": schema.StringAttribute{
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
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"medium": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size_limit": schema.StringAttribute{
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

										"ephemeral": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"volume_claim_template": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"metadata": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"spec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"access_modes": schema.ListAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"data_source": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"api_group": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"kind": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
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

																"data_source_ref": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"api_group": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"kind": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
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

																"resources": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
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

																"selector": schema.SingleNestedAttribute{
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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

																"storage_class_name": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_mode": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_name": schema.StringAttribute{
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
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"lun": schema.Int64Attribute{
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

												"target_ww_ns": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"wwids": schema.ListAttribute{
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

										"flex_volume": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"driver": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"options": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
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

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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
											Required: false,
											Optional: true,
											Computed: false,
										},

										"flocker": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"dataset_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"dataset_uuid": schema.StringAttribute{
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

										"gce_persistent_disk": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"partition": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pd_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"git_repo": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"directory": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"repository": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"revision": schema.StringAttribute{
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

										"glusterfs": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"endpoints": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"host_path": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
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

										"iscsi": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"chap_auth_discovery": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"chap_auth_session": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"initiator_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"iqn": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"iscsi_interface": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"lun": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"portals": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
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

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"target_portal": schema.StringAttribute{
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

										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"nfs": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"server": schema.StringAttribute{
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

										"persistent_volume_claim": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"claim_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"photon_persistent_disk": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pd_id": schema.StringAttribute{
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

										"portworx_volume": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
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

												"volume_id": schema.StringAttribute{
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

										"projected": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"default_mode": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"sources": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config_map": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"items": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
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

															"downward_api": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"items": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"field_ref": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
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
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"resource_field_ref": schema.SingleNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Attributes: map[string]schema.Attribute{
																						"container_name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"divisor": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"resource": schema.StringAttribute{
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
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"items": schema.ListNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"key": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
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

															"service_account_token": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"audience": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"expiration_seconds": schema.Int64Attribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"path": schema.StringAttribute{
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
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"group": schema.StringAttribute{
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

												"registry": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"tenant": schema.StringAttribute{
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

												"volume": schema.StringAttribute{
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

										"rbd": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"keyring": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"monitors": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"pool": schema.StringAttribute{
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

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

										"scale_io": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"gateway": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"protection_domain": schema.StringAttribute{
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

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
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

												"ssl_enabled": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"storage_mode": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"storage_pool": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"system": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"volume_name": schema.StringAttribute{
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
											Description:         "",
											MarkdownDescription: "",
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
																Required:            true,
																Optional:            false,
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

										"storageos": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
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

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"volume_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_namespace": schema.StringAttribute{
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

										"vsphere_volume": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"storage_policy_id": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"storage_policy_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_path": schema.StringAttribute{
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

					"redis_enterprise_image_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"digest_hash": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version_tag": schema.StringAttribute{
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

					"redis_enterprise_node_resources": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

					"redis_enterprise_pod_annotations": schema.MapAttribute{
						Description:         "annotations for redis enterprise pod",
						MarkdownDescription: "annotations for redis enterprise pod",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redis_enterprise_services_configuration": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cm_server": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"operating_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("enabled", "disabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"crdb_coordinator": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"operating_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("enabled", "disabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"crdb_worker": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"operating_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("enabled", "disabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"mdns_server": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"operating_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("enabled", "disabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pdns_server": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"operating_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("enabled", "disabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"saslauthd": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"operating_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("enabled", "disabled"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"stats_archiver": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"operating_mode": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("enabled", "disabled"),
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

					"redis_enterprise_services_rigger_image_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"digest_hash": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repository": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version_tag": schema.StringAttribute{
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

					"redis_enterprise_services_rigger_resources": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

					"redis_enterprise_termination_grace_period_seconds": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redis_enterprise_volume_mounts": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mount_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
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

					"redis_on_flash_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"big_store_driver": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("rocksdb", "speedb"),
								},
							},

							"enabled": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"flash_disk_size": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"flash_storage_engine": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("rocksdb"),
								},
							},

							"storage_class_name": schema.StringAttribute{
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

					"redis_upgrade_policy": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("major", "latest"),
						},
					},

					"service_account_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"services": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"api_service": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
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

							"services_annotations": schema.MapAttribute{
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

					"services_rigger_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"database_service_type": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extra_env_vars": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_map_key_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"field_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"resource_field_ref": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"container_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"divisor": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resource": schema.StringAttribute{
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
													Description:         "",
													MarkdownDescription: "",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_annotations": schema.MapAttribute{
								Description:         "annotations for the service rigger pod",
								MarkdownDescription: "annotations for the service rigger pod",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_naming": schema.StringAttribute{
								Description:         "Used to determine how to name the services created automatically when a database is created. When bdb_name is used, the database name will be also used for the service name. When redis-port is used, the service will be named redis-<port>.",
								MarkdownDescription: "Used to determine how to name the services created automatically when a database is created. When bdb_name is used, the database name will be also used for the service name. When redis-port is used, the service will be named redis-<port>.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("bdb_name", "redis-port"),
								},
							},

							"services_rigger_additional_pod_spec_attributes": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"active_deadline_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"affinity": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
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
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
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
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"operator": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
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

									"automount_service_account_token": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"dns_config": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"nameservers": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"options": schema.ListNestedAttribute{
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

											"searches": schema.ListAttribute{
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

									"dns_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable_service_links": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_containers": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"args": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"command": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"env": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"config_map_key_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

																	"field_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

																	"resource_field_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"container_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"divisor": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"resource": schema.StringAttribute{
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
																		Description:         "",
																		MarkdownDescription: "",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"env_from": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config_map_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
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

															"prefix": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"image": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_pull_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"lifecycle": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"post_start": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

																"http_get": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
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

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
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
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

																"http_get": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
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

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"grpc": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service": schema.StringAttribute{
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"termination_grace_period_seconds": schema.Int64Attribute{
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

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"ports": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"container_port": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"host_ip": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"host_port": schema.Int64Attribute{
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

															"protocol": schema.StringAttribute{
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

												"readiness_probe": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"grpc": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service": schema.StringAttribute{
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"termination_grace_period_seconds": schema.Int64Attribute{
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"claims": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
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

												"security_context": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"startup_probe": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"grpc": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service": schema.StringAttribute{
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"termination_grace_period_seconds": schema.Int64Attribute{
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

												"stdin": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"stdin_once": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"target_container_name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"termination_message_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"termination_message_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tty": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_devices": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"device_path": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
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

												"volume_mounts": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"mount_path": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
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

												"working_dir": schema.StringAttribute{
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

									"host_aliases": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
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

									"host_ipc": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_network": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_pid": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"host_users": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"hostname": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_pull_secrets": schema.ListNestedAttribute{
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

									"init_containers": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"args": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"command": schema.ListAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"env": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"config_map_key_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

																	"field_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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

																	"resource_field_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"container_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"divisor": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"resource": schema.StringAttribute{
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
																		Description:         "",
																		MarkdownDescription: "",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"env_from": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"config_map_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
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

															"prefix": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_ref": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"image": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image_pull_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"lifecycle": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"post_start": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

																"http_get": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
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

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
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
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"exec": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
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

																"http_get": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"http_headers": schema.ListNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "",
																						MarkdownDescription: "",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"value": schema.StringAttribute{
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

																"tcp_socket": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"host": schema.StringAttribute{
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"grpc": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service": schema.StringAttribute{
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"termination_grace_period_seconds": schema.Int64Attribute{
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

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"ports": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"container_port": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"host_ip": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"host_port": schema.Int64Attribute{
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

															"protocol": schema.StringAttribute{
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

												"readiness_probe": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"grpc": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service": schema.StringAttribute{
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"termination_grace_period_seconds": schema.Int64Attribute{
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"claims": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
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

												"security_context": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"startup_probe": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"grpc": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"port": schema.Int64Attribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"service": schema.StringAttribute{
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

														"http_get": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
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

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
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
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"termination_grace_period_seconds": schema.Int64Attribute{
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

												"stdin": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"stdin_once": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"termination_message_path": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"termination_message_policy": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tty": schema.BoolAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_devices": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"device_path": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
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

												"volume_mounts": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"mount_path": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
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

												"working_dir": schema.StringAttribute{
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

									"node_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"node_selector": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"os": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
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

									"overhead": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"preemption_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"priority": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"priority_class_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"readiness_gates": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"condition_type": schema.StringAttribute{
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

									"resource_claims": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"source": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"resource_claim_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"resource_claim_template_name": schema.StringAttribute{
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

									"restart_policy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"runtime_class_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheduler_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheduling_gates": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
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

									"security_context": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
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
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"value": schema.StringAttribute{
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

									"service_account": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_account_name": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"set_hostname_as_fqdn": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"share_process_namespace": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"subdomain": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"termination_grace_period_seconds": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
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

									"volumes": schema.ListNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"aws_elastic_block_store": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"partition": schema.Int64Attribute{
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

														"volume_id": schema.StringAttribute{
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

												"azure_disk": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"caching_mode": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"disk_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"disk_uri": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"kind": schema.StringAttribute{
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

												"azure_file": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"read_only": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"share_name": schema.StringAttribute{
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

												"cephfs": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"monitors": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"path": schema.StringAttribute{
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

														"secret_file": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

												"cinder": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
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

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"volume_id": schema.StringAttribute{
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

												"config_map": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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
																		Required:            true,
																		Optional:            false,
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

												"csi": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"driver": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_publish_secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"read_only": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"volume_attributes": schema.MapAttribute{
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

												"downward_api": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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
																	"field_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"resource_field_ref": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"container_name": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"divisor": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"resource": schema.StringAttribute{
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"medium": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"size_limit": schema.StringAttribute{
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

												"ephemeral": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"volume_claim_template": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"metadata": schema.MapAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"spec": schema.SingleNestedAttribute{
																	Description:         "",
																	MarkdownDescription: "",
																	Attributes: map[string]schema.Attribute{
																		"access_modes": schema.ListAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"data_source": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"api_group": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"kind": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
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

																		"data_source_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"api_group": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"kind": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            true,
																					Optional:            false,
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

																		"resources": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"claims": schema.ListNestedAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
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

																		"selector": schema.SingleNestedAttribute{
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
																								Required:            true,
																								Optional:            false,
																								Computed:            false,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "",
																								MarkdownDescription: "",
																								Required:            true,
																								Optional:            false,
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

																		"storage_class_name": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_mode": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_name": schema.StringAttribute{
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"lun": schema.Int64Attribute{
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

														"target_ww_ns": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"wwids": schema.ListAttribute{
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

												"flex_volume": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"driver": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"options": schema.MapAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
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

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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
													Required: false,
													Optional: true,
													Computed: false,
												},

												"flocker": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"dataset_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"dataset_uuid": schema.StringAttribute{
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

												"gce_persistent_disk": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"partition": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pd_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"git_repo": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"directory": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"repository": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"revision": schema.StringAttribute{
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

												"glusterfs": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"endpoints": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"path": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"host_path": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"path": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
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

												"iscsi": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"chap_auth_discovery": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"chap_auth_session": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"initiator_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"iqn": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"iscsi_interface": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"lun": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"portals": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
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

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"target_portal": schema.StringAttribute{
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

												"name": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"nfs": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"path": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"server": schema.StringAttribute{
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

												"persistent_volume_claim": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"claim_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
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
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"photon_persistent_disk": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"pd_id": schema.StringAttribute{
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

												"portworx_volume": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
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

														"volume_id": schema.StringAttribute{
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

												"projected": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"default_mode": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sources": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"config_map": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"items": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            true,
																							Optional:            false,
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

																	"downward_api": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"items": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"field_ref": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
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
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"resource_field_ref": schema.SingleNestedAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Attributes: map[string]schema.Attribute{
																								"container_name": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"divisor": schema.StringAttribute{
																									Description:         "",
																									MarkdownDescription: "",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"resource": schema.StringAttribute{
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
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"items": schema.ListNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            true,
																							Optional:            false,
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

																	"service_account_token": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"audience": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"expiration_seconds": schema.Int64Attribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.StringAttribute{
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
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"group": schema.StringAttribute{
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

														"registry": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"tenant": schema.StringAttribute{
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

														"volume": schema.StringAttribute{
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

												"rbd": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"image": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"keyring": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"monitors": schema.ListAttribute{
															Description:         "",
															MarkdownDescription: "",
															ElementType:         types.StringType,
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"pool": schema.StringAttribute{
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

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

												"scale_io": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"gateway": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"protection_domain": schema.StringAttribute{
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

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
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

														"ssl_enabled": schema.BoolAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"storage_mode": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"storage_pool": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"system": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"volume_name": schema.StringAttribute{
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
													Description:         "",
													MarkdownDescription: "",
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
																		Required:            true,
																		Optional:            false,
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

												"storageos": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
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

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
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

														"volume_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"volume_namespace": schema.StringAttribute{
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

												"vsphere_volume": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"storage_policy_id": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"storage_policy_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"volume_path": schema.StringAttribute{
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

					"side_containers_spec": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"args": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"command": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"env": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
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
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"config_map_key_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"field_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
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

													"resource_field_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"container_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"divisor": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resource": schema.StringAttribute{
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
														Description:         "",
														MarkdownDescription: "",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"env_from": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"config_map_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
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

											"prefix": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"image": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"image_pull_policy": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"lifecycle": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"post_start": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"http_get": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
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
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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

												"http_get": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
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
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"grpc": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"service": schema.StringAttribute{
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

										"http_get": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http_headers": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.StringAttribute{
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

										"tcp_socket": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"termination_grace_period_seconds": schema.Int64Attribute{
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

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"ports": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"container_port": schema.Int64Attribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"host_ip": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"host_port": schema.Int64Attribute{
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

											"protocol": schema.StringAttribute{
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

								"readiness_probe": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"grpc": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"service": schema.StringAttribute{
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

										"http_get": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http_headers": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.StringAttribute{
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

										"tcp_socket": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"termination_grace_period_seconds": schema.Int64Attribute{
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
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"claims": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
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

								"security_context": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
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

								"startup_probe": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"exec": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"grpc": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"port": schema.Int64Attribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"service": schema.StringAttribute{
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

										"http_get": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"http_headers": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"value": schema.StringAttribute{
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

										"tcp_socket": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"host": schema.StringAttribute{
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"termination_grace_period_seconds": schema.Int64Attribute{
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

								"stdin": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"stdin_once": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"termination_message_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"termination_message_policy": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tty": schema.BoolAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"volume_devices": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"device_path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
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

								"volume_mounts": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"mount_path": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
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

								"working_dir": schema.StringAttribute{
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

					"slave_ha": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"slave_ha_grace_period": schema.Int64Attribute{
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

					"ui_annotations": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ui_service_type": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("ClusterIP", "NodePort", "LoadBalancer", "ExternalName"),
						},
					},

					"upgrade_spec": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"auto_upgrade_redis_enterprise": schema.BoolAttribute{
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

					"username": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"vault_ca_secret": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volumes": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"aws_elastic_block_store": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"partition": schema.Int64Attribute{
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

										"volume_id": schema.StringAttribute{
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

								"azure_disk": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"caching_mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"disk_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"disk_uri": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"kind": schema.StringAttribute{
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

								"azure_file": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"read_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"share_name": schema.StringAttribute{
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

								"cephfs": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"monitors": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"path": schema.StringAttribute{
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

										"secret_file": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"cinder": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
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

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"volume_id": schema.StringAttribute{
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

								"config_map": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
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
														Required:            true,
														Optional:            false,
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

								"csi": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"driver": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_publish_secret_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"read_only": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_attributes": schema.MapAttribute{
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

								"downward_api": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
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
													"field_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
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
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"resource_field_ref": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"container_name": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"divisor": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resource": schema.StringAttribute{
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
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"medium": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size_limit": schema.StringAttribute{
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

								"fc": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"lun": schema.Int64Attribute{
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

										"target_ww_ns": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"wwids": schema.ListAttribute{
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

								"flex_volume": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"driver": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"options": schema.MapAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
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

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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
									Required: false,
									Optional: true,
									Computed: false,
								},

								"flocker": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"dataset_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"dataset_uuid": schema.StringAttribute{
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

								"gce_persistent_disk": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"partition": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pd_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"git_repo": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"directory": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"repository": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"revision": schema.StringAttribute{
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

								"glusterfs": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"endpoints": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"host_path": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
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

								"iscsi": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"chap_auth_discovery": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"chap_auth_session": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"initiator_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"iqn": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"iscsi_interface": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"lun": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"portals": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
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

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"target_portal": schema.StringAttribute{
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

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"nfs": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"path": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"server": schema.StringAttribute{
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

								"persistent_volume_claim": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"claim_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"photon_persistent_disk": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pd_id": schema.StringAttribute{
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

								"portworx_volume": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
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

										"volume_id": schema.StringAttribute{
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

								"projected": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"default_mode": schema.Int64Attribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sources": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config_map": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"items": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
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

													"downward_api": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"items": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"field_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
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
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"resource_field_ref": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"container_name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"divisor": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"resource": schema.StringAttribute{
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
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"items": schema.ListNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Required:            true,
																			Optional:            false,
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

													"service_account_token": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"audience": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"expiration_seconds": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
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

								"quobyte": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"group": schema.StringAttribute{
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

										"registry": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"tenant": schema.StringAttribute{
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

										"volume": schema.StringAttribute{
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

								"rbd": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"keyring": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"monitors": schema.ListAttribute{
											Description:         "",
											MarkdownDescription: "",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"pool": schema.StringAttribute{
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

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

								"scale_io": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"gateway": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"protection_domain": schema.StringAttribute{
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

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
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

										"ssl_enabled": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"storage_mode": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"storage_pool": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"system": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"volume_name": schema.StringAttribute{
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
									Description:         "",
									MarkdownDescription: "",
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
														Required:            true,
														Optional:            false,
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

								"storageos": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
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

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
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

										"volume_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_namespace": schema.StringAttribute{
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

								"vsphere_volume": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"storage_policy_id": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"storage_policy_name": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_path": schema.StringAttribute{
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

func (r *AppRedislabsComRedisEnterpriseClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_redislabs_com_redis_enterprise_cluster_v1alpha1_manifest")

	var model AppRedislabsComRedisEnterpriseClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("app.redislabs.com/v1alpha1")
	model.Kind = pointer.String("RedisEnterpriseCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
