/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package awx_ansible_com_v1beta1

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
	_ datasource.DataSource = &AwxAnsibleComAwxV1Beta1Manifest{}
)

func NewAwxAnsibleComAwxV1Beta1Manifest() datasource.DataSource {
	return &AwxAnsibleComAwxV1Beta1Manifest{}
}

type AwxAnsibleComAwxV1Beta1Manifest struct{}

type AwxAnsibleComAwxV1Beta1ManifestData struct {
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
		Additional_labels     *[]string `tfsdk:"additional_labels" json:"additional_labels,omitempty"`
		Admin_email           *string   `tfsdk:"admin_email" json:"admin_email,omitempty"`
		Admin_password_secret *string   `tfsdk:"admin_password_secret" json:"admin_password_secret,omitempty"`
		Admin_user            *string   `tfsdk:"admin_user" json:"admin_user,omitempty"`
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
		Annotations                  *string `tfsdk:"annotations" json:"annotations,omitempty"`
		Api_version                  *string `tfsdk:"api_version" json:"api_version,omitempty"`
		Auto_upgrade                 *bool   `tfsdk:"auto_upgrade" json:"auto_upgrade,omitempty"`
		Broadcast_websocket_secret   *string `tfsdk:"broadcast_websocket_secret" json:"broadcast_websocket_secret,omitempty"`
		Bundle_cacert_secret         *string `tfsdk:"bundle_cacert_secret" json:"bundle_cacert_secret,omitempty"`
		Ca_trust_bundle              *string `tfsdk:"ca_trust_bundle" json:"ca_trust_bundle,omitempty"`
		Control_plane_ee_image       *string `tfsdk:"control_plane_ee_image" json:"control_plane_ee_image,omitempty"`
		Control_plane_priority_class *string `tfsdk:"control_plane_priority_class" json:"control_plane_priority_class,omitempty"`
		Create_preload_data          *bool   `tfsdk:"create_preload_data" json:"create_preload_data,omitempty"`
		Csrf_cookie_secure           *string `tfsdk:"csrf_cookie_secure" json:"csrf_cookie_secure,omitempty"`
		Deployment_type              *string `tfsdk:"deployment_type" json:"deployment_type,omitempty"`
		Development_mode             *bool   `tfsdk:"development_mode" json:"development_mode,omitempty"`
		Ee_extra_env                 *string `tfsdk:"ee_extra_env" json:"ee_extra_env,omitempty"`
		Ee_extra_volume_mounts       *string `tfsdk:"ee_extra_volume_mounts" json:"ee_extra_volume_mounts,omitempty"`
		Ee_images                    *[]struct {
			Image *string `tfsdk:"image" json:"image,omitempty"`
			Name  *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"ee_images" json:"ee_images,omitempty"`
		Ee_pull_credentials_secret *string `tfsdk:"ee_pull_credentials_secret" json:"ee_pull_credentials_secret,omitempty"`
		Ee_resource_requirements   *struct {
			Limits *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"ee_resource_requirements" json:"ee_resource_requirements,omitempty"`
		Extra_settings *[]struct {
			Setting *string            `tfsdk:"setting" json:"setting,omitempty"`
			Value   *map[string]string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"extra_settings" json:"extra_settings,omitempty"`
		Extra_volumes           *string `tfsdk:"extra_volumes" json:"extra_volumes,omitempty"`
		Garbage_collect_secrets *bool   `tfsdk:"garbage_collect_secrets" json:"garbage_collect_secrets,omitempty"`
		Host_aliases            *[]struct {
			Hostnames *[]string `tfsdk:"hostnames" json:"hostnames,omitempty"`
			Ip        *string   `tfsdk:"ip" json:"ip,omitempty"`
		} `tfsdk:"host_aliases" json:"host_aliases,omitempty"`
		Hostname            *string   `tfsdk:"hostname" json:"hostname,omitempty"`
		Image               *string   `tfsdk:"image" json:"image,omitempty"`
		Image_pull_policy   *string   `tfsdk:"image_pull_policy" json:"image_pull_policy,omitempty"`
		Image_pull_secret   *string   `tfsdk:"image_pull_secret" json:"image_pull_secret,omitempty"`
		Image_pull_secrets  *[]string `tfsdk:"image_pull_secrets" json:"image_pull_secrets,omitempty"`
		Image_version       *string   `tfsdk:"image_version" json:"image_version,omitempty"`
		Ingress_annotations *string   `tfsdk:"ingress_annotations" json:"ingress_annotations,omitempty"`
		Ingress_api_version *string   `tfsdk:"ingress_api_version" json:"ingress_api_version,omitempty"`
		Ingress_class_name  *string   `tfsdk:"ingress_class_name" json:"ingress_class_name,omitempty"`
		Ingress_controller  *string   `tfsdk:"ingress_controller" json:"ingress_controller,omitempty"`
		Ingress_hosts       *[]struct {
			Hostname   *string `tfsdk:"hostname" json:"hostname,omitempty"`
			Tls_secret *string `tfsdk:"tls_secret" json:"tls_secret,omitempty"`
		} `tfsdk:"ingress_hosts" json:"ingress_hosts,omitempty"`
		Ingress_path                         *string `tfsdk:"ingress_path" json:"ingress_path,omitempty"`
		Ingress_path_type                    *string `tfsdk:"ingress_path_type" json:"ingress_path_type,omitempty"`
		Ingress_tls_secret                   *string `tfsdk:"ingress_tls_secret" json:"ingress_tls_secret,omitempty"`
		Ingress_type                         *string `tfsdk:"ingress_type" json:"ingress_type,omitempty"`
		Init_container_extra_commands        *string `tfsdk:"init_container_extra_commands" json:"init_container_extra_commands,omitempty"`
		Init_container_extra_volume_mounts   *string `tfsdk:"init_container_extra_volume_mounts" json:"init_container_extra_volume_mounts,omitempty"`
		Init_container_image                 *string `tfsdk:"init_container_image" json:"init_container_image,omitempty"`
		Init_container_image_version         *string `tfsdk:"init_container_image_version" json:"init_container_image_version,omitempty"`
		Init_container_resource_requirements *struct {
			Limits *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"init_container_resource_requirements" json:"init_container_resource_requirements,omitempty"`
		Init_projects_container_image                 *string   `tfsdk:"init_projects_container_image" json:"init_projects_container_image,omitempty"`
		Ipv6_disabled                                 *bool     `tfsdk:"ipv6_disabled" json:"ipv6_disabled,omitempty"`
		Kind                                          *string   `tfsdk:"kind" json:"kind,omitempty"`
		Ldap_cacert_secret                            *string   `tfsdk:"ldap_cacert_secret" json:"ldap_cacert_secret,omitempty"`
		Ldap_password_secret                          *string   `tfsdk:"ldap_password_secret" json:"ldap_password_secret,omitempty"`
		Loadbalancer_class                            *string   `tfsdk:"loadbalancer_class" json:"loadbalancer_class,omitempty"`
		Loadbalancer_ip                               *string   `tfsdk:"loadbalancer_ip" json:"loadbalancer_ip,omitempty"`
		Loadbalancer_port                             *int64    `tfsdk:"loadbalancer_port" json:"loadbalancer_port,omitempty"`
		Loadbalancer_protocol                         *string   `tfsdk:"loadbalancer_protocol" json:"loadbalancer_protocol,omitempty"`
		Metrics_utility_configmap                     *string   `tfsdk:"metrics_utility_configmap" json:"metrics_utility_configmap,omitempty"`
		Metrics_utility_console_enabled               *bool     `tfsdk:"metrics_utility_console_enabled" json:"metrics_utility_console_enabled,omitempty"`
		Metrics_utility_cronjob_gather_schedule       *string   `tfsdk:"metrics_utility_cronjob_gather_schedule" json:"metrics_utility_cronjob_gather_schedule,omitempty"`
		Metrics_utility_cronjob_report_schedule       *string   `tfsdk:"metrics_utility_cronjob_report_schedule" json:"metrics_utility_cronjob_report_schedule,omitempty"`
		Metrics_utility_enabled                       *bool     `tfsdk:"metrics_utility_enabled" json:"metrics_utility_enabled,omitempty"`
		Metrics_utility_image                         *string   `tfsdk:"metrics_utility_image" json:"metrics_utility_image,omitempty"`
		Metrics_utility_image_pull_policy             *string   `tfsdk:"metrics_utility_image_pull_policy" json:"metrics_utility_image_pull_policy,omitempty"`
		Metrics_utility_image_version                 *string   `tfsdk:"metrics_utility_image_version" json:"metrics_utility_image_version,omitempty"`
		Metrics_utility_pvc_claim                     *string   `tfsdk:"metrics_utility_pvc_claim" json:"metrics_utility_pvc_claim,omitempty"`
		Metrics_utility_pvc_claim_size                *string   `tfsdk:"metrics_utility_pvc_claim_size" json:"metrics_utility_pvc_claim_size,omitempty"`
		Metrics_utility_pvc_claim_storage_class       *string   `tfsdk:"metrics_utility_pvc_claim_storage_class" json:"metrics_utility_pvc_claim_storage_class,omitempty"`
		Metrics_utility_secret                        *string   `tfsdk:"metrics_utility_secret" json:"metrics_utility_secret,omitempty"`
		Metrics_utility_ship_target                   *string   `tfsdk:"metrics_utility_ship_target" json:"metrics_utility_ship_target,omitempty"`
		Nginx_listen_queue_size                       *int64    `tfsdk:"nginx_listen_queue_size" json:"nginx_listen_queue_size,omitempty"`
		Nginx_worker_connections                      *int64    `tfsdk:"nginx_worker_connections" json:"nginx_worker_connections,omitempty"`
		Nginx_worker_cpu_affinity                     *string   `tfsdk:"nginx_worker_cpu_affinity" json:"nginx_worker_cpu_affinity,omitempty"`
		Nginx_worker_processes                        *int64    `tfsdk:"nginx_worker_processes" json:"nginx_worker_processes,omitempty"`
		No_log                                        *bool     `tfsdk:"no_log" json:"no_log,omitempty"`
		Node_selector                                 *string   `tfsdk:"node_selector" json:"node_selector,omitempty"`
		Nodeport_port                                 *int64    `tfsdk:"nodeport_port" json:"nodeport_port,omitempty"`
		Old_postgres_configuration_secret             *string   `tfsdk:"old_postgres_configuration_secret" json:"old_postgres_configuration_secret,omitempty"`
		Postgres_configuration_secret                 *string   `tfsdk:"postgres_configuration_secret" json:"postgres_configuration_secret,omitempty"`
		Postgres_data_volume_init                     *bool     `tfsdk:"postgres_data_volume_init" json:"postgres_data_volume_init,omitempty"`
		Postgres_extra_args                           *[]string `tfsdk:"postgres_extra_args" json:"postgres_extra_args,omitempty"`
		Postgres_extra_volume_mounts                  *string   `tfsdk:"postgres_extra_volume_mounts" json:"postgres_extra_volume_mounts,omitempty"`
		Postgres_extra_volumes                        *string   `tfsdk:"postgres_extra_volumes" json:"postgres_extra_volumes,omitempty"`
		Postgres_image                                *string   `tfsdk:"postgres_image" json:"postgres_image,omitempty"`
		Postgres_image_version                        *string   `tfsdk:"postgres_image_version" json:"postgres_image_version,omitempty"`
		Postgres_init_container_commands              *string   `tfsdk:"postgres_init_container_commands" json:"postgres_init_container_commands,omitempty"`
		Postgres_init_container_resource_requirements *struct {
			Limits *struct {
				Cpu     *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory  *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu     *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory  *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"postgres_init_container_resource_requirements" json:"postgres_init_container_resource_requirements,omitempty"`
		Postgres_keep_pvc_after_upgrade *bool   `tfsdk:"postgres_keep_pvc_after_upgrade" json:"postgres_keep_pvc_after_upgrade,omitempty"`
		Postgres_keepalives             *bool   `tfsdk:"postgres_keepalives" json:"postgres_keepalives,omitempty"`
		Postgres_keepalives_count       *int64  `tfsdk:"postgres_keepalives_count" json:"postgres_keepalives_count,omitempty"`
		Postgres_keepalives_idle        *int64  `tfsdk:"postgres_keepalives_idle" json:"postgres_keepalives_idle,omitempty"`
		Postgres_keepalives_interval    *int64  `tfsdk:"postgres_keepalives_interval" json:"postgres_keepalives_interval,omitempty"`
		Postgres_label_selector         *string `tfsdk:"postgres_label_selector" json:"postgres_label_selector,omitempty"`
		Postgres_priority_class         *string `tfsdk:"postgres_priority_class" json:"postgres_priority_class,omitempty"`
		Postgres_resource_requirements  *struct {
			Limits *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"postgres_resource_requirements" json:"postgres_resource_requirements,omitempty"`
		Postgres_security_context_settings *map[string]string `tfsdk:"postgres_security_context_settings" json:"postgres_security_context_settings,omitempty"`
		Postgres_selector                  *string            `tfsdk:"postgres_selector" json:"postgres_selector,omitempty"`
		Postgres_storage_class             *string            `tfsdk:"postgres_storage_class" json:"postgres_storage_class,omitempty"`
		Postgres_storage_requirements      *struct {
			Limits *struct {
				Storage *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Storage *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"postgres_storage_requirements" json:"postgres_storage_requirements,omitempty"`
		Postgres_tolerations         *string   `tfsdk:"postgres_tolerations" json:"postgres_tolerations,omitempty"`
		Projects_existing_claim      *string   `tfsdk:"projects_existing_claim" json:"projects_existing_claim,omitempty"`
		Projects_persistence         *bool     `tfsdk:"projects_persistence" json:"projects_persistence,omitempty"`
		Projects_storage_access_mode *string   `tfsdk:"projects_storage_access_mode" json:"projects_storage_access_mode,omitempty"`
		Projects_storage_class       *string   `tfsdk:"projects_storage_class" json:"projects_storage_class,omitempty"`
		Projects_storage_size        *string   `tfsdk:"projects_storage_size" json:"projects_storage_size,omitempty"`
		Projects_use_existing_claim  *string   `tfsdk:"projects_use_existing_claim" json:"projects_use_existing_claim,omitempty"`
		Receptor_log_level           *string   `tfsdk:"receptor_log_level" json:"receptor_log_level,omitempty"`
		Redis_capabilities           *[]string `tfsdk:"redis_capabilities" json:"redis_capabilities,omitempty"`
		Redis_image                  *string   `tfsdk:"redis_image" json:"redis_image,omitempty"`
		Redis_image_version          *string   `tfsdk:"redis_image_version" json:"redis_image_version,omitempty"`
		Redis_resource_requirements  *struct {
			Limits *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"redis_resource_requirements" json:"redis_resource_requirements,omitempty"`
		Replicas                        *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
		Route_api_version               *string   `tfsdk:"route_api_version" json:"route_api_version,omitempty"`
		Route_host                      *string   `tfsdk:"route_host" json:"route_host,omitempty"`
		Route_tls_secret                *string   `tfsdk:"route_tls_secret" json:"route_tls_secret,omitempty"`
		Route_tls_termination_mechanism *string   `tfsdk:"route_tls_termination_mechanism" json:"route_tls_termination_mechanism,omitempty"`
		Rsyslog_args                    *[]string `tfsdk:"rsyslog_args" json:"rsyslog_args,omitempty"`
		Rsyslog_command                 *[]string `tfsdk:"rsyslog_command" json:"rsyslog_command,omitempty"`
		Rsyslog_extra_env               *string   `tfsdk:"rsyslog_extra_env" json:"rsyslog_extra_env,omitempty"`
		Rsyslog_extra_volume_mounts     *string   `tfsdk:"rsyslog_extra_volume_mounts" json:"rsyslog_extra_volume_mounts,omitempty"`
		Rsyslog_resource_requirements   *struct {
			Limits *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"rsyslog_resource_requirements" json:"rsyslog_resource_requirements,omitempty"`
		Secret_key_secret           *string            `tfsdk:"secret_key_secret" json:"secret_key_secret,omitempty"`
		Security_context_settings   *map[string]string `tfsdk:"security_context_settings" json:"security_context_settings,omitempty"`
		Service_account_annotations *string            `tfsdk:"service_account_annotations" json:"service_account_annotations,omitempty"`
		Service_annotations         *string            `tfsdk:"service_annotations" json:"service_annotations,omitempty"`
		Service_labels              *string            `tfsdk:"service_labels" json:"service_labels,omitempty"`
		Service_type                *string            `tfsdk:"service_type" json:"service_type,omitempty"`
		Session_cookie_secure       *string            `tfsdk:"session_cookie_secure" json:"session_cookie_secure,omitempty"`
		Set_self_labels             *bool              `tfsdk:"set_self_labels" json:"set_self_labels,omitempty"`
		Task_affinity               *struct {
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
		} `tfsdk:"task_affinity" json:"task_affinity,omitempty"`
		Task_annotations                 *string   `tfsdk:"task_annotations" json:"task_annotations,omitempty"`
		Task_args                        *[]string `tfsdk:"task_args" json:"task_args,omitempty"`
		Task_command                     *[]string `tfsdk:"task_command" json:"task_command,omitempty"`
		Task_extra_env                   *string   `tfsdk:"task_extra_env" json:"task_extra_env,omitempty"`
		Task_extra_volume_mounts         *string   `tfsdk:"task_extra_volume_mounts" json:"task_extra_volume_mounts,omitempty"`
		Task_liveness_failure_threshold  *int64    `tfsdk:"task_liveness_failure_threshold" json:"task_liveness_failure_threshold,omitempty"`
		Task_liveness_initial_delay      *int64    `tfsdk:"task_liveness_initial_delay" json:"task_liveness_initial_delay,omitempty"`
		Task_liveness_period             *int64    `tfsdk:"task_liveness_period" json:"task_liveness_period,omitempty"`
		Task_liveness_timeout            *int64    `tfsdk:"task_liveness_timeout" json:"task_liveness_timeout,omitempty"`
		Task_node_selector               *string   `tfsdk:"task_node_selector" json:"task_node_selector,omitempty"`
		Task_privileged                  *bool     `tfsdk:"task_privileged" json:"task_privileged,omitempty"`
		Task_readiness_failure_threshold *int64    `tfsdk:"task_readiness_failure_threshold" json:"task_readiness_failure_threshold,omitempty"`
		Task_readiness_initial_delay     *int64    `tfsdk:"task_readiness_initial_delay" json:"task_readiness_initial_delay,omitempty"`
		Task_readiness_period            *int64    `tfsdk:"task_readiness_period" json:"task_readiness_period,omitempty"`
		Task_readiness_timeout           *int64    `tfsdk:"task_readiness_timeout" json:"task_readiness_timeout,omitempty"`
		Task_replicas                    *int64    `tfsdk:"task_replicas" json:"task_replicas,omitempty"`
		Task_resource_requirements       *struct {
			Limits *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"task_resource_requirements" json:"task_resource_requirements,omitempty"`
		Task_tolerations                 *string `tfsdk:"task_tolerations" json:"task_tolerations,omitempty"`
		Task_topology_spread_constraints *string `tfsdk:"task_topology_spread_constraints" json:"task_topology_spread_constraints,omitempty"`
		Termination_grace_period_seconds *int64  `tfsdk:"termination_grace_period_seconds" json:"termination_grace_period_seconds,omitempty"`
		Tolerations                      *string `tfsdk:"tolerations" json:"tolerations,omitempty"`
		Topology_spread_constraints      *string `tfsdk:"topology_spread_constraints" json:"topology_spread_constraints,omitempty"`
		Uwsgi_listen_queue_size          *int64  `tfsdk:"uwsgi_listen_queue_size" json:"uwsgi_listen_queue_size,omitempty"`
		Uwsgi_processes                  *int64  `tfsdk:"uwsgi_processes" json:"uwsgi_processes,omitempty"`
		Web_affinity                     *struct {
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
		} `tfsdk:"web_affinity" json:"web_affinity,omitempty"`
		Web_annotations                 *string   `tfsdk:"web_annotations" json:"web_annotations,omitempty"`
		Web_args                        *[]string `tfsdk:"web_args" json:"web_args,omitempty"`
		Web_command                     *[]string `tfsdk:"web_command" json:"web_command,omitempty"`
		Web_extra_env                   *string   `tfsdk:"web_extra_env" json:"web_extra_env,omitempty"`
		Web_extra_volume_mounts         *string   `tfsdk:"web_extra_volume_mounts" json:"web_extra_volume_mounts,omitempty"`
		Web_liveness_failure_threshold  *int64    `tfsdk:"web_liveness_failure_threshold" json:"web_liveness_failure_threshold,omitempty"`
		Web_liveness_initial_delay      *int64    `tfsdk:"web_liveness_initial_delay" json:"web_liveness_initial_delay,omitempty"`
		Web_liveness_period             *int64    `tfsdk:"web_liveness_period" json:"web_liveness_period,omitempty"`
		Web_liveness_timeout            *int64    `tfsdk:"web_liveness_timeout" json:"web_liveness_timeout,omitempty"`
		Web_node_selector               *string   `tfsdk:"web_node_selector" json:"web_node_selector,omitempty"`
		Web_readiness_failure_threshold *int64    `tfsdk:"web_readiness_failure_threshold" json:"web_readiness_failure_threshold,omitempty"`
		Web_readiness_initial_delay     *int64    `tfsdk:"web_readiness_initial_delay" json:"web_readiness_initial_delay,omitempty"`
		Web_readiness_period            *int64    `tfsdk:"web_readiness_period" json:"web_readiness_period,omitempty"`
		Web_readiness_timeout           *int64    `tfsdk:"web_readiness_timeout" json:"web_readiness_timeout,omitempty"`
		Web_replicas                    *int64    `tfsdk:"web_replicas" json:"web_replicas,omitempty"`
		Web_resource_requirements       *struct {
			Limits *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu               *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Ephemeral_storage *string `tfsdk:"ephemeral_storage" json:"ephemeral-storage,omitempty"`
				Memory            *string `tfsdk:"memory" json:"memory,omitempty"`
				Storage           *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"web_resource_requirements" json:"web_resource_requirements,omitempty"`
		Web_tolerations                 *string `tfsdk:"web_tolerations" json:"web_tolerations,omitempty"`
		Web_topology_spread_constraints *string `tfsdk:"web_topology_spread_constraints" json:"web_topology_spread_constraints,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AwxAnsibleComAwxV1Beta1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_awx_ansible_com_awx_v1beta1_manifest"
}

func (r *AwxAnsibleComAwxV1Beta1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Schema validation for the AWX CRD",
		MarkdownDescription: "Schema validation for the AWX CRD",
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
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"additional_labels": schema.ListAttribute{
						Description:         "Additional labels defined on the resource, which should be propagated to child resources",
						MarkdownDescription: "Additional labels defined on the resource, which should be propagated to child resources",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"admin_email": schema.StringAttribute{
						Description:         "The admin user email",
						MarkdownDescription: "The admin user email",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"admin_password_secret": schema.StringAttribute{
						Description:         "Secret where the admin password can be found",
						MarkdownDescription: "Secret where the admin password can be found",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][-a-zA-Z0-9]{0,253}[a-zA-Z0-9]$`), ""),
						},
					},

					"admin_user": schema.StringAttribute{
						Description:         "Username to use for the admin account",
						MarkdownDescription: "Username to use for the admin account",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"affinity": schema.SingleNestedAttribute{
						Description:         "If specified, the pod's scheduling constraints",
						MarkdownDescription: "If specified, the pod's scheduling constraints",
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

					"annotations": schema.StringAttribute{
						Description:         "Common annotations for both Web and Task deployments.",
						MarkdownDescription: "Common annotations for both Web and Task deployments.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"api_version": schema.StringAttribute{
						Description:         "apiVersion of the deployment type",
						MarkdownDescription: "apiVersion of the deployment type",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"auto_upgrade": schema.BoolAttribute{
						Description:         "Should AWX instances be automatically upgraded when operator gets upgraded",
						MarkdownDescription: "Should AWX instances be automatically upgraded when operator gets upgraded",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"broadcast_websocket_secret": schema.StringAttribute{
						Description:         "Secret where the broadcast websocket secret can be found",
						MarkdownDescription: "Secret where the broadcast websocket secret can be found",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][-a-zA-Z0-9]{0,253}[a-zA-Z0-9]$`), ""),
						},
					},

					"bundle_cacert_secret": schema.StringAttribute{
						Description:         "Secret where can be found the trusted Certificate Authority Bundle",
						MarkdownDescription: "Secret where can be found the trusted Certificate Authority Bundle",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ca_trust_bundle": schema.StringAttribute{
						Description:         "Path where the trusted CA bundle is available",
						MarkdownDescription: "Path where the trusted CA bundle is available",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"control_plane_ee_image": schema.StringAttribute{
						Description:         "Registry path to the Execution Environment container image to use on control plane pods",
						MarkdownDescription: "Registry path to the Execution Environment container image to use on control plane pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"control_plane_priority_class": schema.StringAttribute{
						Description:         "Assign a preexisting priority class to the control plane pods",
						MarkdownDescription: "Assign a preexisting priority class to the control plane pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"create_preload_data": schema.BoolAttribute{
						Description:         "Whether or not to preload data upon instance creation",
						MarkdownDescription: "Whether or not to preload data upon instance creation",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"csrf_cookie_secure": schema.StringAttribute{
						Description:         "Set csrf cookie secure mode for web",
						MarkdownDescription: "Set csrf cookie secure mode for web",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"deployment_type": schema.StringAttribute{
						Description:         "Name of the deployment type",
						MarkdownDescription: "Name of the deployment type",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"development_mode": schema.BoolAttribute{
						Description:         "If the deployment should be done in development mode",
						MarkdownDescription: "If the deployment should be done in development mode",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ee_extra_env": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ee_extra_volume_mounts": schema.StringAttribute{
						Description:         "Specify volume mounts to be added to Execution container",
						MarkdownDescription: "Specify volume mounts to be added to Execution container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ee_images": schema.ListNestedAttribute{
						Description:         "Registry path to the Execution Environment container to use",
						MarkdownDescription: "Registry path to the Execution Environment container to use",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"image": schema.StringAttribute{
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ee_pull_credentials_secret": schema.StringAttribute{
						Description:         "Secret where pull credentials for registered ees can be found",
						MarkdownDescription: "Secret where pull credentials for registered ees can be found",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ee_resource_requirements": schema.SingleNestedAttribute{
						Description:         "Resource requirements for the ee container",
						MarkdownDescription: "Resource requirements for the ee container",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

					"extra_settings": schema.ListNestedAttribute{
						Description:         "Extra settings to specify for AWX",
						MarkdownDescription: "Extra settings to specify for AWX",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"setting": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.MapAttribute{
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

					"extra_volumes": schema.StringAttribute{
						Description:         "Specify extra volumes to add to the application pod",
						MarkdownDescription: "Specify extra volumes to add to the application pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"garbage_collect_secrets": schema.BoolAttribute{
						Description:         "Whether or not to remove secrets upon instance removal",
						MarkdownDescription: "Whether or not to remove secrets upon instance removal",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"host_aliases": schema.ListNestedAttribute{
						Description:         "HostAliases for app containers",
						MarkdownDescription: "HostAliases for app containers",
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

					"hostname": schema.StringAttribute{
						Description:         "(Deprecated) The hostname of the instance",
						MarkdownDescription: "(Deprecated) The hostname of the instance",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image": schema.StringAttribute{
						Description:         "Registry path to the application container to use",
						MarkdownDescription: "Registry path to the application container to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "The image pull policy",
						MarkdownDescription: "The image pull policy",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Always", "always", "Never", "never", "IfNotPresent", "ifnotpresent"),
						},
					},

					"image_pull_secret": schema.StringAttribute{
						Description:         "(Deprecated) Image pull secret for app and database containers",
						MarkdownDescription: "(Deprecated) Image pull secret for app and database containers",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListAttribute{
						Description:         "Image pull secrets for app and database containers",
						MarkdownDescription: "Image pull secrets for app and database containers",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_version": schema.StringAttribute{
						Description:         "Application container image version to use",
						MarkdownDescription: "Application container image version to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_annotations": schema.StringAttribute{
						Description:         "Annotations to add to the Ingress Controller",
						MarkdownDescription: "Annotations to add to the Ingress Controller",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_api_version": schema.StringAttribute{
						Description:         "The Ingress API version to use",
						MarkdownDescription: "The Ingress API version to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_class_name": schema.StringAttribute{
						Description:         "The name of ingress class to use instead of the cluster default.",
						MarkdownDescription: "The name of ingress class to use instead of the cluster default.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_controller": schema.StringAttribute{
						Description:         "Special configuration for specific Ingress Controllers",
						MarkdownDescription: "Special configuration for specific Ingress Controllers",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_hosts": schema.ListNestedAttribute{
						Description:         "Ingress hostnames of the instance",
						MarkdownDescription: "Ingress hostnames of the instance",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"hostname": schema.StringAttribute{
									Description:         "Hostname of the instance",
									MarkdownDescription: "Hostname of the instance",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tls_secret": schema.StringAttribute{
									Description:         "Secret where the Ingress TLS secret can be found",
									MarkdownDescription: "Secret where the Ingress TLS secret can be found",
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

					"ingress_path": schema.StringAttribute{
						Description:         "The ingress path used to reach the deployed service",
						MarkdownDescription: "The ingress path used to reach the deployed service",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_path_type": schema.StringAttribute{
						Description:         "The ingress path type for the deployed service",
						MarkdownDescription: "The ingress path type for the deployed service",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_tls_secret": schema.StringAttribute{
						Description:         "(Deprecated) Secret where the Ingress TLS secret can be found",
						MarkdownDescription: "(Deprecated) Secret where the Ingress TLS secret can be found",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ingress_type": schema.StringAttribute{
						Description:         "The ingress type to use to reach the deployed instance",
						MarkdownDescription: "The ingress type to use to reach the deployed instance",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("none", "Ingress", "ingress", "Route", "route"),
						},
					},

					"init_container_extra_commands": schema.StringAttribute{
						Description:         "Extra commands for the init container",
						MarkdownDescription: "Extra commands for the init container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"init_container_extra_volume_mounts": schema.StringAttribute{
						Description:         "Specify volume mounts to be added to the init container",
						MarkdownDescription: "Specify volume mounts to be added to the init container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"init_container_image": schema.StringAttribute{
						Description:         "Registry path to the init container to use",
						MarkdownDescription: "Registry path to the init container to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"init_container_image_version": schema.StringAttribute{
						Description:         "Init container image version to use",
						MarkdownDescription: "Init container image version to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"init_container_resource_requirements": schema.SingleNestedAttribute{
						Description:         "Resource requirements for the init container",
						MarkdownDescription: "Resource requirements for the init container",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

					"init_projects_container_image": schema.StringAttribute{
						Description:         "Registry path to the init projects container to use",
						MarkdownDescription: "Registry path to the init projects container to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ipv6_disabled": schema.BoolAttribute{
						Description:         "Disable web container's nginx ipv6 listener",
						MarkdownDescription: "Disable web container's nginx ipv6 listener",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kind": schema.StringAttribute{
						Description:         "Kind of the deployment type",
						MarkdownDescription: "Kind of the deployment type",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ldap_cacert_secret": schema.StringAttribute{
						Description:         "Secret where can be found the LDAP trusted Certificate Authority Bundle",
						MarkdownDescription: "Secret where can be found the LDAP trusted Certificate Authority Bundle",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ldap_password_secret": schema.StringAttribute{
						Description:         "Secret where can be found the LDAP bind password",
						MarkdownDescription: "Secret where can be found the LDAP bind password",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"loadbalancer_class": schema.StringAttribute{
						Description:         "Class of LoadBalancer to use",
						MarkdownDescription: "Class of LoadBalancer to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"loadbalancer_ip": schema.StringAttribute{
						Description:         "Assign LoadBalancer IP address",
						MarkdownDescription: "Assign LoadBalancer IP address",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"loadbalancer_port": schema.Int64Attribute{
						Description:         "Port to use for the loadbalancer",
						MarkdownDescription: "Port to use for the loadbalancer",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"loadbalancer_protocol": schema.StringAttribute{
						Description:         "Protocol to use for the loadbalancer",
						MarkdownDescription: "Protocol to use for the loadbalancer",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("http", "https"),
						},
					},

					"metrics_utility_configmap": schema.StringAttribute{
						Description:         "Metrics-Utility ConfigMap",
						MarkdownDescription: "Metrics-Utility ConfigMap",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_console_enabled": schema.BoolAttribute{
						Description:         "Enable metrics utility shipping to Red Hat Hybrid Cloud Console",
						MarkdownDescription: "Enable metrics utility shipping to Red Hat Hybrid Cloud Console",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_cronjob_gather_schedule": schema.StringAttribute{
						Description:         "Metrics-Utility Gather Data CronJob Schedule",
						MarkdownDescription: "Metrics-Utility Gather Data CronJob Schedule",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_cronjob_report_schedule": schema.StringAttribute{
						Description:         "Metrics-Utility Report CronJob Schedule",
						MarkdownDescription: "Metrics-Utility Report CronJob Schedule",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_enabled": schema.BoolAttribute{
						Description:         "Enable metrics utility",
						MarkdownDescription: "Enable metrics utility",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_image": schema.StringAttribute{
						Description:         "Metrics-Utility Image",
						MarkdownDescription: "Metrics-Utility Image",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_image_pull_policy": schema.StringAttribute{
						Description:         "Metrics-Utility Image PullPolicy",
						MarkdownDescription: "Metrics-Utility Image PullPolicy",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_image_version": schema.StringAttribute{
						Description:         "Metrics-Utility Image Version",
						MarkdownDescription: "Metrics-Utility Image Version",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_pvc_claim": schema.StringAttribute{
						Description:         "Metrics-Utility PVC Claim",
						MarkdownDescription: "Metrics-Utility PVC Claim",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_pvc_claim_size": schema.StringAttribute{
						Description:         "Metrics-Utility PVC Claim Size",
						MarkdownDescription: "Metrics-Utility PVC Claim Size",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_pvc_claim_storage_class": schema.StringAttribute{
						Description:         "Metrics-Utility PVC Claim Storage Class",
						MarkdownDescription: "Metrics-Utility PVC Claim Storage Class",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_secret": schema.StringAttribute{
						Description:         "Metrics-Utility Secret",
						MarkdownDescription: "Metrics-Utility Secret",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"metrics_utility_ship_target": schema.StringAttribute{
						Description:         "Metrics-Utility Ship Target",
						MarkdownDescription: "Metrics-Utility Ship Target",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nginx_listen_queue_size": schema.Int64Attribute{
						Description:         "Set the socket listen queue size for nginx (defaults to same as uwsgi)",
						MarkdownDescription: "Set the socket listen queue size for nginx (defaults to same as uwsgi)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nginx_worker_connections": schema.Int64Attribute{
						Description:         "Set the number of connections per worker for nginx",
						MarkdownDescription: "Set the number of connections per worker for nginx",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nginx_worker_cpu_affinity": schema.StringAttribute{
						Description:         "Set the CPU affinity for nginx workers",
						MarkdownDescription: "Set the CPU affinity for nginx workers",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nginx_worker_processes": schema.Int64Attribute{
						Description:         "Set the number of workers for nginx",
						MarkdownDescription: "Set the number of workers for nginx",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"no_log": schema.BoolAttribute{
						Description:         "Configure no_log for no_log tasks",
						MarkdownDescription: "Configure no_log for no_log tasks",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"node_selector": schema.StringAttribute{
						Description:         "nodeSelector for the pods",
						MarkdownDescription: "nodeSelector for the pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"nodeport_port": schema.Int64Attribute{
						Description:         "Port to use for the nodeport",
						MarkdownDescription: "Port to use for the nodeport",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"old_postgres_configuration_secret": schema.StringAttribute{
						Description:         "Secret where the old database configuration can be found for data migration",
						MarkdownDescription: "Secret where the old database configuration can be found for data migration",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][-a-zA-Z0-9]{0,253}[a-zA-Z0-9]$`), ""),
						},
					},

					"postgres_configuration_secret": schema.StringAttribute{
						Description:         "Secret where the database configuration can be found",
						MarkdownDescription: "Secret where the database configuration can be found",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_data_volume_init": schema.BoolAttribute{
						Description:         "Sets permissions on the /var/lib/pgdata/data for postgres container using an init container (not Openshift)",
						MarkdownDescription: "Sets permissions on the /var/lib/pgdata/data for postgres container using an init container (not Openshift)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_extra_args": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_extra_volume_mounts": schema.StringAttribute{
						Description:         "Specify volume mounts to be added to Postgres container",
						MarkdownDescription: "Specify volume mounts to be added to Postgres container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_extra_volumes": schema.StringAttribute{
						Description:         "Specify extra volumes to add to the application pod",
						MarkdownDescription: "Specify extra volumes to add to the application pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_image": schema.StringAttribute{
						Description:         "Registry path to the PostgreSQL container to use",
						MarkdownDescription: "Registry path to the PostgreSQL container to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_image_version": schema.StringAttribute{
						Description:         "PostgreSQL container image version to use",
						MarkdownDescription: "PostgreSQL container image version to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_init_container_commands": schema.StringAttribute{
						Description:         "Customize the postgres init container commands (Non Openshift)",
						MarkdownDescription: "Customize the postgres init container commands (Non Openshift)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_init_container_resource_requirements": schema.SingleNestedAttribute{
						Description:         "(Deprecated, use postgres_resource_requirements parameter) Resource requirements for the postgres init container",
						MarkdownDescription: "(Deprecated, use postgres_resource_requirements parameter) Resource requirements for the postgres init container",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

					"postgres_keep_pvc_after_upgrade": schema.BoolAttribute{
						Description:         "Specify whether or not to keep the old PVC after PostgreSQL upgrades",
						MarkdownDescription: "Specify whether or not to keep the old PVC after PostgreSQL upgrades",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_keepalives": schema.BoolAttribute{
						Description:         "Controls whether client-side TCP keepalives are used for Postgres connections.",
						MarkdownDescription: "Controls whether client-side TCP keepalives are used for Postgres connections.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_keepalives_count": schema.Int64Attribute{
						Description:         "Controls the number of TCP keepalives that can be lost before the client's connection to the server is considered dead.",
						MarkdownDescription: "Controls the number of TCP keepalives that can be lost before the client's connection to the server is considered dead.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_keepalives_idle": schema.Int64Attribute{
						Description:         "Controls the number of seconds of inactivity after which TCP should send a keepalive message to the server.",
						MarkdownDescription: "Controls the number of seconds of inactivity after which TCP should send a keepalive message to the server.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_keepalives_interval": schema.Int64Attribute{
						Description:         "Controls the number of seconds after which a TCP keepalive message that is not acknowledged by the server should be retransmitted.",
						MarkdownDescription: "Controls the number of seconds after which a TCP keepalive message that is not acknowledged by the server should be retransmitted.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_label_selector": schema.StringAttribute{
						Description:         "Label selector used to identify postgres pod for data migration",
						MarkdownDescription: "Label selector used to identify postgres pod for data migration",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_priority_class": schema.StringAttribute{
						Description:         "Assign a preexisting priority class to the postgres pod",
						MarkdownDescription: "Assign a preexisting priority class to the postgres pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_resource_requirements": schema.SingleNestedAttribute{
						Description:         "Resource requirements for the PostgreSQL container",
						MarkdownDescription: "Resource requirements for the PostgreSQL container",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
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

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
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

					"postgres_security_context_settings": schema.MapAttribute{
						Description:         "Key/values that will be set under the pod-level securityContext field",
						MarkdownDescription: "Key/values that will be set under the pod-level securityContext field",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_selector": schema.StringAttribute{
						Description:         "nodeSelector for the Postgres pods",
						MarkdownDescription: "nodeSelector for the Postgres pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_storage_class": schema.StringAttribute{
						Description:         "Storage class to use for the PostgreSQL PVC",
						MarkdownDescription: "Storage class to use for the PostgreSQL PVC",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_storage_requirements": schema.SingleNestedAttribute{
						Description:         "Storage requirements for the PostgreSQL container",
						MarkdownDescription: "Storage requirements for the PostgreSQL container",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"storage": schema.StringAttribute{
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

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"storage": schema.StringAttribute{
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

					"postgres_tolerations": schema.StringAttribute{
						Description:         "node tolerations for the Postgres pods",
						MarkdownDescription: "node tolerations for the Postgres pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"projects_existing_claim": schema.StringAttribute{
						Description:         "PersistentVolumeClaim to mount /var/lib/projects directory",
						MarkdownDescription: "PersistentVolumeClaim to mount /var/lib/projects directory",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"projects_persistence": schema.BoolAttribute{
						Description:         "Whether or not the /var/lib/projects directory will be persistent",
						MarkdownDescription: "Whether or not the /var/lib/projects directory will be persistent",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"projects_storage_access_mode": schema.StringAttribute{
						Description:         "AccessMode for the /var/lib/projects PersistentVolumeClaim",
						MarkdownDescription: "AccessMode for the /var/lib/projects PersistentVolumeClaim",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"projects_storage_class": schema.StringAttribute{
						Description:         "Storage class for the /var/lib/projects PersistentVolumeClaim",
						MarkdownDescription: "Storage class for the /var/lib/projects PersistentVolumeClaim",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"projects_storage_size": schema.StringAttribute{
						Description:         "Size for the /var/lib/projects PersistentVolumeClaim",
						MarkdownDescription: "Size for the /var/lib/projects PersistentVolumeClaim",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"projects_use_existing_claim": schema.StringAttribute{
						Description:         "Using existing PersistentVolumeClaim",
						MarkdownDescription: "Using existing PersistentVolumeClaim",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("_Yes_", "_No_"),
						},
					},

					"receptor_log_level": schema.StringAttribute{
						Description:         "Set log level of receptor service",
						MarkdownDescription: "Set log level of receptor service",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redis_capabilities": schema.ListAttribute{
						Description:         "Redis container capabilities",
						MarkdownDescription: "Redis container capabilities",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redis_image": schema.StringAttribute{
						Description:         "Registry path to the redis container to use",
						MarkdownDescription: "Registry path to the redis container to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redis_image_version": schema.StringAttribute{
						Description:         "Redis container image version to use",
						MarkdownDescription: "Redis container image version to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"redis_resource_requirements": schema.SingleNestedAttribute{
						Description:         "Resource requirements for the redis container",
						MarkdownDescription: "Resource requirements for the redis container",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

					"replicas": schema.Int64Attribute{
						Description:         "Number of instance replicas",
						MarkdownDescription: "Number of instance replicas",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_api_version": schema.StringAttribute{
						Description:         "The route API version to use",
						MarkdownDescription: "The route API version to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_host": schema.StringAttribute{
						Description:         "The DNS to use to points to the instance",
						MarkdownDescription: "The DNS to use to points to the instance",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_tls_secret": schema.StringAttribute{
						Description:         "Secret where the TLS related credentials are stored",
						MarkdownDescription: "Secret where the TLS related credentials are stored",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"route_tls_termination_mechanism": schema.StringAttribute{
						Description:         "The secure TLS termination mechanism to use",
						MarkdownDescription: "The secure TLS termination mechanism to use",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Edge", "edge", "Passthrough", "passthrough"),
						},
					},

					"rsyslog_args": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rsyslog_command": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rsyslog_extra_env": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rsyslog_extra_volume_mounts": schema.StringAttribute{
						Description:         "Specify volume mounts to be added to the Rsyslog container",
						MarkdownDescription: "Specify volume mounts to be added to the Rsyslog container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"rsyslog_resource_requirements": schema.SingleNestedAttribute{
						Description:         "Resource requirements for the rsyslog container",
						MarkdownDescription: "Resource requirements for the rsyslog container",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

					"secret_key_secret": schema.StringAttribute{
						Description:         "Secret where the secret key can be found",
						MarkdownDescription: "Secret where the secret key can be found",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(255),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-zA-Z0-9][-a-zA-Z0-9]{0,253}[a-zA-Z0-9]$`), ""),
						},
					},

					"security_context_settings": schema.MapAttribute{
						Description:         "Key/values that will be set under the pod-level securityContext field",
						MarkdownDescription: "Key/values that will be set under the pod-level securityContext field",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_account_annotations": schema.StringAttribute{
						Description:         "ServiceAccount annotations",
						MarkdownDescription: "ServiceAccount annotations",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_annotations": schema.StringAttribute{
						Description:         "Annotations to add to the service",
						MarkdownDescription: "Annotations to add to the service",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_labels": schema.StringAttribute{
						Description:         "Additional labels to apply to the service",
						MarkdownDescription: "Additional labels to apply to the service",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_type": schema.StringAttribute{
						Description:         "The service type to be used on the deployed instance",
						MarkdownDescription: "The service type to be used on the deployed instance",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("LoadBalancer", "loadbalancer", "ClusterIP", "clusterip", "NodePort", "nodeport"),
						},
					},

					"session_cookie_secure": schema.StringAttribute{
						Description:         "Set session cookie secure mode for web",
						MarkdownDescription: "Set session cookie secure mode for web",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"set_self_labels": schema.BoolAttribute{
						Description:         "Maintain some of the recommended 'app.kubernetes.io/*' labels on the resource (self)",
						MarkdownDescription: "Maintain some of the recommended 'app.kubernetes.io/*' labels on the resource (self)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_affinity": schema.SingleNestedAttribute{
						Description:         "If specified, the pod's scheduling constraints",
						MarkdownDescription: "If specified, the pod's scheduling constraints",
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

					"task_annotations": schema.StringAttribute{
						Description:         "Task deployment annotations. This will override the general annotations parameter for the Task deployment.",
						MarkdownDescription: "Task deployment annotations. This will override the general annotations parameter for the Task deployment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_args": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_command": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_extra_env": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_extra_volume_mounts": schema.StringAttribute{
						Description:         "Specify volume mounts to be added to Task container",
						MarkdownDescription: "Specify volume mounts to be added to Task container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_liveness_failure_threshold": schema.Int64Attribute{
						Description:         "Number of consecutive failure events to identify failure of task pod",
						MarkdownDescription: "Number of consecutive failure events to identify failure of task pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_liveness_initial_delay": schema.Int64Attribute{
						Description:         "Initial delay before starting liveness checks on task pod",
						MarkdownDescription: "Initial delay before starting liveness checks on task pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_liveness_period": schema.Int64Attribute{
						Description:         "Time period in seconds between each liveness check for the task pod",
						MarkdownDescription: "Time period in seconds between each liveness check for the task pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_liveness_timeout": schema.Int64Attribute{
						Description:         "Number of seconds to wait for a probe response from task pod",
						MarkdownDescription: "Number of seconds to wait for a probe response from task pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_node_selector": schema.StringAttribute{
						Description:         "nodeSelector for the task pods",
						MarkdownDescription: "nodeSelector for the task pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_privileged": schema.BoolAttribute{
						Description:         "If a privileged security context should be enabled",
						MarkdownDescription: "If a privileged security context should be enabled",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_readiness_failure_threshold": schema.Int64Attribute{
						Description:         "Number of consecutive failure events to identify failure of task pod",
						MarkdownDescription: "Number of consecutive failure events to identify failure of task pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_readiness_initial_delay": schema.Int64Attribute{
						Description:         "Initial delay before starting readiness checks on task pod",
						MarkdownDescription: "Initial delay before starting readiness checks on task pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_readiness_period": schema.Int64Attribute{
						Description:         "Time period in seconds between each readiness check for the task pod",
						MarkdownDescription: "Time period in seconds between each readiness check for the task pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_readiness_timeout": schema.Int64Attribute{
						Description:         "Number of seconds to wait for a probe response from task pod",
						MarkdownDescription: "Number of seconds to wait for a probe response from task pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_replicas": schema.Int64Attribute{
						Description:         "Number of task instance replicas",
						MarkdownDescription: "Number of task instance replicas",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_resource_requirements": schema.SingleNestedAttribute{
						Description:         "Resource requirements for the task container",
						MarkdownDescription: "Resource requirements for the task container",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

					"task_tolerations": schema.StringAttribute{
						Description:         "node tolerations for the task pods",
						MarkdownDescription: "node tolerations for the task pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"task_topology_spread_constraints": schema.StringAttribute{
						Description:         "topology rule(s) for the task pods",
						MarkdownDescription: "topology rule(s) for the task pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"termination_grace_period_seconds": schema.Int64Attribute{
						Description:         "Optional duration in seconds pods needs to terminate gracefully",
						MarkdownDescription: "Optional duration in seconds pods needs to terminate gracefully",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tolerations": schema.StringAttribute{
						Description:         "node tolerations for the pods",
						MarkdownDescription: "node tolerations for the pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"topology_spread_constraints": schema.StringAttribute{
						Description:         "topology rule(s) for the pods",
						MarkdownDescription: "topology rule(s) for the pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uwsgi_listen_queue_size": schema.Int64Attribute{
						Description:         "Set the socket listen queue size for uwsgi",
						MarkdownDescription: "Set the socket listen queue size for uwsgi",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"uwsgi_processes": schema.Int64Attribute{
						Description:         "Set the number of uwsgi processes to run in a web container",
						MarkdownDescription: "Set the number of uwsgi processes to run in a web container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_affinity": schema.SingleNestedAttribute{
						Description:         "If specified, the pod's scheduling constraints",
						MarkdownDescription: "If specified, the pod's scheduling constraints",
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

					"web_annotations": schema.StringAttribute{
						Description:         "Web deployment annotations. This will override the general annotations parameter for the Web deployment.",
						MarkdownDescription: "Web deployment annotations. This will override the general annotations parameter for the Web deployment.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_args": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_command": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_extra_env": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_extra_volume_mounts": schema.StringAttribute{
						Description:         "Specify volume mounts to be added to the Web container",
						MarkdownDescription: "Specify volume mounts to be added to the Web container",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_liveness_failure_threshold": schema.Int64Attribute{
						Description:         "Number of consecutive failure events to identify failure of web pod",
						MarkdownDescription: "Number of consecutive failure events to identify failure of web pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_liveness_initial_delay": schema.Int64Attribute{
						Description:         "Initial delay before starting liveness checks on web pod",
						MarkdownDescription: "Initial delay before starting liveness checks on web pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_liveness_period": schema.Int64Attribute{
						Description:         "Time period in seconds between each liveness check for the web pod",
						MarkdownDescription: "Time period in seconds between each liveness check for the web pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_liveness_timeout": schema.Int64Attribute{
						Description:         "Number of seconds to wait for a probe response from web pod",
						MarkdownDescription: "Number of seconds to wait for a probe response from web pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_node_selector": schema.StringAttribute{
						Description:         "nodeSelector for the web pods",
						MarkdownDescription: "nodeSelector for the web pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_readiness_failure_threshold": schema.Int64Attribute{
						Description:         "Number of consecutive failure events to identify failure of web pod",
						MarkdownDescription: "Number of consecutive failure events to identify failure of web pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_readiness_initial_delay": schema.Int64Attribute{
						Description:         "Initial delay before starting readiness checks on web pod",
						MarkdownDescription: "Initial delay before starting readiness checks on web pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_readiness_period": schema.Int64Attribute{
						Description:         "Time period in seconds between each readiness check for the web pod",
						MarkdownDescription: "Time period in seconds between each readiness check for the web pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_readiness_timeout": schema.Int64Attribute{
						Description:         "Number of seconds to wait for a probe response from web pod",
						MarkdownDescription: "Number of seconds to wait for a probe response from web pod",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_replicas": schema.Int64Attribute{
						Description:         "Number of web instance replicas",
						MarkdownDescription: "Number of web instance replicas",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_resource_requirements": schema.SingleNestedAttribute{
						Description:         "Resource requirements for the web container",
						MarkdownDescription: "Resource requirements for the web container",
						Attributes: map[string]schema.Attribute{
							"limits": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

							"requests": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"cpu": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ephemeral_storage": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storage": schema.StringAttribute{
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

					"web_tolerations": schema.StringAttribute{
						Description:         "node tolerations for the web pods",
						MarkdownDescription: "node tolerations for the web pods",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"web_topology_spread_constraints": schema.StringAttribute{
						Description:         "topology rule(s) for the web pods",
						MarkdownDescription: "topology rule(s) for the web pods",
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

func (r *AwxAnsibleComAwxV1Beta1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_awx_ansible_com_awx_v1beta1_manifest")

	var model AwxAnsibleComAwxV1Beta1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("awx.ansible.com/v1beta1")
	model.Kind = pointer.String("AWX")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
