/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package acid_zalan_do_v1

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
	"regexp"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &AcidZalanDoOperatorConfigurationV1Manifest{}
)

func NewAcidZalanDoOperatorConfigurationV1Manifest() datasource.DataSource {
	return &AcidZalanDoOperatorConfigurationV1Manifest{}
}

type AcidZalanDoOperatorConfigurationV1Manifest struct{}

type AcidZalanDoOperatorConfigurationV1ManifestData struct {
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Configuration *struct {
		Aws_or_gcp *struct {
			Additional_secret_mount           *string `tfsdk:"additional_secret_mount" json:"additional_secret_mount,omitempty"`
			Additional_secret_mount_path      *string `tfsdk:"additional_secret_mount_path" json:"additional_secret_mount_path,omitempty"`
			Aws_region                        *string `tfsdk:"aws_region" json:"aws_region,omitempty"`
			Enable_ebs_gp3_migration          *bool   `tfsdk:"enable_ebs_gp3_migration" json:"enable_ebs_gp3_migration,omitempty"`
			Enable_ebs_gp3_migration_max_size *int64  `tfsdk:"enable_ebs_gp3_migration_max_size" json:"enable_ebs_gp3_migration_max_size,omitempty"`
			Gcp_credentials                   *string `tfsdk:"gcp_credentials" json:"gcp_credentials,omitempty"`
			Kube_iam_role                     *string `tfsdk:"kube_iam_role" json:"kube_iam_role,omitempty"`
			Log_s3_bucket                     *string `tfsdk:"log_s3_bucket" json:"log_s3_bucket,omitempty"`
			Wal_az_storage_account            *string `tfsdk:"wal_az_storage_account" json:"wal_az_storage_account,omitempty"`
			Wal_gs_bucket                     *string `tfsdk:"wal_gs_bucket" json:"wal_gs_bucket,omitempty"`
			Wal_s3_bucket                     *string `tfsdk:"wal_s3_bucket" json:"wal_s3_bucket,omitempty"`
		} `tfsdk:"aws_or_gcp" json:"aws_or_gcp,omitempty"`
		Connection_pooler *struct {
			Connection_pooler_default_cpu_limit      *string `tfsdk:"connection_pooler_default_cpu_limit" json:"connection_pooler_default_cpu_limit,omitempty"`
			Connection_pooler_default_cpu_request    *string `tfsdk:"connection_pooler_default_cpu_request" json:"connection_pooler_default_cpu_request,omitempty"`
			Connection_pooler_default_memory_limit   *string `tfsdk:"connection_pooler_default_memory_limit" json:"connection_pooler_default_memory_limit,omitempty"`
			Connection_pooler_default_memory_request *string `tfsdk:"connection_pooler_default_memory_request" json:"connection_pooler_default_memory_request,omitempty"`
			Connection_pooler_image                  *string `tfsdk:"connection_pooler_image" json:"connection_pooler_image,omitempty"`
			Connection_pooler_max_db_connections     *int64  `tfsdk:"connection_pooler_max_db_connections" json:"connection_pooler_max_db_connections,omitempty"`
			Connection_pooler_mode                   *string `tfsdk:"connection_pooler_mode" json:"connection_pooler_mode,omitempty"`
			Connection_pooler_number_of_instances    *int64  `tfsdk:"connection_pooler_number_of_instances" json:"connection_pooler_number_of_instances,omitempty"`
			Connection_pooler_schema                 *string `tfsdk:"connection_pooler_schema" json:"connection_pooler_schema,omitempty"`
			Connection_pooler_user                   *string `tfsdk:"connection_pooler_user" json:"connection_pooler_user,omitempty"`
		} `tfsdk:"connection_pooler" json:"connection_pooler,omitempty"`
		Crd_categories *[]string `tfsdk:"crd_categories" json:"crd_categories,omitempty"`
		Debug          *struct {
			Debug_logging          *bool `tfsdk:"debug_logging" json:"debug_logging,omitempty"`
			Enable_database_access *bool `tfsdk:"enable_database_access" json:"enable_database_access,omitempty"`
		} `tfsdk:"debug" json:"debug,omitempty"`
		Docker_image                          *string `tfsdk:"docker_image" json:"docker_image,omitempty"`
		Enable_crd_registration               *bool   `tfsdk:"enable_crd_registration" json:"enable_crd_registration,omitempty"`
		Enable_crd_validation                 *bool   `tfsdk:"enable_crd_validation" json:"enable_crd_validation,omitempty"`
		Enable_lazy_spilo_upgrade             *bool   `tfsdk:"enable_lazy_spilo_upgrade" json:"enable_lazy_spilo_upgrade,omitempty"`
		Enable_pgversion_env_var              *bool   `tfsdk:"enable_pgversion_env_var" json:"enable_pgversion_env_var,omitempty"`
		Enable_shm_volume                     *bool   `tfsdk:"enable_shm_volume" json:"enable_shm_volume,omitempty"`
		Enable_spilo_wal_path_compat          *bool   `tfsdk:"enable_spilo_wal_path_compat" json:"enable_spilo_wal_path_compat,omitempty"`
		Enable_team_id_clustername_prefix     *bool   `tfsdk:"enable_team_id_clustername_prefix" json:"enable_team_id_clustername_prefix,omitempty"`
		Etcd_host                             *string `tfsdk:"etcd_host" json:"etcd_host,omitempty"`
		Ignore_instance_limits_annotation_key *string `tfsdk:"ignore_instance_limits_annotation_key" json:"ignore_instance_limits_annotation_key,omitempty"`
		Kubernetes                            *struct {
			Additional_pod_capabilities             *[]string          `tfsdk:"additional_pod_capabilities" json:"additional_pod_capabilities,omitempty"`
			Cluster_domain                          *string            `tfsdk:"cluster_domain" json:"cluster_domain,omitempty"`
			Cluster_labels                          *map[string]string `tfsdk:"cluster_labels" json:"cluster_labels,omitempty"`
			Cluster_name_label                      *string            `tfsdk:"cluster_name_label" json:"cluster_name_label,omitempty"`
			Custom_pod_annotations                  *map[string]string `tfsdk:"custom_pod_annotations" json:"custom_pod_annotations,omitempty"`
			Delete_annotation_date_key              *string            `tfsdk:"delete_annotation_date_key" json:"delete_annotation_date_key,omitempty"`
			Delete_annotation_name_key              *string            `tfsdk:"delete_annotation_name_key" json:"delete_annotation_name_key,omitempty"`
			Downscaler_annotations                  *[]string          `tfsdk:"downscaler_annotations" json:"downscaler_annotations,omitempty"`
			Enable_cross_namespace_secret           *bool              `tfsdk:"enable_cross_namespace_secret" json:"enable_cross_namespace_secret,omitempty"`
			Enable_finalizers                       *bool              `tfsdk:"enable_finalizers" json:"enable_finalizers,omitempty"`
			Enable_init_containers                  *bool              `tfsdk:"enable_init_containers" json:"enable_init_containers,omitempty"`
			Enable_owner_references                 *bool              `tfsdk:"enable_owner_references" json:"enable_owner_references,omitempty"`
			Enable_persistent_volume_claim_deletion *bool              `tfsdk:"enable_persistent_volume_claim_deletion" json:"enable_persistent_volume_claim_deletion,omitempty"`
			Enable_pod_antiaffinity                 *bool              `tfsdk:"enable_pod_antiaffinity" json:"enable_pod_antiaffinity,omitempty"`
			Enable_pod_disruption_budget            *bool              `tfsdk:"enable_pod_disruption_budget" json:"enable_pod_disruption_budget,omitempty"`
			Enable_readiness_probe                  *bool              `tfsdk:"enable_readiness_probe" json:"enable_readiness_probe,omitempty"`
			Enable_secrets_deletion                 *bool              `tfsdk:"enable_secrets_deletion" json:"enable_secrets_deletion,omitempty"`
			Enable_sidecars                         *bool              `tfsdk:"enable_sidecars" json:"enable_sidecars,omitempty"`
			Ignored_annotations                     *[]string          `tfsdk:"ignored_annotations" json:"ignored_annotations,omitempty"`
			Infrastructure_roles_secret_name        *string            `tfsdk:"infrastructure_roles_secret_name" json:"infrastructure_roles_secret_name,omitempty"`
			Infrastructure_roles_secrets            *[]struct {
				Defaultrolevalue *string `tfsdk:"defaultrolevalue" json:"defaultrolevalue,omitempty"`
				Defaultuservalue *string `tfsdk:"defaultuservalue" json:"defaultuservalue,omitempty"`
				Details          *string `tfsdk:"details" json:"details,omitempty"`
				Passwordkey      *string `tfsdk:"passwordkey" json:"passwordkey,omitempty"`
				Rolekey          *string `tfsdk:"rolekey" json:"rolekey,omitempty"`
				Secretname       *string `tfsdk:"secretname" json:"secretname,omitempty"`
				Template         *bool   `tfsdk:"template" json:"template,omitempty"`
				Userkey          *string `tfsdk:"userkey" json:"userkey,omitempty"`
			} `tfsdk:"infrastructure_roles_secrets" json:"infrastructure_roles_secrets,omitempty"`
			Inherited_annotations                    *[]string          `tfsdk:"inherited_annotations" json:"inherited_annotations,omitempty"`
			Inherited_labels                         *[]string          `tfsdk:"inherited_labels" json:"inherited_labels,omitempty"`
			Master_pod_move_timeout                  *string            `tfsdk:"master_pod_move_timeout" json:"master_pod_move_timeout,omitempty"`
			Node_readiness_label                     *map[string]string `tfsdk:"node_readiness_label" json:"node_readiness_label,omitempty"`
			Node_readiness_label_merge               *string            `tfsdk:"node_readiness_label_merge" json:"node_readiness_label_merge,omitempty"`
			Oauth_token_secret_name                  *string            `tfsdk:"oauth_token_secret_name" json:"oauth_token_secret_name,omitempty"`
			Pdb_master_label_selector                *bool              `tfsdk:"pdb_master_label_selector" json:"pdb_master_label_selector,omitempty"`
			Pdb_name_format                          *string            `tfsdk:"pdb_name_format" json:"pdb_name_format,omitempty"`
			Persistent_volume_claim_retention_policy *struct {
				When_deleted *string `tfsdk:"when_deleted" json:"when_deleted,omitempty"`
				When_scaled  *string `tfsdk:"when_scaled" json:"when_scaled,omitempty"`
			} `tfsdk:"persistent_volume_claim_retention_policy" json:"persistent_volume_claim_retention_policy,omitempty"`
			Pod_antiaffinity_preferred_during_scheduling *bool              `tfsdk:"pod_antiaffinity_preferred_during_scheduling" json:"pod_antiaffinity_preferred_during_scheduling,omitempty"`
			Pod_antiaffinity_topology_key                *string            `tfsdk:"pod_antiaffinity_topology_key" json:"pod_antiaffinity_topology_key,omitempty"`
			Pod_environment_configmap                    *string            `tfsdk:"pod_environment_configmap" json:"pod_environment_configmap,omitempty"`
			Pod_environment_secret                       *string            `tfsdk:"pod_environment_secret" json:"pod_environment_secret,omitempty"`
			Pod_management_policy                        *string            `tfsdk:"pod_management_policy" json:"pod_management_policy,omitempty"`
			Pod_priority_class_name                      *string            `tfsdk:"pod_priority_class_name" json:"pod_priority_class_name,omitempty"`
			Pod_role_label                               *string            `tfsdk:"pod_role_label" json:"pod_role_label,omitempty"`
			Pod_service_account_definition               *string            `tfsdk:"pod_service_account_definition" json:"pod_service_account_definition,omitempty"`
			Pod_service_account_name                     *string            `tfsdk:"pod_service_account_name" json:"pod_service_account_name,omitempty"`
			Pod_service_account_role_binding_definition  *string            `tfsdk:"pod_service_account_role_binding_definition" json:"pod_service_account_role_binding_definition,omitempty"`
			Pod_terminate_grace_period                   *string            `tfsdk:"pod_terminate_grace_period" json:"pod_terminate_grace_period,omitempty"`
			Secret_name_template                         *string            `tfsdk:"secret_name_template" json:"secret_name_template,omitempty"`
			Share_pgsocket_with_sidecars                 *bool              `tfsdk:"share_pgsocket_with_sidecars" json:"share_pgsocket_with_sidecars,omitempty"`
			Spilo_allow_privilege_escalation             *bool              `tfsdk:"spilo_allow_privilege_escalation" json:"spilo_allow_privilege_escalation,omitempty"`
			Spilo_fsgroup                                *int64             `tfsdk:"spilo_fsgroup" json:"spilo_fsgroup,omitempty"`
			Spilo_privileged                             *bool              `tfsdk:"spilo_privileged" json:"spilo_privileged,omitempty"`
			Spilo_runasgroup                             *int64             `tfsdk:"spilo_runasgroup" json:"spilo_runasgroup,omitempty"`
			Spilo_runasuser                              *int64             `tfsdk:"spilo_runasuser" json:"spilo_runasuser,omitempty"`
			Storage_resize_mode                          *string            `tfsdk:"storage_resize_mode" json:"storage_resize_mode,omitempty"`
			Toleration                                   *map[string]string `tfsdk:"toleration" json:"toleration,omitempty"`
			Watched_namespace                            *string            `tfsdk:"watched_namespace" json:"watched_namespace,omitempty"`
		} `tfsdk:"kubernetes" json:"kubernetes,omitempty"`
		Kubernetes_use_configmaps *bool `tfsdk:"kubernetes_use_configmaps" json:"kubernetes_use_configmaps,omitempty"`
		Load_balancer             *struct {
			Custom_service_annotations          *map[string]string `tfsdk:"custom_service_annotations" json:"custom_service_annotations,omitempty"`
			Db_hosted_zone                      *string            `tfsdk:"db_hosted_zone" json:"db_hosted_zone,omitempty"`
			Enable_master_load_balancer         *bool              `tfsdk:"enable_master_load_balancer" json:"enable_master_load_balancer,omitempty"`
			Enable_master_pooler_load_balancer  *bool              `tfsdk:"enable_master_pooler_load_balancer" json:"enable_master_pooler_load_balancer,omitempty"`
			Enable_replica_load_balancer        *bool              `tfsdk:"enable_replica_load_balancer" json:"enable_replica_load_balancer,omitempty"`
			Enable_replica_pooler_load_balancer *bool              `tfsdk:"enable_replica_pooler_load_balancer" json:"enable_replica_pooler_load_balancer,omitempty"`
			External_traffic_policy             *string            `tfsdk:"external_traffic_policy" json:"external_traffic_policy,omitempty"`
			Master_dns_name_format              *string            `tfsdk:"master_dns_name_format" json:"master_dns_name_format,omitempty"`
			Master_legacy_dns_name_format       *string            `tfsdk:"master_legacy_dns_name_format" json:"master_legacy_dns_name_format,omitempty"`
			Replica_dns_name_format             *string            `tfsdk:"replica_dns_name_format" json:"replica_dns_name_format,omitempty"`
			Replica_legacy_dns_name_format      *string            `tfsdk:"replica_legacy_dns_name_format" json:"replica_legacy_dns_name_format,omitempty"`
		} `tfsdk:"load_balancer" json:"load_balancer,omitempty"`
		Logging_rest_api *struct {
			Api_port                *int64 `tfsdk:"api_port" json:"api_port,omitempty"`
			Cluster_history_entries *int64 `tfsdk:"cluster_history_entries" json:"cluster_history_entries,omitempty"`
			Ring_log_lines          *int64 `tfsdk:"ring_log_lines" json:"ring_log_lines,omitempty"`
		} `tfsdk:"logging_rest_api" json:"logging_rest_api,omitempty"`
		Logical_backup *struct {
			Logical_backup_azure_storage_account_key      *string `tfsdk:"logical_backup_azure_storage_account_key" json:"logical_backup_azure_storage_account_key,omitempty"`
			Logical_backup_azure_storage_account_name     *string `tfsdk:"logical_backup_azure_storage_account_name" json:"logical_backup_azure_storage_account_name,omitempty"`
			Logical_backup_azure_storage_container        *string `tfsdk:"logical_backup_azure_storage_container" json:"logical_backup_azure_storage_container,omitempty"`
			Logical_backup_cpu_limit                      *string `tfsdk:"logical_backup_cpu_limit" json:"logical_backup_cpu_limit,omitempty"`
			Logical_backup_cpu_request                    *string `tfsdk:"logical_backup_cpu_request" json:"logical_backup_cpu_request,omitempty"`
			Logical_backup_cronjob_environment_secret     *string `tfsdk:"logical_backup_cronjob_environment_secret" json:"logical_backup_cronjob_environment_secret,omitempty"`
			Logical_backup_docker_image                   *string `tfsdk:"logical_backup_docker_image" json:"logical_backup_docker_image,omitempty"`
			Logical_backup_google_application_credentials *string `tfsdk:"logical_backup_google_application_credentials" json:"logical_backup_google_application_credentials,omitempty"`
			Logical_backup_job_prefix                     *string `tfsdk:"logical_backup_job_prefix" json:"logical_backup_job_prefix,omitempty"`
			Logical_backup_memory_limit                   *string `tfsdk:"logical_backup_memory_limit" json:"logical_backup_memory_limit,omitempty"`
			Logical_backup_memory_request                 *string `tfsdk:"logical_backup_memory_request" json:"logical_backup_memory_request,omitempty"`
			Logical_backup_provider                       *string `tfsdk:"logical_backup_provider" json:"logical_backup_provider,omitempty"`
			Logical_backup_s3_access_key_id               *string `tfsdk:"logical_backup_s3_access_key_id" json:"logical_backup_s3_access_key_id,omitempty"`
			Logical_backup_s3_bucket                      *string `tfsdk:"logical_backup_s3_bucket" json:"logical_backup_s3_bucket,omitempty"`
			Logical_backup_s3_bucket_prefix               *string `tfsdk:"logical_backup_s3_bucket_prefix" json:"logical_backup_s3_bucket_prefix,omitempty"`
			Logical_backup_s3_endpoint                    *string `tfsdk:"logical_backup_s3_endpoint" json:"logical_backup_s3_endpoint,omitempty"`
			Logical_backup_s3_region                      *string `tfsdk:"logical_backup_s3_region" json:"logical_backup_s3_region,omitempty"`
			Logical_backup_s3_retention_time              *string `tfsdk:"logical_backup_s3_retention_time" json:"logical_backup_s3_retention_time,omitempty"`
			Logical_backup_s3_secret_access_key           *string `tfsdk:"logical_backup_s3_secret_access_key" json:"logical_backup_s3_secret_access_key,omitempty"`
			Logical_backup_s3_sse                         *string `tfsdk:"logical_backup_s3_sse" json:"logical_backup_s3_sse,omitempty"`
			Logical_backup_schedule                       *string `tfsdk:"logical_backup_schedule" json:"logical_backup_schedule,omitempty"`
		} `tfsdk:"logical_backup" json:"logical_backup,omitempty"`
		Major_version_upgrade *struct {
			Major_version_upgrade_mode            *string   `tfsdk:"major_version_upgrade_mode" json:"major_version_upgrade_mode,omitempty"`
			Major_version_upgrade_team_allow_list *[]string `tfsdk:"major_version_upgrade_team_allow_list" json:"major_version_upgrade_team_allow_list,omitempty"`
			Minimal_major_version                 *string   `tfsdk:"minimal_major_version" json:"minimal_major_version,omitempty"`
			Target_major_version                  *string   `tfsdk:"target_major_version" json:"target_major_version,omitempty"`
		} `tfsdk:"major_version_upgrade" json:"major_version_upgrade,omitempty"`
		Max_instances *int64 `tfsdk:"max_instances" json:"max_instances,omitempty"`
		Min_instances *int64 `tfsdk:"min_instances" json:"min_instances,omitempty"`
		Patroni       *struct {
			Enable_patroni_failsafe_mode *bool `tfsdk:"enable_patroni_failsafe_mode" json:"enable_patroni_failsafe_mode,omitempty"`
		} `tfsdk:"patroni" json:"patroni,omitempty"`
		Postgres_pod_resources *struct {
			Default_cpu_limit      *string `tfsdk:"default_cpu_limit" json:"default_cpu_limit,omitempty"`
			Default_cpu_request    *string `tfsdk:"default_cpu_request" json:"default_cpu_request,omitempty"`
			Default_memory_limit   *string `tfsdk:"default_memory_limit" json:"default_memory_limit,omitempty"`
			Default_memory_request *string `tfsdk:"default_memory_request" json:"default_memory_request,omitempty"`
			Max_cpu_request        *string `tfsdk:"max_cpu_request" json:"max_cpu_request,omitempty"`
			Max_memory_request     *string `tfsdk:"max_memory_request" json:"max_memory_request,omitempty"`
			Min_cpu_limit          *string `tfsdk:"min_cpu_limit" json:"min_cpu_limit,omitempty"`
			Min_memory_limit       *string `tfsdk:"min_memory_limit" json:"min_memory_limit,omitempty"`
		} `tfsdk:"postgres_pod_resources" json:"postgres_pod_resources,omitempty"`
		Repair_period *string `tfsdk:"repair_period" json:"repair_period,omitempty"`
		Resync_period *string `tfsdk:"resync_period" json:"resync_period,omitempty"`
		Scalyr        *struct {
			Scalyr_api_key        *string `tfsdk:"scalyr_api_key" json:"scalyr_api_key,omitempty"`
			Scalyr_cpu_limit      *string `tfsdk:"scalyr_cpu_limit" json:"scalyr_cpu_limit,omitempty"`
			Scalyr_cpu_request    *string `tfsdk:"scalyr_cpu_request" json:"scalyr_cpu_request,omitempty"`
			Scalyr_image          *string `tfsdk:"scalyr_image" json:"scalyr_image,omitempty"`
			Scalyr_memory_limit   *string `tfsdk:"scalyr_memory_limit" json:"scalyr_memory_limit,omitempty"`
			Scalyr_memory_request *string `tfsdk:"scalyr_memory_request" json:"scalyr_memory_request,omitempty"`
			Scalyr_server_url     *string `tfsdk:"scalyr_server_url" json:"scalyr_server_url,omitempty"`
		} `tfsdk:"scalyr" json:"scalyr,omitempty"`
		Set_memory_request_to_limit *bool                `tfsdk:"set_memory_request_to_limit" json:"set_memory_request_to_limit,omitempty"`
		Sidecar_docker_images       *map[string]string   `tfsdk:"sidecar_docker_images" json:"sidecar_docker_images,omitempty"`
		Sidecars                    *[]map[string]string `tfsdk:"sidecars" json:"sidecars,omitempty"`
		Teams_api                   *struct {
			Enable_admin_role_for_users         *bool              `tfsdk:"enable_admin_role_for_users" json:"enable_admin_role_for_users,omitempty"`
			Enable_postgres_team_crd            *bool              `tfsdk:"enable_postgres_team_crd" json:"enable_postgres_team_crd,omitempty"`
			Enable_postgres_team_crd_superusers *bool              `tfsdk:"enable_postgres_team_crd_superusers" json:"enable_postgres_team_crd_superusers,omitempty"`
			Enable_team_member_deprecation      *bool              `tfsdk:"enable_team_member_deprecation" json:"enable_team_member_deprecation,omitempty"`
			Enable_team_superuser               *bool              `tfsdk:"enable_team_superuser" json:"enable_team_superuser,omitempty"`
			Enable_teams_api                    *bool              `tfsdk:"enable_teams_api" json:"enable_teams_api,omitempty"`
			Pam_configuration                   *string            `tfsdk:"pam_configuration" json:"pam_configuration,omitempty"`
			Pam_role_name                       *string            `tfsdk:"pam_role_name" json:"pam_role_name,omitempty"`
			Postgres_superuser_teams            *[]string          `tfsdk:"postgres_superuser_teams" json:"postgres_superuser_teams,omitempty"`
			Protected_role_names                *[]string          `tfsdk:"protected_role_names" json:"protected_role_names,omitempty"`
			Role_deletion_suffix                *string            `tfsdk:"role_deletion_suffix" json:"role_deletion_suffix,omitempty"`
			Team_admin_role                     *string            `tfsdk:"team_admin_role" json:"team_admin_role,omitempty"`
			Team_api_role_configuration         *map[string]string `tfsdk:"team_api_role_configuration" json:"team_api_role_configuration,omitempty"`
			Teams_api_url                       *string            `tfsdk:"teams_api_url" json:"teams_api_url,omitempty"`
		} `tfsdk:"teams_api" json:"teams_api,omitempty"`
		Timeouts *struct {
			Patroni_api_check_interval *string `tfsdk:"patroni_api_check_interval" json:"patroni_api_check_interval,omitempty"`
			Patroni_api_check_timeout  *string `tfsdk:"patroni_api_check_timeout" json:"patroni_api_check_timeout,omitempty"`
			Pod_deletion_wait_timeout  *string `tfsdk:"pod_deletion_wait_timeout" json:"pod_deletion_wait_timeout,omitempty"`
			Pod_label_wait_timeout     *string `tfsdk:"pod_label_wait_timeout" json:"pod_label_wait_timeout,omitempty"`
			Ready_wait_interval        *string `tfsdk:"ready_wait_interval" json:"ready_wait_interval,omitempty"`
			Ready_wait_timeout         *string `tfsdk:"ready_wait_timeout" json:"ready_wait_timeout,omitempty"`
			Resource_check_interval    *string `tfsdk:"resource_check_interval" json:"resource_check_interval,omitempty"`
			Resource_check_timeout     *string `tfsdk:"resource_check_timeout" json:"resource_check_timeout,omitempty"`
		} `tfsdk:"timeouts" json:"timeouts,omitempty"`
		Users *struct {
			Additional_owner_roles           *[]string `tfsdk:"additional_owner_roles" json:"additional_owner_roles,omitempty"`
			Enable_password_rotation         *bool     `tfsdk:"enable_password_rotation" json:"enable_password_rotation,omitempty"`
			Password_rotation_interval       *int64    `tfsdk:"password_rotation_interval" json:"password_rotation_interval,omitempty"`
			Password_rotation_user_retention *int64    `tfsdk:"password_rotation_user_retention" json:"password_rotation_user_retention,omitempty"`
			Replication_username             *string   `tfsdk:"replication_username" json:"replication_username,omitempty"`
			Super_username                   *string   `tfsdk:"super_username" json:"super_username,omitempty"`
		} `tfsdk:"users" json:"users,omitempty"`
		Workers *int64 `tfsdk:"workers" json:"workers,omitempty"`
	} `tfsdk:"configuration" json:"configuration,omitempty"`
}

