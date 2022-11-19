/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

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

type AppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*AppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource)(nil)
)

type AppRedislabsComRedisEnterpriseDatabaseV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type AppRedislabsComRedisEnterpriseDatabaseV1Alpha1GoModel struct {
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
		ActiveActiveName *string `tfsdk:"active_active_name" yaml:"activeActiveName,omitempty"`

		AlertSettings *struct {
			Bdb_backup_delayed *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_backup_delayed" yaml:"bdb_backup_delayed,omitempty"`

			Bdb_crdt_src_high_syncer_lag *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_crdt_src_high_syncer_lag" yaml:"bdb_crdt_src_high_syncer_lag,omitempty"`

			Bdb_crdt_src_syncer_connection_error *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_crdt_src_syncer_connection_error" yaml:"bdb_crdt_src_syncer_connection_error,omitempty"`

			Bdb_crdt_src_syncer_general_error *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_crdt_src_syncer_general_error" yaml:"bdb_crdt_src_syncer_general_error,omitempty"`

			Bdb_high_latency *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_high_latency" yaml:"bdb_high_latency,omitempty"`

			Bdb_high_throughput *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_high_throughput" yaml:"bdb_high_throughput,omitempty"`

			Bdb_long_running_action *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_long_running_action" yaml:"bdb_long_running_action,omitempty"`

			Bdb_low_throughput *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_low_throughput" yaml:"bdb_low_throughput,omitempty"`

			Bdb_ram_dataset_overhead *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_ram_dataset_overhead" yaml:"bdb_ram_dataset_overhead,omitempty"`

			Bdb_ram_values *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_ram_values" yaml:"bdb_ram_values,omitempty"`

			Bdb_replica_src_high_syncer_lag *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_replica_src_high_syncer_lag" yaml:"bdb_replica_src_high_syncer_lag,omitempty"`

			Bdb_replica_src_syncer_connection_error *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_replica_src_syncer_connection_error" yaml:"bdb_replica_src_syncer_connection_error,omitempty"`

			Bdb_shard_num_ram_values *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_shard_num_ram_values" yaml:"bdb_shard_num_ram_values,omitempty"`

			Bdb_size *struct {
				Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`
			} `tfsdk:"bdb_size" yaml:"bdb_size,omitempty"`
		} `tfsdk:"alert_settings" yaml:"alertSettings,omitempty"`

		Backup *struct {
			Abs *struct {
				AbsSecretName *string `tfsdk:"abs_secret_name" yaml:"absSecretName,omitempty"`

				Container *string `tfsdk:"container" yaml:"container,omitempty"`

				Subdir *string `tfsdk:"subdir" yaml:"subdir,omitempty"`
			} `tfsdk:"abs" yaml:"abs,omitempty"`

			Ftp *struct {
				Url *string `tfsdk:"url" yaml:"url,omitempty"`
			} `tfsdk:"ftp" yaml:"ftp,omitempty"`

			Gcs *struct {
				BucketName *string `tfsdk:"bucket_name" yaml:"bucketName,omitempty"`

				GcsSecretName *string `tfsdk:"gcs_secret_name" yaml:"gcsSecretName,omitempty"`

				Subdir *string `tfsdk:"subdir" yaml:"subdir,omitempty"`
			} `tfsdk:"gcs" yaml:"gcs,omitempty"`

			Interval *int64 `tfsdk:"interval" yaml:"interval,omitempty"`

			Mount *struct {
				Path *string `tfsdk:"path" yaml:"path,omitempty"`
			} `tfsdk:"mount" yaml:"mount,omitempty"`

			S3 *struct {
				AwsSecretName *string `tfsdk:"aws_secret_name" yaml:"awsSecretName,omitempty"`

				BucketName *string `tfsdk:"bucket_name" yaml:"bucketName,omitempty"`

				Subdir *string `tfsdk:"subdir" yaml:"subdir,omitempty"`
			} `tfsdk:"s3" yaml:"s3,omitempty"`

			Sftp *struct {
				SftpSecretName *string `tfsdk:"sftp_secret_name" yaml:"sftpSecretName,omitempty"`

				Sftp_url *string `tfsdk:"sftp_url" yaml:"sftp_url,omitempty"`
			} `tfsdk:"sftp" yaml:"sftp,omitempty"`

			Swift *struct {
				Auth_url *string `tfsdk:"auth_url" yaml:"auth_url,omitempty"`

				Container *string `tfsdk:"container" yaml:"container,omitempty"`

				Prefix *string `tfsdk:"prefix" yaml:"prefix,omitempty"`

				SwiftSecretName *string `tfsdk:"swift_secret_name" yaml:"swiftSecretName,omitempty"`
			} `tfsdk:"swift" yaml:"swift,omitempty"`
		} `tfsdk:"backup" yaml:"backup,omitempty"`

		ClientAuthenticationCertificates *[]string `tfsdk:"client_authentication_certificates" yaml:"clientAuthenticationCertificates,omitempty"`

		DataInternodeEncryption *bool `tfsdk:"data_internode_encryption" yaml:"dataInternodeEncryption,omitempty"`

		DatabasePort *int64 `tfsdk:"database_port" yaml:"databasePort,omitempty"`

		DatabaseSecretName *string `tfsdk:"database_secret_name" yaml:"databaseSecretName,omitempty"`

		DefaultUser *bool `tfsdk:"default_user" yaml:"defaultUser,omitempty"`

		EvictionPolicy *string `tfsdk:"eviction_policy" yaml:"evictionPolicy,omitempty"`

		GlobalConfigurations *bool `tfsdk:"global_configurations" yaml:"globalConfigurations,omitempty"`

		IsRof *bool `tfsdk:"is_rof" yaml:"isRof,omitempty"`

		MemcachedSaslSecretName *string `tfsdk:"memcached_sasl_secret_name" yaml:"memcachedSaslSecretName,omitempty"`

		MemorySize *string `tfsdk:"memory_size" yaml:"memorySize,omitempty"`

		ModulesList *[]struct {
			Config *string `tfsdk:"config" yaml:"config,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Version *string `tfsdk:"version" yaml:"version,omitempty"`
		} `tfsdk:"modules_list" yaml:"modulesList,omitempty"`

		OssCluster *bool `tfsdk:"oss_cluster" yaml:"ossCluster,omitempty"`

		Persistence *string `tfsdk:"persistence" yaml:"persistence,omitempty"`

		ProxyPolicy *string `tfsdk:"proxy_policy" yaml:"proxyPolicy,omitempty"`

		RackAware *bool `tfsdk:"rack_aware" yaml:"rackAware,omitempty"`

		RedisEnterpriseCluster *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"redis_enterprise_cluster" yaml:"redisEnterpriseCluster,omitempty"`

		RedisVersion *string `tfsdk:"redis_version" yaml:"redisVersion,omitempty"`

		ReplicaSources *[]struct {
			ClientKeySecret *string `tfsdk:"client_key_secret" yaml:"clientKeySecret,omitempty"`

			Compression *int64 `tfsdk:"compression" yaml:"compression,omitempty"`

			ReplicaSourceName *string `tfsdk:"replica_source_name" yaml:"replicaSourceName,omitempty"`

			ReplicaSourceType *string `tfsdk:"replica_source_type" yaml:"replicaSourceType,omitempty"`

			ServerCertSecret *string `tfsdk:"server_cert_secret" yaml:"serverCertSecret,omitempty"`

			TlsSniName *string `tfsdk:"tls_sni_name" yaml:"tlsSniName,omitempty"`
		} `tfsdk:"replica_sources" yaml:"replicaSources,omitempty"`

		Replication *bool `tfsdk:"replication" yaml:"replication,omitempty"`

		RofRamSize *string `tfsdk:"rof_ram_size" yaml:"rofRamSize,omitempty"`

		RolesPermissions *[]struct {
			Acl *string `tfsdk:"acl" yaml:"acl,omitempty"`

			Role *string `tfsdk:"role" yaml:"role,omitempty"`

			Type *string `tfsdk:"type" yaml:"type,omitempty"`
		} `tfsdk:"roles_permissions" yaml:"rolesPermissions,omitempty"`

		ShardCount *int64 `tfsdk:"shard_count" yaml:"shardCount,omitempty"`

		ShardsPlacement *string `tfsdk:"shards_placement" yaml:"shardsPlacement,omitempty"`

		TlsMode *string `tfsdk:"tls_mode" yaml:"tlsMode,omitempty"`

		Type *string `tfsdk:"type" yaml:"type,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewAppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource() resource.Resource {
	return &AppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource{}
}

func (r *AppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_app_redislabs_com_redis_enterprise_database_v1alpha1"
}

func (r *AppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "RedisEnterpriseDatabase is the Schema for the redisenterprisedatabases API",
		MarkdownDescription: "RedisEnterpriseDatabase is the Schema for the redisenterprisedatabases API",
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
				Description:         "RedisEnterpriseDatabaseSpec defines the desired state of RedisEnterpriseDatabase",
				MarkdownDescription: "RedisEnterpriseDatabaseSpec defines the desired state of RedisEnterpriseDatabase",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"active_active_name": {
						Description:         "The Redis Enterprise Active Active Peering custom resource name this Resource is associated with, also, the corresponding active active database name. In case this resource is created manually at the active active database creation this field must be filled via the user, otherwise, the operator will assign this field automatically. Note: this feature is currently unsupported.",
						MarkdownDescription: "The Redis Enterprise Active Active Peering custom resource name this Resource is associated with, also, the corresponding active active database name. In case this resource is created manually at the active active database creation this field must be filled via the user, otherwise, the operator will assign this field automatically. Note: this feature is currently unsupported.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"alert_settings": {
						Description:         "Settings for database alerts",
						MarkdownDescription: "Settings for database alerts",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"bdb_backup_delayed": {
								Description:         "Periodic backup has been delayed for longer than specified threshold value [minutes]. -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Periodic backup has been delayed for longer than specified threshold value [minutes]. -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_crdt_src_high_syncer_lag": {
								Description:         "Active-active source - sync lag is higher than specified threshold value [seconds] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Active-active source - sync lag is higher than specified threshold value [seconds] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_crdt_src_syncer_connection_error": {
								Description:         "Active-active source - sync has connection error while trying to connect replica source -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Active-active source - sync has connection error while trying to connect replica source -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_crdt_src_syncer_general_error": {
								Description:         "Active-active source - sync encountered in general error -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Active-active source - sync encountered in general error -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_high_latency": {
								Description:         "Latency is higher than specified threshold value [micro-sec] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Latency is higher than specified threshold value [micro-sec] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_high_throughput": {
								Description:         "Throughput is higher than specified threshold value [requests / sec.] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Throughput is higher than specified threshold value [requests / sec.] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_long_running_action": {
								Description:         "An alert for state-machines that are running for too long -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "An alert for state-machines that are running for too long -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_low_throughput": {
								Description:         "Throughput is lower than specified threshold value [requests / sec.] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Throughput is lower than specified threshold value [requests / sec.] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_ram_dataset_overhead": {
								Description:         "Dataset RAM overhead of a shard has reached the threshold value [% of its RAM limit] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Dataset RAM overhead of a shard has reached the threshold value [% of its RAM limit] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_ram_values": {
								Description:         "Percent of values kept in a shard's RAM is lower than [% of its key count] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Percent of values kept in a shard's RAM is lower than [% of its key count] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_replica_src_high_syncer_lag": {
								Description:         "Replica-of source - sync lag is higher than specified threshold value [seconds] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Replica-of source - sync lag is higher than specified threshold value [seconds] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_replica_src_syncer_connection_error": {
								Description:         "Replica-of source - sync has connection error while trying to connect replica source -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Replica-of source - sync has connection error while trying to connect replica source -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_shard_num_ram_values": {
								Description:         "Number of values kept in a shard's RAM is lower than [values] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Number of values kept in a shard's RAM is lower than [values] -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

							"bdb_size": {
								Description:         "Dataset size has reached the threshold value [% of the memory limit] expected fields: -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",
								MarkdownDescription: "Dataset size has reached the threshold value [% of the memory limit] expected fields: -Note threshold is commented (allow string/int/float and support backwards compatibility) but is required",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"enabled": {
										Description:         "Alert enabled or disabled",
										MarkdownDescription: "Alert enabled or disabled",

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

					"backup": {
						Description:         "Target for automatic database backups.",
						MarkdownDescription: "Target for automatic database backups.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"abs": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"abs_secret_name": {
										Description:         "The name of the K8s secret that holds ABS credentials. The secret must contain the keys 'AccountName' and 'AccountKey', and these must hold the corresponding credentials",
										MarkdownDescription: "The name of the K8s secret that holds ABS credentials. The secret must contain the keys 'AccountName' and 'AccountKey', and these must hold the corresponding credentials",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"container": {
										Description:         "Azure Blob Storage container name.",
										MarkdownDescription: "Azure Blob Storage container name.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"subdir": {
										Description:         "Optional. Azure Blob Storage subdir under container.",
										MarkdownDescription: "Optional. Azure Blob Storage subdir under container.",

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

							"ftp": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"url": {
										Description:         "a URI of the ftps://[USER[:PASSWORD]@]HOST[:PORT]/PATH[/]",
										MarkdownDescription: "a URI of the ftps://[USER[:PASSWORD]@]HOST[:PORT]/PATH[/]",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`ftps?://(([^@]+)@)?([^@/:]+)(:(\d+))?([/\.]/?[^@/\.]+)*?/?$`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"gcs": {
								Description:         "GoogleStorage",
								MarkdownDescription: "GoogleStorage",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"bucket_name": {
										Description:         "Google Storage bucket name.",
										MarkdownDescription: "Google Storage bucket name.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"gcs_secret_name": {
										Description:         "The name of the K8s secret that holds the Google Cloud Storage credentials. The secret must contain the keys 'CLIENT_ID', 'PRIVATE_KEY', 'PRIVATE_KEY_ID', 'CLIENT_EMAIL' and these must hold the corresponding credentials. The keys should correspond to the values in the key JSON.",
										MarkdownDescription: "The name of the K8s secret that holds the Google Cloud Storage credentials. The secret must contain the keys 'CLIENT_ID', 'PRIVATE_KEY', 'PRIVATE_KEY_ID', 'CLIENT_EMAIL' and these must hold the corresponding credentials. The keys should correspond to the values in the key JSON.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"subdir": {
										Description:         "Optional. Google Storage subdir under bucket.",
										MarkdownDescription: "Optional. Google Storage subdir under bucket.",

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

							"interval": {
								Description:         "Backup Interval in seconds",
								MarkdownDescription: "Backup Interval in seconds",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"mount": {
								Description:         "MountPointStorage",
								MarkdownDescription: "MountPointStorage",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"path": {
										Description:         "Path to the local mount point. You must create the mount point on all nodes, and the redislabs:redislabs user must have read and write permissions on the local mount point.",
										MarkdownDescription: "Path to the local mount point. You must create the mount point on all nodes, and the redislabs:redislabs user must have read and write permissions on the local mount point.",

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

							"s3": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"aws_secret_name": {
										Description:         "The name of the K8s secret that holds the AWS credentials. The secret must contain the keys 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY', and these must hold the corresponding credentials.",
										MarkdownDescription: "The name of the K8s secret that holds the AWS credentials. The secret must contain the keys 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY', and these must hold the corresponding credentials.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"bucket_name": {
										Description:         "Amazon S3 bucket name.",
										MarkdownDescription: "Amazon S3 bucket name.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"subdir": {
										Description:         "Optional. Amazon S3 subdir under bucket.",
										MarkdownDescription: "Optional. Amazon S3 subdir under bucket.",

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

							"sftp": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"sftp_secret_name": {
										Description:         "The name of the K8s secret that holds SFTP credentials. The secret must contain the 'Key' key, which is the SSH private key for connecting to the sftp server.",
										MarkdownDescription: "The name of the K8s secret that holds SFTP credentials. The secret must contain the 'Key' key, which is the SSH private key for connecting to the sftp server.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"sftp_url": {
										Description:         "SFTP url",
										MarkdownDescription: "SFTP url",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^sftp://(([^@]+)@)?([^@/:]+)(:(\d+))?(/([^@/\.]+[/\.]?)*)?$`), ""),
										},
									},
								}),

								Required: false,
								Optional: true,
								Computed: false,
							},

							"swift": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"auth_url": {
										Description:         "Swift service authentication URL.",
										MarkdownDescription: "Swift service authentication URL.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^https?://(([^@]+)@)?([^@/:]+)(:(\d+))?([/\.]([^@/\.]+))*?/?$`), ""),
										},
									},

									"container": {
										Description:         "Swift object store container for storing the backup files.",
										MarkdownDescription: "Swift object store container for storing the backup files.",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"prefix": {
										Description:         "Optional. Prefix (path) of backup files in the swift container.",
										MarkdownDescription: "Optional. Prefix (path) of backup files in the swift container.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"swift_secret_name": {
										Description:         "The name of the K8s secret that holds Swift credentials. The secret must contain the keys 'Key' and 'User', and these must hold the corresponding credentials: service access key and service user name (pattern for the latter does not allow special characters &,<,>,')",
										MarkdownDescription: "The name of the K8s secret that holds Swift credentials. The secret must contain the keys 'Key' and 'User', and these must hold the corresponding credentials: service access key and service user name (pattern for the latter does not allow special characters &,<,>,')",

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

					"client_authentication_certificates": {
						Description:         "The Secrets containing TLS Client Certificate to use for Authentication",
						MarkdownDescription: "The Secrets containing TLS Client Certificate to use for Authentication",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"data_internode_encryption": {
						Description:         "Internode encryption (INE) setting. An optional boolean setting, overriding a similar cluster-wide policy. If set to False, INE is guaranteed to be turned off for this DB (regardless of cluster-wide policy). If set to True, INE will be turned on, unless the capability is not supported by the DB ( in such a case we will get an error and database creation will fail). If left unspecified, will be disabled if internode encryption is not supported by the DB (regardless of cluster default). Deleting this property after explicitly setting its value shall have no effect.",
						MarkdownDescription: "Internode encryption (INE) setting. An optional boolean setting, overriding a similar cluster-wide policy. If set to False, INE is guaranteed to be turned off for this DB (regardless of cluster-wide policy). If set to True, INE will be turned on, unless the capability is not supported by the DB ( in such a case we will get an error and database creation will fail). If left unspecified, will be disabled if internode encryption is not supported by the DB (regardless of cluster default). Deleting this property after explicitly setting its value shall have no effect.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"database_port": {
						Description:         "Database port number. TCP port on which the database is available. Will be generated automatically if omitted. can not be changed after creation",
						MarkdownDescription: "Database port number. TCP port on which the database is available. Will be generated automatically if omitted. can not be changed after creation",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"database_secret_name": {
						Description:         "The name of the K8s secret that holds the password to the database.",
						MarkdownDescription: "The name of the K8s secret that holds the password to the database.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"default_user": {
						Description:         "Is connecting with a default user allowed?",
						MarkdownDescription: "Is connecting with a default user allowed?",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"eviction_policy": {
						Description:         "Database eviction policy. see more https://docs.redislabs.com/latest/rs/administering/database-operations/eviction-policy/",
						MarkdownDescription: "Database eviction policy. see more https://docs.redislabs.com/latest/rs/administering/database-operations/eviction-policy/",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"global_configurations": {
						Description:         "Flag that determines if this is the default global configurations for active active database. In case this resource is created manually at the active active database creation this field must be filled via the user otherwise the operator will create the global configurations REDB automatically with default values. Note: this feature is currently unsupported.",
						MarkdownDescription: "Flag that determines if this is the default global configurations for active active database. In case this resource is created manually at the active active database creation this field must be filled via the user otherwise the operator will create the global configurations REDB automatically with default values. Note: this feature is currently unsupported.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"is_rof": {
						Description:         "Whether it is an RoF database or not. Applicable only for databases of type 'REDIS'. Assumed to be false if left blank.",
						MarkdownDescription: "Whether it is an RoF database or not. Applicable only for databases of type 'REDIS'. Assumed to be false if left blank.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"memcached_sasl_secret_name": {
						Description:         "Credentials used for binary authentication in memcached databases. The credentials should be saved as an opaque secret and the name of that secret should be configured using this field. For username, use 'username' as the key and the actual username as the value. For password, use 'password' as the key and the actual password as the value. Note that connections are not encrypted.",
						MarkdownDescription: "Credentials used for binary authentication in memcached databases. The credentials should be saved as an opaque secret and the name of that secret should be configured using this field. For username, use 'username' as the key and the actual username as the value. For password, use 'password' as the key and the actual password as the value. Note that connections are not encrypted.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"memory_size": {
						Description:         "memory size of database. use formats like 100MB, 0.1GB. minimum value in 100MB. When redis on flash (RoF) is enabled, this value refers to RAM+Flash memory, and it must not be below 1GB.",
						MarkdownDescription: "memory size of database. use formats like 100MB, 0.1GB. minimum value in 100MB. When redis on flash (RoF) is enabled, this value refers to RAM+Flash memory, and it must not be below 1GB.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"modules_list": {
						Description:         "List of modules associated with database",
						MarkdownDescription: "List of modules associated with database",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"config": {
								Description:         "Module command line arguments e.g. VKEY_MAX_ENTITY_COUNT 30",
								MarkdownDescription: "Module command line arguments e.g. VKEY_MAX_ENTITY_COUNT 30",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "The module's name e.g 'ft' for redissearch",
								MarkdownDescription: "The module's name e.g 'ft' for redissearch",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"version": {
								Description:         "Module's semantic version e.g '1.6.12'",
								MarkdownDescription: "Module's semantic version e.g '1.6.12'",

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

					"oss_cluster": {
						Description:         "OSS Cluster mode option. Note that not all client libraries support OSS cluster mode.",
						MarkdownDescription: "OSS Cluster mode option. Note that not all client libraries support OSS cluster mode.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"persistence": {
						Description:         "Database on-disk persistence policy",
						MarkdownDescription: "Database on-disk persistence policy",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("disabled", "aofEverySecond", "aofAlways", "snapshotEvery1Hour", "snapshotEvery6Hour", "snapshotEvery12Hour"),
						},
					},

					"proxy_policy": {
						Description:         "The policy used for proxy binding to the endpoint. Supported proxy policies are: single/all-master-shards/all-nodes When left blank, the default value will be chosen according to the value of ossCluster - single if disabled, all-master-shards when enabled",
						MarkdownDescription: "The policy used for proxy binding to the endpoint. Supported proxy policies are: single/all-master-shards/all-nodes When left blank, the default value will be chosen according to the value of ossCluster - single if disabled, all-master-shards when enabled",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rack_aware": {
						Description:         "Whether database should be rack aware. This improves availability - more information: https://docs.redislabs.com/latest/rs/concepts/high-availability/rack-zone-awareness/",
						MarkdownDescription: "Whether database should be rack aware. This improves availability - more information: https://docs.redislabs.com/latest/rs/concepts/high-availability/rack-zone-awareness/",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"redis_enterprise_cluster": {
						Description:         "Connection to Redis Enterprise Cluster",
						MarkdownDescription: "Connection to Redis Enterprise Cluster",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "The name of the Redis Enterprise Cluster where the database should be stored.",
								MarkdownDescription: "The name of the Redis Enterprise Cluster where the database should be stored.",

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

					"redis_version": {
						Description:         "Redis OSS version. For existing databases - Upgrade Redis OSS version. For new databases - the version which the database will be created with. If set to 'major' - will always upgrade to the most recent major Redis version. If set to 'latest' - will always upgrade to the most recent Redis version. Depends on 'redisUpgradePolicy' - if you want to set the value to 'latest' for some databases, you must set redisUpgradePolicy on the cluster before. Possible values are 'major' or 'latest' When using upgrade - make sure to backup the database before. This value is used only for database type 'redis'",
						MarkdownDescription: "Redis OSS version. For existing databases - Upgrade Redis OSS version. For new databases - the version which the database will be created with. If set to 'major' - will always upgrade to the most recent major Redis version. If set to 'latest' - will always upgrade to the most recent Redis version. Depends on 'redisUpgradePolicy' - if you want to set the value to 'latest' for some databases, you must set redisUpgradePolicy on the cluster before. Possible values are 'major' or 'latest' When using upgrade - make sure to backup the database before. This value is used only for database type 'redis'",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("major", "latest"),
						},
					},

					"replica_sources": {
						Description:         "What databases to replicate from",
						MarkdownDescription: "What databases to replicate from",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"client_key_secret": {
								Description:         "Secret that defines the client certificate and key used by the syncer in the target database cluster. The secret must have 2 keys in its map: 'cert' which is the PEM encoded certificate, and 'key' which is the PEM encoded private key.",
								MarkdownDescription: "Secret that defines the client certificate and key used by the syncer in the target database cluster. The secret must have 2 keys in its map: 'cert' which is the PEM encoded certificate, and 'key' which is the PEM encoded private key.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"compression": {
								Description:         "GZIP compression level (0-6) to use for replication.",
								MarkdownDescription: "GZIP compression level (0-6) to use for replication.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"replica_source_name": {
								Description:         "The name of the resource from which the source database URI is derived. The type of resource must match the type specified in the ReplicaSourceType field.",
								MarkdownDescription: "The name of the resource from which the source database URI is derived. The type of resource must match the type specified in the ReplicaSourceType field.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"replica_source_type": {
								Description:         "The type of resource from which the source database URI is derived. If set to 'SECRET', the source database URI is derived from the secret named in the ReplicaSourceName field. The secret must have a key named 'uri' that defines the URI of the source database in the form of 'redis://...'. The type of secret (kubernetes, vault, ...) is determined by the secret mechanism used by the underlying REC object. If set to 'REDB', the source database URI is derived from the RedisEnterpriseDatabase resource named in the ReplicaSourceName field.",
								MarkdownDescription: "The type of resource from which the source database URI is derived. If set to 'SECRET', the source database URI is derived from the secret named in the ReplicaSourceName field. The secret must have a key named 'uri' that defines the URI of the source database in the form of 'redis://...'. The type of secret (kubernetes, vault, ...) is determined by the secret mechanism used by the underlying REC object. If set to 'REDB', the source database URI is derived from the RedisEnterpriseDatabase resource named in the ReplicaSourceName field.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"server_cert_secret": {
								Description:         "Secret that defines the server certificate used by the proxy in the source database cluster. The secret must have 1 key in its map: 'cert' which is the PEM encoded certificate.",
								MarkdownDescription: "Secret that defines the server certificate used by the proxy in the source database cluster. The secret must have 1 key in its map: 'cert' which is the PEM encoded certificate.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"tls_sni_name": {
								Description:         "TLS SNI name to use for the replication link.",
								MarkdownDescription: "TLS SNI name to use for the replication link.",

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

					"replication": {
						Description:         "In-memory database replication. When enabled, database will have replica shard for every master - leading to higher availability.",
						MarkdownDescription: "In-memory database replication. When enabled, database will have replica shard for every master - leading to higher availability.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"rof_ram_size": {
						Description:         "The size of the RAM portion of an RoF database. Similarly to 'memorySize' use formats like 100MB, 0.1GB. It must be at least 10% of combined memory size (RAM and Flash), as specified by 'memorySize'.",
						MarkdownDescription: "The size of the RAM portion of an RoF database. Similarly to 'memorySize' use formats like 100MB, 0.1GB. It must be at least 10% of combined memory size (RAM and Flash), as specified by 'memorySize'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"roles_permissions": {
						Description:         "List of Redis Enteprise ACL and Role bindings to apply",
						MarkdownDescription: "List of Redis Enteprise ACL and Role bindings to apply",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"acl": {
								Description:         "Acl Name of RolePermissionType",
								MarkdownDescription: "Acl Name of RolePermissionType",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"role": {
								Description:         "Role Name of RolePermissionType",
								MarkdownDescription: "Role Name of RolePermissionType",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"type": {
								Description:         "Type of Redis Enterprise Database Role Permission",
								MarkdownDescription: "Type of Redis Enterprise Database Role Permission",

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

					"shard_count": {
						Description:         "Number of database server-side shards",
						MarkdownDescription: "Number of database server-side shards",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"shards_placement": {
						Description:         "Control the density of shards - should they reside on as few or as many nodes as possible. Available options are 'dense' or 'sparse'. If left unset, defaults to 'dense'.",
						MarkdownDescription: "Control the density of shards - should they reside on as few or as many nodes as possible. Available options are 'dense' or 'sparse'. If left unset, defaults to 'dense'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("dense", "sparse"),
						},
					},

					"tls_mode": {
						Description:         "Require SSL authenticated and encrypted connections to the database. enabled - all incoming connections to the Database must use SSL. disabled - no incoming connection to the Database should use SSL. replica_ssl - databases that replicate from this one need to use SSL.",
						MarkdownDescription: "Require SSL authenticated and encrypted connections to the database. enabled - all incoming connections to the Database must use SSL. disabled - no incoming connection to the Database should use SSL. replica_ssl - databases that replicate from this one need to use SSL.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("disabled", "enabled", "replica_ssl"),
						},
					},

					"type": {
						Description:         "The type of the database (redis or memcached). Defaults to 'redis'.",
						MarkdownDescription: "The type of the database (redis or memcached). Defaults to 'redis'.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.OneOf("redis", "memcached"),
						},
					},
				}),

				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}, nil
}

func (r *AppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_app_redislabs_com_redis_enterprise_database_v1alpha1")

	var state AppRedislabsComRedisEnterpriseDatabaseV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppRedislabsComRedisEnterpriseDatabaseV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("app.redislabs.com/v1alpha1")
	goModel.Kind = utilities.Ptr("RedisEnterpriseDatabase")

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

func (r *AppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_app_redislabs_com_redis_enterprise_database_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *AppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_app_redislabs_com_redis_enterprise_database_v1alpha1")

	var state AppRedislabsComRedisEnterpriseDatabaseV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel AppRedislabsComRedisEnterpriseDatabaseV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("app.redislabs.com/v1alpha1")
	goModel.Kind = utilities.Ptr("RedisEnterpriseDatabase")

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

func (r *AppRedislabsComRedisEnterpriseDatabaseV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_app_redislabs_com_redis_enterprise_database_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
