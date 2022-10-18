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

type ArgoprojIoArgoCDV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*ArgoprojIoArgoCDV1Alpha1Resource)(nil)
)

type ArgoprojIoArgoCDV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ArgoprojIoArgoCDV1Alpha1GoModel struct {
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
		ApplicationInstanceLabelKey *string `tfsdk:"application_instance_label_key" yaml:"applicationInstanceLabelKey,omitempty"`

		ApplicationSet *struct {
			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`

			WebhookServer *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Ingress *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

					IngressClassName *string `tfsdk:"ingress_class_name" yaml:"ingressClassName,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Tls *[]struct {
						Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"tls" yaml:"tls,omitempty"`
				} `tfsdk:"ingress" yaml:"ingress,omitempty"`

				Route *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

					Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Tls *struct {
						CaCertificate *string `tfsdk:"ca_certificate" yaml:"caCertificate,omitempty"`

						Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

						DestinationCACertificate *string `tfsdk:"destination_ca_certificate" yaml:"destinationCACertificate,omitempty"`

						InsecureEdgeTerminationPolicy *string `tfsdk:"insecure_edge_termination_policy" yaml:"insecureEdgeTerminationPolicy,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Termination *string `tfsdk:"termination" yaml:"termination,omitempty"`
					} `tfsdk:"tls" yaml:"tls,omitempty"`

					WildcardPolicy *string `tfsdk:"wildcard_policy" yaml:"wildcardPolicy,omitempty"`
				} `tfsdk:"route" yaml:"route,omitempty"`
			} `tfsdk:"webhook_server" yaml:"webhookServer,omitempty"`
		} `tfsdk:"application_set" yaml:"applicationSet,omitempty"`

		Banner *struct {
			Content *string `tfsdk:"content" yaml:"content,omitempty"`

			Url *string `tfsdk:"url" yaml:"url,omitempty"`
		} `tfsdk:"banner" yaml:"banner,omitempty"`

		ConfigManagementPlugins *string `tfsdk:"config_management_plugins" yaml:"configManagementPlugins,omitempty"`

		Controller *struct {
			AppSync *string `tfsdk:"app_sync" yaml:"appSync,omitempty"`

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

			LogFormat *string `tfsdk:"log_format" yaml:"logFormat,omitempty"`

			LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

			ParallelismLimit *int64 `tfsdk:"parallelism_limit" yaml:"parallelismLimit,omitempty"`

			Processors *struct {
				Operation *int64 `tfsdk:"operation" yaml:"operation,omitempty"`

				Status *int64 `tfsdk:"status" yaml:"status,omitempty"`
			} `tfsdk:"processors" yaml:"processors,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Sharding *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`
			} `tfsdk:"sharding" yaml:"sharding,omitempty"`
		} `tfsdk:"controller" yaml:"controller,omitempty"`

		Dex *struct {
			Config *string `tfsdk:"config" yaml:"config,omitempty"`

			Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			OpenShiftOAuth *bool `tfsdk:"open_shift_o_auth" yaml:"openShiftOAuth,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"dex" yaml:"dex,omitempty"`

		DisableAdmin *bool `tfsdk:"disable_admin" yaml:"disableAdmin,omitempty"`

		ExtraConfig *map[string]string `tfsdk:"extra_config" yaml:"extraConfig,omitempty"`

		GaAnonymizeUsers *bool `tfsdk:"ga_anonymize_users" yaml:"gaAnonymizeUsers,omitempty"`

		GaTrackingID *string `tfsdk:"ga_tracking_id" yaml:"gaTrackingID,omitempty"`

		Grafana *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			Ingress *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				IngressClassName *string `tfsdk:"ingress_class_name" yaml:"ingressClassName,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Tls *[]struct {
					Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`
			} `tfsdk:"ingress" yaml:"ingress,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Route *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Tls *struct {
					CaCertificate *string `tfsdk:"ca_certificate" yaml:"caCertificate,omitempty"`

					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					DestinationCACertificate *string `tfsdk:"destination_ca_certificate" yaml:"destinationCACertificate,omitempty"`

					InsecureEdgeTerminationPolicy *string `tfsdk:"insecure_edge_termination_policy" yaml:"insecureEdgeTerminationPolicy,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Termination *string `tfsdk:"termination" yaml:"termination,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`

				WildcardPolicy *string `tfsdk:"wildcard_policy" yaml:"wildcardPolicy,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`

			Size *int64 `tfsdk:"size" yaml:"size,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"grafana" yaml:"grafana,omitempty"`

		Ha *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			RedisProxyImage *string `tfsdk:"redis_proxy_image" yaml:"redisProxyImage,omitempty"`

			RedisProxyVersion *string `tfsdk:"redis_proxy_version" yaml:"redisProxyVersion,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`
		} `tfsdk:"ha" yaml:"ha,omitempty"`

		HelpChatText *string `tfsdk:"help_chat_text" yaml:"helpChatText,omitempty"`

		HelpChatURL *string `tfsdk:"help_chat_url" yaml:"helpChatURL,omitempty"`

		Image *string `tfsdk:"image" yaml:"image,omitempty"`

		Import *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
		} `tfsdk:"import" yaml:"import,omitempty"`

		InitialRepositories *string `tfsdk:"initial_repositories" yaml:"initialRepositories,omitempty"`

		InitialSSHKnownHosts *struct {
			Excludedefaulthosts *bool `tfsdk:"excludedefaulthosts" yaml:"excludedefaulthosts,omitempty"`

			Keys *string `tfsdk:"keys" yaml:"keys,omitempty"`
		} `tfsdk:"initial_ssh_known_hosts" yaml:"initialSSHKnownHosts,omitempty"`

		KustomizeBuildOptions *string `tfsdk:"kustomize_build_options" yaml:"kustomizeBuildOptions,omitempty"`

		KustomizeVersions *[]struct {
			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"kustomize_versions" yaml:"kustomizeVersions,omitempty"`

		NodePlacement *struct {
			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

			Tolerations *[]struct {
				Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

				Key *string `tfsdk:"key" yaml:"key,omitempty"`

				Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

				TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

				Value *string `tfsdk:"value" yaml:"value,omitempty"`
			} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
		} `tfsdk:"node_placement" yaml:"nodePlacement,omitempty"`

		Notifications *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

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

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"notifications" yaml:"notifications,omitempty"`

		OidcConfig *string `tfsdk:"oidc_config" yaml:"oidcConfig,omitempty"`

		Prometheus *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Ingress *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				IngressClassName *string `tfsdk:"ingress_class_name" yaml:"ingressClassName,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Tls *[]struct {
					Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`
			} `tfsdk:"ingress" yaml:"ingress,omitempty"`

			Route *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Tls *struct {
					CaCertificate *string `tfsdk:"ca_certificate" yaml:"caCertificate,omitempty"`

					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					DestinationCACertificate *string `tfsdk:"destination_ca_certificate" yaml:"destinationCACertificate,omitempty"`

					InsecureEdgeTerminationPolicy *string `tfsdk:"insecure_edge_termination_policy" yaml:"insecureEdgeTerminationPolicy,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Termination *string `tfsdk:"termination" yaml:"termination,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`

				WildcardPolicy *string `tfsdk:"wildcard_policy" yaml:"wildcardPolicy,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`

			Size *int64 `tfsdk:"size" yaml:"size,omitempty"`
		} `tfsdk:"prometheus" yaml:"prometheus,omitempty"`

		Rbac *struct {
			DefaultPolicy *string `tfsdk:"default_policy" yaml:"defaultPolicy,omitempty"`

			Policy *string `tfsdk:"policy" yaml:"policy,omitempty"`

			PolicyMatcherMode *string `tfsdk:"policy_matcher_mode" yaml:"policyMatcherMode,omitempty"`

			Scopes *string `tfsdk:"scopes" yaml:"scopes,omitempty"`
		} `tfsdk:"rbac" yaml:"rbac,omitempty"`

		Redis *struct {
			Autotls *string `tfsdk:"autotls" yaml:"autotls,omitempty"`

			DisableTLSVerification *bool `tfsdk:"disable_tls_verification" yaml:"disableTLSVerification,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"redis" yaml:"redis,omitempty"`

		Repo *struct {
			Autotls *string `tfsdk:"autotls" yaml:"autotls,omitempty"`

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

			ExecTimeout *int64 `tfsdk:"exec_timeout" yaml:"execTimeout,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

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

			LogFormat *string `tfsdk:"log_format" yaml:"logFormat,omitempty"`

			LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

			Mountsatoken *bool `tfsdk:"mountsatoken" yaml:"mountsatoken,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Serviceaccount *string `tfsdk:"serviceaccount" yaml:"serviceaccount,omitempty"`

			SidecarContainers *[]struct {
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
			} `tfsdk:"sidecar_containers" yaml:"sidecarContainers,omitempty"`

			Verifytls *bool `tfsdk:"verifytls" yaml:"verifytls,omitempty"`

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

							Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

							Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`
					} `tfsdk:"items" yaml:"items,omitempty"`
				} `tfsdk:"downward_api" yaml:"downwardAPI,omitempty"`

				EmptyDir *struct {
					Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

					SizeLimit utilities.IntOrString `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
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

									Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

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
		} `tfsdk:"repo" yaml:"repo,omitempty"`

		RepositoryCredentials *string `tfsdk:"repository_credentials" yaml:"repositoryCredentials,omitempty"`

		ResourceCustomizations *string `tfsdk:"resource_customizations" yaml:"resourceCustomizations,omitempty"`

		ResourceExclusions *string `tfsdk:"resource_exclusions" yaml:"resourceExclusions,omitempty"`

		ResourceInclusions *string `tfsdk:"resource_inclusions" yaml:"resourceInclusions,omitempty"`

		ResourceTrackingMethod *string `tfsdk:"resource_tracking_method" yaml:"resourceTrackingMethod,omitempty"`

		Server *struct {
			Autoscale *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Hpa *struct {
					MaxReplicas *int64 `tfsdk:"max_replicas" yaml:"maxReplicas,omitempty"`

					MinReplicas *int64 `tfsdk:"min_replicas" yaml:"minReplicas,omitempty"`

					ScaleTargetRef *struct {
						ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

						Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"scale_target_ref" yaml:"scaleTargetRef,omitempty"`

					TargetCPUUtilizationPercentage *int64 `tfsdk:"target_cpu_utilization_percentage" yaml:"targetCPUUtilizationPercentage,omitempty"`
				} `tfsdk:"hpa" yaml:"hpa,omitempty"`
			} `tfsdk:"autoscale" yaml:"autoscale,omitempty"`

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

			ExtraCommandArgs *[]string `tfsdk:"extra_command_args" yaml:"extraCommandArgs,omitempty"`

			Grpc *struct {
				Host *string `tfsdk:"host" yaml:"host,omitempty"`

				Ingress *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

					IngressClassName *string `tfsdk:"ingress_class_name" yaml:"ingressClassName,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Tls *[]struct {
						Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

						SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
					} `tfsdk:"tls" yaml:"tls,omitempty"`
				} `tfsdk:"ingress" yaml:"ingress,omitempty"`
			} `tfsdk:"grpc" yaml:"grpc,omitempty"`

			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Ingress *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				IngressClassName *string `tfsdk:"ingress_class_name" yaml:"ingressClassName,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Tls *[]struct {
					Hosts *[]string `tfsdk:"hosts" yaml:"hosts,omitempty"`

					SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`
			} `tfsdk:"ingress" yaml:"ingress,omitempty"`

			Insecure *bool `tfsdk:"insecure" yaml:"insecure,omitempty"`

			LogFormat *string `tfsdk:"log_format" yaml:"logFormat,omitempty"`

			LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			Route *struct {
				Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Tls *struct {
					CaCertificate *string `tfsdk:"ca_certificate" yaml:"caCertificate,omitempty"`

					Certificate *string `tfsdk:"certificate" yaml:"certificate,omitempty"`

					DestinationCACertificate *string `tfsdk:"destination_ca_certificate" yaml:"destinationCACertificate,omitempty"`

					InsecureEdgeTerminationPolicy *string `tfsdk:"insecure_edge_termination_policy" yaml:"insecureEdgeTerminationPolicy,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Termination *string `tfsdk:"termination" yaml:"termination,omitempty"`
				} `tfsdk:"tls" yaml:"tls,omitempty"`

				WildcardPolicy *string `tfsdk:"wildcard_policy" yaml:"wildcardPolicy,omitempty"`
			} `tfsdk:"route" yaml:"route,omitempty"`

			Service *struct {
				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"service" yaml:"service,omitempty"`
		} `tfsdk:"server" yaml:"server,omitempty"`

		Sso *struct {
			Dex *struct {
				Config *string `tfsdk:"config" yaml:"config,omitempty"`

				Groups *[]string `tfsdk:"groups" yaml:"groups,omitempty"`

				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				OpenShiftOAuth *bool `tfsdk:"open_shift_o_auth" yaml:"openShiftOAuth,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"dex" yaml:"dex,omitempty"`

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			Keycloak *struct {
				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				RootCA *string `tfsdk:"root_ca" yaml:"rootCA,omitempty"`

				VerifyTLS *bool `tfsdk:"verify_tls" yaml:"verifyTLS,omitempty"`

				Version *string `tfsdk:"version" yaml:"version,omitempty"`
			} `tfsdk:"keycloak" yaml:"keycloak,omitempty"`

			Provider *string `tfsdk:"provider" yaml:"provider,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			VerifyTLS *bool `tfsdk:"verify_tls" yaml:"verifyTLS,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"sso" yaml:"sso,omitempty"`

		StatusBadgeEnabled *bool `tfsdk:"status_badge_enabled" yaml:"statusBadgeEnabled,omitempty"`

		Tls *struct {
			Ca *struct {
				ConfigMapName *string `tfsdk:"config_map_name" yaml:"configMapName,omitempty"`

				SecretName *string `tfsdk:"secret_name" yaml:"secretName,omitempty"`
			} `tfsdk:"ca" yaml:"ca,omitempty"`

			InitialCerts *map[string]string `tfsdk:"initial_certs" yaml:"initialCerts,omitempty"`
		} `tfsdk:"tls" yaml:"tls,omitempty"`

		UsersAnonymousEnabled *bool `tfsdk:"users_anonymous_enabled" yaml:"usersAnonymousEnabled,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewArgoprojIoArgoCDV1Alpha1Resource() resource.Resource {
	return &ArgoprojIoArgoCDV1Alpha1Resource{}
}

func (r *ArgoprojIoArgoCDV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_argoproj_io_argo_cd_v1alpha1"
}

func (r *ArgoprojIoArgoCDV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ArgoCD is the Schema for the argocds API",
		MarkdownDescription: "ArgoCD is the Schema for the argocds API",
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
				Description:         "ArgoCDSpec defines the desired state of ArgoCD",
				MarkdownDescription: "ArgoCDSpec defines the desired state of ArgoCD",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"application_instance_label_key": {
						Description:         "ApplicationInstanceLabelKey is the key name where Argo CD injects the app name as a tracking label.",
						MarkdownDescription: "ApplicationInstanceLabelKey is the key name where Argo CD injects the app name as a tracking label.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"application_set": {
						Description:         "ArgoCDApplicationSet defines whether the Argo CD ApplicationSet controller should be installed.",
						MarkdownDescription: "ArgoCDApplicationSet defines whether the Argo CD ApplicationSet controller should be installed.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"image": {
								Description:         "Image is the Argo CD ApplicationSet image (optional)",
								MarkdownDescription: "Image is the Argo CD ApplicationSet image (optional)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_level": {
								Description:         "LogLevel describes the log level that should be used by the ApplicationSet controller. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug,info, error, and warn.",
								MarkdownDescription: "LogLevel describes the log level that should be used by the ApplicationSet controller. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug,info, error, and warn.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources defines the Compute Resources required by the container for ApplicationSet.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for ApplicationSet.",

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

							"version": {
								Description:         "Version is the Argo CD ApplicationSet image tag. (optional)",
								MarkdownDescription: "Version is the Argo CD ApplicationSet image tag. (optional)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"webhook_server": {
								Description:         "WebhookServerSpec defines the options for the ApplicationSet Webhook Server component.",
								MarkdownDescription: "WebhookServerSpec defines the options for the ApplicationSet Webhook Server component.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Host is the hostname to use for Ingress/Route resources.",
										MarkdownDescription: "Host is the hostname to use for Ingress/Route resources.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ingress": {
										Description:         "Ingress defines the desired state for an Ingress for the Application set webhook component.",
										MarkdownDescription: "Ingress defines the desired state for an Ingress for the Application set webhook component.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations is the map of annotations to apply to the Ingress.",
												MarkdownDescription: "Annotations is the map of annotations to apply to the Ingress.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enabled": {
												Description:         "Enabled will toggle the creation of the Ingress.",
												MarkdownDescription: "Enabled will toggle the creation of the Ingress.",

												Type: types.BoolType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"ingress_class_name": {
												Description:         "IngressClassName for the Ingress resource.",
												MarkdownDescription: "IngressClassName for the Ingress resource.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Path used for the Ingress resource.",
												MarkdownDescription: "Path used for the Ingress resource.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls": {
												Description:         "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
												MarkdownDescription: "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"hosts": {
														Description:         "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
														MarkdownDescription: "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_name": {
														Description:         "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
														MarkdownDescription: "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",

														Type: types.StringType,

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

									"route": {
										Description:         "Route defines the desired state for an OpenShift Route for the Application set webhook component.",
										MarkdownDescription: "Route defines the desired state for an OpenShift Route for the Application set webhook component.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations is the map of annotations to use for the Route resource.",
												MarkdownDescription: "Annotations is the map of annotations to use for the Route resource.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enabled": {
												Description:         "Enabled will toggle the creation of the OpenShift Route.",
												MarkdownDescription: "Enabled will toggle the creation of the OpenShift Route.",

												Type: types.BoolType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"labels": {
												Description:         "Labels is the map of labels to use for the Route resource",
												MarkdownDescription: "Labels is the map of labels to use for the Route resource",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Path the router watches for, to route traffic for to the service.",
												MarkdownDescription: "Path the router watches for, to route traffic for to the service.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls": {
												Description:         "TLS provides the ability to configure certificates and termination for the Route.",
												MarkdownDescription: "TLS provides the ability to configure certificates and termination for the Route.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"ca_certificate": {
														Description:         "caCertificate provides the cert authority certificate contents",
														MarkdownDescription: "caCertificate provides the cert authority certificate contents",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"certificate": {
														Description:         "certificate provides certificate contents",
														MarkdownDescription: "certificate provides certificate contents",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"destination_ca_certificate": {
														Description:         "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",
														MarkdownDescription: "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_edge_termination_policy": {
														Description:         "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (default) * Disable - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",
														MarkdownDescription: "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (default) * Disable - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": {
														Description:         "key provides key file contents",
														MarkdownDescription: "key provides key file contents",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"termination": {
														Description:         "termination indicates termination type.",
														MarkdownDescription: "termination indicates termination type.",

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

											"wildcard_policy": {
												Description:         "WildcardPolicy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",
												MarkdownDescription: "WildcardPolicy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",

												Type: types.StringType,

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

					"banner": {
						Description:         "Banner defines an additional banner to be displayed in Argo CD UI",
						MarkdownDescription: "Banner defines an additional banner to be displayed in Argo CD UI",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"content": {
								Description:         "Content defines the banner message content to display",
								MarkdownDescription: "Content defines the banner message content to display",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"url": {
								Description:         "URL defines an optional URL to be used as banner message link",
								MarkdownDescription: "URL defines an optional URL to be used as banner message link",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"config_management_plugins": {
						Description:         "ConfigManagementPlugins is used to specify additional config management plugins.",
						MarkdownDescription: "ConfigManagementPlugins is used to specify additional config management plugins.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"controller": {
						Description:         "Controller defines the Application Controller options for ArgoCD.",
						MarkdownDescription: "Controller defines the Application Controller options for ArgoCD.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"app_sync": {
								Description:         "AppSync is used to control the sync frequency, by default the ArgoCD controller polls Git every 3m.  Set this to a duration, e.g. 10m or 600s to control the synchronisation frequency.",
								MarkdownDescription: "AppSync is used to control the sync frequency, by default the ArgoCD controller polls Git every 3m.  Set this to a duration, e.g. 10m or 600s to control the synchronisation frequency.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"env": {
								Description:         "Env lets you specify environment for application controller pods",
								MarkdownDescription: "Env lets you specify environment for application controller pods",

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

							"log_format": {
								Description:         "LogFormat refers to the log format used by the Application Controller component. Defaults to ArgoCDDefaultLogFormat if not configured. Valid options are text or json.",
								MarkdownDescription: "LogFormat refers to the log format used by the Application Controller component. Defaults to ArgoCDDefaultLogFormat if not configured. Valid options are text or json.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_level": {
								Description:         "LogLevel refers to the log level used by the Application Controller component. Defaults to ArgoCDDefaultLogLevel if not configured. Valid options are debug, info, error, and warn.",
								MarkdownDescription: "LogLevel refers to the log level used by the Application Controller component. Defaults to ArgoCDDefaultLogLevel if not configured. Valid options are debug, info, error, and warn.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"parallelism_limit": {
								Description:         "ParallelismLimit defines the limit for parallel kubectl operations",
								MarkdownDescription: "ParallelismLimit defines the limit for parallel kubectl operations",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"processors": {
								Description:         "Processors contains the options for the Application Controller processors.",
								MarkdownDescription: "Processors contains the options for the Application Controller processors.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"operation": {
										Description:         "Operation is the number of application operation processors.",
										MarkdownDescription: "Operation is the number of application operation processors.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"status": {
										Description:         "Status is the number of application status processors.",
										MarkdownDescription: "Status is the number of application status processors.",

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
								Description:         "Resources defines the Compute Resources required by the container for the Application Controller.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for the Application Controller.",

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

							"sharding": {
								Description:         "Sharding contains the options for the Application Controller sharding configuration.",
								MarkdownDescription: "Sharding contains the options for the Application Controller sharding configuration.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Enabled defines whether sharding should be enabled on the Application Controller component.",
										MarkdownDescription: "Enabled defines whether sharding should be enabled on the Application Controller component.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"replicas": {
										Description:         "Replicas defines the number of replicas to run in the Application controller shard.",
										MarkdownDescription: "Replicas defines the number of replicas to run in the Application controller shard.",

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

					"dex": {
						Description:         "Dex defines the Dex server options for ArgoCD.",
						MarkdownDescription: "Dex defines the Dex server options for ArgoCD.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"config": {
								Description:         "Config is the dex connector configuration.",
								MarkdownDescription: "Config is the dex connector configuration.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"groups": {
								Description:         "Optional list of required groups a user must be a member of",
								MarkdownDescription: "Optional list of required groups a user must be a member of",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "Image is the Dex container image.",
								MarkdownDescription: "Image is the Dex container image.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"open_shift_o_auth": {
								Description:         "OpenShiftOAuth enables OpenShift OAuth authentication for the Dex server.",
								MarkdownDescription: "OpenShiftOAuth enables OpenShift OAuth authentication for the Dex server.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources defines the Compute Resources required by the container for Dex.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for Dex.",

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

							"version": {
								Description:         "Version is the Dex container image tag.",
								MarkdownDescription: "Version is the Dex container image tag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"disable_admin": {
						Description:         "DisableAdmin will disable the admin user.",
						MarkdownDescription: "DisableAdmin will disable the admin user.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"extra_config": {
						Description:         "ExtraConfig can be used to add fields to Argo CD configmap that are not supported by Argo CD CRD.  Note: ExtraConfig takes precedence over Argo CD CRD. For example, A user sets 'argocd.Spec.DisableAdmin' = true and also 'a.Spec.ExtraConfig['admin.enabled']' = true. In this case, operator updates Argo CD Configmap as follows -> argocd-cm.Data['admin.enabled'] = true.",
						MarkdownDescription: "ExtraConfig can be used to add fields to Argo CD configmap that are not supported by Argo CD CRD.  Note: ExtraConfig takes precedence over Argo CD CRD. For example, A user sets 'argocd.Spec.DisableAdmin' = true and also 'a.Spec.ExtraConfig['admin.enabled']' = true. In this case, operator updates Argo CD Configmap as follows -> argocd-cm.Data['admin.enabled'] = true.",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ga_anonymize_users": {
						Description:         "GAAnonymizeUsers toggles user IDs being hashed before sending to google analytics.",
						MarkdownDescription: "GAAnonymizeUsers toggles user IDs being hashed before sending to google analytics.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ga_tracking_id": {
						Description:         "GATrackingID is the google analytics tracking ID to use.",
						MarkdownDescription: "GATrackingID is the google analytics tracking ID to use.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"grafana": {
						Description:         "Grafana defines the Grafana server options for ArgoCD.",
						MarkdownDescription: "Grafana defines the Grafana server options for ArgoCD.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enabled will toggle Grafana support globally for ArgoCD.",
								MarkdownDescription: "Enabled will toggle Grafana support globally for ArgoCD.",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"host": {
								Description:         "Host is the hostname to use for Ingress/Route resources.",
								MarkdownDescription: "Host is the hostname to use for Ingress/Route resources.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "Image is the Grafana container image.",
								MarkdownDescription: "Image is the Grafana container image.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingress": {
								Description:         "Ingress defines the desired state for an Ingress for the Grafana component.",
								MarkdownDescription: "Ingress defines the desired state for an Ingress for the Grafana component.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is the map of annotations to apply to the Ingress.",
										MarkdownDescription: "Annotations is the map of annotations to apply to the Ingress.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Enabled will toggle the creation of the Ingress.",
										MarkdownDescription: "Enabled will toggle the creation of the Ingress.",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"ingress_class_name": {
										Description:         "IngressClassName for the Ingress resource.",
										MarkdownDescription: "IngressClassName for the Ingress resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path used for the Ingress resource.",
										MarkdownDescription: "Path used for the Ingress resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
										MarkdownDescription: "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"hosts": {
												Description:         "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
												MarkdownDescription: "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_name": {
												Description:         "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
												MarkdownDescription: "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",

												Type: types.StringType,

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

							"resources": {
								Description:         "Resources defines the Compute Resources required by the container for Grafana.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for Grafana.",

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

							"route": {
								Description:         "Route defines the desired state for an OpenShift Route for the Grafana component.",
								MarkdownDescription: "Route defines the desired state for an OpenShift Route for the Grafana component.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is the map of annotations to use for the Route resource.",
										MarkdownDescription: "Annotations is the map of annotations to use for the Route resource.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Enabled will toggle the creation of the OpenShift Route.",
										MarkdownDescription: "Enabled will toggle the creation of the OpenShift Route.",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"labels": {
										Description:         "Labels is the map of labels to use for the Route resource",
										MarkdownDescription: "Labels is the map of labels to use for the Route resource",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path the router watches for, to route traffic for to the service.",
										MarkdownDescription: "Path the router watches for, to route traffic for to the service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "TLS provides the ability to configure certificates and termination for the Route.",
										MarkdownDescription: "TLS provides the ability to configure certificates and termination for the Route.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca_certificate": {
												Description:         "caCertificate provides the cert authority certificate contents",
												MarkdownDescription: "caCertificate provides the cert authority certificate contents",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"certificate": {
												Description:         "certificate provides certificate contents",
												MarkdownDescription: "certificate provides certificate contents",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"destination_ca_certificate": {
												Description:         "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",
												MarkdownDescription: "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_edge_termination_policy": {
												Description:         "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (default) * Disable - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",
												MarkdownDescription: "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (default) * Disable - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "key provides key file contents",
												MarkdownDescription: "key provides key file contents",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination": {
												Description:         "termination indicates termination type.",
												MarkdownDescription: "termination indicates termination type.",

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

									"wildcard_policy": {
										Description:         "WildcardPolicy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",
										MarkdownDescription: "WildcardPolicy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",

										Type: types.StringType,

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
								Description:         "Size is the replica count for the Grafana Deployment.",
								MarkdownDescription: "Size is the replica count for the Grafana Deployment.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Version is the Grafana container image tag.",
								MarkdownDescription: "Version is the Grafana container image tag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ha": {
						Description:         "HA options for High Availability support for the Redis component.",
						MarkdownDescription: "HA options for High Availability support for the Redis component.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enabled will toggle HA support globally for Argo CD.",
								MarkdownDescription: "Enabled will toggle HA support globally for Argo CD.",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"redis_proxy_image": {
								Description:         "RedisProxyImage is the Redis HAProxy container image.",
								MarkdownDescription: "RedisProxyImage is the Redis HAProxy container image.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"redis_proxy_version": {
								Description:         "RedisProxyVersion is the Redis HAProxy container image tag.",
								MarkdownDescription: "RedisProxyVersion is the Redis HAProxy container image tag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources defines the Compute Resources required by the container for HA.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for HA.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"help_chat_text": {
						Description:         "HelpChatText is the text for getting chat help, defaults to 'Chat now!'",
						MarkdownDescription: "HelpChatText is the text for getting chat help, defaults to 'Chat now!'",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"help_chat_url": {
						Description:         "HelpChatURL is the URL for getting chat help, this will typically be your Slack channel for support.",
						MarkdownDescription: "HelpChatURL is the URL for getting chat help, this will typically be your Slack channel for support.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"image": {
						Description:         "Image is the ArgoCD container image for all ArgoCD components.",
						MarkdownDescription: "Image is the ArgoCD container image for all ArgoCD components.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"import": {
						Description:         "Import is the import/restore options for ArgoCD.",
						MarkdownDescription: "Import is the import/restore options for ArgoCD.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "Name of an ArgoCDExport from which to import data.",
								MarkdownDescription: "Name of an ArgoCDExport from which to import data.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace for the ArgoCDExport, defaults to the same namespace as the ArgoCD.",
								MarkdownDescription: "Namespace for the ArgoCDExport, defaults to the same namespace as the ArgoCD.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"initial_repositories": {
						Description:         "InitialRepositories to configure Argo CD with upon creation of the cluster.",
						MarkdownDescription: "InitialRepositories to configure Argo CD with upon creation of the cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"initial_ssh_known_hosts": {
						Description:         "InitialSSHKnownHosts defines the SSH known hosts data upon creation of the cluster for connecting Git repositories via SSH.",
						MarkdownDescription: "InitialSSHKnownHosts defines the SSH known hosts data upon creation of the cluster for connecting Git repositories via SSH.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"excludedefaulthosts": {
								Description:         "ExcludeDefaultHosts describes whether you would like to include the default list of SSH Known Hosts provided by ArgoCD.",
								MarkdownDescription: "ExcludeDefaultHosts describes whether you would like to include the default list of SSH Known Hosts provided by ArgoCD.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"keys": {
								Description:         "Keys describes a custom set of SSH Known Hosts that you would like to have included in your ArgoCD server.",
								MarkdownDescription: "Keys describes a custom set of SSH Known Hosts that you would like to have included in your ArgoCD server.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kustomize_build_options": {
						Description:         "KustomizeBuildOptions is used to specify build options/parameters to use with 'kustomize build'.",
						MarkdownDescription: "KustomizeBuildOptions is used to specify build options/parameters to use with 'kustomize build'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kustomize_versions": {
						Description:         "KustomizeVersions is a listing of configured versions of Kustomize to be made available within ArgoCD.",
						MarkdownDescription: "KustomizeVersions is a listing of configured versions of Kustomize to be made available within ArgoCD.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"path": {
								Description:         "Path is the path to a configured kustomize version on the filesystem of your repo server.",
								MarkdownDescription: "Path is the path to a configured kustomize version on the filesystem of your repo server.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Version is a configured kustomize version in the format of vX.Y.Z",
								MarkdownDescription: "Version is a configured kustomize version in the format of vX.Y.Z",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"node_placement": {
						Description:         "NodePlacement defines NodeSelectors and Taints for Argo CD workloads",
						MarkdownDescription: "NodePlacement defines NodeSelectors and Taints for Argo CD workloads",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"node_selector": {
								Description:         "NodeSelector is a field of PodSpec, it is a map of key value pairs used for node selection",
								MarkdownDescription: "NodeSelector is a field of PodSpec, it is a map of key value pairs used for node selection",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tolerations": {
								Description:         "Tolerations allow the pods to schedule onto nodes with matching taints",
								MarkdownDescription: "Tolerations allow the pods to schedule onto nodes with matching taints",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"notifications": {
						Description:         "Notifications defines whether the Argo CD Notifications controller should be installed.",
						MarkdownDescription: "Notifications defines whether the Argo CD Notifications controller should be installed.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enabled defines whether argocd-notifications controller should be deployed or not",
								MarkdownDescription: "Enabled defines whether argocd-notifications controller should be deployed or not",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"env": {
								Description:         "Env let you specify environment variables for Notifications pods",
								MarkdownDescription: "Env let you specify environment variables for Notifications pods",

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

							"image": {
								Description:         "Image is the Argo CD Notifications image (optional)",
								MarkdownDescription: "Image is the Argo CD Notifications image (optional)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_level": {
								Description:         "LogLevel describes the log level that should be used by the argocd-notifications. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug,info, error, and warn.",
								MarkdownDescription: "LogLevel describes the log level that should be used by the argocd-notifications. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug,info, error, and warn.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": {
								Description:         "Replicas defines the number of replicas to run for notifications-controller",
								MarkdownDescription: "Replicas defines the number of replicas to run for notifications-controller",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources defines the Compute Resources required by the container for Argo CD Notifications.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for Argo CD Notifications.",

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

							"version": {
								Description:         "Version is the Argo CD Notifications image tag. (optional)",
								MarkdownDescription: "Version is the Argo CD Notifications image tag. (optional)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"oidc_config": {
						Description:         "OIDCConfig is the OIDC configuration as an alternative to dex.",
						MarkdownDescription: "OIDCConfig is the OIDC configuration as an alternative to dex.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"prometheus": {
						Description:         "Prometheus defines the Prometheus server options for ArgoCD.",
						MarkdownDescription: "Prometheus defines the Prometheus server options for ArgoCD.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enabled will toggle Prometheus support globally for ArgoCD.",
								MarkdownDescription: "Enabled will toggle Prometheus support globally for ArgoCD.",

								Type: types.BoolType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"host": {
								Description:         "Host is the hostname to use for Ingress/Route resources.",
								MarkdownDescription: "Host is the hostname to use for Ingress/Route resources.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingress": {
								Description:         "Ingress defines the desired state for an Ingress for the Prometheus component.",
								MarkdownDescription: "Ingress defines the desired state for an Ingress for the Prometheus component.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is the map of annotations to apply to the Ingress.",
										MarkdownDescription: "Annotations is the map of annotations to apply to the Ingress.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Enabled will toggle the creation of the Ingress.",
										MarkdownDescription: "Enabled will toggle the creation of the Ingress.",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"ingress_class_name": {
										Description:         "IngressClassName for the Ingress resource.",
										MarkdownDescription: "IngressClassName for the Ingress resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path used for the Ingress resource.",
										MarkdownDescription: "Path used for the Ingress resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
										MarkdownDescription: "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"hosts": {
												Description:         "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
												MarkdownDescription: "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_name": {
												Description:         "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
												MarkdownDescription: "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",

												Type: types.StringType,

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

							"route": {
								Description:         "Route defines the desired state for an OpenShift Route for the Prometheus component.",
								MarkdownDescription: "Route defines the desired state for an OpenShift Route for the Prometheus component.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is the map of annotations to use for the Route resource.",
										MarkdownDescription: "Annotations is the map of annotations to use for the Route resource.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Enabled will toggle the creation of the OpenShift Route.",
										MarkdownDescription: "Enabled will toggle the creation of the OpenShift Route.",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"labels": {
										Description:         "Labels is the map of labels to use for the Route resource",
										MarkdownDescription: "Labels is the map of labels to use for the Route resource",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path the router watches for, to route traffic for to the service.",
										MarkdownDescription: "Path the router watches for, to route traffic for to the service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "TLS provides the ability to configure certificates and termination for the Route.",
										MarkdownDescription: "TLS provides the ability to configure certificates and termination for the Route.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca_certificate": {
												Description:         "caCertificate provides the cert authority certificate contents",
												MarkdownDescription: "caCertificate provides the cert authority certificate contents",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"certificate": {
												Description:         "certificate provides certificate contents",
												MarkdownDescription: "certificate provides certificate contents",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"destination_ca_certificate": {
												Description:         "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",
												MarkdownDescription: "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_edge_termination_policy": {
												Description:         "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (default) * Disable - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",
												MarkdownDescription: "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (default) * Disable - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "key provides key file contents",
												MarkdownDescription: "key provides key file contents",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination": {
												Description:         "termination indicates termination type.",
												MarkdownDescription: "termination indicates termination type.",

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

									"wildcard_policy": {
										Description:         "WildcardPolicy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",
										MarkdownDescription: "WildcardPolicy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",

										Type: types.StringType,

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
								Description:         "Size is the replica count for the Prometheus StatefulSet.",
								MarkdownDescription: "Size is the replica count for the Prometheus StatefulSet.",

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

					"rbac": {
						Description:         "RBAC defines the RBAC configuration for Argo CD.",
						MarkdownDescription: "RBAC defines the RBAC configuration for Argo CD.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"default_policy": {
								Description:         "DefaultPolicy is the name of the default role which Argo CD will falls back to, when authorizing API requests (optional). If omitted or empty, users may be still be able to login, but will see no apps, projects, etc...",
								MarkdownDescription: "DefaultPolicy is the name of the default role which Argo CD will falls back to, when authorizing API requests (optional). If omitted or empty, users may be still be able to login, but will see no apps, projects, etc...",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"policy": {
								Description:         "Policy is CSV containing user-defined RBAC policies and role definitions. Policy rules are in the form:   p, subject, resource, action, object, effect Role definitions and bindings are in the form:   g, subject, inherited-subject See https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/rbac.md for additional information.",
								MarkdownDescription: "Policy is CSV containing user-defined RBAC policies and role definitions. Policy rules are in the form:   p, subject, resource, action, object, effect Role definitions and bindings are in the form:   g, subject, inherited-subject See https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/rbac.md for additional information.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"policy_matcher_mode": {
								Description:         "PolicyMatcherMode configures the matchers function mode for casbin. There are two options for this, 'glob' for glob matcher or 'regex' for regex matcher.",
								MarkdownDescription: "PolicyMatcherMode configures the matchers function mode for casbin. There are two options for this, 'glob' for glob matcher or 'regex' for regex matcher.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scopes": {
								Description:         "Scopes controls which OIDC scopes to examine during rbac enforcement (in addition to 'sub' scope). If omitted, defaults to: '[groups]'.",
								MarkdownDescription: "Scopes controls which OIDC scopes to examine during rbac enforcement (in addition to 'sub' scope). If omitted, defaults to: '[groups]'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis": {
						Description:         "Redis defines the Redis server options for ArgoCD.",
						MarkdownDescription: "Redis defines the Redis server options for ArgoCD.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"autotls": {
								Description:         "AutoTLS specifies the method to use for automatic TLS configuration for the redis server The value specified here can currently be: - openshift - Use the OpenShift service CA to request TLS config",
								MarkdownDescription: "AutoTLS specifies the method to use for automatic TLS configuration for the redis server The value specified here can currently be: - openshift - Use the OpenShift service CA to request TLS config",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"disable_tls_verification": {
								Description:         "DisableTLSVerification defines whether redis server API should be accessed using strict TLS validation",
								MarkdownDescription: "DisableTLSVerification defines whether redis server API should be accessed using strict TLS validation",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "Image is the Redis container image.",
								MarkdownDescription: "Image is the Redis container image.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources defines the Compute Resources required by the container for Redis.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for Redis.",

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

							"version": {
								Description:         "Version is the Redis container image tag.",
								MarkdownDescription: "Version is the Redis container image tag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"repo": {
						Description:         "Repo defines the repo server options for Argo CD.",
						MarkdownDescription: "Repo defines the repo server options for Argo CD.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"autotls": {
								Description:         "AutoTLS specifies the method to use for automatic TLS configuration for the repo server The value specified here can currently be: - openshift - Use the OpenShift service CA to request TLS config",
								MarkdownDescription: "AutoTLS specifies the method to use for automatic TLS configuration for the repo server The value specified here can currently be: - openshift - Use the OpenShift service CA to request TLS config",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"env": {
								Description:         "Env lets you specify environment for repo server pods",
								MarkdownDescription: "Env lets you specify environment for repo server pods",

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

							"exec_timeout": {
								Description:         "ExecTimeout specifies the timeout in seconds for tool execution",
								MarkdownDescription: "ExecTimeout specifies the timeout in seconds for tool execution",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image": {
								Description:         "Image is the ArgoCD Repo Server container image.",
								MarkdownDescription: "Image is the ArgoCD Repo Server container image.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"init_containers": {
								Description:         "InitContainers defines the list of initialization containers for the repo server deployment",
								MarkdownDescription: "InitContainers defines the list of initialization containers for the repo server deployment",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"args": {
										Description:         "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
										MarkdownDescription: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"command": {
										Description:         "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
										MarkdownDescription: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

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
										Description:         "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
										MarkdownDescription: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",

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
												Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
												MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

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
										Description:         "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",
										MarkdownDescription: "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",

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
												Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
												MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

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
												Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
												MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

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

							"log_format": {
								Description:         "LogFormat describes the log format that should be used by the Repo Server. Defaults to ArgoCDDefaultLogFormat if not configured. Valid options are text or json.",
								MarkdownDescription: "LogFormat describes the log format that should be used by the Repo Server. Defaults to ArgoCDDefaultLogFormat if not configured. Valid options are text or json.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_level": {
								Description:         "LogLevel describes the log level that should be used by the Repo Server. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug, info, error, and warn.",
								MarkdownDescription: "LogLevel describes the log level that should be used by the Repo Server. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug, info, error, and warn.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mountsatoken": {
								Description:         "MountSAToken describes whether you would like to have the Repo server mount the service account token",
								MarkdownDescription: "MountSAToken describes whether you would like to have the Repo server mount the service account token",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": {
								Description:         "Replicas defines the number of replicas for argocd-repo-server. Value should be greater than or equal to 0. Default is nil.",
								MarkdownDescription: "Replicas defines the number of replicas for argocd-repo-server. Value should be greater than or equal to 0. Default is nil.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources defines the Compute Resources required by the container for Redis.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for Redis.",

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

							"serviceaccount": {
								Description:         "ServiceAccount defines the ServiceAccount user that you would like the Repo server to use",
								MarkdownDescription: "ServiceAccount defines the ServiceAccount user that you would like the Repo server to use",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"sidecar_containers": {
								Description:         "SidecarContainers defines the list of sidecar containers for the repo server deployment",
								MarkdownDescription: "SidecarContainers defines the list of sidecar containers for the repo server deployment",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"args": {
										Description:         "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
										MarkdownDescription: "Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"command": {
										Description:         "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",
										MarkdownDescription: "Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell",

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
										Description:         "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",
										MarkdownDescription: "Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images This field is optional to allow higher level config management to default or override container images in workload controllers like Deployments and StatefulSets.",

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
												Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
												MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

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
										Description:         "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",
										MarkdownDescription: "List of ports to expose from the container. Exposing a port here gives the system additional information about the network connections a container uses, but is primarily informational. Not specifying a port here DOES NOT prevent that port from being exposed. Any port which is listening on the default '0.0.0.0' address inside a container will be accessible from the network. Cannot be updated.",

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
												Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
												MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

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
												Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
												MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

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

							"verifytls": {
								Description:         "VerifyTLS defines whether repo server API should be accessed using strict TLS validation",
								MarkdownDescription: "VerifyTLS defines whether repo server API should be accessed using strict TLS validation",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Version is the ArgoCD Repo Server container image tag.",
								MarkdownDescription: "Version is the ArgoCD Repo Server container image tag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_mounts": {
								Description:         "VolumeMounts adds volumeMounts to the repo server container",
								MarkdownDescription: "VolumeMounts adds volumeMounts to the repo server container",

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
								Description:         "Volumes adds volumes to the repo server deployment",
								MarkdownDescription: "Volumes adds volumes to the repo server deployment",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"aws_elastic_block_store": {
										Description:         "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
										MarkdownDescription: "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
												MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"partition": {
												Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
												MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
												MarkdownDescription: "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_id": {
												Description:         "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
												MarkdownDescription: "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

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
										Description:         "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
										MarkdownDescription: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"caching_mode": {
												Description:         "Host Caching mode: None, Read Only, Read Write.",
												MarkdownDescription: "Host Caching mode: None, Read Only, Read Write.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"disk_name": {
												Description:         "The Name of the data disk in the blob storage",
												MarkdownDescription: "The Name of the data disk in the blob storage",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"disk_uri": {
												Description:         "The URI the data disk in the blob storage",
												MarkdownDescription: "The URI the data disk in the blob storage",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"fs_type": {
												Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
												MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"kind": {
												Description:         "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
												MarkdownDescription: "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

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
										Description:         "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
										MarkdownDescription: "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"read_only": {
												Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_name": {
												Description:         "the name of secret that contains Azure Storage Account Name and Key",
												MarkdownDescription: "the name of secret that contains Azure Storage Account Name and Key",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"share_name": {
												Description:         "Share Name",
												MarkdownDescription: "Share Name",

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
										Description:         "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
										MarkdownDescription: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"monitors": {
												Description:         "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
												MarkdownDescription: "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"path": {
												Description:         "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
												MarkdownDescription: "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
												MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_file": {
												Description:         "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
												MarkdownDescription: "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
												MarkdownDescription: "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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
												Description:         "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
												MarkdownDescription: "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

												Type: types.StringType,

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
										Description:         "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
										MarkdownDescription: "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "Optional: points to a secret object containing parameters used to connect to OpenStack.",
												MarkdownDescription: "Optional: points to a secret object containing parameters used to connect to OpenStack.",

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
												Description:         "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												MarkdownDescription: "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

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
										Description:         "ConfigMap represents a configMap that should populate this volume",
										MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default_mode": {
												Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
												MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"items": {
												Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
												MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

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
														Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
														MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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

									"csi": {
										Description:         "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
										MarkdownDescription: "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"driver": {
												Description:         "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
												MarkdownDescription: "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"fs_type": {
												Description:         "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
												MarkdownDescription: "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_publish_secret_ref": {
												Description:         "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
												MarkdownDescription: "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",

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
												Description:         "Specifies a read-only configuration for the volume. Defaults to false (read/write).",
												MarkdownDescription: "Specifies a read-only configuration for the volume. Defaults to false (read/write).",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_attributes": {
												Description:         "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
												MarkdownDescription: "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",

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
										Description:         "DownwardAPI represents downward API about the pod that should populate this volume",
										MarkdownDescription: "DownwardAPI represents downward API about the pod that should populate this volume",

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
												}),

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
										Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
										MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"medium": {
												Description:         "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												MarkdownDescription: "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"size_limit": {
												Description:         "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
												MarkdownDescription: "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",

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

									"ephemeral": {
										Description:         "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
										MarkdownDescription: "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",

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
																Description:         "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																MarkdownDescription: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"data_source": {
																Description:         "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																MarkdownDescription: "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",

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
																Description:         "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																MarkdownDescription: "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",

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
																Description:         "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																MarkdownDescription: "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",

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
																Description:         "A label query over volumes to consider for binding.",
																MarkdownDescription: "A label query over volumes to consider for binding.",

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
																Description:         "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																MarkdownDescription: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",

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
																Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
																MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",

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
										Description:         "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
										MarkdownDescription: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
												MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"lun": {
												Description:         "Optional: FC target lun number",
												MarkdownDescription: "Optional: FC target lun number",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"target_ww_ns": {
												Description:         "Optional: FC target worldwide names (WWNs)",
												MarkdownDescription: "Optional: FC target worldwide names (WWNs)",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"wwids": {
												Description:         "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
												MarkdownDescription: "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",

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
										Description:         "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
										MarkdownDescription: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"driver": {
												Description:         "Driver is the name of the driver to use for this volume.",
												MarkdownDescription: "Driver is the name of the driver to use for this volume.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"fs_type": {
												Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
												MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"options": {
												Description:         "Optional: Extra command options if any.",
												MarkdownDescription: "Optional: Extra command options if any.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
												MarkdownDescription: "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",

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
										Description:         "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
										MarkdownDescription: "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"dataset_name": {
												Description:         "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
												MarkdownDescription: "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"dataset_uuid": {
												Description:         "UUID of the dataset. This is unique identifier of a Flocker dataset",
												MarkdownDescription: "UUID of the dataset. This is unique identifier of a Flocker dataset",

												Type: types.StringType,

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
										Description:         "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
										MarkdownDescription: "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
												MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"partition": {
												Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
												MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pd_name": {
												Description:         "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
												MarkdownDescription: "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"read_only": {
												Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
												MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

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
										Description:         "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
										MarkdownDescription: "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"directory": {
												Description:         "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
												MarkdownDescription: "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"repository": {
												Description:         "Repository URL",
												MarkdownDescription: "Repository URL",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"revision": {
												Description:         "Commit hash for the specified revision.",
												MarkdownDescription: "Commit hash for the specified revision.",

												Type: types.StringType,

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
										Description:         "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
										MarkdownDescription: "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"endpoints": {
												Description:         "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
												MarkdownDescription: "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"path": {
												Description:         "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
												MarkdownDescription: "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"read_only": {
												Description:         "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
												MarkdownDescription: "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

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
										Description:         "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
										MarkdownDescription: "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"path": {
												Description:         "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
												MarkdownDescription: "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
												MarkdownDescription: "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

												Type: types.StringType,

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
										Description:         "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
										MarkdownDescription: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"chap_auth_discovery": {
												Description:         "whether support iSCSI Discovery CHAP authentication",
												MarkdownDescription: "whether support iSCSI Discovery CHAP authentication",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"chap_auth_session": {
												Description:         "whether support iSCSI Session CHAP authentication",
												MarkdownDescription: "whether support iSCSI Session CHAP authentication",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"fs_type": {
												Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
												MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"initiator_name": {
												Description:         "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
												MarkdownDescription: "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"iqn": {
												Description:         "Target iSCSI Qualified Name.",
												MarkdownDescription: "Target iSCSI Qualified Name.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"iscsi_interface": {
												Description:         "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
												MarkdownDescription: "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"lun": {
												Description:         "iSCSI Target Lun number.",
												MarkdownDescription: "iSCSI Target Lun number.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"portals": {
												Description:         "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
												MarkdownDescription: "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
												MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "CHAP Secret for iSCSI target and initiator authentication",
												MarkdownDescription: "CHAP Secret for iSCSI target and initiator authentication",

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
												Description:         "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
												MarkdownDescription: "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

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
										Description:         "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"nfs": {
										Description:         "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
										MarkdownDescription: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"path": {
												Description:         "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
												MarkdownDescription: "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"read_only": {
												Description:         "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
												MarkdownDescription: "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"server": {
												Description:         "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
												MarkdownDescription: "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

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
										Description:         "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
										MarkdownDescription: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"claim_name": {
												Description:         "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
												MarkdownDescription: "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"read_only": {
												Description:         "Will force the ReadOnly setting in VolumeMounts. Default false.",
												MarkdownDescription: "Will force the ReadOnly setting in VolumeMounts. Default false.",

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
										Description:         "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
										MarkdownDescription: "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
												MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"pd_id": {
												Description:         "ID that identifies Photon Controller persistent disk",
												MarkdownDescription: "ID that identifies Photon Controller persistent disk",

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
										Description:         "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
										MarkdownDescription: "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
												MarkdownDescription: "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_id": {
												Description:         "VolumeID uniquely identifies a Portworx volume",
												MarkdownDescription: "VolumeID uniquely identifies a Portworx volume",

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
										Description:         "Items for all in one resources secrets, configmaps, and downward API",
										MarkdownDescription: "Items for all in one resources secrets, configmaps, and downward API",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default_mode": {
												Description:         "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
												MarkdownDescription: "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"sources": {
												Description:         "list of volume projections",
												MarkdownDescription: "list of volume projections",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"config_map": {
														Description:         "information about the configMap data to project",
														MarkdownDescription: "information about the configMap data to project",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"items": {
																Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

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
																		Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																		MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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

													"downward_api": {
														Description:         "information about the downwardAPI data to project",
														MarkdownDescription: "information about the downwardAPI data to project",

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
																}),

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
														Description:         "information about the secret data to project",
														MarkdownDescription: "information about the secret data to project",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"items": {
																Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

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
																		Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																		MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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

													"service_account_token": {
														Description:         "information about the serviceAccountToken data to project",
														MarkdownDescription: "information about the serviceAccountToken data to project",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"audience": {
																Description:         "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																MarkdownDescription: "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"expiration_seconds": {
																Description:         "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																MarkdownDescription: "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "Path is the path relative to the mount point of the file to project the token into.",
																MarkdownDescription: "Path is the path relative to the mount point of the file to project the token into.",

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
										Description:         "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
										MarkdownDescription: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"group": {
												Description:         "Group to map volume access to Default is no group",
												MarkdownDescription: "Group to map volume access to Default is no group",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
												MarkdownDescription: "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"registry": {
												Description:         "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
												MarkdownDescription: "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"tenant": {
												Description:         "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
												MarkdownDescription: "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"user": {
												Description:         "User to map volume access to Defaults to serivceaccount user",
												MarkdownDescription: "User to map volume access to Defaults to serivceaccount user",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume": {
												Description:         "Volume is a string that references an already created Quobyte volume by name.",
												MarkdownDescription: "Volume is a string that references an already created Quobyte volume by name.",

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
										Description:         "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
										MarkdownDescription: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
												MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"image": {
												Description:         "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
												MarkdownDescription: "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"keyring": {
												Description:         "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
												MarkdownDescription: "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"monitors": {
												Description:         "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
												MarkdownDescription: "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

												Type: types.ListType{ElemType: types.StringType},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"pool": {
												Description:         "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
												MarkdownDescription: "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
												MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
												MarkdownDescription: "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

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
												Description:         "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
												MarkdownDescription: "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

												Type: types.StringType,

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
										Description:         "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
										MarkdownDescription: "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
												MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"gateway": {
												Description:         "The host address of the ScaleIO API Gateway.",
												MarkdownDescription: "The host address of the ScaleIO API Gateway.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"protection_domain": {
												Description:         "The name of the ScaleIO Protection Domain for the configured storage.",
												MarkdownDescription: "The name of the ScaleIO Protection Domain for the configured storage.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
												MarkdownDescription: "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",

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
												Description:         "Flag to enable/disable SSL communication with Gateway, default false",
												MarkdownDescription: "Flag to enable/disable SSL communication with Gateway, default false",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_mode": {
												Description:         "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
												MarkdownDescription: "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_pool": {
												Description:         "The ScaleIO Storage Pool associated with the protection domain.",
												MarkdownDescription: "The ScaleIO Storage Pool associated with the protection domain.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"system": {
												Description:         "The name of the storage system as configured in ScaleIO.",
												MarkdownDescription: "The name of the storage system as configured in ScaleIO.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"volume_name": {
												Description:         "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
												MarkdownDescription: "The name of a volume already created in the ScaleIO system that is associated with this volume source.",

												Type: types.StringType,

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
										Description:         "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
										MarkdownDescription: "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"default_mode": {
												Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
												MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"items": {
												Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
												MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

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
														Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"path": {
														Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
														MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
												Description:         "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
												MarkdownDescription: "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

												Type: types.StringType,

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
										Description:         "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
										MarkdownDescription: "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
												MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"read_only": {
												Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_ref": {
												Description:         "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
												MarkdownDescription: "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",

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
												Description:         "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
												MarkdownDescription: "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_namespace": {
												Description:         "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
												MarkdownDescription: "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",

												Type: types.StringType,

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
										Description:         "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
										MarkdownDescription: "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"fs_type": {
												Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
												MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_policy_id": {
												Description:         "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
												MarkdownDescription: "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_policy_name": {
												Description:         "Storage Policy Based Management (SPBM) profile name.",
												MarkdownDescription: "Storage Policy Based Management (SPBM) profile name.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"volume_path": {
												Description:         "Path that identifies vSphere volume vmdk",
												MarkdownDescription: "Path that identifies vSphere volume vmdk",

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

					"repository_credentials": {
						Description:         "RepositoryCredentials are the Git pull credentials to configure Argo CD with upon creation of the cluster.",
						MarkdownDescription: "RepositoryCredentials are the Git pull credentials to configure Argo CD with upon creation of the cluster.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_customizations": {
						Description:         "ResourceCustomizations customizes resource behavior. Keys are in the form: group/Kind.",
						MarkdownDescription: "ResourceCustomizations customizes resource behavior. Keys are in the form: group/Kind.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_exclusions": {
						Description:         "ResourceExclusions is used to completely ignore entire classes of resource group/kinds.",
						MarkdownDescription: "ResourceExclusions is used to completely ignore entire classes of resource group/kinds.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_inclusions": {
						Description:         "ResourceInclusions is used to only include specific group/kinds in the reconciliation process.",
						MarkdownDescription: "ResourceInclusions is used to only include specific group/kinds in the reconciliation process.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_tracking_method": {
						Description:         "ResourceTrackingMethod defines how Argo CD should track resources that it manages",
						MarkdownDescription: "ResourceTrackingMethod defines how Argo CD should track resources that it manages",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"server": {
						Description:         "Server defines the options for the ArgoCD Server component.",
						MarkdownDescription: "Server defines the options for the ArgoCD Server component.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"autoscale": {
								Description:         "Autoscale defines the autoscale options for the Argo CD Server component.",
								MarkdownDescription: "Autoscale defines the autoscale options for the Argo CD Server component.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Enabled will toggle autoscaling support for the Argo CD Server component.",
										MarkdownDescription: "Enabled will toggle autoscaling support for the Argo CD Server component.",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"hpa": {
										Description:         "HPA defines the HorizontalPodAutoscaler options for the Argo CD Server component.",
										MarkdownDescription: "HPA defines the HorizontalPodAutoscaler options for the Argo CD Server component.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"max_replicas": {
												Description:         "upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",
												MarkdownDescription: "upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"min_replicas": {
												Description:         "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",
												MarkdownDescription: "minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scale_target_ref": {
												Description:         "reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.",
												MarkdownDescription: "reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"api_version": {
														Description:         "API version of the referent",
														MarkdownDescription: "API version of the referent",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",
														MarkdownDescription: "Kind of the referent; More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds'",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"name": {
														Description:         "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",
														MarkdownDescription: "Name of the referent; More info: http://kubernetes.io/docs/user-guide/identifiers#names",

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

											"target_cpu_utilization_percentage": {
												Description:         "target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.",
												MarkdownDescription: "target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.",

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

							"env": {
								Description:         "Env lets you specify environment for API server pods",
								MarkdownDescription: "Env lets you specify environment for API server pods",

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

							"extra_command_args": {
								Description:         "Extra Command arguments that would append to the Argo CD server command. ExtraCommandArgs will not be added, if one of these commands is already part of the server command with same or different value.",
								MarkdownDescription: "Extra Command arguments that would append to the Argo CD server command. ExtraCommandArgs will not be added, if one of these commands is already part of the server command with same or different value.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"grpc": {
								Description:         "GRPC defines the state for the Argo CD Server GRPC options.",
								MarkdownDescription: "GRPC defines the state for the Argo CD Server GRPC options.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"host": {
										Description:         "Host is the hostname to use for Ingress/Route resources.",
										MarkdownDescription: "Host is the hostname to use for Ingress/Route resources.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"ingress": {
										Description:         "Ingress defines the desired state for the Argo CD Server GRPC Ingress.",
										MarkdownDescription: "Ingress defines the desired state for the Argo CD Server GRPC Ingress.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "Annotations is the map of annotations to apply to the Ingress.",
												MarkdownDescription: "Annotations is the map of annotations to apply to the Ingress.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enabled": {
												Description:         "Enabled will toggle the creation of the Ingress.",
												MarkdownDescription: "Enabled will toggle the creation of the Ingress.",

												Type: types.BoolType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"ingress_class_name": {
												Description:         "IngressClassName for the Ingress resource.",
												MarkdownDescription: "IngressClassName for the Ingress resource.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"path": {
												Description:         "Path used for the Ingress resource.",
												MarkdownDescription: "Path used for the Ingress resource.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls": {
												Description:         "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
												MarkdownDescription: "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"hosts": {
														Description:         "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
														MarkdownDescription: "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_name": {
														Description:         "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
														MarkdownDescription: "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",

														Type: types.StringType,

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

							"host": {
								Description:         "Host is the hostname to use for Ingress/Route resources.",
								MarkdownDescription: "Host is the hostname to use for Ingress/Route resources.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingress": {
								Description:         "Ingress defines the desired state for an Ingress for the Argo CD Server component.",
								MarkdownDescription: "Ingress defines the desired state for an Ingress for the Argo CD Server component.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is the map of annotations to apply to the Ingress.",
										MarkdownDescription: "Annotations is the map of annotations to apply to the Ingress.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Enabled will toggle the creation of the Ingress.",
										MarkdownDescription: "Enabled will toggle the creation of the Ingress.",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"ingress_class_name": {
										Description:         "IngressClassName for the Ingress resource.",
										MarkdownDescription: "IngressClassName for the Ingress resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path used for the Ingress resource.",
										MarkdownDescription: "Path used for the Ingress resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",
										MarkdownDescription: "TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"hosts": {
												Description:         "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",
												MarkdownDescription: "Hosts are a list of hosts included in the TLS certificate. The values in this list must match the name/s used in the tlsSecret. Defaults to the wildcard host setting for the loadbalancer controller fulfilling this Ingress, if left unspecified.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_name": {
												Description:         "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",
												MarkdownDescription: "SecretName is the name of the secret used to terminate TLS traffic on port 443. Field is left optional to allow TLS routing based on SNI hostname alone. If the SNI host in a listener conflicts with the 'Host' header field used by an IngressRule, the SNI host is used for termination and value of the Host header is used for routing.",

												Type: types.StringType,

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

							"insecure": {
								Description:         "Insecure toggles the insecure flag.",
								MarkdownDescription: "Insecure toggles the insecure flag.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_format": {
								Description:         "LogFormat refers to the log level to be used by the ArgoCD Server component. Defaults to ArgoCDDefaultLogFormat if not configured. Valid options are text or json.",
								MarkdownDescription: "LogFormat refers to the log level to be used by the ArgoCD Server component. Defaults to ArgoCDDefaultLogFormat if not configured. Valid options are text or json.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_level": {
								Description:         "LogLevel refers to the log level to be used by the ArgoCD Server component. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug, info, error, and warn.",
								MarkdownDescription: "LogLevel refers to the log level to be used by the ArgoCD Server component. Defaults to ArgoCDDefaultLogLevel if not set.  Valid options are debug, info, error, and warn.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": {
								Description:         "Replicas defines the number of replicas for argocd-server. Default is nil. Value should be greater than or equal to 0. Value will be ignored if Autoscaler is enabled.",
								MarkdownDescription: "Replicas defines the number of replicas for argocd-server. Default is nil. Value should be greater than or equal to 0. Value will be ignored if Autoscaler is enabled.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources defines the Compute Resources required by the container for the Argo CD server component.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for the Argo CD server component.",

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

							"route": {
								Description:         "Route defines the desired state for an OpenShift Route for the Argo CD Server component.",
								MarkdownDescription: "Route defines the desired state for an OpenShift Route for the Argo CD Server component.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"annotations": {
										Description:         "Annotations is the map of annotations to use for the Route resource.",
										MarkdownDescription: "Annotations is the map of annotations to use for the Route resource.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Enabled will toggle the creation of the OpenShift Route.",
										MarkdownDescription: "Enabled will toggle the creation of the OpenShift Route.",

										Type: types.BoolType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"labels": {
										Description:         "Labels is the map of labels to use for the Route resource",
										MarkdownDescription: "Labels is the map of labels to use for the Route resource",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"path": {
										Description:         "Path the router watches for, to route traffic for to the service.",
										MarkdownDescription: "Path the router watches for, to route traffic for to the service.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tls": {
										Description:         "TLS provides the ability to configure certificates and termination for the Route.",
										MarkdownDescription: "TLS provides the ability to configure certificates and termination for the Route.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca_certificate": {
												Description:         "caCertificate provides the cert authority certificate contents",
												MarkdownDescription: "caCertificate provides the cert authority certificate contents",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"certificate": {
												Description:         "certificate provides certificate contents",
												MarkdownDescription: "certificate provides certificate contents",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"destination_ca_certificate": {
												Description:         "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",
												MarkdownDescription: "destinationCACertificate provides the contents of the ca certificate of the final destination.  When using reencrypt termination this file should be provided in order to have routers use it for health checks on the secure connection. If this field is not specified, the router may provide its own destination CA and perform hostname validation using the short service name (service.namespace.svc), which allows infrastructure generated certificates to automatically verify.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_edge_termination_policy": {
												Description:         "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (default) * Disable - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",
												MarkdownDescription: "insecureEdgeTerminationPolicy indicates the desired behavior for insecure connections to a route. While each router may make its own decisions on which ports to expose, this is normally port 80.  * Allow - traffic is sent to the server on the insecure port (default) * Disable - no traffic is allowed on the insecure port. * Redirect - clients are redirected to the secure port.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "key provides key file contents",
												MarkdownDescription: "key provides key file contents",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"termination": {
												Description:         "termination indicates termination type.",
												MarkdownDescription: "termination indicates termination type.",

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

									"wildcard_policy": {
										Description:         "WildcardPolicy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",
										MarkdownDescription: "WildcardPolicy if any for the route. Currently only 'Subdomain' or 'None' is allowed.",

										Type: types.StringType,

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
								Description:         "Service defines the options for the Service backing the ArgoCD Server component.",
								MarkdownDescription: "Service defines the options for the Service backing the ArgoCD Server component.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"type": {
										Description:         "Type is the ServiceType to use for the Service resource.",
										MarkdownDescription: "Type is the ServiceType to use for the Service resource.",

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

					"sso": {
						Description:         "SSO defines the Single Sign-on configuration for Argo CD",
						MarkdownDescription: "SSO defines the Single Sign-on configuration for Argo CD",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dex": {
								Description:         "Dex contains the configuration for Argo CD dex authentication",
								MarkdownDescription: "Dex contains the configuration for Argo CD dex authentication",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config": {
										Description:         "Config is the dex connector configuration.",
										MarkdownDescription: "Config is the dex connector configuration.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"groups": {
										Description:         "Optional list of required groups a user must be a member of",
										MarkdownDescription: "Optional list of required groups a user must be a member of",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"image": {
										Description:         "Image is the Dex container image.",
										MarkdownDescription: "Image is the Dex container image.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"open_shift_o_auth": {
										Description:         "OpenShiftOAuth enables OpenShift OAuth authentication for the Dex server.",
										MarkdownDescription: "OpenShiftOAuth enables OpenShift OAuth authentication for the Dex server.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "Resources defines the Compute Resources required by the container for Dex.",
										MarkdownDescription: "Resources defines the Compute Resources required by the container for Dex.",

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

									"version": {
										Description:         "Version is the Dex container image tag.",
										MarkdownDescription: "Version is the Dex container image tag.",

										Type: types.StringType,

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
								Description:         "Image is the SSO container image.",
								MarkdownDescription: "Image is the SSO container image.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"keycloak": {
								Description:         "Keycloak contains the configuration for Argo CD keycloak authentication",
								MarkdownDescription: "Keycloak contains the configuration for Argo CD keycloak authentication",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image": {
										Description:         "Image is the Keycloak container image.",
										MarkdownDescription: "Image is the Keycloak container image.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": {
										Description:         "Resources defines the Compute Resources required by the container for Keycloak.",
										MarkdownDescription: "Resources defines the Compute Resources required by the container for Keycloak.",

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

									"root_ca": {
										Description:         "Custom root CA certificate for communicating with the Keycloak OIDC provider",
										MarkdownDescription: "Custom root CA certificate for communicating with the Keycloak OIDC provider",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"verify_tls": {
										Description:         "VerifyTLS set to false disables strict TLS validation.",
										MarkdownDescription: "VerifyTLS set to false disables strict TLS validation.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"version": {
										Description:         "Version is the Keycloak container image tag.",
										MarkdownDescription: "Version is the Keycloak container image tag.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"provider": {
								Description:         "Provider installs and configures the given SSO Provider with Argo CD.",
								MarkdownDescription: "Provider installs and configures the given SSO Provider with Argo CD.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources defines the Compute Resources required by the container for SSO.",
								MarkdownDescription: "Resources defines the Compute Resources required by the container for SSO.",

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

							"verify_tls": {
								Description:         "VerifyTLS set to false disables strict TLS validation.",
								MarkdownDescription: "VerifyTLS set to false disables strict TLS validation.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version": {
								Description:         "Version is the SSO container image tag.",
								MarkdownDescription: "Version is the SSO container image tag.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"status_badge_enabled": {
						Description:         "StatusBadgeEnabled toggles application status badge feature.",
						MarkdownDescription: "StatusBadgeEnabled toggles application status badge feature.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"tls": {
						Description:         "TLS defines the TLS options for ArgoCD.",
						MarkdownDescription: "TLS defines the TLS options for ArgoCD.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"ca": {
								Description:         "CA defines the CA options.",
								MarkdownDescription: "CA defines the CA options.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_map_name": {
										Description:         "ConfigMapName is the name of the ConfigMap containing the CA Certificate.",
										MarkdownDescription: "ConfigMapName is the name of the ConfigMap containing the CA Certificate.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_name": {
										Description:         "SecretName is the name of the Secret containing the CA Certificate and Key.",
										MarkdownDescription: "SecretName is the name of the Secret containing the CA Certificate and Key.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"initial_certs": {
								Description:         "InitialCerts defines custom TLS certificates upon creation of the cluster for connecting Git repositories via HTTPS.",
								MarkdownDescription: "InitialCerts defines custom TLS certificates upon creation of the cluster for connecting Git repositories via HTTPS.",

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

					"users_anonymous_enabled": {
						Description:         "UsersAnonymousEnabled toggles anonymous user access. The anonymous users get default role permissions specified argocd-rbac-cm.",
						MarkdownDescription: "UsersAnonymousEnabled toggles anonymous user access. The anonymous users get default role permissions specified argocd-rbac-cm.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": {
						Description:         "Version is the tag to use with the ArgoCD container image for all ArgoCD components.",
						MarkdownDescription: "Version is the tag to use with the ArgoCD container image for all ArgoCD components.",

						Type: types.StringType,

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

func (r *ArgoprojIoArgoCDV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_argoproj_io_argo_cd_v1alpha1")

	var state ArgoprojIoArgoCDV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ArgoprojIoArgoCDV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("argoproj.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ArgoCD")

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

func (r *ArgoprojIoArgoCDV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_argoproj_io_argo_cd_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *ArgoprojIoArgoCDV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_argoproj_io_argo_cd_v1alpha1")

	var state ArgoprojIoArgoCDV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ArgoprojIoArgoCDV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("argoproj.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ArgoCD")

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

func (r *ArgoprojIoArgoCDV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_argoproj_io_argo_cd_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
