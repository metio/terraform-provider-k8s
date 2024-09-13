/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package asdb_aerospike_com_v1

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
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AsdbAerospikeComAerospikeClusterV1Manifest{}
)

func NewAsdbAerospikeComAerospikeClusterV1Manifest() datasource.DataSource {
	return &AsdbAerospikeComAerospikeClusterV1Manifest{}
}

type AsdbAerospikeComAerospikeClusterV1Manifest struct{}

type AsdbAerospikeComAerospikeClusterV1ManifestData struct {
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
		AerospikeAccessControl *struct {
			AdminPolicy *struct {
				Timeout *int64 `tfsdk:"timeout" json:"timeout,omitempty"`
			} `tfsdk:"admin_policy" json:"adminPolicy,omitempty"`
			Roles *[]struct {
				Name       *string   `tfsdk:"name" json:"name,omitempty"`
				Privileges *[]string `tfsdk:"privileges" json:"privileges,omitempty"`
				ReadQuota  *int64    `tfsdk:"read_quota" json:"readQuota,omitempty"`
				Whitelist  *[]string `tfsdk:"whitelist" json:"whitelist,omitempty"`
				WriteQuota *int64    `tfsdk:"write_quota" json:"writeQuota,omitempty"`
			} `tfsdk:"roles" json:"roles,omitempty"`
			Users *[]struct {
				Name       *string   `tfsdk:"name" json:"name,omitempty"`
				Roles      *[]string `tfsdk:"roles" json:"roles,omitempty"`
				SecretName *string   `tfsdk:"secret_name" json:"secretName,omitempty"`
			} `tfsdk:"users" json:"users,omitempty"`
		} `tfsdk:"aerospike_access_control" json:"aerospikeAccessControl,omitempty"`
		AerospikeConfig        *map[string]string `tfsdk:"aerospike_config" json:"aerospikeConfig,omitempty"`
		AerospikeNetworkPolicy *struct {
			Access                               *string   `tfsdk:"access" json:"access,omitempty"`
			AlternateAccess                      *string   `tfsdk:"alternate_access" json:"alternateAccess,omitempty"`
			CustomAccessNetworkNames             *[]string `tfsdk:"custom_access_network_names" json:"customAccessNetworkNames,omitempty"`
			CustomAlternateAccessNetworkNames    *[]string `tfsdk:"custom_alternate_access_network_names" json:"customAlternateAccessNetworkNames,omitempty"`
			CustomFabricNetworkNames             *[]string `tfsdk:"custom_fabric_network_names" json:"customFabricNetworkNames,omitempty"`
			CustomTLSAccessNetworkNames          *[]string `tfsdk:"custom_tls_access_network_names" json:"customTLSAccessNetworkNames,omitempty"`
			CustomTLSAlternateAccessNetworkNames *[]string `tfsdk:"custom_tls_alternate_access_network_names" json:"customTLSAlternateAccessNetworkNames,omitempty"`
			CustomTLSFabricNetworkNames          *[]string `tfsdk:"custom_tls_fabric_network_names" json:"customTLSFabricNetworkNames,omitempty"`
			Fabric                               *string   `tfsdk:"fabric" json:"fabric,omitempty"`
			TlsAccess                            *string   `tfsdk:"tls_access" json:"tlsAccess,omitempty"`
			TlsAlternateAccess                   *string   `tfsdk:"tls_alternate_access" json:"tlsAlternateAccess,omitempty"`
			TlsFabric                            *string   `tfsdk:"tls_fabric" json:"tlsFabric,omitempty"`
		} `tfsdk:"aerospike_network_policy" json:"aerospikeNetworkPolicy,omitempty"`
		DisablePDB                *bool     `tfsdk:"disable_pdb" json:"disablePDB,omitempty"`
		EnableDynamicConfigUpdate *bool     `tfsdk:"enable_dynamic_config_update" json:"enableDynamicConfigUpdate,omitempty"`
		Image                     *string   `tfsdk:"image" json:"image,omitempty"`
		K8sNodeBlockList          *[]string `tfsdk:"k8s_node_block_list" json:"k8sNodeBlockList,omitempty"`
		MaxUnavailable            *string   `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
		Operations                *[]struct {
			Id      *string   `tfsdk:"id" json:"id,omitempty"`
			Kind    *string   `tfsdk:"kind" json:"kind,omitempty"`
			PodList *[]string `tfsdk:"pod_list" json:"podList,omitempty"`
		} `tfsdk:"operations" json:"operations,omitempty"`
		OperatorClientCert *struct {
			CertPathInOperator *struct {
				CaCertsPath    *string `tfsdk:"ca_certs_path" json:"caCertsPath,omitempty"`
				ClientCertPath *string `tfsdk:"client_cert_path" json:"clientCertPath,omitempty"`
				ClientKeyPath  *string `tfsdk:"client_key_path" json:"clientKeyPath,omitempty"`
			} `tfsdk:"cert_path_in_operator" json:"certPathInOperator,omitempty"`
			SecretCertSource *struct {
				CaCertsFilename *string `tfsdk:"ca_certs_filename" json:"caCertsFilename,omitempty"`
				CaCertsSource   *struct {
					SecretName      *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					SecretNamespace *string `tfsdk:"secret_namespace" json:"secretNamespace,omitempty"`
				} `tfsdk:"ca_certs_source" json:"caCertsSource,omitempty"`
				ClientCertFilename *string `tfsdk:"client_cert_filename" json:"clientCertFilename,omitempty"`
				ClientKeyFilename  *string `tfsdk:"client_key_filename" json:"clientKeyFilename,omitempty"`
				SecretName         *string `tfsdk:"secret_name" json:"secretName,omitempty"`
				SecretNamespace    *string `tfsdk:"secret_namespace" json:"secretNamespace,omitempty"`
			} `tfsdk:"secret_cert_source" json:"secretCertSource,omitempty"`
			TlsClientName *string `tfsdk:"tls_client_name" json:"tlsClientName,omitempty"`
		} `tfsdk:"operator_client_cert" json:"operatorClientCert,omitempty"`
		Paused  *bool `tfsdk:"paused" json:"paused,omitempty"`
		PodSpec *struct {
			AerospikeContainer *struct {
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
			} `tfsdk:"aerospike_container" json:"aerospikeContainer,omitempty"`
			AerospikeInitContainer *struct {
				ImageNameAndTag        *string `tfsdk:"image_name_and_tag" json:"imageNameAndTag,omitempty"`
				ImageRegistry          *string `tfsdk:"image_registry" json:"imageRegistry,omitempty"`
				ImageRegistryNamespace *string `tfsdk:"image_registry_namespace" json:"imageRegistryNamespace,omitempty"`
				Resources              *struct {
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
			} `tfsdk:"aerospike_init_container" json:"aerospikeInitContainer,omitempty"`
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
			DnsConfig *struct {
				Nameservers *[]string `tfsdk:"nameservers" json:"nameservers,omitempty"`
				Options     *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"options" json:"options,omitempty"`
				Searches *[]string `tfsdk:"searches" json:"searches,omitempty"`
			} `tfsdk:"dns_config" json:"dnsConfig,omitempty"`
			DnsPolicy          *string `tfsdk:"dns_policy" json:"dnsPolicy,omitempty"`
			EffectiveDNSPolicy *string `tfsdk:"effective_dns_policy" json:"effectiveDNSPolicy,omitempty"`
			HostNetwork        *bool   `tfsdk:"host_network" json:"hostNetwork,omitempty"`
			ImagePullSecrets   *[]struct {
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
						Sleep *struct {
							Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
						} `tfsdk:"sleep" json:"sleep,omitempty"`
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
						Sleep *struct {
							Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
						} `tfsdk:"sleep" json:"sleep,omitempty"`
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
				RestartPolicy   *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
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
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
			MultiPodPerHost *bool              `tfsdk:"multi_pod_per_host" json:"multiPodPerHost,omitempty"`
			NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
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
			Sidecars *[]struct {
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
						Sleep *struct {
							Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
						} `tfsdk:"sleep" json:"sleep,omitempty"`
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
						Sleep *struct {
							Seconds *int64 `tfsdk:"seconds" json:"seconds,omitempty"`
						} `tfsdk:"sleep" json:"sleep,omitempty"`
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
				RestartPolicy   *string `tfsdk:"restart_policy" json:"restartPolicy,omitempty"`
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
			} `tfsdk:"sidecars" json:"sidecars,omitempty"`
			Tolerations *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		} `tfsdk:"pod_spec" json:"podSpec,omitempty"`
		RackConfig *struct {
			MaxIgnorablePods *string   `tfsdk:"max_ignorable_pods" json:"maxIgnorablePods,omitempty"`
			Namespaces       *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
			Racks            *[]struct {
				AerospikeConfig          *map[string]string `tfsdk:"aerospike_config" json:"aerospikeConfig,omitempty"`
				EffectiveAerospikeConfig *map[string]string `tfsdk:"effective_aerospike_config" json:"effectiveAerospikeConfig,omitempty"`
				EffectivePodSpec         *struct {
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
					NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					Tolerations  *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"effective_pod_spec" json:"effectivePodSpec,omitempty"`
				EffectiveStorage *struct {
					BlockVolumePolicy *struct {
						CascadeDelete          *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
						EffectiveCascadeDelete *bool   `tfsdk:"effective_cascade_delete" json:"effectiveCascadeDelete,omitempty"`
						EffectiveInitMethod    *string `tfsdk:"effective_init_method" json:"effectiveInitMethod,omitempty"`
						EffectiveWipeMethod    *string `tfsdk:"effective_wipe_method" json:"effectiveWipeMethod,omitempty"`
						InitMethod             *string `tfsdk:"init_method" json:"initMethod,omitempty"`
						WipeMethod             *string `tfsdk:"wipe_method" json:"wipeMethod,omitempty"`
					} `tfsdk:"block_volume_policy" json:"blockVolumePolicy,omitempty"`
					CleanupThreads         *int64 `tfsdk:"cleanup_threads" json:"cleanupThreads,omitempty"`
					FilesystemVolumePolicy *struct {
						CascadeDelete          *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
						EffectiveCascadeDelete *bool   `tfsdk:"effective_cascade_delete" json:"effectiveCascadeDelete,omitempty"`
						EffectiveInitMethod    *string `tfsdk:"effective_init_method" json:"effectiveInitMethod,omitempty"`
						EffectiveWipeMethod    *string `tfsdk:"effective_wipe_method" json:"effectiveWipeMethod,omitempty"`
						InitMethod             *string `tfsdk:"init_method" json:"initMethod,omitempty"`
						WipeMethod             *string `tfsdk:"wipe_method" json:"wipeMethod,omitempty"`
					} `tfsdk:"filesystem_volume_policy" json:"filesystemVolumePolicy,omitempty"`
					LocalStorageClasses *[]string `tfsdk:"local_storage_classes" json:"localStorageClasses,omitempty"`
					Volumes             *[]struct {
						Aerospike *struct {
							MountOptions *struct {
								MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" json:"mountOptions,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"aerospike" json:"aerospike,omitempty"`
						CascadeDelete          *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
						EffectiveCascadeDelete *bool   `tfsdk:"effective_cascade_delete" json:"effectiveCascadeDelete,omitempty"`
						EffectiveInitMethod    *string `tfsdk:"effective_init_method" json:"effectiveInitMethod,omitempty"`
						EffectiveWipeMethod    *string `tfsdk:"effective_wipe_method" json:"effectiveWipeMethod,omitempty"`
						InitContainers         *[]struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							MountOptions  *struct {
								MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" json:"mountOptions,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"init_containers" json:"initContainers,omitempty"`
						InitMethod *string `tfsdk:"init_method" json:"initMethod,omitempty"`
						Name       *string `tfsdk:"name" json:"name,omitempty"`
						Sidecars   *[]struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							MountOptions  *struct {
								MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" json:"mountOptions,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"sidecars" json:"sidecars,omitempty"`
						Source *struct {
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
								SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
							} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
							PersistentVolume *struct {
								AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
								Metadata    *struct {
									Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
									Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								Selector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"selector" json:"selector,omitempty"`
								Size         *string `tfsdk:"size" json:"size,omitempty"`
								StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
								VolumeMode   *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
							} `tfsdk:"persistent_volume" json:"persistentVolume,omitempty"`
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
						} `tfsdk:"source" json:"source,omitempty"`
						WipeMethod *string `tfsdk:"wipe_method" json:"wipeMethod,omitempty"`
					} `tfsdk:"volumes" json:"volumes,omitempty"`
				} `tfsdk:"effective_storage" json:"effectiveStorage,omitempty"`
				Id       *int64  `tfsdk:"id" json:"id,omitempty"`
				NodeName *string `tfsdk:"node_name" json:"nodeName,omitempty"`
				PodSpec  *struct {
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
					NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					Tolerations  *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"pod_spec" json:"podSpec,omitempty"`
				RackLabel *string `tfsdk:"rack_label" json:"rackLabel,omitempty"`
				Region    *string `tfsdk:"region" json:"region,omitempty"`
				Storage   *struct {
					BlockVolumePolicy *struct {
						CascadeDelete          *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
						EffectiveCascadeDelete *bool   `tfsdk:"effective_cascade_delete" json:"effectiveCascadeDelete,omitempty"`
						EffectiveInitMethod    *string `tfsdk:"effective_init_method" json:"effectiveInitMethod,omitempty"`
						EffectiveWipeMethod    *string `tfsdk:"effective_wipe_method" json:"effectiveWipeMethod,omitempty"`
						InitMethod             *string `tfsdk:"init_method" json:"initMethod,omitempty"`
						WipeMethod             *string `tfsdk:"wipe_method" json:"wipeMethod,omitempty"`
					} `tfsdk:"block_volume_policy" json:"blockVolumePolicy,omitempty"`
					CleanupThreads         *int64 `tfsdk:"cleanup_threads" json:"cleanupThreads,omitempty"`
					FilesystemVolumePolicy *struct {
						CascadeDelete          *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
						EffectiveCascadeDelete *bool   `tfsdk:"effective_cascade_delete" json:"effectiveCascadeDelete,omitempty"`
						EffectiveInitMethod    *string `tfsdk:"effective_init_method" json:"effectiveInitMethod,omitempty"`
						EffectiveWipeMethod    *string `tfsdk:"effective_wipe_method" json:"effectiveWipeMethod,omitempty"`
						InitMethod             *string `tfsdk:"init_method" json:"initMethod,omitempty"`
						WipeMethod             *string `tfsdk:"wipe_method" json:"wipeMethod,omitempty"`
					} `tfsdk:"filesystem_volume_policy" json:"filesystemVolumePolicy,omitempty"`
					LocalStorageClasses *[]string `tfsdk:"local_storage_classes" json:"localStorageClasses,omitempty"`
					Volumes             *[]struct {
						Aerospike *struct {
							MountOptions *struct {
								MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" json:"mountOptions,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"aerospike" json:"aerospike,omitempty"`
						CascadeDelete          *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
						EffectiveCascadeDelete *bool   `tfsdk:"effective_cascade_delete" json:"effectiveCascadeDelete,omitempty"`
						EffectiveInitMethod    *string `tfsdk:"effective_init_method" json:"effectiveInitMethod,omitempty"`
						EffectiveWipeMethod    *string `tfsdk:"effective_wipe_method" json:"effectiveWipeMethod,omitempty"`
						InitContainers         *[]struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							MountOptions  *struct {
								MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" json:"mountOptions,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"init_containers" json:"initContainers,omitempty"`
						InitMethod *string `tfsdk:"init_method" json:"initMethod,omitempty"`
						Name       *string `tfsdk:"name" json:"name,omitempty"`
						Sidecars   *[]struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							MountOptions  *struct {
								MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" json:"mountOptions,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"sidecars" json:"sidecars,omitempty"`
						Source *struct {
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
								SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
							} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
							PersistentVolume *struct {
								AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
								Metadata    *struct {
									Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
									Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								Selector *struct {
									MatchExpressions *[]struct {
										Key      *string   `tfsdk:"key" json:"key,omitempty"`
										Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
										Values   *[]string `tfsdk:"values" json:"values,omitempty"`
									} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
									MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
								} `tfsdk:"selector" json:"selector,omitempty"`
								Size         *string `tfsdk:"size" json:"size,omitempty"`
								StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
								VolumeMode   *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
							} `tfsdk:"persistent_volume" json:"persistentVolume,omitempty"`
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
						} `tfsdk:"source" json:"source,omitempty"`
						WipeMethod *string `tfsdk:"wipe_method" json:"wipeMethod,omitempty"`
					} `tfsdk:"volumes" json:"volumes,omitempty"`
				} `tfsdk:"storage" json:"storage,omitempty"`
				Zone *string `tfsdk:"zone" json:"zone,omitempty"`
			} `tfsdk:"racks" json:"racks,omitempty"`
			RollingUpdateBatchSize *string `tfsdk:"rolling_update_batch_size" json:"rollingUpdateBatchSize,omitempty"`
			ScaleDownBatchSize     *string `tfsdk:"scale_down_batch_size" json:"scaleDownBatchSize,omitempty"`
		} `tfsdk:"rack_config" json:"rackConfig,omitempty"`
		RosterNodeBlockList *[]string `tfsdk:"roster_node_block_list" json:"rosterNodeBlockList,omitempty"`
		SeedsFinderServices *struct {
			LoadBalancer *struct {
				Annotations              *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				ExternalTrafficPolicy    *string            `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
				LoadBalancerSourceRanges *[]string          `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
				Port                     *int64             `tfsdk:"port" json:"port,omitempty"`
				PortName                 *string            `tfsdk:"port_name" json:"portName,omitempty"`
				TargetPort               *int64             `tfsdk:"target_port" json:"targetPort,omitempty"`
			} `tfsdk:"load_balancer" json:"loadBalancer,omitempty"`
		} `tfsdk:"seeds_finder_services" json:"seedsFinderServices,omitempty"`
		Size    *int64 `tfsdk:"size" json:"size,omitempty"`
		Storage *struct {
			BlockVolumePolicy *struct {
				CascadeDelete          *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
				EffectiveCascadeDelete *bool   `tfsdk:"effective_cascade_delete" json:"effectiveCascadeDelete,omitempty"`
				EffectiveInitMethod    *string `tfsdk:"effective_init_method" json:"effectiveInitMethod,omitempty"`
				EffectiveWipeMethod    *string `tfsdk:"effective_wipe_method" json:"effectiveWipeMethod,omitempty"`
				InitMethod             *string `tfsdk:"init_method" json:"initMethod,omitempty"`
				WipeMethod             *string `tfsdk:"wipe_method" json:"wipeMethod,omitempty"`
			} `tfsdk:"block_volume_policy" json:"blockVolumePolicy,omitempty"`
			CleanupThreads         *int64 `tfsdk:"cleanup_threads" json:"cleanupThreads,omitempty"`
			FilesystemVolumePolicy *struct {
				CascadeDelete          *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
				EffectiveCascadeDelete *bool   `tfsdk:"effective_cascade_delete" json:"effectiveCascadeDelete,omitempty"`
				EffectiveInitMethod    *string `tfsdk:"effective_init_method" json:"effectiveInitMethod,omitempty"`
				EffectiveWipeMethod    *string `tfsdk:"effective_wipe_method" json:"effectiveWipeMethod,omitempty"`
				InitMethod             *string `tfsdk:"init_method" json:"initMethod,omitempty"`
				WipeMethod             *string `tfsdk:"wipe_method" json:"wipeMethod,omitempty"`
			} `tfsdk:"filesystem_volume_policy" json:"filesystemVolumePolicy,omitempty"`
			LocalStorageClasses *[]string `tfsdk:"local_storage_classes" json:"localStorageClasses,omitempty"`
			Volumes             *[]struct {
				Aerospike *struct {
					MountOptions *struct {
						MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
						ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
						SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
					} `tfsdk:"mount_options" json:"mountOptions,omitempty"`
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"aerospike" json:"aerospike,omitempty"`
				CascadeDelete          *bool   `tfsdk:"cascade_delete" json:"cascadeDelete,omitempty"`
				EffectiveCascadeDelete *bool   `tfsdk:"effective_cascade_delete" json:"effectiveCascadeDelete,omitempty"`
				EffectiveInitMethod    *string `tfsdk:"effective_init_method" json:"effectiveInitMethod,omitempty"`
				EffectiveWipeMethod    *string `tfsdk:"effective_wipe_method" json:"effectiveWipeMethod,omitempty"`
				InitContainers         *[]struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					MountOptions  *struct {
						MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
						ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
						SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
					} `tfsdk:"mount_options" json:"mountOptions,omitempty"`
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"init_containers" json:"initContainers,omitempty"`
				InitMethod *string `tfsdk:"init_method" json:"initMethod,omitempty"`
				Name       *string `tfsdk:"name" json:"name,omitempty"`
				Sidecars   *[]struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					MountOptions  *struct {
						MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
						ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
						SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
					} `tfsdk:"mount_options" json:"mountOptions,omitempty"`
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"sidecars" json:"sidecars,omitempty"`
				Source *struct {
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
						SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
					} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
					PersistentVolume *struct {
						AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
						Metadata    *struct {
							Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						Selector *struct {
							MatchExpressions *[]struct {
								Key      *string   `tfsdk:"key" json:"key,omitempty"`
								Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
								Values   *[]string `tfsdk:"values" json:"values,omitempty"`
							} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
							MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
						} `tfsdk:"selector" json:"selector,omitempty"`
						Size         *string `tfsdk:"size" json:"size,omitempty"`
						StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
						VolumeMode   *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					} `tfsdk:"persistent_volume" json:"persistentVolume,omitempty"`
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
				} `tfsdk:"source" json:"source,omitempty"`
				WipeMethod *string `tfsdk:"wipe_method" json:"wipeMethod,omitempty"`
			} `tfsdk:"volumes" json:"volumes,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		ValidationPolicy *struct {
			SkipWorkDirValidate     *bool `tfsdk:"skip_work_dir_validate" json:"skipWorkDirValidate,omitempty"`
			SkipXdrDlogFileValidate *bool `tfsdk:"skip_xdr_dlog_file_validate" json:"skipXdrDlogFileValidate,omitempty"`
		} `tfsdk:"validation_policy" json:"validationPolicy,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AsdbAerospikeComAerospikeClusterV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_asdb_aerospike_com_aerospike_cluster_v1_manifest"
}

func (r *AsdbAerospikeComAerospikeClusterV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AerospikeCluster is the schema for the AerospikeCluster API",
		MarkdownDescription: "AerospikeCluster is the schema for the AerospikeCluster API",
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
				Description:         "AerospikeClusterSpec defines the desired state of AerospikeCluster",
				MarkdownDescription: "AerospikeClusterSpec defines the desired state of AerospikeCluster",
				Attributes: map[string]schema.Attribute{
					"aerospike_access_control": schema.SingleNestedAttribute{
						Description:         "Has the Aerospike roles and users definitions. Required if aerospike cluster security is enabled.",
						MarkdownDescription: "Has the Aerospike roles and users definitions. Required if aerospike cluster security is enabled.",
						Attributes: map[string]schema.Attribute{
							"admin_policy": schema.SingleNestedAttribute{
								Description:         "AerospikeClientAdminPolicy specify the aerospike client admin policy for access control operations.",
								MarkdownDescription: "AerospikeClientAdminPolicy specify the aerospike client admin policy for access control operations.",
								Attributes: map[string]schema.Attribute{
									"timeout": schema.Int64Attribute{
										Description:         "Timeout for admin client policy in milliseconds.",
										MarkdownDescription: "Timeout for admin client policy in milliseconds.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"roles": schema.ListNestedAttribute{
								Description:         "Roles is the set of roles to allow on the Aerospike cluster.",
								MarkdownDescription: "Roles is the set of roles to allow on the Aerospike cluster.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of this role.",
											MarkdownDescription: "Name of this role.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"privileges": schema.ListAttribute{
											Description:         "Privileges granted to this role.",
											MarkdownDescription: "Privileges granted to this role.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"read_quota": schema.Int64Attribute{
											Description:         "ReadQuota specifies permitted rate of read records for current role (the value is in RPS)",
											MarkdownDescription: "ReadQuota specifies permitted rate of read records for current role (the value is in RPS)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"whitelist": schema.ListAttribute{
											Description:         "Whitelist of host address allowed for this role.",
											MarkdownDescription: "Whitelist of host address allowed for this role.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"write_quota": schema.Int64Attribute{
											Description:         "WriteQuota specifies permitted rate of write records for current role (the value is in RPS)",
											MarkdownDescription: "WriteQuota specifies permitted rate of write records for current role (the value is in RPS)",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"users": schema.ListNestedAttribute{
								Description:         "Users is the set of users to allow on the Aerospike cluster.",
								MarkdownDescription: "Users is the set of users to allow on the Aerospike cluster.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name is the user's username.",
											MarkdownDescription: "Name is the user's username.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"roles": schema.ListAttribute{
											Description:         "Roles is the list of roles granted to the user.",
											MarkdownDescription: "Roles is the list of roles granted to the user.",
											ElementType:         types.StringType,
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "SecretName has secret info created by user. User needs to create this secret from password literal.eg: kubectl create secret generic dev-db-secret --from-literal=password='password'",
											MarkdownDescription: "SecretName has secret info created by user. User needs to create this secret from password literal.eg: kubectl create secret generic dev-db-secret --from-literal=password='password'",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"aerospike_config": schema.MapAttribute{
						Description:         "Sets config in aerospike.conf file. Other configs are taken as default",
						MarkdownDescription: "Sets config in aerospike.conf file. Other configs are taken as default",
						ElementType:         types.StringType,
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"aerospike_network_policy": schema.SingleNestedAttribute{
						Description:         "AerospikeNetworkPolicy specifies how clients and tools access the Aerospike cluster.",
						MarkdownDescription: "AerospikeNetworkPolicy specifies how clients and tools access the Aerospike cluster.",
						Attributes: map[string]schema.Attribute{
							"access": schema.StringAttribute{
								Description:         "AccessType is the type of network address to use for Aerospike access address.Defaults to hostInternal.",
								MarkdownDescription: "AccessType is the type of network address to use for Aerospike access address.Defaults to hostInternal.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("pod", "hostInternal", "hostExternal", "configuredIP", "customInterface"),
								},
							},

							"alternate_access": schema.StringAttribute{
								Description:         "AlternateAccessType is the type of network address to use for Aerospike alternate access address.Defaults to hostExternal.",
								MarkdownDescription: "AlternateAccessType is the type of network address to use for Aerospike alternate access address.Defaults to hostExternal.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("pod", "hostInternal", "hostExternal", "configuredIP", "customInterface"),
								},
							},

							"custom_access_network_names": schema.ListAttribute{
								Description:         "CustomAccessNetworkNames is the list of the pod's network interfaces used for Aerospike access address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' access type.",
								MarkdownDescription: "CustomAccessNetworkNames is the list of the pod's network interfaces used for Aerospike access address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' access type.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_alternate_access_network_names": schema.ListAttribute{
								Description:         "CustomAlternateAccessNetworkNames is the list of the pod's network interfaces used for Aerospikealternate access address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' alternateAccess type",
								MarkdownDescription: "CustomAlternateAccessNetworkNames is the list of the pod's network interfaces used for Aerospikealternate access address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' alternateAccess type",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_fabric_network_names": schema.ListAttribute{
								Description:         "CustomFabricNetworkNames is the list of the pod's network interfaces used for Aerospike fabric address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' fabric type",
								MarkdownDescription: "CustomFabricNetworkNames is the list of the pod's network interfaces used for Aerospike fabric address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' fabric type",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_tls_access_network_names": schema.ListAttribute{
								Description:         "CustomTLSAccessNetworkNames is the list of the pod's network interfaces used for Aerospike TLS access address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' tlsAccess type",
								MarkdownDescription: "CustomTLSAccessNetworkNames is the list of the pod's network interfaces used for Aerospike TLS access address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' tlsAccess type",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_tls_alternate_access_network_names": schema.ListAttribute{
								Description:         "CustomTLSAlternateAccessNetworkNames is the list of the pod's network interfaces used for Aerospike TLSalternate access address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' tlsAlternateAccess type",
								MarkdownDescription: "CustomTLSAlternateAccessNetworkNames is the list of the pod's network interfaces used for Aerospike TLSalternate access address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assignnetwork interfaces to the pod.Required with 'customInterface' tlsAlternateAccess type",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_tls_fabric_network_names": schema.ListAttribute{
								Description:         "CustomTLSFabricNetworkNames is the list of the pod's network interfaces used for Aerospike TLS fabric address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assign networkinterfaces to the pod.Required with 'customInterface' tlsFabric type",
								MarkdownDescription: "CustomTLSFabricNetworkNames is the list of the pod's network interfaces used for Aerospike TLS fabric address.Each element in the list is specified with a namespace and the name of a NetworkAttachmentDefinition,separated by a forward slash (/).These elements must be defined in the pod annotation k8s.v1.cni.cncf.io/networks in order to assign networkinterfaces to the pod.Required with 'customInterface' tlsFabric type",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"fabric": schema.StringAttribute{
								Description:         "FabricType is the type of network address to use for Aerospike fabric address.Defaults is empty meaning all interfaces 'any'.",
								MarkdownDescription: "FabricType is the type of network address to use for Aerospike fabric address.Defaults is empty meaning all interfaces 'any'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("customInterface"),
								},
							},

							"tls_access": schema.StringAttribute{
								Description:         "TLSAccessType is the type of network address to use for Aerospike TLS access address.Defaults to hostInternal.",
								MarkdownDescription: "TLSAccessType is the type of network address to use for Aerospike TLS access address.Defaults to hostInternal.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("pod", "hostInternal", "hostExternal", "configuredIP", "customInterface"),
								},
							},

							"tls_alternate_access": schema.StringAttribute{
								Description:         "TLSAlternateAccessType is the type of network address to use for Aerospike TLS alternate access address.Defaults to hostExternal.",
								MarkdownDescription: "TLSAlternateAccessType is the type of network address to use for Aerospike TLS alternate access address.Defaults to hostExternal.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("pod", "hostInternal", "hostExternal", "configuredIP", "customInterface"),
								},
							},

							"tls_fabric": schema.StringAttribute{
								Description:         "TLSFabricType is the type of network address to use for Aerospike TLS fabric address.Defaults is empty meaning all interfaces 'any'.",
								MarkdownDescription: "TLSFabricType is the type of network address to use for Aerospike TLS fabric address.Defaults is empty meaning all interfaces 'any'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("customInterface"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_pdb": schema.BoolAttribute{
						Description:         "Disable the PodDisruptionBudget creation for the Aerospike cluster.",
						MarkdownDescription: "Disable the PodDisruptionBudget creation for the Aerospike cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_dynamic_config_update": schema.BoolAttribute{
						Description:         "EnableDynamicConfigUpdate enables dynamic config update flow of the operator.If enabled, operator will try to update the Aerospike config dynamically.In case of inconsistent state during dynamic config update, operator falls back to rolling restart.",
						MarkdownDescription: "EnableDynamicConfigUpdate enables dynamic config update flow of the operator.If enabled, operator will try to update the Aerospike config dynamically.In case of inconsistent state during dynamic config update, operator falls back to rolling restart.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "Aerospike server image",
						MarkdownDescription: "Aerospike server image",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"k8s_node_block_list": schema.ListAttribute{
						Description:         "K8sNodeBlockList is a list of Kubernetes nodes which are not used for Aerospike pods. Pods are not scheduled onthese nodes. Pods are migrated from these nodes if already present. This is useful for the maintenance ofKubernetes nodes.",
						MarkdownDescription: "K8sNodeBlockList is a list of Kubernetes nodes which are not used for Aerospike pods. Pods are not scheduled onthese nodes. Pods are migrated from these nodes if already present. This is useful for the maintenance ofKubernetes nodes.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"max_unavailable": schema.StringAttribute{
						Description:         "MaxUnavailable is the percentage/number of pods that can be allowed to go down or unavailable before applicationdisruption. This value is used to create PodDisruptionBudget. Defaults to 1.Refer Aerospike documentation for more details.",
						MarkdownDescription: "MaxUnavailable is the percentage/number of pods that can be allowed to go down or unavailable before applicationdisruption. This value is used to create PodDisruptionBudget. Defaults to 1.Refer Aerospike documentation for more details.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"operations": schema.ListNestedAttribute{
						Description:         "Operations is a list of on-demand operations to be performed on the Aerospike cluster.",
						MarkdownDescription: "Operations is a list of on-demand operations to be performed on the Aerospike cluster.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtLeast(1),
										stringvalidator.LengthAtMost(20),
									},
								},

								"kind": schema.StringAttribute{
									Description:         "Kind is the type of operation to be performed on the Aerospike cluster.",
									MarkdownDescription: "Kind is the type of operation to be performed on the Aerospike cluster.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("WarmRestart", "PodRestart"),
									},
								},

								"pod_list": schema.ListAttribute{
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

					"operator_client_cert": schema.SingleNestedAttribute{
						Description:         "Certificates to connect to Aerospike.",
						MarkdownDescription: "Certificates to connect to Aerospike.",
						Attributes: map[string]schema.Attribute{
							"cert_path_in_operator": schema.SingleNestedAttribute{
								Description:         "AerospikeCertPathInOperatorSource contain configuration for certificates used by operator to connect to aerospikecluster.All paths are on operator's filesystem.",
								MarkdownDescription: "AerospikeCertPathInOperatorSource contain configuration for certificates used by operator to connect to aerospikecluster.All paths are on operator's filesystem.",
								Attributes: map[string]schema.Attribute{
									"ca_certs_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_cert_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_key_path": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_cert_source": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"ca_certs_filename": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ca_certs_source": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"secret_name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"secret_namespace": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_cert_filename": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_key_filename": schema.StringAttribute{
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

									"secret_namespace": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_client_name": schema.StringAttribute{
								Description:         "If specified, this name will be added to tls-authenticate-client list by the operator",
								MarkdownDescription: "If specified, this name will be added to tls-authenticate-client list by the operator",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"paused": schema.BoolAttribute{
						Description:         "Paused flag is used to pause the reconciliation for the AerospikeCluster.",
						MarkdownDescription: "Paused flag is used to pause the reconciliation for the AerospikeCluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_spec": schema.SingleNestedAttribute{
						Description:         "Specify additional configuration for the Aerospike pods",
						MarkdownDescription: "Specify additional configuration for the Aerospike pods",
						Attributes: map[string]schema.Attribute{
							"aerospike_container": schema.SingleNestedAttribute{
								Description:         "AerospikeContainerSpec configures the aerospike-server containercreated by the operator.",
								MarkdownDescription: "AerospikeContainerSpec configures the aerospike-server containercreated by the operator.",
								Attributes: map[string]schema.Attribute{
									"resources": schema.SingleNestedAttribute{
										Description:         "Define resources requests and limits for Aerospike Server Container.Please contact aerospike for proper sizing exerciseOnly Memory and Cpu resources can be givenResources.Limits should be more than Resources.Requests.",
										MarkdownDescription: "Define resources requests and limits for Aerospike Server Container.Please contact aerospike for proper sizing exerciseOnly Memory and Cpu resources can be givenResources.Limits should be more than Resources.Requests.",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
															MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
												Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
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
										Description:         "SecurityContext that will be added to aerospike-server container created by operator.",
										MarkdownDescription: "SecurityContext that will be added to aerospike-server container created by operator.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"capabilities": schema.SingleNestedAttribute{
												Description:         "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
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
												Description:         "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"proc_mount": schema.StringAttribute{
												Description:         "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"read_only_root_filesystem": schema.BoolAttribute{
												Description:         "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"se_linux_options": schema.SingleNestedAttribute{
												Description:         "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
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
												Description:         "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
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
												Description:         "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
												MarkdownDescription: "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
												Attributes: map[string]schema.Attribute{
													"gmsa_credential_spec": schema.StringAttribute{
														Description:         "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
														MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
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
														Description:         "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
														MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_user_name": schema.StringAttribute{
														Description:         "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
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

							"aerospike_init_container": schema.SingleNestedAttribute{
								Description:         "AerospikeInitContainerSpec configures the aerospike-init containercreated by the operator.",
								MarkdownDescription: "AerospikeInitContainerSpec configures the aerospike-init containercreated by the operator.",
								Attributes: map[string]schema.Attribute{
									"image_name_and_tag": schema.StringAttribute{
										Description:         "ImageNameAndTag is the name:tag of aerospike-init container image",
										MarkdownDescription: "ImageNameAndTag is the name:tag of aerospike-init container image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_registry": schema.StringAttribute{
										Description:         "ImageRegistry is the name of image registry for aerospike-init container imageImageRegistry, e.g. docker.io, redhat.access.com",
										MarkdownDescription: "ImageRegistry is the name of image registry for aerospike-init container imageImageRegistry, e.g. docker.io, redhat.access.com",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"image_registry_namespace": schema.StringAttribute{
										Description:         "ImageRegistryNamespace is the name of namespace in registry for aerospike-init container image",
										MarkdownDescription: "ImageRegistryNamespace is the name of namespace in registry for aerospike-init container image",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "Define resources requests and limits for Aerospike init Container.Only Memory and Cpu resources can be givenResources.Limits should be more than Resources.Requests.",
										MarkdownDescription: "Define resources requests and limits for Aerospike init Container.Only Memory and Cpu resources can be givenResources.Limits should be more than Resources.Requests.",
										Attributes: map[string]schema.Attribute{
											"claims": schema.ListNestedAttribute{
												Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
												MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
															MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
												Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
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
										Description:         "SecurityContext that will be added to aerospike-init container created by operator.",
										MarkdownDescription: "SecurityContext that will be added to aerospike-init container created by operator.",
										Attributes: map[string]schema.Attribute{
											"allow_privilege_escalation": schema.BoolAttribute{
												Description:         "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"capabilities": schema.SingleNestedAttribute{
												Description:         "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
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
												Description:         "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"proc_mount": schema.StringAttribute{
												Description:         "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"read_only_root_filesystem": schema.BoolAttribute{
												Description:         "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_group": schema.Int64Attribute{
												Description:         "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_non_root": schema.BoolAttribute{
												Description:         "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user": schema.Int64Attribute{
												Description:         "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"se_linux_options": schema.SingleNestedAttribute{
												Description:         "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
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
												Description:         "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
												MarkdownDescription: "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
												Attributes: map[string]schema.Attribute{
													"localhost_profile": schema.StringAttribute{
														Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
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
												Description:         "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
												MarkdownDescription: "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
												Attributes: map[string]schema.Attribute{
													"gmsa_credential_spec": schema.StringAttribute{
														Description:         "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
														MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
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
														Description:         "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
														MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"run_as_user_name": schema.StringAttribute{
														Description:         "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
														MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
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

							"affinity": schema.SingleNestedAttribute{
								Description:         "Affinity rules for pod placement.",
								MarkdownDescription: "Affinity rules for pod placement.",
								Attributes: map[string]schema.Attribute{
									"node_affinity": schema.SingleNestedAttribute{
										Description:         "Describes node affinity scheduling rules for the pod.",
										MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"preference": schema.SingleNestedAttribute{
															Description:         "A node selector term, associated with the corresponding weight.",
															MarkdownDescription: "A node selector term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's labels.",
																	MarkdownDescription: "A list of node selector requirements by node's labels.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																	Description:         "A list of node selector requirements by node's fields.",
																	MarkdownDescription: "A list of node selector requirements by node's fields.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
															Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
															MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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
												Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
												MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
												Attributes: map[string]schema.Attribute{
													"node_selector_terms": schema.ListNestedAttribute{
														Description:         "Required. A list of node selector terms. The terms are ORed.",
														MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "A list of node selector requirements by node's labels.",
																	MarkdownDescription: "A list of node selector requirements by node's labels.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																	Description:         "A list of node selector requirements by node's fields.",
																	MarkdownDescription: "A list of node selector requirements by node's fields.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
										Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
										MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "Required. A pod affinity term, associated with the corresponding weight.",
															MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																						Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			ElementType:         types.StringType,
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
																	Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"mismatch_label_keys": schema.ListAttribute{
																	Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																						Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			ElementType:         types.StringType,
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
																	Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
															Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
															MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
												Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
												MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
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
															Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
										Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
										MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "Required. A pod affinity term, associated with the corresponding weight.",
															MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																						Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			ElementType:         types.StringType,
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
																	Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"mismatch_label_keys": schema.ListAttribute{
																	Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"namespace_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																						Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			ElementType:         types.StringType,
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
																	Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
															Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
															MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
												Description:         "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
												MarkdownDescription: "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
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
															Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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

							"dns_config": schema.SingleNestedAttribute{
								Description:         "DNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.This is required field when dnsPolicy is set to 'None'",
								MarkdownDescription: "DNSConfig defines the DNS parameters of a pod in addition to those generated from DNSPolicy.This is required field when dnsPolicy is set to 'None'",
								Attributes: map[string]schema.Attribute{
									"nameservers": schema.ListAttribute{
										Description:         "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
										MarkdownDescription: "A list of DNS name server IP addresses.This will be appended to the base nameservers generated from DNSPolicy.Duplicated nameservers will be removed.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"options": schema.ListNestedAttribute{
										Description:         "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
										MarkdownDescription: "A list of DNS resolver options.This will be merged with the base options generated from DNSPolicy.Duplicated entries will be removed. Resolution options given in Optionswill override those that appear in the base DNSPolicy.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Required.",
													MarkdownDescription: "Required.",
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
										Description:         "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
										MarkdownDescription: "A list of DNS search domains for host-name lookup.This will be appended to the base search paths generated from DNSPolicy.Duplicated search paths will be removed.",
										ElementType:         types.StringType,
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
								Description:         "DnsPolicy same as https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy.If hostNetwork is true and policy is not specified, it defaults to ClusterFirstWithHostNet",
								MarkdownDescription: "DnsPolicy same as https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-s-dns-policy.If hostNetwork is true and policy is not specified, it defaults to ClusterFirstWithHostNet",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"effective_dns_policy": schema.StringAttribute{
								Description:         "Effective value of the DNSPolicy",
								MarkdownDescription: "Effective value of the DNSPolicy",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"host_network": schema.BoolAttribute{
								Description:         "HostNetwork enables host networking for the pod.To enable hostNetwork multiPodPerHost must be false.",
								MarkdownDescription: "HostNetwork enables host networking for the pod.To enable hostNetwork multiPodPerHost must be false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_secrets": schema.ListNestedAttribute{
								Description:         "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any ofthe images used by this PodSpec.More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",
								MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the same namespace to use for pulling any ofthe images used by this PodSpec.More info: https://kubernetes.io/docs/concepts/containers/images#specifying-imagepullsecrets-on-a-pod",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
											Required:            false,
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
								Description:         "InitContainers to add to the pods.",
								MarkdownDescription: "InitContainers to add to the pods.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"args": schema.ListAttribute{
											Description:         "Arguments to the entrypoint.The container image's CMD is used if this is not provided.Variable references $(VAR_NAME) are expanded using the container's environment. If a variablecannot be resolved, the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' willproduce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardlessof whether the variable exists or not. Cannot be updated.More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											MarkdownDescription: "Arguments to the entrypoint.The container image's CMD is used if this is not provided.Variable references $(VAR_NAME) are expanded using the container's environment. If a variablecannot be resolved, the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' willproduce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardlessof whether the variable exists or not. Cannot be updated.More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"command": schema.ListAttribute{
											Description:         "Entrypoint array. Not executed within a shell.The container image's ENTRYPOINT is used if this is not provided.Variable references $(VAR_NAME) are expanded using the container's environment. If a variablecannot be resolved, the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' willproduce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardlessof whether the variable exists or not. Cannot be updated.More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											MarkdownDescription: "Entrypoint array. Not executed within a shell.The container image's ENTRYPOINT is used if this is not provided.Variable references $(VAR_NAME) are expanded using the container's environment. If a variablecannot be resolved, the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' willproduce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardlessof whether the variable exists or not. Cannot be updated.More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"env": schema.ListNestedAttribute{
											Description:         "List of environment variables to set in the container.Cannot be updated.",
											MarkdownDescription: "List of environment variables to set in the container.Cannot be updated.",
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
														Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
														Required:            false,
														Optional:            true,
														Computed:            false,
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
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
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																Attributes: map[string]schema.Attribute{
																	"api_version": schema.StringAttribute{
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"field_path": schema.StringAttribute{
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",
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
																Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																Attributes: map[string]schema.Attribute{
																	"container_name": schema.StringAttribute{
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"divisor": schema.StringAttribute{
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"resource": schema.StringAttribute{
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",
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
																Description:         "Selects a key of a secret in the pod's namespace",
																MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
											Description:         "List of sources to populate environment variables in the container.The keys defined within a source must be a C_IDENTIFIER. All invalid keyswill be reported as an event when the container is starting. When a key exists in multiplesources, the value associated with the last source will take precedence.Values defined by an Env with a duplicate key will take precedence.Cannot be updated.",
											MarkdownDescription: "List of sources to populate environment variables in the container.The keys defined within a source must be a C_IDENTIFIER. All invalid keyswill be reported as an event when the container is starting. When a key exists in multiplesources, the value associated with the last source will take precedence.Values defined by an Env with a duplicate key will take precedence.Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "The ConfigMap to select from",
														MarkdownDescription: "The ConfigMap to select from",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"optional": schema.BoolAttribute{
																Description:         "Specify whether the ConfigMap must be defined",
																MarkdownDescription: "Specify whether the ConfigMap must be defined",
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
														Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "The Secret to select from",
														MarkdownDescription: "The Secret to select from",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"optional": schema.BoolAttribute{
																Description:         "Specify whether the Secret must be defined",
																MarkdownDescription: "Specify whether the Secret must be defined",
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
											Description:         "Container image name.More info: https://kubernetes.io/docs/concepts/containers/imagesThis field is optional to allow higher level config management to default or overridecontainer images in workload controllers like Deployments and StatefulSets.",
											MarkdownDescription: "Container image name.More info: https://kubernetes.io/docs/concepts/containers/imagesThis field is optional to allow higher level config management to default or overridecontainer images in workload controllers like Deployments and StatefulSets.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image_pull_policy": schema.StringAttribute{
											Description:         "Image pull policy.One of Always, Never, IfNotPresent.Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
											MarkdownDescription: "Image pull policy.One of Always, Never, IfNotPresent.Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"lifecycle": schema.SingleNestedAttribute{
											Description:         "Actions that the management system should take in response to container lifecycle events.Cannot be updated.",
											MarkdownDescription: "Actions that the management system should take in response to container lifecycle events.Cannot be updated.",
											Attributes: map[string]schema.Attribute{
												"post_start": schema.SingleNestedAttribute{
													Description:         "PostStart is called immediately after a container is created. If the handler fails,the container is terminated and restarted according to its restart policy.Other management of the container blocks until the hook completes.More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails,the container is terminated and restarted according to its restart policy.Other management of the container blocks until the hook completes.More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "Exec specifies the action to take.",
															MarkdownDescription: "Exec specifies the action to take.",
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	ElementType:         types.StringType,
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
															Description:         "HTTPGet specifies the http request to perform.",
															MarkdownDescription: "HTTPGet specifies the http request to perform.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																	MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																	MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",
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
																	Description:         "Path to access on the HTTP server.",
																	MarkdownDescription: "Path to access on the HTTP server.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
																	MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"sleep": schema.SingleNestedAttribute{
															Description:         "Sleep represents the duration that the container should sleep before being terminated.",
															MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
															Attributes: map[string]schema.Attribute{
																"seconds": schema.Int64Attribute{
																	Description:         "Seconds is the number of seconds to sleep.",
																	MarkdownDescription: "Seconds is the number of seconds to sleep.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and keptfor the backward compatibility. There are no validation of this field andlifecycle hooks will fail in runtime when tcp handler is specified.",
															MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and keptfor the backward compatibility. There are no validation of this field andlifecycle hooks will fail in runtime when tcp handler is specified.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																	MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
													Description:         "PreStop is called immediately before a container is terminated due to anAPI request or management event such as liveness/startup probe failure,preemption, resource contention, etc. The handler is not called if thecontainer crashes or exits. The Pod's termination grace period countdown begins before thePreStop hook is executed. Regardless of the outcome of the handler, thecontainer will eventually terminate within the Pod's termination graceperiod (unless delayed by finalizers). Other management of the container blocks until the hook completesor until the termination grace period is reached.More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													MarkdownDescription: "PreStop is called immediately before a container is terminated due to anAPI request or management event such as liveness/startup probe failure,preemption, resource contention, etc. The handler is not called if thecontainer crashes or exits. The Pod's termination grace period countdown begins before thePreStop hook is executed. Regardless of the outcome of the handler, thecontainer will eventually terminate within the Pod's termination graceperiod (unless delayed by finalizers). Other management of the container blocks until the hook completesor until the termination grace period is reached.More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "Exec specifies the action to take.",
															MarkdownDescription: "Exec specifies the action to take.",
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	ElementType:         types.StringType,
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
															Description:         "HTTPGet specifies the http request to perform.",
															MarkdownDescription: "HTTPGet specifies the http request to perform.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																	MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																	MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",
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
																	Description:         "Path to access on the HTTP server.",
																	MarkdownDescription: "Path to access on the HTTP server.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
																	MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"sleep": schema.SingleNestedAttribute{
															Description:         "Sleep represents the duration that the container should sleep before being terminated.",
															MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
															Attributes: map[string]schema.Attribute{
																"seconds": schema.Int64Attribute{
																	Description:         "Seconds is the number of seconds to sleep.",
																	MarkdownDescription: "Seconds is the number of seconds to sleep.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and keptfor the backward compatibility. There are no validation of this field andlifecycle hooks will fail in runtime when tcp handler is specified.",
															MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and keptfor the backward compatibility. There are no validation of this field andlifecycle hooks will fail in runtime when tcp handler is specified.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																	MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
											Description:         "Periodic probe of container liveness.Container will be restarted if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											MarkdownDescription: "Periodic probe of container liveness.Container will be restarted if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															ElementType:         types.StringType,
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
													Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grpc": schema.SingleNestedAttribute{
													Description:         "GRPC specifies an action involving a GRPC port.",
													MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
															Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
															MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
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
													Description:         "HTTPGet specifies the http request to perform.",
													MarkdownDescription: "HTTPGet specifies the http request to perform.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
															MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",
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
															Description:         "Path to access on the HTTP server.",
															MarkdownDescription: "Path to access on the HTTP server.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"scheme": schema.StringAttribute{
															Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
															MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
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
													Description:         "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"period_seconds": schema.Int64Attribute{
													Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"success_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "TCPSocket specifies an action involving a TCP port.",
													MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Optional: Host name to connect to, defaults to the pod IP.",
															MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
													Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
											Description:         "Name of the container specified as a DNS_LABEL.Each container in a pod must have a unique name (DNS_LABEL).Cannot be updated.",
											MarkdownDescription: "Name of the container specified as a DNS_LABEL.Each container in a pod must have a unique name (DNS_LABEL).Cannot be updated.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"ports": schema.ListNestedAttribute{
											Description:         "List of ports to expose from the container. Not specifying a port hereDOES NOT prevent that port from being exposed. Any port which islistening on the default '0.0.0.0' address inside a container will beaccessible from the network.Modifying this array with strategic merge patch may corrupt the data.For more information See https://github.com/kubernetes/kubernetes/issues/108255.Cannot be updated.",
											MarkdownDescription: "List of ports to expose from the container. Not specifying a port hereDOES NOT prevent that port from being exposed. Any port which islistening on the default '0.0.0.0' address inside a container will beaccessible from the network.Modifying this array with strategic merge patch may corrupt the data.For more information See https://github.com/kubernetes/kubernetes/issues/108255.Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_port": schema.Int64Attribute{
														Description:         "Number of port to expose on the pod's IP address.This must be a valid port number, 0 < x < 65536.",
														MarkdownDescription: "Number of port to expose on the pod's IP address.This must be a valid port number, 0 < x < 65536.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"host_ip": schema.StringAttribute{
														Description:         "What host IP to bind the external port to.",
														MarkdownDescription: "What host IP to bind the external port to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host_port": schema.Int64Attribute{
														Description:         "Number of port to expose on the host.If specified, this must be a valid port number, 0 < x < 65536.If HostNetwork is specified, this must match ContainerPort.Most containers do not need this.",
														MarkdownDescription: "Number of port to expose on the host.If specified, this must be a valid port number, 0 < x < 65536.If HostNetwork is specified, this must match ContainerPort.Most containers do not need this.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Eachnamed port in a pod must have a unique name. Name for the port that can bereferred to by services.",
														MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Eachnamed port in a pod must have a unique name. Name for the port that can bereferred to by services.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol for port. Must be UDP, TCP, or SCTP.Defaults to 'TCP'.",
														MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP.Defaults to 'TCP'.",
														Required:            false,
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
											Description:         "Periodic probe of container service readiness.Container will be removed from service endpoints if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											MarkdownDescription: "Periodic probe of container service readiness.Container will be removed from service endpoints if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															ElementType:         types.StringType,
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
													Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grpc": schema.SingleNestedAttribute{
													Description:         "GRPC specifies an action involving a GRPC port.",
													MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
															Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
															MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
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
													Description:         "HTTPGet specifies the http request to perform.",
													MarkdownDescription: "HTTPGet specifies the http request to perform.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
															MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",
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
															Description:         "Path to access on the HTTP server.",
															MarkdownDescription: "Path to access on the HTTP server.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"scheme": schema.StringAttribute{
															Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
															MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
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
													Description:         "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"period_seconds": schema.Int64Attribute{
													Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"success_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "TCPSocket specifies an action involving a TCP port.",
													MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Optional: Host name to connect to, defaults to the pod IP.",
															MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
													Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"resize_policy": schema.ListNestedAttribute{
											Description:         "Resources resize policy for the container.",
											MarkdownDescription: "Resources resize policy for the container.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"resource_name": schema.StringAttribute{
														Description:         "Name of the resource to which this resource resize policy applies.Supported values: cpu, memory.",
														MarkdownDescription: "Name of the resource to which this resource resize policy applies.Supported values: cpu, memory.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"restart_policy": schema.StringAttribute{
														Description:         "Restart policy to apply when specified resource is resized.If not specified, it defaults to NotRequired.",
														MarkdownDescription: "Restart policy to apply when specified resource is resized.If not specified, it defaults to NotRequired.",
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

										"resources": schema.SingleNestedAttribute{
											Description:         "Compute Resources required by this container.Cannot be updated.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Compute Resources required by this container.Cannot be updated.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											Attributes: map[string]schema.Attribute{
												"claims": schema.ListNestedAttribute{
													Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
													MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
													Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"requests": schema.MapAttribute{
													Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"restart_policy": schema.StringAttribute{
											Description:         "RestartPolicy defines the restart behavior of individual containers in a pod.This field may only be set for init containers, and the only allowed value is 'Always'.For non-init containers or when this field is not specified,the restart behavior is defined by the Pod's restart policy and the container type.Setting the RestartPolicy as 'Always' for the init container will have the following effect:this init container will be continually restarted onexit until all regular containers have terminated. Once all regularcontainers have completed, all init containers with restartPolicy 'Always'will be shut down. This lifecycle differs from normal init containers andis often referred to as a 'sidecar' container. Although this initcontainer still starts in the init container sequence, it does not waitfor the container to complete before proceeding to the next initcontainer. Instead, the next init container starts immediately after thisinit container is started, or after any startupProbe has successfullycompleted.",
											MarkdownDescription: "RestartPolicy defines the restart behavior of individual containers in a pod.This field may only be set for init containers, and the only allowed value is 'Always'.For non-init containers or when this field is not specified,the restart behavior is defined by the Pod's restart policy and the container type.Setting the RestartPolicy as 'Always' for the init container will have the following effect:this init container will be continually restarted onexit until all regular containers have terminated. Once all regularcontainers have completed, all init containers with restartPolicy 'Always'will be shut down. This lifecycle differs from normal init containers andis often referred to as a 'sidecar' container. Although this initcontainer still starts in the init container sequence, it does not waitfor the container to complete before proceeding to the next initcontainer. Instead, the next init container starts immediately after thisinit container is started, or after any startupProbe has successfullycompleted.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"security_context": schema.SingleNestedAttribute{
											Description:         "SecurityContext defines the security options the container should be run with.If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
											MarkdownDescription: "SecurityContext defines the security options the container should be run with.If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
											Attributes: map[string]schema.Attribute{
												"allow_privilege_escalation": schema.BoolAttribute{
													Description:         "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"capabilities": schema.SingleNestedAttribute{
													Description:         "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
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
													Description:         "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proc_mount": schema.StringAttribute{
													Description:         "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_only_root_filesystem": schema.BoolAttribute{
													Description:         "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"run_as_group": schema.Int64Attribute{
													Description:         "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"run_as_non_root": schema.BoolAttribute{
													Description:         "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
													MarkdownDescription: "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"run_as_user": schema.Int64Attribute{
													Description:         "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"se_linux_options": schema.SingleNestedAttribute{
													Description:         "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
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
													Description:         "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
													Attributes: map[string]schema.Attribute{
														"localhost_profile": schema.StringAttribute{
															Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
															MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
															MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
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
													Description:         "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
													MarkdownDescription: "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
													Attributes: map[string]schema.Attribute{
														"gmsa_credential_spec": schema.StringAttribute{
															Description:         "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
															MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
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
															Description:         "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
															MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"run_as_user_name": schema.StringAttribute{
															Description:         "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
															MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
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
											Description:         "StartupProbe indicates that the Pod has successfully initialized.If specified, no other probes are executed until this completes successfully.If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,when it might take a long time to load data or warm a cache, than during steady-state operation.This cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized.If specified, no other probes are executed until this completes successfully.If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,when it might take a long time to load data or warm a cache, than during steady-state operation.This cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															ElementType:         types.StringType,
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
													Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grpc": schema.SingleNestedAttribute{
													Description:         "GRPC specifies an action involving a GRPC port.",
													MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
															Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
															MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
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
													Description:         "HTTPGet specifies the http request to perform.",
													MarkdownDescription: "HTTPGet specifies the http request to perform.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
															MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",
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
															Description:         "Path to access on the HTTP server.",
															MarkdownDescription: "Path to access on the HTTP server.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"scheme": schema.StringAttribute{
															Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
															MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
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
													Description:         "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"period_seconds": schema.Int64Attribute{
													Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"success_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "TCPSocket specifies an action involving a TCP port.",
													MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Optional: Host name to connect to, defaults to the pod IP.",
															MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
													Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
											Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If thisis not set, reads from stdin in the container will always result in EOF.Default is false.",
											MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If thisis not set, reads from stdin in the container will always result in EOF.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"stdin_once": schema.BoolAttribute{
											Description:         "Whether the container runtime should close the stdin channel after it has been opened bya single attach. When stdin is true the stdin stream will remain open across multiple attachsessions. If stdinOnce is set to true, stdin is opened on container start, is empty until thefirst client attaches to stdin, and then remains open and accepts data until the client disconnects,at which time stdin is closed and remains closed until the container is restarted. If thisflag is false, a container processes that reads from stdin will never receive an EOF.Default is false",
											MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened bya single attach. When stdin is true the stdin stream will remain open across multiple attachsessions. If stdinOnce is set to true, stdin is opened on container start, is empty until thefirst client attaches to stdin, and then remains open and accepts data until the client disconnects,at which time stdin is closed and remains closed until the container is restarted. If thisflag is false, a container processes that reads from stdin will never receive an EOF.Default is false",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"termination_message_path": schema.StringAttribute{
											Description:         "Optional: Path at which the file to which the container's termination messagewill be written is mounted into the container's filesystem.Message written is intended to be brief final status, such as an assertion failure message.Will be truncated by the node if greater than 4096 bytes. The total message length acrossall containers will be limited to 12kb.Defaults to /dev/termination-log.Cannot be updated.",
											MarkdownDescription: "Optional: Path at which the file to which the container's termination messagewill be written is mounted into the container's filesystem.Message written is intended to be brief final status, such as an assertion failure message.Will be truncated by the node if greater than 4096 bytes. The total message length acrossall containers will be limited to 12kb.Defaults to /dev/termination-log.Cannot be updated.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"termination_message_policy": schema.StringAttribute{
											Description:         "Indicate how the termination message should be populated. File will use the contents ofterminationMessagePath to populate the container status message on both success and failure.FallbackToLogsOnError will use the last chunk of container log output if the terminationmessage file is empty and the container exited with an error.The log output is limited to 2048 bytes or 80 lines, whichever is smaller.Defaults to File.Cannot be updated.",
											MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents ofterminationMessagePath to populate the container status message on both success and failure.FallbackToLogsOnError will use the last chunk of container log output if the terminationmessage file is empty and the container exited with an error.The log output is limited to 2048 bytes or 80 lines, whichever is smaller.Defaults to File.Cannot be updated.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tty": schema.BoolAttribute{
											Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.Default is false.",
											MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_devices": schema.ListNestedAttribute{
											Description:         "volumeDevices is the list of block devices to be used by the container.",
											MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"device_path": schema.StringAttribute{
														Description:         "devicePath is the path inside of the container that the device will be mapped to.",
														MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "name must match the name of a persistentVolumeClaim in the pod",
														MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",
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
											Description:         "Pod volumes to mount into the container's filesystem.Cannot be updated.",
											MarkdownDescription: "Pod volumes to mount into the container's filesystem.Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"mount_path": schema.StringAttribute{
														Description:         "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
														MarkdownDescription: "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"mount_propagation": schema.StringAttribute{
														Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
														Required:            false,
														Optional:            true,
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
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sub_path": schema.StringAttribute{
														Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
														MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sub_path_expr": schema.StringAttribute{
														Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
														MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
														Required:            false,
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
											Description:         "Container's working directory.If not specified, the container runtime's default will be used, whichmight be configured in the container image.Cannot be updated.",
											MarkdownDescription: "Container's working directory.If not specified, the container runtime's default will be used, whichmight be configured in the container image.Cannot be updated.",
											Required:            false,
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
								Description:         "MetaData to add to the pod.",
								MarkdownDescription: "MetaData to add to the pod.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Key - Value pair that may be set by external tools to store and retrieve arbitrary metadata",
										MarkdownDescription: "Key - Value pair that may be set by external tools to store and retrieve arbitrary metadata",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Key - Value pairs that can be used to organize and categorize scope and select objects",
										MarkdownDescription: "Key - Value pairs that can be used to organize and categorize scope and select objects",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"multi_pod_per_host": schema.BoolAttribute{
								Description:         "If set true then multiple pods can be created per Kubernetes Node.This will create a NodePort service for each Pod if aerospikeNetworkPolicy definedhas one of the network types: 'hostInternal', 'hostExternal', 'configuredIP'NodePort, as the name implies, opens a specific port on all the Kubernetes Nodes ,and any traffic that is sent to this port is forwarded to the service.Here service picks a random port in range (30000-32767), so these port should be open.If set false then only single pod can be created per Kubernetes Node.This will create Pods using hostPort setting.The container port will be exposed to the external network at <hostIP>:<hostPort>,where the hostIP is the IP address of the Kubernetes Node where the container is running andthe hostPort is the port requested by the user.",
								MarkdownDescription: "If set true then multiple pods can be created per Kubernetes Node.This will create a NodePort service for each Pod if aerospikeNetworkPolicy definedhas one of the network types: 'hostInternal', 'hostExternal', 'configuredIP'NodePort, as the name implies, opens a specific port on all the Kubernetes Nodes ,and any traffic that is sent to this port is forwarded to the service.Here service picks a random port in range (30000-32767), so these port should be open.If set false then only single pod can be created per Kubernetes Node.This will create Pods using hostPort setting.The container port will be exposed to the external network at <hostIP>:<hostPort>,where the hostIP is the IP address of the Kubernetes Node where the container is running andthe hostPort is the port requested by the user.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector constraints for this pod.",
								MarkdownDescription: "NodeSelector constraints for this pod.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context": schema.SingleNestedAttribute{
								Description:         "SecurityContext holds pod-level security attributes and common container settings.Optional: Defaults to empty.  See type description for default values of each field.",
								MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings.Optional: Defaults to empty.  See type description for default values of each field.",
								Attributes: map[string]schema.Attribute{
									"fs_group": schema.Int64Attribute{
										Description:         "A special supplemental group that applies to all containers in a pod.Some volume types allow the Kubelet to change the ownership of that volumeto be owned by the pod:1. The owning GID will be the FSGroup2. The setgid bit is set (new files created in the volume will be owned by FSGroup)3. The permission bits are OR'd with rw-rw----If unset, the Kubelet will not modify the ownership and permissions of any volume.Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A special supplemental group that applies to all containers in a pod.Some volume types allow the Kubelet to change the ownership of that volumeto be owned by the pod:1. The owning GID will be the FSGroup2. The setgid bit is set (new files created in the volume will be owned by FSGroup)3. The permission bits are OR'd with rw-rw----If unset, the Kubelet will not modify the ownership and permissions of any volume.Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"fs_group_change_policy": schema.StringAttribute{
										Description:         "fsGroupChangePolicy defines behavior of changing ownership and permission of the volumebefore being exposed inside Pod. This field will only apply tovolume types which support fsGroup based ownership(and permissions).It will have no effect on ephemeral volume types such as: secret, configmapsand emptydir.Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used.Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permission of the volumebefore being exposed inside Pod. This field will only apply tovolume types which support fsGroup based ownership(and permissions).It will have no effect on ephemeral volume types such as: secret, configmapsand emptydir.Valid values are 'OnRootMismatch' and 'Always'. If not specified, 'Always' is used.Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_group": schema.Int64Attribute{
										Description:         "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_non_root": schema.BoolAttribute{
										Description:         "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
										MarkdownDescription: "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"run_as_user": schema.Int64Attribute{
										Description:         "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in SecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedencefor that container.Note that this field cannot be set when spec.os.name is windows.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"se_linux_options": schema.SingleNestedAttribute{
										Description:         "The SELinux context to be applied to all containers.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in SecurityContext.  If set inboth SecurityContext and PodSecurityContext, the value specified in SecurityContexttakes precedence for that container.Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The SELinux context to be applied to all containers.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in SecurityContext.  If set inboth SecurityContext and PodSecurityContext, the value specified in SecurityContexttakes precedence for that container.Note that this field cannot be set when spec.os.name is windows.",
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
										Description:         "The seccomp options to use by the containers in this pod.Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "The seccomp options to use by the containers in this pod.Note that this field cannot be set when spec.os.name is windows.",
										Attributes: map[string]schema.Attribute{
											"localhost_profile": schema.StringAttribute{
												Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
												MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
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
										Description:         "A list of groups applied to the first process run in each container, in additionto the container's primary GID, the fsGroup (if specified), and group membershipsdefined in the container image for the uid of the container process. If unspecified,no additional groups are added to any container. Note that group membershipsdefined in the container image for the uid of the container process are still effective,even if they are not included in this list.Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "A list of groups applied to the first process run in each container, in additionto the container's primary GID, the fsGroup (if specified), and group membershipsdefined in the container image for the uid of the container process. If unspecified,no additional groups are added to any container. Note that group membershipsdefined in the container image for the uid of the container process are still effective,even if they are not included in this list.Note that this field cannot be set when spec.os.name is windows.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"sysctls": schema.ListNestedAttribute{
										Description:         "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupportedsysctls (by the container runtime) might fail to launch.Note that this field cannot be set when spec.os.name is windows.",
										MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupportedsysctls (by the container runtime) might fail to launch.Note that this field cannot be set when spec.os.name is windows.",
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
										Description:         "The Windows specific settings applied to all containers.If unspecified, the options within a container's SecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
										MarkdownDescription: "The Windows specific settings applied to all containers.If unspecified, the options within a container's SecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
										Attributes: map[string]schema.Attribute{
											"gmsa_credential_spec": schema.StringAttribute{
												Description:         "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
												MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
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
												Description:         "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
												MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"run_as_user_name": schema.StringAttribute{
												Description:         "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
												MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
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

							"sidecars": schema.ListNestedAttribute{
								Description:         "Sidecars to add to the pod.",
								MarkdownDescription: "Sidecars to add to the pod.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"args": schema.ListAttribute{
											Description:         "Arguments to the entrypoint.The container image's CMD is used if this is not provided.Variable references $(VAR_NAME) are expanded using the container's environment. If a variablecannot be resolved, the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' willproduce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardlessof whether the variable exists or not. Cannot be updated.More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											MarkdownDescription: "Arguments to the entrypoint.The container image's CMD is used if this is not provided.Variable references $(VAR_NAME) are expanded using the container's environment. If a variablecannot be resolved, the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' willproduce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardlessof whether the variable exists or not. Cannot be updated.More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"command": schema.ListAttribute{
											Description:         "Entrypoint array. Not executed within a shell.The container image's ENTRYPOINT is used if this is not provided.Variable references $(VAR_NAME) are expanded using the container's environment. If a variablecannot be resolved, the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' willproduce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardlessof whether the variable exists or not. Cannot be updated.More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											MarkdownDescription: "Entrypoint array. Not executed within a shell.The container image's ENTRYPOINT is used if this is not provided.Variable references $(VAR_NAME) are expanded using the container's environment. If a variablecannot be resolved, the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' willproduce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardlessof whether the variable exists or not. Cannot be updated.More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"env": schema.ListNestedAttribute{
											Description:         "List of environment variables to set in the container.Cannot be updated.",
											MarkdownDescription: "List of environment variables to set in the container.Cannot be updated.",
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
														Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
														Required:            false,
														Optional:            true,
														Computed:            false,
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
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
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
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																Attributes: map[string]schema.Attribute{
																	"api_version": schema.StringAttribute{
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"field_path": schema.StringAttribute{
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",
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
																Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																Attributes: map[string]schema.Attribute{
																	"container_name": schema.StringAttribute{
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"divisor": schema.StringAttribute{
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"resource": schema.StringAttribute{
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",
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
																Description:         "Selects a key of a secret in the pod's namespace",
																MarkdownDescription: "Selects a key of a secret in the pod's namespace",
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"optional": schema.BoolAttribute{
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
											Description:         "List of sources to populate environment variables in the container.The keys defined within a source must be a C_IDENTIFIER. All invalid keyswill be reported as an event when the container is starting. When a key exists in multiplesources, the value associated with the last source will take precedence.Values defined by an Env with a duplicate key will take precedence.Cannot be updated.",
											MarkdownDescription: "List of sources to populate environment variables in the container.The keys defined within a source must be a C_IDENTIFIER. All invalid keyswill be reported as an event when the container is starting. When a key exists in multiplesources, the value associated with the last source will take precedence.Values defined by an Env with a duplicate key will take precedence.Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "The ConfigMap to select from",
														MarkdownDescription: "The ConfigMap to select from",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"optional": schema.BoolAttribute{
																Description:         "Specify whether the ConfigMap must be defined",
																MarkdownDescription: "Specify whether the ConfigMap must be defined",
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
														Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"secret_ref": schema.SingleNestedAttribute{
														Description:         "The Secret to select from",
														MarkdownDescription: "The Secret to select from",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"optional": schema.BoolAttribute{
																Description:         "Specify whether the Secret must be defined",
																MarkdownDescription: "Specify whether the Secret must be defined",
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
											Description:         "Container image name.More info: https://kubernetes.io/docs/concepts/containers/imagesThis field is optional to allow higher level config management to default or overridecontainer images in workload controllers like Deployments and StatefulSets.",
											MarkdownDescription: "Container image name.More info: https://kubernetes.io/docs/concepts/containers/imagesThis field is optional to allow higher level config management to default or overridecontainer images in workload controllers like Deployments and StatefulSets.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image_pull_policy": schema.StringAttribute{
											Description:         "Image pull policy.One of Always, Never, IfNotPresent.Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
											MarkdownDescription: "Image pull policy.One of Always, Never, IfNotPresent.Defaults to Always if :latest tag is specified, or IfNotPresent otherwise.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"lifecycle": schema.SingleNestedAttribute{
											Description:         "Actions that the management system should take in response to container lifecycle events.Cannot be updated.",
											MarkdownDescription: "Actions that the management system should take in response to container lifecycle events.Cannot be updated.",
											Attributes: map[string]schema.Attribute{
												"post_start": schema.SingleNestedAttribute{
													Description:         "PostStart is called immediately after a container is created. If the handler fails,the container is terminated and restarted according to its restart policy.Other management of the container blocks until the hook completes.More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails,the container is terminated and restarted according to its restart policy.Other management of the container blocks until the hook completes.More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "Exec specifies the action to take.",
															MarkdownDescription: "Exec specifies the action to take.",
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	ElementType:         types.StringType,
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
															Description:         "HTTPGet specifies the http request to perform.",
															MarkdownDescription: "HTTPGet specifies the http request to perform.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																	MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																	MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",
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
																	Description:         "Path to access on the HTTP server.",
																	MarkdownDescription: "Path to access on the HTTP server.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
																	MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"sleep": schema.SingleNestedAttribute{
															Description:         "Sleep represents the duration that the container should sleep before being terminated.",
															MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
															Attributes: map[string]schema.Attribute{
																"seconds": schema.Int64Attribute{
																	Description:         "Seconds is the number of seconds to sleep.",
																	MarkdownDescription: "Seconds is the number of seconds to sleep.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and keptfor the backward compatibility. There are no validation of this field andlifecycle hooks will fail in runtime when tcp handler is specified.",
															MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and keptfor the backward compatibility. There are no validation of this field andlifecycle hooks will fail in runtime when tcp handler is specified.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																	MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
													Description:         "PreStop is called immediately before a container is terminated due to anAPI request or management event such as liveness/startup probe failure,preemption, resource contention, etc. The handler is not called if thecontainer crashes or exits. The Pod's termination grace period countdown begins before thePreStop hook is executed. Regardless of the outcome of the handler, thecontainer will eventually terminate within the Pod's termination graceperiod (unless delayed by finalizers). Other management of the container blocks until the hook completesor until the termination grace period is reached.More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													MarkdownDescription: "PreStop is called immediately before a container is terminated due to anAPI request or management event such as liveness/startup probe failure,preemption, resource contention, etc. The handler is not called if thecontainer crashes or exits. The Pod's termination grace period countdown begins before thePreStop hook is executed. Regardless of the outcome of the handler, thecontainer will eventually terminate within the Pod's termination graceperiod (unless delayed by finalizers). Other management of the container blocks until the hook completesor until the termination grace period is reached.More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
													Attributes: map[string]schema.Attribute{
														"exec": schema.SingleNestedAttribute{
															Description:         "Exec specifies the action to take.",
															MarkdownDescription: "Exec specifies the action to take.",
															Attributes: map[string]schema.Attribute{
																"command": schema.ListAttribute{
																	Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
																	ElementType:         types.StringType,
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
															Description:         "HTTPGet specifies the http request to perform.",
															MarkdownDescription: "HTTPGet specifies the http request to perform.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																	MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"http_headers": schema.ListNestedAttribute{
																	Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
																	MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"name": schema.StringAttribute{
																				Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "The header field value",
																				MarkdownDescription: "The header field value",
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
																	Description:         "Path to access on the HTTP server.",
																	MarkdownDescription: "Path to access on the HTTP server.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"scheme": schema.StringAttribute{
																	Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
																	MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"sleep": schema.SingleNestedAttribute{
															Description:         "Sleep represents the duration that the container should sleep before being terminated.",
															MarkdownDescription: "Sleep represents the duration that the container should sleep before being terminated.",
															Attributes: map[string]schema.Attribute{
																"seconds": schema.Int64Attribute{
																	Description:         "Seconds is the number of seconds to sleep.",
																	MarkdownDescription: "Seconds is the number of seconds to sleep.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"tcp_socket": schema.SingleNestedAttribute{
															Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and keptfor the backward compatibility. There are no validation of this field andlifecycle hooks will fail in runtime when tcp handler is specified.",
															MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and keptfor the backward compatibility. There are no validation of this field andlifecycle hooks will fail in runtime when tcp handler is specified.",
															Attributes: map[string]schema.Attribute{
																"host": schema.StringAttribute{
																	Description:         "Optional: Host name to connect to, defaults to the pod IP.",
																	MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"port": schema.StringAttribute{
																	Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
																	MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
											Description:         "Periodic probe of container liveness.Container will be restarted if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											MarkdownDescription: "Periodic probe of container liveness.Container will be restarted if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															ElementType:         types.StringType,
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
													Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grpc": schema.SingleNestedAttribute{
													Description:         "GRPC specifies an action involving a GRPC port.",
													MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
															Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
															MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
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
													Description:         "HTTPGet specifies the http request to perform.",
													MarkdownDescription: "HTTPGet specifies the http request to perform.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
															MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",
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
															Description:         "Path to access on the HTTP server.",
															MarkdownDescription: "Path to access on the HTTP server.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"scheme": schema.StringAttribute{
															Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
															MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
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
													Description:         "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"period_seconds": schema.Int64Attribute{
													Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"success_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "TCPSocket specifies an action involving a TCP port.",
													MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Optional: Host name to connect to, defaults to the pod IP.",
															MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
													Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
											Description:         "Name of the container specified as a DNS_LABEL.Each container in a pod must have a unique name (DNS_LABEL).Cannot be updated.",
											MarkdownDescription: "Name of the container specified as a DNS_LABEL.Each container in a pod must have a unique name (DNS_LABEL).Cannot be updated.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"ports": schema.ListNestedAttribute{
											Description:         "List of ports to expose from the container. Not specifying a port hereDOES NOT prevent that port from being exposed. Any port which islistening on the default '0.0.0.0' address inside a container will beaccessible from the network.Modifying this array with strategic merge patch may corrupt the data.For more information See https://github.com/kubernetes/kubernetes/issues/108255.Cannot be updated.",
											MarkdownDescription: "List of ports to expose from the container. Not specifying a port hereDOES NOT prevent that port from being exposed. Any port which islistening on the default '0.0.0.0' address inside a container will beaccessible from the network.Modifying this array with strategic merge patch may corrupt the data.For more information See https://github.com/kubernetes/kubernetes/issues/108255.Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_port": schema.Int64Attribute{
														Description:         "Number of port to expose on the pod's IP address.This must be a valid port number, 0 < x < 65536.",
														MarkdownDescription: "Number of port to expose on the pod's IP address.This must be a valid port number, 0 < x < 65536.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"host_ip": schema.StringAttribute{
														Description:         "What host IP to bind the external port to.",
														MarkdownDescription: "What host IP to bind the external port to.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"host_port": schema.Int64Attribute{
														Description:         "Number of port to expose on the host.If specified, this must be a valid port number, 0 < x < 65536.If HostNetwork is specified, this must match ContainerPort.Most containers do not need this.",
														MarkdownDescription: "Number of port to expose on the host.If specified, this must be a valid port number, 0 < x < 65536.If HostNetwork is specified, this must match ContainerPort.Most containers do not need this.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Eachnamed port in a pod must have a unique name. Name for the port that can bereferred to by services.",
														MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Eachnamed port in a pod must have a unique name. Name for the port that can bereferred to by services.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol for port. Must be UDP, TCP, or SCTP.Defaults to 'TCP'.",
														MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP.Defaults to 'TCP'.",
														Required:            false,
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
											Description:         "Periodic probe of container service readiness.Container will be removed from service endpoints if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											MarkdownDescription: "Periodic probe of container service readiness.Container will be removed from service endpoints if the probe fails.Cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															ElementType:         types.StringType,
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
													Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grpc": schema.SingleNestedAttribute{
													Description:         "GRPC specifies an action involving a GRPC port.",
													MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
															Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
															MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
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
													Description:         "HTTPGet specifies the http request to perform.",
													MarkdownDescription: "HTTPGet specifies the http request to perform.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
															MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",
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
															Description:         "Path to access on the HTTP server.",
															MarkdownDescription: "Path to access on the HTTP server.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"scheme": schema.StringAttribute{
															Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
															MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
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
													Description:         "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"period_seconds": schema.Int64Attribute{
													Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"success_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "TCPSocket specifies an action involving a TCP port.",
													MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Optional: Host name to connect to, defaults to the pod IP.",
															MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
													Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"resize_policy": schema.ListNestedAttribute{
											Description:         "Resources resize policy for the container.",
											MarkdownDescription: "Resources resize policy for the container.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"resource_name": schema.StringAttribute{
														Description:         "Name of the resource to which this resource resize policy applies.Supported values: cpu, memory.",
														MarkdownDescription: "Name of the resource to which this resource resize policy applies.Supported values: cpu, memory.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"restart_policy": schema.StringAttribute{
														Description:         "Restart policy to apply when specified resource is resized.If not specified, it defaults to NotRequired.",
														MarkdownDescription: "Restart policy to apply when specified resource is resized.If not specified, it defaults to NotRequired.",
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

										"resources": schema.SingleNestedAttribute{
											Description:         "Compute Resources required by this container.Cannot be updated.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Compute Resources required by this container.Cannot be updated.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											Attributes: map[string]schema.Attribute{
												"claims": schema.ListNestedAttribute{
													Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
													MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
																MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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
													Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"requests": schema.MapAttribute{
													Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"restart_policy": schema.StringAttribute{
											Description:         "RestartPolicy defines the restart behavior of individual containers in a pod.This field may only be set for init containers, and the only allowed value is 'Always'.For non-init containers or when this field is not specified,the restart behavior is defined by the Pod's restart policy and the container type.Setting the RestartPolicy as 'Always' for the init container will have the following effect:this init container will be continually restarted onexit until all regular containers have terminated. Once all regularcontainers have completed, all init containers with restartPolicy 'Always'will be shut down. This lifecycle differs from normal init containers andis often referred to as a 'sidecar' container. Although this initcontainer still starts in the init container sequence, it does not waitfor the container to complete before proceeding to the next initcontainer. Instead, the next init container starts immediately after thisinit container is started, or after any startupProbe has successfullycompleted.",
											MarkdownDescription: "RestartPolicy defines the restart behavior of individual containers in a pod.This field may only be set for init containers, and the only allowed value is 'Always'.For non-init containers or when this field is not specified,the restart behavior is defined by the Pod's restart policy and the container type.Setting the RestartPolicy as 'Always' for the init container will have the following effect:this init container will be continually restarted onexit until all regular containers have terminated. Once all regularcontainers have completed, all init containers with restartPolicy 'Always'will be shut down. This lifecycle differs from normal init containers andis often referred to as a 'sidecar' container. Although this initcontainer still starts in the init container sequence, it does not waitfor the container to complete before proceeding to the next initcontainer. Instead, the next init container starts immediately after thisinit container is started, or after any startupProbe has successfullycompleted.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"security_context": schema.SingleNestedAttribute{
											Description:         "SecurityContext defines the security options the container should be run with.If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
											MarkdownDescription: "SecurityContext defines the security options the container should be run with.If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
											Attributes: map[string]schema.Attribute{
												"allow_privilege_escalation": schema.BoolAttribute{
													Description:         "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain moreprivileges than its parent process. This bool directly controls ifthe no_new_privs flag will be set on the container process.AllowPrivilegeEscalation is true always when the container is:1) run as Privileged2) has CAP_SYS_ADMINNote that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"capabilities": schema.SingleNestedAttribute{
													Description:         "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The capabilities to add/drop when running containers.Defaults to the default set of capabilities granted by the container runtime.Note that this field cannot be set when spec.os.name is windows.",
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
													Description:         "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "Run container in privileged mode.Processes in privileged containers are essentially equivalent to root on the host.Defaults to false.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"proc_mount": schema.StringAttribute{
													Description:         "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.The default is DefaultProcMount which uses the container runtime defaults forreadonly paths and masked paths.This requires the ProcMountType feature flag to be enabled.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_only_root_filesystem": schema.BoolAttribute{
													Description:         "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "Whether this container has a read-only root filesystem.Default is false.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"run_as_group": schema.Int64Attribute{
													Description:         "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The GID to run the entrypoint of the container process.Uses runtime default if unset.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"run_as_non_root": schema.BoolAttribute{
													Description:         "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
													MarkdownDescription: "Indicates that the container must run as a non-root user.If true, the Kubelet will validate the image at runtime to ensure that itdoes not run as UID 0 (root) and fail to start the container if it does.If unset or false, no such validation will be performed.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"run_as_user": schema.Int64Attribute{
													Description:         "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The UID to run the entrypoint of the container process.Defaults to user specified in image metadata if unspecified.May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"se_linux_options": schema.SingleNestedAttribute{
													Description:         "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The SELinux context to be applied to the container.If unspecified, the container runtime will allocate a random SELinux context for eachcontainer.  May also be set in PodSecurityContext.  If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is windows.",
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
													Description:         "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
													MarkdownDescription: "The seccomp options to use by this container. If seccomp options areprovided at both the pod & container level, the container optionsoverride the pod options.Note that this field cannot be set when spec.os.name is windows.",
													Attributes: map[string]schema.Attribute{
														"localhost_profile": schema.StringAttribute{
															Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
															MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"type": schema.StringAttribute{
															Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
															MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
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
													Description:         "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
													MarkdownDescription: "The Windows specific settings applied to all containers.If unspecified, the options from the PodSecurityContext will be used.If set in both SecurityContext and PodSecurityContext, the value specified in SecurityContext takes precedence.Note that this field cannot be set when spec.os.name is linux.",
													Attributes: map[string]schema.Attribute{
														"gmsa_credential_spec": schema.StringAttribute{
															Description:         "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
															MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook(https://github.com/kubernetes-sigs/windows-gmsa) inlines the contents of theGMSA credential spec named by the GMSACredentialSpecName field.",
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
															Description:         "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
															MarkdownDescription: "HostProcess determines if a container should be run as a 'Host Process' container.All of a Pod's containers must have the same effective HostProcess value(it is not allowed to have a mix of HostProcess containers and non-HostProcess containers).In addition, if HostProcess is true then HostNetwork must also be set to true.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"run_as_user_name": schema.StringAttribute{
															Description:         "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
															MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process.Defaults to the user specified in image metadata if unspecified.May also be set in PodSecurityContext. If set in both SecurityContext andPodSecurityContext, the value specified in SecurityContext takes precedence.",
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
											Description:         "StartupProbe indicates that the Pod has successfully initialized.If specified, no other probes are executed until this completes successfully.If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,when it might take a long time to load data or warm a cache, than during steady-state operation.This cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized.If specified, no other probes are executed until this completes successfully.If this probe fails, the Pod will be restarted, just as if the livenessProbe failed.This can be used to provide different probe parameters at the beginning of a Pod's lifecycle,when it might take a long time to load data or warm a cache, than during steady-state operation.This cannot be updated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
											Attributes: map[string]schema.Attribute{
												"exec": schema.SingleNestedAttribute{
													Description:         "Exec specifies the action to take.",
													MarkdownDescription: "Exec specifies the action to take.",
													Attributes: map[string]schema.Attribute{
														"command": schema.ListAttribute{
															Description:         "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															MarkdownDescription: "Command is the command line to execute inside the container, the working directory for thecommand  is root ('/') in the container's filesystem. The command is simply exec'd, it isnot run inside a shell, so traditional shell instructions ('|', etc) won't work. To usea shell, you need to explicitly call out to that shell.Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
															ElementType:         types.StringType,
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
													Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded.Defaults to 3. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"grpc": schema.SingleNestedAttribute{
													Description:         "GRPC specifies an action involving a GRPC port.",
													MarkdownDescription: "GRPC specifies an action involving a GRPC port.",
													Attributes: map[string]schema.Attribute{
														"port": schema.Int64Attribute{
															Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"service": schema.StringAttribute{
															Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
															MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest(see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).If this is not specified, the default behavior is defined by gRPC.",
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
													Description:         "HTTPGet specifies the http request to perform.",
													MarkdownDescription: "HTTPGet specifies the http request to perform.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set'Host' in httpHeaders instead.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"http_headers": schema.ListNestedAttribute{
															Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
															MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"name": schema.StringAttribute{
																		Description:         "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		MarkdownDescription: "The header field name.This will be canonicalized upon output, so case-variant names will be understood as the same header.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "The header field value",
																		MarkdownDescription: "The header field value",
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
															Description:         "Path to access on the HTTP server.",
															MarkdownDescription: "Path to access on the HTTP server.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Name or number of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"scheme": schema.StringAttribute{
															Description:         "Scheme to use for connecting to the host.Defaults to HTTP.",
															MarkdownDescription: "Scheme to use for connecting to the host.Defaults to HTTP.",
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
													Description:         "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"period_seconds": schema.Int64Attribute{
													Description:         "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													MarkdownDescription: "How often (in seconds) to perform the probe.Default to 10 seconds. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"success_threshold": schema.Int64Attribute{
													Description:         "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed.Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tcp_socket": schema.SingleNestedAttribute{
													Description:         "TCPSocket specifies an action involving a TCP port.",
													MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",
													Attributes: map[string]schema.Attribute{
														"host": schema.StringAttribute{
															Description:         "Optional: Host name to connect to, defaults to the pod IP.",
															MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
															MarkdownDescription: "Number or name of the port to access on the container.Number must be in the range 1 to 65535.Name must be an IANA_SVC_NAME.",
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
													Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure.The grace period is the duration in seconds after the processes running in the pod are senta termination signal and the time when the processes are forcibly halted with a kill signal.Set this value longer than the expected cleanup time for your process.If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, thisvalue overrides the value provided by the pod spec.Value must be non-negative integer. The value zero indicates stop immediately viathe kill signal (no opportunity to shut down).This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate.Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"timeout_seconds": schema.Int64Attribute{
													Description:         "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
													MarkdownDescription: "Number of seconds after which the probe times out.Defaults to 1 second. Minimum value is 1.More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
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
											Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If thisis not set, reads from stdin in the container will always result in EOF.Default is false.",
											MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If thisis not set, reads from stdin in the container will always result in EOF.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"stdin_once": schema.BoolAttribute{
											Description:         "Whether the container runtime should close the stdin channel after it has been opened bya single attach. When stdin is true the stdin stream will remain open across multiple attachsessions. If stdinOnce is set to true, stdin is opened on container start, is empty until thefirst client attaches to stdin, and then remains open and accepts data until the client disconnects,at which time stdin is closed and remains closed until the container is restarted. If thisflag is false, a container processes that reads from stdin will never receive an EOF.Default is false",
											MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened bya single attach. When stdin is true the stdin stream will remain open across multiple attachsessions. If stdinOnce is set to true, stdin is opened on container start, is empty until thefirst client attaches to stdin, and then remains open and accepts data until the client disconnects,at which time stdin is closed and remains closed until the container is restarted. If thisflag is false, a container processes that reads from stdin will never receive an EOF.Default is false",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"termination_message_path": schema.StringAttribute{
											Description:         "Optional: Path at which the file to which the container's termination messagewill be written is mounted into the container's filesystem.Message written is intended to be brief final status, such as an assertion failure message.Will be truncated by the node if greater than 4096 bytes. The total message length acrossall containers will be limited to 12kb.Defaults to /dev/termination-log.Cannot be updated.",
											MarkdownDescription: "Optional: Path at which the file to which the container's termination messagewill be written is mounted into the container's filesystem.Message written is intended to be brief final status, such as an assertion failure message.Will be truncated by the node if greater than 4096 bytes. The total message length acrossall containers will be limited to 12kb.Defaults to /dev/termination-log.Cannot be updated.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"termination_message_policy": schema.StringAttribute{
											Description:         "Indicate how the termination message should be populated. File will use the contents ofterminationMessagePath to populate the container status message on both success and failure.FallbackToLogsOnError will use the last chunk of container log output if the terminationmessage file is empty and the container exited with an error.The log output is limited to 2048 bytes or 80 lines, whichever is smaller.Defaults to File.Cannot be updated.",
											MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents ofterminationMessagePath to populate the container status message on both success and failure.FallbackToLogsOnError will use the last chunk of container log output if the terminationmessage file is empty and the container exited with an error.The log output is limited to 2048 bytes or 80 lines, whichever is smaller.Defaults to File.Cannot be updated.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tty": schema.BoolAttribute{
											Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.Default is false.",
											MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true.Default is false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"volume_devices": schema.ListNestedAttribute{
											Description:         "volumeDevices is the list of block devices to be used by the container.",
											MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"device_path": schema.StringAttribute{
														Description:         "devicePath is the path inside of the container that the device will be mapped to.",
														MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "name must match the name of a persistentVolumeClaim in the pod",
														MarkdownDescription: "name must match the name of a persistentVolumeClaim in the pod",
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
											Description:         "Pod volumes to mount into the container's filesystem.Cannot be updated.",
											MarkdownDescription: "Pod volumes to mount into the container's filesystem.Cannot be updated.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"mount_path": schema.StringAttribute{
														Description:         "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
														MarkdownDescription: "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"mount_propagation": schema.StringAttribute{
														Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
														Required:            false,
														Optional:            true,
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
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sub_path": schema.StringAttribute{
														Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
														MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"sub_path_expr": schema.StringAttribute{
														Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
														MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
														Required:            false,
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
											Description:         "Container's working directory.If not specified, the container runtime's default will be used, whichmight be configured in the container image.Cannot be updated.",
											MarkdownDescription: "Container's working directory.If not specified, the container runtime's default will be used, whichmight be configured in the container image.Cannot be updated.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations for this pod.",
								MarkdownDescription: "Tolerations for this pod.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
											Required:            false,
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

					"rack_config": schema.SingleNestedAttribute{
						Description:         "RackConfig Configures the operator to deploy rack aware Aerospike cluster.Pods will be deployed in given racks based on given configuration",
						MarkdownDescription: "RackConfig Configures the operator to deploy rack aware Aerospike cluster.Pods will be deployed in given racks based on given configuration",
						Attributes: map[string]schema.Attribute{
							"max_ignorable_pods": schema.StringAttribute{
								Description:         "MaxIgnorablePods is the maximum number/percentage of pending/failed pods in a rack that are ignored whileassessing cluster stability. Pods identified using this value are not considered part of the cluster.Additionally, in SC mode clusters, these pods are removed from the roster.This is particularly useful when some pods are stuck in pending/failed state due to any scheduling issues andcannot be fixed by simply updating the CR.It enables the operator to perform specific operations on the cluster, like changing Aerospike configurations,without being hindered by these problematic pods.Remember to set MaxIgnorablePods back to 0 once the required operation is done.This makes sure that later on, all pods are properly counted when evaluating the cluster stability.",
								MarkdownDescription: "MaxIgnorablePods is the maximum number/percentage of pending/failed pods in a rack that are ignored whileassessing cluster stability. Pods identified using this value are not considered part of the cluster.Additionally, in SC mode clusters, these pods are removed from the roster.This is particularly useful when some pods are stuck in pending/failed state due to any scheduling issues andcannot be fixed by simply updating the CR.It enables the operator to perform specific operations on the cluster, like changing Aerospike configurations,without being hindered by these problematic pods.Remember to set MaxIgnorablePods back to 0 once the required operation is done.This makes sure that later on, all pods are properly counted when evaluating the cluster stability.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"namespaces": schema.ListAttribute{
								Description:         "List of Aerospike namespaces for which rack feature will be enabled",
								MarkdownDescription: "List of Aerospike namespaces for which rack feature will be enabled",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"racks": schema.ListNestedAttribute{
								Description:         "Racks is the list of all racks",
								MarkdownDescription: "Racks is the list of all racks",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"aerospike_config": schema.MapAttribute{
											Description:         "AerospikeConfig overrides the common AerospikeConfig for this Rack. This is merged with global Aerospike config.",
											MarkdownDescription: "AerospikeConfig overrides the common AerospikeConfig for this Rack. This is merged with global Aerospike config.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"effective_aerospike_config": schema.MapAttribute{
											Description:         "Effective/operative Aerospike config. The resultant is a merge of rack Aerospike config and the globalAerospike config",
											MarkdownDescription: "Effective/operative Aerospike config. The resultant is a merge of rack Aerospike config and the globalAerospike config",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"effective_pod_spec": schema.SingleNestedAttribute{
											Description:         "Effective/operative PodSpec. The resultant is user input if specified else global PodSpec",
											MarkdownDescription: "Effective/operative PodSpec. The resultant is user input if specified else global PodSpec",
											Attributes: map[string]schema.Attribute{
												"affinity": schema.SingleNestedAttribute{
													Description:         "Affinity rules for pod placement.",
													MarkdownDescription: "Affinity rules for pod placement.",
													Attributes: map[string]schema.Attribute{
														"node_affinity": schema.SingleNestedAttribute{
															Description:         "Describes node affinity scheduling rules for the pod.",
															MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
															Attributes: map[string]schema.Attribute{
																"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																	Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
																	MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"preference": schema.SingleNestedAttribute{
																				Description:         "A node selector term, associated with the corresponding weight.",
																				MarkdownDescription: "A node selector term, associated with the corresponding weight.",
																				Attributes: map[string]schema.Attribute{
																					"match_expressions": schema.ListNestedAttribute{
																						Description:         "A list of node selector requirements by node's labels.",
																						MarkdownDescription: "A list of node selector requirements by node's labels.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "The label key that the selector applies to.",
																									MarkdownDescription: "The label key that the selector applies to.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"operator": schema.StringAttribute{
																									Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																						Description:         "A list of node selector requirements by node's fields.",
																						MarkdownDescription: "A list of node selector requirements by node's fields.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "The label key that the selector applies to.",
																									MarkdownDescription: "The label key that the selector applies to.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"operator": schema.StringAttribute{
																									Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																				Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																				MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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
																	Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
																	MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
																	Attributes: map[string]schema.Attribute{
																		"node_selector_terms": schema.ListNestedAttribute{
																			Description:         "Required. A list of node selector terms. The terms are ORed.",
																			MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"match_expressions": schema.ListNestedAttribute{
																						Description:         "A list of node selector requirements by node's labels.",
																						MarkdownDescription: "A list of node selector requirements by node's labels.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "The label key that the selector applies to.",
																									MarkdownDescription: "The label key that the selector applies to.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"operator": schema.StringAttribute{
																									Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																						Description:         "A list of node selector requirements by node's fields.",
																						MarkdownDescription: "A list of node selector requirements by node's fields.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "The label key that the selector applies to.",
																									MarkdownDescription: "The label key that the selector applies to.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"operator": schema.StringAttribute{
																									Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
															Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
															MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
															Attributes: map[string]schema.Attribute{
																"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																	Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																	MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"pod_affinity_term": schema.SingleNestedAttribute{
																				Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																				MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																				Attributes: map[string]schema.Attribute{
																					"label_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																						MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								ElementType:         types.StringType,
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
																						Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"mismatch_label_keys": schema.ListAttribute{
																						Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"namespace_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																						MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								ElementType:         types.StringType,
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
																						Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"topology_key": schema.StringAttribute{
																						Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																						MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
																				Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																				MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
																	Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
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
																				Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"mismatch_label_keys": schema.ListAttribute{
																				Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"namespace_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
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
																				Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
															Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
															MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
															Attributes: map[string]schema.Attribute{
																"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																	Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																	MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"pod_affinity_term": schema.SingleNestedAttribute{
																				Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																				MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																				Attributes: map[string]schema.Attribute{
																					"label_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																						MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								ElementType:         types.StringType,
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
																						Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"mismatch_label_keys": schema.ListAttribute{
																						Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"namespace_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																						MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								ElementType:         types.StringType,
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
																						Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"topology_key": schema.StringAttribute{
																						Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																						MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
																				Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																				MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
																	Description:         "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	MarkdownDescription: "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
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
																				Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"mismatch_label_keys": schema.ListAttribute{
																				Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"namespace_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
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
																				Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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

												"node_selector": schema.MapAttribute{
													Description:         "NodeSelector constraints for this pod.",
													MarkdownDescription: "NodeSelector constraints for this pod.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tolerations": schema.ListNestedAttribute{
													Description:         "Tolerations for this pod.",
													MarkdownDescription: "Tolerations for this pod.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"effect": schema.StringAttribute{
																Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key": schema.StringAttribute{
																Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
																MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"toleration_seconds": schema.Int64Attribute{
																Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
																MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
																MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
																Required:            false,
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

										"effective_storage": schema.SingleNestedAttribute{
											Description:         "Effective/operative storage. The resultant is user input if specified else global storage",
											MarkdownDescription: "Effective/operative storage. The resultant is user input if specified else global storage",
											Attributes: map[string]schema.Attribute{
												"block_volume_policy": schema.SingleNestedAttribute{
													Description:         "BlockVolumePolicy contains default policies for block volumes.",
													MarkdownDescription: "BlockVolumePolicy contains default policies for block volumes.",
													Attributes: map[string]schema.Attribute{
														"cascade_delete": schema.BoolAttribute{
															Description:         "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
															MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"effective_cascade_delete": schema.BoolAttribute{
															Description:         "Effective/operative value to use for cascade delete after applying defaults.",
															MarkdownDescription: "Effective/operative value to use for cascade delete after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"effective_init_method": schema.StringAttribute{
															Description:         "Effective/operative value to use as the volume init method after applying defaults.",
															MarkdownDescription: "Effective/operative value to use as the volume init method after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"effective_wipe_method": schema.StringAttribute{
															Description:         "Effective/operative value to use as the volume wipe method after applying defaults.",
															MarkdownDescription: "Effective/operative value to use as the volume wipe method after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"init_method": schema.StringAttribute{
															Description:         "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
															MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"wipe_method": schema.StringAttribute{
															Description:         "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
															MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"cleanup_threads": schema.Int64Attribute{
													Description:         "CleanupThreads contains the maximum number of cleanup threads(dd or blkdiscard) per init container.",
													MarkdownDescription: "CleanupThreads contains the maximum number of cleanup threads(dd or blkdiscard) per init container.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"filesystem_volume_policy": schema.SingleNestedAttribute{
													Description:         "FileSystemVolumePolicy contains default policies for filesystem volumes.",
													MarkdownDescription: "FileSystemVolumePolicy contains default policies for filesystem volumes.",
													Attributes: map[string]schema.Attribute{
														"cascade_delete": schema.BoolAttribute{
															Description:         "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
															MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"effective_cascade_delete": schema.BoolAttribute{
															Description:         "Effective/operative value to use for cascade delete after applying defaults.",
															MarkdownDescription: "Effective/operative value to use for cascade delete after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"effective_init_method": schema.StringAttribute{
															Description:         "Effective/operative value to use as the volume init method after applying defaults.",
															MarkdownDescription: "Effective/operative value to use as the volume init method after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"effective_wipe_method": schema.StringAttribute{
															Description:         "Effective/operative value to use as the volume wipe method after applying defaults.",
															MarkdownDescription: "Effective/operative value to use as the volume wipe method after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"init_method": schema.StringAttribute{
															Description:         "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
															MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"wipe_method": schema.StringAttribute{
															Description:         "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
															MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"local_storage_classes": schema.ListAttribute{
													Description:         "LocalStorageClasses contains a list of storage classes which provisions local volumes.",
													MarkdownDescription: "LocalStorageClasses contains a list of storage classes which provisions local volumes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volumes": schema.ListNestedAttribute{
													Description:         "Volumes list to attach to created pods.",
													MarkdownDescription: "Volumes list to attach to created pods.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"aerospike": schema.SingleNestedAttribute{
																Description:         "Aerospike attachment of this volume on Aerospike server container.",
																MarkdownDescription: "Aerospike attachment of this volume on Aerospike server container.",
																Attributes: map[string]schema.Attribute{
																	"mount_options": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"mount_propagation": schema.StringAttribute{
																				Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																				MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"read_only": schema.BoolAttribute{
																				Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																				MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"sub_path": schema.StringAttribute{
																				Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																				MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"sub_path_expr": schema.StringAttribute{
																				Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																				MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": schema.StringAttribute{
																		Description:         "Path to attach the volume on the Aerospike server container.",
																		MarkdownDescription: "Path to attach the volume on the Aerospike server container.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"cascade_delete": schema.BoolAttribute{
																Description:         "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
																MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"effective_cascade_delete": schema.BoolAttribute{
																Description:         "Effective/operative value to use for cascade delete after applying defaults.",
																MarkdownDescription: "Effective/operative value to use for cascade delete after applying defaults.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"effective_init_method": schema.StringAttribute{
																Description:         "Effective/operative value to use as the volume init method after applying defaults.",
																MarkdownDescription: "Effective/operative value to use as the volume init method after applying defaults.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
																},
															},

															"effective_wipe_method": schema.StringAttribute{
																Description:         "Effective/operative value to use as the volume wipe method after applying defaults.",
																MarkdownDescription: "Effective/operative value to use as the volume wipe method after applying defaults.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
																},
															},

															"init_containers": schema.ListNestedAttribute{
																Description:         "InitContainers are additional init containers where this volume will be mounted",
																MarkdownDescription: "InitContainers are additional init containers where this volume will be mounted",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"container_name": schema.StringAttribute{
																			Description:         "ContainerName is the name of the container to attach this volume to.",
																			MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mount_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"mount_propagation": schema.StringAttribute{
																					Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																					MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"read_only": schema.BoolAttribute{
																					Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																					MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"sub_path": schema.StringAttribute{
																					Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																					MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"sub_path_expr": schema.StringAttribute{
																					Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																					MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "Path to attach the volume on the container.",
																			MarkdownDescription: "Path to attach the volume on the container.",
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

															"init_method": schema.StringAttribute{
																Description:         "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
																MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "Name for this volume, Name or path should be given.",
																MarkdownDescription: "Name for this volume, Name or path should be given.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"sidecars": schema.ListNestedAttribute{
																Description:         "Sidecars are side containers where this volume will be mounted",
																MarkdownDescription: "Sidecars are side containers where this volume will be mounted",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"container_name": schema.StringAttribute{
																			Description:         "ContainerName is the name of the container to attach this volume to.",
																			MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mount_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"mount_propagation": schema.StringAttribute{
																					Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																					MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"read_only": schema.BoolAttribute{
																					Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																					MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"sub_path": schema.StringAttribute{
																					Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																					MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"sub_path_expr": schema.StringAttribute{
																					Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																					MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "Path to attach the volume on the container.",
																			MarkdownDescription: "Path to attach the volume on the container.",
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

															"source": schema.SingleNestedAttribute{
																Description:         "Source of this volume.",
																MarkdownDescription: "Source of this volume.",
																Attributes: map[string]schema.Attribute{
																	"config_map": schema.SingleNestedAttribute{
																		Description:         "ConfigMap represents a configMap that should populate this volume",
																		MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",
																		Attributes: map[string]schema.Attribute{
																			"default_mode": schema.Int64Attribute{
																				Description:         "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"items": schema.ListNestedAttribute{
																				Description:         "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																				MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "key is the key to project.",
																							MarkdownDescription: "key is the key to project.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"mode": schema.Int64Attribute{
																							Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
																							Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																							MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "optional specify whether the ConfigMap or its keys must be defined",
																				MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
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
																		Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																		MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																		Attributes: map[string]schema.Attribute{
																			"medium": schema.StringAttribute{
																				Description:         "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																				MarkdownDescription: "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"size_limit": schema.StringAttribute{
																				Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																				MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"persistent_volume": schema.SingleNestedAttribute{
																		Description:         "PersistentVolumeSpec describes a persistent volume to claim and attach to Aerospike pods.",
																		MarkdownDescription: "PersistentVolumeSpec describes a persistent volume to claim and attach to Aerospike pods.",
																		Attributes: map[string]schema.Attribute{
																			"access_modes": schema.ListAttribute{
																				Description:         "Name for creating PVC for this volume, Name or path should be givenName string 'json:'name''",
																				MarkdownDescription: "Name for creating PVC for this volume, Name or path should be givenName string 'json:'name''",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"metadata": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"annotations": schema.MapAttribute{
																						Description:         "Key - Value pair that may be set by external tools to store and retrieve arbitrary metadata",
																						MarkdownDescription: "Key - Value pair that may be set by external tools to store and retrieve arbitrary metadata",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"labels": schema.MapAttribute{
																						Description:         "Key - Value pairs that can be used to organize and categorize scope and select objects",
																						MarkdownDescription: "Key - Value pairs that can be used to organize and categorize scope and select objects",
																						ElementType:         types.StringType,
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
																				Description:         "A label query over volumes to consider for binding.",
																				MarkdownDescription: "A label query over volumes to consider for binding.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"size": schema.StringAttribute{
																				Description:         "Size of volume.",
																				MarkdownDescription: "Size of volume.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"storage_class": schema.StringAttribute{
																				Description:         "StorageClass should be pre-created by user.",
																				MarkdownDescription: "StorageClass should be pre-created by user.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"volume_mode": schema.StringAttribute{
																				Description:         "VolumeMode specifies if the volume is block/raw or a filesystem.",
																				MarkdownDescription: "VolumeMode specifies if the volume is block/raw or a filesystem.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret": schema.SingleNestedAttribute{
																		Description:         "Adapts a Secret into a volume.The contents of the target Secret's Data field will be presented in a volumeas files using the keys in the Data field as the file names.Secret volumes support ownership management and SELinux relabeling.",
																		MarkdownDescription: "Adapts a Secret into a volume.The contents of the target Secret's Data field will be presented in a volumeas files using the keys in the Data field as the file names.Secret volumes support ownership management and SELinux relabeling.",
																		Attributes: map[string]schema.Attribute{
																			"default_mode": schema.Int64Attribute{
																				Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"items": schema.ListNestedAttribute{
																				Description:         "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																				MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "key is the key to project.",
																							MarkdownDescription: "key is the key to project.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"mode": schema.Int64Attribute{
																							Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
																							Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																							MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
																				Description:         "optional field specify whether the Secret or its keys must be defined",
																				MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"secret_name": schema.StringAttribute{
																				Description:         "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																				MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
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

															"wipe_method": schema.StringAttribute{
																Description:         "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
																MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
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

										"id": schema.Int64Attribute{
											Description:         "Identifier for the rack",
											MarkdownDescription: "Identifier for the rack",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"node_name": schema.StringAttribute{
											Description:         "K8s Node name for setting rack affinity. Rack pods will be deployed in given k8s Node",
											MarkdownDescription: "K8s Node name for setting rack affinity. Rack pods will be deployed in given k8s Node",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pod_spec": schema.SingleNestedAttribute{
											Description:         "PodSpec to use for the pods in this rack. This value overwrites the global storage config",
											MarkdownDescription: "PodSpec to use for the pods in this rack. This value overwrites the global storage config",
											Attributes: map[string]schema.Attribute{
												"affinity": schema.SingleNestedAttribute{
													Description:         "Affinity rules for pod placement.",
													MarkdownDescription: "Affinity rules for pod placement.",
													Attributes: map[string]schema.Attribute{
														"node_affinity": schema.SingleNestedAttribute{
															Description:         "Describes node affinity scheduling rules for the pod.",
															MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
															Attributes: map[string]schema.Attribute{
																"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																	Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
																	MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"preference": schema.SingleNestedAttribute{
																				Description:         "A node selector term, associated with the corresponding weight.",
																				MarkdownDescription: "A node selector term, associated with the corresponding weight.",
																				Attributes: map[string]schema.Attribute{
																					"match_expressions": schema.ListNestedAttribute{
																						Description:         "A list of node selector requirements by node's labels.",
																						MarkdownDescription: "A list of node selector requirements by node's labels.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "The label key that the selector applies to.",
																									MarkdownDescription: "The label key that the selector applies to.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"operator": schema.StringAttribute{
																									Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																						Description:         "A list of node selector requirements by node's fields.",
																						MarkdownDescription: "A list of node selector requirements by node's fields.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "The label key that the selector applies to.",
																									MarkdownDescription: "The label key that the selector applies to.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"operator": schema.StringAttribute{
																									Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																				Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																				MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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
																	Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
																	MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
																	Attributes: map[string]schema.Attribute{
																		"node_selector_terms": schema.ListNestedAttribute{
																			Description:         "Required. A list of node selector terms. The terms are ORed.",
																			MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"match_expressions": schema.ListNestedAttribute{
																						Description:         "A list of node selector requirements by node's labels.",
																						MarkdownDescription: "A list of node selector requirements by node's labels.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "The label key that the selector applies to.",
																									MarkdownDescription: "The label key that the selector applies to.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"operator": schema.StringAttribute{
																									Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
																						Description:         "A list of node selector requirements by node's fields.",
																						MarkdownDescription: "A list of node selector requirements by node's fields.",
																						NestedObject: schema.NestedAttributeObject{
																							Attributes: map[string]schema.Attribute{
																								"key": schema.StringAttribute{
																									Description:         "The label key that the selector applies to.",
																									MarkdownDescription: "The label key that the selector applies to.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"operator": schema.StringAttribute{
																									Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
															Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
															MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
															Attributes: map[string]schema.Attribute{
																"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																	Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																	MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"pod_affinity_term": schema.SingleNestedAttribute{
																				Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																				MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																				Attributes: map[string]schema.Attribute{
																					"label_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																						MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								ElementType:         types.StringType,
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
																						Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"mismatch_label_keys": schema.ListAttribute{
																						Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"namespace_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																						MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								ElementType:         types.StringType,
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
																						Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"topology_key": schema.StringAttribute{
																						Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																						MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
																				Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																				MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
																	Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
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
																				Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"mismatch_label_keys": schema.ListAttribute{
																				Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"namespace_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
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
																				Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
															Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
															MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
															Attributes: map[string]schema.Attribute{
																"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																	Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																	MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"pod_affinity_term": schema.SingleNestedAttribute{
																				Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																				MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																				Attributes: map[string]schema.Attribute{
																					"label_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																						MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								ElementType:         types.StringType,
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
																						Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"mismatch_label_keys": schema.ListAttribute{
																						Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"namespace_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																						MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																											Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								ElementType:         types.StringType,
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
																						Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"topology_key": schema.StringAttribute{
																						Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																						MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
																				Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
																				MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
																	Description:         "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	MarkdownDescription: "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
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
																				Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.Also, MatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"mismatch_label_keys": schema.ListAttribute{
																				Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'LabelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both MismatchLabelKeys and LabelSelector.Also, MismatchLabelKeys cannot be set when LabelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"namespace_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
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
																				Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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

												"node_selector": schema.MapAttribute{
													Description:         "NodeSelector constraints for this pod.",
													MarkdownDescription: "NodeSelector constraints for this pod.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tolerations": schema.ListNestedAttribute{
													Description:         "Tolerations for this pod.",
													MarkdownDescription: "Tolerations for this pod.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"effect": schema.StringAttribute{
																Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key": schema.StringAttribute{
																Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
																MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"toleration_seconds": schema.Int64Attribute{
																Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
																MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
																MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
																Required:            false,
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

										"rack_label": schema.StringAttribute{
											Description:         "RackLabel for setting rack affinity.Rack pods will be deployed in k8s nodes having rackLabel {aerospike.com/rack-label: <rack-label>}",
											MarkdownDescription: "RackLabel for setting rack affinity.Rack pods will be deployed in k8s nodes having rackLabel {aerospike.com/rack-label: <rack-label>}",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"region": schema.StringAttribute{
											Description:         "Region name for setting rack affinity. Rack pods will be deployed to given Region",
											MarkdownDescription: "Region name for setting rack affinity. Rack pods will be deployed to given Region",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"storage": schema.SingleNestedAttribute{
											Description:         "Storage specify persistent storage to use for the pods in this rack. This value overwrites the global storage config",
											MarkdownDescription: "Storage specify persistent storage to use for the pods in this rack. This value overwrites the global storage config",
											Attributes: map[string]schema.Attribute{
												"block_volume_policy": schema.SingleNestedAttribute{
													Description:         "BlockVolumePolicy contains default policies for block volumes.",
													MarkdownDescription: "BlockVolumePolicy contains default policies for block volumes.",
													Attributes: map[string]schema.Attribute{
														"cascade_delete": schema.BoolAttribute{
															Description:         "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
															MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"effective_cascade_delete": schema.BoolAttribute{
															Description:         "Effective/operative value to use for cascade delete after applying defaults.",
															MarkdownDescription: "Effective/operative value to use for cascade delete after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"effective_init_method": schema.StringAttribute{
															Description:         "Effective/operative value to use as the volume init method after applying defaults.",
															MarkdownDescription: "Effective/operative value to use as the volume init method after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"effective_wipe_method": schema.StringAttribute{
															Description:         "Effective/operative value to use as the volume wipe method after applying defaults.",
															MarkdownDescription: "Effective/operative value to use as the volume wipe method after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"init_method": schema.StringAttribute{
															Description:         "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
															MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"wipe_method": schema.StringAttribute{
															Description:         "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
															MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"cleanup_threads": schema.Int64Attribute{
													Description:         "CleanupThreads contains the maximum number of cleanup threads(dd or blkdiscard) per init container.",
													MarkdownDescription: "CleanupThreads contains the maximum number of cleanup threads(dd or blkdiscard) per init container.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"filesystem_volume_policy": schema.SingleNestedAttribute{
													Description:         "FileSystemVolumePolicy contains default policies for filesystem volumes.",
													MarkdownDescription: "FileSystemVolumePolicy contains default policies for filesystem volumes.",
													Attributes: map[string]schema.Attribute{
														"cascade_delete": schema.BoolAttribute{
															Description:         "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
															MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"effective_cascade_delete": schema.BoolAttribute{
															Description:         "Effective/operative value to use for cascade delete after applying defaults.",
															MarkdownDescription: "Effective/operative value to use for cascade delete after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"effective_init_method": schema.StringAttribute{
															Description:         "Effective/operative value to use as the volume init method after applying defaults.",
															MarkdownDescription: "Effective/operative value to use as the volume init method after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"effective_wipe_method": schema.StringAttribute{
															Description:         "Effective/operative value to use as the volume wipe method after applying defaults.",
															MarkdownDescription: "Effective/operative value to use as the volume wipe method after applying defaults.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"init_method": schema.StringAttribute{
															Description:         "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
															MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},

														"wipe_method": schema.StringAttribute{
															Description:         "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
															MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
															Required:            false,
															Optional:            true,
															Computed:            false,
															Validators: []validator.String{
																stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
															},
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"local_storage_classes": schema.ListAttribute{
													Description:         "LocalStorageClasses contains a list of storage classes which provisions local volumes.",
													MarkdownDescription: "LocalStorageClasses contains a list of storage classes which provisions local volumes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volumes": schema.ListNestedAttribute{
													Description:         "Volumes list to attach to created pods.",
													MarkdownDescription: "Volumes list to attach to created pods.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"aerospike": schema.SingleNestedAttribute{
																Description:         "Aerospike attachment of this volume on Aerospike server container.",
																MarkdownDescription: "Aerospike attachment of this volume on Aerospike server container.",
																Attributes: map[string]schema.Attribute{
																	"mount_options": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"mount_propagation": schema.StringAttribute{
																				Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																				MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"read_only": schema.BoolAttribute{
																				Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																				MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"sub_path": schema.StringAttribute{
																				Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																				MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"sub_path_expr": schema.StringAttribute{
																				Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																				MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": schema.StringAttribute{
																		Description:         "Path to attach the volume on the Aerospike server container.",
																		MarkdownDescription: "Path to attach the volume on the Aerospike server container.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"cascade_delete": schema.BoolAttribute{
																Description:         "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
																MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"effective_cascade_delete": schema.BoolAttribute{
																Description:         "Effective/operative value to use for cascade delete after applying defaults.",
																MarkdownDescription: "Effective/operative value to use for cascade delete after applying defaults.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"effective_init_method": schema.StringAttribute{
																Description:         "Effective/operative value to use as the volume init method after applying defaults.",
																MarkdownDescription: "Effective/operative value to use as the volume init method after applying defaults.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
																},
															},

															"effective_wipe_method": schema.StringAttribute{
																Description:         "Effective/operative value to use as the volume wipe method after applying defaults.",
																MarkdownDescription: "Effective/operative value to use as the volume wipe method after applying defaults.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
																},
															},

															"init_containers": schema.ListNestedAttribute{
																Description:         "InitContainers are additional init containers where this volume will be mounted",
																MarkdownDescription: "InitContainers are additional init containers where this volume will be mounted",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"container_name": schema.StringAttribute{
																			Description:         "ContainerName is the name of the container to attach this volume to.",
																			MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mount_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"mount_propagation": schema.StringAttribute{
																					Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																					MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"read_only": schema.BoolAttribute{
																					Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																					MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"sub_path": schema.StringAttribute{
																					Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																					MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"sub_path_expr": schema.StringAttribute{
																					Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																					MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "Path to attach the volume on the container.",
																			MarkdownDescription: "Path to attach the volume on the container.",
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

															"init_method": schema.StringAttribute{
																Description:         "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
																MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
																},
															},

															"name": schema.StringAttribute{
																Description:         "Name for this volume, Name or path should be given.",
																MarkdownDescription: "Name for this volume, Name or path should be given.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"sidecars": schema.ListNestedAttribute{
																Description:         "Sidecars are side containers where this volume will be mounted",
																MarkdownDescription: "Sidecars are side containers where this volume will be mounted",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"container_name": schema.StringAttribute{
																			Description:         "ContainerName is the name of the container to attach this volume to.",
																			MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"mount_options": schema.SingleNestedAttribute{
																			Description:         "",
																			MarkdownDescription: "",
																			Attributes: map[string]schema.Attribute{
																				"mount_propagation": schema.StringAttribute{
																					Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																					MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"read_only": schema.BoolAttribute{
																					Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																					MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"sub_path": schema.StringAttribute{
																					Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																					MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"sub_path_expr": schema.StringAttribute{
																					Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																					MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "Path to attach the volume on the container.",
																			MarkdownDescription: "Path to attach the volume on the container.",
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

															"source": schema.SingleNestedAttribute{
																Description:         "Source of this volume.",
																MarkdownDescription: "Source of this volume.",
																Attributes: map[string]schema.Attribute{
																	"config_map": schema.SingleNestedAttribute{
																		Description:         "ConfigMap represents a configMap that should populate this volume",
																		MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",
																		Attributes: map[string]schema.Attribute{
																			"default_mode": schema.Int64Attribute{
																				Description:         "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"items": schema.ListNestedAttribute{
																				Description:         "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																				MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "key is the key to project.",
																							MarkdownDescription: "key is the key to project.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"mode": schema.Int64Attribute{
																							Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
																							Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																							MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
																				Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "optional specify whether the ConfigMap or its keys must be defined",
																				MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
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
																		Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																		MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																		Attributes: map[string]schema.Attribute{
																			"medium": schema.StringAttribute{
																				Description:         "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																				MarkdownDescription: "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"size_limit": schema.StringAttribute{
																				Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																				MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"persistent_volume": schema.SingleNestedAttribute{
																		Description:         "PersistentVolumeSpec describes a persistent volume to claim and attach to Aerospike pods.",
																		MarkdownDescription: "PersistentVolumeSpec describes a persistent volume to claim and attach to Aerospike pods.",
																		Attributes: map[string]schema.Attribute{
																			"access_modes": schema.ListAttribute{
																				Description:         "Name for creating PVC for this volume, Name or path should be givenName string 'json:'name''",
																				MarkdownDescription: "Name for creating PVC for this volume, Name or path should be givenName string 'json:'name''",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"metadata": schema.SingleNestedAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Attributes: map[string]schema.Attribute{
																					"annotations": schema.MapAttribute{
																						Description:         "Key - Value pair that may be set by external tools to store and retrieve arbitrary metadata",
																						MarkdownDescription: "Key - Value pair that may be set by external tools to store and retrieve arbitrary metadata",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"labels": schema.MapAttribute{
																						Description:         "Key - Value pairs that can be used to organize and categorize scope and select objects",
																						MarkdownDescription: "Key - Value pairs that can be used to organize and categorize scope and select objects",
																						ElementType:         types.StringType,
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
																				Description:         "A label query over volumes to consider for binding.",
																				MarkdownDescription: "A label query over volumes to consider for binding.",
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
																									Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},
																				},
																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"size": schema.StringAttribute{
																				Description:         "Size of volume.",
																				MarkdownDescription: "Size of volume.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"storage_class": schema.StringAttribute{
																				Description:         "StorageClass should be pre-created by user.",
																				MarkdownDescription: "StorageClass should be pre-created by user.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"volume_mode": schema.StringAttribute{
																				Description:         "VolumeMode specifies if the volume is block/raw or a filesystem.",
																				MarkdownDescription: "VolumeMode specifies if the volume is block/raw or a filesystem.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},
																		},
																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"secret": schema.SingleNestedAttribute{
																		Description:         "Adapts a Secret into a volume.The contents of the target Secret's Data field will be presented in a volumeas files using the keys in the Data field as the file names.Secret volumes support ownership management and SELinux relabeling.",
																		MarkdownDescription: "Adapts a Secret into a volume.The contents of the target Secret's Data field will be presented in a volumeas files using the keys in the Data field as the file names.Secret volumes support ownership management and SELinux relabeling.",
																		Attributes: map[string]schema.Attribute{
																			"default_mode": schema.Int64Attribute{
																				Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"items": schema.ListNestedAttribute{
																				Description:         "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																				MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "key is the key to project.",
																							MarkdownDescription: "key is the key to project.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"mode": schema.Int64Attribute{
																							Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
																							Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																							MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
																				Description:         "optional field specify whether the Secret or its keys must be defined",
																				MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"secret_name": schema.StringAttribute{
																				Description:         "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																				MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
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

															"wipe_method": schema.StringAttribute{
																Description:         "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
																MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
																Required:            false,
																Optional:            true,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
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

										"zone": schema.StringAttribute{
											Description:         "Zone name for setting rack affinity. Rack pods will be deployed to given Zone",
											MarkdownDescription: "Zone name for setting rack affinity. Rack pods will be deployed to given Zone",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rolling_update_batch_size": schema.StringAttribute{
								Description:         "RollingUpdateBatchSize is the percentage/number of rack pods that can be restarted simultaneously",
								MarkdownDescription: "RollingUpdateBatchSize is the percentage/number of rack pods that can be restarted simultaneously",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scale_down_batch_size": schema.StringAttribute{
								Description:         "ScaleDownBatchSize is the percentage/number of rack pods that can be scaled down simultaneously",
								MarkdownDescription: "ScaleDownBatchSize is the percentage/number of rack pods that can be scaled down simultaneously",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"roster_node_block_list": schema.ListAttribute{
						Description:         "RosterNodeBlockList is a list of blocked nodeIDs from roster in a strong-consistency setup",
						MarkdownDescription: "RosterNodeBlockList is a list of blocked nodeIDs from roster in a strong-consistency setup",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"seeds_finder_services": schema.SingleNestedAttribute{
						Description:         "SeedsFinderServices creates additional Kubernetes service that allowclients to discover Aerospike cluster nodes.",
						MarkdownDescription: "SeedsFinderServices creates additional Kubernetes service that allowclients to discover Aerospike cluster nodes.",
						Attributes: map[string]schema.Attribute{
							"load_balancer": schema.SingleNestedAttribute{
								Description:         "LoadBalancer created to discover Aerospike Cluster nodes from outside ofKubernetes cluster.",
								MarkdownDescription: "LoadBalancer created to discover Aerospike Cluster nodes from outside ofKubernetes cluster.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_traffic_policy": schema.StringAttribute{
										Description:         "ServiceExternalTrafficPolicy describes how nodes distribute service traffic theyreceive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs,and LoadBalancer IPs.",
										MarkdownDescription: "ServiceExternalTrafficPolicy describes how nodes distribute service traffic theyreceive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs,and LoadBalancer IPs.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("Local", "Cluster"),
										},
									},

									"load_balancer_source_ranges": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Port Exposed port on load balancer. If not specified TargetPort is used.",
										MarkdownDescription: "Port Exposed port on load balancer. If not specified TargetPort is used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1024),
											int64validator.AtMost(65535),
										},
									},

									"port_name": schema.StringAttribute{
										Description:         "The name of the port exposed on load balancer service.",
										MarkdownDescription: "The name of the port exposed on load balancer service.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"target_port": schema.Int64Attribute{
										Description:         "TargetPort Target port. If not specified the tls-port of network.service stanza is used from Aerospike config.If there is no tls port configured then regular port from network.service is used.",
										MarkdownDescription: "TargetPort Target port. If not specified the tls-port of network.service stanza is used from Aerospike config.If there is no tls port configured then regular port from network.service is used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1024),
											int64validator.AtMost(65535),
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

					"size": schema.Int64Attribute{
						Description:         "Aerospike cluster size",
						MarkdownDescription: "Aerospike cluster size",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Storage specify persistent storage to use for the Aerospike pods",
						MarkdownDescription: "Storage specify persistent storage to use for the Aerospike pods",
						Attributes: map[string]schema.Attribute{
							"block_volume_policy": schema.SingleNestedAttribute{
								Description:         "BlockVolumePolicy contains default policies for block volumes.",
								MarkdownDescription: "BlockVolumePolicy contains default policies for block volumes.",
								Attributes: map[string]schema.Attribute{
									"cascade_delete": schema.BoolAttribute{
										Description:         "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
										MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"effective_cascade_delete": schema.BoolAttribute{
										Description:         "Effective/operative value to use for cascade delete after applying defaults.",
										MarkdownDescription: "Effective/operative value to use for cascade delete after applying defaults.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"effective_init_method": schema.StringAttribute{
										Description:         "Effective/operative value to use as the volume init method after applying defaults.",
										MarkdownDescription: "Effective/operative value to use as the volume init method after applying defaults.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"effective_wipe_method": schema.StringAttribute{
										Description:         "Effective/operative value to use as the volume wipe method after applying defaults.",
										MarkdownDescription: "Effective/operative value to use as the volume wipe method after applying defaults.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"init_method": schema.StringAttribute{
										Description:         "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
										MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"wipe_method": schema.StringAttribute{
										Description:         "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
										MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"cleanup_threads": schema.Int64Attribute{
								Description:         "CleanupThreads contains the maximum number of cleanup threads(dd or blkdiscard) per init container.",
								MarkdownDescription: "CleanupThreads contains the maximum number of cleanup threads(dd or blkdiscard) per init container.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"filesystem_volume_policy": schema.SingleNestedAttribute{
								Description:         "FileSystemVolumePolicy contains default policies for filesystem volumes.",
								MarkdownDescription: "FileSystemVolumePolicy contains default policies for filesystem volumes.",
								Attributes: map[string]schema.Attribute{
									"cascade_delete": schema.BoolAttribute{
										Description:         "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
										MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"effective_cascade_delete": schema.BoolAttribute{
										Description:         "Effective/operative value to use for cascade delete after applying defaults.",
										MarkdownDescription: "Effective/operative value to use for cascade delete after applying defaults.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"effective_init_method": schema.StringAttribute{
										Description:         "Effective/operative value to use as the volume init method after applying defaults.",
										MarkdownDescription: "Effective/operative value to use as the volume init method after applying defaults.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"effective_wipe_method": schema.StringAttribute{
										Description:         "Effective/operative value to use as the volume wipe method after applying defaults.",
										MarkdownDescription: "Effective/operative value to use as the volume wipe method after applying defaults.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"init_method": schema.StringAttribute{
										Description:         "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
										MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"wipe_method": schema.StringAttribute{
										Description:         "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
										MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"local_storage_classes": schema.ListAttribute{
								Description:         "LocalStorageClasses contains a list of storage classes which provisions local volumes.",
								MarkdownDescription: "LocalStorageClasses contains a list of storage classes which provisions local volumes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volumes": schema.ListNestedAttribute{
								Description:         "Volumes list to attach to created pods.",
								MarkdownDescription: "Volumes list to attach to created pods.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"aerospike": schema.SingleNestedAttribute{
											Description:         "Aerospike attachment of this volume on Aerospike server container.",
											MarkdownDescription: "Aerospike attachment of this volume on Aerospike server container.",
											Attributes: map[string]schema.Attribute{
												"mount_options": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"mount_propagation": schema.StringAttribute{
															Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
															MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
															MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sub_path": schema.StringAttribute{
															Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
															MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sub_path_expr": schema.StringAttribute{
															Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
															MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"path": schema.StringAttribute{
													Description:         "Path to attach the volume on the Aerospike server container.",
													MarkdownDescription: "Path to attach the volume on the Aerospike server container.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"cascade_delete": schema.BoolAttribute{
											Description:         "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
											MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after the pod this volume binds to isterminated and removed from the cluster.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"effective_cascade_delete": schema.BoolAttribute{
											Description:         "Effective/operative value to use for cascade delete after applying defaults.",
											MarkdownDescription: "Effective/operative value to use for cascade delete after applying defaults.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"effective_init_method": schema.StringAttribute{
											Description:         "Effective/operative value to use as the volume init method after applying defaults.",
											MarkdownDescription: "Effective/operative value to use as the volume init method after applying defaults.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
											},
										},

										"effective_wipe_method": schema.StringAttribute{
											Description:         "Effective/operative value to use as the volume wipe method after applying defaults.",
											MarkdownDescription: "Effective/operative value to use as the volume wipe method after applying defaults.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
											},
										},

										"init_containers": schema.ListNestedAttribute{
											Description:         "InitContainers are additional init containers where this volume will be mounted",
											MarkdownDescription: "InitContainers are additional init containers where this volume will be mounted",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_name": schema.StringAttribute{
														Description:         "ContainerName is the name of the container to attach this volume to.",
														MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"mount_options": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"mount_propagation": schema.StringAttribute{
																Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sub_path": schema.StringAttribute{
																Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sub_path_expr": schema.StringAttribute{
																Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": schema.StringAttribute{
														Description:         "Path to attach the volume on the container.",
														MarkdownDescription: "Path to attach the volume on the container.",
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

										"init_method": schema.StringAttribute{
											Description:         "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
											MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods are initialized when the pods come up thefirst time. Defaults to 'none'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
											},
										},

										"name": schema.StringAttribute{
											Description:         "Name for this volume, Name or path should be given.",
											MarkdownDescription: "Name for this volume, Name or path should be given.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"sidecars": schema.ListNestedAttribute{
											Description:         "Sidecars are side containers where this volume will be mounted",
											MarkdownDescription: "Sidecars are side containers where this volume will be mounted",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"container_name": schema.StringAttribute{
														Description:         "ContainerName is the name of the container to attach this volume to.",
														MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"mount_options": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"mount_propagation": schema.StringAttribute{
																Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"read_only": schema.BoolAttribute{
																Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sub_path": schema.StringAttribute{
																Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"sub_path_expr": schema.StringAttribute{
																Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": schema.StringAttribute{
														Description:         "Path to attach the volume on the container.",
														MarkdownDescription: "Path to attach the volume on the container.",
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

										"source": schema.SingleNestedAttribute{
											Description:         "Source of this volume.",
											MarkdownDescription: "Source of this volume.",
											Attributes: map[string]schema.Attribute{
												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap represents a configMap that should populate this volume",
													MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",
													Attributes: map[string]schema.Attribute{
														"default_mode": schema.Int64Attribute{
															Description:         "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
															MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"items": schema.ListNestedAttribute{
															Description:         "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
															MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the key to project.",
																		MarkdownDescription: "key is the key to project.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"mode": schema.Int64Attribute{
																		Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"path": schema.StringAttribute{
																		Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																		MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
															Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"optional": schema.BoolAttribute{
															Description:         "optional specify whether the ConfigMap or its keys must be defined",
															MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
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
													Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
													MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
													Attributes: map[string]schema.Attribute{
														"medium": schema.StringAttribute{
															Description:         "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															MarkdownDescription: "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"size_limit": schema.StringAttribute{
															Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"persistent_volume": schema.SingleNestedAttribute{
													Description:         "PersistentVolumeSpec describes a persistent volume to claim and attach to Aerospike pods.",
													MarkdownDescription: "PersistentVolumeSpec describes a persistent volume to claim and attach to Aerospike pods.",
													Attributes: map[string]schema.Attribute{
														"access_modes": schema.ListAttribute{
															Description:         "Name for creating PVC for this volume, Name or path should be givenName string 'json:'name''",
															MarkdownDescription: "Name for creating PVC for this volume, Name or path should be givenName string 'json:'name''",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"metadata": schema.SingleNestedAttribute{
															Description:         "",
															MarkdownDescription: "",
															Attributes: map[string]schema.Attribute{
																"annotations": schema.MapAttribute{
																	Description:         "Key - Value pair that may be set by external tools to store and retrieve arbitrary metadata",
																	MarkdownDescription: "Key - Value pair that may be set by external tools to store and retrieve arbitrary metadata",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"labels": schema.MapAttribute{
																	Description:         "Key - Value pairs that can be used to organize and categorize scope and select objects",
																	MarkdownDescription: "Key - Value pairs that can be used to organize and categorize scope and select objects",
																	ElementType:         types.StringType,
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
															Description:         "A label query over volumes to consider for binding.",
															MarkdownDescription: "A label query over volumes to consider for binding.",
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
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"size": schema.StringAttribute{
															Description:         "Size of volume.",
															MarkdownDescription: "Size of volume.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"storage_class": schema.StringAttribute{
															Description:         "StorageClass should be pre-created by user.",
															MarkdownDescription: "StorageClass should be pre-created by user.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"volume_mode": schema.StringAttribute{
															Description:         "VolumeMode specifies if the volume is block/raw or a filesystem.",
															MarkdownDescription: "VolumeMode specifies if the volume is block/raw or a filesystem.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"secret": schema.SingleNestedAttribute{
													Description:         "Adapts a Secret into a volume.The contents of the target Secret's Data field will be presented in a volumeas files using the keys in the Data field as the file names.Secret volumes support ownership management and SELinux relabeling.",
													MarkdownDescription: "Adapts a Secret into a volume.The contents of the target Secret's Data field will be presented in a volumeas files using the keys in the Data field as the file names.Secret volumes support ownership management and SELinux relabeling.",
													Attributes: map[string]schema.Attribute{
														"default_mode": schema.Int64Attribute{
															Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
															MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"items": schema.ListNestedAttribute{
															Description:         "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
															MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the key to project.",
																		MarkdownDescription: "key is the key to project.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"mode": schema.Int64Attribute{
																		Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"path": schema.StringAttribute{
																		Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																		MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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
															Description:         "optional field specify whether the Secret or its keys must be defined",
															MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"secret_name": schema.StringAttribute{
															Description:         "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
															MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
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

										"wipe_method": schema.StringAttribute{
											Description:         "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
											MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods are wiped for dealing with storage formatchanges.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
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

					"validation_policy": schema.SingleNestedAttribute{
						Description:         "ValidationPolicy controls validation of the Aerospike cluster resource.",
						MarkdownDescription: "ValidationPolicy controls validation of the Aerospike cluster resource.",
						Attributes: map[string]schema.Attribute{
							"skip_work_dir_validate": schema.BoolAttribute{
								Description:         "skipWorkDirValidate validates that Aerospike work directory is mounted on a persistent file storage.Defaults to false.",
								MarkdownDescription: "skipWorkDirValidate validates that Aerospike work directory is mounted on a persistent file storage.Defaults to false.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"skip_xdr_dlog_file_validate": schema.BoolAttribute{
								Description:         "ValidateXdrDigestLogFile validates that xdr digest log file is mounted on a persistent file storage.Defaults to false.",
								MarkdownDescription: "ValidateXdrDigestLogFile validates that xdr digest log file is mounted on a persistent file storage.Defaults to false.",
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
	}
}

func (r *AsdbAerospikeComAerospikeClusterV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_asdb_aerospike_com_aerospike_cluster_v1_manifest")

	var model AsdbAerospikeComAerospikeClusterV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("asdb.aerospike.com/v1")
	model.Kind = pointer.String("AerospikeCluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
