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

type ScyllaScylladbComScyllaClusterV1Resource struct{}

var (
	_ resource.Resource = (*ScyllaScylladbComScyllaClusterV1Resource)(nil)
)

type ScyllaScylladbComScyllaClusterV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type ScyllaScylladbComScyllaClusterV1GoModel struct {
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
		AgentRepository *string `tfsdk:"agent_repository" yaml:"agentRepository,omitempty"`

		AgentVersion *string `tfsdk:"agent_version" yaml:"agentVersion,omitempty"`

		Alternator *struct {
			Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

			WriteIsolation *string `tfsdk:"write_isolation" yaml:"writeIsolation,omitempty"`
		} `tfsdk:"alternator" yaml:"alternator,omitempty"`

		AutomaticOrphanedNodeCleanup *bool `tfsdk:"automatic_orphaned_node_cleanup" yaml:"automaticOrphanedNodeCleanup,omitempty"`

		Backups *[]struct {
			Dc *[]string `tfsdk:"dc" yaml:"dc,omitempty"`

			Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

			Keyspace *[]string `tfsdk:"keyspace" yaml:"keyspace,omitempty"`

			Location *[]string `tfsdk:"location" yaml:"location,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			NumRetries *int64 `tfsdk:"num_retries" yaml:"numRetries,omitempty"`

			RateLimit *[]string `tfsdk:"rate_limit" yaml:"rateLimit,omitempty"`

			Retention *int64 `tfsdk:"retention" yaml:"retention,omitempty"`

			SnapshotParallel *[]string `tfsdk:"snapshot_parallel" yaml:"snapshotParallel,omitempty"`

			StartDate *string `tfsdk:"start_date" yaml:"startDate,omitempty"`

			UploadParallel *[]string `tfsdk:"upload_parallel" yaml:"uploadParallel,omitempty"`
		} `tfsdk:"backups" yaml:"backups,omitempty"`

		Cpuset *bool `tfsdk:"cpuset" yaml:"cpuset,omitempty"`

		Datacenter *struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Racks *[]struct {
				AgentResources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"agent_resources" yaml:"agentResources,omitempty"`

				AgentVolumeMounts *[]struct {
					MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

					MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

					SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

					SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
				} `tfsdk:"agent_volume_mounts" yaml:"agentVolumeMounts,omitempty"`

				Members *int64 `tfsdk:"members" yaml:"members,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Placement *struct {
					NodeAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							Preference *struct {
								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchFields *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
							} `tfsdk:"preference" yaml:"preference,omitempty"`

							Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

						RequiredDuringSchedulingIgnoredDuringExecution *struct {
							NodeSelectorTerms *[]struct {
								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchFields *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_fields" yaml:"matchFields,omitempty"`
							} `tfsdk:"node_selector_terms" yaml:"nodeSelectorTerms,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`

					PodAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							PodAffinityTerm *struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

							Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

						RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

							NamespaceSelector *struct {
								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

							Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

							TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_affinity" yaml:"podAffinity,omitempty"`

					PodAntiAffinity *struct {
						PreferredDuringSchedulingIgnoredDuringExecution *[]struct {
							PodAffinityTerm *struct {
								LabelSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

								NamespaceSelector *struct {
									MatchExpressions *[]struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

										Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
									} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

									MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
								} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

								Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

								TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
							} `tfsdk:"pod_affinity_term" yaml:"podAffinityTerm,omitempty"`

							Weight *int64 `tfsdk:"weight" yaml:"weight,omitempty"`
						} `tfsdk:"preferred_during_scheduling_ignored_during_execution" yaml:"preferredDuringSchedulingIgnoredDuringExecution,omitempty"`

						RequiredDuringSchedulingIgnoredDuringExecution *[]struct {
							LabelSelector *struct {
								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

							NamespaceSelector *struct {
								MatchExpressions *[]struct {
									Key *string `tfsdk:"key" yaml:"key,omitempty"`

									Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

									Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
								} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

								MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
							} `tfsdk:"namespace_selector" yaml:"namespaceSelector,omitempty"`

							Namespaces *[]string `tfsdk:"namespaces" yaml:"namespaces,omitempty"`

							TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`
						} `tfsdk:"required_during_scheduling_ignored_during_execution" yaml:"requiredDuringSchedulingIgnoredDuringExecution,omitempty"`
					} `tfsdk:"pod_anti_affinity" yaml:"podAntiAffinity,omitempty"`

					Tolerations *[]struct {
						Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
				} `tfsdk:"placement" yaml:"placement,omitempty"`

				Resources *struct {
					Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

					Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
				} `tfsdk:"resources" yaml:"resources,omitempty"`

				ScyllaAgentConfig *string `tfsdk:"scylla_agent_config" yaml:"scyllaAgentConfig,omitempty"`

				ScyllaConfig *string `tfsdk:"scylla_config" yaml:"scyllaConfig,omitempty"`

				Storage *struct {
					Capacity *string `tfsdk:"capacity" yaml:"capacity,omitempty"`

					StorageClassName *string `tfsdk:"storage_class_name" yaml:"storageClassName,omitempty"`
				} `tfsdk:"storage" yaml:"storage,omitempty"`

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
			} `tfsdk:"racks" yaml:"racks,omitempty"`
		} `tfsdk:"datacenter" yaml:"datacenter,omitempty"`

		DeveloperMode *bool `tfsdk:"developer_mode" yaml:"developerMode,omitempty"`

		DnsDomains *[]string `tfsdk:"dns_domains" yaml:"dnsDomains,omitempty"`

		ExposeOptions *struct {
			Cql *struct {
				Ingress *struct {
					Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

					Disabled *bool `tfsdk:"disabled" yaml:"disabled,omitempty"`

					IngressClassName *string `tfsdk:"ingress_class_name" yaml:"ingressClassName,omitempty"`
				} `tfsdk:"ingress" yaml:"ingress,omitempty"`
			} `tfsdk:"cql" yaml:"cql,omitempty"`
		} `tfsdk:"expose_options" yaml:"exposeOptions,omitempty"`

		ForceRedeploymentReason *string `tfsdk:"force_redeployment_reason" yaml:"forceRedeploymentReason,omitempty"`

		GenericUpgrade *struct {
			FailureStrategy *string `tfsdk:"failure_strategy" yaml:"failureStrategy,omitempty"`

			PollInterval *string `tfsdk:"poll_interval" yaml:"pollInterval,omitempty"`
		} `tfsdk:"generic_upgrade" yaml:"genericUpgrade,omitempty"`

		ImagePullSecrets *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`
		} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

		Network *struct {
			DnsPolicy *string `tfsdk:"dns_policy" yaml:"dnsPolicy,omitempty"`

			HostNetworking *bool `tfsdk:"host_networking" yaml:"hostNetworking,omitempty"`
		} `tfsdk:"network" yaml:"network,omitempty"`

		Repairs *[]struct {
			Dc *[]string `tfsdk:"dc" yaml:"dc,omitempty"`

			FailFast *bool `tfsdk:"fail_fast" yaml:"failFast,omitempty"`

			Host *string `tfsdk:"host" yaml:"host,omitempty"`

			Intensity *string `tfsdk:"intensity" yaml:"intensity,omitempty"`

			Interval *string `tfsdk:"interval" yaml:"interval,omitempty"`

			Keyspace *[]string `tfsdk:"keyspace" yaml:"keyspace,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			NumRetries *int64 `tfsdk:"num_retries" yaml:"numRetries,omitempty"`

			Parallel *int64 `tfsdk:"parallel" yaml:"parallel,omitempty"`

			SmallTableThreshold *string `tfsdk:"small_table_threshold" yaml:"smallTableThreshold,omitempty"`

			StartDate *string `tfsdk:"start_date" yaml:"startDate,omitempty"`
		} `tfsdk:"repairs" yaml:"repairs,omitempty"`

		Repository *string `tfsdk:"repository" yaml:"repository,omitempty"`

		ScyllaArgs *string `tfsdk:"scylla_args" yaml:"scyllaArgs,omitempty"`

		Sysctls *[]string `tfsdk:"sysctls" yaml:"sysctls,omitempty"`

		Version *string `tfsdk:"version" yaml:"version,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewScyllaScylladbComScyllaClusterV1Resource() resource.Resource {
	return &ScyllaScylladbComScyllaClusterV1Resource{}
}

func (r *ScyllaScylladbComScyllaClusterV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_scylla_scylladb_com_scylla_cluster_v1"
}

func (r *ScyllaScylladbComScyllaClusterV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ScyllaCluster defines a Scylla cluster.",
		MarkdownDescription: "ScyllaCluster defines a Scylla cluster.",
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
				Description:         "spec defines the desired state of this scylla cluster.",
				MarkdownDescription: "spec defines the desired state of this scylla cluster.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"agent_repository": {
						Description:         "agentRepository is the repository to pull the agent image from.",
						MarkdownDescription: "agentRepository is the repository to pull the agent image from.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"agent_version": {
						Description:         "agentVersion indicates the version of Scylla Manager Agent to use.",
						MarkdownDescription: "agentVersion indicates the version of Scylla Manager Agent to use.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"alternator": {
						Description:         "alternator designates this cluster an Alternator cluster.",
						MarkdownDescription: "alternator designates this cluster an Alternator cluster.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"port": {
								Description:         "port is the port number used to bind the Alternator API.",
								MarkdownDescription: "port is the port number used to bind the Alternator API.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"write_isolation": {
								Description:         "writeIsolation indicates the isolation level.",
								MarkdownDescription: "writeIsolation indicates the isolation level.",

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

					"automatic_orphaned_node_cleanup": {
						Description:         "automaticOrphanedNodeCleanup controls if automatic orphan node cleanup should be performed.",
						MarkdownDescription: "automaticOrphanedNodeCleanup controls if automatic orphan node cleanup should be performed.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"backups": {
						Description:         "backups specifies backup tasks in Scylla Manager. When Scylla Manager is not installed, these will be ignored.",
						MarkdownDescription: "backups specifies backup tasks in Scylla Manager. When Scylla Manager is not installed, these will be ignored.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"dc": {
								Description:         "dc is a list of datacenter glob patterns, e.g. 'dc1,!otherdc*' used to specify the DCs to include or exclude from backup.",
								MarkdownDescription: "dc is a list of datacenter glob patterns, e.g. 'dc1,!otherdc*' used to specify the DCs to include or exclude from backup.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interval": {
								Description:         "interval represents a task schedule interval e.g. 3d2h10m, valid units are d, h, m, s.",
								MarkdownDescription: "interval represents a task schedule interval e.g. 3d2h10m, valid units are d, h, m, s.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"keyspace": {
								Description:         "keyspace is a list of keyspace/tables glob patterns, e.g. 'keyspace,!keyspace.table_prefix_*' used to include or exclude keyspaces from repair.",
								MarkdownDescription: "keyspace is a list of keyspace/tables glob patterns, e.g. 'keyspace,!keyspace.table_prefix_*' used to include or exclude keyspaces from repair.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"location": {
								Description:         "location is a list of backup locations in the format [<dc>:]<provider>:<name> ex. s3:my-bucket. The <dc>: part is optional and is only needed when different datacenters are being used to upload data to different locations. <name> must be an alphanumeric string and may contain a dash and or a dot, but other characters are forbidden. The only supported storage <provider> at the moment are s3 and gcs.",
								MarkdownDescription: "location is a list of backup locations in the format [<dc>:]<provider>:<name> ex. s3:my-bucket. The <dc>: part is optional and is only needed when different datacenters are being used to upload data to different locations. <name> must be an alphanumeric string and may contain a dash and or a dot, but other characters are forbidden. The only supported storage <provider> at the moment are s3 and gcs.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "name is a unique name of a task.",
								MarkdownDescription: "name is a unique name of a task.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"num_retries": {
								Description:         "numRetries indicates how many times a scheduled task will be retried before failing.",
								MarkdownDescription: "numRetries indicates how many times a scheduled task will be retried before failing.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"rate_limit": {
								Description:         "rateLimit is a list of megabytes (MiB) per second rate limits expressed in the format [<dc>:]<limit>. The <dc>: part is optional and only needed when different datacenters need different upload limits. Set to 0 for no limit (default 100).",
								MarkdownDescription: "rateLimit is a list of megabytes (MiB) per second rate limits expressed in the format [<dc>:]<limit>. The <dc>: part is optional and only needed when different datacenters need different upload limits. Set to 0 for no limit (default 100).",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"retention": {
								Description:         "retention is the number of backups which are to be stored.",
								MarkdownDescription: "retention is the number of backups which are to be stored.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"snapshot_parallel": {
								Description:         "snapshotParallel is a list of snapshot parallelism limits in the format [<dc>:]<limit>. The <dc>: part is optional and allows for specifying different limits in selected datacenters. If The <dc>: part is not set, the limit is global (e.g. 'dc1:2,5') the runs are parallel in n nodes (2 in dc1) and n nodes in all the other datacenters.",
								MarkdownDescription: "snapshotParallel is a list of snapshot parallelism limits in the format [<dc>:]<limit>. The <dc>: part is optional and allows for specifying different limits in selected datacenters. If The <dc>: part is not set, the limit is global (e.g. 'dc1:2,5') the runs are parallel in n nodes (2 in dc1) and n nodes in all the other datacenters.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"start_date": {
								Description:         "startDate specifies the task start date expressed in the RFC3339 format or now[+duration], e.g. now+3d2h10m, valid units are d, h, m, s.",
								MarkdownDescription: "startDate specifies the task start date expressed in the RFC3339 format or now[+duration], e.g. now+3d2h10m, valid units are d, h, m, s.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"upload_parallel": {
								Description:         "uploadParallel is a list of upload parallelism limits in the format [<dc>:]<limit>. The <dc>: part is optional and allows for specifying different limits in selected datacenters. If The <dc>: part is not set the limit is global (e.g. 'dc1:2,5') the runs are parallel in n nodes (2 in dc1) and n nodes in all the other datacenters.",
								MarkdownDescription: "uploadParallel is a list of upload parallelism limits in the format [<dc>:]<limit>. The <dc>: part is optional and allows for specifying different limits in selected datacenters. If The <dc>: part is not set the limit is global (e.g. 'dc1:2,5') the runs are parallel in n nodes (2 in dc1) and n nodes in all the other datacenters.",

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

					"cpuset": {
						Description:         "cpuset determines if the cluster will use cpu-pinning for max performance.",
						MarkdownDescription: "cpuset determines if the cluster will use cpu-pinning for max performance.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"datacenter": {
						Description:         "datacenter holds a specification of a datacenter.",
						MarkdownDescription: "datacenter holds a specification of a datacenter.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "name is the name of the scylla datacenter. Used in the cassandra-rackdc.properties file.",
								MarkdownDescription: "name is the name of the scylla datacenter. Used in the cassandra-rackdc.properties file.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"racks": {
								Description:         "racks specify the racks in the datacenter.",
								MarkdownDescription: "racks specify the racks in the datacenter.",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"agent_resources": {
										Description:         "agentResources specify the resources for the Agent container.",
										MarkdownDescription: "agentResources specify the resources for the Agent container.",

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

									"agent_volume_mounts": {
										Description:         "AgentVolumeMounts to be added to Agent container.",
										MarkdownDescription: "AgentVolumeMounts to be added to Agent container.",

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

									"members": {
										Description:         "members is the number of Scylla instances in this rack.",
										MarkdownDescription: "members is the number of Scylla instances in this rack.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "name is the name of the Scylla Rack. Used in the cassandra-rackdc.properties file.",
										MarkdownDescription: "name is the name of the Scylla Rack. Used in the cassandra-rackdc.properties file.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"placement": {
										Description:         "placement describes restrictions for the nodes Scylla is scheduled on.",
										MarkdownDescription: "placement describes restrictions for the nodes Scylla is scheduled on.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"node_affinity": {
												Description:         "nodeAffinity describes node affinity scheduling rules for the pod.",
												MarkdownDescription: "nodeAffinity describes node affinity scheduling rules for the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"preferred_during_scheduling_ignored_during_execution": {
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"preference": {
																Description:         "A node selector term, associated with the corresponding weight.",
																MarkdownDescription: "A node selector term, associated with the corresponding weight.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"match_expressions": {
																		Description:         "A list of node selector requirements by node's labels.",
																		MarkdownDescription: "A list of node selector requirements by node's labels.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"operator": {
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

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

																	"match_fields": {
																		Description:         "A list of node selector requirements by node's fields.",
																		MarkdownDescription: "A list of node selector requirements by node's fields.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"operator": {
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

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
																}),

																Required: true,
																Optional: false,
																Computed: false,
															},

															"weight": {
																Description:         "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",
																MarkdownDescription: "Weight associated with matching the corresponding nodeSelectorTerm, in the range 1-100.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"required_during_scheduling_ignored_during_execution": {
														Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
														MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"node_selector_terms": {
																Description:         "Required. A list of node selector terms. The terms are ORed.",
																MarkdownDescription: "Required. A list of node selector terms. The terms are ORed.",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"match_expressions": {
																		Description:         "A list of node selector requirements by node's labels.",
																		MarkdownDescription: "A list of node selector requirements by node's labels.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"operator": {
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

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

																	"match_fields": {
																		Description:         "A list of node selector requirements by node's fields.",
																		MarkdownDescription: "A list of node selector requirements by node's fields.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "The label key that the selector applies to.",
																				MarkdownDescription: "The label key that the selector applies to.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"operator": {
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"values": {
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",

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

											"pod_affinity": {
												Description:         "podAffinity describes pod affinity scheduling rules.",
												MarkdownDescription: "podAffinity describes pod affinity scheduling rules.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"preferred_during_scheduling_ignored_during_execution": {
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"pod_affinity_term": {
																Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"label_selector": {
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",

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

																	"namespace_selector": {
																		Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

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

																	"namespaces": {
																		Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"topology_key": {
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

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

															"weight": {
																Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"required_during_scheduling_ignored_during_execution": {
														Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
														MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"label_selector": {
																Description:         "A label query over a set of resources, in this case pods.",
																MarkdownDescription: "A label query over a set of resources, in this case pods.",

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

															"namespace_selector": {
																Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

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

															"namespaces": {
																Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"topology_key": {
																Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

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

											"pod_anti_affinity": {
												Description:         "podAntiAffinity describes pod anti-affinity scheduling rules.",
												MarkdownDescription: "podAntiAffinity describes pod anti-affinity scheduling rules.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"preferred_during_scheduling_ignored_during_execution": {
														Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
														MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"pod_affinity_term": {
																Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"label_selector": {
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",

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

																	"namespace_selector": {
																		Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

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

																	"namespaces": {
																		Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"topology_key": {
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

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

															"weight": {
																Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",

																Type: types.Int64Type,

																Required: true,
																Optional: false,
																Computed: false,
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"required_during_scheduling_ignored_during_execution": {
														Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
														MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"label_selector": {
																Description:         "A label query over a set of resources, in this case pods.",
																MarkdownDescription: "A label query over a set of resources, in this case pods.",

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

															"namespace_selector": {
																Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",

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

															"namespaces": {
																Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"topology_key": {
																Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",

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

											"tolerations": {
												Description:         "tolerations allow the pod to tolerate any taint that matches the triple <key,value,effect> using the matching operator.",
												MarkdownDescription: "tolerations allow the pod to tolerate any taint that matches the triple <key,value,effect> using the matching operator.",

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

									"resources": {
										Description:         "resources the Scylla container will use.",
										MarkdownDescription: "resources the Scylla container will use.",

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

									"scylla_agent_config": {
										Description:         "Scylla config map name to customize scylla manager agent",
										MarkdownDescription: "Scylla config map name to customize scylla manager agent",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"scylla_config": {
										Description:         "Scylla config map name to customize scylla.yaml",
										MarkdownDescription: "Scylla config map name to customize scylla.yaml",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"storage": {
										Description:         "storage describes the underlying storage that Scylla will consume.",
										MarkdownDescription: "storage describes the underlying storage that Scylla will consume.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"capacity": {
												Description:         "capacity describes the requested size of each persistent volume.",
												MarkdownDescription: "capacity describes the requested size of each persistent volume.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"storage_class_name": {
												Description:         "storageClassName is the name of a storageClass to request.",
												MarkdownDescription: "storageClassName is the name of a storageClass to request.",

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

									"volume_mounts": {
										Description:         "VolumeMounts to be added to Scylla container.",
										MarkdownDescription: "VolumeMounts to be added to Scylla container.",

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
										Description:         "Volumes added to Scylla Pod.",
										MarkdownDescription: "Volumes added to Scylla Pod.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"aws_elastic_block_store": {
												Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
												MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"partition": {
														Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
														MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_id": {
														Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
														MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

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
												Description:         "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
												MarkdownDescription: "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"caching_mode": {
														Description:         "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
														MarkdownDescription: "cachingMode is the Host Caching mode: None, Read Only, Read Write.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"disk_name": {
														Description:         "diskName is the Name of the data disk in the blob storage",
														MarkdownDescription: "diskName is the Name of the data disk in the blob storage",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"disk_uri": {
														Description:         "diskURI is the URI of data disk in the blob storage",
														MarkdownDescription: "diskURI is the URI of data disk in the blob storage",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"fs_type": {
														Description:         "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"kind": {
														Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
														MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

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
												Description:         "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
												MarkdownDescription: "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"read_only": {
														Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_name": {
														Description:         "secretName is the  name of secret that contains Azure Storage Account Name and Key",
														MarkdownDescription: "secretName is the  name of secret that contains Azure Storage Account Name and Key",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"share_name": {
														Description:         "shareName is the azure share Name",
														MarkdownDescription: "shareName is the azure share Name",

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
												Description:         "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
												MarkdownDescription: "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"monitors": {
														Description:         "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"path": {
														Description:         "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
														MarkdownDescription: "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_file": {
														Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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
														Description:         "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
														MarkdownDescription: "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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
												Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
														MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",

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
														Description:         "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
														MarkdownDescription: "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

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
												Description:         "configMap represents a configMap that should populate this volume",
												MarkdownDescription: "configMap represents a configMap that should populate this volume",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mode": {
																Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
														Description:         "optional specify whether the ConfigMap or its keys must be defined",
														MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",

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
												Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
												MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"driver": {
														Description:         "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
														MarkdownDescription: "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"fs_type": {
														Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
														MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"node_publish_secret_ref": {
														Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
														MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",

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
														Description:         "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
														MarkdownDescription: "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_attributes": {
														Description:         "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
														MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",

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
												Description:         "downwardAPI represents downward API about the pod that should populate this volume",
												MarkdownDescription: "downwardAPI represents downward API about the pod that should populate this volume",

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
												Description:         "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
												MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"medium": {
														Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
														MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"size_limit": {
														Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
														MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",

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
												Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
												MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity    tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through    a PersistentVolumeClaim (see EphemeralVolumeSource for more    information on the connection between this volume type    and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",

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
																		Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																		MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"data_source": {
																		Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																		MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",

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
																		Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																		MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef   allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef   preserves all values, and generates an error if a disallowed value is   specified. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",

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
																		Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																		MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",

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
																		Description:         "selector is a label query over volumes to consider for binding.",
																		MarkdownDescription: "selector is a label query over volumes to consider for binding.",

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
																		Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																		MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",

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
																		Description:         "volumeName is the binding reference to the PersistentVolume backing this claim.",
																		MarkdownDescription: "volumeName is the binding reference to the PersistentVolume backing this claim.",

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
												Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
												MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"lun": {
														Description:         "lun is Optional: FC target lun number",
														MarkdownDescription: "lun is Optional: FC target lun number",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"target_ww_ns": {
														Description:         "targetWWNs is Optional: FC target worldwide names (WWNs)",
														MarkdownDescription: "targetWWNs is Optional: FC target worldwide names (WWNs)",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"wwids": {
														Description:         "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
														MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",

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
												Description:         "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
												MarkdownDescription: "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"driver": {
														Description:         "driver is the name of the driver to use for this volume.",
														MarkdownDescription: "driver is the name of the driver to use for this volume.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"options": {
														Description:         "options is Optional: this field holds extra command options if any.",
														MarkdownDescription: "options is Optional: this field holds extra command options if any.",

														Type: types.MapType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
														MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",

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
												Description:         "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
												MarkdownDescription: "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"dataset_name": {
														Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
														MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"dataset_uuid": {
														Description:         "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
														MarkdownDescription: "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",

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
												Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
												MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"partition": {
														Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pd_name": {
														Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
														MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

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
												Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
												MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"directory": {
														Description:         "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
														MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"repository": {
														Description:         "repository is the URL",
														MarkdownDescription: "repository is the URL",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"revision": {
														Description:         "revision is the commit hash for the specified revision.",
														MarkdownDescription: "revision is the commit hash for the specified revision.",

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
												Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
												MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"endpoints": {
														Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"path": {
														Description:         "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
														MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

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
												Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
												MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"type": {
														Description:         "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
														MarkdownDescription: "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

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
												Description:         "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
												MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"chap_auth_discovery": {
														Description:         "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
														MarkdownDescription: "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"chap_auth_session": {
														Description:         "chapAuthSession defines whether support iSCSI Session CHAP authentication",
														MarkdownDescription: "chapAuthSession defines whether support iSCSI Session CHAP authentication",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"fs_type": {
														Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"initiator_name": {
														Description:         "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
														MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"iqn": {
														Description:         "iqn is the target iSCSI Qualified Name.",
														MarkdownDescription: "iqn is the target iSCSI Qualified Name.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"iscsi_interface": {
														Description:         "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
														MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"lun": {
														Description:         "lun represents iSCSI Target Lun number.",
														MarkdownDescription: "lun represents iSCSI Target Lun number.",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"portals": {
														Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
														MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
														MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
														MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",

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
														Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
														MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

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
												Description:         "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"nfs": {
												Description:         "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
												MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"path": {
														Description:         "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"server": {
														Description:         "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
														MarkdownDescription: "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

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
												Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
												MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"claim_name": {
														Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
														MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
														MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",

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
												Description:         "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
												MarkdownDescription: "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"pd_id": {
														Description:         "pdID is the ID that identifies Photon Controller persistent disk",
														MarkdownDescription: "pdID is the ID that identifies Photon Controller persistent disk",

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
												Description:         "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
												MarkdownDescription: "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_id": {
														Description:         "volumeID uniquely identifies a Portworx volume",
														MarkdownDescription: "volumeID uniquely identifies a Portworx volume",

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
												Description:         "projected items for all in one resources secrets, configmaps, and downward API",
												MarkdownDescription: "projected items for all in one resources secrets, configmaps, and downward API",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"sources": {
														Description:         "sources is the list of volume projections",
														MarkdownDescription: "sources is the list of volume projections",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"config_map": {
																Description:         "configMap information about the configMap data to project",
																MarkdownDescription: "configMap information about the configMap data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"items": {
																		Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "key is the key to project.",
																				MarkdownDescription: "key is the key to project.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"mode": {
																				Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
																		Description:         "optional specify whether the ConfigMap or its keys must be defined",
																		MarkdownDescription: "optional specify whether the ConfigMap or its keys must be defined",

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
																Description:         "downwardAPI information about the downwardAPI data to project",
																MarkdownDescription: "downwardAPI information about the downwardAPI data to project",

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
																Description:         "secret information about the secret data to project",
																MarkdownDescription: "secret information about the secret data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"items": {
																		Description:         "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																		Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																			"key": {
																				Description:         "key is the key to project.",
																				MarkdownDescription: "key is the key to project.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"mode": {
																				Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"path": {
																				Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
																		Description:         "optional field specify whether the Secret or its key must be defined",
																		MarkdownDescription: "optional field specify whether the Secret or its key must be defined",

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
																Description:         "serviceAccountToken is information about the serviceAccountToken data to project",
																MarkdownDescription: "serviceAccountToken is information about the serviceAccountToken data to project",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"audience": {
																		Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																		MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"expiration_seconds": {
																		Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																		MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",

																		Type: types.Int64Type,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"path": {
																		Description:         "path is the path relative to the mount point of the file to project the token into.",
																		MarkdownDescription: "path is the path relative to the mount point of the file to project the token into.",

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
												Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
												MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"group": {
														Description:         "group to map volume access to Default is no group",
														MarkdownDescription: "group to map volume access to Default is no group",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
														MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"registry": {
														Description:         "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
														MarkdownDescription: "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"tenant": {
														Description:         "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
														MarkdownDescription: "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"user": {
														Description:         "user to map volume access to Defaults to serivceaccount user",
														MarkdownDescription: "user to map volume access to Defaults to serivceaccount user",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume": {
														Description:         "volume is a string that references an already created Quobyte volume by name.",
														MarkdownDescription: "volume is a string that references an already created Quobyte volume by name.",

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
												Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
												MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
														MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"image": {
														Description:         "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"keyring": {
														Description:         "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"monitors": {
														Description:         "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.ListType{ElemType: types.StringType},

														Required: true,
														Optional: false,
														Computed: false,
													},

													"pool": {
														Description:         "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

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
														Description:         "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
														MarkdownDescription: "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

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
												Description:         "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
												MarkdownDescription: "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"gateway": {
														Description:         "gateway is the host address of the ScaleIO API Gateway.",
														MarkdownDescription: "gateway is the host address of the ScaleIO API Gateway.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"protection_domain": {
														Description:         "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
														MarkdownDescription: "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
														MarkdownDescription: "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",

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
														Description:         "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
														MarkdownDescription: "sslEnabled Flag enable/disable SSL communication with Gateway, default false",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_mode": {
														Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
														MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_pool": {
														Description:         "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
														MarkdownDescription: "storagePool is the ScaleIO Storage Pool associated with the protection domain.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"system": {
														Description:         "system is the name of the storage system as configured in ScaleIO.",
														MarkdownDescription: "system is the name of the storage system as configured in ScaleIO.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"volume_name": {
														Description:         "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
														MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",

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
												Description:         "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
												MarkdownDescription: "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"default_mode": {
														Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"items": {
														Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

														Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

															"key": {
																Description:         "key is the key to project.",
																MarkdownDescription: "key is the key to project.",

																Type: types.StringType,

																Required: true,
																Optional: false,
																Computed: false,
															},

															"mode": {
																Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																Type: types.Int64Type,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"path": {
																Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
														Description:         "optional field specify whether the Secret or its keys must be defined",
														MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_name": {
														Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
														MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

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
												Description:         "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
												MarkdownDescription: "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"read_only": {
														Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
														MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"secret_ref": {
														Description:         "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
														MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",

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
														Description:         "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
														MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_namespace": {
														Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
														MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",

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
												Description:         "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
												MarkdownDescription: "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"fs_type": {
														Description:         "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
														MarkdownDescription: "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_policy_id": {
														Description:         "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
														MarkdownDescription: "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"storage_policy_name": {
														Description:         "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
														MarkdownDescription: "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"volume_path": {
														Description:         "volumePath is the path that identifies vSphere volume vmdk",
														MarkdownDescription: "volumePath is the path that identifies vSphere volume vmdk",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"developer_mode": {
						Description:         "developerMode determines if the cluster runs in developer-mode.",
						MarkdownDescription: "developerMode determines if the cluster runs in developer-mode.",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"dns_domains": {
						Description:         "dnsDomains is a list of DNS domains this cluster is reachable by. These domains are used when setting up the infrastructure, like certificates. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",
						MarkdownDescription: "dnsDomains is a list of DNS domains this cluster is reachable by. These domains are used when setting up the infrastructure, like certificates. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"expose_options": {
						Description:         "exposeOptions specifies options for exposing ScyllaCluster services. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",
						MarkdownDescription: "exposeOptions specifies options for exposing ScyllaCluster services. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"cql": {
								Description:         "cql specifies expose options for CQL SSL backend. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",
								MarkdownDescription: "cql specifies expose options for CQL SSL backend. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"ingress": {
										Description:         "ingress is an Ingress configuration options. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",
										MarkdownDescription: "ingress is an Ingress configuration options. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"annotations": {
												Description:         "annotations specifies custom annotations merged into every Ingress object. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",
												MarkdownDescription: "annotations specifies custom annotations merged into every Ingress object. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"disabled": {
												Description:         "disabled controls if Ingress object creation is disabled. Unless disabled, there is an Ingress objects created for every Scylla node. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",
												MarkdownDescription: "disabled controls if Ingress object creation is disabled. Unless disabled, there is an Ingress objects created for every Scylla node. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",

												Type: types.BoolType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"ingress_class_name": {
												Description:         "ingressClassName specifies Ingress class name. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",
												MarkdownDescription: "ingressClassName specifies Ingress class name. EXPERIMENTAL. Do not rely on any particular behaviour controlled by this field.",

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

					"force_redeployment_reason": {
						Description:         "forceRedeploymentReason can be used to force a rolling update of all racks by providing a unique string.",
						MarkdownDescription: "forceRedeploymentReason can be used to force a rolling update of all racks by providing a unique string.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"generic_upgrade": {
						Description:         "genericUpgrade allows to configure behavior of generic upgrade logic.",
						MarkdownDescription: "genericUpgrade allows to configure behavior of generic upgrade logic.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"failure_strategy": {
								Description:         "failureStrategy specifies which logic is executed when upgrade failure happens. Currently only Retry is supported.",
								MarkdownDescription: "failureStrategy specifies which logic is executed when upgrade failure happens. Currently only Retry is supported.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"poll_interval": {
								Description:         "pollInterval specifies how often upgrade logic polls on state updates. Increasing this value should lower number of requests sent to apiserver, but it may affect overall time spent during upgrade. DEPRECATED.",
								MarkdownDescription: "pollInterval specifies how often upgrade logic polls on state updates. Increasing this value should lower number of requests sent to apiserver, but it may affect overall time spent during upgrade. DEPRECATED.",

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

					"image_pull_secrets": {
						Description:         "imagePullSecrets is an optional list of references to secrets in the same namespace used for pulling Scylla and Agent images.",
						MarkdownDescription: "imagePullSecrets is an optional list of references to secrets in the same namespace used for pulling Scylla and Agent images.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

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

					"network": {
						Description:         "network holds the networking config.",
						MarkdownDescription: "network holds the networking config.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"dns_policy": {
								Description:         "dnsPolicy defines how a pod's DNS will be configured.",
								MarkdownDescription: "dnsPolicy defines how a pod's DNS will be configured.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host_networking": {
								Description:         "hostNetworking determines if scylla uses the host's network namespace. Setting this option avoids going through Kubernetes SDN and exposes scylla on node's IP.",
								MarkdownDescription: "hostNetworking determines if scylla uses the host's network namespace. Setting this option avoids going through Kubernetes SDN and exposes scylla on node's IP.",

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

					"repairs": {
						Description:         "repairs specify repair tasks in Scylla Manager. When Scylla Manager is not installed, these will be ignored.",
						MarkdownDescription: "repairs specify repair tasks in Scylla Manager. When Scylla Manager is not installed, these will be ignored.",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"dc": {
								Description:         "dc is a list of datacenter glob patterns, e.g. 'dc1', '!otherdc*' used to specify the DCs to include or exclude from backup.",
								MarkdownDescription: "dc is a list of datacenter glob patterns, e.g. 'dc1', '!otherdc*' used to specify the DCs to include or exclude from backup.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"fail_fast": {
								Description:         "failFast indicates if a repair should be stopped on first error.",
								MarkdownDescription: "failFast indicates if a repair should be stopped on first error.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"host": {
								Description:         "host specifies a host to repair. If empty, all hosts are repaired.",
								MarkdownDescription: "host specifies a host to repair. If empty, all hosts are repaired.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"intensity": {
								Description:         "intensity indicates how many token ranges (per shard) to repair in a single Scylla repair job. By default this is 1. If you set it to 0 the number of token ranges is adjusted to the maximum supported by node (see max_repair_ranges_in_parallel in Scylla logs). Valid values are 0 and integers >= 1. Higher values will result in increased cluster load and slightly faster repairs. Changing the intensity impacts repair granularity if you need to resume it, the higher the value the more work on resume. For Scylla clusters that *do not support row-level repair*, intensity can be a decimal between (0,1). In that case it specifies percent of shards that can be repaired in parallel on a repair master node. For Scylla clusters that are row-level repair enabled, setting intensity below 1 has the same effect as setting intensity 1.",
								MarkdownDescription: "intensity indicates how many token ranges (per shard) to repair in a single Scylla repair job. By default this is 1. If you set it to 0 the number of token ranges is adjusted to the maximum supported by node (see max_repair_ranges_in_parallel in Scylla logs). Valid values are 0 and integers >= 1. Higher values will result in increased cluster load and slightly faster repairs. Changing the intensity impacts repair granularity if you need to resume it, the higher the value the more work on resume. For Scylla clusters that *do not support row-level repair*, intensity can be a decimal between (0,1). In that case it specifies percent of shards that can be repaired in parallel on a repair master node. For Scylla clusters that are row-level repair enabled, setting intensity below 1 has the same effect as setting intensity 1.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"interval": {
								Description:         "interval represents a task schedule interval e.g. 3d2h10m, valid units are d, h, m, s.",
								MarkdownDescription: "interval represents a task schedule interval e.g. 3d2h10m, valid units are d, h, m, s.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"keyspace": {
								Description:         "keyspace is a list of keyspace/tables glob patterns, e.g. 'keyspace,!keyspace.table_prefix_*' used to include or exclude keyspaces from repair.",
								MarkdownDescription: "keyspace is a list of keyspace/tables glob patterns, e.g. 'keyspace,!keyspace.table_prefix_*' used to include or exclude keyspaces from repair.",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "name is a unique name of a task.",
								MarkdownDescription: "name is a unique name of a task.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"num_retries": {
								Description:         "numRetries indicates how many times a scheduled task will be retried before failing.",
								MarkdownDescription: "numRetries indicates how many times a scheduled task will be retried before failing.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"parallel": {
								Description:         "parallel is the maximum number of Scylla repair jobs that can run at the same time (on different token ranges and replicas). Each node can take part in at most one repair at any given moment. By default the maximum possible parallelism is used. The effective parallelism depends on a keyspace replication factor (RF) and the number of nodes. The formula to calculate it is as follows: number of nodes / RF, ex. for 6 node cluster with RF=3 the maximum parallelism is 2.",
								MarkdownDescription: "parallel is the maximum number of Scylla repair jobs that can run at the same time (on different token ranges and replicas). Each node can take part in at most one repair at any given moment. By default the maximum possible parallelism is used. The effective parallelism depends on a keyspace replication factor (RF) and the number of nodes. The formula to calculate it is as follows: number of nodes / RF, ex. for 6 node cluster with RF=3 the maximum parallelism is 2.",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"small_table_threshold": {
								Description:         "smallTableThreshold enable small table optimization for tables of size lower than given threshold. Supported units [B, MiB, GiB, TiB].",
								MarkdownDescription: "smallTableThreshold enable small table optimization for tables of size lower than given threshold. Supported units [B, MiB, GiB, TiB].",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"start_date": {
								Description:         "startDate specifies the task start date expressed in the RFC3339 format or now[+duration], e.g. now+3d2h10m, valid units are d, h, m, s.",
								MarkdownDescription: "startDate specifies the task start date expressed in the RFC3339 format or now[+duration], e.g. now+3d2h10m, valid units are d, h, m, s.",

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

					"repository": {
						Description:         "repository is the image repository to pull the Scylla image from.",
						MarkdownDescription: "repository is the image repository to pull the Scylla image from.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"scylla_args": {
						Description:         "scyllaArgs will be appended to Scylla binary during startup. This is supported from 4.2.0 Scylla version.",
						MarkdownDescription: "scyllaArgs will be appended to Scylla binary during startup. This is supported from 4.2.0 Scylla version.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"sysctls": {
						Description:         "sysctls holds the sysctl properties to be applied during initialization given as a list of key=value pairs. Example: fs.aio-max-nr=232323",
						MarkdownDescription: "sysctls holds the sysctl properties to be applied during initialization given as a list of key=value pairs. Example: fs.aio-max-nr=232323",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"version": {
						Description:         "version is a version tag of Scylla to use.",
						MarkdownDescription: "version is a version tag of Scylla to use.",

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

func (r *ScyllaScylladbComScyllaClusterV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_scylla_scylladb_com_scylla_cluster_v1")

	var state ScyllaScylladbComScyllaClusterV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ScyllaScylladbComScyllaClusterV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("scylla.scylladb.com/v1")
	goModel.Kind = utilities.Ptr("ScyllaCluster")

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

func (r *ScyllaScylladbComScyllaClusterV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_scylla_scylladb_com_scylla_cluster_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *ScyllaScylladbComScyllaClusterV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_scylla_scylladb_com_scylla_cluster_v1")

	var state ScyllaScylladbComScyllaClusterV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel ScyllaScylladbComScyllaClusterV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("scylla.scylladb.com/v1")
	goModel.Kind = utilities.Ptr("ScyllaCluster")

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

func (r *ScyllaScylladbComScyllaClusterV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_scylla_scylladb_com_scylla_cluster_v1")
	// NO-OP: Terraform removes the state automatically for us
}
