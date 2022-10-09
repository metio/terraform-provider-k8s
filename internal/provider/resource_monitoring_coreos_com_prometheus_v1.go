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

type MonitoringCoreosComPrometheusV1Resource struct{}

var (
	_ resource.Resource = (*MonitoringCoreosComPrometheusV1Resource)(nil)
)

type MonitoringCoreosComPrometheusV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type MonitoringCoreosComPrometheusV1GoModel struct {
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
		AdditionalAlertManagerConfigs *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
		} `tfsdk:"additional_alert_manager_configs" yaml:"additionalAlertManagerConfigs,omitempty"`

		AdditionalAlertRelabelConfigs *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
		} `tfsdk:"additional_alert_relabel_configs" yaml:"additionalAlertRelabelConfigs,omitempty"`

		AdditionalArgs *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Value *string `tfsdk:"value" yaml:"value,omitempty"`
		} `tfsdk:"additional_args" yaml:"additionalArgs,omitempty"`

		AdditionalScrapeConfigs *struct {
			Key *string `tfsdk:"key" yaml:"key,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
		} `tfsdk:"additional_scrape_configs" yaml:"additionalScrapeConfigs,omitempty"`

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

		Alerting *struct {
			Alertmanagers *[]struct {
				ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

				Authorization *struct {
					Credentials *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"credentials" yaml:"credentials,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"authorization" yaml:"authorization,omitempty"`

				BearerTokenFile *string `tfsdk:"bearer_token_file" yaml:"bearerTokenFile,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				PathPrefix *string `tfsdk:"path_prefix" yaml:"pathPrefix,omitempty"`

				Port *string `tfsdk:"port" yaml:"port,omitempty"`

				Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`

				Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`

				TlsConfig *struct {
					Ca *struct {
						ConfigMap *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"config_map" yaml:"configMap,omitempty"`

						Secret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"secret" yaml:"secret,omitempty"`
					} `tfsdk:"ca" yaml:"ca,omitempty"`

					CaFile *string `tfsdk:"ca_file" yaml:"caFile,omitempty"`

					Cert *struct {
						ConfigMap *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"config_map" yaml:"configMap,omitempty"`

						Secret *struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Name *string `tfsdk:"name" yaml:"name,omitempty"`

							Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
						} `tfsdk:"secret" yaml:"secret,omitempty"`
					} `tfsdk:"cert" yaml:"cert,omitempty"`

					CertFile *string `tfsdk:"cert_file" yaml:"certFile,omitempty"`

					InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

					KeyFile *string `tfsdk:"key_file" yaml:"keyFile,omitempty"`

					KeySecret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

					ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
				} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
			} `tfsdk:"alertmanagers" yaml:"alertmanagers,omitempty"`
		} `tfsdk:"alerting" yaml:"alerting,omitempty"`

		AllowOverlappingBlocks *bool `tfsdk:"allow_overlapping_blocks" yaml:"allowOverlappingBlocks,omitempty"`

		ApiserverConfig *struct {
			Authorization *struct {
				Credentials *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"credentials" yaml:"credentials,omitempty"`

				CredentialsFile *string `tfsdk:"credentials_file" yaml:"credentialsFile,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"authorization" yaml:"authorization,omitempty"`

			BasicAuth *struct {
				Password *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"password" yaml:"password,omitempty"`

				Username *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"username" yaml:"username,omitempty"`
			} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

			BearerToken *string `tfsdk:"bearer_token" yaml:"bearerToken,omitempty"`

			BearerTokenFile *string `tfsdk:"bearer_token_file" yaml:"bearerTokenFile,omitempty"`

			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			TlsConfig *struct {
				Ca *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				CaFile *string `tfsdk:"ca_file" yaml:"caFile,omitempty"`

				Cert *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"cert" yaml:"cert,omitempty"`

				CertFile *string `tfsdk:"cert_file" yaml:"certFile,omitempty"`

				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

				KeyFile *string `tfsdk:"key_file" yaml:"keyFile,omitempty"`

				KeySecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

				ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
			} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
		} `tfsdk:"apiserver_config" yaml:"apiserverConfig,omitempty"`

		ArbitraryFSAccessThroughSMs *struct {
			Deny *bool `tfsdk:"deny" yaml:"deny,omitempty"`
		} `tfsdk:"arbitrary_fs_access_through_s_ms" yaml:"arbitraryFSAccessThroughSMs,omitempty"`

		BaseImage *string `tfsdk:"base_image" yaml:"baseImage,omitempty"`

		ConfigMaps *[]string `tfsdk:"config_maps" yaml:"configMaps,omitempty"`

		Containers *[]struct {
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

						Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

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

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
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

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
				} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
			} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

			LivenessProbe *struct {
				Exec *struct {
					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
				} `tfsdk:"exec" yaml:"exec,omitempty"`

				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`
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

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`
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

					HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

					RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
			} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

			StartupProbe *struct {
				Exec *struct {
					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
				} `tfsdk:"exec" yaml:"exec,omitempty"`

				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`
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
		} `tfsdk:"containers" yaml:"containers,omitempty"`

		DisableCompaction *bool `tfsdk:"disable_compaction" yaml:"disableCompaction,omitempty"`

		EnableAdminAPI *bool `tfsdk:"enable_admin_api" yaml:"enableAdminAPI,omitempty"`

		EnableFeatures *[]string `tfsdk:"enable_features" yaml:"enableFeatures,omitempty"`

		EnableRemoteWriteReceiver *bool `tfsdk:"enable_remote_write_receiver" yaml:"enableRemoteWriteReceiver,omitempty"`

		EnforcedBodySizeLimit *string `tfsdk:"enforced_body_size_limit" yaml:"enforcedBodySizeLimit,omitempty"`

		EnforcedLabelLimit *int64 `tfsdk:"enforced_label_limit" yaml:"enforcedLabelLimit,omitempty"`

		EnforcedLabelNameLengthLimit *int64 `tfsdk:"enforced_label_name_length_limit" yaml:"enforcedLabelNameLengthLimit,omitempty"`

		EnforcedLabelValueLengthLimit *int64 `tfsdk:"enforced_label_value_length_limit" yaml:"enforcedLabelValueLengthLimit,omitempty"`

		EnforcedNamespaceLabel *string `tfsdk:"enforced_namespace_label" yaml:"enforcedNamespaceLabel,omitempty"`

		EnforcedSampleLimit *int64 `tfsdk:"enforced_sample_limit" yaml:"enforcedSampleLimit,omitempty"`

		EnforcedTargetLimit *int64 `tfsdk:"enforced_target_limit" yaml:"enforcedTargetLimit,omitempty"`

		EvaluationInterval *string `tfsdk:"evaluation_interval" yaml:"evaluationInterval,omitempty"`

		ExcludedFromEnforcement *[]struct {
			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
		} `tfsdk:"excluded_from_enforcement" yaml:"excludedFromEnforcement,omitempty"`

		Exemplars *struct {
			MaxSize *int64 `tfsdk:"max_size" yaml:"maxSize,omitempty"`
		} `tfsdk:"exemplars" yaml:"exemplars,omitempty"`

		ExternalLabels *map[string]string `tfsdk:"external_labels" yaml:"externalLabels,omitempty"`

		ExternalUrl *string `tfsdk:"external_url" yaml:"externalUrl,omitempty"`

		HostAliases *[]struct {
			Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

			Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
		} `tfsdk:"host_aliases" yaml:"hostAliases,omitempty"`

		IgnoreNamespaceSelectors *bool `tfsdk:"ignore_namespace_selectors" yaml:"ignoreNamespaceSelectors,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

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

						Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

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

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
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

						Port *string `tfsdk:"port" yaml:"port,omitempty"`

						Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
					} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

					TcpSocket *struct {
						Host *string `tfsdk:"host" yaml:"host,omitempty"`

						Port *string `tfsdk:"port" yaml:"port,omitempty"`
					} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`
				} `tfsdk:"pre_stop" yaml:"preStop,omitempty"`
			} `tfsdk:"lifecycle" yaml:"lifecycle,omitempty"`

			LivenessProbe *struct {
				Exec *struct {
					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
				} `tfsdk:"exec" yaml:"exec,omitempty"`

				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`
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

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`
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

					HostProcess *bool `tfsdk:"host_process" yaml:"hostProcess,omitempty"`

					RunAsUserName *string `tfsdk:"run_as_user_name" yaml:"runAsUserName,omitempty"`
				} `tfsdk:"windows_options" yaml:"windowsOptions,omitempty"`
			} `tfsdk:"security_context" yaml:"securityContext,omitempty"`

			StartupProbe *struct {
				Exec *struct {
					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
				} `tfsdk:"exec" yaml:"exec,omitempty"`

				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port *string `tfsdk:"port" yaml:"port,omitempty"`
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

		ListenLocal *bool `tfsdk:"listen_local" yaml:"listenLocal,omitempty"`

		LogFormat *string `tfsdk:"log_format" yaml:"logFormat,omitempty"`

		LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

		MinReadySeconds *int64 `tfsdk:"min_ready_seconds" yaml:"minReadySeconds,omitempty"`

		NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

		OverrideHonorLabels *bool `tfsdk:"override_honor_labels" yaml:"overrideHonorLabels,omitempty"`

		OverrideHonorTimestamps *bool `tfsdk:"override_honor_timestamps" yaml:"overrideHonorTimestamps,omitempty"`

		Paused *bool `tfsdk:"paused" yaml:"paused,omitempty"`

		PodMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"pod_metadata" yaml:"podMetadata,omitempty"`

		PodMonitorNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"pod_monitor_namespace_selector" yaml:"podMonitorNamespaceSelector,omitempty"`

		PodMonitorSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"pod_monitor_selector" yaml:"podMonitorSelector,omitempty"`

		PortName *string `tfsdk:"port_name" yaml:"portName,omitempty"`

		PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

		ProbeNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"probe_namespace_selector" yaml:"probeNamespaceSelector,omitempty"`

		ProbeSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"probe_selector" yaml:"probeSelector,omitempty"`

		PrometheusExternalLabelName *string `tfsdk:"prometheus_external_label_name" yaml:"prometheusExternalLabelName,omitempty"`

		PrometheusRulesExcludedFromEnforce *[]struct {
			RuleName *string `tfsdk:"rule_name" yaml:"ruleName,omitempty"`

			RuleNamespace *string `tfsdk:"rule_namespace" yaml:"ruleNamespace,omitempty"`
		} `tfsdk:"prometheus_rules_excluded_from_enforce" yaml:"prometheusRulesExcludedFromEnforce,omitempty"`

		Query *struct {
			LookbackDelta *string `tfsdk:"lookback_delta" yaml:"lookbackDelta,omitempty"`

			MaxConcurrency *int64 `tfsdk:"max_concurrency" yaml:"maxConcurrency,omitempty"`

			MaxSamples *int64 `tfsdk:"max_samples" yaml:"maxSamples,omitempty"`

			Timeout *string `tfsdk:"timeout" yaml:"timeout,omitempty"`
		} `tfsdk:"query" yaml:"query,omitempty"`

		QueryLogFile *string `tfsdk:"query_log_file" yaml:"queryLogFile,omitempty"`

		RemoteRead *[]struct {
			Authorization *struct {
				Credentials *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"credentials" yaml:"credentials,omitempty"`

				CredentialsFile *string `tfsdk:"credentials_file" yaml:"credentialsFile,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"authorization" yaml:"authorization,omitempty"`

			BasicAuth *struct {
				Password *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"password" yaml:"password,omitempty"`

				Username *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"username" yaml:"username,omitempty"`
			} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

			BearerToken *string `tfsdk:"bearer_token" yaml:"bearerToken,omitempty"`

			BearerTokenFile *string `tfsdk:"bearer_token_file" yaml:"bearerTokenFile,omitempty"`

			Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Oauth2 *struct {
				ClientId *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"client_id" yaml:"clientId,omitempty"`

				ClientSecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

				EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

				Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

				TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
			} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

			ProxyUrl *string `tfsdk:"proxy_url" yaml:"proxyUrl,omitempty"`

			ReadRecent *bool `tfsdk:"read_recent" yaml:"readRecent,omitempty"`

			RemoteTimeout *string `tfsdk:"remote_timeout" yaml:"remoteTimeout,omitempty"`

			RequiredMatchers *map[string]string `tfsdk:"required_matchers" yaml:"requiredMatchers,omitempty"`

			TlsConfig *struct {
				Ca *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				CaFile *string `tfsdk:"ca_file" yaml:"caFile,omitempty"`

				Cert *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"cert" yaml:"cert,omitempty"`

				CertFile *string `tfsdk:"cert_file" yaml:"certFile,omitempty"`

				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

				KeyFile *string `tfsdk:"key_file" yaml:"keyFile,omitempty"`

				KeySecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

				ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
			} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`
		} `tfsdk:"remote_read" yaml:"remoteRead,omitempty"`

		RemoteWrite *[]struct {
			Authorization *struct {
				Credentials *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"credentials" yaml:"credentials,omitempty"`

				CredentialsFile *string `tfsdk:"credentials_file" yaml:"credentialsFile,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"authorization" yaml:"authorization,omitempty"`

			BasicAuth *struct {
				Password *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"password" yaml:"password,omitempty"`

				Username *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"username" yaml:"username,omitempty"`
			} `tfsdk:"basic_auth" yaml:"basicAuth,omitempty"`

			BearerToken *string `tfsdk:"bearer_token" yaml:"bearerToken,omitempty"`

			BearerTokenFile *string `tfsdk:"bearer_token_file" yaml:"bearerTokenFile,omitempty"`

			Headers *map[string]string `tfsdk:"headers" yaml:"headers,omitempty"`

			MetadataConfig *struct {
				Send *bool `tfsdk:"send" yaml:"send,omitempty"`

				SendInterval *string `tfsdk:"send_interval" yaml:"sendInterval,omitempty"`
			} `tfsdk:"metadata_config" yaml:"metadataConfig,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Oauth2 *struct {
				ClientId *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"client_id" yaml:"clientId,omitempty"`

				ClientSecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"client_secret" yaml:"clientSecret,omitempty"`

				EndpointParams *map[string]string `tfsdk:"endpoint_params" yaml:"endpointParams,omitempty"`

				Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

				TokenUrl *string `tfsdk:"token_url" yaml:"tokenUrl,omitempty"`
			} `tfsdk:"oauth2" yaml:"oauth2,omitempty"`

			ProxyUrl *string `tfsdk:"proxy_url" yaml:"proxyUrl,omitempty"`

			QueueConfig *struct {
				BatchSendDeadline *string `tfsdk:"batch_send_deadline" yaml:"batchSendDeadline,omitempty"`

				Capacity *int64 `tfsdk:"capacity" yaml:"capacity,omitempty"`

				MaxBackoff *string `tfsdk:"max_backoff" yaml:"maxBackoff,omitempty"`

				MaxRetries *int64 `tfsdk:"max_retries" yaml:"maxRetries,omitempty"`

				MaxSamplesPerSend *int64 `tfsdk:"max_samples_per_send" yaml:"maxSamplesPerSend,omitempty"`

				MaxShards *int64 `tfsdk:"max_shards" yaml:"maxShards,omitempty"`

				MinBackoff *string `tfsdk:"min_backoff" yaml:"minBackoff,omitempty"`

				MinShards *int64 `tfsdk:"min_shards" yaml:"minShards,omitempty"`

				RetryOnRateLimit *bool `tfsdk:"retry_on_rate_limit" yaml:"retryOnRateLimit,omitempty"`
			} `tfsdk:"queue_config" yaml:"queueConfig,omitempty"`

			RemoteTimeout *string `tfsdk:"remote_timeout" yaml:"remoteTimeout,omitempty"`

			SendExemplars *bool `tfsdk:"send_exemplars" yaml:"sendExemplars,omitempty"`

			Sigv4 *struct {
				AccessKey *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"access_key" yaml:"accessKey,omitempty"`

				Profile *string `tfsdk:"profile" yaml:"profile,omitempty"`

				Region *string `tfsdk:"region" yaml:"region,omitempty"`

				RoleArn *string `tfsdk:"role_arn" yaml:"roleArn,omitempty"`

				SecretKey *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"secret_key" yaml:"secretKey,omitempty"`
			} `tfsdk:"sigv4" yaml:"sigv4,omitempty"`

			TlsConfig *struct {
				Ca *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				CaFile *string `tfsdk:"ca_file" yaml:"caFile,omitempty"`

				Cert *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"cert" yaml:"cert,omitempty"`

				CertFile *string `tfsdk:"cert_file" yaml:"certFile,omitempty"`

				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

				KeyFile *string `tfsdk:"key_file" yaml:"keyFile,omitempty"`

				KeySecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

				ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
			} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`

			WriteRelabelConfigs *[]struct {
				Action *string `tfsdk:"action" yaml:"action,omitempty"`

				Modulus *int64 `tfsdk:"modulus" yaml:"modulus,omitempty"`

				Regex *string `tfsdk:"regex" yaml:"regex,omitempty"`

				Replacement *string `tfsdk:"replacement" yaml:"replacement,omitempty"`

				Separator *string `tfsdk:"separator" yaml:"separator,omitempty"`

				SourceLabels *[]string `tfsdk:"source_labels" yaml:"sourceLabels,omitempty"`

				TargetLabel *string `tfsdk:"target_label" yaml:"targetLabel,omitempty"`
			} `tfsdk:"write_relabel_configs" yaml:"writeRelabelConfigs,omitempty"`
		} `tfsdk:"remote_write" yaml:"remoteWrite,omitempty"`

		ReplicaExternalLabelName *string `tfsdk:"replica_external_label_name" yaml:"replicaExternalLabelName,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

		Resources *struct {
			Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

			Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
		} `tfsdk:"resources" yaml:"resources,omitempty"`

		Retention *string `tfsdk:"retention" yaml:"retention,omitempty"`

		RetentionSize *string `tfsdk:"retention_size" yaml:"retentionSize,omitempty"`

		RoutePrefix *string `tfsdk:"route_prefix" yaml:"routePrefix,omitempty"`

		RuleNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"rule_namespace_selector" yaml:"ruleNamespaceSelector,omitempty"`

		RuleSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"rule_selector" yaml:"ruleSelector,omitempty"`

		Rules *struct {
			Alert *struct {
				ForGracePeriod *string `tfsdk:"for_grace_period" yaml:"forGracePeriod,omitempty"`

				ForOutageTolerance *string `tfsdk:"for_outage_tolerance" yaml:"forOutageTolerance,omitempty"`

				ResendDelay *string `tfsdk:"resend_delay" yaml:"resendDelay,omitempty"`
			} `tfsdk:"alert" yaml:"alert,omitempty"`
		} `tfsdk:"rules" yaml:"rules,omitempty"`

		ScrapeInterval *string `tfsdk:"scrape_interval" yaml:"scrapeInterval,omitempty"`

		ScrapeTimeout *string `tfsdk:"scrape_timeout" yaml:"scrapeTimeout,omitempty"`

		Secrets *[]string `tfsdk:"secrets" yaml:"secrets,omitempty"`

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

		ServiceAccountName *string `tfsdk:"service_account_name" yaml:"serviceAccountName,omitempty"`

		ServiceMonitorNamespaceSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"service_monitor_namespace_selector" yaml:"serviceMonitorNamespaceSelector,omitempty"`

		ServiceMonitorSelector *struct {
			MatchExpressions *[]struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
			} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

			MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
		} `tfsdk:"service_monitor_selector" yaml:"serviceMonitorSelector,omitempty"`

		Sha *string `tfsdk:"sha" yaml:"sha,omitempty"`

		Shards *int64 `tfsdk:"shards" yaml:"shards,omitempty"`

		Storage *struct {
			DisableMountSubPath *bool `tfsdk:"disable_mount_sub_path" yaml:"disableMountSubPath,omitempty"`

			EmptyDir *struct {
				Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

				SizeLimit *string `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
			} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

			Ephemeral *struct {
				VolumeClaimTemplate *struct {
					Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

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

			VolumeClaimTemplate *struct {
				ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Metadata *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
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
			} `tfsdk:"volume_claim_template" yaml:"volumeClaimTemplate,omitempty"`
		} `tfsdk:"storage" yaml:"storage,omitempty"`

		Tag *string `tfsdk:"tag" yaml:"tag,omitempty"`

		Thanos *struct {
			AdditionalArgs *[]struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"additional_args" yaml:"additionalArgs,omitempty"`

			BaseImage *string `tfsdk:"base_image" yaml:"baseImage,omitempty"`

			GrpcServerTlsConfig *struct {
				Ca *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"ca" yaml:"ca,omitempty"`

				CaFile *string `tfsdk:"ca_file" yaml:"caFile,omitempty"`

				Cert *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"cert" yaml:"cert,omitempty"`

				CertFile *string `tfsdk:"cert_file" yaml:"certFile,omitempty"`

				InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

				KeyFile *string `tfsdk:"key_file" yaml:"keyFile,omitempty"`

				KeySecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

				ServerName *string `tfsdk:"server_name" yaml:"serverName,omitempty"`
			} `tfsdk:"grpc_server_tls_config" yaml:"grpcServerTlsConfig,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			ListenLocal *bool `tfsdk:"listen_local" yaml:"listenLocal,omitempty"`

			LogFormat *string `tfsdk:"log_format" yaml:"logFormat,omitempty"`

			LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

			MinTime *string `tfsdk:"min_time" yaml:"minTime,omitempty"`

			ObjectStorageConfig *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"object_storage_config" yaml:"objectStorageConfig,omitempty"`

			ObjectStorageConfigFile *string `tfsdk:"object_storage_config_file" yaml:"objectStorageConfigFile,omitempty"`

			ReadyTimeout *string `tfsdk:"ready_timeout" yaml:"readyTimeout,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Sha *string `tfsdk:"sha" yaml:"sha,omitempty"`

			Tag *string `tfsdk:"tag" yaml:"tag,omitempty"`

			TracingConfig *struct {
				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"tracing_config" yaml:"tracingConfig,omitempty"`

			TracingConfigFile *string `tfsdk:"tracing_config_file" yaml:"tracingConfigFile,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`

			VolumeMounts *[]struct {
				MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

				MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

				SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`
		} `tfsdk:"thanos" yaml:"thanos,omitempty"`

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

				MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

			MatchLabelKeys *[]string `tfsdk:"match_label_keys" yaml:"matchLabelKeys,omitempty"`

			MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

			MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

			NodeAffinityPolicy *string `tfsdk:"node_affinity_policy" yaml:"nodeAffinityPolicy,omitempty"`

			NodeTaintsPolicy *string `tfsdk:"node_taints_policy" yaml:"nodeTaintsPolicy,omitempty"`

			TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

			WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
		} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`

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

						Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
					} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`
				} `tfsdk:"items" yaml:"items,omitempty"`
			} `tfsdk:"downward_api" yaml:"downwardAPI,omitempty"`

			EmptyDir *struct {
				Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

				SizeLimit *string `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
			} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

			Ephemeral *struct {
				VolumeClaimTemplate *struct {
					Metadata *map[string]string `tfsdk:"metadata" yaml:"metadata,omitempty"`

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

								Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

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

		WalCompression *bool `tfsdk:"wal_compression" yaml:"walCompression,omitempty"`

		Web *struct {
			HttpConfig *struct {
				Headers *struct {
					ContentSecurityPolicy *string `tfsdk:"content_security_policy" yaml:"contentSecurityPolicy,omitempty"`

					StrictTransportSecurity *string `tfsdk:"strict_transport_security" yaml:"strictTransportSecurity,omitempty"`

					XContentTypeOptions *string `tfsdk:"x_content_type_options" yaml:"xContentTypeOptions,omitempty"`

					XFrameOptions *string `tfsdk:"x_frame_options" yaml:"xFrameOptions,omitempty"`

					XXSSProtection *string `tfsdk:"x_xss_protection" yaml:"xXSSProtection,omitempty"`
				} `tfsdk:"headers" yaml:"headers,omitempty"`

				Http2 *bool `tfsdk:"http2" yaml:"http2,omitempty"`
			} `tfsdk:"http_config" yaml:"httpConfig,omitempty"`

			PageTitle *string `tfsdk:"page_title" yaml:"pageTitle,omitempty"`

			TlsConfig *struct {
				Cert *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"cert" yaml:"cert,omitempty"`

				CipherSuites *[]string `tfsdk:"cipher_suites" yaml:"cipherSuites,omitempty"`

				ClientAuthType *string `tfsdk:"client_auth_type" yaml:"clientAuthType,omitempty"`

				Client_ca *struct {
					ConfigMap *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"config_map" yaml:"configMap,omitempty"`

					Secret *struct {
						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
					} `tfsdk:"secret" yaml:"secret,omitempty"`
				} `tfsdk:"client_ca" yaml:"client_ca,omitempty"`

				CurvePreferences *[]string `tfsdk:"curve_preferences" yaml:"curvePreferences,omitempty"`

				KeySecret *struct {
					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
				} `tfsdk:"key_secret" yaml:"keySecret,omitempty"`

				MaxVersion *string `tfsdk:"max_version" yaml:"maxVersion,omitempty"`

				MinVersion *string `tfsdk:"min_version" yaml:"minVersion,omitempty"`

				PreferServerCipherSuites *bool `tfsdk:"prefer_server_cipher_suites" yaml:"preferServerCipherSuites,omitempty"`
			} `tfsdk:"tls_config" yaml:"tlsConfig,omitempty"`
		} `tfsdk:"web" yaml:"web,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewMonitoringCoreosComPrometheusV1Resource() resource.Resource {
	return &MonitoringCoreosComPrometheusV1Resource{}
}

func (r *MonitoringCoreosComPrometheusV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_monitoring_coreos_com_prometheus_v1"
}

func (r *MonitoringCoreosComPrometheusV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "Prometheus defines a Prometheus deployment.",
		MarkdownDescription: "Prometheus defines a Prometheus deployment.",
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
				Description:         "Specification of the desired behavior of the Prometheus cluster. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				MarkdownDescription: "Specification of the desired behavior of the Prometheus cluster. More info: https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"additional_alert_manager_configs": {
						Description:         "AdditionalAlertManagerConfigs allows specifying a key of a Secret containing additional Prometheus AlertManager configurations. AlertManager configurations specified are appended to the configurations generated by the Prometheus Operator. Job configurations specified must have the form as specified in the official Prometheus documentation: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#alertmanager_config. As AlertManager configs are appended, the user is responsible to make sure it is valid. Note that using this feature may expose the possibility to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible AlertManager configs are going to break Prometheus after the upgrade.",
						MarkdownDescription: "AdditionalAlertManagerConfigs allows specifying a key of a Secret containing additional Prometheus AlertManager configurations. AlertManager configurations specified are appended to the configurations generated by the Prometheus Operator. Job configurations specified must have the form as specified in the official Prometheus documentation: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#alertmanager_config. As AlertManager configs are appended, the user is responsible to make sure it is valid. Note that using this feature may expose the possibility to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible AlertManager configs are going to break Prometheus after the upgrade.",

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

					"additional_alert_relabel_configs": {
						Description:         "AdditionalAlertRelabelConfigs allows specifying a key of a Secret containing additional Prometheus alert relabel configurations. Alert relabel configurations specified are appended to the configurations generated by the Prometheus Operator. Alert relabel configurations specified must have the form as specified in the official Prometheus documentation: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#alert_relabel_configs. As alert relabel configs are appended, the user is responsible to make sure it is valid. Note that using this feature may expose the possibility to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible alert relabel configs are going to break Prometheus after the upgrade.",
						MarkdownDescription: "AdditionalAlertRelabelConfigs allows specifying a key of a Secret containing additional Prometheus alert relabel configurations. Alert relabel configurations specified are appended to the configurations generated by the Prometheus Operator. Alert relabel configurations specified must have the form as specified in the official Prometheus documentation: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#alert_relabel_configs. As alert relabel configs are appended, the user is responsible to make sure it is valid. Note that using this feature may expose the possibility to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible alert relabel configs are going to break Prometheus after the upgrade.",

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

					"additional_args": {
						Description:         "AdditionalArgs allows setting additional arguments for the Prometheus container. It is intended for e.g. activating hidden flags which are not supported by the dedicated configuration options yet. The arguments are passed as-is to the Prometheus container which may cause issues if they are invalid or not supporeted by the given Prometheus version. In case of an argument conflict (e.g. an argument which is already set by the operator itself) or when providing an invalid argument the reconciliation will fail and an error will be logged.",
						MarkdownDescription: "AdditionalArgs allows setting additional arguments for the Prometheus container. It is intended for e.g. activating hidden flags which are not supported by the dedicated configuration options yet. The arguments are passed as-is to the Prometheus container which may cause issues if they are invalid or not supporeted by the given Prometheus version. In case of an argument conflict (e.g. an argument which is already set by the operator itself) or when providing an invalid argument the reconciliation will fail and an error will be logged.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of the argument, e.g. 'scrape.discovery-reload-interval'.",
								MarkdownDescription: "Name of the argument, e.g. 'scrape.discovery-reload-interval'.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"value": {
								Description:         "Argument value, e.g. 30s. Can be empty for name-only arguments (e.g. --storage.tsdb.no-lockfile)",
								MarkdownDescription: "Argument value, e.g. 30s. Can be empty for name-only arguments (e.g. --storage.tsdb.no-lockfile)",

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

					"additional_scrape_configs": {
						Description:         "AdditionalScrapeConfigs allows specifying a key of a Secret containing additional Prometheus scrape configurations. Scrape configurations specified are appended to the configurations generated by the Prometheus Operator. Job configurations specified must have the form as specified in the official Prometheus documentation: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config. As scrape configs are appended, the user is responsible to make sure it is valid. Note that using this feature may expose the possibility to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible scrape configs are going to break Prometheus after the upgrade.",
						MarkdownDescription: "AdditionalScrapeConfigs allows specifying a key of a Secret containing additional Prometheus scrape configurations. Scrape configurations specified are appended to the configurations generated by the Prometheus Operator. Job configurations specified must have the form as specified in the official Prometheus documentation: https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config. As scrape configs are appended, the user is responsible to make sure it is valid. Note that using this feature may expose the possibility to break upgrades of Prometheus. It is advised to review Prometheus release notes to ensure that no incompatible scrape configs are going to break Prometheus after the upgrade.",

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

					"affinity": {
						Description:         "If specified, the pod's scheduling constraints.",
						MarkdownDescription: "If specified, the pod's scheduling constraints.",

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

					"alerting": {
						Description:         "Define details regarding alerting.",
						MarkdownDescription: "Define details regarding alerting.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"alertmanagers": {
								Description:         "AlertmanagerEndpoints Prometheus should fire alerts against.",
								MarkdownDescription: "AlertmanagerEndpoints Prometheus should fire alerts against.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"api_version": {
										Description:         "Version of the Alertmanager API that Prometheus uses to send alerts. It can be 'v1' or 'v2'.",
										MarkdownDescription: "Version of the Alertmanager API that Prometheus uses to send alerts. It can be 'v1' or 'v2'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorization": {
										Description:         "Authorization section for this alertmanager endpoint",
										MarkdownDescription: "Authorization section for this alertmanager endpoint",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"credentials": {
												Description:         "The secret's key that contains the credentials of the request",
												MarkdownDescription: "The secret's key that contains the credentials of the request",

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

											"type": {
												Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
												MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

									"bearer_token_file": {
										Description:         "BearerTokenFile to read from filesystem to use when authenticating to Alertmanager.",
										MarkdownDescription: "BearerTokenFile to read from filesystem to use when authenticating to Alertmanager.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of Endpoints object in Namespace.",
										MarkdownDescription: "Name of Endpoints object in Namespace.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace of Endpoints object.",
										MarkdownDescription: "Namespace of Endpoints object.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"path_prefix": {
										Description:         "Prefix for the HTTP path alerts are pushed to.",
										MarkdownDescription: "Prefix for the HTTP path alerts are pushed to.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": {
										Description:         "Port the Alertmanager API is exposed on.",
										MarkdownDescription: "Port the Alertmanager API is exposed on.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"scheme": {
										Description:         "Scheme to use when firing alerts.",
										MarkdownDescription: "Scheme to use when firing alerts.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout": {
										Description:         "Timeout is a per-target Alertmanager timeout when pushing alerts.",
										MarkdownDescription: "Timeout is a per-target Alertmanager timeout when pushing alerts.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls_config": {
										Description:         "TLS Config to use for alertmanager connection.",
										MarkdownDescription: "TLS Config to use for alertmanager connection.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca": {
												Description:         "Struct containing the CA cert to use for the targets.",
												MarkdownDescription: "Struct containing the CA cert to use for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config_map": {
														Description:         "ConfigMap containing data to use for the targets.",
														MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

													"secret": {
														Description:         "Secret containing data to use for the targets.",
														MarkdownDescription: "Secret containing data to use for the targets.",

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

											"ca_file": {
												Description:         "Path to the CA cert in the Prometheus container to use for the targets.",
												MarkdownDescription: "Path to the CA cert in the Prometheus container to use for the targets.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cert": {
												Description:         "Struct containing the client cert file for the targets.",
												MarkdownDescription: "Struct containing the client cert file for the targets.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"config_map": {
														Description:         "ConfigMap containing data to use for the targets.",
														MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

													"secret": {
														Description:         "Secret containing data to use for the targets.",
														MarkdownDescription: "Secret containing data to use for the targets.",

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

											"cert_file": {
												Description:         "Path to the client cert file in the Prometheus container for the targets.",
												MarkdownDescription: "Path to the client cert file in the Prometheus container for the targets.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_skip_verify": {
												Description:         "Disable target certificate validation.",
												MarkdownDescription: "Disable target certificate validation.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key_file": {
												Description:         "Path to the client key file in the Prometheus container for the targets.",
												MarkdownDescription: "Path to the client key file in the Prometheus container for the targets.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key_secret": {
												Description:         "Secret containing the client key file for the targets.",
												MarkdownDescription: "Secret containing the client key file for the targets.",

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

											"server_name": {
												Description:         "Used to verify the hostname for the targets.",
												MarkdownDescription: "Used to verify the hostname for the targets.",

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

								Required: true,
								Optional: false,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"allow_overlapping_blocks": {
						Description:         "AllowOverlappingBlocks enables vertical compaction and vertical query merge in Prometheus. This is still experimental in Prometheus so it may change in any upcoming release.",
						MarkdownDescription: "AllowOverlappingBlocks enables vertical compaction and vertical query merge in Prometheus. This is still experimental in Prometheus so it may change in any upcoming release.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"apiserver_config": {
						Description:         "APIServerConfig allows specifying a host and auth methods to access apiserver. If left empty, Prometheus is assumed to run inside of the cluster and will discover API servers automatically and use the pod's CA certificate and bearer token file at /var/run/secrets/kubernetes.io/serviceaccount/.",
						MarkdownDescription: "APIServerConfig allows specifying a host and auth methods to access apiserver. If left empty, Prometheus is assumed to run inside of the cluster and will discover API servers automatically and use the pod's CA certificate and bearer token file at /var/run/secrets/kubernetes.io/serviceaccount/.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"authorization": {
								Description:         "Authorization section for accessing apiserver",
								MarkdownDescription: "Authorization section for accessing apiserver",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials": {
										Description:         "The secret's key that contains the credentials of the request",
										MarkdownDescription: "The secret's key that contains the credentials of the request",

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

									"credentials_file": {
										Description:         "File to read a secret from, mutually exclusive with Credentials (from SafeAuthorization)",
										MarkdownDescription: "File to read a secret from, mutually exclusive with Credentials (from SafeAuthorization)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
										MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

							"basic_auth": {
								Description:         "BasicAuth allow an endpoint to authenticate over basic authentication",
								MarkdownDescription: "BasicAuth allow an endpoint to authenticate over basic authentication",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"password": {
										Description:         "The secret in the service monitor namespace that contains the password for authentication.",
										MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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

									"username": {
										Description:         "The secret in the service monitor namespace that contains the username for authentication.",
										MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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

							"bearer_token": {
								Description:         "Bearer token for accessing apiserver.",
								MarkdownDescription: "Bearer token for accessing apiserver.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bearer_token_file": {
								Description:         "File to read bearer token for accessing apiserver.",
								MarkdownDescription: "File to read bearer token for accessing apiserver.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host": {
								Description:         "Host of apiserver. A valid string consisting of a hostname or IP followed by an optional port number",
								MarkdownDescription: "Host of apiserver. A valid string consisting of a hostname or IP followed by an optional port number",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"tls_config": {
								Description:         "TLS Config to use for accessing apiserver.",
								MarkdownDescription: "TLS Config to use for accessing apiserver.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "Struct containing the CA cert to use for the targets.",
										MarkdownDescription: "Struct containing the CA cert to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

									"ca_file": {
										Description:         "Path to the CA cert in the Prometheus container to use for the targets.",
										MarkdownDescription: "Path to the CA cert in the Prometheus container to use for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cert": {
										Description:         "Struct containing the client cert file for the targets.",
										MarkdownDescription: "Struct containing the client cert file for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

									"cert_file": {
										Description:         "Path to the client cert file in the Prometheus container for the targets.",
										MarkdownDescription: "Path to the client cert file in the Prometheus container for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure_skip_verify": {
										Description:         "Disable target certificate validation.",
										MarkdownDescription: "Disable target certificate validation.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_file": {
										Description:         "Path to the client key file in the Prometheus container for the targets.",
										MarkdownDescription: "Path to the client key file in the Prometheus container for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_secret": {
										Description:         "Secret containing the client key file for the targets.",
										MarkdownDescription: "Secret containing the client key file for the targets.",

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

									"server_name": {
										Description:         "Used to verify the hostname for the targets.",
										MarkdownDescription: "Used to verify the hostname for the targets.",

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

					"arbitrary_fs_access_through_s_ms": {
						Description:         "ArbitraryFSAccessThroughSMs configures whether configuration based on a service monitor can access arbitrary files on the file system of the Prometheus container e.g. bearer token files.",
						MarkdownDescription: "ArbitraryFSAccessThroughSMs configures whether configuration based on a service monitor can access arbitrary files on the file system of the Prometheus container e.g. bearer token files.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"deny": {
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

					"base_image": {
						Description:         "Base image to use for a Prometheus deployment. Deprecated: use 'image' instead",
						MarkdownDescription: "Base image to use for a Prometheus deployment. Deprecated: use 'image' instead",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"config_maps": {
						Description:         "ConfigMaps is a list of ConfigMaps in the same namespace as the Prometheus object, which shall be mounted into the Prometheus Pods. The ConfigMaps are mounted into /etc/prometheus/configmaps/<configmap-name>.",
						MarkdownDescription: "ConfigMaps is a list of ConfigMaps in the same namespace as the Prometheus object, which shall be mounted into the Prometheus Pods. The ConfigMaps are mounted into /etc/prometheus/configmaps/<configmap-name>.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"containers": {
						Description:         "Containers allows injecting additional containers or modifying operator generated containers. This can be used to allow adding an authentication proxy to a Prometheus pod or to change the behavior of an operator generated container. Containers described here modify an operator generated container if they share the same name and modifications are done via a strategic merge patch. The current container names are: 'prometheus', 'config-reloader', and 'thanos-sidecar'. Overriding containers is entirely outside the scope of what the maintainers will support and by doing so, you accept that this behaviour may break at any time without notice.",
						MarkdownDescription: "Containers allows injecting additional containers or modifying operator generated containers. This can be used to allow adding an authentication proxy to a Prometheus pod or to change the behavior of an operator generated container. Containers described here modify an operator generated container if they share the same name and modifications are done via a strategic merge patch. The current container names are: 'prometheus', 'config-reloader', and 'thanos-sidecar'. Overriding containers is entirely outside the scope of what the maintainers will support and by doing so, you accept that this behaviour may break at any time without notice.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"args": {
								Description:         "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
								MarkdownDescription: "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"command": {
								Description:         "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
								MarkdownDescription: "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"env": {
								Description:         "List of environment variables to set in the container. Cannot be updated.",
								MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",

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

														Type: types.StringType,

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
								Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
								MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",

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

							"image": {
								Description:         "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
								MarkdownDescription: "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_policy": {
								Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
								MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"lifecycle": {
								Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
								MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"post_start": {
										Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
										MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "Exec specifies the action to take.",
												MarkdownDescription: "Exec specifies the action to take.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
														Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

														Type: types.StringType,

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
												Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
												MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

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
														Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

									"pre_stop": {
										Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
										MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "Exec specifies the action to take.",
												MarkdownDescription: "Exec specifies the action to take.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
														Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

														Type: types.StringType,

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
												Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
												MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

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
														Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

							"liveness_probe": {
								Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: types.StringType,

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
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

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
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
								Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
								MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"ports": {
								Description:         "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
								MarkdownDescription: "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"container_port": {
										Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
										MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",

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
										Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
										MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
										MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",

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
								Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: types.StringType,

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
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

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
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
								Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

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
								Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
								MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",

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

							"startup_probe": {
								Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: types.StringType,

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
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

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
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
								Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
								MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"stdin_once": {
								Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
								MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"termination_message_path": {
								Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
								MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"termination_message_policy": {
								Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
								MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tty": {
								Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
								MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_devices": {
								Description:         "volumeDevices is the list of block devices to be used by the container.",
								MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"device_path": {
										Description:         "devicePath is the path inside of the container that the device will be mapped to.",
										MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",

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
								Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
								MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",

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

							"working_dir": {
								Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
								MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",

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

					"disable_compaction": {
						Description:         "Disable prometheus compaction.",
						MarkdownDescription: "Disable prometheus compaction.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_admin_api": {
						Description:         "Enable access to prometheus web admin API. Defaults to the value of 'false'. WARNING: Enabling the admin APIs enables mutating endpoints, to delete data, shutdown Prometheus, and more. Enabling this should be done with care and the user is advised to add additional authentication authorization via a proxy to ensure only clients authorized to perform these actions can do so. For more information see https://prometheus.io/docs/prometheus/latest/querying/api/#tsdb-admin-apis",
						MarkdownDescription: "Enable access to prometheus web admin API. Defaults to the value of 'false'. WARNING: Enabling the admin APIs enables mutating endpoints, to delete data, shutdown Prometheus, and more. Enabling this should be done with care and the user is advised to add additional authentication authorization via a proxy to ensure only clients authorized to perform these actions can do so. For more information see https://prometheus.io/docs/prometheus/latest/querying/api/#tsdb-admin-apis",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_features": {
						Description:         "Enable access to Prometheus disabled features. By default, no features are enabled. Enabling disabled features is entirely outside the scope of what the maintainers will support and by doing so, you accept that this behaviour may break at any time without notice. For more information see https://prometheus.io/docs/prometheus/latest/disabled_features/",
						MarkdownDescription: "Enable access to Prometheus disabled features. By default, no features are enabled. Enabling disabled features is entirely outside the scope of what the maintainers will support and by doing so, you accept that this behaviour may break at any time without notice. For more information see https://prometheus.io/docs/prometheus/latest/disabled_features/",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_remote_write_receiver": {
						Description:         "Enable Prometheus to be used as a receiver for the Prometheus remote write protocol. Defaults to the value of 'false'. WARNING: This is not considered an efficient way of ingesting samples. Use it with caution for specific low-volume use cases. It is not suitable for replacing the ingestion via scraping and turning Prometheus into a push-based metrics collection system. For more information see https://prometheus.io/docs/prometheus/latest/querying/api/#remote-write-receiver Only valid in Prometheus versions 2.33.0 and newer.",
						MarkdownDescription: "Enable Prometheus to be used as a receiver for the Prometheus remote write protocol. Defaults to the value of 'false'. WARNING: This is not considered an efficient way of ingesting samples. Use it with caution for specific low-volume use cases. It is not suitable for replacing the ingestion via scraping and turning Prometheus into a push-based metrics collection system. For more information see https://prometheus.io/docs/prometheus/latest/querying/api/#remote-write-receiver Only valid in Prometheus versions 2.33.0 and newer.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enforced_body_size_limit": {
						Description:         "EnforcedBodySizeLimit defines the maximum size of uncompressed response body that will be accepted by Prometheus. Targets responding with a body larger than this many bytes will cause the scrape to fail. Example: 100MB. If defined, the limit will apply to all service/pod monitors and probes. This is an experimental feature, this behaviour could change or be removed in the future. Only valid in Prometheus versions 2.28.0 and newer.",
						MarkdownDescription: "EnforcedBodySizeLimit defines the maximum size of uncompressed response body that will be accepted by Prometheus. Targets responding with a body larger than this many bytes will cause the scrape to fail. Example: 100MB. If defined, the limit will apply to all service/pod monitors and probes. This is an experimental feature, this behaviour could change or be removed in the future. Only valid in Prometheus versions 2.28.0 and newer.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enforced_label_limit": {
						Description:         "Per-scrape limit on number of labels that will be accepted for a sample. If more than this number of labels are present post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on number of labels that will be accepted for a sample. If more than this number of labels are present post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. Only valid in Prometheus versions 2.27.0 and newer.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enforced_label_name_length_limit": {
						Description:         "Per-scrape limit on length of labels name that will be accepted for a sample. If a label name is longer than this number post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on length of labels name that will be accepted for a sample. If a label name is longer than this number post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. Only valid in Prometheus versions 2.27.0 and newer.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enforced_label_value_length_limit": {
						Description:         "Per-scrape limit on length of labels value that will be accepted for a sample. If a label value is longer than this number post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. Only valid in Prometheus versions 2.27.0 and newer.",
						MarkdownDescription: "Per-scrape limit on length of labels value that will be accepted for a sample. If a label value is longer than this number post metric-relabeling, the entire scrape will be treated as failed. 0 means no limit. Only valid in Prometheus versions 2.27.0 and newer.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enforced_namespace_label": {
						Description:         "EnforcedNamespaceLabel If set, a label will be added to  1. all user-metrics (created by 'ServiceMonitor', 'PodMonitor' and 'Probe' objects) and 2. in all 'PrometheusRule' objects (except the ones excluded in 'prometheusRulesExcludedFromEnforce') to * alerting & recording rules and * the metrics used in their expressions ('expr').  Label name is this field's value. Label value is the namespace of the created object (mentioned above).",
						MarkdownDescription: "EnforcedNamespaceLabel If set, a label will be added to  1. all user-metrics (created by 'ServiceMonitor', 'PodMonitor' and 'Probe' objects) and 2. in all 'PrometheusRule' objects (except the ones excluded in 'prometheusRulesExcludedFromEnforce') to * alerting & recording rules and * the metrics used in their expressions ('expr').  Label name is this field's value. Label value is the namespace of the created object (mentioned above).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enforced_sample_limit": {
						Description:         "EnforcedSampleLimit defines global limit on number of scraped samples that will be accepted. This overrides any SampleLimit set per ServiceMonitor or/and PodMonitor. It is meant to be used by admins to enforce the SampleLimit to keep overall number of samples/series under the desired limit. Note that if SampleLimit is lower that value will be taken instead.",
						MarkdownDescription: "EnforcedSampleLimit defines global limit on number of scraped samples that will be accepted. This overrides any SampleLimit set per ServiceMonitor or/and PodMonitor. It is meant to be used by admins to enforce the SampleLimit to keep overall number of samples/series under the desired limit. Note that if SampleLimit is lower that value will be taken instead.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enforced_target_limit": {
						Description:         "EnforcedTargetLimit defines a global limit on the number of scraped targets.  This overrides any TargetLimit set per ServiceMonitor or/and PodMonitor.  It is meant to be used by admins to enforce the TargetLimit to keep the overall number of targets under the desired limit. Note that if TargetLimit is lower, that value will be taken instead, except if either value is zero, in which case the non-zero value will be used.  If both values are zero, no limit is enforced.",
						MarkdownDescription: "EnforcedTargetLimit defines a global limit on the number of scraped targets.  This overrides any TargetLimit set per ServiceMonitor or/and PodMonitor.  It is meant to be used by admins to enforce the TargetLimit to keep the overall number of targets under the desired limit. Note that if TargetLimit is lower, that value will be taken instead, except if either value is zero, in which case the non-zero value will be used.  If both values are zero, no limit is enforced.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"evaluation_interval": {
						Description:         "Interval between consecutive evaluations. Default: '30s'",
						MarkdownDescription: "Interval between consecutive evaluations. Default: '30s'",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"excluded_from_enforcement": {
						Description:         "List of references to PodMonitor, ServiceMonitor, Probe and PrometheusRule objects to be excluded from enforcing a namespace label of origin. Applies only if enforcedNamespaceLabel set to true.",
						MarkdownDescription: "List of references to PodMonitor, ServiceMonitor, Probe and PrometheusRule objects to be excluded from enforcing a namespace label of origin. Applies only if enforcedNamespaceLabel set to true.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"group": {
								Description:         "Group of the referent. When not specified, it defaults to 'monitoring.coreos.com'",
								MarkdownDescription: "Group of the referent. When not specified, it defaults to 'monitoring.coreos.com'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent. When not set, all resources are matched.",
								MarkdownDescription: "Name of the referent. When not set, all resources are matched.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"resource": {
								Description:         "Resource of the referent.",
								MarkdownDescription: "Resource of the referent.",

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

					"exemplars": {
						Description:         "Exemplars related settings that are runtime reloadable. It requires to enable the exemplar storage feature to be effective.",
						MarkdownDescription: "Exemplars related settings that are runtime reloadable. It requires to enable the exemplar storage feature to be effective.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"max_size": {
								Description:         "Maximum number of exemplars stored in memory for all series. If not set, Prometheus uses its default value. A value of zero or less than zero disables the storage.",
								MarkdownDescription: "Maximum number of exemplars stored in memory for all series. If not set, Prometheus uses its default value. A value of zero or less than zero disables the storage.",

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

					"external_labels": {
						Description:         "The labels to add to any time series or alerts when communicating with external systems (federation, remote storage, Alertmanager).",
						MarkdownDescription: "The labels to add to any time series or alerts when communicating with external systems (federation, remote storage, Alertmanager).",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_url": {
						Description:         "The external URL the Prometheus instances will be available under. This is necessary to generate correct URLs. This is necessary if Prometheus is not served from root of a DNS name.",
						MarkdownDescription: "The external URL the Prometheus instances will be available under. This is necessary to generate correct URLs. This is necessary if Prometheus is not served from root of a DNS name.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"host_aliases": {
						Description:         "Pods' hostAliases configuration",
						MarkdownDescription: "Pods' hostAliases configuration",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"hostnames": {
								Description:         "Hostnames for the above IP address.",
								MarkdownDescription: "Hostnames for the above IP address.",

								Type: types.ListType{ElemType: types.StringType},

								Required: true,
								Optional: false,
								Computed: false,
							},

							"ip": {
								Description:         "IP address of the host file entry.",
								MarkdownDescription: "IP address of the host file entry.",

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

					"ignore_namespace_selectors": {
						Description:         "IgnoreNamespaceSelectors if set to true will ignore NamespaceSelector settings from all PodMonitor, ServiceMonitor and Probe objects. They will only discover endpoints within the namespace of the PodMonitor, ServiceMonitor and Probe objects. Defaults to false.",
						MarkdownDescription: "IgnoreNamespaceSelectors if set to true will ignore NamespaceSelector settings from all PodMonitor, ServiceMonitor and Probe objects. They will only discover endpoints within the namespace of the PodMonitor, ServiceMonitor and Probe objects. Defaults to false.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "Image if specified has precedence over baseImage, tag and sha combinations. Specifying the version is still necessary to ensure the Prometheus Operator knows what version of Prometheus is being configured.",
						MarkdownDescription: "Image if specified has precedence over baseImage, tag and sha combinations. Specifying the version is still necessary to ensure the Prometheus Operator knows what version of Prometheus is being configured.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_pull_secrets": {
						Description:         "An optional list of references to secrets in the same namespace to use for pulling prometheus and alertmanager images from registries see http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod",
						MarkdownDescription: "An optional list of references to secrets in the same namespace to use for pulling prometheus and alertmanager images from registries see http://kubernetes.io/docs/user-guide/images#specifying-imagepullsecrets-on-a-pod",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

					"init_containers": {
						Description:         "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g. fetch secrets for injection into the Prometheus configuration from external sources. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ InitContainers described here modify an operator generated init containers if they share the same name and modifications are done via a strategic merge patch. The current init container name is: 'init-config-reloader'. Overriding init containers is entirely outside the scope of what the maintainers will support and by doing so, you accept that this behaviour may break at any time without notice.",
						MarkdownDescription: "InitContainers allows adding initContainers to the pod definition. Those can be used to e.g. fetch secrets for injection into the Prometheus configuration from external sources. Any errors during the execution of an initContainer will lead to a restart of the Pod. More info: https://kubernetes.io/docs/concepts/workloads/pods/init-containers/ InitContainers described here modify an operator generated init containers if they share the same name and modifications are done via a strategic merge patch. The current init container name is: 'init-config-reloader'. Overriding init containers is entirely outside the scope of what the maintainers will support and by doing so, you accept that this behaviour may break at any time without notice.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"args": {
								Description:         "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
								MarkdownDescription: "Arguments to the entrypoint. The container image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"command": {
								Description:         "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
								MarkdownDescription: "Entrypoint array. Not executed within a shell. The container image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"env": {
								Description:         "List of environment variables to set in the container. Cannot be updated.",
								MarkdownDescription: "List of environment variables to set in the container. Cannot be updated.",

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

														Type: types.StringType,

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
								Description:         "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",
								MarkdownDescription: "List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.",

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

							"image": {
								Description:         "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
								MarkdownDescription: "Container image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_policy": {
								Description:         "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
								MarkdownDescription: "Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"lifecycle": {
								Description:         "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",
								MarkdownDescription: "Actions that the management system should take in response to container lifecycle events. Cannot be updated.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"post_start": {
										Description:         "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
										MarkdownDescription: "PostStart is called immediately after a container is created. If the handler fails, the container is terminated and restarted according to its restart policy. Other management of the container blocks until the hook completes. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "Exec specifies the action to take.",
												MarkdownDescription: "Exec specifies the action to take.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
														Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

														Type: types.StringType,

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
												Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
												MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

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
														Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

									"pre_stop": {
										Description:         "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",
										MarkdownDescription: "PreStop is called immediately before a container is terminated due to an API request or management event such as liveness/startup probe failure, preemption, resource contention, etc. The handler is not called if the container crashes or exits. The Pod's termination grace period countdown begins before the PreStop hook is executed. Regardless of the outcome of the handler, the container will eventually terminate within the Pod's termination grace period (unless delayed by finalizers). Other management of the container blocks until the hook completes or until the termination grace period is reached. More info: https://kubernetes.io/docs/concepts/containers/container-lifecycle-hooks/#container-hooks",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"exec": {
												Description:         "Exec specifies the action to take.",
												MarkdownDescription: "Exec specifies the action to take.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
														MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
														Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
														MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
														Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

														Type: types.StringType,

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
												Description:         "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",
												MarkdownDescription: "Deprecated. TCPSocket is NOT supported as a LifecycleHandler and kept for the backward compatibility. There are no validation of this field and lifecycle hooks will fail in runtime when tcp handler is specified.",

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
														Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
														MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

							"liveness_probe": {
								Description:         "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Periodic probe of container liveness. Container will be restarted if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: types.StringType,

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
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

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
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
								Description:         "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",
								MarkdownDescription: "Name of the container specified as a DNS_LABEL. Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"ports": {
								Description:         "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",
								MarkdownDescription: "List of ports to expose from the container. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Modifying this array with strategic merge patch may corrupt the data. For more information See https://github.com/kubernetes/kubernetes/issues/108255. Cannot be updated.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"container_port": {
										Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
										MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",

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
										Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
										MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
										MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",

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
								Description:         "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "Periodic probe of container service readiness. Container will be removed from service endpoints if the probe fails. Cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: types.StringType,

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
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

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
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
								Description:         "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Compute Resources required by this container. Cannot be updated. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",

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
								Description:         "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",
								MarkdownDescription: "SecurityContext defines the security options the container should be run with. If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext. More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/",

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

							"startup_probe": {
								Description:         "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
								MarkdownDescription: "StartupProbe indicates that the Pod has successfully initialized. If specified, no other probes are executed until this completes successfully. If this probe fails, the Pod will be restarted, just as if the livenessProbe failed. This can be used to provide different probe parameters at the beginning of a Pod's lifecycle, when it might take a long time to load data or warm a cache, than during steady-state operation. This cannot be updated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is a beta field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

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
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: types.StringType,

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
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

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
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

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

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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
								Description:         "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",
								MarkdownDescription: "Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"stdin_once": {
								Description:         "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",
								MarkdownDescription: "Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"termination_message_path": {
								Description:         "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",
								MarkdownDescription: "Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"termination_message_policy": {
								Description:         "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",
								MarkdownDescription: "Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tty": {
								Description:         "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",
								MarkdownDescription: "Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_devices": {
								Description:         "volumeDevices is the list of block devices to be used by the container.",
								MarkdownDescription: "volumeDevices is the list of block devices to be used by the container.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"device_path": {
										Description:         "devicePath is the path inside of the container that the device will be mapped to.",
										MarkdownDescription: "devicePath is the path inside of the container that the device will be mapped to.",

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
								Description:         "Pod volumes to mount into the container's filesystem. Cannot be updated.",
								MarkdownDescription: "Pod volumes to mount into the container's filesystem. Cannot be updated.",

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

							"working_dir": {
								Description:         "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",
								MarkdownDescription: "Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.",

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

					"listen_local": {
						Description:         "ListenLocal makes the Prometheus server listen on loopback, so that it does not bind against the Pod IP.",
						MarkdownDescription: "ListenLocal makes the Prometheus server listen on loopback, so that it does not bind against the Pod IP.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_format": {
						Description:         "Log format for Prometheus to be configured with.",
						MarkdownDescription: "Log format for Prometheus to be configured with.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"log_level": {
						Description:         "Log level for Prometheus to be configured with.",
						MarkdownDescription: "Log level for Prometheus to be configured with.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"min_ready_seconds": {
						Description:         "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate.",
						MarkdownDescription: "Minimum number of seconds for which a newly created pod should be ready without any of its container crashing for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) This is an alpha field and requires enabling StatefulSetMinReadySeconds feature gate.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_selector": {
						Description:         "Define which Nodes the Pods are scheduled on.",
						MarkdownDescription: "Define which Nodes the Pods are scheduled on.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"override_honor_labels": {
						Description:         "When true, Prometheus resolves label conflicts by renaming the labels in the scraped data to 'exported_<label value>' for all targets created from service and pod monitors. Otherwise the HonorLabels field of the service or pod monitor applies.",
						MarkdownDescription: "When true, Prometheus resolves label conflicts by renaming the labels in the scraped data to 'exported_<label value>' for all targets created from service and pod monitors. Otherwise the HonorLabels field of the service or pod monitor applies.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"override_honor_timestamps": {
						Description:         "When true, Prometheus ignores the timestamps for all the targets created from service and pod monitors. Otherwise the HonorTimestamps field of the service or pod monitor applies.",
						MarkdownDescription: "When true, Prometheus ignores the timestamps for all the targets created from service and pod monitors. Otherwise the HonorTimestamps field of the service or pod monitor applies.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"paused": {
						Description:         "When a Prometheus deployment is paused, no actions except for deletion will be performed on the underlying objects.",
						MarkdownDescription: "When a Prometheus deployment is paused, no actions except for deletion will be performed on the underlying objects.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"pod_metadata": {
						Description:         "PodMetadata configures Labels and Annotations which are propagated to the prometheus pods.",
						MarkdownDescription: "PodMetadata configures Labels and Annotations which are propagated to the prometheus pods.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"annotations": {
								Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
								MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"labels": {
								Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
								MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
								MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",

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

					"pod_monitor_namespace_selector": {
						Description:         "Namespace's labels to match for PodMonitor discovery. If nil, only check own namespace.",
						MarkdownDescription: "Namespace's labels to match for PodMonitor discovery. If nil, only check own namespace.",

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

					"pod_monitor_selector": {
						Description:         "*Experimental* PodMonitors to be selected for target discovery. *Deprecated:* if neither this nor serviceMonitorSelector are specified, configuration is unmanaged.",
						MarkdownDescription: "*Experimental* PodMonitors to be selected for target discovery. *Deprecated:* if neither this nor serviceMonitorSelector are specified, configuration is unmanaged.",

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

					"port_name": {
						Description:         "Port name used for the pods and governing service. This defaults to web",
						MarkdownDescription: "Port name used for the pods and governing service. This defaults to web",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"priority_class_name": {
						Description:         "Priority class assigned to the Pods",
						MarkdownDescription: "Priority class assigned to the Pods",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"probe_namespace_selector": {
						Description:         "*Experimental* Namespaces to be selected for Probe discovery. If nil, only check own namespace.",
						MarkdownDescription: "*Experimental* Namespaces to be selected for Probe discovery. If nil, only check own namespace.",

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

					"probe_selector": {
						Description:         "*Experimental* Probes to be selected for target discovery.",
						MarkdownDescription: "*Experimental* Probes to be selected for target discovery.",

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

					"prometheus_external_label_name": {
						Description:         "Name of Prometheus external label used to denote Prometheus instance name. Defaults to the value of 'prometheus'. External label will _not_ be added when value is set to empty string ('''').",
						MarkdownDescription: "Name of Prometheus external label used to denote Prometheus instance name. Defaults to the value of 'prometheus'. External label will _not_ be added when value is set to empty string ('''').",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus_rules_excluded_from_enforce": {
						Description:         "PrometheusRulesExcludedFromEnforce - list of prometheus rules to be excluded from enforcing of adding namespace labels. Works only if enforcedNamespaceLabel set to true. Make sure both ruleNamespace and ruleName are set for each pair. Deprecated: use excludedFromEnforcement instead.",
						MarkdownDescription: "PrometheusRulesExcludedFromEnforce - list of prometheus rules to be excluded from enforcing of adding namespace labels. Works only if enforcedNamespaceLabel set to true. Make sure both ruleNamespace and ruleName are set for each pair. Deprecated: use excludedFromEnforcement instead.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"rule_name": {
								Description:         "RuleNamespace - name of excluded rule",
								MarkdownDescription: "RuleNamespace - name of excluded rule",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"rule_namespace": {
								Description:         "RuleNamespace - namespace of excluded rule",
								MarkdownDescription: "RuleNamespace - namespace of excluded rule",

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

					"query": {
						Description:         "QuerySpec defines the query command line flags when starting Prometheus.",
						MarkdownDescription: "QuerySpec defines the query command line flags when starting Prometheus.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"lookback_delta": {
								Description:         "The delta difference allowed for retrieving metrics during expression evaluations.",
								MarkdownDescription: "The delta difference allowed for retrieving metrics during expression evaluations.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_concurrency": {
								Description:         "Number of concurrent queries that can be run at once.",
								MarkdownDescription: "Number of concurrent queries that can be run at once.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_samples": {
								Description:         "Maximum number of samples a single query can load into memory. Note that queries will fail if they would load more samples than this into memory, so this also limits the number of samples a query can return.",
								MarkdownDescription: "Maximum number of samples a single query can load into memory. Note that queries will fail if they would load more samples than this into memory, so this also limits the number of samples a query can return.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"timeout": {
								Description:         "Maximum time a query may take before being aborted.",
								MarkdownDescription: "Maximum time a query may take before being aborted.",

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

					"query_log_file": {
						Description:         "QueryLogFile specifies the file to which PromQL queries are logged. If the filename has an empty path, e.g. 'query.log', prometheus-operator will mount the file into an emptyDir volume at '/var/log/prometheus'. If a full path is provided, e.g. /var/log/prometheus/query.log, you must mount a volume in the specified directory and it must be writable. This is because the prometheus container runs with a read-only root filesystem for security reasons. Alternatively, the location can be set to a stdout location such as '/dev/stdout' to log query information to the default Prometheus log stream. This is only available in versions of Prometheus >= 2.16.0. For more details, see the Prometheus docs (https://prometheus.io/docs/guides/query-log/)",
						MarkdownDescription: "QueryLogFile specifies the file to which PromQL queries are logged. If the filename has an empty path, e.g. 'query.log', prometheus-operator will mount the file into an emptyDir volume at '/var/log/prometheus'. If a full path is provided, e.g. /var/log/prometheus/query.log, you must mount a volume in the specified directory and it must be writable. This is because the prometheus container runs with a read-only root filesystem for security reasons. Alternatively, the location can be set to a stdout location such as '/dev/stdout' to log query information to the default Prometheus log stream. This is only available in versions of Prometheus >= 2.16.0. For more details, see the Prometheus docs (https://prometheus.io/docs/guides/query-log/)",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"remote_read": {
						Description:         "remoteRead is the list of remote read configurations.",
						MarkdownDescription: "remoteRead is the list of remote read configurations.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"authorization": {
								Description:         "Authorization section for remote read",
								MarkdownDescription: "Authorization section for remote read",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials": {
										Description:         "The secret's key that contains the credentials of the request",
										MarkdownDescription: "The secret's key that contains the credentials of the request",

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

									"credentials_file": {
										Description:         "File to read a secret from, mutually exclusive with Credentials (from SafeAuthorization)",
										MarkdownDescription: "File to read a secret from, mutually exclusive with Credentials (from SafeAuthorization)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
										MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

							"basic_auth": {
								Description:         "BasicAuth for the URL.",
								MarkdownDescription: "BasicAuth for the URL.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"password": {
										Description:         "The secret in the service monitor namespace that contains the password for authentication.",
										MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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

									"username": {
										Description:         "The secret in the service monitor namespace that contains the username for authentication.",
										MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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

							"bearer_token": {
								Description:         "Bearer token for remote read.",
								MarkdownDescription: "Bearer token for remote read.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bearer_token_file": {
								Description:         "File to read bearer token for remote read.",
								MarkdownDescription: "File to read bearer token for remote read.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"headers": {
								Description:         "Custom HTTP headers to be sent along with each remote read request. Be aware that headers that are set by Prometheus itself can't be overwritten. Only valid in Prometheus versions 2.26.0 and newer.",
								MarkdownDescription: "Custom HTTP headers to be sent along with each remote read request. Be aware that headers that are set by Prometheus itself can't be overwritten. Only valid in Prometheus versions 2.26.0 and newer.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "The name of the remote read queue, it must be unique if specified. The name is used in metrics and logging in order to differentiate read configurations.  Only valid in Prometheus versions 2.15.0 and newer.",
								MarkdownDescription: "The name of the remote read queue, it must be unique if specified. The name is used in metrics and logging in order to differentiate read configurations.  Only valid in Prometheus versions 2.15.0 and newer.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"oauth2": {
								Description:         "OAuth2 for the URL. Only valid in Prometheus versions 2.27.0 and newer.",
								MarkdownDescription: "OAuth2 for the URL. Only valid in Prometheus versions 2.27.0 and newer.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client_id": {
										Description:         "The secret or configmap containing the OAuth2 client id",
										MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

										Required: true,
										Optional: false,
										Computed: false,
									},

									"client_secret": {
										Description:         "The secret containing the OAuth2 client secret",
										MarkdownDescription: "The secret containing the OAuth2 client secret",

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

										Required: true,
										Optional: false,
										Computed: false,
									},

									"endpoint_params": {
										Description:         "Parameters to append to the token URL",
										MarkdownDescription: "Parameters to append to the token URL",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scopes": {
										Description:         "OAuth2 scopes used for the token request",
										MarkdownDescription: "OAuth2 scopes used for the token request",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"token_url": {
										Description:         "The URL to fetch the token from",
										MarkdownDescription: "The URL to fetch the token from",

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

							"proxy_url": {
								Description:         "Optional ProxyURL.",
								MarkdownDescription: "Optional ProxyURL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"read_recent": {
								Description:         "Whether reads should be made for queries for time ranges that the local storage should have complete data for.",
								MarkdownDescription: "Whether reads should be made for queries for time ranges that the local storage should have complete data for.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"remote_timeout": {
								Description:         "Timeout for requests to the remote read endpoint.",
								MarkdownDescription: "Timeout for requests to the remote read endpoint.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"required_matchers": {
								Description:         "An optional list of equality matchers which have to be present in a selector to query the remote read endpoint.",
								MarkdownDescription: "An optional list of equality matchers which have to be present in a selector to query the remote read endpoint.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_config": {
								Description:         "TLS Config to use for remote read.",
								MarkdownDescription: "TLS Config to use for remote read.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "Struct containing the CA cert to use for the targets.",
										MarkdownDescription: "Struct containing the CA cert to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

									"ca_file": {
										Description:         "Path to the CA cert in the Prometheus container to use for the targets.",
										MarkdownDescription: "Path to the CA cert in the Prometheus container to use for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cert": {
										Description:         "Struct containing the client cert file for the targets.",
										MarkdownDescription: "Struct containing the client cert file for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

									"cert_file": {
										Description:         "Path to the client cert file in the Prometheus container for the targets.",
										MarkdownDescription: "Path to the client cert file in the Prometheus container for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure_skip_verify": {
										Description:         "Disable target certificate validation.",
										MarkdownDescription: "Disable target certificate validation.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_file": {
										Description:         "Path to the client key file in the Prometheus container for the targets.",
										MarkdownDescription: "Path to the client key file in the Prometheus container for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_secret": {
										Description:         "Secret containing the client key file for the targets.",
										MarkdownDescription: "Secret containing the client key file for the targets.",

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

									"server_name": {
										Description:         "Used to verify the hostname for the targets.",
										MarkdownDescription: "Used to verify the hostname for the targets.",

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

							"url": {
								Description:         "The URL of the endpoint to query from.",
								MarkdownDescription: "The URL of the endpoint to query from.",

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

					"remote_write": {
						Description:         "remoteWrite is the list of remote write configurations.",
						MarkdownDescription: "remoteWrite is the list of remote write configurations.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"authorization": {
								Description:         "Authorization section for remote write",
								MarkdownDescription: "Authorization section for remote write",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"credentials": {
										Description:         "The secret's key that contains the credentials of the request",
										MarkdownDescription: "The secret's key that contains the credentials of the request",

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

									"credentials_file": {
										Description:         "File to read a secret from, mutually exclusive with Credentials (from SafeAuthorization)",
										MarkdownDescription: "File to read a secret from, mutually exclusive with Credentials (from SafeAuthorization)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": {
										Description:         "Set the authentication type. Defaults to Bearer, Basic will cause an error",
										MarkdownDescription: "Set the authentication type. Defaults to Bearer, Basic will cause an error",

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

							"basic_auth": {
								Description:         "BasicAuth for the URL.",
								MarkdownDescription: "BasicAuth for the URL.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"password": {
										Description:         "The secret in the service monitor namespace that contains the password for authentication.",
										MarkdownDescription: "The secret in the service monitor namespace that contains the password for authentication.",

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

									"username": {
										Description:         "The secret in the service monitor namespace that contains the username for authentication.",
										MarkdownDescription: "The secret in the service monitor namespace that contains the username for authentication.",

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

							"bearer_token": {
								Description:         "Bearer token for remote write.",
								MarkdownDescription: "Bearer token for remote write.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"bearer_token_file": {
								Description:         "File to read bearer token for remote write.",
								MarkdownDescription: "File to read bearer token for remote write.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"headers": {
								Description:         "Custom HTTP headers to be sent along with each remote write request. Be aware that headers that are set by Prometheus itself can't be overwritten. Only valid in Prometheus versions 2.25.0 and newer.",
								MarkdownDescription: "Custom HTTP headers to be sent along with each remote write request. Be aware that headers that are set by Prometheus itself can't be overwritten. Only valid in Prometheus versions 2.25.0 and newer.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"metadata_config": {
								Description:         "MetadataConfig configures the sending of series metadata to the remote storage.",
								MarkdownDescription: "MetadataConfig configures the sending of series metadata to the remote storage.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"send": {
										Description:         "Whether metric metadata is sent to the remote storage or not.",
										MarkdownDescription: "Whether metric metadata is sent to the remote storage or not.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"send_interval": {
										Description:         "How frequently metric metadata is sent to the remote storage.",
										MarkdownDescription: "How frequently metric metadata is sent to the remote storage.",

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
								Description:         "The name of the remote write queue, it must be unique if specified. The name is used in metrics and logging in order to differentiate queues. Only valid in Prometheus versions 2.15.0 and newer.",
								MarkdownDescription: "The name of the remote write queue, it must be unique if specified. The name is used in metrics and logging in order to differentiate queues. Only valid in Prometheus versions 2.15.0 and newer.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"oauth2": {
								Description:         "OAuth2 for the URL. Only valid in Prometheus versions 2.27.0 and newer.",
								MarkdownDescription: "OAuth2 for the URL. Only valid in Prometheus versions 2.27.0 and newer.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client_id": {
										Description:         "The secret or configmap containing the OAuth2 client id",
										MarkdownDescription: "The secret or configmap containing the OAuth2 client id",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

										Required: true,
										Optional: false,
										Computed: false,
									},

									"client_secret": {
										Description:         "The secret containing the OAuth2 client secret",
										MarkdownDescription: "The secret containing the OAuth2 client secret",

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

										Required: true,
										Optional: false,
										Computed: false,
									},

									"endpoint_params": {
										Description:         "Parameters to append to the token URL",
										MarkdownDescription: "Parameters to append to the token URL",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scopes": {
										Description:         "OAuth2 scopes used for the token request",
										MarkdownDescription: "OAuth2 scopes used for the token request",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"token_url": {
										Description:         "The URL to fetch the token from",
										MarkdownDescription: "The URL to fetch the token from",

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

							"proxy_url": {
								Description:         "Optional ProxyURL.",
								MarkdownDescription: "Optional ProxyURL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"queue_config": {
								Description:         "QueueConfig allows tuning of the remote write queue parameters.",
								MarkdownDescription: "QueueConfig allows tuning of the remote write queue parameters.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"batch_send_deadline": {
										Description:         "BatchSendDeadline is the maximum time a sample will wait in buffer.",
										MarkdownDescription: "BatchSendDeadline is the maximum time a sample will wait in buffer.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"capacity": {
										Description:         "Capacity is the number of samples to buffer per shard before we start dropping them.",
										MarkdownDescription: "Capacity is the number of samples to buffer per shard before we start dropping them.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_backoff": {
										Description:         "MaxBackoff is the maximum retry delay.",
										MarkdownDescription: "MaxBackoff is the maximum retry delay.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_retries": {
										Description:         "MaxRetries is the maximum number of times to retry a batch on recoverable errors.",
										MarkdownDescription: "MaxRetries is the maximum number of times to retry a batch on recoverable errors.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_samples_per_send": {
										Description:         "MaxSamplesPerSend is the maximum number of samples per send.",
										MarkdownDescription: "MaxSamplesPerSend is the maximum number of samples per send.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"max_shards": {
										Description:         "MaxShards is the maximum number of shards, i.e. amount of concurrency.",
										MarkdownDescription: "MaxShards is the maximum number of shards, i.e. amount of concurrency.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_backoff": {
										Description:         "MinBackoff is the initial retry delay. Gets doubled for every retry.",
										MarkdownDescription: "MinBackoff is the initial retry delay. Gets doubled for every retry.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_shards": {
										Description:         "MinShards is the minimum number of shards, i.e. amount of concurrency.",
										MarkdownDescription: "MinShards is the minimum number of shards, i.e. amount of concurrency.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"retry_on_rate_limit": {
										Description:         "Retry upon receiving a 429 status code from the remote-write storage. This is experimental feature and might change in the future.",
										MarkdownDescription: "Retry upon receiving a 429 status code from the remote-write storage. This is experimental feature and might change in the future.",

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

							"remote_timeout": {
								Description:         "Timeout for requests to the remote write endpoint.",
								MarkdownDescription: "Timeout for requests to the remote write endpoint.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"send_exemplars": {
								Description:         "Enables sending of exemplars over remote write. Note that exemplar-storage itself must be enabled using the enableFeature option for exemplars to be scraped in the first place.  Only valid in Prometheus versions 2.27.0 and newer.",
								MarkdownDescription: "Enables sending of exemplars over remote write. Note that exemplar-storage itself must be enabled using the enableFeature option for exemplars to be scraped in the first place.  Only valid in Prometheus versions 2.27.0 and newer.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sigv4": {
								Description:         "Sigv4 allows to configures AWS's Signature Verification 4",
								MarkdownDescription: "Sigv4 allows to configures AWS's Signature Verification 4",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"access_key": {
										Description:         "AccessKey is the AWS API key. If blank, the environment variable 'AWS_ACCESS_KEY_ID' is used.",
										MarkdownDescription: "AccessKey is the AWS API key. If blank, the environment variable 'AWS_ACCESS_KEY_ID' is used.",

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

									"profile": {
										Description:         "Profile is the named AWS profile used to authenticate.",
										MarkdownDescription: "Profile is the named AWS profile used to authenticate.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"region": {
										Description:         "Region is the AWS region. If blank, the region from the default credentials chain used.",
										MarkdownDescription: "Region is the AWS region. If blank, the region from the default credentials chain used.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"role_arn": {
										Description:         "RoleArn is the named AWS profile used to authenticate.",
										MarkdownDescription: "RoleArn is the named AWS profile used to authenticate.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_key": {
										Description:         "SecretKey is the AWS API secret. If blank, the environment variable 'AWS_SECRET_ACCESS_KEY' is used.",
										MarkdownDescription: "SecretKey is the AWS API secret. If blank, the environment variable 'AWS_SECRET_ACCESS_KEY' is used.",

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

							"tls_config": {
								Description:         "TLS Config to use for remote write.",
								MarkdownDescription: "TLS Config to use for remote write.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "Struct containing the CA cert to use for the targets.",
										MarkdownDescription: "Struct containing the CA cert to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

									"ca_file": {
										Description:         "Path to the CA cert in the Prometheus container to use for the targets.",
										MarkdownDescription: "Path to the CA cert in the Prometheus container to use for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cert": {
										Description:         "Struct containing the client cert file for the targets.",
										MarkdownDescription: "Struct containing the client cert file for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

									"cert_file": {
										Description:         "Path to the client cert file in the Prometheus container for the targets.",
										MarkdownDescription: "Path to the client cert file in the Prometheus container for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure_skip_verify": {
										Description:         "Disable target certificate validation.",
										MarkdownDescription: "Disable target certificate validation.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_file": {
										Description:         "Path to the client key file in the Prometheus container for the targets.",
										MarkdownDescription: "Path to the client key file in the Prometheus container for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_secret": {
										Description:         "Secret containing the client key file for the targets.",
										MarkdownDescription: "Secret containing the client key file for the targets.",

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

									"server_name": {
										Description:         "Used to verify the hostname for the targets.",
										MarkdownDescription: "Used to verify the hostname for the targets.",

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

							"url": {
								Description:         "The URL of the endpoint to send samples to.",
								MarkdownDescription: "The URL of the endpoint to send samples to.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"write_relabel_configs": {
								Description:         "The list of remote write relabel configurations.",
								MarkdownDescription: "The list of remote write relabel configurations.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"action": {
										Description:         "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",
										MarkdownDescription: "Action to perform based on regex matching. Default is 'replace'. uppercase and lowercase actions require Prometheus >= 2.36.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"modulus": {
										Description:         "Modulus to take of the hash of the source label values.",
										MarkdownDescription: "Modulus to take of the hash of the source label values.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"regex": {
										Description:         "Regular expression against which the extracted value is matched. Default is '(.*)'",
										MarkdownDescription: "Regular expression against which the extracted value is matched. Default is '(.*)'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replacement": {
										Description:         "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",
										MarkdownDescription: "Replacement value against which a regex replace is performed if the regular expression matches. Regex capture groups are available. Default is '$1'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"separator": {
										Description:         "Separator placed between concatenated source label values. default is ';'.",
										MarkdownDescription: "Separator placed between concatenated source label values. default is ';'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"source_labels": {
										Description:         "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",
										MarkdownDescription: "The source labels select values from existing labels. Their content is concatenated using the configured separator and matched against the configured regular expression for the replace, keep, and drop actions.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"target_label": {
										Description:         "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",
										MarkdownDescription: "Label to which the resulting value is written in a replace action. It is mandatory for replace actions. Regex capture groups are available.",

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

					"replica_external_label_name": {
						Description:         "Name of Prometheus external label used to denote replica name. Defaults to the value of 'prometheus_replica'. External label will _not_ be added when value is set to empty string ('''').",
						MarkdownDescription: "Name of Prometheus external label used to denote replica name. Defaults to the value of 'prometheus_replica'. External label will _not_ be added when value is set to empty string ('''').",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": {
						Description:         "Number of replicas of each shard to deploy for a Prometheus deployment. Number of replicas multiplied by shards is the total number of Pods created.",
						MarkdownDescription: "Number of replicas of each shard to deploy for a Prometheus deployment. Number of replicas multiplied by shards is the total number of Pods created.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": {
						Description:         "Define resources requests and limits for single Pods.",
						MarkdownDescription: "Define resources requests and limits for single Pods.",

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

					"retention": {
						Description:         "Time duration Prometheus shall retain data for. Default is '24h' if retentionSize is not set, and must match the regular expression '[0-9]+(ms|s|m|h|d|w|y)' (milliseconds seconds minutes hours days weeks years).",
						MarkdownDescription: "Time duration Prometheus shall retain data for. Default is '24h' if retentionSize is not set, and must match the regular expression '[0-9]+(ms|s|m|h|d|w|y)' (milliseconds seconds minutes hours days weeks years).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"retention_size": {
						Description:         "Maximum amount of disk space used by blocks.",
						MarkdownDescription: "Maximum amount of disk space used by blocks.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"route_prefix": {
						Description:         "The route prefix Prometheus registers HTTP handlers for. This is useful, if using ExternalURL and a proxy is rewriting HTTP routes of a request, and the actual ExternalURL is still true, but the server serves requests under a different route prefix. For example for use with 'kubectl proxy'.",
						MarkdownDescription: "The route prefix Prometheus registers HTTP handlers for. This is useful, if using ExternalURL and a proxy is rewriting HTTP routes of a request, and the actual ExternalURL is still true, but the server serves requests under a different route prefix. For example for use with 'kubectl proxy'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rule_namespace_selector": {
						Description:         "Namespaces to be selected for PrometheusRules discovery. If unspecified, only the same namespace as the Prometheus object is in is used.",
						MarkdownDescription: "Namespaces to be selected for PrometheusRules discovery. If unspecified, only the same namespace as the Prometheus object is in is used.",

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

					"rule_selector": {
						Description:         "A selector to select which PrometheusRules to mount for loading alerting/recording rules from. Until (excluding) Prometheus Operator v0.24.0 Prometheus Operator will migrate any legacy rule ConfigMaps to PrometheusRule custom resources selected by RuleSelector. Make sure it does not match any config maps that you do not want to be migrated.",
						MarkdownDescription: "A selector to select which PrometheusRules to mount for loading alerting/recording rules from. Until (excluding) Prometheus Operator v0.24.0 Prometheus Operator will migrate any legacy rule ConfigMaps to PrometheusRule custom resources selected by RuleSelector. Make sure it does not match any config maps that you do not want to be migrated.",

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

					"rules": {
						Description:         "/--rules.*/ command-line arguments.",
						MarkdownDescription: "/--rules.*/ command-line arguments.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"alert": {
								Description:         "/--rules.alert.*/ command-line arguments",
								MarkdownDescription: "/--rules.alert.*/ command-line arguments",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"for_grace_period": {
										Description:         "Minimum duration between alert and restored 'for' state. This is maintained only for alerts with configured 'for' time greater than grace period.",
										MarkdownDescription: "Minimum duration between alert and restored 'for' state. This is maintained only for alerts with configured 'for' time greater than grace period.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"for_outage_tolerance": {
										Description:         "Max time to tolerate prometheus outage for restoring 'for' state of alert.",
										MarkdownDescription: "Max time to tolerate prometheus outage for restoring 'for' state of alert.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resend_delay": {
										Description:         "Minimum amount of time to wait before resending an alert to Alertmanager.",
										MarkdownDescription: "Minimum amount of time to wait before resending an alert to Alertmanager.",

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

					"scrape_interval": {
						Description:         "Interval between consecutive scrapes. Default: '30s'",
						MarkdownDescription: "Interval between consecutive scrapes. Default: '30s'",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"scrape_timeout": {
						Description:         "Number of seconds to wait for target to respond before erroring.",
						MarkdownDescription: "Number of seconds to wait for target to respond before erroring.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"secrets": {
						Description:         "Secrets is a list of Secrets in the same namespace as the Prometheus object, which shall be mounted into the Prometheus Pods. The Secrets are mounted into /etc/prometheus/secrets/<secret-name>.",
						MarkdownDescription: "Secrets is a list of Secrets in the same namespace as the Prometheus object, which shall be mounted into the Prometheus Pods. The Secrets are mounted into /etc/prometheus/secrets/<secret-name>.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"security_context": {
						Description:         "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",
						MarkdownDescription: "SecurityContext holds pod-level security attributes and common container settings. This defaults to the default PodSecurityContext.",

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

					"service_account_name": {
						Description:         "ServiceAccountName is the name of the ServiceAccount to use to run the Prometheus Pods.",
						MarkdownDescription: "ServiceAccountName is the name of the ServiceAccount to use to run the Prometheus Pods.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_monitor_namespace_selector": {
						Description:         "Namespace's labels to match for ServiceMonitor discovery. If nil, only check own namespace.",
						MarkdownDescription: "Namespace's labels to match for ServiceMonitor discovery. If nil, only check own namespace.",

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

					"service_monitor_selector": {
						Description:         "ServiceMonitors to be selected for target discovery. *Deprecated:* if neither this nor podMonitorSelector are specified, configuration is unmanaged.",
						MarkdownDescription: "ServiceMonitors to be selected for target discovery. *Deprecated:* if neither this nor podMonitorSelector are specified, configuration is unmanaged.",

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

					"sha": {
						Description:         "SHA of Prometheus container image to be deployed. Defaults to the value of 'version'. Similar to a tag, but the SHA explicitly deploys an immutable container image. Version and Tag are ignored if SHA is set. Deprecated: use 'image' instead.  The image digest can be specified as part of the image URL.",
						MarkdownDescription: "SHA of Prometheus container image to be deployed. Defaults to the value of 'version'. Similar to a tag, but the SHA explicitly deploys an immutable container image. Version and Tag are ignored if SHA is set. Deprecated: use 'image' instead.  The image digest can be specified as part of the image URL.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"shards": {
						Description:         "EXPERIMENTAL: Number of shards to distribute targets onto. Number of replicas multiplied by shards is the total number of Pods created. Note that scaling down shards will not reshard data onto remaining instances, it must be manually moved. Increasing shards will not reshard data either but it will continue to be available from the same instances. To query globally use Thanos sidecar and Thanos querier or remote write data to a central location. Sharding is done on the content of the '__address__' target meta-label.",
						MarkdownDescription: "EXPERIMENTAL: Number of shards to distribute targets onto. Number of replicas multiplied by shards is the total number of Pods created. Note that scaling down shards will not reshard data onto remaining instances, it must be manually moved. Increasing shards will not reshard data either but it will continue to be available from the same instances. To query globally use Thanos sidecar and Thanos querier or remote write data to a central location. Sharding is done on the content of the '__address__' target meta-label.",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"storage": {
						Description:         "Storage spec to specify how storage shall be used.",
						MarkdownDescription: "Storage spec to specify how storage shall be used.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"disable_mount_sub_path": {
								Description:         "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary. DisableMountSubPath allows to remove any subPath usage in volume mounts.",
								MarkdownDescription: "Deprecated: subPath usage will be disabled by default in a future release, this option will become unnecessary. DisableMountSubPath allows to remove any subPath usage in volume mounts.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"empty_dir": {
								Description:         "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",
								MarkdownDescription: "EmptyDirVolumeSource to be used by the Prometheus StatefulSets. If specified, used in place of any volumeClaimTemplate. More info: https://kubernetes.io/docs/concepts/storage/volumes/#emptydir",

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

							"ephemeral": {
								Description:         "EphemeralVolumeSource to be used by the Prometheus StatefulSets. This is a beta field in k8s 1.21, for lower versions, starting with k8s 1.19, it requires enabling the GenericEphemeralVolume feature gate. More info: https://kubernetes.io/docs/concepts/storage/ephemeral-volumes/#generic-ephemeral-volumes",
								MarkdownDescription: "EphemeralVolumeSource to be used by the Prometheus StatefulSets. This is a beta field in k8s 1.21, for lower versions, starting with k8s 1.19, it requires enabling the GenericEphemeralVolume feature gate. More info: https://kubernetes.io/docs/concepts/storage/ephemeral-volumes/#generic-ephemeral-volumes",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"volume_claim_template": {
										Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
										MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
												MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",

												Type: types.MapType{ElemType: types.StringType},

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

							"volume_claim_template": {
								Description:         "A PVC spec to be used by the Prometheus StatefulSets.",
								MarkdownDescription: "A PVC spec to be used by the Prometheus StatefulSets.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
										Description:         "EmbeddedMetadata contains metadata relevant to an EmbeddedResource.",
										MarkdownDescription: "EmbeddedMetadata contains metadata relevant to an EmbeddedResource.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",
												MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. They are not queryable and should be preserved when modifying objects. More info: http://kubernetes.io/docs/user-guide/annotations",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"labels": {
												Description:         "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",
												MarkdownDescription: "Map of string keys and values that can be used to organize and categorize (scope and select) objects. May match selectors of replication controllers and services. More info: http://kubernetes.io/docs/user-guide/labels",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": {
												Description:         "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",
												MarkdownDescription: "Name must be unique within a namespace. Is required when creating resources, although some resources may allow a client to request the generation of an appropriate name automatically. Name is primarily intended for creation idempotence and configuration definition. Cannot be updated. More info: http://kubernetes.io/docs/user-guide/identifiers#names",

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
										Description:         "Spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
										MarkdownDescription: "Spec defines the desired characteristics of a volume requested by a pod author. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

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
										Description:         "Status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
										MarkdownDescription: "Status represents the current information/status of a persistent volume claim. Read-only. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

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
													},

													"last_transition_time": {
														Description:         "lastTransitionTime is the time the condition transitioned from one status to another.",
														MarkdownDescription: "lastTransitionTime is the time the condition transitioned from one status to another.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tag": {
						Description:         "Tag of Prometheus container image to be deployed. Defaults to the value of 'version'. Version is ignored if Tag is set. Deprecated: use 'image' instead.  The image tag can be specified as part of the image URL.",
						MarkdownDescription: "Tag of Prometheus container image to be deployed. Defaults to the value of 'version'. Version is ignored if Tag is set. Deprecated: use 'image' instead.  The image tag can be specified as part of the image URL.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"thanos": {
						Description:         "Thanos configuration allows configuring various aspects of a Prometheus server in a Thanos environment.  This section is experimental, it may change significantly without deprecation notice in any release.  This is experimental and may change significantly without backward compatibility in any release.",
						MarkdownDescription: "Thanos configuration allows configuring various aspects of a Prometheus server in a Thanos environment.  This section is experimental, it may change significantly without deprecation notice in any release.  This is experimental and may change significantly without backward compatibility in any release.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"additional_args": {
								Description:         "AdditionalArgs allows setting additional arguments for the Thanos container. The arguments are passed as-is to the Thanos container which may cause issues if they are invalid or not supporeted the given Thanos version. In case of an argument conflict (e.g. an argument which is already set by the operator itself) or when providing an invalid argument the reconciliation will fail and an error will be logged.",
								MarkdownDescription: "AdditionalArgs allows setting additional arguments for the Thanos container. The arguments are passed as-is to the Thanos container which may cause issues if they are invalid or not supporeted the given Thanos version. In case of an argument conflict (e.g. an argument which is already set by the operator itself) or when providing an invalid argument the reconciliation will fail and an error will be logged.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "Name of the argument, e.g. 'scrape.discovery-reload-interval'.",
										MarkdownDescription: "Name of the argument, e.g. 'scrape.discovery-reload-interval'.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"value": {
										Description:         "Argument value, e.g. 30s. Can be empty for name-only arguments (e.g. --storage.tsdb.no-lockfile)",
										MarkdownDescription: "Argument value, e.g. 30s. Can be empty for name-only arguments (e.g. --storage.tsdb.no-lockfile)",

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

							"base_image": {
								Description:         "Thanos base image if other than default. Deprecated: use 'image' instead",
								MarkdownDescription: "Thanos base image if other than default. Deprecated: use 'image' instead",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"grpc_server_tls_config": {
								Description:         "GRPCServerTLSConfig configures the gRPC server from which Thanos Querier reads recorded rule data. Note: Currently only the CAFile, CertFile, and KeyFile fields are supported. Maps to the '--grpc-server-tls-*' CLI args.",
								MarkdownDescription: "GRPCServerTLSConfig configures the gRPC server from which Thanos Querier reads recorded rule data. Note: Currently only the CAFile, CertFile, and KeyFile fields are supported. Maps to the '--grpc-server-tls-*' CLI args.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ca": {
										Description:         "Struct containing the CA cert to use for the targets.",
										MarkdownDescription: "Struct containing the CA cert to use for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

									"ca_file": {
										Description:         "Path to the CA cert in the Prometheus container to use for the targets.",
										MarkdownDescription: "Path to the CA cert in the Prometheus container to use for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cert": {
										Description:         "Struct containing the client cert file for the targets.",
										MarkdownDescription: "Struct containing the client cert file for the targets.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

									"cert_file": {
										Description:         "Path to the client cert file in the Prometheus container for the targets.",
										MarkdownDescription: "Path to the client cert file in the Prometheus container for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure_skip_verify": {
										Description:         "Disable target certificate validation.",
										MarkdownDescription: "Disable target certificate validation.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_file": {
										Description:         "Path to the client key file in the Prometheus container for the targets.",
										MarkdownDescription: "Path to the client key file in the Prometheus container for the targets.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_secret": {
										Description:         "Secret containing the client key file for the targets.",
										MarkdownDescription: "Secret containing the client key file for the targets.",

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

									"server_name": {
										Description:         "Used to verify the hostname for the targets.",
										MarkdownDescription: "Used to verify the hostname for the targets.",

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

							"image": {
								Description:         "Image if specified has precedence over baseImage, tag and sha combinations. Specifying the version is still necessary to ensure the Prometheus Operator knows what version of Thanos is being configured.",
								MarkdownDescription: "Image if specified has precedence over baseImage, tag and sha combinations. Specifying the version is still necessary to ensure the Prometheus Operator knows what version of Thanos is being configured.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"listen_local": {
								Description:         "ListenLocal makes the Thanos sidecar listen on loopback, so that it does not bind against the Pod IP.",
								MarkdownDescription: "ListenLocal makes the Thanos sidecar listen on loopback, so that it does not bind against the Pod IP.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_format": {
								Description:         "LogFormat for Thanos sidecar to be configured with.",
								MarkdownDescription: "LogFormat for Thanos sidecar to be configured with.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_level": {
								Description:         "LogLevel for Thanos sidecar to be configured with.",
								MarkdownDescription: "LogLevel for Thanos sidecar to be configured with.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"min_time": {
								Description:         "MinTime for Thanos sidecar to be configured with. Option can be a constant time in RFC3339 format or time duration relative to current time, such as -1d or 2h45m. Valid duration units are ms, s, m, h, d, w, y.",
								MarkdownDescription: "MinTime for Thanos sidecar to be configured with. Option can be a constant time in RFC3339 format or time duration relative to current time, such as -1d or 2h45m. Valid duration units are ms, s, m, h, d, w, y.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"object_storage_config": {
								Description:         "ObjectStorageConfig configures object storage in Thanos. Alternative to ObjectStorageConfigFile, and lower order priority.",
								MarkdownDescription: "ObjectStorageConfig configures object storage in Thanos. Alternative to ObjectStorageConfigFile, and lower order priority.",

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

							"object_storage_config_file": {
								Description:         "ObjectStorageConfigFile specifies the path of the object storage configuration file. When used alongside with ObjectStorageConfig, ObjectStorageConfigFile takes precedence.",
								MarkdownDescription: "ObjectStorageConfigFile specifies the path of the object storage configuration file. When used alongside with ObjectStorageConfig, ObjectStorageConfigFile takes precedence.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ready_timeout": {
								Description:         "ReadyTimeout is the maximum time Thanos sidecar will wait for Prometheus to start. Eg 10m",
								MarkdownDescription: "ReadyTimeout is the maximum time Thanos sidecar will wait for Prometheus to start. Eg 10m",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources defines the resource requirements for the Thanos sidecar. If not provided, no requests/limits will be set",
								MarkdownDescription: "Resources defines the resource requirements for the Thanos sidecar. If not provided, no requests/limits will be set",

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

							"sha": {
								Description:         "SHA of Thanos container image to be deployed. Defaults to the value of 'version'. Similar to a tag, but the SHA explicitly deploys an immutable container image. Version and Tag are ignored if SHA is set. Deprecated: use 'image' instead.  The image digest can be specified as part of the image URL.",
								MarkdownDescription: "SHA of Thanos container image to be deployed. Defaults to the value of 'version'. Similar to a tag, but the SHA explicitly deploys an immutable container image. Version and Tag are ignored if SHA is set. Deprecated: use 'image' instead.  The image digest can be specified as part of the image URL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tag": {
								Description:         "Tag of Thanos sidecar container image to be deployed. Defaults to the value of 'version'. Version is ignored if Tag is set. Deprecated: use 'image' instead.  The image tag can be specified as part of the image URL.",
								MarkdownDescription: "Tag of Thanos sidecar container image to be deployed. Defaults to the value of 'version'. Version is ignored if Tag is set. Deprecated: use 'image' instead.  The image tag can be specified as part of the image URL.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tracing_config": {
								Description:         "TracingConfig configures tracing in Thanos. This is an experimental feature, it may change in any upcoming release in a breaking way.",
								MarkdownDescription: "TracingConfig configures tracing in Thanos. This is an experimental feature, it may change in any upcoming release in a breaking way.",

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

							"tracing_config_file": {
								Description:         "TracingConfig specifies the path of the tracing configuration file. When used alongside with TracingConfig, TracingConfigFile takes precedence.",
								MarkdownDescription: "TracingConfig specifies the path of the tracing configuration file. When used alongside with TracingConfig, TracingConfigFile takes precedence.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Version describes the version of Thanos to use.",
								MarkdownDescription: "Version describes the version of Thanos to use.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_mounts": {
								Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the thanos-sidecar container.",
								MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the thanos-sidecar container.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tolerations": {
						Description:         "If specified, the pod's tolerations.",
						MarkdownDescription: "If specified, the pod's tolerations.",

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

					"topology_spread_constraints": {
						Description:         "If specified, the pod's topology spread constraints.",
						MarkdownDescription: "If specified, the pod's topology spread constraints.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"label_selector": {
								Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
								MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",

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

							"match_label_keys": {
								Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.",
								MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"max_skew": {
								Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
								MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"min_domains": {
								Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
								MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_affinity_policy": {
								Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
								MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_taints_policy": {
								Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
								MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"topology_key": {
								Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
								MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"when_unsatisfiable": {
								Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
								MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",

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

					"version": {
						Description:         "Version of Prometheus to be deployed.",
						MarkdownDescription: "Version of Prometheus to be deployed.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"volume_mounts": {
						Description:         "VolumeMounts allows configuration of additional VolumeMounts on the output StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the prometheus container, that are generated as a result of StorageSpec objects.",
						MarkdownDescription: "VolumeMounts allows configuration of additional VolumeMounts on the output StatefulSet definition. VolumeMounts specified will be appended to other VolumeMounts in the prometheus container, that are generated as a result of StorageSpec objects.",

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
						Description:         "Volumes allows configuration of additional volumes on the output StatefulSet definition. Volumes specified will be appended to other volumes that are generated as a result of StorageSpec objects.",
						MarkdownDescription: "Volumes allows configuration of additional volumes on the output StatefulSet definition. Volumes specified will be appended to other volumes that are generated as a result of StorageSpec objects.",

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

														Type: types.StringType,

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

												Type: types.MapType{ElemType: types.StringType},

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

																		Type: types.StringType,

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

					"wal_compression": {
						Description:         "Enable compression of the write-ahead log using Snappy. This flag is only available in versions of Prometheus >= 2.11.0.",
						MarkdownDescription: "Enable compression of the write-ahead log using Snappy. This flag is only available in versions of Prometheus >= 2.11.0.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"web": {
						Description:         "Defines the web command line flags when starting Prometheus.",
						MarkdownDescription: "Defines the web command line flags when starting Prometheus.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"http_config": {
								Description:         "Defines HTTP parameters for web server.",
								MarkdownDescription: "Defines HTTP parameters for web server.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"headers": {
										Description:         "List of headers that can be added to HTTP responses.",
										MarkdownDescription: "List of headers that can be added to HTTP responses.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"content_security_policy": {
												Description:         "Set the Content-Security-Policy header to HTTP responses. Unset if blank.",
												MarkdownDescription: "Set the Content-Security-Policy header to HTTP responses. Unset if blank.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"strict_transport_security": {
												Description:         "Set the Strict-Transport-Security header to HTTP responses. Unset if blank. Please make sure that you use this with care as this header might force browsers to load Prometheus and the other applications hosted on the same domain and subdomains over HTTPS. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security",
												MarkdownDescription: "Set the Strict-Transport-Security header to HTTP responses. Unset if blank. Please make sure that you use this with care as this header might force browsers to load Prometheus and the other applications hosted on the same domain and subdomains over HTTPS. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Strict-Transport-Security",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"x_content_type_options": {
												Description:         "Set the X-Content-Type-Options header to HTTP responses. Unset if blank. Accepted value is nosniff. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options",
												MarkdownDescription: "Set the X-Content-Type-Options header to HTTP responses. Unset if blank. Accepted value is nosniff. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"x_frame_options": {
												Description:         "Set the X-Frame-Options header to HTTP responses. Unset if blank. Accepted values are deny and sameorigin. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options",
												MarkdownDescription: "Set the X-Frame-Options header to HTTP responses. Unset if blank. Accepted values are deny and sameorigin. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"x_xss_protection": {
												Description:         "Set the X-XSS-Protection header to all responses. Unset if blank. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection",
												MarkdownDescription: "Set the X-XSS-Protection header to all responses. Unset if blank. https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection",

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

									"http2": {
										Description:         "Enable HTTP/2 support. Note that HTTP/2 is only supported with TLS. When TLSConfig is not configured, HTTP/2 will be disabled. Whenever the value of the field changes, a rolling update will be triggered.",
										MarkdownDescription: "Enable HTTP/2 support. Note that HTTP/2 is only supported with TLS. When TLSConfig is not configured, HTTP/2 will be disabled. Whenever the value of the field changes, a rolling update will be triggered.",

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

							"page_title": {
								Description:         "The prometheus web page title",
								MarkdownDescription: "The prometheus web page title",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_config": {
								Description:         "Defines the TLS parameters for HTTPS.",
								MarkdownDescription: "Defines the TLS parameters for HTTPS.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"cert": {
										Description:         "Contains the TLS certificate for the server.",
										MarkdownDescription: "Contains the TLS certificate for the server.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

										Required: true,
										Optional: false,
										Computed: false,
									},

									"cipher_suites": {
										Description:         "List of supported cipher suites for TLS versions up to TLS 1.2. If empty, Go default cipher suites are used. Available cipher suites are documented in the go documentation: https://golang.org/pkg/crypto/tls/#pkg-constants",
										MarkdownDescription: "List of supported cipher suites for TLS versions up to TLS 1.2. If empty, Go default cipher suites are used. Available cipher suites are documented in the go documentation: https://golang.org/pkg/crypto/tls/#pkg-constants",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_auth_type": {
										Description:         "Server policy for client authentication. Maps to ClientAuth Policies. For more detail on clientAuth options: https://golang.org/pkg/crypto/tls/#ClientAuthType",
										MarkdownDescription: "Server policy for client authentication. Maps to ClientAuth Policies. For more detail on clientAuth options: https://golang.org/pkg/crypto/tls/#ClientAuthType",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_ca": {
										Description:         "Contains the CA certificate for client certificate authentication to the server.",
										MarkdownDescription: "Contains the CA certificate for client certificate authentication to the server.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_map": {
												Description:         "ConfigMap containing data to use for the targets.",
												MarkdownDescription: "ConfigMap containing data to use for the targets.",

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

											"secret": {
												Description:         "Secret containing data to use for the targets.",
												MarkdownDescription: "Secret containing data to use for the targets.",

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

									"curve_preferences": {
										Description:         "Elliptic curves that will be used in an ECDHE handshake, in preference order. Available curves are documented in the go documentation: https://golang.org/pkg/crypto/tls/#CurveID",
										MarkdownDescription: "Elliptic curves that will be used in an ECDHE handshake, in preference order. Available curves are documented in the go documentation: https://golang.org/pkg/crypto/tls/#CurveID",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"key_secret": {
										Description:         "Secret containing the TLS key for the server.",
										MarkdownDescription: "Secret containing the TLS key for the server.",

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

										Required: true,
										Optional: false,
										Computed: false,
									},

									"max_version": {
										Description:         "Maximum TLS version that is acceptable. Defaults to TLS13.",
										MarkdownDescription: "Maximum TLS version that is acceptable. Defaults to TLS13.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"min_version": {
										Description:         "Minimum TLS version that is acceptable. Defaults to TLS12.",
										MarkdownDescription: "Minimum TLS version that is acceptable. Defaults to TLS12.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prefer_server_cipher_suites": {
										Description:         "Controls whether the server selects the client's most preferred cipher suite, or the server's most preferred cipher suite. If true then the server's preference, as expressed in the order of elements in cipherSuites, is used.",
										MarkdownDescription: "Controls whether the server selects the client's most preferred cipher suite, or the server's most preferred cipher suite. If true then the server's preference, as expressed in the order of elements in cipherSuites, is used.",

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

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *MonitoringCoreosComPrometheusV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_monitoring_coreos_com_prometheus_v1")

	var state MonitoringCoreosComPrometheusV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MonitoringCoreosComPrometheusV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("monitoring.coreos.com/v1")
	goModel.Kind = utilities.Ptr("Prometheus")

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

func (r *MonitoringCoreosComPrometheusV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_monitoring_coreos_com_prometheus_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *MonitoringCoreosComPrometheusV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_monitoring_coreos_com_prometheus_v1")

	var state MonitoringCoreosComPrometheusV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel MonitoringCoreosComPrometheusV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("monitoring.coreos.com/v1")
	goModel.Kind = utilities.Ptr("Prometheus")

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

func (r *MonitoringCoreosComPrometheusV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_monitoring_coreos_com_prometheus_v1")
	// NO-OP: Terraform removes the state automatically for us
}
