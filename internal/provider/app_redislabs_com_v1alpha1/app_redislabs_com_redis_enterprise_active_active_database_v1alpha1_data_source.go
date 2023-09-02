/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package app_redislabs_com_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
)

var (
	_ datasource.DataSource              = &AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSource{}
)

func NewAppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSource() datasource.DataSource {
	return &AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSource{}
}

type AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSourceData struct {
	ID types.String `tfsdk:"id" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		GlobalConfigurations *struct {
			ActiveActive *struct {
				Name                     *string `tfsdk:"name" json:"name,omitempty"`
				ParticipatingClusterName *string `tfsdk:"participating_cluster_name" json:"participatingClusterName,omitempty"`
			} `tfsdk:"active_active" json:"activeActive,omitempty"`
			AlertSettings *struct {
				Bdb_backup_delayed *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_backup_delayed" json:"bdb_backup_delayed,omitempty"`
				Bdb_crdt_src_high_syncer_lag *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_crdt_src_high_syncer_lag" json:"bdb_crdt_src_high_syncer_lag,omitempty"`
				Bdb_crdt_src_syncer_connection_error *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_crdt_src_syncer_connection_error" json:"bdb_crdt_src_syncer_connection_error,omitempty"`
				Bdb_crdt_src_syncer_general_error *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_crdt_src_syncer_general_error" json:"bdb_crdt_src_syncer_general_error,omitempty"`
				Bdb_high_latency *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_high_latency" json:"bdb_high_latency,omitempty"`
				Bdb_high_throughput *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_high_throughput" json:"bdb_high_throughput,omitempty"`
				Bdb_long_running_action *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_long_running_action" json:"bdb_long_running_action,omitempty"`
				Bdb_low_throughput *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_low_throughput" json:"bdb_low_throughput,omitempty"`
				Bdb_ram_dataset_overhead *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_ram_dataset_overhead" json:"bdb_ram_dataset_overhead,omitempty"`
				Bdb_ram_values *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_ram_values" json:"bdb_ram_values,omitempty"`
				Bdb_replica_src_high_syncer_lag *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_replica_src_high_syncer_lag" json:"bdb_replica_src_high_syncer_lag,omitempty"`
				Bdb_replica_src_syncer_connection_error *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_replica_src_syncer_connection_error" json:"bdb_replica_src_syncer_connection_error,omitempty"`
				Bdb_shard_num_ram_values *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_shard_num_ram_values" json:"bdb_shard_num_ram_values,omitempty"`
				Bdb_size *struct {
					Enabled   *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
					Threshold *string `tfsdk:"threshold" json:"threshold,omitempty"`
				} `tfsdk:"bdb_size" json:"bdb_size,omitempty"`
			} `tfsdk:"alert_settings" json:"alertSettings,omitempty"`
			Backup *struct {
				Abs *struct {
					AbsSecretName *string `tfsdk:"abs_secret_name" json:"absSecretName,omitempty"`
					Container     *string `tfsdk:"container" json:"container,omitempty"`
					Subdir        *string `tfsdk:"subdir" json:"subdir,omitempty"`
				} `tfsdk:"abs" json:"abs,omitempty"`
				Ftp *struct {
					Url *string `tfsdk:"url" json:"url,omitempty"`
				} `tfsdk:"ftp" json:"ftp,omitempty"`
				Gcs *struct {
					BucketName    *string `tfsdk:"bucket_name" json:"bucketName,omitempty"`
					GcsSecretName *string `tfsdk:"gcs_secret_name" json:"gcsSecretName,omitempty"`
					Subdir        *string `tfsdk:"subdir" json:"subdir,omitempty"`
				} `tfsdk:"gcs" json:"gcs,omitempty"`
				Interval *int64 `tfsdk:"interval" json:"interval,omitempty"`
				Mount    *struct {
					Path *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"mount" json:"mount,omitempty"`
				S3 *struct {
					AwsSecretName *string `tfsdk:"aws_secret_name" json:"awsSecretName,omitempty"`
					BucketName    *string `tfsdk:"bucket_name" json:"bucketName,omitempty"`
					Subdir        *string `tfsdk:"subdir" json:"subdir,omitempty"`
				} `tfsdk:"s3" json:"s3,omitempty"`
				Sftp *struct {
					SftpSecretName *string `tfsdk:"sftp_secret_name" json:"sftpSecretName,omitempty"`
					Sftp_url       *string `tfsdk:"sftp_url" json:"sftp_url,omitempty"`
				} `tfsdk:"sftp" json:"sftp,omitempty"`
				Swift *struct {
					Auth_url        *string `tfsdk:"auth_url" json:"auth_url,omitempty"`
					Container       *string `tfsdk:"container" json:"container,omitempty"`
					Prefix          *string `tfsdk:"prefix" json:"prefix,omitempty"`
					SwiftSecretName *string `tfsdk:"swift_secret_name" json:"swiftSecretName,omitempty"`
				} `tfsdk:"swift" json:"swift,omitempty"`
			} `tfsdk:"backup" json:"backup,omitempty"`
			ClientAuthenticationCertificates *[]string `tfsdk:"client_authentication_certificates" json:"clientAuthenticationCertificates,omitempty"`
			DataInternodeEncryption          *bool     `tfsdk:"data_internode_encryption" json:"dataInternodeEncryption,omitempty"`
			DatabasePort                     *int64    `tfsdk:"database_port" json:"databasePort,omitempty"`
			DatabaseSecretName               *string   `tfsdk:"database_secret_name" json:"databaseSecretName,omitempty"`
			DefaultUser                      *bool     `tfsdk:"default_user" json:"defaultUser,omitempty"`
			EvictionPolicy                   *string   `tfsdk:"eviction_policy" json:"evictionPolicy,omitempty"`
			IsRof                            *bool     `tfsdk:"is_rof" json:"isRof,omitempty"`
			MemcachedSaslSecretName          *string   `tfsdk:"memcached_sasl_secret_name" json:"memcachedSaslSecretName,omitempty"`
			MemorySize                       *string   `tfsdk:"memory_size" json:"memorySize,omitempty"`
			ModulesList                      *[]struct {
				Config  *string `tfsdk:"config" json:"config,omitempty"`
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Uid     *string `tfsdk:"uid" json:"uid,omitempty"`
				Version *string `tfsdk:"version" json:"version,omitempty"`
			} `tfsdk:"modules_list" json:"modulesList,omitempty"`
			OssCluster             *bool   `tfsdk:"oss_cluster" json:"ossCluster,omitempty"`
			Persistence            *string `tfsdk:"persistence" json:"persistence,omitempty"`
			ProxyPolicy            *string `tfsdk:"proxy_policy" json:"proxyPolicy,omitempty"`
			RackAware              *bool   `tfsdk:"rack_aware" json:"rackAware,omitempty"`
			RedisEnterpriseCluster *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"redis_enterprise_cluster" json:"redisEnterpriseCluster,omitempty"`
			RedisVersion   *string `tfsdk:"redis_version" json:"redisVersion,omitempty"`
			ReplicaSources *[]struct {
				ClientKeySecret   *string `tfsdk:"client_key_secret" json:"clientKeySecret,omitempty"`
				Compression       *int64  `tfsdk:"compression" json:"compression,omitempty"`
				ReplicaSourceName *string `tfsdk:"replica_source_name" json:"replicaSourceName,omitempty"`
				ReplicaSourceType *string `tfsdk:"replica_source_type" json:"replicaSourceType,omitempty"`
				ServerCertSecret  *string `tfsdk:"server_cert_secret" json:"serverCertSecret,omitempty"`
				TlsSniName        *string `tfsdk:"tls_sni_name" json:"tlsSniName,omitempty"`
			} `tfsdk:"replica_sources" json:"replicaSources,omitempty"`
			Replication      *bool   `tfsdk:"replication" json:"replication,omitempty"`
			RofRamSize       *string `tfsdk:"rof_ram_size" json:"rofRamSize,omitempty"`
			RolesPermissions *[]struct {
				Acl  *string `tfsdk:"acl" json:"acl,omitempty"`
				Role *string `tfsdk:"role" json:"role,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"roles_permissions" json:"rolesPermissions,omitempty"`
			ShardCount      *int64  `tfsdk:"shard_count" json:"shardCount,omitempty"`
			ShardsPlacement *string `tfsdk:"shards_placement" json:"shardsPlacement,omitempty"`
			TlsMode         *string `tfsdk:"tls_mode" json:"tlsMode,omitempty"`
			Type            *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"global_configurations" json:"globalConfigurations,omitempty"`
		ParticipatingClusters *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"participating_clusters" json:"participatingClusters,omitempty"`
		RedisEnterpriseCluster *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"redis_enterprise_cluster" json:"redisEnterpriseCluster,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_app_redislabs_com_redis_enterprise_active_active_database_v1alpha1"
}

func (r *AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "RedisEnterpriseActiveActiveDatabase is the Schema for the redisenterpriseactiveactivedatabase API",
		MarkdownDescription: "RedisEnterpriseActiveActiveDatabase is the Schema for the redisenterpriseactiveactivedatabase API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
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
						Optional:            false,
						Computed:            true,
					},
					"annotations": schema.MapAttribute{
						Description:         "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						MarkdownDescription: "Keys and values that can be used by external tooling to store and retrieve arbitrary metadata about this object. See https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},
				},
			},

			"spec": schema.SingleNestedAttribute{
				Description:         "RedisEnterpriseActiveActiveDatabaseSpec defines the desired state of RedisEnterpriseActiveActiveDatabase",
				MarkdownDescription: "RedisEnterpriseActiveActiveDatabaseSpec defines the desired state of RedisEnterpriseActiveActiveDatabase",
				Attributes: map[string]schema.Attribute{
					"global_configurations": schema.SingleNestedAttribute{
						Description:         "The Active-Active database global configurations, contains the global properties for each of the participating clusters/ instances databases within the Active-Active database.",
						MarkdownDescription: "The Active-Active database global configurations, contains the global properties for each of the participating clusters/ instances databases within the Active-Active database.",
						Attributes: map[string]schema.Attribute{
							"active_active": schema.SingleNestedAttribute{
								Description:         "Connection/ association to the Active-Active database.",
								MarkdownDescription: "Connection/ association to the Active-Active database.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "The the corresponding Active-Active database name, Redis Enterprise Active Active Database custom resource name, this Resource is associated with. In case this resource is created manually at the active active database creation this field must be filled via the user, otherwise, the operator will assign this field automatically. Note: this feature is currently unsupported.",
										MarkdownDescription: "The the corresponding Active-Active database name, Redis Enterprise Active Active Database custom resource name, this Resource is associated with. In case this resource is created manually at the active active database creation this field must be filled via the user, otherwise, the operator will assign this field automatically. Note: this feature is currently unsupported.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"participating_cluster_name": schema.StringAttribute{
										Description:         "The corresponding participating cluster name, Redis Enterprise Remote Cluster custom resource name, in the Active-Active database, In case this resource is created manually at the active active database creation this field must be filled via the user, otherwise, the operator will assign this field automatically. Note: this feature is currently unsupported.",
										MarkdownDescription: "The corresponding participating cluster name, Redis Enterprise Remote Cluster custom resource name, in the Active-Active database, In case this resource is created manually at the active active database creation this field must be filled via the user, otherwise, the operator will assign this field automatically. Note: this feature is currently unsupported.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"alert_settings": schema.SingleNestedAttribute{
								Description:         "Settings for database alerts",
								MarkdownDescription: "Settings for database alerts",
								Attributes: map[string]schema.Attribute{
									"bdb_backup_delayed": schema.SingleNestedAttribute{
										Description:         "Periodic backup has been delayed for longer than specified threshold value [minutes]",
										MarkdownDescription: "Periodic backup has been delayed for longer than specified threshold value [minutes]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_crdt_src_high_syncer_lag": schema.SingleNestedAttribute{
										Description:         "Active-active source - sync lag is higher than specified threshold value [seconds]",
										MarkdownDescription: "Active-active source - sync lag is higher than specified threshold value [seconds]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_crdt_src_syncer_connection_error": schema.SingleNestedAttribute{
										Description:         "Active-active source - sync has connection error while trying to connect replica source",
										MarkdownDescription: "Active-active source - sync has connection error while trying to connect replica source",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_crdt_src_syncer_general_error": schema.SingleNestedAttribute{
										Description:         "Active-active source - sync encountered in general error",
										MarkdownDescription: "Active-active source - sync encountered in general error",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_high_latency": schema.SingleNestedAttribute{
										Description:         "Latency is higher than specified threshold value [micro-sec]",
										MarkdownDescription: "Latency is higher than specified threshold value [micro-sec]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_high_throughput": schema.SingleNestedAttribute{
										Description:         "Throughput is higher than specified threshold value [requests / sec.]",
										MarkdownDescription: "Throughput is higher than specified threshold value [requests / sec.]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_long_running_action": schema.SingleNestedAttribute{
										Description:         "An alert for state-machines that are running for too long",
										MarkdownDescription: "An alert for state-machines that are running for too long",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_low_throughput": schema.SingleNestedAttribute{
										Description:         "Throughput is lower than specified threshold value [requests / sec.]",
										MarkdownDescription: "Throughput is lower than specified threshold value [requests / sec.]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_ram_dataset_overhead": schema.SingleNestedAttribute{
										Description:         "Dataset RAM overhead of a shard has reached the threshold value [% of its RAM limit]",
										MarkdownDescription: "Dataset RAM overhead of a shard has reached the threshold value [% of its RAM limit]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_ram_values": schema.SingleNestedAttribute{
										Description:         "Percent of values kept in a shard's RAM is lower than [% of its key count]",
										MarkdownDescription: "Percent of values kept in a shard's RAM is lower than [% of its key count]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_replica_src_high_syncer_lag": schema.SingleNestedAttribute{
										Description:         "Replica-of source - sync lag is higher than specified threshold value [seconds]",
										MarkdownDescription: "Replica-of source - sync lag is higher than specified threshold value [seconds]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_replica_src_syncer_connection_error": schema.SingleNestedAttribute{
										Description:         "Replica-of source - sync has connection error while trying to connect replica source",
										MarkdownDescription: "Replica-of source - sync has connection error while trying to connect replica source",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_shard_num_ram_values": schema.SingleNestedAttribute{
										Description:         "Number of values kept in a shard's RAM is lower than [values]",
										MarkdownDescription: "Number of values kept in a shard's RAM is lower than [values]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"bdb_size": schema.SingleNestedAttribute{
										Description:         "Dataset size has reached the threshold value [% of the memory limit]",
										MarkdownDescription: "Dataset size has reached the threshold value [% of the memory limit]",
										Attributes: map[string]schema.Attribute{
											"enabled": schema.BoolAttribute{
												Description:         "Alert enabled or disabled",
												MarkdownDescription: "Alert enabled or disabled",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"threshold": schema.StringAttribute{
												Description:         "Threshold for alert going on/off",
												MarkdownDescription: "Threshold for alert going on/off",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"backup": schema.SingleNestedAttribute{
								Description:         "Target for automatic database backups.",
								MarkdownDescription: "Target for automatic database backups.",
								Attributes: map[string]schema.Attribute{
									"abs": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"abs_secret_name": schema.StringAttribute{
												Description:         "The name of the secret that holds ABS credentials. The secret must contain the keys 'AccountName' and 'AccountKey', and these must hold the corresponding credentials",
												MarkdownDescription: "The name of the secret that holds ABS credentials. The secret must contain the keys 'AccountName' and 'AccountKey', and these must hold the corresponding credentials",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"container": schema.StringAttribute{
												Description:         "Azure Blob Storage container name.",
												MarkdownDescription: "Azure Blob Storage container name.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"subdir": schema.StringAttribute{
												Description:         "Optional. Azure Blob Storage subdir under container.",
												MarkdownDescription: "Optional. Azure Blob Storage subdir under container.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"ftp": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"url": schema.StringAttribute{
												Description:         "a URI of the 'ftps://[USER[:PASSWORD]@]HOST[:PORT]/PATH[/]' format",
												MarkdownDescription: "a URI of the 'ftps://[USER[:PASSWORD]@]HOST[:PORT]/PATH[/]' format",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"gcs": schema.SingleNestedAttribute{
										Description:         "GoogleStorage",
										MarkdownDescription: "GoogleStorage",
										Attributes: map[string]schema.Attribute{
											"bucket_name": schema.StringAttribute{
												Description:         "Google Storage bucket name.",
												MarkdownDescription: "Google Storage bucket name.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"gcs_secret_name": schema.StringAttribute{
												Description:         "The name of the secret that holds the Google Cloud Storage credentials. The secret must contain the keys 'CLIENT_ID', 'PRIVATE_KEY', 'PRIVATE_KEY_ID', 'CLIENT_EMAIL' and these must hold the corresponding credentials. The keys should correspond to the values in the key JSON.",
												MarkdownDescription: "The name of the secret that holds the Google Cloud Storage credentials. The secret must contain the keys 'CLIENT_ID', 'PRIVATE_KEY', 'PRIVATE_KEY_ID', 'CLIENT_EMAIL' and these must hold the corresponding credentials. The keys should correspond to the values in the key JSON.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"subdir": schema.StringAttribute{
												Description:         "Optional. Google Storage subdir under bucket.",
												MarkdownDescription: "Optional. Google Storage subdir under bucket.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"interval": schema.Int64Attribute{
										Description:         "Backup Interval in seconds",
										MarkdownDescription: "Backup Interval in seconds",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"mount": schema.SingleNestedAttribute{
										Description:         "MountPointStorage",
										MarkdownDescription: "MountPointStorage",
										Attributes: map[string]schema.Attribute{
											"path": schema.StringAttribute{
												Description:         "Path to the local mount point. You must create the mount point on all nodes, and the redislabs:redislabs user must have read and write permissions on the local mount point.",
												MarkdownDescription: "Path to the local mount point. You must create the mount point on all nodes, and the redislabs:redislabs user must have read and write permissions on the local mount point.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"s3": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"aws_secret_name": schema.StringAttribute{
												Description:         "The name of the secret that holds the AWS credentials. The secret must contain the keys 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY', and these must hold the corresponding credentials.",
												MarkdownDescription: "The name of the secret that holds the AWS credentials. The secret must contain the keys 'AWS_ACCESS_KEY_ID' and 'AWS_SECRET_ACCESS_KEY', and these must hold the corresponding credentials.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"bucket_name": schema.StringAttribute{
												Description:         "Amazon S3 bucket name.",
												MarkdownDescription: "Amazon S3 bucket name.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"subdir": schema.StringAttribute{
												Description:         "Optional. Amazon S3 subdir under bucket.",
												MarkdownDescription: "Optional. Amazon S3 subdir under bucket.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"sftp": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"sftp_secret_name": schema.StringAttribute{
												Description:         "The name of the secret that holds SFTP credentials. The secret must contain the 'Key' key, which is the SSH private key for connecting to the sftp server.",
												MarkdownDescription: "The name of the secret that holds SFTP credentials. The secret must contain the 'Key' key, which is the SSH private key for connecting to the sftp server.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"sftp_url": schema.StringAttribute{
												Description:         "SFTP url",
												MarkdownDescription: "SFTP url",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},

									"swift": schema.SingleNestedAttribute{
										Description:         "",
										MarkdownDescription: "",
										Attributes: map[string]schema.Attribute{
											"auth_url": schema.StringAttribute{
												Description:         "Swift service authentication URL.",
												MarkdownDescription: "Swift service authentication URL.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"container": schema.StringAttribute{
												Description:         "Swift object store container for storing the backup files.",
												MarkdownDescription: "Swift object store container for storing the backup files.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"prefix": schema.StringAttribute{
												Description:         "Optional. Prefix (path) of backup files in the swift container.",
												MarkdownDescription: "Optional. Prefix (path) of backup files in the swift container.",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},

											"swift_secret_name": schema.StringAttribute{
												Description:         "The name of the secret that holds Swift credentials. The secret must contain the keys 'Key' and 'User', and these must hold the corresponding credentials: service access key and service user name (pattern for the latter does not allow special characters &,<,>,')",
												MarkdownDescription: "The name of the secret that holds Swift credentials. The secret must contain the keys 'Key' and 'User', and these must hold the corresponding credentials: service access key and service user name (pattern for the latter does not allow special characters &,<,>,')",
												Required:            false,
												Optional:            false,
												Computed:            true,
											},
										},
										Required: false,
										Optional: false,
										Computed: true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"client_authentication_certificates": schema.ListAttribute{
								Description:         "The Secrets containing TLS Client Certificate to use for Authentication",
								MarkdownDescription: "The Secrets containing TLS Client Certificate to use for Authentication",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"data_internode_encryption": schema.BoolAttribute{
								Description:         "Internode encryption (INE) setting. An optional boolean setting, overriding a similar cluster-wide policy. If set to False, INE is guaranteed to be turned off for this DB (regardless of cluster-wide policy). If set to True, INE will be turned on, unless the capability is not supported by the DB ( in such a case we will get an error and database creation will fail). If left unspecified, will be disabled if internode encryption is not supported by the DB (regardless of cluster default). Deleting this property after explicitly setting its value shall have no effect.",
								MarkdownDescription: "Internode encryption (INE) setting. An optional boolean setting, overriding a similar cluster-wide policy. If set to False, INE is guaranteed to be turned off for this DB (regardless of cluster-wide policy). If set to True, INE will be turned on, unless the capability is not supported by the DB ( in such a case we will get an error and database creation will fail). If left unspecified, will be disabled if internode encryption is not supported by the DB (regardless of cluster default). Deleting this property after explicitly setting its value shall have no effect.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"database_port": schema.Int64Attribute{
								Description:         "Database port number. TCP port on which the database is available. Will be generated automatically if omitted. can not be changed after creation",
								MarkdownDescription: "Database port number. TCP port on which the database is available. Will be generated automatically if omitted. can not be changed after creation",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"database_secret_name": schema.StringAttribute{
								Description:         "The name of the secret that holds the password to the database (redis databases only). If secret does not exist, it will be created. To define the password, create an opaque secret and set the name in the spec. The password will be taken from the value of the 'password' key. Use an empty string as value within the secret to disable authentication for the database. Notes - For Active-Active databases this secret will not be automatically created, and also, memcached databases must not be set with a value, and a secret/password will not be automatically created for them. Use the memcachedSaslSecretName field to set authentication parameters for memcached databases.",
								MarkdownDescription: "The name of the secret that holds the password to the database (redis databases only). If secret does not exist, it will be created. To define the password, create an opaque secret and set the name in the spec. The password will be taken from the value of the 'password' key. Use an empty string as value within the secret to disable authentication for the database. Notes - For Active-Active databases this secret will not be automatically created, and also, memcached databases must not be set with a value, and a secret/password will not be automatically created for them. Use the memcachedSaslSecretName field to set authentication parameters for memcached databases.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"default_user": schema.BoolAttribute{
								Description:         "Is connecting with a default user allowed?  If disabled, the DatabaseSecret will not be created or updated",
								MarkdownDescription: "Is connecting with a default user allowed?  If disabled, the DatabaseSecret will not be created or updated",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"eviction_policy": schema.StringAttribute{
								Description:         "Database eviction policy. see more https://docs.redislabs.com/latest/rs/administering/database-operations/eviction-policy/",
								MarkdownDescription: "Database eviction policy. see more https://docs.redislabs.com/latest/rs/administering/database-operations/eviction-policy/",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"is_rof": schema.BoolAttribute{
								Description:         "Whether it is an RoF database or not. Applicable only for databases of type 'REDIS'. Assumed to be false if left blank.",
								MarkdownDescription: "Whether it is an RoF database or not. Applicable only for databases of type 'REDIS'. Assumed to be false if left blank.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"memcached_sasl_secret_name": schema.StringAttribute{
								Description:         "Credentials used for binary authentication in memcached databases. The credentials should be saved as an opaque secret and the name of that secret should be configured using this field. For username, use 'username' as the key and the actual username as the value. For password, use 'password' as the key and the actual password as the value. Note that connections are not encrypted.",
								MarkdownDescription: "Credentials used for binary authentication in memcached databases. The credentials should be saved as an opaque secret and the name of that secret should be configured using this field. For username, use 'username' as the key and the actual username as the value. For password, use 'password' as the key and the actual password as the value. Note that connections are not encrypted.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"memory_size": schema.StringAttribute{
								Description:         "memory size of database. use formats like 100MB, 0.1GB. minimum value in 100MB. When redis on flash (RoF) is enabled, this value refers to RAM+Flash memory, and it must not be below 1GB.",
								MarkdownDescription: "memory size of database. use formats like 100MB, 0.1GB. minimum value in 100MB. When redis on flash (RoF) is enabled, this value refers to RAM+Flash memory, and it must not be below 1GB.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"modules_list": schema.ListNestedAttribute{
								Description:         "List of modules associated with database. Note - For Active-Active databases this feature is currently in preview. For this feature to take effect for Active-Active databases, set a boolean environment variable with the name 'ENABLE_ALPHA_FEATURES' to True. This variable can be set via the redis-enterprise-operator pod spec, or through the operator-environment-config Config Map.",
								MarkdownDescription: "List of modules associated with database. Note - For Active-Active databases this feature is currently in preview. For this feature to take effect for Active-Active databases, set a boolean environment variable with the name 'ENABLE_ALPHA_FEATURES' to True. This variable can be set via the redis-enterprise-operator pod spec, or through the operator-environment-config Config Map.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"config": schema.StringAttribute{
											Description:         "Module command line arguments e.g. VKEY_MAX_ENTITY_COUNT 30",
											MarkdownDescription: "Module command line arguments e.g. VKEY_MAX_ENTITY_COUNT 30",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "The module's name e.g 'ft' for redissearch",
											MarkdownDescription: "The module's name e.g 'ft' for redissearch",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"uid": schema.StringAttribute{
											Description:         "Module's uid - do not set, for system use only nolint:staticcheck // custom json tag unknown to the linter",
											MarkdownDescription: "Module's uid - do not set, for system use only nolint:staticcheck // custom json tag unknown to the linter",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"version": schema.StringAttribute{
											Description:         "Module's semantic version e.g '1.6.12'",
											MarkdownDescription: "Module's semantic version e.g '1.6.12'",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"oss_cluster": schema.BoolAttribute{
								Description:         "OSS Cluster mode option. Note that not all client libraries support OSS cluster mode.",
								MarkdownDescription: "OSS Cluster mode option. Note that not all client libraries support OSS cluster mode.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"persistence": schema.StringAttribute{
								Description:         "Database on-disk persistence policy",
								MarkdownDescription: "Database on-disk persistence policy",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"proxy_policy": schema.StringAttribute{
								Description:         "The policy used for proxy binding to the endpoint. Supported proxy policies are: single/all-master-shards/all-nodes When left blank, the default value will be chosen according to the value of ossCluster - single if disabled, all-master-shards when enabled",
								MarkdownDescription: "The policy used for proxy binding to the endpoint. Supported proxy policies are: single/all-master-shards/all-nodes When left blank, the default value will be chosen according to the value of ossCluster - single if disabled, all-master-shards when enabled",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"rack_aware": schema.BoolAttribute{
								Description:         "Whether database should be rack aware. This improves availability - more information: https://docs.redislabs.com/latest/rs/concepts/high-availability/rack-zone-awareness/",
								MarkdownDescription: "Whether database should be rack aware. This improves availability - more information: https://docs.redislabs.com/latest/rs/concepts/high-availability/rack-zone-awareness/",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"redis_enterprise_cluster": schema.SingleNestedAttribute{
								Description:         "Connection to Redis Enterprise Cluster",
								MarkdownDescription: "Connection to Redis Enterprise Cluster",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "The name of the Redis Enterprise Cluster where the database should be stored.",
										MarkdownDescription: "The name of the Redis Enterprise Cluster where the database should be stored.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"redis_version": schema.StringAttribute{
								Description:         "Redis OSS version. For existing databases - Upgrade Redis OSS version. For new databases - the version which the database will be created with. If set to 'major' - will always upgrade to the most recent major Redis version. If set to 'latest' - will always upgrade to the most recent Redis version. Depends on 'redisUpgradePolicy' - if you want to set the value to 'latest' for some databases, you must set redisUpgradePolicy on the cluster before. Possible values are 'major' or 'latest' When using upgrade - make sure to backup the database before. This value is used only for database type 'redis'",
								MarkdownDescription: "Redis OSS version. For existing databases - Upgrade Redis OSS version. For new databases - the version which the database will be created with. If set to 'major' - will always upgrade to the most recent major Redis version. If set to 'latest' - will always upgrade to the most recent Redis version. Depends on 'redisUpgradePolicy' - if you want to set the value to 'latest' for some databases, you must set redisUpgradePolicy on the cluster before. Possible values are 'major' or 'latest' When using upgrade - make sure to backup the database before. This value is used only for database type 'redis'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replica_sources": schema.ListNestedAttribute{
								Description:         "What databases to replicate from",
								MarkdownDescription: "What databases to replicate from",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"client_key_secret": schema.StringAttribute{
											Description:         "Secret that defines the client certificate and key used by the syncer in the target database cluster. The secret must have 2 keys in its map: 'cert' which is the PEM encoded certificate, and 'key' which is the PEM encoded private key.",
											MarkdownDescription: "Secret that defines the client certificate and key used by the syncer in the target database cluster. The secret must have 2 keys in its map: 'cert' which is the PEM encoded certificate, and 'key' which is the PEM encoded private key.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"compression": schema.Int64Attribute{
											Description:         "GZIP compression level (0-6) to use for replication.",
											MarkdownDescription: "GZIP compression level (0-6) to use for replication.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"replica_source_name": schema.StringAttribute{
											Description:         "The name of the resource from which the source database URI is derived. The type of resource must match the type specified in the ReplicaSourceType field.",
											MarkdownDescription: "The name of the resource from which the source database URI is derived. The type of resource must match the type specified in the ReplicaSourceType field.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"replica_source_type": schema.StringAttribute{
											Description:         "The type of resource from which the source database URI is derived. If set to 'SECRET', the source database URI is derived from the secret named in the ReplicaSourceName field. The secret must have a key named 'uri' that defines the URI of the source database in the form of 'redis://...'. The type of secret (kubernetes, vault, ...) is determined by the secret mechanism used by the underlying REC object. If set to 'REDB', the source database URI is derived from the RedisEnterpriseDatabase resource named in the ReplicaSourceName field.",
											MarkdownDescription: "The type of resource from which the source database URI is derived. If set to 'SECRET', the source database URI is derived from the secret named in the ReplicaSourceName field. The secret must have a key named 'uri' that defines the URI of the source database in the form of 'redis://...'. The type of secret (kubernetes, vault, ...) is determined by the secret mechanism used by the underlying REC object. If set to 'REDB', the source database URI is derived from the RedisEnterpriseDatabase resource named in the ReplicaSourceName field.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"server_cert_secret": schema.StringAttribute{
											Description:         "Secret that defines the server certificate used by the proxy in the source database cluster. The secret must have 1 key in its map: 'cert' which is the PEM encoded certificate.",
											MarkdownDescription: "Secret that defines the server certificate used by the proxy in the source database cluster. The secret must have 1 key in its map: 'cert' which is the PEM encoded certificate.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"tls_sni_name": schema.StringAttribute{
											Description:         "TLS SNI name to use for the replication link.",
											MarkdownDescription: "TLS SNI name to use for the replication link.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"replication": schema.BoolAttribute{
								Description:         "In-memory database replication. When enabled, database will have replica shard for every master - leading to higher availability.",
								MarkdownDescription: "In-memory database replication. When enabled, database will have replica shard for every master - leading to higher availability.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"rof_ram_size": schema.StringAttribute{
								Description:         "The size of the RAM portion of an RoF database. Similarly to 'memorySize' use formats like 100MB, 0.1GB It must be at least 10% of combined memory size (RAM+Flash), as specified by 'memorySize'.",
								MarkdownDescription: "The size of the RAM portion of an RoF database. Similarly to 'memorySize' use formats like 100MB, 0.1GB It must be at least 10% of combined memory size (RAM+Flash), as specified by 'memorySize'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"roles_permissions": schema.ListNestedAttribute{
								Description:         "List of Redis Enteprise ACL and Role bindings to apply",
								MarkdownDescription: "List of Redis Enteprise ACL and Role bindings to apply",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"acl": schema.StringAttribute{
											Description:         "Acl Name of RolePermissionType (note: use exact name of the ACL from the Redis Enterprise ACL list, case sensitive)",
											MarkdownDescription: "Acl Name of RolePermissionType (note: use exact name of the ACL from the Redis Enterprise ACL list, case sensitive)",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"role": schema.StringAttribute{
											Description:         "Role Name of RolePermissionType (note: use exact name of the role from the Redis Enterprise role list, case sensitive)",
											MarkdownDescription: "Role Name of RolePermissionType (note: use exact name of the role from the Redis Enterprise role list, case sensitive)",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"type": schema.StringAttribute{
											Description:         "Type of Redis Enterprise Database Role Permission",
											MarkdownDescription: "Type of Redis Enterprise Database Role Permission",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"shard_count": schema.Int64Attribute{
								Description:         "Number of database server-side shards",
								MarkdownDescription: "Number of database server-side shards",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"shards_placement": schema.StringAttribute{
								Description:         "Control the density of shards - should they reside on as few or as many nodes as possible. Available options are 'dense' or 'sparse'. If left unset, defaults to 'dense'.",
								MarkdownDescription: "Control the density of shards - should they reside on as few or as many nodes as possible. Available options are 'dense' or 'sparse'. If left unset, defaults to 'dense'.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"tls_mode": schema.StringAttribute{
								Description:         "Require SSL authenticated and encrypted connections to the database. enabled - all incoming connections to the Database must use SSL. disabled - no incoming connection to the Database should use SSL. replica_ssl - databases that replicate from this one need to use SSL.",
								MarkdownDescription: "Require SSL authenticated and encrypted connections to the database. enabled - all incoming connections to the Database must use SSL. disabled - no incoming connection to the Database should use SSL. replica_ssl - databases that replicate from this one need to use SSL.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"type": schema.StringAttribute{
								Description:         "The type of the database.",
								MarkdownDescription: "The type of the database.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"participating_clusters": schema.ListNestedAttribute{
						Description:         "The list of instances/ clusters specifications and configurations.",
						MarkdownDescription: "The list of instances/ clusters specifications and configurations.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "The name of the remote cluster CR to link.",
									MarkdownDescription: "The name of the remote cluster CR to link.",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"redis_enterprise_cluster": schema.SingleNestedAttribute{
						Description:         "Connection to Redis Enterprise Cluster",
						MarkdownDescription: "Connection to Redis Enterprise Cluster",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "The name of the Redis Enterprise Cluster where the database should be stored.",
								MarkdownDescription: "The name of the Redis Enterprise Cluster where the database should be stored.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},
				},
				Required: false,
				Optional: false,
				Computed: true,
			},
		},
	}
}

func (r *AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *provider.DataSourceData, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_app_redislabs_com_redis_enterprise_active_active_database_v1alpha1")

	var data AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "app.redislabs.com", Version: "v1alpha1", Resource: "RedisEnterpriseActiveActiveDatabase"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse AppRedislabsComRedisEnterpriseActiveActiveDatabaseV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Name, data.Metadata.Namespace))
	data.ApiVersion = pointer.String("app.redislabs.com/v1alpha1")
	data.Kind = pointer.String("RedisEnterpriseActiveActiveDatabase")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}