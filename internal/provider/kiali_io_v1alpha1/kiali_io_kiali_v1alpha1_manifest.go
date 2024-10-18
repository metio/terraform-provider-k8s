/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package kiali_io_v1alpha1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &KialiIoKialiV1Alpha1Manifest{}
)

func NewKialiIoKialiV1Alpha1Manifest() datasource.DataSource {
	return &KialiIoKialiV1Alpha1Manifest{}
}

type KialiIoKialiV1Alpha1Manifest struct{}

type KialiIoKialiV1Alpha1ManifestData struct {
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
		Additional_display_details *[]struct {
			Annotation      *string `tfsdk:"annotation" json:"annotation,omitempty"`
			Icon_annotation *string `tfsdk:"icon_annotation" json:"icon_annotation,omitempty"`
			Title           *string `tfsdk:"title" json:"title,omitempty"`
		} `tfsdk:"additional_display_details" json:"additional_display_details,omitempty"`
		Auth *struct {
			Openid *struct {
				Additional_request_params *map[string]string `tfsdk:"additional_request_params" json:"additional_request_params,omitempty"`
				Allowed_domains           *[]string          `tfsdk:"allowed_domains" json:"allowed_domains,omitempty"`
				Api_proxy                 *string            `tfsdk:"api_proxy" json:"api_proxy,omitempty"`
				Api_proxy_ca_data         *string            `tfsdk:"api_proxy_ca_data" json:"api_proxy_ca_data,omitempty"`
				Api_token                 *string            `tfsdk:"api_token" json:"api_token,omitempty"`
				Authentication_timeout    *int64             `tfsdk:"authentication_timeout" json:"authentication_timeout,omitempty"`
				Authorization_endpoint    *string            `tfsdk:"authorization_endpoint" json:"authorization_endpoint,omitempty"`
				Client_id                 *string            `tfsdk:"client_id" json:"client_id,omitempty"`
				Disable_rbac              *bool              `tfsdk:"disable_rbac" json:"disable_rbac,omitempty"`
				Http_proxy                *string            `tfsdk:"http_proxy" json:"http_proxy,omitempty"`
				Https_proxy               *string            `tfsdk:"https_proxy" json:"https_proxy,omitempty"`
				Insecure_skip_verify_tls  *bool              `tfsdk:"insecure_skip_verify_tls" json:"insecure_skip_verify_tls,omitempty"`
				Issuer_uri                *string            `tfsdk:"issuer_uri" json:"issuer_uri,omitempty"`
				Scopes                    *[]string          `tfsdk:"scopes" json:"scopes,omitempty"`
				Username_claim            *string            `tfsdk:"username_claim" json:"username_claim,omitempty"`
			} `tfsdk:"openid" json:"openid,omitempty"`
			Openshift *struct {
				Token_inactivity_timeout *int64 `tfsdk:"token_inactivity_timeout" json:"token_inactivity_timeout,omitempty"`
				Token_max_age            *int64 `tfsdk:"token_max_age" json:"token_max_age,omitempty"`
			} `tfsdk:"openshift" json:"openshift,omitempty"`
			Strategy *string `tfsdk:"strategy" json:"strategy,omitempty"`
		} `tfsdk:"auth" json:"auth,omitempty"`
		Clustering *struct {
			Autodetect_secrets *struct {
				Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Label   *string `tfsdk:"label" json:"label,omitempty"`
			} `tfsdk:"autodetect_secrets" json:"autodetect_secrets,omitempty"`
			Clusters *[]struct {
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				Secret_name *string `tfsdk:"secret_name" json:"secret_name,omitempty"`
			} `tfsdk:"clusters" json:"clusters,omitempty"`
			Kiali_urls *[]struct {
				Cluster_name  *string `tfsdk:"cluster_name" json:"cluster_name,omitempty"`
				Instance_name *string `tfsdk:"instance_name" json:"instance_name,omitempty"`
				Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Url           *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"kiali_urls" json:"kiali_urls,omitempty"`
		} `tfsdk:"clustering" json:"clustering,omitempty"`
		Custom_dashboards *[]map[string]string `tfsdk:"custom_dashboards" json:"custom_dashboards,omitempty"`
		Deployment        *struct {
			Additional_service_yaml *map[string]string `tfsdk:"additional_service_yaml" json:"additional_service_yaml,omitempty"`
			Affinity                *struct {
				Node     *map[string]string `tfsdk:"node" json:"node,omitempty"`
				Pod      *map[string]string `tfsdk:"pod" json:"pod,omitempty"`
				Pod_anti *map[string]string `tfsdk:"pod_anti" json:"pod_anti,omitempty"`
			} `tfsdk:"affinity" json:"affinity,omitempty"`
			Cluster_wide_access   *bool              `tfsdk:"cluster_wide_access" json:"cluster_wide_access,omitempty"`
			Configmap_annotations *map[string]string `tfsdk:"configmap_annotations" json:"configmap_annotations,omitempty"`
			Custom_envs           *[]struct {
				Name  *string `tfsdk:"name" json:"name,omitempty"`
				Value *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"custom_envs" json:"custom_envs,omitempty"`
			Custom_secrets *[]struct {
				Csi      *map[string]string `tfsdk:"csi" json:"csi,omitempty"`
				Mount    *string            `tfsdk:"mount" json:"mount,omitempty"`
				Name     *string            `tfsdk:"name" json:"name,omitempty"`
				Optional *bool              `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"custom_secrets" json:"custom_secrets,omitempty"`
			Discovery_selectors *struct {
				Default *[]struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"default" json:"default,omitempty"`
				Overrides *map[string]string `tfsdk:"overrides" json:"overrides,omitempty"`
			} `tfsdk:"discovery_selectors" json:"discovery_selectors,omitempty"`
			Dns *struct {
				Config *map[string]string `tfsdk:"config" json:"config,omitempty"`
				Policy *string            `tfsdk:"policy" json:"policy,omitempty"`
			} `tfsdk:"dns" json:"dns,omitempty"`
			Host_aliases *[]struct {
				Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
				Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
			} `tfsdk:"host_aliases" json:"host_aliases,omitempty"`
			Hpa *struct {
				Api_version *string            `tfsdk:"api_version" json:"api_version,omitempty"`
				Spec        *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"hpa" json:"hpa,omitempty"`
			Image_digest       *string   `tfsdk:"image_digest" json:"image_digest,omitempty"`
			Image_name         *string   `tfsdk:"image_name" json:"image_name,omitempty"`
			Image_pull_policy  *string   `tfsdk:"image_pull_policy" json:"image_pull_policy,omitempty"`
			Image_pull_secrets *[]string `tfsdk:"image_pull_secrets" json:"image_pull_secrets,omitempty"`
			Image_version      *string   `tfsdk:"image_version" json:"image_version,omitempty"`
			Ingress            *struct {
				Additional_labels *map[string]string `tfsdk:"additional_labels" json:"additional_labels,omitempty"`
				Class_name        *string            `tfsdk:"class_name" json:"class_name,omitempty"`
				Enabled           *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				Override_yaml     *struct {
					Metadata *struct {
						Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					} `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec *map[string]string `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"override_yaml" json:"override_yaml,omitempty"`
			} `tfsdk:"ingress" json:"ingress,omitempty"`
			Instance_name *string `tfsdk:"instance_name" json:"instance_name,omitempty"`
			Logger        *struct {
				Log_format        *string `tfsdk:"log_format" json:"log_format,omitempty"`
				Log_level         *string `tfsdk:"log_level" json:"log_level,omitempty"`
				Sampler_rate      *string `tfsdk:"sampler_rate" json:"sampler_rate,omitempty"`
				Time_field_format *string `tfsdk:"time_field_format" json:"time_field_format,omitempty"`
			} `tfsdk:"logger" json:"logger,omitempty"`
			Namespace           *string              `tfsdk:"namespace" json:"namespace,omitempty"`
			Node_selector       *map[string]string   `tfsdk:"node_selector" json:"node_selector,omitempty"`
			Pod_annotations     *map[string]string   `tfsdk:"pod_annotations" json:"pod_annotations,omitempty"`
			Pod_labels          *map[string]string   `tfsdk:"pod_labels" json:"pod_labels,omitempty"`
			Priority_class_name *string              `tfsdk:"priority_class_name" json:"priority_class_name,omitempty"`
			Replicas            *int64               `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources           *map[string]string   `tfsdk:"resources" json:"resources,omitempty"`
			Secret_name         *string              `tfsdk:"secret_name" json:"secret_name,omitempty"`
			Security_context    *map[string]string   `tfsdk:"security_context" json:"security_context,omitempty"`
			Service_annotations *map[string]string   `tfsdk:"service_annotations" json:"service_annotations,omitempty"`
			Service_type        *string              `tfsdk:"service_type" json:"service_type,omitempty"`
			Tolerations         *[]map[string]string `tfsdk:"tolerations" json:"tolerations,omitempty"`
			Version_label       *string              `tfsdk:"version_label" json:"version_label,omitempty"`
			View_only_mode      *bool                `tfsdk:"view_only_mode" json:"view_only_mode,omitempty"`
		} `tfsdk:"deployment" json:"deployment,omitempty"`
		Extensions *[]struct {
			Enabled *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Name    *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"extensions" json:"extensions,omitempty"`
		External_services *struct {
			Custom_dashboards *struct {
				Discovery_auto_threshold *int64  `tfsdk:"discovery_auto_threshold" json:"discovery_auto_threshold,omitempty"`
				Discovery_enabled        *string `tfsdk:"discovery_enabled" json:"discovery_enabled,omitempty"`
				Enabled                  *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				Is_core                  *bool   `tfsdk:"is_core" json:"is_core,omitempty"`
				Namespace_label          *string `tfsdk:"namespace_label" json:"namespace_label,omitempty"`
				Prometheus               *struct {
					Auth *struct {
						Ca_file              *string `tfsdk:"ca_file" json:"ca_file,omitempty"`
						Insecure_skip_verify *bool   `tfsdk:"insecure_skip_verify" json:"insecure_skip_verify,omitempty"`
						Password             *string `tfsdk:"password" json:"password,omitempty"`
						Token                *string `tfsdk:"token" json:"token,omitempty"`
						Type                 *string `tfsdk:"type" json:"type,omitempty"`
						Use_kiali_token      *bool   `tfsdk:"use_kiali_token" json:"use_kiali_token,omitempty"`
						Username             *string `tfsdk:"username" json:"username,omitempty"`
					} `tfsdk:"auth" json:"auth,omitempty"`
					Cache_duration   *int64             `tfsdk:"cache_duration" json:"cache_duration,omitempty"`
					Cache_enabled    *bool              `tfsdk:"cache_enabled" json:"cache_enabled,omitempty"`
					Cache_expiration *int64             `tfsdk:"cache_expiration" json:"cache_expiration,omitempty"`
					Custom_headers   *map[string]string `tfsdk:"custom_headers" json:"custom_headers,omitempty"`
					Health_check_url *string            `tfsdk:"health_check_url" json:"health_check_url,omitempty"`
					Is_core          *bool              `tfsdk:"is_core" json:"is_core,omitempty"`
					Query_scope      *map[string]string `tfsdk:"query_scope" json:"query_scope,omitempty"`
					Thanos_proxy     *struct {
						Enabled          *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
						Retention_period *string `tfsdk:"retention_period" json:"retention_period,omitempty"`
						Scrape_interval  *string `tfsdk:"scrape_interval" json:"scrape_interval,omitempty"`
					} `tfsdk:"thanos_proxy" json:"thanos_proxy,omitempty"`
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"prometheus" json:"prometheus,omitempty"`
			} `tfsdk:"custom_dashboards" json:"custom_dashboards,omitempty"`
			Grafana *struct {
				Auth *struct {
					Ca_file              *string `tfsdk:"ca_file" json:"ca_file,omitempty"`
					Insecure_skip_verify *bool   `tfsdk:"insecure_skip_verify" json:"insecure_skip_verify,omitempty"`
					Password             *string `tfsdk:"password" json:"password,omitempty"`
					Token                *string `tfsdk:"token" json:"token,omitempty"`
					Type                 *string `tfsdk:"type" json:"type,omitempty"`
					Use_kiali_token      *bool   `tfsdk:"use_kiali_token" json:"use_kiali_token,omitempty"`
					Username             *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				Dashboards *[]struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Variables *struct {
						App       *string `tfsdk:"app" json:"app,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Service   *string `tfsdk:"service" json:"service,omitempty"`
						Workload  *string `tfsdk:"workload" json:"workload,omitempty"`
					} `tfsdk:"variables" json:"variables,omitempty"`
				} `tfsdk:"dashboards" json:"dashboards,omitempty"`
				Enabled          *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				External_url     *string `tfsdk:"external_url" json:"external_url,omitempty"`
				Health_check_url *string `tfsdk:"health_check_url" json:"health_check_url,omitempty"`
				Internal_url     *string `tfsdk:"internal_url" json:"internal_url,omitempty"`
				Is_core          *bool   `tfsdk:"is_core" json:"is_core,omitempty"`
			} `tfsdk:"grafana" json:"grafana,omitempty"`
			Istio *struct {
				Component_status *struct {
					Components *[]struct {
						App_label *string `tfsdk:"app_label" json:"app_label,omitempty"`
						Is_core   *bool   `tfsdk:"is_core" json:"is_core,omitempty"`
						Is_proxy  *bool   `tfsdk:"is_proxy" json:"is_proxy,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"components" json:"components,omitempty"`
					Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
				} `tfsdk:"component_status" json:"component_status,omitempty"`
				Config_map_name          *string `tfsdk:"config_map_name" json:"config_map_name,omitempty"`
				Egress_gateway_namespace *string `tfsdk:"egress_gateway_namespace" json:"egress_gateway_namespace,omitempty"`
				Envoy_admin_local_port   *int64  `tfsdk:"envoy_admin_local_port" json:"envoy_admin_local_port,omitempty"`
				Gateway_api_classes      *[]struct {
					Class_name *string `tfsdk:"class_name" json:"class_name,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"gateway_api_classes" json:"gateway_api_classes,omitempty"`
				Ingress_gateway_namespace *string `tfsdk:"ingress_gateway_namespace" json:"ingress_gateway_namespace,omitempty"`
				Istio_api_enabled         *bool   `tfsdk:"istio_api_enabled" json:"istio_api_enabled,omitempty"`
				Istio_canary_revision     *struct {
					Current *string `tfsdk:"current" json:"current,omitempty"`
					Upgrade *string `tfsdk:"upgrade" json:"upgrade,omitempty"`
				} `tfsdk:"istio_canary_revision" json:"istio_canary_revision,omitempty"`
				Istio_identity_domain                  *string `tfsdk:"istio_identity_domain" json:"istio_identity_domain,omitempty"`
				Istio_injection_annotation             *string `tfsdk:"istio_injection_annotation" json:"istio_injection_annotation,omitempty"`
				Istio_sidecar_annotation               *string `tfsdk:"istio_sidecar_annotation" json:"istio_sidecar_annotation,omitempty"`
				Istio_sidecar_injector_config_map_name *string `tfsdk:"istio_sidecar_injector_config_map_name" json:"istio_sidecar_injector_config_map_name,omitempty"`
				Istiod_deployment_name                 *string `tfsdk:"istiod_deployment_name" json:"istiod_deployment_name,omitempty"`
				Istiod_pod_monitoring_port             *int64  `tfsdk:"istiod_pod_monitoring_port" json:"istiod_pod_monitoring_port,omitempty"`
				Root_namespace                         *string `tfsdk:"root_namespace" json:"root_namespace,omitempty"`
				Url_service_version                    *string `tfsdk:"url_service_version" json:"url_service_version,omitempty"`
			} `tfsdk:"istio" json:"istio,omitempty"`
			Prometheus *struct {
				Auth *struct {
					Ca_file              *string `tfsdk:"ca_file" json:"ca_file,omitempty"`
					Insecure_skip_verify *bool   `tfsdk:"insecure_skip_verify" json:"insecure_skip_verify,omitempty"`
					Password             *string `tfsdk:"password" json:"password,omitempty"`
					Token                *string `tfsdk:"token" json:"token,omitempty"`
					Type                 *string `tfsdk:"type" json:"type,omitempty"`
					Use_kiali_token      *bool   `tfsdk:"use_kiali_token" json:"use_kiali_token,omitempty"`
					Username             *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				Cache_duration   *int64             `tfsdk:"cache_duration" json:"cache_duration,omitempty"`
				Cache_enabled    *bool              `tfsdk:"cache_enabled" json:"cache_enabled,omitempty"`
				Cache_expiration *int64             `tfsdk:"cache_expiration" json:"cache_expiration,omitempty"`
				Custom_headers   *map[string]string `tfsdk:"custom_headers" json:"custom_headers,omitempty"`
				Health_check_url *string            `tfsdk:"health_check_url" json:"health_check_url,omitempty"`
				Is_core          *bool              `tfsdk:"is_core" json:"is_core,omitempty"`
				Query_scope      *map[string]string `tfsdk:"query_scope" json:"query_scope,omitempty"`
				Thanos_proxy     *struct {
					Enabled          *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Retention_period *string `tfsdk:"retention_period" json:"retention_period,omitempty"`
					Scrape_interval  *string `tfsdk:"scrape_interval" json:"scrape_interval,omitempty"`
				} `tfsdk:"thanos_proxy" json:"thanos_proxy,omitempty"`
				Url *string `tfsdk:"url" json:"url,omitempty"`
			} `tfsdk:"prometheus" json:"prometheus,omitempty"`
			Tracing *struct {
				Auth *struct {
					Ca_file              *string `tfsdk:"ca_file" json:"ca_file,omitempty"`
					Insecure_skip_verify *bool   `tfsdk:"insecure_skip_verify" json:"insecure_skip_verify,omitempty"`
					Password             *string `tfsdk:"password" json:"password,omitempty"`
					Token                *string `tfsdk:"token" json:"token,omitempty"`
					Type                 *string `tfsdk:"type" json:"type,omitempty"`
					Use_kiali_token      *bool   `tfsdk:"use_kiali_token" json:"use_kiali_token,omitempty"`
					Username             *string `tfsdk:"username" json:"username,omitempty"`
				} `tfsdk:"auth" json:"auth,omitempty"`
				Custom_headers     *map[string]string `tfsdk:"custom_headers" json:"custom_headers,omitempty"`
				Enabled            *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
				External_url       *string            `tfsdk:"external_url" json:"external_url,omitempty"`
				Grpc_port          *int64             `tfsdk:"grpc_port" json:"grpc_port,omitempty"`
				Health_check_url   *string            `tfsdk:"health_check_url" json:"health_check_url,omitempty"`
				Internal_url       *string            `tfsdk:"internal_url" json:"internal_url,omitempty"`
				Is_core            *bool              `tfsdk:"is_core" json:"is_core,omitempty"`
				Namespace_selector *bool              `tfsdk:"namespace_selector" json:"namespace_selector,omitempty"`
				Provider           *string            `tfsdk:"provider" json:"provider,omitempty"`
				Query_scope        *map[string]string `tfsdk:"query_scope" json:"query_scope,omitempty"`
				Query_timeout      *int64             `tfsdk:"query_timeout" json:"query_timeout,omitempty"`
				Tempo_config       *struct {
					Datasource_uid *string `tfsdk:"datasource_uid" json:"datasource_uid,omitempty"`
					Org_id         *string `tfsdk:"org_id" json:"org_id,omitempty"`
					Url_format     *string `tfsdk:"url_format" json:"url_format,omitempty"`
				} `tfsdk:"tempo_config" json:"tempo_config,omitempty"`
				Use_grpc               *bool     `tfsdk:"use_grpc" json:"use_grpc,omitempty"`
				Whitelist_istio_system *[]string `tfsdk:"whitelist_istio_system" json:"whitelist_istio_system,omitempty"`
			} `tfsdk:"tracing" json:"tracing,omitempty"`
		} `tfsdk:"external_services" json:"external_services,omitempty"`
		Health_config *struct {
			Rate *[]struct {
				Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				Tolerance *[]struct {
					Code      *string `tfsdk:"code" json:"code,omitempty"`
					Degraded  *int64  `tfsdk:"degraded" json:"degraded,omitempty"`
					Direction *string `tfsdk:"direction" json:"direction,omitempty"`
					Failure   *int64  `tfsdk:"failure" json:"failure,omitempty"`
					Protocol  *string `tfsdk:"protocol" json:"protocol,omitempty"`
				} `tfsdk:"tolerance" json:"tolerance,omitempty"`
			} `tfsdk:"rate" json:"rate,omitempty"`
		} `tfsdk:"health_config" json:"health_config,omitempty"`
		Identity *struct {
			Cert_file        *string `tfsdk:"cert_file" json:"cert_file,omitempty"`
			Private_key_file *string `tfsdk:"private_key_file" json:"private_key_file,omitempty"`
		} `tfsdk:"identity" json:"identity,omitempty"`
		Installation_tag *string `tfsdk:"installation_tag" json:"installation_tag,omitempty"`
		Istio_labels     *struct {
			App_label_name        *string `tfsdk:"app_label_name" json:"app_label_name,omitempty"`
			Egress_gateway_label  *string `tfsdk:"egress_gateway_label" json:"egress_gateway_label,omitempty"`
			Ingress_gateway_label *string `tfsdk:"ingress_gateway_label" json:"ingress_gateway_label,omitempty"`
			Injection_label_name  *string `tfsdk:"injection_label_name" json:"injection_label_name,omitempty"`
			Injection_label_rev   *string `tfsdk:"injection_label_rev" json:"injection_label_rev,omitempty"`
			Version_label_name    *string `tfsdk:"version_label_name" json:"version_label_name,omitempty"`
		} `tfsdk:"istio_labels" json:"istio_labels,omitempty"`
		Istio_namespace     *string `tfsdk:"istio_namespace" json:"istio_namespace,omitempty"`
		Kiali_feature_flags *struct {
			Disabled_features       *[]string `tfsdk:"disabled_features" json:"disabled_features,omitempty"`
			Istio_annotation_action *bool     `tfsdk:"istio_annotation_action" json:"istio_annotation_action,omitempty"`
			Istio_injection_action  *bool     `tfsdk:"istio_injection_action" json:"istio_injection_action,omitempty"`
			Istio_upgrade_action    *bool     `tfsdk:"istio_upgrade_action" json:"istio_upgrade_action,omitempty"`
			Ui_defaults             *struct {
				Graph *struct {
					Find_options *[]struct {
						Auto_select *bool   `tfsdk:"auto_select" json:"auto_select,omitempty"`
						Description *string `tfsdk:"description" json:"description,omitempty"`
						Expression  *string `tfsdk:"expression" json:"expression,omitempty"`
					} `tfsdk:"find_options" json:"find_options,omitempty"`
					Hide_options *[]struct {
						Auto_select *bool   `tfsdk:"auto_select" json:"auto_select,omitempty"`
						Description *string `tfsdk:"description" json:"description,omitempty"`
						Expression  *string `tfsdk:"expression" json:"expression,omitempty"`
					} `tfsdk:"hide_options" json:"hide_options,omitempty"`
					Traffic *struct {
						Grpc *string `tfsdk:"grpc" json:"grpc,omitempty"`
						Http *string `tfsdk:"http" json:"http,omitempty"`
						Tcp  *string `tfsdk:"tcp" json:"tcp,omitempty"`
					} `tfsdk:"traffic" json:"traffic,omitempty"`
				} `tfsdk:"graph" json:"graph,omitempty"`
				I18n *struct {
					Language      *string `tfsdk:"language" json:"language,omitempty"`
					Show_selector *bool   `tfsdk:"show_selector" json:"show_selector,omitempty"`
				} `tfsdk:"i18n" json:"i18n,omitempty"`
				List *struct {
					Include_health          *bool `tfsdk:"include_health" json:"include_health,omitempty"`
					Include_istio_resources *bool `tfsdk:"include_istio_resources" json:"include_istio_resources,omitempty"`
					Include_validations     *bool `tfsdk:"include_validations" json:"include_validations,omitempty"`
					Show_include_toggles    *bool `tfsdk:"show_include_toggles" json:"show_include_toggles,omitempty"`
				} `tfsdk:"list" json:"list,omitempty"`
				Metrics_inbound *struct {
					Aggregations *[]struct {
						Display_name *string `tfsdk:"display_name" json:"display_name,omitempty"`
						Label        *string `tfsdk:"label" json:"label,omitempty"`
					} `tfsdk:"aggregations" json:"aggregations,omitempty"`
				} `tfsdk:"metrics_inbound" json:"metrics_inbound,omitempty"`
				Metrics_outbound *struct {
					Aggregations *[]struct {
						Display_name *string `tfsdk:"display_name" json:"display_name,omitempty"`
						Label        *string `tfsdk:"label" json:"label,omitempty"`
					} `tfsdk:"aggregations" json:"aggregations,omitempty"`
				} `tfsdk:"metrics_outbound" json:"metrics_outbound,omitempty"`
				Metrics_per_refresh *string   `tfsdk:"metrics_per_refresh" json:"metrics_per_refresh,omitempty"`
				Namespaces          *[]string `tfsdk:"namespaces" json:"namespaces,omitempty"`
				Refresh_interval    *string   `tfsdk:"refresh_interval" json:"refresh_interval,omitempty"`
			} `tfsdk:"ui_defaults" json:"ui_defaults,omitempty"`
			Validations *struct {
				Ignore                      *[]string `tfsdk:"ignore" json:"ignore,omitempty"`
				Skip_wildcard_gateway_hosts *bool     `tfsdk:"skip_wildcard_gateway_hosts" json:"skip_wildcard_gateway_hosts,omitempty"`
			} `tfsdk:"validations" json:"validations,omitempty"`
		} `tfsdk:"kiali_feature_flags" json:"kiali_feature_flags,omitempty"`
		Kubernetes_config *struct {
			Burst                          *int64    `tfsdk:"burst" json:"burst,omitempty"`
			Cache_duration                 *int64    `tfsdk:"cache_duration" json:"cache_duration,omitempty"`
			Cache_token_namespace_duration *int64    `tfsdk:"cache_token_namespace_duration" json:"cache_token_namespace_duration,omitempty"`
			Cluster_name                   *string   `tfsdk:"cluster_name" json:"cluster_name,omitempty"`
			Excluded_workloads             *[]string `tfsdk:"excluded_workloads" json:"excluded_workloads,omitempty"`
			Qps                            *int64    `tfsdk:"qps" json:"qps,omitempty"`
		} `tfsdk:"kubernetes_config" json:"kubernetes_config,omitempty"`
		Login_token *struct {
			Expiration_seconds *int64  `tfsdk:"expiration_seconds" json:"expiration_seconds,omitempty"`
			Signing_key        *string `tfsdk:"signing_key" json:"signing_key,omitempty"`
		} `tfsdk:"login_token" json:"login_token,omitempty"`
		Server *struct {
			Address        *string `tfsdk:"address" json:"address,omitempty"`
			Audit_log      *bool   `tfsdk:"audit_log" json:"audit_log,omitempty"`
			Cors_allow_all *bool   `tfsdk:"cors_allow_all" json:"cors_allow_all,omitempty"`
			Gzip_enabled   *bool   `tfsdk:"gzip_enabled" json:"gzip_enabled,omitempty"`
			Node_port      *int64  `tfsdk:"node_port" json:"node_port,omitempty"`
			Observability  *struct {
				Metrics *struct {
					Enabled *bool  `tfsdk:"enabled" json:"enabled,omitempty"`
					Port    *int64 `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
				Tracing *struct {
					Collector_type *string `tfsdk:"collector_type" json:"collector_type,omitempty"`
					Collector_url  *string `tfsdk:"collector_url" json:"collector_url,omitempty"`
					Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Otel           *struct {
						Ca_name     *string `tfsdk:"ca_name" json:"ca_name,omitempty"`
						Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
						Skip_verify *bool   `tfsdk:"skip_verify" json:"skip_verify,omitempty"`
						Tls_enabled *bool   `tfsdk:"tls_enabled" json:"tls_enabled,omitempty"`
					} `tfsdk:"otel" json:"otel,omitempty"`
				} `tfsdk:"tracing" json:"tracing,omitempty"`
			} `tfsdk:"observability" json:"observability,omitempty"`
			Port     *int64 `tfsdk:"port" json:"port,omitempty"`
			Profiler *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"profiler" json:"profiler,omitempty"`
			Web_fqdn         *string `tfsdk:"web_fqdn" json:"web_fqdn,omitempty"`
			Web_history_mode *string `tfsdk:"web_history_mode" json:"web_history_mode,omitempty"`
			Web_port         *string `tfsdk:"web_port" json:"web_port,omitempty"`
			Web_root         *string `tfsdk:"web_root" json:"web_root,omitempty"`
			Web_schema       *string `tfsdk:"web_schema" json:"web_schema,omitempty"`
			Write_timeout    *int64  `tfsdk:"write_timeout" json:"write_timeout,omitempty"`
		} `tfsdk:"server" json:"server,omitempty"`
		Version *string `tfsdk:"version" json:"version,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *KialiIoKialiV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_kiali_io_kiali_v1alpha1_manifest"
}

func (r *KialiIoKialiV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "This is the CRD for the resources called Kiali CRs. The Kiali Operator will watch for resources of this type and when it detects a Kiali CR has been added, deleted, or modified, it will install, uninstall, and update the associated Kiali Server installation. The settings here will configure the Kiali Server as well as the Kiali Operator. All of these settings will be stored in the Kiali ConfigMap. Do not modify the ConfigMap; it will be managed by the Kiali Operator. Only modify the Kiali CR when you want to change a configuration setting.",
				MarkdownDescription: "This is the CRD for the resources called Kiali CRs. The Kiali Operator will watch for resources of this type and when it detects a Kiali CR has been added, deleted, or modified, it will install, uninstall, and update the associated Kiali Server installation. The settings here will configure the Kiali Server as well as the Kiali Operator. All of these settings will be stored in the Kiali ConfigMap. Do not modify the ConfigMap; it will be managed by the Kiali Operator. Only modify the Kiali CR when you want to change a configuration setting.",
				Attributes: map[string]schema.Attribute{
					"additional_display_details": schema.ListNestedAttribute{
						Description:         "A list of additional details that Kiali will look for in annotations. When found on any workload or service, Kiali will display the additional details in the respective workload or service details page. This is typically used to inject some CI metadata or documentation links into Kiali views. For example, by default, Kiali will recognize these annotations on a service or workload (e.g. a Deployment, StatefulSet, etc.): ''' annotations: kiali.io/api-spec: http://list/to/my/api/doc kiali.io/api-type: rest ''' Note that if you change this setting for your own custom annotations, keep in mind that it would override the current default. So you would have to add the default setting as shown in the example CR if you want to preserve the default links. ",
						MarkdownDescription: "A list of additional details that Kiali will look for in annotations. When found on any workload or service, Kiali will display the additional details in the respective workload or service details page. This is typically used to inject some CI metadata or documentation links into Kiali views. For example, by default, Kiali will recognize these annotations on a service or workload (e.g. a Deployment, StatefulSet, etc.): ''' annotations: kiali.io/api-spec: http://list/to/my/api/doc kiali.io/api-type: rest ''' Note that if you change this setting for your own custom annotations, keep in mind that it would override the current default. So you would have to add the default setting as shown in the example CR if you want to preserve the default links. ",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"annotation": schema.StringAttribute{
									Description:         "The name of the annotation whose value is a URL to additional documentation useful to the user.",
									MarkdownDescription: "The name of the annotation whose value is a URL to additional documentation useful to the user.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"icon_annotation": schema.StringAttribute{
									Description:         "The name of the annotation whose value is used to determine what icon to display. The annotation name itself can be anything, but note that the value of that annotation must be one of: 'rest', 'grpc', and 'graphql' - any other value is ignored.",
									MarkdownDescription: "The name of the annotation whose value is used to determine what icon to display. The annotation name itself can be anything, but note that the value of that annotation must be one of: 'rest', 'grpc', and 'graphql' - any other value is ignored.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"title": schema.StringAttribute{
									Description:         "The title of the link that Kiali will display. The link will go to the URL specified in the value of the configured 'annotation'.",
									MarkdownDescription: "The title of the link that Kiali will display. The link will go to the URL specified in the value of the configured 'annotation'.",
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

					"auth": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"openid": schema.SingleNestedAttribute{
								Description:         "To learn more about these settings and how to configure the OpenId authentication strategy, read the documentation at https://kiali.io/docs/configuration/authentication/openid/",
								MarkdownDescription: "To learn more about these settings and how to configure the OpenId authentication strategy, read the documentation at https://kiali.io/docs/configuration/authentication/openid/",
								Attributes: map[string]schema.Attribute{
									"additional_request_params": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"allowed_domains": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"api_proxy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"api_proxy_ca_data": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"api_token": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"authentication_timeout": schema.Int64Attribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"authorization_endpoint": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"client_id": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"disable_rbac": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"http_proxy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"https_proxy": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"insecure_skip_verify_tls": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"issuer_uri": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scopes": schema.ListAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username_claim": schema.StringAttribute{
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

							"openshift": schema.SingleNestedAttribute{
								Description:         "To learn more about these settings and how to configure the OpenShift authentication strategy, read the documentation at https://kiali.io/docs/configuration/authentication/openshift/",
								MarkdownDescription: "To learn more about these settings and how to configure the OpenShift authentication strategy, read the documentation at https://kiali.io/docs/configuration/authentication/openshift/",
								Attributes: map[string]schema.Attribute{
									"token_inactivity_timeout": schema.Int64Attribute{
										Description:         "Timeout that overrides the default OpenShift token inactivity timeout. This value represents the maximum amount of time in seconds that can occur between consecutive uses of the token. Tokens become invalid if they are not used within this temporal window. If 0, the Kiali tokens never timeout. OpenShift may have a minimum allowed value - see the OpenShift documentation specific for the version of OpenShift you are using. WARNING: existing tokens will not be affected by changing this setting.",
										MarkdownDescription: "Timeout that overrides the default OpenShift token inactivity timeout. This value represents the maximum amount of time in seconds that can occur between consecutive uses of the token. Tokens become invalid if they are not used within this temporal window. If 0, the Kiali tokens never timeout. OpenShift may have a minimum allowed value - see the OpenShift documentation specific for the version of OpenShift you are using. WARNING: existing tokens will not be affected by changing this setting.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"token_max_age": schema.Int64Attribute{
										Description:         "A time duration in seconds that overrides the default OpenShift access token max age. If 0 then there will be no expiration of tokens.",
										MarkdownDescription: "A time duration in seconds that overrides the default OpenShift access token max age. If 0 then there will be no expiration of tokens.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"strategy": schema.StringAttribute{
								Description:         "Determines what authentication strategy to use when users log into Kiali. Options are 'anonymous', 'token', 'openshift', 'openid', or 'header'. * Choose 'anonymous' to allow full access to Kiali without requiring any credentials. * Choose 'token' to allow access to Kiali using service account tokens, which controls access based on RBAC roles assigned to the service account. * Choose 'openshift' to use the OpenShift OAuth login which controls access based on the individual's RBAC roles in OpenShift. Not valid for non-OpenShift environments. * Choose 'openid' to enable OpenID Connect-based authentication. Your cluster is required to be configured to accept the tokens issued by your IdP. There are additional required configurations for this strategy. See below for the additional OpenID configuration section. * Choose 'header' when Kiali is running behind a reverse proxy that will inject an Authorization header and potentially impersonation headers. When empty, this value will default to 'openshift' on OpenShift and 'token' on other Kubernetes environments. ",
								MarkdownDescription: "Determines what authentication strategy to use when users log into Kiali. Options are 'anonymous', 'token', 'openshift', 'openid', or 'header'. * Choose 'anonymous' to allow full access to Kiali without requiring any credentials. * Choose 'token' to allow access to Kiali using service account tokens, which controls access based on RBAC roles assigned to the service account. * Choose 'openshift' to use the OpenShift OAuth login which controls access based on the individual's RBAC roles in OpenShift. Not valid for non-OpenShift environments. * Choose 'openid' to enable OpenID Connect-based authentication. Your cluster is required to be configured to accept the tokens issued by your IdP. There are additional required configurations for this strategy. See below for the additional OpenID configuration section. * Choose 'header' when Kiali is running behind a reverse proxy that will inject an Authorization header and potentially impersonation headers. When empty, this value will default to 'openshift' on OpenShift and 'token' on other Kubernetes environments. ",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "anonymous", "token", "openshift", "openid", "header"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"clustering": schema.SingleNestedAttribute{
						Description:         "Multi-cluster related features.",
						MarkdownDescription: "Multi-cluster related features.",
						Attributes: map[string]schema.Attribute{
							"autodetect_secrets": schema.SingleNestedAttribute{
								Description:         "Settings to allow cluster secrets to be auto-detected. Secrets must exist in the Kiali deployment namespace.",
								MarkdownDescription: "Settings to allow cluster secrets to be auto-detected. Secrets must exist in the Kiali deployment namespace.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "If true then remote cluster secrets will be autodetected during the installation of the Kiali Server Deployment. Any remote cluster secrets found in the Kiali deployment namespace will be mounted to the Kiali Server's file system. If false, you can still manually specify the remote cluster secret information in the 'clusters' setting if you wish to utilize multicluster features.",
										MarkdownDescription: "If true then remote cluster secrets will be autodetected during the installation of the Kiali Server Deployment. Any remote cluster secrets found in the Kiali deployment namespace will be mounted to the Kiali Server's file system. If false, you can still manually specify the remote cluster secret information in the 'clusters' setting if you wish to utilize multicluster features.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"label": schema.StringAttribute{
										Description:         "The name and value of a label that exists on all remote cluster secrets. Default is 'kiali.io/multiCluster=true'.",
										MarkdownDescription: "The name and value of a label that exists on all remote cluster secrets. Default is 'kiali.io/multiCluster=true'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"clusters": schema.ListNestedAttribute{
								Description:         "A list of clusters that the Kiali Server can access. You need to specify the remote clusters here if 'autodetect_secrets.enabled' is false.",
								MarkdownDescription: "A list of clusters that the Kiali Server can access. You need to specify the remote clusters here if 'autodetect_secrets.enabled' is false.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "The name of the cluster.",
											MarkdownDescription: "The name of the cluster.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secret_name": schema.StringAttribute{
											Description:         "The name of the secret that contains the credentials necessary to connect to the remote cluster. This secret must exist in the Kiali deployment namespace. If a secret name is not provided then it's assumed that the cluster is inaccessible.",
											MarkdownDescription: "The name of the secret that contains the credentials necessary to connect to the remote cluster. This secret must exist in the Kiali deployment namespace. If a secret name is not provided then it's assumed that the cluster is inaccessible.",
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

							"kiali_urls": schema.ListNestedAttribute{
								Description:         "A map between cluster name, instance name and namespace to a Kiali URL. Will be used showing the Mesh page's Kiali URLs. The Kiali service's 'kiali.io/external-url' annotation will be overridden when this property is set.",
								MarkdownDescription: "A map between cluster name, instance name and namespace to a Kiali URL. Will be used showing the Mesh page's Kiali URLs. The Kiali service's 'kiali.io/external-url' annotation will be overridden when this property is set.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cluster_name": schema.StringAttribute{
											Description:         "The name of the cluster.",
											MarkdownDescription: "The name of the cluster.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"instance_name": schema.StringAttribute{
											Description:         "The instance name of this Kiali installation. This should be the value used in 'deployment.instance_name' for Kiali resource name.",
											MarkdownDescription: "The instance name of this Kiali installation. This should be the value used in 'deployment.instance_name' for Kiali resource name.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "The namespace into which Kiali is installed.",
											MarkdownDescription: "The namespace into which Kiali is installed.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"url": schema.StringAttribute{
											Description:         "The URL of Kiali in the cluster.",
											MarkdownDescription: "The URL of Kiali in the cluster.",
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

					"custom_dashboards": schema.ListAttribute{
						Description:         "A list of user-defined custom monitoring dashboards that you can use to generate metrics charts for your applications. The server has some built-in dashboards; if you define a custom dashboard here with the same name as a built-in dashboard, your custom dashboard takes precedence and will overwrite the built-in dashboard. You can disable one or more of the built-in dashboards by simply defining an empty dashboard. An example of an additional user-defined dashboard, ''' - name: myapp title: My App Metrics items: - chart: name: 'Thread Count' spans: 4 metricName: 'thread-count' dataType: 'raw' ''' An example of disabling a built-in dashboard (in this case, disabling the Envoy dashboard), ''' - name: envoy ''' To learn more about custom monitoring dashboards, see the documentation at https://kiali.io/docs/configuration/custom-dashboard/ ",
						MarkdownDescription: "A list of user-defined custom monitoring dashboards that you can use to generate metrics charts for your applications. The server has some built-in dashboards; if you define a custom dashboard here with the same name as a built-in dashboard, your custom dashboard takes precedence and will overwrite the built-in dashboard. You can disable one or more of the built-in dashboards by simply defining an empty dashboard. An example of an additional user-defined dashboard, ''' - name: myapp title: My App Metrics items: - chart: name: 'Thread Count' spans: 4 metricName: 'thread-count' dataType: 'raw' ''' An example of disabling a built-in dashboard (in this case, disabling the Envoy dashboard), ''' - name: envoy ''' To learn more about custom monitoring dashboards, see the documentation at https://kiali.io/docs/configuration/custom-dashboard/ ",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deployment": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"additional_service_yaml": schema.MapAttribute{
								Description:         "Additional custom yaml to add to the service definition. This is used mainly to customize the service type. For example, if the 'deployment.service_type' is set to 'LoadBalancer' and you want to set the loadBalancerIP, you can do so here with: 'additional_service_yaml: { 'loadBalancerIP': '78.11.24.19' }'. Another example would be if the 'deployment.service_type' is set to 'ExternalName' you will need to configure the name via: 'additional_service_yaml: { 'externalName': 'my.kiali.example.com' }'. A final example would be if external IPs need to be set: 'additional_service_yaml: { 'externalIPs': ['80.11.12.10'] }'",
								MarkdownDescription: "Additional custom yaml to add to the service definition. This is used mainly to customize the service type. For example, if the 'deployment.service_type' is set to 'LoadBalancer' and you want to set the loadBalancerIP, you can do so here with: 'additional_service_yaml: { 'loadBalancerIP': '78.11.24.19' }'. Another example would be if the 'deployment.service_type' is set to 'ExternalName' you will need to configure the name via: 'additional_service_yaml: { 'externalName': 'my.kiali.example.com' }'. A final example would be if external IPs need to be set: 'additional_service_yaml: { 'externalIPs': ['80.11.12.10'] }'",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"affinity": schema.SingleNestedAttribute{
								Description:         "Affinity definitions that are to be used to define the nodes where the Kiali pod should be constrained. See the Kubernetes documentation on Assigning Pods to Nodes for the proper syntax for these three different affinity types.",
								MarkdownDescription: "Affinity definitions that are to be used to define the nodes where the Kiali pod should be constrained. See the Kubernetes documentation on Assigning Pods to Nodes for the proper syntax for these three different affinity types.",
								Attributes: map[string]schema.Attribute{
									"node": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"pod_anti": schema.MapAttribute{
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

							"cluster_wide_access": schema.BoolAttribute{
								Description:         "Determines if the Kiali server will be granted cluster-wide permissions to see all namespaces. When true, this provides more efficient caching within the Kiali server. It must be 'true' if 'deployment.discovery_selectors.default' is left unset. To limit the namespaces for which Kiali has permissions, set to 'false' and define the desired selectors in 'deployment.discovery_selectors.default'. When not set, this value will default to 'true'.",
								MarkdownDescription: "Determines if the Kiali server will be granted cluster-wide permissions to see all namespaces. When true, this provides more efficient caching within the Kiali server. It must be 'true' if 'deployment.discovery_selectors.default' is left unset. To limit the namespaces for which Kiali has permissions, set to 'false' and define the desired selectors in 'deployment.discovery_selectors.default'. When not set, this value will default to 'true'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"configmap_annotations": schema.MapAttribute{
								Description:         "Custom annotations to be created on the Kiali ConfigMap.",
								MarkdownDescription: "Custom annotations to be created on the Kiali ConfigMap.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_envs": schema.ListNestedAttribute{
								Description:         "Defines additional environment variables to be set in the Kiali server pod. This is typically used for (but not limited to) setting proxy environment variables such as HTTP_PROXY, HTTPS_PROXY, and/or NO_PROXY.",
								MarkdownDescription: "Defines additional environment variables to be set in the Kiali server pod. This is typically used for (but not limited to) setting proxy environment variables such as HTTP_PROXY, HTTPS_PROXY, and/or NO_PROXY.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "The name of the custom environment variable.",
											MarkdownDescription: "The name of the custom environment variable.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "The value of the custom environment variable.",
											MarkdownDescription: "The value of the custom environment variable.",
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

							"custom_secrets": schema.ListNestedAttribute{
								Description:         "Defines additional secrets that are to be mounted in the Kiali pod. These are useful to contain certs that are used by Kiali to securely connect to third party systems (for example, see 'external_services.tracing.auth.ca_file'). These secrets must be created by an external mechanism. Kiali will not generate these secrets; it is assumed these secrets are externally managed. You can define 0, 1, or more secrets. An example configuration is, ''' custom_secrets: - name: mysecret mount: /mysecret-path - name: my-other-secret mount: /my-other-secret-location optional: true ''' ",
								MarkdownDescription: "Defines additional secrets that are to be mounted in the Kiali pod. These are useful to contain certs that are used by Kiali to securely connect to third party systems (for example, see 'external_services.tracing.auth.ca_file'). These secrets must be created by an external mechanism. Kiali will not generate these secrets; it is assumed these secrets are externally managed. You can define 0, 1, or more secrets. An example configuration is, ''' custom_secrets: - name: mysecret mount: /mysecret-path - name: my-other-secret mount: /my-other-secret-location optional: true ''' ",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"csi": schema.MapAttribute{
											Description:         "Defines CSI-specific settings that allows a secret from an external CSI secret store to be injected in the pod via a volume mount. For details, see https://secrets-store-csi-driver.sigs.k8s.io/",
											MarkdownDescription: "Defines CSI-specific settings that allows a secret from an external CSI secret store to be injected in the pod via a volume mount. For details, see https://secrets-store-csi-driver.sigs.k8s.io/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"mount": schema.StringAttribute{
											Description:         "The file path location where the secret content will be mounted. The custom secret cannot be mounted on a path that the operator will use to mount its secrets. Make sure you set your custom secret mount path to a unique, unused path. Paths such as '/kiali-configuration', '/kiali-cert', '/kiali-cabundle', and '/kiali-secret' should not be used as mount paths for custom secrets because the operator may want to use one of those paths.",
											MarkdownDescription: "The file path location where the secret content will be mounted. The custom secret cannot be mounted on a path that the operator will use to mount its secrets. Make sure you set your custom secret mount path to a unique, unused path. Paths such as '/kiali-configuration', '/kiali-cert', '/kiali-cabundle', and '/kiali-secret' should not be used as mount paths for custom secrets because the operator may want to use one of those paths.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "The name of the secret that is to be mounted to the Kiali pod's file system. The name of the custom secret must not be the same name as one created by the operator. Names such as 'kiali', 'kiali-cert-secret', and 'kiali-cabundle' should not be used as a custom secret name because the operator may want to create one with one of those names.",
											MarkdownDescription: "The name of the secret that is to be mounted to the Kiali pod's file system. The name of the custom secret must not be the same name as one created by the operator. Names such as 'kiali', 'kiali-cert-secret', and 'kiali-cabundle' should not be used as a custom secret name because the operator may want to create one with one of those names.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"optional": schema.BoolAttribute{
											Description:         "Indicates if the secret may or may not exist at the time the Kiali pod starts. This will default to 'false' if not specified. This is ignored if 'csi' is specified - CSI secrets must exist when specified.",
											MarkdownDescription: "Indicates if the secret may or may not exist at the time the Kiali pod starts. This will default to 'false' if not specified. This is ignored if 'csi' is specified - CSI secrets must exist when specified.",
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

							"discovery_selectors": schema.SingleNestedAttribute{
								Description:         "Discovery selectors used to determine which namespaces are accessible to Kiali and which namespaces are visible to Kiali users. You can define discovery selectors to match namespaces on the local cluster as well as remote clusters. The list of namespaces that a user can access is a subset of these namespaces, given that user's RBAC permissions. These selectors will have similar semantics as defined by Istio ( https://istio.io/latest/docs/reference/config/istio.mesh.v1alpha1/#MeshConfig ) and the syntax of the equality-based and set-based label selectors are documented by Kubernetes here ( https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#resources-that-support-set-based-requirements ) ",
								MarkdownDescription: "Discovery selectors used to determine which namespaces are accessible to Kiali and which namespaces are visible to Kiali users. You can define discovery selectors to match namespaces on the local cluster as well as remote clusters. The list of namespaces that a user can access is a subset of these namespaces, given that user's RBAC permissions. These selectors will have similar semantics as defined by Istio ( https://istio.io/latest/docs/reference/config/istio.mesh.v1alpha1/#MeshConfig ) and the syntax of the equality-based and set-based label selectors are documented by Kubernetes here ( https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#resources-that-support-set-based-requirements ) ",
								Attributes: map[string]schema.Attribute{
									"default": schema.ListNestedAttribute{
										Description:         "These are label selectors for the Kiali local cluster and for all remote clusters that do not have overrides. Namespaces that match these selectors are visible to Kiali users. When 'cluster_wide_access=false' these 'default' selectors are used to restrict which namespaces Kiali will have access to. If there are no default discovery selectors, then 'cluster_wide_access' should be 'true' in which case Kiali will have permissions to access all namespaces. ",
										MarkdownDescription: "These are label selectors for the Kiali local cluster and for all remote clusters that do not have overrides. Namespaces that match these selectors are visible to Kiali users. When 'cluster_wide_access=false' these 'default' selectors are used to restrict which namespaces Kiali will have access to. If there are no default discovery selectors, then 'cluster_wide_access' should be 'true' in which case Kiali will have permissions to access all namespaces. ",
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
																Validators: []validator.String{
																	stringvalidator.OneOf("In", "NotIn", "Exists", "DoesNotExist"),
																},
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"overrides": schema.MapAttribute{
										Description:         "If a remote cluster has different namespaces than the local cluster, these overrides provide a way for you to match those remote namespaces. Kiali will make these remote namespaces visible to users. The name of the overrides section is the name of the remote cluster. Note that the 'default' selectors are ignored when matching namespaces on a remote cluster if that remote cluster has overrides defined.",
										MarkdownDescription: "If a remote cluster has different namespaces than the local cluster, these overrides provide a way for you to match those remote namespaces. Kiali will make these remote namespaces visible to users. The name of the overrides section is the name of the remote cluster. Note that the 'default' selectors are ignored when matching namespaces on a remote cluster if that remote cluster has overrides defined.",
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

							"dns": schema.SingleNestedAttribute{
								Description:         "The Kiali server pod's DNS configuration. Kubernetes supports different DNS policies and configurations. For further details, consult the Kubernetes documentation - https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/ ",
								MarkdownDescription: "The Kiali server pod's DNS configuration. Kubernetes supports different DNS policies and configurations. For further details, consult the Kubernetes documentation - https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/ ",
								Attributes: map[string]schema.Attribute{
									"config": schema.MapAttribute{
										Description:         "DNS configuration that is applied to the DNS policy. See the Kubernetes documentation for the different configuration settings that are supported.",
										MarkdownDescription: "DNS configuration that is applied to the DNS policy. See the Kubernetes documentation for the different configuration settings that are supported.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"policy": schema.StringAttribute{
										Description:         "DNS policy. See the Kubernetes documentation for the different policies that are supported.",
										MarkdownDescription: "DNS policy. See the Kubernetes documentation for the different policies that are supported.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_aliases": schema.ListNestedAttribute{
								Description:         "This is content for the Kubernetes 'hostAliases' setting for the Kiali server. This allows you to modify the Kiali server pod '/etc/hosts' file. A typical way to configure this setting is, ''' host_aliases: - ip: 192.168.1.100 hostnames: - 'foo.local' - 'bar.local' ''' For details on the content of this setting, see https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/#adding-additional-entries-with-hostaliases ",
								MarkdownDescription: "This is content for the Kubernetes 'hostAliases' setting for the Kiali server. This allows you to modify the Kiali server pod '/etc/hosts' file. A typical way to configure this setting is, ''' host_aliases: - ip: 192.168.1.100 hostnames: - 'foo.local' - 'bar.local' ''' For details on the content of this setting, see https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/#adding-additional-entries-with-hostaliases ",
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

							"hpa": schema.SingleNestedAttribute{
								Description:         "Determines what (if any) HorizontalPodAutoscaler should be created to autoscale the Kiali pod. A typical way to configure HPA for Kiali is, ''' hpa: api_version: 'autoscaling/v2' spec: maxReplicas: 2 minReplicas: 1 metrics: - type: Resource resource: name: cpu target: type: Utilization averageUtilization: 50 ''' ",
								MarkdownDescription: "Determines what (if any) HorizontalPodAutoscaler should be created to autoscale the Kiali pod. A typical way to configure HPA for Kiali is, ''' hpa: api_version: 'autoscaling/v2' spec: maxReplicas: 2 minReplicas: 1 metrics: - type: Resource resource: name: cpu target: type: Utilization averageUtilization: 50 ''' ",
								Attributes: map[string]schema.Attribute{
									"api_version": schema.StringAttribute{
										Description:         "A specific HPA API version that can be specified in case there is some HPA feature you want to use that is only supported in that specific version. If value is an empty string, an attempt will be made to determine a valid version.",
										MarkdownDescription: "A specific HPA API version that can be specified in case there is some HPA feature you want to use that is only supported in that specific version. If value is an empty string, an attempt will be made to determine a valid version.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"spec": schema.MapAttribute{
										Description:         "The 'spec' specified here will be placed in the created HPA resource's 'spec' section. If 'spec' is left empty, no HPA resource will be created. Note that you must not specify the 'scaleTargetRef' section in 'spec'; the Kiali Operator will populate that for you.",
										MarkdownDescription: "The 'spec' specified here will be placed in the created HPA resource's 'spec' section. If 'spec' is left empty, no HPA resource will be created. Note that you must not specify the 'scaleTargetRef' section in 'spec'; the Kiali Operator will populate that for you.",
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

							"image_digest": schema.StringAttribute{
								Description:         "If 'deployment.image_version' is a digest hash, this value indicates what type of digest it is. A typical value would be 'sha256'. Note: do NOT prefix this value with a '@'.",
								MarkdownDescription: "If 'deployment.image_version' is a digest hash, this value indicates what type of digest it is. A typical value would be 'sha256'. Note: do NOT prefix this value with a '@'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_name": schema.StringAttribute{
								Description:         "Determines which Kiali image to download and install. If you set this to a specific name (i.e. you do not leave it as the default empty string), you must make sure that image is supported by the operator. If empty, the operator will use a known supported image name based on which 'version' was defined. Note that, as a security measure, a cluster admin may have configured the Kiali operator to ignore this setting. A cluster admin may do this to ensure the Kiali operator only installs a single, specific Kiali version, thus this setting may have no effect depending on how the operator itself was configured.",
								MarkdownDescription: "Determines which Kiali image to download and install. If you set this to a specific name (i.e. you do not leave it as the default empty string), you must make sure that image is supported by the operator. If empty, the operator will use a known supported image name based on which 'version' was defined. Note that, as a security measure, a cluster admin may have configured the Kiali operator to ignore this setting. A cluster admin may do this to ensure the Kiali operator only installs a single, specific Kiali version, thus this setting may have no effect depending on how the operator itself was configured.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "The Kubernetes pull policy for the Kiali deployment. This is overridden to be 'Always' if 'deployment.image_version' is set to 'latest'.",
								MarkdownDescription: "The Kubernetes pull policy for the Kiali deployment. This is overridden to be 'Always' if 'deployment.image_version' is set to 'latest'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_pull_secrets": schema.ListAttribute{
								Description:         "The names of the secrets to be used when container images are to be pulled.",
								MarkdownDescription: "The names of the secrets to be used when container images are to be pulled.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"image_version": schema.StringAttribute{
								Description:         "Determines which version of Kiali to install. Choose 'lastrelease' to use the last Kiali release. Choose 'latest' to use the latest image (which may or may not be a released version of Kiali). Choose 'operator_version' to use the image whose version is the same as the operator version. Otherwise, you can set this to any valid Kiali version (such as 'v1.0') or any valid Kiali digest hash (if you set this to a digest hash, you must indicate the digest in 'deployment.image_digest'). Note that if this is set to 'latest' then the 'deployment.image_pull_policy' will be set to 'Always'. If you set this to a specific version (i.e. you do not leave it as the default empty string), you must make sure that image is supported by the operator. If empty, the operator will use a known supported image version based on which 'version' was defined. Note that, as a security measure, a cluster admin may have configured the Kiali operator to ignore this setting. A cluster admin may do this to ensure the Kiali operator only installs a single, specific Kiali version, thus this setting may have no effect depending on how the operator itself was configured. ",
								MarkdownDescription: "Determines which version of Kiali to install. Choose 'lastrelease' to use the last Kiali release. Choose 'latest' to use the latest image (which may or may not be a released version of Kiali). Choose 'operator_version' to use the image whose version is the same as the operator version. Otherwise, you can set this to any valid Kiali version (such as 'v1.0') or any valid Kiali digest hash (if you set this to a digest hash, you must indicate the digest in 'deployment.image_digest'). Note that if this is set to 'latest' then the 'deployment.image_pull_policy' will be set to 'Always'. If you set this to a specific version (i.e. you do not leave it as the default empty string), you must make sure that image is supported by the operator. If empty, the operator will use a known supported image version based on which 'version' was defined. Note that, as a security measure, a cluster admin may have configured the Kiali operator to ignore this setting. A cluster admin may do this to ensure the Kiali operator only installs a single, specific Kiali version, thus this setting may have no effect depending on how the operator itself was configured. ",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress": schema.SingleNestedAttribute{
								Description:         "Configures if/how the Kiali endpoint should be exposed externally.",
								MarkdownDescription: "Configures if/how the Kiali endpoint should be exposed externally.",
								Attributes: map[string]schema.Attribute{
									"additional_labels": schema.MapAttribute{
										Description:         "Additional labels to add to the Ingress (or Route if on OpenShift). These are added to the labels that are created by default; these do not override the default labels.",
										MarkdownDescription: "Additional labels to add to the Ingress (or Route if on OpenShift). These are added to the labels that are created by default; these do not override the default labels.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"class_name": schema.StringAttribute{
										Description:         "If 'class_name' is a non-empty string, it will be used as the 'spec.ingressClassName' in the created Kubernetes Ingress resource. This setting is ignored if on OpenShift. This is also ignored if 'override_yaml.spec' is defined (i.e. you must define the 'ingressClassName' directly in your override yaml).",
										MarkdownDescription: "If 'class_name' is a non-empty string, it will be used as the 'spec.ingressClassName' in the created Kubernetes Ingress resource. This setting is ignored if on OpenShift. This is also ignored if 'override_yaml.spec' is defined (i.e. you must define the 'ingressClassName' directly in your override yaml).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Determines if the Kiali endpoint should be exposed externally. If 'true', an Ingress will be created if on Kubernetes or a Route if on OpenShift. If left undefined, this will be 'false' on Kubernetes and 'true' on OpenShift.",
										MarkdownDescription: "Determines if the Kiali endpoint should be exposed externally. If 'true', an Ingress will be created if on Kubernetes or a Route if on OpenShift. If left undefined, this will be 'false' on Kubernetes and 'true' on OpenShift.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"override_yaml": schema.SingleNestedAttribute{
										Description:         "Because an Ingress into a cluster can vary wildly in its desired configuration, this setting provides a way to override complete portions of the Ingress resource configuration (Ingress on Kubernetes and Route on OpenShift). It is up to the user to ensure this override YAML configuration is valid and supports the cluster environment since the operator will blindly copy this custom configuration into the resource it creates. This setting is not used if 'deployment.ingress.enabled' is set to 'false'. Note that only 'metadata.annotations' and 'spec' is valid and only they will be used to override those same sections in the created resource. You can define either one or both. Note that 'override_yaml.metadata.labels' is not allowed - you cannot override the labels; to add labels to the default set of labels, use the 'deployment.ingress.additional_labels' setting. Example, ''' override_yaml: metadata: annotations: nginx.ingress.kubernetes.io/secure-backends: 'true' nginx.ingress.kubernetes.io/backend-protocol: 'HTTPS' spec: rules: - http: paths: - path: /kiali pathType: Prefix backend: service name: 'kiali' port: number: 20001 ''' ",
										MarkdownDescription: "Because an Ingress into a cluster can vary wildly in its desired configuration, this setting provides a way to override complete portions of the Ingress resource configuration (Ingress on Kubernetes and Route on OpenShift). It is up to the user to ensure this override YAML configuration is valid and supports the cluster environment since the operator will blindly copy this custom configuration into the resource it creates. This setting is not used if 'deployment.ingress.enabled' is set to 'false'. Note that only 'metadata.annotations' and 'spec' is valid and only they will be used to override those same sections in the created resource. You can define either one or both. Note that 'override_yaml.metadata.labels' is not allowed - you cannot override the labels; to add labels to the default set of labels, use the 'deployment.ingress.additional_labels' setting. Example, ''' override_yaml: metadata: annotations: nginx.ingress.kubernetes.io/secure-backends: 'true' nginx.ingress.kubernetes.io/backend-protocol: 'HTTPS' spec: rules: - http: paths: - path: /kiali pathType: Prefix backend: service name: 'kiali' port: number: 20001 ''' ",
										Attributes: map[string]schema.Attribute{
											"metadata": schema.SingleNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
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

											"spec": schema.MapAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"instance_name": schema.StringAttribute{
								Description:         "The instance name of this Kiali installation. This instance name will be the prefix prepended to the names of all Kiali resources created by the operator and will be used to label those resources as belonging to this Kiali installation instance. You cannot change this instance name after a Kiali CR is created. If you attempt to change it, the operator will abort with an error. If you want to change it, you must first delete the original Kiali CR and create a new one. Note that this does not affect the name of the auto-generated signing key secret. If you do not supply a signing key, the operator will create one for you in a secret, but that secret will always be named 'kiali-signing-key' and shared across all Kiali instances in the same deployment namespace. If you want a different signing key secret, you are free to create your own and tell the operator about it via 'login_token.signing_key'. See the docs on that setting for more details. Note also that if you are setting this value, you may also want to change the 'installation_tag' setting, but this is not required.",
								MarkdownDescription: "The instance name of this Kiali installation. This instance name will be the prefix prepended to the names of all Kiali resources created by the operator and will be used to label those resources as belonging to this Kiali installation instance. You cannot change this instance name after a Kiali CR is created. If you attempt to change it, the operator will abort with an error. If you want to change it, you must first delete the original Kiali CR and create a new one. Note that this does not affect the name of the auto-generated signing key secret. If you do not supply a signing key, the operator will create one for you in a secret, but that secret will always be named 'kiali-signing-key' and shared across all Kiali instances in the same deployment namespace. If you want a different signing key secret, you are free to create your own and tell the operator about it via 'login_token.signing_key'. See the docs on that setting for more details. Note also that if you are setting this value, you may also want to change the 'installation_tag' setting, but this is not required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logger": schema.SingleNestedAttribute{
								Description:         "Configures the logger that emits messages to the Kiali server pod logs.",
								MarkdownDescription: "Configures the logger that emits messages to the Kiali server pod logs.",
								Attributes: map[string]schema.Attribute{
									"log_format": schema.StringAttribute{
										Description:         "Indicates if the logs should be written with one log message per line or using a JSON format. Must be one of: 'text' or 'json'.",
										MarkdownDescription: "Indicates if the logs should be written with one log message per line or using a JSON format. Must be one of: 'text' or 'json'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("text", "json"),
										},
									},

									"log_level": schema.StringAttribute{
										Description:         "The lowest priority of messages to log. Must be one of: 'trace', 'debug', 'info', 'warn', 'error', or 'fatal'.",
										MarkdownDescription: "The lowest priority of messages to log. Must be one of: 'trace', 'debug', 'info', 'warn', 'error', or 'fatal'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("trace", "debug", "info", "warn", "error", "fatal"),
										},
									},

									"sampler_rate": schema.StringAttribute{
										Description:         "With this setting every sampler_rate-th message will be logged. By default, every message is logged. As an example, setting this to ''2'' means every other message will be logged. The value of this setting is a string but must be parsable as an integer.",
										MarkdownDescription: "With this setting every sampler_rate-th message will be logged. By default, every message is logged. As an example, setting this to ''2'' means every other message will be logged. The value of this setting is a string but must be parsable as an integer.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9]+$`), ""),
										},
									},

									"time_field_format": schema.StringAttribute{
										Description:         "The log message timestamp format. This supports a golang time format (see https://golang.org/pkg/time/)",
										MarkdownDescription: "The log message timestamp format. This supports a golang time format (see https://golang.org/pkg/time/)",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": schema.StringAttribute{
								Description:         "The namespace into which Kiali is to be installed. If this is empty or not defined, the default will be the namespace where the Kiali CR is located.",
								MarkdownDescription: "The namespace into which Kiali is to be installed. If this is empty or not defined, the default will be the namespace where the Kiali CR is located.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "A set of node labels that dictate onto which node the Kiali pod will be deployed.",
								MarkdownDescription: "A set of node labels that dictate onto which node the Kiali pod will be deployed.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_annotations": schema.MapAttribute{
								Description:         "Custom annotations to be created on the Kiali pod.",
								MarkdownDescription: "Custom annotations to be created on the Kiali pod.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_labels": schema.MapAttribute{
								Description:         "Custom labels to be created on the Kiali pod. An example use for this setting is to inject an Istio sidecar such as, ''' sidecar.istio.io/inject: 'true' ''' ",
								MarkdownDescription: "Custom labels to be created on the Kiali pod. An example use for this setting is to inject an Istio sidecar such as, ''' sidecar.istio.io/inject: 'true' ''' ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"priority_class_name": schema.StringAttribute{
								Description:         "The priorityClassName used to assign the priority of the Kiali pod.",
								MarkdownDescription: "The priorityClassName used to assign the priority of the Kiali pod.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "The replica count for the Kiail deployment. If 'deployment.hpa' is specified, this setting is ignored.",
								MarkdownDescription: "The replica count for the Kiail deployment. If 'deployment.hpa' is specified, this setting is ignored.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resources": schema.MapAttribute{
								Description:         "Defines compute resources that are to be given to the Kiali pod's container. The value is a dict as defined by Kubernetes. See the Kubernetes documentation (https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container). If you set this to an empty dict ('{}') then no resources will be defined in the Deployment. If you do not set this at all, the default is, ''' requests: cpu: '10m' memory: '64Mi' limits: memory: '1Gi' ''' ",
								MarkdownDescription: "Defines compute resources that are to be given to the Kiali pod's container. The value is a dict as defined by Kubernetes. See the Kubernetes documentation (https://kubernetes.io/docs/concepts/configuration/manage-compute-resources-container). If you set this to an empty dict ('{}') then no resources will be defined in the Deployment. If you do not set this at all, the default is, ''' requests: cpu: '10m' memory: '64Mi' limits: memory: '1Gi' ''' ",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "The name of a secret used by the Kiali. This secret is optionally used when configuring the OpenID authentication strategy. Consult the OpenID docs for more information at https://kiali.io/docs/configuration/authentication/openid/",
								MarkdownDescription: "The name of a secret used by the Kiali. This secret is optionally used when configuring the OpenID authentication strategy. Consult the OpenID docs for more information at https://kiali.io/docs/configuration/authentication/openid/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"security_context": schema.MapAttribute{
								Description:         "Custom security context to be placed on the server container. The entire security context on the container will be the value of this setting if the operator is configured to allow it. Note that, as a security measure, a cluster admin may have configured the Kiali operator to not allow portions of this override setting - in this case you can specify additional security context settings but you cannot replace existing, default ones.",
								MarkdownDescription: "Custom security context to be placed on the server container. The entire security context on the container will be the value of this setting if the operator is configured to allow it. Note that, as a security measure, a cluster admin may have configured the Kiali operator to not allow portions of this override setting - in this case you can specify additional security context settings but you cannot replace existing, default ones.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_annotations": schema.MapAttribute{
								Description:         "Custom annotations to be created on the Kiali Service resource.",
								MarkdownDescription: "Custom annotations to be created on the Kiali Service resource.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_type": schema.StringAttribute{
								Description:         "The Kiali service type. Kubernetes determines what values are valid. Common values are 'NodePort', 'ClusterIP', and 'LoadBalancer'.",
								MarkdownDescription: "The Kiali service type. Kubernetes determines what values are valid. Common values are 'NodePort', 'ClusterIP', and 'LoadBalancer'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListAttribute{
								Description:         "A list of tolerations which declare which node taints Kiali can tolerate. See the Kubernetes documentation on Taints and Tolerations for more details.",
								MarkdownDescription: "A list of tolerations which declare which node taints Kiali can tolerate. See the Kubernetes documentation on Taints and Tolerations for more details.",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version_label": schema.StringAttribute{
								Description:         "Kiali resources will be assigned a 'version' label when they are deployed. This setting determines what value those 'version' labels will have. When empty, its default will be determined as follows, * If 'deployment.image_version' is 'latest', 'version_label' will be fixed to 'master'. * If 'deployment.image_version' is 'lastrelease', 'version_label' will be fixed to the last Kiali release version string. * If 'deployment.image_version' is anything else, 'version_label' will be that value, too. ",
								MarkdownDescription: "Kiali resources will be assigned a 'version' label when they are deployed. This setting determines what value those 'version' labels will have. When empty, its default will be determined as follows, * If 'deployment.image_version' is 'latest', 'version_label' will be fixed to 'master'. * If 'deployment.image_version' is 'lastrelease', 'version_label' will be fixed to the last Kiali release version string. * If 'deployment.image_version' is anything else, 'version_label' will be that value, too. ",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"view_only_mode": schema.BoolAttribute{
								Description:         "When true, Kiali will be in 'view only' mode, allowing the user to view and retrieve management and monitoring data for the service mesh, but not allow the user to modify the service mesh.",
								MarkdownDescription: "When true, Kiali will be in 'view only' mode, allowing the user to view and retrieve management and monitoring data for the service mesh, but not allow the user to modify the service mesh.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"extensions": schema.ListNestedAttribute{
						Description:         "Defines third-party extensions whose metrics can be integrated into the Kiali traffic graph. ",
						MarkdownDescription: "Defines third-party extensions whose metrics can be integrated into the Kiali traffic graph. ",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description:         "Determines if the Kiali traffic graph should incorporate the extension's metrics.",
									MarkdownDescription: "Determines if the Kiali traffic graph should incorporate the extension's metrics.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "The name that is used to identify the metric time series for the extension.",
									MarkdownDescription: "The name that is used to identify the metric time series for the extension.",
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

					"external_services": schema.SingleNestedAttribute{
						Description:         "These external service configuration settings define how to connect to the external services like Prometheus, Grafana, and Jaeger. Regarding sensitive values in the external_services 'auth' sections: Some external services configured below support an 'auth' sub-section in order to tell Kiali how it should authenticate with the external services. Credentials used to authenticate Kiali to those external services can be defined in the 'auth.password' and 'auth.token' values within the 'auth' sub-section. Because these are sensitive values, you may not want to declare the actual credentials here in the Kiali CR. In this case, you may store the actual password or token string in a Kubernetes secret. If you do, you need to set the 'auth.password' or 'auth.token' to a value in the format 'secret:<secretName>:<secretKey>' where '<secretName>' is the name of the secret object that Kiali can access, and '<secretKey>' is the name of the key within the named secret that contains the actual password or token string. For example, if Grafana requires a password, you can store that password in a secret named 'myGrafanaCredentials' in a key named 'myGrafanaPw'. In this case, you would set 'external_services.grafana.auth.password' to 'secret:myGrafanaCredentials:myGrafanaPw'. ",
						MarkdownDescription: "These external service configuration settings define how to connect to the external services like Prometheus, Grafana, and Jaeger. Regarding sensitive values in the external_services 'auth' sections: Some external services configured below support an 'auth' sub-section in order to tell Kiali how it should authenticate with the external services. Credentials used to authenticate Kiali to those external services can be defined in the 'auth.password' and 'auth.token' values within the 'auth' sub-section. Because these are sensitive values, you may not want to declare the actual credentials here in the Kiali CR. In this case, you may store the actual password or token string in a Kubernetes secret. If you do, you need to set the 'auth.password' or 'auth.token' to a value in the format 'secret:<secretName>:<secretKey>' where '<secretName>' is the name of the secret object that Kiali can access, and '<secretKey>' is the name of the key within the named secret that contains the actual password or token string. For example, if Grafana requires a password, you can store that password in a secret named 'myGrafanaCredentials' in a key named 'myGrafanaPw'. In this case, you would set 'external_services.grafana.auth.password' to 'secret:myGrafanaCredentials:myGrafanaPw'. ",
						Attributes: map[string]schema.Attribute{
							"custom_dashboards": schema.SingleNestedAttribute{
								Description:         "Settings for enabling and discovering custom dashboards.",
								MarkdownDescription: "Settings for enabling and discovering custom dashboards.",
								Attributes: map[string]schema.Attribute{
									"discovery_auto_threshold": schema.Int64Attribute{
										Description:         "Threshold of the number of pods, for a given Application or Workload, above which dashboards discovery will be skipped. This setting only takes effect when 'discovery_enabled' is set to 'auto'.",
										MarkdownDescription: "Threshold of the number of pods, for a given Application or Workload, above which dashboards discovery will be skipped. This setting only takes effect when 'discovery_enabled' is set to 'auto'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"discovery_enabled": schema.StringAttribute{
										Description:         "Enable, disable or set 'auto' mode to the dashboards discovery process. If set to 'true', Kiali will always try to discover dashboards based on metrics. Note that this can generate performance penalties while discovering dashboards for workloads having many pods (thus many metrics). When set to 'auto', Kiali will skip dashboards discovery for workloads with more than a configured threshold of pods (see 'discovery_auto_threshold'). When discovery is disabled or auto/skipped, it is still possible to tie workloads with dashboards through annotations on pods (refer to the doc https://kiali.io/docs/configuration/custom-dashboard/#pod-annotations). Value must be one of: 'true', 'false', 'auto'.",
										MarkdownDescription: "Enable, disable or set 'auto' mode to the dashboards discovery process. If set to 'true', Kiali will always try to discover dashboards based on metrics. Note that this can generate performance penalties while discovering dashboards for workloads having many pods (thus many metrics). When set to 'auto', Kiali will skip dashboards discovery for workloads with more than a configured threshold of pods (see 'discovery_auto_threshold'). When discovery is disabled or auto/skipped, it is still possible to tie workloads with dashboards through annotations on pods (refer to the doc https://kiali.io/docs/configuration/custom-dashboard/#pod-annotations). Value must be one of: 'true', 'false', 'auto'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "Enable or disable custom dashboards, including the dashboards discovery process.",
										MarkdownDescription: "Enable or disable custom dashboards, including the dashboards discovery process.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"is_core": schema.BoolAttribute{
										Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace_label": schema.StringAttribute{
										Description:         "The Prometheus label name used for identifying namespaces in metrics for custom dashboards. The default is 'namespace' but you may want to use 'kubernetes_namespace' depending on your Prometheus configuration.",
										MarkdownDescription: "The Prometheus label name used for identifying namespaces in metrics for custom dashboards. The default is 'namespace' but you may want to use 'kubernetes_namespace' depending on your Prometheus configuration.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"prometheus": schema.SingleNestedAttribute{
										Description:         "The Prometheus configuration defined here refers to the Prometheus instance that is dedicated to fetching metrics for custom dashboards. This means you can obtain these metrics for the custom dashboards from a Prometheus instance that is different from the one that Istio uses. If this section is omitted, the same Prometheus that is used to obtain the Istio metrics will also be used for retrieving custom dashboard metrics.",
										MarkdownDescription: "The Prometheus configuration defined here refers to the Prometheus instance that is dedicated to fetching metrics for custom dashboards. This means you can obtain these metrics for the custom dashboards from a Prometheus instance that is different from the one that Istio uses. If this section is omitted, the same Prometheus that is used to obtain the Istio metrics will also be used for retrieving custom dashboard metrics.",
										Attributes: map[string]schema.Attribute{
											"auth": schema.SingleNestedAttribute{
												Description:         "Settings used to authenticate with the Prometheus instance.",
												MarkdownDescription: "Settings used to authenticate with the Prometheus instance.",
												Attributes: map[string]schema.Attribute{
													"ca_file": schema.StringAttribute{
														Description:         "The certificate authority file to use when accessing Prometheus using https. An empty string means no extra certificate authority file is used.",
														MarkdownDescription: "The certificate authority file to use when accessing Prometheus using https. An empty string means no extra certificate authority file is used.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"insecure_skip_verify": schema.BoolAttribute{
														Description:         "Set true to skip verifying certificate validity when Kiali contacts Prometheus over https.",
														MarkdownDescription: "Set true to skip verifying certificate validity when Kiali contacts Prometheus over https.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"password": schema.StringAttribute{
														Description:         "Password to be used when making requests to Prometheus, for basic authentication. May refer to a secret.",
														MarkdownDescription: "Password to be used when making requests to Prometheus, for basic authentication. May refer to a secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"token": schema.StringAttribute{
														Description:         "Token / API key to access Prometheus, for token-based authentication. May refer to a secret.",
														MarkdownDescription: "Token / API key to access Prometheus, for token-based authentication. May refer to a secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"type": schema.StringAttribute{
														Description:         "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Prometheus server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
														MarkdownDescription: "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Prometheus server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"use_kiali_token": schema.BoolAttribute{
														Description:         "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Prometheus (in this case, 'auth.token' config is ignored).",
														MarkdownDescription: "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Prometheus (in this case, 'auth.token' config is ignored).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"username": schema.StringAttribute{
														Description:         "Username to be used when making requests to Prometheus with 'basic' authentication. May refer to a secret.",
														MarkdownDescription: "Username to be used when making requests to Prometheus with 'basic' authentication. May refer to a secret.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"cache_duration": schema.Int64Attribute{
												Description:         "Prometheus caching duration expressed in seconds.",
												MarkdownDescription: "Prometheus caching duration expressed in seconds.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cache_enabled": schema.BoolAttribute{
												Description:         "Enable/disable Prometheus caching used for Health services.",
												MarkdownDescription: "Enable/disable Prometheus caching used for Health services.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cache_expiration": schema.Int64Attribute{
												Description:         "Prometheus caching expiration expressed in seconds.",
												MarkdownDescription: "Prometheus caching expiration expressed in seconds.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"custom_headers": schema.MapAttribute{
												Description:         "A set of name/value settings that will be passed as headers when requests are sent to Prometheus.",
												MarkdownDescription: "A set of name/value settings that will be passed as headers when requests are sent to Prometheus.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"health_check_url": schema.StringAttribute{
												Description:         "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",
												MarkdownDescription: "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"is_core": schema.BoolAttribute{
												Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
												MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"query_scope": schema.MapAttribute{
												Description:         "A set of labelName/labelValue settings applied to every Prometheus query. Used to narrow unified metrics to only those scoped to the Kiali instance.",
												MarkdownDescription: "A set of labelName/labelValue settings applied to every Prometheus query. Used to narrow unified metrics to only those scoped to the Kiali instance.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"thanos_proxy": schema.SingleNestedAttribute{
												Description:         "Define this section if Prometheus is to be queried through a Thanos proxy. Kiali will still use the 'url' setting to query for Prometheus metrics so make sure that is set appropriately.",
												MarkdownDescription: "Define this section if Prometheus is to be queried through a Thanos proxy. Kiali will still use the 'url' setting to query for Prometheus metrics so make sure that is set appropriately.",
												Attributes: map[string]schema.Attribute{
													"enabled": schema.BoolAttribute{
														Description:         "Set to true when a Thanos proxy is in front of Prometheus.",
														MarkdownDescription: "Set to true when a Thanos proxy is in front of Prometheus.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"retention_period": schema.StringAttribute{
														Description:         "Thanos Retention period value expresed as a string.",
														MarkdownDescription: "Thanos Retention period value expresed as a string.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"scrape_interval": schema.StringAttribute{
														Description:         "Thanos Scrape interval value expresed as a string.",
														MarkdownDescription: "Thanos Scrape interval value expresed as a string.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"url": schema.StringAttribute{
												Description:         "The URL used to query the Prometheus Server. This URL must be accessible from the Kiali pod. If empty, the default will assume Prometheus is in the Istio control plane namespace; e.g. 'http://prometheus.<istio_namespace>:9090'.",
												MarkdownDescription: "The URL used to query the Prometheus Server. This URL must be accessible from the Kiali pod. If empty, the default will assume Prometheus is in the Istio control plane namespace; e.g. 'http://prometheus.<istio_namespace>:9090'.",
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

							"grafana": schema.SingleNestedAttribute{
								Description:         "Configuration used to access the Grafana dashboards.",
								MarkdownDescription: "Configuration used to access the Grafana dashboards.",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Settings used to authenticate with the Grafana instance.",
										MarkdownDescription: "Settings used to authenticate with the Grafana instance.",
										Attributes: map[string]schema.Attribute{
											"ca_file": schema.StringAttribute{
												Description:         "The certificate authority file to use when accessing Grafana using https. An empty string means no extra certificate authority file is used.",
												MarkdownDescription: "The certificate authority file to use when accessing Grafana using https. An empty string means no extra certificate authority file is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "Set true to skip verifying certificate validity when Kiali contacts Grafana over https.",
												MarkdownDescription: "Set true to skip verifying certificate validity when Kiali contacts Grafana over https.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"password": schema.StringAttribute{
												Description:         "Password to be used when making requests to Grafana, for basic authentication. May refer to a secret.",
												MarkdownDescription: "Password to be used when making requests to Grafana, for basic authentication. May refer to a secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"token": schema.StringAttribute{
												Description:         "Token / API key to access Grafana, for token-based authentication. May refer to a secret.",
												MarkdownDescription: "Token / API key to access Grafana, for token-based authentication. May refer to a secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Grafana server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
												MarkdownDescription: "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Grafana server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"use_kiali_token": schema.BoolAttribute{
												Description:         "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Grafana (in this case, 'auth.token' config is ignored).",
												MarkdownDescription: "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Grafana (in this case, 'auth.token' config is ignored).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"username": schema.StringAttribute{
												Description:         "Username to be used when making requests to Grafana with 'basic' authentication. May refer to a secret.",
												MarkdownDescription: "Username to be used when making requests to Grafana with 'basic' authentication. May refer to a secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"dashboards": schema.ListNestedAttribute{
										Description:         "A list of Grafana dashboards that Kiali can link to.",
										MarkdownDescription: "A list of Grafana dashboards that Kiali can link to.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The name of the Grafana dashboard.",
													MarkdownDescription: "The name of the Grafana dashboard.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"variables": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"app": schema.StringAttribute{
															Description:         "The name of a variable that holds the app name, if used in that dashboard (else it must be omitted).",
															MarkdownDescription: "The name of a variable that holds the app name, if used in that dashboard (else it must be omitted).",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "The name of a variable that holds the namespace, if used in that dashboard (else it must be omitted).",
															MarkdownDescription: "The name of a variable that holds the namespace, if used in that dashboard (else it must be omitted).",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"service": schema.StringAttribute{
															Description:         "The name of a variable that holds the service name, if used in that dashboard (else it must be omitted).",
															MarkdownDescription: "The name of a variable that holds the service name, if used in that dashboard (else it must be omitted).",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"workload": schema.StringAttribute{
															Description:         "The name of a variable that holds the workload name, if used in that dashboard (else it must be omitted).",
															MarkdownDescription: "The name of a variable that holds the workload name, if used in that dashboard (else it must be omitted).",
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

									"enabled": schema.BoolAttribute{
										Description:         "When true, Grafana support will be enabled in Kiali.",
										MarkdownDescription: "When true, Grafana support will be enabled in Kiali.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_url": schema.StringAttribute{
										Description:         "The URL that the Kiali UI uses when displaying Grafana links to the user. This URL must be accessible to clients external to the cluster (e.g. a browser) in order for the integration to work properly. If empty, an attempt to auto-discover it is made. This URL can contain query parameters if needed, such as '?orgId=1'.",
										MarkdownDescription: "The URL that the Kiali UI uses when displaying Grafana links to the user. This URL must be accessible to clients external to the cluster (e.g. a browser) in order for the integration to work properly. If empty, an attempt to auto-discover it is made. This URL can contain query parameters if needed, such as '?orgId=1'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"health_check_url": schema.StringAttribute{
										Description:         "Used in the Components health feature. This is the URL which Kiali will ping to determine whether the component is reachable or not. It defaults to 'internal_url' when not provided.",
										MarkdownDescription: "Used in the Components health feature. This is the URL which Kiali will ping to determine whether the component is reachable or not. It defaults to 'internal_url' when not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"internal_url": schema.StringAttribute{
										Description:         "The URL used by Kiali to perform requests and queries to Grafana. An example would be 'http://grafana.istio-system:3000'. This URL can contain query parameters if needed, such as '?orgId=1'. If not defined, it will default to 'http://grafana.<istio_namespace>:3000'.",
										MarkdownDescription: "The URL used by Kiali to perform requests and queries to Grafana. An example would be 'http://grafana.istio-system:3000'. This URL can contain query parameters if needed, such as '?orgId=1'. If not defined, it will default to 'http://grafana.<istio_namespace>:3000'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"is_core": schema.BoolAttribute{
										Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"istio": schema.SingleNestedAttribute{
								Description:         "Istio configuration that Kiali needs to know about in order to observe the mesh.",
								MarkdownDescription: "Istio configuration that Kiali needs to know about in order to observe the mesh.",
								Attributes: map[string]schema.Attribute{
									"component_status": schema.SingleNestedAttribute{
										Description:         "Istio components whose status will be monitored by Kiali.",
										MarkdownDescription: "Istio components whose status will be monitored by Kiali.",
										Attributes: map[string]schema.Attribute{
											"components": schema.ListNestedAttribute{
												Description:         "A specific Istio component whose status will be monitored by Kiali.",
												MarkdownDescription: "A specific Istio component whose status will be monitored by Kiali.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"app_label": schema.StringAttribute{
															Description:         "Istio component pod app label.",
															MarkdownDescription: "Istio component pod app label.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"is_core": schema.BoolAttribute{
															Description:         "Whether the component is to be considered a core component for your deployment.",
															MarkdownDescription: "Whether the component is to be considered a core component for your deployment.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"is_proxy": schema.BoolAttribute{
															Description:         "Whether the component is a native Envoy proxy.",
															MarkdownDescription: "Whether the component is a native Envoy proxy.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "The namespace where the component is installed. It defaults to the Istio control plane namespace (e.g. 'istio_namespace') setting. Note that the Istio documentation suggests you install the ingress and egress to different namespaces, so you most likely will want to explicitly set this namespace value for the ingress and egress components.",
															MarkdownDescription: "The namespace where the component is installed. It defaults to the Istio control plane namespace (e.g. 'istio_namespace') setting. Note that the Istio documentation suggests you install the ingress and egress to different namespaces, so you most likely will want to explicitly set this namespace value for the ingress and egress components.",
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

											"enabled": schema.BoolAttribute{
												Description:         "Determines if Istio component statuses will be displayed in the Kiali masthead indicator.",
												MarkdownDescription: "Determines if Istio component statuses will be displayed in the Kiali masthead indicator.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"config_map_name": schema.StringAttribute{
										Description:         "The name of the istio control plane config map.",
										MarkdownDescription: "The name of the istio control plane config map.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"egress_gateway_namespace": schema.StringAttribute{
										Description:         "The namespace where Istio EgressGateway component is read for a status check. When left empty, then 'istio_namespace' value is used.",
										MarkdownDescription: "The namespace where Istio EgressGateway component is read for a status check. When left empty, then 'istio_namespace' value is used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"envoy_admin_local_port": schema.Int64Attribute{
										Description:         "The port which kiali will open to fetch envoy config data information.",
										MarkdownDescription: "The port which kiali will open to fetch envoy config data information.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"gateway_api_classes": schema.ListNestedAttribute{
										Description:         "A list declaring all the Gateways used in Istio. If left empty or undefined, the default is a single list item whose name is 'Istio' and class_name is 'istio'.",
										MarkdownDescription: "A list declaring all the Gateways used in Istio. If left empty or undefined, the default is a single list item whose name is 'Istio' and class_name is 'istio'.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"class_name": schema.StringAttribute{
													Description:         "The name of the GatewayClass.",
													MarkdownDescription: "The name of the GatewayClass.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "The name of the Gateway API implementation.",
													MarkdownDescription: "The name of the Gateway API implementation.",
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

									"ingress_gateway_namespace": schema.StringAttribute{
										Description:         "The namespace where Istio IngressGateway component is read for a status check. When left empty, then 'istio_namespace' value is used.",
										MarkdownDescription: "The namespace where Istio IngressGateway component is read for a status check. When left empty, then 'istio_namespace' value is used.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"istio_api_enabled": schema.BoolAttribute{
										Description:         "Indicates if Kiali has access to istiod. true by default.",
										MarkdownDescription: "Indicates if Kiali has access to istiod. true by default.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"istio_canary_revision": schema.SingleNestedAttribute{
										Description:         "These values are used in Canary upgrade/downgrade functionality when 'istio_upgrade_action' is true.",
										MarkdownDescription: "These values are used in Canary upgrade/downgrade functionality when 'istio_upgrade_action' is true.",
										Attributes: map[string]schema.Attribute{
											"current": schema.StringAttribute{
												Description:         "The currently installed Istio revision.",
												MarkdownDescription: "The currently installed Istio revision.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"upgrade": schema.StringAttribute{
												Description:         "The installed Istio canary revision to upgrade to.",
												MarkdownDescription: "The installed Istio canary revision to upgrade to.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"istio_identity_domain": schema.StringAttribute{
										Description:         "The annotation used by Istio to identify domains.",
										MarkdownDescription: "The annotation used by Istio to identify domains.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"istio_injection_annotation": schema.StringAttribute{
										Description:         "The name of the field that annotates a workload to indicate a sidecar should be automatically injected by Istio. This is the name of a Kubernetes annotation. Note that some Istio implementations also support labels by the same name. In other words, if a workload has a Kubernetes label with this name, that may also trigger automatic sidecar injection.",
										MarkdownDescription: "The name of the field that annotates a workload to indicate a sidecar should be automatically injected by Istio. This is the name of a Kubernetes annotation. Note that some Istio implementations also support labels by the same name. In other words, if a workload has a Kubernetes label with this name, that may also trigger automatic sidecar injection.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"istio_sidecar_annotation": schema.StringAttribute{
										Description:         "The pod annotation used by Istio to identify the sidecar.",
										MarkdownDescription: "The pod annotation used by Istio to identify the sidecar.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"istio_sidecar_injector_config_map_name": schema.StringAttribute{
										Description:         "The name of the istio-sidecar-injector config map.",
										MarkdownDescription: "The name of the istio-sidecar-injector config map.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"istiod_deployment_name": schema.StringAttribute{
										Description:         "The name of the istiod deployment.",
										MarkdownDescription: "The name of the istiod deployment.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"istiod_pod_monitoring_port": schema.Int64Attribute{
										Description:         "The monitoring port of the IstioD pod (not the Service).",
										MarkdownDescription: "The monitoring port of the IstioD pod (not the Service).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"root_namespace": schema.StringAttribute{
										Description:         "The namespace to treat as the administrative root namespace for Istio configuration.",
										MarkdownDescription: "The namespace to treat as the administrative root namespace for Istio configuration.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"url_service_version": schema.StringAttribute{
										Description:         "The Istio service used to determine the Istio version. If empty, assumes the URL for the well-known Istio version endpoint.",
										MarkdownDescription: "The Istio service used to determine the Istio version. If empty, assumes the URL for the well-known Istio version endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"prometheus": schema.SingleNestedAttribute{
								Description:         "The Prometheus configuration defined here refers to the Prometheus instance that is used by Istio to store its telemetry.",
								MarkdownDescription: "The Prometheus configuration defined here refers to the Prometheus instance that is used by Istio to store its telemetry.",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Settings used to authenticate with the Prometheus instance.",
										MarkdownDescription: "Settings used to authenticate with the Prometheus instance.",
										Attributes: map[string]schema.Attribute{
											"ca_file": schema.StringAttribute{
												Description:         "The certificate authority file to use when accessing Prometheus using https. An empty string means no extra certificate authority file is used.",
												MarkdownDescription: "The certificate authority file to use when accessing Prometheus using https. An empty string means no extra certificate authority file is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "Set true to skip verifying certificate validity when Kiali contacts Prometheus over https.",
												MarkdownDescription: "Set true to skip verifying certificate validity when Kiali contacts Prometheus over https.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"password": schema.StringAttribute{
												Description:         "Password to be used when making requests to Prometheus, for basic authentication. May refer to a secret.",
												MarkdownDescription: "Password to be used when making requests to Prometheus, for basic authentication. May refer to a secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"token": schema.StringAttribute{
												Description:         "Token / API key to access Prometheus, for token-based authentication. May refer to a secret.",
												MarkdownDescription: "Token / API key to access Prometheus, for token-based authentication. May refer to a secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Prometheus server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
												MarkdownDescription: "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Prometheus server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"use_kiali_token": schema.BoolAttribute{
												Description:         "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Prometheus (in this case, 'auth.token' config is ignored).",
												MarkdownDescription: "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to Prometheus (in this case, 'auth.token' config is ignored).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"username": schema.StringAttribute{
												Description:         "Username to be used when making requests to Prometheus with 'basic' authentication. May refer to a secret.",
												MarkdownDescription: "Username to be used when making requests to Prometheus with 'basic' authentication. May refer to a secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"cache_duration": schema.Int64Attribute{
										Description:         "Prometheus caching duration expressed in seconds.",
										MarkdownDescription: "Prometheus caching duration expressed in seconds.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cache_enabled": schema.BoolAttribute{
										Description:         "Enable/disable Prometheus caching used for Health services.",
										MarkdownDescription: "Enable/disable Prometheus caching used for Health services.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cache_expiration": schema.Int64Attribute{
										Description:         "Prometheus caching expiration expressed in seconds.",
										MarkdownDescription: "Prometheus caching expiration expressed in seconds.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"custom_headers": schema.MapAttribute{
										Description:         "A set of name/value settings that will be passed as headers when requests are sent to Prometheus.",
										MarkdownDescription: "A set of name/value settings that will be passed as headers when requests are sent to Prometheus.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"health_check_url": schema.StringAttribute{
										Description:         "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",
										MarkdownDescription: "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"is_core": schema.BoolAttribute{
										Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query_scope": schema.MapAttribute{
										Description:         "A set of labelName/labelValue settings applied to every Prometheus query. Used to narrow unified metrics to only those scoped to the Kiali instance.",
										MarkdownDescription: "A set of labelName/labelValue settings applied to every Prometheus query. Used to narrow unified metrics to only those scoped to the Kiali instance.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"thanos_proxy": schema.SingleNestedAttribute{
										Description:         "Define this section if Prometheus is to be queried through a Thanos proxy. Kiali will still use the 'url' setting to query for Prometheus metrics so make sure that is set appropriately.",
										MarkdownDescription: "Define this section if Prometheus is to be queried through a Thanos proxy. Kiali will still use the 'url' setting to query for Prometheus metrics so make sure that is set appropriately.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Set to true when a Thanos proxy is in front of Prometheus.",
												MarkdownDescription: "Set to true when a Thanos proxy is in front of Prometheus.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"retention_period": schema.StringAttribute{
												Description:         "Thanos Retention period value expresed as a string.",
												MarkdownDescription: "Thanos Retention period value expresed as a string.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"scrape_interval": schema.StringAttribute{
												Description:         "Thanos Scrape interval value expresed as a string.",
												MarkdownDescription: "Thanos Scrape interval value expresed as a string.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"url": schema.StringAttribute{
										Description:         "The URL used to query the Prometheus Server. This URL must be accessible from the Kiali pod. If empty, the default will assume Prometheus is in the Istio control plane namespace; e.g. 'http://prometheus.<istio_namespace>:9090'.",
										MarkdownDescription: "The URL used to query the Prometheus Server. This URL must be accessible from the Kiali pod. If empty, the default will assume Prometheus is in the Istio control plane namespace; e.g. 'http://prometheus.<istio_namespace>:9090'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tracing": schema.SingleNestedAttribute{
								Description:         "Configuration used to access the Tracing (Jaeger or Tempo) dashboards.",
								MarkdownDescription: "Configuration used to access the Tracing (Jaeger or Tempo) dashboards.",
								Attributes: map[string]schema.Attribute{
									"auth": schema.SingleNestedAttribute{
										Description:         "Settings used to authenticate with the Tracing server instance.",
										MarkdownDescription: "Settings used to authenticate with the Tracing server instance.",
										Attributes: map[string]schema.Attribute{
											"ca_file": schema.StringAttribute{
												Description:         "The certificate authority file to use when accessing the Tracing server using https. An empty string means no extra certificate authority file is used.",
												MarkdownDescription: "The certificate authority file to use when accessing the Tracing server using https. An empty string means no extra certificate authority file is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure_skip_verify": schema.BoolAttribute{
												Description:         "Set true to skip verifying certificate validity when Kiali contacts the Tracing server over https.",
												MarkdownDescription: "Set true to skip verifying certificate validity when Kiali contacts the Tracing server over https.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"password": schema.StringAttribute{
												Description:         "Password to be used when making requests to the Tracing server, for basic authentication. May refer to a secret.",
												MarkdownDescription: "Password to be used when making requests to the Tracing server, for basic authentication. May refer to a secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"token": schema.StringAttribute{
												Description:         "Token / API key to access the Tracing server, for token-based authentication. May refer to a secret.",
												MarkdownDescription: "Token / API key to access the Tracing server, for token-based authentication. May refer to a secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"type": schema.StringAttribute{
												Description:         "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Tracing server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
												MarkdownDescription: "The type of authentication to use when contacting the server. Use 'bearer' to send the token to the Tracing server. Use 'basic' to connect with username and password credentials. Use 'none' to not use any authentication (this is the default).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"use_kiali_token": schema.BoolAttribute{
												Description:         "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to the Tracing server (in this case, 'auth.token' config is ignored).",
												MarkdownDescription: "When true and if 'auth.type' is 'bearer', Kiali Service Account token will be used for the API calls to the Tracing server (in this case, 'auth.token' config is ignored).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"username": schema.StringAttribute{
												Description:         "Username to be used when making requests to the Tracing server with 'basic' authentication. May refer to a secret.",
												MarkdownDescription: "Username to be used when making requests to the Tracing server with 'basic' authentication. May refer to a secret.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"custom_headers": schema.MapAttribute{
										Description:         "A set of name/value settings that will be passed as headers when requests are sent to the Tracing backend.",
										MarkdownDescription: "A set of name/value settings that will be passed as headers when requests are sent to the Tracing backend.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enabled": schema.BoolAttribute{
										Description:         "When true, connections to the Tracing server are enabled. 'internal_url' and/or 'external_url' need to be provided.",
										MarkdownDescription: "When true, connections to the Tracing server are enabled. 'internal_url' and/or 'external_url' need to be provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"external_url": schema.StringAttribute{
										Description:         "The URL that the Kiali UI uses when displaying Tracing UI links to the user. This URL must be accessible to clients external to the cluster (e.g. a browser) in order to generate valid links. If the tracing service is deployed with a QUERY_BASE_PATH set, set this URL like https://<hostname>/<QUERY_BASE_PATH>; for example, https://tracing-service:8080/jaeger",
										MarkdownDescription: "The URL that the Kiali UI uses when displaying Tracing UI links to the user. This URL must be accessible to clients external to the cluster (e.g. a browser) in order to generate valid links. If the tracing service is deployed with a QUERY_BASE_PATH set, set this URL like https://<hostname>/<QUERY_BASE_PATH>; for example, https://tracing-service:8080/jaeger",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"grpc_port": schema.Int64Attribute{
										Description:         "Set port number when 'use_grpc' is true and 'provider' is 'tempo'. By default is '9095'",
										MarkdownDescription: "Set port number when 'use_grpc' is true and 'provider' is 'tempo'. By default is '9095'",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"health_check_url": schema.StringAttribute{
										Description:         "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",
										MarkdownDescription: "Used in the Components health feature. This is the url which Kiali will ping to determine whether the component is reachable or not. It defaults to 'url' when not provided.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"internal_url": schema.StringAttribute{
										Description:         "The URL used by Kiali to perform requests and queries to the tracing backend which enables further integration between Kiali and the tracing server. When not provided, Kiali will only show external links using the 'external_url' setting. Note: Jaeger v1.20+ has separated ports for GRPC(16685) and HTTP(16686) requests. Make sure you use the appropriate port according to the 'use_grpc' value. Example: http://tracing.istio-system:16685",
										MarkdownDescription: "The URL used by Kiali to perform requests and queries to the tracing backend which enables further integration between Kiali and the tracing server. When not provided, Kiali will only show external links using the 'external_url' setting. Note: Jaeger v1.20+ has separated ports for GRPC(16685) and HTTP(16686) requests. Make sure you use the appropriate port according to the 'use_grpc' value. Example: http://tracing.istio-system:16685",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"is_core": schema.BoolAttribute{
										Description:         "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										MarkdownDescription: "Used in the Components health feature. When true, the unhealthy scenarios will be raised as errors. Otherwise, they will be raised as a warning.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"namespace_selector": schema.BoolAttribute{
										Description:         "Kiali use this boolean to find traces with a namespace selector : service.namespace.",
										MarkdownDescription: "Kiali use this boolean to find traces with a namespace selector : service.namespace.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"provider": schema.StringAttribute{
										Description:         "The trace provider to get the traces from. Value must be one of: 'jaeger' or 'tempo'.",
										MarkdownDescription: "The trace provider to get the traces from. Value must be one of: 'jaeger' or 'tempo'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query_scope": schema.MapAttribute{
										Description:         "A set of tagKey/tagValue settings applied to every Jaeger query. Used to narrow unified traces to only those scoped to the Kiali instance.",
										MarkdownDescription: "A set of tagKey/tagValue settings applied to every Jaeger query. Used to narrow unified traces to only those scoped to the Kiali instance.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"query_timeout": schema.Int64Attribute{
										Description:         "The amount of time in seconds Kiali will wait for a response from 'jaeger-query' service when fetching traces.",
										MarkdownDescription: "The amount of time in seconds Kiali will wait for a response from 'jaeger-query' service when fetching traces.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tempo_config": schema.SingleNestedAttribute{
										Description:         "Settings used to configure the access url to the Tempo Datasource in Grafana.",
										MarkdownDescription: "Settings used to configure the access url to the Tempo Datasource in Grafana.",
										Attributes: map[string]schema.Attribute{
											"datasource_uid": schema.StringAttribute{
												Description:         "The unique identifier (uid) of the Tempo datasource in Grafana.",
												MarkdownDescription: "The unique identifier (uid) of the Tempo datasource in Grafana.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"org_id": schema.StringAttribute{
												Description:         "The Id of the organization that the dashboard is in. Default to 1 (the first and default organization).",
												MarkdownDescription: "The Id of the organization that the dashboard is in. Default to 1 (the first and default organization).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"url_format": schema.StringAttribute{
												Description:         "The URL format for the external url. Can be 'jaeger' or 'grafana'. Default to 'grafana'. Grafana will need a Grafana url in the Grafana settings.",
												MarkdownDescription: "The URL format for the external url. Can be 'jaeger' or 'grafana'. Default to 'grafana'. Grafana will need a Grafana url in the Grafana settings.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"use_grpc": schema.BoolAttribute{
										Description:         "Set to true in order to enable GRPC connections between Kiali and Jaeger which will speed up the queries. In some setups you might not be able to use GRPC (e.g. if Jaeger is behind some reverse proxy that doesn't support it). If not specified, this will defalt to 'true'.",
										MarkdownDescription: "Set to true in order to enable GRPC connections between Kiali and Jaeger which will speed up the queries. In some setups you might not be able to use GRPC (e.g. if Jaeger is behind some reverse proxy that doesn't support it). If not specified, this will defalt to 'true'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"whitelist_istio_system": schema.ListAttribute{
										Description:         "Kiali will get the traces of these services found in the Istio control plane namespace.",
										MarkdownDescription: "Kiali will get the traces of these services found in the Istio control plane namespace.",
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

					"health_config": schema.SingleNestedAttribute{
						Description:         "This section defines what it means for nodes to be healthy. For more details, see https://kiali.io/docs/configuration/health/",
						MarkdownDescription: "This section defines what it means for nodes to be healthy. For more details, see https://kiali.io/docs/configuration/health/",
						Attributes: map[string]schema.Attribute{
							"rate": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "The type of resource that this configuration applies to. This is a regular expression.",
											MarkdownDescription: "The type of resource that this configuration applies to. This is a regular expression.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "The name of a resource that this configuration applies to. This is a regular expression.",
											MarkdownDescription: "The name of a resource that this configuration applies to. This is a regular expression.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"namespace": schema.StringAttribute{
											Description:         "The name of the namespace that this configuration applies to. This is a regular expression.",
											MarkdownDescription: "The name of the namespace that this configuration applies to. This is a regular expression.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tolerance": schema.ListNestedAttribute{
											Description:         "A list of tolerances for this configuration.",
											MarkdownDescription: "A list of tolerances for this configuration.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"code": schema.StringAttribute{
														Description:         "The status code that applies for this tolerance. This is a regular expression.",
														MarkdownDescription: "The status code that applies for this tolerance. This is a regular expression.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"degraded": schema.Int64Attribute{
														Description:         "Health will be considered degraded when the telemetry reaches this value (specified as an integer representing a percentage).",
														MarkdownDescription: "Health will be considered degraded when the telemetry reaches this value (specified as an integer representing a percentage).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"direction": schema.StringAttribute{
														Description:         "The direction that applies for this tolerance (e.g. inbound or outbound). This is a regular expression.",
														MarkdownDescription: "The direction that applies for this tolerance (e.g. inbound or outbound). This is a regular expression.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"failure": schema.Int64Attribute{
														Description:         "A failure status will be shown when the telemetry reaches this value (specified as an integer representing a percentage).",
														MarkdownDescription: "A failure status will be shown when the telemetry reaches this value (specified as an integer representing a percentage).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "The protocol that applies for this tolerance (e.g. grpc or http). This is a regular expression.",
														MarkdownDescription: "The protocol that applies for this tolerance (e.g. grpc or http). This is a regular expression.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"identity": schema.SingleNestedAttribute{
						Description:         "Settings that define the Kiali server identity.",
						MarkdownDescription: "Settings that define the Kiali server identity.",
						Attributes: map[string]schema.Attribute{
							"cert_file": schema.StringAttribute{
								Description:         "Certificate file used to identify the Kiali server. If set, you must go over https to access Kiali. The Kiali operator will set this if it deploys Kiali behind https. When left undefined, the operator will attempt to generate a cluster-specific cert file that provides https by default (today, this auto-generation of a cluster-specific cert is only supported on OpenShift). When set to an empty string, https will be disabled.",
								MarkdownDescription: "Certificate file used to identify the Kiali server. If set, you must go over https to access Kiali. The Kiali operator will set this if it deploys Kiali behind https. When left undefined, the operator will attempt to generate a cluster-specific cert file that provides https by default (today, this auto-generation of a cluster-specific cert is only supported on OpenShift). When set to an empty string, https will be disabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"private_key_file": schema.StringAttribute{
								Description:         "Private key file used to identify the Kiali server. If set, you must go over https to access Kiali. When left undefined, the Kiali operator will attempt to generate a cluster-specific private key file that provides https by default (today, this auto-generation of a cluster-specific private key is only supported on OpenShift). When set to an empty string, https will be disabled.",
								MarkdownDescription: "Private key file used to identify the Kiali server. If set, you must go over https to access Kiali. When left undefined, the Kiali operator will attempt to generate a cluster-specific private key file that provides https by default (today, this auto-generation of a cluster-specific private key is only supported on OpenShift). When set to an empty string, https will be disabled.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"installation_tag": schema.StringAttribute{
						Description:         "Tag used to identify a particular instance/installation of the Kiali server. This is merely a human-readable string that will be used within Kiali to help a user identify the Kiali being used (e.g. in the Kiali UI title bar). See 'deployment.instance_name' for the setting used to customize Kiali resource names that are created.",
						MarkdownDescription: "Tag used to identify a particular instance/installation of the Kiali server. This is merely a human-readable string that will be used within Kiali to help a user identify the Kiali being used (e.g. in the Kiali UI title bar). See 'deployment.instance_name' for the setting used to customize Kiali resource names that are created.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"istio_labels": schema.SingleNestedAttribute{
						Description:         "Defines specific labels used by Istio that Kiali needs to know about.",
						MarkdownDescription: "Defines specific labels used by Istio that Kiali needs to know about.",
						Attributes: map[string]schema.Attribute{
							"app_label_name": schema.StringAttribute{
								Description:         "The name of the label used to define what application a workload belongs to. This is typically something like 'app' or 'app.kubernetes.io/name'.",
								MarkdownDescription: "The name of the label used to define what application a workload belongs to. This is typically something like 'app' or 'app.kubernetes.io/name'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"egress_gateway_label": schema.StringAttribute{
								Description:         "The selector label for Egress Gateway workload. This is typically 'istio=egressgateway'.",
								MarkdownDescription: "The selector label for Egress Gateway workload. This is typically 'istio=egressgateway'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ingress_gateway_label": schema.StringAttribute{
								Description:         "The selector label for Ingress Gateway workload. This is typically 'istio=ingressgateway'.",
								MarkdownDescription: "The selector label for Ingress Gateway workload. This is typically 'istio=ingressgateway'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"injection_label_name": schema.StringAttribute{
								Description:         "The name of the label used to instruct Istio to automatically inject sidecar proxies when applications are deployed.",
								MarkdownDescription: "The name of the label used to instruct Istio to automatically inject sidecar proxies when applications are deployed.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"injection_label_rev": schema.StringAttribute{
								Description:         "The label used to identify the Istio revision.",
								MarkdownDescription: "The label used to identify the Istio revision.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version_label_name": schema.StringAttribute{
								Description:         "The name of the label used to define what version of the application a workload belongs to. This is typically something like 'version' or 'app.kubernetes.io/version'.",
								MarkdownDescription: "The name of the label used to define what version of the application a workload belongs to. This is typically something like 'version' or 'app.kubernetes.io/version'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"istio_namespace": schema.StringAttribute{
						Description:         "The namespace where Istio is installed. If left empty, it is assumed to be the same namespace as where Kiali is installed (i.e. 'deployment.namespace').",
						MarkdownDescription: "The namespace where Istio is installed. If left empty, it is assumed to be the same namespace as where Kiali is installed (i.e. 'deployment.namespace').",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kiali_feature_flags": schema.SingleNestedAttribute{
						Description:         "Kiali features that can be enabled or disabled.",
						MarkdownDescription: "Kiali features that can be enabled or disabled.",
						Attributes: map[string]schema.Attribute{
							"disabled_features": schema.ListAttribute{
								Description:         "There may be some features that admins do not want to be accessible to users (even in 'view only' mode). In this case, this setting allows you to disable one or more of those features entirely.",
								MarkdownDescription: "There may be some features that admins do not want to be accessible to users (even in 'view only' mode). In this case, this setting allows you to disable one or more of those features entirely.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"istio_annotation_action": schema.BoolAttribute{
								Description:         "Flag to enable/disable an Action to edit annotations.",
								MarkdownDescription: "Flag to enable/disable an Action to edit annotations.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"istio_injection_action": schema.BoolAttribute{
								Description:         "Flag to enable/disable an Action to label a namespace for automatic Istio Sidecar injection.",
								MarkdownDescription: "Flag to enable/disable an Action to label a namespace for automatic Istio Sidecar injection.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"istio_upgrade_action": schema.BoolAttribute{
								Description:         "Flag to activate the Kiali functionality of upgrading namespaces to point to an installed Istio Canary revision. Related Canary upgrade and current revisions of Istio should be defined in 'istio_canary_revision' section.",
								MarkdownDescription: "Flag to activate the Kiali functionality of upgrading namespaces to point to an installed Istio Canary revision. Related Canary upgrade and current revisions of Istio should be defined in 'istio_canary_revision' section.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ui_defaults": schema.SingleNestedAttribute{
								Description:         "Default settings for the UI. These defaults apply to all users.",
								MarkdownDescription: "Default settings for the UI. These defaults apply to all users.",
								Attributes: map[string]schema.Attribute{
									"graph": schema.SingleNestedAttribute{
										Description:         "Default settings for the Graph UI.",
										MarkdownDescription: "Default settings for the Graph UI.",
										Attributes: map[string]schema.Attribute{
											"find_options": schema.ListNestedAttribute{
												Description:         "A list of commonly used and useful find expressions that will be provided to the user out-of-box.",
												MarkdownDescription: "A list of commonly used and useful find expressions that will be provided to the user out-of-box.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"auto_select": schema.BoolAttribute{
															Description:         "If true this option will be selected and take effect automatically. Note that only one option in the list can have this value be set to true.",
															MarkdownDescription: "If true this option will be selected and take effect automatically. Note that only one option in the list can have this value be set to true.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "Human-readable text to let the user know what the expression does.",
															MarkdownDescription: "Human-readable text to let the user know what the expression does.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression": schema.StringAttribute{
															Description:         "The find expression.",
															MarkdownDescription: "The find expression.",
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

											"hide_options": schema.ListNestedAttribute{
												Description:         "A list of commonly used and useful hide expressions that will be provided to the user out-of-box.",
												MarkdownDescription: "A list of commonly used and useful hide expressions that will be provided to the user out-of-box.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"auto_select": schema.BoolAttribute{
															Description:         "If true this option will be selected and take effect automatically. Note that only one option in the list can have this value be set to true.",
															MarkdownDescription: "If true this option will be selected and take effect automatically. Note that only one option in the list can have this value be set to true.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"description": schema.StringAttribute{
															Description:         "Human-readable text to let the user know what the expression does.",
															MarkdownDescription: "Human-readable text to let the user know what the expression does.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"expression": schema.StringAttribute{
															Description:         "The hide expression.",
															MarkdownDescription: "The hide expression.",
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

											"traffic": schema.SingleNestedAttribute{
												Description:         "These settings determine which rates are used to determine graph traffic.",
												MarkdownDescription: "These settings determine which rates are used to determine graph traffic.",
												Attributes: map[string]schema.Attribute{
													"grpc": schema.StringAttribute{
														Description:         "gRPC traffic is measured in requests or sent/received/total messages. Value must be one of: 'none', 'requests', 'sent', 'received', or 'total'.",
														MarkdownDescription: "gRPC traffic is measured in requests or sent/received/total messages. Value must be one of: 'none', 'requests', 'sent', 'received', or 'total'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("none", "requests", "sent", "received", "total"),
														},
													},

													"http": schema.StringAttribute{
														Description:         "HTTP traffic is measured in requests. Value must be one of: 'none' or 'requests'.",
														MarkdownDescription: "HTTP traffic is measured in requests. Value must be one of: 'none' or 'requests'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("none", "requests"),
														},
													},

													"tcp": schema.StringAttribute{
														Description:         "TCP traffic is measured in sent/received/total bytes. Only request traffic supplies response codes. Value must be one of: 'none', 'sent', 'received', or 'total'.",
														MarkdownDescription: "TCP traffic is measured in sent/received/total bytes. Only request traffic supplies response codes. Value must be one of: 'none', 'sent', 'received', or 'total'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("none", "sent", "received", "total"),
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

									"i18n": schema.SingleNestedAttribute{
										Description:         "Default settings for the i18n values.",
										MarkdownDescription: "Default settings for the i18n values.",
										Attributes: map[string]schema.Attribute{
											"language": schema.StringAttribute{
												Description:         "Default language used in Kiali application.",
												MarkdownDescription: "Default language used in Kiali application.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"show_selector": schema.BoolAttribute{
												Description:         "If true Kiali masthead displays language selector icon. Default is false.",
												MarkdownDescription: "If true Kiali masthead displays language selector icon. Default is false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"list": schema.SingleNestedAttribute{
										Description:         "Default settings for the List views (Apps, Workloads, etc).",
										MarkdownDescription: "Default settings for the List views (Apps, Workloads, etc).",
										Attributes: map[string]schema.Attribute{
											"include_health": schema.BoolAttribute{
												Description:         "Include Health column (by default) for applicable list views. Setting to false can improve performance.",
												MarkdownDescription: "Include Health column (by default) for applicable list views. Setting to false can improve performance.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"include_istio_resources": schema.BoolAttribute{
												Description:         "Include Istio resources (by default) in Details column for applicable list views. Setting to false can improve performance.",
												MarkdownDescription: "Include Istio resources (by default) in Details column for applicable list views. Setting to false can improve performance.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"include_validations": schema.BoolAttribute{
												Description:         "Include Configuration validation column (by default) for applicable list views. Setting to false can improve performance.",
												MarkdownDescription: "Include Configuration validation column (by default) for applicable list views. Setting to false can improve performance.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"show_include_toggles": schema.BoolAttribute{
												Description:         "If true list pages display checkbox toggles for the include options, Otherwise the configured settings are applied but can not be changed by the user. Default is false.",
												MarkdownDescription: "If true list pages display checkbox toggles for the include options, Otherwise the configured settings are applied but can not be changed by the user. Default is false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"metrics_inbound": schema.SingleNestedAttribute{
										Description:         "Additional label aggregation for inbound metric pages in detail pages. You will see these configurations in the 'Metric Settings' drop-down. An example, ''' metrics_inbound: aggregations: - display_name: Istio Network label: topology_istio_io_network - display_name: Istio Revision label: istio_io_rev ''' ",
										MarkdownDescription: "Additional label aggregation for inbound metric pages in detail pages. You will see these configurations in the 'Metric Settings' drop-down. An example, ''' metrics_inbound: aggregations: - display_name: Istio Network label: topology_istio_io_network - display_name: Istio Revision label: istio_io_rev ''' ",
										Attributes: map[string]schema.Attribute{
											"aggregations": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"display_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label": schema.StringAttribute{
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

									"metrics_outbound": schema.SingleNestedAttribute{
										Description:         "Additional label aggregation for outbound metric pages in detail pages. You will see these configurations in the 'Metric Settings' drop-down. An example, ''' metrics_outbound: aggregations: - display_name: Istio Network label: topology_istio_io_network - display_name: Istio Revision label: istio_io_rev ''' ",
										MarkdownDescription: "Additional label aggregation for outbound metric pages in detail pages. You will see these configurations in the 'Metric Settings' drop-down. An example, ''' metrics_outbound: aggregations: - display_name: Istio Network label: topology_istio_io_network - display_name: Istio Revision label: istio_io_rev ''' ",
										Attributes: map[string]schema.Attribute{
											"aggregations": schema.ListNestedAttribute{
												Description:         "",
												MarkdownDescription: "",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"display_name": schema.StringAttribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"label": schema.StringAttribute{
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

									"metrics_per_refresh": schema.StringAttribute{
										Description:         "Duration of metrics to fetch on each refresh. Value must be one of: '1m', '2m', '5m', '10m', '30m', '1h', '3h', '6h', '12h', '1d', '7d', or '30d'",
										MarkdownDescription: "Duration of metrics to fetch on each refresh. Value must be one of: '1m', '2m', '5m', '10m', '30m', '1h', '3h', '6h', '12h', '1d', '7d', or '30d'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("1m", "2m", "5m", "10m", "30m", "1h", "3h", "6h", "12h", "1d", "7d", "30d"),
										},
									},

									"namespaces": schema.ListAttribute{
										Description:         "Default selections for the namespace selection dropdown. Non-existent or inaccessible namespaces will be ignored. Omit or set to an empty array for no default namespaces.",
										MarkdownDescription: "Default selections for the namespace selection dropdown. Non-existent or inaccessible namespaces will be ignored. Omit or set to an empty array for no default namespaces.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"refresh_interval": schema.StringAttribute{
										Description:         "The automatic refresh interval for pages offering automatic refresh. Value must be one of: 'pause', '10s', '15s', '30s', '1m', '5m' or '15m'",
										MarkdownDescription: "The automatic refresh interval for pages offering automatic refresh. Value must be one of: 'pause', '10s', '15s', '30s', '1m', '5m' or '15m'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("pause", "10s", "15s", "30s", "1m", "5m", "15m"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"validations": schema.SingleNestedAttribute{
								Description:         "Features specific to the validations subsystem.",
								MarkdownDescription: "Features specific to the validations subsystem.",
								Attributes: map[string]schema.Attribute{
									"ignore": schema.ListAttribute{
										Description:         "A list of one or more validation codes whose errors are to be ignored.",
										MarkdownDescription: "A list of one or more validation codes whose errors are to be ignored.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"skip_wildcard_gateway_hosts": schema.BoolAttribute{
										Description:         "The KIA0301 validation checks duplicity of host and port combinations across all Istio Gateways. This includes also Gateways with '*' in hosts. But Istio considers such a Gateway with a wildcard in hosts as the last in order, after the Gateways with FQDN in hosts. This option is to skip Gateways with wildcards in hosts from the KIA0301 validations but still keep Gateways with FQDN hosts.",
										MarkdownDescription: "The KIA0301 validation checks duplicity of host and port combinations across all Istio Gateways. This includes also Gateways with '*' in hosts. But Istio considers such a Gateway with a wildcard in hosts as the last in order, after the Gateways with FQDN in hosts. This option is to skip Gateways with wildcards in hosts from the KIA0301 validations but still keep Gateways with FQDN hosts.",
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

					"kubernetes_config": schema.SingleNestedAttribute{
						Description:         "Configuration of Kiali's access of the Kubernetes API.",
						MarkdownDescription: "Configuration of Kiali's access of the Kubernetes API.",
						Attributes: map[string]schema.Attribute{
							"burst": schema.Int64Attribute{
								Description:         "The Burst value of the Kubernetes client.",
								MarkdownDescription: "The Burst value of the Kubernetes client.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_duration": schema.Int64Attribute{
								Description:         "The ratio interval (expressed in seconds) used for the cache to perform a full refresh. Only used when 'cache_enabled' is true.",
								MarkdownDescription: "The ratio interval (expressed in seconds) used for the cache to perform a full refresh. Only used when 'cache_enabled' is true.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cache_token_namespace_duration": schema.Int64Attribute{
								Description:         "This Kiali cache is a list of namespaces per user. This is typically a short-lived cache compared with the duration of the namespace cache defined by the 'cache_duration' setting. This is specified in seconds.",
								MarkdownDescription: "This Kiali cache is a list of namespaces per user. This is typically a short-lived cache compared with the duration of the namespace cache defined by the 'cache_duration' setting. This is specified in seconds.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_name": schema.StringAttribute{
								Description:         "The name of the cluster Kiali is deployed in. This is only used in multi cluster environments. If not set, Kiali will try to auto detect the cluster name from the Istiod deployment or use the default 'Kubernetes'.",
								MarkdownDescription: "The name of the cluster Kiali is deployed in. This is only used in multi cluster environments. If not set, Kiali will try to auto detect the cluster name from the Istiod deployment or use the default 'Kubernetes'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"excluded_workloads": schema.ListAttribute{
								Description:         "List of controllers that won't be used for Workload calculation. Kiali queries Deployment, ReplicaSet, ReplicationController, DeploymentConfig, StatefulSet, Job and CronJob controllers. Deployment and ReplicaSet will be always queried, but ReplicationController, DeploymentConfig, StatefulSet, Job and CronJobs can be skipped from Kiali workloads queries if they are present in this list.",
								MarkdownDescription: "List of controllers that won't be used for Workload calculation. Kiali queries Deployment, ReplicaSet, ReplicationController, DeploymentConfig, StatefulSet, Job and CronJob controllers. Deployment and ReplicaSet will be always queried, but ReplicationController, DeploymentConfig, StatefulSet, Job and CronJobs can be skipped from Kiali workloads queries if they are present in this list.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"qps": schema.Int64Attribute{
								Description:         "The QPS value of the Kubernetes client.",
								MarkdownDescription: "The QPS value of the Kubernetes client.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"login_token": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"expiration_seconds": schema.Int64Attribute{
								Description:         "A user's login token expiration specified in seconds. This is applicable to token and header auth strategies only.",
								MarkdownDescription: "A user's login token expiration specified in seconds. This is applicable to token and header auth strategies only.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"signing_key": schema.StringAttribute{
								Description:         "The signing key used to generate tokens for user authentication. Because this is potentially sensitive, you have the option to store this value in a secret. If you store this signing key value in a secret, you must indicate what key in what secret by setting this value to a string in the form of 'secret:<secretName>:<secretKey>'. If left as an empty string, a secret with a random signing key will be generated for you. The signing key must be 16, 24 or 32 byte long.",
								MarkdownDescription: "The signing key used to generate tokens for user authentication. Because this is potentially sensitive, you have the option to store this value in a secret. If you store this signing key value in a secret, you must indicate what key in what secret by setting this value to a string in the form of 'secret:<secretName>:<secretKey>'. If left as an empty string, a secret with a random signing key will be generated for you. The signing key must be 16, 24 or 32 byte long.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"server": schema.SingleNestedAttribute{
						Description:         "Configuration that controls some core components within the Kiali Server.",
						MarkdownDescription: "Configuration that controls some core components within the Kiali Server.",
						Attributes: map[string]schema.Attribute{
							"address": schema.StringAttribute{
								Description:         "Where the Kiali server is bound. The console and API server are accessible on this host.",
								MarkdownDescription: "Where the Kiali server is bound. The console and API server are accessible on this host.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"audit_log": schema.BoolAttribute{
								Description:         "When true, allows additional audit logging on write operations.",
								MarkdownDescription: "When true, allows additional audit logging on write operations.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cors_allow_all": schema.BoolAttribute{
								Description:         "When true, allows the web console to send requests to other domains other than where the console came from. Typically used for development environments only.",
								MarkdownDescription: "When true, allows the web console to send requests to other domains other than where the console came from. Typically used for development environments only.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gzip_enabled": schema.BoolAttribute{
								Description:         "When true, Kiali serves http requests with gzip enabled (if the browser supports it) when the requests are over 1400 bytes.",
								MarkdownDescription: "When true, Kiali serves http requests with gzip enabled (if the browser supports it) when the requests are over 1400 bytes.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_port": schema.Int64Attribute{
								Description:         "If 'deployment.service_type' is 'NodePort' and this value is set, then this is the node port that the Kiali service will listen to.",
								MarkdownDescription: "If 'deployment.service_type' is 'NodePort' and this value is set, then this is the node port that the Kiali service will listen to.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"observability": schema.SingleNestedAttribute{
								Description:         "Settings to enable observability into the Kiali server itself.",
								MarkdownDescription: "Settings to enable observability into the Kiali server itself.",
								Attributes: map[string]schema.Attribute{
									"metrics": schema.SingleNestedAttribute{
										Description:         "Settings that control how Kiali itself emits its own metrics.",
										MarkdownDescription: "Settings that control how Kiali itself emits its own metrics.",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "When true, the metrics endpoint will be available for Prometheus to scrape.",
												MarkdownDescription: "When true, the metrics endpoint will be available for Prometheus to scrape.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"port": schema.Int64Attribute{
												Description:         "The port that the server will bind to in order to receive metric requests. This is the port Prometheus will need to scrape when collecting metrics from Kiali.",
												MarkdownDescription: "The port that the server will bind to in order to receive metric requests. This is the port Prometheus will need to scrape when collecting metrics from Kiali.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"tracing": schema.SingleNestedAttribute{
										Description:         "Settings that control how the Kiali server itself emits its own tracing data.",
										MarkdownDescription: "Settings that control how the Kiali server itself emits its own tracing data.",
										Attributes: map[string]schema.Attribute{
											"collector_type": schema.StringAttribute{
												Description:         "The collector type to use. Value must be one of: 'jaeger' or 'otel'.",
												MarkdownDescription: "The collector type to use. Value must be one of: 'jaeger' or 'otel'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"collector_url": schema.StringAttribute{
												Description:         "The URL used to determine where the Kiali server tracing data will be stored.",
												MarkdownDescription: "The URL used to determine where the Kiali server tracing data will be stored.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"enabled": schema.BoolAttribute{
												Description:         "When true, the Kiali server itself will product its own tracing data.",
												MarkdownDescription: "When true, the Kiali server itself will product its own tracing data.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"otel": schema.SingleNestedAttribute{
												Description:         "Specific properties when the collector type is 'otel'.",
												MarkdownDescription: "Specific properties when the collector type is 'otel'.",
												Attributes: map[string]schema.Attribute{
													"ca_name": schema.StringAttribute{
														Description:         "The name of the CA cert; this is used when 'tls_enabled' is 'true' and 'skip_verify' is 'false'.",
														MarkdownDescription: "The name of the CA cert; this is used when 'tls_enabled' is 'true' and 'skip_verify' is 'false'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"protocol": schema.StringAttribute{
														Description:         "Protocol. Supported values are: 'http', 'https' or 'grpc'.",
														MarkdownDescription: "Protocol. Supported values are: 'http', 'https' or 'grpc'.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"skip_verify": schema.BoolAttribute{
														Description:         "If true, TLS certificate verification will not be performed. This is an unsecure option and is recommended only for testing.",
														MarkdownDescription: "If true, TLS certificate verification will not be performed. This is an unsecure option and is recommended only for testing.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"tls_enabled": schema.BoolAttribute{
														Description:         "Enable TLS for the collector. This must be specified when 'protocol' is 'https' or 'grpc'. When you set this to 'true', you must also set a 'ca_name' or set 'skip_verify' to true.",
														MarkdownDescription: "Enable TLS for the collector. This must be specified when 'protocol' is 'https' or 'grpc'. When you set this to 'true', you must also set a 'ca_name' or set 'skip_verify' to true.",
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

							"port": schema.Int64Attribute{
								Description:         "The port that the server will bind to in order to receive console and API requests.",
								MarkdownDescription: "The port that the server will bind to in order to receive console and API requests.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"profiler": schema.SingleNestedAttribute{
								Description:         "Controls the internal profiler used to debug the internals of Kiali",
								MarkdownDescription: "Controls the internal profiler used to debug the internals of Kiali",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "When 'true', the profiler will be enabled and accessible at /debug/pprof/ on the Kiali endpoint.",
										MarkdownDescription: "When 'true', the profiler will be enabled and accessible at /debug/pprof/ on the Kiali endpoint.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"web_fqdn": schema.StringAttribute{
								Description:         "Defines the public domain where Kiali is being served. This is the 'domain' part of the URL (usually it's a fully-qualified domain name). For example, 'kiali.example.org'. When empty, Kiali will try to guess this value from HTTP headers. On non-OpenShift clusters, you must populate this value if you want to enable cross-linking between Kiali instances in a multi-cluster setup.",
								MarkdownDescription: "Defines the public domain where Kiali is being served. This is the 'domain' part of the URL (usually it's a fully-qualified domain name). For example, 'kiali.example.org'. When empty, Kiali will try to guess this value from HTTP headers. On non-OpenShift clusters, you must populate this value if you want to enable cross-linking between Kiali instances in a multi-cluster setup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_history_mode": schema.StringAttribute{
								Description:         "Define the history mode of kiali UI. Value must be one of: 'browser' or 'hash'.",
								MarkdownDescription: "Define the history mode of kiali UI. Value must be one of: 'browser' or 'hash'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_port": schema.StringAttribute{
								Description:         "Defines the ingress port where the connections come from. This is usually necessary when the application responds through a proxy/ingress, and it does not forward the correct headers (when this happens, Kiali cannot guess the port). When empty, Kiali will try to guess this value from HTTP headers.",
								MarkdownDescription: "Defines the ingress port where the connections come from. This is usually necessary when the application responds through a proxy/ingress, and it does not forward the correct headers (when this happens, Kiali cannot guess the port). When empty, Kiali will try to guess this value from HTTP headers.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_root": schema.StringAttribute{
								Description:         "Defines the context root path for the Kiali console and API endpoints and readiness probes. When providing a context root path that is not '/', do not add a trailing slash (i.e. use '/kiali' not '/kiali/'). When empty, this will default to '/' on OpenShift and '/kiali' on other Kubernetes environments.",
								MarkdownDescription: "Defines the context root path for the Kiali console and API endpoints and readiness probes. When providing a context root path that is not '/', do not add a trailing slash (i.e. use '/kiali' not '/kiali/'). When empty, this will default to '/' on OpenShift and '/kiali' on other Kubernetes environments.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"web_schema": schema.StringAttribute{
								Description:         "Defines the public HTTP schema used to serve Kiali. Value must be one of: 'http' or 'https'. When empty, Kiali will try to guess this value from HTTP headers. On non-OpenShift clusters, you must populate this value if you want to enable cross-linking between Kiali instances in a multi-cluster setup.",
								MarkdownDescription: "Defines the public HTTP schema used to serve Kiali. Value must be one of: 'http' or 'https'. When empty, Kiali will try to guess this value from HTTP headers. On non-OpenShift clusters, you must populate this value if you want to enable cross-linking between Kiali instances in a multi-cluster setup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("", "http", "https"),
								},
							},

							"write_timeout": schema.Int64Attribute{
								Description:         "The maximum duration, in seconds, before timing out writes of the HTTP response back to the client. Default is 30. In OpenShift clusters, the route request time out should be also increased as the default is 30 seconds. This can be done by annotating the specific route with 'haproxy.router.openshift.io/timeout'. See https://docs.openshift.com/container-platform/4.16/networking/routes/route-configuration.html#nw-configuring-route-timeouts_route-configuration for further details. ",
								MarkdownDescription: "The maximum duration, in seconds, before timing out writes of the HTTP response back to the client. Default is 30. In OpenShift clusters, the route request time out should be also increased as the default is 30 seconds. This can be done by annotating the specific route with 'haproxy.router.openshift.io/timeout'. See https://docs.openshift.com/container-platform/4.16/networking/routes/route-configuration.html#nw-configuring-route-timeouts_route-configuration for further details. ",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": schema.StringAttribute{
						Description:         "The version of the Ansible role that will be executed in order to install Kiali. This also indirectly determines the version of Kiali that will be installed. You normally will want to use 'default' since this is the only officially supported value today. If not specified, the value of 'default' is assumed which means the most recent Ansible role is used; thus the most recent release of Kiali will be installed. Refer to this file to see what the valid values are for this 'version' field (as defined in the master branch), https://github.com/kiali/kiali-operator/blob/master/playbooks/kiali-default-supported-images.yml This 'version' setting affects the defaults of the 'deployment.image_name' and 'deployment.image_version' settings. See the documentation for those settings below for additional details. In short, this 'version' setting will dictate which version of the Kiali image will be deployed by default. However, if you explicitly set 'deployment.image_name' and/or 'deployment.image_version' to reference your own custom image, that will override the default Kiali image to be installed; therefore, you are responsible for ensuring those settings are compatible with the Ansible role that will be executed in order to install Kiali (i.e. your custom Kiali image must be compatible with the rest of the configuration and resources the operator will install). ",
						MarkdownDescription: "The version of the Ansible role that will be executed in order to install Kiali. This also indirectly determines the version of Kiali that will be installed. You normally will want to use 'default' since this is the only officially supported value today. If not specified, the value of 'default' is assumed which means the most recent Ansible role is used; thus the most recent release of Kiali will be installed. Refer to this file to see what the valid values are for this 'version' field (as defined in the master branch), https://github.com/kiali/kiali-operator/blob/master/playbooks/kiali-default-supported-images.yml This 'version' setting affects the defaults of the 'deployment.image_name' and 'deployment.image_version' settings. See the documentation for those settings below for additional details. In short, this 'version' setting will dictate which version of the Kiali image will be deployed by default. However, if you explicitly set 'deployment.image_name' and/or 'deployment.image_version' to reference your own custom image, that will override the default Kiali image to be installed; therefore, you are responsible for ensuring those settings are compatible with the Ansible role that will be executed in order to install Kiali (i.e. your custom Kiali image must be compatible with the rest of the configuration and resources the operator will install). ",
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
	}
}

func (r *KialiIoKialiV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_kiali_io_kiali_v1alpha1_manifest")

	var model KialiIoKialiV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("kiali.io/v1alpha1")
	model.Kind = pointer.String("Kiali")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
