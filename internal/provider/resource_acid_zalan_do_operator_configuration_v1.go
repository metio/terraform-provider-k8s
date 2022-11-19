/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

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

type AcidZalanDoOperatorConfigurationV1Resource struct{}

var (
	_ resource.Resource = (*AcidZalanDoOperatorConfigurationV1Resource)(nil)
)

type AcidZalanDoOperatorConfigurationV1TerraformModel struct {
	Id            types.Int64  `tfsdk:"id"`
	YAML          types.String `tfsdk:"yaml"`
	ApiVersion    types.String `tfsdk:"api_version"`
	Kind          types.String `tfsdk:"kind"`
	Metadata      types.Object `tfsdk:"metadata"`
	Configuration types.Object `tfsdk:"configuration"`
}

type AcidZalanDoOperatorConfigurationV1GoModel struct {
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

	Configuration *struct {
		Aws_or_gcp *struct {
			Additional_secret_mount *string `tfsdk:"additional_secret_mount" yaml:"additional_secret_mount,omitempty"`

			Additional_secret_mount_path *string `tfsdk:"additional_secret_mount_path" yaml:"additional_secret_mount_path,omitempty"`

			Aws_region *string `tfsdk:"aws_region" yaml:"aws_region,omitempty"`

			Enable_ebs_gp3_migration *bool `tfsdk:"enable_ebs_gp3_migration" yaml:"enable_ebs_gp3_migration,omitempty"`

			Enable_ebs_gp3_migration_max_size *int64 `tfsdk:"enable_ebs_gp3_migration_max_size" yaml:"enable_ebs_gp3_migration_max_size,omitempty"`

			Gcp_credentials *string `tfsdk:"gcp_credentials" yaml:"gcp_credentials,omitempty"`

			Kube_iam_role *string `tfsdk:"kube_iam_role" yaml:"kube_iam_role,omitempty"`

			Log_s3_bucket *string `tfsdk:"log_s3_bucket" yaml:"log_s3_bucket,omitempty"`

			Wal_az_storage_account *string `tfsdk:"wal_az_storage_account" yaml:"wal_az_storage_account,omitempty"`

			Wal_gs_bucket *string `tfsdk:"wal_gs_bucket" yaml:"wal_gs_bucket,omitempty"`

			Wal_s3_bucket *string `tfsdk:"wal_s3_bucket" yaml:"wal_s3_bucket,omitempty"`
		} `tfsdk:"aws_or_gcp" yaml:"aws_or_gcp,omitempty"`

		Connection_pooler *struct {
			Connection_pooler_default_cpu_limit *string `tfsdk:"connection_pooler_default_cpu_limit" yaml:"connection_pooler_default_cpu_limit,omitempty"`

			Connection_pooler_default_cpu_request *string `tfsdk:"connection_pooler_default_cpu_request" yaml:"connection_pooler_default_cpu_request,omitempty"`

			Connection_pooler_default_memory_limit *string `tfsdk:"connection_pooler_default_memory_limit" yaml:"connection_pooler_default_memory_limit,omitempty"`

			Connection_pooler_default_memory_request *string `tfsdk:"connection_pooler_default_memory_request" yaml:"connection_pooler_default_memory_request,omitempty"`

			Connection_pooler_image *string `tfsdk:"connection_pooler_image" yaml:"connection_pooler_image,omitempty"`

			Connection_pooler_max_db_connections *int64 `tfsdk:"connection_pooler_max_db_connections" yaml:"connection_pooler_max_db_connections,omitempty"`

			Connection_pooler_mode *string `tfsdk:"connection_pooler_mode" yaml:"connection_pooler_mode,omitempty"`

			Connection_pooler_number_of_instances *int64 `tfsdk:"connection_pooler_number_of_instances" yaml:"connection_pooler_number_of_instances,omitempty"`

			Connection_pooler_schema *string `tfsdk:"connection_pooler_schema" yaml:"connection_pooler_schema,omitempty"`

			Connection_pooler_user *string `tfsdk:"connection_pooler_user" yaml:"connection_pooler_user,omitempty"`
		} `tfsdk:"connection_pooler" yaml:"connection_pooler,omitempty"`

		Crd_categories *[]string `tfsdk:"crd_categories" yaml:"crd_categories,omitempty"`

		Debug *struct {
			Debug_logging *bool `tfsdk:"debug_logging" yaml:"debug_logging,omitempty"`

			Enable_database_access *bool `tfsdk:"enable_database_access" yaml:"enable_database_access,omitempty"`
		} `tfsdk:"debug" yaml:"debug,omitempty"`

		Docker_image *string `tfsdk:"docker_image" yaml:"docker_image,omitempty"`

		Enable_crd_registration *bool `tfsdk:"enable_crd_registration" yaml:"enable_crd_registration,omitempty"`

		Enable_crd_validation *bool `tfsdk:"enable_crd_validation" yaml:"enable_crd_validation,omitempty"`

		Enable_lazy_spilo_upgrade *bool `tfsdk:"enable_lazy_spilo_upgrade" yaml:"enable_lazy_spilo_upgrade,omitempty"`

		Enable_pgversion_env_var *bool `tfsdk:"enable_pgversion_env_var" yaml:"enable_pgversion_env_var,omitempty"`

		Enable_shm_volume *bool `tfsdk:"enable_shm_volume" yaml:"enable_shm_volume,omitempty"`

		Enable_spilo_wal_path_compat *bool `tfsdk:"enable_spilo_wal_path_compat" yaml:"enable_spilo_wal_path_compat,omitempty"`

		Enable_team_id_clustername_prefix *bool `tfsdk:"enable_team_id_clustername_prefix" yaml:"enable_team_id_clustername_prefix,omitempty"`

		Etcd_host *string `tfsdk:"etcd_host" yaml:"etcd_host,omitempty"`

		Ignore_instance_limits_annotation_key *string `tfsdk:"ignore_instance_limits_annotation_key" yaml:"ignore_instance_limits_annotation_key,omitempty"`

		Kubernetes *struct {
			Additional_pod_capabilities *[]string `tfsdk:"additional_pod_capabilities" yaml:"additional_pod_capabilities,omitempty"`

			Cluster_domain *string `tfsdk:"cluster_domain" yaml:"cluster_domain,omitempty"`

			Cluster_labels *map[string]string `tfsdk:"cluster_labels" yaml:"cluster_labels,omitempty"`

			Cluster_name_label *string `tfsdk:"cluster_name_label" yaml:"cluster_name_label,omitempty"`

			Custom_pod_annotations *map[string]string `tfsdk:"custom_pod_annotations" yaml:"custom_pod_annotations,omitempty"`

			Delete_annotation_date_key *string `tfsdk:"delete_annotation_date_key" yaml:"delete_annotation_date_key,omitempty"`

			Delete_annotation_name_key *string `tfsdk:"delete_annotation_name_key" yaml:"delete_annotation_name_key,omitempty"`

			Downscaler_annotations *[]string `tfsdk:"downscaler_annotations" yaml:"downscaler_annotations,omitempty"`

			Enable_cross_namespace_secret *bool `tfsdk:"enable_cross_namespace_secret" yaml:"enable_cross_namespace_secret,omitempty"`

			Enable_init_containers *bool `tfsdk:"enable_init_containers" yaml:"enable_init_containers,omitempty"`

			Enable_pod_antiaffinity *bool `tfsdk:"enable_pod_antiaffinity" yaml:"enable_pod_antiaffinity,omitempty"`

			Enable_pod_disruption_budget *bool `tfsdk:"enable_pod_disruption_budget" yaml:"enable_pod_disruption_budget,omitempty"`

			Enable_readiness_probe *bool `tfsdk:"enable_readiness_probe" yaml:"enable_readiness_probe,omitempty"`

			Enable_sidecars *bool `tfsdk:"enable_sidecars" yaml:"enable_sidecars,omitempty"`

			Ignored_annotations *[]string `tfsdk:"ignored_annotations" yaml:"ignored_annotations,omitempty"`

			Infrastructure_roles_secret_name *string `tfsdk:"infrastructure_roles_secret_name" yaml:"infrastructure_roles_secret_name,omitempty"`

			Infrastructure_roles_secrets *[]struct {
				Defaultrolevalue *string `tfsdk:"defaultrolevalue" yaml:"defaultrolevalue,omitempty"`

				Defaultuservalue *string `tfsdk:"defaultuservalue" yaml:"defaultuservalue,omitempty"`

				Details *string `tfsdk:"details" yaml:"details,omitempty"`

				Passwordkey *string `tfsdk:"passwordkey" yaml:"passwordkey,omitempty"`

				Rolekey *string `tfsdk:"rolekey" yaml:"rolekey,omitempty"`

				Secretname *string `tfsdk:"secretname" yaml:"secretname,omitempty"`

				Template *bool `tfsdk:"template" yaml:"template,omitempty"`

				Userkey *string `tfsdk:"userkey" yaml:"userkey,omitempty"`
			} `tfsdk:"infrastructure_roles_secrets" yaml:"infrastructure_roles_secrets,omitempty"`

			Inherited_annotations *[]string `tfsdk:"inherited_annotations" yaml:"inherited_annotations,omitempty"`

			Inherited_labels *[]string `tfsdk:"inherited_labels" yaml:"inherited_labels,omitempty"`

			Master_pod_move_timeout *string `tfsdk:"master_pod_move_timeout" yaml:"master_pod_move_timeout,omitempty"`

			Node_readiness_label *map[string]string `tfsdk:"node_readiness_label" yaml:"node_readiness_label,omitempty"`

			Node_readiness_label_merge *string `tfsdk:"node_readiness_label_merge" yaml:"node_readiness_label_merge,omitempty"`

			Oauth_token_secret_name *string `tfsdk:"oauth_token_secret_name" yaml:"oauth_token_secret_name,omitempty"`

			Pdb_name_format *string `tfsdk:"pdb_name_format" yaml:"pdb_name_format,omitempty"`

			Pod_antiaffinity_topology_key *string `tfsdk:"pod_antiaffinity_topology_key" yaml:"pod_antiaffinity_topology_key,omitempty"`

			Pod_environment_configmap *string `tfsdk:"pod_environment_configmap" yaml:"pod_environment_configmap,omitempty"`

			Pod_environment_secret *string `tfsdk:"pod_environment_secret" yaml:"pod_environment_secret,omitempty"`

			Pod_management_policy *string `tfsdk:"pod_management_policy" yaml:"pod_management_policy,omitempty"`

			Pod_priority_class_name *string `tfsdk:"pod_priority_class_name" yaml:"pod_priority_class_name,omitempty"`

			Pod_role_label *string `tfsdk:"pod_role_label" yaml:"pod_role_label,omitempty"`

			Pod_service_account_definition *string `tfsdk:"pod_service_account_definition" yaml:"pod_service_account_definition,omitempty"`

			Pod_service_account_name *string `tfsdk:"pod_service_account_name" yaml:"pod_service_account_name,omitempty"`

			Pod_service_account_role_binding_definition *string `tfsdk:"pod_service_account_role_binding_definition" yaml:"pod_service_account_role_binding_definition,omitempty"`

			Pod_terminate_grace_period *string `tfsdk:"pod_terminate_grace_period" yaml:"pod_terminate_grace_period,omitempty"`

			Secret_name_template *string `tfsdk:"secret_name_template" yaml:"secret_name_template,omitempty"`

			Spilo_allow_privilege_escalation *bool `tfsdk:"spilo_allow_privilege_escalation" yaml:"spilo_allow_privilege_escalation,omitempty"`

			Spilo_fsgroup *int64 `tfsdk:"spilo_fsgroup" yaml:"spilo_fsgroup,omitempty"`

			Spilo_privileged *bool `tfsdk:"spilo_privileged" yaml:"spilo_privileged,omitempty"`

			Spilo_runasgroup *int64 `tfsdk:"spilo_runasgroup" yaml:"spilo_runasgroup,omitempty"`

			Spilo_runasuser *int64 `tfsdk:"spilo_runasuser" yaml:"spilo_runasuser,omitempty"`

			Storage_resize_mode *string `tfsdk:"storage_resize_mode" yaml:"storage_resize_mode,omitempty"`

			Toleration *map[string]string `tfsdk:"toleration" yaml:"toleration,omitempty"`

			Watched_namespace *string `tfsdk:"watched_namespace" yaml:"watched_namespace,omitempty"`
		} `tfsdk:"kubernetes" yaml:"kubernetes,omitempty"`

		Kubernetes_use_configmaps *bool `tfsdk:"kubernetes_use_configmaps" yaml:"kubernetes_use_configmaps,omitempty"`

		Load_balancer *struct {
			Custom_service_annotations *map[string]string `tfsdk:"custom_service_annotations" yaml:"custom_service_annotations,omitempty"`

			Db_hosted_zone *string `tfsdk:"db_hosted_zone" yaml:"db_hosted_zone,omitempty"`

			Enable_master_load_balancer *bool `tfsdk:"enable_master_load_balancer" yaml:"enable_master_load_balancer,omitempty"`

			Enable_master_pooler_load_balancer *bool `tfsdk:"enable_master_pooler_load_balancer" yaml:"enable_master_pooler_load_balancer,omitempty"`

			Enable_replica_load_balancer *bool `tfsdk:"enable_replica_load_balancer" yaml:"enable_replica_load_balancer,omitempty"`

			Enable_replica_pooler_load_balancer *bool `tfsdk:"enable_replica_pooler_load_balancer" yaml:"enable_replica_pooler_load_balancer,omitempty"`

			External_traffic_policy *string `tfsdk:"external_traffic_policy" yaml:"external_traffic_policy,omitempty"`

			Master_dns_name_format *string `tfsdk:"master_dns_name_format" yaml:"master_dns_name_format,omitempty"`

			Replica_dns_name_format *string `tfsdk:"replica_dns_name_format" yaml:"replica_dns_name_format,omitempty"`
		} `tfsdk:"load_balancer" yaml:"load_balancer,omitempty"`

		Logging_rest_api *struct {
			Api_port *int64 `tfsdk:"api_port" yaml:"api_port,omitempty"`

			Cluster_history_entries *int64 `tfsdk:"cluster_history_entries" yaml:"cluster_history_entries,omitempty"`

			Ring_log_lines *int64 `tfsdk:"ring_log_lines" yaml:"ring_log_lines,omitempty"`
		} `tfsdk:"logging_rest_api" yaml:"logging_rest_api,omitempty"`

		Logical_backup *struct {
			Logical_backup_docker_image *string `tfsdk:"logical_backup_docker_image" yaml:"logical_backup_docker_image,omitempty"`

			Logical_backup_google_application_credentials *string `tfsdk:"logical_backup_google_application_credentials" yaml:"logical_backup_google_application_credentials,omitempty"`

			Logical_backup_job_prefix *string `tfsdk:"logical_backup_job_prefix" yaml:"logical_backup_job_prefix,omitempty"`

			Logical_backup_provider *string `tfsdk:"logical_backup_provider" yaml:"logical_backup_provider,omitempty"`

			Logical_backup_s3_access_key_id *string `tfsdk:"logical_backup_s3_access_key_id" yaml:"logical_backup_s3_access_key_id,omitempty"`

			Logical_backup_s3_bucket *string `tfsdk:"logical_backup_s3_bucket" yaml:"logical_backup_s3_bucket,omitempty"`

			Logical_backup_s3_endpoint *string `tfsdk:"logical_backup_s3_endpoint" yaml:"logical_backup_s3_endpoint,omitempty"`

			Logical_backup_s3_region *string `tfsdk:"logical_backup_s3_region" yaml:"logical_backup_s3_region,omitempty"`

			Logical_backup_s3_retention_time *string `tfsdk:"logical_backup_s3_retention_time" yaml:"logical_backup_s3_retention_time,omitempty"`

			Logical_backup_s3_secret_access_key *string `tfsdk:"logical_backup_s3_secret_access_key" yaml:"logical_backup_s3_secret_access_key,omitempty"`

			Logical_backup_s3_sse *string `tfsdk:"logical_backup_s3_sse" yaml:"logical_backup_s3_sse,omitempty"`

			Logical_backup_schedule *string `tfsdk:"logical_backup_schedule" yaml:"logical_backup_schedule,omitempty"`
		} `tfsdk:"logical_backup" yaml:"logical_backup,omitempty"`

		Major_version_upgrade *struct {
			Major_version_upgrade_mode *string `tfsdk:"major_version_upgrade_mode" yaml:"major_version_upgrade_mode,omitempty"`

			Major_version_upgrade_team_allow_list *[]string `tfsdk:"major_version_upgrade_team_allow_list" yaml:"major_version_upgrade_team_allow_list,omitempty"`

			Minimal_major_version *string `tfsdk:"minimal_major_version" yaml:"minimal_major_version,omitempty"`

			Target_major_version *string `tfsdk:"target_major_version" yaml:"target_major_version,omitempty"`
		} `tfsdk:"major_version_upgrade" yaml:"major_version_upgrade,omitempty"`

		Max_instances *int64 `tfsdk:"max_instances" yaml:"max_instances,omitempty"`

		Min_instances *int64 `tfsdk:"min_instances" yaml:"min_instances,omitempty"`

		Postgres_pod_resources *struct {
			Default_cpu_limit *string `tfsdk:"default_cpu_limit" yaml:"default_cpu_limit,omitempty"`

			Default_cpu_request *string `tfsdk:"default_cpu_request" yaml:"default_cpu_request,omitempty"`

			Default_memory_limit *string `tfsdk:"default_memory_limit" yaml:"default_memory_limit,omitempty"`

			Default_memory_request *string `tfsdk:"default_memory_request" yaml:"default_memory_request,omitempty"`

			Max_cpu_request *string `tfsdk:"max_cpu_request" yaml:"max_cpu_request,omitempty"`

			Max_memory_request *string `tfsdk:"max_memory_request" yaml:"max_memory_request,omitempty"`

			Min_cpu_limit *string `tfsdk:"min_cpu_limit" yaml:"min_cpu_limit,omitempty"`

			Min_memory_limit *string `tfsdk:"min_memory_limit" yaml:"min_memory_limit,omitempty"`
		} `tfsdk:"postgres_pod_resources" yaml:"postgres_pod_resources,omitempty"`

		Repair_period *string `tfsdk:"repair_period" yaml:"repair_period,omitempty"`

		Resync_period *string `tfsdk:"resync_period" yaml:"resync_period,omitempty"`

		Scalyr *struct {
			Scalyr_api_key *string `tfsdk:"scalyr_api_key" yaml:"scalyr_api_key,omitempty"`

			Scalyr_cpu_limit *string `tfsdk:"scalyr_cpu_limit" yaml:"scalyr_cpu_limit,omitempty"`

			Scalyr_cpu_request *string `tfsdk:"scalyr_cpu_request" yaml:"scalyr_cpu_request,omitempty"`

			Scalyr_image *string `tfsdk:"scalyr_image" yaml:"scalyr_image,omitempty"`

			Scalyr_memory_limit *string `tfsdk:"scalyr_memory_limit" yaml:"scalyr_memory_limit,omitempty"`

			Scalyr_memory_request *string `tfsdk:"scalyr_memory_request" yaml:"scalyr_memory_request,omitempty"`

			Scalyr_server_url *string `tfsdk:"scalyr_server_url" yaml:"scalyr_server_url,omitempty"`
		} `tfsdk:"scalyr" yaml:"scalyr,omitempty"`

		Set_memory_request_to_limit *bool `tfsdk:"set_memory_request_to_limit" yaml:"set_memory_request_to_limit,omitempty"`

		Sidecar_docker_images *map[string]string `tfsdk:"sidecar_docker_images" yaml:"sidecar_docker_images,omitempty"`

		Sidecars *[]map[string]string `tfsdk:"sidecars" yaml:"sidecars,omitempty"`

		Teams_api *struct {
			Enable_admin_role_for_users *bool `tfsdk:"enable_admin_role_for_users" yaml:"enable_admin_role_for_users,omitempty"`

			Enable_postgres_team_crd *bool `tfsdk:"enable_postgres_team_crd" yaml:"enable_postgres_team_crd,omitempty"`

			Enable_postgres_team_crd_superusers *bool `tfsdk:"enable_postgres_team_crd_superusers" yaml:"enable_postgres_team_crd_superusers,omitempty"`

			Enable_team_member_deprecation *bool `tfsdk:"enable_team_member_deprecation" yaml:"enable_team_member_deprecation,omitempty"`

			Enable_team_superuser *bool `tfsdk:"enable_team_superuser" yaml:"enable_team_superuser,omitempty"`

			Enable_teams_api *bool `tfsdk:"enable_teams_api" yaml:"enable_teams_api,omitempty"`

			Pam_configuration *string `tfsdk:"pam_configuration" yaml:"pam_configuration,omitempty"`

			Pam_role_name *string `tfsdk:"pam_role_name" yaml:"pam_role_name,omitempty"`

			Postgres_superuser_teams *[]string `tfsdk:"postgres_superuser_teams" yaml:"postgres_superuser_teams,omitempty"`

			Protected_role_names *[]string `tfsdk:"protected_role_names" yaml:"protected_role_names,omitempty"`

			Role_deletion_suffix *string `tfsdk:"role_deletion_suffix" yaml:"role_deletion_suffix,omitempty"`

			Team_admin_role *string `tfsdk:"team_admin_role" yaml:"team_admin_role,omitempty"`

			Team_api_role_configuration *map[string]string `tfsdk:"team_api_role_configuration" yaml:"team_api_role_configuration,omitempty"`

			Teams_api_url *string `tfsdk:"teams_api_url" yaml:"teams_api_url,omitempty"`
		} `tfsdk:"teams_api" yaml:"teams_api,omitempty"`

		Timeouts *struct {
			Patroni_api_check_interval *string `tfsdk:"patroni_api_check_interval" yaml:"patroni_api_check_interval,omitempty"`

			Patroni_api_check_timeout *string `tfsdk:"patroni_api_check_timeout" yaml:"patroni_api_check_timeout,omitempty"`

			Pod_deletion_wait_timeout *string `tfsdk:"pod_deletion_wait_timeout" yaml:"pod_deletion_wait_timeout,omitempty"`

			Pod_label_wait_timeout *string `tfsdk:"pod_label_wait_timeout" yaml:"pod_label_wait_timeout,omitempty"`

			Ready_wait_interval *string `tfsdk:"ready_wait_interval" yaml:"ready_wait_interval,omitempty"`

			Ready_wait_timeout *string `tfsdk:"ready_wait_timeout" yaml:"ready_wait_timeout,omitempty"`

			Resource_check_interval *string `tfsdk:"resource_check_interval" yaml:"resource_check_interval,omitempty"`

			Resource_check_timeout *string `tfsdk:"resource_check_timeout" yaml:"resource_check_timeout,omitempty"`
		} `tfsdk:"timeouts" yaml:"timeouts,omitempty"`

		Users *struct {
			Additional_owner_roles *[]string `tfsdk:"additional_owner_roles" yaml:"additional_owner_roles,omitempty"`

			Enable_password_rotation *bool `tfsdk:"enable_password_rotation" yaml:"enable_password_rotation,omitempty"`

			Password_rotation_interval *int64 `tfsdk:"password_rotation_interval" yaml:"password_rotation_interval,omitempty"`

			Password_rotation_user_retention *int64 `tfsdk:"password_rotation_user_retention" yaml:"password_rotation_user_retention,omitempty"`

			Replication_username *string `tfsdk:"replication_username" yaml:"replication_username,omitempty"`

			Super_username *string `tfsdk:"super_username" yaml:"super_username,omitempty"`
		} `tfsdk:"users" yaml:"users,omitempty"`

		Workers *int64 `tfsdk:"workers" yaml:"workers,omitempty"`
	} `tfsdk:"configuration" yaml:"configuration,omitempty"`
}

