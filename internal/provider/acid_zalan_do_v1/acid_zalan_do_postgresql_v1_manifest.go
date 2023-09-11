/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package acid_zalan_do_v1

import (
	"context"
	"fmt"
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
	_ datasource.DataSource = &AcidZalanDoPostgresqlV1Manifest{}
)

func NewAcidZalanDoPostgresqlV1Manifest() datasource.DataSource {
	return &AcidZalanDoPostgresqlV1Manifest{}
}

type AcidZalanDoPostgresqlV1Manifest struct{}

type AcidZalanDoPostgresqlV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
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
		AdditionalVolumes *[]struct {
			MountPath        *string            `tfsdk:"mount_path" json:"mountPath,omitempty"`
			Name             *string            `tfsdk:"name" json:"name,omitempty"`
			SubPath          *string            `tfsdk:"sub_path" json:"subPath,omitempty"`
			TargetContainers *[]string          `tfsdk:"target_containers" json:"targetContainers,omitempty"`
			VolumeSource     *map[string]string `tfsdk:"volume_source" json:"volumeSource,omitempty"`
		} `tfsdk:"additional_volumes" json:"additionalVolumes,omitempty"`
		AllowedSourceRanges *[]string `tfsdk:"allowed_source_ranges" json:"allowedSourceRanges,omitempty"`
		Clone               *struct {
			Cluster              *string `tfsdk:"cluster" json:"cluster,omitempty"`
			S3_access_key_id     *string `tfsdk:"s3_access_key_id" json:"s3_access_key_id,omitempty"`
			S3_endpoint          *string `tfsdk:"s3_endpoint" json:"s3_endpoint,omitempty"`
			S3_force_path_style  *bool   `tfsdk:"s3_force_path_style" json:"s3_force_path_style,omitempty"`
			S3_secret_access_key *string `tfsdk:"s3_secret_access_key" json:"s3_secret_access_key,omitempty"`
			S3_wal_path          *string `tfsdk:"s3_wal_path" json:"s3_wal_path,omitempty"`
			Timestamp            *string `tfsdk:"timestamp" json:"timestamp,omitempty"`
			Uid                  *string `tfsdk:"uid" json:"uid,omitempty"`
		} `tfsdk:"clone" json:"clone,omitempty"`
		ConnectionPooler *struct {
			DockerImage       *string `tfsdk:"docker_image" json:"dockerImage,omitempty"`
			MaxDBConnections  *int64  `tfsdk:"max_db_connections" json:"maxDBConnections,omitempty"`
			Mode              *string `tfsdk:"mode" json:"mode,omitempty"`
			NumberOfInstances *int64  `tfsdk:"number_of_instances" json:"numberOfInstances,omitempty"`
			Resources         *struct {
				Limits *struct {
					Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *string `tfsdk:"memory" json:"memory,omitempty"`
				} `tfsdk:"limits" json:"limits,omitempty"`
				Requests *struct {
					Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
					Memory *string `tfsdk:"memory" json:"memory,omitempty"`
				} `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			Schema *string `tfsdk:"schema" json:"schema,omitempty"`
			User   *string `tfsdk:"user" json:"user,omitempty"`
		} `tfsdk:"connection_pooler" json:"connectionPooler,omitempty"`
		Databases                       *map[string]string   `tfsdk:"databases" json:"databases,omitempty"`
		DockerImage                     *string              `tfsdk:"docker_image" json:"dockerImage,omitempty"`
		EnableConnectionPooler          *bool                `tfsdk:"enable_connection_pooler" json:"enableConnectionPooler,omitempty"`
		EnableLogicalBackup             *bool                `tfsdk:"enable_logical_backup" json:"enableLogicalBackup,omitempty"`
		EnableMasterLoadBalancer        *bool                `tfsdk:"enable_master_load_balancer" json:"enableMasterLoadBalancer,omitempty"`
		EnableMasterPoolerLoadBalancer  *bool                `tfsdk:"enable_master_pooler_load_balancer" json:"enableMasterPoolerLoadBalancer,omitempty"`
		EnableReplicaConnectionPooler   *bool                `tfsdk:"enable_replica_connection_pooler" json:"enableReplicaConnectionPooler,omitempty"`
		EnableReplicaLoadBalancer       *bool                `tfsdk:"enable_replica_load_balancer" json:"enableReplicaLoadBalancer,omitempty"`
		EnableReplicaPoolerLoadBalancer *bool                `tfsdk:"enable_replica_pooler_load_balancer" json:"enableReplicaPoolerLoadBalancer,omitempty"`
		EnableShmVolume                 *bool                `tfsdk:"enable_shm_volume" json:"enableShmVolume,omitempty"`
		Env                             *[]map[string]string `tfsdk:"env" json:"env,omitempty"`
		InitContainers                  *[]map[string]string `tfsdk:"init_containers" json:"initContainers,omitempty"`
		LogicalBackupSchedule           *string              `tfsdk:"logical_backup_schedule" json:"logicalBackupSchedule,omitempty"`
		MaintenanceWindows              *[]string            `tfsdk:"maintenance_windows" json:"maintenanceWindows,omitempty"`
		MasterServiceAnnotations        *map[string]string   `tfsdk:"master_service_annotations" json:"masterServiceAnnotations,omitempty"`
		NodeAffinity                    *struct {
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
		NumberOfInstances *int64 `tfsdk:"number_of_instances" json:"numberOfInstances,omitempty"`
		Patroni           *struct {
			Failsafe_mode           *bool                         `tfsdk:"failsafe_mode" json:"failsafe_mode,omitempty"`
			Initdb                  *map[string]string            `tfsdk:"initdb" json:"initdb,omitempty"`
			Loop_wait               *int64                        `tfsdk:"loop_wait" json:"loop_wait,omitempty"`
			Maximum_lag_on_failover *int64                        `tfsdk:"maximum_lag_on_failover" json:"maximum_lag_on_failover,omitempty"`
			Pg_hba                  *[]string                     `tfsdk:"pg_hba" json:"pg_hba,omitempty"`
			Retry_timeout           *int64                        `tfsdk:"retry_timeout" json:"retry_timeout,omitempty"`
			Slots                   *map[string]map[string]string `tfsdk:"slots" json:"slots,omitempty"`
			Synchronous_mode        *bool                         `tfsdk:"synchronous_mode" json:"synchronous_mode,omitempty"`
			Synchronous_mode_strict *bool                         `tfsdk:"synchronous_mode_strict" json:"synchronous_mode_strict,omitempty"`
			Synchronous_node_count  *int64                        `tfsdk:"synchronous_node_count" json:"synchronous_node_count,omitempty"`
			Ttl                     *int64                        `tfsdk:"ttl" json:"ttl,omitempty"`
		} `tfsdk:"patroni" json:"patroni,omitempty"`
		PodAnnotations       *map[string]string `tfsdk:"pod_annotations" json:"podAnnotations,omitempty"`
		PodPriorityClassName *string            `tfsdk:"pod_priority_class_name" json:"podPriorityClassName,omitempty"`
		Postgresql           *struct {
			Parameters *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
			Version    *string            `tfsdk:"version" json:"version,omitempty"`
		} `tfsdk:"postgresql" json:"postgresql,omitempty"`
		PreparedDatabases *struct {
			DefaultUsers *bool              `tfsdk:"default_users" json:"defaultUsers,omitempty"`
			Extensions   *map[string]string `tfsdk:"extensions" json:"extensions,omitempty"`
			Schemas      *struct {
				DefaultRoles *bool `tfsdk:"default_roles" json:"defaultRoles,omitempty"`
				DefaultUsers *bool `tfsdk:"default_users" json:"defaultUsers,omitempty"`
			} `tfsdk:"schemas" json:"schemas,omitempty"`
			SecretNamespace *string `tfsdk:"secret_namespace" json:"secretNamespace,omitempty"`
		} `tfsdk:"prepared_databases" json:"preparedDatabases,omitempty"`
		ReplicaLoadBalancer       *bool              `tfsdk:"replica_load_balancer" json:"replicaLoadBalancer,omitempty"`
		ReplicaServiceAnnotations *map[string]string `tfsdk:"replica_service_annotations" json:"replicaServiceAnnotations,omitempty"`
		Resources                 *struct {
			Limits *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"limits" json:"limits,omitempty"`
			Requests *struct {
				Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
				Memory *string `tfsdk:"memory" json:"memory,omitempty"`
			} `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		SchedulerName      *string              `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
		ServiceAnnotations *map[string]string   `tfsdk:"service_annotations" json:"serviceAnnotations,omitempty"`
		Sidecars           *[]map[string]string `tfsdk:"sidecars" json:"sidecars,omitempty"`
		SpiloFSGroup       *int64               `tfsdk:"spilo_fs_group" json:"spiloFSGroup,omitempty"`
		SpiloRunAsGroup    *int64               `tfsdk:"spilo_run_as_group" json:"spiloRunAsGroup,omitempty"`
		SpiloRunAsUser     *int64               `tfsdk:"spilo_run_as_user" json:"spiloRunAsUser,omitempty"`
		Standby            *struct {
			Gs_wal_path  *string `tfsdk:"gs_wal_path" json:"gs_wal_path,omitempty"`
			S3_wal_path  *string `tfsdk:"s3_wal_path" json:"s3_wal_path,omitempty"`
			Standby_host *string `tfsdk:"standby_host" json:"standby_host,omitempty"`
			Standby_port *string `tfsdk:"standby_port" json:"standby_port,omitempty"`
		} `tfsdk:"standby" json:"standby,omitempty"`
		Streams *[]struct {
			ApplicationId *string            `tfsdk:"application_id" json:"applicationId,omitempty"`
			BatchSize     *int64             `tfsdk:"batch_size" json:"batchSize,omitempty"`
			Database      *string            `tfsdk:"database" json:"database,omitempty"`
			Filter        *map[string]string `tfsdk:"filter" json:"filter,omitempty"`
			Tables        *struct {
				EventType     *string `tfsdk:"event_type" json:"eventType,omitempty"`
				IdColumn      *string `tfsdk:"id_column" json:"idColumn,omitempty"`
				PayloadColumn *string `tfsdk:"payload_column" json:"payloadColumn,omitempty"`
			} `tfsdk:"tables" json:"tables,omitempty"`
		} `tfsdk:"streams" json:"streams,omitempty"`
		TeamId *string `tfsdk:"team_id" json:"teamId,omitempty"`
		Tls    *struct {
			CaFile          *string `tfsdk:"ca_file" json:"caFile,omitempty"`
			CaSecretName    *string `tfsdk:"ca_secret_name" json:"caSecretName,omitempty"`
			CertificateFile *string `tfsdk:"certificate_file" json:"certificateFile,omitempty"`
			PrivateKeyFile  *string `tfsdk:"private_key_file" json:"privateKeyFile,omitempty"`
			SecretName      *string `tfsdk:"secret_name" json:"secretName,omitempty"`
		} `tfsdk:"tls" json:"tls,omitempty"`
		Tolerations *[]struct {
			Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
			Key               *string `tfsdk:"key" json:"key,omitempty"`
			Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
			TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
			Value             *string `tfsdk:"value" json:"value,omitempty"`
		} `tfsdk:"tolerations" json:"tolerations,omitempty"`
		UseLoadBalancer                *bool                `tfsdk:"use_load_balancer" json:"useLoadBalancer,omitempty"`
		Users                          *map[string][]string `tfsdk:"users" json:"users,omitempty"`
		UsersWithInPlaceSecretRotation *[]string            `tfsdk:"users_with_in_place_secret_rotation" json:"usersWithInPlaceSecretRotation,omitempty"`
		UsersWithSecretRotation        *[]string            `tfsdk:"users_with_secret_rotation" json:"usersWithSecretRotation,omitempty"`
		Volume                         *struct {
			Iops     *int64 `tfsdk:"iops" json:"iops,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
			Size         *string `tfsdk:"size" json:"size,omitempty"`
			StorageClass *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
			SubPath      *string `tfsdk:"sub_path" json:"subPath,omitempty"`
			Throughput   *int64  `tfsdk:"throughput" json:"throughput,omitempty"`
		} `tfsdk:"volume" json:"volume,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AcidZalanDoPostgresqlV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_acid_zalan_do_postgresql_v1_manifest"
}

func (r *AcidZalanDoPostgresqlV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

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
					"additional_volumes": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"mount_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"sub_path": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"target_containers": schema.ListAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"volume_source": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
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

					"allowed_source_ranges": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"clone": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"cluster": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"s3_access_key_id": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_endpoint": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_force_path_style": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_secret_access_key": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_wal_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"timestamp": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9]+)-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])[Tt]([01][0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9]|60)(\.[0-9]+)?(([+-]([01][0-9]|2[0-3]):[0-5][0-9]))$`), ""),
								},
							},

							"uid": schema.StringAttribute{
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
							"docker_image": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"max_db_connections": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"mode": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("session", "transaction"),
								},
							},

							"number_of_instances": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
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
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
												},
											},

											"memory": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
												},
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
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
												},
											},

											"memory": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
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

							"schema": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"user": schema.StringAttribute{
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

					"databases": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"docker_image": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_connection_pooler": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_logical_backup": schema.BoolAttribute{
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

					"enable_replica_connection_pooler": schema.BoolAttribute{
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

					"enable_shm_volume": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"env": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"init_containers": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.MapType{ElemType: types.StringType},
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

					"maintenance_windows": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"master_service_annotations": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

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

					"number_of_instances": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"patroni": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"failsafe_mode": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"initdb": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"loop_wait": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"maximum_lag_on_failover": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pg_hba": schema.ListAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retry_timeout": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"slots": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.MapType{ElemType: types.StringType},
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"synchronous_mode": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"synchronous_mode_strict": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"synchronous_node_count": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ttl": schema.Int64Attribute{
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

					"pod_annotations": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"pod_priority_class_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgresql": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"parameters": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"version": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("10", "11", "12", "13", "14", "15"),
								},
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},

					"prepared_databases": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"default_users": schema.BoolAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"extensions": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"schemas": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"default_roles": schema.BoolAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"default_users": schema.BoolAttribute{
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

							"secret_namespace": schema.StringAttribute{
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

					"replica_load_balancer": schema.BoolAttribute{
						Description:         "deprecated",
						MarkdownDescription: "deprecated",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"replica_service_annotations": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
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
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
										},
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
										},
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
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+m|\d+(\.\d{1,3})?)$`), ""),
										},
									},

									"memory": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
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

					"scheduler_name": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"service_annotations": schema.MapAttribute{
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

					"spilo_fs_group": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"spilo_run_as_group": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"spilo_run_as_user": schema.Int64Attribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"standby": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"gs_wal_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"s3_wal_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"standby_host": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"standby_port": schema.StringAttribute{
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

					"streams": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"application_id": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"batch_size": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"database": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"filter": schema.MapAttribute{
									Description:         "",
									MarkdownDescription: "",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tables": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"event_type": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"id_column": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"payload_column": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"team_id": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"tls": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"ca_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ca_secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"certificate_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"private_key_file": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"secret_name": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tolerations": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"effect": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("NoExecute", "NoSchedule", "PreferNoSchedule"),
									},
								},

								"key": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"operator": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Equal", "Exists"),
									},
								},

								"toleration_seconds": schema.Int64Attribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value": schema.StringAttribute{
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

					"use_load_balancer": schema.BoolAttribute{
						Description:         "deprecated",
						MarkdownDescription: "deprecated",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"users": schema.MapAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.ListType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"users_with_in_place_secret_rotation": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"users_with_secret_rotation": schema.ListAttribute{
						Description:         "",
						MarkdownDescription: "",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"iops": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"selector": schema.SingleNestedAttribute{
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
													Validators: []validator.String{
														stringvalidator.OneOf("DoesNotExist", "Exists", "In", "NotIn"),
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"size": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(\d+(e\d+)?|\d+(\.\d+)?(e\d+)?[EPTGMK]i?)$`), ""),
								},
							},

							"storage_class": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sub_path": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"throughput": schema.Int64Attribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: true,
						Optional: false,
						Computed: false,
					},
				},
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *AcidZalanDoPostgresqlV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_acid_zalan_do_postgresql_v1_manifest")

	var model AcidZalanDoPostgresqlV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Namespace, model.Metadata.Name))
	model.ApiVersion = pointer.String("acid.zalan.do/v1")
	model.Kind = pointer.String("postgresql")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