func (r *AcidZalanDoOperatorConfigurationV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_acid_zalan_do_operator_configuration_v1_manifest"
}

func (r *AcidZalanDoOperatorConfigurationV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
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

			"configuration": schema.SingleNestedAttribute{
				Description:         "",
				MarkdownDescription: "",
				Attributes: map[string]schema.Attribute{
					"aws_or_gcp": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"additional_secret_mount": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"additional_secret_mount_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"aws_region": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_ebs_gp3_migration": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_ebs_gp3_migration_max_size": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gcp_credentials": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kube_iam_role": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"log_s3_bucket": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wal_az_storage_account": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wal_gs_bucket": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"wal_s3_bucket": schema.StringAttribute{
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

					"connection_pooler": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"connection_pooler_default_cpu_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"connection_pooler_default_cpu_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"connection_pooler_default_memory_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"connection_pooler_default_memory_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"connection_pooler_image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"connection_pooler_max_db_connections": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"connection_pooler_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("session", "transaction"),
								},
							},

							"connection_pooler_number_of_instances": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"connection_pooler_schema": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"connection_pooler_user": schema.StringAttribute{
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

					"crd_categories": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"debug": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"debug_logging": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_database_access": schema.BoolAttribute{
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

					"docker_image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_crd_registration": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_crd_validation": schema.BoolAttribute{
						Description:         "deprecated",
						MarkdownDescription: "deprecated",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_lazy_spilo_upgrade": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_pgversion_env_var": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_shm_volume": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_spilo_wal_path_compat": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_team_id_clustername_prefix": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"etcd_host": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ignore_instance_limits_annotation_key": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"kubernetes": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"additional_pod_capabilities": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_domain": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_labels": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_name_label": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"custom_pod_annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delete_annotation_date_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"delete_annotation_name_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"downscaler_annotations": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_cross_namespace_secret": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_finalizers": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_init_containers": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_owner_references": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_persistent_volume_claim_deletion": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_pod_antiaffinity": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_pod_disruption_budget": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_readiness_probe": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_secrets_deletion": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_sidecars": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ignored_annotations": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"infrastructure_roles_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"infrastructure_roles_secrets": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"defaultrolevalue": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"defaultuservalue": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"details": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"passwordkey": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"rolekey": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secretname": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"template": schema.BoolAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"userkey": schema.StringAttribute{
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

							"inherited_annotations": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"inherited_labels": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"master_pod_move_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_readiness_label": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_readiness_label_merge": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("AND", "OR"),
								},
							},

							"oauth_token_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pdb_master_label_selector": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pdb_name_format": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"persistent_volume_claim_retention_policy": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"when_deleted": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("delete", "retain"),
										},
									},

									"when_scaled": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("delete", "retain"),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_antiaffinity_preferred_during_scheduling": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_antiaffinity_topology_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_environment_configmap": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_environment_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_management_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ordered_ready", "parallel"),
								},
							},

							"pod_priority_class_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_role_label": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_service_account_definition": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_service_account_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_service_account_role_binding_definition": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_terminate_grace_period": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name_template": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"share_pgsocket_with_sidecars": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spilo_allow_privilege_escalation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spilo_fsgroup": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spilo_privileged": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spilo_runasgroup": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"spilo_runasuser": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_resize_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("ebs", "mixed", "pvc", "off"),
								},
							},

							"toleration": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"watched_namespace": schema.StringAttribute{
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

					"kubernetes_use_configmaps": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"load_balancer": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"custom_service_annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"db_hosted_zone": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_master_load_balancer": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_master_pooler_load_balancer": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_replica_load_balancer": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_replica_pooler_load_balancer": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"external_traffic_policy": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Cluster", "Local"),
								},
							},

							"master_dns_name_format": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"master_legacy_dns_name_format": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica_dns_name_format": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replica_legacy_dns_name_format": schema.StringAttribute{
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

					"logging_rest_api": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"api_port": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"cluster_history_entries": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ring_log_lines": schema.Int64Attribute{
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

					"logical_backup": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"logical_backup_azure_storage_account_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_azure_storage_account_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_azure_storage_container": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_cpu_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"logical_backup_cpu_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"logical_backup_cronjob_environment_secret": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_docker_image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_google_application_credentials": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_job_prefix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_memory_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"logical_backup_memory_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"logical_backup_provider": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("az", "gcs", "s3"),
								},
							},

							"logical_backup_s3_access_key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_s3_bucket": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_s3_bucket_prefix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_s3_endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_s3_region": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_s3_retention_time": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_s3_secret_access_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_s3_sse": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"logical_backup_schedule": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+|\*)(/\d+)?(\s+(\d+|\*)(/\d+)?){4}$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"major_version_upgrade": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"major_version_upgrade_mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"major_version_upgrade_team_allow_list": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"minimal_major_version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"target_major_version": schema.StringAttribute{
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

					"max_instances": schema.Int64Attribute{
						Description:         "-1 = disabled",
						MarkdownDescription: "-1 = disabled",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},

					"min_instances": schema.Int64Attribute{
						Description:         "-1 = disabled",
						MarkdownDescription: "-1 = disabled",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(-1),
						},
					},

					"patroni": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"enable_patroni_failsafe_mode": schema.BoolAttribute{
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

					"postgres_pod_resources": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"default_cpu_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$|^$`), ""),
								},
							},

							"default_cpu_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$|^$`), ""),
								},
							},

							"default_memory_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$|^$`), ""),
								},
							},

							"default_memory_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$|^$`), ""),
								},
							},

							"max_cpu_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$|^$`), ""),
								},
							},

							"max_memory_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$|^$`), ""),
								},
							},

							"min_cpu_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$|^$`), ""),
								},
							},

							"min_memory_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$|^$`), ""),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"repair_period": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resync_period": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scalyr": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"scalyr_api_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scalyr_cpu_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"scalyr_cpu_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"scalyr_image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scalyr_memory_limit": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"scalyr_memory_request": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"scalyr_server_url": schema.StringAttribute{
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

					"set_memory_request_to_limit": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sidecar_docker_images": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sidecars": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"teams_api": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"enable_admin_role_for_users": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_postgres_team_crd": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_postgres_team_crd_superusers": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_team_member_deprecation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_team_superuser": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_teams_api": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pam_configuration": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pam_role_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"postgres_superuser_teams": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"protected_role_names": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"role_deletion_suffix": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"team_admin_role": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"team_api_role_configuration": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"teams_api_url": schema.StringAttribute{
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

					"timeouts": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"patroni_api_check_interval": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"patroni_api_check_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_deletion_wait_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_label_wait_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ready_wait_interval": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ready_wait_timeout": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_check_interval": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"resource_check_timeout": schema.StringAttribute{
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

					"users": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"additional_owner_roles": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_password_rotation": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password_rotation_interval": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"password_rotation_user_retention": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replication_username": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"super_username": schema.StringAttribute{
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

					"workers": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *AcidZalanDoOperatorConfigurationV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_acid_zalan_do_operator_configuration_v1_manifest")

	var model AcidZalanDoOperatorConfigurationV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("acid.zalan.do/v1")
	model.Kind = pointer.String("OperatorConfiguration")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