func NewAcidZalanDoOperatorConfigurationV1Resource() resource.Resource {
	return &AcidZalanDoOperatorConfigurationV1Resource{}
}

func (r *AcidZalanDoOperatorConfigurationV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_acid_zalan_do_operator_configuration_v1"
}

func (r *AcidZalanDoOperatorConfigurationV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
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

			"configuration": {
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"aws_or_gcp": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"additional_secret_mount": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"additional_secret_mount_path": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"aws_region": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_ebs_gp3_migration": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_ebs_gp3_migration_max_size": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gcp_credentials": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kube_iam_role": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_s3_bucket": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"wal_az_storage_account": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"wal_gs_bucket": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"wal_s3_bucket": {
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

					"connection_pooler": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"connection_pooler_default_cpu_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"connection_pooler_default_cpu_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"connection_pooler_default_memory_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"connection_pooler_default_memory_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"connection_pooler_image": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"connection_pooler_max_db_connections": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"connection_pooler_mode": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("session", "transaction"),
								},
							},

							"connection_pooler_number_of_instances": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"connection_pooler_schema": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"connection_pooler_user": {
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

					"crd_categories": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"debug": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"debug_logging": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_database_access": {
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

					"docker_image": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_crd_registration": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_crd_validation": {
						Description:         "deprecated",
						MarkdownDescription: "deprecated",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_lazy_spilo_upgrade": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_pgversion_env_var": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_shm_volume": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_spilo_wal_path_compat": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"enable_team_id_clustername_prefix": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"etcd_host": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"ignore_instance_limits_annotation_key": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"kubernetes": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"additional_pod_capabilities": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_domain": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_labels": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_name_label": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"custom_pod_annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"delete_annotation_date_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"delete_annotation_name_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"downscaler_annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_cross_namespace_secret": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_init_containers": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_pod_antiaffinity": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_pod_disruption_budget": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_readiness_probe": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_sidecars": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ignored_annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"infrastructure_roles_secret_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"infrastructure_roles_secrets": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"defaultrolevalue": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"defaultuservalue": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"details": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"passwordkey": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"rolekey": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"secretname": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"template": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.BoolType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"userkey": {
										Description:         "",
										MarkdownDescription: "",

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

							"inherited_annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"inherited_labels": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"master_pod_move_timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_readiness_label": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_readiness_label_merge": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("AND", "OR"),
								},
							},

							"oauth_token_secret_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pdb_name_format": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_antiaffinity_topology_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_environment_configmap": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_environment_secret": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_management_policy": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("ordered_ready", "parallel"),
								},
							},

							"pod_priority_class_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_role_label": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_service_account_definition": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_service_account_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_service_account_role_binding_definition": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_terminate_grace_period": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_name_template": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spilo_allow_privilege_escalation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spilo_fsgroup": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spilo_privileged": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spilo_runasgroup": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spilo_runasuser": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"storage_resize_mode": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("ebs", "mixed", "pvc", "off"),
								},
							},

							"toleration": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"watched_namespace": {
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

					"kubernetes_use_configmaps": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"load_balancer": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"custom_service_annotations": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"db_hosted_zone": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_master_load_balancer": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_master_pooler_load_balancer": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_replica_load_balancer": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_replica_pooler_load_balancer": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"external_traffic_policy": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("Cluster", "Local"),
								},
							},

							"master_dns_name_format": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replica_dns_name_format": {
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

					"logging_rest_api": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_port": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"cluster_history_entries": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ring_log_lines": {
								Description:         "",
								MarkdownDescription: "",

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

					"logical_backup": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"logical_backup_docker_image": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_google_application_credentials": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_job_prefix": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_provider": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_s3_access_key_id": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_s3_bucket": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_s3_endpoint": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_s3_region": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_s3_retention_time": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_s3_secret_access_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_s3_sse": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"logical_backup_schedule": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+|\*)(/\d+)?(\s+(\d+|\*)(/\d+)?){4}$`), ""),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"major_version_upgrade": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"major_version_upgrade_mode": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"major_version_upgrade_team_allow_list": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"minimal_major_version": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"target_major_version": {
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

					"max_instances": {
						Description:         "-1 = disabled",
						MarkdownDescription: "-1 = disabled",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(-1),
						},
					},

					"min_instances": {
						Description:         "-1 = disabled",
						MarkdownDescription: "-1 = disabled",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(-1),
						},
					},

					"postgres_pod_resources": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"default_cpu_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"default_cpu_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"default_memory_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"default_memory_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"max_cpu_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"max_memory_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"min_cpu_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"min_memory_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"repair_period": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"resync_period": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"scalyr": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"scalyr_api_key": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scalyr_cpu_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"scalyr_cpu_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
								},
							},

							"scalyr_image": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"scalyr_memory_limit": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"scalyr_memory_request": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"scalyr_server_url": {
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

					"set_memory_request_to_limit": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sidecar_docker_images": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sidecars": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.ListType{ElemType: types.MapType{ElemType: types.StringType}},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"teams_api": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enable_admin_role_for_users": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_postgres_team_crd": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_postgres_team_crd_superusers": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_team_member_deprecation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_team_superuser": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_teams_api": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pam_configuration": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pam_role_name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"postgres_superuser_teams": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"protected_role_names": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"role_deletion_suffix": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"team_admin_role": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"team_api_role_configuration": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"teams_api_url": {
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

					"timeouts": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"patroni_api_check_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"patroni_api_check_timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_deletion_wait_timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_label_wait_timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ready_wait_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ready_wait_timeout": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_check_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_check_timeout": {
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

					"users": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"additional_owner_roles": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_password_rotation": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password_rotation_interval": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"password_rotation_user_retention": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replication_username": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"super_username": {
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

					"workers": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							int64validator.AtLeast(1),
						},
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *AcidZalanDoOperatorConfigurationV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_acid_zalan_do_operator_configuration_v1")

	var state AcidZalanDoOperatorConfigurationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AcidZalanDoOperatorConfigurationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("acid.zalan.do/v1")
	goModel.Kind = utilities.Ptr("OperatorConfiguration")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *AcidZalanDoOperatorConfigurationV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_acid_zalan_do_operator_configuration_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *AcidZalanDoOperatorConfigurationV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_acid_zalan_do_operator_configuration_v1")

	var state AcidZalanDoOperatorConfigurationV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AcidZalanDoOperatorConfigurationV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("acid.zalan.do/v1")
	goModel.Kind = utilities.Ptr("OperatorConfiguration")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *AcidZalanDoOperatorConfigurationV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_acid_zalan_do_operator_configuration_v1")
	// NO-OP: Terraform removes the state automatically for us
}
