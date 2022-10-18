/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

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

type AsdbAerospikeComAerospikeClusterV1Beta1Resource struct{}

var (
	_ resource.Resource = (*AsdbAerospikeComAerospikeClusterV1Beta1Resource)(nil)
)

type AsdbAerospikeComAerospikeClusterV1Beta1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AsdbAerospikeComAerospikeClusterV1Beta1GoModel struct {
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
		AerospikeAccessControl *struct {
			AdminPolicy *struct {
				Timeout *int64 `tfsdk:"timeout" yaml:"timeout,omitempty"`
			} `tfsdk:"admin_policy" yaml:"adminPolicy,omitempty"`

			Roles *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Privileges *[]string `tfsdk:"privileges" yaml:"privileges,omitempty"`

				ReadQuota *int64 `tfsdk:"read_quota" yaml:"readQuota,omitempty"`

				Whitelist *[]string `tfsdk:"whitelist" yaml:"whitelist,omitempty"`

				WriteQuota *int64 `tfsdk:"write_quota" yaml:"writeQuota,omitempty"`
			} `tfsdk:"roles" yaml:"roles,omitempty"`

			Users *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Roles *[]string `tfsdk:"roles" yaml:"roles,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"users" yaml:"users,omitempty"`
		} `tfsdk:"aerospike_access_control" yaml:"aerospikeAccessControl,omitempty"`

		AerospikeConfig utilities.Dynamic `tfsdk:"aerospike_config" yaml:"aerospikeConfig,omitempty"`

		AerospikeNetworkPolicy *struct {
			Access *string `tfsdk:"access" yaml:"access,omitempty"`

			AlternateAccess *string `tfsdk:"alternate_access" yaml:"alternateAccess,omitempty"`

			TlsAccess *string `tfsdk:"tls_access" yaml:"tlsAccess,omitempty"`

			TlsAlternateAccess *string `tfsdk:"tls_alternate_access" yaml:"tlsAlternateAccess,omitempty"`
		} `tfsdk:"aerospike_network_policy" yaml:"aerospikeNetworkPolicy,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		OperatorClientCert *struct {
			CertPathInOperator *struct {
				CaCertsPath *string `tfsdk:"ca_certs_path" yaml:"caCertsPath,omitempty"`

				ClientCertPath *string `tfsdk:"client_cert_path" yaml:"clientCertPath,omitempty"`

				ClientKeyPath *string `tfsdk:"client_key_path" yaml:"clientKeyPath,omitempty"`
			} `tfsdk:"cert_path_in_operator" yaml:"certPathInOperator,omitempty"`

			SecretCertSource *struct {
				CaCertsFilename *string `tfsdk:"ca_certs_filename" yaml:"caCertsFilename,omitempty"`

				ClientCertFilename *string `tfsdk:"client_cert_filename" yaml:"clientCertFilename,omitempty"`

				ClientKeyFilename *string `tfsdk:"client_key_filename" yaml:"clientKeyFilename,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`

				SecretNamespace *string `tfsdk:"secret_namespace" yaml:"secretNamespace,omitempty"`
			} `tfsdk:"secret_cert_source" yaml:"secretCertSource,omitempty"`

			TlsClientName *string `tfsdk:"tls_client_name" yaml:"tlsClientName,omitempty"`
		} `tfsdk:"operator_client_cert" yaml:"operatorClientCert,omitempty"`

		PodSpec *struct {
			AerospikeContainer *struct {
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

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
			} `tfsdk:"aerospike_container" yaml:"aerospikeContainer,omitempty"`

			AerospikeInitContainer *struct {
				ImageRegistry *string `tfsdk:"image_registry" yaml:"imageRegistry,omitempty"`

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

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`
			} `tfsdk:"aerospike_init_container" yaml:"aerospikeInitContainer,omitempty"`

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

			DnsPolicy *string `tfsdk:"dns_policy" yaml:"dnsPolicy,omitempty"`

			EffectiveDNSPolicy *string `tfsdk:"effective_dns_policy" yaml:"effectiveDNSPolicy,omitempty"`

			HostNetwork *bool `tfsdk:"host_network" yaml:"hostNetwork,omitempty"`

			ImagePullSecrets *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

			InitContainers *[]struct {
				Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

				Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

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

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

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

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"post_start" yaml:"postStart,omitempty"`

					PreStop *struct {
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

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
				} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

				LivenessProbe *struct {
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

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Ports *[]struct {
					ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

					HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

					HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				ReadinessProbe *struct {
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

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

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

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

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

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

				Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

				StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

				TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

				TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

				Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

				VolumeDevices *[]struct {
					DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

				VolumeMounts *[]struct {
					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

					MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

					SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

					SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
				} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

				WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
			} `tfsdk:"init_containers" yaml:"initContainers,omitempty"`

			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
			} `tfsdk:"metadata" yaml:"metadata,omitempty"`

			MultiPodPerHost *bool `tfsdk:"multi_pod_per_host" yaml:"multiPodPerHost,omitempty"`

			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

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

					RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
			} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

			Sidecars *[]struct {
				Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

				Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

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

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

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

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"post_start" yaml:"postStart,omitempty"`

					PreStop *struct {
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

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

							Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
						} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

						TcpSocket *struct {
							Host *string `tfsdk:"host" yaml:"host,omitempty"`

							Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
						} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
					} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
				} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

				LivenessProbe *struct {
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

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Ports *[]struct {
					ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

					HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

					HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"ports" yaml:"ports,omitempty"`

				ReadinessProbe *struct {
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

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

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

						RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
					} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
				} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

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

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

					PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

					SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

					TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

					TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
				} `tfsdk:"startup_probe" yaml:"startupProbe,omitempty"`

				Stdin *bool `tfsdk:"stdin" yaml:"stdin,omitempty"`

				StdinOnce *bool `tfsdk:"stdin_once" yaml:"stdinOnce,omitempty"`

				TerminationMessagePath *string `tfsdk:"termination_message_path" yaml:"terminationMessagePath,omitempty"`

				TerminationMessagePolicy *string `tfsdk:"termination_message_policy" yaml:"terminationMessagePolicy,omitempty"`

				Tty *bool `tfsdk:"tty" yaml:"tty,omitempty"`

				VolumeDevices *[]struct {
					DevicePath *string `tfsdk:"device_path" yaml:"devicePath,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"volume_devices" yaml:"volumeDevices,omitempty"`

				VolumeMounts *[]struct {
					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

					MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

					SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

					SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
				} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`

				WorkingDir *string `tfsdk:"working_dir" yaml:"workingDir,omitempty"`
			} `tfsdk:"sidecars" yaml:"sidecars,omitempty"`

			Tolerations *[]struct {
				Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
		} `tfsdk:"pod_spec" yaml:"podSpec,omitempty"`

		RackConfig *struct {
			Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

			Racks *[]struct {
				AerospikeConfig utilities.Dynamic `tfsdk:"aerospike_config" yaml:"aerospikeConfig,omitempty"`

				EffectiveAerospikeConfig utilities.Dynamic `tfsdk:"effective_aerospike_config" yaml:"effectiveAerospikeConfig,omitempty"`

				EffectivePodSpec *struct {
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

					NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

					Tolerations *[]struct {
						Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
				} `tfsdk:"effective_pod_spec" yaml:"effectivePodSpec,omitempty"`

				EffectiveStorage *struct {
					BlockVolumePolicy *struct {
						CascadeDelete *bool `tfsdk:"cascade_delete" yaml:"cascadeDelete,omitempty"`

						EffectiveCascadeDelete *bool `tfsdk:"effective_cascade_delete" yaml:"effectiveCascadeDelete,omitempty"`

						EffectiveInitMethod *string `tfsdk:"effective_init_method" yaml:"effectiveInitMethod,omitempty"`

						EffectiveWipeMethod *string `tfsdk:"effective_wipe_method" yaml:"effectiveWipeMethod,omitempty"`

						InitMethod *string `tfsdk:"init_method" yaml:"initMethod,omitempty"`

						WipeMethod *string `tfsdk:"wipe_method" yaml:"wipeMethod,omitempty"`
					} `tfsdk:"block_volume_policy" yaml:"blockVolumePolicy,omitempty"`

					FilesystemVolumePolicy *struct {
						CascadeDelete *bool `tfsdk:"cascade_delete" yaml:"cascadeDelete,omitempty"`

						EffectiveCascadeDelete *bool `tfsdk:"effective_cascade_delete" yaml:"effectiveCascadeDelete,omitempty"`

						EffectiveInitMethod *string `tfsdk:"effective_init_method" yaml:"effectiveInitMethod,omitempty"`

						EffectiveWipeMethod *string `tfsdk:"effective_wipe_method" yaml:"effectiveWipeMethod,omitempty"`

						InitMethod *string `tfsdk:"init_method" yaml:"initMethod,omitempty"`

						WipeMethod *string `tfsdk:"wipe_method" yaml:"wipeMethod,omitempty"`
					} `tfsdk:"filesystem_volume_policy" yaml:"filesystemVolumePolicy,omitempty"`

					Volumes *[]struct {
						Aerospike *struct {
							MountOptions *struct {
								MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

								ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

								SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

								SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"aerospike" yaml:"aerospike,omitempty"`

						CascadeDelete *bool `tfsdk:"cascade_delete" yaml:"cascadeDelete,omitempty"`

						EffectiveCascadeDelete *bool `tfsdk:"effective_cascade_delete" yaml:"effectiveCascadeDelete,omitempty"`

						EffectiveInitMethod *string `tfsdk:"effective_init_method" yaml:"effectiveInitMethod,omitempty"`

						EffectiveWipeMethod *string `tfsdk:"effective_wipe_method" yaml:"effectiveWipeMethod,omitempty"`

						InitContainers *[]struct {
							ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

							MountOptions *struct {
								MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

								ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

								SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

								SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"init_containers" yaml:"initContainers,omitempty"`

						InitMethod *string `tfsdk:"init_method" yaml:"initMethod,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Sidecars *[]struct {
							ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

							MountOptions *struct {
								MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

								ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

								SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

								SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"sidecars" yaml:"sidecars,omitempty"`

						Source *struct {
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

							EmptyDir *struct {
								Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

								SizeLimit utilities.IntOrString `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
							} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

							PersistentVolume *struct {
								AccessModes *[]string `tfsdk:"access_modes" yaml:"accessModes,omitempty"`

								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

									Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
								} `tfsdk:"metadata" yaml:"metadata,omitempty"`

								Selector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"selector" yaml:"selector,omitempty"`

								Size utilities.IntOrString `tfsdk:"size" yaml:"size,omitempty"`

								StorageClass *string `tfsdk:"storage_class" yaml:"storageClass,omitempty"`

								VolumeMode *string `tfsdk:"volume_mode" yaml:"volumeMode,omitempty"`
							} `tfsdk:"persistent_volume" yaml:"persistentVolume,omitempty"`

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
						} `tfsdk:"source" yaml:"source,omitempty"`

						WipeMethod *string `tfsdk:"wipe_method" yaml:"wipeMethod,omitempty"`
					} `tfsdk:"volumes" yaml:"volumes,omitempty"`
				} `tfsdk:"effective_storage" yaml:"effectiveStorage,omitempty"`

				Id *int64 `tfsdk:"id" yaml:"id,omitempty"`

				NodeName *string `tfsdk:"node_name" yaml:"nodeName,omitempty"`

				PodSpec *struct {
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

					NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

					Tolerations *[]struct {
						Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
				} `tfsdk:"pod_spec" yaml:"podSpec,omitempty"`

				RackLabel *string `tfsdk:"rack_label" yaml:"rackLabel,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`

				Storage *struct {
					BlockVolumePolicy *struct {
						CascadeDelete *bool `tfsdk:"cascade_delete" yaml:"cascadeDelete,omitempty"`

						EffectiveCascadeDelete *bool `tfsdk:"effective_cascade_delete" yaml:"effectiveCascadeDelete,omitempty"`

						EffectiveInitMethod *string `tfsdk:"effective_init_method" yaml:"effectiveInitMethod,omitempty"`

						EffectiveWipeMethod *string `tfsdk:"effective_wipe_method" yaml:"effectiveWipeMethod,omitempty"`

						InitMethod *string `tfsdk:"init_method" yaml:"initMethod,omitempty"`

						WipeMethod *string `tfsdk:"wipe_method" yaml:"wipeMethod,omitempty"`
					} `tfsdk:"block_volume_policy" yaml:"blockVolumePolicy,omitempty"`

					FilesystemVolumePolicy *struct {
						CascadeDelete *bool `tfsdk:"cascade_delete" yaml:"cascadeDelete,omitempty"`

						EffectiveCascadeDelete *bool `tfsdk:"effective_cascade_delete" yaml:"effectiveCascadeDelete,omitempty"`

						EffectiveInitMethod *string `tfsdk:"effective_init_method" yaml:"effectiveInitMethod,omitempty"`

						EffectiveWipeMethod *string `tfsdk:"effective_wipe_method" yaml:"effectiveWipeMethod,omitempty"`

						InitMethod *string `tfsdk:"init_method" yaml:"initMethod,omitempty"`

						WipeMethod *string `tfsdk:"wipe_method" yaml:"wipeMethod,omitempty"`
					} `tfsdk:"filesystem_volume_policy" yaml:"filesystemVolumePolicy,omitempty"`

					Volumes *[]struct {
						Aerospike *struct {
							MountOptions *struct {
								MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

								ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

								SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

								SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"aerospike" yaml:"aerospike,omitempty"`

						CascadeDelete *bool `tfsdk:"cascade_delete" yaml:"cascadeDelete,omitempty"`

						EffectiveCascadeDelete *bool `tfsdk:"effective_cascade_delete" yaml:"effectiveCascadeDelete,omitempty"`

						EffectiveInitMethod *string `tfsdk:"effective_init_method" yaml:"effectiveInitMethod,omitempty"`

						EffectiveWipeMethod *string `tfsdk:"effective_wipe_method" yaml:"effectiveWipeMethod,omitempty"`

						InitContainers *[]struct {
							ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

							MountOptions *struct {
								MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

								ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

								SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

								SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"init_containers" yaml:"initContainers,omitempty"`

						InitMethod *string `tfsdk:"init_method" yaml:"initMethod,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Sidecars *[]struct {
							ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

							MountOptions *struct {
								MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

								ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

								SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

								SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
							} `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

							Path *string `tfsdk:"path" yaml:"path,omitempty"`
						} `tfsdk:"sidecars" yaml:"sidecars,omitempty"`

						Source *struct {
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

							EmptyDir *struct {
								Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

								SizeLimit utilities.IntOrString `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
							} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

							PersistentVolume *struct {
								AccessModes *[]string `tfsdk:"access_modes" yaml:"accessModes,omitempty"`

								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

									Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
								} `tfsdk:"metadata" yaml:"metadata,omitempty"`

								Selector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"selector" yaml:"selector,omitempty"`

								Size utilities.IntOrString `tfsdk:"size" yaml:"size,omitempty"`

								StorageClass *string `tfsdk:"storage_class" yaml:"storageClass,omitempty"`

								VolumeMode *string `tfsdk:"volume_mode" yaml:"volumeMode,omitempty"`
							} `tfsdk:"persistent_volume" yaml:"persistentVolume,omitempty"`

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
						} `tfsdk:"source" yaml:"source,omitempty"`

						WipeMethod *string `tfsdk:"wipe_method" yaml:"wipeMethod,omitempty"`
					} `tfsdk:"volumes" yaml:"volumes,omitempty"`
				} `tfsdk:"storage" yaml:"storage,omitempty"`

				Zone *string `tfsdk:"zone" yaml:"zone,omitempty"`
			} `tfsdk:"racks" yaml:"racks,omitempty"`
		} `tfsdk:"rack_config" yaml:"rackConfig,omitempty"`

		SeedsFinderServices *struct {
			LoadBalancer *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				ExternalTrafficPolicy *string `tfsdk:"external_traffic_policy" yaml:"externalTrafficPolicy,omitempty"`

				LoadBalancerSourceRanges *[]string `tfsdk:"load_balancer_source_ranges" yaml:"loadBalancerSourceRanges,omitempty"`

				Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

				TargetPort *int64 `tfsdk:"target_port" yaml:"targetPort,omitempty"`
			} `tfsdk:"load_balancer" yaml:"loadBalancer,omitempty"`
		} `tfsdk:"seeds_finder_services" yaml:"seedsFinderServices,omitempty"`

		Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

		Storage *struct {
			BlockVolumePolicy *struct {
				CascadeDelete *bool `tfsdk:"cascade_delete" yaml:"cascadeDelete,omitempty"`

				EffectiveCascadeDelete *bool `tfsdk:"effective_cascade_delete" yaml:"effectiveCascadeDelete,omitempty"`

				EffectiveInitMethod *string `tfsdk:"effective_init_method" yaml:"effectiveInitMethod,omitempty"`

				EffectiveWipeMethod *string `tfsdk:"effective_wipe_method" yaml:"effectiveWipeMethod,omitempty"`

				InitMethod *string `tfsdk:"init_method" yaml:"initMethod,omitempty"`

				WipeMethod *string `tfsdk:"wipe_method" yaml:"wipeMethod,omitempty"`
			} `tfsdk:"block_volume_policy" yaml:"blockVolumePolicy,omitempty"`

			FilesystemVolumePolicy *struct {
				CascadeDelete *bool `tfsdk:"cascade_delete" yaml:"cascadeDelete,omitempty"`

				EffectiveCascadeDelete *bool `tfsdk:"effective_cascade_delete" yaml:"effectiveCascadeDelete,omitempty"`

				EffectiveInitMethod *string `tfsdk:"effective_init_method" yaml:"effectiveInitMethod,omitempty"`

				EffectiveWipeMethod *string `tfsdk:"effective_wipe_method" yaml:"effectiveWipeMethod,omitempty"`

				InitMethod *string `tfsdk:"init_method" yaml:"initMethod,omitempty"`

				WipeMethod *string `tfsdk:"wipe_method" yaml:"wipeMethod,omitempty"`
			} `tfsdk:"filesystem_volume_policy" yaml:"filesystemVolumePolicy,omitempty"`

			Volumes *[]struct {
				Aerospike *struct {
					MountOptions *struct {
						MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

						SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
					} `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"aerospike" yaml:"aerospike,omitempty"`

				CascadeDelete *bool `tfsdk:"cascade_delete" yaml:"cascadeDelete,omitempty"`

				EffectiveCascadeDelete *bool `tfsdk:"effective_cascade_delete" yaml:"effectiveCascadeDelete,omitempty"`

				EffectiveInitMethod *string `tfsdk:"effective_init_method" yaml:"effectiveInitMethod,omitempty"`

				EffectiveWipeMethod *string `tfsdk:"effective_wipe_method" yaml:"effectiveWipeMethod,omitempty"`

				InitContainers *[]struct {
					ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

					MountOptions *struct {
						MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

						SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
					} `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"init_containers" yaml:"initContainers,omitempty"`

				InitMethod *string `tfsdk:"init_method" yaml:"initMethod,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Sidecars *[]struct {
					ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

					MountOptions *struct {
						MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

						ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

						SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

						SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
					} `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`
				} `tfsdk:"sidecars" yaml:"sidecars,omitempty"`

				Source *struct {
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

					EmptyDir *struct {
						Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

						SizeLimit utilities.IntOrString `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
					} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

					PersistentVolume *struct {
						AccessModes *[]string `tfsdk:"access_modes" yaml:"accessModes,omitempty"`

						Metadata *struct {
							Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

							Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`
						} `tfsdk:"metadata" yaml:"metadata,omitempty"`

						Selector *struct {
							MatchExpressions *[]struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

								Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
							} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

							MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
						} `tfsdk:"selector" yaml:"selector,omitempty"`

						Size utilities.IntOrString `tfsdk:"size" yaml:"size,omitempty"`

						StorageClass *string `tfsdk:"storage_class" yaml:"storageClass,omitempty"`

						VolumeMode *string `tfsdk:"volume_mode" yaml:"volumeMode,omitempty"`
					} `tfsdk:"persistent_volume" yaml:"persistentVolume,omitempty"`

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
				} `tfsdk:"source" yaml:"source,omitempty"`

				WipeMethod *string `tfsdk:"wipe_method" yaml:"wipeMethod,omitempty"`
			} `tfsdk:"volumes" yaml:"volumes,omitempty"`
		} `tfsdk:"storage" yaml:"storage,omitempty"`

		ValidationPolicy *struct {
			SkipWorkDirValidate *bool `tfsdk:"skip_work_dir_validate" yaml:"skipWorkDirValidate,omitempty"`

			SkipXdrDlogFileValidate *bool `tfsdk:"skip_xdr_dlog_file_validate" yaml:"skipXdrDlogFileValidate,omitempty"`
		} `tfsdk:"validation_policy" yaml:"validationPolicy,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAsdbAerospikeComAerospikeClusterV1Beta1Resource() resource.Resource {
	return &AsdbAerospikeComAerospikeClusterV1Beta1Resource{}
}

func (r *AsdbAerospikeComAerospikeClusterV1Beta1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asdb_aerospike_com_aerospike_cluster_v1beta1"
}

func (r *AsdbAerospikeComAerospikeClusterV1Beta1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "AerospikeCluster is the schema for the AerospikeCluster API",
		MarkdownDescription: "AerospikeCluster is the schema for the AerospikeCluster API",
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
				Description:         "AerospikeClusterSpec defines the desired state of AerospikeCluster",
				MarkdownDescription: "AerospikeClusterSpec defines the desired state of AerospikeCluster",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"aerospike_access_control": {
						Description:         "Has the Aerospike roles and users definitions.",
						MarkdownDescription: "Has the Aerospike roles and users definitions.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"admin_policy": {
								Description:         "AerospikeClientAdminPolicy specify the aerospike client admin policy f",
								MarkdownDescription: "AerospikeClientAdminPolicy specify the aerospike client admin policy f",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"timeout": {
										Description:         "Timeout for admin client policy in milliseconds.",
										MarkdownDescription: "Timeout for admin client policy in milliseconds.",

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

							"roles": {
								Description:         "Roles is the set of roles to allow on the Aerospike cluster.",
								MarkdownDescription: "Roles is the set of roles to allow on the Aerospike cluster.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of this role.",
										MarkdownDescription: "Name of this role.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"privileges": {
										Description:         "Privileges granted to this role.",
										MarkdownDescription: "Privileges granted to this role.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"read_quota": {
										Description:         "ReadQuota specifies permitted rate of read records for current role (t",
										MarkdownDescription: "ReadQuota specifies permitted rate of read records for current role (t",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"whitelist": {
										Description:         "Whitelist of host address allowed for this role.",
										MarkdownDescription: "Whitelist of host address allowed for this role.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"write_quota": {
										Description:         "WriteQuota specifies permitted rate of write records for current role ",
										MarkdownDescription: "WriteQuota specifies permitted rate of write records for current role ",

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

							"users": {
								Description:         "Users is the set of users to allow on the Aerospike cluster.",
								MarkdownDescription: "Users is the set of users to allow on the Aerospike cluster.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name is the user's username.",
										MarkdownDescription: "Name is the user's username.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"roles": {
										Description:         "Roles is the list of roles granted to the user.",
										MarkdownDescription: "Roles is the list of roles granted to the user.",

										Type: types.ListType{ElemType: types.StringType},

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_name": {
										Description:         "SecretName has secret info created by user.",
										MarkdownDescription: "SecretName has secret info created by user.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"aerospike_config": {
						Description:         "Sets config in aerospike.conf file. Other configs are taken as default",
						MarkdownDescription: "Sets config in aerospike.conf file. Other configs are taken as default",

						Type: utilities.DynamicType{},

						Required: true,
						Optional: false,
						Computed: false,
					},

					"aerospike_network_policy": {
						Description:         "AerospikeNetworkPolicy specifies how clients and tools access the Aero",
						MarkdownDescription: "AerospikeNetworkPolicy specifies how clients and tools access the Aero",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"access": {
								Description:         "AccessType is the type of network address to use for Aerospike access ",
								MarkdownDescription: "AccessType is the type of network address to use for Aerospike access ",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("pod", "hostInternal", "hostExternal"),
								},
							},

							"alternate_access": {
								Description:         "AlternateAccessType is the type of network address to use for Aerospik",
								MarkdownDescription: "AlternateAccessType is the type of network address to use for Aerospik",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("pod", "hostInternal", "hostExternal"),
								},
							},

							"tls_access": {
								Description:         "TLSAccessType is the type of network address to use for Aerospike TLS ",
								MarkdownDescription: "TLSAccessType is the type of network address to use for Aerospike TLS ",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("pod", "hostInternal", "hostExternal"),
								},
							},

							"tls_alternate_access": {
								Description:         "TLSAlternateAccessType is the type of network address to use for Aeros",
								MarkdownDescription: "TLSAlternateAccessType is the type of network address to use for Aeros",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("pod", "hostInternal", "hostExternal"),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "Aerospike server image",
						MarkdownDescription: "Aerospike server image",

						Type: types.StringType,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"operator_client_cert": {
						Description:         "Certificates to connect to Aerospike.",
						MarkdownDescription: "Certificates to connect to Aerospike.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cert_path_in_operator": {
								Description:         "AerospikeCertPathInOperatorSource contain configuration for certificat",
								MarkdownDescription: "AerospikeCertPathInOperatorSource contain configuration for certificat",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_certs_path": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_cert_path": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_key_path": {
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

							"secret_cert_source": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca_certs_filename": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_cert_filename": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_key_filename": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"secret_namespace": {
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

							"tls_client_name": {
								Description:         "If specified, this name will be added to tls-authenticate-client list ",
								MarkdownDescription: "If specified, this name will be added to tls-authenticate-client list ",

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

					"pod_spec": {
						Description:         "Specify additional configuration for the Aerospike pods",
						MarkdownDescription: "Specify additional configuration for the Aerospike pods",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"aerospike_container": {
								Description:         "AerospikeContainerSpec configures the aerospike-server container creat",
								MarkdownDescription: "AerospikeContainerSpec configures the aerospike-server container creat",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"resources": {
										Description:         "Define resources requests and limits for Aerospike Server Container.",
										MarkdownDescription: "Define resources requests and limits for Aerospike Server Container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed.",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required.",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required.",

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
										Description:         "SecurityContext that will be added to aerospike-server container creat",
										MarkdownDescription: "SecurityContext that will be added to aerospike-server container creat",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_privilege_escalation": {
												Description:         "AllowPrivilegeEscalation controls whether a process can gain more priv",
												MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more priv",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"capabilities": {
												Description:         "The capabilities to add/drop when running containers.",
												MarkdownDescription: "The capabilities to add/drop when running containers.",

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
												Description:         "Run container in privileged mode.",
												MarkdownDescription: "Run container in privileged mode.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proc_mount": {
												Description:         "procMount denotes the type of proc mount to use for the containers.",
												MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only_root_filesystem": {
												Description:         "Whether this container has a read-only root filesystem.",
												MarkdownDescription: "Whether this container has a read-only root filesystem.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_group": {
												Description:         "The GID to run the entrypoint of the container process.",
												MarkdownDescription: "The GID to run the entrypoint of the container process.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_non_root": {
												Description:         "Indicates that the container must run as a non-root user.",
												MarkdownDescription: "Indicates that the container must run as a non-root user.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_user": {
												Description:         "The UID to run the entrypoint of the container process.",
												MarkdownDescription: "The UID to run the entrypoint of the container process.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"se_linux_options": {
												Description:         "The SELinux context to be applied to the container.",
												MarkdownDescription: "The SELinux context to be applied to the container.",

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
												Description:         "The seccomp options to use by this container.",
												MarkdownDescription: "The seccomp options to use by this container.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"localhost_profile": {
														Description:         "localhostProfile indicates a profile defined in a file on the node sho",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node sho",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "type indicates which kind of seccomp profile will be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied.",

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
												Description:         "The Windows specific settings applied to all containers.",
												MarkdownDescription: "The Windows specific settings applied to all containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"gmsa_credential_spec": {
														Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github",
														MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github",

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
														Description:         "The UserName in Windows to run the entrypoint of the container process",
														MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process",

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

							"aerospike_init_container": {
								Description:         "AerospikeInitContainerSpec configures the aerospike-init container cre",
								MarkdownDescription: "AerospikeInitContainerSpec configures the aerospike-init container cre",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image_registry": {
										Description:         "ImageRegistry is the name of image registry for aerospike-init contain",
										MarkdownDescription: "ImageRegistry is the name of image registry for aerospike-init contain",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "Define resources requests and limits for Aerospike init Container.",
										MarkdownDescription: "Define resources requests and limits for Aerospike init Container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed.",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required.",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required.",

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
										Description:         "SecurityContext that will be added to aerospike-init container created",
										MarkdownDescription: "SecurityContext that will be added to aerospike-init container created",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_privilege_escalation": {
												Description:         "AllowPrivilegeEscalation controls whether a process can gain more priv",
												MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more priv",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"capabilities": {
												Description:         "The capabilities to add/drop when running containers.",
												MarkdownDescription: "The capabilities to add/drop when running containers.",

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
												Description:         "Run container in privileged mode.",
												MarkdownDescription: "Run container in privileged mode.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proc_mount": {
												Description:         "procMount denotes the type of proc mount to use for the containers.",
												MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only_root_filesystem": {
												Description:         "Whether this container has a read-only root filesystem.",
												MarkdownDescription: "Whether this container has a read-only root filesystem.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_group": {
												Description:         "The GID to run the entrypoint of the container process.",
												MarkdownDescription: "The GID to run the entrypoint of the container process.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_non_root": {
												Description:         "Indicates that the container must run as a non-root user.",
												MarkdownDescription: "Indicates that the container must run as a non-root user.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_user": {
												Description:         "The UID to run the entrypoint of the container process.",
												MarkdownDescription: "The UID to run the entrypoint of the container process.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"se_linux_options": {
												Description:         "The SELinux context to be applied to the container.",
												MarkdownDescription: "The SELinux context to be applied to the container.",

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
												Description:         "The seccomp options to use by this container.",
												MarkdownDescription: "The seccomp options to use by this container.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"localhost_profile": {
														Description:         "localhostProfile indicates a profile defined in a file on the node sho",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node sho",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "type indicates which kind of seccomp profile will be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied.",

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
												Description:         "The Windows specific settings applied to all containers.",
												MarkdownDescription: "The Windows specific settings applied to all containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"gmsa_credential_spec": {
														Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github",
														MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github",

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
														Description:         "The UserName in Windows to run the entrypoint of the container process",
														MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process",

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

							"affinity": {
								Description:         "Affinity rules for pod placement.",
								MarkdownDescription: "Affinity rules for pod placement.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_affinity": {
										Description:         "Describes node affinity scheduling rules for the pod.",
										MarkdownDescription: "Describes node affinity scheduling rules for the pod.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"preferred_during_scheduling_ignored_during_execution": {
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the a",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the a",

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
																		Description:         "Represents a key's relationship to a set of values.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "An array of string values.",
																		MarkdownDescription: "An array of string values.",

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
																		Description:         "Represents a key's relationship to a set of values.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "An array of string values.",
																		MarkdownDescription: "An array of string values.",

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
														Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in",
														MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in",

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
												Description:         "If the affinity requirements specified by this field are not met at sc",
												MarkdownDescription: "If the affinity requirements specified by this field are not met at sc",

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
																		Description:         "Represents a key's relationship to a set of values.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "An array of string values.",
																		MarkdownDescription: "An array of string values.",

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
																		Description:         "Represents a key's relationship to a set of values.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "An array of string values.",
																		MarkdownDescription: "An array of string values.",

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
										Description:         "Describes pod affinity scheduling rules (e.g.",
										MarkdownDescription: "Describes pod affinity scheduling rules (e.g.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"preferred_during_scheduling_ignored_during_execution": {
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the a",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the a",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"pod_affinity_term": {
														Description:         "Required.",
														MarkdownDescription: "Required.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"label_selector": {
																Description:         "A label query over a set of resources, in this case pods.",
																MarkdownDescription: "A label query over a set of resources, in this case pods.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"match_expressions": {
																		Description:         "matchExpressions is a list of label selector requirements.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																				Description:         "operator represents a key's relationship to a set of values.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "values is an array of string values.",
																				MarkdownDescription: "values is an array of string values.",

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
																		Description:         "matchLabels is a map of {key,value} pairs.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																Description:         "A label query over the set of namespaces that the term applies to.",
																MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"match_expressions": {
																		Description:         "matchExpressions is a list of label selector requirements.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																				Description:         "operator represents a key's relationship to a set of values.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "values is an array of string values.",
																				MarkdownDescription: "values is an array of string values.",

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
																		Description:         "matchLabels is a map of {key,value} pairs.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																Description:         "namespaces specifies a static list of namespace names that the term ap",
																MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"topology_key": {
																Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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
														Description:         "weight associated with matching the corresponding podAffinityTerm, in ",
														MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in ",

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
												Description:         "If the affinity requirements specified by this field are not met at sc",
												MarkdownDescription: "If the affinity requirements specified by this field are not met at sc",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"label_selector": {
														Description:         "A label query over a set of resources, in this case pods.",
														MarkdownDescription: "A label query over a set of resources, in this case pods.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																		Description:         "operator represents a key's relationship to a set of values.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values.",
																		MarkdownDescription: "values is an array of string values.",

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
																Description:         "matchLabels is a map of {key,value} pairs.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
														Description:         "A label query over the set of namespaces that the term applies to.",
														MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																		Description:         "operator represents a key's relationship to a set of values.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values.",
																		MarkdownDescription: "values is an array of string values.",

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
																Description:         "matchLabels is a map of {key,value} pairs.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
														Description:         "namespaces specifies a static list of namespace names that the term ap",
														MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
														MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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
										Description:         "Describes pod anti-affinity scheduling rules (e.g.",
										MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"preferred_during_scheduling_ignored_during_execution": {
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the a",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the a",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"pod_affinity_term": {
														Description:         "Required.",
														MarkdownDescription: "Required.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"label_selector": {
																Description:         "A label query over a set of resources, in this case pods.",
																MarkdownDescription: "A label query over a set of resources, in this case pods.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"match_expressions": {
																		Description:         "matchExpressions is a list of label selector requirements.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																				Description:         "operator represents a key's relationship to a set of values.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "values is an array of string values.",
																				MarkdownDescription: "values is an array of string values.",

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
																		Description:         "matchLabels is a map of {key,value} pairs.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																Description:         "A label query over the set of namespaces that the term applies to.",
																MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"match_expressions": {
																		Description:         "matchExpressions is a list of label selector requirements.",
																		MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																				Description:         "operator represents a key's relationship to a set of values.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "values is an array of string values.",
																				MarkdownDescription: "values is an array of string values.",

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
																		Description:         "matchLabels is a map of {key,value} pairs.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																Description:         "namespaces specifies a static list of namespace names that the term ap",
																MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"topology_key": {
																Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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
														Description:         "weight associated with matching the corresponding podAffinityTerm, in ",
														MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in ",

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
												Description:         "If the anti-affinity requirements specified by this field are not met ",
												MarkdownDescription: "If the anti-affinity requirements specified by this field are not met ",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"label_selector": {
														Description:         "A label query over a set of resources, in this case pods.",
														MarkdownDescription: "A label query over a set of resources, in this case pods.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																		Description:         "operator represents a key's relationship to a set of values.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values.",
																		MarkdownDescription: "values is an array of string values.",

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
																Description:         "matchLabels is a map of {key,value} pairs.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
														Description:         "A label query over the set of namespaces that the term applies to.",
														MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																		Description:         "operator represents a key's relationship to a set of values.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values.",
																		MarkdownDescription: "values is an array of string values.",

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
																Description:         "matchLabels is a map of {key,value} pairs.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
														Description:         "namespaces specifies a static list of namespace names that the term ap",
														MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"topology_key": {
														Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
														MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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

							"dns_policy": {
								Description:         "DnsPolicy same as https://kubernetes.",
								MarkdownDescription: "DnsPolicy same as https://kubernetes.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"effective_dns_policy": {
								Description:         "Effective value of the DNSPolicy",
								MarkdownDescription: "Effective value of the DNSPolicy",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_network": {
								Description:         "HostNetwork enables host networking for the pod.",
								MarkdownDescription: "HostNetwork enables host networking for the pod.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_secrets": {
								Description:         "ImagePullSecrets is an optional list of references to secrets in the s",
								MarkdownDescription: "ImagePullSecrets is an optional list of references to secrets in the s",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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

							"init_containers": {
								Description:         "InitContainers to add to the pods.",
								MarkdownDescription: "InitContainers to add to the pods.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"args": {
										Description:         "Arguments to the entrypoint.",
										MarkdownDescription: "Arguments to the entrypoint.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"command": {
										Description:         "Entrypoint array. Not executed within a shell.",
										MarkdownDescription: "Entrypoint array. Not executed within a shell.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"env": {
										Description:         "List of environment variables to set in the container.",
										MarkdownDescription: "List of environment variables to set in the container.",

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
												Description:         "Variable references $(VAR_NAME) are expanded using the previous define",
												MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous define",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "Source for the environment variable's value.",
												MarkdownDescription: "Source for the environment variable's value.",

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
																Description:         "Name of the referent. More info: https://kubernetes.",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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
														Description:         "Selects a field of the pod: supports metadata.name, metadata.",
														MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api_version": {
																Description:         "Version of the schema the FieldPath is written in terms of, defaults t",
																MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults t",

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
														Description:         "Selects a resource of the container: only resources limits and request",
														MarkdownDescription: "Selects a resource of the container: only resources limits and request",

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
																Description:         "Name of the referent. More info: https://kubernetes.",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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
										Description:         "List of sources to populate environment variables in the container.",
										MarkdownDescription: "List of sources to populate environment variables in the container.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_ref": {
												Description:         "The ConfigMap to select from",
												MarkdownDescription: "The ConfigMap to select from",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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
												Description:         "An optional identifier to prepend to each key in the ConfigMap.",
												MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap.",

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
														Description:         "Name of the referent. More info: https://kubernetes.",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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

									"image": {
										Description:         "Docker image name. More info: https://kubernetes.",
										MarkdownDescription: "Docker image name. More info: https://kubernetes.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_pull_policy": {
										Description:         "Image pull policy. One of Always, Never, IfNotPresent.",
										MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"lifecycle": {
										Description:         "Actions that the management system should take in response to containe",
										MarkdownDescription: "Actions that the management system should take in response to containe",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"post_start": {
												Description:         "PostStart is called immediately after a container is created.",
												MarkdownDescription: "PostStart is called immediately after a container is created.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "One and only one of the following should be specified.",
														MarkdownDescription: "One and only one of the following should be specified.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the worki",
																MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
																Description:         "Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
																Description:         "Name or number of the port to access on the container.",
																MarkdownDescription: "Name or number of the port to access on the container.",

																Type: utilities.IntOrStringType{},

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
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
																Description:         "Number or name of the port to access on the container.",
																MarkdownDescription: "Number or name of the port to access on the container.",

																Type: utilities.IntOrStringType{},

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
												Description:         "PreStop is called immediately before a container is terminated due to ",
												MarkdownDescription: "PreStop is called immediately before a container is terminated due to ",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "One and only one of the following should be specified.",
														MarkdownDescription: "One and only one of the following should be specified.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the worki",
																MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
																Description:         "Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
																Description:         "Name or number of the port to access on the container.",
																MarkdownDescription: "Name or number of the port to access on the container.",

																Type: utilities.IntOrStringType{},

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
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
																Description:         "Number or name of the port to access on the container.",
																MarkdownDescription: "Number or name of the port to access on the container.",

																Type: utilities.IntOrStringType{},

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
										Description:         "Periodic probe of container liveness.",
										MarkdownDescription: "Periodic probe of container liveness.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "One and only one of the following should be specified.",
												MarkdownDescription: "One and only one of the following should be specified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the worki",
														MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
												Description:         "Minimum consecutive failures for the probe to be considered failed aft",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed aft",

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
														Description:         "Host name to connect to, defaults to the pod IP.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
														Description:         "Name or number of the port to access on the container.",
														MarkdownDescription: "Name or number of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Number of seconds after the container has started before liveness prob",
												MarkdownDescription: "Number of seconds after the container has started before liveness prob",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successfu",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successfu",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tcp_socket": {
												Description:         "TCPSocket specifies an action involving a TCP port.",
												MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
														Description:         "Number or name of the port to access on the container.",
														MarkdownDescription: "Number or name of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Optional duration in seconds the pod needs to terminate gracefully upo",
												MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upo",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"timeout_seconds": {
												Description:         "Number of seconds after which the probe times out.",
												MarkdownDescription: "Number of seconds after which the probe times out.",

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
										Description:         "Name of the container specified as a DNS_LABEL.",
										MarkdownDescription: "Name of the container specified as a DNS_LABEL.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"ports": {
										Description:         "List of ports to expose from the container.",
										MarkdownDescription: "List of ports to expose from the container.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"container_port": {
												Description:         "Number of port to expose on the pod's IP address.",
												MarkdownDescription: "Number of port to expose on the pod's IP address.",

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
												Description:         "Number of port to expose on the host.",
												MarkdownDescription: "Number of port to expose on the host.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod.",
												MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_probe": {
										Description:         "Periodic probe of container service readiness.",
										MarkdownDescription: "Periodic probe of container service readiness.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "One and only one of the following should be specified.",
												MarkdownDescription: "One and only one of the following should be specified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the worki",
														MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
												Description:         "Minimum consecutive failures for the probe to be considered failed aft",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed aft",

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
														Description:         "Host name to connect to, defaults to the pod IP.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
														Description:         "Name or number of the port to access on the container.",
														MarkdownDescription: "Name or number of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Number of seconds after the container has started before liveness prob",
												MarkdownDescription: "Number of seconds after the container has started before liveness prob",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successfu",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successfu",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tcp_socket": {
												Description:         "TCPSocket specifies an action involving a TCP port.",
												MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
														Description:         "Number or name of the port to access on the container.",
														MarkdownDescription: "Number or name of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Optional duration in seconds the pod needs to terminate gracefully upo",
												MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upo",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"timeout_seconds": {
												Description:         "Number of seconds after which the probe times out.",
												MarkdownDescription: "Number of seconds after which the probe times out.",

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

									"resources": {
										Description:         "Compute Resources required by this container. Cannot be updated.",
										MarkdownDescription: "Compute Resources required by this container. Cannot be updated.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed.",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required.",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required.",

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
										Description:         "Security options the pod should run with.",
										MarkdownDescription: "Security options the pod should run with.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_privilege_escalation": {
												Description:         "AllowPrivilegeEscalation controls whether a process can gain more priv",
												MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more priv",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"capabilities": {
												Description:         "The capabilities to add/drop when running containers.",
												MarkdownDescription: "The capabilities to add/drop when running containers.",

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
												Description:         "Run container in privileged mode.",
												MarkdownDescription: "Run container in privileged mode.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proc_mount": {
												Description:         "procMount denotes the type of proc mount to use for the containers.",
												MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only_root_filesystem": {
												Description:         "Whether this container has a read-only root filesystem.",
												MarkdownDescription: "Whether this container has a read-only root filesystem.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_group": {
												Description:         "The GID to run the entrypoint of the container process.",
												MarkdownDescription: "The GID to run the entrypoint of the container process.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_non_root": {
												Description:         "Indicates that the container must run as a non-root user.",
												MarkdownDescription: "Indicates that the container must run as a non-root user.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_user": {
												Description:         "The UID to run the entrypoint of the container process.",
												MarkdownDescription: "The UID to run the entrypoint of the container process.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"se_linux_options": {
												Description:         "The SELinux context to be applied to the container.",
												MarkdownDescription: "The SELinux context to be applied to the container.",

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
												Description:         "The seccomp options to use by this container.",
												MarkdownDescription: "The seccomp options to use by this container.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"localhost_profile": {
														Description:         "localhostProfile indicates a profile defined in a file on the node sho",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node sho",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "type indicates which kind of seccomp profile will be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied.",

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
												Description:         "The Windows specific settings applied to all containers.",
												MarkdownDescription: "The Windows specific settings applied to all containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"gmsa_credential_spec": {
														Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github",
														MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github",

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
														Description:         "The UserName in Windows to run the entrypoint of the container process",
														MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process",

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
										Description:         "StartupProbe indicates that the Pod has successfully initialized.",
										MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "One and only one of the following should be specified.",
												MarkdownDescription: "One and only one of the following should be specified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the worki",
														MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
												Description:         "Minimum consecutive failures for the probe to be considered failed aft",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed aft",

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
														Description:         "Host name to connect to, defaults to the pod IP.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
														Description:         "Name or number of the port to access on the container.",
														MarkdownDescription: "Name or number of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Number of seconds after the container has started before liveness prob",
												MarkdownDescription: "Number of seconds after the container has started before liveness prob",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successfu",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successfu",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tcp_socket": {
												Description:         "TCPSocket specifies an action involving a TCP port.",
												MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
														Description:         "Number or name of the port to access on the container.",
														MarkdownDescription: "Number or name of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Optional duration in seconds the pod needs to terminate gracefully upo",
												MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upo",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"timeout_seconds": {
												Description:         "Number of seconds after which the probe times out.",
												MarkdownDescription: "Number of seconds after which the probe times out.",

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
										Description:         "Whether this container should allocate a buffer for stdin in the conta",
										MarkdownDescription: "Whether this container should allocate a buffer for stdin in the conta",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"stdin_once": {
										Description:         "Whether the container runtime should close the stdin channel after it ",
										MarkdownDescription: "Whether the container runtime should close the stdin channel after it ",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"termination_message_path": {
										Description:         "Optional: Path at which the file to which the container's termination ",
										MarkdownDescription: "Optional: Path at which the file to which the container's termination ",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"termination_message_policy": {
										Description:         "Indicate how the termination message should be populated.",
										MarkdownDescription: "Indicate how the termination message should be populated.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tty": {
										Description:         "Whether this container should allocate a TTY for itself, also requires",
										MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_devices": {
										Description:         "volumeDevices is the list of block devices to be used by the container",
										MarkdownDescription: "volumeDevices is the list of block devices to be used by the container",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"device_path": {
												Description:         "devicePath is the path inside of the container that the device will be",
												MarkdownDescription: "devicePath is the path inside of the container that the device will be",

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
										Description:         "Pod volumes to mount into the container's filesystem.",
										MarkdownDescription: "Pod volumes to mount into the container's filesystem.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"mount_path": {
												Description:         "Path within the container at which the volume should be mounted.",
												MarkdownDescription: "Path within the container at which the volume should be mounted.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"mount_propagation": {
												Description:         "mountPropagation determines how mounts are propagated from the host to",
												MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

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
												Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
												MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sub_path": {
												Description:         "Path within the volume from which the container's volume should be mou",
												MarkdownDescription: "Path within the volume from which the container's volume should be mou",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sub_path_expr": {
												Description:         "Expanded path within the volume from which the container's volume shou",
												MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

									"working_dir": {
										Description:         "Container's working directory.",
										MarkdownDescription: "Container's working directory.",

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
								Description:         "MetaData to add to the pod.",
								MarkdownDescription: "MetaData to add to the pod.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Key - Value pair that may be set by external tools to store and retrie",
										MarkdownDescription: "Key - Value pair that may be set by external tools to store and retrie",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"labels": {
										Description:         "Key - Value pairs that can be used to organize and categorize scope an",
										MarkdownDescription: "Key - Value pairs that can be used to organize and categorize scope an",

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

							"multi_pod_per_host": {
								Description:         "If set true then multiple pods can be created per Kubernetes Node.",
								MarkdownDescription: "If set true then multiple pods can be created per Kubernetes Node.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_selector": {
								Description:         "NodeSelector constraints for this pod.",
								MarkdownDescription: "NodeSelector constraints for this pod.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"security_context": {
								Description:         "SecurityContext holds pod-level security attributes and common contain",
								MarkdownDescription: "SecurityContext holds pod-level security attributes and common contain",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"fs_group": {
										Description:         "A special supplemental group that applies to all containers in a pod.",
										MarkdownDescription: "A special supplemental group that applies to all containers in a pod.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"fs_group_change_policy": {
										Description:         "fsGroupChangePolicy defines behavior of changing ownership and permiss",
										MarkdownDescription: "fsGroupChangePolicy defines behavior of changing ownership and permiss",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"run_as_group": {
										Description:         "The GID to run the entrypoint of the container process.",
										MarkdownDescription: "The GID to run the entrypoint of the container process.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"run_as_non_root": {
										Description:         "Indicates that the container must run as a non-root user.",
										MarkdownDescription: "Indicates that the container must run as a non-root user.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"run_as_user": {
										Description:         "The UID to run the entrypoint of the container process.",
										MarkdownDescription: "The UID to run the entrypoint of the container process.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"se_linux_options": {
										Description:         "The SELinux context to be applied to all containers.",
										MarkdownDescription: "The SELinux context to be applied to all containers.",

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
										Description:         "The seccomp options to use by the containers in this pod.",
										MarkdownDescription: "The seccomp options to use by the containers in this pod.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"localhost_profile": {
												Description:         "localhostProfile indicates a profile defined in a file on the node sho",
												MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node sho",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "type indicates which kind of seccomp profile will be applied.",
												MarkdownDescription: "type indicates which kind of seccomp profile will be applied.",

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
										Description:         "A list of groups applied to the first process run in each container, i",
										MarkdownDescription: "A list of groups applied to the first process run in each container, i",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sysctls": {
										Description:         "Sysctls hold a list of namespaced sysctls used for the pod.",
										MarkdownDescription: "Sysctls hold a list of namespaced sysctls used for the pod.",

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
										Description:         "The Windows specific settings applied to all containers.",
										MarkdownDescription: "The Windows specific settings applied to all containers.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"gmsa_credential_spec": {
												Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github",
												MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github",

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
												Description:         "The UserName in Windows to run the entrypoint of the container process",
												MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process",

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

							"sidecars": {
								Description:         "Sidecars to add to the pod.",
								MarkdownDescription: "Sidecars to add to the pod.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"args": {
										Description:         "Arguments to the entrypoint.",
										MarkdownDescription: "Arguments to the entrypoint.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"command": {
										Description:         "Entrypoint array. Not executed within a shell.",
										MarkdownDescription: "Entrypoint array. Not executed within a shell.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"env": {
										Description:         "List of environment variables to set in the container.",
										MarkdownDescription: "List of environment variables to set in the container.",

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
												Description:         "Variable references $(VAR_NAME) are expanded using the previous define",
												MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous define",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value_from": {
												Description:         "Source for the environment variable's value.",
												MarkdownDescription: "Source for the environment variable's value.",

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
																Description:         "Name of the referent. More info: https://kubernetes.",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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
														Description:         "Selects a field of the pod: supports metadata.name, metadata.",
														MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"api_version": {
																Description:         "Version of the schema the FieldPath is written in terms of, defaults t",
																MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults t",

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
														Description:         "Selects a resource of the container: only resources limits and request",
														MarkdownDescription: "Selects a resource of the container: only resources limits and request",

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
																Description:         "Name of the referent. More info: https://kubernetes.",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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
										Description:         "List of sources to populate environment variables in the container.",
										MarkdownDescription: "List of sources to populate environment variables in the container.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"config_map_ref": {
												Description:         "The ConfigMap to select from",
												MarkdownDescription: "The ConfigMap to select from",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of the referent. More info: https://kubernetes.",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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
												Description:         "An optional identifier to prepend to each key in the ConfigMap.",
												MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap.",

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
														Description:         "Name of the referent. More info: https://kubernetes.",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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

									"image": {
										Description:         "Docker image name. More info: https://kubernetes.",
										MarkdownDescription: "Docker image name. More info: https://kubernetes.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image_pull_policy": {
										Description:         "Image pull policy. One of Always, Never, IfNotPresent.",
										MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"lifecycle": {
										Description:         "Actions that the management system should take in response to containe",
										MarkdownDescription: "Actions that the management system should take in response to containe",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"post_start": {
												Description:         "PostStart is called immediately after a container is created.",
												MarkdownDescription: "PostStart is called immediately after a container is created.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "One and only one of the following should be specified.",
														MarkdownDescription: "One and only one of the following should be specified.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the worki",
																MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
																Description:         "Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
																Description:         "Name or number of the port to access on the container.",
																MarkdownDescription: "Name or number of the port to access on the container.",

																Type: utilities.IntOrStringType{},

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
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
																Description:         "Number or name of the port to access on the container.",
																MarkdownDescription: "Number or name of the port to access on the container.",

																Type: utilities.IntOrStringType{},

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
												Description:         "PreStop is called immediately before a container is terminated due to ",
												MarkdownDescription: "PreStop is called immediately before a container is terminated due to ",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"exec": {
														Description:         "One and only one of the following should be specified.",
														MarkdownDescription: "One and only one of the following should be specified.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"command": {
																Description:         "Command is the command line to execute inside the container, the worki",
																MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
																Description:         "Host name to connect to, defaults to the pod IP.",
																MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
																Description:         "Name or number of the port to access on the container.",
																MarkdownDescription: "Name or number of the port to access on the container.",

																Type: utilities.IntOrStringType{},

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
														Description:         "TCPSocket specifies an action involving a TCP port.",
														MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
																Description:         "Number or name of the port to access on the container.",
																MarkdownDescription: "Number or name of the port to access on the container.",

																Type: utilities.IntOrStringType{},

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
										Description:         "Periodic probe of container liveness.",
										MarkdownDescription: "Periodic probe of container liveness.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "One and only one of the following should be specified.",
												MarkdownDescription: "One and only one of the following should be specified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the worki",
														MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
												Description:         "Minimum consecutive failures for the probe to be considered failed aft",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed aft",

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
														Description:         "Host name to connect to, defaults to the pod IP.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
														Description:         "Name or number of the port to access on the container.",
														MarkdownDescription: "Name or number of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Number of seconds after the container has started before liveness prob",
												MarkdownDescription: "Number of seconds after the container has started before liveness prob",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successfu",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successfu",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tcp_socket": {
												Description:         "TCPSocket specifies an action involving a TCP port.",
												MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
														Description:         "Number or name of the port to access on the container.",
														MarkdownDescription: "Number or name of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Optional duration in seconds the pod needs to terminate gracefully upo",
												MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upo",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"timeout_seconds": {
												Description:         "Number of seconds after which the probe times out.",
												MarkdownDescription: "Number of seconds after which the probe times out.",

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
										Description:         "Name of the container specified as a DNS_LABEL.",
										MarkdownDescription: "Name of the container specified as a DNS_LABEL.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"ports": {
										Description:         "List of ports to expose from the container.",
										MarkdownDescription: "List of ports to expose from the container.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"container_port": {
												Description:         "Number of port to expose on the pod's IP address.",
												MarkdownDescription: "Number of port to expose on the pod's IP address.",

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
												Description:         "Number of port to expose on the host.",
												MarkdownDescription: "Number of port to expose on the host.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod.",
												MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod.",

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
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"readiness_probe": {
										Description:         "Periodic probe of container service readiness.",
										MarkdownDescription: "Periodic probe of container service readiness.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "One and only one of the following should be specified.",
												MarkdownDescription: "One and only one of the following should be specified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the worki",
														MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
												Description:         "Minimum consecutive failures for the probe to be considered failed aft",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed aft",

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
														Description:         "Host name to connect to, defaults to the pod IP.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
														Description:         "Name or number of the port to access on the container.",
														MarkdownDescription: "Name or number of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Number of seconds after the container has started before liveness prob",
												MarkdownDescription: "Number of seconds after the container has started before liveness prob",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successfu",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successfu",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tcp_socket": {
												Description:         "TCPSocket specifies an action involving a TCP port.",
												MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
														Description:         "Number or name of the port to access on the container.",
														MarkdownDescription: "Number or name of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Optional duration in seconds the pod needs to terminate gracefully upo",
												MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upo",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"timeout_seconds": {
												Description:         "Number of seconds after which the probe times out.",
												MarkdownDescription: "Number of seconds after which the probe times out.",

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

									"resources": {
										Description:         "Compute Resources required by this container. Cannot be updated.",
										MarkdownDescription: "Compute Resources required by this container. Cannot be updated.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"limits": {
												Description:         "Limits describes the maximum amount of compute resources allowed.",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"requests": {
												Description:         "Requests describes the minimum amount of compute resources required.",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required.",

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
										Description:         "Security options the pod should run with.",
										MarkdownDescription: "Security options the pod should run with.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"allow_privilege_escalation": {
												Description:         "AllowPrivilegeEscalation controls whether a process can gain more priv",
												MarkdownDescription: "AllowPrivilegeEscalation controls whether a process can gain more priv",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"capabilities": {
												Description:         "The capabilities to add/drop when running containers.",
												MarkdownDescription: "The capabilities to add/drop when running containers.",

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
												Description:         "Run container in privileged mode.",
												MarkdownDescription: "Run container in privileged mode.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"proc_mount": {
												Description:         "procMount denotes the type of proc mount to use for the containers.",
												MarkdownDescription: "procMount denotes the type of proc mount to use for the containers.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only_root_filesystem": {
												Description:         "Whether this container has a read-only root filesystem.",
												MarkdownDescription: "Whether this container has a read-only root filesystem.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_group": {
												Description:         "The GID to run the entrypoint of the container process.",
												MarkdownDescription: "The GID to run the entrypoint of the container process.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_non_root": {
												Description:         "Indicates that the container must run as a non-root user.",
												MarkdownDescription: "Indicates that the container must run as a non-root user.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"run_as_user": {
												Description:         "The UID to run the entrypoint of the container process.",
												MarkdownDescription: "The UID to run the entrypoint of the container process.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"se_linux_options": {
												Description:         "The SELinux context to be applied to the container.",
												MarkdownDescription: "The SELinux context to be applied to the container.",

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
												Description:         "The seccomp options to use by this container.",
												MarkdownDescription: "The seccomp options to use by this container.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"localhost_profile": {
														Description:         "localhostProfile indicates a profile defined in a file on the node sho",
														MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node sho",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "type indicates which kind of seccomp profile will be applied.",
														MarkdownDescription: "type indicates which kind of seccomp profile will be applied.",

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
												Description:         "The Windows specific settings applied to all containers.",
												MarkdownDescription: "The Windows specific settings applied to all containers.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"gmsa_credential_spec": {
														Description:         "GMSACredentialSpec is where the GMSA admission webhook (https://github",
														MarkdownDescription: "GMSACredentialSpec is where the GMSA admission webhook (https://github",

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
														Description:         "The UserName in Windows to run the entrypoint of the container process",
														MarkdownDescription: "The UserName in Windows to run the entrypoint of the container process",

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
										Description:         "StartupProbe indicates that the Pod has successfully initialized.",
										MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "One and only one of the following should be specified.",
												MarkdownDescription: "One and only one of the following should be specified.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the worki",
														MarkdownDescription: "Command is the command line to execute inside the container, the worki",

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
												Description:         "Minimum consecutive failures for the probe to be considered failed aft",
												MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed aft",

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
														Description:         "Host name to connect to, defaults to the pod IP.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP.",

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
														Description:         "Name or number of the port to access on the container.",
														MarkdownDescription: "Name or number of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Number of seconds after the container has started before liveness prob",
												MarkdownDescription: "Number of seconds after the container has started before liveness prob",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"period_seconds": {
												Description:         "How often (in seconds) to perform the probe. Default to 10 seconds.",
												MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"success_threshold": {
												Description:         "Minimum consecutive successes for the probe to be considered successfu",
												MarkdownDescription: "Minimum consecutive successes for the probe to be considered successfu",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tcp_socket": {
												Description:         "TCPSocket specifies an action involving a TCP port.",
												MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

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
														Description:         "Number or name of the port to access on the container.",
														MarkdownDescription: "Number or name of the port to access on the container.",

														Type: utilities.IntOrStringType{},

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
												Description:         "Optional duration in seconds the pod needs to terminate gracefully upo",
												MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upo",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"timeout_seconds": {
												Description:         "Number of seconds after which the probe times out.",
												MarkdownDescription: "Number of seconds after which the probe times out.",

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
										Description:         "Whether this container should allocate a buffer for stdin in the conta",
										MarkdownDescription: "Whether this container should allocate a buffer for stdin in the conta",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"stdin_once": {
										Description:         "Whether the container runtime should close the stdin channel after it ",
										MarkdownDescription: "Whether the container runtime should close the stdin channel after it ",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"termination_message_path": {
										Description:         "Optional: Path at which the file to which the container's termination ",
										MarkdownDescription: "Optional: Path at which the file to which the container's termination ",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"termination_message_policy": {
										Description:         "Indicate how the termination message should be populated.",
										MarkdownDescription: "Indicate how the termination message should be populated.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tty": {
										Description:         "Whether this container should allocate a TTY for itself, also requires",
										MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"volume_devices": {
										Description:         "volumeDevices is the list of block devices to be used by the container",
										MarkdownDescription: "volumeDevices is the list of block devices to be used by the container",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"device_path": {
												Description:         "devicePath is the path inside of the container that the device will be",
												MarkdownDescription: "devicePath is the path inside of the container that the device will be",

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
										Description:         "Pod volumes to mount into the container's filesystem.",
										MarkdownDescription: "Pod volumes to mount into the container's filesystem.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"mount_path": {
												Description:         "Path within the container at which the volume should be mounted.",
												MarkdownDescription: "Path within the container at which the volume should be mounted.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"mount_propagation": {
												Description:         "mountPropagation determines how mounts are propagated from the host to",
												MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

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
												Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
												MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sub_path": {
												Description:         "Path within the volume from which the container's volume should be mou",
												MarkdownDescription: "Path within the volume from which the container's volume should be mou",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sub_path_expr": {
												Description:         "Expanded path within the volume from which the container's volume shou",
												MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

									"working_dir": {
										Description:         "Container's working directory.",
										MarkdownDescription: "Container's working directory.",

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
								Description:         "Tolerations for this pod.",
								MarkdownDescription: "Tolerations for this pod.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"effect": {
										Description:         "Effect indicates the taint effect to match.",
										MarkdownDescription: "Effect indicates the taint effect to match.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key": {
										Description:         "Key is the taint key that the toleration applies to.",
										MarkdownDescription: "Key is the taint key that the toleration applies to.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"operator": {
										Description:         "Operator represents a key's relationship to the value.",
										MarkdownDescription: "Operator represents a key's relationship to the value.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"toleration_seconds": {
										Description:         "TolerationSeconds represents the period of time the toleration (which ",
										MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which ",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"value": {
										Description:         "Value is the taint value the toleration matches to.",
										MarkdownDescription: "Value is the taint value the toleration matches to.",

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

					"rack_config": {
						Description:         "RackConfig Configures the operator to deploy rack aware Aerospike clus",
						MarkdownDescription: "RackConfig Configures the operator to deploy rack aware Aerospike clus",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"namespaces": {
								Description:         "List of Aerospike namespaces for which rack feature will be enabled",
								MarkdownDescription: "List of Aerospike namespaces for which rack feature will be enabled",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"racks": {
								Description:         "Racks is the list of all racks",
								MarkdownDescription: "Racks is the list of all racks",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"aerospike_config": {
										Description:         "AerospikeConfig overrides the common AerospikeConfig for this Rack.",
										MarkdownDescription: "AerospikeConfig overrides the common AerospikeConfig for this Rack.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"effective_aerospike_config": {
										Description:         "Effective/operative Aerospike config.",
										MarkdownDescription: "Effective/operative Aerospike config.",

										Type: utilities.DynamicType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"effective_pod_spec": {
										Description:         "Effective/operative PodSpec.",
										MarkdownDescription: "Effective/operative PodSpec.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"affinity": {
												Description:         "Affinity rules for pod placement.",
												MarkdownDescription: "Affinity rules for pod placement.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"node_affinity": {
														Description:         "Describes node affinity scheduling rules for the pod.",
														MarkdownDescription: "Describes node affinity scheduling rules for the pod.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"preferred_during_scheduling_ignored_during_execution": {
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the a",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the a",

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
																						Description:         "Represents a key's relationship to a set of values.",
																						MarkdownDescription: "Represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "An array of string values.",
																						MarkdownDescription: "An array of string values.",

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
																						Description:         "Represents a key's relationship to a set of values.",
																						MarkdownDescription: "Represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "An array of string values.",
																						MarkdownDescription: "An array of string values.",

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
																		Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in",
																		MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in",

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
																Description:         "If the affinity requirements specified by this field are not met at sc",
																MarkdownDescription: "If the affinity requirements specified by this field are not met at sc",

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
																						Description:         "Represents a key's relationship to a set of values.",
																						MarkdownDescription: "Represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "An array of string values.",
																						MarkdownDescription: "An array of string values.",

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
																						Description:         "Represents a key's relationship to a set of values.",
																						MarkdownDescription: "Represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "An array of string values.",
																						MarkdownDescription: "An array of string values.",

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
														Description:         "Describes pod affinity scheduling rules (e.g.",
														MarkdownDescription: "Describes pod affinity scheduling rules (e.g.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"preferred_during_scheduling_ignored_during_execution": {
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the a",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the a",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"pod_affinity_term": {
																		Description:         "Required.",
																		MarkdownDescription: "Required.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"label_selector": {
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "matchExpressions is a list of label selector requirements.",
																						MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																								Description:         "operator represents a key's relationship to a set of values.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "values is an array of string values.",
																								MarkdownDescription: "values is an array of string values.",

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
																						Description:         "matchLabels is a map of {key,value} pairs.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																				Description:         "A label query over the set of namespaces that the term applies to.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "matchExpressions is a list of label selector requirements.",
																						MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																								Description:         "operator represents a key's relationship to a set of values.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "values is an array of string values.",
																								MarkdownDescription: "values is an array of string values.",

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
																						Description:         "matchLabels is a map of {key,value} pairs.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																				Description:         "namespaces specifies a static list of namespace names that the term ap",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"topology_key": {
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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
																		Description:         "weight associated with matching the corresponding podAffinityTerm, in ",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in ",

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
																Description:         "If the affinity requirements specified by this field are not met at sc",
																MarkdownDescription: "If the affinity requirements specified by this field are not met at sc",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"label_selector": {
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																		Description:         "A label query over the set of namespaces that the term applies to.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																		Description:         "namespaces specifies a static list of namespace names that the term ap",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"topology_key": {
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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
														Description:         "Describes pod anti-affinity scheduling rules (e.g.",
														MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"preferred_during_scheduling_ignored_during_execution": {
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the a",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the a",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"pod_affinity_term": {
																		Description:         "Required.",
																		MarkdownDescription: "Required.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"label_selector": {
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "matchExpressions is a list of label selector requirements.",
																						MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																								Description:         "operator represents a key's relationship to a set of values.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "values is an array of string values.",
																								MarkdownDescription: "values is an array of string values.",

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
																						Description:         "matchLabels is a map of {key,value} pairs.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																				Description:         "A label query over the set of namespaces that the term applies to.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "matchExpressions is a list of label selector requirements.",
																						MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																								Description:         "operator represents a key's relationship to a set of values.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "values is an array of string values.",
																								MarkdownDescription: "values is an array of string values.",

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
																						Description:         "matchLabels is a map of {key,value} pairs.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																				Description:         "namespaces specifies a static list of namespace names that the term ap",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"topology_key": {
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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
																		Description:         "weight associated with matching the corresponding podAffinityTerm, in ",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in ",

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
																Description:         "If the anti-affinity requirements specified by this field are not met ",
																MarkdownDescription: "If the anti-affinity requirements specified by this field are not met ",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"label_selector": {
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																		Description:         "A label query over the set of namespaces that the term applies to.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																		Description:         "namespaces specifies a static list of namespace names that the term ap",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"topology_key": {
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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

											"node_selector": {
												Description:         "NodeSelector constraints for this pod.",
												MarkdownDescription: "NodeSelector constraints for this pod.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tolerations": {
												Description:         "Tolerations for this pod.",
												MarkdownDescription: "Tolerations for this pod.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"effect": {
														Description:         "Effect indicates the taint effect to match.",
														MarkdownDescription: "Effect indicates the taint effect to match.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": {
														Description:         "Key is the taint key that the toleration applies to.",
														MarkdownDescription: "Key is the taint key that the toleration applies to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"operator": {
														Description:         "Operator represents a key's relationship to the value.",
														MarkdownDescription: "Operator represents a key's relationship to the value.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"toleration_seconds": {
														Description:         "TolerationSeconds represents the period of time the toleration (which ",
														MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which ",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "Value is the taint value the toleration matches to.",
														MarkdownDescription: "Value is the taint value the toleration matches to.",

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

									"effective_storage": {
										Description:         "Effective/operative storage.",
										MarkdownDescription: "Effective/operative storage.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"block_volume_policy": {
												Description:         "BlockVolumePolicy contains default policies for block volumes.",
												MarkdownDescription: "BlockVolumePolicy contains default policies for block volumes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cascade_delete": {
														Description:         "CascadeDelete determines if the persistent volumes are deleted after t",
														MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after t",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_cascade_delete": {
														Description:         "Effective/operative value to use for cascade delete after applying def",
														MarkdownDescription: "Effective/operative value to use for cascade delete after applying def",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_init_method": {
														Description:         "Effective/operative value to use as the volume init method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume init method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"effective_wipe_method": {
														Description:         "Effective/operative value to use as the volume wipe method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume wipe method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"init_method": {
														Description:         "InitMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"wipe_method": {
														Description:         "WipeMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"filesystem_volume_policy": {
												Description:         "FileSystemVolumePolicy contains default policies for filesystem volume",
												MarkdownDescription: "FileSystemVolumePolicy contains default policies for filesystem volume",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cascade_delete": {
														Description:         "CascadeDelete determines if the persistent volumes are deleted after t",
														MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after t",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_cascade_delete": {
														Description:         "Effective/operative value to use for cascade delete after applying def",
														MarkdownDescription: "Effective/operative value to use for cascade delete after applying def",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_init_method": {
														Description:         "Effective/operative value to use as the volume init method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume init method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"effective_wipe_method": {
														Description:         "Effective/operative value to use as the volume wipe method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume wipe method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"init_method": {
														Description:         "InitMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"wipe_method": {
														Description:         "WipeMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volumes": {
												Description:         "Volumes list to attach to created pods.",
												MarkdownDescription: "Volumes list to attach to created pods.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"aerospike": {
														Description:         "Aerospike attachment of this volume on Aerospike server container.",
														MarkdownDescription: "Aerospike attachment of this volume on Aerospike server container.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"mount_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"mount_propagation": {
																		Description:         "mountPropagation determines how mounts are propagated from the host to",
																		MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"read_only": {
																		Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
																		MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path": {
																		Description:         "Path within the volume from which the container's volume should be mou",
																		MarkdownDescription: "Path within the volume from which the container's volume should be mou",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path_expr": {
																		Description:         "Expanded path within the volume from which the container's volume shou",
																		MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

															"path": {
																Description:         "Path to attach the volume on the Aerospike server container.",
																MarkdownDescription: "Path to attach the volume on the Aerospike server container.",

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

													"cascade_delete": {
														Description:         "CascadeDelete determines if the persistent volumes are deleted after t",
														MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after t",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_cascade_delete": {
														Description:         "Effective/operative value to use for cascade delete after applying def",
														MarkdownDescription: "Effective/operative value to use for cascade delete after applying def",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_init_method": {
														Description:         "Effective/operative value to use as the volume init method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume init method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"effective_wipe_method": {
														Description:         "Effective/operative value to use as the volume wipe method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume wipe method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"init_containers": {
														Description:         "InitContainers are additional init containers where this volume will b",
														MarkdownDescription: "InitContainers are additional init containers where this volume will b",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"container_name": {
																Description:         "ContainerName is the name of the container to attach this volume to.",
																MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mount_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"mount_propagation": {
																		Description:         "mountPropagation determines how mounts are propagated from the host to",
																		MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"read_only": {
																		Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
																		MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path": {
																		Description:         "Path within the volume from which the container's volume should be mou",
																		MarkdownDescription: "Path within the volume from which the container's volume should be mou",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path_expr": {
																		Description:         "Expanded path within the volume from which the container's volume shou",
																		MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

															"path": {
																Description:         "Path to attach the volume on the container.",
																MarkdownDescription: "Path to attach the volume on the container.",

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

													"init_method": {
														Description:         "InitMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"name": {
														Description:         "Name for this volume, Name or path should be given.",
														MarkdownDescription: "Name for this volume, Name or path should be given.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"sidecars": {
														Description:         "Sidecars are side containers where this volume will be mounted",
														MarkdownDescription: "Sidecars are side containers where this volume will be mounted",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"container_name": {
																Description:         "ContainerName is the name of the container to attach this volume to.",
																MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mount_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"mount_propagation": {
																		Description:         "mountPropagation determines how mounts are propagated from the host to",
																		MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"read_only": {
																		Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
																		MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path": {
																		Description:         "Path within the volume from which the container's volume should be mou",
																		MarkdownDescription: "Path within the volume from which the container's volume should be mou",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path_expr": {
																		Description:         "Expanded path within the volume from which the container's volume shou",
																		MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

															"path": {
																Description:         "Path to attach the volume on the container.",
																MarkdownDescription: "Path to attach the volume on the container.",

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

													"source": {
														Description:         "Source of this volume.",
														MarkdownDescription: "Source of this volume.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap represents a configMap that should populate this volume",
																MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_mode": {
																		Description:         "Optional: mode bits used to set permissions on created files by defaul",
																		MarkdownDescription: "Optional: mode bits used to set permissions on created files by defaul",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"items": {
																		Description:         "If unspecified, each key-value pair in the Data field of the reference",
																		MarkdownDescription: "If unspecified, each key-value pair in the Data field of the reference",

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
																				Description:         "Optional: mode bits used to set permissions on this file.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "The relative path of the file to map the key to.",
																				MarkdownDescription: "The relative path of the file to map the key to.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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

															"empty_dir": {
																Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime",
																MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"medium": {
																		Description:         "What type of storage medium should back this directory.",
																		MarkdownDescription: "What type of storage medium should back this directory.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"size_limit": {
																		Description:         "Total amount of local storage required for this EmptyDir volume.",
																		MarkdownDescription: "Total amount of local storage required for this EmptyDir volume.",

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

															"persistent_volume": {
																Description:         "PersistentVolumeSpec describes a persistent volume to claim and attach",
																MarkdownDescription: "PersistentVolumeSpec describes a persistent volume to claim and attach",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"access_modes": {
																		Description:         "Name for creating PVC for this volume, Name or path should be given Na",
																		MarkdownDescription: "Name for creating PVC for this volume, Name or path should be given Na",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"annotations": {
																				Description:         "Key - Value pair that may be set by external tools to store and retrie",
																				MarkdownDescription: "Key - Value pair that may be set by external tools to store and retrie",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"labels": {
																				Description:         "Key - Value pairs that can be used to organize and categorize scope an",
																				MarkdownDescription: "Key - Value pairs that can be used to organize and categorize scope an",

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
																		Description:         "A label query over volumes to consider for binding.",
																		MarkdownDescription: "A label query over volumes to consider for binding.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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

																	"size": {
																		Description:         "Size of volume.",
																		MarkdownDescription: "Size of volume.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"storage_class": {
																		Description:         "StorageClass should be pre-created by user.",
																		MarkdownDescription: "StorageClass should be pre-created by user.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"volume_mode": {
																		Description:         "VolumeMode specifies if the volume is block/raw or a filesystem.",
																		MarkdownDescription: "VolumeMode specifies if the volume is block/raw or a filesystem.",

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

															"secret": {
																Description:         "Adapts a Secret into a volume.",
																MarkdownDescription: "Adapts a Secret into a volume.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_mode": {
																		Description:         "Optional: mode bits used to set permissions on created files by defaul",
																		MarkdownDescription: "Optional: mode bits used to set permissions on created files by defaul",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"items": {
																		Description:         "If unspecified, each key-value pair in the Data field of the reference",
																		MarkdownDescription: "If unspecified, each key-value pair in the Data field of the reference",

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
																				Description:         "Optional: mode bits used to set permissions on this file.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "The relative path of the file to map the key to.",
																				MarkdownDescription: "The relative path of the file to map the key to.",

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
																		Description:         "Name of the secret in the pod's namespace to use.",
																		MarkdownDescription: "Name of the secret in the pod's namespace to use.",

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

													"wipe_method": {
														Description:         "WipeMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
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

									"id": {
										Description:         "Identifier for the rack",
										MarkdownDescription: "Identifier for the rack",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"node_name": {
										Description:         "K8s Node name for setting rack affinity.",
										MarkdownDescription: "K8s Node name for setting rack affinity.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_spec": {
										Description:         "PodSpec to use for the pods in this rack.",
										MarkdownDescription: "PodSpec to use for the pods in this rack.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"affinity": {
												Description:         "Affinity rules for pod placement.",
												MarkdownDescription: "Affinity rules for pod placement.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"node_affinity": {
														Description:         "Describes node affinity scheduling rules for the pod.",
														MarkdownDescription: "Describes node affinity scheduling rules for the pod.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"preferred_during_scheduling_ignored_during_execution": {
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the a",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the a",

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
																						Description:         "Represents a key's relationship to a set of values.",
																						MarkdownDescription: "Represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "An array of string values.",
																						MarkdownDescription: "An array of string values.",

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
																						Description:         "Represents a key's relationship to a set of values.",
																						MarkdownDescription: "Represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "An array of string values.",
																						MarkdownDescription: "An array of string values.",

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
																		Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in",
																		MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in",

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
																Description:         "If the affinity requirements specified by this field are not met at sc",
																MarkdownDescription: "If the affinity requirements specified by this field are not met at sc",

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
																						Description:         "Represents a key's relationship to a set of values.",
																						MarkdownDescription: "Represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "An array of string values.",
																						MarkdownDescription: "An array of string values.",

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
																						Description:         "Represents a key's relationship to a set of values.",
																						MarkdownDescription: "Represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "An array of string values.",
																						MarkdownDescription: "An array of string values.",

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
														Description:         "Describes pod affinity scheduling rules (e.g.",
														MarkdownDescription: "Describes pod affinity scheduling rules (e.g.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"preferred_during_scheduling_ignored_during_execution": {
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the a",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the a",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"pod_affinity_term": {
																		Description:         "Required.",
																		MarkdownDescription: "Required.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"label_selector": {
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "matchExpressions is a list of label selector requirements.",
																						MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																								Description:         "operator represents a key's relationship to a set of values.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "values is an array of string values.",
																								MarkdownDescription: "values is an array of string values.",

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
																						Description:         "matchLabels is a map of {key,value} pairs.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																				Description:         "A label query over the set of namespaces that the term applies to.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "matchExpressions is a list of label selector requirements.",
																						MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																								Description:         "operator represents a key's relationship to a set of values.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "values is an array of string values.",
																								MarkdownDescription: "values is an array of string values.",

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
																						Description:         "matchLabels is a map of {key,value} pairs.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																				Description:         "namespaces specifies a static list of namespace names that the term ap",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"topology_key": {
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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
																		Description:         "weight associated with matching the corresponding podAffinityTerm, in ",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in ",

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
																Description:         "If the affinity requirements specified by this field are not met at sc",
																MarkdownDescription: "If the affinity requirements specified by this field are not met at sc",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"label_selector": {
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																		Description:         "A label query over the set of namespaces that the term applies to.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																		Description:         "namespaces specifies a static list of namespace names that the term ap",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"topology_key": {
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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
														Description:         "Describes pod anti-affinity scheduling rules (e.g.",
														MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"preferred_during_scheduling_ignored_during_execution": {
																Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the a",
																MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the a",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"pod_affinity_term": {
																		Description:         "Required.",
																		MarkdownDescription: "Required.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"label_selector": {
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "matchExpressions is a list of label selector requirements.",
																						MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																								Description:         "operator represents a key's relationship to a set of values.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "values is an array of string values.",
																								MarkdownDescription: "values is an array of string values.",

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
																						Description:         "matchLabels is a map of {key,value} pairs.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																				Description:         "A label query over the set of namespaces that the term applies to.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"match_expressions": {
																						Description:         "matchExpressions is a list of label selector requirements.",
																						MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																								Description:         "operator represents a key's relationship to a set of values.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values.",

																								Type: types.StringType,

																								Required: true,
																								Optional: false,
																								Computed: false,
																							},

																							"values": {
																								Description:         "values is an array of string values.",
																								MarkdownDescription: "values is an array of string values.",

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
																						Description:         "matchLabels is a map of {key,value} pairs.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																				Description:         "namespaces specifies a static list of namespace names that the term ap",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"topology_key": {
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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
																		Description:         "weight associated with matching the corresponding podAffinityTerm, in ",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in ",

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
																Description:         "If the anti-affinity requirements specified by this field are not met ",
																MarkdownDescription: "If the anti-affinity requirements specified by this field are not met ",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"label_selector": {
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																		Description:         "A label query over the set of namespaces that the term applies to.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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
																		Description:         "namespaces specifies a static list of namespace names that the term ap",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term ap",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"topology_key": {
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affin",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affin",

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

											"node_selector": {
												Description:         "NodeSelector constraints for this pod.",
												MarkdownDescription: "NodeSelector constraints for this pod.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tolerations": {
												Description:         "Tolerations for this pod.",
												MarkdownDescription: "Tolerations for this pod.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"effect": {
														Description:         "Effect indicates the taint effect to match.",
														MarkdownDescription: "Effect indicates the taint effect to match.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": {
														Description:         "Key is the taint key that the toleration applies to.",
														MarkdownDescription: "Key is the taint key that the toleration applies to.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"operator": {
														Description:         "Operator represents a key's relationship to the value.",
														MarkdownDescription: "Operator represents a key's relationship to the value.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"toleration_seconds": {
														Description:         "TolerationSeconds represents the period of time the toleration (which ",
														MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which ",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "Value is the taint value the toleration matches to.",
														MarkdownDescription: "Value is the taint value the toleration matches to.",

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

									"rack_label": {
										Description:         "RackLabel for setting rack affinity.",
										MarkdownDescription: "RackLabel for setting rack affinity.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"region": {
										Description:         "Region name for setting rack affinity.",
										MarkdownDescription: "Region name for setting rack affinity.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage": {
										Description:         "Storage specify persistent storage to use for the pods in this rack.",
										MarkdownDescription: "Storage specify persistent storage to use for the pods in this rack.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"block_volume_policy": {
												Description:         "BlockVolumePolicy contains default policies for block volumes.",
												MarkdownDescription: "BlockVolumePolicy contains default policies for block volumes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cascade_delete": {
														Description:         "CascadeDelete determines if the persistent volumes are deleted after t",
														MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after t",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_cascade_delete": {
														Description:         "Effective/operative value to use for cascade delete after applying def",
														MarkdownDescription: "Effective/operative value to use for cascade delete after applying def",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_init_method": {
														Description:         "Effective/operative value to use as the volume init method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume init method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"effective_wipe_method": {
														Description:         "Effective/operative value to use as the volume wipe method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume wipe method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"init_method": {
														Description:         "InitMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"wipe_method": {
														Description:         "WipeMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"filesystem_volume_policy": {
												Description:         "FileSystemVolumePolicy contains default policies for filesystem volume",
												MarkdownDescription: "FileSystemVolumePolicy contains default policies for filesystem volume",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"cascade_delete": {
														Description:         "CascadeDelete determines if the persistent volumes are deleted after t",
														MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after t",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_cascade_delete": {
														Description:         "Effective/operative value to use for cascade delete after applying def",
														MarkdownDescription: "Effective/operative value to use for cascade delete after applying def",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_init_method": {
														Description:         "Effective/operative value to use as the volume init method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume init method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"effective_wipe_method": {
														Description:         "Effective/operative value to use as the volume wipe method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume wipe method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"init_method": {
														Description:         "InitMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"wipe_method": {
														Description:         "WipeMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volumes": {
												Description:         "Volumes list to attach to created pods.",
												MarkdownDescription: "Volumes list to attach to created pods.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"aerospike": {
														Description:         "Aerospike attachment of this volume on Aerospike server container.",
														MarkdownDescription: "Aerospike attachment of this volume on Aerospike server container.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"mount_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"mount_propagation": {
																		Description:         "mountPropagation determines how mounts are propagated from the host to",
																		MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"read_only": {
																		Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
																		MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path": {
																		Description:         "Path within the volume from which the container's volume should be mou",
																		MarkdownDescription: "Path within the volume from which the container's volume should be mou",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path_expr": {
																		Description:         "Expanded path within the volume from which the container's volume shou",
																		MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

															"path": {
																Description:         "Path to attach the volume on the Aerospike server container.",
																MarkdownDescription: "Path to attach the volume on the Aerospike server container.",

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

													"cascade_delete": {
														Description:         "CascadeDelete determines if the persistent volumes are deleted after t",
														MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after t",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_cascade_delete": {
														Description:         "Effective/operative value to use for cascade delete after applying def",
														MarkdownDescription: "Effective/operative value to use for cascade delete after applying def",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"effective_init_method": {
														Description:         "Effective/operative value to use as the volume init method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume init method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"effective_wipe_method": {
														Description:         "Effective/operative value to use as the volume wipe method after apply",
														MarkdownDescription: "Effective/operative value to use as the volume wipe method after apply",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"init_containers": {
														Description:         "InitContainers are additional init containers where this volume will b",
														MarkdownDescription: "InitContainers are additional init containers where this volume will b",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"container_name": {
																Description:         "ContainerName is the name of the container to attach this volume to.",
																MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mount_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"mount_propagation": {
																		Description:         "mountPropagation determines how mounts are propagated from the host to",
																		MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"read_only": {
																		Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
																		MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path": {
																		Description:         "Path within the volume from which the container's volume should be mou",
																		MarkdownDescription: "Path within the volume from which the container's volume should be mou",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path_expr": {
																		Description:         "Expanded path within the volume from which the container's volume shou",
																		MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

															"path": {
																Description:         "Path to attach the volume on the container.",
																MarkdownDescription: "Path to attach the volume on the container.",

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

													"init_method": {
														Description:         "InitMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
														},
													},

													"name": {
														Description:         "Name for this volume, Name or path should be given.",
														MarkdownDescription: "Name for this volume, Name or path should be given.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"sidecars": {
														Description:         "Sidecars are side containers where this volume will be mounted",
														MarkdownDescription: "Sidecars are side containers where this volume will be mounted",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"container_name": {
																Description:         "ContainerName is the name of the container to attach this volume to.",
																MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mount_options": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"mount_propagation": {
																		Description:         "mountPropagation determines how mounts are propagated from the host to",
																		MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"read_only": {
																		Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
																		MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

																		Type: types.BoolType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path": {
																		Description:         "Path within the volume from which the container's volume should be mou",
																		MarkdownDescription: "Path within the volume from which the container's volume should be mou",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"sub_path_expr": {
																		Description:         "Expanded path within the volume from which the container's volume shou",
																		MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

															"path": {
																Description:         "Path to attach the volume on the container.",
																MarkdownDescription: "Path to attach the volume on the container.",

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

													"source": {
														Description:         "Source of this volume.",
														MarkdownDescription: "Source of this volume.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "ConfigMap represents a configMap that should populate this volume",
																MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_mode": {
																		Description:         "Optional: mode bits used to set permissions on created files by defaul",
																		MarkdownDescription: "Optional: mode bits used to set permissions on created files by defaul",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"items": {
																		Description:         "If unspecified, each key-value pair in the Data field of the reference",
																		MarkdownDescription: "If unspecified, each key-value pair in the Data field of the reference",

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
																				Description:         "Optional: mode bits used to set permissions on this file.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "The relative path of the file to map the key to.",
																				MarkdownDescription: "The relative path of the file to map the key to.",

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
																		Description:         "Name of the referent. More info: https://kubernetes.",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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

															"empty_dir": {
																Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime",
																MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"medium": {
																		Description:         "What type of storage medium should back this directory.",
																		MarkdownDescription: "What type of storage medium should back this directory.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"size_limit": {
																		Description:         "Total amount of local storage required for this EmptyDir volume.",
																		MarkdownDescription: "Total amount of local storage required for this EmptyDir volume.",

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

															"persistent_volume": {
																Description:         "PersistentVolumeSpec describes a persistent volume to claim and attach",
																MarkdownDescription: "PersistentVolumeSpec describes a persistent volume to claim and attach",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"access_modes": {
																		Description:         "Name for creating PVC for this volume, Name or path should be given Na",
																		MarkdownDescription: "Name for creating PVC for this volume, Name or path should be given Na",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"metadata": {
																		Description:         "",
																		MarkdownDescription: "",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"annotations": {
																				Description:         "Key - Value pair that may be set by external tools to store and retrie",
																				MarkdownDescription: "Key - Value pair that may be set by external tools to store and retrie",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"labels": {
																				Description:         "Key - Value pairs that can be used to organize and categorize scope an",
																				MarkdownDescription: "Key - Value pairs that can be used to organize and categorize scope an",

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
																		Description:         "A label query over volumes to consider for binding.",
																		MarkdownDescription: "A label query over volumes to consider for binding.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"match_expressions": {
																				Description:         "matchExpressions is a list of label selector requirements.",
																				MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																						Description:         "operator represents a key's relationship to a set of values.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"values": {
																						Description:         "values is an array of string values.",
																						MarkdownDescription: "values is an array of string values.",

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
																				Description:         "matchLabels is a map of {key,value} pairs.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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

																	"size": {
																		Description:         "Size of volume.",
																		MarkdownDescription: "Size of volume.",

																		Type: utilities.IntOrStringType{},

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"storage_class": {
																		Description:         "StorageClass should be pre-created by user.",
																		MarkdownDescription: "StorageClass should be pre-created by user.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"volume_mode": {
																		Description:         "VolumeMode specifies if the volume is block/raw or a filesystem.",
																		MarkdownDescription: "VolumeMode specifies if the volume is block/raw or a filesystem.",

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

															"secret": {
																Description:         "Adapts a Secret into a volume.",
																MarkdownDescription: "Adapts a Secret into a volume.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"default_mode": {
																		Description:         "Optional: mode bits used to set permissions on created files by defaul",
																		MarkdownDescription: "Optional: mode bits used to set permissions on created files by defaul",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"items": {
																		Description:         "If unspecified, each key-value pair in the Data field of the reference",
																		MarkdownDescription: "If unspecified, each key-value pair in the Data field of the reference",

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
																				Description:         "Optional: mode bits used to set permissions on this file.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "The relative path of the file to map the key to.",
																				MarkdownDescription: "The relative path of the file to map the key to.",

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
																		Description:         "Name of the secret in the pod's namespace to use.",
																		MarkdownDescription: "Name of the secret in the pod's namespace to use.",

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

													"wipe_method": {
														Description:         "WipeMethod determines how volumes attached to Aerospike server pods ar",
														MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods ar",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
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

									"zone": {
										Description:         "Zone name for setting rack affinity.",
										MarkdownDescription: "Zone name for setting rack affinity.",

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

					"seeds_finder_services": {
						Description:         "SeedsFinderServices creates additional Kubernetes service that allow c",
						MarkdownDescription: "SeedsFinderServices creates additional Kubernetes service that allow c",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"load_balancer": {
								Description:         "LoadBalancer created to discover Aerospike Cluster nodes from outside ",
								MarkdownDescription: "LoadBalancer created to discover Aerospike Cluster nodes from outside ",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"external_traffic_policy": {
										Description:         "Service External Traffic Policy Type string",
										MarkdownDescription: "Service External Traffic Policy Type string",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("Local", "Cluster"),
										},
									},

									"load_balancer_source_ranges": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Port Exposed port on load balancer.",
										MarkdownDescription: "Port Exposed port on load balancer.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1024),

											int64validator.AtMost(65535),
										},
									},

									"target_port": {
										Description:         "TargetPort Target port. If not specified the tls-port of network.",
										MarkdownDescription: "TargetPort Target port. If not specified the tls-port of network.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											int64validator.AtLeast(1024),

											int64validator.AtMost(65535),
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

					"size": {
						Description:         "Aerospike cluster size",
						MarkdownDescription: "Aerospike cluster size",

						Type: types.Int64Type,

						Required: true,
						Optional: false,
						Computed: false,
					},

					"storage": {
						Description:         "Storage specify persistent storage to use for the Aerospike pods",
						MarkdownDescription: "Storage specify persistent storage to use for the Aerospike pods",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"block_volume_policy": {
								Description:         "BlockVolumePolicy contains default policies for block volumes.",
								MarkdownDescription: "BlockVolumePolicy contains default policies for block volumes.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cascade_delete": {
										Description:         "CascadeDelete determines if the persistent volumes are deleted after t",
										MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after t",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"effective_cascade_delete": {
										Description:         "Effective/operative value to use for cascade delete after applying def",
										MarkdownDescription: "Effective/operative value to use for cascade delete after applying def",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"effective_init_method": {
										Description:         "Effective/operative value to use as the volume init method after apply",
										MarkdownDescription: "Effective/operative value to use as the volume init method after apply",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"effective_wipe_method": {
										Description:         "Effective/operative value to use as the volume wipe method after apply",
										MarkdownDescription: "Effective/operative value to use as the volume wipe method after apply",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"init_method": {
										Description:         "InitMethod determines how volumes attached to Aerospike server pods ar",
										MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods ar",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"wipe_method": {
										Description:         "WipeMethod determines how volumes attached to Aerospike server pods ar",
										MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods ar",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"filesystem_volume_policy": {
								Description:         "FileSystemVolumePolicy contains default policies for filesystem volume",
								MarkdownDescription: "FileSystemVolumePolicy contains default policies for filesystem volume",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cascade_delete": {
										Description:         "CascadeDelete determines if the persistent volumes are deleted after t",
										MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after t",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"effective_cascade_delete": {
										Description:         "Effective/operative value to use for cascade delete after applying def",
										MarkdownDescription: "Effective/operative value to use for cascade delete after applying def",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"effective_init_method": {
										Description:         "Effective/operative value to use as the volume init method after apply",
										MarkdownDescription: "Effective/operative value to use as the volume init method after apply",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"effective_wipe_method": {
										Description:         "Effective/operative value to use as the volume wipe method after apply",
										MarkdownDescription: "Effective/operative value to use as the volume wipe method after apply",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"init_method": {
										Description:         "InitMethod determines how volumes attached to Aerospike server pods ar",
										MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods ar",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"wipe_method": {
										Description:         "WipeMethod determines how volumes attached to Aerospike server pods ar",
										MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods ar",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volumes": {
								Description:         "Volumes list to attach to created pods.",
								MarkdownDescription: "Volumes list to attach to created pods.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"aerospike": {
										Description:         "Aerospike attachment of this volume on Aerospike server container.",
										MarkdownDescription: "Aerospike attachment of this volume on Aerospike server container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"mount_options": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"mount_propagation": {
														Description:         "mountPropagation determines how mounts are propagated from the host to",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path": {
														Description:         "Path within the volume from which the container's volume should be mou",
														MarkdownDescription: "Path within the volume from which the container's volume should be mou",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path_expr": {
														Description:         "Expanded path within the volume from which the container's volume shou",
														MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

											"path": {
												Description:         "Path to attach the volume on the Aerospike server container.",
												MarkdownDescription: "Path to attach the volume on the Aerospike server container.",

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

									"cascade_delete": {
										Description:         "CascadeDelete determines if the persistent volumes are deleted after t",
										MarkdownDescription: "CascadeDelete determines if the persistent volumes are deleted after t",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"effective_cascade_delete": {
										Description:         "Effective/operative value to use for cascade delete after applying def",
										MarkdownDescription: "Effective/operative value to use for cascade delete after applying def",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"effective_init_method": {
										Description:         "Effective/operative value to use as the volume init method after apply",
										MarkdownDescription: "Effective/operative value to use as the volume init method after apply",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"effective_wipe_method": {
										Description:         "Effective/operative value to use as the volume wipe method after apply",
										MarkdownDescription: "Effective/operative value to use as the volume wipe method after apply",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"init_containers": {
										Description:         "InitContainers are additional init containers where this volume will b",
										MarkdownDescription: "InitContainers are additional init containers where this volume will b",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"container_name": {
												Description:         "ContainerName is the name of the container to attach this volume to.",
												MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"mount_options": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"mount_propagation": {
														Description:         "mountPropagation determines how mounts are propagated from the host to",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path": {
														Description:         "Path within the volume from which the container's volume should be mou",
														MarkdownDescription: "Path within the volume from which the container's volume should be mou",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path_expr": {
														Description:         "Expanded path within the volume from which the container's volume shou",
														MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

											"path": {
												Description:         "Path to attach the volume on the container.",
												MarkdownDescription: "Path to attach the volume on the container.",

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

									"init_method": {
										Description:         "InitMethod determines how volumes attached to Aerospike server pods ar",
										MarkdownDescription: "InitMethod determines how volumes attached to Aerospike server pods ar",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
										},
									},

									"name": {
										Description:         "Name for this volume, Name or path should be given.",
										MarkdownDescription: "Name for this volume, Name or path should be given.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"sidecars": {
										Description:         "Sidecars are side containers where this volume will be mounted",
										MarkdownDescription: "Sidecars are side containers where this volume will be mounted",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"container_name": {
												Description:         "ContainerName is the name of the container to attach this volume to.",
												MarkdownDescription: "ContainerName is the name of the container to attach this volume to.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"mount_options": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"mount_propagation": {
														Description:         "mountPropagation determines how mounts are propagated from the host to",
														MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "Mounted read-only if true, read-write otherwise (false or unspecified)",
														MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified)",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path": {
														Description:         "Path within the volume from which the container's volume should be mou",
														MarkdownDescription: "Path within the volume from which the container's volume should be mou",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sub_path_expr": {
														Description:         "Expanded path within the volume from which the container's volume shou",
														MarkdownDescription: "Expanded path within the volume from which the container's volume shou",

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

											"path": {
												Description:         "Path to attach the volume on the container.",
												MarkdownDescription: "Path to attach the volume on the container.",

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

									"source": {
										Description:         "Source of this volume.",
										MarkdownDescription: "Source of this volume.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap represents a configMap that should populate this volume",
												MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "Optional: mode bits used to set permissions on created files by defaul",
														MarkdownDescription: "Optional: mode bits used to set permissions on created files by defaul",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "If unspecified, each key-value pair in the Data field of the reference",
														MarkdownDescription: "If unspecified, each key-value pair in the Data field of the reference",

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
																Description:         "Optional: mode bits used to set permissions on this file.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "The relative path of the file to map the key to.",
																MarkdownDescription: "The relative path of the file to map the key to.",

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
														Description:         "Name of the referent. More info: https://kubernetes.",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.",

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

											"empty_dir": {
												Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime",
												MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"medium": {
														Description:         "What type of storage medium should back this directory.",
														MarkdownDescription: "What type of storage medium should back this directory.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size_limit": {
														Description:         "Total amount of local storage required for this EmptyDir volume.",
														MarkdownDescription: "Total amount of local storage required for this EmptyDir volume.",

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

											"persistent_volume": {
												Description:         "PersistentVolumeSpec describes a persistent volume to claim and attach",
												MarkdownDescription: "PersistentVolumeSpec describes a persistent volume to claim and attach",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"access_modes": {
														Description:         "Name for creating PVC for this volume, Name or path should be given Na",
														MarkdownDescription: "Name for creating PVC for this volume, Name or path should be given Na",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"metadata": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"annotations": {
																Description:         "Key - Value pair that may be set by external tools to store and retrie",
																MarkdownDescription: "Key - Value pair that may be set by external tools to store and retrie",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"labels": {
																Description:         "Key - Value pairs that can be used to organize and categorize scope an",
																MarkdownDescription: "Key - Value pairs that can be used to organize and categorize scope an",

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
														Description:         "A label query over volumes to consider for binding.",
														MarkdownDescription: "A label query over volumes to consider for binding.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"match_expressions": {
																Description:         "matchExpressions is a list of label selector requirements.",
																MarkdownDescription: "matchExpressions is a list of label selector requirements.",

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
																		Description:         "operator represents a key's relationship to a set of values.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"values": {
																		Description:         "values is an array of string values.",
																		MarkdownDescription: "values is an array of string values.",

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
																Description:         "matchLabels is a map of {key,value} pairs.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs.",

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

													"size": {
														Description:         "Size of volume.",
														MarkdownDescription: "Size of volume.",

														Type: utilities.IntOrStringType{},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"storage_class": {
														Description:         "StorageClass should be pre-created by user.",
														MarkdownDescription: "StorageClass should be pre-created by user.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"volume_mode": {
														Description:         "VolumeMode specifies if the volume is block/raw or a filesystem.",
														MarkdownDescription: "VolumeMode specifies if the volume is block/raw or a filesystem.",

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

											"secret": {
												Description:         "Adapts a Secret into a volume.",
												MarkdownDescription: "Adapts a Secret into a volume.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "Optional: mode bits used to set permissions on created files by defaul",
														MarkdownDescription: "Optional: mode bits used to set permissions on created files by defaul",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "If unspecified, each key-value pair in the Data field of the reference",
														MarkdownDescription: "If unspecified, each key-value pair in the Data field of the reference",

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
																Description:         "Optional: mode bits used to set permissions on this file.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "The relative path of the file to map the key to.",
																MarkdownDescription: "The relative path of the file to map the key to.",

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
														Description:         "Name of the secret in the pod's namespace to use.",
														MarkdownDescription: "Name of the secret in the pod's namespace to use.",

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

									"wipe_method": {
										Description:         "WipeMethod determines how volumes attached to Aerospike server pods ar",
										MarkdownDescription: "WipeMethod determines how volumes attached to Aerospike server pods ar",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("none", "dd", "blkdiscard", "deleteFiles"),
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

					"validation_policy": {
						Description:         "ValidationPolicy controls validation of the Aerospike cluster resource",
						MarkdownDescription: "ValidationPolicy controls validation of the Aerospike cluster resource",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"skip_work_dir_validate": {
								Description:         "skipWorkDirValidate validates that Aerospike work directory is mounted",
								MarkdownDescription: "skipWorkDirValidate validates that Aerospike work directory is mounted",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"skip_xdr_dlog_file_validate": {
								Description:         "ValidateXdrDigestLogFile validates that xdr digest log file is mounted",
								MarkdownDescription: "ValidateXdrDigestLogFile validates that xdr digest log file is mounted",

								Type: types.BoolType,

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
		},
	}, nil
}

func (r *AsdbAerospikeComAerospikeClusterV1Beta1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_asdb_aerospike_com_aerospike_cluster_v1beta1")

	var state AsdbAerospikeComAerospikeClusterV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AsdbAerospikeComAerospikeClusterV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("asdb.aerospike.com/v1beta1")
	goModel.Kind = utilities.Ptr("AerospikeCluster")

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

func (r *AsdbAerospikeComAerospikeClusterV1Beta1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_asdb_aerospike_com_aerospike_cluster_v1beta1")
	// NO-OP: All data is already in Terraform state
}

func (r *AsdbAerospikeComAerospikeClusterV1Beta1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_asdb_aerospike_com_aerospike_cluster_v1beta1")

	var state AsdbAerospikeComAerospikeClusterV1Beta1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AsdbAerospikeComAerospikeClusterV1Beta1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("asdb.aerospike.com/v1beta1")
	goModel.Kind = utilities.Ptr("AerospikeCluster")

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

func (r *AsdbAerospikeComAerospikeClusterV1Beta1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_asdb_aerospike_com_aerospike_cluster_v1beta1")
	// NO-OP: Terraform removes the state automatically for us
}
