/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package apps_kubeblocks_io_v1alpha1

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
	_ datasource.DataSource = &AppsKubeblocksIoClusterV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoClusterV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoClusterV1Alpha1Manifest{}
}

type AppsKubeblocksIoClusterV1Alpha1Manifest struct{}

type AppsKubeblocksIoClusterV1Alpha1ManifestData struct {
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
			NodeLabels      *map[string]string `tfsdk:"node_labels" json:"nodeLabels,omitempty"`
			PodAntiAffinity *string            `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
			Tenancy         *string            `tfsdk:"tenancy" json:"tenancy,omitempty"`
			TopologyKeys    *[]string          `tfsdk:"topology_keys" json:"topologyKeys,omitempty"`
		} `tfsdk:"affinity" json:"affinity,omitempty"`
		AvailabilityPolicy *string `tfsdk:"availability_policy" json:"availabilityPolicy,omitempty"`
		Backup             *struct {
			CronExpression          *string `tfsdk:"cron_expression" json:"cronExpression,omitempty"`
			Enabled                 *bool   `tfsdk:"enabled" json:"enabled,omitempty"`
			Method                  *string `tfsdk:"method" json:"method,omitempty"`
			PitrEnabled             *bool   `tfsdk:"pitr_enabled" json:"pitrEnabled,omitempty"`
			RepoName                *string `tfsdk:"repo_name" json:"repoName,omitempty"`
			RetentionPeriod         *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
			StartingDeadlineMinutes *int64  `tfsdk:"starting_deadline_minutes" json:"startingDeadlineMinutes,omitempty"`
		} `tfsdk:"backup" json:"backup,omitempty"`
		ClusterDefinitionRef *string `tfsdk:"cluster_definition_ref" json:"clusterDefinitionRef,omitempty"`
		ClusterVersionRef    *string `tfsdk:"cluster_version_ref" json:"clusterVersionRef,omitempty"`
		ComponentSpecs       *[]struct {
			Affinity *struct {
				NodeLabels      *map[string]string `tfsdk:"node_labels" json:"nodeLabels,omitempty"`
				PodAntiAffinity *string            `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
				Tenancy         *string            `tfsdk:"tenancy" json:"tenancy,omitempty"`
				TopologyKeys    *[]string          `tfsdk:"topology_keys" json:"topologyKeys,omitempty"`
			} `tfsdk:"affinity" json:"affinity,omitempty"`
			ComponentDef    *string `tfsdk:"component_def" json:"componentDef,omitempty"`
			ComponentDefRef *string `tfsdk:"component_def_ref" json:"componentDefRef,omitempty"`
			Configs         *[]struct {
				ConfigMap *struct {
					DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
					Items       *[]struct {
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
						Path *string `tfsdk:"path" json:"path,omitempty"`
					} `tfsdk:"items" json:"items,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map" json:"configMap,omitempty"`
				Name *string `tfsdk:"name" json:"name,omitempty"`
			} `tfsdk:"configs" json:"configs,omitempty"`
			DisableExporter *bool     `tfsdk:"disable_exporter" json:"disableExporter,omitempty"`
			EnabledLogs     *[]string `tfsdk:"enabled_logs" json:"enabledLogs,omitempty"`
			Instances       *[]struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Env         *[]struct {
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
				Image        *string            `tfsdk:"image" json:"image,omitempty"`
				Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
				Name         *string            `tfsdk:"name" json:"name,omitempty"`
				NodeName     *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources    *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				Tolerations *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				VolumeClaimTemplates *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Spec *struct {
						AccessModes *map[string]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
						Resources   *struct {
							Claims *[]struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"claims" json:"claims,omitempty"`
							Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
							Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
						StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
						VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
				VolumeMounts *[]struct {
					MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
					Name             *string `tfsdk:"name" json:"name,omitempty"`
					ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
				} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
				Volumes *[]struct {
					AwsElasticBlockStore *struct {
						FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						VolumeID  *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
					} `tfsdk:"aws_elastic_block_store" json:"awsElasticBlockStore,omitempty"`
					AzureDisk *struct {
						CachingMode *string `tfsdk:"caching_mode" json:"cachingMode,omitempty"`
						DiskName    *string `tfsdk:"disk_name" json:"diskName,omitempty"`
						DiskURI     *string `tfsdk:"disk_uri" json:"diskURI,omitempty"`
						FsType      *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
						ReadOnly    *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"azure_disk" json:"azureDisk,omitempty"`
					AzureFile *struct {
						ReadOnly   *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						ShareName  *string `tfsdk:"share_name" json:"shareName,omitempty"`
					} `tfsdk:"azure_file" json:"azureFile,omitempty"`
					Cephfs *struct {
						Monitors   *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
						Path       *string   `tfsdk:"path" json:"path,omitempty"`
						ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretFile *string   `tfsdk:"secret_file" json:"secretFile,omitempty"`
						SecretRef  *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						User *string `tfsdk:"user" json:"user,omitempty"`
					} `tfsdk:"cephfs" json:"cephfs,omitempty"`
					Cinder *struct {
						FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
					} `tfsdk:"cinder" json:"cinder,omitempty"`
					ConfigMap *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map" json:"configMap,omitempty"`
					Csi *struct {
						Driver               *string `tfsdk:"driver" json:"driver,omitempty"`
						FsType               *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						NodePublishSecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"node_publish_secret_ref" json:"nodePublishSecretRef,omitempty"`
						ReadOnly         *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
						VolumeAttributes *map[string]string `tfsdk:"volume_attributes" json:"volumeAttributes,omitempty"`
					} `tfsdk:"csi" json:"csi,omitempty"`
					DownwardAPI *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
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
					EmptyDir *struct {
						Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
						SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
					} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
					Ephemeral *struct {
						VolumeClaimTemplate *struct {
							Metadata *struct {
								Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
								Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
								Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
								Name        *string            `tfsdk:"name" json:"name,omitempty"`
								Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
							} `tfsdk:"metadata" json:"metadata,omitempty"`
							Spec *struct {
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
									Claims *[]struct {
										Name *string `tfsdk:"name" json:"name,omitempty"`
									} `tfsdk:"claims" json:"claims,omitempty"`
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
								StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
								VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
								VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
							} `tfsdk:"spec" json:"spec,omitempty"`
						} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
					} `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
					Fc *struct {
						FsType     *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
						Lun        *int64    `tfsdk:"lun" json:"lun,omitempty"`
						ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
						TargetWWNs *[]string `tfsdk:"target_ww_ns" json:"targetWWNs,omitempty"`
						Wwids      *[]string `tfsdk:"wwids" json:"wwids,omitempty"`
					} `tfsdk:"fc" json:"fc,omitempty"`
					FlexVolume *struct {
						Driver    *string            `tfsdk:"driver" json:"driver,omitempty"`
						FsType    *string            `tfsdk:"fs_type" json:"fsType,omitempty"`
						Options   *map[string]string `tfsdk:"options" json:"options,omitempty"`
						ReadOnly  *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					} `tfsdk:"flex_volume" json:"flexVolume,omitempty"`
					Flocker *struct {
						DatasetName *string `tfsdk:"dataset_name" json:"datasetName,omitempty"`
						DatasetUUID *string `tfsdk:"dataset_uuid" json:"datasetUUID,omitempty"`
					} `tfsdk:"flocker" json:"flocker,omitempty"`
					GcePersistentDisk *struct {
						FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
						PdName    *string `tfsdk:"pd_name" json:"pdName,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"gce_persistent_disk" json:"gcePersistentDisk,omitempty"`
					GitRepo *struct {
						Directory  *string `tfsdk:"directory" json:"directory,omitempty"`
						Repository *string `tfsdk:"repository" json:"repository,omitempty"`
						Revision   *string `tfsdk:"revision" json:"revision,omitempty"`
					} `tfsdk:"git_repo" json:"gitRepo,omitempty"`
					Glusterfs *struct {
						Endpoints *string `tfsdk:"endpoints" json:"endpoints,omitempty"`
						Path      *string `tfsdk:"path" json:"path,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"glusterfs" json:"glusterfs,omitempty"`
					HostPath *struct {
						Path *string `tfsdk:"path" json:"path,omitempty"`
						Type *string `tfsdk:"type" json:"type,omitempty"`
					} `tfsdk:"host_path" json:"hostPath,omitempty"`
					Iscsi *struct {
						ChapAuthDiscovery *bool     `tfsdk:"chap_auth_discovery" json:"chapAuthDiscovery,omitempty"`
						ChapAuthSession   *bool     `tfsdk:"chap_auth_session" json:"chapAuthSession,omitempty"`
						FsType            *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
						InitiatorName     *string   `tfsdk:"initiator_name" json:"initiatorName,omitempty"`
						Iqn               *string   `tfsdk:"iqn" json:"iqn,omitempty"`
						IscsiInterface    *string   `tfsdk:"iscsi_interface" json:"iscsiInterface,omitempty"`
						Lun               *int64    `tfsdk:"lun" json:"lun,omitempty"`
						Portals           *[]string `tfsdk:"portals" json:"portals,omitempty"`
						ReadOnly          *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef         *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						TargetPortal *string `tfsdk:"target_portal" json:"targetPortal,omitempty"`
					} `tfsdk:"iscsi" json:"iscsi,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Nfs  *struct {
						Path     *string `tfsdk:"path" json:"path,omitempty"`
						ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						Server   *string `tfsdk:"server" json:"server,omitempty"`
					} `tfsdk:"nfs" json:"nfs,omitempty"`
					PersistentVolumeClaim *struct {
						ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
					PhotonPersistentDisk *struct {
						FsType *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						PdID   *string `tfsdk:"pd_id" json:"pdID,omitempty"`
					} `tfsdk:"photon_persistent_disk" json:"photonPersistentDisk,omitempty"`
					PortworxVolume *struct {
						FsType   *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
					} `tfsdk:"portworx_volume" json:"portworxVolume,omitempty"`
					Projected *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Sources     *[]struct {
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
					} `tfsdk:"projected" json:"projected,omitempty"`
					Quobyte *struct {
						Group    *string `tfsdk:"group" json:"group,omitempty"`
						ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						Registry *string `tfsdk:"registry" json:"registry,omitempty"`
						Tenant   *string `tfsdk:"tenant" json:"tenant,omitempty"`
						User     *string `tfsdk:"user" json:"user,omitempty"`
						Volume   *string `tfsdk:"volume" json:"volume,omitempty"`
					} `tfsdk:"quobyte" json:"quobyte,omitempty"`
					Rbd *struct {
						FsType    *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
						Image     *string   `tfsdk:"image" json:"image,omitempty"`
						Keyring   *string   `tfsdk:"keyring" json:"keyring,omitempty"`
						Monitors  *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
						Pool      *string   `tfsdk:"pool" json:"pool,omitempty"`
						ReadOnly  *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						User *string `tfsdk:"user" json:"user,omitempty"`
					} `tfsdk:"rbd" json:"rbd,omitempty"`
					ScaleIO *struct {
						FsType           *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						Gateway          *string `tfsdk:"gateway" json:"gateway,omitempty"`
						ProtectionDomain *string `tfsdk:"protection_domain" json:"protectionDomain,omitempty"`
						ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef        *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						SslEnabled  *bool   `tfsdk:"ssl_enabled" json:"sslEnabled,omitempty"`
						StorageMode *string `tfsdk:"storage_mode" json:"storageMode,omitempty"`
						StoragePool *string `tfsdk:"storage_pool" json:"storagePool,omitempty"`
						System      *string `tfsdk:"system" json:"system,omitempty"`
						VolumeName  *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
					} `tfsdk:"scale_io" json:"scaleIO,omitempty"`
					Secret *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
						SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					Storageos *struct {
						FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SecretRef *struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						VolumeName      *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
						VolumeNamespace *string `tfsdk:"volume_namespace" json:"volumeNamespace,omitempty"`
					} `tfsdk:"storageos" json:"storageos,omitempty"`
					VsphereVolume *struct {
						FsType            *string `tfsdk:"fs_type" json:"fsType,omitempty"`
						StoragePolicyID   *string `tfsdk:"storage_policy_id" json:"storagePolicyID,omitempty"`
						StoragePolicyName *string `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
						VolumePath        *string `tfsdk:"volume_path" json:"volumePath,omitempty"`
					} `tfsdk:"vsphere_volume" json:"vsphereVolume,omitempty"`
				} `tfsdk:"volumes" json:"volumes,omitempty"`
			} `tfsdk:"instances" json:"instances,omitempty"`
			Issuer *struct {
				Name      *string `tfsdk:"name" json:"name,omitempty"`
				SecretRef *struct {
					Ca   *string `tfsdk:"ca" json:"ca,omitempty"`
					Cert *string `tfsdk:"cert" json:"cert,omitempty"`
					Key  *string `tfsdk:"key" json:"key,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"issuer" json:"issuer,omitempty"`
			Monitor          *bool     `tfsdk:"monitor" json:"monitor,omitempty"`
			Name             *string   `tfsdk:"name" json:"name,omitempty"`
			OfflineInstances *[]string `tfsdk:"offline_instances" json:"offlineInstances,omitempty"`
			Replicas         *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources        *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			SchedulingPolicy *struct {
				Affinity *struct {
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
				NodeName      *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
				NodeSelector  *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				SchedulerName *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
				Tolerations   *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
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
			} `tfsdk:"scheduling_policy" json:"schedulingPolicy,omitempty"`
			ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
			ServiceRefs        *[]struct {
				Cluster                *string `tfsdk:"cluster" json:"cluster,omitempty"`
				ClusterServiceSelector *struct {
					Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
					Credential *struct {
						Component *string `tfsdk:"component" json:"component,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"credential" json:"credential,omitempty"`
					Service *struct {
						Component *string `tfsdk:"component" json:"component,omitempty"`
						Port      *string `tfsdk:"port" json:"port,omitempty"`
						Service   *string `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"service" json:"service,omitempty"`
				} `tfsdk:"cluster_service_selector" json:"clusterServiceSelector,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				Namespace         *string `tfsdk:"namespace" json:"namespace,omitempty"`
				ServiceDescriptor *string `tfsdk:"service_descriptor" json:"serviceDescriptor,omitempty"`
			} `tfsdk:"service_refs" json:"serviceRefs,omitempty"`
			ServiceVersion *string `tfsdk:"service_version" json:"serviceVersion,omitempty"`
			Services       *[]struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Name        *string            `tfsdk:"name" json:"name,omitempty"`
				PodService  *bool              `tfsdk:"pod_service" json:"podService,omitempty"`
				ServiceType *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
			SwitchPolicy *struct {
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"switch_policy" json:"switchPolicy,omitempty"`
			Tls              *bool              `tfsdk:"tls" json:"tls,omitempty"`
			Tolerations      *map[string]string `tfsdk:"tolerations" json:"tolerations,omitempty"`
			UpdateStrategy   *string            `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
			UserResourceRefs *struct {
				ConfigMapRefs *[]struct {
					AsVolumeFrom *[]string `tfsdk:"as_volume_from" json:"asVolumeFrom,omitempty"`
					ConfigMap    *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map" json:"configMap,omitempty"`
					MountPoint *string `tfsdk:"mount_point" json:"mountPoint,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					SubPath    *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				} `tfsdk:"config_map_refs" json:"configMapRefs,omitempty"`
				SecretRefs *[]struct {
					AsVolumeFrom *[]string `tfsdk:"as_volume_from" json:"asVolumeFrom,omitempty"`
					MountPoint   *string   `tfsdk:"mount_point" json:"mountPoint,omitempty"`
					Name         *string   `tfsdk:"name" json:"name,omitempty"`
					Secret       *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
						SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					} `tfsdk:"secret" json:"secret,omitempty"`
					SubPath *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
			} `tfsdk:"user_resource_refs" json:"userResourceRefs,omitempty"`
			VolumeClaimTemplates *[]struct {
				Name *string `tfsdk:"name" json:"name,omitempty"`
				Spec *struct {
					AccessModes *map[string]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
					Resources   *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
					VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				} `tfsdk:"spec" json:"spec,omitempty"`
			} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		} `tfsdk:"component_specs" json:"componentSpecs,omitempty"`
		Network *struct {
			HostNetworkAccessible *bool `tfsdk:"host_network_accessible" json:"hostNetworkAccessible,omitempty"`
			PubliclyAccessible    *bool `tfsdk:"publicly_accessible" json:"publiclyAccessible,omitempty"`
		} `tfsdk:"network" json:"network,omitempty"`
		Replicas  *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		Resources *struct {
			Cpu    *string `tfsdk:"cpu" json:"cpu,omitempty"`
			Memory *string `tfsdk:"memory" json:"memory,omitempty"`
		} `tfsdk:"resources" json:"resources,omitempty"`
		RuntimeClassName *string `tfsdk:"runtime_class_name" json:"runtimeClassName,omitempty"`
		SchedulingPolicy *struct {
			Affinity *struct {
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
			NodeName      *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
			NodeSelector  *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			SchedulerName *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
			Tolerations   *[]struct {
				Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
				Key               *string `tfsdk:"key" json:"key,omitempty"`
				Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
				TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
				Value             *string `tfsdk:"value" json:"value,omitempty"`
			} `tfsdk:"tolerations" json:"tolerations,omitempty"`
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
		} `tfsdk:"scheduling_policy" json:"schedulingPolicy,omitempty"`
		Services      *map[string]string `tfsdk:"services" json:"services,omitempty"`
		ShardingSpecs *[]struct {
			Name     *string `tfsdk:"name" json:"name,omitempty"`
			Shards   *int64  `tfsdk:"shards" json:"shards,omitempty"`
			Template *struct {
				Affinity *struct {
					NodeLabels      *map[string]string `tfsdk:"node_labels" json:"nodeLabels,omitempty"`
					PodAntiAffinity *string            `tfsdk:"pod_anti_affinity" json:"podAntiAffinity,omitempty"`
					Tenancy         *string            `tfsdk:"tenancy" json:"tenancy,omitempty"`
					TopologyKeys    *[]string          `tfsdk:"topology_keys" json:"topologyKeys,omitempty"`
				} `tfsdk:"affinity" json:"affinity,omitempty"`
				ComponentDef    *string `tfsdk:"component_def" json:"componentDef,omitempty"`
				ComponentDefRef *string `tfsdk:"component_def_ref" json:"componentDefRef,omitempty"`
				Configs         *[]struct {
					ConfigMap *struct {
						DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
						Items       *[]struct {
							Key  *string `tfsdk:"key" json:"key,omitempty"`
							Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
							Path *string `tfsdk:"path" json:"path,omitempty"`
						} `tfsdk:"items" json:"items,omitempty"`
						Name     *string `tfsdk:"name" json:"name,omitempty"`
						Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
					} `tfsdk:"config_map" json:"configMap,omitempty"`
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"configs" json:"configs,omitempty"`
				DisableExporter *bool     `tfsdk:"disable_exporter" json:"disableExporter,omitempty"`
				EnabledLogs     *[]string `tfsdk:"enabled_logs" json:"enabledLogs,omitempty"`
				Instances       *[]struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Env         *[]struct {
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
					Image        *string            `tfsdk:"image" json:"image,omitempty"`
					Labels       *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
					Name         *string            `tfsdk:"name" json:"name,omitempty"`
					NodeName     *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
					NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					Replicas     *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
					Resources    *struct {
						Claims *[]struct {
							Name *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"claims" json:"claims,omitempty"`
						Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
						Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
					} `tfsdk:"resources" json:"resources,omitempty"`
					Tolerations *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
					VolumeClaimTemplates *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Spec *struct {
							AccessModes *map[string]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
							Resources   *struct {
								Claims *[]struct {
									Name *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"claims" json:"claims,omitempty"`
								Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
								Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
							} `tfsdk:"resources" json:"resources,omitempty"`
							StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
							VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
						} `tfsdk:"spec" json:"spec,omitempty"`
					} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
					VolumeMounts *[]struct {
						MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
						MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
						Name             *string `tfsdk:"name" json:"name,omitempty"`
						ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
						SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
					} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
					Volumes *[]struct {
						AwsElasticBlockStore *struct {
							FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							VolumeID  *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
						} `tfsdk:"aws_elastic_block_store" json:"awsElasticBlockStore,omitempty"`
						AzureDisk *struct {
							CachingMode *string `tfsdk:"caching_mode" json:"cachingMode,omitempty"`
							DiskName    *string `tfsdk:"disk_name" json:"diskName,omitempty"`
							DiskURI     *string `tfsdk:"disk_uri" json:"diskURI,omitempty"`
							FsType      *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							Kind        *string `tfsdk:"kind" json:"kind,omitempty"`
							ReadOnly    *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"azure_disk" json:"azureDisk,omitempty"`
						AzureFile *struct {
							ReadOnly   *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
							ShareName  *string `tfsdk:"share_name" json:"shareName,omitempty"`
						} `tfsdk:"azure_file" json:"azureFile,omitempty"`
						Cephfs *struct {
							Monitors   *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
							Path       *string   `tfsdk:"path" json:"path,omitempty"`
							ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretFile *string   `tfsdk:"secret_file" json:"secretFile,omitempty"`
							SecretRef  *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							User *string `tfsdk:"user" json:"user,omitempty"`
						} `tfsdk:"cephfs" json:"cephfs,omitempty"`
						Cinder *struct {
							FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
						} `tfsdk:"cinder" json:"cinder,omitempty"`
						ConfigMap *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						Csi *struct {
							Driver               *string `tfsdk:"driver" json:"driver,omitempty"`
							FsType               *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							NodePublishSecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"node_publish_secret_ref" json:"nodePublishSecretRef,omitempty"`
							ReadOnly         *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
							VolumeAttributes *map[string]string `tfsdk:"volume_attributes" json:"volumeAttributes,omitempty"`
						} `tfsdk:"csi" json:"csi,omitempty"`
						DownwardAPI *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
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
						EmptyDir *struct {
							Medium    *string `tfsdk:"medium" json:"medium,omitempty"`
							SizeLimit *string `tfsdk:"size_limit" json:"sizeLimit,omitempty"`
						} `tfsdk:"empty_dir" json:"emptyDir,omitempty"`
						Ephemeral *struct {
							VolumeClaimTemplate *struct {
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
									Finalizers  *[]string          `tfsdk:"finalizers" json:"finalizers,omitempty"`
									Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
									Name        *string            `tfsdk:"name" json:"name,omitempty"`
									Namespace   *string            `tfsdk:"namespace" json:"namespace,omitempty"`
								} `tfsdk:"metadata" json:"metadata,omitempty"`
								Spec *struct {
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
										Claims *[]struct {
											Name *string `tfsdk:"name" json:"name,omitempty"`
										} `tfsdk:"claims" json:"claims,omitempty"`
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
									StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
									VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
									VolumeName       *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
								} `tfsdk:"spec" json:"spec,omitempty"`
							} `tfsdk:"volume_claim_template" json:"volumeClaimTemplate,omitempty"`
						} `tfsdk:"ephemeral" json:"ephemeral,omitempty"`
						Fc *struct {
							FsType     *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
							Lun        *int64    `tfsdk:"lun" json:"lun,omitempty"`
							ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
							TargetWWNs *[]string `tfsdk:"target_ww_ns" json:"targetWWNs,omitempty"`
							Wwids      *[]string `tfsdk:"wwids" json:"wwids,omitempty"`
						} `tfsdk:"fc" json:"fc,omitempty"`
						FlexVolume *struct {
							Driver    *string            `tfsdk:"driver" json:"driver,omitempty"`
							FsType    *string            `tfsdk:"fs_type" json:"fsType,omitempty"`
							Options   *map[string]string `tfsdk:"options" json:"options,omitempty"`
							ReadOnly  *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
						} `tfsdk:"flex_volume" json:"flexVolume,omitempty"`
						Flocker *struct {
							DatasetName *string `tfsdk:"dataset_name" json:"datasetName,omitempty"`
							DatasetUUID *string `tfsdk:"dataset_uuid" json:"datasetUUID,omitempty"`
						} `tfsdk:"flocker" json:"flocker,omitempty"`
						GcePersistentDisk *struct {
							FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							Partition *int64  `tfsdk:"partition" json:"partition,omitempty"`
							PdName    *string `tfsdk:"pd_name" json:"pdName,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"gce_persistent_disk" json:"gcePersistentDisk,omitempty"`
						GitRepo *struct {
							Directory  *string `tfsdk:"directory" json:"directory,omitempty"`
							Repository *string `tfsdk:"repository" json:"repository,omitempty"`
							Revision   *string `tfsdk:"revision" json:"revision,omitempty"`
						} `tfsdk:"git_repo" json:"gitRepo,omitempty"`
						Glusterfs *struct {
							Endpoints *string `tfsdk:"endpoints" json:"endpoints,omitempty"`
							Path      *string `tfsdk:"path" json:"path,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"glusterfs" json:"glusterfs,omitempty"`
						HostPath *struct {
							Path *string `tfsdk:"path" json:"path,omitempty"`
							Type *string `tfsdk:"type" json:"type,omitempty"`
						} `tfsdk:"host_path" json:"hostPath,omitempty"`
						Iscsi *struct {
							ChapAuthDiscovery *bool     `tfsdk:"chap_auth_discovery" json:"chapAuthDiscovery,omitempty"`
							ChapAuthSession   *bool     `tfsdk:"chap_auth_session" json:"chapAuthSession,omitempty"`
							FsType            *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
							InitiatorName     *string   `tfsdk:"initiator_name" json:"initiatorName,omitempty"`
							Iqn               *string   `tfsdk:"iqn" json:"iqn,omitempty"`
							IscsiInterface    *string   `tfsdk:"iscsi_interface" json:"iscsiInterface,omitempty"`
							Lun               *int64    `tfsdk:"lun" json:"lun,omitempty"`
							Portals           *[]string `tfsdk:"portals" json:"portals,omitempty"`
							ReadOnly          *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef         *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							TargetPortal *string `tfsdk:"target_portal" json:"targetPortal,omitempty"`
						} `tfsdk:"iscsi" json:"iscsi,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
						Nfs  *struct {
							Path     *string `tfsdk:"path" json:"path,omitempty"`
							ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							Server   *string `tfsdk:"server" json:"server,omitempty"`
						} `tfsdk:"nfs" json:"nfs,omitempty"`
						PersistentVolumeClaim *struct {
							ClaimName *string `tfsdk:"claim_name" json:"claimName,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
						} `tfsdk:"persistent_volume_claim" json:"persistentVolumeClaim,omitempty"`
						PhotonPersistentDisk *struct {
							FsType *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							PdID   *string `tfsdk:"pd_id" json:"pdID,omitempty"`
						} `tfsdk:"photon_persistent_disk" json:"photonPersistentDisk,omitempty"`
						PortworxVolume *struct {
							FsType   *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
						} `tfsdk:"portworx_volume" json:"portworxVolume,omitempty"`
						Projected *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Sources     *[]struct {
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
						} `tfsdk:"projected" json:"projected,omitempty"`
						Quobyte *struct {
							Group    *string `tfsdk:"group" json:"group,omitempty"`
							ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							Registry *string `tfsdk:"registry" json:"registry,omitempty"`
							Tenant   *string `tfsdk:"tenant" json:"tenant,omitempty"`
							User     *string `tfsdk:"user" json:"user,omitempty"`
							Volume   *string `tfsdk:"volume" json:"volume,omitempty"`
						} `tfsdk:"quobyte" json:"quobyte,omitempty"`
						Rbd *struct {
							FsType    *string   `tfsdk:"fs_type" json:"fsType,omitempty"`
							Image     *string   `tfsdk:"image" json:"image,omitempty"`
							Keyring   *string   `tfsdk:"keyring" json:"keyring,omitempty"`
							Monitors  *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
							Pool      *string   `tfsdk:"pool" json:"pool,omitempty"`
							ReadOnly  *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							User *string `tfsdk:"user" json:"user,omitempty"`
						} `tfsdk:"rbd" json:"rbd,omitempty"`
						ScaleIO *struct {
							FsType           *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							Gateway          *string `tfsdk:"gateway" json:"gateway,omitempty"`
							ProtectionDomain *string `tfsdk:"protection_domain" json:"protectionDomain,omitempty"`
							ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef        *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							SslEnabled  *bool   `tfsdk:"ssl_enabled" json:"sslEnabled,omitempty"`
							StorageMode *string `tfsdk:"storage_mode" json:"storageMode,omitempty"`
							StoragePool *string `tfsdk:"storage_pool" json:"storagePool,omitempty"`
							System      *string `tfsdk:"system" json:"system,omitempty"`
							VolumeName  *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
						} `tfsdk:"scale_io" json:"scaleIO,omitempty"`
						Secret *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						Storageos *struct {
							FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
							SecretRef *struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
							VolumeName      *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
							VolumeNamespace *string `tfsdk:"volume_namespace" json:"volumeNamespace,omitempty"`
						} `tfsdk:"storageos" json:"storageos,omitempty"`
						VsphereVolume *struct {
							FsType            *string `tfsdk:"fs_type" json:"fsType,omitempty"`
							StoragePolicyID   *string `tfsdk:"storage_policy_id" json:"storagePolicyID,omitempty"`
							StoragePolicyName *string `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
							VolumePath        *string `tfsdk:"volume_path" json:"volumePath,omitempty"`
						} `tfsdk:"vsphere_volume" json:"vsphereVolume,omitempty"`
					} `tfsdk:"volumes" json:"volumes,omitempty"`
				} `tfsdk:"instances" json:"instances,omitempty"`
				Issuer *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					SecretRef *struct {
						Ca   *string `tfsdk:"ca" json:"ca,omitempty"`
						Cert *string `tfsdk:"cert" json:"cert,omitempty"`
						Key  *string `tfsdk:"key" json:"key,omitempty"`
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
				} `tfsdk:"issuer" json:"issuer,omitempty"`
				Monitor          *bool     `tfsdk:"monitor" json:"monitor,omitempty"`
				Name             *string   `tfsdk:"name" json:"name,omitempty"`
				OfflineInstances *[]string `tfsdk:"offline_instances" json:"offlineInstances,omitempty"`
				Replicas         *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
				Resources        *struct {
					Claims *[]struct {
						Name *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"claims" json:"claims,omitempty"`
					Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
					Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
				} `tfsdk:"resources" json:"resources,omitempty"`
				SchedulingPolicy *struct {
					Affinity *struct {
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
					NodeName      *string            `tfsdk:"node_name" json:"nodeName,omitempty"`
					NodeSelector  *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					SchedulerName *string            `tfsdk:"scheduler_name" json:"schedulerName,omitempty"`
					Tolerations   *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
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
				} `tfsdk:"scheduling_policy" json:"schedulingPolicy,omitempty"`
				ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
				ServiceRefs        *[]struct {
					Cluster                *string `tfsdk:"cluster" json:"cluster,omitempty"`
					ClusterServiceSelector *struct {
						Cluster    *string `tfsdk:"cluster" json:"cluster,omitempty"`
						Credential *struct {
							Component *string `tfsdk:"component" json:"component,omitempty"`
							Name      *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"credential" json:"credential,omitempty"`
						Service *struct {
							Component *string `tfsdk:"component" json:"component,omitempty"`
							Port      *string `tfsdk:"port" json:"port,omitempty"`
							Service   *string `tfsdk:"service" json:"service,omitempty"`
						} `tfsdk:"service" json:"service,omitempty"`
					} `tfsdk:"cluster_service_selector" json:"clusterServiceSelector,omitempty"`
					Name              *string `tfsdk:"name" json:"name,omitempty"`
					Namespace         *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ServiceDescriptor *string `tfsdk:"service_descriptor" json:"serviceDescriptor,omitempty"`
				} `tfsdk:"service_refs" json:"serviceRefs,omitempty"`
				ServiceVersion *string `tfsdk:"service_version" json:"serviceVersion,omitempty"`
				Services       *[]struct {
					Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
					Name        *string            `tfsdk:"name" json:"name,omitempty"`
					PodService  *bool              `tfsdk:"pod_service" json:"podService,omitempty"`
					ServiceType *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
				} `tfsdk:"services" json:"services,omitempty"`
				SwitchPolicy *struct {
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"switch_policy" json:"switchPolicy,omitempty"`
				Tls              *bool              `tfsdk:"tls" json:"tls,omitempty"`
				Tolerations      *map[string]string `tfsdk:"tolerations" json:"tolerations,omitempty"`
				UpdateStrategy   *string            `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
				UserResourceRefs *struct {
					ConfigMapRefs *[]struct {
						AsVolumeFrom *[]string `tfsdk:"as_volume_from" json:"asVolumeFrom,omitempty"`
						ConfigMap    *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
							Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
						} `tfsdk:"config_map" json:"configMap,omitempty"`
						MountPoint *string `tfsdk:"mount_point" json:"mountPoint,omitempty"`
						Name       *string `tfsdk:"name" json:"name,omitempty"`
						SubPath    *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					} `tfsdk:"config_map_refs" json:"configMapRefs,omitempty"`
					SecretRefs *[]struct {
						AsVolumeFrom *[]string `tfsdk:"as_volume_from" json:"asVolumeFrom,omitempty"`
						MountPoint   *string   `tfsdk:"mount_point" json:"mountPoint,omitempty"`
						Name         *string   `tfsdk:"name" json:"name,omitempty"`
						Secret       *struct {
							DefaultMode *int64 `tfsdk:"default_mode" json:"defaultMode,omitempty"`
							Items       *[]struct {
								Key  *string `tfsdk:"key" json:"key,omitempty"`
								Mode *int64  `tfsdk:"mode" json:"mode,omitempty"`
								Path *string `tfsdk:"path" json:"path,omitempty"`
							} `tfsdk:"items" json:"items,omitempty"`
							Optional   *bool   `tfsdk:"optional" json:"optional,omitempty"`
							SecretName *string `tfsdk:"secret_name" json:"secretName,omitempty"`
						} `tfsdk:"secret" json:"secret,omitempty"`
						SubPath *string `tfsdk:"sub_path" json:"subPath,omitempty"`
					} `tfsdk:"secret_refs" json:"secretRefs,omitempty"`
				} `tfsdk:"user_resource_refs" json:"userResourceRefs,omitempty"`
				VolumeClaimTemplates *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
					Spec *struct {
						AccessModes *map[string]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
						Resources   *struct {
							Claims *[]struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"claims" json:"claims,omitempty"`
							Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
							Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
						} `tfsdk:"resources" json:"resources,omitempty"`
						StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
						VolumeMode       *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
					} `tfsdk:"spec" json:"spec,omitempty"`
				} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
			} `tfsdk:"template" json:"template,omitempty"`
		} `tfsdk:"sharding_specs" json:"shardingSpecs,omitempty"`
		Storage *struct {
			Size *string `tfsdk:"size" json:"size,omitempty"`
		} `tfsdk:"storage" json:"storage,omitempty"`
		Tenancy           *string            `tfsdk:"tenancy" json:"tenancy,omitempty"`
		TerminationPolicy *string            `tfsdk:"termination_policy" json:"terminationPolicy,omitempty"`
		Tolerations       *map[string]string `tfsdk:"tolerations" json:"tolerations,omitempty"`
		Topology          *string            `tfsdk:"topology" json:"topology,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoClusterV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_cluster_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoClusterV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "Cluster offers a unified management interface for a wide variety of database and storage systems:  - Relational databases: MySQL, PostgreSQL, MariaDB - NoSQL databases: Redis, MongoDB - KV stores: ZooKeeper, etcd - Analytics systems: ElasticSearch, OpenSearch, ClickHouse, Doris, StarRocks, Solr - Message queues: Kafka, Pulsar - Distributed SQL: TiDB, OceanBase - Vector databases: Qdrant, Milvus, Weaviate - Object storage: Minio  KubeBlocks utilizes an abstraction layer to encapsulate the characteristics of these diverse systems. A Cluster is composed of multiple Components, each defined by vendors or KubeBlocks Addon developers via ComponentDefinition, arranged in Directed Acyclic Graph (DAG) topologies. The topologies, defined in a ClusterDefinition, coordinate reconciliation across Cluster's lifecycle phases: Creating, Running, Updating, Stopping, Stopped, Deleting. Lifecycle management ensures that each Component operates in harmony, executing appropriate actions at each lifecycle stage.  For sharded-nothing architecture, the Cluster supports managing multiple shards, each shard managed by a separate Component, supporting dynamic resharding.  The Cluster object is aimed to maintain the overall integrity and availability of a database cluster, serves as the central control point, abstracting the complexity of multiple-component management, and providing a unified interface for cluster-wide operations.",
		MarkdownDescription: "Cluster offers a unified management interface for a wide variety of database and storage systems:  - Relational databases: MySQL, PostgreSQL, MariaDB - NoSQL databases: Redis, MongoDB - KV stores: ZooKeeper, etcd - Analytics systems: ElasticSearch, OpenSearch, ClickHouse, Doris, StarRocks, Solr - Message queues: Kafka, Pulsar - Distributed SQL: TiDB, OceanBase - Vector databases: Qdrant, Milvus, Weaviate - Object storage: Minio  KubeBlocks utilizes an abstraction layer to encapsulate the characteristics of these diverse systems. A Cluster is composed of multiple Components, each defined by vendors or KubeBlocks Addon developers via ComponentDefinition, arranged in Directed Acyclic Graph (DAG) topologies. The topologies, defined in a ClusterDefinition, coordinate reconciliation across Cluster's lifecycle phases: Creating, Running, Updating, Stopping, Stopped, Deleting. Lifecycle management ensures that each Component operates in harmony, executing appropriate actions at each lifecycle stage.  For sharded-nothing architecture, the Cluster supports managing multiple shards, each shard managed by a separate Component, supporting dynamic resharding.  The Cluster object is aimed to maintain the overall integrity and availability of a database cluster, serves as the central control point, abstracting the complexity of multiple-component management, and providing a unified interface for cluster-wide operations.",
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
				Description:         "ClusterSpec defines the desired state of Cluster.",
				MarkdownDescription: "ClusterSpec defines the desired state of Cluster.",
				Attributes: map[string]schema.Attribute{
					"affinity": schema.SingleNestedAttribute{
						Description:         "Defines a set of node affinity scheduling rules for the Cluster's Pods. This field helps control the placement of Pods on nodes within the Cluster.  Deprecated since v0.10. Use the 'schedulingPolicy' field instead.",
						MarkdownDescription: "Defines a set of node affinity scheduling rules for the Cluster's Pods. This field helps control the placement of Pods on nodes within the Cluster.  Deprecated since v0.10. Use the 'schedulingPolicy' field instead.",
						Attributes: map[string]schema.Attribute{
							"node_labels": schema.MapAttribute{
								Description:         "Indicates the node labels that must be present on nodes for pods to be scheduled on them. It is a map where the keys are the label keys and the values are the corresponding label values. Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.  For example, if NodeLabels is set to {'nodeType': 'ssd', 'environment': 'production'}, pods will only be scheduled on nodes that have both the 'nodeType' label with value 'ssd' and the 'environment' label with value 'production'.  This field allows users to control Pod placement based on specific node labels. It can be used to ensure that Pods are scheduled on nodes with certain characteristics, such as specific hardware (e.g., SSD), environment (e.g., production, staging), or any other custom labels assigned to nodes.",
								MarkdownDescription: "Indicates the node labels that must be present on nodes for pods to be scheduled on them. It is a map where the keys are the label keys and the values are the corresponding label values. Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.  For example, if NodeLabels is set to {'nodeType': 'ssd', 'environment': 'production'}, pods will only be scheduled on nodes that have both the 'nodeType' label with value 'ssd' and the 'environment' label with value 'production'.  This field allows users to control Pod placement based on specific node labels. It can be used to ensure that Pods are scheduled on nodes with certain characteristics, such as specific hardware (e.g., SSD), environment (e.g., production, staging), or any other custom labels assigned to nodes.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_anti_affinity": schema.StringAttribute{
								Description:         "Specifies the anti-affinity level of Pods within a Component. It determines how pods should be spread across nodes to improve availability and performance. It can have the following values: 'Preferred' and 'Required'. The default value is 'Preferred'.",
								MarkdownDescription: "Specifies the anti-affinity level of Pods within a Component. It determines how pods should be spread across nodes to improve availability and performance. It can have the following values: 'Preferred' and 'Required'. The default value is 'Preferred'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Preferred", "Required"),
								},
							},

							"tenancy": schema.StringAttribute{
								Description:         "Determines the level of resource isolation between Pods. It can have the following values: 'SharedNode' and 'DedicatedNode'.  - SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s. - DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node. In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node. Which provides a higher level of isolation and resource guarantee for Pods.  The default value is 'SharedNode'.",
								MarkdownDescription: "Determines the level of resource isolation between Pods. It can have the following values: 'SharedNode' and 'DedicatedNode'.  - SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s. - DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node. In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node. Which provides a higher level of isolation and resource guarantee for Pods.  The default value is 'SharedNode'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("SharedNode", "DedicatedNode"),
								},
							},

							"topology_keys": schema.ListAttribute{
								Description:         "Represents the key of node labels used to define the topology domain for Pod anti-affinity and Pod spread constraints.  In K8s, a topology domain is a set of nodes that have the same value for a specific label key. Nodes with labels containing any of the specified TopologyKeys and identical values are considered to be in the same topology domain.  Note: The concept of topology in the context of K8s TopologyKeys is different from the concept of topology in the ClusterDefinition.  When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule the Pod on nodes with different values for the specified TopologyKeys. This ensures that Pods are spread across different topology domains, promoting high availability and reducing the impact of node failures.  Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey. These keys represent the hostname and zone of a node, respectively. By including these keys in the TopologyKeys list, Pods will be spread across nodes with different hostnames or zones.  In addition to the well-known keys, users can also specify custom label keys as TopologyKeys. This allows for more flexible and custom topology definitions based on the specific needs of the application or environment.  The TopologyKeys field is a slice of strings, where each string represents a label key. The order of the keys in the slice does not matter.",
								MarkdownDescription: "Represents the key of node labels used to define the topology domain for Pod anti-affinity and Pod spread constraints.  In K8s, a topology domain is a set of nodes that have the same value for a specific label key. Nodes with labels containing any of the specified TopologyKeys and identical values are considered to be in the same topology domain.  Note: The concept of topology in the context of K8s TopologyKeys is different from the concept of topology in the ClusterDefinition.  When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule the Pod on nodes with different values for the specified TopologyKeys. This ensures that Pods are spread across different topology domains, promoting high availability and reducing the impact of node failures.  Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey. These keys represent the hostname and zone of a node, respectively. By including these keys in the TopologyKeys list, Pods will be spread across nodes with different hostnames or zones.  In addition to the well-known keys, users can also specify custom label keys as TopologyKeys. This allows for more flexible and custom topology definitions based on the specific needs of the application or environment.  The TopologyKeys field is a slice of strings, where each string represents a label key. The order of the keys in the slice does not matter.",
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

					"availability_policy": schema.StringAttribute{
						Description:         "Describes the availability policy, including zone, node, and none.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						MarkdownDescription: "Describes the availability policy, including zone, node, and none.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("zone", "node", "none"),
						},
					},

					"backup": schema.SingleNestedAttribute{
						Description:         "Specifies the backup configuration of the Cluster.",
						MarkdownDescription: "Specifies the backup configuration of the Cluster.",
						Attributes: map[string]schema.Attribute{
							"cron_expression": schema.StringAttribute{
								Description:         "The cron expression for the schedule. The timezone is in UTC. See https://en.wikipedia.org/wiki/Cron.",
								MarkdownDescription: "The cron expression for the schedule. The timezone is in UTC. See https://en.wikipedia.org/wiki/Cron.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"enabled": schema.BoolAttribute{
								Description:         "Specifies whether automated backup is enabled for the Cluster.",
								MarkdownDescription: "Specifies whether automated backup is enabled for the Cluster.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"method": schema.StringAttribute{
								Description:         "Specifies the backup method to use, as defined in backupPolicy.",
								MarkdownDescription: "Specifies the backup method to use, as defined in backupPolicy.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"pitr_enabled": schema.BoolAttribute{
								Description:         "Specifies whether to enable point-in-time recovery.",
								MarkdownDescription: "Specifies whether to enable point-in-time recovery.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"repo_name": schema.StringAttribute{
								Description:         "Specifies the name of the backupRepo. If not set, the default backupRepo will be used.",
								MarkdownDescription: "Specifies the name of the backupRepo. If not set, the default backupRepo will be used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retention_period": schema.StringAttribute{
								Description:         "Determines the duration to retain backups. Backups older than this period are automatically removed.  For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format:  - years: 	2y - months: 	6mo - days: 		30d - hours: 	12h - minutes: 	30m  You can also combine the above durations. For example: 30d12h30m. Default value is 7d.",
								MarkdownDescription: "Determines the duration to retain backups. Backups older than this period are automatically removed.  For example, RetentionPeriod of '30d' will keep only the backups of last 30 days. Sample duration format:  - years: 	2y - months: 	6mo - days: 		30d - hours: 	12h - minutes: 	30m  You can also combine the above durations. For example: 30d12h30m. Default value is 7d.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"starting_deadline_minutes": schema.Int64Attribute{
								Description:         "Specifies the maximum time in minutes that the system will wait to start a missed backup job. If the scheduled backup time is missed for any reason, the backup job must start within this deadline. Values must be between 0 (immediate execution) and 1440 (one day).",
								MarkdownDescription: "Specifies the maximum time in minutes that the system will wait to start a missed backup job. If the scheduled backup time is missed for any reason, the backup job must start within this deadline. Values must be between 0 (immediate execution) and 1440 (one day).",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(1440),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cluster_definition_ref": schema.StringAttribute{
						Description:         "Specifies the name of the ClusterDefinition to use when creating a Cluster.  This field enables users to create a Cluster based on a specific ClusterDefinition. Which, in conjunction with the 'topology' field, determine:  - The Components to be included in the Cluster. - The sequences in which the Components are created, updated, and terminate.  This facilitates multiple-components management with predefined ClusterDefinition.  Users with advanced requirements can bypass this general setting and specify more precise control over the composition of the Cluster by directly referencing specific ComponentDefinitions for each component within 'componentSpecs[*].componentDef'.  If this field is not provided, each component must be explicitly defined in 'componentSpecs[*].componentDef'.  Note: Once set, this field cannot be modified; it is immutable.",
						MarkdownDescription: "Specifies the name of the ClusterDefinition to use when creating a Cluster.  This field enables users to create a Cluster based on a specific ClusterDefinition. Which, in conjunction with the 'topology' field, determine:  - The Components to be included in the Cluster. - The sequences in which the Components are created, updated, and terminate.  This facilitates multiple-components management with predefined ClusterDefinition.  Users with advanced requirements can bypass this general setting and specify more precise control over the composition of the Cluster by directly referencing specific ComponentDefinitions for each component within 'componentSpecs[*].componentDef'.  If this field is not provided, each component must be explicitly defined in 'componentSpecs[*].componentDef'.  Note: Once set, this field cannot be modified; it is immutable.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
						},
					},

					"cluster_version_ref": schema.StringAttribute{
						Description:         "Refers to the ClusterVersion name.  Deprecated since v0.9, use ComponentVersion instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						MarkdownDescription: "Refers to the ClusterVersion name.  Deprecated since v0.9, use ComponentVersion instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(63),
							stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
						},
					},

					"component_specs": schema.ListNestedAttribute{
						Description:         "Specifies a list of ClusterComponentSpec objects used to define the individual Components that make up a Cluster. This field allows for detailed configuration of each Component within the Cluster.  Note: 'shardingSpecs' and 'componentSpecs' cannot both be empty; at least one must be defined to configure a Cluster.",
						MarkdownDescription: "Specifies a list of ClusterComponentSpec objects used to define the individual Components that make up a Cluster. This field allows for detailed configuration of each Component within the Cluster.  Note: 'shardingSpecs' and 'componentSpecs' cannot both be empty; at least one must be defined to configure a Cluster.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"affinity": schema.SingleNestedAttribute{
									Description:         "Specifies a group of affinity scheduling rules for the Component. It allows users to control how the Component's Pods are scheduled onto nodes in the K8s cluster.  Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
									MarkdownDescription: "Specifies a group of affinity scheduling rules for the Component. It allows users to control how the Component's Pods are scheduled onto nodes in the K8s cluster.  Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
									Attributes: map[string]schema.Attribute{
										"node_labels": schema.MapAttribute{
											Description:         "Indicates the node labels that must be present on nodes for pods to be scheduled on them. It is a map where the keys are the label keys and the values are the corresponding label values. Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.  For example, if NodeLabels is set to {'nodeType': 'ssd', 'environment': 'production'}, pods will only be scheduled on nodes that have both the 'nodeType' label with value 'ssd' and the 'environment' label with value 'production'.  This field allows users to control Pod placement based on specific node labels. It can be used to ensure that Pods are scheduled on nodes with certain characteristics, such as specific hardware (e.g., SSD), environment (e.g., production, staging), or any other custom labels assigned to nodes.",
											MarkdownDescription: "Indicates the node labels that must be present on nodes for pods to be scheduled on them. It is a map where the keys are the label keys and the values are the corresponding label values. Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.  For example, if NodeLabels is set to {'nodeType': 'ssd', 'environment': 'production'}, pods will only be scheduled on nodes that have both the 'nodeType' label with value 'ssd' and the 'environment' label with value 'production'.  This field allows users to control Pod placement based on specific node labels. It can be used to ensure that Pods are scheduled on nodes with certain characteristics, such as specific hardware (e.g., SSD), environment (e.g., production, staging), or any other custom labels assigned to nodes.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"pod_anti_affinity": schema.StringAttribute{
											Description:         "Specifies the anti-affinity level of Pods within a Component. It determines how pods should be spread across nodes to improve availability and performance. It can have the following values: 'Preferred' and 'Required'. The default value is 'Preferred'.",
											MarkdownDescription: "Specifies the anti-affinity level of Pods within a Component. It determines how pods should be spread across nodes to improve availability and performance. It can have the following values: 'Preferred' and 'Required'. The default value is 'Preferred'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Preferred", "Required"),
											},
										},

										"tenancy": schema.StringAttribute{
											Description:         "Determines the level of resource isolation between Pods. It can have the following values: 'SharedNode' and 'DedicatedNode'.  - SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s. - DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node. In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node. Which provides a higher level of isolation and resource guarantee for Pods.  The default value is 'SharedNode'.",
											MarkdownDescription: "Determines the level of resource isolation between Pods. It can have the following values: 'SharedNode' and 'DedicatedNode'.  - SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s. - DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node. In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node. Which provides a higher level of isolation and resource guarantee for Pods.  The default value is 'SharedNode'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("SharedNode", "DedicatedNode"),
											},
										},

										"topology_keys": schema.ListAttribute{
											Description:         "Represents the key of node labels used to define the topology domain for Pod anti-affinity and Pod spread constraints.  In K8s, a topology domain is a set of nodes that have the same value for a specific label key. Nodes with labels containing any of the specified TopologyKeys and identical values are considered to be in the same topology domain.  Note: The concept of topology in the context of K8s TopologyKeys is different from the concept of topology in the ClusterDefinition.  When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule the Pod on nodes with different values for the specified TopologyKeys. This ensures that Pods are spread across different topology domains, promoting high availability and reducing the impact of node failures.  Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey. These keys represent the hostname and zone of a node, respectively. By including these keys in the TopologyKeys list, Pods will be spread across nodes with different hostnames or zones.  In addition to the well-known keys, users can also specify custom label keys as TopologyKeys. This allows for more flexible and custom topology definitions based on the specific needs of the application or environment.  The TopologyKeys field is a slice of strings, where each string represents a label key. The order of the keys in the slice does not matter.",
											MarkdownDescription: "Represents the key of node labels used to define the topology domain for Pod anti-affinity and Pod spread constraints.  In K8s, a topology domain is a set of nodes that have the same value for a specific label key. Nodes with labels containing any of the specified TopologyKeys and identical values are considered to be in the same topology domain.  Note: The concept of topology in the context of K8s TopologyKeys is different from the concept of topology in the ClusterDefinition.  When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule the Pod on nodes with different values for the specified TopologyKeys. This ensures that Pods are spread across different topology domains, promoting high availability and reducing the impact of node failures.  Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey. These keys represent the hostname and zone of a node, respectively. By including these keys in the TopologyKeys list, Pods will be spread across nodes with different hostnames or zones.  In addition to the well-known keys, users can also specify custom label keys as TopologyKeys. This allows for more flexible and custom topology definitions based on the specific needs of the application or environment.  The TopologyKeys field is a slice of strings, where each string represents a label key. The order of the keys in the slice does not matter.",
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

								"component_def": schema.StringAttribute{
									Description:         "References the name of a ComponentDefinition object. The ComponentDefinition specifies the behavior and characteristics of the Component. If both 'componentDefRef' and 'componentDef' are provided, the 'componentDef' will take precedence over 'componentDefRef'.",
									MarkdownDescription: "References the name of a ComponentDefinition object. The ComponentDefinition specifies the behavior and characteristics of the Component. If both 'componentDefRef' and 'componentDef' are provided, the 'componentDef' will take precedence over 'componentDefRef'.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(64),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"component_def_ref": schema.StringAttribute{
									Description:         "References a ClusterComponentDefinition defined in the 'clusterDefinition.spec.componentDef' field. Must comply with the IANA service naming rule.  Deprecated since v0.9, because defining Components in 'clusterDefinition.spec.componentDef' field has been deprecated. This field is replaced by the 'componentDef' field, use 'componentDef' instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDefRef is immutable'",
									MarkdownDescription: "References a ClusterComponentDefinition defined in the 'clusterDefinition.spec.componentDef' field. Must comply with the IANA service naming rule.  Deprecated since v0.9, because defining Components in 'clusterDefinition.spec.componentDef' field has been deprecated. This field is replaced by the 'componentDef' field, use 'componentDef' instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDefRef is immutable'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(22),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
									},
								},

								"configs": schema.ListNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"config_map": schema.SingleNestedAttribute{
												Description:         "ConfigMap source for the config.",
												MarkdownDescription: "ConfigMap source for the config.",
												Attributes: map[string]schema.Attribute{
													"default_mode": schema.Int64Attribute{
														Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"items": schema.ListNestedAttribute{
														Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
														MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																	Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																	MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

											"name": schema.StringAttribute{
												Description:         "The name of the config.",
												MarkdownDescription: "The name of the config.",
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

								"disable_exporter": schema.BoolAttribute{
									Description:         "Determines whether metrics exporter information is annotated on the Component's headless Service.  If set to true, the following annotations will not be patched into the Service:  - 'monitor.kubeblocks.io/path' - 'monitor.kubeblocks.io/port' - 'monitor.kubeblocks.io/scheme'  These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
									MarkdownDescription: "Determines whether metrics exporter information is annotated on the Component's headless Service.  If set to true, the following annotations will not be patched into the Service:  - 'monitor.kubeblocks.io/path' - 'monitor.kubeblocks.io/port' - 'monitor.kubeblocks.io/scheme'  These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"enabled_logs": schema.ListAttribute{
									Description:         "Specifies which types of logs should be collected for the Component. The log types are defined in the 'componentDefinition.spec.logConfigs' field with the LogConfig entries.  The elements in the 'enabledLogs' array correspond to the names of the LogConfig entries. For example, if the 'componentDefinition.spec.logConfigs' defines LogConfig entries with names 'slow_query_log' and 'error_log', you can enable the collection of these logs by including their names in the 'enabledLogs' array: '''yaml enabledLogs: - slow_query_log - error_log '''",
									MarkdownDescription: "Specifies which types of logs should be collected for the Component. The log types are defined in the 'componentDefinition.spec.logConfigs' field with the LogConfig entries.  The elements in the 'enabledLogs' array correspond to the names of the LogConfig entries. For example, if the 'componentDefinition.spec.logConfigs' defines LogConfig entries with names 'slow_query_log' and 'error_log', you can enable the collection of these logs by including their names in the 'enabledLogs' array: '''yaml enabledLogs: - slow_query_log - error_log '''",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"instances": schema.ListNestedAttribute{
									Description:         "Allows for the customization of configuration values for each instance within a Component. An instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps). While instances typically share a common configuration as defined in the ClusterComponentSpec, they can require unique settings in various scenarios:  For example: - A database Component might require different resource allocations for primary and secondary instances, with primaries needing more resources. - During a rolling upgrade, a Component may first update the image for one or a few instances, and then update the remaining instances after verifying that the updated instances are functioning correctly.  InstanceTemplate allows for specifying these unique configurations per instance. Each instance's name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal), starting with an ordinal of 0. It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.  The sum of replicas across all InstanceTemplates should not exceed the total number of replicas specified for the Component. Any remaining replicas will be generated using the default template and will follow the default naming rules.",
									MarkdownDescription: "Allows for the customization of configuration values for each instance within a Component. An instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps). While instances typically share a common configuration as defined in the ClusterComponentSpec, they can require unique settings in various scenarios:  For example: - A database Component might require different resource allocations for primary and secondary instances, with primaries needing more resources. - During a rolling upgrade, a Component may first update the image for one or a few instances, and then update the remaining instances after verifying that the updated instances are functioning correctly.  InstanceTemplate allows for specifying these unique configurations per instance. Each instance's name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal), starting with an ordinal of 0. It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.  The sum of replicas across all InstanceTemplates should not exceed the total number of replicas specified for the Component. Any remaining replicas will be generated using the default template and will follow the default naming rules.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Specifies a map of key-value pairs to be merged into the Pod's existing annotations. Existing keys will have their values overwritten, while new keys will be added to the annotations.",
												MarkdownDescription: "Specifies a map of key-value pairs to be merged into the Pod's existing annotations. Existing keys will have their values overwritten, while new keys will be added to the annotations.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"env": schema.ListNestedAttribute{
												Description:         "Defines Env to override. Add new or override existing envs.",
												MarkdownDescription: "Defines Env to override. Add new or override existing envs.",
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
															Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
															MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
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
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																	Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																	MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
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
																	Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																	MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
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
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

											"image": schema.StringAttribute{
												Description:         "Specifies an override for the first container's image in the Pod.",
												MarkdownDescription: "Specifies an override for the first container's image in the Pod.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"labels": schema.MapAttribute{
												Description:         "Specifies a map of key-value pairs that will be merged into the Pod's existing labels. Values for existing keys will be overwritten, and new keys will be added.",
												MarkdownDescription: "Specifies a map of key-value pairs that will be merged into the Pod's existing labels. Values for existing keys will be overwritten, and new keys will be added.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name specifies the unique name of the instance Pod created using this InstanceTemplate. This name is constructed by concatenating the Component's name, the template's name, and the instance's ordinal using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0. The specified name overrides any default naming conventions or patterns.",
												MarkdownDescription: "Name specifies the unique name of the instance Pod created using this InstanceTemplate. This name is constructed by concatenating the Component's name, the template's name, and the instance's ordinal using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0. The specified name overrides any default naming conventions or patterns.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(54),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
												},
											},

											"node_name": schema.StringAttribute{
												Description:         "Specifies the name of the node where the Pod should be scheduled. If set, the Pod will be directly assigned to the specified node, bypassing the Kubernetes scheduler. This is useful for controlling Pod placement on specific nodes.  Important considerations: - 'nodeName' bypasses default scheduling constraints (e.g., resource requirements, node selectors, affinity rules). - It is the user's responsibility to ensure the node is suitable for the Pod. - If the node is unavailable, the Pod will remain in 'Pending' state until the node is available or the Pod is deleted.",
												MarkdownDescription: "Specifies the name of the node where the Pod should be scheduled. If set, the Pod will be directly assigned to the specified node, bypassing the Kubernetes scheduler. This is useful for controlling Pod placement on specific nodes.  Important considerations: - 'nodeName' bypasses default scheduling constraints (e.g., resource requirements, node selectors, affinity rules). - It is the user's responsibility to ensure the node is suitable for the Pod. - If the node is unavailable, the Pod will remain in 'Pending' state until the node is available or the Pod is deleted.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_selector": schema.MapAttribute{
												Description:         "Defines NodeSelector to override.",
												MarkdownDescription: "Defines NodeSelector to override.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"replicas": schema.Int64Attribute{
												Description:         "Specifies the number of instances (Pods) to create from this InstanceTemplate. This field allows setting how many replicated instances of the Component, with the specific overrides in the InstanceTemplate, are created. The default value is 1. A value of 0 disables instance creation.",
												MarkdownDescription: "Specifies the number of instances (Pods) to create from this InstanceTemplate. This field allows setting how many replicated instances of the Component, with the specific overrides in the InstanceTemplate, are created. The default value is 1. A value of 0 disables instance creation.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(0),
												},
											},

											"resources": schema.SingleNestedAttribute{
												Description:         "Specifies an override for the resource requirements of the first container in the Pod. This field allows for customizing resource allocation (CPU, memory, etc.) for the container.",
												MarkdownDescription: "Specifies an override for the resource requirements of the first container in the Pod. This field allows for customizing resource allocation (CPU, memory, etc.) for the container.",
												Attributes: map[string]schema.Attribute{
													"claims": schema.ListNestedAttribute{
														Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																	MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
														Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"requests": schema.MapAttribute{
														Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
														MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

											"tolerations": schema.ListNestedAttribute{
												Description:         "Tolerations specifies a list of tolerations to be applied to the Pod, allowing it to tolerate node taints. This field can be used to add new tolerations or override existing ones.",
												MarkdownDescription: "Tolerations specifies a list of tolerations to be applied to the Pod, allowing it to tolerate node taints. This field can be used to add new tolerations or override existing ones.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"effect": schema.StringAttribute{
															Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
															MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
															MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"operator": schema.StringAttribute{
															Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
															MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"toleration_seconds": schema.Int64Attribute{
															Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
															MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"value": schema.StringAttribute{
															Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
															MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

											"volume_claim_templates": schema.ListNestedAttribute{
												Description:         "Defines VolumeClaimTemplates to override. Add new or override existing volume claim templates.",
												MarkdownDescription: "Defines VolumeClaimTemplates to override. Add new or override existing volume claim templates.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Refers to the name of a volumeMount defined in either:  - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)  The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
															MarkdownDescription: "Refers to the name of a volumeMount defined in either:  - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)  The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"spec": schema.SingleNestedAttribute{
															Description:         "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.  When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
															MarkdownDescription: "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.  When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
															Attributes: map[string]schema.Attribute{
																"access_modes": schema.MapAttribute{
																	Description:         "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																	MarkdownDescription: "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"resources": schema.SingleNestedAttribute{
																	Description:         "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
																	MarkdownDescription: "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
																	Attributes: map[string]schema.Attribute{
																		"claims": schema.ListNestedAttribute{
																			Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"name": schema.StringAttribute{
																						Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																						MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																			Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"requests": schema.MapAttribute{
																			Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																			MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
																	Description:         "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																	MarkdownDescription: "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_mode": schema.StringAttribute{
																	Description:         "Defines what type of volume is required by the claim, either Block or Filesystem.",
																	MarkdownDescription: "Defines what type of volume is required by the claim, either Block or Filesystem.",
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

											"volume_mounts": schema.ListNestedAttribute{
												Description:         "Defines VolumeMounts to override. Add new or override existing volume mounts of the first container in the Pod.",
												MarkdownDescription: "Defines VolumeMounts to override. Add new or override existing volume mounts of the first container in the Pod.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"mount_path": schema.StringAttribute{
															Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
															MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"mount_propagation": schema.StringAttribute{
															Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
															MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "This must match the Name of a Volume.",
															MarkdownDescription: "This must match the Name of a Volume.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
															MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sub_path": schema.StringAttribute{
															Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
															MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"sub_path_expr": schema.StringAttribute{
															Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
															MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

											"volumes": schema.ListNestedAttribute{
												Description:         "Defines Volumes to override. Add new or override existing volumes.",
												MarkdownDescription: "Defines Volumes to override. Add new or override existing volumes.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"aws_elastic_block_store": schema.SingleNestedAttribute{
															Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
															MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
																	MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"partition": schema.Int64Attribute{
																	Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																	MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_id": schema.StringAttribute{
																	Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"azure_disk": schema.SingleNestedAttribute{
															Description:         "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
															MarkdownDescription: "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
															Attributes: map[string]schema.Attribute{
																"caching_mode": schema.StringAttribute{
																	Description:         "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																	MarkdownDescription: "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"disk_name": schema.StringAttribute{
																	Description:         "diskName is the Name of the data disk in the blob storage",
																	MarkdownDescription: "diskName is the Name of the data disk in the blob storage",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"disk_uri": schema.StringAttribute{
																	Description:         "diskURI is the URI of data disk in the blob storage",
																	MarkdownDescription: "diskURI is the URI of data disk in the blob storage",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"fs_type": schema.StringAttribute{
																	Description:         "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"kind": schema.StringAttribute{
																	Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																	MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"azure_file": schema.SingleNestedAttribute{
															Description:         "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
															MarkdownDescription: "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
															Attributes: map[string]schema.Attribute{
																"read_only": schema.BoolAttribute{
																	Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_name": schema.StringAttribute{
																	Description:         "secretName is the  name of secret that contains Azure Storage Account Name and Key",
																	MarkdownDescription: "secretName is the  name of secret that contains Azure Storage Account Name and Key",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"share_name": schema.StringAttribute{
																	Description:         "shareName is the azure share Name",
																	MarkdownDescription: "shareName is the azure share Name",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"cephfs": schema.SingleNestedAttribute{
															Description:         "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
															MarkdownDescription: "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
															Attributes: map[string]schema.Attribute{
																"monitors": schema.ListAttribute{
																	Description:         "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																	MarkdownDescription: "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_file": schema.StringAttribute{
																	Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"user": schema.StringAttribute{
																	Description:         "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	MarkdownDescription: "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"cinder": schema.SingleNestedAttribute{
															Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																	MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"volume_id": schema.StringAttribute{
																	Description:         "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	MarkdownDescription: "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"config_map": schema.SingleNestedAttribute{
															Description:         "configMap represents a configMap that should populate this volume",
															MarkdownDescription: "configMap represents a configMap that should populate this volume",
															Attributes: map[string]schema.Attribute{
																"default_mode": schema.Int64Attribute{
																	Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"items": schema.ListNestedAttribute{
																	Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																	MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																				Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.StringAttribute{
																				Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

														"csi": schema.SingleNestedAttribute{
															Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
															MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
															Attributes: map[string]schema.Attribute{
																"driver": schema.StringAttribute{
																	Description:         "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																	MarkdownDescription: "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"fs_type": schema.StringAttribute{
																	Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																	MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"node_publish_secret_ref": schema.SingleNestedAttribute{
																	Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																	MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																	MarkdownDescription: "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_attributes": schema.MapAttribute{
																	Description:         "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
																	MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
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

														"downward_api": schema.SingleNestedAttribute{
															Description:         "downwardAPI represents downward API about the pod that should populate this volume",
															MarkdownDescription: "downwardAPI represents downward API about the pod that should populate this volume",
															Attributes: map[string]schema.Attribute{
																"default_mode": schema.Int64Attribute{
																	Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"items": schema.ListNestedAttribute{
																	Description:         "Items is a list of downward API volume file",
																	MarkdownDescription: "Items is a list of downward API volume file",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"field_ref": schema.SingleNestedAttribute{
																				Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																				MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
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
																				Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
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
																				Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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

														"empty_dir": schema.SingleNestedAttribute{
															Description:         "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															Attributes: map[string]schema.Attribute{
																"medium": schema.StringAttribute{
																	Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																	MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"size_limit": schema.StringAttribute{
																	Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																	MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"ephemeral": schema.SingleNestedAttribute{
															Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
															MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
															Attributes: map[string]schema.Attribute{
																"volume_claim_template": schema.SingleNestedAttribute{
																	Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
																	MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
																	Attributes: map[string]schema.Attribute{
																		"metadata": schema.SingleNestedAttribute{
																			Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																			MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																			Attributes: map[string]schema.Attribute{
																				"annotations": schema.MapAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"finalizers": schema.ListAttribute{
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

																				"name": schema.StringAttribute{
																					Description:         "",
																					MarkdownDescription: "",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"namespace": schema.StringAttribute{
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

																		"spec": schema.SingleNestedAttribute{
																			Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																			MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																			Attributes: map[string]schema.Attribute{
																				"access_modes": schema.ListAttribute{
																					Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																					MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"data_source": schema.SingleNestedAttribute{
																					Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																					MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																					Attributes: map[string]schema.Attribute{
																						"api_group": schema.StringAttribute{
																							Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																							MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
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
																					Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																					MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																					Attributes: map[string]schema.Attribute{
																						"api_group": schema.StringAttribute{
																							Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																							MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
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
																							Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																							MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
																					Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																					MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																					Attributes: map[string]schema.Attribute{
																						"claims": schema.ListNestedAttribute{
																							Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																							MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"name": schema.StringAttribute{
																										Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																										MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																							Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																							MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"requests": schema.MapAttribute{
																							Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																							MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
																										Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																										MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"values": schema.ListAttribute{
																										Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																										MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																							Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																							MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																					Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																					MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"volume_mode": schema.StringAttribute{
																					Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																					MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
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

														"fc": schema.SingleNestedAttribute{
															Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
															MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"lun": schema.Int64Attribute{
																	Description:         "lun is Optional: FC target lun number",
																	MarkdownDescription: "lun is Optional: FC target lun number",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"target_ww_ns": schema.ListAttribute{
																	Description:         "targetWWNs is Optional: FC target worldwide names (WWNs)",
																	MarkdownDescription: "targetWWNs is Optional: FC target worldwide names (WWNs)",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"wwids": schema.ListAttribute{
																	Description:         "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
																	MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
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

														"flex_volume": schema.SingleNestedAttribute{
															Description:         "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
															MarkdownDescription: "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
															Attributes: map[string]schema.Attribute{
																"driver": schema.StringAttribute{
																	Description:         "driver is the name of the driver to use for this volume.",
																	MarkdownDescription: "driver is the name of the driver to use for this volume.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"options": schema.MapAttribute{
																	Description:         "options is Optional: this field holds extra command options if any.",
																	MarkdownDescription: "options is Optional: this field holds extra command options if any.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																	MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

														"flocker": schema.SingleNestedAttribute{
															Description:         "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
															MarkdownDescription: "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
															Attributes: map[string]schema.Attribute{
																"dataset_name": schema.StringAttribute{
																	Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																	MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"dataset_uuid": schema.StringAttribute{
																	Description:         "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																	MarkdownDescription: "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"gce_persistent_disk": schema.SingleNestedAttribute{
															Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																	MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"partition": schema.Int64Attribute{
																	Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"pd_name": schema.StringAttribute{
																	Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"git_repo": schema.SingleNestedAttribute{
															Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
															MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
															Attributes: map[string]schema.Attribute{
																"directory": schema.StringAttribute{
																	Description:         "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																	MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"repository": schema.StringAttribute{
																	Description:         "repository is the URL",
																	MarkdownDescription: "repository is the URL",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"revision": schema.StringAttribute{
																	Description:         "revision is the commit hash for the specified revision.",
																	MarkdownDescription: "revision is the commit hash for the specified revision.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"glusterfs": schema.SingleNestedAttribute{
															Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
															MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
															Attributes: map[string]schema.Attribute{
																"endpoints": schema.StringAttribute{
																	Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"path": schema.StringAttribute{
																	Description:         "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	MarkdownDescription: "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"host_path": schema.SingleNestedAttribute{
															Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
															MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
															Attributes: map[string]schema.Attribute{
																"path": schema.StringAttribute{
																	Description:         "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																	MarkdownDescription: "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"type": schema.StringAttribute{
																	Description:         "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																	MarkdownDescription: "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"iscsi": schema.SingleNestedAttribute{
															Description:         "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
															MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
															Attributes: map[string]schema.Attribute{
																"chap_auth_discovery": schema.BoolAttribute{
																	Description:         "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																	MarkdownDescription: "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"chap_auth_session": schema.BoolAttribute{
																	Description:         "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																	MarkdownDescription: "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																	MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"initiator_name": schema.StringAttribute{
																	Description:         "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																	MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"iqn": schema.StringAttribute{
																	Description:         "iqn is the target iSCSI Qualified Name.",
																	MarkdownDescription: "iqn is the target iSCSI Qualified Name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"iscsi_interface": schema.StringAttribute{
																	Description:         "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																	MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"lun": schema.Int64Attribute{
																	Description:         "lun represents iSCSI Target Lun number.",
																	MarkdownDescription: "lun represents iSCSI Target Lun number.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"portals": schema.ListAttribute{
																	Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																	MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																	MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																	MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"target_portal": schema.StringAttribute{
																	Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																	MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
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
															Description:         "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															MarkdownDescription: "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"nfs": schema.SingleNestedAttribute{
															Description:         "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															Attributes: map[string]schema.Attribute{
																"path": schema.StringAttribute{
																	Description:         "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	MarkdownDescription: "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"server": schema.StringAttribute{
																	Description:         "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	MarkdownDescription: "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"persistent_volume_claim": schema.SingleNestedAttribute{
															Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
															MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
															Attributes: map[string]schema.Attribute{
																"claim_name": schema.StringAttribute{
																	Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																	MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
																	MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"photon_persistent_disk": schema.SingleNestedAttribute{
															Description:         "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
															MarkdownDescription: "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"pd_id": schema.StringAttribute{
																	Description:         "pdID is the ID that identifies Photon Controller persistent disk",
																	MarkdownDescription: "pdID is the ID that identifies Photon Controller persistent disk",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"portworx_volume": schema.SingleNestedAttribute{
															Description:         "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
															MarkdownDescription: "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_id": schema.StringAttribute{
																	Description:         "volumeID uniquely identifies a Portworx volume",
																	MarkdownDescription: "volumeID uniquely identifies a Portworx volume",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"projected": schema.SingleNestedAttribute{
															Description:         "projected items for all in one resources secrets, configmaps, and downward API",
															MarkdownDescription: "projected items for all in one resources secrets, configmaps, and downward API",
															Attributes: map[string]schema.Attribute{
																"default_mode": schema.Int64Attribute{
																	Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"sources": schema.ListNestedAttribute{
																	Description:         "sources is the list of volume projections",
																	MarkdownDescription: "sources is the list of volume projections",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"config_map": schema.SingleNestedAttribute{
																				Description:         "configMap information about the configMap data to project",
																				MarkdownDescription: "configMap information about the configMap data to project",
																				Attributes: map[string]schema.Attribute{
																					"items": schema.ListNestedAttribute{
																						Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																						MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																									Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"path": schema.StringAttribute{
																									Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																									MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																									Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																									MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
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
																									Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
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
																									Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																									MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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
																						Description:         "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																						MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																									Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"path": schema.StringAttribute{
																									Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																									MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																						Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																						MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																						Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																						MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"expiration_seconds": schema.Int64Attribute{
																						Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																						MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "path is the path relative to the mount point of the file to project the token into.",
																						MarkdownDescription: "path is the path relative to the mount point of the file to project the token into.",
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

														"quobyte": schema.SingleNestedAttribute{
															Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
															MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
															Attributes: map[string]schema.Attribute{
																"group": schema.StringAttribute{
																	Description:         "group to map volume access to Default is no group",
																	MarkdownDescription: "group to map volume access to Default is no group",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																	MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"registry": schema.StringAttribute{
																	Description:         "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																	MarkdownDescription: "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"tenant": schema.StringAttribute{
																	Description:         "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																	MarkdownDescription: "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"user": schema.StringAttribute{
																	Description:         "user to map volume access to Defaults to serivceaccount user",
																	MarkdownDescription: "user to map volume access to Defaults to serivceaccount user",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume": schema.StringAttribute{
																	Description:         "volume is a string that references an already created Quobyte volume by name.",
																	MarkdownDescription: "volume is a string that references an already created Quobyte volume by name.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"rbd": schema.SingleNestedAttribute{
															Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
															MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																	MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"image": schema.StringAttribute{
																	Description:         "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"keyring": schema.StringAttribute{
																	Description:         "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"monitors": schema.ListAttribute{
																	Description:         "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	ElementType:         types.StringType,
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"pool": schema.StringAttribute{
																	Description:         "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"user": schema.StringAttribute{
																	Description:         "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	MarkdownDescription: "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"scale_io": schema.SingleNestedAttribute{
															Description:         "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
															MarkdownDescription: "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"gateway": schema.StringAttribute{
																	Description:         "gateway is the host address of the ScaleIO API Gateway.",
																	MarkdownDescription: "gateway is the host address of the ScaleIO API Gateway.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"protection_domain": schema.StringAttribute{
																	Description:         "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																	MarkdownDescription: "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																	MarkdownDescription: "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: true,
																	Optional: false,
																	Computed: false,
																},

																"ssl_enabled": schema.BoolAttribute{
																	Description:         "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																	MarkdownDescription: "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"storage_mode": schema.StringAttribute{
																	Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																	MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"storage_pool": schema.StringAttribute{
																	Description:         "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																	MarkdownDescription: "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"system": schema.StringAttribute{
																	Description:         "system is the name of the storage system as configured in ScaleIO.",
																	MarkdownDescription: "system is the name of the storage system as configured in ScaleIO.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"volume_name": schema.StringAttribute{
																	Description:         "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
																	MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
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
															Description:         "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
															MarkdownDescription: "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
															Attributes: map[string]schema.Attribute{
																"default_mode": schema.Int64Attribute{
																	Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"items": schema.ListNestedAttribute{
																	Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																	MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																				Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"path": schema.StringAttribute{
																				Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																				MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																"optional": schema.BoolAttribute{
																	Description:         "optional field specify whether the Secret or its keys must be defined",
																	MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_name": schema.StringAttribute{
																	Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																	MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"storageos": schema.SingleNestedAttribute{
															Description:         "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
															MarkdownDescription: "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"secret_ref": schema.SingleNestedAttribute{
																	Description:         "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
																	MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"volume_name": schema.StringAttribute{
																	Description:         "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
																	MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_namespace": schema.StringAttribute{
																	Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																	MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},
															},
															Required: false,
															Optional: true,
															Computed: false,
														},

														"vsphere_volume": schema.SingleNestedAttribute{
															Description:         "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
															MarkdownDescription: "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
															Attributes: map[string]schema.Attribute{
																"fs_type": schema.StringAttribute{
																	Description:         "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	MarkdownDescription: "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"storage_policy_id": schema.StringAttribute{
																	Description:         "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																	MarkdownDescription: "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"storage_policy_name": schema.StringAttribute{
																	Description:         "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																	MarkdownDescription: "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"volume_path": schema.StringAttribute{
																	Description:         "volumePath is the path that identifies vSphere volume vmdk",
																	MarkdownDescription: "volumePath is the path that identifies vSphere volume vmdk",
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
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"issuer": schema.SingleNestedAttribute{
									Description:         "Specifies the configuration for the TLS certificates issuer. It allows defining the issuer name and the reference to the secret containing the TLS certificates and key. The secret should contain the CA certificate, TLS certificate, and private key in the specified keys. Required when TLS is enabled.",
									MarkdownDescription: "Specifies the configuration for the TLS certificates issuer. It allows defining the issuer name and the reference to the secret containing the TLS certificates and key. The secret should contain the CA certificate, TLS certificate, and private key in the specified keys. Required when TLS is enabled.",
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Description:         "The issuer for TLS certificates. It only allows two enum values: 'KubeBlocks' and 'UserProvided'.  - 'KubeBlocks' indicates that the self-signed TLS certificates generated by the KubeBlocks Operator will be used. - 'UserProvided' means that the user is responsible for providing their own CA, Cert, and Key. In this case, the user-provided CA certificate, server certificate, and private key will be used for TLS communication.",
											MarkdownDescription: "The issuer for TLS certificates. It only allows two enum values: 'KubeBlocks' and 'UserProvided'.  - 'KubeBlocks' indicates that the self-signed TLS certificates generated by the KubeBlocks Operator will be used. - 'UserProvided' means that the user is responsible for providing their own CA, Cert, and Key. In this case, the user-provided CA certificate, server certificate, and private key will be used for TLS communication.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "SecretRef is the reference to the secret that contains user-provided certificates. It is required when the issuer is set to 'UserProvided'.",
											MarkdownDescription: "SecretRef is the reference to the secret that contains user-provided certificates. It is required when the issuer is set to 'UserProvided'.",
											Attributes: map[string]schema.Attribute{
												"ca": schema.StringAttribute{
													Description:         "Key of CA cert in Secret",
													MarkdownDescription: "Key of CA cert in Secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"cert": schema.StringAttribute{
													Description:         "Key of Cert in Secret",
													MarkdownDescription: "Key of Cert in Secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Key of TLS private key in Secret",
													MarkdownDescription: "Key of TLS private key in Secret",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"name": schema.StringAttribute{
													Description:         "Name of the Secret that contains user-provided certificates.",
													MarkdownDescription: "Name of the Secret that contains user-provided certificates.",
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

								"monitor": schema.BoolAttribute{
									Description:         "Deprecated since v0.9 Determines whether metrics exporter information is annotated on the Component's headless Service.  If set to true, the following annotations will be patched into the Service:  - 'monitor.kubeblocks.io/path' - 'monitor.kubeblocks.io/port' - 'monitor.kubeblocks.io/scheme'  These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
									MarkdownDescription: "Deprecated since v0.9 Determines whether metrics exporter information is annotated on the Component's headless Service.  If set to true, the following annotations will be patched into the Service:  - 'monitor.kubeblocks.io/path' - 'monitor.kubeblocks.io/port' - 'monitor.kubeblocks.io/scheme'  These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"name": schema.StringAttribute{
									Description:         "Specifies the Component's name. It's part of the Service DNS name and must comply with the IANA service naming rule. The name is optional when ClusterComponentSpec is used as a template (e.g., in 'shardingSpec'), but required otherwise.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='name is immutable'",
									MarkdownDescription: "Specifies the Component's name. It's part of the Service DNS name and must comply with the IANA service naming rule. The name is optional when ClusterComponentSpec is used as a template (e.g., in 'shardingSpec'), but required otherwise.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='name is immutable'",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(22),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
									},
								},

								"offline_instances": schema.ListAttribute{
									Description:         "Specifies the names of instances to be transitioned to offline status.  Marking an instance as offline results in the following:  1. The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential future reuse or data recovery, but it is no longer actively used. 2. The ordinal number assigned to this instance is preserved, ensuring it remains unique and avoiding conflicts with new instances.  Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining ordinal consistency within the Cluster. Note that offline instances and their associated resources, such as PVCs, are not automatically deleted. The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.",
									MarkdownDescription: "Specifies the names of instances to be transitioned to offline status.  Marking an instance as offline results in the following:  1. The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential future reuse or data recovery, but it is no longer actively used. 2. The ordinal number assigned to this instance is preserved, ensuring it remains unique and avoiding conflicts with new instances.  Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining ordinal consistency within the Cluster. Note that offline instances and their associated resources, such as PVCs, are not automatically deleted. The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replicas": schema.Int64Attribute{
									Description:         "Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.",
									MarkdownDescription: "Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},

								"resources": schema.SingleNestedAttribute{
									Description:         "Specifies the resources required by the Component. It allows defining the CPU, memory requirements and limits for the Component's containers.",
									MarkdownDescription: "Specifies the resources required by the Component. It allows defining the CPU, memory requirements and limits for the Component's containers.",
									Attributes: map[string]schema.Attribute{
										"claims": schema.ListNestedAttribute{
											Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
											MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
														MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
											Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"requests": schema.MapAttribute{
											Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
											MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

								"scheduling_policy": schema.SingleNestedAttribute{
									Description:         "Specifies the scheduling policy for the Component.",
									MarkdownDescription: "Specifies the scheduling policy for the Component.",
									Attributes: map[string]schema.Attribute{
										"affinity": schema.SingleNestedAttribute{
											Description:         "Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.",
											MarkdownDescription: "Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.",
											Attributes: map[string]schema.Attribute{
												"node_affinity": schema.SingleNestedAttribute{
													Description:         "Describes node affinity scheduling rules for the pod.",
													MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
													Attributes: map[string]schema.Attribute{
														"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
															Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
															MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
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
																							Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																							Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
															Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
															MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
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
																							Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																							Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
													Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
													MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
													Attributes: map[string]schema.Attribute{
														"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
															Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
															MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"pod_affinity_term": schema.SingleNestedAttribute{
																		Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																		MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																									Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																				Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																									Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																				Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
																		Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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
															Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
															MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"label_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																							Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																							MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																				Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																		Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																							Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																							MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																				Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																		Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"topology_key": schema.StringAttribute{
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
													Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
													MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
													Attributes: map[string]schema.Attribute{
														"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
															Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
															MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"pod_affinity_term": schema.SingleNestedAttribute{
																		Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																		MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																									Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																				Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																									Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																				Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
																		Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																		MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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
															Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
															MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"label_selector": schema.SingleNestedAttribute{
																		Description:         "A label query over a set of resources, in this case pods.",
																		MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																							Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																							MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																				Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																		Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																		MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																							Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																							MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"values": schema.ListAttribute{
																							Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																							MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																				Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																				MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																		Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"topology_key": schema.StringAttribute{
																		Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																		MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

										"node_name": schema.StringAttribute{
											Description:         "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty, the scheduler simply schedules this Pod onto that node, assuming that it fits resource requirements.",
											MarkdownDescription: "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty, the scheduler simply schedules this Pod onto that node, assuming that it fits resource requirements.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_selector": schema.MapAttribute{
											Description:         "NodeSelector is a selector which must be true for the Pod to fit on a node. Selector which must match a node's labels for the Pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
											MarkdownDescription: "NodeSelector is a selector which must be true for the Pod to fit on a node. Selector which must match a node's labels for the Pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"scheduler_name": schema.StringAttribute{
											Description:         "If specified, the Pod will be dispatched by specified scheduler. If not specified, the Pod will be dispatched by default scheduler.",
											MarkdownDescription: "If specified, the Pod will be dispatched by specified scheduler. If not specified, the Pod will be dispatched by default scheduler.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tolerations": schema.ListNestedAttribute{
											Description:         "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
											MarkdownDescription: "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"effect": schema.StringAttribute{
														Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
														MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
														MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"operator": schema.StringAttribute{
														Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
														MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"toleration_seconds": schema.Int64Attribute{
														Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
														MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
														MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
											Description:         "TopologySpreadConstraints describes how a group of Pods ought to spread across topology domains. Scheduler will schedule Pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
											MarkdownDescription: "TopologySpreadConstraints describes how a group of Pods ought to spread across topology domains. Scheduler will schedule Pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"label_selector": schema.SingleNestedAttribute{
														Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
														MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
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
																			Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																			MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"values": schema.ListAttribute{
																			Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																			MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
														Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
														MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_skew": schema.Int64Attribute{
														Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
														MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"min_domains": schema.Int64Attribute{
														Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
														MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_affinity_policy": schema.StringAttribute{
														Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
														MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_taints_policy": schema.StringAttribute{
														Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
														MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"topology_key": schema.StringAttribute{
														Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
														MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"when_unsatisfiable": schema.StringAttribute{
														Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
														MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
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

								"service_account_name": schema.StringAttribute{
									Description:         "Specifies the name of the ServiceAccount required by the running Component. This ServiceAccount is used to grant necessary permissions for the Component's Pods to interact with other Kubernetes resources, such as modifying Pod labels or sending events.  Defaults: If not specified, KubeBlocks automatically assigns a default ServiceAccount named 'kb-{cluster.name}', bound to a default role installed together with KubeBlocks.  Future Changes: Future versions might change the default ServiceAccount creation strategy to one per Component, potentially revising the naming to 'kb-{cluster.name}-{component.name}'.  Users can override the automatic ServiceAccount assignment by explicitly setting the name of an existed ServiceAccount in this field.",
									MarkdownDescription: "Specifies the name of the ServiceAccount required by the running Component. This ServiceAccount is used to grant necessary permissions for the Component's Pods to interact with other Kubernetes resources, such as modifying Pod labels or sending events.  Defaults: If not specified, KubeBlocks automatically assigns a default ServiceAccount named 'kb-{cluster.name}', bound to a default role installed together with KubeBlocks.  Future Changes: Future versions might change the default ServiceAccount creation strategy to one per Component, potentially revising the naming to 'kb-{cluster.name}-{component.name}'.  Users can override the automatic ServiceAccount assignment by explicitly setting the name of an existed ServiceAccount in this field.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"service_refs": schema.ListNestedAttribute{
									Description:         "Defines a list of ServiceRef for a Component, enabling access to both external services and Services provided by other Clusters.  Types of services:  - External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator; Require a ServiceDescriptor for connection details. - Services provided by a Cluster: Managed by the same KubeBlocks operator; identified using Cluster, Component and Service names.  ServiceRefs with identical 'serviceRef.name' in the same Cluster are considered the same.  Example: '''yaml serviceRefs: - name: 'redis-sentinel' serviceDescriptor: name: 'external-redis-sentinel' - name: 'postgres-cluster' clusterServiceSelector: cluster: 'my-postgres-cluster' service: component: 'postgresql' ''' The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.",
									MarkdownDescription: "Defines a list of ServiceRef for a Component, enabling access to both external services and Services provided by other Clusters.  Types of services:  - External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator; Require a ServiceDescriptor for connection details. - Services provided by a Cluster: Managed by the same KubeBlocks operator; identified using Cluster, Component and Service names.  ServiceRefs with identical 'serviceRef.name' in the same Cluster are considered the same.  Example: '''yaml serviceRefs: - name: 'redis-sentinel' serviceDescriptor: name: 'external-redis-sentinel' - name: 'postgres-cluster' clusterServiceSelector: cluster: 'my-postgres-cluster' service: component: 'postgresql' ''' The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"cluster": schema.StringAttribute{
												Description:         "Specifies the name of the KubeBlocks Cluster being referenced. This is used when services from another KubeBlocks Cluster are consumed.  By default, the referenced KubeBlocks Cluster's 'clusterDefinition.spec.connectionCredential' will be utilized to bind to the current Component. This credential should include: 'endpoint', 'port', 'username', and 'password'.  Note:  - The 'ServiceKind' and 'ServiceVersion' specified in the service reference within the ClusterDefinition are not validated when using this approach. - If both 'cluster' and 'serviceDescriptor' are present, 'cluster' will take precedence.  Deprecated since v0.9 since 'clusterDefinition.spec.connectionCredential' is deprecated, use 'clusterServiceSelector' instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
												MarkdownDescription: "Specifies the name of the KubeBlocks Cluster being referenced. This is used when services from another KubeBlocks Cluster are consumed.  By default, the referenced KubeBlocks Cluster's 'clusterDefinition.spec.connectionCredential' will be utilized to bind to the current Component. This credential should include: 'endpoint', 'port', 'username', and 'password'.  Note:  - The 'ServiceKind' and 'ServiceVersion' specified in the service reference within the ClusterDefinition are not validated when using this approach. - If both 'cluster' and 'serviceDescriptor' are present, 'cluster' will take precedence.  Deprecated since v0.9 since 'clusterDefinition.spec.connectionCredential' is deprecated, use 'clusterServiceSelector' instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cluster_service_selector": schema.SingleNestedAttribute{
												Description:         "References a service provided by another KubeBlocks Cluster. It specifies the ClusterService and the account credentials needed for access.",
												MarkdownDescription: "References a service provided by another KubeBlocks Cluster. It specifies the ClusterService and the account credentials needed for access.",
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "The name of the Cluster being referenced.",
														MarkdownDescription: "The name of the Cluster being referenced.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"credential": schema.SingleNestedAttribute{
														Description:         "Specifies the SystemAccount to authenticate and establish a connection with the referenced Cluster. The SystemAccount should be defined in 'componentDefinition.spec.systemAccounts' of the Component providing the service in the referenced Cluster.",
														MarkdownDescription: "Specifies the SystemAccount to authenticate and establish a connection with the referenced Cluster. The SystemAccount should be defined in 'componentDefinition.spec.systemAccounts' of the Component providing the service in the referenced Cluster.",
														Attributes: map[string]schema.Attribute{
															"component": schema.StringAttribute{
																Description:         "The name of the Component where the credential resides in.",
																MarkdownDescription: "The name of the Component where the credential resides in.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"name": schema.StringAttribute{
																Description:         "The name of the credential (SystemAccount) to reference.",
																MarkdownDescription: "The name of the credential (SystemAccount) to reference.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"service": schema.SingleNestedAttribute{
														Description:         "Identifies a ClusterService from the list of Services defined in 'cluster.spec.services' of the referenced Cluster.",
														MarkdownDescription: "Identifies a ClusterService from the list of Services defined in 'cluster.spec.services' of the referenced Cluster.",
														Attributes: map[string]schema.Attribute{
															"component": schema.StringAttribute{
																Description:         "The name of the Component where the Service resides in.  It is required when referencing a Component's Service.",
																MarkdownDescription: "The name of the Component where the Service resides in.  It is required when referencing a Component's Service.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"port": schema.StringAttribute{
																Description:         "The port name of the Service to be referenced.  If there is a non-zero node-port exist for the matched Service port, the node-port will be selected first.  If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched, and the resolved value will be presented in the following format: service1.name:port1,service2.name:port2...",
																MarkdownDescription: "The port name of the Service to be referenced.  If there is a non-zero node-port exist for the matched Service port, the node-port will be selected first.  If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched, and the resolved value will be presented in the following format: service1.name:port1,service2.name:port2...",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"service": schema.StringAttribute{
																Description:         "The name of the Service to be referenced.  Leave it empty to reference the default Service. Set it to 'headless' to reference the default headless Service.  If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched, and the resolved value will be presented in the following format: service1.name,service2.name...",
																MarkdownDescription: "The name of the Service to be referenced.  Leave it empty to reference the default Service. Set it to 'headless' to reference the default headless Service.  If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched, and the resolved value will be presented in the following format: service1.name,service2.name...",
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

											"name": schema.StringAttribute{
												Description:         "Specifies the identifier of the service reference declaration. It corresponds to the serviceRefDeclaration name defined in either:  - 'componentDefinition.spec.serviceRefDeclarations[*].name' - 'clusterDefinition.spec.componentDefs[*].serviceRefDeclarations[*].name' (deprecated)",
												MarkdownDescription: "Specifies the identifier of the service reference declaration. It corresponds to the serviceRefDeclaration name defined in either:  - 'componentDefinition.spec.serviceRefDeclarations[*].name' - 'clusterDefinition.spec.componentDefs[*].serviceRefDeclarations[*].name' (deprecated)",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object. If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current Cluster by default.",
												MarkdownDescription: "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object. If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current Cluster by default.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_descriptor": schema.StringAttribute{
												Description:         "Specifies the name of the ServiceDescriptor object that describes a service provided by external sources.  When referencing a service provided by external sources, a ServiceDescriptor object is required to establish the service binding. The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKind and serviceVersion declared in the definition.  If both 'cluster' and 'serviceDescriptor' are specified, the 'cluster' takes precedence.",
												MarkdownDescription: "Specifies the name of the ServiceDescriptor object that describes a service provided by external sources.  When referencing a service provided by external sources, a ServiceDescriptor object is required to establish the service binding. The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKind and serviceVersion declared in the definition.  If both 'cluster' and 'serviceDescriptor' are specified, the 'cluster' takes precedence.",
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

								"service_version": schema.StringAttribute{
									Description:         "ServiceVersion specifies the version of the Service expected to be provisioned by this Component. The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/). If no version is specified, the latest available version will be used.",
									MarkdownDescription: "ServiceVersion specifies the version of the Service expected to be provisioned by this Component. The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/). If no version is specified, the latest available version will be used.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(32),
									},
								},

								"services": schema.ListNestedAttribute{
									Description:         "Overrides services defined in referenced ComponentDefinition and expose endpoints that can be accessed by clients.",
									MarkdownDescription: "Overrides services defined in referenced ComponentDefinition and expose endpoints that can be accessed by clients.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "If ServiceType is LoadBalancer, cloud provider related parameters can be put here. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												MarkdownDescription: "If ServiceType is LoadBalancer, cloud provider related parameters can be put here. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "References the ComponentService name defined in the 'componentDefinition.spec.services[*].name'.",
												MarkdownDescription: "References the ComponentService name defined in the 'componentDefinition.spec.services[*].name'.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(25),
												},
											},

											"pod_service": schema.BoolAttribute{
												Description:         "Indicates whether to generate individual Services for each Pod. If set to true, a separate Service will be created for each Pod in the Cluster.",
												MarkdownDescription: "Indicates whether to generate individual Services for each Pod. If set to true, a separate Service will be created for each Pod in the Cluster.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_type": schema.StringAttribute{
												Description:         "Determines how the Service is exposed. Valid options are 'ClusterIP', 'NodePort', and 'LoadBalancer'.  - 'ClusterIP' allocates a Cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the ClusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the ClusterIP.  Note: although K8s Service type allows the 'ExternalName' type, it is not a valid option for ClusterComponentService.  For more info, see: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
												MarkdownDescription: "Determines how the Service is exposed. Valid options are 'ClusterIP', 'NodePort', and 'LoadBalancer'.  - 'ClusterIP' allocates a Cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the ClusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the ClusterIP.  Note: although K8s Service type allows the 'ExternalName' type, it is not a valid option for ClusterComponentService.  For more info, see: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("ClusterIP", "NodePort", "LoadBalancer"),
												},
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"switch_policy": schema.SingleNestedAttribute{
									Description:         "Defines the strategy for switchover and failover when workloadType is Replication.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
									MarkdownDescription: "Defines the strategy for switchover and failover when workloadType is Replication.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
									Attributes: map[string]schema.Attribute{
										"type": schema.StringAttribute{
											Description:         "Type specifies the type of switch policy to be applied.",
											MarkdownDescription: "Type specifies the type of switch policy to be applied.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Noop"),
											},
										},
									},
									Required: false,
									Optional: true,
									Computed: false,
								},

								"tls": schema.BoolAttribute{
									Description:         "A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication. When set to true, the Component will be configured to use TLS encryption for its network connections. This ensures that the data transmitted between the Component and its clients or other Components is encrypted and protected from unauthorized access. If TLS is enabled, the Component may require additional configuration, such as specifying TLS certificates and keys, to properly set up the secure communication channel.",
									MarkdownDescription: "A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication. When set to true, the Component will be configured to use TLS encryption for its network connections. This ensures that the data transmitted between the Component and its clients or other Components is encrypted and protected from unauthorized access. If TLS is enabled, the Component may require additional configuration, such as specifying TLS certificates and keys, to properly set up the secure communication channel.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"tolerations": schema.MapAttribute{
									Description:         "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.  Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
									MarkdownDescription: "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.  Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"update_strategy": schema.StringAttribute{
									Description:         "Defines the update strategy for the Component.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
									MarkdownDescription: "Defines the update strategy for the Component.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Serial", "BestEffortParallel", "Parallel"),
									},
								},

								"user_resource_refs": schema.SingleNestedAttribute{
									Description:         "Allows users to specify custom ConfigMaps and Secrets to be mounted as volumes in the Cluster's Pods. This is useful in scenarios where users need to provide additional resources to the Cluster, such as:  - Mounting custom scripts or configuration files during Cluster startup. - Mounting Secrets as volumes to provide sensitive information, like S3 AK/SK, to the Cluster.",
									MarkdownDescription: "Allows users to specify custom ConfigMaps and Secrets to be mounted as volumes in the Cluster's Pods. This is useful in scenarios where users need to provide additional resources to the Cluster, such as:  - Mounting custom scripts or configuration files during Cluster startup. - Mounting Secrets as volumes to provide sensitive information, like S3 AK/SK, to the Cluster.",
									Attributes: map[string]schema.Attribute{
										"config_map_refs": schema.ListNestedAttribute{
											Description:         "ConfigMapRefs defines the user-defined ConfigMaps.",
											MarkdownDescription: "ConfigMapRefs defines the user-defined ConfigMaps.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"as_volume_from": schema.ListAttribute{
														Description:         "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
														MarkdownDescription: "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"config_map": schema.SingleNestedAttribute{
														Description:         "ConfigMap specifies the ConfigMap to be mounted as a volume.",
														MarkdownDescription: "ConfigMap specifies the ConfigMap to be mounted as a volume.",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																			Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																			MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
														Required: true,
														Optional: false,
														Computed: false,
													},

													"mount_point": schema.StringAttribute{
														Description:         "MountPoint is the filesystem path where the volume will be mounted.",
														MarkdownDescription: "MountPoint is the filesystem path where the volume will be mounted.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(256),
															stringvalidator.RegexMatches(regexp.MustCompile(`^/[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
														MarkdownDescription: "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
														},
													},

													"sub_path": schema.StringAttribute{
														Description:         "SubPath specifies a path within the volume from which to mount.",
														MarkdownDescription: "SubPath specifies a path within the volume from which to mount.",
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

										"secret_refs": schema.ListNestedAttribute{
											Description:         "SecretRefs defines the user-defined Secrets.",
											MarkdownDescription: "SecretRefs defines the user-defined Secrets.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"as_volume_from": schema.ListAttribute{
														Description:         "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
														MarkdownDescription: "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"mount_point": schema.StringAttribute{
														Description:         "MountPoint is the filesystem path where the volume will be mounted.",
														MarkdownDescription: "MountPoint is the filesystem path where the volume will be mounted.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(256),
															stringvalidator.RegexMatches(regexp.MustCompile(`^/[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
														MarkdownDescription: "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(63),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
														},
													},

													"secret": schema.SingleNestedAttribute{
														Description:         "Secret specifies the Secret to be mounted as a volume.",
														MarkdownDescription: "Secret specifies the Secret to be mounted as a volume.",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																			Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																			MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

															"optional": schema.BoolAttribute{
																Description:         "optional field specify whether the Secret or its keys must be defined",
																MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"secret_name": schema.StringAttribute{
																Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},
														},
														Required: true,
														Optional: false,
														Computed: false,
													},

													"sub_path": schema.StringAttribute{
														Description:         "SubPath specifies a path within the volume from which to mount.",
														MarkdownDescription: "SubPath specifies a path within the volume from which to mount.",
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

								"volume_claim_templates": schema.ListNestedAttribute{
									Description:         "Specifies a list of PersistentVolumeClaim templates that represent the storage requirements for the Component. Each template specifies the desired characteristics of a persistent volume, such as storage class, size, and access modes. These templates are used to dynamically provision persistent volumes for the Component.",
									MarkdownDescription: "Specifies a list of PersistentVolumeClaim templates that represent the storage requirements for the Component. Each template specifies the desired characteristics of a persistent volume, such as storage class, size, and access modes. These templates are used to dynamically provision persistent volumes for the Component.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Refers to the name of a volumeMount defined in either:  - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)  The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
												MarkdownDescription: "Refers to the name of a volumeMount defined in either:  - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)  The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"spec": schema.SingleNestedAttribute{
												Description:         "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.  When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
												MarkdownDescription: "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.  When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
												Attributes: map[string]schema.Attribute{
													"access_modes": schema.MapAttribute{
														Description:         "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
														MarkdownDescription: "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
														MarkdownDescription: "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
														Attributes: map[string]schema.Attribute{
															"claims": schema.ListNestedAttribute{
																Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																			MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"requests": schema.MapAttribute{
																Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
														Description:         "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
														MarkdownDescription: "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"volume_mode": schema.StringAttribute{
														Description:         "Defines what type of volume is required by the claim, either Block or Filesystem.",
														MarkdownDescription: "Defines what type of volume is required by the claim, either Block or Filesystem.",
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
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"network": schema.SingleNestedAttribute{
						Description:         "The configuration of network.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						MarkdownDescription: "The configuration of network.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						Attributes: map[string]schema.Attribute{
							"host_network_accessible": schema.BoolAttribute{
								Description:         "Indicates whether the host network can be accessed. By default, this is set to false.",
								MarkdownDescription: "Indicates whether the host network can be accessed. By default, this is set to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"publicly_accessible": schema.BoolAttribute{
								Description:         "Indicates whether the network is accessible to the public. By default, this is set to false.",
								MarkdownDescription: "Indicates whether the network is accessible to the public. By default, this is set to false.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": schema.Int64Attribute{
						Description:         "Specifies the replicas of the first componentSpec, if the replicas of the first componentSpec is specified, this value will be ignored.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						MarkdownDescription: "Specifies the replicas of the first componentSpec, if the replicas of the first componentSpec is specified, this value will be ignored.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"resources": schema.SingleNestedAttribute{
						Description:         "Specifies the resources of the first componentSpec, if the resources of the first componentSpec is specified, this value will be ignored.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						MarkdownDescription: "Specifies the resources of the first componentSpec, if the resources of the first componentSpec is specified, this value will be ignored.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						Attributes: map[string]schema.Attribute{
							"cpu": schema.StringAttribute{
								Description:         "Specifies the amount of CPU resource the Cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Specifies the amount of CPU resource the Cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"memory": schema.StringAttribute{
								Description:         "Specifies the amount of memory resource the Cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Specifies the amount of memory resource the Cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"runtime_class_name": schema.StringAttribute{
						Description:         "Specifies runtimeClassName for all Pods managed by this Cluster.",
						MarkdownDescription: "Specifies runtimeClassName for all Pods managed by this Cluster.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"scheduling_policy": schema.SingleNestedAttribute{
						Description:         "Specifies the scheduling policy for the Cluster.",
						MarkdownDescription: "Specifies the scheduling policy for the Cluster.",
						Attributes: map[string]schema.Attribute{
							"affinity": schema.SingleNestedAttribute{
								Description:         "Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.",
								MarkdownDescription: "Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.",
								Attributes: map[string]schema.Attribute{
									"node_affinity": schema.SingleNestedAttribute{
										Description:         "Describes node affinity scheduling rules for the pod.",
										MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
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
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
												Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
												MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
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
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																				Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
										Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
										MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "Required. A pod affinity term, associated with the corresponding weight.",
															MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																	Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																	Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
															Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
															MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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
												Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
										Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
										MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
										Attributes: map[string]schema.Attribute{
											"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
												Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"pod_affinity_term": schema.SingleNestedAttribute{
															Description:         "Required. A pod affinity term, associated with the corresponding weight.",
															MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
															Attributes: map[string]schema.Attribute{
																"label_selector": schema.SingleNestedAttribute{
																	Description:         "A label query over a set of resources, in this case pods.",
																	MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																	Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																	MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																						Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																						Required:            true,
																						Optional:            false,
																						Computed:            false,
																					},

																					"values": schema.ListAttribute{
																						Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																						MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																			Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																			MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																	Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"topology_key": schema.StringAttribute{
																	Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																	MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
															Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
															MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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
												Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"label_selector": schema.SingleNestedAttribute{
															Description:         "A label query over a set of resources, in this case pods.",
															MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
															MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
															Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"topology_key": schema.StringAttribute{
															Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
															MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

							"node_name": schema.StringAttribute{
								Description:         "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty, the scheduler simply schedules this Pod onto that node, assuming that it fits resource requirements.",
								MarkdownDescription: "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty, the scheduler simply schedules this Pod onto that node, assuming that it fits resource requirements.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the Pod to fit on a node. Selector which must match a node's labels for the Pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
								MarkdownDescription: "NodeSelector is a selector which must be true for the Pod to fit on a node. Selector which must match a node's labels for the Pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"scheduler_name": schema.StringAttribute{
								Description:         "If specified, the Pod will be dispatched by specified scheduler. If not specified, the Pod will be dispatched by default scheduler.",
								MarkdownDescription: "If specified, the Pod will be dispatched by specified scheduler. If not specified, the Pod will be dispatched by default scheduler.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"tolerations": schema.ListNestedAttribute{
								Description:         "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
								MarkdownDescription: "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"effect": schema.StringAttribute{
											Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"key": schema.StringAttribute{
											Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"operator": schema.StringAttribute{
											Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"toleration_seconds": schema.Int64Attribute{
											Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"value": schema.StringAttribute{
											Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
											MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
								Description:         "TopologySpreadConstraints describes how a group of Pods ought to spread across topology domains. Scheduler will schedule Pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
								MarkdownDescription: "TopologySpreadConstraints describes how a group of Pods ought to spread across topology domains. Scheduler will schedule Pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"label_selector": schema.SingleNestedAttribute{
											Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
											MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
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
																Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"values": schema.ListAttribute{
																Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
													Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
													MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
											Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
											MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"max_skew": schema.Int64Attribute{
											Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
											MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"min_domains": schema.Int64Attribute{
											Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
											MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_affinity_policy": schema.StringAttribute{
											Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"node_taints_policy": schema.StringAttribute{
											Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"topology_key": schema.StringAttribute{
											Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
											MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"when_unsatisfiable": schema.StringAttribute{
											Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
											MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
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

					"services": schema.MapAttribute{
						Description:         "Defines a list of additional Services that are exposed by a Cluster. This field allows Services of selected Components, either from 'componentSpecs' or 'shardingSpecs' to be exposed, alongside Services defined with ComponentService.  Services defined here can be referenced by other clusters using the ServiceRefClusterSelector.",
						MarkdownDescription: "Defines a list of additional Services that are exposed by a Cluster. This field allows Services of selected Components, either from 'componentSpecs' or 'shardingSpecs' to be exposed, alongside Services defined with ComponentService.  Services defined here can be referenced by other clusters using the ServiceRefClusterSelector.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"sharding_specs": schema.ListNestedAttribute{
						Description:         "Specifies a list of ShardingSpec objects that manage the sharding topology for Cluster Components. Each ShardingSpec organizes components into shards, with each shard corresponding to a Component. Components within a shard are all based on a common ClusterComponentSpec template, ensuring uniform configurations.  This field supports dynamic resharding by facilitating the addition or removal of shards through the 'shards' field in ShardingSpec.  Note: 'shardingSpecs' and 'componentSpecs' cannot both be empty; at least one must be defined to configure a Cluster.",
						MarkdownDescription: "Specifies a list of ShardingSpec objects that manage the sharding topology for Cluster Components. Each ShardingSpec organizes components into shards, with each shard corresponding to a Component. Components within a shard are all based on a common ClusterComponentSpec template, ensuring uniform configurations.  This field supports dynamic resharding by facilitating the addition or removal of shards through the 'shards' field in ShardingSpec.  Note: 'shardingSpecs' and 'componentSpecs' cannot both be empty; at least one must be defined to configure a Cluster.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "Represents the common parent part of all shard names. This identifier is included as part of the Service DNS name and must comply with IANA service naming rules. It is used to generate the names of underlying Components following the pattern '$(shardingSpec.name)-$(ShardID)'. ShardID is a random string that is appended to the Name to generate unique identifiers for each shard. For example, if the sharding specification name is 'my-shard' and the ShardID is 'abc', the resulting Component name would be 'my-shard-abc'.  Note that the name defined in Component template('shardingSpec.template.name') will be disregarded when generating the Component names of the shards. The 'shardingSpec.name' field takes precedence.",
									MarkdownDescription: "Represents the common parent part of all shard names. This identifier is included as part of the Service DNS name and must comply with IANA service naming rules. It is used to generate the names of underlying Components following the pattern '$(shardingSpec.name)-$(ShardID)'. ShardID is a random string that is appended to the Name to generate unique identifiers for each shard. For example, if the sharding specification name is 'my-shard' and the ShardID is 'abc', the resulting Component name would be 'my-shard-abc'.  Note that the name defined in Component template('shardingSpec.template.name') will be disregarded when generating the Component names of the shards. The 'shardingSpec.name' field takes precedence.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.LengthAtMost(15),
										stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
									},
								},

								"shards": schema.Int64Attribute{
									Description:         "Specifies the desired number of shards. Users can declare the desired number of shards through this field. KubeBlocks dynamically creates and deletes Components based on the difference between the desired and actual number of shards. KubeBlocks provides lifecycle management for sharding, including:  - Executing the postProvision Action defined in the ComponentDefinition when the number of shards increases. This allows for custom actions to be performed after a new shard is provisioned. - Executing the preTerminate Action defined in the ComponentDefinition when the number of shards decreases. This enables custom cleanup or data migration tasks to be executed before a shard is terminated. Resources and data associated with the corresponding Component will also be deleted.",
									MarkdownDescription: "Specifies the desired number of shards. Users can declare the desired number of shards through this field. KubeBlocks dynamically creates and deletes Components based on the difference between the desired and actual number of shards. KubeBlocks provides lifecycle management for sharding, including:  - Executing the postProvision Action defined in the ComponentDefinition when the number of shards increases. This allows for custom actions to be performed after a new shard is provisioned. - Executing the preTerminate Action defined in the ComponentDefinition when the number of shards decreases. This enables custom cleanup or data migration tasks to be executed before a shard is terminated. Resources and data associated with the corresponding Component will also be deleted.",
									Required:            false,
									Optional:            true,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
										int64validator.AtMost(2048),
									},
								},

								"template": schema.SingleNestedAttribute{
									Description:         "The template for generating Components for shards, where each shard consists of one Component. This field is of type ClusterComponentSpec, which encapsulates all the required details and definitions for creating and managing the Components. KubeBlocks uses this template to generate a set of identical Components or shards. All the generated Components will have the same specifications and definitions as specified in the 'template' field.  This allows for the creation of multiple Components with consistent configurations, enabling sharding and distribution of workloads across Components.",
									MarkdownDescription: "The template for generating Components for shards, where each shard consists of one Component. This field is of type ClusterComponentSpec, which encapsulates all the required details and definitions for creating and managing the Components. KubeBlocks uses this template to generate a set of identical Components or shards. All the generated Components will have the same specifications and definitions as specified in the 'template' field.  This allows for the creation of multiple Components with consistent configurations, enabling sharding and distribution of workloads across Components.",
									Attributes: map[string]schema.Attribute{
										"affinity": schema.SingleNestedAttribute{
											Description:         "Specifies a group of affinity scheduling rules for the Component. It allows users to control how the Component's Pods are scheduled onto nodes in the K8s cluster.  Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
											MarkdownDescription: "Specifies a group of affinity scheduling rules for the Component. It allows users to control how the Component's Pods are scheduled onto nodes in the K8s cluster.  Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
											Attributes: map[string]schema.Attribute{
												"node_labels": schema.MapAttribute{
													Description:         "Indicates the node labels that must be present on nodes for pods to be scheduled on them. It is a map where the keys are the label keys and the values are the corresponding label values. Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.  For example, if NodeLabels is set to {'nodeType': 'ssd', 'environment': 'production'}, pods will only be scheduled on nodes that have both the 'nodeType' label with value 'ssd' and the 'environment' label with value 'production'.  This field allows users to control Pod placement based on specific node labels. It can be used to ensure that Pods are scheduled on nodes with certain characteristics, such as specific hardware (e.g., SSD), environment (e.g., production, staging), or any other custom labels assigned to nodes.",
													MarkdownDescription: "Indicates the node labels that must be present on nodes for pods to be scheduled on them. It is a map where the keys are the label keys and the values are the corresponding label values. Pods will only be scheduled on nodes that have all the specified labels with the corresponding values.  For example, if NodeLabels is set to {'nodeType': 'ssd', 'environment': 'production'}, pods will only be scheduled on nodes that have both the 'nodeType' label with value 'ssd' and the 'environment' label with value 'production'.  This field allows users to control Pod placement based on specific node labels. It can be used to ensure that Pods are scheduled on nodes with certain characteristics, such as specific hardware (e.g., SSD), environment (e.g., production, staging), or any other custom labels assigned to nodes.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pod_anti_affinity": schema.StringAttribute{
													Description:         "Specifies the anti-affinity level of Pods within a Component. It determines how pods should be spread across nodes to improve availability and performance. It can have the following values: 'Preferred' and 'Required'. The default value is 'Preferred'.",
													MarkdownDescription: "Specifies the anti-affinity level of Pods within a Component. It determines how pods should be spread across nodes to improve availability and performance. It can have the following values: 'Preferred' and 'Required'. The default value is 'Preferred'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Preferred", "Required"),
													},
												},

												"tenancy": schema.StringAttribute{
													Description:         "Determines the level of resource isolation between Pods. It can have the following values: 'SharedNode' and 'DedicatedNode'.  - SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s. - DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node. In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node. Which provides a higher level of isolation and resource guarantee for Pods.  The default value is 'SharedNode'.",
													MarkdownDescription: "Determines the level of resource isolation between Pods. It can have the following values: 'SharedNode' and 'DedicatedNode'.  - SharedNode: Allow that multiple Pods may share the same node, which is the default behavior of K8s. - DedicatedNode: Each Pod runs on a dedicated node, ensuring that no two Pods share the same node. In other words, if a Pod is already running on a node, no other Pods will be scheduled on that node. Which provides a higher level of isolation and resource guarantee for Pods.  The default value is 'SharedNode'.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("SharedNode", "DedicatedNode"),
													},
												},

												"topology_keys": schema.ListAttribute{
													Description:         "Represents the key of node labels used to define the topology domain for Pod anti-affinity and Pod spread constraints.  In K8s, a topology domain is a set of nodes that have the same value for a specific label key. Nodes with labels containing any of the specified TopologyKeys and identical values are considered to be in the same topology domain.  Note: The concept of topology in the context of K8s TopologyKeys is different from the concept of topology in the ClusterDefinition.  When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule the Pod on nodes with different values for the specified TopologyKeys. This ensures that Pods are spread across different topology domains, promoting high availability and reducing the impact of node failures.  Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey. These keys represent the hostname and zone of a node, respectively. By including these keys in the TopologyKeys list, Pods will be spread across nodes with different hostnames or zones.  In addition to the well-known keys, users can also specify custom label keys as TopologyKeys. This allows for more flexible and custom topology definitions based on the specific needs of the application or environment.  The TopologyKeys field is a slice of strings, where each string represents a label key. The order of the keys in the slice does not matter.",
													MarkdownDescription: "Represents the key of node labels used to define the topology domain for Pod anti-affinity and Pod spread constraints.  In K8s, a topology domain is a set of nodes that have the same value for a specific label key. Nodes with labels containing any of the specified TopologyKeys and identical values are considered to be in the same topology domain.  Note: The concept of topology in the context of K8s TopologyKeys is different from the concept of topology in the ClusterDefinition.  When a Pod has anti-affinity or spread constraints specified, Kubernetes will attempt to schedule the Pod on nodes with different values for the specified TopologyKeys. This ensures that Pods are spread across different topology domains, promoting high availability and reducing the impact of node failures.  Some well-known label keys, such as 'kubernetes.io/hostname' and 'topology.kubernetes.io/zone', are often used as TopologyKey. These keys represent the hostname and zone of a node, respectively. By including these keys in the TopologyKeys list, Pods will be spread across nodes with different hostnames or zones.  In addition to the well-known keys, users can also specify custom label keys as TopologyKeys. This allows for more flexible and custom topology definitions based on the specific needs of the application or environment.  The TopologyKeys field is a slice of strings, where each string represents a label key. The order of the keys in the slice does not matter.",
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

										"component_def": schema.StringAttribute{
											Description:         "References the name of a ComponentDefinition object. The ComponentDefinition specifies the behavior and characteristics of the Component. If both 'componentDefRef' and 'componentDef' are provided, the 'componentDef' will take precedence over 'componentDefRef'.",
											MarkdownDescription: "References the name of a ComponentDefinition object. The ComponentDefinition specifies the behavior and characteristics of the Component. If both 'componentDefRef' and 'componentDef' are provided, the 'componentDef' will take precedence over 'componentDefRef'.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(64),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"component_def_ref": schema.StringAttribute{
											Description:         "References a ClusterComponentDefinition defined in the 'clusterDefinition.spec.componentDef' field. Must comply with the IANA service naming rule.  Deprecated since v0.9, because defining Components in 'clusterDefinition.spec.componentDef' field has been deprecated. This field is replaced by the 'componentDef' field, use 'componentDef' instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDefRef is immutable'",
											MarkdownDescription: "References a ClusterComponentDefinition defined in the 'clusterDefinition.spec.componentDef' field. Must comply with the IANA service naming rule.  Deprecated since v0.9, because defining Components in 'clusterDefinition.spec.componentDef' field has been deprecated. This field is replaced by the 'componentDef' field, use 'componentDef' instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='componentDefRef is immutable'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(22),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
										},

										"configs": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config_map": schema.SingleNestedAttribute{
														Description:         "ConfigMap source for the config.",
														MarkdownDescription: "ConfigMap source for the config.",
														Attributes: map[string]schema.Attribute{
															"default_mode": schema.Int64Attribute{
																Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"items": schema.ListNestedAttribute{
																Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																			Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																			MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

													"name": schema.StringAttribute{
														Description:         "The name of the config.",
														MarkdownDescription: "The name of the config.",
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

										"disable_exporter": schema.BoolAttribute{
											Description:         "Determines whether metrics exporter information is annotated on the Component's headless Service.  If set to true, the following annotations will not be patched into the Service:  - 'monitor.kubeblocks.io/path' - 'monitor.kubeblocks.io/port' - 'monitor.kubeblocks.io/scheme'  These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
											MarkdownDescription: "Determines whether metrics exporter information is annotated on the Component's headless Service.  If set to true, the following annotations will not be patched into the Service:  - 'monitor.kubeblocks.io/path' - 'monitor.kubeblocks.io/port' - 'monitor.kubeblocks.io/scheme'  These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"enabled_logs": schema.ListAttribute{
											Description:         "Specifies which types of logs should be collected for the Component. The log types are defined in the 'componentDefinition.spec.logConfigs' field with the LogConfig entries.  The elements in the 'enabledLogs' array correspond to the names of the LogConfig entries. For example, if the 'componentDefinition.spec.logConfigs' defines LogConfig entries with names 'slow_query_log' and 'error_log', you can enable the collection of these logs by including their names in the 'enabledLogs' array: '''yaml enabledLogs: - slow_query_log - error_log '''",
											MarkdownDescription: "Specifies which types of logs should be collected for the Component. The log types are defined in the 'componentDefinition.spec.logConfigs' field with the LogConfig entries.  The elements in the 'enabledLogs' array correspond to the names of the LogConfig entries. For example, if the 'componentDefinition.spec.logConfigs' defines LogConfig entries with names 'slow_query_log' and 'error_log', you can enable the collection of these logs by including their names in the 'enabledLogs' array: '''yaml enabledLogs: - slow_query_log - error_log '''",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"instances": schema.ListNestedAttribute{
											Description:         "Allows for the customization of configuration values for each instance within a Component. An instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps). While instances typically share a common configuration as defined in the ClusterComponentSpec, they can require unique settings in various scenarios:  For example: - A database Component might require different resource allocations for primary and secondary instances, with primaries needing more resources. - During a rolling upgrade, a Component may first update the image for one or a few instances, and then update the remaining instances after verifying that the updated instances are functioning correctly.  InstanceTemplate allows for specifying these unique configurations per instance. Each instance's name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal), starting with an ordinal of 0. It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.  The sum of replicas across all InstanceTemplates should not exceed the total number of replicas specified for the Component. Any remaining replicas will be generated using the default template and will follow the default naming rules.",
											MarkdownDescription: "Allows for the customization of configuration values for each instance within a Component. An instance represent a single replica (Pod and associated K8s resources like PVCs, Services, and ConfigMaps). While instances typically share a common configuration as defined in the ClusterComponentSpec, they can require unique settings in various scenarios:  For example: - A database Component might require different resource allocations for primary and secondary instances, with primaries needing more resources. - During a rolling upgrade, a Component may first update the image for one or a few instances, and then update the remaining instances after verifying that the updated instances are functioning correctly.  InstanceTemplate allows for specifying these unique configurations per instance. Each instance's name is constructed using the pattern: $(component.name)-$(template.name)-$(ordinal), starting with an ordinal of 0. It is crucial to maintain unique names for each InstanceTemplate to avoid conflicts.  The sum of replicas across all InstanceTemplates should not exceed the total number of replicas specified for the Component. Any remaining replicas will be generated using the default template and will follow the default naming rules.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "Specifies a map of key-value pairs to be merged into the Pod's existing annotations. Existing keys will have their values overwritten, while new keys will be added to the annotations.",
														MarkdownDescription: "Specifies a map of key-value pairs to be merged into the Pod's existing annotations. Existing keys will have their values overwritten, while new keys will be added to the annotations.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"env": schema.ListNestedAttribute{
														Description:         "Defines Env to override. Add new or override existing envs.",
														MarkdownDescription: "Defines Env to override. Add new or override existing envs.",
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
																	Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																	MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
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
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																			Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
																			MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
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
																			Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																			MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
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
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

													"image": schema.StringAttribute{
														Description:         "Specifies an override for the first container's image in the Pod.",
														MarkdownDescription: "Specifies an override for the first container's image in the Pod.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"labels": schema.MapAttribute{
														Description:         "Specifies a map of key-value pairs that will be merged into the Pod's existing labels. Values for existing keys will be overwritten, and new keys will be added.",
														MarkdownDescription: "Specifies a map of key-value pairs that will be merged into the Pod's existing labels. Values for existing keys will be overwritten, and new keys will be added.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name specifies the unique name of the instance Pod created using this InstanceTemplate. This name is constructed by concatenating the Component's name, the template's name, and the instance's ordinal using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0. The specified name overrides any default naming conventions or patterns.",
														MarkdownDescription: "Name specifies the unique name of the instance Pod created using this InstanceTemplate. This name is constructed by concatenating the Component's name, the template's name, and the instance's ordinal using the pattern: $(cluster.name)-$(component.name)-$(template.name)-$(ordinal). Ordinals start from 0. The specified name overrides any default naming conventions or patterns.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(54),
															stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
														},
													},

													"node_name": schema.StringAttribute{
														Description:         "Specifies the name of the node where the Pod should be scheduled. If set, the Pod will be directly assigned to the specified node, bypassing the Kubernetes scheduler. This is useful for controlling Pod placement on specific nodes.  Important considerations: - 'nodeName' bypasses default scheduling constraints (e.g., resource requirements, node selectors, affinity rules). - It is the user's responsibility to ensure the node is suitable for the Pod. - If the node is unavailable, the Pod will remain in 'Pending' state until the node is available or the Pod is deleted.",
														MarkdownDescription: "Specifies the name of the node where the Pod should be scheduled. If set, the Pod will be directly assigned to the specified node, bypassing the Kubernetes scheduler. This is useful for controlling Pod placement on specific nodes.  Important considerations: - 'nodeName' bypasses default scheduling constraints (e.g., resource requirements, node selectors, affinity rules). - It is the user's responsibility to ensure the node is suitable for the Pod. - If the node is unavailable, the Pod will remain in 'Pending' state until the node is available or the Pod is deleted.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"node_selector": schema.MapAttribute{
														Description:         "Defines NodeSelector to override.",
														MarkdownDescription: "Defines NodeSelector to override.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"replicas": schema.Int64Attribute{
														Description:         "Specifies the number of instances (Pods) to create from this InstanceTemplate. This field allows setting how many replicated instances of the Component, with the specific overrides in the InstanceTemplate, are created. The default value is 1. A value of 0 disables instance creation.",
														MarkdownDescription: "Specifies the number of instances (Pods) to create from this InstanceTemplate. This field allows setting how many replicated instances of the Component, with the specific overrides in the InstanceTemplate, are created. The default value is 1. A value of 0 disables instance creation.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
														},
													},

													"resources": schema.SingleNestedAttribute{
														Description:         "Specifies an override for the resource requirements of the first container in the Pod. This field allows for customizing resource allocation (CPU, memory, etc.) for the container.",
														MarkdownDescription: "Specifies an override for the resource requirements of the first container in the Pod. This field allows for customizing resource allocation (CPU, memory, etc.) for the container.",
														Attributes: map[string]schema.Attribute{
															"claims": schema.ListNestedAttribute{
																Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"name": schema.StringAttribute{
																			Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																			MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"requests": schema.MapAttribute{
																Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

													"tolerations": schema.ListNestedAttribute{
														Description:         "Tolerations specifies a list of tolerations to be applied to the Pod, allowing it to tolerate node taints. This field can be used to add new tolerations or override existing ones.",
														MarkdownDescription: "Tolerations specifies a list of tolerations to be applied to the Pod, allowing it to tolerate node taints. This field can be used to add new tolerations or override existing ones.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"effect": schema.StringAttribute{
																	Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																	MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"key": schema.StringAttribute{
																	Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																	MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"operator": schema.StringAttribute{
																	Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																	MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"toleration_seconds": schema.Int64Attribute{
																	Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																	MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
																	MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

													"volume_claim_templates": schema.ListNestedAttribute{
														Description:         "Defines VolumeClaimTemplates to override. Add new or override existing volume claim templates.",
														MarkdownDescription: "Defines VolumeClaimTemplates to override. Add new or override existing volume claim templates.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Refers to the name of a volumeMount defined in either:  - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)  The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
																	MarkdownDescription: "Refers to the name of a volumeMount defined in either:  - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)  The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"spec": schema.SingleNestedAttribute{
																	Description:         "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.  When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
																	MarkdownDescription: "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.  When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
																	Attributes: map[string]schema.Attribute{
																		"access_modes": schema.MapAttribute{
																			Description:         "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																			MarkdownDescription: "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"resources": schema.SingleNestedAttribute{
																			Description:         "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
																			MarkdownDescription: "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
																			Attributes: map[string]schema.Attribute{
																				"claims": schema.ListNestedAttribute{
																					Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																					MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"name": schema.StringAttribute{
																								Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																								MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																					Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																					MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"requests": schema.MapAttribute{
																					Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																					MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
																			Description:         "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																			MarkdownDescription: "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_mode": schema.StringAttribute{
																			Description:         "Defines what type of volume is required by the claim, either Block or Filesystem.",
																			MarkdownDescription: "Defines what type of volume is required by the claim, either Block or Filesystem.",
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

													"volume_mounts": schema.ListNestedAttribute{
														Description:         "Defines VolumeMounts to override. Add new or override existing volume mounts of the first container in the Pod.",
														MarkdownDescription: "Defines VolumeMounts to override. Add new or override existing volume mounts of the first container in the Pod.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"mount_path": schema.StringAttribute{
																	Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
																	MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"mount_propagation": schema.StringAttribute{
																	Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
																	MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"name": schema.StringAttribute{
																	Description:         "This must match the Name of a Volume.",
																	MarkdownDescription: "This must match the Name of a Volume.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"read_only": schema.BoolAttribute{
																	Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
																	MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"sub_path": schema.StringAttribute{
																	Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
																	MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
																	Required:            false,
																	Optional:            true,
																	Computed:            false,
																},

																"sub_path_expr": schema.StringAttribute{
																	Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
																	MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

													"volumes": schema.ListNestedAttribute{
														Description:         "Defines Volumes to override. Add new or override existing volumes.",
														MarkdownDescription: "Defines Volumes to override. Add new or override existing volumes.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"aws_elastic_block_store": schema.SingleNestedAttribute{
																	Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"partition": schema.Int64Attribute{
																			Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																			MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																			MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_id": schema.StringAttribute{
																			Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																			MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"azure_disk": schema.SingleNestedAttribute{
																	Description:         "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
																	MarkdownDescription: "azureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
																	Attributes: map[string]schema.Attribute{
																		"caching_mode": schema.StringAttribute{
																			Description:         "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																			MarkdownDescription: "cachingMode is the Host Caching mode: None, Read Only, Read Write.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"disk_name": schema.StringAttribute{
																			Description:         "diskName is the Name of the data disk in the blob storage",
																			MarkdownDescription: "diskName is the Name of the data disk in the blob storage",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"disk_uri": schema.StringAttribute{
																			Description:         "diskURI is the URI of data disk in the blob storage",
																			MarkdownDescription: "diskURI is the URI of data disk in the blob storage",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fsType is Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"kind": schema.StringAttribute{
																			Description:         "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																			MarkdownDescription: "kind expected values are Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"azure_file": schema.SingleNestedAttribute{
																	Description:         "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
																	MarkdownDescription: "azureFile represents an Azure File Service mount on the host and bind mount to the pod.",
																	Attributes: map[string]schema.Attribute{
																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_name": schema.StringAttribute{
																			Description:         "secretName is the  name of secret that contains Azure Storage Account Name and Key",
																			MarkdownDescription: "secretName is the  name of secret that contains Azure Storage Account Name and Key",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"share_name": schema.StringAttribute{
																			Description:         "shareName is the azure share Name",
																			MarkdownDescription: "shareName is the azure share Name",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"cephfs": schema.SingleNestedAttribute{
																	Description:         "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
																	MarkdownDescription: "cephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
																	Attributes: map[string]schema.Attribute{
																		"monitors": schema.ListAttribute{
																			Description:         "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																			MarkdownDescription: "path is Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_file": schema.StringAttribute{
																			Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"user": schema.StringAttribute{
																			Description:         "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			MarkdownDescription: "user is optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"cinder": schema.SingleNestedAttribute{
																	Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																			MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connect to OpenStack.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"volume_id": schema.StringAttribute{
																			Description:         "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			MarkdownDescription: "volumeID used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"config_map": schema.SingleNestedAttribute{
																	Description:         "configMap represents a configMap that should populate this volume",
																	MarkdownDescription: "configMap represents a configMap that should populate this volume",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																			MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																						Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																						MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																			Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																			MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

																"csi": schema.SingleNestedAttribute{
																	Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
																	MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
																	Attributes: map[string]schema.Attribute{
																		"driver": schema.StringAttribute{
																			Description:         "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																			MarkdownDescription: "driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																			MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"node_publish_secret_ref": schema.SingleNestedAttribute{
																			Description:         "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																			MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																			MarkdownDescription: "readOnly specifies a read-only configuration for the volume. Defaults to false (read/write).",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_attributes": schema.MapAttribute{
																			Description:         "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
																			MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
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

																"downward_api": schema.SingleNestedAttribute{
																	Description:         "downwardAPI represents downward API about the pod that should populate this volume",
																	MarkdownDescription: "downwardAPI represents downward API about the pod that should populate this volume",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "Items is a list of downward API volume file",
																			MarkdownDescription: "Items is a list of downward API volume file",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"field_ref": schema.SingleNestedAttribute{
																						Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																						MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
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
																						Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
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
																						Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																						MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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

																"empty_dir": schema.SingleNestedAttribute{
																	Description:         "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																	MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																	Attributes: map[string]schema.Attribute{
																		"medium": schema.StringAttribute{
																			Description:         "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																			MarkdownDescription: "medium represents what type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"size_limit": schema.StringAttribute{
																			Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																			MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"ephemeral": schema.SingleNestedAttribute{
																	Description:         "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
																	MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
																	Attributes: map[string]schema.Attribute{
																		"volume_claim_template": schema.SingleNestedAttribute{
																			Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
																			MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
																			Attributes: map[string]schema.Attribute{
																				"metadata": schema.SingleNestedAttribute{
																					Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																					MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																					Attributes: map[string]schema.Attribute{
																						"annotations": schema.MapAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"finalizers": schema.ListAttribute{
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

																						"name": schema.StringAttribute{
																							Description:         "",
																							MarkdownDescription: "",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"namespace": schema.StringAttribute{
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

																				"spec": schema.SingleNestedAttribute{
																					Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																					MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																					Attributes: map[string]schema.Attribute{
																						"access_modes": schema.ListAttribute{
																							Description:         "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																							MarkdownDescription: "accessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"data_source": schema.SingleNestedAttribute{
																							Description:         "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																							MarkdownDescription: "dataSource field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. When the AnyVolumeDataSource feature gate is enabled, dataSource contents will be copied to dataSourceRef, and dataSourceRef contents will be copied to dataSource when dataSourceRef.namespace is not specified. If the namespace is specified, then dataSourceRef will not be copied to dataSource.",
																							Attributes: map[string]schema.Attribute{
																								"api_group": schema.StringAttribute{
																									Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																									MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
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
																							Description:         "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																							MarkdownDescription: "dataSourceRef specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the dataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, when namespace isn't specified in dataSourceRef, both fields (dataSource and dataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. When namespace is specified in dataSourceRef, dataSource isn't set to the same value and must be empty. There are three important differences between dataSource and dataSourceRef: * While dataSource only allows two specific types of objects, dataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While dataSource ignores disallowed values (dropping them), dataSourceRef preserves all values, and generates an error if a disallowed value is specified. * While dataSource only allows local objects, dataSourceRef allows objects in any namespaces. (Beta) Using this field requires the AnyVolumeDataSource feature gate to be enabled. (Alpha) Using the namespace field of dataSourceRef requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																							Attributes: map[string]schema.Attribute{
																								"api_group": schema.StringAttribute{
																									Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																									MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
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
																									Description:         "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
																									MarkdownDescription: "Namespace is the namespace of resource being referenced Note that when a namespace is specified, a gateway.networking.k8s.io/ReferenceGrant object is required in the referent namespace to allow that namespace's owner to accept the reference. See the ReferenceGrant documentation for details. (Alpha) This field requires the CrossNamespaceVolumeDataSource feature gate to be enabled.",
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
																							Description:         "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																							MarkdownDescription: "resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																							Attributes: map[string]schema.Attribute{
																								"claims": schema.ListNestedAttribute{
																									Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																									MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																									NestedObject: schema.NestedAttributeObject{
																										Attributes: map[string]schema.Attribute{
																											"name": schema.StringAttribute{
																												Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																												MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																									Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																									MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																									ElementType:         types.StringType,
																									Required:            false,
																									Optional:            true,
																									Computed:            false,
																								},

																								"requests": schema.MapAttribute{
																									Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																									MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
																												Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																												MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																												Required:            true,
																												Optional:            false,
																												Computed:            false,
																											},

																											"values": schema.ListAttribute{
																												Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																												MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																									Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																									MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																							Description:         "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																							MarkdownDescription: "storageClassName is the name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_mode": schema.StringAttribute{
																							Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																							MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
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

																"fc": schema.SingleNestedAttribute{
																	Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
																	MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"lun": schema.Int64Attribute{
																			Description:         "lun is Optional: FC target lun number",
																			MarkdownDescription: "lun is Optional: FC target lun number",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"target_ww_ns": schema.ListAttribute{
																			Description:         "targetWWNs is Optional: FC target worldwide names (WWNs)",
																			MarkdownDescription: "targetWWNs is Optional: FC target worldwide names (WWNs)",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"wwids": schema.ListAttribute{
																			Description:         "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
																			MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
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

																"flex_volume": schema.SingleNestedAttribute{
																	Description:         "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
																	MarkdownDescription: "flexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
																	Attributes: map[string]schema.Attribute{
																		"driver": schema.StringAttribute{
																			Description:         "driver is the name of the driver to use for this volume.",
																			MarkdownDescription: "driver is the name of the driver to use for this volume.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"options": schema.MapAttribute{
																			Description:         "options is Optional: this field holds extra command options if any.",
																			MarkdownDescription: "options is Optional: this field holds extra command options if any.",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																			MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

																"flocker": schema.SingleNestedAttribute{
																	Description:         "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
																	MarkdownDescription: "flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
																	Attributes: map[string]schema.Attribute{
																		"dataset_name": schema.StringAttribute{
																			Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																			MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"dataset_uuid": schema.StringAttribute{
																			Description:         "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																			MarkdownDescription: "datasetUUID is the UUID of the dataset. This is unique identifier of a Flocker dataset",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"gce_persistent_disk": schema.SingleNestedAttribute{
																	Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"partition": schema.Int64Attribute{
																			Description:         "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			MarkdownDescription: "partition is the partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"pd_name": schema.StringAttribute{
																			Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"git_repo": schema.SingleNestedAttribute{
																	Description:         "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
																	MarkdownDescription: "gitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
																	Attributes: map[string]schema.Attribute{
																		"directory": schema.StringAttribute{
																			Description:         "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																			MarkdownDescription: "directory is the target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"repository": schema.StringAttribute{
																			Description:         "repository is the URL",
																			MarkdownDescription: "repository is the URL",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"revision": schema.StringAttribute{
																			Description:         "revision is the commit hash for the specified revision.",
																			MarkdownDescription: "revision is the commit hash for the specified revision.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"glusterfs": schema.SingleNestedAttribute{
																	Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
																	MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
																	Attributes: map[string]schema.Attribute{
																		"endpoints": schema.StringAttribute{
																			Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"path": schema.StringAttribute{
																			Description:         "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			MarkdownDescription: "path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"host_path": schema.SingleNestedAttribute{
																	Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
																	MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
																	Attributes: map[string]schema.Attribute{
																		"path": schema.StringAttribute{
																			Description:         "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																			MarkdownDescription: "path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"type": schema.StringAttribute{
																			Description:         "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																			MarkdownDescription: "type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"iscsi": schema.SingleNestedAttribute{
																	Description:         "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
																	MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
																	Attributes: map[string]schema.Attribute{
																		"chap_auth_discovery": schema.BoolAttribute{
																			Description:         "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																			MarkdownDescription: "chapAuthDiscovery defines whether support iSCSI Discovery CHAP authentication",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"chap_auth_session": schema.BoolAttribute{
																			Description:         "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																			MarkdownDescription: "chapAuthSession defines whether support iSCSI Session CHAP authentication",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"initiator_name": schema.StringAttribute{
																			Description:         "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																			MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"iqn": schema.StringAttribute{
																			Description:         "iqn is the target iSCSI Qualified Name.",
																			MarkdownDescription: "iqn is the target iSCSI Qualified Name.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"iscsi_interface": schema.StringAttribute{
																			Description:         "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																			MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"lun": schema.Int64Attribute{
																			Description:         "lun represents iSCSI Target Lun number.",
																			MarkdownDescription: "lun represents iSCSI Target Lun number.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"portals": schema.ListAttribute{
																			Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																			MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																			MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																			MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"target_portal": schema.StringAttribute{
																			Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																			MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
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
																	Description:         "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	MarkdownDescription: "name of the volume. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"nfs": schema.SingleNestedAttribute{
																	Description:         "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																	Attributes: map[string]schema.Attribute{
																		"path": schema.StringAttribute{
																			Description:         "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			MarkdownDescription: "path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"server": schema.StringAttribute{
																			Description:         "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			MarkdownDescription: "server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"persistent_volume_claim": schema.SingleNestedAttribute{
																	Description:         "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																	MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																	Attributes: map[string]schema.Attribute{
																		"claim_name": schema.StringAttribute{
																			Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																			MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
																			MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts. Default false.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"photon_persistent_disk": schema.SingleNestedAttribute{
																	Description:         "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
																	MarkdownDescription: "photonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"pd_id": schema.StringAttribute{
																			Description:         "pdID is the ID that identifies Photon Controller persistent disk",
																			MarkdownDescription: "pdID is the ID that identifies Photon Controller persistent disk",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"portworx_volume": schema.SingleNestedAttribute{
																	Description:         "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
																	MarkdownDescription: "portworxVolume represents a portworx volume attached and mounted on kubelets host machine",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_id": schema.StringAttribute{
																			Description:         "volumeID uniquely identifies a Portworx volume",
																			MarkdownDescription: "volumeID uniquely identifies a Portworx volume",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"projected": schema.SingleNestedAttribute{
																	Description:         "projected items for all in one resources secrets, configmaps, and downward API",
																	MarkdownDescription: "projected items for all in one resources secrets, configmaps, and downward API",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "defaultMode are the mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"sources": schema.ListNestedAttribute{
																			Description:         "sources is the list of volume projections",
																			MarkdownDescription: "sources is the list of volume projections",
																			NestedObject: schema.NestedAttributeObject{
																				Attributes: map[string]schema.Attribute{
																					"config_map": schema.SingleNestedAttribute{
																						Description:         "configMap information about the configMap data to project",
																						MarkdownDescription: "configMap information about the configMap data to project",
																						Attributes: map[string]schema.Attribute{
																							"items": schema.ListNestedAttribute{
																								Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																								MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																											Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
																											Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																											MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																											Description:         "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
																											MarkdownDescription: "Required: Selects a field of the pod: only annotations, labels, name and namespace are supported.",
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
																											Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
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
																											Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																											MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
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
																								Description:         "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																								MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																											Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																											Required:            false,
																											Optional:            true,
																											Computed:            false,
																										},

																										"path": schema.StringAttribute{
																											Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																											MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																								Description:         "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																								MarkdownDescription: "audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"expiration_seconds": schema.Int64Attribute{
																								Description:         "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																								MarkdownDescription: "expirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																								Required:            false,
																								Optional:            true,
																								Computed:            false,
																							},

																							"path": schema.StringAttribute{
																								Description:         "path is the path relative to the mount point of the file to project the token into.",
																								MarkdownDescription: "path is the path relative to the mount point of the file to project the token into.",
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

																"quobyte": schema.SingleNestedAttribute{
																	Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
																	MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
																	Attributes: map[string]schema.Attribute{
																		"group": schema.StringAttribute{
																			Description:         "group to map volume access to Default is no group",
																			MarkdownDescription: "group to map volume access to Default is no group",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																			MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"registry": schema.StringAttribute{
																			Description:         "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																			MarkdownDescription: "registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"tenant": schema.StringAttribute{
																			Description:         "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																			MarkdownDescription: "tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"user": schema.StringAttribute{
																			Description:         "user to map volume access to Defaults to serivceaccount user",
																			MarkdownDescription: "user to map volume access to Defaults to serivceaccount user",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume": schema.StringAttribute{
																			Description:         "volume is a string that references an already created Quobyte volume by name.",
																			MarkdownDescription: "volume is a string that references an already created Quobyte volume by name.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"rbd": schema.SingleNestedAttribute{
																	Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
																	MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																			MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"image": schema.StringAttribute{
																			Description:         "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "image is the rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"keyring": schema.StringAttribute{
																			Description:         "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"monitors": schema.ListAttribute{
																			Description:         "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "monitors is a collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			ElementType:         types.StringType,
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"pool": schema.StringAttribute{
																			Description:         "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "pool is the rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"user": schema.StringAttribute{
																			Description:         "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			MarkdownDescription: "user is the rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"scale_io": schema.SingleNestedAttribute{
																	Description:         "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
																	MarkdownDescription: "scaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"gateway": schema.StringAttribute{
																			Description:         "gateway is the host address of the ScaleIO API Gateway.",
																			MarkdownDescription: "gateway is the host address of the ScaleIO API Gateway.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"protection_domain": schema.StringAttribute{
																			Description:         "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																			MarkdownDescription: "protectionDomain is the name of the ScaleIO Protection Domain for the configured storage.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																			MarkdownDescription: "secretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: true,
																			Optional: false,
																			Computed: false,
																		},

																		"ssl_enabled": schema.BoolAttribute{
																			Description:         "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																			MarkdownDescription: "sslEnabled Flag enable/disable SSL communication with Gateway, default false",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_mode": schema.StringAttribute{
																			Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																			MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_pool": schema.StringAttribute{
																			Description:         "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																			MarkdownDescription: "storagePool is the ScaleIO Storage Pool associated with the protection domain.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"system": schema.StringAttribute{
																			Description:         "system is the name of the storage system as configured in ScaleIO.",
																			MarkdownDescription: "system is the name of the storage system as configured in ScaleIO.",
																			Required:            true,
																			Optional:            false,
																			Computed:            false,
																		},

																		"volume_name": schema.StringAttribute{
																			Description:         "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
																			MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO system that is associated with this volume source.",
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
																	Description:         "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																	MarkdownDescription: "secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																	Attributes: map[string]schema.Attribute{
																		"default_mode": schema.Int64Attribute{
																			Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"items": schema.ListNestedAttribute{
																			Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																			MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																						Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"path": schema.StringAttribute{
																						Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																						MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																		"optional": schema.BoolAttribute{
																			Description:         "optional field specify whether the Secret or its keys must be defined",
																			MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_name": schema.StringAttribute{
																			Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																			MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"storageos": schema.SingleNestedAttribute{
																	Description:         "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
																	MarkdownDescription: "storageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"read_only": schema.BoolAttribute{
																			Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"secret_ref": schema.SingleNestedAttribute{
																			Description:         "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
																			MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},
																			},
																			Required: false,
																			Optional: true,
																			Computed: false,
																		},

																		"volume_name": schema.StringAttribute{
																			Description:         "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
																			MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_namespace": schema.StringAttribute{
																			Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																			MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},
																	},
																	Required: false,
																	Optional: true,
																	Computed: false,
																},

																"vsphere_volume": schema.SingleNestedAttribute{
																	Description:         "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
																	MarkdownDescription: "vsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
																	Attributes: map[string]schema.Attribute{
																		"fs_type": schema.StringAttribute{
																			Description:         "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			MarkdownDescription: "fsType is filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_policy_id": schema.StringAttribute{
																			Description:         "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																			MarkdownDescription: "storagePolicyID is the storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"storage_policy_name": schema.StringAttribute{
																			Description:         "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																			MarkdownDescription: "storagePolicyName is the storage Policy Based Management (SPBM) profile name.",
																			Required:            false,
																			Optional:            true,
																			Computed:            false,
																		},

																		"volume_path": schema.StringAttribute{
																			Description:         "volumePath is the path that identifies vSphere volume vmdk",
																			MarkdownDescription: "volumePath is the path that identifies vSphere volume vmdk",
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
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"issuer": schema.SingleNestedAttribute{
											Description:         "Specifies the configuration for the TLS certificates issuer. It allows defining the issuer name and the reference to the secret containing the TLS certificates and key. The secret should contain the CA certificate, TLS certificate, and private key in the specified keys. Required when TLS is enabled.",
											MarkdownDescription: "Specifies the configuration for the TLS certificates issuer. It allows defining the issuer name and the reference to the secret containing the TLS certificates and key. The secret should contain the CA certificate, TLS certificate, and private key in the specified keys. Required when TLS is enabled.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "The issuer for TLS certificates. It only allows two enum values: 'KubeBlocks' and 'UserProvided'.  - 'KubeBlocks' indicates that the self-signed TLS certificates generated by the KubeBlocks Operator will be used. - 'UserProvided' means that the user is responsible for providing their own CA, Cert, and Key. In this case, the user-provided CA certificate, server certificate, and private key will be used for TLS communication.",
													MarkdownDescription: "The issuer for TLS certificates. It only allows two enum values: 'KubeBlocks' and 'UserProvided'.  - 'KubeBlocks' indicates that the self-signed TLS certificates generated by the KubeBlocks Operator will be used. - 'UserProvided' means that the user is responsible for providing their own CA, Cert, and Key. In this case, the user-provided CA certificate, server certificate, and private key will be used for TLS communication.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "SecretRef is the reference to the secret that contains user-provided certificates. It is required when the issuer is set to 'UserProvided'.",
													MarkdownDescription: "SecretRef is the reference to the secret that contains user-provided certificates. It is required when the issuer is set to 'UserProvided'.",
													Attributes: map[string]schema.Attribute{
														"ca": schema.StringAttribute{
															Description:         "Key of CA cert in Secret",
															MarkdownDescription: "Key of CA cert in Secret",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"cert": schema.StringAttribute{
															Description:         "Key of Cert in Secret",
															MarkdownDescription: "Key of Cert in Secret",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Key of TLS private key in Secret",
															MarkdownDescription: "Key of TLS private key in Secret",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "Name of the Secret that contains user-provided certificates.",
															MarkdownDescription: "Name of the Secret that contains user-provided certificates.",
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

										"monitor": schema.BoolAttribute{
											Description:         "Deprecated since v0.9 Determines whether metrics exporter information is annotated on the Component's headless Service.  If set to true, the following annotations will be patched into the Service:  - 'monitor.kubeblocks.io/path' - 'monitor.kubeblocks.io/port' - 'monitor.kubeblocks.io/scheme'  These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
											MarkdownDescription: "Deprecated since v0.9 Determines whether metrics exporter information is annotated on the Component's headless Service.  If set to true, the following annotations will be patched into the Service:  - 'monitor.kubeblocks.io/path' - 'monitor.kubeblocks.io/port' - 'monitor.kubeblocks.io/scheme'  These annotations allow the Prometheus installed by KubeBlocks to discover and scrape metrics from the exporter.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"name": schema.StringAttribute{
											Description:         "Specifies the Component's name. It's part of the Service DNS name and must comply with the IANA service naming rule. The name is optional when ClusterComponentSpec is used as a template (e.g., in 'shardingSpec'), but required otherwise.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='name is immutable'",
											MarkdownDescription: "Specifies the Component's name. It's part of the Service DNS name and must comply with the IANA service naming rule. The name is optional when ClusterComponentSpec is used as a template (e.g., in 'shardingSpec'), but required otherwise.  TODO +kubebuilder:validation:XValidation:rule='self == oldSelf',message='name is immutable'",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(22),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
											},
										},

										"offline_instances": schema.ListAttribute{
											Description:         "Specifies the names of instances to be transitioned to offline status.  Marking an instance as offline results in the following:  1. The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential future reuse or data recovery, but it is no longer actively used. 2. The ordinal number assigned to this instance is preserved, ensuring it remains unique and avoiding conflicts with new instances.  Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining ordinal consistency within the Cluster. Note that offline instances and their associated resources, such as PVCs, are not automatically deleted. The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.",
											MarkdownDescription: "Specifies the names of instances to be transitioned to offline status.  Marking an instance as offline results in the following:  1. The associated Pod is stopped, and its PersistentVolumeClaim (PVC) is retained for potential future reuse or data recovery, but it is no longer actively used. 2. The ordinal number assigned to this instance is preserved, ensuring it remains unique and avoiding conflicts with new instances.  Setting instances to offline allows for a controlled scale-in process, preserving their data and maintaining ordinal consistency within the Cluster. Note that offline instances and their associated resources, such as PVCs, are not automatically deleted. The administrator must manually manage the cleanup and removal of these resources when they are no longer needed.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"replicas": schema.Int64Attribute{
											Description:         "Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.",
											MarkdownDescription: "Specifies the desired number of replicas in the Component for enhancing availability and durability, or load balancing.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.Int64{
												int64validator.AtLeast(0),
											},
										},

										"resources": schema.SingleNestedAttribute{
											Description:         "Specifies the resources required by the Component. It allows defining the CPU, memory requirements and limits for the Component's containers.",
											MarkdownDescription: "Specifies the resources required by the Component. It allows defining the CPU, memory requirements and limits for the Component's containers.",
											Attributes: map[string]schema.Attribute{
												"claims": schema.ListNestedAttribute{
													Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
													MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
													Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"requests": schema.MapAttribute{
													Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
													MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

										"scheduling_policy": schema.SingleNestedAttribute{
											Description:         "Specifies the scheduling policy for the Component.",
											MarkdownDescription: "Specifies the scheduling policy for the Component.",
											Attributes: map[string]schema.Attribute{
												"affinity": schema.SingleNestedAttribute{
													Description:         "Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.",
													MarkdownDescription: "Specifies a group of affinity scheduling rules of the Cluster, including NodeAffinity, PodAffinity, and PodAntiAffinity.",
													Attributes: map[string]schema.Attribute{
														"node_affinity": schema.SingleNestedAttribute{
															Description:         "Describes node affinity scheduling rules for the pod.",
															MarkdownDescription: "Describes node affinity scheduling rules for the pod.",
															Attributes: map[string]schema.Attribute{
																"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																	Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
																	MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node matches the corresponding matchExpressions; the node(s) with the highest sum are the most preferred.",
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
																									Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																									Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																	Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
																	MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to an update), the system may or may not try to eventually evict the pod from its node.",
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
																									Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
																									Description:         "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									MarkdownDescription: "Represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist. Gt, and Lt.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "An array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. If the operator is Gt or Lt, the values array must have a single element, which will be interpreted as an integer. This array is replaced during a strategic merge patch.",
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
															Description:         "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
															MarkdownDescription: "Describes pod affinity scheduling rules (e.g. co-locate this pod in the same node, zone, etc. as some other pod(s)).",
															Attributes: map[string]schema.Attribute{
																"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																	Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																	MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"pod_affinity_term": schema.SingleNestedAttribute{
																				Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																				MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																				Attributes: map[string]schema.Attribute{
																					"label_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over a set of resources, in this case pods.",
																						MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																						Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																						MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																						Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"topology_key": schema.StringAttribute{
																						Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																						MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
																				Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																				MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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
																	Description:         "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	MarkdownDescription: "If the affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																									Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																				Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																									Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																				Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
															Description:         "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
															MarkdownDescription: "Describes pod anti-affinity scheduling rules (e.g. avoid putting this pod in the same node, zone, etc. as some other pod(s)).",
															Attributes: map[string]schema.Attribute{
																"preferred_during_scheduling_ignored_during_execution": schema.ListNestedAttribute{
																	Description:         "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																	MarkdownDescription: "The scheduler will prefer to schedule pods to nodes that satisfy the anti-affinity expressions specified by this field, but it may choose a node that violates one or more of the expressions. The node that is most preferred is the one with the greatest sum of weights, i.e. for each node that meets all of the scheduling requirements (resource request, requiredDuringScheduling anti-affinity expressions, etc.), compute a sum by iterating through the elements of this field and adding 'weight' to the sum if the node has pods which matches the corresponding podAffinityTerm; the node(s) with the highest sum are the most preferred.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"pod_affinity_term": schema.SingleNestedAttribute{
																				Description:         "Required. A pod affinity term, associated with the corresponding weight.",
																				MarkdownDescription: "Required. A pod affinity term, associated with the corresponding weight.",
																				Attributes: map[string]schema.Attribute{
																					"label_selector": schema.SingleNestedAttribute{
																						Description:         "A label query over a set of resources, in this case pods.",
																						MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																						Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																						MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																											Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																											MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																											Required:            true,
																											Optional:            false,
																											Computed:            false,
																										},

																										"values": schema.ListAttribute{
																											Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																											MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																								Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																								MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																						Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																						ElementType:         types.StringType,
																						Required:            false,
																						Optional:            true,
																						Computed:            false,
																					},

																					"topology_key": schema.StringAttribute{
																						Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																						MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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
																				Description:         "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
																				MarkdownDescription: "weight associated with matching the corresponding podAffinityTerm, in the range 1-100.",
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
																	Description:         "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	MarkdownDescription: "If the anti-affinity requirements specified by this field are not met at scheduling time, the pod will not be scheduled onto the node. If the anti-affinity requirements specified by this field cease to be met at some point during pod execution (e.g. due to a pod label update), the system may or may not try to eventually evict the pod from its node. When there are multiple elements, the lists of nodes corresponding to each podAffinityTerm are intersected, i.e. all terms must be satisfied.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"label_selector": schema.SingleNestedAttribute{
																				Description:         "A label query over a set of resources, in this case pods.",
																				MarkdownDescription: "A label query over a set of resources, in this case pods.",
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
																									Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																				Description:         "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
																				MarkdownDescription: "A label query over the set of namespaces that the term applies to. The term is applied to the union of the namespaces selected by this field and the ones listed in the namespaces field. null selector and null or empty namespaces list means 'this pod's namespace'. An empty selector ({}) matches all namespaces.",
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
																									Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																									Required:            true,
																									Optional:            false,
																									Computed:            false,
																								},

																								"values": schema.ListAttribute{
																									Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																									MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																						Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																						MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																				Description:         "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				MarkdownDescription: "namespaces specifies a static list of namespace names that the term applies to. The term is applied to the union of the namespaces listed in this field and the ones selected by namespaceSelector. null or empty namespaces list and null namespaceSelector means 'this pod's namespace'.",
																				ElementType:         types.StringType,
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"topology_key": schema.StringAttribute{
																				Description:         "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
																				MarkdownDescription: "This pod should be co-located (affinity) or not co-located (anti-affinity) with the pods matching the labelSelector in the specified namespaces, where co-located is defined as running on a node whose value of the label with key topologyKey matches that of any node on which any of the selected pods is running. Empty topologyKey is not allowed.",
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

												"node_name": schema.StringAttribute{
													Description:         "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty, the scheduler simply schedules this Pod onto that node, assuming that it fits resource requirements.",
													MarkdownDescription: "NodeName is a request to schedule this Pod onto a specific node. If it is non-empty, the scheduler simply schedules this Pod onto that node, assuming that it fits resource requirements.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selector": schema.MapAttribute{
													Description:         "NodeSelector is a selector which must be true for the Pod to fit on a node. Selector which must match a node's labels for the Pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
													MarkdownDescription: "NodeSelector is a selector which must be true for the Pod to fit on a node. Selector which must match a node's labels for the Pod to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"scheduler_name": schema.StringAttribute{
													Description:         "If specified, the Pod will be dispatched by specified scheduler. If not specified, the Pod will be dispatched by default scheduler.",
													MarkdownDescription: "If specified, the Pod will be dispatched by specified scheduler. If not specified, the Pod will be dispatched by default scheduler.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tolerations": schema.ListNestedAttribute{
													Description:         "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
													MarkdownDescription: "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"effect": schema.StringAttribute{
																Description:         "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																MarkdownDescription: "Effect indicates the taint effect to match. Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key": schema.StringAttribute{
																Description:         "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																MarkdownDescription: "Key is the taint key that the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists; this combination means to match all values and all keys.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																MarkdownDescription: "Operator represents a key's relationship to the value. Valid operators are Exists and Equal. Defaults to Equal. Exists is equivalent to wildcard for value, so that a pod can tolerate all taints of a particular category.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"toleration_seconds": schema.Int64Attribute{
																Description:         "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																MarkdownDescription: "TolerationSeconds represents the period of time the toleration (which must be of effect NoExecute, otherwise this field is ignored) tolerates the taint. By default, it is not set, which means tolerate the taint forever (do not evict). Zero and negative values will be treated as 0 (evict immediately) by the system.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
																MarkdownDescription: "Value is the taint value the toleration matches to. If the operator is Exists, the value should be empty, otherwise just a regular string.",
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
													Description:         "TopologySpreadConstraints describes how a group of Pods ought to spread across topology domains. Scheduler will schedule Pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
													MarkdownDescription: "TopologySpreadConstraints describes how a group of Pods ought to spread across topology domains. Scheduler will schedule Pods in a way which abides by the constraints. All topologySpreadConstraints are ANDed.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"label_selector": schema.SingleNestedAttribute{
																Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
																MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
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
																					Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"values": schema.ListAttribute{
																					Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																					MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
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
																		Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																		MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
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
																Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
																MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. The same key is forbidden to exist in both MatchLabelKeys and LabelSelector. MatchLabelKeys cannot be set when LabelSelector isn't set. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.  This is a beta field and requires the MatchLabelKeysInPodTopologySpread feature gate to be enabled (enabled by default).",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"max_skew": schema.Int64Attribute{
																Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
																MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"min_domains": schema.Int64Attribute{
																Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
																MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"node_affinity_policy": schema.StringAttribute{
																Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"node_taints_policy": schema.StringAttribute{
																Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a beta-level feature default enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"topology_key": schema.StringAttribute{
																Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
																MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"when_unsatisfiable": schema.StringAttribute{
																Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
																MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location, but giving higher precedence to topologies that would help reduce the skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
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

										"service_account_name": schema.StringAttribute{
											Description:         "Specifies the name of the ServiceAccount required by the running Component. This ServiceAccount is used to grant necessary permissions for the Component's Pods to interact with other Kubernetes resources, such as modifying Pod labels or sending events.  Defaults: If not specified, KubeBlocks automatically assigns a default ServiceAccount named 'kb-{cluster.name}', bound to a default role installed together with KubeBlocks.  Future Changes: Future versions might change the default ServiceAccount creation strategy to one per Component, potentially revising the naming to 'kb-{cluster.name}-{component.name}'.  Users can override the automatic ServiceAccount assignment by explicitly setting the name of an existed ServiceAccount in this field.",
											MarkdownDescription: "Specifies the name of the ServiceAccount required by the running Component. This ServiceAccount is used to grant necessary permissions for the Component's Pods to interact with other Kubernetes resources, such as modifying Pod labels or sending events.  Defaults: If not specified, KubeBlocks automatically assigns a default ServiceAccount named 'kb-{cluster.name}', bound to a default role installed together with KubeBlocks.  Future Changes: Future versions might change the default ServiceAccount creation strategy to one per Component, potentially revising the naming to 'kb-{cluster.name}-{component.name}'.  Users can override the automatic ServiceAccount assignment by explicitly setting the name of an existed ServiceAccount in this field.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"service_refs": schema.ListNestedAttribute{
											Description:         "Defines a list of ServiceRef for a Component, enabling access to both external services and Services provided by other Clusters.  Types of services:  - External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator; Require a ServiceDescriptor for connection details. - Services provided by a Cluster: Managed by the same KubeBlocks operator; identified using Cluster, Component and Service names.  ServiceRefs with identical 'serviceRef.name' in the same Cluster are considered the same.  Example: '''yaml serviceRefs: - name: 'redis-sentinel' serviceDescriptor: name: 'external-redis-sentinel' - name: 'postgres-cluster' clusterServiceSelector: cluster: 'my-postgres-cluster' service: component: 'postgresql' ''' The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.",
											MarkdownDescription: "Defines a list of ServiceRef for a Component, enabling access to both external services and Services provided by other Clusters.  Types of services:  - External services: Not managed by KubeBlocks or managed by a different KubeBlocks operator; Require a ServiceDescriptor for connection details. - Services provided by a Cluster: Managed by the same KubeBlocks operator; identified using Cluster, Component and Service names.  ServiceRefs with identical 'serviceRef.name' in the same Cluster are considered the same.  Example: '''yaml serviceRefs: - name: 'redis-sentinel' serviceDescriptor: name: 'external-redis-sentinel' - name: 'postgres-cluster' clusterServiceSelector: cluster: 'my-postgres-cluster' service: component: 'postgresql' ''' The example above includes ServiceRefs to an external Redis Sentinel service and a PostgreSQL Cluster.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"cluster": schema.StringAttribute{
														Description:         "Specifies the name of the KubeBlocks Cluster being referenced. This is used when services from another KubeBlocks Cluster are consumed.  By default, the referenced KubeBlocks Cluster's 'clusterDefinition.spec.connectionCredential' will be utilized to bind to the current Component. This credential should include: 'endpoint', 'port', 'username', and 'password'.  Note:  - The 'ServiceKind' and 'ServiceVersion' specified in the service reference within the ClusterDefinition are not validated when using this approach. - If both 'cluster' and 'serviceDescriptor' are present, 'cluster' will take precedence.  Deprecated since v0.9 since 'clusterDefinition.spec.connectionCredential' is deprecated, use 'clusterServiceSelector' instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
														MarkdownDescription: "Specifies the name of the KubeBlocks Cluster being referenced. This is used when services from another KubeBlocks Cluster are consumed.  By default, the referenced KubeBlocks Cluster's 'clusterDefinition.spec.connectionCredential' will be utilized to bind to the current Component. This credential should include: 'endpoint', 'port', 'username', and 'password'.  Note:  - The 'ServiceKind' and 'ServiceVersion' specified in the service reference within the ClusterDefinition are not validated when using this approach. - If both 'cluster' and 'serviceDescriptor' are present, 'cluster' will take precedence.  Deprecated since v0.9 since 'clusterDefinition.spec.connectionCredential' is deprecated, use 'clusterServiceSelector' instead. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cluster_service_selector": schema.SingleNestedAttribute{
														Description:         "References a service provided by another KubeBlocks Cluster. It specifies the ClusterService and the account credentials needed for access.",
														MarkdownDescription: "References a service provided by another KubeBlocks Cluster. It specifies the ClusterService and the account credentials needed for access.",
														Attributes: map[string]schema.Attribute{
															"cluster": schema.StringAttribute{
																Description:         "The name of the Cluster being referenced.",
																MarkdownDescription: "The name of the Cluster being referenced.",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"credential": schema.SingleNestedAttribute{
																Description:         "Specifies the SystemAccount to authenticate and establish a connection with the referenced Cluster. The SystemAccount should be defined in 'componentDefinition.spec.systemAccounts' of the Component providing the service in the referenced Cluster.",
																MarkdownDescription: "Specifies the SystemAccount to authenticate and establish a connection with the referenced Cluster. The SystemAccount should be defined in 'componentDefinition.spec.systemAccounts' of the Component providing the service in the referenced Cluster.",
																Attributes: map[string]schema.Attribute{
																	"component": schema.StringAttribute{
																		Description:         "The name of the Component where the credential resides in.",
																		MarkdownDescription: "The name of the Component where the credential resides in.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"name": schema.StringAttribute{
																		Description:         "The name of the credential (SystemAccount) to reference.",
																		MarkdownDescription: "The name of the credential (SystemAccount) to reference.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"service": schema.SingleNestedAttribute{
																Description:         "Identifies a ClusterService from the list of Services defined in 'cluster.spec.services' of the referenced Cluster.",
																MarkdownDescription: "Identifies a ClusterService from the list of Services defined in 'cluster.spec.services' of the referenced Cluster.",
																Attributes: map[string]schema.Attribute{
																	"component": schema.StringAttribute{
																		Description:         "The name of the Component where the Service resides in.  It is required when referencing a Component's Service.",
																		MarkdownDescription: "The name of the Component where the Service resides in.  It is required when referencing a Component's Service.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"port": schema.StringAttribute{
																		Description:         "The port name of the Service to be referenced.  If there is a non-zero node-port exist for the matched Service port, the node-port will be selected first.  If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched, and the resolved value will be presented in the following format: service1.name:port1,service2.name:port2...",
																		MarkdownDescription: "The port name of the Service to be referenced.  If there is a non-zero node-port exist for the matched Service port, the node-port will be selected first.  If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched, and the resolved value will be presented in the following format: service1.name:port1,service2.name:port2...",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"service": schema.StringAttribute{
																		Description:         "The name of the Service to be referenced.  Leave it empty to reference the default Service. Set it to 'headless' to reference the default headless Service.  If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched, and the resolved value will be presented in the following format: service1.name,service2.name...",
																		MarkdownDescription: "The name of the Service to be referenced.  Leave it empty to reference the default Service. Set it to 'headless' to reference the default headless Service.  If the referenced Service is of pod-service type (a Service per Pod), there will be multiple Service objects matched, and the resolved value will be presented in the following format: service1.name,service2.name...",
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

													"name": schema.StringAttribute{
														Description:         "Specifies the identifier of the service reference declaration. It corresponds to the serviceRefDeclaration name defined in either:  - 'componentDefinition.spec.serviceRefDeclarations[*].name' - 'clusterDefinition.spec.componentDefs[*].serviceRefDeclarations[*].name' (deprecated)",
														MarkdownDescription: "Specifies the identifier of the service reference declaration. It corresponds to the serviceRefDeclaration name defined in either:  - 'componentDefinition.spec.serviceRefDeclarations[*].name' - 'clusterDefinition.spec.componentDefs[*].serviceRefDeclarations[*].name' (deprecated)",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object. If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current Cluster by default.",
														MarkdownDescription: "Specifies the namespace of the referenced Cluster or the namespace of the referenced ServiceDescriptor object. If not provided, the referenced Cluster and ServiceDescriptor will be searched in the namespace of the current Cluster by default.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_descriptor": schema.StringAttribute{
														Description:         "Specifies the name of the ServiceDescriptor object that describes a service provided by external sources.  When referencing a service provided by external sources, a ServiceDescriptor object is required to establish the service binding. The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKind and serviceVersion declared in the definition.  If both 'cluster' and 'serviceDescriptor' are specified, the 'cluster' takes precedence.",
														MarkdownDescription: "Specifies the name of the ServiceDescriptor object that describes a service provided by external sources.  When referencing a service provided by external sources, a ServiceDescriptor object is required to establish the service binding. The 'serviceDescriptor.spec.serviceKind' and 'serviceDescriptor.spec.serviceVersion' should match the serviceKind and serviceVersion declared in the definition.  If both 'cluster' and 'serviceDescriptor' are specified, the 'cluster' takes precedence.",
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

										"service_version": schema.StringAttribute{
											Description:         "ServiceVersion specifies the version of the Service expected to be provisioned by this Component. The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/). If no version is specified, the latest available version will be used.",
											MarkdownDescription: "ServiceVersion specifies the version of the Service expected to be provisioned by this Component. The version should follow the syntax and semantics of the 'Semantic Versioning' specification (http://semver.org/). If no version is specified, the latest available version will be used.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(32),
											},
										},

										"services": schema.ListNestedAttribute{
											Description:         "Overrides services defined in referenced ComponentDefinition and expose endpoints that can be accessed by clients.",
											MarkdownDescription: "Overrides services defined in referenced ComponentDefinition and expose endpoints that can be accessed by clients.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"annotations": schema.MapAttribute{
														Description:         "If ServiceType is LoadBalancer, cloud provider related parameters can be put here. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
														MarkdownDescription: "If ServiceType is LoadBalancer, cloud provider related parameters can be put here. More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "References the ComponentService name defined in the 'componentDefinition.spec.services[*].name'.",
														MarkdownDescription: "References the ComponentService name defined in the 'componentDefinition.spec.services[*].name'.",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtMost(25),
														},
													},

													"pod_service": schema.BoolAttribute{
														Description:         "Indicates whether to generate individual Services for each Pod. If set to true, a separate Service will be created for each Pod in the Cluster.",
														MarkdownDescription: "Indicates whether to generate individual Services for each Pod. If set to true, a separate Service will be created for each Pod in the Cluster.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"service_type": schema.StringAttribute{
														Description:         "Determines how the Service is exposed. Valid options are 'ClusterIP', 'NodePort', and 'LoadBalancer'.  - 'ClusterIP' allocates a Cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the ClusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the ClusterIP.  Note: although K8s Service type allows the 'ExternalName' type, it is not a valid option for ClusterComponentService.  For more info, see: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
														MarkdownDescription: "Determines how the Service is exposed. Valid options are 'ClusterIP', 'NodePort', and 'LoadBalancer'.  - 'ClusterIP' allocates a Cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. - 'NodePort' builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the ClusterIP. - 'LoadBalancer' builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the ClusterIP.  Note: although K8s Service type allows the 'ExternalName' type, it is not a valid option for ClusterComponentService.  For more info, see: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.OneOf("ClusterIP", "NodePort", "LoadBalancer"),
														},
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"switch_policy": schema.SingleNestedAttribute{
											Description:         "Defines the strategy for switchover and failover when workloadType is Replication.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
											MarkdownDescription: "Defines the strategy for switchover and failover when workloadType is Replication.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
											Attributes: map[string]schema.Attribute{
												"type": schema.StringAttribute{
													Description:         "Type specifies the type of switch policy to be applied.",
													MarkdownDescription: "Type specifies the type of switch policy to be applied.",
													Required:            false,
													Optional:            true,
													Computed:            false,
													Validators: []validator.String{
														stringvalidator.OneOf("Noop"),
													},
												},
											},
											Required: false,
											Optional: true,
											Computed: false,
										},

										"tls": schema.BoolAttribute{
											Description:         "A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication. When set to true, the Component will be configured to use TLS encryption for its network connections. This ensures that the data transmitted between the Component and its clients or other Components is encrypted and protected from unauthorized access. If TLS is enabled, the Component may require additional configuration, such as specifying TLS certificates and keys, to properly set up the secure communication channel.",
											MarkdownDescription: "A boolean flag that indicates whether the Component should use Transport Layer Security (TLS) for secure communication. When set to true, the Component will be configured to use TLS encryption for its network connections. This ensures that the data transmitted between the Component and its clients or other Components is encrypted and protected from unauthorized access. If TLS is enabled, the Component may require additional configuration, such as specifying TLS certificates and keys, to properly set up the secure communication channel.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"tolerations": schema.MapAttribute{
											Description:         "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.  Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
											MarkdownDescription: "Allows Pods to be scheduled onto nodes with matching taints. Each toleration in the array allows the Pod to tolerate node taints based on specified 'key', 'value', 'effect', and 'operator'.  - The 'key', 'value', and 'effect' identify the taint that the toleration matches. - The 'operator' determines how the toleration matches the taint.  Pods with matching tolerations are allowed to be scheduled on tainted nodes, typically reserved for specific purposes.  Deprecated since v0.10, replaced by the 'schedulingPolicy' field.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"update_strategy": schema.StringAttribute{
											Description:         "Defines the update strategy for the Component.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
											MarkdownDescription: "Defines the update strategy for the Component.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("Serial", "BestEffortParallel", "Parallel"),
											},
										},

										"user_resource_refs": schema.SingleNestedAttribute{
											Description:         "Allows users to specify custom ConfigMaps and Secrets to be mounted as volumes in the Cluster's Pods. This is useful in scenarios where users need to provide additional resources to the Cluster, such as:  - Mounting custom scripts or configuration files during Cluster startup. - Mounting Secrets as volumes to provide sensitive information, like S3 AK/SK, to the Cluster.",
											MarkdownDescription: "Allows users to specify custom ConfigMaps and Secrets to be mounted as volumes in the Cluster's Pods. This is useful in scenarios where users need to provide additional resources to the Cluster, such as:  - Mounting custom scripts or configuration files during Cluster startup. - Mounting Secrets as volumes to provide sensitive information, like S3 AK/SK, to the Cluster.",
											Attributes: map[string]schema.Attribute{
												"config_map_refs": schema.ListNestedAttribute{
													Description:         "ConfigMapRefs defines the user-defined ConfigMaps.",
													MarkdownDescription: "ConfigMapRefs defines the user-defined ConfigMaps.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"as_volume_from": schema.ListAttribute{
																Description:         "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
																MarkdownDescription: "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"config_map": schema.SingleNestedAttribute{
																Description:         "ConfigMap specifies the ConfigMap to be mounted as a volume.",
																MarkdownDescription: "ConfigMap specifies the ConfigMap to be mounted as a volume.",
																Attributes: map[string]schema.Attribute{
																	"default_mode": schema.Int64Attribute{
																		Description:         "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"items": schema.ListNestedAttribute{
																		Description:         "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items if unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																					Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
																					Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																					MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																		Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																		MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																Required: true,
																Optional: false,
																Computed: false,
															},

															"mount_point": schema.StringAttribute{
																Description:         "MountPoint is the filesystem path where the volume will be mounted.",
																MarkdownDescription: "MountPoint is the filesystem path where the volume will be mounted.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(256),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^/[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
																},
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
																MarkdownDescription: "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
																},
															},

															"sub_path": schema.StringAttribute{
																Description:         "SubPath specifies a path within the volume from which to mount.",
																MarkdownDescription: "SubPath specifies a path within the volume from which to mount.",
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

												"secret_refs": schema.ListNestedAttribute{
													Description:         "SecretRefs defines the user-defined Secrets.",
													MarkdownDescription: "SecretRefs defines the user-defined Secrets.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"as_volume_from": schema.ListAttribute{
																Description:         "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
																MarkdownDescription: "AsVolumeFrom lists the names of containers in which the volume should be mounted.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"mount_point": schema.StringAttribute{
																Description:         "MountPoint is the filesystem path where the volume will be mounted.",
																MarkdownDescription: "MountPoint is the filesystem path where the volume will be mounted.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(256),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^/[a-z]([a-z0-9\-]*[a-z0-9])?$`), ""),
																},
															},

															"name": schema.StringAttribute{
																Description:         "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
																MarkdownDescription: "Name is the name of the referenced ConfigMap or Secret object. It must conform to DNS label standards.",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtMost(63),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
																},
															},

															"secret": schema.SingleNestedAttribute{
																Description:         "Secret specifies the Secret to be mounted as a volume.",
																MarkdownDescription: "Secret specifies the Secret to be mounted as a volume.",
																Attributes: map[string]schema.Attribute{
																	"default_mode": schema.Int64Attribute{
																		Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"items": schema.ListNestedAttribute{
																		Description:         "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																		MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
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
																					Description:         "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					MarkdownDescription: "mode is Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																					Required:            false,
																					Optional:            true,
																					Computed:            false,
																				},

																				"path": schema.StringAttribute{
																					Description:         "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																					MarkdownDescription: "path is the relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																	"optional": schema.BoolAttribute{
																		Description:         "optional field specify whether the Secret or its keys must be defined",
																		MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"secret_name": schema.StringAttribute{
																		Description:         "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																		MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"sub_path": schema.StringAttribute{
																Description:         "SubPath specifies a path within the volume from which to mount.",
																MarkdownDescription: "SubPath specifies a path within the volume from which to mount.",
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

										"volume_claim_templates": schema.ListNestedAttribute{
											Description:         "Specifies a list of PersistentVolumeClaim templates that represent the storage requirements for the Component. Each template specifies the desired characteristics of a persistent volume, such as storage class, size, and access modes. These templates are used to dynamically provision persistent volumes for the Component.",
											MarkdownDescription: "Specifies a list of PersistentVolumeClaim templates that represent the storage requirements for the Component. Each template specifies the desired characteristics of a persistent volume, such as storage class, size, and access modes. These templates are used to dynamically provision persistent volumes for the Component.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Refers to the name of a volumeMount defined in either:  - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)  The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
														MarkdownDescription: "Refers to the name of a volumeMount defined in either:  - 'componentDefinition.spec.runtime.containers[*].volumeMounts' - 'clusterDefinition.spec.componentDefs[*].podSpec.containers[*].volumeMounts' (deprecated)  The value of 'name' must match the 'name' field of a volumeMount specified in the corresponding 'volumeMounts' array.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"spec": schema.SingleNestedAttribute{
														Description:         "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.  When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
														MarkdownDescription: "Defines the desired characteristics of a PersistentVolumeClaim that will be created for the volume with the mount name specified in the 'name' field.  When a Pod is created for this ClusterComponent, a new PVC will be created based on the specification defined in the 'spec' field. The PVC will be associated with the volume mount specified by the 'name' field.",
														Attributes: map[string]schema.Attribute{
															"access_modes": schema.MapAttribute{
																Description:         "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																MarkdownDescription: "Contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1.",
																ElementType:         types.StringType,
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"resources": schema.SingleNestedAttribute{
																Description:         "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
																MarkdownDescription: "Represents the minimum resources the volume should have. If the RecoverVolumeExpansionFailure feature is enabled, users are allowed to specify resource requirements that are lower than the previous value but must still be higher than the capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources.",
																Attributes: map[string]schema.Attribute{
																	"claims": schema.ListNestedAttribute{
																		Description:         "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																		MarkdownDescription: "Claims lists the names of resources, defined in spec.resourceClaims, that are used by this container.  This is an alpha field and requires enabling the DynamicResourceAllocation feature gate.  This field is immutable. It can only be set for containers.",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
																					MarkdownDescription: "Name must match the name of one entry in pod.spec.resourceClaims of the Pod where this field is used. It makes that resource available inside a container.",
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
																		Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																		MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"requests": schema.MapAttribute{
																		Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																		MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. Requests cannot exceed Limits. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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
																Description:         "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																MarkdownDescription: "The name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"volume_mode": schema.StringAttribute{
																Description:         "Defines what type of volume is required by the claim, either Block or Filesystem.",
																MarkdownDescription: "Defines what type of volume is required by the claim, either Block or Filesystem.",
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

					"storage": schema.SingleNestedAttribute{
						Description:         "Specifies the storage of the first componentSpec, if the storage of the first componentSpec is specified, this value will be ignored.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						MarkdownDescription: "Specifies the storage of the first componentSpec, if the storage of the first componentSpec is specified, this value will be ignored.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						Attributes: map[string]schema.Attribute{
							"size": schema.StringAttribute{
								Description:         "Specifies the amount of storage the Cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Specifies the amount of storage the Cluster needs. For more information, refer to: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"tenancy": schema.StringAttribute{
						Description:         "Describes how Pods are distributed across node.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						MarkdownDescription: "Describes how Pods are distributed across node.  Deprecated since v0.9. This field is maintained for backward compatibility and its use is discouraged. Existing usage should be updated to the current preferred approach to avoid compatibility issues in future releases.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("SharedNode", "DedicatedNode"),
						},
					},

					"termination_policy": schema.StringAttribute{
						Description:         "Specifies the behavior when a Cluster is deleted. It defines how resources, data, and backups associated with a Cluster are managed during termination. Choose a policy based on the desired level of resource cleanup and data preservation:  - 'DoNotTerminate': Prevents deletion of the Cluster. This policy ensures that all resources remain intact. - 'Halt': Deletes Cluster resources like Pods and Services but retains Persistent Volume Claims (PVCs), allowing for data preservation while stopping other operations. - 'Delete': Extends the 'Halt' policy by also removing PVCs, leading to a thorough cleanup while removing all persistent data. - 'WipeOut': An aggressive policy that deletes all Cluster resources, including volume snapshots and backups in external storage. This results in complete data removal and should be used cautiously, primarily in non-production environments to avoid irreversible data loss.  Warning: Choosing an inappropriate termination policy can result in data loss. The 'WipeOut' policy is particularly risky in production environments due to its irreversible nature.",
						MarkdownDescription: "Specifies the behavior when a Cluster is deleted. It defines how resources, data, and backups associated with a Cluster are managed during termination. Choose a policy based on the desired level of resource cleanup and data preservation:  - 'DoNotTerminate': Prevents deletion of the Cluster. This policy ensures that all resources remain intact. - 'Halt': Deletes Cluster resources like Pods and Services but retains Persistent Volume Claims (PVCs), allowing for data preservation while stopping other operations. - 'Delete': Extends the 'Halt' policy by also removing PVCs, leading to a thorough cleanup while removing all persistent data. - 'WipeOut': An aggressive policy that deletes all Cluster resources, including volume snapshots and backups in external storage. This results in complete data removal and should be used cautiously, primarily in non-production environments to avoid irreversible data loss.  Warning: Choosing an inappropriate termination policy can result in data loss. The 'WipeOut' policy is particularly risky in production environments due to its irreversible nature.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("DoNotTerminate", "Halt", "Delete", "WipeOut"),
						},
					},

					"tolerations": schema.MapAttribute{
						Description:         "An array that specifies tolerations attached to the Cluster's Pods, allowing them to be scheduled onto nodes with matching taints.  Deprecated since v0.10. Use the 'schedulingPolicy' field instead.",
						MarkdownDescription: "An array that specifies tolerations attached to the Cluster's Pods, allowing them to be scheduled onto nodes with matching taints.  Deprecated since v0.10. Use the 'schedulingPolicy' field instead.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"topology": schema.StringAttribute{
						Description:         "Specifies the name of the ClusterTopology to be used when creating the Cluster.  This field defines which set of Components, as outlined in the ClusterDefinition, will be used to construct the Cluster based on the named topology. The ClusterDefinition may list multiple topologies under 'clusterdefinition.spec.topologies[*]', each tailored to different use cases or environments.  If 'topology' is not specified, the Cluster will use the default topology defined in the ClusterDefinition.  Note: Once set during the Cluster creation, the 'topology' field cannot be modified. It establishes the initial composition and structure of the Cluster and is intended for one-time configuration.",
						MarkdownDescription: "Specifies the name of the ClusterTopology to be used when creating the Cluster.  This field defines which set of Components, as outlined in the ClusterDefinition, will be used to construct the Cluster based on the named topology. The ClusterDefinition may list multiple topologies under 'clusterdefinition.spec.topologies[*]', each tailored to different use cases or environments.  If 'topology' is not specified, the Cluster will use the default topology defined in the ClusterDefinition.  Note: Once set during the Cluster creation, the 'topology' field cannot be modified. It establishes the initial composition and structure of the Cluster and is intended for one-time configuration.",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.LengthAtMost(32),
						},
					},
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppsKubeblocksIoClusterV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_cluster_v1alpha1_manifest")

	var model AppsKubeblocksIoClusterV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("Cluster")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
