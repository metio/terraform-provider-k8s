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

type KialiIoKialiV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*KialiIoKialiV1Alpha1Resource)(nil)
)

type KialiIoKialiV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type KialiIoKialiV1Alpha1GoModel struct {
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
		Istio_namespace *string `tfsdk:"istio_namespace" yaml:"istio_namespace,omitempty"`

		Kiali_feature_flags *struct {
			Validations *struct {
				Ignore *[]string `tfsdk:"ignore" yaml:"ignore,omitempty"`
			} `tfsdk:"validations" yaml:"validations,omitempty"`

			Certificates_information_indicators *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Secrets *[]string `tfsdk:"secrets" yaml:"secrets,omitempty"`
			} `tfsdk:"certificates_information_indicators" yaml:"certificates_information_indicators,omitempty"`

			Clustering *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"clustering" yaml:"clustering,omitempty"`

			Disabled_features *[]string `tfsdk:"disabled_features" yaml:"disabled_features,omitempty"`

			Istio_injection_action *bool `tfsdk:"istio_injection_action" yaml:"istio_injection_action,omitempty"`

			Istio_upgrade_action *bool `tfsdk:"istio_upgrade_action" yaml:"istio_upgrade_action,omitempty"`

			Ui_defaults *struct {
				Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

				Refresh_interval *string `tfsdk:"refresh_interval" yaml:"refresh_interval,omitempty"`

				Graph *struct {
					Find_options *[]struct {
						Description *string `tfsdk:"description" yaml:"description,omitempty"`

						Expression *string `tfsdk:"expression" yaml:"expression,omitempty"`
					} `tfsdk:"find_options" yaml:"find_options,omitempty"`

					Hide_options *[]struct {
						Description *string `tfsdk:"description" yaml:"description,omitempty"`

						Expression *string `tfsdk:"expression" yaml:"expression,omitempty"`
					} `tfsdk:"hide_options" yaml:"hide_options,omitempty"`

					Traffic *struct {
						Grpc *string `tfsdk:"grpc" yaml:"grpc,omitempty"`

						Http *string `tfsdk:"http" yaml:"http,omitempty"`

						Tcp *string `tfsdk:"tcp" yaml:"tcp,omitempty"`
					} `tfsdk:"traffic" yaml:"traffic,omitempty"`
				} `tfsdk:"graph" yaml:"graph,omitempty"`

				Metrics_inbound *struct {
					Aggregations *[]struct {
						Label *string `tfsdk:"label" yaml:"label,omitempty"`

						Display_name *string `tfsdk:"display_name" yaml:"display_name,omitempty"`
					} `tfsdk:"aggregations" yaml:"aggregations,omitempty"`
				} `tfsdk:"metrics_inbound" yaml:"metrics_inbound,omitempty"`

				Metrics_outbound *struct {
					Aggregations *[]struct {
						Display_name *string `tfsdk:"display_name" yaml:"display_name,omitempty"`

						Label *string `tfsdk:"label" yaml:"label,omitempty"`
					} `tfsdk:"aggregations" yaml:"aggregations,omitempty"`
				} `tfsdk:"metrics_outbound" yaml:"metrics_outbound,omitempty"`

				Metrics_per_refresh *string `tfsdk:"metrics_per_refresh" yaml:"metrics_per_refresh,omitempty"`
			} `tfsdk:"ui_defaults" yaml:"ui_defaults,omitempty"`
		} `tfsdk:"kiali_feature_flags" yaml:"kiali_feature_flags,omitempty"`

		Login_token *struct {
			Expiration_seconds *int64 `tfsdk:"expiration_seconds" yaml:"expiration_seconds,omitempty"`

			Signing_key *string `tfsdk:"signing_key" yaml:"signing_key,omitempty"`
		} `tfsdk:"login_token" yaml:"login_token,omitempty"`

		Additional_display_details *[]struct {
			Annotation *string `tfsdk:"annotation" yaml:"annotation,omitempty"`

			Icon_annotation *string `tfsdk:"icon_annotation" yaml:"icon_annotation,omitempty"`

			Title *string `tfsdk:"title" yaml:"title,omitempty"`
		} `tfsdk:"additional_display_details" yaml:"additional_display_details,omitempty"`

		Custom_dashboards *[]map[string]string `tfsdk:"custom_dashboards" yaml:"custom_dashboards,omitempty"`

		External_services *struct {
			Grafana *struct {
				Dashboards *[]struct {
					Variables *struct {
						Workload *string `tfsdk:"workload" yaml:"workload,omitempty"`

						App *string `tfsdk:"app" yaml:"app,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

						Service *string `tfsdk:"service" yaml:"service,omitempty"`
					} `tfsdk:"variables" yaml:"variables,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`
				} `tfsdk:"dashboards" yaml:"dashboards,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Health_check_url *string `tfsdk:"health_check_url" yaml:"health_check_url,omitempty"`

				In_cluster_url *string `tfsdk:"in_cluster_url" yaml:"in_cluster_url,omitempty"`

				Is_core *bool `tfsdk:"is_core" yaml:"is_core,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`

				Auth *struct {
					Use_kiali_token *bool `tfsdk:"use_kiali_token" yaml:"use_kiali_token,omitempty"`

					Username *string `tfsdk:"username" yaml:"username,omitempty"`

					Ca_file *string `tfsdk:"ca_file" yaml:"ca_file,omitempty"`

					Insecure_skip_verify *bool `tfsdk:"insecure_skip_verify" yaml:"insecure_skip_verify,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					Token *string `tfsdk:"token" yaml:"token,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`
			} `tfsdk:"grafana" yaml:"grafana,omitempty"`

			Istio *struct {
				Envoy_admin_local_port *int64 `tfsdk:"envoy_admin_local_port" yaml:"envoy_admin_local_port,omitempty"`

				Istio_canary_revision *struct {
					Current *string `tfsdk:"current" yaml:"current,omitempty"`

					Upgrade *string `tfsdk:"upgrade" yaml:"upgrade,omitempty"`
				} `tfsdk:"istio_canary_revision" yaml:"istio_canary_revision,omitempty"`

				Istio_injection_annotation *string `tfsdk:"istio_injection_annotation" yaml:"istio_injection_annotation,omitempty"`

				Istio_sidecar_annotation *string `tfsdk:"istio_sidecar_annotation" yaml:"istio_sidecar_annotation,omitempty"`

				Istio_sidecar_injector_config_map_name *string `tfsdk:"istio_sidecar_injector_config_map_name" yaml:"istio_sidecar_injector_config_map_name,omitempty"`

				Istiod_deployment_name *string `tfsdk:"istiod_deployment_name" yaml:"istiod_deployment_name,omitempty"`

				Istiod_pod_monitoring_port *int64 `tfsdk:"istiod_pod_monitoring_port" yaml:"istiod_pod_monitoring_port,omitempty"`

				Root_namespace *string `tfsdk:"root_namespace" yaml:"root_namespace,omitempty"`

				Component_status *struct {
					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

					Components *[]struct {
						App_label *string `tfsdk:"app_label" yaml:"app_label,omitempty"`

						Is_core *bool `tfsdk:"is_core" yaml:"is_core,omitempty"`

						Is_proxy *bool `tfsdk:"is_proxy" yaml:"is_proxy,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
					} `tfsdk:"components" yaml:"components,omitempty"`
				} `tfsdk:"component_status" yaml:"component_status,omitempty"`

				Config_map_name *string `tfsdk:"config_map_name" yaml:"config_map_name,omitempty"`

				Istio_identity_domain *string `tfsdk:"istio_identity_domain" yaml:"istio_identity_domain,omitempty"`

				Url_service_version *string `tfsdk:"url_service_version" yaml:"url_service_version,omitempty"`
			} `tfsdk:"istio" yaml:"istio,omitempty"`

			Prometheus *struct {
				Health_check_url *string `tfsdk:"health_check_url" yaml:"health_check_url,omitempty"`

				Is_core *bool `tfsdk:"is_core" yaml:"is_core,omitempty"`

				Query_scope *map[string]string `tfsdk:"query_scope" yaml:"query_scope,omitempty"`

				Url *string `tfsdk:"url" yaml:"url,omitempty"`

				Cache_duration *int64 `tfsdk:"cache_duration" yaml:"cache_duration,omitempty"`

				Cache_enabled *bool `tfsdk:"cache_enabled" yaml:"cache_enabled,omitempty"`

				Cache_expiration *int64 `tfsdk:"cache_expiration" yaml:"cache_expiration,omitempty"`

				Custom_headers *map[string]string `tfsdk:"custom_headers" yaml:"custom_headers,omitempty"`

				Thanos_proxy *struct {
					Retention_period *string `tfsdk:"retention_period" yaml:"retention_period,omitempty"`

					Scrape_interval *string `tfsdk:"scrape_interval" yaml:"scrape_interval,omitempty"`

					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
				} `tfsdk:"thanos_proxy" yaml:"thanos_proxy,omitempty"`

				Auth *struct {
					Ca_file *string `tfsdk:"ca_file" yaml:"ca_file,omitempty"`

					Insecure_skip_verify *bool `tfsdk:"insecure_skip_verify" yaml:"insecure_skip_verify,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					Token *string `tfsdk:"token" yaml:"token,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					Use_kiali_token *bool `tfsdk:"use_kiali_token" yaml:"use_kiali_token,omitempty"`

					Username *string `tfsdk:"username" yaml:"username,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`
			} `tfsdk:"prometheus" yaml:"prometheus,omitempty"`

			Tracing *struct {
				Url *string `tfsdk:"url" yaml:"url,omitempty"`

				Auth *struct {
					Type *string `tfsdk:"type" yaml:"type,omitempty"`

					Use_kiali_token *bool `tfsdk:"use_kiali_token" yaml:"use_kiali_token,omitempty"`

					Username *string `tfsdk:"username" yaml:"username,omitempty"`

					Ca_file *string `tfsdk:"ca_file" yaml:"ca_file,omitempty"`

					Insecure_skip_verify *bool `tfsdk:"insecure_skip_verify" yaml:"insecure_skip_verify,omitempty"`

					Password *string `tfsdk:"password" yaml:"password,omitempty"`

					Token *string `tfsdk:"token" yaml:"token,omitempty"`
				} `tfsdk:"auth" yaml:"auth,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				In_cluster_url *string `tfsdk:"in_cluster_url" yaml:"in_cluster_url,omitempty"`

				Namespace_selector *bool `tfsdk:"namespace_selector" yaml:"namespace_selector,omitempty"`

				Query_scope *map[string]string `tfsdk:"query_scope" yaml:"query_scope,omitempty"`

				Is_core *bool `tfsdk:"is_core" yaml:"is_core,omitempty"`

				Use_grpc *bool `tfsdk:"use_grpc" yaml:"use_grpc,omitempty"`

				Whitelist_istio_system *[]string `tfsdk:"whitelist_istio_system" yaml:"whitelist_istio_system,omitempty"`
			} `tfsdk:"tracing" yaml:"tracing,omitempty"`

			Custom_dashboards *struct {
				Discovery_enabled *string `tfsdk:"discovery_enabled" yaml:"discovery_enabled,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Is_core *bool `tfsdk:"is_core" yaml:"is_core,omitempty"`

				Namespace_label *string `tfsdk:"namespace_label" yaml:"namespace_label,omitempty"`

				Prometheus *struct {
					Cache_enabled *bool `tfsdk:"cache_enabled" yaml:"cache_enabled,omitempty"`

					Thanos_proxy *struct {
						Scrape_interval *string `tfsdk:"scrape_interval" yaml:"scrape_interval,omitempty"`

						Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

						Retention_period *string `tfsdk:"retention_period" yaml:"retention_period,omitempty"`
					} `tfsdk:"thanos_proxy" yaml:"thanos_proxy,omitempty"`

					Url *string `tfsdk:"url" yaml:"url,omitempty"`

					Auth *struct {
						Token *string `tfsdk:"token" yaml:"token,omitempty"`

						Type *string `tfsdk:"type" yaml:"type,omitempty"`

						Use_kiali_token *bool `tfsdk:"use_kiali_token" yaml:"use_kiali_token,omitempty"`

						Username *string `tfsdk:"username" yaml:"username,omitempty"`

						Ca_file *string `tfsdk:"ca_file" yaml:"ca_file,omitempty"`

						Insecure_skip_verify *bool `tfsdk:"insecure_skip_verify" yaml:"insecure_skip_verify,omitempty"`

						Password *string `tfsdk:"password" yaml:"password,omitempty"`
					} `tfsdk:"auth" yaml:"auth,omitempty"`

					Cache_duration *int64 `tfsdk:"cache_duration" yaml:"cache_duration,omitempty"`

					Cache_expiration *int64 `tfsdk:"cache_expiration" yaml:"cache_expiration,omitempty"`

					Custom_headers *map[string]string `tfsdk:"custom_headers" yaml:"custom_headers,omitempty"`

					Health_check_url *string `tfsdk:"health_check_url" yaml:"health_check_url,omitempty"`

					Is_core *bool `tfsdk:"is_core" yaml:"is_core,omitempty"`

					Query_scope *map[string]string `tfsdk:"query_scope" yaml:"query_scope,omitempty"`
				} `tfsdk:"prometheus" yaml:"prometheus,omitempty"`

				Discovery_auto_threshold *int64 `tfsdk:"discovery_auto_threshold" yaml:"discovery_auto_threshold,omitempty"`
			} `tfsdk:"custom_dashboards" yaml:"custom_dashboards,omitempty"`
		} `tfsdk:"external_services" yaml:"external_services,omitempty"`

		Server *struct {
			Gzip_enabled *bool `tfsdk:"gzip_enabled" yaml:"gzip_enabled,omitempty"`

			Observability *struct {
				Tracing *struct {
					Collector_url *string `tfsdk:"collector_url" yaml:"collector_url,omitempty"`

					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
				} `tfsdk:"tracing" yaml:"tracing,omitempty"`

				Metrics *struct {
					Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"metrics" yaml:"metrics,omitempty"`
			} `tfsdk:"observability" yaml:"observability,omitempty"`

			Web_history_mode *string `tfsdk:"web_history_mode" yaml:"web_history_mode,omitempty"`

			Web_schema *string `tfsdk:"web_schema" yaml:"web_schema,omitempty"`

			Address *string `tfsdk:"address" yaml:"address,omitempty"`

			Audit_log *bool `tfsdk:"audit_log" yaml:"audit_log,omitempty"`

			Cors_allow_all *bool `tfsdk:"cors_allow_all" yaml:"cors_allow_all,omitempty"`

			Web_root *string `tfsdk:"web_root" yaml:"web_root,omitempty"`

			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			Web_fqdn *string `tfsdk:"web_fqdn" yaml:"web_fqdn,omitempty"`

			Web_port *string `tfsdk:"web_port" yaml:"web_port,omitempty"`
		} `tfsdk:"server" yaml:"server,omitempty"`

		Api *struct {
			Namespaces *struct {
				Exclude *[]string `tfsdk:"exclude" yaml:"exclude,omitempty"`

				Label_selector *string `tfsdk:"label_selector" yaml:"label_selector,omitempty"`
			} `tfsdk:"namespaces" yaml:"namespaces,omitempty"`
		} `tfsdk:"api" yaml:"api,omitempty"`

		Deployment *struct {
			Instance_name *string `tfsdk:"instance_name" yaml:"instance_name,omitempty"`

			Priority_class_name *string `tfsdk:"priority_class_name" yaml:"priority_class_name,omitempty"`

			Service_type *string `tfsdk:"service_type" yaml:"service_type,omitempty"`

			Tolerations *[]map[string]string `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			Service_annotations *map[string]string `tfsdk:"service_annotations" yaml:"service_annotations,omitempty"`

			Version_label *string `tfsdk:"version_label" yaml:"version_label,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			Resources *map[string]string `tfsdk:"resources" yaml:"resources,omitempty"`

			Security_context *map[string]string `tfsdk:"security_context" yaml:"security_context,omitempty"`

			Affinity *struct {
				Node *map[string]string `tfsdk:"node" yaml:"node,omitempty"`

				Pod *map[string]string `tfsdk:"pod" yaml:"pod,omitempty"`

				Pod_anti *map[string]string `tfsdk:"pod_anti" yaml:"pod_anti,omitempty"`
			} `tfsdk:"affinity" yaml:"affinity,omitempty"`

			Custom_secrets *[]struct {
				Mount *string `tfsdk:"mount" yaml:"mount,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
			} `tfsdk:"custom_secrets" yaml:"custom_secrets,omitempty"`

			Image_name *string `tfsdk:"image_name" yaml:"image_name,omitempty"`

			Image_pull_policy *string `tfsdk:"image_pull_policy" yaml:"image_pull_policy,omitempty"`

			Configmap_annotations *map[string]string `tfsdk:"configmap_annotations" yaml:"configmap_annotations,omitempty"`

			Host_aliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" yaml:"hostnames,omitempty"`

				Ip *string `tfsdk:"ip" yaml:"ip,omitempty"`
			} `tfsdk:"host_aliases" yaml:"host_aliases,omitempty"`

			Secret_name *string `tfsdk:"secret_name" yaml:"secret_name,omitempty"`

			Image_digest *string `tfsdk:"image_digest" yaml:"image_digest,omitempty"`

			Image_version *string `tfsdk:"image_version" yaml:"image_version,omitempty"`

			Logger *struct {
				Time_field_format *string `tfsdk:"time_field_format" yaml:"time_field_format,omitempty"`

				Log_format *string `tfsdk:"log_format" yaml:"log_format,omitempty"`

				Log_level *string `tfsdk:"log_level" yaml:"log_level,omitempty"`

				Sampler_rate *string `tfsdk:"sampler_rate" yaml:"sampler_rate,omitempty"`
			} `tfsdk:"logger" yaml:"logger,omitempty"`

			Pod_annotations *map[string]string `tfsdk:"pod_annotations" yaml:"pod_annotations,omitempty"`

			Image_pull_secrets *[]string `tfsdk:"image_pull_secrets" yaml:"image_pull_secrets,omitempty"`

			Ingress *struct {
				Additional_labels *map[string]string `tfsdk:"additional_labels" yaml:"additional_labels,omitempty"`

				Class_name *string `tfsdk:"class_name" yaml:"class_name,omitempty"`

				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

				Override_yaml *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`
					} `tfsdk:"metadata" yaml:"metadata,omitempty"`

					Spec *map[string]string `tfsdk:"spec" yaml:"spec,omitempty"`
				} `tfsdk:"override_yaml" yaml:"override_yaml,omitempty"`
			} `tfsdk:"ingress" yaml:"ingress,omitempty"`

			Pod_labels *map[string]string `tfsdk:"pod_labels" yaml:"pod_labels,omitempty"`

			Verbose_mode *string `tfsdk:"verbose_mode" yaml:"verbose_mode,omitempty"`

			View_only_mode *bool `tfsdk:"view_only_mode" yaml:"view_only_mode,omitempty"`

			Accessible_namespaces *[]string `tfsdk:"accessible_namespaces" yaml:"accessible_namespaces,omitempty"`

			Additional_service_yaml *map[string]string `tfsdk:"additional_service_yaml" yaml:"additional_service_yaml,omitempty"`

			Hpa *struct {
				Spec *map[string]string `tfsdk:"spec" yaml:"spec,omitempty"`

				Api_version *string `tfsdk:"api_version" yaml:"api_version,omitempty"`
			} `tfsdk:"hpa" yaml:"hpa,omitempty"`

			Node_selector *map[string]string `tfsdk:"node_selector" yaml:"node_selector,omitempty"`
		} `tfsdk:"deployment" yaml:"deployment,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`

		Istio_labels *struct {
			Injection_label_rev *string `tfsdk:"injection_label_rev" yaml:"injection_label_rev,omitempty"`

			Version_label_name *string `tfsdk:"version_label_name" yaml:"version_label_name,omitempty"`

			App_label_name *string `tfsdk:"app_label_name" yaml:"app_label_name,omitempty"`

			Injection_label_name *string `tfsdk:"injection_label_name" yaml:"injection_label_name,omitempty"`
		} `tfsdk:"istio_labels" yaml:"istio_labels,omitempty"`

		Kubernetes_config *struct {
			Qps *int64 `tfsdk:"qps" yaml:"qps,omitempty"`

			Burst *int64 `tfsdk:"burst" yaml:"burst,omitempty"`

			Cache_duration *int64 `tfsdk:"cache_duration" yaml:"cache_duration,omitempty"`

			Cache_enabled *bool `tfsdk:"cache_enabled" yaml:"cache_enabled,omitempty"`

			Cache_istio_types *[]string `tfsdk:"cache_istio_types" yaml:"cache_istio_types,omitempty"`

			Cache_namespaces *[]string `tfsdk:"cache_namespaces" yaml:"cache_namespaces,omitempty"`

			Cache_token_namespace_duration *int64 `tfsdk:"cache_token_namespace_duration" yaml:"cache_token_namespace_duration,omitempty"`

			Excluded_workloads *[]string `tfsdk:"excluded_workloads" yaml:"excluded_workloads,omitempty"`
		} `tfsdk:"kubernetes_config" yaml:"kubernetes_config,omitempty"`

		Identity *struct {
			Cert_file *string `tfsdk:"cert_file" yaml:"cert_file,omitempty"`

			Private_key_file *string `tfsdk:"private_key_file" yaml:"private_key_file,omitempty"`
		} `tfsdk:"identity" yaml:"identity,omitempty"`

		Installation_tag *string `tfsdk:"installation_tag" yaml:"installation_tag,omitempty"`

		Auth *struct {
			Openshift *struct {
				Client_id_prefix *string `tfsdk:"client_id_prefix" yaml:"client_id_prefix,omitempty"`
			} `tfsdk:"openshift" yaml:"openshift,omitempty"`

			Strategy *string `tfsdk:"strategy" yaml:"strategy,omitempty"`

			Openid *struct {
				Api_proxy *string `tfsdk:"api_proxy" yaml:"api_proxy,omitempty"`

				Http_proxy *string `tfsdk:"http_proxy" yaml:"http_proxy,omitempty"`

				Https_proxy *string `tfsdk:"https_proxy" yaml:"https_proxy,omitempty"`

				Insecure_skip_verify_tls *bool `tfsdk:"insecure_skip_verify_tls" yaml:"insecure_skip_verify_tls,omitempty"`

				Username_claim *string `tfsdk:"username_claim" yaml:"username_claim,omitempty"`

				Additional_request_params *map[string]string `tfsdk:"additional_request_params" yaml:"additional_request_params,omitempty"`

				Authorization_endpoint *string `tfsdk:"authorization_endpoint" yaml:"authorization_endpoint,omitempty"`

				Client_id *string `tfsdk:"client_id" yaml:"client_id,omitempty"`

				Allowed_domains *[]string `tfsdk:"allowed_domains" yaml:"allowed_domains,omitempty"`

				Api_proxy_ca_data *string `tfsdk:"api_proxy_ca_data" yaml:"api_proxy_ca_data,omitempty"`

				Api_token *string `tfsdk:"api_token" yaml:"api_token,omitempty"`

				Disable_rbac *bool `tfsdk:"disable_rbac" yaml:"disable_rbac,omitempty"`

				Scopes *[]string `tfsdk:"scopes" yaml:"scopes,omitempty"`

				Authentication_timeout *int64 `tfsdk:"authentication_timeout" yaml:"authentication_timeout,omitempty"`

				Issuer_uri *string `tfsdk:"issuer_uri" yaml:"issuer_uri,omitempty"`
			} `tfsdk:"openid" yaml:"openid,omitempty"`
		} `tfsdk:"auth" yaml:"auth,omitempty"`

		Health_config *struct {
			Rate *[]struct {
				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				Tolerance *[]struct {
					Code *string `tfsdk:"code" yaml:"code,omitempty"`

					Degraded *int64 `tfsdk:"degraded" yaml:"degraded,omitempty"`

					Direction *string `tfsdk:"direction" yaml:"direction,omitempty"`

					Failure *int64 `tfsdk:"failure" yaml:"failure,omitempty"`

					Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
				} `tfsdk:"tolerance" yaml:"tolerance,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`
			} `tfsdk:"rate" yaml:"rate,omitempty"`
		} `tfsdk:"health_config" yaml:"health_config,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewKialiIoKialiV1Alpha1Resource() resource.Resource {
	return &KialiIoKialiV1Alpha1Resource{}
}

func (r *KialiIoKialiV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_kiali_io_kiali_v1alpha1"
}

func (r *KialiIoKialiV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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
				Description:         "This is the CRD for the resources called Kiali CRs. The Kiali Operator will watch for resources of this type and when it detects a Kiali CR has been added, deleted, or modified, it will install, uninstall, and update the associated Kiali Server installation. The settings here will configure the Kiali Server as well as the Kiali Operator. All of these settings will be stored in the Kiali ConfigMap. Do not modify the ConfigMap; it will be managed by the Kiali Operator. Only modify the Kiali CR when you want to change a configuration setting.",
				MarkdownDescription: "This is the CRD for the resources called Kiali CRs. The Kiali Operator will watch for resources of this type and when it detects a Kiali CR has been added, deleted, or modified, it will install, uninstall, and update the associated Kiali Server installation. The settings here will configure the Kiali Server as well as the Kiali Operator. All of these settings will be stored in the Kiali ConfigMap. Do not modify the ConfigMap; it will be managed by the Kiali Operator. Only modify the Kiali CR when you want to change a configuration setting.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"istio_namespace": {
						Description:         "The namespace where Istio is installed. If left empty, it is assumed to be the same namespace as where Kiali is installed (i.e. 'deployment.namespace').",
						MarkdownDescription: "The namespace where Istio is installed. If left empty, it is assumed to be the same namespace as where Kiali is installed (i.e. 'deployment.namespace').",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kiali_feature_flags": {
						Description:         "Kiali features that can be enabled or disabled.",
						MarkdownDescription: "Kiali features that can be enabled or disabled.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"validations": {
								Description:         "Features specific to the validations subsystem.",
								MarkdownDescription: "Features specific to the validations subsystem.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ignore": {
										Description:         "A list of one or more validation codes whose errors are to be ignored.",
										MarkdownDescription: "A list of one or more validation codes whose errors are to be ignored.",

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

							"certificates_information_indicators": {
								Description:         "Flag to enable/disable displaying certificates information and which secrets to grant read permissions.",
								MarkdownDescription: "Flag to enable/disable displaying certificates information and which secrets to grant read permissions.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secrets": {
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

							"clustering": {
								Description:         "Clustering and federation related features.",
								MarkdownDescription: "Clustering and federation related features.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Flag to enable/disable clustering and federation related features.",
										MarkdownDescription: "Flag to enable/disable clustering and federation related features.",

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

							"disabled_features": {
								Description:         "There may be some features that admins do not want to be accessible to users (even in 'view only' mode). In this case, this setting allows you to disable one or more of those features entirely.",
								MarkdownDescription: "There may be some features that admins do not want to be accessible to users (even in 'view only' mode). In this case, this setting allows you to disable one or more of those features entirely.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"istio_injection_action": {
								Description:         "Flag to enable/disable an Action to label a namespace for automatic Istio Sidecar injection.",
								MarkdownDescription: "Flag to enable/disable an Action to label a namespace for automatic Istio Sidecar injection.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"istio_upgrade_action": {
								Description:         "Flag to activate the Kiali functionality of upgrading namespaces to point to an installed Istio Canary revision. Related Canary upgrade and current revisions of Istio should be defined in 'istio_canary_revision' section.",
								MarkdownDescription: "Flag to activate the Kiali functionality of upgrading namespaces to point to an installed Istio Canary revision. Related Canary upgrade and current revisions of Istio should be defined in 'istio_canary_revision' section.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ui_defaults": {
								Description:         "Default settings for the UI. These defaults apply to all users.",
								MarkdownDescription: "Default settings for the UI. These defaults apply to all users.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"namespaces": {
										Description:         "Default selections for the namespace selection dropdown. Non-existent or inaccessible namespaces will be ignored. Omit or set to an empty array for no default namespaces.",
										MarkdownDescription: "Default selections for the namespace selection dropdown. Non-existent or inaccessible namespaces will be ignored. Omit or set to an empty array for no default namespaces.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"refresh_interval": {
										Description:         "The automatic refresh interval for pages offering automatic refresh. Value must be one of: 'pause', '10s', '15s', '30s', '1m', '5m' or '15m'",
										MarkdownDescription: "The automatic refresh interval for pages offering automatic refresh. Value must be one of: 'pause', '10s', '15s', '30s', '1m', '5m' or '15m'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"graph": {
										Description:         "Default settings for the Graph UI.",
										MarkdownDescription: "Default settings for the Graph UI.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"find_options": {
												Description:         "A list of commonly used and useful find expressions that will be provided to the user out-of-box.",
												MarkdownDescription: "A list of commonly used and useful find expressions that will be provided to the user out-of-box.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"description": {
														Description:         "Human-readable text to let the user know what the expression does.",
														MarkdownDescription: "Human-readable text to let the user know what the expression does.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression": {
														Description:         "The find expression.",
														MarkdownDescription: "The find expression.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"hide_options": {
												Description:         "A list of commonly used and useful hide expressions that will be provided to the user out-of-box.",
												MarkdownDescription: "A list of commonly used and useful hide expressions that will be provided to the user out-of-box.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"description": {
														Description:         "Human-readable text to let the user know what the expression does.",
														MarkdownDescription: "Human-readable text to let the user know what the expression does.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"expression": {
														Description:         "The hide expression.",
														MarkdownDescription: "The hide expression.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"traffic": {
												Description:         "These settings determine which rates are used to determine graph traffic.",
												MarkdownDescription: "These settings determine which rates are used to determine graph traffic.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"grpc": {
														Description:         "gRPC traffic is measured in requests or sent/received/total messages. Value must be one of: 'none', 'requests', 'sent', 'received', or 'total'.",
														MarkdownDescription: "gRPC traffic is measured in requests or sent/received/total messages. Value must be one of: 'none', 'requests', 'sent', 'received', or 'total'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"http": {
														Description:         "HTTP traffic is measured in requests. Value must be one of: 'none' or 'requests'.",
														MarkdownDescription: "HTTP traffic is measured in requests. Value must be one of: 'none' or 'requests'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"tcp": {
														Description:         "TCP traffic is measured in sent/received/total bytes. Only request traffic supplies response codes. Value must be one of: 'none', 'sent', 'received', or 'total'.",
														MarkdownDescription: "TCP traffic is measured in sent/received/total bytes. Only request traffic supplies response codes. Value must be one of: 'none', 'sent', 'received', or 'total'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"metrics_inbound": {
										Description:         "Additional label aggregation for inbound metric pages in detail pages.You will see these configurations in the 'Metric Settings' drop-down.An example,'''metrics_inbound:  aggregations:  - display_name: Istio Network    label: topology_istio_io_network  - display_name: Istio Revision    label: istio_io_rev'''",
										MarkdownDescription: "Additional label aggregation for inbound metric pages in detail pages.You will see these configurations in the 'Metric Settings' drop-down.An example,'''metrics_inbound:  aggregations:  - display_name: Istio Network    label: topology_istio_io_network  - display_name: Istio Revision    label: istio_io_rev'''",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"aggregations": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"label": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"display_name": {
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

									"metrics_outbound": {
										Description:         "Additional label aggregation for outbound metric pages in detail pages.You will see these configurations in the 'Metric Settings' drop-down.An example,'''metrics_outbound:  aggregations:  - display_name: Istio Network    label: topology_istio_io_network  - display_name: Istio Revision    label: istio_io_rev'''",
										MarkdownDescription: "Additional label aggregation for outbound metric pages in detail pages.You will see these configurations in the 'Metric Settings' drop-down.An example,'''metrics_outbound:  aggregations:  - display_name: Istio Network    label: topology_istio_io_network  - display_name: Istio Revision    label: istio_io_rev'''",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"aggregations": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"display_name": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label": {
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

									"metrics_per_refresh": {
										Description:         "Duration of metrics to fetch on each refresh. Value must be one of: '1m', '2m', '5m', '10m', '30m', '1h', '3h', '6h', '12h', '1d', '7d', or '30d'",
										MarkdownDescription: "Duration of metrics to fetch on each refresh. Value must be one of: '1m', '2m', '5m', '10m', '30m', '1h', '3h', '6h', '12h', '1d', '7d', or '30d'",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"login_token": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"expiration_seconds": {
								Description:         "A user's login token expiration specified in seconds. This is applicable to token and header auth strategies only.",
								MarkdownDescription: "A user's login token expiration specified in seconds. This is applicable to token and header auth strategies only.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"signing_key": {
								Description:         "The signing key used to generate tokens for user authentication. Because this is potentially sensitive, you have the option to store this value in a secret. If you store this signing key value in a secret, you must indicate what key in what secret by setting this value to a string in the form of 'secret:<secretName>:<secretKey>'. If left as an empty string, a secret with a random signing key will be generated for you. The signing key must be 16, 24 or 32 byte long.",
								MarkdownDescription: "The signing key used to generate tokens for user authentication. Because this is potentially sensitive, you have the option to store this value in a secret. If you store this signing key value in a secret, you must indicate what key in what secret by setting this value to a string in the form of 'secret:<secretName>:<secretKey>'. If left as an empty string, a secret with a random signing key will be generated for you. The signing key must be 16, 24 or 32 byte long.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"additional_display_details": {
						Description:         "A list of additional details that Kiali will look for in annotations. When found on any workload or service, Kiali will display the additional details in the respective workload or service details page. This is typically used to inject some CI metadata or documentation links into Kiali views. For example, by default, Kiali will recognize these annotations on a service or workload (e.g. a Deployment, StatefulSet, etc.):'''annotations:  kiali.io/api-spec: http://list/to/my/api/doc  kiali.io/api-type: rest'''Note that if you change this setting for your own custom annotations, keep in mind that it would override the current default. So you would have to add the default setting as shown in the example CR if you want to preserve the default links.",
						MarkdownDescription: "A list of additional details that Kiali will look for in annotations. When found on any workload or service, Kiali will display the additional details in the respective workload or service details page. This is typically used to inject some CI metadata or documentation links into Kiali views. For example, by default, Kiali will recognize these annotations on a service or workload (e.g. a Deployment, StatefulSet, etc.):'''annotations:  kiali.io/api-spec: http://list/to/my/api/doc  kiali.io/api-type: rest'''Note that if you change this setting for your own custom annotations, keep in mind that it would override the current default. So you would have to add the default setting as shown in the example CR if you want to preserve the default links.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"annotation": {
								Description:         "The name of the annotation whose value is a URL to additional documentation useful to the user.",
								MarkdownDescription: "The name of the annotation whose value is a URL to additional documentation useful to the user.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"icon_annotation": {
								Description:         "The name of the annotation whose value is used to determine what icon to display. The annotation name itself can be anything, but note that the value of that annotation must be one of: 'rest', 'grpc', and 'graphql' - any other value is ignored.",
								MarkdownDescription: "The name of the annotation whose value is used to determine what icon to display. The annotation name itself can be anything, but note that the value of that annotation must be one of: 'rest', 'grpc', and 'graphql' - any other value is ignored.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"title": {
								Description:         "The title of the link that Kiali will display. The link will go to the URL specified in the value of the configured 'annotation'.",
								MarkdownDescription: "The title of the link that Kiali will display. The link will go to the URL specified in the value of the configured 'annotation'.",

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

					"custom_dashboards": {
						Description:         "A list of user-defined custom monitoring dashboards that you can use to generate metrics chartsfor your applications. The server has some built-in dashboards; if you define a custom dashboard herewith the same name as a built-in dashboard, your custom dashboard takes precedence and will overwritethe built-in dashboard. You can disable one or more of the built-in dashboards by simply defining anempty dashboard.An example of an additional user-defined dashboard,'''- name: myapp  title: My App Metrics  items:  - chart:      name: 'Thread Count'      spans: 4      metricName: 'thread-count'      dataType: 'raw''''An example of disabling a built-in dashboard (in this case, disabling the Envoy dashboard),'''- name: envoy'''To learn more about custom monitoring dashboards, see the documentation at https://kiali.io/docs/configuration/custom-dashboard/",
						MarkdownDescription: "A list of user-defined custom monitoring dashboards that you can use to generate metrics chartsfor your applications. The server has some built-in dashboards; if you define a custom dashboard herewith the same name as a built-in dashboard, your custom dashboard takes precedence and will overwritethe built-in dashboard. You can disable one or more of the built-in dashboards by simply defining anempty dashboard.An example of an additional user-defined dashboard,'''- name: myapp  title: My App Metrics  items:  - chart:      name: 'Thread Count'      spans: 4      metricName: 'thread-count'      dataType: 'raw''''An example of disabling a built-in dashboard (in this case, disabling the Envoy dashboard),'''- name: envoy'''To learn more about custom monitoring dashboards, see the documentation at https://kiali.io/docs/configuration/custom-dashboard/",

						Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_services": {
						Description:         "These external service configuration settings define how to connect to the external serviceslike Prometheus, Grafana, and Jaeger.Regarding sensitive values in the external_services 'auth' sections:Some external services configured below support an 'auth' sub-section in order to tell Kialihow it should authenticate with the external services. Credentials used to authenticate Kialito those external services can be defined in the 'auth.password' and 'auth.token' valueswithin the 'auth' sub-section. Because these are sensitive values, you may not want to declarethe actual credentials here in the Kiali CR. In this case, you may store the actual passwordor token string in a Kubernetes secret. If you do, you need to set the 'auth.password' or'auth.token' to a value in the format 'secret:<secretName>:<secretKey>' where '<secretName>'is the name of the secret object that Kiali can access, and '<secretKey>' is the name of thekey within the named secret that contains the actual password or token string. For example,if Grafana requires a password, you can store that password in a secret named 'myGrafanaCredentials'in a key named 'myGrafanaPw'. In this case, you would set 'external_services.grafana.auth.password'to 'secret:myGrafanaCredentials:myGrafanaPw'.",
						MarkdownDescription: "These external service configuration settings define how to connect to the external serviceslike Prometheus, Grafana, and Jaeger.Regarding sensitive values in the external_services 'auth' sections:Some external services configured below support an 'auth' sub-section in order to tell Kialihow it should authenticate with the external services. Credentials used to authenticate Kialito those external services can be defined in the 'auth.password' and 'auth.token' valueswithin the 'auth' sub-section. Because these are sensitive values, you may not want to declarethe actual credentials here in the Kiali CR. In this case, you may store the actual passwordor token string in a Kubernetes secret. If you do, you need to set the 'auth.password' or'auth.token' to a value in the format 'secret:<secretName>:<secretKey>' where '<secretName>'is the name of the secret object that Kiali can access, and '<secretKey>' is the name of thekey within the named secret that contains the actual password or token string. For example,if Grafana requires a password, you can store that password in a secret named 'myGrafanaCredentials'in a key named 'myGrafanaPw'. In this case, you would set 'external_services.grafana.auth.password'to 'secret:myGrafanaCredentials:myGrafanaPw'.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"grafana": {
								Description:         "Configuration used to access the Grafana dashboards.",
								MarkdownDescription: "Configuration used to access the Grafana dashboards.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"dashboards": {
										Description:         "A list of Grafana dashboards that Kiali can link to.",
										MarkdownDescription: "A list of Grafana dashboards that Kiali can link to.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"variables": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"workload": {
														Description:         "The name of a variable that holds the workload name, if used in that dashboard (else it must be omitted).",
														MarkdownDescription: "The name of a variable that holds the workload name, if used in that dashboard (else it must be omitted).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"app": {
														Description:         "The name of a variable that holds the app name, if used in that dashboard (else it must be omitted).",
														MarkdownDescription: "The name of a variable that holds the app name, if used in that dashboard (else it must be omitted).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "The name of a variable that holds the namespace, if used in that dashboard (else it must be omitted).",
														MarkdownDescription: "The name of a variable that holds the namespace, if used in that dashboard (else it must be omitted).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"service": {
														Description:         "The name of a variable that holds the service name, if used in that dashboard (else it must be omitted).",
														MarkdownDescription: "The name of a variable that holds the service name, if used in that dashboard (else it must be omitted).",

														Type: types.StringType,

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
												Description:         "The name of the Grafana dashboard.",
												MarkdownDescription: "The name of the Grafana dashboard.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "When true, Grafana support will be enabled in Kiali.",
										MarkdownDescription: "When true, Grafana support will be enabled in Kiali.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"health_check_url": {
										Description:         "Used in the Components health feature. This is the URL which Kiali will ping to determine whether the component is reachable or not. It defaults to 'in_cluster_url' when not provided.",
										MarkdownDescription: "Used in the Components health feature. This is the URL which Kiali will ping to determine whether the component is reachable or not. It defaults to 'in_cluster_url' when not provided.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"in_cluster_url": {
										Description:         "The URL used for in-cluster access. An example would be 'http://grafana.istio-system:3000'. This URL can contain query parameters if needed, such as '?orgId=1'. If not defined, it will default to 'http://grafana.<istio_namespace>:3000'.",
										MarkdownDescription: "The URL used for in-cluster access. An example would be 'http://grafana.istio-system:3000'. This URL can contain query parameters if needed, such as '?orgId=1'. If not defined, it will default to 'http://grafana.<istio_namespace>:3000'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"is_core": {
										Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "The URL that Kiali uses when integrating with Grafana. This URL must be accessible to clients external to the cluster in order for the integration to work properly. If empty, an attempt to auto-discover it is made. This URL can contain query parameters if needed, such as '?orgId=1'.",
										MarkdownDescription: "The URL that Kiali uses when integrating with Grafana. This URL must be accessible to clients external to the cluster in order for the integration to work properly. If empty, an attempt to auto-discover it is made. This URL can contain query parameters if needed, such as '?orgId=1'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"auth": {
										Description:         "Settings used to authenticate with the Grafana instance.",
										MarkdownDescription: "Settings used to authenticate with the Grafana instance.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"use_kiali_token": {
												Description:         "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Grafana (in this case, 'auth.token' config is ignored).",
												MarkdownDescription: "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Grafana (in this case, 'auth.token' config is ignored).",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"username": {
												Description:         "Username to be used when making requests to Grafana with 'basic' authentication.",
												MarkdownDescription: "Username to be used when making requests to Grafana with 'basic' authentication.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ca_file": {
												Description:         "The certificate authority file to use when accessing Grafana using https. An empty string means no extra certificate authority file is used.",
												MarkdownDescription: "The certificate authority file to use when accessing Grafana using https. An empty string means no extra certificate authority file is used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_skip_verify": {
												Description:         "Set true to skip verifying certificate validity when Kiali contacts Grafana over https.",
												MarkdownDescription: "Set true to skip verifying certificate validity when Kiali contacts Grafana over https.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "Password to be used when making requests to Grafana, for basic authentication. May refer to a secret.",
												MarkdownDescription: "Password to be used when making requests to Grafana, for basic authentication. May refer to a secret.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"token": {
												Description:         "Token / API key to access Grafana, for token-based authentication. May refer to a secret.",
												MarkdownDescription: "Token / API key to access Grafana, for token-based authentication. May refer to a secret.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Grafana server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
												MarkdownDescription: "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Grafana server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"istio": {
								Description:         "Istio configuration that Kiali needs to know about in order to observe the mesh.",
								MarkdownDescription: "Istio configuration that Kiali needs to know about in order to observe the mesh.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"envoy_admin_local_port": {
										Description:         "The port which kiali will open to fetch envoy config data information.",
										MarkdownDescription: "The port which kiali will open to fetch envoy config data information.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"istio_canary_revision": {
										Description:         "These values are used in Canary upgrade/downgrade functionality when 'istio_upgrade_action' is true.",
										MarkdownDescription: "These values are used in Canary upgrade/downgrade functionality when 'istio_upgrade_action' is true.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"current": {
												Description:         "The currently installed Istio revision.",
												MarkdownDescription: "The currently installed Istio revision.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"upgrade": {
												Description:         "The installed Istio canary revision to upgrade to.",
												MarkdownDescription: "The installed Istio canary revision to upgrade to.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"istio_injection_annotation": {
										Description:         "The name of the field that annotates a workload to indicate a sidecar should be automatically injected by Istio. This is the name of a Kubernetes annotation. Note that some Istio implementations also support labels by the same name. In other words, if a workload has a Kubernetes label with this name, that may also trigger automatic sidecar injection.",
										MarkdownDescription: "The name of the field that annotates a workload to indicate a sidecar should be automatically injected by Istio. This is the name of a Kubernetes annotation. Note that some Istio implementations also support labels by the same name. In other words, if a workload has a Kubernetes label with this name, that may also trigger automatic sidecar injection.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"istio_sidecar_annotation": {
										Description:         "The pod annotation used by Istio to identify the sidecar.",
										MarkdownDescription: "The pod annotation used by Istio to identify the sidecar.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"istio_sidecar_injector_config_map_name": {
										Description:         "The name of the istio-sidecar-injector config map.",
										MarkdownDescription: "The name of the istio-sidecar-injector config map.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"istiod_deployment_name": {
										Description:         "The name of the istiod deployment.",
										MarkdownDescription: "The name of the istiod deployment.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"istiod_pod_monitoring_port": {
										Description:         "The monitoring port of the IstioD pod (not the Service).",
										MarkdownDescription: "The monitoring port of the IstioD pod (not the Service).",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"root_namespace": {
										Description:         "The namespace to treat as the administrative root namespace for Istio configuration.",
										MarkdownDescription: "The namespace to treat as the administrative root namespace for Istio configuration.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"component_status": {
										Description:         "Istio components whose status will be monitored by Kiali.",
										MarkdownDescription: "Istio components whose status will be monitored by Kiali.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enabled": {
												Description:         "Determines if Istio component statuses will be displayed in the Kiali masthead indicator.",
												MarkdownDescription: "Determines if Istio component statuses will be displayed in the Kiali masthead indicator.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"components": {
												Description:         "A specific Istio component whose status will be monitored by Kiali.",
												MarkdownDescription: "A specific Istio component whose status will be monitored by Kiali.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"app_label": {
														Description:         "Istio component pod app label.",
														MarkdownDescription: "Istio component pod app label.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"is_core": {
														Description:         "Whether the component is to be considered a core component for your deployment.",
														MarkdownDescription: "Whether the component is to be considered a core component for your deployment.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"is_proxy": {
														Description:         "Whether the component is a native Envoy proxy.",
														MarkdownDescription: "Whether the component is a native Envoy proxy.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"namespace": {
														Description:         "The namespace where the component is installed. It defaults to the Istio control plane namespace (e.g. 'istio_namespace') setting. Note that the Istio documentation suggests you install the ingress and egress to different namespaces, so you most likely will want to explicitly set this namespace value for the ingress and egress components.",
														MarkdownDescription: "The namespace where the component is installed. It defaults to the Istio control plane namespace (e.g. 'istio_namespace') setting. Note that the Istio documentation suggests you install the ingress and egress to different namespaces, so you most likely will want to explicitly set this namespace value for the ingress and egress components.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"config_map_name": {
										Description:         "The name of the istio control plane config map.",
										MarkdownDescription: "The name of the istio control plane config map.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"istio_identity_domain": {
										Description:         "The annotation used by Istio to identify domains.",
										MarkdownDescription: "The annotation used by Istio to identify domains.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url_service_version": {
										Description:         "The Istio service used to determine the Istio version. If empty, assumes the URL for the well-known Istio version endpoint.",
										MarkdownDescription: "The Istio service used to determine the Istio version. If empty, assumes the URL for the well-known Istio version endpoint.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"prometheus": {
								Description:         "The Prometheus configuration defined here refers to the Prometheus instance that is used by Istio to store its telemetry.",
								MarkdownDescription: "The Prometheus configuration defined here refers to the Prometheus instance that is used by Istio to store its telemetry.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"health_check_url": {
										Description:         "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",
										MarkdownDescription: "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"is_core": {
										Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_scope": {
										Description:         "A set of labelName/labelValue settings applied to every Prometheus query. Used to narrow unified metrics to only those scoped to the Kiali instance.",
										MarkdownDescription: "A set of labelName/labelValue settings applied to every Prometheus query. Used to narrow unified metrics to only those scoped to the Kiali instance.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": {
										Description:         "The URL used to query the Prometheus Server. This URL must be accessible from the Kiali pod. If empty, the default will assume Prometheus is in the Istio control plane namespace; e.g. 'http://prometheus.<istio_namespace>:9090'.",
										MarkdownDescription: "The URL used to query the Prometheus Server. This URL must be accessible from the Kiali pod. If empty, the default will assume Prometheus is in the Istio control plane namespace; e.g. 'http://prometheus.<istio_namespace>:9090'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cache_duration": {
										Description:         "Prometheus caching duration expressed in seconds.",
										MarkdownDescription: "Prometheus caching duration expressed in seconds.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cache_enabled": {
										Description:         "Enable/disable Prometheus caching used for Health services.",
										MarkdownDescription: "Enable/disable Prometheus caching used for Health services.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"cache_expiration": {
										Description:         "Prometheus caching expiration expressed in seconds.",
										MarkdownDescription: "Prometheus caching expiration expressed in seconds.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"custom_headers": {
										Description:         "A set of name/value settings that will be passed as headers when requests are sent to Prometheus.",
										MarkdownDescription: "A set of name/value settings that will be passed as headers when requests are sent to Prometheus.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"thanos_proxy": {
										Description:         "Define this section if Prometheus is to be queried through a Thanos proxy. Kiali will still use the 'url' setting to query for Prometheus metrics so make sure that is set appropriately.",
										MarkdownDescription: "Define this section if Prometheus is to be queried through a Thanos proxy. Kiali will still use the 'url' setting to query for Prometheus metrics so make sure that is set appropriately.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"retention_period": {
												Description:         "Thanos Retention period value expresed as a string.",
												MarkdownDescription: "Thanos Retention period value expresed as a string.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"scrape_interval": {
												Description:         "Thanos Scrape interval value expresed as a string.",
												MarkdownDescription: "Thanos Scrape interval value expresed as a string.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enabled": {
												Description:         "Set to true when a Thanos proxy is in front of Prometheus.",
												MarkdownDescription: "Set to true when a Thanos proxy is in front of Prometheus.",

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

									"auth": {
										Description:         "Settings used to authenticate with the Prometheus instance.",
										MarkdownDescription: "Settings used to authenticate with the Prometheus instance.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"ca_file": {
												Description:         "The certificate authority file to use when accessing Prometheus using https. An empty string means no extra certificate authority file is used.",
												MarkdownDescription: "The certificate authority file to use when accessing Prometheus using https. An empty string means no extra certificate authority file is used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_skip_verify": {
												Description:         "Set true to skip verifying certificate validity when Kiali contacts Prometheus over https.",
												MarkdownDescription: "Set true to skip verifying certificate validity when Kiali contacts Prometheus over https.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "Password to be used when making requests to Prometheus, for basic authentication. May refer to a secret.",
												MarkdownDescription: "Password to be used when making requests to Prometheus, for basic authentication. May refer to a secret.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"token": {
												Description:         "Token / API key to access Prometheus, for token-based authentication. May refer to a secret.",
												MarkdownDescription: "Token / API key to access Prometheus, for token-based authentication. May refer to a secret.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": {
												Description:         "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Prometheus server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
												MarkdownDescription: "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Prometheus server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"use_kiali_token": {
												Description:         "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Prometheus (in this case, 'auth.token' config is ignored).",
												MarkdownDescription: "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Prometheus (in this case, 'auth.token' config is ignored).",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"username": {
												Description:         "Username to be used when making requests to Prometheus with 'basic' authentication.",
												MarkdownDescription: "Username to be used when making requests to Prometheus with 'basic' authentication.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

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
								Description:         "Configuration used to access the Tracing (Jaeger) dashboards.",
								MarkdownDescription: "Configuration used to access the Tracing (Jaeger) dashboards.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"url": {
										Description:         "The external URL that will be used to generate links to Jaeger. It must be accessible to clients external to the cluster (e.g: a browser) in order to generate valid links. If the tracing service is deployed with a QUERY_BASE_PATH set, set this URL like https://<hostname>/<QUERY_BASE_PATH>. For example, https://tracing-service:8080/jaeger",
										MarkdownDescription: "The external URL that will be used to generate links to Jaeger. It must be accessible to clients external to the cluster (e.g: a browser) in order to generate valid links. If the tracing service is deployed with a QUERY_BASE_PATH set, set this URL like https://<hostname>/<QUERY_BASE_PATH>. For example, https://tracing-service:8080/jaeger",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"auth": {
										Description:         "Settings used to authenticate with the Tracing server instance.",
										MarkdownDescription: "Settings used to authenticate with the Tracing server instance.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"type": {
												Description:         "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Tracing server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
												MarkdownDescription: "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Tracing server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"use_kiali_token": {
												Description:         "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to the Tracing server (in this case, 'auth.token' config is ignored).",
												MarkdownDescription: "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to the Tracing server (in this case, 'auth.token' config is ignored).",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"username": {
												Description:         "Username to be used when making requests to the Tracing server with 'basic' authentication.",
												MarkdownDescription: "Username to be used when making requests to the Tracing server with 'basic' authentication.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ca_file": {
												Description:         "The certificate authority file to use when accessing the Tracing server using https. An empty string means no extra certificate authority file is used.",
												MarkdownDescription: "The certificate authority file to use when accessing the Tracing server using https. An empty string means no extra certificate authority file is used.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"insecure_skip_verify": {
												Description:         "Set true to skip verifying certificate validity when Kiali contacts the Tracing server over https.",
												MarkdownDescription: "Set true to skip verifying certificate validity when Kiali contacts the Tracing server over https.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"password": {
												Description:         "Password to be used when making requests to the Tracing server, for basic authentication. May refer to a secret.",
												MarkdownDescription: "Password to be used when making requests to the Tracing server, for basic authentication. May refer to a secret.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"token": {
												Description:         "Token / API key to access the Tracing server, for token-based authentication. May refer to a secret.",
												MarkdownDescription: "Token / API key to access the Tracing server, for token-based authentication. May refer to a secret.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "When true, connections to the Tracing server are enabled. 'in_cluster_url' and/or 'url' need to be provided.",
										MarkdownDescription: "When true, connections to the Tracing server are enabled. 'in_cluster_url' and/or 'url' need to be provided.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"in_cluster_url": {
										Description:         "Set URL for in-cluster access, which enables further integration between Kiali and Jaeger. When not provided, Kiali will only show external links using the 'url' setting. Note: Jaeger v1.20+ has separated ports for GRPC(16685) and HTTP(16686) requests. Make sure you use the appropriate port according to the 'use_grpc' value. Example: http://tracing.istio-system:16685",
										MarkdownDescription: "Set URL for in-cluster access, which enables further integration between Kiali and Jaeger. When not provided, Kiali will only show external links using the 'url' setting. Note: Jaeger v1.20+ has separated ports for GRPC(16685) and HTTP(16686) requests. Make sure you use the appropriate port according to the 'use_grpc' value. Example: http://tracing.istio-system:16685",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace_selector": {
										Description:         "Kiali use this boolean to find traces with a namespace selector : service.namespace.",
										MarkdownDescription: "Kiali use this boolean to find traces with a namespace selector : service.namespace.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"query_scope": {
										Description:         "A set of tagKey/tagValue settings applied to every Jaeger query. Used to narrow unified traces to only those scoped to the Kiali instance.",
										MarkdownDescription: "A set of tagKey/tagValue settings applied to every Jaeger query. Used to narrow unified traces to only those scoped to the Kiali instance.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"is_core": {
										Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_grpc": {
										Description:         "Set to true in order to enable GRPC connections between Kiali and Jaeger which will speed up the queries. In some setups you might not be able to use GRPC (e.g. if Jaeger is behind some reverse proxy that doesn't support it). If not specified, this will defalt to 'false' if deployed within a Maistra/OSSM+OpenShift environment, 'true' otherwise.",
										MarkdownDescription: "Set to true in order to enable GRPC connections between Kiali and Jaeger which will speed up the queries. In some setups you might not be able to use GRPC (e.g. if Jaeger is behind some reverse proxy that doesn't support it). If not specified, this will defalt to 'false' if deployed within a Maistra/OSSM+OpenShift environment, 'true' otherwise.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"whitelist_istio_system": {
										Description:         "Kiali will get the traces of these services found in the Istio control plane namespace.",
										MarkdownDescription: "Kiali will get the traces of these services found in the Istio control plane namespace.",

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

							"custom_dashboards": {
								Description:         "Settings for enabling and discovering custom dashboards.",
								MarkdownDescription: "Settings for enabling and discovering custom dashboards.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"discovery_enabled": {
										Description:         "Enable, disable or set 'auto' mode to the dashboards discovery process. If set to 'true', Kiali will always try to discover dashboards based on metrics. Note that this can generate performance penalties while discovering dashboards for workloads having many pods (thus many metrics). When set to 'auto', Kiali will skip dashboards discovery for workloads with more than a configured threshold of pods (see 'discovery_auto_threshold'). When discovery is disabled or auto/skipped, it is still possible to tie workloads with dashboards through annotations on pods (refer to the doc https://kiali.io/docs/configuration/custom-dashboard/#pod-annotations). Value must be one of: 'true', 'false', 'auto'.",
										MarkdownDescription: "Enable, disable or set 'auto' mode to the dashboards discovery process. If set to 'true', Kiali will always try to discover dashboards based on metrics. Note that this can generate performance penalties while discovering dashboards for workloads having many pods (thus many metrics). When set to 'auto', Kiali will skip dashboards discovery for workloads with more than a configured threshold of pods (see 'discovery_auto_threshold'). When discovery is disabled or auto/skipped, it is still possible to tie workloads with dashboards through annotations on pods (refer to the doc https://kiali.io/docs/configuration/custom-dashboard/#pod-annotations). Value must be one of: 'true', 'false', 'auto'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Enable or disable custom dashboards, including the dashboards discovery process.",
										MarkdownDescription: "Enable or disable custom dashboards, including the dashboards discovery process.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"is_core": {
										Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace_label": {
										Description:         "The Prometheus label name used for identifying namespaces in metrics for custom dashboards. The default is 'namespace' but you may want to use 'kubernetes_namespace' depending on your Prometheus configuration.",
										MarkdownDescription: "The Prometheus label name used for identifying namespaces in metrics for custom dashboards. The default is 'namespace' but you may want to use 'kubernetes_namespace' depending on your Prometheus configuration.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"prometheus": {
										Description:         "The Prometheus configuration defined here refers to the Prometheus instance that is dedicated to fetching metrics for custom dashboards. This means you can obtain these metrics for the custom dashboards from a Prometheus instance that is different from the one that Istio uses. If this section is omitted, the same Prometheus that is used to obtain the Istio metrics will also be used for retrieving custom dashboard metrics.",
										MarkdownDescription: "The Prometheus configuration defined here refers to the Prometheus instance that is dedicated to fetching metrics for custom dashboards. This means you can obtain these metrics for the custom dashboards from a Prometheus instance that is different from the one that Istio uses. If this section is omitted, the same Prometheus that is used to obtain the Istio metrics will also be used for retrieving custom dashboard metrics.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"cache_enabled": {
												Description:         "Enable/disable Prometheus caching used for Health services.",
												MarkdownDescription: "Enable/disable Prometheus caching used for Health services.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"thanos_proxy": {
												Description:         "Define this section if Prometheus is to be queried through a Thanos proxy. Kiali will still use the 'url' setting to query for Prometheus metrics so make sure that is set appropriately.",
												MarkdownDescription: "Define this section if Prometheus is to be queried through a Thanos proxy. Kiali will still use the 'url' setting to query for Prometheus metrics so make sure that is set appropriately.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"scrape_interval": {
														Description:         "Thanos Scrape interval value expresed as a string.",
														MarkdownDescription: "Thanos Scrape interval value expresed as a string.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"enabled": {
														Description:         "Set to true when a Thanos proxy is in front of Prometheus.",
														MarkdownDescription: "Set to true when a Thanos proxy is in front of Prometheus.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"retention_period": {
														Description:         "Thanos Retention period value expresed as a string.",
														MarkdownDescription: "Thanos Retention period value expresed as a string.",

														Type: types.StringType,

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
												Description:         "The URL used to query the Prometheus Server. This URL must be accessible from the Kiali pod. If empty, the default will assume Prometheus is in the Istio control plane namespace; e.g. 'http://prometheus.<istio_namespace>:9090'.",
												MarkdownDescription: "The URL used to query the Prometheus Server. This URL must be accessible from the Kiali pod. If empty, the default will assume Prometheus is in the Istio control plane namespace; e.g. 'http://prometheus.<istio_namespace>:9090'.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"auth": {
												Description:         "Settings used to authenticate with the Prometheus instance.",
												MarkdownDescription: "Settings used to authenticate with the Prometheus instance.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"token": {
														Description:         "Token / API key to access Prometheus, for token-based authentication. May refer to a secret.",
														MarkdownDescription: "Token / API key to access Prometheus, for token-based authentication. May refer to a secret.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": {
														Description:         "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Prometheus server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
														MarkdownDescription: "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Prometheus server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"use_kiali_token": {
														Description:         "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Prometheus (in this case, 'auth.token' config is ignored).",
														MarkdownDescription: "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Prometheus (in this case, 'auth.token' config is ignored).",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"username": {
														Description:         "Username to be used when making requests to Prometheus with 'basic' authentication.",
														MarkdownDescription: "Username to be used when making requests to Prometheus with 'basic' authentication.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"ca_file": {
														Description:         "The certificate authority file to use when accessing Prometheus using https. An empty string means no extra certificate authority file is used.",
														MarkdownDescription: "The certificate authority file to use when accessing Prometheus using https. An empty string means no extra certificate authority file is used.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"insecure_skip_verify": {
														Description:         "Set true to skip verifying certificate validity when Kiali contacts Prometheus over https.",
														MarkdownDescription: "Set true to skip verifying certificate validity when Kiali contacts Prometheus over https.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"password": {
														Description:         "Password to be used when making requests to Prometheus, for basic authentication. May refer to a secret.",
														MarkdownDescription: "Password to be used when making requests to Prometheus, for basic authentication. May refer to a secret.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cache_duration": {
												Description:         "Prometheus caching duration expressed in seconds.",
												MarkdownDescription: "Prometheus caching duration expressed in seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"cache_expiration": {
												Description:         "Prometheus caching expiration expressed in seconds.",
												MarkdownDescription: "Prometheus caching expiration expressed in seconds.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"custom_headers": {
												Description:         "A set of name/value settings that will be passed as headers when requests are sent to Prometheus.",
												MarkdownDescription: "A set of name/value settings that will be passed as headers when requests are sent to Prometheus.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"health_check_url": {
												Description:         "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",
												MarkdownDescription: "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"is_core": {
												Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
												MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"query_scope": {
												Description:         "A set of labelName/labelValue settings applied to every Prometheus query. Used to narrow unified metrics to only those scoped to the Kiali instance.",
												MarkdownDescription: "A set of labelName/labelValue settings applied to every Prometheus query. Used to narrow unified metrics to only those scoped to the Kiali instance.",

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

									"discovery_auto_threshold": {
										Description:         "Threshold of the number of pods, for a given Application or Workload, above which dashboards discovery will be skipped. This setting only takes effect when 'discovery_enabled' is set to 'auto'.",
										MarkdownDescription: "Threshold of the number of pods, for a given Application or Workload, above which dashboards discovery will be skipped. This setting only takes effect when 'discovery_enabled' is set to 'auto'.",

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

					"server": {
						Description:         "Configuration that controls some core components within the Kiali Server.",
						MarkdownDescription: "Configuration that controls some core components within the Kiali Server.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"gzip_enabled": {
								Description:         "When true, Kiali serves http requests with gzip enabled (if the browser supports it) when the requests are over 1400 bytes.",
								MarkdownDescription: "When true, Kiali serves http requests with gzip enabled (if the browser supports it) when the requests are over 1400 bytes.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"observability": {
								Description:         "Settings to enable observability into the Kiali server itself.",
								MarkdownDescription: "Settings to enable observability into the Kiali server itself.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"tracing": {
										Description:         "Settings that control how the Kiali server itself emits its own tracing data.",
										MarkdownDescription: "Settings that control how the Kiali server itself emits its own tracing data.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"collector_url": {
												Description:         "The URL used to determine where the Kiali server tracing data will be stored.",
												MarkdownDescription: "The URL used to determine where the Kiali server tracing data will be stored.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"enabled": {
												Description:         "When true, the Kiali server itself will product its own tracing data.",
												MarkdownDescription: "When true, the Kiali server itself will product its own tracing data.",

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

									"metrics": {
										Description:         "Settings that control how Kiali itself emits its own metrics.",
										MarkdownDescription: "Settings that control how Kiali itself emits its own metrics.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"enabled": {
												Description:         "When true, the metrics endpoint will be available for Prometheus to scrape.",
												MarkdownDescription: "When true, the metrics endpoint will be available for Prometheus to scrape.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "The port that the server will bind to in order to receive metric requests. This is the port Prometheus will need to scrape when collecting metrics from Kiali.",
												MarkdownDescription: "The port that the server will bind to in order to receive metric requests. This is the port Prometheus will need to scrape when collecting metrics from Kiali.",

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

							"web_history_mode": {
								Description:         "Define the history mode of kiali UI. Value must be one of: 'browser' or 'hash'.",
								MarkdownDescription: "Define the history mode of kiali UI. Value must be one of: 'browser' or 'hash'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"web_schema": {
								Description:         "Defines the public HTTP schema used to serve Kiali. Value must be one of: 'http' or 'https'. When empty, Kiali will try to guess this value from HTTP headers. On non-OpenShift clusters, you must populate this value if you want to enable cross-linking between Kiali instances in a multi-cluster setup.",
								MarkdownDescription: "Defines the public HTTP schema used to serve Kiali. Value must be one of: 'http' or 'https'. When empty, Kiali will try to guess this value from HTTP headers. On non-OpenShift clusters, you must populate this value if you want to enable cross-linking between Kiali instances in a multi-cluster setup.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"address": {
								Description:         "Where the Kiali server is bound. The console and API server are accessible on this host.",
								MarkdownDescription: "Where the Kiali server is bound. The console and API server are accessible on this host.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"audit_log": {
								Description:         "When true, allows additional audit logging on write operations.",
								MarkdownDescription: "When true, allows additional audit logging on write operations.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cors_allow_all": {
								Description:         "When true, allows the web console to send requests to other domains other than where the console came from. Typically used for development environments only.",
								MarkdownDescription: "When true, allows the web console to send requests to other domains other than where the console came from. Typically used for development environments only.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"web_root": {
								Description:         "Defines the context root path for the Kiali console and API endpoints and readiness probes. When providing a context root path that is not '/', do not add a trailing slash (i.e. use '/kiali' not '/kiali/'). When empty, this will default to '/' on OpenShift and '/kiali' on other Kubernetes environments.",
								MarkdownDescription: "Defines the context root path for the Kiali console and API endpoints and readiness probes. When providing a context root path that is not '/', do not add a trailing slash (i.e. use '/kiali' not '/kiali/'). When empty, this will default to '/' on OpenShift and '/kiali' on other Kubernetes environments.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"port": {
								Description:         "The port that the server will bind to in order to receive console and API requests.",
								MarkdownDescription: "The port that the server will bind to in order to receive console and API requests.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"web_fqdn": {
								Description:         "Defines the public domain where Kiali is being served. This is the 'domain' part of the URL (usually it's a fully-qualified domain name). For example, 'kiali.example.org'. When empty, Kiali will try to guess this value from HTTP headers. On non-OpenShift clusters, you must populate this value if you want to enable cross-linking between Kiali instances in a multi-cluster setup.",
								MarkdownDescription: "Defines the public domain where Kiali is being served. This is the 'domain' part of the URL (usually it's a fully-qualified domain name). For example, 'kiali.example.org'. When empty, Kiali will try to guess this value from HTTP headers. On non-OpenShift clusters, you must populate this value if you want to enable cross-linking between Kiali instances in a multi-cluster setup.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"web_port": {
								Description:         "Defines the ingress port where the connections come from. This is usually necessary when the application responds through a proxy/ingress, and it does not forward the correct headers (when this happens, Kiali cannot guess the port). When empty, Kiali will try to guess this value from HTTP headers.",
								MarkdownDescription: "Defines the ingress port where the connections come from. This is usually necessary when the application responds through a proxy/ingress, and it does not forward the correct headers (when this happens, Kiali cannot guess the port). When empty, Kiali will try to guess this value from HTTP headers.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"api": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"namespaces": {
								Description:         "Settings that control what namespaces are returned by Kiali.",
								MarkdownDescription: "Settings that control what namespaces are returned by Kiali.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exclude": {
										Description:         "A list of namespaces to be excluded from the list of namespaces provided by the Kiali API and Kiali UI. Regex is supported. This does not affect explicit namespace access.",
										MarkdownDescription: "A list of namespaces to be excluded from the list of namespaces provided by the Kiali API and Kiali UI. Regex is supported. This does not affect explicit namespace access.",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"label_selector": {
										Description:         "A Kubernetes label selector (e.g. 'myLabel=myValue') which is used when fetching the list ofavailable namespaces. This does not affect explicit namespace access.If 'deployment.accessible_namespaces' does not have the special value of ''**''then the Kiali operator will add a new label to all accessible namespaces - that newlabel will be this 'label_selector'.Note that if you do not set this 'label_selector' setting but 'deployment.accessible_namespaces'does not have the special 'all namespaces' entry of ''**'' then this 'label_selector' will be setto a default value of 'kiali.io/[<deployment.instance_name>.]member-of=<deployment.namespace>'where '[<deployment.instance_name>.]' is the instance name assigned to the Kiali installationif it is not the default 'kiali' (otherwise, this is omitted) and '<deployment.namespace>'is the namespace where Kiali is to be installed.",
										MarkdownDescription: "A Kubernetes label selector (e.g. 'myLabel=myValue') which is used when fetching the list ofavailable namespaces. This does not affect explicit namespace access.If 'deployment.accessible_namespaces' does not have the special value of ''**''then the Kiali operator will add a new label to all accessible namespaces - that newlabel will be this 'label_selector'.Note that if you do not set this 'label_selector' setting but 'deployment.accessible_namespaces'does not have the special 'all namespaces' entry of ''**'' then this 'label_selector' will be setto a default value of 'kiali.io/[<deployment.instance_name>.]member-of=<deployment.namespace>'where '[<deployment.instance_name>.]' is the instance name assigned to the Kiali installationif it is not the default 'kiali' (otherwise, this is omitted) and '<deployment.namespace>'is the namespace where Kiali is to be installed.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

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
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"instance_name": {
								Description:         "The instance name of this Kiali installation. This instance name will be the prefix prepended to the names of all Kiali resources created by the operator and will be used to label those resources as belonging to this Kiali installation instance. You cannot change this instance name after a Kiali CR is created. If you attempt to change it, the operator will abort with an error. If you want to change it, you must first delete the original Kiali CR and create a new one. Note that this does not affect the name of the auto-generated signing key secret. If you do not supply a signing key, the operator will create one for you in a secret, but that secret will always be named 'kiali-signing-key' and shared across all Kiali instances in the same deployment namespace. If you want a different signing key secret, you are free to create your own and tell the operator about it via 'login_token.signing_key'. See the docs on that setting for more details. Note also that if you are setting this value, you may also want to change the 'installation_tag' setting, but this is not required.",
								MarkdownDescription: "The instance name of this Kiali installation. This instance name will be the prefix prepended to the names of all Kiali resources created by the operator and will be used to label those resources as belonging to this Kiali installation instance. You cannot change this instance name after a Kiali CR is created. If you attempt to change it, the operator will abort with an error. If you want to change it, you must first delete the original Kiali CR and create a new one. Note that this does not affect the name of the auto-generated signing key secret. If you do not supply a signing key, the operator will create one for you in a secret, but that secret will always be named 'kiali-signing-key' and shared across all Kiali instances in the same deployment namespace. If you want a different signing key secret, you are free to create your own and tell the operator about it via 'login_token.signing_key'. See the docs on that setting for more details. Note also that if you are setting this value, you may also want to change the 'installation_tag' setting, but this is not required.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"priority_class_name": {
								Description:         "The priorityClassName used to assign the priority of the Kiali pod.",
								MarkdownDescription: "The priorityClassName used to assign the priority of the Kiali pod.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_type": {
								Description:         "The Kiali service type. Kubernetes determines what values are valid. Common values are 'NodePort', 'ClusterIP', and 'LoadBalancer'.",
								MarkdownDescription: "The Kiali service type. Kubernetes determines what values are valid. Common values are 'NodePort', 'ClusterIP', and 'LoadBalancer'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tolerations": {
								Description:         "A list of tolerations which declare which node taints Kiali can tolerate. See the Kubernetes documentation on Taints and Tolerations for more details.",
								MarkdownDescription: "A list of tolerations which declare which node taints Kiali can tolerate. See the Kubernetes documentation on Taints and Tolerations for more details.",

								Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replicas": {
								Description:         "The replica count for the Kiail deployment.",
								MarkdownDescription: "The replica count for the Kiail deployment.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"service_annotations": {
								Description:         "Custom annotations to be created on the Kiali Service resource.",
								MarkdownDescription: "Custom annotations to be created on the Kiali Service resource.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version_label": {
								Description:         "Kiali resources will be assigned a 'version' label when they are deployed.This setting determines what value those 'version' labels will have.When empty, its default will be determined as follows,* If 'deployment.image_version' is 'latest', 'version_label' will be fixed to 'master'.* If 'deployment.image_version' is 'lastrelease', 'version_label' will be fixed to the last Kiali release version string.* If 'deployment.image_version' is anything else, 'version_label' will be that value, too.",
								MarkdownDescription: "Kiali resources will be assigned a 'version' label when they are deployed.This setting determines what value those 'version' labels will have.When empty, its default will be determined as follows,* If 'deployment.image_version' is 'latest', 'version_label' will be fixed to 'master'.* If 'deployment.image_version' is 'lastrelease', 'version_label' will be fixed to the last Kiali release version string.* If 'deployment.image_version' is anything else, 'version_label' will be that value, too.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "The namespace into which Kiali is to be installed. If this is empty or not defined, the default will be the namespace where the Kiali CR is located.",
								MarkdownDescription: "The namespace into which Kiali is to be installed. If this is empty or not defined, the default will be the namespace where the Kiali CR is located.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Defines compute resources that are to be given to the Kiali pod's container. The value is a dict as defined by Kubernetes. See the Kubernetes documentation (https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container).If you set this to an empty dict ('{}') then no resources will be defined in the Deployment.If you do not set this at all, the default is,'''requests:  cpu: '10m'  memory: '64Mi'limits:  memory: '1Gi''''",
								MarkdownDescription: "Defines compute resources that are to be given to the Kiali pod's container. The value is a dict as defined by Kubernetes. See the Kubernetes documentation (https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container).If you set this to an empty dict ('{}') then no resources will be defined in the Deployment.If you do not set this at all, the default is,'''requests:  cpu: '10m'  memory: '64Mi'limits:  memory: '1Gi''''",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"security_context": {
								Description:         "Custom security context to be placed on the server container. The entire security context on the container will be the value of this setting if the operator is configured to allow it. Note that, as a security measure, a cluster admin may have configured the Kiali operator to not allow portions of this override setting - in this case you can specify additional security context settings but you cannot replace existing, default ones.",
								MarkdownDescription: "Custom security context to be placed on the server container. The entire security context on the container will be the value of this setting if the operator is configured to allow it. Note that, as a security measure, a cluster admin may have configured the Kiali operator to not allow portions of this override setting - in this case you can specify additional security context settings but you cannot replace existing, default ones.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"affinity": {
								Description:         "Affinity definitions that are to be used to define the nodes where the Kiali pod should be constrained. See the Kubernetes documentation on Assigning Pods to Nodes for the proper syntax for these three different affinity types.",
								MarkdownDescription: "Affinity definitions that are to be used to define the nodes where the Kiali pod should be constrained. See the Kubernetes documentation on Assigning Pods to Nodes for the proper syntax for these three different affinity types.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"pod_anti": {
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

							"custom_secrets": {
								Description:         "Defines additional secrets that are to be mounted in the Kiali pod.These are useful to contain certs that are used by Kiali to securely connect to third party systems(for example, see 'external_services.tracing.auth.ca_file').These secrets must be created by an external mechanism. Kiali will not generate these secrets; itis assumed these secrets are externally managed. You can define 0, 1, or more secrets.An example configuration is,'''custom_secrets:- name: mysecret  mount: /mysecret-path- name: my-other-secret  mount: /my-other-secret-location  optional: true'''",
								MarkdownDescription: "Defines additional secrets that are to be mounted in the Kiali pod.These are useful to contain certs that are used by Kiali to securely connect to third party systems(for example, see 'external_services.tracing.auth.ca_file').These secrets must be created by an external mechanism. Kiali will not generate these secrets; itis assumed these secrets are externally managed. You can define 0, 1, or more secrets.An example configuration is,'''custom_secrets:- name: mysecret  mount: /mysecret-path- name: my-other-secret  mount: /my-other-secret-location  optional: true'''",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"mount": {
										Description:         "The file path location where the secret content will be mounted. The custom secret cannot be mounted on a path that the operator will use to mount its secrets. Make sure you set your custom secret mount path to a unique, unused path. Paths such as '/kiali-configuration', '/kiali-cert', '/kiali-cabundle', and '/kiali-secret' should not be used as mount paths for custom secrets because the operator may want to use one of those paths.",
										MarkdownDescription: "The file path location where the secret content will be mounted. The custom secret cannot be mounted on a path that the operator will use to mount its secrets. Make sure you set your custom secret mount path to a unique, unused path. Paths such as '/kiali-configuration', '/kiali-cert', '/kiali-cabundle', and '/kiali-secret' should not be used as mount paths for custom secrets because the operator may want to use one of those paths.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"name": {
										Description:         "The name of the secret that is to be mounted to the Kiali pod's file system. The name of the custom secret must not be the same name as one created by the operator. Names such as 'kiali', 'kiali-cert-secret', and 'kiali-cabundle' should not be used as a custom secret name because the operator may want to create one with one of those names.",
										MarkdownDescription: "The name of the secret that is to be mounted to the Kiali pod's file system. The name of the custom secret must not be the same name as one created by the operator. Names such as 'kiali', 'kiali-cert-secret', and 'kiali-cabundle' should not be used as a custom secret name because the operator may want to create one with one of those names.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"optional": {
										Description:         "Indicates if the secret may or may not exist at the time the Kiali pod starts. This will default to 'false' if not specified.",
										MarkdownDescription: "Indicates if the secret may or may not exist at the time the Kiali pod starts. This will default to 'false' if not specified.",

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

							"image_name": {
								Description:         "Determines which Kiali image to download and install. If you set this to a specific name (i.e. you do not leave it as the default empty string), you must make sure that image is supported by the operator. If empty, the operator will use a known supported image name based on which 'version' was defined. Note that, as a security measure, a cluster admin may have configured the Kiali operator to ignore this setting. A cluster admin may do this to ensure the Kiali operator only installs a single, specific Kiali version, thus this setting may have no effect depending on how the operator itself was configured.",
								MarkdownDescription: "Determines which Kiali image to download and install. If you set this to a specific name (i.e. you do not leave it as the default empty string), you must make sure that image is supported by the operator. If empty, the operator will use a known supported image name based on which 'version' was defined. Note that, as a security measure, a cluster admin may have configured the Kiali operator to ignore this setting. A cluster admin may do this to ensure the Kiali operator only installs a single, specific Kiali version, thus this setting may have no effect depending on how the operator itself was configured.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_policy": {
								Description:         "The Kubernetes pull policy for the Kiali deployment. This is overridden to be 'Always' if 'deployment.image_version' is set to 'latest'.",
								MarkdownDescription: "The Kubernetes pull policy for the Kiali deployment. This is overridden to be 'Always' if 'deployment.image_version' is set to 'latest'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"configmap_annotations": {
								Description:         "Custom annotations to be created on the Kiali ConfigMap.",
								MarkdownDescription: "Custom annotations to be created on the Kiali ConfigMap.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_aliases": {
								Description:         "This is content for the Kubernetes 'hostAliases' setting for the Kiali server.This allows you to modify the Kiali server pod '/etc/hosts' file.A typical way to configure this setting is,'''host_aliases:- ip: 192.168.1.100  hostnames:  - 'foo.local'  - 'bar.local''''For details on the content of this setting, see https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/#adding-additional-entries-with-hostaliases",
								MarkdownDescription: "This is content for the Kubernetes 'hostAliases' setting for the Kiali server.This allows you to modify the Kiali server pod '/etc/hosts' file.A typical way to configure this setting is,'''host_aliases:- ip: 192.168.1.100  hostnames:  - 'foo.local'  - 'bar.local''''For details on the content of this setting, see https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/#adding-additional-entries-with-hostaliases",

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

							"secret_name": {
								Description:         "The name of a secret used by the Kiali. This secret is optionally used when configuring the OpenID authentication strategy. Consult the OpenID docs for more information at https://kiali.io/docs/configuration/authentication/openid/",
								MarkdownDescription: "The name of a secret used by the Kiali. This secret is optionally used when configuring the OpenID authentication strategy. Consult the OpenID docs for more information at https://kiali.io/docs/configuration/authentication/openid/",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_digest": {
								Description:         "If 'deployment.image_version' is a digest hash, this value indicates what type of digest it is. A typical value would be 'sha256'. Note: do NOT prefix this value with a '@'.",
								MarkdownDescription: "If 'deployment.image_version' is a digest hash, this value indicates what type of digest it is. A typical value would be 'sha256'. Note: do NOT prefix this value with a '@'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_version": {
								Description:         "Determines which version of Kiali to install.Choose 'lastrelease' to use the last Kiali release.Choose 'latest' to use the latest image (which may or may not be a released version of Kiali).Choose 'operator_version' to use the image whose version is the same as the operator version.Otherwise, you can set this to any valid Kiali version (such as 'v1.0') or any valid Kialidigest hash (if you set this to a digest hash, you must indicate the digest in 'deployment.image_digest').Note that if this is set to 'latest' then the 'deployment.image_pull_policy' will be set to 'Always'.If you set this to a specific version (i.e. you do not leave it as the default empty string),you must make sure that image is supported by the operator.If empty, the operator will use a known supported image version based on which 'version' was defined.Note that, as a security measure, a cluster admin may have configured the Kiali operator toignore this setting. A cluster admin may do this to ensure the Kiali operator only installsa single, specific Kiali version, thus this setting may have no effect depending on how theoperator itself was configured.",
								MarkdownDescription: "Determines which version of Kiali to install.Choose 'lastrelease' to use the last Kiali release.Choose 'latest' to use the latest image (which may or may not be a released version of Kiali).Choose 'operator_version' to use the image whose version is the same as the operator version.Otherwise, you can set this to any valid Kiali version (such as 'v1.0') or any valid Kialidigest hash (if you set this to a digest hash, you must indicate the digest in 'deployment.image_digest').Note that if this is set to 'latest' then the 'deployment.image_pull_policy' will be set to 'Always'.If you set this to a specific version (i.e. you do not leave it as the default empty string),you must make sure that image is supported by the operator.If empty, the operator will use a known supported image version based on which 'version' was defined.Note that, as a security measure, a cluster admin may have configured the Kiali operator toignore this setting. A cluster admin may do this to ensure the Kiali operator only installsa single, specific Kiali version, thus this setting may have no effect depending on how theoperator itself was configured.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logger": {
								Description:         "Configures the logger that emits messages to the Kiali server pod logs.",
								MarkdownDescription: "Configures the logger that emits messages to the Kiali server pod logs.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"time_field_format": {
										Description:         "The log message timestamp format. This supports a golang time format (see https://golang.org/pkg/time/)",
										MarkdownDescription: "The log message timestamp format. This supports a golang time format (see https://golang.org/pkg/time/)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_format": {
										Description:         "Indicates if the logs should be written with one log message per line or using a JSON format. Must be one of: 'text' or 'json'.",
										MarkdownDescription: "Indicates if the logs should be written with one log message per line or using a JSON format. Must be one of: 'text' or 'json'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"log_level": {
										Description:         "The lowest priority of messages to log. Must be one of: 'trace', 'debug', 'info', 'warn', 'error', or 'fatal'.",
										MarkdownDescription: "The lowest priority of messages to log. Must be one of: 'trace', 'debug', 'info', 'warn', 'error', or 'fatal'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"sampler_rate": {
										Description:         "With this setting every sampler_rate-th message will be logged. By default, every message is logged. As an example, setting this to ''2'' means every other message will be logged. The value of this setting is a string but must be parsable as an integer.",
										MarkdownDescription: "With this setting every sampler_rate-th message will be logged. By default, every message is logged. As an example, setting this to ''2'' means every other message will be logged. The value of this setting is a string but must be parsable as an integer.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_annotations": {
								Description:         "Custom annotations to be created on the Kiali pod.",
								MarkdownDescription: "Custom annotations to be created on the Kiali pod.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_secrets": {
								Description:         "The names of the secrets to be used when container images are to be pulled.",
								MarkdownDescription: "The names of the secrets to be used when container images are to be pulled.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ingress": {
								Description:         "Configures if/how the Kiali endpoint should be exposed externally.",
								MarkdownDescription: "Configures if/how the Kiali endpoint should be exposed externally.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"additional_labels": {
										Description:         "Additional labels to add to the Ingress (or Route if on OpenShift). These are added to the labels that are created by default; these do not override the default labels.",
										MarkdownDescription: "Additional labels to add to the Ingress (or Route if on OpenShift). These are added to the labels that are created by default; these do not override the default labels.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"class_name": {
										Description:         "If 'class_name' is a non-empty string, it will be used as the 'spec.ingressClassName' in the created Kubernetes Ingress resource. This setting is ignored if on OpenShift. This is also ignored if 'override_yaml.spec' is defined (i.e. you must define the 'ingressClassName' directly in your override yaml).",
										MarkdownDescription: "If 'class_name' is a non-empty string, it will be used as the 'spec.ingressClassName' in the created Kubernetes Ingress resource. This setting is ignored if on OpenShift. This is also ignored if 'override_yaml.spec' is defined (i.e. you must define the 'ingressClassName' directly in your override yaml).",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"enabled": {
										Description:         "Determines if the Kiali endpoint should be exposed externally. If 'true', an Ingress will be created if on Kubernetes or a Route if on OpenShift. If left undefined, this will be 'false' on Kubernetes and 'true' on OpenShift.",
										MarkdownDescription: "Determines if the Kiali endpoint should be exposed externally. If 'true', an Ingress will be created if on Kubernetes or a Route if on OpenShift. If left undefined, this will be 'false' on Kubernetes and 'true' on OpenShift.",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"override_yaml": {
										Description:         "Because an Ingress into a cluster can vary wildly in its desired configuration,this setting provides a way to override complete portions of the Ingress resourceconfiguration (Ingress on Kubernetes and Route on OpenShift). It is up to the userto ensure this override YAML configuration is valid and supports the cluster environmentsince the operator will blindly copy this custom configuration into the resource itcreates.This setting is not used if 'deployment.ingress.enabled' is set to 'false'.Note that only 'metadata.annotations' and 'spec' is valid and only they willbe used to override those same sections in the created resource. You can defineeither one or both.Note that 'override_yaml.metadata.labels' is not allowed - you cannot override the labels; to addlabels to the default set of labels, use the 'deployment.ingress.additional_labels' setting.Example,'''override_yaml:  metadata:    annotations:      nginx.ingress.kubernetes.io/secure-backends: 'true'      nginx.ingress.kubernetes.io/backend-protocol: 'HTTPS'  spec:    rules:    - http:        paths:        - path: /kiali          pathType: Prefix          backend:            service              name: 'kiali'              port:                number: 20001'''",
										MarkdownDescription: "Because an Ingress into a cluster can vary wildly in its desired configuration,this setting provides a way to override complete portions of the Ingress resourceconfiguration (Ingress on Kubernetes and Route on OpenShift). It is up to the userto ensure this override YAML configuration is valid and supports the cluster environmentsince the operator will blindly copy this custom configuration into the resource itcreates.This setting is not used if 'deployment.ingress.enabled' is set to 'false'.Note that only 'metadata.annotations' and 'spec' is valid and only they willbe used to override those same sections in the created resource. You can defineeither one or both.Note that 'override_yaml.metadata.labels' is not allowed - you cannot override the labels; to addlabels to the default set of labels, use the 'deployment.ingress.additional_labels' setting.Example,'''override_yaml:  metadata:    annotations:      nginx.ingress.kubernetes.io/secure-backends: 'true'      nginx.ingress.kubernetes.io/backend-protocol: 'HTTPS'  spec:    rules:    - http:        paths:        - path: /kiali          pathType: Prefix          backend:            service              name: 'kiali'              port:                number: 20001'''",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"metadata": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"annotations": {
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

											"spec": {
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

							"pod_labels": {
								Description:         "Custom labels to be created on the Kiali pod.An example use for this setting is to inject an Istio sidecar such as,'''sidecar.istio.io/inject: 'true''''",
								MarkdownDescription: "Custom labels to be created on the Kiali pod.An example use for this setting is to inject an Istio sidecar such as,'''sidecar.istio.io/inject: 'true''''",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"verbose_mode": {
								Description:         "DEPRECATED! Determines which priority levels of log messages Kiali will output. Use 'deployment.logger' settings instead.",
								MarkdownDescription: "DEPRECATED! Determines which priority levels of log messages Kiali will output. Use 'deployment.logger' settings instead.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"view_only_mode": {
								Description:         "When true, Kiali will be in 'view only' mode, allowing the user to view and retrieve management and monitoring data for the service mesh, but not allow the user to modify the service mesh.",
								MarkdownDescription: "When true, Kiali will be in 'view only' mode, allowing the user to view and retrieve management and monitoring data for the service mesh, but not allow the user to modify the service mesh.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"accessible_namespaces": {
								Description:         "A list of namespaces Kiali is to be given access to. These namespaces have service mesh components that are to be observed by Kiali. You can provide names using regex expressions matched against all namespaces the operator can see. The default makes all namespaces accessible except for some internal namespaces that typically should be ignored. NOTE! If this has an entry with the special value of ''**'' (two asterisks), that will denote you want Kiali to be given access to all namespaces via a single cluster role (if using this special value of ''**'', you are required to have already granted the operator permissions to create cluster roles and cluster role bindings).",
								MarkdownDescription: "A list of namespaces Kiali is to be given access to. These namespaces have service mesh components that are to be observed by Kiali. You can provide names using regex expressions matched against all namespaces the operator can see. The default makes all namespaces accessible except for some internal namespaces that typically should be ignored. NOTE! If this has an entry with the special value of ''**'' (two asterisks), that will denote you want Kiali to be given access to all namespaces via a single cluster role (if using this special value of ''**'', you are required to have already granted the operator permissions to create cluster roles and cluster role bindings).",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"additional_service_yaml": {
								Description:         "Additional custom yaml to add to the service definition. This is used mainly to customize the service type. For example, if the 'deployment.service_type' is set to 'LoadBalancer' and you want to set the loadBalancerIP, you can do so here with: 'additional_service_yaml: { 'loadBalancerIP': '78.11.24.19' }'. Another example would be if the 'deployment.service_type' is set to 'ExternalName' you will need to configure the name via: 'additional_service_yaml: { 'externalName': 'my.kiali.example.com' }'. A final example would be if external IPs need to be set: 'additional_service_yaml: { 'externalIPs': ['80.11.12.10'] }'",
								MarkdownDescription: "Additional custom yaml to add to the service definition. This is used mainly to customize the service type. For example, if the 'deployment.service_type' is set to 'LoadBalancer' and you want to set the loadBalancerIP, you can do so here with: 'additional_service_yaml: { 'loadBalancerIP': '78.11.24.19' }'. Another example would be if the 'deployment.service_type' is set to 'ExternalName' you will need to configure the name via: 'additional_service_yaml: { 'externalName': 'my.kiali.example.com' }'. A final example would be if external IPs need to be set: 'additional_service_yaml: { 'externalIPs': ['80.11.12.10'] }'",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"hpa": {
								Description:         "Determines what (if any) HorizontalPodAutoscaler should be created to autoscale the Kiali pod.A typical way to configure HPA for Kiali is,'''hpa:  api_version: 'autoscaling/v2'  spec:    maxReplicas: 2    minReplicas: 1    metrics:    - type: Resource      resource:        name: cpu        target:          type: Utilization          averageUtilization: 50'''",
								MarkdownDescription: "Determines what (if any) HorizontalPodAutoscaler should be created to autoscale the Kiali pod.A typical way to configure HPA for Kiali is,'''hpa:  api_version: 'autoscaling/v2'  spec:    maxReplicas: 2    minReplicas: 1    metrics:    - type: Resource      resource:        name: cpu        target:          type: Utilization          averageUtilization: 50'''",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"spec": {
										Description:         "The 'spec' specified here will be placed in the created HPA resource's 'spec' section. If 'spec' is left empty, no HPA resource will be created. Note that you must not specify the 'scaleTargetRef' section in 'spec'; the Kiali Operator will populate that for you.",
										MarkdownDescription: "The 'spec' specified here will be placed in the created HPA resource's 'spec' section. If 'spec' is left empty, no HPA resource will be created. Note that you must not specify the 'scaleTargetRef' section in 'spec'; the Kiali Operator will populate that for you.",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"api_version": {
										Description:         "A specific HPA API version that can be specified in case there is some HPA feature you want to use that is only supported in that specific version. If value is an empty string, an attempt will be made to determine a valid version.",
										MarkdownDescription: "A specific HPA API version that can be specified in case there is some HPA feature you want to use that is only supported in that specific version. If value is an empty string, an attempt will be made to determine a valid version.",

										Type: types.StringType,

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
								Description:         "A set of node labels that dictate onto which node the Kiali pod will be deployed.",
								MarkdownDescription: "A set of node labels that dictate onto which node the Kiali pod will be deployed.",

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
						Description:         "The version of the Ansible playbook to execute in order to install that version of Kiali.It is rare you will want to set this - if you are thinking of setting this, know what you are doing first.The only supported value today is 'default'.If not specified, a default version of Kiali will be installed which will be the most recent release of Kiali.Refer to this file to see where these values are defined in the master branch,https://github.com/kiali/kiali-operator/tree/master/playbooks/default-supported-images.ymlThis version setting affects the defaults of the deployment.image_name anddeployment.image_version settings. See the comments for those settingsbelow for additional details. But in short, this version setting willdictate which version of the Kiali image will be deployed by default.Note that if you explicitly set deployment.image_name and/ordeployment.image_version you are responsible for ensuring those settingsare compatible with this setting (i.e. the Kiali image must be compatiblewith the rest of the configuration and resources the operator will install).",
						MarkdownDescription: "The version of the Ansible playbook to execute in order to install that version of Kiali.It is rare you will want to set this - if you are thinking of setting this, know what you are doing first.The only supported value today is 'default'.If not specified, a default version of Kiali will be installed which will be the most recent release of Kiali.Refer to this file to see where these values are defined in the master branch,https://github.com/kiali/kiali-operator/tree/master/playbooks/default-supported-images.ymlThis version setting affects the defaults of the deployment.image_name anddeployment.image_version settings. See the comments for those settingsbelow for additional details. But in short, this version setting willdictate which version of the Kiali image will be deployed by default.Note that if you explicitly set deployment.image_name and/ordeployment.image_version you are responsible for ensuring those settingsare compatible with this setting (i.e. the Kiali image must be compatiblewith the rest of the configuration and resources the operator will install).",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"istio_labels": {
						Description:         "Defines specific labels used by Istio that Kiali needs to know about.",
						MarkdownDescription: "Defines specific labels used by Istio that Kiali needs to know about.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"injection_label_rev": {
								Description:         "The label used to identify the Istio revision.",
								MarkdownDescription: "The label used to identify the Istio revision.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"version_label_name": {
								Description:         "The name of the label used to define what version of the application a workload belongs to. This is typically something like 'version' or 'app.kubernetes.io/version'.",
								MarkdownDescription: "The name of the label used to define what version of the application a workload belongs to. This is typically something like 'version' or 'app.kubernetes.io/version'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"app_label_name": {
								Description:         "The name of the label used to define what application a workload belongs to. This is typically something like 'app' or 'app.kubernetes.io/name'.",
								MarkdownDescription: "The name of the label used to define what application a workload belongs to. This is typically something like 'app' or 'app.kubernetes.io/name'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"injection_label_name": {
								Description:         "The name of the label used to instruct Istio to automatically inject sidecar proxies when applications are deployed.",
								MarkdownDescription: "The name of the label used to instruct Istio to automatically inject sidecar proxies when applications are deployed.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kubernetes_config": {
						Description:         "Configuration of Kiali's access of the Kubernetes API.",
						MarkdownDescription: "Configuration of Kiali's access of the Kubernetes API.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"qps": {
								Description:         "The QPS value of the Kubernetes client.",
								MarkdownDescription: "The QPS value of the Kubernetes client.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"burst": {
								Description:         "The Burst value of the Kubernetes client.",
								MarkdownDescription: "The Burst value of the Kubernetes client.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cache_duration": {
								Description:         "The ratio interval (expressed in seconds) used for the cache to perform a full refresh. Only used when 'cache_enabled' is true.",
								MarkdownDescription: "The ratio interval (expressed in seconds) used for the cache to perform a full refresh. Only used when 'cache_enabled' is true.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cache_enabled": {
								Description:         "Flag to use a Kubernetes cache for watching changes and updating pods and controllers data asynchronously.",
								MarkdownDescription: "Flag to use a Kubernetes cache for watching changes and updating pods and controllers data asynchronously.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cache_istio_types": {
								Description:         "Kiali can cache VirtualService, DestinationRule, Gateway and ServiceEntry Istio resources if they are present on this list of Istio types. Other Istio types are not yet supported.",
								MarkdownDescription: "Kiali can cache VirtualService, DestinationRule, Gateway and ServiceEntry Istio resources if they are present on this list of Istio types. Other Istio types are not yet supported.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cache_namespaces": {
								Description:         "List of namespaces or regex defining namespaces to include in a cache.",
								MarkdownDescription: "List of namespaces or regex defining namespaces to include in a cache.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cache_token_namespace_duration": {
								Description:         "This Kiali cache is a list of namespaces per user. This is typically a short-lived cache compared with the duration of the namespace cache defined by the 'cache_duration' setting. This is specified in seconds.",
								MarkdownDescription: "This Kiali cache is a list of namespaces per user. This is typically a short-lived cache compared with the duration of the namespace cache defined by the 'cache_duration' setting. This is specified in seconds.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"excluded_workloads": {
								Description:         "List of controllers that won't be used for Workload calculation. Kiali queries Deployment, ReplicaSet, ReplicationController, DeploymentConfig, StatefulSet, Job and CronJob controllers. Deployment and ReplicaSet will be always queried, but ReplicationController, DeploymentConfig, StatefulSet, Job and CronJobs can be skipped from Kiali workloads queries if they are present in this list.",
								MarkdownDescription: "List of controllers that won't be used for Workload calculation. Kiali queries Deployment, ReplicaSet, ReplicationController, DeploymentConfig, StatefulSet, Job and CronJob controllers. Deployment and ReplicaSet will be always queried, but ReplicationController, DeploymentConfig, StatefulSet, Job and CronJobs can be skipped from Kiali workloads queries if they are present in this list.",

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

					"identity": {
						Description:         "Settings that define the Kiali server identity.",
						MarkdownDescription: "Settings that define the Kiali server identity.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cert_file": {
								Description:         "Certificate file used to identify the Kiali server. If set, you must go over https to access Kiali. The Kiali operator will set this if it deploys Kiali behind https. When left undefined, the operator will attempt to generate a cluster-specific cert file that provides https by default (today, this auto-generation of a cluster-specific cert is only supported on OpenShift). When set to an empty string, https will be disabled.",
								MarkdownDescription: "Certificate file used to identify the Kiali server. If set, you must go over https to access Kiali. The Kiali operator will set this if it deploys Kiali behind https. When left undefined, the operator will attempt to generate a cluster-specific cert file that provides https by default (today, this auto-generation of a cluster-specific cert is only supported on OpenShift). When set to an empty string, https will be disabled.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"private_key_file": {
								Description:         "Private key file used to identify the Kiali server. If set, you must go over https to access Kiali. When left undefined, the Kiali operator will attempt to generate a cluster-specific private key file that provides https by default (today, this auto-generation of a cluster-specific private key is only supported on OpenShift). When set to an empty string, https will be disabled.",
								MarkdownDescription: "Private key file used to identify the Kiali server. If set, you must go over https to access Kiali. When left undefined, the Kiali operator will attempt to generate a cluster-specific private key file that provides https by default (today, this auto-generation of a cluster-specific private key is only supported on OpenShift). When set to an empty string, https will be disabled.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"installation_tag": {
						Description:         "Tag used to identify a particular instance/installation of the Kiali server. This is merely a human-readable string that will be used within Kiali to help a user identify the Kiali being used (e.g. in the Kiali UI title bar). See 'deployment.instance_name' for the setting used to customize Kiali resource names that are created.",
						MarkdownDescription: "Tag used to identify a particular instance/installation of the Kiali server. This is merely a human-readable string that will be used within Kiali to help a user identify the Kiali being used (e.g. in the Kiali UI title bar). See 'deployment.instance_name' for the setting used to customize Kiali resource names that are created.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"auth": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"openshift": {
								Description:         "To learn more about these settings and how to configure the OpenShift authentication strategy, read the documentation at https://kiali.io/docs/configuration/authentication/openshift/",
								MarkdownDescription: "To learn more about these settings and how to configure the OpenShift authentication strategy, read the documentation at https://kiali.io/docs/configuration/authentication/openshift/",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"client_id_prefix": {
										Description:         "The Route resource name and OAuthClient resource name will have this value as its prefix. This value normally should never change. The installer will ensure this value is set correctly.",
										MarkdownDescription: "The Route resource name and OAuthClient resource name will have this value as its prefix. This value normally should never change. The installer will ensure this value is set correctly.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"strategy": {
								Description:         "Determines what authentication strategy to use when users log into Kiali.Options are 'anonymous', 'token', 'openshift', 'openid', or 'header'.* Choose 'anonymous' to allow full access to Kiali without requiring any credentials.* Choose 'token' to allow access to Kiali using service account tokens, which controlsaccess based on RBAC roles assigned to the service account.* Choose 'openshift' to use the OpenShift OAuth login which controls access based onthe individual's RBAC roles in OpenShift. Not valid for non-OpenShift environments.* Choose 'openid' to enable OpenID Connect-based authentication. Your cluster is required tobe configured to accept the tokens issued by your IdP. There are additional requiredconfigurations for this strategy. See below for the additional OpenID configuration section.* Choose 'header' when Kiali is running behind a reverse proxy that will inject anAuthorization header and potentially impersonation headers.When empty, this value will default to 'openshift' on OpenShift and 'token' on other Kubernetes environments.",
								MarkdownDescription: "Determines what authentication strategy to use when users log into Kiali.Options are 'anonymous', 'token', 'openshift', 'openid', or 'header'.* Choose 'anonymous' to allow full access to Kiali without requiring any credentials.* Choose 'token' to allow access to Kiali using service account tokens, which controlsaccess based on RBAC roles assigned to the service account.* Choose 'openshift' to use the OpenShift OAuth login which controls access based onthe individual's RBAC roles in OpenShift. Not valid for non-OpenShift environments.* Choose 'openid' to enable OpenID Connect-based authentication. Your cluster is required tobe configured to accept the tokens issued by your IdP. There are additional requiredconfigurations for this strategy. See below for the additional OpenID configuration section.* Choose 'header' when Kiali is running behind a reverse proxy that will inject anAuthorization header and potentially impersonation headers.When empty, this value will default to 'openshift' on OpenShift and 'token' on other Kubernetes environments.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"openid": {
								Description:         "To learn more about these settings and how to configure the OpenId authentication strategy, read the documentation at https://kiali.io/docs/configuration/authentication/openid/",
								MarkdownDescription: "To learn more about these settings and how to configure the OpenId authentication strategy, read the documentation at https://kiali.io/docs/configuration/authentication/openid/",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"api_proxy": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"http_proxy": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"https_proxy": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"insecure_skip_verify_tls": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"username_claim": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"additional_request_params": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authorization_endpoint": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"client_id": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"allowed_domains": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"api_proxy_ca_data": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"api_token": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"disable_rbac": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scopes": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.ListType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"authentication_timeout": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"issuer_uri": {
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

					"health_config": {
						Description:         "This section defines what it means for nodes to be healthy. For more details, see https://kiali.io/docs/configuration/health/",
						MarkdownDescription: "This section defines what it means for nodes to be healthy. For more details, see https://kiali.io/docs/configuration/health/",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"rate": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"namespace": {
										Description:         "The name of the namespace that this configuration applies to. This is a regular expression.",
										MarkdownDescription: "The name of the namespace that this configuration applies to. This is a regular expression.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerance": {
										Description:         "A list of tolerances for this configuration.",
										MarkdownDescription: "A list of tolerances for this configuration.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"code": {
												Description:         "The status code that applies for this tolerance. This is a regular expression.",
												MarkdownDescription: "The status code that applies for this tolerance. This is a regular expression.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"degraded": {
												Description:         "Health will be considered degraded when the telemetry reaches this value (specified as an integer representing a percentage).",
												MarkdownDescription: "Health will be considered degraded when the telemetry reaches this value (specified as an integer representing a percentage).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"direction": {
												Description:         "The direction that applies for this tolerance (e.g. inbound or outbound). This is a regular expression.",
												MarkdownDescription: "The direction that applies for this tolerance (e.g. inbound or outbound). This is a regular expression.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"failure": {
												Description:         "A failure status will be shown when the telemetry reaches this value (specified as an integer representing a percentage).",
												MarkdownDescription: "A failure status will be shown when the telemetry reaches this value (specified as an integer representing a percentage).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"protocol": {
												Description:         "The protocol that applies for this tolerance (e.g. grpc or http). This is a regular expression.",
												MarkdownDescription: "The protocol that applies for this tolerance (e.g. grpc or http). This is a regular expression.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "The type of resource that this configuration applies to. This is a regular expression.",
										MarkdownDescription: "The type of resource that this configuration applies to. This is a regular expression.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "The name of a resource that this configuration applies to. This is a regular expression.",
										MarkdownDescription: "The name of a resource that this configuration applies to. This is a regular expression.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},
						}),

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

func (r *KialiIoKialiV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_kiali_io_kiali_v1alpha1")

	var state KialiIoKialiV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KialiIoKialiV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kiali.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Kiali")

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

func (r *KialiIoKialiV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kiali_io_kiali_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *KialiIoKialiV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_kiali_io_kiali_v1alpha1")

	var state KialiIoKialiV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel KialiIoKialiV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("kiali.io/v1alpha1")
	goModel.Kind = utilities.Ptr("Kiali")

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

func (r *KialiIoKialiV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_kiali_io_kiali_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
