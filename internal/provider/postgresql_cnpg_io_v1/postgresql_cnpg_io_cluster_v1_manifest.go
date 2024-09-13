/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package postgresql_cnpg_io_v1

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
	_ datasource.DataSource = &PostgresqlCnpgIoClusterV1Manifest{}
)

func NewPostgresqlCnpgIoClusterV1Manifest() datasource.DataSource {
	return &PostgresqlCnpgIoClusterV1Manifest{}
}

type PostgresqlCnpgIoClusterV1Manifest struct{}

type PostgresqlCnpgIoClusterV1ManifestData struct {
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
		Affinity *struct {
			AdditionalPodAffinity *struct {
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
						MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
						MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
					MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
					MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
			} `tfsdk:"additional_pod_affinity" json:"additionalPodAffinity,omitempty"`
			AdditionalPodAntiAffinity *struct {
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
						MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
						MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
					MatchLabelKeys    *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
					MismatchLabelKeys *[]string `tfsdk:"mismatch_label_keys" json:"mismatchLabelKeys,omitempty"`
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
			} `tfsdk:"additional_pod_anti_affinity" json:"additionalPodAntiAffinity,omitempty"`
			EnablePodAntiAffinity *bool `tfsdk:"enable_pod_anti_affinity" json:"enablePodAntiAffinity,omitempty"`
			NodeAffinity          *struct {
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
			NodeSelector        *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PodAntiAffinityType *string            `tfsdk:"pod_anti_affinity_type" json:"podAntiAffinityType,omitempty"`
			Tolerations         *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			TopologyKey *string `tfsdk:"topology_key" json:"topologyKey,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		Backup *struct {
			BarmanObjectStore *struct {
				AzureCredentials *struct {
					ConnectionString *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"connection_string" json:"connectionString,omitempty"`
					InheritFromAzureAD *bool `tfsdk:"inherit_from_azure_ad" json:"inheritFromAzureAD,omitempty"`
					StorageAccount     *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"storage_account" json:"storageAccount,omitempty"`
					StorageKey *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"storage_key" json:"storageKey,omitempty"`
					StorageSasToken *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"storage_sas_token" json:"storageSasToken,omitempty"`
				} `tfsdk:"azure_credentials" json:"azureCredentials,omitempty"`
				Data *struct {
					AdditionalCommandArgs *[]string `tfsdk:"additional_command_args" json:"additionalCommandArgs,omitempty"`
					Compression           *string   `tfsdk:"compression" json:"compression,omitempty"`
					Encryption            *string   `tfsdk:"encryption" json:"encryption,omitempty"`
					ImmediateCheckpoint   *bool     `tfsdk:"immediate_checkpoint" json:"immediateCheckpoint,omitempty"`
					Jobs                  *int64    `tfsdk:"jobs" json:"jobs,omitempty"`
				} `tfsdk:"data" json:"data,omitempty"`
				DestinationPath *string `tfsdk:"destination_path" json:"destinationPath,omitempty"`
				EndpointCA      *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"endpoint_ca" json:"endpointCA,omitempty"`
				EndpointURL       *string `tfsdk:"endpoint_url" json:"endpointURL,omitempty"`
				GoogleCredentials *struct {
					ApplicationCredentials *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"application_credentials" json:"applicationCredentials,omitempty"`
					GkeEnvironment *bool `tfsdk:"gke_environment" json:"gkeEnvironment,omitempty"`
				} `tfsdk:"google_credentials" json:"googleCredentials,omitempty"`
				HistoryTags   *map[string]string `tfsdk:"history_tags" json:"historyTags,omitempty"`
				S3Credentials *struct {
					AccessKeyId *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
					InheritFromIAMRole *bool `tfsdk:"inherit_from_iam_role" json:"inheritFromIAMRole,omitempty"`
					Region             *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"region" json:"region,omitempty"`
					SecretAccessKey *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
					SessionToken *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"session_token" json:"sessionToken,omitempty"`
				} `tfsdk:"s3_credentials" json:"s3Credentials,omitempty"`
				ServerName *string            `tfsdk:"server_name" json:"serverName,omitempty"`
				Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
				Wal        *struct {
					ArchiveAdditionalCommandArgs *[]string `tfsdk:"archive_additional_command_args" json:"archiveAdditionalCommandArgs,omitempty"`
					Compression                  *string   `tfsdk:"compression" json:"compression,omitempty"`
					Encryption                   *string   `tfsdk:"encryption" json:"encryption,omitempty"`
					MaxParallel                  *int64    `tfsdk:"max_parallel" json:"maxParallel,omitempty"`
					RestoreAdditionalCommandArgs *[]string `tfsdk:"restore_additional_command_args" json:"restoreAdditionalCommandArgs,omitempty"`
				} `tfsdk:"wal" json:"wal,omitempty"`
			} `tfsdk:"barman_object_store" json:"barmanObjectStore,omitempty"`
			RetentionPolicy *string `tfsdk:"retention_policy" json:"retentionPolicy,omitempty"`
			Target          *string `tfsdk:"target" json:"target,omitempty"`
			VolumeSnapshot  *struct {
				Annotations         *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				ClassName           *string            `tfsdk:"class_name" json:"className,omitempty"`
				Labels              *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Online              *bool              `tfsdk:"online" json:"online,omitempty"`
				OnlineConfiguration *struct {
					ImmediateCheckpoint *bool `tfsdk:"immediate_checkpoint" json:"immediateCheckpoint,omitempty"`
					WaitForArchive      *bool `tfsdk:"wait_for_archive" json:"waitForArchive,omitempty"`
				} `tfsdk:"online_configuration" json:"onlineConfiguration,omitempty"`
				SnapshotOwnerReference *string            `tfsdk:"snapshot_owner_reference" json:"snapshotOwnerReference,omitempty"`
				TablespaceClassName    *map[string]string `tfsdk:"tablespace_class_name" json:"tablespaceClassName,omitempty"`
				WalClassName           *string            `tfsdk:"wal_class_name" json:"walClassName,omitempty"`
			} `tfsdk:"volume_snapshot" json:"volumeSnapshot,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		Bootstrap *struct {
			Initdb *struct {
				DataChecksums *bool   `tfsdk:"data_checksums" json:"dataChecksums,omitempty"`
				Database      *string `tfsdk:"database" json:"database,omitempty"`
				Encoding      *string `tfsdk:"encoding" json:"encoding,omitempty"`
				Import        *struct {
					Databases                *[]string `tfsdk:"databases" json:"databases,omitempty"`
					PostImportApplicationSQL *[]string `tfsdk:"post_import_application_sql" json:"postImportApplicationSQL,omitempty"`
					Roles                    *[]string `tfsdk:"roles" json:"roles,omitempty"`
					SchemaOnly               *bool     `tfsdk:"schema_only" json:"schemaOnly,omitempty"`
					Source                   *struct {
						ExternalCluster *string `tfsdk:"external_cluster" json:"externalCluster,omitempty"`
					} `tfsdk:"source" json:"source,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"import" json:"import,omitempty"`
				LocaleCType                *string   `tfsdk:"locale_c_type" json:"localeCType,omitempty"`
				LocaleCollate              *string   `tfsdk:"locale_collate" json:"localeCollate,omitempty"`
				Options                    *[]string `tfsdk:"options" json:"options,omitempty"`
				Owner                      *string   `tfsdk:"owner" json:"owner,omitempty"`
				PostInitApplicationSQL     *[]string `tfsdk:"post_init_application_sql" json:"postInitApplicationSQL,omitempty"`
				PostInitApplicationSQLRefs *struct {
					ConfigMapRefs *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_refs" json:"configMapRefs,omitempty"`
					SecretRefs *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
				} `tfsdk:"post_init_application_sql_refs" json:"postInitApplicationSQLRefs,omitempty"`
				PostInitSQL     *[]string `tfsdk:"post_init_sql" json:"postInitSQL,omitempty"`
				PostInitSQLRefs *struct {
					ConfigMapRefs *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_refs" json:"configMapRefs,omitempty"`
					SecretRefs *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
				} `tfsdk:"post_init_sql_refs" json:"postInitSQLRefs,omitempty"`
				PostInitTemplateSQL     *[]string `tfsdk:"post_init_template_sql" json:"postInitTemplateSQL,omitempty"`
				PostInitTemplateSQLRefs *struct {
					ConfigMapRefs *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_map_refs" json:"configMapRefs,omitempty"`
					SecretRefs *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
				} `tfsdk:"post_init_template_sql_refs" json:"postInitTemplateSQLRefs,omitempty"`
				Secret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				WalSegmentSize *int64 `tfsdk:"wal_segment_size" json:"walSegmentSize,omitempty"`
			} `tfsdk:"initdb" json:"initdb,omitempty"`
			Pg_basebackup *struct {
				Database *string `tfsdk:"database" json:"database,omitempty"`
				Owner    *string `tfsdk:"owner" json:"owner,omitempty"`
				Secret   *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Source *string `tfsdk:"source" json:"source,omitempty"`
			} `tfsdk:"pg_basebackup" json:"pg_basebackup,omitempty"`
			Recovery *struct {
				Backup *struct {
					EndpointCA *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"endpoint_ca" json:"endpointCA,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"backup" json:"backup,omitempty"`
				Database       *string `tfsdk:"database" json:"database,omitempty"`
				Owner          *string `tfsdk:"owner" json:"owner,omitempty"`
				RecoveryTarget *struct {
					BackupID        *string `tfsdk:"backup_id" json:"backupID,omitempty"`
					Exclusive       *bool   `tfsdk:"exclusive" json:"exclusive,omitempty"`
					TargetImmediate *bool   `tfsdk:"target_immediate" json:"targetImmediate,omitempty"`
					TargetLSN       *string `tfsdk:"target_lsn" json:"targetLSN,omitempty"`
					TargetName      *string `tfsdk:"target_name" json:"targetName,omitempty"`
					TargetTLI       *string `tfsdk:"target_tli" json:"targetTLI,omitempty"`
					TargetTime      *string `tfsdk:"target_time" json:"targetTime,omitempty"`
					TargetXID       *string `tfsdk:"target_xid" json:"targetXID,omitempty"`
				} `tfsdk:"recovery_target" json:"recoveryTarget,omitempty"`
				Secret *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				Source          *string `tfsdk:"source" json:"source,omitempty"`
				VolumeSnapshots *struct {
					Storage *struct {
						ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"storage" json:"storage,omitempty"`
					TablespaceStorage *struct {
						ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"tablespace_storage" json:"tablespaceStorage,omitempty"`
					WalStorage *struct {
						ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"wal_storage" json:"walStorage,omitempty"`
				} `tfsdk:"volume_snapshots" json:"volumeSnapshots,omitempty"`
			} `tfsdk:"recovery" json:"recovery,omitempty"`
		} `tfsdk:"bootstrap" json:"bootstrap,omitempty"`
		Certificates *struct {
			ClientCASecret       *string   `tfsdk:"client_ca_secret" json:"clientCASecret,omitempty"`
			ReplicationTLSSecret *string   `tfsdk:"replication_tls_secret" json:"replicationTLSSecret,omitempty"`
			ServerAltDNSNames    *[]string `tfsdk:"server_alt_dns_names" json:"serverAltDNSNames,omitempty"`
			ServerCASecret       *string   `tfsdk:"server_ca_secret" json:"serverCASecret,omitempty"`
			ServerTLSSecret      *string   `tfsdk:"server_tls_secret" json:"serverTLSSecret,omitempty"`
		} `tfsdk:"certificates" json:"certificates,omitempty"`
		Description           *string `tfsdk:"description" json:"description,omitempty"`
		EnablePDB             *bool   `tfsdk:"enable_pdb" json:"enablePDB,omitempty"`
		EnableSuperuserAccess *bool   `tfsdk:"enable_superuser_access" json:"enableSuperuserAccess,omitempty"`
		Env                   *[]struct {
			Name      *string `tfsdk:"name" json:"name,omitempty"`
			Value     *string `tfsdk:"value" json:"value,omitempty"`
			ValueFrom *struct {
				ConfigMapKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_key_ref" json:"configMapKeyRef,omitempty"`
				FieldRef *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
				} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
				ResourceFieldRef *struct {
					ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
					Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
					Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
				} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
				SecretKeyRef *struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_key_ref" json:"secretKeyRef,omitempty"`
			} `tfsdk:"value_from" json:"valueFrom,omitempty"`
		} `tfsdk:"env" json:"env,omitempty"`
		EnvFrom *[]struct {
			ConfigMapRef *struct {
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
			Prefix    *string `tfsdk:"prefix" json:"prefix,omitempty"`
			SecretRef *struct {
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
		} `tfsdk:"env_from" json:"envFrom,omitempty"`
		EphemeralVolumeSource *struct {
			VolumeClaimTemplate *struct {
				Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
				Spec     *struct {
					AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
					DataSource  *struct {
						ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"data_source" json:"dataSource,omitempty"`
					DataSourceRef *struct {
						ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
					Resources *struct {
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Selector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					StorageClassName          *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
					VolumeAttributesClassName *string `tfsdk:"volume_attributes_class_name" json:"volumeAttributesClassName,omitempty"`
					VolumeMode                *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					VolumeName                *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
		} `tfsdk:"ephemeral_volume_source" json:"ephemeralVolumeSource,omitempty"`
		EphemeralVolumesSizeLimit *struct {
			Shm           *string `tfsdk:"shm" json:"shm,omitempty"`
			TemporaryData *string `tfsdk:"temporary_data" json:"temporaryData,omitempty"`
		} `tfsdk:"ephemeral_volumes_size_limit" json:"ephemeralVolumesSizeLimit,omitempty"`
		ExternalClusters *[]struct {
			BarmanObjectStore *struct {
				AzureCredentials *struct {
					ConnectionString *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"connection_string" json:"connectionString,omitempty"`
					InheritFromAzureAD *bool `tfsdk:"inherit_from_azure_ad" json:"inheritFromAzureAD,omitempty"`
					StorageAccount     *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"storage_account" json:"storageAccount,omitempty"`
					StorageKey *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"storage_key" json:"storageKey,omitempty"`
					StorageSasToken *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"storage_sas_token" json:"storageSasToken,omitempty"`
				} `tfsdk:"azure_credentials" json:"azureCredentials,omitempty"`
				Data *struct {
					AdditionalCommandArgs *[]string `tfsdk:"additional_command_args" json:"additionalCommandArgs,omitempty"`
					Compression           *string   `tfsdk:"compression" json:"compression,omitempty"`
					Encryption            *string   `tfsdk:"encryption" json:"encryption,omitempty"`
					ImmediateCheckpoint   *bool     `tfsdk:"immediate_checkpoint" json:"immediateCheckpoint,omitempty"`
					Jobs                  *int64    `tfsdk:"jobs" json:"jobs,omitempty"`
				} `tfsdk:"data" json:"data,omitempty"`
				DestinationPath *string `tfsdk:"destination_path" json:"destinationPath,omitempty"`
				EndpointCA      *struct {
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"endpoint_ca" json:"endpointCA,omitempty"`
				EndpointURL       *string `tfsdk:"endpoint_url" json:"endpointURL,omitempty"`
				GoogleCredentials *struct {
					ApplicationCredentials *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"application_credentials" json:"applicationCredentials,omitempty"`
					GkeEnvironment *bool `tfsdk:"gke_environment" json:"gkeEnvironment,omitempty"`
				} `tfsdk:"google_credentials" json:"googleCredentials,omitempty"`
				HistoryTags   *map[string]string `tfsdk:"history_tags" json:"historyTags,omitempty"`
				S3Credentials *struct {
					AccessKeyId *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"access_key_id" json:"accessKeyId,omitempty"`
					InheritFromIAMRole *bool `tfsdk:"inherit_from_iam_role" json:"inheritFromIAMRole,omitempty"`
					Region             *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"region" json:"region,omitempty"`
					SecretAccessKey *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_access_key" json:"secretAccessKey,omitempty"`
					SessionToken *struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"session_token" json:"sessionToken,omitempty"`
				} `tfsdk:"s3_credentials" json:"s3Credentials,omitempty"`
				ServerName *string            `tfsdk:"server_name" json:"serverName,omitempty"`
				Tags       *map[string]string `tfsdk:"tags" json:"tags,omitempty"`
				Wal        *struct {
					ArchiveAdditionalCommandArgs *[]string `tfsdk:"archive_additional_command_args" json:"archiveAdditionalCommandArgs,omitempty"`
					Compression                  *string   `tfsdk:"compression" json:"compression,omitempty"`
					Encryption                   *string   `tfsdk:"encryption" json:"encryption,omitempty"`
					MaxParallel                  *int64    `tfsdk:"max_parallel" json:"maxParallel,omitempty"`
					RestoreAdditionalCommandArgs *[]string `tfsdk:"restore_additional_command_args" json:"restoreAdditionalCommandArgs,omitempty"`
				} `tfsdk:"wal" json:"wal,omitempty"`
			} `tfsdk:"barman_object_store" json:"barmanObjectStore,omitempty"`
			ConnectionParameters *map[string]string `tfsdk:"connection_parameters" json:"connectionParameters,omitempty"`
			Name                 *string            `tfsdk:"name" json:"name,omitempty"`
			Password             *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"password" json:"password,omitempty"`
			SslCert *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"ssl_cert" json:"sslCert,omitempty"`
			SslKey *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"ssl_key" json:"sslKey,omitempty"`
			SslRootCert *struct {
				Key      *string `tfsdk:"key" json:"key,omitempty"`
				Name     *string `tfsdk:"name" json:"name,omitempty"`
				Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
			} `tfsdk:"ssl_root_cert" json:"sslRootCert,omitempty"`
		} `tfsdk:"external_clusters" json:"externalClusters,omitempty"`
		FailoverDelay   *int64 `tfsdk:"failover_delay" json:"failoverDelay,omitempty"`
		ImageCatalogRef *struct {
			ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
			Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
			Major    *int64  `tfsdk:"major" json:"major,omitempty"`
			Name     *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_catalog_ref" json:"imageCatalogRef,omitempty"`
		ImageName        *string `tfsdk:"image_name" json:"imageName,omitempty"`
		ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
		InheritedMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"inherited_metadata" json:"inheritedMetadata,omitempty"`
		Instances            *int64  `tfsdk:"instances" json:"instances,omitempty"`
		LivenessProbeTimeout *int64  `tfsdk:"liveness_probe_timeout" json:"livenessProbeTimeout,omitempty"`
		LogLevel             *string `tfsdk:"log_level" json:"logLevel,omitempty"`
		Managed              *struct {
			Roles *[]struct {
				Bypassrls       *bool     `tfsdk:"bypassrls" json:"bypassrls,omitempty"`
				Comment         *string   `tfsdk:"comment" json:"comment,omitempty"`
				ConnectionLimit *int64    `tfsdk:"connection_limit" json:"connectionLimit,omitempty"`
				Createdb        *bool     `tfsdk:"createdb" json:"createdb,omitempty"`
				Createrole      *bool     `tfsdk:"createrole" json:"createrole,omitempty"`
				DisablePassword *bool     `tfsdk:"disable_password" json:"disablePassword,omitempty"`
				Ensure          *string   `tfsdk:"ensure" json:"ensure,omitempty"`
				InRoles         *[]string `tfsdk:"in_roles" json:"inRoles,omitempty"`
				Inherit         *bool     `tfsdk:"inherit" json:"inherit,omitempty"`
				Login           *bool     `tfsdk:"login" json:"login,omitempty"`
				Name            *string   `tfsdk:"name" json:"name,omitempty"`
				PasswordSecret  *struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"password_secret" json:"passwordSecret,omitempty"`
				Replication *bool   `tfsdk:"replication" json:"replication,omitempty"`
				Superuser   *bool   `tfsdk:"superuser" json:"superuser,omitempty"`
				ValidUntil  *string `tfsdk:"valid_until" json:"validUntil,omitempty"`
			} `tfsdk:"roles" json:"roles,omitempty"`
			Services *struct {
				Additional *[]struct {
					SelectorType    *string `tfsdk:"selector_type" json:"selectorType,omitempty"`
					ServiceTemplate *struct {
						Metadata *struct {
							Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
							Name        *string            `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"metadata" json:"metadata,omitempty"`
						Spec *struct {
							AllocateLoadBalancerNodePorts *bool     `tfsdk:"allocate_load_balancer_node_ports" json:"allocateLoadBalancerNodePorts,omitempty"`
							ClusterIP                     *string   `tfsdk:"cluster_ip" json:"clusterIP,omitempty"`
							ClusterIPs                    *[]string `tfsdk:"cluster_i_ps" json:"clusterIPs,omitempty"`
							ExternalIPs                   *[]string `tfsdk:"external_i_ps" json:"externalIPs,omitempty"`
							ExternalName                  *string   `tfsdk:"external_name" json:"externalName,omitempty"`
							ExternalTrafficPolicy         *string   `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
							HealthCheckNodePort           *int64    `tfsdk:"health_check_node_port" json:"healthCheckNodePort,omitempty"`
							InternalTrafficPolicy         *string   `tfsdk:"internal_traffic_policy" json:"internalTrafficPolicy,omitempty"`
							IpFamilies                    *[]string `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
							IpFamilyPolicy                *string   `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
							LoadBalancerClass             *string   `tfsdk:"load_balancer_class" json:"loadBalancerClass,omitempty"`
							LoadBalancerIP                *string   `tfsdk:"load_balancer_ip" json:"loadBalancerIP,omitempty"`
							LoadBalancerSourceRanges      *[]string `tfsdk:"load_balancer_source_ranges" json:"loadBalancerSourceRanges,omitempty"`
							Ports                         *[]struct {
								AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
								Name        *string `tfsdk:"name" json:"name,omitempty"`
								NodePort    *int64  `tfsdk:"node_port" json:"nodePort,omitempty"`
								Port        *int64  `tfsdk:"port" json:"port,omitempty"`
								Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
								TargetPort  *string `tfsdk:"target_port" json:"targetPort,omitempty"`
							} `tfsdk:"ports" json:"ports,omitempty"`
							PublishNotReadyAddresses *bool              `tfsdk:"publish_not_ready_addresses" json:"publishNotReadyAddresses,omitempty"`
							Selector                 *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
							SessionAffinity          *string            `tfsdk:"session_affinity" json:"sessionAffinity,omitempty"`
							SessionAffinityConfig    *struct {
								ClientIP *struct {
									TimeoutSeconds *int64 `tfsdk:"timeout_seconds" json:"timeoutSeconds,omitempty"`
								} `tfsdk:"client_ip" json:"clientIP,omitempty"`
							} `tfsdk:"session_affinity_config" json:"sessionAffinityConfig,omitempty"`
							TrafficDistribution *string `tfsdk:"traffic_distribution" json:"trafficDistribution,omitempty"`
							Type                *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"service_template" json:"serviceTemplate,omitempty"`
					UpdateStrategy *string `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
				} `tfsdk:"additional" json:"additional,omitempty"`
				DisabledDefaultServices *[]string `tfsdk:"disabled_default_services" json:"disabledDefaultServices,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
		} `tfsdk:"managed" json:"managed,omitempty"`
		MaxSyncReplicas *int64 `tfsdk:"max_sync_replicas" json:"maxSyncReplicas,omitempty"`
		MinSyncReplicas *int64 `tfsdk:"min_sync_replicas" json:"minSyncReplicas,omitempty"`
		Monitoring      *struct {
			CustomQueriesConfigMap *[]struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"custom_queries_config_map" json:"customQueriesConfigMap,omitempty"`
			CustomQueriesSecret *[]struct {
				Key  *string `tfsdk:"key" json:"key,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"custom_queries_secret" json:"customQueriesSecret,omitempty"`
			DisableDefaultQueries       *bool `tfsdk:"disable_default_queries" json:"disableDefaultQueries,omitempty"`
			EnablePodMonitor            *bool `tfsdk:"enable_pod_monitor" json:"enablePodMonitor,omitempty"`
			PodMonitorMetricRelabelings *[]struct {
				Action       *string   `tfsdk:"action" json:"action,omitempty"`
				Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
				Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
				Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
				Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
				SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
			} `tfsdk:"pod_monitor_metric_relabelings" json:"podMonitorMetricRelabelings,omitempty"`
			PodMonitorRelabelings *[]struct {
				Action       *string   `tfsdk:"action" json:"action,omitempty"`
				Modulus      *int64    `tfsdk:"modulus" json:"modulus,omitempty"`
				Regex        *string   `tfsdk:"regex" json:"regex,omitempty"`
				Replacement  *string   `tfsdk:"replacement" json:"replacement,omitempty"`
				Separator    *string   `tfsdk:"separator" json:"separator,omitempty"`
				SourceLabels *[]string `tfsdk:"source_labels" json:"sourceLabels,omitempty"`
				TargetLabel  *string   `tfsdk:"target_label" json:"targetLabel,omitempty"`
			} `tfsdk:"pod_monitor_relabelings" json:"podMonitorRelabelings,omitempty"`
			Tls *struct {
				Enabled *bool `tfsdk:"enabled" json:"enabled,omitempty"`
			} `tfsdk:"tls" json:"tls,omitempty"`
		} `tfsdk:"monitoring" json:"monitoring,omitempty"`
		NodeMaintenanceWindow *struct {
			InProgress *bool `tfsdk:"in_progress" json:"inProgress,omitempty"`
			ReusePVC   *bool `tfsdk:"reuse_pvc" json:"reusePVC,omitempty"`
		} `tfsdk:"node_maintenance_window" json:"nodeMaintenanceWindow,omitempty"`
		Plugins *[]struct {
			Enabled    *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			Name       *string            `tfsdk:"name" json:"name,omitempty"`
			Parameters *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
		} `tfsdk:"plugins" json:"plugins,omitempty"`
		PostgresGID *int64 `tfsdk:"postgres_gid" json:"postgresGID,omitempty"`
		PostgresUID *int64 `tfsdk:"postgres_uid" json:"postgresUID,omitempty"`
		Postgresql  *struct {
			EnableAlterSystem *bool `tfsdk:"enable_alter_system" json:"enableAlterSystem,omitempty"`
			Ldap              *struct {
				BindAsAuth *struct {
					Prefix *string `tfsdk:"prefix" json:"prefix,omitempty"`
					Suffix *string `tfsdk:"suffix" json:"suffix,omitempty"`
				} `tfsdk:"bind_as_auth" json:"bindAsAuth,omitempty"`
				BindSearchAuth *struct {
					BaseDN       *string `tfsdk:"base_dn" json:"baseDN,omitempty"`
					BindDN       *string `tfsdk:"bind_dn" json:"bindDN,omitempty"`
					BindPassword *struct {
						Key      *string `tfsdk:"key" json:"key,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"bind_password" json:"bindPassword,omitempty"`
					SearchAttribute *string `tfsdk:"search_attribute" json:"searchAttribute,omitempty"`
					SearchFilter    *string `tfsdk:"search_filter" json:"searchFilter,omitempty"`
				} `tfsdk:"bind_search_auth" json:"bindSearchAuth,omitempty"`
				Port   *int64  `tfsdk:"port" json:"port,omitempty"`
				Scheme *string `tfsdk:"scheme" json:"scheme,omitempty"`
				Server *string `tfsdk:"server" json:"server,omitempty"`
				Tls    *bool   `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"ldap" json:"ldap,omitempty"`
			Parameters                    *map[string]string `tfsdk:"parameters" json:"parameters,omitempty"`
			Pg_hba                        *[]string          `tfsdk:"pg_hba" json:"pg_hba,omitempty"`
			Pg_ident                      *[]string          `tfsdk:"pg_ident" json:"pg_ident,omitempty"`
			PromotionTimeout              *int64             `tfsdk:"promotion_timeout" json:"promotionTimeout,omitempty"`
			Shared_preload_libraries      *[]string          `tfsdk:"shared_preload_libraries" json:"shared_preload_libraries,omitempty"`
			SyncReplicaElectionConstraint *struct {
				Enabled                *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				NodeLabelsAntiAffinity *[]string `tfsdk:"node_labels_anti_affinity" json:"nodeLabelsAntiAffinity,omitempty"`
			} `tfsdk:"sync_replica_election_constraint" json:"syncReplicaElectionConstraint,omitempty"`
			Synchronous *struct {
				MaxStandbyNamesFromCluster *int64    `tfsdk:"max_standby_names_from_cluster" json:"maxStandbyNamesFromCluster,omitempty"`
				Method                     *string   `tfsdk:"method" json:"method,omitempty"`
				Number                     *int64    `tfsdk:"number" json:"number,omitempty"`
				StandbyNamesPost           *[]string `tfsdk:"standby_names_post" json:"standbyNamesPost,omitempty"`
				StandbyNamesPre            *[]string `tfsdk:"standby_names_pre" json:"standbyNamesPre,omitempty"`
			} `tfsdk:"synchronous" json:"synchronous,omitempty"`
		} `tfsdk:"postgresql" json:"postgresql,omitempty"`
		PrimaryUpdateMethod     *string `tfsdk:"primary_update_method" json:"primaryUpdateMethod,omitempty"`
		PrimaryUpdateStrategy   *string `tfsdk:"primary_update_strategy" json:"primaryUpdateStrategy,omitempty"`
		PriorityClassName       *string `tfsdk:"priority_class_name" json:"priorityClassName,omitempty"`
		ProjectedVolumeTemplate *struct {
			DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
			Sources     *[]struct {
				ClusterTrustBundle *struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
					Path       *string `tfsdk:"path" json:"path,omitempty"`
					SignerName *string `tfsdk:"signer_name" json:"signerName,omitempty"`
				} `tfsdk:"cluster_trust_bundle" json:"clusterTrustBundle,omitempty"`
				ConfigMap *struct {
					Items *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				DownwardAPI *struct {
					Items *[]struct {
						FieldRef *struct {
							ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
							FieldPath  *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						} `tfsdk:"field_ref" json:"fieldRef,omitempty"`
						Mode             *int64  `tfsdk:"mode" json:"mode,omitempty"`
						Path             *string `tfsdk:"path" json:"path,omitempty"`
						ResourceFieldRef *struct {
							ContainerName *string `tfsdk:"container_name" json:"containerName,omitempty"`
							Divisor       *string `tfsdk:"divisor" json:"divisor,omitempty"`
							Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
						} `tfsdk:"resource_field_ref" json:"resourceFieldRef,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
				} `tfsdk:"downward_api" json:"downwardAPI,omitempty"`
				Secret *struct {
					Items *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret" json:"secret,omitempty"`
				ServiceAccountToken *struct {
					Audience          *string `tfsdk:"audience" json:"audience,omitempty"`
					ExpirationSeconds *int64  `tfsdk:"expiration_seconds" json:"expirationSeconds,omitempty"`
					Path              *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"service_account_token" json:"serviceAccountToken,omitempty"`
			} `tfsdk:"sources" json:"sources,omitempty"`
		} `tfsdk:"projected_volume_template" json:"projectedVolumeTemplate,omitempty"`
		Replica *struct {
			Enabled        *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			MinApplyDelay  *string `tfsdk:"min_apply_delay" json:"minApplyDelay,omitempty"`
			Primary        *string `tfsdk:"primary" json:"primary,omitempty"`
			PromotionToken *string `tfsdk:"promotion_token" json:"promotionToken,omitempty"`
			Self           *string `tfsdk:"self" json:"self,omitempty"`
			Source         *string `tfsdk:"source" json:"source,omitempty"`
		} `tfsdk:"replica" json:"replica,omitempty"`
		ReplicationSlots *struct {
			HighAvailability *struct {
				Enabled    *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
				SlotPrefix *string `tfsdk:"slot_prefix" json:"slotPrefix,omitempty"`
			} `tfsdk:"high_availability" json:"highAvailability,omitempty"`
			SynchronizeReplicas *struct {
				Enabled         *bool     `tfsdk:"enabled" json:"enabled,omitempty"`
				ExcludePatterns *[]string `tfsdk:"exclude_patterns" json:"excludePatterns,omitempty"`
			} `tfsdk:"synchronize_replicas" json:"synchronizeReplicas,omitempty"`
			UpdateInterval *int64 `tfsdk:"update_interval" json:"updateInterval,omitempty"`
		} `tfsdk:"replication_slots" json:"replicationSlots,omitempty"`
		Resources *struct {
			Claims *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"claims" json:"claims,omitempty"`
			Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
			Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		SchedulerName  *string `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
		SeccompProfile *struct {
			LocalhostProfile *string `tfsdk:"localhost_profile" json:"localhostProfile,omitempty"`
			Type             *string `tfsdk:"type" json:"type,omitempty"`
		} `tfsdk:"seccomp_profile" json:"seccompProfile,omitempty"`
		ServiceAccountTemplate *struct {
			Metadata *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"metadata" json:"metadata,omitempty"`
		} `tfsdk:"service_account_template" json:"serviceAccountTemplate,omitempty"`
		SmartShutdownTimeout *int64 `tfsdk:"smart_shutdown_timeout" json:"smartShutdownTimeout,omitempty"`
		StartDelay           *int64 `tfsdk:"start_delay" json:"startDelay,omitempty"`
		StopDelay            *int64 `tfsdk:"stop_delay" json:"stopDelay,omitempty"`
		Storage              *struct {
			PvcTemplate *struct {
				AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
				DataSource  *struct {
					ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"data_source" json:"dataSource,omitempty"`
				DataSourceRef *struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
				Resources *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				StorageClassName          *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				VolumeAttributesClassName *string `tfsdk:"volume_attributes_class_name" json:"volumeAttributesClassName,omitempty"`
				VolumeMode                *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				VolumeName                *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"pvc_template" json:"pvcTemplate,omitempty"`
			ResizeInUseVolumes *bool   `tfsdk:"resize_in_use_volumes" json:"resizeInUseVolumes,omitempty"`
			Size               *string `tfsdk:"size" json:"size,omitempty"`
			StorageClass       *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		SuperuserSecret *struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
		} `tfsdk:"superuser_secret" json:"superuserSecret,omitempty"`
		SwitchoverDelay *int64 `tfsdk:"switchover_delay" json:"switchoverDelay,omitempty"`
		Tablespaces     *[]struct {
			Name  *string `tfsdk:"name" json:"name,omitempty"`
			Owner *struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"owner" json:"owner,omitempty"`
			Storage *struct {
				PvcTemplate *struct {
					AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
					DataSource  *struct {
						ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"data_source" json:"dataSource,omitempty"`
					DataSourceRef *struct {
						ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
						Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
					Resources *struct {
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Selector *struct {
						MatchExpressions *[]struct {
							Key      *string   `tfsdk:"key" json:"key,omitempty"`
							Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
							Values   *[]string `tfsdk:"values" json:"values,omitempty"`
						} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
						MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
					} `tfsdk:"selector" json:"selector,omitempty"`
					StorageClassName          *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
					VolumeAttributesClassName *string `tfsdk:"volume_attributes_class_name" json:"volumeAttributesClassName,omitempty"`
					VolumeMode                *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					VolumeName                *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
				} `tfsdk:"pvc_template" json:"pvcTemplate,omitempty"`
				ResizeInUseVolumes *bool   `tfsdk:"resize_in_use_volumes" json:"resizeInUseVolumes,omitempty"`
				Size               *string `tfsdk:"size" json:"size,omitempty"`
				StorageClass       *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
			} `tfsdk:"storage" json:"storage,omitempty"`
			Temporary *bool `tfsdk:"temporary" json:"temporary,omitempty"`
		} `tfsdk:"tablespaces" json:"tablespaces,omitempty"`
		TopologySpreadConstraints *[]struct {
			LabelSelector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"label_selector" json:"labelSelector,omitempty"`
			MatchLabelKeys     *[]string `tfsdk:"match_label_keys" json:"matchLabelKeys,omitempty"`
			MaxSkew            *int64    `tfsdk:"max_skew" json:"maxSkew,omitempty"`
			MinDomains         *int64    `tfsdk:"min_domains" json:"minDomains,omitempty"`
			NodeAffinityPolicy *string   `tfsdk:"node_affinity_policy" json:"nodeAffinityPolicy,omitempty"`
			NodeTaintsPolicy   *string   `tfsdk:"node_taints_policy" json:"nodeTaintsPolicy,omitempty"`
			TopologyKey        *string   `tfsdk:"topology_key" json:"topologyKey,omitempty"`
			WhenUnsatisfiable  *string   `tfsdk:"when_unsatisfiable" json:"whenUnsatisfiable,omitempty"`
		} `tfsdk:"topology_spread_constraints" json:"topologySpreadConstraints,omitempty"`
		WalStorage *struct {
			PvcTemplate *struct {
				AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
				DataSource  *struct {
					ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"data_source" json:"dataSource,omitempty"`
				DataSourceRef *struct {
					ApiGroup  *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
					Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"data_source_ref" json:"dataSourceRef,omitempty"`
				Resources *struct {
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Selector *struct {
					MatchExpressions *[]struct {
						Key      *string   `tfsdk:"key" json:"key,omitempty"`
						Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
						Values   *[]string `tfsdk:"values" json:"values,omitempty"`
					} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
					MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
				} `tfsdk:"selector" json:"selector,omitempty"`
				StorageClassName          *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				VolumeAttributesClassName *string `tfsdk:"volume_attributes_class_name" json:"volumeAttributesClassName,omitempty"`
				VolumeMode                *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				VolumeName                *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
			} `tfsdk:"pvc_template" json:"pvcTemplate,omitempty"`
			ResizeInUseVolumes *bool   `tfsdk:"resize_in_use_volumes" json:"resizeInUseVolumes,omitempty"`
			Size               *string `tfsdk:"size" json:"size,omitempty"`
			StorageClass       *string `tfsdk:"storage_class" json:"storageClass,omitempty"`
		} `tfsdk:"wal_storage" json:"walStorage,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *PostgresqlCnpgIoClusterV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_postgresql_cnpg_io_cluster_v1_manifest"
}

func (r *PostgresqlCnpgIoClusterV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Cluster is the Schema for the PostgreSQL API",
		MarkdownDescription: "Cluster is the Schema for the PostgreSQL API",
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
				Description:         "Specification of the desired behavior of the cluster.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				MarkdownDescription: "Specification of the desired behavior of the cluster.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "Affinity/Anti-affinity rules for Pods",
						MarkdownDescription: "Affinity/Anti-affinity rules for Pods",
						Attributes: map[string]schema.Attribute{
							"additional_pod_affinity": schema.SingleNestedAttribute{
								Description:         "AdditionalPodAffinity allows to specify pod affinity terms to be passed to all the cluster's pods.",
								MarkdownDescription: "AdditionalPodAffinity allows to specify pod affinity terms to be passed to all the cluster's pods.",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Required. A pod affinity term, associated with the corresponding weight.",
													MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"match_label_keys": schema.ListAttribute{
															Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
													Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
										Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"match_label_keys": schema.ListAttribute{
													Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mismatch_label_keys": schema.ListAttribute{
													Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
													Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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

							"additional_pod_anti_affinity": schema.SingleNestedAttribute{
								Description:         "AdditionalPodAntiAffinity allows to specify pod anti-affinity terms to be added to the ones generatedby the operator if EnablePodAntiAffinity is set to true (default) or to be used exclusively if set to false.",
								MarkdownDescription: "AdditionalPodAntiAffinity allows to specify pod anti-affinity terms to be added to the ones generatedby the operator if EnablePodAntiAffinity is set to true (default) or to be used exclusively if set to false.",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe anti-affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling anti-affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; thenode(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"pod_affinity_term": schema.SingleNestedAttribute{
													Description:         "Required. A pod affinity term, associated with the corresponding weight.",
													MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

														"match_label_keys": schema.ListAttribute{
															Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"mismatch_label_keys": schema.ListAttribute{
															Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"namespace_selector": schema.SingleNestedAttribute{
															Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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
													Description:         "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
													MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm,in the range 1-100.",
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
										Description:         "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
										MarkdownDescription: "If the anti-affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the anti-affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to a pod label update), thesystem may or may not try to eventually evict the pod from its node.When there are multiple elements, the lists of nodes corresponding to eachpodAffinityTerm are intersected, i.e. all terms must be satisfied.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
													MarkdownDescription: "A label query over a set of resources, in this case pods.If it's null, this PodAffinityTerm matches with no Pods.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"match_label_keys": schema.ListAttribute{
													Description:         "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key in (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both matchLabelKeys and labelSelector.Also, matchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"mismatch_label_keys": schema.ListAttribute{
													Description:         "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													MarkdownDescription: "MismatchLabelKeys is a set of pod label keys to select which pods willbe taken into consideration. The keys are used to lookup values from theincoming pod labels, those key-value labels are merged with 'labelSelector' as 'key notin (value)'to select the group of existing pods which pods will be taken into considerationfor the incoming pod's pod (anti) affinity. Keys that don't exist in the incomingpod labels will be ignored. The default value is empty.The same key is forbidden to exist in both mismatchLabelKeys and labelSelector.Also, mismatchLabelKeys cannot be set when labelSelector isn't set.This is an alpha field and requires enabling MatchLabelKeysInPodAffinity feature gate.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace_selector": schema.SingleNestedAttribute{
													Description:         "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
													MarkdownDescription: "A label query over the set of namespaces that the term applies to.The term is applied to the union of the namespaces selected by this fieldand the ones listed in the namespaces field.null selector and null or empty namespaces list means 'this pod's namespace'.An empty selector ({}) matches all namespaces.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
													Description:         "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to.The term is applied to the union of the namespaces listed in this fieldand the ones selected by namespaceSelector.null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"topology_key": schema.StringAttribute{
													Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
													MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matchingthe labelSelector in the specified namespaces, where co-located is defined as running on a nodewhose value of the label with key topologyKey matches that of any node on which any of theselected pods is running.Empty topologyKey is not allowed.",
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

							"enable_pod_anti_affinity": schema.BoolAttribute{
								Description:         "Activates anti-affinity for the pods. The operator will define podsanti-affinity unless this field is explicitly set to false",
								MarkdownDescription: "Activates anti-affinity for the pods. The operator will define podsanti-affinity unless this field is explicitly set to false",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_affinity": schema.SingleNestedAttribute{
								Description:         "NodeAffinity describes node affinity scheduling rules for the pod.More info: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity",
								MarkdownDescription: "NodeAffinity describes node affinity scheduling rules for the pod.More info: https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#node-affinity",
								Attributes: map[string]schema.Attribute{
									"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
										Description:         "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
										MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfythe affinity expressions specified by this field, but it may choosea node that violates one or more of the expressions. The node that ismost preferred is the one with the greatest sum of weights, i.e.for each node that meets all of the scheduling requirements (resourcerequest, requiredDuringScheduling affinity expressions, etc.),compute a sum by iterating through the elements of this field and adding'weight' to the sum if the node matches the corresponding matchExpressions; thenode(s) with the highest sum are the most preferred.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"preference": schema.SingleNestedAttribute{
													Description:         "A node selector term, associated with the corresponding weight.",
													MarkdownDescription: "A node selector term, associated with the corresponding weight.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's labels.",
															MarkdownDescription: "A list of node selector requirements by node's labels.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
													Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
													MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
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
										Description:         "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
										MarkdownDescription: "If the affinity requirements specified by this field are not met atscheduling time, the pod will not be scheduled onto the node.If the affinity requirements specified by this field cease to be metat some point during pod execution (e.g. due to an update), the systemmay or may not try to eventually evict the pod from its node.",
										Attributes: map[string]schema.Attribute{
											"node_selector_terms": schema.ListNestedAttribute{
												Description:         "Required. A list of node selector terms. The terms are ORed.",
												MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "A list of node selector requirements by node's labels.",
															MarkdownDescription: "A list of node selector requirements by node's labels.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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
															Description:         "A list of node selector requirements by node's fields.",
															MarkdownDescription: "A list of node selector requirements by node's fields.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The label key that the selector applies to.",
																		MarkdownDescription: "The label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		MarkdownDescription: "Represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
																		MarkdownDescription: "An array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. If the operator is Gt or Lt, the valuesarray must have a single element, which will be interpreted as an integer.This array is replaced during a strategic merge patch.",
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

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is map of key-value pairs used to define the nodes on whichthe pods can run.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
								MarkdownDescription: "NodeSelector is map of key-value pairs used to define the nodes on whichthe pods can run.More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_anti_affinity_type": schema.StringAttribute{
								Description:         "PodAntiAffinityType allows the user to decide whether pod anti-affinity between cluster instance has to beconsidered a strong requirement during scheduling or not. Allowed values are: 'preferred' (default if empty) or'required'. Setting it to 'required', could lead to instances remaining pending until new kubernetes nodes areadded if all the existing nodes don't match the required pod anti-affinity rule.More info:https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity",
								MarkdownDescription: "PodAntiAffinityType allows the user to decide whether pod anti-affinity between cluster instance has to beconsidered a strong requirement during scheduling or not. Allowed values are: 'preferred' (default if empty) or'required'. Setting it to 'required', could lead to instances remaining pending until new kubernetes nodes areadded if all the existing nodes don't match the required pod anti-affinity rule.More info:https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#inter-pod-affinity-and-anti-affinity",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Tolerations is a list of Tolerations that should be set for all the pods, in order to allow them to runon tainted nodes.More info: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/",
								MarkdownDescription: "Tolerations is a list of Tolerations that should be set for all the pods, in order to allow them to runon tainted nodes.More info: https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects.When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys.If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value.Valid operators are Exists and Equal. Defaults to Equal.Exists is equivalent to wildcard for value, so that a pod cantolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must beof effect NoExecute, otherwise this field is ignored) tolerates the taint. By default,it is not set, which means tolerate the taint forever (do not evict). Zero andnegative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to.If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

							"topology_key": schema.StringAttribute{
								Description:         "TopologyKey to use for anti-affinity configuration. See k8s documentationfor more info on that",
								MarkdownDescription: "TopologyKey to use for anti-affinity configuration. See k8s documentationfor more info on that",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"backup": schema.SingleNestedAttribute{
						Description:         "The configuration to be used for backups",
						MarkdownDescription: "The configuration to be used for backups",
						Attributes: map[string]schema.Attribute{
							"barman_object_store": schema.SingleNestedAttribute{
								Description:         "The configuration for the barman-cloud tool suite",
								MarkdownDescription: "The configuration for the barman-cloud tool suite",
								Attributes: map[string]schema.Attribute{
									"azure_credentials": schema.SingleNestedAttribute{
										Description:         "The credentials to use to upload data to Azure Blob Storage",
										MarkdownDescription: "The credentials to use to upload data to Azure Blob Storage",
										Attributes: map[string]schema.Attribute{
											"connection_string": schema.SingleNestedAttribute{
												Description:         "The connection string to be used",
												MarkdownDescription: "The connection string to be used",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"inherit_from_azure_ad": schema.BoolAttribute{
												Description:         "Use the Azure AD based authentication without providing explicitly the keys.",
												MarkdownDescription: "Use the Azure AD based authentication without providing explicitly the keys.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"storage_account": schema.SingleNestedAttribute{
												Description:         "The storage account where to upload data",
												MarkdownDescription: "The storage account where to upload data",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_key": schema.SingleNestedAttribute{
												Description:         "The storage account key to be used in conjunctionwith the storage account name",
												MarkdownDescription: "The storage account key to be used in conjunctionwith the storage account name",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_sas_token": schema.SingleNestedAttribute{
												Description:         "A shared-access-signature to be used in conjunction withthe storage account name",
												MarkdownDescription: "A shared-access-signature to be used in conjunction withthe storage account name",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
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

									"data": schema.SingleNestedAttribute{
										Description:         "The configuration to be used to backup the data filesWhen not defined, base backups files will be stored uncompressed and maybe unencrypted in the object store, according to the bucket defaultpolicy.",
										MarkdownDescription: "The configuration to be used to backup the data filesWhen not defined, base backups files will be stored uncompressed and maybe unencrypted in the object store, according to the bucket defaultpolicy.",
										Attributes: map[string]schema.Attribute{
											"additional_command_args": schema.ListAttribute{
												Description:         "AdditionalCommandArgs represents additional arguments that can be appendedto the 'barman-cloud-backup' command-line invocation. These argumentsprovide flexibility to customize the backup process further according tospecific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-backup' command, to avoid potential errors or unintendedbehavior during execution.",
												MarkdownDescription: "AdditionalCommandArgs represents additional arguments that can be appendedto the 'barman-cloud-backup' command-line invocation. These argumentsprovide flexibility to customize the backup process further according tospecific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-backup' command, to avoid potential errors or unintendedbehavior during execution.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"compression": schema.StringAttribute{
												Description:         "Compress a backup file (a tar file per tablespace) while streaming itto the object store. Available options are empty string (nocompression, default), 'gzip', 'bzip2' or 'snappy'.",
												MarkdownDescription: "Compress a backup file (a tar file per tablespace) while streaming itto the object store. Available options are empty string (nocompression, default), 'gzip', 'bzip2' or 'snappy'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("gzip", "bzip2", "snappy"),
												},
											},

											"encryption": schema.StringAttribute{
												Description:         "Whenever to force the encryption of files (if the bucket isnot already configured for that).Allowed options are empty string (use the bucket policy, default),'AES256' and 'aws:kms'",
												MarkdownDescription: "Whenever to force the encryption of files (if the bucket isnot already configured for that).Allowed options are empty string (use the bucket policy, default),'AES256' and 'aws:kms'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("AES256", "aws:kms"),
												},
											},

											"immediate_checkpoint": schema.BoolAttribute{
												Description:         "Control whether the I/O workload for the backup initial checkpoint willbe limited, according to the 'checkpoint_completion_target' setting onthe PostgreSQL server. If set to true, an immediate checkpoint will beused, meaning PostgreSQL will complete the checkpoint as soon aspossible. 'false' by default.",
												MarkdownDescription: "Control whether the I/O workload for the backup initial checkpoint willbe limited, according to the 'checkpoint_completion_target' setting onthe PostgreSQL server. If set to true, an immediate checkpoint will beused, meaning PostgreSQL will complete the checkpoint as soon aspossible. 'false' by default.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"jobs": schema.Int64Attribute{
												Description:         "The number of parallel jobs to be used to upload the backup, defaultsto 2",
												MarkdownDescription: "The number of parallel jobs to be used to upload the backup, defaultsto 2",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"destination_path": schema.StringAttribute{
										Description:         "The path where to store the backup (i.e. s3://bucket/path/to/folder)this path, with different destination folders, will be used for WALsand for data",
										MarkdownDescription: "The path where to store the backup (i.e. s3://bucket/path/to/folder)this path, with different destination folders, will be used for WALsand for data",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"endpoint_ca": schema.SingleNestedAttribute{
										Description:         "EndpointCA store the CA bundle of the barman endpoint.Useful when using self-signed certificates to avoiderrors with certificate issuer and barman-cloud-wal-archive",
										MarkdownDescription: "EndpointCA store the CA bundle of the barman endpoint.Useful when using self-signed certificates to avoiderrors with certificate issuer and barman-cloud-wal-archive",
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Description:         "The key to select",
												MarkdownDescription: "The key to select",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"endpoint_url": schema.StringAttribute{
										Description:         "Endpoint to be used to upload data to the cloud,overriding the automatic endpoint discovery",
										MarkdownDescription: "Endpoint to be used to upload data to the cloud,overriding the automatic endpoint discovery",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"google_credentials": schema.SingleNestedAttribute{
										Description:         "The credentials to use to upload data to Google Cloud Storage",
										MarkdownDescription: "The credentials to use to upload data to Google Cloud Storage",
										Attributes: map[string]schema.Attribute{
											"application_credentials": schema.SingleNestedAttribute{
												Description:         "The secret containing the Google Cloud Storage JSON file with the credentials",
												MarkdownDescription: "The secret containing the Google Cloud Storage JSON file with the credentials",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"gke_environment": schema.BoolAttribute{
												Description:         "If set to true, will presume that it's running inside a GKE environment,default to false.",
												MarkdownDescription: "If set to true, will presume that it's running inside a GKE environment,default to false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"history_tags": schema.MapAttribute{
										Description:         "HistoryTags is a list of key value pairs that will be passed to theBarman --history-tags option.",
										MarkdownDescription: "HistoryTags is a list of key value pairs that will be passed to theBarman --history-tags option.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"s3_credentials": schema.SingleNestedAttribute{
										Description:         "The credentials to use to upload data to S3",
										MarkdownDescription: "The credentials to use to upload data to S3",
										Attributes: map[string]schema.Attribute{
											"access_key_id": schema.SingleNestedAttribute{
												Description:         "The reference to the access key id",
												MarkdownDescription: "The reference to the access key id",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"inherit_from_iam_role": schema.BoolAttribute{
												Description:         "Use the role based authentication without providing explicitly the keys.",
												MarkdownDescription: "Use the role based authentication without providing explicitly the keys.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"region": schema.SingleNestedAttribute{
												Description:         "The reference to the secret containing the region name",
												MarkdownDescription: "The reference to the secret containing the region name",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"secret_access_key": schema.SingleNestedAttribute{
												Description:         "The reference to the secret access key",
												MarkdownDescription: "The reference to the secret access key",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"session_token": schema.SingleNestedAttribute{
												Description:         "The references to the session key",
												MarkdownDescription: "The references to the session key",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
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

									"server_name": schema.StringAttribute{
										Description:         "The server name on S3, the cluster name is used if thisparameter is omitted",
										MarkdownDescription: "The server name on S3, the cluster name is used if thisparameter is omitted",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tags": schema.MapAttribute{
										Description:         "Tags is a list of key value pairs that will be passed to theBarman --tags option.",
										MarkdownDescription: "Tags is a list of key value pairs that will be passed to theBarman --tags option.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"wal": schema.SingleNestedAttribute{
										Description:         "The configuration for the backup of the WAL stream.When not defined, WAL files will be stored uncompressed and may beunencrypted in the object store, according to the bucket default policy.",
										MarkdownDescription: "The configuration for the backup of the WAL stream.When not defined, WAL files will be stored uncompressed and may beunencrypted in the object store, according to the bucket default policy.",
										Attributes: map[string]schema.Attribute{
											"archive_additional_command_args": schema.ListAttribute{
												Description:         "Additional arguments that can be appended to the 'barman-cloud-wal-archive'command-line invocation. These arguments provide flexibility to customizethe WAL archive process further, according to specific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-wal-archive' command, to avoid potential errors or unintendedbehavior during execution.",
												MarkdownDescription: "Additional arguments that can be appended to the 'barman-cloud-wal-archive'command-line invocation. These arguments provide flexibility to customizethe WAL archive process further, according to specific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-wal-archive' command, to avoid potential errors or unintendedbehavior during execution.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"compression": schema.StringAttribute{
												Description:         "Compress a WAL file before sending it to the object store. Availableoptions are empty string (no compression, default), 'gzip', 'bzip2' or 'snappy'.",
												MarkdownDescription: "Compress a WAL file before sending it to the object store. Availableoptions are empty string (no compression, default), 'gzip', 'bzip2' or 'snappy'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("gzip", "bzip2", "snappy"),
												},
											},

											"encryption": schema.StringAttribute{
												Description:         "Whenever to force the encryption of files (if the bucket isnot already configured for that).Allowed options are empty string (use the bucket policy, default),'AES256' and 'aws:kms'",
												MarkdownDescription: "Whenever to force the encryption of files (if the bucket isnot already configured for that).Allowed options are empty string (use the bucket policy, default),'AES256' and 'aws:kms'",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("AES256", "aws:kms"),
												},
											},

											"max_parallel": schema.Int64Attribute{
												Description:         "Number of WAL files to be either archived in parallel (when thePostgreSQL instance is archiving to a backup object store) orrestored in parallel (when a PostgreSQL standby is fetching WALfiles from a recovery object store). If not specified, WAL fileswill be processed one at a time. It accepts a positive integer as avalue - with 1 being the minimum accepted value.",
												MarkdownDescription: "Number of WAL files to be either archived in parallel (when thePostgreSQL instance is archiving to a backup object store) orrestored in parallel (when a PostgreSQL standby is fetching WALfiles from a recovery object store). If not specified, WAL fileswill be processed one at a time. It accepts a positive integer as avalue - with 1 being the minimum accepted value.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"restore_additional_command_args": schema.ListAttribute{
												Description:         "Additional arguments that can be appended to the 'barman-cloud-wal-restore'command-line invocation. These arguments provide flexibility to customizethe WAL restore process further, according to specific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-wal-restore' command, to avoid potential errors or unintendedbehavior during execution.",
												MarkdownDescription: "Additional arguments that can be appended to the 'barman-cloud-wal-restore'command-line invocation. These arguments provide flexibility to customizethe WAL restore process further, according to specific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-wal-restore' command, to avoid potential errors or unintendedbehavior during execution.",
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

							"retention_policy": schema.StringAttribute{
								Description:         "RetentionPolicy is the retention policy to be used for backupsand WALs (i.e. '60d'). The retention policy is expressed in the formof 'XXu' where 'XX' is a positive integer and 'u' is in '[dwm]' -days, weeks, months.It's currently only applicable when using the BarmanObjectStore method.",
								MarkdownDescription: "RetentionPolicy is the retention policy to be used for backupsand WALs (i.e. '60d'). The retention policy is expressed in the formof 'XXu' where 'XX' is a positive integer and 'u' is in '[dwm]' -days, weeks, months.It's currently only applicable when using the BarmanObjectStore method.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^[1-9][0-9]*[dwm]$`), ""),
								},
							},

							"target": schema.StringAttribute{
								Description:         "The policy to decide which instance should perform backups. Availableoptions are empty string, which will default to 'prefer-standby' policy,'primary' to have backups run always on primary instances, 'prefer-standby'to have backups run preferably on the most updated standby, if available.",
								MarkdownDescription: "The policy to decide which instance should perform backups. Availableoptions are empty string, which will default to 'prefer-standby' policy,'primary' to have backups run always on primary instances, 'prefer-standby'to have backups run preferably on the most updated standby, if available.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("primary", "prefer-standby"),
								},
							},

							"volume_snapshot": schema.SingleNestedAttribute{
								Description:         "VolumeSnapshot provides the configuration for the execution of volume snapshot backups.",
								MarkdownDescription: "VolumeSnapshot provides the configuration for the execution of volume snapshot backups.",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations key-value pairs that will be added to .metadata.annotations snapshot resources.",
										MarkdownDescription: "Annotations key-value pairs that will be added to .metadata.annotations snapshot resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"class_name": schema.StringAttribute{
										Description:         "ClassName specifies the Snapshot Class to be used for PG_DATA PersistentVolumeClaim.It is the default class for the other types if no specific class is present",
										MarkdownDescription: "ClassName specifies the Snapshot Class to be used for PG_DATA PersistentVolumeClaim.It is the default class for the other types if no specific class is present",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels are key-value pairs that will be added to .metadata.labels snapshot resources.",
										MarkdownDescription: "Labels are key-value pairs that will be added to .metadata.labels snapshot resources.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"online": schema.BoolAttribute{
										Description:         "Whether the default type of backup with volume snapshots isonline/hot ('true', default) or offline/cold ('false')",
										MarkdownDescription: "Whether the default type of backup with volume snapshots isonline/hot ('true', default) or offline/cold ('false')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"online_configuration": schema.SingleNestedAttribute{
										Description:         "Configuration parameters to control the online/hot backup with volume snapshots",
										MarkdownDescription: "Configuration parameters to control the online/hot backup with volume snapshots",
										Attributes: map[string]schema.Attribute{
											"immediate_checkpoint": schema.BoolAttribute{
												Description:         "Control whether the I/O workload for the backup initial checkpoint willbe limited, according to the 'checkpoint_completion_target' setting onthe PostgreSQL server. If set to true, an immediate checkpoint will beused, meaning PostgreSQL will complete the checkpoint as soon aspossible. 'false' by default.",
												MarkdownDescription: "Control whether the I/O workload for the backup initial checkpoint willbe limited, according to the 'checkpoint_completion_target' setting onthe PostgreSQL server. If set to true, an immediate checkpoint will beused, meaning PostgreSQL will complete the checkpoint as soon aspossible. 'false' by default.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"wait_for_archive": schema.BoolAttribute{
												Description:         "If false, the function will return immediately after the backup is completed,without waiting for WAL to be archived.This behavior is only useful with backup software that independently monitors WAL archiving.Otherwise, WAL required to make the backup consistent might be missing and make the backup useless.By default, or when this parameter is true, pg_backup_stop will wait for WAL to be archived when archiving isenabled.On a standby, this means that it will wait only when archive_mode = always.If write activity on the primary is low, it may be useful to run pg_switch_wal on the primary in order to triggeran immediate segment switch.",
												MarkdownDescription: "If false, the function will return immediately after the backup is completed,without waiting for WAL to be archived.This behavior is only useful with backup software that independently monitors WAL archiving.Otherwise, WAL required to make the backup consistent might be missing and make the backup useless.By default, or when this parameter is true, pg_backup_stop will wait for WAL to be archived when archiving isenabled.On a standby, this means that it will wait only when archive_mode = always.If write activity on the primary is low, it may be useful to run pg_switch_wal on the primary in order to triggeran immediate segment switch.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"snapshot_owner_reference": schema.StringAttribute{
										Description:         "SnapshotOwnerReference indicates the type of owner reference the snapshot should have",
										MarkdownDescription: "SnapshotOwnerReference indicates the type of owner reference the snapshot should have",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("none", "cluster", "backup"),
										},
									},

									"tablespace_class_name": schema.MapAttribute{
										Description:         "TablespaceClassName specifies the Snapshot Class to be used for the tablespaces.defaults to the PGDATA Snapshot Class, if set",
										MarkdownDescription: "TablespaceClassName specifies the Snapshot Class to be used for the tablespaces.defaults to the PGDATA Snapshot Class, if set",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"wal_class_name": schema.StringAttribute{
										Description:         "WalClassName specifies the Snapshot Class to be used for the PG_WAL PersistentVolumeClaim.",
										MarkdownDescription: "WalClassName specifies the Snapshot Class to be used for the PG_WAL PersistentVolumeClaim.",
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

					"bootstrap": schema.SingleNestedAttribute{
						Description:         "Instructions to bootstrap this cluster",
						MarkdownDescription: "Instructions to bootstrap this cluster",
						Attributes: map[string]schema.Attribute{
							"initdb": schema.SingleNestedAttribute{
								Description:         "Bootstrap the cluster via initdb",
								MarkdownDescription: "Bootstrap the cluster via initdb",
								Attributes: map[string]schema.Attribute{
									"data_checksums": schema.BoolAttribute{
										Description:         "Whether the '-k' option should be passed to initdb,enabling checksums on data pages (default: 'false')",
										MarkdownDescription: "Whether the '-k' option should be passed to initdb,enabling checksums on data pages (default: 'false')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"database": schema.StringAttribute{
										Description:         "Name of the database used by the application. Default: 'app'.",
										MarkdownDescription: "Name of the database used by the application. Default: 'app'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"encoding": schema.StringAttribute{
										Description:         "The value to be passed as option '--encoding' for initdb (default:'UTF8')",
										MarkdownDescription: "The value to be passed as option '--encoding' for initdb (default:'UTF8')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"import": schema.SingleNestedAttribute{
										Description:         "Bootstraps the new cluster by importing data from an existing PostgreSQLinstance using logical backup ('pg_dump' and 'pg_restore')",
										MarkdownDescription: "Bootstraps the new cluster by importing data from an existing PostgreSQLinstance using logical backup ('pg_dump' and 'pg_restore')",
										Attributes: map[string]schema.Attribute{
											"databases": schema.ListAttribute{
												Description:         "The databases to import",
												MarkdownDescription: "The databases to import",
												ElementType:         types.StringType,
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"post_import_application_sql": schema.ListAttribute{
												Description:         "List of SQL queries to be executed as a superuser in the applicationdatabase right after is imported - to be used with extreme care(by default empty). Only available in microservice type.",
												MarkdownDescription: "List of SQL queries to be executed as a superuser in the applicationdatabase right after is imported - to be used with extreme care(by default empty). Only available in microservice type.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"roles": schema.ListAttribute{
												Description:         "The roles to import",
												MarkdownDescription: "The roles to import",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"schema_only": schema.BoolAttribute{
												Description:         "When set to true, only the 'pre-data' and 'post-data' sections of'pg_restore' are invoked, avoiding data import. Default: 'false'.",
												MarkdownDescription: "When set to true, only the 'pre-data' and 'post-data' sections of'pg_restore' are invoked, avoiding data import. Default: 'false'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"source": schema.SingleNestedAttribute{
												Description:         "The source of the import",
												MarkdownDescription: "The source of the import",
												Attributes: map[string]schema.Attribute{
													"external_cluster": schema.StringAttribute{
														Description:         "The name of the externalCluster used for import",
														MarkdownDescription: "The name of the externalCluster used for import",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "The import type. Can be 'microservice' or 'monolith'.",
												MarkdownDescription: "The import type. Can be 'microservice' or 'monolith'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("microservice", "monolith"),
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"locale_c_type": schema.StringAttribute{
										Description:         "The value to be passed as option '--lc-ctype' for initdb (default:'C')",
										MarkdownDescription: "The value to be passed as option '--lc-ctype' for initdb (default:'C')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"locale_collate": schema.StringAttribute{
										Description:         "The value to be passed as option '--lc-collate' for initdb (default:'C')",
										MarkdownDescription: "The value to be passed as option '--lc-collate' for initdb (default:'C')",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"options": schema.ListAttribute{
										Description:         "The list of options that must be passed to initdb when creating the cluster.Deprecated: This could lead to inconsistent configurations,please use the explicit provided parameters instead.If defined, explicit values will be ignored.",
										MarkdownDescription: "The list of options that must be passed to initdb when creating the cluster.Deprecated: This could lead to inconsistent configurations,please use the explicit provided parameters instead.If defined, explicit values will be ignored.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"owner": schema.StringAttribute{
										Description:         "Name of the owner of the database in the instance to be usedby applications. Defaults to the value of the 'database' key.",
										MarkdownDescription: "Name of the owner of the database in the instance to be usedby applications. Defaults to the value of the 'database' key.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"post_init_application_sql": schema.ListAttribute{
										Description:         "List of SQL queries to be executed as a superuser in the applicationdatabase right after the cluster has been created - to be used with extreme care(by default empty)",
										MarkdownDescription: "List of SQL queries to be executed as a superuser in the applicationdatabase right after the cluster has been created - to be used with extreme care(by default empty)",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"post_init_application_sql_refs": schema.SingleNestedAttribute{
										Description:         "List of references to ConfigMaps or Secrets containing SQL filesto be executed as a superuser in the application database right afterthe cluster has been created. The references are processed in a specific order:first, all Secrets are processed, followed by all ConfigMaps.Within each group, the processing order follows the sequence specifiedin their respective arrays.(by default empty)",
										MarkdownDescription: "List of references to ConfigMaps or Secrets containing SQL filesto be executed as a superuser in the application database right afterthe cluster has been created. The references are processed in a specific order:first, all Secrets are processed, followed by all ConfigMaps.Within each group, the processing order follows the sequence specifiedin their respective arrays.(by default empty)",
										Attributes: map[string]schema.Attribute{
											"config_map_refs": schema.ListNestedAttribute{
												Description:         "ConfigMapRefs holds a list of references to ConfigMaps",
												MarkdownDescription: "ConfigMapRefs holds a list of references to ConfigMaps",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
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

											"secret_refs": schema.ListNestedAttribute{
												Description:         "SecretRefs holds a list of references to Secrets",
												MarkdownDescription: "SecretRefs holds a list of references to Secrets",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
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

									"post_init_sql": schema.ListAttribute{
										Description:         "List of SQL queries to be executed as a superuser in the 'postgres'database right after the cluster has been created - to be used with extreme care(by default empty)",
										MarkdownDescription: "List of SQL queries to be executed as a superuser in the 'postgres'database right after the cluster has been created - to be used with extreme care(by default empty)",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"post_init_sql_refs": schema.SingleNestedAttribute{
										Description:         "List of references to ConfigMaps or Secrets containing SQL filesto be executed as a superuser in the 'postgres' database right afterthe cluster has been created. The references are processed in a specific order:first, all Secrets are processed, followed by all ConfigMaps.Within each group, the processing order follows the sequence specifiedin their respective arrays.(by default empty)",
										MarkdownDescription: "List of references to ConfigMaps or Secrets containing SQL filesto be executed as a superuser in the 'postgres' database right afterthe cluster has been created. The references are processed in a specific order:first, all Secrets are processed, followed by all ConfigMaps.Within each group, the processing order follows the sequence specifiedin their respective arrays.(by default empty)",
										Attributes: map[string]schema.Attribute{
											"config_map_refs": schema.ListNestedAttribute{
												Description:         "ConfigMapRefs holds a list of references to ConfigMaps",
												MarkdownDescription: "ConfigMapRefs holds a list of references to ConfigMaps",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
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

											"secret_refs": schema.ListNestedAttribute{
												Description:         "SecretRefs holds a list of references to Secrets",
												MarkdownDescription: "SecretRefs holds a list of references to Secrets",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
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

									"post_init_template_sql": schema.ListAttribute{
										Description:         "List of SQL queries to be executed as a superuser in the 'template1'database right after the cluster has been created - to be used with extreme care(by default empty)",
										MarkdownDescription: "List of SQL queries to be executed as a superuser in the 'template1'database right after the cluster has been created - to be used with extreme care(by default empty)",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"post_init_template_sql_refs": schema.SingleNestedAttribute{
										Description:         "List of references to ConfigMaps or Secrets containing SQL filesto be executed as a superuser in the 'template1' database right afterthe cluster has been created. The references are processed in a specific order:first, all Secrets are processed, followed by all ConfigMaps.Within each group, the processing order follows the sequence specifiedin their respective arrays.(by default empty)",
										MarkdownDescription: "List of references to ConfigMaps or Secrets containing SQL filesto be executed as a superuser in the 'template1' database right afterthe cluster has been created. The references are processed in a specific order:first, all Secrets are processed, followed by all ConfigMaps.Within each group, the processing order follows the sequence specifiedin their respective arrays.(by default empty)",
										Attributes: map[string]schema.Attribute{
											"config_map_refs": schema.ListNestedAttribute{
												Description:         "ConfigMapRefs holds a list of references to ConfigMaps",
												MarkdownDescription: "ConfigMapRefs holds a list of references to ConfigMaps",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
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

											"secret_refs": schema.ListNestedAttribute{
												Description:         "SecretRefs holds a list of references to Secrets",
												MarkdownDescription: "SecretRefs holds a list of references to Secrets",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
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

									"secret": schema.SingleNestedAttribute{
										Description:         "Name of the secret containing the initial credentials for theowner of the user database. If empty a new secret will becreated from scratch",
										MarkdownDescription: "Name of the secret containing the initial credentials for theowner of the user database. If empty a new secret will becreated from scratch",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"wal_segment_size": schema.Int64Attribute{
										Description:         "The value in megabytes (1 to 1024) to be passed to the '--wal-segsize'option for initdb (default: empty, resulting in PostgreSQL default: 16MB)",
										MarkdownDescription: "The value in megabytes (1 to 1024) to be passed to the '--wal-segsize'option for initdb (default: empty, resulting in PostgreSQL default: 16MB)",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
											int64validator.AtMost(1024),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pg_basebackup": schema.SingleNestedAttribute{
								Description:         "Bootstrap the cluster taking a physical backup of another compatiblePostgreSQL instance",
								MarkdownDescription: "Bootstrap the cluster taking a physical backup of another compatiblePostgreSQL instance",
								Attributes: map[string]schema.Attribute{
									"database": schema.StringAttribute{
										Description:         "Name of the database used by the application. Default: 'app'.",
										MarkdownDescription: "Name of the database used by the application. Default: 'app'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"owner": schema.StringAttribute{
										Description:         "Name of the owner of the database in the instance to be usedby applications. Defaults to the value of the 'database' key.",
										MarkdownDescription: "Name of the owner of the database in the instance to be usedby applications. Defaults to the value of the 'database' key.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "Name of the secret containing the initial credentials for theowner of the user database. If empty a new secret will becreated from scratch",
										MarkdownDescription: "Name of the secret containing the initial credentials for theowner of the user database. If empty a new secret will becreated from scratch",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"source": schema.StringAttribute{
										Description:         "The name of the server of which we need to take a physical backup",
										MarkdownDescription: "The name of the server of which we need to take a physical backup",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"recovery": schema.SingleNestedAttribute{
								Description:         "Bootstrap the cluster from a backup",
								MarkdownDescription: "Bootstrap the cluster from a backup",
								Attributes: map[string]schema.Attribute{
									"backup": schema.SingleNestedAttribute{
										Description:         "The backup object containing the physical base backup from which toinitiate the recovery procedure.Mutually exclusive with 'source' and 'volumeSnapshots'.",
										MarkdownDescription: "The backup object containing the physical base backup from which toinitiate the recovery procedure.Mutually exclusive with 'source' and 'volumeSnapshots'.",
										Attributes: map[string]schema.Attribute{
											"endpoint_ca": schema.SingleNestedAttribute{
												Description:         "EndpointCA store the CA bundle of the barman endpoint.Useful when using self-signed certificates to avoiderrors with certificate issuer and barman-cloud-wal-archive.",
												MarkdownDescription: "EndpointCA store the CA bundle of the barman endpoint.Useful when using self-signed certificates to avoiderrors with certificate issuer and barman-cloud-wal-archive.",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to select",
														MarkdownDescription: "The key to select",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.",
														MarkdownDescription: "Name of the referent.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"database": schema.StringAttribute{
										Description:         "Name of the database used by the application. Default: 'app'.",
										MarkdownDescription: "Name of the database used by the application. Default: 'app'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"owner": schema.StringAttribute{
										Description:         "Name of the owner of the database in the instance to be usedby applications. Defaults to the value of the 'database' key.",
										MarkdownDescription: "Name of the owner of the database in the instance to be usedby applications. Defaults to the value of the 'database' key.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"recovery_target": schema.SingleNestedAttribute{
										Description:         "By default, the recovery process applies all the availableWAL files in the archive (full recovery). However, you can alsoend the recovery as soon as a consistent state is reached orrecover to a point-in-time (PITR) by specifying a 'RecoveryTarget' object,as expected by PostgreSQL (i.e., timestamp, transaction Id, LSN, ...).More info: https://www.postgresql.org/docs/current/runtime-config-wal.html#RUNTIME-CONFIG-WAL-RECOVERY-TARGET",
										MarkdownDescription: "By default, the recovery process applies all the availableWAL files in the archive (full recovery). However, you can alsoend the recovery as soon as a consistent state is reached orrecover to a point-in-time (PITR) by specifying a 'RecoveryTarget' object,as expected by PostgreSQL (i.e., timestamp, transaction Id, LSN, ...).More info: https://www.postgresql.org/docs/current/runtime-config-wal.html#RUNTIME-CONFIG-WAL-RECOVERY-TARGET",
										Attributes: map[string]schema.Attribute{
											"backup_id": schema.StringAttribute{
												Description:         "The ID of the backup from which to start the recovery process.If empty (default) the operator will automatically detect the backupbased on targetTime or targetLSN if specified. Otherwise use thelatest available backup in chronological order.",
												MarkdownDescription: "The ID of the backup from which to start the recovery process.If empty (default) the operator will automatically detect the backupbased on targetTime or targetLSN if specified. Otherwise use thelatest available backup in chronological order.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"exclusive": schema.BoolAttribute{
												Description:         "Set the target to be exclusive. If omitted, defaults to false, so thatin Postgres, 'recovery_target_inclusive' will be true",
												MarkdownDescription: "Set the target to be exclusive. If omitted, defaults to false, so thatin Postgres, 'recovery_target_inclusive' will be true",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_immediate": schema.BoolAttribute{
												Description:         "End recovery as soon as a consistent state is reached",
												MarkdownDescription: "End recovery as soon as a consistent state is reached",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_lsn": schema.StringAttribute{
												Description:         "The target LSN (Log Sequence Number)",
												MarkdownDescription: "The target LSN (Log Sequence Number)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_name": schema.StringAttribute{
												Description:         "The target name (to be previously createdwith 'pg_create_restore_point')",
												MarkdownDescription: "The target name (to be previously createdwith 'pg_create_restore_point')",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_tli": schema.StringAttribute{
												Description:         "The target timeline ('latest' or a positive integer)",
												MarkdownDescription: "The target timeline ('latest' or a positive integer)",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_time": schema.StringAttribute{
												Description:         "The target time as a timestamp in the RFC3339 standard",
												MarkdownDescription: "The target time as a timestamp in the RFC3339 standard",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"target_xid": schema.StringAttribute{
												Description:         "The target transaction ID",
												MarkdownDescription: "The target transaction ID",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret": schema.SingleNestedAttribute{
										Description:         "Name of the secret containing the initial credentials for theowner of the user database. If empty a new secret will becreated from scratch",
										MarkdownDescription: "Name of the secret containing the initial credentials for theowner of the user database. If empty a new secret will becreated from scratch",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Name of the referent.",
												MarkdownDescription: "Name of the referent.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"source": schema.StringAttribute{
										Description:         "The external cluster whose backup we will restore. This is alsoused as the name of the folder under which the backup is stored,so it must be set to the name of the source clusterMutually exclusive with 'backup'.",
										MarkdownDescription: "The external cluster whose backup we will restore. This is alsoused as the name of the folder under which the backup is stored,so it must be set to the name of the source clusterMutually exclusive with 'backup'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_snapshots": schema.SingleNestedAttribute{
										Description:         "The static PVC data source(s) from which to initiate therecovery procedure. Currently supporting 'VolumeSnapshot'and 'PersistentVolumeClaim' resources that map an existingPVC group, compatible with CloudNativePG, and taken witha cold backup copy on a fenced Postgres instance (limitationwhich will be removed in the future when online backupwill be implemented).Mutually exclusive with 'backup'.",
										MarkdownDescription: "The static PVC data source(s) from which to initiate therecovery procedure. Currently supporting 'VolumeSnapshot'and 'PersistentVolumeClaim' resources that map an existingPVC group, compatible with CloudNativePG, and taken witha cold backup copy on a fenced Postgres instance (limitationwhich will be removed in the future when online backupwill be implemented).Mutually exclusive with 'backup'.",
										Attributes: map[string]schema.Attribute{
											"storage": schema.SingleNestedAttribute{
												Description:         "Configuration of the storage of the instances",
												MarkdownDescription: "Configuration of the storage of the instances",
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is the type of resource being referenced",
														MarkdownDescription: "Kind is the type of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of resource being referenced",
														MarkdownDescription: "Name is the name of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: true,
												Optional: false,
												Computed: false,
											},

											"tablespace_storage": schema.SingleNestedAttribute{
												Description:         "Configuration of the storage for PostgreSQL tablespaces",
												MarkdownDescription: "Configuration of the storage for PostgreSQL tablespaces",
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is the type of resource being referenced",
														MarkdownDescription: "Kind is the type of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of resource being referenced",
														MarkdownDescription: "Name is the name of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"wal_storage": schema.SingleNestedAttribute{
												Description:         "Configuration of the storage for PostgreSQL WAL (Write-Ahead Log)",
												MarkdownDescription: "Configuration of the storage for PostgreSQL WAL (Write-Ahead Log)",
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is the type of resource being referenced",
														MarkdownDescription: "Kind is the type of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of resource being referenced",
														MarkdownDescription: "Name is the name of resource being referenced",
														Required:            true,
														Optional:            false,
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"certificates": schema.SingleNestedAttribute{
						Description:         "The configuration for the CA and related certificates",
						MarkdownDescription: "The configuration for the CA and related certificates",
						Attributes: map[string]schema.Attribute{
							"client_ca_secret": schema.StringAttribute{
								Description:         "The secret containing the Client CA certificate. If not defined, a new secret will be createdwith a self-signed CA and will be used to generate all the client certificates.<br /><br />Contains:<br /><br />- 'ca.crt': CA that should be used to validate the client certificates,used as 'ssl_ca_file' of all the instances.<br />- 'ca.key': key used to generate client certificates, if ReplicationTLSSecret is provided,this can be omitted.<br />",
								MarkdownDescription: "The secret containing the Client CA certificate. If not defined, a new secret will be createdwith a self-signed CA and will be used to generate all the client certificates.<br /><br />Contains:<br /><br />- 'ca.crt': CA that should be used to validate the client certificates,used as 'ssl_ca_file' of all the instances.<br />- 'ca.key': key used to generate client certificates, if ReplicationTLSSecret is provided,this can be omitted.<br />",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replication_tls_secret": schema.StringAttribute{
								Description:         "The secret of type kubernetes.io/tls containing the client certificate to authenticate asthe 'streaming_replica' user.If not defined, ClientCASecret must provide also 'ca.key', and a new secret will becreated using the provided CA.",
								MarkdownDescription: "The secret of type kubernetes.io/tls containing the client certificate to authenticate asthe 'streaming_replica' user.If not defined, ClientCASecret must provide also 'ca.key', and a new secret will becreated using the provided CA.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_alt_dns_names": schema.ListAttribute{
								Description:         "The list of the server alternative DNS names to be added to the generated server TLS certificates, when required.",
								MarkdownDescription: "The list of the server alternative DNS names to be added to the generated server TLS certificates, when required.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_ca_secret": schema.StringAttribute{
								Description:         "The secret containing the Server CA certificate. If not defined, a new secret will be createdwith a self-signed CA and will be used to generate the TLS certificate ServerTLSSecret.<br /><br />Contains:<br /><br />- 'ca.crt': CA that should be used to validate the server certificate,used as 'sslrootcert' in client connection strings.<br />- 'ca.key': key used to generate Server SSL certs, if ServerTLSSecret is provided,this can be omitted.<br />",
								MarkdownDescription: "The secret containing the Server CA certificate. If not defined, a new secret will be createdwith a self-signed CA and will be used to generate the TLS certificate ServerTLSSecret.<br /><br />Contains:<br /><br />- 'ca.crt': CA that should be used to validate the server certificate,used as 'sslrootcert' in client connection strings.<br />- 'ca.key': key used to generate Server SSL certs, if ServerTLSSecret is provided,this can be omitted.<br />",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"server_tls_secret": schema.StringAttribute{
								Description:         "The secret of type kubernetes.io/tls containing the server TLS certificate and key that will be set as'ssl_cert_file' and 'ssl_key_file' so that clients can connect to postgres securely.If not defined, ServerCASecret must provide also 'ca.key' and a new secret will becreated using the provided CA.",
								MarkdownDescription: "The secret of type kubernetes.io/tls containing the server TLS certificate and key that will be set as'ssl_cert_file' and 'ssl_key_file' so that clients can connect to postgres securely.If not defined, ServerCASecret must provide also 'ca.key' and a new secret will becreated using the provided CA.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"description": schema.StringAttribute{
						Description:         "Description of this PostgreSQL cluster",
						MarkdownDescription: "Description of this PostgreSQL cluster",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_pdb": schema.BoolAttribute{
						Description:         "Manage the 'PodDisruptionBudget' resources within the cluster. Whenconfigured as 'true' (default setting), the pod disruption budgetswill safeguard the primary node from being terminated. Conversely,setting it to 'false' will result in the absence of any'PodDisruptionBudget' resource, permitting the shutdown of all nodeshosting the PostgreSQL cluster. This latter configuration isadvisable for any PostgreSQL cluster employed fordevelopment/staging purposes.",
						MarkdownDescription: "Manage the 'PodDisruptionBudget' resources within the cluster. Whenconfigured as 'true' (default setting), the pod disruption budgetswill safeguard the primary node from being terminated. Conversely,setting it to 'false' will result in the absence of any'PodDisruptionBudget' resource, permitting the shutdown of all nodeshosting the PostgreSQL cluster. This latter configuration isadvisable for any PostgreSQL cluster employed fordevelopment/staging purposes.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"enable_superuser_access": schema.BoolAttribute{
						Description:         "When this option is enabled, the operator will use the 'SuperuserSecret'to update the 'postgres' user password (if the secret isnot present, the operator will automatically create one). When thisoption is disabled, the operator will ignore the 'SuperuserSecret' content, deleteit when automatically created, and then blank the password of the 'postgres'user by setting it to 'NULL'. Disabled by default.",
						MarkdownDescription: "When this option is enabled, the operator will use the 'SuperuserSecret'to update the 'postgres' user password (if the secret isnot present, the operator will automatically create one). When thisoption is disabled, the operator will ignore the 'SuperuserSecret' content, deleteit when automatically created, and then blank the password of the 'postgres'user by setting it to 'NULL'. Disabled by default.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"env": schema.ListNestedAttribute{
						Description:         "Env follows the Env format to pass environment variablesto the pods created in the cluster",
						MarkdownDescription: "Env follows the Env format to pass environment variablesto the pods created in the cluster",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
									MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"value": schema.StringAttribute{
									Description:         "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
									MarkdownDescription: "Variable references $(VAR_NAME) are expandedusing the previously defined environment variables in the container andany service environment variables. If a variable cannot be resolved,the reference in the input string will be unchanged. Double $$ are reducedto a single $, which allows for escaping the $(VAR_NAME) syntax: i.e.'$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'.Escaped references will never be expanded, regardless of whether the variableexists or not.Defaults to ''.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"value_from": schema.SingleNestedAttribute{
									Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
									MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",
									Attributes: map[string]schema.Attribute{
										"config_map_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a ConfigMap.",
											MarkdownDescription: "Selects a key of a ConfigMap.",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key to select.",
													MarkdownDescription: "The key to select.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the ConfigMap or its key must be defined",
													MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"field_ref": schema.SingleNestedAttribute{
											Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']',spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
											Attributes: map[string]schema.Attribute{
												"api_version": schema.StringAttribute{
													Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
													MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"field_path": schema.StringAttribute{
													Description:         "Path of the field to select in the specified API version.",
													MarkdownDescription: "Path of the field to select in the specified API version.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"resource_field_ref": schema.SingleNestedAttribute{
											Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
											Attributes: map[string]schema.Attribute{
												"container_name": schema.StringAttribute{
													Description:         "Container name: required for volumes, optional for env vars",
													MarkdownDescription: "Container name: required for volumes, optional for env vars",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"divisor": schema.StringAttribute{
													Description:         "Specifies the output format of the exposed resources, defaults to '1'",
													MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"resource": schema.StringAttribute{
													Description:         "Required: resource to select",
													MarkdownDescription: "Required: resource to select",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"secret_key_ref": schema.SingleNestedAttribute{
											Description:         "Selects a key of a secret in the pod's namespace",
											MarkdownDescription: "Selects a key of a secret in the pod's namespace",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key of the secret to select from.  Must be a valid secret key.",
													MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "Specify whether the Secret or its key must be defined",
													MarkdownDescription: "Specify whether the Secret or its key must be defined",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"env_from": schema.ListNestedAttribute{
						Description:         "EnvFrom follows the EnvFrom format to pass environment variablessources to the pods to be used by Env",
						MarkdownDescription: "EnvFrom follows the EnvFrom format to pass environment variablessources to the pods to be used by Env",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"config_map_ref": schema.SingleNestedAttribute{
									Description:         "The ConfigMap to select from",
									MarkdownDescription: "The ConfigMap to select from",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the ConfigMap must be defined",
											MarkdownDescription: "Specify whether the ConfigMap must be defined",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"prefix": schema.StringAttribute{
									Description:         "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
									MarkdownDescription: "An optional identifier to prepend to each key in the ConfigMap. Must be a C_IDENTIFIER.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"secret_ref": schema.SingleNestedAttribute{
									Description:         "The Secret to select from",
									MarkdownDescription: "The Secret to select from",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the Secret must be defined",
											MarkdownDescription: "Specify whether the Secret must be defined",
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

					"ephemeral_volume_source": schema.SingleNestedAttribute{
						Description:         "EphemeralVolumeSource allows the user to configure the source of ephemeral volumes.",
						MarkdownDescription: "EphemeralVolumeSource allows the user to configure the source of ephemeral volumes.",
						Attributes: map[string]schema.Attribute{
							"volume_claim_template": schema.SingleNestedAttribute{
								Description:         "Will be used to create a stand-alone PVC to provision the volume.The pod in which this EphemeralVolumeSource is embedded will be theowner of the PVC, i.e. the PVC will be deleted together with thepod.  The name of the PVC will be '<pod name>-<volume name>' where'<volume name>' is the name from the 'PodSpec.Volumes' arrayentry. Pod validation will reject the pod if the concatenated nameis not valid for a PVC (for example, too long).An existing PVC with that name that is not owned by the podwill *not* be used for the pod to avoid using an unrelatedvolume by mistake. Starting the pod is then blocked untilthe unrelated PVC is removed. If such a pre-created PVC ismeant to be used by the pod, the PVC has to updated with anowner reference to the pod once the pod exists. Normallythis should not be necessary, but it may be useful whenmanually reconstructing a broken cluster.This field is read-only and no changes will be made by Kubernetesto the PVC after it has been created.Required, must not be nil.",
								MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume.The pod in which this EphemeralVolumeSource is embedded will be theowner of the PVC, i.e. the PVC will be deleted together with thepod.  The name of the PVC will be '<pod name>-<volume name>' where'<volume name>' is the name from the 'PodSpec.Volumes' arrayentry. Pod validation will reject the pod if the concatenated nameis not valid for a PVC (for example, too long).An existing PVC with that name that is not owned by the podwill *not* be used for the pod to avoid using an unrelatedvolume by mistake. Starting the pod is then blocked untilthe unrelated PVC is removed. If such a pre-created PVC ismeant to be used by the pod, the PVC has to updated with anowner reference to the pod once the pod exists. Normallythis should not be necessary, but it may be useful whenmanually reconstructing a broken cluster.This field is read-only and no changes will be made by Kubernetesto the PVC after it has been created.Required, must not be nil.",
								Attributes: map[string]schema.Attribute{
									"metadata": schema.MapAttribute{
										Description:         "May contain labels and annotations that will be copied into the PVCwhen creating it. No other fields are allowed and will be rejected duringvalidation.",
										MarkdownDescription: "May contain labels and annotations that will be copied into the PVCwhen creating it. No other fields are allowed and will be rejected duringvalidation.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"spec": schema.SingleNestedAttribute{
										Description:         "The specification for the PersistentVolumeClaim. The entire content iscopied unchanged into the PVC that gets created from thistemplate. The same fields as in a PersistentVolumeClaimare also valid here.",
										MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content iscopied unchanged into the PVC that gets created from thistemplate. The same fields as in a PersistentVolumeClaimare also valid here.",
										Attributes: map[string]schema.Attribute{
											"access_modes": schema.ListAttribute{
												Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
												MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"data_source": schema.SingleNestedAttribute{
												Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
												MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is the type of resource being referenced",
														MarkdownDescription: "Kind is the type of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of resource being referenced",
														MarkdownDescription: "Name is the name of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"data_source_ref": schema.SingleNestedAttribute{
												Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
												MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
												Attributes: map[string]schema.Attribute{
													"api_group": schema.StringAttribute{
														Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind is the type of resource being referenced",
														MarkdownDescription: "Kind is the type of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of resource being referenced",
														MarkdownDescription: "Name is the name of resource being referenced",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
														MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
												MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
												Attributes: map[string]schema.Attribute{
													"limits": schema.MapAttribute{
														Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.MapAttribute{
														Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

											"selector": schema.SingleNestedAttribute{
												Description:         "selector is a label query over volumes to consider for binding.",
												MarkdownDescription: "selector is a label query over volumes to consider for binding.",
												Attributes: map[string]schema.Attribute{
													"match_expressions": schema.ListNestedAttribute{
														Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "key is the label key that the selector applies to.",
																	MarkdownDescription: "key is the label key that the selector applies to.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"operator": schema.StringAttribute{
																	Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																	MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"values": schema.ListAttribute{
																	Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																	MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
														Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
														MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

											"storage_class_name": schema.StringAttribute{
												Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
												MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"volume_attributes_class_name": schema.StringAttribute{
												Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
												MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"volume_mode": schema.StringAttribute{
												Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
												MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"volume_name": schema.StringAttribute{
												Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
												MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"ephemeral_volumes_size_limit": schema.SingleNestedAttribute{
						Description:         "EphemeralVolumesSizeLimit allows the user to set the limits for the ephemeralvolumes",
						MarkdownDescription: "EphemeralVolumesSizeLimit allows the user to set the limits for the ephemeralvolumes",
						Attributes: map[string]schema.Attribute{
							"shm": schema.StringAttribute{
								Description:         "Shm is the size limit of the shared memory volume",
								MarkdownDescription: "Shm is the size limit of the shared memory volume",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"temporary_data": schema.StringAttribute{
								Description:         "TemporaryData is the size limit of the temporary data volume",
								MarkdownDescription: "TemporaryData is the size limit of the temporary data volume",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"external_clusters": schema.ListNestedAttribute{
						Description:         "The list of external clusters which are used in the configuration",
						MarkdownDescription: "The list of external clusters which are used in the configuration",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"barman_object_store": schema.SingleNestedAttribute{
									Description:         "The configuration for the barman-cloud tool suite",
									MarkdownDescription: "The configuration for the barman-cloud tool suite",
									Attributes: map[string]schema.Attribute{
										"azure_credentials": schema.SingleNestedAttribute{
											Description:         "The credentials to use to upload data to Azure Blob Storage",
											MarkdownDescription: "The credentials to use to upload data to Azure Blob Storage",
											Attributes: map[string]schema.Attribute{
												"connection_string": schema.SingleNestedAttribute{
													Description:         "The connection string to be used",
													MarkdownDescription: "The connection string to be used",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"inherit_from_azure_ad": schema.BoolAttribute{
													Description:         "Use the Azure AD based authentication without providing explicitly the keys.",
													MarkdownDescription: "Use the Azure AD based authentication without providing explicitly the keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"storage_account": schema.SingleNestedAttribute{
													Description:         "The storage account where to upload data",
													MarkdownDescription: "The storage account where to upload data",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"storage_key": schema.SingleNestedAttribute{
													Description:         "The storage account key to be used in conjunctionwith the storage account name",
													MarkdownDescription: "The storage account key to be used in conjunctionwith the storage account name",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"storage_sas_token": schema.SingleNestedAttribute{
													Description:         "A shared-access-signature to be used in conjunction withthe storage account name",
													MarkdownDescription: "A shared-access-signature to be used in conjunction withthe storage account name",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
															Required:            true,
															Optional:            false,
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

										"data": schema.SingleNestedAttribute{
											Description:         "The configuration to be used to backup the data filesWhen not defined, base backups files will be stored uncompressed and maybe unencrypted in the object store, according to the bucket defaultpolicy.",
											MarkdownDescription: "The configuration to be used to backup the data filesWhen not defined, base backups files will be stored uncompressed and maybe unencrypted in the object store, according to the bucket defaultpolicy.",
											Attributes: map[string]schema.Attribute{
												"additional_command_args": schema.ListAttribute{
													Description:         "AdditionalCommandArgs represents additional arguments that can be appendedto the 'barman-cloud-backup' command-line invocation. These argumentsprovide flexibility to customize the backup process further according tospecific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-backup' command, to avoid potential errors or unintendedbehavior during execution.",
													MarkdownDescription: "AdditionalCommandArgs represents additional arguments that can be appendedto the 'barman-cloud-backup' command-line invocation. These argumentsprovide flexibility to customize the backup process further according tospecific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-backup' command, to avoid potential errors or unintendedbehavior during execution.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"compression": schema.StringAttribute{
													Description:         "Compress a backup file (a tar file per tablespace) while streaming itto the object store. Available options are empty string (nocompression, default), 'gzip', 'bzip2' or 'snappy'.",
													MarkdownDescription: "Compress a backup file (a tar file per tablespace) while streaming itto the object store. Available options are empty string (nocompression, default), 'gzip', 'bzip2' or 'snappy'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("gzip", "bzip2", "snappy"),
													},
												},

												"encryption": schema.StringAttribute{
													Description:         "Whenever to force the encryption of files (if the bucket isnot already configured for that).Allowed options are empty string (use the bucket policy, default),'AES256' and 'aws:kms'",
													MarkdownDescription: "Whenever to force the encryption of files (if the bucket isnot already configured for that).Allowed options are empty string (use the bucket policy, default),'AES256' and 'aws:kms'",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("AES256", "aws:kms"),
													},
												},

												"immediate_checkpoint": schema.BoolAttribute{
													Description:         "Control whether the I/O workload for the backup initial checkpoint willbe limited, according to the 'checkpoint_completion_target' setting onthe PostgreSQL server. If set to true, an immediate checkpoint will beused, meaning PostgreSQL will complete the checkpoint as soon aspossible. 'false' by default.",
													MarkdownDescription: "Control whether the I/O workload for the backup initial checkpoint willbe limited, according to the 'checkpoint_completion_target' setting onthe PostgreSQL server. If set to true, an immediate checkpoint will beused, meaning PostgreSQL will complete the checkpoint as soon aspossible. 'false' by default.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"jobs": schema.Int64Attribute{
													Description:         "The number of parallel jobs to be used to upload the backup, defaultsto 2",
													MarkdownDescription: "The number of parallel jobs to be used to upload the backup, defaultsto 2",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"destination_path": schema.StringAttribute{
											Description:         "The path where to store the backup (i.e. s3://bucket/path/to/folder)this path, with different destination folders, will be used for WALsand for data",
											MarkdownDescription: "The path where to store the backup (i.e. s3://bucket/path/to/folder)this path, with different destination folders, will be used for WALsand for data",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtLeast(1),
											},
										},

										"endpoint_ca": schema.SingleNestedAttribute{
											Description:         "EndpointCA store the CA bundle of the barman endpoint.Useful when using self-signed certificates to avoiderrors with certificate issuer and barman-cloud-wal-archive",
											MarkdownDescription: "EndpointCA store the CA bundle of the barman endpoint.Useful when using self-signed certificates to avoiderrors with certificate issuer and barman-cloud-wal-archive",
											Attributes: map[string]schema.Attribute{
												"key": schema.StringAttribute{
													Description:         "The key to select",
													MarkdownDescription: "The key to select",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the referent.",
													MarkdownDescription: "Name of the referent.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"endpoint_url": schema.StringAttribute{
											Description:         "Endpoint to be used to upload data to the cloud,overriding the automatic endpoint discovery",
											MarkdownDescription: "Endpoint to be used to upload data to the cloud,overriding the automatic endpoint discovery",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"google_credentials": schema.SingleNestedAttribute{
											Description:         "The credentials to use to upload data to Google Cloud Storage",
											MarkdownDescription: "The credentials to use to upload data to Google Cloud Storage",
											Attributes: map[string]schema.Attribute{
												"application_credentials": schema.SingleNestedAttribute{
													Description:         "The secret containing the Google Cloud Storage JSON file with the credentials",
													MarkdownDescription: "The secret containing the Google Cloud Storage JSON file with the credentials",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"gke_environment": schema.BoolAttribute{
													Description:         "If set to true, will presume that it's running inside a GKE environment,default to false.",
													MarkdownDescription: "If set to true, will presume that it's running inside a GKE environment,default to false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"history_tags": schema.MapAttribute{
											Description:         "HistoryTags is a list of key value pairs that will be passed to theBarman --history-tags option.",
											MarkdownDescription: "HistoryTags is a list of key value pairs that will be passed to theBarman --history-tags option.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"s3_credentials": schema.SingleNestedAttribute{
											Description:         "The credentials to use to upload data to S3",
											MarkdownDescription: "The credentials to use to upload data to S3",
											Attributes: map[string]schema.Attribute{
												"access_key_id": schema.SingleNestedAttribute{
													Description:         "The reference to the access key id",
													MarkdownDescription: "The reference to the access key id",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"inherit_from_iam_role": schema.BoolAttribute{
													Description:         "Use the role based authentication without providing explicitly the keys.",
													MarkdownDescription: "Use the role based authentication without providing explicitly the keys.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"region": schema.SingleNestedAttribute{
													Description:         "The reference to the secret containing the region name",
													MarkdownDescription: "The reference to the secret containing the region name",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"secret_access_key": schema.SingleNestedAttribute{
													Description:         "The reference to the secret access key",
													MarkdownDescription: "The reference to the secret access key",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"session_token": schema.SingleNestedAttribute{
													Description:         "The references to the session key",
													MarkdownDescription: "The references to the session key",
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "The key to select",
															MarkdownDescription: "The key to select",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the referent.",
															MarkdownDescription: "Name of the referent.",
															Required:            true,
															Optional:            false,
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

										"server_name": schema.StringAttribute{
											Description:         "The server name on S3, the cluster name is used if thisparameter is omitted",
											MarkdownDescription: "The server name on S3, the cluster name is used if thisparameter is omitted",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tags": schema.MapAttribute{
											Description:         "Tags is a list of key value pairs that will be passed to theBarman --tags option.",
											MarkdownDescription: "Tags is a list of key value pairs that will be passed to theBarman --tags option.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"wal": schema.SingleNestedAttribute{
											Description:         "The configuration for the backup of the WAL stream.When not defined, WAL files will be stored uncompressed and may beunencrypted in the object store, according to the bucket default policy.",
											MarkdownDescription: "The configuration for the backup of the WAL stream.When not defined, WAL files will be stored uncompressed and may beunencrypted in the object store, according to the bucket default policy.",
											Attributes: map[string]schema.Attribute{
												"archive_additional_command_args": schema.ListAttribute{
													Description:         "Additional arguments that can be appended to the 'barman-cloud-wal-archive'command-line invocation. These arguments provide flexibility to customizethe WAL archive process further, according to specific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-wal-archive' command, to avoid potential errors or unintendedbehavior during execution.",
													MarkdownDescription: "Additional arguments that can be appended to the 'barman-cloud-wal-archive'command-line invocation. These arguments provide flexibility to customizethe WAL archive process further, according to specific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-wal-archive' command, to avoid potential errors or unintendedbehavior during execution.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"compression": schema.StringAttribute{
													Description:         "Compress a WAL file before sending it to the object store. Availableoptions are empty string (no compression, default), 'gzip', 'bzip2' or 'snappy'.",
													MarkdownDescription: "Compress a WAL file before sending it to the object store. Availableoptions are empty string (no compression, default), 'gzip', 'bzip2' or 'snappy'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("gzip", "bzip2", "snappy"),
													},
												},

												"encryption": schema.StringAttribute{
													Description:         "Whenever to force the encryption of files (if the bucket isnot already configured for that).Allowed options are empty string (use the bucket policy, default),'AES256' and 'aws:kms'",
													MarkdownDescription: "Whenever to force the encryption of files (if the bucket isnot already configured for that).Allowed options are empty string (use the bucket policy, default),'AES256' and 'aws:kms'",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("AES256", "aws:kms"),
													},
												},

												"max_parallel": schema.Int64Attribute{
													Description:         "Number of WAL files to be either archived in parallel (when thePostgreSQL instance is archiving to a backup object store) orrestored in parallel (when a PostgreSQL standby is fetching WALfiles from a recovery object store). If not specified, WAL fileswill be processed one at a time. It accepts a positive integer as avalue - with 1 being the minimum accepted value.",
													MarkdownDescription: "Number of WAL files to be either archived in parallel (when thePostgreSQL instance is archiving to a backup object store) orrestored in parallel (when a PostgreSQL standby is fetching WALfiles from a recovery object store). If not specified, WAL fileswill be processed one at a time. It accepts a positive integer as avalue - with 1 being the minimum accepted value.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.Int64{
														int64validator.AtLeast(1),
													},
												},

												"restore_additional_command_args": schema.ListAttribute{
													Description:         "Additional arguments that can be appended to the 'barman-cloud-wal-restore'command-line invocation. These arguments provide flexibility to customizethe WAL restore process further, according to specific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-wal-restore' command, to avoid potential errors or unintendedbehavior during execution.",
													MarkdownDescription: "Additional arguments that can be appended to the 'barman-cloud-wal-restore'command-line invocation. These arguments provide flexibility to customizethe WAL restore process further, according to specific requirements or configurations.Example:In a scenario where specialized backup options are required, such as settinga specific timeout or defining custom behavior, users can use this fieldto specify additional command arguments.Note:It's essential to ensure that the provided arguments are valid and supportedby the 'barman-cloud-wal-restore' command, to avoid potential errors or unintendedbehavior during execution.",
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

								"connection_parameters": schema.MapAttribute{
									Description:         "The list of connection parameters, such as dbname, host, username, etc",
									MarkdownDescription: "The list of connection parameters, such as dbname, host, username, etc",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "The server name, required",
									MarkdownDescription: "The server name, required",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"password": schema.SingleNestedAttribute{
									Description:         "The reference to the password to be used to connect to the server.If a password is provided, CloudNativePG creates a PostgreSQLpassfile at '/controller/external/NAME/pass' (where 'NAME' is thecluster's name). This passfile is automatically referenced in theconnection string when establishing a connection to the remotePostgreSQL server from the current PostgreSQL 'Cluster'. This ensuressecure and efficient password management for external clusters.",
									MarkdownDescription: "The reference to the password to be used to connect to the server.If a password is provided, CloudNativePG creates a PostgreSQLpassfile at '/controller/external/NAME/pass' (where 'NAME' is thecluster's name). This passfile is automatically referenced in theconnection string when establishing a connection to the remotePostgreSQL server from the current PostgreSQL 'Cluster'. This ensuressecure and efficient password management for external clusters.",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key of the secret to select from.  Must be a valid secret key.",
											MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the Secret or its key must be defined",
											MarkdownDescription: "Specify whether the Secret or its key must be defined",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ssl_cert": schema.SingleNestedAttribute{
									Description:         "The reference to an SSL certificate to be used to connect to thisinstance",
									MarkdownDescription: "The reference to an SSL certificate to be used to connect to thisinstance",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key of the secret to select from.  Must be a valid secret key.",
											MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the Secret or its key must be defined",
											MarkdownDescription: "Specify whether the Secret or its key must be defined",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ssl_key": schema.SingleNestedAttribute{
									Description:         "The reference to an SSL private key to be used to connect to thisinstance",
									MarkdownDescription: "The reference to an SSL private key to be used to connect to thisinstance",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key of the secret to select from.  Must be a valid secret key.",
											MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the Secret or its key must be defined",
											MarkdownDescription: "Specify whether the Secret or its key must be defined",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"ssl_root_cert": schema.SingleNestedAttribute{
									Description:         "The reference to an SSL CA public key to be used to connect to thisinstance",
									MarkdownDescription: "The reference to an SSL CA public key to be used to connect to thisinstance",
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key of the secret to select from.  Must be a valid secret key.",
											MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the Secret or its key must be defined",
											MarkdownDescription: "Specify whether the Secret or its key must be defined",
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

					"failover_delay": schema.Int64Attribute{
						Description:         "The amount of time (in seconds) to wait before triggering a failoverafter the primary PostgreSQL instance in the cluster was detectedto be unhealthy",
						MarkdownDescription: "The amount of time (in seconds) to wait before triggering a failoverafter the primary PostgreSQL instance in the cluster was detectedto be unhealthy",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_catalog_ref": schema.SingleNestedAttribute{
						Description:         "Defines the major PostgreSQL version we want to use within an ImageCatalog",
						MarkdownDescription: "Defines the major PostgreSQL version we want to use within an ImageCatalog",
						Attributes: map[string]schema.Attribute{
							"api_group": schema.StringAttribute{
								Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
								MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kind": schema.StringAttribute{
								Description:         "Kind is the type of resource being referenced",
								MarkdownDescription: "Kind is the type of resource being referenced",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"major": schema.Int64Attribute{
								Description:         "The major version of PostgreSQL we want to use from the ImageCatalog",
								MarkdownDescription: "The major version of PostgreSQL we want to use from the ImageCatalog",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"name": schema.StringAttribute{
								Description:         "Name is the name of resource being referenced",
								MarkdownDescription: "Name is the name of resource being referenced",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"image_name": schema.StringAttribute{
						Description:         "Name of the container image, supporting both tags ('<image>:<tag>')and digests for deterministic and repeatable deployments('<image>:<tag>@sha256:<digestValue>')",
						MarkdownDescription: "Name of the container image, supporting both tags ('<image>:<tag>')and digests for deterministic and repeatable deployments('<image>:<tag>@sha256:<digestValue>')",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_policy": schema.StringAttribute{
						Description:         "Image pull policy.One of 'Always', 'Never' or 'IfNotPresent'.If not defined, it defaults to 'IfNotPresent'.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						MarkdownDescription: "Image pull policy.One of 'Always', 'Never' or 'IfNotPresent'.If not defined, it defaults to 'IfNotPresent'.Cannot be updated.More info: https://kubernetes.io/docs/concepts/containers/images#updating-images",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"image_pull_secrets": schema.ListNestedAttribute{
						Description:         "The list of pull secrets to be used to pull the images",
						MarkdownDescription: "The list of pull secrets to be used to pull the images",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Name of the referent.",
									MarkdownDescription: "Name of the referent.",
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

					"inherited_metadata": schema.SingleNestedAttribute{
						Description:         "Metadata that will be inherited by all objects related to the Cluster",
						MarkdownDescription: "Metadata that will be inherited by all objects related to the Cluster",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "",
								MarkdownDescription: "",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"labels": schema.MapAttribute{
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

					"instances": schema.Int64Attribute{
						Description:         "Number of instances required in the cluster",
						MarkdownDescription: "Number of instances required in the cluster",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
						},
					},

					"liveness_probe_timeout": schema.Int64Attribute{
						Description:         "LivenessProbeTimeout is the time (in seconds) that is allowed for a PostgreSQL instanceto successfully respond to the liveness probe (default 30).The Liveness probe failure threshold is derived from this value using the formula:ceiling(livenessProbe / 10).",
						MarkdownDescription: "LivenessProbeTimeout is the time (in seconds) that is allowed for a PostgreSQL instanceto successfully respond to the liveness probe (default 30).The Liveness probe failure threshold is derived from this value using the formula:ceiling(livenessProbe / 10).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"log_level": schema.StringAttribute{
						Description:         "The instances' log level, one of the following values: error, warning, info (default), debug, trace",
						MarkdownDescription: "The instances' log level, one of the following values: error, warning, info (default), debug, trace",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("error", "warning", "info", "debug", "trace"),
						},
					},

					"managed": schema.SingleNestedAttribute{
						Description:         "The configuration that is used by the portions of PostgreSQL that are managed by the instance manager",
						MarkdownDescription: "The configuration that is used by the portions of PostgreSQL that are managed by the instance manager",
						Attributes: map[string]schema.Attribute{
							"roles": schema.ListNestedAttribute{
								Description:         "Database roles managed by the 'Cluster'",
								MarkdownDescription: "Database roles managed by the 'Cluster'",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"bypassrls": schema.BoolAttribute{
											Description:         "Whether a role bypasses every row-level security (RLS) policy.Default is 'false'.",
											MarkdownDescription: "Whether a role bypasses every row-level security (RLS) policy.Default is 'false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"comment": schema.StringAttribute{
											Description:         "Description of the role",
											MarkdownDescription: "Description of the role",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"connection_limit": schema.Int64Attribute{
											Description:         "If the role can log in, this specifies how many concurrentconnections the role can make. '-1' (the default) means no limit.",
											MarkdownDescription: "If the role can log in, this specifies how many concurrentconnections the role can make. '-1' (the default) means no limit.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"createdb": schema.BoolAttribute{
											Description:         "When set to 'true', the role being defined will be allowed to createnew databases. Specifying 'false' (default) will deny a role theability to create databases.",
											MarkdownDescription: "When set to 'true', the role being defined will be allowed to createnew databases. Specifying 'false' (default) will deny a role theability to create databases.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"createrole": schema.BoolAttribute{
											Description:         "Whether the role will be permitted to create, alter, drop, commenton, change the security label for, and grant or revoke membership inother roles. Default is 'false'.",
											MarkdownDescription: "Whether the role will be permitted to create, alter, drop, commenton, change the security label for, and grant or revoke membership inother roles. Default is 'false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"disable_password": schema.BoolAttribute{
											Description:         "DisablePassword indicates that a role's password should be set to NULL in Postgres",
											MarkdownDescription: "DisablePassword indicates that a role's password should be set to NULL in Postgres",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"ensure": schema.StringAttribute{
											Description:         "Ensure the role is 'present' or 'absent' - defaults to 'present'",
											MarkdownDescription: "Ensure the role is 'present' or 'absent' - defaults to 'present'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("present", "absent"),
											},
										},

										"in_roles": schema.ListAttribute{
											Description:         "List of one or more existing roles to which this role will beimmediately added as a new member. Default empty.",
											MarkdownDescription: "List of one or more existing roles to which this role will beimmediately added as a new member. Default empty.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"inherit": schema.BoolAttribute{
											Description:         "Whether a role 'inherits' the privileges of roles it is a member of.Defaults is 'true'.",
											MarkdownDescription: "Whether a role 'inherits' the privileges of roles it is a member of.Defaults is 'true'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"login": schema.BoolAttribute{
											Description:         "Whether the role is allowed to log in. A role having the 'login'attribute can be thought of as a user. Roles without this attributeare useful for managing database privileges, but are not users inthe usual sense of the word. Default is 'false'.",
											MarkdownDescription: "Whether the role is allowed to log in. A role having the 'login'attribute can be thought of as a user. Roles without this attributeare useful for managing database privileges, but are not users inthe usual sense of the word. Default is 'false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the role",
											MarkdownDescription: "Name of the role",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"password_secret": schema.SingleNestedAttribute{
											Description:         "Secret containing the password of the role (if present)If null, the password will be ignored unless DisablePassword is set",
											MarkdownDescription: "Secret containing the password of the role (if present)If null, the password will be ignored unless DisablePassword is set",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent.",
													MarkdownDescription: "Name of the referent.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"replication": schema.BoolAttribute{
											Description:         "Whether a role is a replication role. A role must have thisattribute (or be a superuser) in order to be able to connect to theserver in replication mode (physical or logical replication) and inorder to be able to create or drop replication slots. A role havingthe 'replication' attribute is a very highly privileged role, andshould only be used on roles actually used for replication. Defaultis 'false'.",
											MarkdownDescription: "Whether a role is a replication role. A role must have thisattribute (or be a superuser) in order to be able to connect to theserver in replication mode (physical or logical replication) and inorder to be able to create or drop replication slots. A role havingthe 'replication' attribute is a very highly privileged role, andshould only be used on roles actually used for replication. Defaultis 'false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"superuser": schema.BoolAttribute{
											Description:         "Whether the role is a 'superuser' who can override all accessrestrictions within the database - superuser status is dangerous andshould be used only when really needed. You must yourself be asuperuser to create a new superuser. Defaults is 'false'.",
											MarkdownDescription: "Whether the role is a 'superuser' who can override all accessrestrictions within the database - superuser status is dangerous andshould be used only when really needed. You must yourself be asuperuser to create a new superuser. Defaults is 'false'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"valid_until": schema.StringAttribute{
											Description:         "Date and time after which the role's password is no longer valid.When omitted, the password will never expire (default).",
											MarkdownDescription: "Date and time after which the role's password is no longer valid.When omitted, the password will never expire (default).",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												validators.DateTime64Validator(),
											},
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"services": schema.SingleNestedAttribute{
								Description:         "Services roles managed by the 'Cluster'",
								MarkdownDescription: "Services roles managed by the 'Cluster'",
								Attributes: map[string]schema.Attribute{
									"additional": schema.ListNestedAttribute{
										Description:         "Additional is a list of additional managed services specified by the user.",
										MarkdownDescription: "Additional is a list of additional managed services specified by the user.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"selector_type": schema.StringAttribute{
													Description:         "SelectorType specifies the type of selectors that the service will have.Valid values are 'rw', 'r', and 'ro', representing read-write, read, and read-only services.",
													MarkdownDescription: "SelectorType specifies the type of selectors that the service will have.Valid values are 'rw', 'r', and 'ro', representing read-write, read, and read-only services.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"service_template": schema.SingleNestedAttribute{
													Description:         "ServiceTemplate is the template specification for the service.",
													MarkdownDescription: "ServiceTemplate is the template specification for the service.",
													Attributes: map[string]schema.Attribute{
														"metadata": schema.SingleNestedAttribute{
															Description:         "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
															MarkdownDescription: "Standard object's metadata.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
															Attributes: map[string]schema.Attribute{
																"annotations": schema.MapAttribute{
																	Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
																	MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"labels": schema.MapAttribute{
																	Description:         "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
																	MarkdownDescription: "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "The name of the resource. Only supported for certain types",
																	MarkdownDescription: "The name of the resource. Only supported for certain types",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"spec": schema.SingleNestedAttribute{
															Description:         "Specification of the desired behavior of the service.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
															MarkdownDescription: "Specification of the desired behavior of the service.More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status",
															Attributes: map[string]schema.Attribute{
																"allocate_load_balancer_node_ports": schema.BoolAttribute{
																	Description:         "allocateLoadBalancerNodePorts defines if NodePorts will be automaticallyallocated for services with type LoadBalancer.  Default is 'true'. Itmay be set to 'false' if the cluster load-balancer does not rely onNodePorts.  If the caller requests specific NodePorts (by specifying avalue), those requests will be respected, regardless of this field.This field may only be set for services with type LoadBalancer and willbe cleared if the type is changed to any other type.",
																	MarkdownDescription: "allocateLoadBalancerNodePorts defines if NodePorts will be automaticallyallocated for services with type LoadBalancer.  Default is 'true'. Itmay be set to 'false' if the cluster load-balancer does not rely onNodePorts.  If the caller requests specific NodePorts (by specifying avalue), those requests will be respected, regardless of this field.This field may only be set for services with type LoadBalancer and willbe cleared if the type is changed to any other type.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"cluster_ip": schema.StringAttribute{
																	Description:         "clusterIP is the IP address of the service and is usually assignedrandomly. If an address is specified manually, is in-range (as persystem configuration), and is not in use, it will be allocated to theservice; otherwise creation of the service will fail. This field may notbe changed through updates unless the type field is also being changedto ExternalName (which requires this field to be blank) or the typefield is being changed from ExternalName (in which case this field mayoptionally be specified, as describe above).  Valid values are 'None',empty string (''), or a valid IP address. Setting this to 'None' makes a'headless service' (no virtual IP), which is useful when direct endpointconnections are preferred and proxying is not required.  Only applies totypes ClusterIP, NodePort, and LoadBalancer. If this field is specifiedwhen creating a Service of type ExternalName, creation will fail. Thisfield will be wiped when updating a Service to type ExternalName.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
																	MarkdownDescription: "clusterIP is the IP address of the service and is usually assignedrandomly. If an address is specified manually, is in-range (as persystem configuration), and is not in use, it will be allocated to theservice; otherwise creation of the service will fail. This field may notbe changed through updates unless the type field is also being changedto ExternalName (which requires this field to be blank) or the typefield is being changed from ExternalName (in which case this field mayoptionally be specified, as describe above).  Valid values are 'None',empty string (''), or a valid IP address. Setting this to 'None' makes a'headless service' (no virtual IP), which is useful when direct endpointconnections are preferred and proxying is not required.  Only applies totypes ClusterIP, NodePort, and LoadBalancer. If this field is specifiedwhen creating a Service of type ExternalName, creation will fail. Thisfield will be wiped when updating a Service to type ExternalName.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"cluster_i_ps": schema.ListAttribute{
																	Description:         "ClusterIPs is a list of IP addresses assigned to this service, and areusually assigned randomly.  If an address is specified manually, isin-range (as per system configuration), and is not in use, it will beallocated to the service; otherwise creation of the service will fail.This field may not be changed through updates unless the type field isalso being changed to ExternalName (which requires this field to beempty) or the type field is being changed from ExternalName (in whichcase this field may optionally be specified, as describe above).  Validvalues are 'None', empty string (''), or a valid IP address.  Settingthis to 'None' makes a 'headless service' (no virtual IP), which isuseful when direct endpoint connections are preferred and proxying isnot required.  Only applies to types ClusterIP, NodePort, andLoadBalancer. If this field is specified when creating a Service of typeExternalName, creation will fail. This field will be wiped when updatinga Service to type ExternalName.  If this field is not specified, it willbe initialized from the clusterIP field.  If this field is specified,clients must ensure that clusterIPs[0] and clusterIP have the samevalue.This field may hold a maximum of two entries (dual-stack IPs, in either order).These IPs must correspond to the values of the ipFamilies field. BothclusterIPs and ipFamilies are governed by the ipFamilyPolicy field.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
																	MarkdownDescription: "ClusterIPs is a list of IP addresses assigned to this service, and areusually assigned randomly.  If an address is specified manually, isin-range (as per system configuration), and is not in use, it will beallocated to the service; otherwise creation of the service will fail.This field may not be changed through updates unless the type field isalso being changed to ExternalName (which requires this field to beempty) or the type field is being changed from ExternalName (in whichcase this field may optionally be specified, as describe above).  Validvalues are 'None', empty string (''), or a valid IP address.  Settingthis to 'None' makes a 'headless service' (no virtual IP), which isuseful when direct endpoint connections are preferred and proxying isnot required.  Only applies to types ClusterIP, NodePort, andLoadBalancer. If this field is specified when creating a Service of typeExternalName, creation will fail. This field will be wiped when updatinga Service to type ExternalName.  If this field is not specified, it willbe initialized from the clusterIP field.  If this field is specified,clients must ensure that clusterIPs[0] and clusterIP have the samevalue.This field may hold a maximum of two entries (dual-stack IPs, in either order).These IPs must correspond to the values of the ipFamilies field. BothclusterIPs and ipFamilies are governed by the ipFamilyPolicy field.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"external_i_ps": schema.ListAttribute{
																	Description:         "externalIPs is a list of IP addresses for which nodes in the clusterwill also accept traffic for this service.  These IPs are not managed byKubernetes.  The user is responsible for ensuring that traffic arrivesat a node with this IP.  A common example is external load-balancersthat are not part of the Kubernetes system.",
																	MarkdownDescription: "externalIPs is a list of IP addresses for which nodes in the clusterwill also accept traffic for this service.  These IPs are not managed byKubernetes.  The user is responsible for ensuring that traffic arrivesat a node with this IP.  A common example is external load-balancersthat are not part of the Kubernetes system.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"external_name": schema.StringAttribute{
																	Description:         "externalName is the external reference that discovery mechanisms willreturn as an alias for this service (e.g. a DNS CNAME record). Noproxying will be involved.  Must be a lowercase RFC-1123 hostname(https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
																	MarkdownDescription: "externalName is the external reference that discovery mechanisms willreturn as an alias for this service (e.g. a DNS CNAME record). Noproxying will be involved.  Must be a lowercase RFC-1123 hostname(https://tools.ietf.org/html/rfc1123) and requires 'type' to be 'ExternalName'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"external_traffic_policy": schema.StringAttribute{
																	Description:         "externalTrafficPolicy describes how nodes distribute service traffic theyreceive on one of the Service's 'externally-facing' addresses (NodePorts,ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configurethe service in a way that assumes that external load balancers will take careof balancing the service traffic between nodes, and so each node will delivertraffic only to the node-local endpoints of the service, without masqueradingthe client source IP. (Traffic mistakenly sent to a node with no endpoints willbe dropped.) The default value, 'Cluster', uses the standard behavior ofrouting to all endpoints evenly (possibly modified by topology and otherfeatures). Note that traffic sent to an External IP or LoadBalancer IP fromwithin the cluster will always get 'Cluster' semantics, but clients sending toa NodePort from within the cluster may need to take traffic policy into accountwhen picking a node.",
																	MarkdownDescription: "externalTrafficPolicy describes how nodes distribute service traffic theyreceive on one of the Service's 'externally-facing' addresses (NodePorts,ExternalIPs, and LoadBalancer IPs). If set to 'Local', the proxy will configurethe service in a way that assumes that external load balancers will take careof balancing the service traffic between nodes, and so each node will delivertraffic only to the node-local endpoints of the service, without masqueradingthe client source IP. (Traffic mistakenly sent to a node with no endpoints willbe dropped.) The default value, 'Cluster', uses the standard behavior ofrouting to all endpoints evenly (possibly modified by topology and otherfeatures). Note that traffic sent to an External IP or LoadBalancer IP fromwithin the cluster will always get 'Cluster' semantics, but clients sending toa NodePort from within the cluster may need to take traffic policy into accountwhen picking a node.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"health_check_node_port": schema.Int64Attribute{
																	Description:         "healthCheckNodePort specifies the healthcheck nodePort for the service.This only applies when type is set to LoadBalancer andexternalTrafficPolicy is set to Local. If a value is specified, isin-range, and is not in use, it will be used.  If not specified, a valuewill be automatically allocated.  External systems (e.g. load-balancers)can use this port to determine if a given node holds endpoints for thisservice or not.  If this field is specified when creating a Servicewhich does not need it, creation will fail. This field will be wipedwhen updating a Service to no longer need it (e.g. changing type).This field cannot be updated once set.",
																	MarkdownDescription: "healthCheckNodePort specifies the healthcheck nodePort for the service.This only applies when type is set to LoadBalancer andexternalTrafficPolicy is set to Local. If a value is specified, isin-range, and is not in use, it will be used.  If not specified, a valuewill be automatically allocated.  External systems (e.g. load-balancers)can use this port to determine if a given node holds endpoints for thisservice or not.  If this field is specified when creating a Servicewhich does not need it, creation will fail. This field will be wipedwhen updating a Service to no longer need it (e.g. changing type).This field cannot be updated once set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"internal_traffic_policy": schema.StringAttribute{
																	Description:         "InternalTrafficPolicy describes how nodes distribute service traffic theyreceive on the ClusterIP. If set to 'Local', the proxy will assume that podsonly want to talk to endpoints of the service on the same node as the pod,dropping the traffic if there are no local endpoints. The default value,'Cluster', uses the standard behavior of routing to all endpoints evenly(possibly modified by topology and other features).",
																	MarkdownDescription: "InternalTrafficPolicy describes how nodes distribute service traffic theyreceive on the ClusterIP. If set to 'Local', the proxy will assume that podsonly want to talk to endpoints of the service on the same node as the pod,dropping the traffic if there are no local endpoints. The default value,'Cluster', uses the standard behavior of routing to all endpoints evenly(possibly modified by topology and other features).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"ip_families": schema.ListAttribute{
																	Description:         "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to thisservice. This field is usually assigned automatically based on clusterconfiguration and the ipFamilyPolicy field. If this field is specifiedmanually, the requested family is available in the cluster,and ipFamilyPolicy allows it, it will be used; otherwise creation ofthe service will fail. This field is conditionally mutable: it allowsfor adding or removing a secondary IP family, but it does not allowchanging the primary IP family of the Service. Valid values are 'IPv4'and 'IPv6'.  This field only applies to Services of types ClusterIP,NodePort, and LoadBalancer, and does apply to 'headless' services.This field will be wiped when updating a Service to type ExternalName.This field may hold a maximum of two entries (dual-stack families, ineither order).  These families must correspond to the values of theclusterIPs field, if specified. Both clusterIPs and ipFamilies aregoverned by the ipFamilyPolicy field.",
																	MarkdownDescription: "IPFamilies is a list of IP families (e.g. IPv4, IPv6) assigned to thisservice. This field is usually assigned automatically based on clusterconfiguration and the ipFamilyPolicy field. If this field is specifiedmanually, the requested family is available in the cluster,and ipFamilyPolicy allows it, it will be used; otherwise creation ofthe service will fail. This field is conditionally mutable: it allowsfor adding or removing a secondary IP family, but it does not allowchanging the primary IP family of the Service. Valid values are 'IPv4'and 'IPv6'.  This field only applies to Services of types ClusterIP,NodePort, and LoadBalancer, and does apply to 'headless' services.This field will be wiped when updating a Service to type ExternalName.This field may hold a maximum of two entries (dual-stack families, ineither order).  These families must correspond to the values of theclusterIPs field, if specified. Both clusterIPs and ipFamilies aregoverned by the ipFamilyPolicy field.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"ip_family_policy": schema.StringAttribute{
																	Description:         "IPFamilyPolicy represents the dual-stack-ness requested or required bythis Service. If there is no value provided, then this field will be setto SingleStack. Services can be 'SingleStack' (a single IP family),'PreferDualStack' (two IP families on dual-stack configured clusters ora single IP family on single-stack clusters), or 'RequireDualStack'(two IP families on dual-stack configured clusters, otherwise fail). TheipFamilies and clusterIPs fields depend on the value of this field. Thisfield will be wiped when updating a service to type ExternalName.",
																	MarkdownDescription: "IPFamilyPolicy represents the dual-stack-ness requested or required bythis Service. If there is no value provided, then this field will be setto SingleStack. Services can be 'SingleStack' (a single IP family),'PreferDualStack' (two IP families on dual-stack configured clusters ora single IP family on single-stack clusters), or 'RequireDualStack'(two IP families on dual-stack configured clusters, otherwise fail). TheipFamilies and clusterIPs fields depend on the value of this field. Thisfield will be wiped when updating a service to type ExternalName.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"load_balancer_class": schema.StringAttribute{
																	Description:         "loadBalancerClass is the class of the load balancer implementation this Service belongs to.If specified, the value of this field must be a label-style identifier, with an optional prefix,e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users.This field can only be set when the Service type is 'LoadBalancer'. If not set, the default loadbalancer implementation is used, today this is typically done through the cloud provider integration,but should apply for any default implementation. If set, it is assumed that a load balancerimplementation is watching for Services with a matching class. Any default load balancerimplementation (e.g. cloud providers) should ignore Services that set this field.This field can only be set when creating or updating a Service to type 'LoadBalancer'.Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
																	MarkdownDescription: "loadBalancerClass is the class of the load balancer implementation this Service belongs to.If specified, the value of this field must be a label-style identifier, with an optional prefix,e.g. 'internal-vip' or 'example.com/internal-vip'. Unprefixed names are reserved for end-users.This field can only be set when the Service type is 'LoadBalancer'. If not set, the default loadbalancer implementation is used, today this is typically done through the cloud provider integration,but should apply for any default implementation. If set, it is assumed that a load balancerimplementation is watching for Services with a matching class. Any default load balancerimplementation (e.g. cloud providers) should ignore Services that set this field.This field can only be set when creating or updating a Service to type 'LoadBalancer'.Once set, it can not be changed. This field will be wiped when a service is updated to a non 'LoadBalancer' type.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"load_balancer_ip": schema.StringAttribute{
																	Description:         "Only applies to Service Type: LoadBalancer.This feature depends on whether the underlying cloud-provider supports specifyingthe loadBalancerIP when a load balancer is created.This field will be ignored if the cloud-provider does not support the feature.Deprecated: This field was under-specified and its meaning varies across implementations.Using it is non-portable and it may not support dual-stack.Users are encouraged to use implementation-specific annotations when available.",
																	MarkdownDescription: "Only applies to Service Type: LoadBalancer.This feature depends on whether the underlying cloud-provider supports specifyingthe loadBalancerIP when a load balancer is created.This field will be ignored if the cloud-provider does not support the feature.Deprecated: This field was under-specified and its meaning varies across implementations.Using it is non-portable and it may not support dual-stack.Users are encouraged to use implementation-specific annotations when available.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"load_balancer_source_ranges": schema.ListAttribute{
																	Description:         "If specified and supported by the platform, this will restrict traffic through the cloud-providerload-balancer will be restricted to the specified client IPs. This field will be ignored if thecloud-provider does not support the feature.'More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
																	MarkdownDescription: "If specified and supported by the platform, this will restrict traffic through the cloud-providerload-balancer will be restricted to the specified client IPs. This field will be ignored if thecloud-provider does not support the feature.'More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"ports": schema.ListNestedAttribute{
																	Description:         "The list of ports that are exposed by this service.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
																	MarkdownDescription: "The list of ports that are exposed by this service.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"app_protocol": schema.StringAttribute{
																				Description:         "The application protocol for this port.This is used as a hint for implementations to offer richer behavior for protocols that they understand.This field follows standard Kubernetes label syntax.Valid values are either:* Un-prefixed protocol names - reserved for IANA standard service names (as perRFC-6335 and https://www.iana.org/assignments/service-names).* Kubernetes-defined prefixed names:  * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-  * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455  * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455* Other protocols should use implementation-defined prefixed names such asmycompany.com/my-custom-protocol.",
																				MarkdownDescription: "The application protocol for this port.This is used as a hint for implementations to offer richer behavior for protocols that they understand.This field follows standard Kubernetes label syntax.Valid values are either:* Un-prefixed protocol names - reserved for IANA standard service names (as perRFC-6335 and https://www.iana.org/assignments/service-names).* Kubernetes-defined prefixed names:  * 'kubernetes.io/h2c' - HTTP/2 prior knowledge over cleartext as described in https://www.rfc-editor.org/rfc/rfc9113.html#name-starting-http-2-with-prior-  * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455  * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455* Other protocols should use implementation-defined prefixed names such asmycompany.com/my-custom-protocol.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"name": schema.StringAttribute{
																				Description:         "The name of this port within the service. This must be a DNS_LABEL.All ports within a ServiceSpec must have unique names. When consideringthe endpoints for a Service, this must match the 'name' field in theEndpointPort.Optional if only one ServicePort is defined on this service.",
																				MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL.All ports within a ServiceSpec must have unique names. When consideringthe endpoints for a Service, this must match the 'name' field in theEndpointPort.Optional if only one ServicePort is defined on this service.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"node_port": schema.Int64Attribute{
																				Description:         "The port on each node on which this service is exposed when type isNodePort or LoadBalancer.  Usually assigned by the system. If a value isspecified, in-range, and not in use it will be used, otherwise theoperation will fail.  If not specified, a port will be allocated if thisService requires one.  If this field is specified when creating aService which does not need it, creation will fail. This field will bewiped when updating a Service to no longer need it (e.g. changing typefrom NodePort to ClusterIP).More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
																				MarkdownDescription: "The port on each node on which this service is exposed when type isNodePort or LoadBalancer.  Usually assigned by the system. If a value isspecified, in-range, and not in use it will be used, otherwise theoperation will fail.  If not specified, a port will be allocated if thisService requires one.  If this field is specified when creating aService which does not need it, creation will fail. This field will bewiped when updating a Service to no longer need it (e.g. changing typefrom NodePort to ClusterIP).More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"port": schema.Int64Attribute{
																				Description:         "The port that will be exposed by this service.",
																				MarkdownDescription: "The port that will be exposed by this service.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"protocol": schema.StringAttribute{
																				Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'.Default is TCP.",
																				MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'.Default is TCP.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"target_port": schema.StringAttribute{
																				Description:         "Number or name of the port to access on the pods targeted by the service.Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.If this is a string, it will be looked up as a named port in thetarget Pod's container ports. If this is not specified, the valueof the 'port' field is used (an identity map).This field is ignored for services with clusterIP=None, and should beomitted or set equal to the 'port' field.More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
																				MarkdownDescription: "Number or name of the port to access on the pods targeted by the service.Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.If this is a string, it will be looked up as a named port in thetarget Pod's container ports. If this is not specified, the valueof the 'port' field is used (an identity map).This field is ignored for services with clusterIP=None, and should beomitted or set equal to the 'port' field.More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
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

																"publish_not_ready_addresses": schema.BoolAttribute{
																	Description:         "publishNotReadyAddresses indicates that any agent which deals with endpoints for thisService should disregard any indications of ready/not-ready.The primary use case for setting this field is for a StatefulSet's Headless Service topropagate SRV DNS records for its Pods for the purpose of peer discovery.The Kubernetes controllers that generate Endpoints and EndpointSlice resources forServices interpret this to mean that all endpoints are considered 'ready' even if thePods themselves are not. Agents which consume only Kubernetes generated endpointsthrough the Endpoints or EndpointSlice resources can safely assume this behavior.",
																	MarkdownDescription: "publishNotReadyAddresses indicates that any agent which deals with endpoints for thisService should disregard any indications of ready/not-ready.The primary use case for setting this field is for a StatefulSet's Headless Service topropagate SRV DNS records for its Pods for the purpose of peer discovery.The Kubernetes controllers that generate Endpoints and EndpointSlice resources forServices interpret this to mean that all endpoints are considered 'ready' even if thePods themselves are not. Agents which consume only Kubernetes generated endpointsthrough the Endpoints or EndpointSlice resources can safely assume this behavior.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"selector": schema.MapAttribute{
																	Description:         "Route service traffic to pods with label keys and values matching thisselector. If empty or not present, the service is assumed to have anexternal process managing its endpoints, which Kubernetes will notmodify. Only applies to types ClusterIP, NodePort, and LoadBalancer.Ignored if type is ExternalName.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
																	MarkdownDescription: "Route service traffic to pods with label keys and values matching thisselector. If empty or not present, the service is assumed to have anexternal process managing its endpoints, which Kubernetes will notmodify. Only applies to types ClusterIP, NodePort, and LoadBalancer.Ignored if type is ExternalName.More info: https://kubernetes.io/docs/concepts/services-networking/service/",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"session_affinity": schema.StringAttribute{
																	Description:         "Supports 'ClientIP' and 'None'. Used to maintain session affinity.Enable client IP based session affinity.Must be ClientIP or None.Defaults to None.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
																	MarkdownDescription: "Supports 'ClientIP' and 'None'. Used to maintain session affinity.Enable client IP based session affinity.Must be ClientIP or None.Defaults to None.More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"session_affinity_config": schema.SingleNestedAttribute{
																	Description:         "sessionAffinityConfig contains the configurations of session affinity.",
																	MarkdownDescription: "sessionAffinityConfig contains the configurations of session affinity.",
																	Attributes: map[string]schema.Attribute{
																		"client_ip": schema.SingleNestedAttribute{
																			Description:         "clientIP contains the configurations of Client IP based session affinity.",
																			MarkdownDescription: "clientIP contains the configurations of Client IP based session affinity.",
																			Attributes: map[string]schema.Attribute{
																				"timeout_seconds": schema.Int64Attribute{
																					Description:         "timeoutSeconds specifies the seconds of ClientIP type session sticky time.The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'.Default value is 10800(for 3 hours).",
																					MarkdownDescription: "timeoutSeconds specifies the seconds of ClientIP type session sticky time.The value must be >0 && <=86400(for 1 day) if ServiceAffinity == 'ClientIP'.Default value is 10800(for 3 hours).",
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

																"traffic_distribution": schema.StringAttribute{
																	Description:         "TrafficDistribution offers a way to express preferences for how traffic isdistributed to Service endpoints. Implementations can use this field as ahint, but are not required to guarantee strict adherence. If the field isnot set, the implementation will apply its default routing strategy. If setto 'PreferClose', implementations should prioritize endpoints that aretopologically close (e.g., same zone).This is an alpha field and requires enabling ServiceTrafficDistribution feature.",
																	MarkdownDescription: "TrafficDistribution offers a way to express preferences for how traffic isdistributed to Service endpoints. Implementations can use this field as ahint, but are not required to guarantee strict adherence. If the field isnot set, the implementation will apply its default routing strategy. If setto 'PreferClose', implementations should prioritize endpoints that aretopologically close (e.g., same zone).This is an alpha field and requires enabling ServiceTrafficDistribution feature.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type determines how the Service is exposed. Defaults to ClusterIP. Validoptions are ExternalName, ClusterIP, NodePort, and LoadBalancer.'ClusterIP' allocates a cluster-internal IP address for load-balancingto endpoints. Endpoints are determined by the selector or if that is notspecified, by manual construction of an Endpoints object orEndpointSlice objects. If clusterIP is 'None', no virtual IP isallocated and the endpoints are published as a set of endpoints ratherthan a virtual IP.'NodePort' builds on ClusterIP and allocates a port on every node whichroutes to the same endpoints as the clusterIP.'LoadBalancer' builds on NodePort and creates an external load-balancer(if supported in the current cloud) which routes to the same endpointsas the clusterIP.'ExternalName' aliases this service to the specified externalName.Several other fields do not apply to ExternalName services.More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
																	MarkdownDescription: "type determines how the Service is exposed. Defaults to ClusterIP. Validoptions are ExternalName, ClusterIP, NodePort, and LoadBalancer.'ClusterIP' allocates a cluster-internal IP address for load-balancingto endpoints. Endpoints are determined by the selector or if that is notspecified, by manual construction of an Endpoints object orEndpointSlice objects. If clusterIP is 'None', no virtual IP isallocated and the endpoints are published as a set of endpoints ratherthan a virtual IP.'NodePort' builds on ClusterIP and allocates a port on every node whichroutes to the same endpoints as the clusterIP.'LoadBalancer' builds on NodePort and creates an external load-balancer(if supported in the current cloud) which routes to the same endpointsas the clusterIP.'ExternalName' aliases this service to the specified externalName.Several other fields do not apply to ExternalName services.More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types",
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
													Required: true,
													Optional: false,
													Computed: false,
												},

												"update_strategy": schema.StringAttribute{
													Description:         "UpdateStrategy describes how the service differences should be reconciled",
													MarkdownDescription: "UpdateStrategy describes how the service differences should be reconciled",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("patch", "replace"),
													},
												},
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"disabled_default_services": schema.ListAttribute{
										Description:         "DisabledDefaultServices is a list of service types that are disabled by default.Valid values are 'r', and 'ro', representing read, and read-only services.",
										MarkdownDescription: "DisabledDefaultServices is a list of service types that are disabled by default.Valid values are 'r', and 'ro', representing read, and read-only services.",
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

					"max_sync_replicas": schema.Int64Attribute{
						Description:         "The target value for the synchronous replication quorum, that can bedecreased if the number of ready standbys is lower than this.Undefined or 0 disable synchronous replication.",
						MarkdownDescription: "The target value for the synchronous replication quorum, that can bedecreased if the number of ready standbys is lower than this.Undefined or 0 disable synchronous replication.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"min_sync_replicas": schema.Int64Attribute{
						Description:         "Minimum number of instances required in synchronous replication with theprimary. Undefined or 0 allow writes to complete when no standby isavailable.",
						MarkdownDescription: "Minimum number of instances required in synchronous replication with theprimary. Undefined or 0 allow writes to complete when no standby isavailable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.Int64{
							int64validator.AtLeast(0),
						},
					},

					"monitoring": schema.SingleNestedAttribute{
						Description:         "The configuration of the monitoring infrastructure of this cluster",
						MarkdownDescription: "The configuration of the monitoring infrastructure of this cluster",
						Attributes: map[string]schema.Attribute{
							"custom_queries_config_map": schema.ListNestedAttribute{
								Description:         "The list of config maps containing the custom queries",
								MarkdownDescription: "The list of config maps containing the custom queries",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key to select",
											MarkdownDescription: "The key to select",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.",
											MarkdownDescription: "Name of the referent.",
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

							"custom_queries_secret": schema.ListNestedAttribute{
								Description:         "The list of secrets containing the custom queries",
								MarkdownDescription: "The list of secrets containing the custom queries",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"key": schema.StringAttribute{
											Description:         "The key to select",
											MarkdownDescription: "The key to select",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Name of the referent.",
											MarkdownDescription: "Name of the referent.",
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

							"disable_default_queries": schema.BoolAttribute{
								Description:         "Whether the default queries should be injected.Set it to 'true' if you don't want to inject default queries into the cluster.Default: false.",
								MarkdownDescription: "Whether the default queries should be injected.Set it to 'true' if you don't want to inject default queries into the cluster.Default: false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enable_pod_monitor": schema.BoolAttribute{
								Description:         "Enable or disable the 'PodMonitor'",
								MarkdownDescription: "Enable or disable the 'PodMonitor'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_monitor_metric_relabelings": schema.ListNestedAttribute{
								Description:         "The list of metric relabelings for the 'PodMonitor'. Applied to samples before ingestion.",
								MarkdownDescription: "The list of metric relabelings for the 'PodMonitor'. Applied to samples before ingestion.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action to perform based on the regex matching.'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0.'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.Default: 'Replace'",
											MarkdownDescription: "Action to perform based on the regex matching.'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0.'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.Default: 'Replace'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
											},
										},

										"modulus": schema.Int64Attribute{
											Description:         "Modulus to take of the hash of the source label values.Only applicable when the action is 'HashMod'.",
											MarkdownDescription: "Modulus to take of the hash of the source label values.Only applicable when the action is 'HashMod'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex": schema.StringAttribute{
											Description:         "Regular expression against which the extracted value is matched.",
											MarkdownDescription: "Regular expression against which the extracted value is matched.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replacement": schema.StringAttribute{
											Description:         "Replacement value against which a Replace action is performed if theregular expression matches.Regex capture groups are available.",
											MarkdownDescription: "Replacement value against which a Replace action is performed if theregular expression matches.Regex capture groups are available.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"separator": schema.StringAttribute{
											Description:         "Separator is the string between concatenated SourceLabels.",
											MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source_labels": schema.ListAttribute{
											Description:         "The source labels select values from existing labels. Their content isconcatenated using the configured Separator and matched against theconfigured regular expression.",
											MarkdownDescription: "The source labels select values from existing labels. Their content isconcatenated using the configured Separator and matched against theconfigured regular expression.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"target_label": schema.StringAttribute{
											Description:         "Label to which the resulting string is written in a replacement.It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase','KeepEqual' and 'DropEqual' actions.Regex capture groups are available.",
											MarkdownDescription: "Label to which the resulting string is written in a replacement.It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase','KeepEqual' and 'DropEqual' actions.Regex capture groups are available.",
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

							"pod_monitor_relabelings": schema.ListNestedAttribute{
								Description:         "The list of relabelings for the 'PodMonitor'. Applied to samples before scraping.",
								MarkdownDescription: "The list of relabelings for the 'PodMonitor'. Applied to samples before scraping.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"action": schema.StringAttribute{
											Description:         "Action to perform based on the regex matching.'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0.'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.Default: 'Replace'",
											MarkdownDescription: "Action to perform based on the regex matching.'Uppercase' and 'Lowercase' actions require Prometheus >= v2.36.0.'DropEqual' and 'KeepEqual' actions require Prometheus >= v2.41.0.Default: 'Replace'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("replace", "Replace", "keep", "Keep", "drop", "Drop", "hashmod", "HashMod", "labelmap", "LabelMap", "labeldrop", "LabelDrop", "labelkeep", "LabelKeep", "lowercase", "Lowercase", "uppercase", "Uppercase", "keepequal", "KeepEqual", "dropequal", "DropEqual"),
											},
										},

										"modulus": schema.Int64Attribute{
											Description:         "Modulus to take of the hash of the source label values.Only applicable when the action is 'HashMod'.",
											MarkdownDescription: "Modulus to take of the hash of the source label values.Only applicable when the action is 'HashMod'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"regex": schema.StringAttribute{
											Description:         "Regular expression against which the extracted value is matched.",
											MarkdownDescription: "Regular expression against which the extracted value is matched.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replacement": schema.StringAttribute{
											Description:         "Replacement value against which a Replace action is performed if theregular expression matches.Regex capture groups are available.",
											MarkdownDescription: "Replacement value against which a Replace action is performed if theregular expression matches.Regex capture groups are available.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"separator": schema.StringAttribute{
											Description:         "Separator is the string between concatenated SourceLabels.",
											MarkdownDescription: "Separator is the string between concatenated SourceLabels.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"source_labels": schema.ListAttribute{
											Description:         "The source labels select values from existing labels. Their content isconcatenated using the configured Separator and matched against theconfigured regular expression.",
											MarkdownDescription: "The source labels select values from existing labels. Their content isconcatenated using the configured Separator and matched against theconfigured regular expression.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"target_label": schema.StringAttribute{
											Description:         "Label to which the resulting string is written in a replacement.It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase','KeepEqual' and 'DropEqual' actions.Regex capture groups are available.",
											MarkdownDescription: "Label to which the resulting string is written in a replacement.It is mandatory for 'Replace', 'HashMod', 'Lowercase', 'Uppercase','KeepEqual' and 'DropEqual' actions.Regex capture groups are available.",
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

							"tls": schema.SingleNestedAttribute{
								Description:         "Configure TLS communication for the metrics endpoint.Changing tls.enabled option will force a rollout of all instances.",
								MarkdownDescription: "Configure TLS communication for the metrics endpoint.Changing tls.enabled option will force a rollout of all instances.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "Enable TLS for the monitoring endpoint.Changing this option will force a rollout of all instances.",
										MarkdownDescription: "Enable TLS for the monitoring endpoint.Changing this option will force a rollout of all instances.",
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

					"node_maintenance_window": schema.SingleNestedAttribute{
						Description:         "Define a maintenance window for the Kubernetes nodes",
						MarkdownDescription: "Define a maintenance window for the Kubernetes nodes",
						Attributes: map[string]schema.Attribute{
							"in_progress": schema.BoolAttribute{
								Description:         "Is there a node maintenance activity in progress?",
								MarkdownDescription: "Is there a node maintenance activity in progress?",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"reuse_pvc": schema.BoolAttribute{
								Description:         "Reuse the existing PVC (wait for the node to comeup again) or not (recreate it elsewhere - when 'instances' >1)",
								MarkdownDescription: "Reuse the existing PVC (wait for the node to comeup again) or not (recreate it elsewhere - when 'instances' >1)",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"plugins": schema.ListNestedAttribute{
						Description:         "The plugins configuration, containingany plugin to be loaded with the corresponding configuration",
						MarkdownDescription: "The plugins configuration, containingany plugin to be loaded with the corresponding configuration",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"enabled": schema.BoolAttribute{
									Description:         "Enabled is true if this plugin will be used",
									MarkdownDescription: "Enabled is true if this plugin will be used",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Name is the plugin name",
									MarkdownDescription: "Name is the plugin name",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"parameters": schema.MapAttribute{
									Description:         "Parameters is the configuration of the plugin",
									MarkdownDescription: "Parameters is the configuration of the plugin",
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

					"postgres_gid": schema.Int64Attribute{
						Description:         "The GID of the 'postgres' user inside the image, defaults to '26'",
						MarkdownDescription: "The GID of the 'postgres' user inside the image, defaults to '26'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgres_uid": schema.Int64Attribute{
						Description:         "The UID of the 'postgres' user inside the image, defaults to '26'",
						MarkdownDescription: "The UID of the 'postgres' user inside the image, defaults to '26'",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"postgresql": schema.SingleNestedAttribute{
						Description:         "Configuration of the PostgreSQL server",
						MarkdownDescription: "Configuration of the PostgreSQL server",
						Attributes: map[string]schema.Attribute{
							"enable_alter_system": schema.BoolAttribute{
								Description:         "If this parameter is true, the user will be able to invoke 'ALTER SYSTEM'on this CloudNativePG Cluster.This should only be used for debugging and troubleshooting.Defaults to false.",
								MarkdownDescription: "If this parameter is true, the user will be able to invoke 'ALTER SYSTEM'on this CloudNativePG Cluster.This should only be used for debugging and troubleshooting.Defaults to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"ldap": schema.SingleNestedAttribute{
								Description:         "Options to specify LDAP configuration",
								MarkdownDescription: "Options to specify LDAP configuration",
								Attributes: map[string]schema.Attribute{
									"bind_as_auth": schema.SingleNestedAttribute{
										Description:         "Bind as authentication configuration",
										MarkdownDescription: "Bind as authentication configuration",
										Attributes: map[string]schema.Attribute{
											"prefix": schema.StringAttribute{
												Description:         "Prefix for the bind authentication option",
												MarkdownDescription: "Prefix for the bind authentication option",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"suffix": schema.StringAttribute{
												Description:         "Suffix for the bind authentication option",
												MarkdownDescription: "Suffix for the bind authentication option",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"bind_search_auth": schema.SingleNestedAttribute{
										Description:         "Bind+Search authentication configuration",
										MarkdownDescription: "Bind+Search authentication configuration",
										Attributes: map[string]schema.Attribute{
											"base_dn": schema.StringAttribute{
												Description:         "Root DN to begin the user search",
												MarkdownDescription: "Root DN to begin the user search",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"bind_dn": schema.StringAttribute{
												Description:         "DN of the user to bind to the directory",
												MarkdownDescription: "DN of the user to bind to the directory",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"bind_password": schema.SingleNestedAttribute{
												Description:         "Secret with the password for the user to bind to the directory",
												MarkdownDescription: "Secret with the password for the user to bind to the directory",
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key of the secret to select from.  Must be a valid secret key.",
														MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"optional": schema.BoolAttribute{
														Description:         "Specify whether the Secret or its key must be defined",
														MarkdownDescription: "Specify whether the Secret or its key must be defined",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"search_attribute": schema.StringAttribute{
												Description:         "Attribute to match against the username",
												MarkdownDescription: "Attribute to match against the username",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"search_filter": schema.StringAttribute{
												Description:         "Search filter to use when doing the search+bind authentication",
												MarkdownDescription: "Search filter to use when doing the search+bind authentication",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"port": schema.Int64Attribute{
										Description:         "LDAP server port",
										MarkdownDescription: "LDAP server port",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"scheme": schema.StringAttribute{
										Description:         "LDAP schema to be used, possible options are 'ldap' and 'ldaps'",
										MarkdownDescription: "LDAP schema to be used, possible options are 'ldap' and 'ldaps'",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("ldap", "ldaps"),
										},
									},

									"server": schema.StringAttribute{
										Description:         "LDAP hostname or IP address",
										MarkdownDescription: "LDAP hostname or IP address",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.BoolAttribute{
										Description:         "Set to 'true' to enable LDAP over TLS. 'false' is default",
										MarkdownDescription: "Set to 'true' to enable LDAP over TLS. 'false' is default",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"parameters": schema.MapAttribute{
								Description:         "PostgreSQL configuration options (postgresql.conf)",
								MarkdownDescription: "PostgreSQL configuration options (postgresql.conf)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pg_hba": schema.ListAttribute{
								Description:         "PostgreSQL Host Based Authentication rules (lines to be appendedto the pg_hba.conf file)",
								MarkdownDescription: "PostgreSQL Host Based Authentication rules (lines to be appendedto the pg_hba.conf file)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pg_ident": schema.ListAttribute{
								Description:         "PostgreSQL User Name Maps rules (lines to be appendedto the pg_ident.conf file)",
								MarkdownDescription: "PostgreSQL User Name Maps rules (lines to be appendedto the pg_ident.conf file)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"promotion_timeout": schema.Int64Attribute{
								Description:         "Specifies the maximum number of seconds to wait when promoting an instance to primary.Default value is 40000000, greater than one year in seconds,big enough to simulate an infinite timeout",
								MarkdownDescription: "Specifies the maximum number of seconds to wait when promoting an instance to primary.Default value is 40000000, greater than one year in seconds,big enough to simulate an infinite timeout",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"shared_preload_libraries": schema.ListAttribute{
								Description:         "Lists of shared preload libraries to add to the default ones",
								MarkdownDescription: "Lists of shared preload libraries to add to the default ones",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sync_replica_election_constraint": schema.SingleNestedAttribute{
								Description:         "Requirements to be met by sync replicas. This will affect how the 'synchronous_standby_names' parameter will beset up.",
								MarkdownDescription: "Requirements to be met by sync replicas. This will affect how the 'synchronous_standby_names' parameter will beset up.",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "This flag enables the constraints for sync replicas",
										MarkdownDescription: "This flag enables the constraints for sync replicas",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"node_labels_anti_affinity": schema.ListAttribute{
										Description:         "A list of node labels values to extract and compare to evaluate if the pods reside in the same topology or not",
										MarkdownDescription: "A list of node labels values to extract and compare to evaluate if the pods reside in the same topology or not",
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

							"synchronous": schema.SingleNestedAttribute{
								Description:         "Configuration of the PostgreSQL synchronous replication feature",
								MarkdownDescription: "Configuration of the PostgreSQL synchronous replication feature",
								Attributes: map[string]schema.Attribute{
									"max_standby_names_from_cluster": schema.Int64Attribute{
										Description:         "Specifies the maximum number of local cluster pods that can beautomatically included in the 'synchronous_standby_names' option inPostgreSQL.",
										MarkdownDescription: "Specifies the maximum number of local cluster pods that can beautomatically included in the 'synchronous_standby_names' option inPostgreSQL.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"method": schema.StringAttribute{
										Description:         "Method to select synchronous replication standbys from the listedservers, accepting 'any' (quorum-based synchronous replication) or'first' (priority-based synchronous replication) as values.",
										MarkdownDescription: "Method to select synchronous replication standbys from the listedservers, accepting 'any' (quorum-based synchronous replication) or'first' (priority-based synchronous replication) as values.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.OneOf("any", "first"),
										},
									},

									"number": schema.Int64Attribute{
										Description:         "Specifies the number of synchronous standby servers thattransactions must wait for responses from.",
										MarkdownDescription: "Specifies the number of synchronous standby servers thattransactions must wait for responses from.",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"standby_names_post": schema.ListAttribute{
										Description:         "A user-defined list of application names to be added to'synchronous_standby_names' after local cluster pods (the order isonly useful for priority-based synchronous replication).",
										MarkdownDescription: "A user-defined list of application names to be added to'synchronous_standby_names' after local cluster pods (the order isonly useful for priority-based synchronous replication).",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"standby_names_pre": schema.ListAttribute{
										Description:         "A user-defined list of application names to be added to'synchronous_standby_names' before local cluster pods (the order isonly useful for priority-based synchronous replication).",
										MarkdownDescription: "A user-defined list of application names to be added to'synchronous_standby_names' before local cluster pods (the order isonly useful for priority-based synchronous replication).",
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

					"primary_update_method": schema.StringAttribute{
						Description:         "Method to follow to upgrade the primary server during a rollingupdate procedure, after all replicas have been successfully updated:it can be with a switchover ('switchover') or in-place ('restart' - default)",
						MarkdownDescription: "Method to follow to upgrade the primary server during a rollingupdate procedure, after all replicas have been successfully updated:it can be with a switchover ('switchover') or in-place ('restart' - default)",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("switchover", "restart"),
						},
					},

					"primary_update_strategy": schema.StringAttribute{
						Description:         "Deployment strategy to follow to upgrade the primary server during a rollingupdate procedure, after all replicas have been successfully updated:it can be automated ('unsupervised' - default) or manual ('supervised')",
						MarkdownDescription: "Deployment strategy to follow to upgrade the primary server during a rollingupdate procedure, after all replicas have been successfully updated:it can be automated ('unsupervised' - default) or manual ('supervised')",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("unsupervised", "supervised"),
						},
					},

					"priority_class_name": schema.StringAttribute{
						Description:         "Name of the priority class which will be used in every generated Pod, if the PriorityClassspecified does not exist, the pod will not be able to schedule.  Please refer tohttps://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/#priorityclassfor more information",
						MarkdownDescription: "Name of the priority class which will be used in every generated Pod, if the PriorityClassspecified does not exist, the pod will not be able to schedule.  Please refer tohttps://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/#priorityclassfor more information",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"projected_volume_template": schema.SingleNestedAttribute{
						Description:         "Template to be used to define projected volumes, projected volumes will be mountedunder '/projected' base folder",
						MarkdownDescription: "Template to be used to define projected volumes, projected volumes will be mountedunder '/projected' base folder",
						Attributes: map[string]schema.Attribute{
							"default_mode": schema.Int64Attribute{
								Description:         "defaultMode are the mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
								MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"sources": schema.ListNestedAttribute{
								Description:         "sources is the list of volume projections",
								MarkdownDescription: "sources is the list of volume projections",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"cluster_trust_bundle": schema.SingleNestedAttribute{
											Description:         "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' fieldof ClusterTrustBundle objects in an auto-updating file.Alpha, gated by the ClusterTrustBundleProjection feature gate.ClusterTrustBundle objects can either be selected by name, or by thecombination of signer name and a label selector.Kubelet performs aggressive normalization of the PEM contents writteninto the pod filesystem.  Esoteric PEM features such as inter-blockcomments and block headers are stripped.  Certificates are deduplicated.The ordering of certificates within the file is arbitrary, and Kubeletmay change the order over time.",
											MarkdownDescription: "ClusterTrustBundle allows a pod to access the '.spec.trustBundle' fieldof ClusterTrustBundle objects in an auto-updating file.Alpha, gated by the ClusterTrustBundleProjection feature gate.ClusterTrustBundle objects can either be selected by name, or by thecombination of signer name and a label selector.Kubelet performs aggressive normalization of the PEM contents writteninto the pod filesystem.  Esoteric PEM features such as inter-blockcomments and block headers are stripped.  Certificates are deduplicated.The ordering of certificates within the file is arbitrary, and Kubeletmay change the order over time.",
											Attributes: map[string]schema.Attribute{
												"label_selector": schema.SingleNestedAttribute{
													Description:         "Select all ClusterTrustBundles that match this label selector.  Only haseffect if signerName is set.  Mutually-exclusive with name.  If unset,interpreted as 'match nothing'.  If set but empty, interpreted as 'matcheverything'.",
													MarkdownDescription: "Select all ClusterTrustBundles that match this label selector.  Only haseffect if signerName is set.  Mutually-exclusive with name.  If unset,interpreted as 'match nothing'.  If set but empty, interpreted as 'matcheverything'.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"name": schema.StringAttribute{
													Description:         "Select a single ClusterTrustBundle by object name.  Mutually-exclusivewith signerName and labelSelector.",
													MarkdownDescription: "Select a single ClusterTrustBundle by object name.  Mutually-exclusivewith signerName and labelSelector.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "If true, don't block pod startup if the referenced ClusterTrustBundle(s)aren't available.  If using name, then the named ClusterTrustBundle isallowed not to exist.  If using signerName, then the combination ofsignerName and labelSelector is allowed to match zeroClusterTrustBundles.",
													MarkdownDescription: "If true, don't block pod startup if the referenced ClusterTrustBundle(s)aren't available.  If using name, then the named ClusterTrustBundle isallowed not to exist.  If using signerName, then the combination ofsignerName and labelSelector is allowed to match zeroClusterTrustBundles.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "Relative path from the volume root to write the bundle.",
													MarkdownDescription: "Relative path from the volume root to write the bundle.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"signer_name": schema.StringAttribute{
													Description:         "Select all ClusterTrustBundles that match this signer name.Mutually-exclusive with name.  The contents of all selectedClusterTrustBundles will be unified and deduplicated.",
													MarkdownDescription: "Select all ClusterTrustBundles that match this signer name.Mutually-exclusive with name.  The contents of all selectedClusterTrustBundles will be unified and deduplicated.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"config_map": schema.SingleNestedAttribute{
											Description:         "configMap information about the configMap data to project",
											MarkdownDescription: "configMap information about the configMap data to project",
											Attributes: map[string]schema.Attribute{
												"items": schema.ListNestedAttribute{
													Description:         "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
													MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedConfigMap will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the ConfigMap,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"mode": schema.Int64Attribute{
																Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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

												"name": schema.StringAttribute{
													Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "optional specify whether the ConfigMap or its keys must be defined",
													MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"downward_api": schema.SingleNestedAttribute{
											Description:         "downwardAPI information about the downwardAPI data to project",
											MarkdownDescription: "downwardAPI information about the downwardAPI data to project",
											Attributes: map[string]schema.Attribute{
												"items": schema.ListNestedAttribute{
													Description:         "Items is a list of DownwardAPIVolume file",
													MarkdownDescription: "Items is a list of DownwardAPIVolume file",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"field_ref": schema.SingleNestedAttribute{
																Description:         "Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.",
																MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name, namespace and uid are supported.",
																Attributes: map[string]schema.Attribute{
																	"api_version": schema.StringAttribute{
																		Description:         "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		MarkdownDescription: "Version of the schema the FieldPath is written in terms of, defaults to 'v1'.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"field_path": schema.StringAttribute{
																		Description:         "Path of the field to select in the specified API version.",
																		MarkdownDescription: "Path of the field to select in the specified API version.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"mode": schema.Int64Attribute{
																Description:         "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal valuebetween 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"resource_field_ref": schema.SingleNestedAttribute{
																Description:         "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests(limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																Attributes: map[string]schema.Attribute{
																	"container_name": schema.StringAttribute{
																		Description:         "Container name: required for volumes, optional for env vars",
																		MarkdownDescription: "Container name: required for volumes, optional for env vars",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"divisor": schema.StringAttribute{
																		Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																		MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"resource": schema.StringAttribute{
																		Description:         "Required: resource to select",
																		MarkdownDescription: "Required: resource to select",
																		Required:            true,
																		Optional:            false,
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"secret": schema.SingleNestedAttribute{
											Description:         "secret information about the secret data to project",
											MarkdownDescription: "secret information about the secret data to project",
											Attributes: map[string]schema.Attribute{
												"items": schema.ListNestedAttribute{
													Description:         "items if unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
													MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"key": schema.StringAttribute{
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"mode": schema.Int64Attribute{
																Description:         "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.If not specified, the volume defaultMode will be used.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"path": schema.StringAttribute{
																Description:         "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to.May not be an absolute path.May not contain the path element '..'.May not start with the string '..'.",
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

												"name": schema.StringAttribute{
													Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"optional": schema.BoolAttribute{
													Description:         "optional field specify whether the Secret or its key must be defined",
													MarkdownDescription: "optional field specify whether the Secret or its key must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"service_account_token": schema.SingleNestedAttribute{
											Description:         "serviceAccountToken is information about the serviceAccountToken data to project",
											MarkdownDescription: "serviceAccountToken is information about the serviceAccountToken data to project",
											Attributes: map[string]schema.Attribute{
												"audience": schema.StringAttribute{
													Description:         "audience is the intended audience of the token. A recipient of a tokenmust identify itself with an identifier specified in the audience of thetoken, and otherwise should reject the token. The audience defaults to theidentifier of the apiserver.",
													MarkdownDescription: "audience is the intended audience of the token. A recipient of a tokenmust identify itself with an identifier specified in the audience of thetoken, and otherwise should reject the token. The audience defaults to theidentifier of the apiserver.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"expiration_seconds": schema.Int64Attribute{
													Description:         "expirationSeconds is the requested duration of validity of the serviceaccount token. As the token approaches expiration, the kubelet volumeplugin will proactively rotate the service account token. The kubelet willstart trying to rotate the token if the token is older than 80 percent ofits time to live or if the token is older than 24 hours.Defaults to 1 hourand must be at least 10 minutes.",
													MarkdownDescription: "expirationSeconds is the requested duration of validity of the serviceaccount token. As the token approaches expiration, the kubelet volumeplugin will proactively rotate the service account token. The kubelet willstart trying to rotate the token if the token is older than 80 percent ofits time to live or if the token is older than 24 hours.Defaults to 1 hourand must be at least 10 minutes.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "path is the path relative to the mount point of the file to project thetoken into.",
													MarkdownDescription: "path is the path relative to the mount point of the file to project thetoken into.",
													Required:            true,
													Optional:            false,
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replica": schema.SingleNestedAttribute{
						Description:         "Replica cluster configuration",
						MarkdownDescription: "Replica cluster configuration",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "If replica mode is enabled, this cluster will be a replica of anexisting cluster. Replica cluster can be created from a recoveryobject store or via streaming through pg_basebackup.Refer to the Replica clusters page of the documentation for more information.",
								MarkdownDescription: "If replica mode is enabled, this cluster will be a replica of anexisting cluster. Replica cluster can be created from a recoveryobject store or via streaming through pg_basebackup.Refer to the Replica clusters page of the documentation for more information.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"min_apply_delay": schema.StringAttribute{
								Description:         "When replica mode is enabled, this parameter allows you to replaytransactions only when the system time is at least the configuredtime past the commit time. This provides an opportunity to correctdata loss errors. Note that when this parameter is set, a promotiontoken cannot be used.",
								MarkdownDescription: "When replica mode is enabled, this parameter allows you to replaytransactions only when the system time is at least the configuredtime past the commit time. This provides an opportunity to correctdata loss errors. Note that when this parameter is set, a promotiontoken cannot be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"primary": schema.StringAttribute{
								Description:         "Primary defines which Cluster is defined to be the primary in the distributed PostgreSQL cluster, based on thetopology specified in externalClusters",
								MarkdownDescription: "Primary defines which Cluster is defined to be the primary in the distributed PostgreSQL cluster, based on thetopology specified in externalClusters",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"promotion_token": schema.StringAttribute{
								Description:         "A demotion token generated by an external cluster used tocheck if the promotion requirements are met.",
								MarkdownDescription: "A demotion token generated by an external cluster used tocheck if the promotion requirements are met.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"self": schema.StringAttribute{
								Description:         "Self defines the name of this cluster. It is used to determine if this is a primaryor a replica cluster, comparing it with 'primary'",
								MarkdownDescription: "Self defines the name of this cluster. It is used to determine if this is a primaryor a replica cluster, comparing it with 'primary'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"source": schema.StringAttribute{
								Description:         "The name of the external cluster which is the replication origin",
								MarkdownDescription: "The name of the external cluster which is the replication origin",
								Required:            true,
								Optional:            false,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.LengthAtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replication_slots": schema.SingleNestedAttribute{
						Description:         "Replication slots management configuration",
						MarkdownDescription: "Replication slots management configuration",
						Attributes: map[string]schema.Attribute{
							"high_availability": schema.SingleNestedAttribute{
								Description:         "Replication slots for high availability configuration",
								MarkdownDescription: "Replication slots for high availability configuration",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "If enabled (default), the operator will automatically manage replication slotson the primary instance and use them in streaming replicationconnections with all the standby instances that are part of the HAcluster. If disabled, the operator will not take advantageof replication slots in streaming connections with the replicas.This feature also controls replication slots in replica cluster,from the designated primary to its cascading replicas.",
										MarkdownDescription: "If enabled (default), the operator will automatically manage replication slotson the primary instance and use them in streaming replicationconnections with all the standby instances that are part of the HAcluster. If disabled, the operator will not take advantageof replication slots in streaming connections with the replicas.This feature also controls replication slots in replica cluster,from the designated primary to its cascading replicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"slot_prefix": schema.StringAttribute{
										Description:         "Prefix for replication slots managed by the operator for HA.It may only contain lower case letters, numbers, and the underscore character.This can only be set at creation time. By default set to '_cnpg_'.",
										MarkdownDescription: "Prefix for replication slots managed by the operator for HA.It may only contain lower case letters, numbers, and the underscore character.This can only be set at creation time. By default set to '_cnpg_'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9a-z_]*$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"synchronize_replicas": schema.SingleNestedAttribute{
								Description:         "Configures the synchronization of the user defined physical replication slots",
								MarkdownDescription: "Configures the synchronization of the user defined physical replication slots",
								Attributes: map[string]schema.Attribute{
									"enabled": schema.BoolAttribute{
										Description:         "When set to true, every replication slot that is on the primary is synchronized on each standby",
										MarkdownDescription: "When set to true, every replication slot that is on the primary is synchronized on each standby",
										Required:            true,
										Optional:            false,
										Computed:            false,
									},

									"exclude_patterns": schema.ListAttribute{
										Description:         "List of regular expression patterns to match the names of replication slots to be excluded (by default empty)",
										MarkdownDescription: "List of regular expression patterns to match the names of replication slots to be excluded (by default empty)",
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

							"update_interval": schema.Int64Attribute{
								Description:         "Standby will update the status of the local replication slotsevery 'updateInterval' seconds (default 30).",
								MarkdownDescription: "Standby will update the status of the local replication slotsevery 'updateInterval' seconds (default 30).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Resources requirements of every generated Pod. Please refer tohttps://kubernetes.io/docs/concepts/configuration/manage-resources-containers/for more information.",
						MarkdownDescription: "Resources requirements of every generated Pod. Please refer tohttps://kubernetes.io/docs/concepts/configuration/manage-resources-containers/for more information.",
						Attributes: map[string]schema.Attribute{
							"claims": schema.ListNestedAttribute{
								Description:         "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims,that are used by this container.This is an alpha field and requires enabling theDynamicResourceAllocation feature gate.This field is immutable. It can only be set for containers.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
											MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims ofthe Pod where this field is used. It makes that resource availableinside a container.",
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

							"limits": schema.MapAttribute{
								Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"requests": schema.MapAttribute{
								Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

					"scheduler_name": schema.StringAttribute{
						Description:         "If specified, the pod will be dispatched by specified Kubernetesscheduler. If not specified, the pod will be dispatched by the defaultscheduler. More info:https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/",
						MarkdownDescription: "If specified, the pod will be dispatched by specified Kubernetesscheduler. If not specified, the pod will be dispatched by the defaultscheduler. More info:https://kubernetes.io/docs/concepts/scheduling-eviction/kube-scheduler/",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"seccomp_profile": schema.SingleNestedAttribute{
						Description:         "The SeccompProfile applied to every Pod and Container.Defaults to: 'RuntimeDefault'",
						MarkdownDescription: "The SeccompProfile applied to every Pod and Container.Defaults to: 'RuntimeDefault'",
						Attributes: map[string]schema.Attribute{
							"localhost_profile": schema.StringAttribute{
								Description:         "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
								MarkdownDescription: "localhostProfile indicates a profile defined in a file on the node should be used.The profile must be preconfigured on the node to work.Must be a descending path, relative to the kubelet's configured seccomp profile location.Must be set if type is 'Localhost'. Must NOT be set for any other type.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"type": schema.StringAttribute{
								Description:         "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
								MarkdownDescription: "type indicates which kind of seccomp profile will be applied.Valid options are:Localhost - a profile defined in a file on the node should be used.RuntimeDefault - the container runtime default profile should be used.Unconfined - no profile should be applied.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"service_account_template": schema.SingleNestedAttribute{
						Description:         "Configure the generation of the service account",
						MarkdownDescription: "Configure the generation of the service account",
						Attributes: map[string]schema.Attribute{
							"metadata": schema.SingleNestedAttribute{
								Description:         "Metadata are the metadata to be used for the generatedservice account",
								MarkdownDescription: "Metadata are the metadata to be used for the generatedservice account",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
										MarkdownDescription: "Annotations is an unstructured key value map stored with a resource that may beset by external tools to store and retrieve arbitrary metadata. They are notqueryable and should be preserved when modifying objects.More info: http://kubernetes.io/docs/user-guide/annotations",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"labels": schema.MapAttribute{
										Description:         "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
										MarkdownDescription: "Map of string keys and values that can be used to organize and categorize(scope and select) objects. May match selectors of replication controllersand services.More info: http://kubernetes.io/docs/user-guide/labels",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"name": schema.StringAttribute{
										Description:         "The name of the resource. Only supported for certain types",
										MarkdownDescription: "The name of the resource. Only supported for certain types",
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
						Required: false,
						Optional: true,
						Computed: false,
					},

					"smart_shutdown_timeout": schema.Int64Attribute{
						Description:         "The time in seconds that controls the window of time reserved for the smart shutdown of Postgres to complete.Make sure you reserve enough time for the operator to request a fast shutdown of Postgres(that is: 'stopDelay' - 'smartShutdownTimeout').",
						MarkdownDescription: "The time in seconds that controls the window of time reserved for the smart shutdown of Postgres to complete.Make sure you reserve enough time for the operator to request a fast shutdown of Postgres(that is: 'stopDelay' - 'smartShutdownTimeout').",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"start_delay": schema.Int64Attribute{
						Description:         "The time in seconds that is allowed for a PostgreSQL instance tosuccessfully start up (default 3600).The startup probe failure threshold is derived from this value using the formula:ceiling(startDelay / 10).",
						MarkdownDescription: "The time in seconds that is allowed for a PostgreSQL instance tosuccessfully start up (default 3600).The startup probe failure threshold is derived from this value using the formula:ceiling(startDelay / 10).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"stop_delay": schema.Int64Attribute{
						Description:         "The time in seconds that is allowed for a PostgreSQL instance togracefully shutdown (default 1800)",
						MarkdownDescription: "The time in seconds that is allowed for a PostgreSQL instance togracefully shutdown (default 1800)",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"storage": schema.SingleNestedAttribute{
						Description:         "Configuration of the storage of the instances",
						MarkdownDescription: "Configuration of the storage of the instances",
						Attributes: map[string]schema.Attribute{
							"pvc_template": schema.SingleNestedAttribute{
								Description:         "Template to be used to generate the Persistent Volume Claim",
								MarkdownDescription: "Template to be used to generate the Persistent Volume Claim",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"data_source": schema.SingleNestedAttribute{
										Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
										MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the type of resource being referenced",
												MarkdownDescription: "Kind is the type of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of resource being referenced",
												MarkdownDescription: "Name is the name of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"data_source_ref": schema.SingleNestedAttribute{
										Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
										MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the type of resource being referenced",
												MarkdownDescription: "Kind is the type of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of resource being referenced",
												MarkdownDescription: "Name is the name of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
												MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

									"selector": schema.SingleNestedAttribute{
										Description:         "selector is a label query over volumes to consider for binding.",
										MarkdownDescription: "selector is a label query over volumes to consider for binding.",
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

									"storage_class_name": schema.StringAttribute{
										Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_attributes_class_name": schema.StringAttribute{
										Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
										MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_mode": schema.StringAttribute{
										Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
										MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_name": schema.StringAttribute{
										Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
										MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"resize_in_use_volumes": schema.BoolAttribute{
								Description:         "Resize existent PVCs, defaults to true",
								MarkdownDescription: "Resize existent PVCs, defaults to true",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "Size of the storage. Required if not already specified in the PVC template.Changes to this field are automatically reapplied to the created PVCs.Size cannot be decreased.",
								MarkdownDescription: "Size of the storage. Required if not already specified in the PVC template.Changes to this field are automatically reapplied to the created PVCs.Size cannot be decreased.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_class": schema.StringAttribute{
								Description:         "StorageClass to use for PVCs. Applied afterevaluating the PVC template, if available.If not specified, the generated PVCs will use thedefault storage class",
								MarkdownDescription: "StorageClass to use for PVCs. Applied afterevaluating the PVC template, if available.If not specified, the generated PVCs will use thedefault storage class",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"superuser_secret": schema.SingleNestedAttribute{
						Description:         "The secret containing the superuser password. If not defined a newsecret will be created with a randomly generated password",
						MarkdownDescription: "The secret containing the superuser password. If not defined a newsecret will be created with a randomly generated password",
						Attributes: map[string]schema.Attribute{
							"name": schema.StringAttribute{
								Description:         "Name of the referent.",
								MarkdownDescription: "Name of the referent.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"switchover_delay": schema.Int64Attribute{
						Description:         "The time in seconds that is allowed for a primary PostgreSQL instanceto gracefully shutdown during a switchover.Default value is 3600 seconds (1 hour).",
						MarkdownDescription: "The time in seconds that is allowed for a primary PostgreSQL instanceto gracefully shutdown during a switchover.Default value is 3600 seconds (1 hour).",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"tablespaces": schema.ListNestedAttribute{
						Description:         "The tablespaces configuration",
						MarkdownDescription: "The tablespaces configuration",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "The name of the tablespace",
									MarkdownDescription: "The name of the tablespace",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"owner": schema.SingleNestedAttribute{
									Description:         "Owner is the PostgreSQL user owning the tablespace",
									MarkdownDescription: "Owner is the PostgreSQL user owning the tablespace",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
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

								"storage": schema.SingleNestedAttribute{
									Description:         "The storage configuration for the tablespace",
									MarkdownDescription: "The storage configuration for the tablespace",
									Attributes: map[string]schema.Attribute{
										"pvc_template": schema.SingleNestedAttribute{
											Description:         "Template to be used to generate the Persistent Volume Claim",
											MarkdownDescription: "Template to be used to generate the Persistent Volume Claim",
											Attributes: map[string]schema.Attribute{
												"access_modes": schema.ListAttribute{
													Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"data_source": schema.SingleNestedAttribute{
													Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
													MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
													Attributes: map[string]schema.Attribute{
														"api_group": schema.StringAttribute{
															Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
															MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"kind": schema.StringAttribute{
															Description:         "Kind is the type of resource being referenced",
															MarkdownDescription: "Kind is the type of resource being referenced",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name of resource being referenced",
															MarkdownDescription: "Name is the name of resource being referenced",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"data_source_ref": schema.SingleNestedAttribute{
													Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
													MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
													Attributes: map[string]schema.Attribute{
														"api_group": schema.StringAttribute{
															Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
															MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"kind": schema.StringAttribute{
															Description:         "Kind is the type of resource being referenced",
															MarkdownDescription: "Kind is the type of resource being referenced",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name is the name of resource being referenced",
															MarkdownDescription: "Name is the name of resource being referenced",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"namespace": schema.StringAttribute{
															Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
															MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"resources": schema.SingleNestedAttribute{
													Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
													MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
													Attributes: map[string]schema.Attribute{
														"limits": schema.MapAttribute{
															Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"requests": schema.MapAttribute{
															Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
															MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

												"selector": schema.SingleNestedAttribute{
													Description:         "selector is a label query over volumes to consider for binding.",
													MarkdownDescription: "selector is a label query over volumes to consider for binding.",
													Attributes: map[string]schema.Attribute{
														"match_expressions": schema.ListNestedAttribute{
															Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "key is the label key that the selector applies to.",
																		MarkdownDescription: "key is the label key that the selector applies to.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"operator": schema.StringAttribute{
																		Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"values": schema.ListAttribute{
																		Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
																		MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
															Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
															MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

												"storage_class_name": schema.StringAttribute{
													Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
													MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_attributes_class_name": schema.StringAttribute{
													Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
													MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_mode": schema.StringAttribute{
													Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
													MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_name": schema.StringAttribute{
													Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
													MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"resize_in_use_volumes": schema.BoolAttribute{
											Description:         "Resize existent PVCs, defaults to true",
											MarkdownDescription: "Resize existent PVCs, defaults to true",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"size": schema.StringAttribute{
											Description:         "Size of the storage. Required if not already specified in the PVC template.Changes to this field are automatically reapplied to the created PVCs.Size cannot be decreased.",
											MarkdownDescription: "Size of the storage. Required if not already specified in the PVC template.Changes to this field are automatically reapplied to the created PVCs.Size cannot be decreased.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"storage_class": schema.StringAttribute{
											Description:         "StorageClass to use for PVCs. Applied afterevaluating the PVC template, if available.If not specified, the generated PVCs will use thedefault storage class",
											MarkdownDescription: "StorageClass to use for PVCs. Applied afterevaluating the PVC template, if available.If not specified, the generated PVCs will use thedefault storage class",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"temporary": schema.BoolAttribute{
									Description:         "When set to true, the tablespace will be added as a 'temp_tablespaces'entry in PostgreSQL, and will be available to automatically house tempdatabase objects, or other temporary files. Please refer to PostgreSQLdocumentation for more information on the 'temp_tablespaces' GUC.",
									MarkdownDescription: "When set to true, the tablespace will be added as a 'temp_tablespaces'entry in PostgreSQL, and will be available to automatically house tempdatabase objects, or other temporary files. Please refer to PostgreSQLdocumentation for more information on the 'temp_tablespaces' GUC.",
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

					"topology_spread_constraints": schema.ListNestedAttribute{
						Description:         "TopologySpreadConstraints specifies how to spread matching pods among the given topology.More info:https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/",
						MarkdownDescription: "TopologySpreadConstraints specifies how to spread matching pods among the given topology.More info:https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"label_selector": schema.SingleNestedAttribute{
									Description:         "LabelSelector is used to find matching pods.Pods that match this label selector are counted to determine the number of podsin their corresponding topology domain.",
									MarkdownDescription: "LabelSelector is used to find matching pods.Pods that match this label selector are counted to determine the number of podsin their corresponding topology domain.",
									Attributes: map[string]schema.Attribute{
										"match_expressions": schema.ListNestedAttribute{
											Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "key is the label key that the selector applies to.",
														MarkdownDescription: "key is the label key that the selector applies to.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
														MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"values": schema.ListAttribute{
														Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
														MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
											Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
											MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

								"match_label_keys": schema.ListAttribute{
									Description:         "MatchLabelKeys is a set of pod label keys to select the pods over whichspreading will be calculated. The keys are used to lookup values from theincoming pod labels, those key-value labels are ANDed with labelSelectorto select the group of existing pods over which spreading will be calculatedfor the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.MatchLabelKeys cannot be set when LabelSelector isn't set.Keys that don't exist in the incoming pod labels willbe ignored. A null or empty list means only match against labelSelector.This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
									MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over whichspreading will be calculated. The keys are used to lookup values from theincoming pod labels, those key-value labels are ANDed with labelSelectorto select the group of existing pods over which spreading will be calculatedfor the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector.MatchLabelKeys cannot be set when LabelSelector isn't set.Keys that don't exist in the incoming pod labels willbe ignored. A null or empty list means only match against labelSelector.This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"max_skew": schema.Int64Attribute{
									Description:         "MaxSkew describes the degree to which pods may be unevenly distributed.When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted differencebetween the number of matching pods in the target topology and the global minimum.The global minimum is the minimum number of matching pods in an eligible domainor zero if the number of eligible domains is less than MinDomains.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 2/2/1:In this case, the global minimum is 1.| zone1 | zone2 | zone3 ||  P P  |  P P  |   P   |- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)violate MaxSkew(1).- if MaxSkew is 2, incoming pod can be scheduled onto any zone.When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedenceto topologies that satisfy it.It's a required field. Default value is 1 and 0 is not allowed.",
									MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed.When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted differencebetween the number of matching pods in the target topology and the global minimum.The global minimum is the minimum number of matching pods in an eligible domainor zero if the number of eligible domains is less than MinDomains.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 2/2/1:In this case, the global minimum is 1.| zone1 | zone2 | zone3 ||  P P  |  P P  |   P   |- if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2;scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2)violate MaxSkew(1).- if MaxSkew is 2, incoming pod can be scheduled onto any zone.When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedenceto topologies that satisfy it.It's a required field. Default value is 1 and 0 is not allowed.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"min_domains": schema.Int64Attribute{
									Description:         "MinDomains indicates a minimum number of eligible domains.When the number of eligible domains with matching topology keys is less than minDomains,Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed.And when the number of eligible domains with matching topology keys equals or greater than minDomains,this value has no effect on scheduling.As a result, when the number of eligible domains is less than minDomains,scheduler won't schedule more than maxSkew Pods to those domains.If value is nil, the constraint behaves as if MinDomains is equal to 1.Valid values are integers greater than 0.When value is not nil, WhenUnsatisfiable must be DoNotSchedule.For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the samelabelSelector spread as 2/2/2:| zone1 | zone2 | zone3 ||  P P  |  P P  |  P P  |The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0.In this situation, new pod with the same labelSelector cannot be scheduled,because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,it will violate MaxSkew.",
									MarkdownDescription: "MinDomains indicates a minimum number of eligible domains.When the number of eligible domains with matching topology keys is less than minDomains,Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed.And when the number of eligible domains with matching topology keys equals or greater than minDomains,this value has no effect on scheduling.As a result, when the number of eligible domains is less than minDomains,scheduler won't schedule more than maxSkew Pods to those domains.If value is nil, the constraint behaves as if MinDomains is equal to 1.Valid values are integers greater than 0.When value is not nil, WhenUnsatisfiable must be DoNotSchedule.For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the samelabelSelector spread as 2/2/2:| zone1 | zone2 | zone3 ||  P P  |  P P  |  P P  |The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0.In this situation, new pod with the same labelSelector cannot be scheduled,because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones,it will violate MaxSkew.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"node_affinity_policy": schema.StringAttribute{
									Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelectorwhen calculating pod topology spread skew. Options are:- Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.- Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.If this value is nil, the behavior is equivalent to the Honor policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
									MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelectorwhen calculating pod topology spread skew. Options are:- Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations.- Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.If this value is nil, the behavior is equivalent to the Honor policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"node_taints_policy": schema.StringAttribute{
									Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculatingpod topology spread skew. Options are:- Honor: nodes without taints, along with tainted nodes for which the incoming podhas a toleration, are included.- Ignore: node taints are ignored. All nodes are included.If this value is nil, the behavior is equivalent to the Ignore policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
									MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculatingpod topology spread skew. Options are:- Honor: nodes without taints, along with tainted nodes for which the incoming podhas a toleration, are included.- Ignore: node taints are ignored. All nodes are included.If this value is nil, the behavior is equivalent to the Ignore policy.This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"topology_key": schema.StringAttribute{
									Description:         "TopologyKey is the key of node labels. Nodes that have a label with this keyand identical values are considered to be in the same topology.We consider each <key, value> as a 'bucket', and try to put balanced numberof pods into each bucket.We define a domain as a particular instance of a topology.Also, we define an eligible domain as a domain whose nodes meet the requirements ofnodeAffinityPolicy and nodeTaintsPolicy.e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology.And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology.It's a required field.",
									MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this keyand identical values are considered to be in the same topology.We consider each <key, value> as a 'bucket', and try to put balanced numberof pods into each bucket.We define a domain as a particular instance of a topology.Also, we define an eligible domain as a domain whose nodes meet the requirements ofnodeAffinityPolicy and nodeTaintsPolicy.e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology.And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology.It's a required field.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"when_unsatisfiable": schema.StringAttribute{
									Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfythe spread constraint.- DoNotSchedule (default) tells the scheduler not to schedule it.- ScheduleAnyway tells the scheduler to schedule the pod in any location,  but giving higher precedence to topologies that would help reduce the  skew.A constraint is considered 'Unsatisfiable' for an incoming podif and only if every possible node assignment for that pod would violate'MaxSkew' on some topology.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 3/1/1:| zone1 | zone2 | zone3 || P P P |   P   |   P   |If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduledto zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfiesMaxSkew(1). In other words, the cluster can still be imbalanced, but schedulerwon't make it *more* imbalanced.It's a required field.",
									MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfythe spread constraint.- DoNotSchedule (default) tells the scheduler not to schedule it.- ScheduleAnyway tells the scheduler to schedule the pod in any location,  but giving higher precedence to topologies that would help reduce the  skew.A constraint is considered 'Unsatisfiable' for an incoming podif and only if every possible node assignment for that pod would violate'MaxSkew' on some topology.For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the samelabelSelector spread as 3/1/1:| zone1 | zone2 | zone3 || P P P |   P   |   P   |If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduledto zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfiesMaxSkew(1). In other words, the cluster can still be imbalanced, but schedulerwon't make it *more* imbalanced.It's a required field.",
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

					"wal_storage": schema.SingleNestedAttribute{
						Description:         "Configuration of the storage for PostgreSQL WAL (Write-Ahead Log)",
						MarkdownDescription: "Configuration of the storage for PostgreSQL WAL (Write-Ahead Log)",
						Attributes: map[string]schema.Attribute{
							"pvc_template": schema.SingleNestedAttribute{
								Description:         "Template to be used to generate the Persistent Volume Claim",
								MarkdownDescription: "Template to be used to generate the Persistent Volume Claim",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										MarkdownDescription: "accessModes contains the desired access modes the volume should have.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"data_source": schema.SingleNestedAttribute{
										Description:         "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
										MarkdownDescription: "dataSource field can be used to specify either:* An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot)* An existing PVC (PersistentVolumeClaim)If the provisioner or an external controller can support the specified data source,it will create a new volume based on the contents of the specified data source.When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef,and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified.If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the type of resource being referenced",
												MarkdownDescription: "Kind is the type of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of resource being referenced",
												MarkdownDescription: "Name is the name of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"data_source_ref": schema.SingleNestedAttribute{
										Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
										MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-emptyvolume is desired. This may be any object from a non-empty API group (noncore object) or a PersistentVolumeClaim object.When this field is specified, volume binding will only succeed if the type ofthe specified object matches some installed volume populator or dynamicprovisioner.This field will replace the functionality of the dataSource field and as suchif both fields are non-empty, they must have the same value. For backwardscompatibility, when namespace isn't specified in dataSourceRef,both fields (dataSource and dataSourceRef) will be set to the samevalue automatically if one of them is empty and the other is non-empty.When namespace is specified in dataSourceRef,dataSource isn't set to the same value and must be empty.There are three important differences between dataSource and dataSourceRef:* While dataSource only allows two specific types of objects, dataSourceRef  allows any non-core object, as well as PersistentVolumeClaim objects.* While dataSource ignores disallowed values (dropping them), dataSourceRef  preserves all values, and generates an error if a disallowed value is  specified.* While dataSource only allows local objects, dataSourceRef allows objects  in any namespaces.(Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.(Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
										Attributes: map[string]schema.Attribute{
											"api_group": schema.StringAttribute{
												Description:         "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
												MarkdownDescription: "APIGroup is the group for the resource being referenced.If APIGroup is not specified, the specified Kind must be in the core API group.For any other third-party types, APIGroup is required.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind is the type of resource being referenced",
												MarkdownDescription: "Kind is the type of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name is the name of resource being referenced",
												MarkdownDescription: "Name is the name of resource being referenced",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
												MarkdownDescription: "Namespace is the namespace of resource being referencedNote that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details.(Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"resources": schema.SingleNestedAttribute{
										Description:         "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										MarkdownDescription: "resources represents the minimum resources the volume should have.If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirementsthat are lower than previous value but must still be higher than capacity recorded in thestatus field of the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
										Attributes: map[string]schema.Attribute{
											"limits": schema.MapAttribute{
												Description:         "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Limits describes the maximum amount of compute resources allowed.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"requests": schema.MapAttribute{
												Description:         "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
												MarkdownDescription: "Requests describes the minimum amount of compute resources required.If Requests is omitted for a container, it defaults to Limits if that is explicitly specified,otherwise to an implementation-defined value. Requests cannot exceed Limits.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

									"selector": schema.SingleNestedAttribute{
										Description:         "selector is a label query over volumes to consider for binding.",
										MarkdownDescription: "selector is a label query over volumes to consider for binding.",
										Attributes: map[string]schema.Attribute{
											"match_expressions": schema.ListNestedAttribute{
												Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"key": schema.StringAttribute{
															Description:         "key is the label key that the selector applies to.",
															MarkdownDescription: "key is the label key that the selector applies to.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
															MarkdownDescription: "operator represents a key's relationship to a set of values.Valid operators are In, NotIn, Exists and DoesNotExist.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"values": schema.ListAttribute{
															Description:         "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
															MarkdownDescription: "values is an array of string values. If the operator is In or NotIn,the values array must be non-empty. If the operator is Exists or DoesNotExist,the values array must be empty. This array is replaced during a strategicmerge patch.",
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
												Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
												MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabelsmap is equivalent to an element of matchExpressions, whose key field is 'key', theoperator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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

									"storage_class_name": schema.StringAttribute{
										Description:         "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_attributes_class_name": schema.StringAttribute{
										Description:         "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
										MarkdownDescription: "volumeAttributesClassName may be used to set the VolumeAttributesClass used by this claim.If specified, the CSI driver will create or update the volume with the attributes definedin the corresponding VolumeAttributesClass. This has a different purpose than storageClassName,it can be changed after the claim is created. An empty string value means that no VolumeAttributesClasswill be applied to the claim but it's not allowed to reset this field to empty string once it is set.If unspecified and the PersistentVolumeClaim is unbound, the default VolumeAttributesClasswill be set by the persistentvolume controller if it exists.If the resource referred to by volumeAttributesClass does not exist, this PersistentVolumeClaim will beset to a Pending state, as reflected by the modifyVolumeStatus field, until such as a resourceexists.More info: https://kubernetes.io/docs/concepts/storage/volume-attributes-classes/(Alpha) Using this field requires the VolumeAttributesClass feature gate to be enabled.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_mode": schema.StringAttribute{
										Description:         "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
										MarkdownDescription: "volumeMode defines what type of volume is required by the claim.Value of Filesystem is implied when not included in claim spec.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"volume_name": schema.StringAttribute{
										Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
										MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"resize_in_use_volumes": schema.BoolAttribute{
								Description:         "Resize existent PVCs, defaults to true",
								MarkdownDescription: "Resize existent PVCs, defaults to true",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"size": schema.StringAttribute{
								Description:         "Size of the storage. Required if not already specified in the PVC template.Changes to this field are automatically reapplied to the created PVCs.Size cannot be decreased.",
								MarkdownDescription: "Size of the storage. Required if not already specified in the PVC template.Changes to this field are automatically reapplied to the created PVCs.Size cannot be decreased.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"storage_class": schema.StringAttribute{
								Description:         "StorageClass to use for PVCs. Applied afterevaluating the PVC template, if available.If not specified, the generated PVCs will use thedefault storage class",
								MarkdownDescription: "StorageClass to use for PVCs. Applied afterevaluating the PVC template, if available.If not specified, the generated PVCs will use thedefault storage class",
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
				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}
}

func (r *PostgresqlCnpgIoClusterV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_postgresql_cnpg_io_cluster_v1_manifest")

	var model PostgresqlCnpgIoClusterV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("postgresql.cnpg.io/v1")
	model.Kind = pointer.String("Cluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
