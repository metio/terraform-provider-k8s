/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package data_fluid_io_v1alpha1

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
	_ datasource.DataSource              = &DataFluidIoAlluxioRuntimeV1Alpha1DataSource{}
	_ datasource.DataSourceWithConfigure = &DataFluidIoAlluxioRuntimeV1Alpha1DataSource{}
)

func NewDataFluidIoAlluxioRuntimeV1Alpha1DataSource() datasource.DataSource {
	return &DataFluidIoAlluxioRuntimeV1Alpha1DataSource{}
}

type DataFluidIoAlluxioRuntimeV1Alpha1DataSource struct {
	kubernetesClient dynamic.Interface
}

type DataFluidIoAlluxioRuntimeV1Alpha1DataSourceData struct {
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
		AlluxioVersion *struct {
			Image           *string `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			ImageTag        *string `tfsdk:"image_tag" json:"imageTag,omitempty"`
		} `tfsdk:"alluxio_version" json:"alluxioVersion,omitempty"`
		ApiGateway *struct {
			Enabled      *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			Env          *map[string]string `tfsdk:"env" json:"env,omitempty"`
			JvmOptions   *[]string          `tfsdk:"jvm_options" json:"jvmOptions,omitempty"`
			NetworkMode  *string            `tfsdk:"network_mode" json:"networkMode,omitempty"`
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PodMetadata  *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Ports      *map[string]string `tfsdk:"ports" json:"ports,omitempty"`
			Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
			Replicas   *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources  *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			VolumeMounts *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"api_gateway" json:"apiGateway,omitempty"`
		Data *struct {
			Pin      *bool  `tfsdk:"pin" json:"pin,omitempty"`
			Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
		} `tfsdk:"data" json:"data,omitempty"`
		DisablePrometheus *bool `tfsdk:"disable_prometheus" json:"disablePrometheus,omitempty"`
		Fuse              *struct {
			Args            *[]string          `tfsdk:"args" json:"args,omitempty"`
			CleanPolicy     *string            `tfsdk:"clean_policy" json:"cleanPolicy,omitempty"`
			Env             *map[string]string `tfsdk:"env" json:"env,omitempty"`
			Global          *bool              `tfsdk:"global" json:"global,omitempty"`
			Image           *string            `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string            `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			ImageTag        *string            `tfsdk:"image_tag" json:"imageTag,omitempty"`
			JvmOptions      *[]string          `tfsdk:"jvm_options" json:"jvmOptions,omitempty"`
			NetworkMode     *string            `tfsdk:"network_mode" json:"networkMode,omitempty"`
			NodeSelector    *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PodMetadata     *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
			Resources  *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			VolumeMounts *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"fuse" json:"fuse,omitempty"`
		HadoopConfig *string `tfsdk:"hadoop_config" json:"hadoopConfig,omitempty"`
		InitUsers    *struct {
			Env             *map[string]string `tfsdk:"env" json:"env,omitempty"`
			Image           *string            `tfsdk:"image" json:"image,omitempty"`
			ImagePullPolicy *string            `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
			ImageTag        *string            `tfsdk:"image_tag" json:"imageTag,omitempty"`
			Resources       *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
		} `tfsdk:"init_users" json:"initUsers,omitempty"`
		JobMaster *struct {
			Enabled      *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			Env          *map[string]string `tfsdk:"env" json:"env,omitempty"`
			JvmOptions   *[]string          `tfsdk:"jvm_options" json:"jvmOptions,omitempty"`
			NetworkMode  *string            `tfsdk:"network_mode" json:"networkMode,omitempty"`
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PodMetadata  *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Ports      *map[string]string `tfsdk:"ports" json:"ports,omitempty"`
			Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
			Replicas   *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources  *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			VolumeMounts *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"job_master" json:"jobMaster,omitempty"`
		JobWorker *struct {
			Enabled      *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			Env          *map[string]string `tfsdk:"env" json:"env,omitempty"`
			JvmOptions   *[]string          `tfsdk:"jvm_options" json:"jvmOptions,omitempty"`
			NetworkMode  *string            `tfsdk:"network_mode" json:"networkMode,omitempty"`
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PodMetadata  *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Ports      *map[string]string `tfsdk:"ports" json:"ports,omitempty"`
			Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
			Replicas   *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources  *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			VolumeMounts *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"job_worker" json:"jobWorker,omitempty"`
		JvmOptions *[]string `tfsdk:"jvm_options" json:"jvmOptions,omitempty"`
		Management *struct {
			CleanCachePolicy *struct {
				GracePeriodSeconds *int64 `tfsdk:"grace_period_seconds" json:"gracePeriodSeconds,omitempty"`
				MaxRetryAttempts   *int64 `tfsdk:"max_retry_attempts" json:"maxRetryAttempts,omitempty"`
			} `tfsdk:"clean_cache_policy" json:"cleanCachePolicy,omitempty"`
			MetadataSyncPolicy *struct {
				AutoSync *bool `tfsdk:"auto_sync" json:"autoSync,omitempty"`
			} `tfsdk:"metadata_sync_policy" json:"metadataSyncPolicy,omitempty"`
		} `tfsdk:"management" json:"management,omitempty"`
		Master *struct {
			Enabled      *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			Env          *map[string]string `tfsdk:"env" json:"env,omitempty"`
			JvmOptions   *[]string          `tfsdk:"jvm_options" json:"jvmOptions,omitempty"`
			NetworkMode  *string            `tfsdk:"network_mode" json:"networkMode,omitempty"`
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PodMetadata  *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Ports      *map[string]string `tfsdk:"ports" json:"ports,omitempty"`
			Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
			Replicas   *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources  *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			VolumeMounts *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"master" json:"master,omitempty"`
		PodMetadata *struct {
			Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
			Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
		Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
		Replicas   *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
		RunAs      *struct {
			Gid   *int64  `tfsdk:"gid" json:"gid,omitempty"`
			Group *string `tfsdk:"group" json:"group,omitempty"`
			Uid   *int64  `tfsdk:"uid" json:"uid,omitempty"`
			User  *string `tfsdk:"user" json:"user,omitempty"`
		} `tfsdk:"run_as" json:"runAs,omitempty"`
		Tieredstore *struct {
			Levels *[]struct {
				High         *string `tfsdk:"high" json:"high,omitempty"`
				Low          *string `tfsdk:"low" json:"low,omitempty"`
				Mediumtype   *string `tfsdk:"mediumtype" json:"mediumtype,omitempty"`
				Path         *string `tfsdk:"path" json:"path,omitempty"`
				Quota        *string `tfsdk:"quota" json:"quota,omitempty"`
				QuotaList    *string `tfsdk:"quota_list" json:"quotaList,omitempty"`
				VolumeSource *struct {
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
							Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
							Spec     *struct {
								AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
								DataSource  *struct {
									ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
									Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
								} `tfsdk:"data_source" json:"dataSource,omitempty"`
								DataSourceRef *struct {
									ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
									Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
									Name     *string `tfsdk:"name" json:"name,omitempty"`
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
					Nfs *struct {
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
				} `tfsdk:"volume_source" json:"volumeSource,omitempty"`
				VolumeType *string `tfsdk:"volume_type" json:"volumeType,omitempty"`
			} `tfsdk:"levels" json:"levels,omitempty"`
		} `tfsdk:"tieredstore" json:"tieredstore,omitempty"`
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
					Metadata *map[string]string `tfsdk:"metadata" json:"metadata,omitempty"`
					Spec     *struct {
						AccessModes *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
						DataSource  *struct {
							ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
							Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
						} `tfsdk:"data_source" json:"dataSource,omitempty"`
						DataSourceRef *struct {
							ApiGroup *string `tfsdk:"api_group" json:"apiGroup,omitempty"`
							Kind     *string `tfsdk:"kind" json:"kind,omitempty"`
							Name     *string `tfsdk:"name" json:"name,omitempty"`
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
		Worker *struct {
			Enabled      *bool              `tfsdk:"enabled" json:"enabled,omitempty"`
			Env          *map[string]string `tfsdk:"env" json:"env,omitempty"`
			JvmOptions   *[]string          `tfsdk:"jvm_options" json:"jvmOptions,omitempty"`
			NetworkMode  *string            `tfsdk:"network_mode" json:"networkMode,omitempty"`
			NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
			PodMetadata  *struct {
				Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				Labels      *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
			} `tfsdk:"pod_metadata" json:"podMetadata,omitempty"`
			Ports      *map[string]string `tfsdk:"ports" json:"ports,omitempty"`
			Properties *map[string]string `tfsdk:"properties" json:"properties,omitempty"`
			Replicas   *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources  *struct {
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			VolumeMounts *[]struct {
				MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name             *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" json:"volumeMounts,omitempty"`
		} `tfsdk:"worker" json:"worker,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *DataFluidIoAlluxioRuntimeV1Alpha1DataSource) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_data_fluid_io_alluxio_runtime_v1alpha1"
}

func (r *DataFluidIoAlluxioRuntimeV1Alpha1DataSource) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "AlluxioRuntime is the Schema for the alluxioruntimes API",
		MarkdownDescription: "AlluxioRuntime is the Schema for the alluxioruntimes API",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.namespace/metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.namespace/metadata.name`.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"api_version": schema.StringAttribute{
				Description:         "The API group of the requested resource.",
				MarkdownDescription: "The API group of the requested resource.",
				Required:            false,
				Optional:            false,
				Computed:            true,
			},

			"kind": schema.StringAttribute{
				Description:         "The type of the requested resource.",
				MarkdownDescription: "The type of the requested resource.",
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
				Description:         "AlluxioRuntimeSpec defines the desired state of AlluxioRuntime",
				MarkdownDescription: "AlluxioRuntimeSpec defines the desired state of AlluxioRuntime",
				Attributes: map[string]schema.Attribute{
					"alluxio_version": schema.SingleNestedAttribute{
						Description:         "The version information that instructs fluid to orchestrate a particular version of Alluxio.",
						MarkdownDescription: "The version information that instructs fluid to orchestrate a particular version of Alluxio.",
						Attributes: map[string]schema.Attribute{
							"image": schema.StringAttribute{
								Description:         "Image (e.g. alluxio/alluxio)",
								MarkdownDescription: "Image (e.g. alluxio/alluxio)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "One of the three policies: 'Always', 'IfNotPresent', 'Never'",
								MarkdownDescription: "One of the three policies: 'Always', 'IfNotPresent', 'Never'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_tag": schema.StringAttribute{
								Description:         "Image tag (e.g. 2.3.0-SNAPSHOT)",
								MarkdownDescription: "Image tag (e.g. 2.3.0-SNAPSHOT)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"api_gateway": schema.SingleNestedAttribute{
						Description:         "The component spec of Alluxio API Gateway",
						MarkdownDescription: "The component spec of Alluxio API Gateway",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								MarkdownDescription: "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"env": schema.MapAttribute{
								Description:         "Environment variables that will be used by Alluxio component. <br>",
								MarkdownDescription: "Environment variables that will be used by Alluxio component. <br>",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"jvm_options": schema.ListAttribute{
								Description:         "Options for JVM",
								MarkdownDescription: "Options for JVM",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"network_mode": schema.StringAttribute{
								Description:         "Whether to use hostnetwork or not",
								MarkdownDescription: "Whether to use hostnetwork or not",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the master to fit on a node",
								MarkdownDescription: "NodeSelector is a selector which must be true for the master to fit on a node",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								MarkdownDescription: "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations are annotations of pod specification",
										MarkdownDescription: "Annotations are annotations of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels are labels of pod specification",
										MarkdownDescription: "Labels are labels of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ports": schema.MapAttribute{
								Description:         "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								MarkdownDescription: "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"properties": schema.MapAttribute{
								Description:         "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								MarkdownDescription: "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								MarkdownDescription: "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								MarkdownDescription: "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								MarkdownDescription: "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"data": schema.SingleNestedAttribute{
						Description:         "Management strategies for the dataset to which the runtime is bound",
						MarkdownDescription: "Management strategies for the dataset to which the runtime is bound",
						Attributes: map[string]schema.Attribute{
							"pin": schema.BoolAttribute{
								Description:         "Pin the dataset or not. Refer to <a href='https://docs.alluxio.io/os/user/stable/en/operation/User-CLI.html#pin'>Alluxio User-CLI pin</a>",
								MarkdownDescription: "Pin the dataset or not. Refer to <a href='https://docs.alluxio.io/os/user/stable/en/operation/User-CLI.html#pin'>Alluxio User-CLI pin</a>",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replicas": schema.Int64Attribute{
								Description:         "The copies of the dataset",
								MarkdownDescription: "The copies of the dataset",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"disable_prometheus": schema.BoolAttribute{
						Description:         "Disable monitoring for Alluxio Runtime Prometheus is enabled by default",
						MarkdownDescription: "Disable monitoring for Alluxio Runtime Prometheus is enabled by default",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"fuse": schema.SingleNestedAttribute{
						Description:         "The component spec of Alluxio Fuse",
						MarkdownDescription: "The component spec of Alluxio Fuse",
						Attributes: map[string]schema.Attribute{
							"args": schema.ListAttribute{
								Description:         "Arguments that will be passed to Alluxio Fuse",
								MarkdownDescription: "Arguments that will be passed to Alluxio Fuse",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"clean_policy": schema.StringAttribute{
								Description:         "CleanPolicy decides when to clean Alluxio Fuse pods. Currently Fluid supports two policies: OnDemand and OnRuntimeDeleted OnDemand cleans fuse pod once the fuse pod on some node is not needed OnRuntimeDeleted cleans fuse pod only when the cache runtime is deleted Defaults to OnRuntimeDeleted",
								MarkdownDescription: "CleanPolicy decides when to clean Alluxio Fuse pods. Currently Fluid supports two policies: OnDemand and OnRuntimeDeleted OnDemand cleans fuse pod once the fuse pod on some node is not needed OnRuntimeDeleted cleans fuse pod only when the cache runtime is deleted Defaults to OnRuntimeDeleted",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"env": schema.MapAttribute{
								Description:         "Environment variables that will be used by Alluxio Fuse",
								MarkdownDescription: "Environment variables that will be used by Alluxio Fuse",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"global": schema.BoolAttribute{
								Description:         "If the fuse client should be deployed in global mode, otherwise the affinity should be considered",
								MarkdownDescription: "If the fuse client should be deployed in global mode, otherwise the affinity should be considered",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image": schema.StringAttribute{
								Description:         "Image for Alluxio Fuse(e.g. alluxio/alluxio-fuse)",
								MarkdownDescription: "Image for Alluxio Fuse(e.g. alluxio/alluxio-fuse)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "One of the three policies: 'Always', 'IfNotPresent', 'Never'",
								MarkdownDescription: "One of the three policies: 'Always', 'IfNotPresent', 'Never'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_tag": schema.StringAttribute{
								Description:         "Image Tag for Alluxio Fuse(e.g. 2.3.0-SNAPSHOT)",
								MarkdownDescription: "Image Tag for Alluxio Fuse(e.g. 2.3.0-SNAPSHOT)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"jvm_options": schema.ListAttribute{
								Description:         "Options for JVM",
								MarkdownDescription: "Options for JVM",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"network_mode": schema.StringAttribute{
								Description:         "Whether to use hostnetwork or not",
								MarkdownDescription: "Whether to use hostnetwork or not",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the fuse client to fit on a node, this option only effect when global is enabled",
								MarkdownDescription: "NodeSelector is a selector which must be true for the fuse client to fit on a node, this option only effect when global is enabled",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata defines labels and annotations that will be propagated to Alluxio's fuse pods",
								MarkdownDescription: "PodMetadata defines labels and annotations that will be propagated to Alluxio's fuse pods",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations are annotations of pod specification",
										MarkdownDescription: "Annotations are annotations of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels are labels of pod specification",
										MarkdownDescription: "Labels are labels of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"properties": schema.MapAttribute{
								Description:         "Configurable properties for Alluxio System. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								MarkdownDescription: "Configurable properties for Alluxio System. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources that will be requested by Alluxio Fuse. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								MarkdownDescription: "Resources that will be requested by Alluxio Fuse. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								MarkdownDescription: "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"hadoop_config": schema.StringAttribute{
						Description:         "Name of the configMap used to support HDFS configurations when using HDFS as Alluxio's UFS. The configMap must be in the same namespace with the AlluxioRuntime. The configMap should contain user-specific HDFS conf files in it. For now, only 'hdfs-site.xml' and 'core-site.xml' are supported. It must take the filename of the conf file as the key and content of the file as the value.",
						MarkdownDescription: "Name of the configMap used to support HDFS configurations when using HDFS as Alluxio's UFS. The configMap must be in the same namespace with the AlluxioRuntime. The configMap should contain user-specific HDFS conf files in it. For now, only 'hdfs-site.xml' and 'core-site.xml' are supported. It must take the filename of the conf file as the key and content of the file as the value.",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"init_users": schema.SingleNestedAttribute{
						Description:         "The spec of init users",
						MarkdownDescription: "The spec of init users",
						Attributes: map[string]schema.Attribute{
							"env": schema.MapAttribute{
								Description:         "Environment variables that will be used by initialize the users for runtime",
								MarkdownDescription: "Environment variables that will be used by initialize the users for runtime",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image": schema.StringAttribute{
								Description:         "Image for initialize the users for runtime(e.g. alluxio/alluxio-User init)",
								MarkdownDescription: "Image for initialize the users for runtime(e.g. alluxio/alluxio-User init)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_pull_policy": schema.StringAttribute{
								Description:         "One of the three policies: 'Always', 'IfNotPresent', 'Never'",
								MarkdownDescription: "One of the three policies: 'Always', 'IfNotPresent', 'Never'",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"image_tag": schema.StringAttribute{
								Description:         "Image Tag for initialize the users for runtime(e.g. 2.3.0-SNAPSHOT)",
								MarkdownDescription: "Image Tag for initialize the users for runtime(e.g. 2.3.0-SNAPSHOT)",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources that will be requested by initialize the users for runtime. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								MarkdownDescription: "Resources that will be requested by initialize the users for runtime. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
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

					"job_master": schema.SingleNestedAttribute{
						Description:         "The component spec of Alluxio job master",
						MarkdownDescription: "The component spec of Alluxio job master",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								MarkdownDescription: "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"env": schema.MapAttribute{
								Description:         "Environment variables that will be used by Alluxio component. <br>",
								MarkdownDescription: "Environment variables that will be used by Alluxio component. <br>",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"jvm_options": schema.ListAttribute{
								Description:         "Options for JVM",
								MarkdownDescription: "Options for JVM",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"network_mode": schema.StringAttribute{
								Description:         "Whether to use hostnetwork or not",
								MarkdownDescription: "Whether to use hostnetwork or not",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the master to fit on a node",
								MarkdownDescription: "NodeSelector is a selector which must be true for the master to fit on a node",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								MarkdownDescription: "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations are annotations of pod specification",
										MarkdownDescription: "Annotations are annotations of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels are labels of pod specification",
										MarkdownDescription: "Labels are labels of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ports": schema.MapAttribute{
								Description:         "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								MarkdownDescription: "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"properties": schema.MapAttribute{
								Description:         "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								MarkdownDescription: "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								MarkdownDescription: "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								MarkdownDescription: "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								MarkdownDescription: "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"job_worker": schema.SingleNestedAttribute{
						Description:         "The component spec of Alluxio job Worker",
						MarkdownDescription: "The component spec of Alluxio job Worker",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								MarkdownDescription: "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"env": schema.MapAttribute{
								Description:         "Environment variables that will be used by Alluxio component. <br>",
								MarkdownDescription: "Environment variables that will be used by Alluxio component. <br>",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"jvm_options": schema.ListAttribute{
								Description:         "Options for JVM",
								MarkdownDescription: "Options for JVM",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"network_mode": schema.StringAttribute{
								Description:         "Whether to use hostnetwork or not",
								MarkdownDescription: "Whether to use hostnetwork or not",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the master to fit on a node",
								MarkdownDescription: "NodeSelector is a selector which must be true for the master to fit on a node",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								MarkdownDescription: "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations are annotations of pod specification",
										MarkdownDescription: "Annotations are annotations of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels are labels of pod specification",
										MarkdownDescription: "Labels are labels of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ports": schema.MapAttribute{
								Description:         "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								MarkdownDescription: "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"properties": schema.MapAttribute{
								Description:         "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								MarkdownDescription: "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								MarkdownDescription: "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								MarkdownDescription: "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								MarkdownDescription: "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"jvm_options": schema.ListAttribute{
						Description:         "Options for JVM",
						MarkdownDescription: "Options for JVM",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"management": schema.SingleNestedAttribute{
						Description:         "RuntimeManagement defines policies when managing the runtime",
						MarkdownDescription: "RuntimeManagement defines policies when managing the runtime",
						Attributes: map[string]schema.Attribute{
							"clean_cache_policy": schema.SingleNestedAttribute{
								Description:         "CleanCachePolicy defines the policy of cleaning cache when shutting down the runtime",
								MarkdownDescription: "CleanCachePolicy defines the policy of cleaning cache when shutting down the runtime",
								Attributes: map[string]schema.Attribute{
									"grace_period_seconds": schema.Int64Attribute{
										Description:         "Optional duration in seconds the cache needs to clean gracefully. May be decreased in delete runtime request. Value must be non-negative integer. The value zero indicates clean immediately via the timeout command (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with timeout command. Set this value longer than the expected cleanup time for your process.",
										MarkdownDescription: "Optional duration in seconds the cache needs to clean gracefully. May be decreased in delete runtime request. Value must be non-negative integer. The value zero indicates clean immediately via the timeout command (no opportunity to shut down). If this value is nil, the default grace period will be used instead. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with timeout command. Set this value longer than the expected cleanup time for your process.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"max_retry_attempts": schema.Int64Attribute{
										Description:         "Optional max retry Attempts when cleanCache function returns an error after execution, runtime attempts to run it three more times by default. With Maximum Retry Attempts, you can customize the maximum number of retries. This gives you the option to continue processing retries.",
										MarkdownDescription: "Optional max retry Attempts when cleanCache function returns an error after execution, runtime attempts to run it three more times by default. With Maximum Retry Attempts, you can customize the maximum number of retries. This gives you the option to continue processing retries.",
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"metadata_sync_policy": schema.SingleNestedAttribute{
								Description:         "MetadataSyncPolicy defines the policy of syncing metadata when setting up the runtime. If not set,",
								MarkdownDescription: "MetadataSyncPolicy defines the policy of syncing metadata when setting up the runtime. If not set,",
								Attributes: map[string]schema.Attribute{
									"auto_sync": schema.BoolAttribute{
										Description:         "AutoSync enables automatic metadata sync when setting up a runtime. If not set, it defaults to true.",
										MarkdownDescription: "AutoSync enables automatic metadata sync when setting up a runtime. If not set, it defaults to true.",
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

					"master": schema.SingleNestedAttribute{
						Description:         "The component spec of Alluxio master",
						MarkdownDescription: "The component spec of Alluxio master",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								MarkdownDescription: "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"env": schema.MapAttribute{
								Description:         "Environment variables that will be used by Alluxio component. <br>",
								MarkdownDescription: "Environment variables that will be used by Alluxio component. <br>",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"jvm_options": schema.ListAttribute{
								Description:         "Options for JVM",
								MarkdownDescription: "Options for JVM",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"network_mode": schema.StringAttribute{
								Description:         "Whether to use hostnetwork or not",
								MarkdownDescription: "Whether to use hostnetwork or not",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the master to fit on a node",
								MarkdownDescription: "NodeSelector is a selector which must be true for the master to fit on a node",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								MarkdownDescription: "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations are annotations of pod specification",
										MarkdownDescription: "Annotations are annotations of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels are labels of pod specification",
										MarkdownDescription: "Labels are labels of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ports": schema.MapAttribute{
								Description:         "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								MarkdownDescription: "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"properties": schema.MapAttribute{
								Description:         "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								MarkdownDescription: "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								MarkdownDescription: "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								MarkdownDescription: "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								MarkdownDescription: "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"pod_metadata": schema.SingleNestedAttribute{
						Description:         "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
						MarkdownDescription: "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
						Attributes: map[string]schema.Attribute{
							"annotations": schema.MapAttribute{
								Description:         "Annotations are annotations of pod specification",
								MarkdownDescription: "Annotations are annotations of pod specification",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"labels": schema.MapAttribute{
								Description:         "Labels are labels of pod specification",
								MarkdownDescription: "Labels are labels of pod specification",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"properties": schema.MapAttribute{
						Description:         "Configurable properties for Alluxio system. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
						MarkdownDescription: "Configurable properties for Alluxio system. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"replicas": schema.Int64Attribute{
						Description:         "The replicas of the worker, need to be specified",
						MarkdownDescription: "The replicas of the worker, need to be specified",
						Required:            false,
						Optional:            false,
						Computed:            true,
					},

					"run_as": schema.SingleNestedAttribute{
						Description:         "Manage the user to run Alluxio Runtime",
						MarkdownDescription: "Manage the user to run Alluxio Runtime",
						Attributes: map[string]schema.Attribute{
							"gid": schema.Int64Attribute{
								Description:         "The gid to run the alluxio runtime",
								MarkdownDescription: "The gid to run the alluxio runtime",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"group": schema.StringAttribute{
								Description:         "The group name to run the alluxio runtime",
								MarkdownDescription: "The group name to run the alluxio runtime",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"uid": schema.Int64Attribute{
								Description:         "The uid to run the alluxio runtime",
								MarkdownDescription: "The uid to run the alluxio runtime",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"user": schema.StringAttribute{
								Description:         "The user name to run the alluxio runtime",
								MarkdownDescription: "The user name to run the alluxio runtime",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"tieredstore": schema.SingleNestedAttribute{
						Description:         "Tiered storage used by Alluxio",
						MarkdownDescription: "Tiered storage used by Alluxio",
						Attributes: map[string]schema.Attribute{
							"levels": schema.ListNestedAttribute{
								Description:         "configurations for multiple tiers",
								MarkdownDescription: "configurations for multiple tiers",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"high": schema.StringAttribute{
											Description:         "Ratio of high watermark of the tier (e.g. 0.9)",
											MarkdownDescription: "Ratio of high watermark of the tier (e.g. 0.9)",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"low": schema.StringAttribute{
											Description:         "Ratio of low watermark of the tier (e.g. 0.7)",
											MarkdownDescription: "Ratio of low watermark of the tier (e.g. 0.7)",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mediumtype": schema.StringAttribute{
											Description:         "Medium Type of the tier. One of the three types: 'MEM', 'SSD', 'HDD'",
											MarkdownDescription: "Medium Type of the tier. One of the three types: 'MEM', 'SSD', 'HDD'",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"path": schema.StringAttribute{
											Description:         "File paths to be used for the tier. Multiple paths are supported. Multiple paths should be separated with comma. For example: '/mnt/cache1,/mnt/cache2'.",
											MarkdownDescription: "File paths to be used for the tier. Multiple paths are supported. Multiple paths should be separated with comma. For example: '/mnt/cache1,/mnt/cache2'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"quota": schema.StringAttribute{
											Description:         "Quota for the whole tier. (e.g. 100Gi) Please note that if there're multiple paths used for this tierstore, the quota will be equally divided into these paths. If you'd like to set quota for each, path, see QuotaList for more information.",
											MarkdownDescription: "Quota for the whole tier. (e.g. 100Gi) Please note that if there're multiple paths used for this tierstore, the quota will be equally divided into these paths. If you'd like to set quota for each, path, see QuotaList for more information.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"quota_list": schema.StringAttribute{
											Description:         "QuotaList are quotas used to set quota on multiple paths. Quotas should be separated with comma. Quotas in this list will be set to paths with the same order in Path. For example, with Path defined with '/mnt/cache1,/mnt/cache2' and QuotaList set to '100Gi, 50Gi', then we get 100GiB cache storage under '/mnt/cache1' and 50GiB under '/mnt/cache2'. Also note that num of quotas must be consistent with the num of paths defined in Path.",
											MarkdownDescription: "QuotaList are quotas used to set quota on multiple paths. Quotas should be separated with comma. Quotas in this list will be set to paths with the same order in Path. For example, with Path defined with '/mnt/cache1,/mnt/cache2' and QuotaList set to '100Gi, 50Gi', then we get 100GiB cache storage under '/mnt/cache1' and 50GiB under '/mnt/cache2'. Also note that num of quotas must be consistent with the num of paths defined in Path.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"volume_source": schema.SingleNestedAttribute{
											Description:         "VolumeSource is the volume source of the tier. It follows the form of corev1.VolumeSource. For now, users should only specify VolumeSource when VolumeType is set to emptyDir.",
											MarkdownDescription: "VolumeSource is the volume source of the tier. It follows the form of corev1.VolumeSource. For now, users should only specify VolumeSource when VolumeType is set to emptyDir.",
											Attributes: map[string]schema.Attribute{
												"aws_elastic_block_store": schema.SingleNestedAttribute{
													Description:         "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
													MarkdownDescription: "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
															MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"partition": schema.Int64Attribute{
															Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
															MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
															MarkdownDescription: "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume_id": schema.StringAttribute{
															Description:         "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
															MarkdownDescription: "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"azure_disk": schema.SingleNestedAttribute{
													Description:         "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
													MarkdownDescription: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
													Attributes: map[string]schema.Attribute{
														"caching_mode": schema.StringAttribute{
															Description:         "Host Caching mode: None, Read Only, Read Write.",
															MarkdownDescription: "Host Caching mode: None, Read Only, Read Write.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"disk_name": schema.StringAttribute{
															Description:         "The Name of the data disk in the blob storage",
															MarkdownDescription: "The Name of the data disk in the blob storage",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"disk_uri": schema.StringAttribute{
															Description:         "The URI the data disk in the blob storage",
															MarkdownDescription: "The URI the data disk in the blob storage",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
															MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"kind": schema.StringAttribute{
															Description:         "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
															MarkdownDescription: "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"azure_file": schema.SingleNestedAttribute{
													Description:         "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
													MarkdownDescription: "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
													Attributes: map[string]schema.Attribute{
														"read_only": schema.BoolAttribute{
															Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_name": schema.StringAttribute{
															Description:         "the name of secret that contains Azure Storage Account Name and Key",
															MarkdownDescription: "the name of secret that contains Azure Storage Account Name and Key",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"share_name": schema.StringAttribute{
															Description:         "Share Name",
															MarkdownDescription: "Share Name",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"cephfs": schema.SingleNestedAttribute{
													Description:         "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
													MarkdownDescription: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
													Attributes: map[string]schema.Attribute{
														"monitors": schema.ListAttribute{
															Description:         "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															MarkdownDescription: "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"path": schema.StringAttribute{
															Description:         "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
															MarkdownDescription: "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_file": schema.StringAttribute{
															Description:         "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															MarkdownDescription: "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															MarkdownDescription: "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"user": schema.StringAttribute{
															Description:         "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															MarkdownDescription: "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"cinder": schema.SingleNestedAttribute{
													Description:         "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													MarkdownDescription: "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "Optional: points to a secret object containing parameters used to connect to OpenStack.",
															MarkdownDescription: "Optional: points to a secret object containing parameters used to connect to OpenStack.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"volume_id": schema.StringAttribute{
															Description:         "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															MarkdownDescription: "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"config_map": schema.SingleNestedAttribute{
													Description:         "ConfigMap represents a configMap that should populate this volume",
													MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",
													Attributes: map[string]schema.Attribute{
														"default_mode": schema.Int64Attribute{
															Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
															MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"items": schema.ListNestedAttribute{
															Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
															MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key to project.",
																		MarkdownDescription: "The key to project.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"mode": schema.Int64Attribute{
																		Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"path": schema.StringAttribute{
																		Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																		MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

														"name": schema.StringAttribute{
															Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the ConfigMap or its keys must be defined",
															MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"csi": schema.SingleNestedAttribute{
													Description:         "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
													MarkdownDescription: "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
													Attributes: map[string]schema.Attribute{
														"driver": schema.StringAttribute{
															Description:         "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
															MarkdownDescription: "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
															MarkdownDescription: "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"node_publish_secret_ref": schema.SingleNestedAttribute{
															Description:         "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
															MarkdownDescription: "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Specifies a read-only configuration for the volume. Defaults to false (read/write).",
															MarkdownDescription: "Specifies a read-only configuration for the volume. Defaults to false (read/write).",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume_attributes": schema.MapAttribute{
															Description:         "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
															MarkdownDescription: "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"downward_api": schema.SingleNestedAttribute{
													Description:         "DownwardAPI represents downward API about the pod that should populate this volume",
													MarkdownDescription: "DownwardAPI represents downward API about the pod that should populate this volume",
													Attributes: map[string]schema.Attribute{
														"default_mode": schema.Int64Attribute{
															Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
															MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
															Required:            false,
															Optional:            false,
															Computed:            true,
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
																				Optional:            false,
																				Computed:            true,
																			},

																			"field_path": schema.StringAttribute{
																				Description:         "Path of the field to select in the specified API version.",
																				MarkdownDescription: "Path of the field to select in the specified API version.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"mode": schema.Int64Attribute{
																		Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"path": schema.StringAttribute{
																		Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																		MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"resource_field_ref": schema.SingleNestedAttribute{
																		Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																		MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																		Attributes: map[string]schema.Attribute{
																			"container_name": schema.StringAttribute{
																				Description:         "Container name: required for volumes, optional for env vars",
																				MarkdownDescription: "Container name: required for volumes, optional for env vars",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"divisor": schema.StringAttribute{
																				Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																				MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"resource": schema.StringAttribute{
																				Description:         "Required: resource to select",
																				MarkdownDescription: "Required: resource to select",
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

												"empty_dir": schema.SingleNestedAttribute{
													Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
													MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
													Attributes: map[string]schema.Attribute{
														"medium": schema.StringAttribute{
															Description:         "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															MarkdownDescription: "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"size_limit": schema.StringAttribute{
															Description:         "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
															MarkdownDescription: "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"ephemeral": schema.SingleNestedAttribute{
													Description:         "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
													MarkdownDescription: "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
													Attributes: map[string]schema.Attribute{
														"volume_claim_template": schema.SingleNestedAttribute{
															Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
															MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
															Attributes: map[string]schema.Attribute{
																"metadata": schema.MapAttribute{
																	Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																	MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"spec": schema.SingleNestedAttribute{
																	Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																	MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
																	Attributes: map[string]schema.Attribute{
																		"access_modes": schema.ListAttribute{
																			Description:         "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																			MarkdownDescription: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																			ElementType:         types.StringType,
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"data_source": schema.SingleNestedAttribute{
																			Description:         "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																			MarkdownDescription: "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																			Attributes: map[string]schema.Attribute{
																				"api_group": schema.StringAttribute{
																					Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																					MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"kind": schema.StringAttribute{
																					Description:         "Kind is the type of resource being referenced",
																					MarkdownDescription: "Kind is the type of resource being referenced",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name is the name of resource being referenced",
																					MarkdownDescription: "Name is the name of resource being referenced",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"data_source_ref": schema.SingleNestedAttribute{
																			Description:         "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																			MarkdownDescription: "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																			Attributes: map[string]schema.Attribute{
																				"api_group": schema.StringAttribute{
																					Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																					MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"kind": schema.StringAttribute{
																					Description:         "Kind is the type of resource being referenced",
																					MarkdownDescription: "Kind is the type of resource being referenced",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"name": schema.StringAttribute{
																					Description:         "Name is the name of resource being referenced",
																					MarkdownDescription: "Name is the name of resource being referenced",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"resources": schema.SingleNestedAttribute{
																			Description:         "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																			MarkdownDescription: "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																			Attributes: map[string]schema.Attribute{
																				"limits": schema.MapAttribute{
																					Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																					MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"requests": schema.MapAttribute{
																					Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																					MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"selector": schema.SingleNestedAttribute{
																			Description:         "A label query over volumes to consider for binding.",
																			MarkdownDescription: "A label query over volumes to consider for binding.",
																			Attributes: map[string]schema.Attribute{
																				"match_expressions": schema.ListNestedAttribute{
																					Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																					NestedObject: schema.NestedAttributeObject{
																						Attributes: map[string]schema.Attribute{
																							"key": schema.StringAttribute{
																								Description:         "key is the label key that the selector applies to.",
																								MarkdownDescription: "key is the label key that the selector applies to.",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"operator": schema.StringAttribute{
																								Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																								Required:            false,
																								Optional:            false,
																								Computed:            true,
																							},

																							"values": schema.ListAttribute{
																								Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																								MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																								ElementType:         types.StringType,
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

																				"match_labels": schema.MapAttribute{
																					Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																					ElementType:         types.StringType,
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"storage_class_name": schema.StringAttribute{
																			Description:         "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																			MarkdownDescription: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"volume_mode": schema.StringAttribute{
																			Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																			MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"volume_name": schema.StringAttribute{
																			Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
																			MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",
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
													Required: false,
													Optional: false,
													Computed: true,
												},

												"fc": schema.SingleNestedAttribute{
													Description:         "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
													MarkdownDescription: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
															MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"lun": schema.Int64Attribute{
															Description:         "Optional: FC target lun number",
															MarkdownDescription: "Optional: FC target lun number",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"target_ww_ns": schema.ListAttribute{
															Description:         "Optional: FC target worldwide names (WWNs)",
															MarkdownDescription: "Optional: FC target worldwide names (WWNs)",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"wwids": schema.ListAttribute{
															Description:         "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
															MarkdownDescription: "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"flex_volume": schema.SingleNestedAttribute{
													Description:         "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
													MarkdownDescription: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
													Attributes: map[string]schema.Attribute{
														"driver": schema.StringAttribute{
															Description:         "Driver is the name of the driver to use for this volume.",
															MarkdownDescription: "Driver is the name of the driver to use for this volume.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
															MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"options": schema.MapAttribute{
															Description:         "Optional: Extra command options if any.",
															MarkdownDescription: "Optional: Extra command options if any.",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
															MarkdownDescription: "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

												"flocker": schema.SingleNestedAttribute{
													Description:         "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
													MarkdownDescription: "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
													Attributes: map[string]schema.Attribute{
														"dataset_name": schema.StringAttribute{
															Description:         "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
															MarkdownDescription: "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"dataset_uuid": schema.StringAttribute{
															Description:         "UUID of the dataset. This is unique identifier of a Flocker dataset",
															MarkdownDescription: "UUID of the dataset. This is unique identifier of a Flocker dataset",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"gce_persistent_disk": schema.SingleNestedAttribute{
													Description:         "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													MarkdownDescription: "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
															MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"partition": schema.Int64Attribute{
															Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"pd_name": schema.StringAttribute{
															Description:         "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															MarkdownDescription: "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"git_repo": schema.SingleNestedAttribute{
													Description:         "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
													MarkdownDescription: "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
													Attributes: map[string]schema.Attribute{
														"directory": schema.StringAttribute{
															Description:         "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
															MarkdownDescription: "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"repository": schema.StringAttribute{
															Description:         "Repository URL",
															MarkdownDescription: "Repository URL",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"revision": schema.StringAttribute{
															Description:         "Commit hash for the specified revision.",
															MarkdownDescription: "Commit hash for the specified revision.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"glusterfs": schema.SingleNestedAttribute{
													Description:         "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
													MarkdownDescription: "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
													Attributes: map[string]schema.Attribute{
														"endpoints": schema.StringAttribute{
															Description:         "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
															MarkdownDescription: "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"path": schema.StringAttribute{
															Description:         "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
															MarkdownDescription: "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
															MarkdownDescription: "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"host_path": schema.SingleNestedAttribute{
													Description:         "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
													MarkdownDescription: "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
													Attributes: map[string]schema.Attribute{
														"path": schema.StringAttribute{
															Description:         "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
															MarkdownDescription: "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"type": schema.StringAttribute{
															Description:         "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
															MarkdownDescription: "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"iscsi": schema.SingleNestedAttribute{
													Description:         "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
													MarkdownDescription: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
													Attributes: map[string]schema.Attribute{
														"chap_auth_discovery": schema.BoolAttribute{
															Description:         "whether support iSCSI Discovery CHAP authentication",
															MarkdownDescription: "whether support iSCSI Discovery CHAP authentication",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"chap_auth_session": schema.BoolAttribute{
															Description:         "whether support iSCSI Session CHAP authentication",
															MarkdownDescription: "whether support iSCSI Session CHAP authentication",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
															MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"initiator_name": schema.StringAttribute{
															Description:         "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
															MarkdownDescription: "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"iqn": schema.StringAttribute{
															Description:         "Target iSCSI Qualified Name.",
															MarkdownDescription: "Target iSCSI Qualified Name.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"iscsi_interface": schema.StringAttribute{
															Description:         "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
															MarkdownDescription: "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"lun": schema.Int64Attribute{
															Description:         "iSCSI Target Lun number.",
															MarkdownDescription: "iSCSI Target Lun number.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"portals": schema.ListAttribute{
															Description:         "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
															MarkdownDescription: "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
															MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "CHAP Secret for iSCSI target and initiator authentication",
															MarkdownDescription: "CHAP Secret for iSCSI target and initiator authentication",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"target_portal": schema.StringAttribute{
															Description:         "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
															MarkdownDescription: "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"nfs": schema.SingleNestedAttribute{
													Description:         "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													MarkdownDescription: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													Attributes: map[string]schema.Attribute{
														"path": schema.StringAttribute{
															Description:         "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															MarkdownDescription: "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															MarkdownDescription: "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"server": schema.StringAttribute{
															Description:         "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															MarkdownDescription: "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"persistent_volume_claim": schema.SingleNestedAttribute{
													Description:         "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
													MarkdownDescription: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
													Attributes: map[string]schema.Attribute{
														"claim_name": schema.StringAttribute{
															Description:         "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
															MarkdownDescription: "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Will force the ReadOnly setting in VolumeMounts. Default false.",
															MarkdownDescription: "Will force the ReadOnly setting in VolumeMounts. Default false.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"photon_persistent_disk": schema.SingleNestedAttribute{
													Description:         "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
													MarkdownDescription: "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
															MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"pd_id": schema.StringAttribute{
															Description:         "ID that identifies Photon Controller persistent disk",
															MarkdownDescription: "ID that identifies Photon Controller persistent disk",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"portworx_volume": schema.SingleNestedAttribute{
													Description:         "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
													MarkdownDescription: "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
															MarkdownDescription: "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume_id": schema.StringAttribute{
															Description:         "VolumeID uniquely identifies a Portworx volume",
															MarkdownDescription: "VolumeID uniquely identifies a Portworx volume",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"projected": schema.SingleNestedAttribute{
													Description:         "Items for all in one resources secrets, configmaps, and downward API",
													MarkdownDescription: "Items for all in one resources secrets, configmaps, and downward API",
													Attributes: map[string]schema.Attribute{
														"default_mode": schema.Int64Attribute{
															Description:         "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
															MarkdownDescription: "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"sources": schema.ListNestedAttribute{
															Description:         "list of volume projections",
															MarkdownDescription: "list of volume projections",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"config_map": schema.SingleNestedAttribute{
																		Description:         "information about the configMap data to project",
																		MarkdownDescription: "information about the configMap data to project",
																		Attributes: map[string]schema.Attribute{
																			"items": schema.ListNestedAttribute{
																				Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																				MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "The key to project.",
																							MarkdownDescription: "The key to project.",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"mode": schema.Int64Attribute{
																							Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"path": schema.StringAttribute{
																							Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																							MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the ConfigMap or its keys must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"downward_api": schema.SingleNestedAttribute{
																		Description:         "information about the downwardAPI data to project",
																		MarkdownDescription: "information about the downwardAPI data to project",
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
																									Optional:            false,
																									Computed:            true,
																								},

																								"field_path": schema.StringAttribute{
																									Description:         "Path of the field to select in the specified API version.",
																									MarkdownDescription: "Path of the field to select in the specified API version.",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},
																							},
																							Required: false,
																							Optional: false,
																							Computed: true,
																						},

																						"mode": schema.Int64Attribute{
																							Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"path": schema.StringAttribute{
																							Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																							MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"resource_field_ref": schema.SingleNestedAttribute{
																							Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																							MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																							Attributes: map[string]schema.Attribute{
																								"container_name": schema.StringAttribute{
																									Description:         "Container name: required for volumes, optional for env vars",
																									MarkdownDescription: "Container name: required for volumes, optional for env vars",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"divisor": schema.StringAttribute{
																									Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																									MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																									Required:            false,
																									Optional:            false,
																									Computed:            true,
																								},

																								"resource": schema.StringAttribute{
																									Description:         "Required: resource to select",
																									MarkdownDescription: "Required: resource to select",
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

																	"secret": schema.SingleNestedAttribute{
																		Description:         "information about the secret data to project",
																		MarkdownDescription: "information about the secret data to project",
																		Attributes: map[string]schema.Attribute{
																			"items": schema.ListNestedAttribute{
																				Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																				MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"key": schema.StringAttribute{
																							Description:         "The key to project.",
																							MarkdownDescription: "The key to project.",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"mode": schema.Int64Attribute{
																							Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            false,
																							Computed:            true,
																						},

																						"path": schema.StringAttribute{
																							Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																							MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																			"name": schema.StringAttribute{
																				Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"optional": schema.BoolAttribute{
																				Description:         "Specify whether the Secret or its key must be defined",
																				MarkdownDescription: "Specify whether the Secret or its key must be defined",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},
																		},
																		Required: false,
																		Optional: false,
																		Computed: true,
																	},

																	"service_account_token": schema.SingleNestedAttribute{
																		Description:         "information about the serviceAccountToken data to project",
																		MarkdownDescription: "information about the serviceAccountToken data to project",
																		Attributes: map[string]schema.Attribute{
																			"audience": schema.StringAttribute{
																				Description:         "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																				MarkdownDescription: "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"expiration_seconds": schema.Int64Attribute{
																				Description:         "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																				MarkdownDescription: "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"path": schema.StringAttribute{
																				Description:         "Path is the path relative to the mount point of the file to project the token into.",
																				MarkdownDescription: "Path is the path relative to the mount point of the file to project the token into.",
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

												"quobyte": schema.SingleNestedAttribute{
													Description:         "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
													MarkdownDescription: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
													Attributes: map[string]schema.Attribute{
														"group": schema.StringAttribute{
															Description:         "Group to map volume access to Default is no group",
															MarkdownDescription: "Group to map volume access to Default is no group",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
															MarkdownDescription: "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"registry": schema.StringAttribute{
															Description:         "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
															MarkdownDescription: "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"tenant": schema.StringAttribute{
															Description:         "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
															MarkdownDescription: "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"user": schema.StringAttribute{
															Description:         "User to map volume access to Defaults to serivceaccount user",
															MarkdownDescription: "User to map volume access to Defaults to serivceaccount user",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume": schema.StringAttribute{
															Description:         "Volume is a string that references an already created Quobyte volume by name.",
															MarkdownDescription: "Volume is a string that references an already created Quobyte volume by name.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"rbd": schema.SingleNestedAttribute{
													Description:         "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
													MarkdownDescription: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
															MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"image": schema.StringAttribute{
															Description:         "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															MarkdownDescription: "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"keyring": schema.StringAttribute{
															Description:         "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															MarkdownDescription: "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"monitors": schema.ListAttribute{
															Description:         "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															MarkdownDescription: "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"pool": schema.StringAttribute{
															Description:         "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															MarkdownDescription: "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															MarkdownDescription: "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"user": schema.StringAttribute{
															Description:         "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															MarkdownDescription: "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"scale_io": schema.SingleNestedAttribute{
													Description:         "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
													MarkdownDescription: "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
															MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"gateway": schema.StringAttribute{
															Description:         "The host address of the ScaleIO API Gateway.",
															MarkdownDescription: "The host address of the ScaleIO API Gateway.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"protection_domain": schema.StringAttribute{
															Description:         "The name of the ScaleIO Protection Domain for the configured storage.",
															MarkdownDescription: "The name of the ScaleIO Protection Domain for the configured storage.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
															MarkdownDescription: "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"ssl_enabled": schema.BoolAttribute{
															Description:         "Flag to enable/disable SSL communication with Gateway, default false",
															MarkdownDescription: "Flag to enable/disable SSL communication with Gateway, default false",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"storage_mode": schema.StringAttribute{
															Description:         "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
															MarkdownDescription: "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"storage_pool": schema.StringAttribute{
															Description:         "The ScaleIO Storage Pool associated with the protection domain.",
															MarkdownDescription: "The ScaleIO Storage Pool associated with the protection domain.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"system": schema.StringAttribute{
															Description:         "The name of the storage system as configured in ScaleIO.",
															MarkdownDescription: "The name of the storage system as configured in ScaleIO.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume_name": schema.StringAttribute{
															Description:         "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
															MarkdownDescription: "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"secret": schema.SingleNestedAttribute{
													Description:         "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
													MarkdownDescription: "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
													Attributes: map[string]schema.Attribute{
														"default_mode": schema.Int64Attribute{
															Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
															MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"items": schema.ListNestedAttribute{
															Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
															MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "The key to project.",
																		MarkdownDescription: "The key to project.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"mode": schema.Int64Attribute{
																		Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																		Required:            false,
																		Optional:            false,
																		Computed:            true,
																	},

																	"path": schema.StringAttribute{
																		Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																		MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

														"optional": schema.BoolAttribute{
															Description:         "Specify whether the Secret or its keys must be defined",
															MarkdownDescription: "Specify whether the Secret or its keys must be defined",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_name": schema.StringAttribute{
															Description:         "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
															MarkdownDescription: "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"storageos": schema.SingleNestedAttribute{
													Description:         "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
													MarkdownDescription: "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
															MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"read_only": schema.BoolAttribute{
															Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"secret_ref": schema.SingleNestedAttribute{
															Description:         "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
															MarkdownDescription: "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
															Attributes: map[string]schema.Attribute{
																"name": schema.StringAttribute{
																	Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"volume_name": schema.StringAttribute{
															Description:         "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
															MarkdownDescription: "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume_namespace": schema.StringAttribute{
															Description:         "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
															MarkdownDescription: "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},
													},
													Required: false,
													Optional: false,
													Computed: true,
												},

												"vsphere_volume": schema.SingleNestedAttribute{
													Description:         "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
													MarkdownDescription: "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
													Attributes: map[string]schema.Attribute{
														"fs_type": schema.StringAttribute{
															Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
															MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"storage_policy_id": schema.StringAttribute{
															Description:         "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
															MarkdownDescription: "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"storage_policy_name": schema.StringAttribute{
															Description:         "Storage Policy Based Management (SPBM) profile name.",
															MarkdownDescription: "Storage Policy Based Management (SPBM) profile name.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume_path": schema.StringAttribute{
															Description:         "Path that identifies vSphere volume vmdk",
															MarkdownDescription: "Path that identifies vSphere volume vmdk",
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

										"volume_type": schema.StringAttribute{
											Description:         "VolumeType is the volume type of the tier. Should be one of the three types: 'hostPath', 'emptyDir' and 'volumeTemplate'. If not set, defaults to hostPath.",
											MarkdownDescription: "VolumeType is the volume type of the tier. Should be one of the three types: 'hostPath', 'emptyDir' and 'volumeTemplate'. If not set, defaults to hostPath.",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"volumes": schema.ListNestedAttribute{
						Description:         "Volumes is the list of Kubernetes volumes that can be mounted by the alluxio runtime components and/or fuses.",
						MarkdownDescription: "Volumes is the list of Kubernetes volumes that can be mounted by the alluxio runtime components and/or fuses.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"aws_elastic_block_store": schema.SingleNestedAttribute{
									Description:         "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
									MarkdownDescription: "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"partition": schema.Int64Attribute{
											Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
											MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											MarkdownDescription: "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"volume_id": schema.StringAttribute{
											Description:         "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											MarkdownDescription: "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"azure_disk": schema.SingleNestedAttribute{
									Description:         "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
									MarkdownDescription: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
									Attributes: map[string]schema.Attribute{
										"caching_mode": schema.StringAttribute{
											Description:         "Host Caching mode: None, Read Only, Read Write.",
											MarkdownDescription: "Host Caching mode: None, Read Only, Read Write.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"disk_name": schema.StringAttribute{
											Description:         "The Name of the data disk in the blob storage",
											MarkdownDescription: "The Name of the data disk in the blob storage",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"disk_uri": schema.StringAttribute{
											Description:         "The URI the data disk in the blob storage",
											MarkdownDescription: "The URI the data disk in the blob storage",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"kind": schema.StringAttribute{
											Description:         "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
											MarkdownDescription: "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"azure_file": schema.SingleNestedAttribute{
									Description:         "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
									MarkdownDescription: "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
									Attributes: map[string]schema.Attribute{
										"read_only": schema.BoolAttribute{
											Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_name": schema.StringAttribute{
											Description:         "the name of secret that contains Azure Storage Account Name and Key",
											MarkdownDescription: "the name of secret that contains Azure Storage Account Name and Key",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"share_name": schema.StringAttribute{
											Description:         "Share Name",
											MarkdownDescription: "Share Name",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"cephfs": schema.SingleNestedAttribute{
									Description:         "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
									MarkdownDescription: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
									Attributes: map[string]schema.Attribute{
										"monitors": schema.ListAttribute{
											Description:         "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"path": schema.StringAttribute{
											Description:         "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
											MarkdownDescription: "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_file": schema.StringAttribute{
											Description:         "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"user": schema.StringAttribute{
											Description:         "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											MarkdownDescription: "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"cinder": schema.SingleNestedAttribute{
									Description:         "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
									MarkdownDescription: "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "Optional: points to a secret object containing parameters used to connect to OpenStack.",
											MarkdownDescription: "Optional: points to a secret object containing parameters used to connect to OpenStack.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"volume_id": schema.StringAttribute{
											Description:         "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											MarkdownDescription: "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"config_map": schema.SingleNestedAttribute{
									Description:         "ConfigMap represents a configMap that should populate this volume",
									MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",
									Attributes: map[string]schema.Attribute{
										"default_mode": schema.Int64Attribute{
											Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"items": schema.ListNestedAttribute{
											Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
											MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to project.",
														MarkdownDescription: "The key to project.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"mode": schema.Int64Attribute{
														Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"path": schema.StringAttribute{
														Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
														MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

										"name": schema.StringAttribute{
											Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
											MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the ConfigMap or its keys must be defined",
											MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"csi": schema.SingleNestedAttribute{
									Description:         "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
									MarkdownDescription: "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
									Attributes: map[string]schema.Attribute{
										"driver": schema.StringAttribute{
											Description:         "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
											MarkdownDescription: "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
											MarkdownDescription: "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"node_publish_secret_ref": schema.SingleNestedAttribute{
											Description:         "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
											MarkdownDescription: "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Specifies a read-only configuration for the volume. Defaults to false (read/write).",
											MarkdownDescription: "Specifies a read-only configuration for the volume. Defaults to false (read/write).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"volume_attributes": schema.MapAttribute{
											Description:         "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
											MarkdownDescription: "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"downward_api": schema.SingleNestedAttribute{
									Description:         "DownwardAPI represents downward API about the pod that should populate this volume",
									MarkdownDescription: "DownwardAPI represents downward API about the pod that should populate this volume",
									Attributes: map[string]schema.Attribute{
										"default_mode": schema.Int64Attribute{
											Description:         "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											Required:            false,
											Optional:            false,
											Computed:            true,
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
																Optional:            false,
																Computed:            true,
															},

															"field_path": schema.StringAttribute{
																Description:         "Path of the field to select in the specified API version.",
																MarkdownDescription: "Path of the field to select in the specified API version.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"mode": schema.Int64Attribute{
														Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"path": schema.StringAttribute{
														Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
														MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"resource_field_ref": schema.SingleNestedAttribute{
														Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
														MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
														Attributes: map[string]schema.Attribute{
															"container_name": schema.StringAttribute{
																Description:         "Container name: required for volumes, optional for env vars",
																MarkdownDescription: "Container name: required for volumes, optional for env vars",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"divisor": schema.StringAttribute{
																Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"resource": schema.StringAttribute{
																Description:         "Required: resource to select",
																MarkdownDescription: "Required: resource to select",
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

								"empty_dir": schema.SingleNestedAttribute{
									Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
									MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
									Attributes: map[string]schema.Attribute{
										"medium": schema.StringAttribute{
											Description:         "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
											MarkdownDescription: "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"size_limit": schema.StringAttribute{
											Description:         "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
											MarkdownDescription: "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"ephemeral": schema.SingleNestedAttribute{
									Description:         "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
									MarkdownDescription: "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
									Attributes: map[string]schema.Attribute{
										"volume_claim_template": schema.SingleNestedAttribute{
											Description:         "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
											MarkdownDescription: "Will be used to create a stand-alone PVC to provision the volume. The pod in which this EphemeralVolumeSource is embedded will be the owner of the PVC, i.e. the PVC will be deleted together with the pod.  The name of the PVC will be '<pod name>-<volume name>' where '<volume name>' is the name from the 'PodSpec.Volumes' array entry. Pod validation will reject the pod if the concatenated name is not valid for a PVC (for example, too long).  An existing PVC with that name that is not owned by the pod will *not* be used for the pod to avoid using an unrelated volume by mistake. Starting the pod is then blocked until the unrelated PVC is removed. If such a pre-created PVC is meant to be used by the pod, the PVC has to updated with an owner reference to the pod once the pod exists. Normally this should not be necessary, but it may be useful when manually reconstructing a broken cluster.  This field is read-only and no changes will be made by Kubernetes to the PVC after it has been created.  Required, must not be nil.",
											Attributes: map[string]schema.Attribute{
												"metadata": schema.MapAttribute{
													Description:         "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
													MarkdownDescription: "May contain labels and annotations that will be copied into the PVC when creating it. No other fields are allowed and will be rejected during validation.",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            false,
													Computed:            true,
												},

												"spec": schema.SingleNestedAttribute{
													Description:         "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
													MarkdownDescription: "The specification for the PersistentVolumeClaim. The entire content is copied unchanged into the PVC that gets created from this template. The same fields as in a PersistentVolumeClaim are also valid here.",
													Attributes: map[string]schema.Attribute{
														"access_modes": schema.ListAttribute{
															Description:         "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
															MarkdownDescription: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
															ElementType:         types.StringType,
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"data_source": schema.SingleNestedAttribute{
															Description:         "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
															MarkdownDescription: "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
															Attributes: map[string]schema.Attribute{
																"api_group": schema.StringAttribute{
																	Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																	MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind is the type of resource being referenced",
																	MarkdownDescription: "Kind is the type of resource being referenced",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "Name is the name of resource being referenced",
																	MarkdownDescription: "Name is the name of resource being referenced",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"data_source_ref": schema.SingleNestedAttribute{
															Description:         "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
															MarkdownDescription: "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
															Attributes: map[string]schema.Attribute{
																"api_group": schema.StringAttribute{
																	Description:         "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																	MarkdownDescription: "APIGroup is the group for the resource being referenced. If APIGroup is not specified, the specified Kind must be in the core API group. For any other third-party types, APIGroup is required.",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"kind": schema.StringAttribute{
																	Description:         "Kind is the type of resource being referenced",
																	MarkdownDescription: "Kind is the type of resource being referenced",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"name": schema.StringAttribute{
																	Description:         "Name is the name of resource being referenced",
																	MarkdownDescription: "Name is the name of resource being referenced",
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"resources": schema.SingleNestedAttribute{
															Description:         "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
															MarkdownDescription: "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
															Attributes: map[string]schema.Attribute{
																"limits": schema.MapAttribute{
																	Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},

																"requests": schema.MapAttribute{
																	Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"selector": schema.SingleNestedAttribute{
															Description:         "A label query over volumes to consider for binding.",
															MarkdownDescription: "A label query over volumes to consider for binding.",
															Attributes: map[string]schema.Attribute{
																"match_expressions": schema.ListNestedAttribute{
																	Description:         "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	MarkdownDescription: "matchExpressions is a list of label selector requirements. The requirements are ANDed.",
																	NestedObject: schema.NestedAttributeObject{
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "key is the label key that the selector applies to.",
																				MarkdownDescription: "key is the label key that the selector applies to.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"operator": schema.StringAttribute{
																				Description:         "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				MarkdownDescription: "operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.",
																				Required:            false,
																				Optional:            false,
																				Computed:            true,
																			},

																			"values": schema.ListAttribute{
																				Description:         "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				MarkdownDescription: "values is an array of string values. If the operator is In or NotIn, the values array must be non-empty. If the operator is Exists or DoesNotExist, the values array must be empty. This array is replaced during a strategic merge patch.",
																				ElementType:         types.StringType,
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

																"match_labels": schema.MapAttribute{
																	Description:         "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	MarkdownDescription: "matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is 'key', the operator is 'In', and the values array contains only 'value'. The requirements are ANDed.",
																	ElementType:         types.StringType,
																	Required:            false,
																	Optional:            false,
																	Computed:            true,
																},
															},
															Required: false,
															Optional: false,
															Computed: true,
														},

														"storage_class_name": schema.StringAttribute{
															Description:         "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
															MarkdownDescription: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume_mode": schema.StringAttribute{
															Description:         "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
															MarkdownDescription: "volumeMode defines what type of volume is required by the claim. Value of Filesystem is implied when not included in claim spec.",
															Required:            false,
															Optional:            false,
															Computed:            true,
														},

														"volume_name": schema.StringAttribute{
															Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
															MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",
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
									Required: false,
									Optional: false,
									Computed: true,
								},

								"fc": schema.SingleNestedAttribute{
									Description:         "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
									MarkdownDescription: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"lun": schema.Int64Attribute{
											Description:         "Optional: FC target lun number",
											MarkdownDescription: "Optional: FC target lun number",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"target_ww_ns": schema.ListAttribute{
											Description:         "Optional: FC target worldwide names (WWNs)",
											MarkdownDescription: "Optional: FC target worldwide names (WWNs)",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"wwids": schema.ListAttribute{
											Description:         "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
											MarkdownDescription: "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"flex_volume": schema.SingleNestedAttribute{
									Description:         "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
									MarkdownDescription: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
									Attributes: map[string]schema.Attribute{
										"driver": schema.StringAttribute{
											Description:         "Driver is the name of the driver to use for this volume.",
											MarkdownDescription: "Driver is the name of the driver to use for this volume.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
											MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"options": schema.MapAttribute{
											Description:         "Optional: Extra command options if any.",
											MarkdownDescription: "Optional: Extra command options if any.",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
											MarkdownDescription: "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

								"flocker": schema.SingleNestedAttribute{
									Description:         "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
									MarkdownDescription: "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
									Attributes: map[string]schema.Attribute{
										"dataset_name": schema.StringAttribute{
											Description:         "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
											MarkdownDescription: "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"dataset_uuid": schema.StringAttribute{
											Description:         "UUID of the dataset. This is unique identifier of a Flocker dataset",
											MarkdownDescription: "UUID of the dataset. This is unique identifier of a Flocker dataset",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"gce_persistent_disk": schema.SingleNestedAttribute{
									Description:         "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
									MarkdownDescription: "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"partition": schema.Int64Attribute{
											Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"pd_name": schema.StringAttribute{
											Description:         "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											MarkdownDescription: "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"git_repo": schema.SingleNestedAttribute{
									Description:         "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
									MarkdownDescription: "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
									Attributes: map[string]schema.Attribute{
										"directory": schema.StringAttribute{
											Description:         "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
											MarkdownDescription: "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"repository": schema.StringAttribute{
											Description:         "Repository URL",
											MarkdownDescription: "Repository URL",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"revision": schema.StringAttribute{
											Description:         "Commit hash for the specified revision.",
											MarkdownDescription: "Commit hash for the specified revision.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"glusterfs": schema.SingleNestedAttribute{
									Description:         "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
									MarkdownDescription: "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
									Attributes: map[string]schema.Attribute{
										"endpoints": schema.StringAttribute{
											Description:         "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											MarkdownDescription: "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"path": schema.StringAttribute{
											Description:         "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											MarkdownDescription: "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											MarkdownDescription: "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"host_path": schema.SingleNestedAttribute{
									Description:         "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
									MarkdownDescription: "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
									Attributes: map[string]schema.Attribute{
										"path": schema.StringAttribute{
											Description:         "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
											MarkdownDescription: "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"type": schema.StringAttribute{
											Description:         "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
											MarkdownDescription: "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"iscsi": schema.SingleNestedAttribute{
									Description:         "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
									MarkdownDescription: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
									Attributes: map[string]schema.Attribute{
										"chap_auth_discovery": schema.BoolAttribute{
											Description:         "whether support iSCSI Discovery CHAP authentication",
											MarkdownDescription: "whether support iSCSI Discovery CHAP authentication",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"chap_auth_session": schema.BoolAttribute{
											Description:         "whether support iSCSI Session CHAP authentication",
											MarkdownDescription: "whether support iSCSI Session CHAP authentication",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"initiator_name": schema.StringAttribute{
											Description:         "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
											MarkdownDescription: "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"iqn": schema.StringAttribute{
											Description:         "Target iSCSI Qualified Name.",
											MarkdownDescription: "Target iSCSI Qualified Name.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"iscsi_interface": schema.StringAttribute{
											Description:         "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
											MarkdownDescription: "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"lun": schema.Int64Attribute{
											Description:         "iSCSI Target Lun number.",
											MarkdownDescription: "iSCSI Target Lun number.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"portals": schema.ListAttribute{
											Description:         "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
											MarkdownDescription: "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
											MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "CHAP Secret for iSCSI target and initiator authentication",
											MarkdownDescription: "CHAP Secret for iSCSI target and initiator authentication",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"target_portal": schema.StringAttribute{
											Description:         "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
											MarkdownDescription: "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"name": schema.StringAttribute{
									Description:         "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
									MarkdownDescription: "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
									Required:            false,
									Optional:            false,
									Computed:            true,
								},

								"nfs": schema.SingleNestedAttribute{
									Description:         "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
									MarkdownDescription: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
									Attributes: map[string]schema.Attribute{
										"path": schema.StringAttribute{
											Description:         "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											MarkdownDescription: "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											MarkdownDescription: "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"server": schema.StringAttribute{
											Description:         "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											MarkdownDescription: "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"persistent_volume_claim": schema.SingleNestedAttribute{
									Description:         "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									MarkdownDescription: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
									Attributes: map[string]schema.Attribute{
										"claim_name": schema.StringAttribute{
											Description:         "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											MarkdownDescription: "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Will force the ReadOnly setting in VolumeMounts. Default false.",
											MarkdownDescription: "Will force the ReadOnly setting in VolumeMounts. Default false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"photon_persistent_disk": schema.SingleNestedAttribute{
									Description:         "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
									MarkdownDescription: "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"pd_id": schema.StringAttribute{
											Description:         "ID that identifies Photon Controller persistent disk",
											MarkdownDescription: "ID that identifies Photon Controller persistent disk",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"portworx_volume": schema.SingleNestedAttribute{
									Description:         "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
									MarkdownDescription: "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"volume_id": schema.StringAttribute{
											Description:         "VolumeID uniquely identifies a Portworx volume",
											MarkdownDescription: "VolumeID uniquely identifies a Portworx volume",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"projected": schema.SingleNestedAttribute{
									Description:         "Items for all in one resources secrets, configmaps, and downward API",
									MarkdownDescription: "Items for all in one resources secrets, configmaps, and downward API",
									Attributes: map[string]schema.Attribute{
										"default_mode": schema.Int64Attribute{
											Description:         "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sources": schema.ListNestedAttribute{
											Description:         "list of volume projections",
											MarkdownDescription: "list of volume projections",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config_map": schema.SingleNestedAttribute{
														Description:         "information about the configMap data to project",
														MarkdownDescription: "information about the configMap data to project",
														Attributes: map[string]schema.Attribute{
															"items": schema.ListNestedAttribute{
																Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "The key to project.",
																			MarkdownDescription: "The key to project.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"path": schema.StringAttribute{
																			Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																			MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"optional": schema.BoolAttribute{
																Description:         "Specify whether the ConfigMap or its keys must be defined",
																MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"downward_api": schema.SingleNestedAttribute{
														Description:         "information about the downwardAPI data to project",
														MarkdownDescription: "information about the downwardAPI data to project",
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
																					Optional:            false,
																					Computed:            true,
																				},

																				"field_path": schema.StringAttribute{
																					Description:         "Path of the field to select in the specified API version.",
																					MarkdownDescription: "Path of the field to select in the specified API version.",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},
																			},
																			Required: false,
																			Optional: false,
																			Computed: true,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "Optional: mode bits used to set permissions on this file, must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"path": schema.StringAttribute{
																			Description:         "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																			MarkdownDescription: "Required: Path is  the relative path name of the file to be created. Must not be absolute or contain the '..' path. Must be utf-8 encoded. The first item of the relative path must not start with '..'",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"resource_field_ref": schema.SingleNestedAttribute{
																			Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																			MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, requests.cpu and requests.memory) are currently supported.",
																			Attributes: map[string]schema.Attribute{
																				"container_name": schema.StringAttribute{
																					Description:         "Container name: required for volumes, optional for env vars",
																					MarkdownDescription: "Container name: required for volumes, optional for env vars",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"divisor": schema.StringAttribute{
																					Description:         "Specifies the output format of the exposed resources, defaults to '1'",
																					MarkdownDescription: "Specifies the output format of the exposed resources, defaults to '1'",
																					Required:            false,
																					Optional:            false,
																					Computed:            true,
																				},

																				"resource": schema.StringAttribute{
																					Description:         "Required: resource to select",
																					MarkdownDescription: "Required: resource to select",
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

													"secret": schema.SingleNestedAttribute{
														Description:         "information about the secret data to project",
														MarkdownDescription: "information about the secret data to project",
														Attributes: map[string]schema.Attribute{
															"items": schema.ListNestedAttribute{
																Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																NestedObject: schema.NestedAttributeObject{
																	Attributes: map[string]schema.Attribute{
																		"key": schema.StringAttribute{
																			Description:         "The key to project.",
																			MarkdownDescription: "The key to project.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"mode": schema.Int64Attribute{
																			Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																			Required:            false,
																			Optional:            false,
																			Computed:            true,
																		},

																		"path": schema.StringAttribute{
																			Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																			MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"optional": schema.BoolAttribute{
																Description:         "Specify whether the Secret or its key must be defined",
																MarkdownDescription: "Specify whether the Secret or its key must be defined",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},
														},
														Required: false,
														Optional: false,
														Computed: true,
													},

													"service_account_token": schema.SingleNestedAttribute{
														Description:         "information about the serviceAccountToken data to project",
														MarkdownDescription: "information about the serviceAccountToken data to project",
														Attributes: map[string]schema.Attribute{
															"audience": schema.StringAttribute{
																Description:         "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																MarkdownDescription: "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"expiration_seconds": schema.Int64Attribute{
																Description:         "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																MarkdownDescription: "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																Required:            false,
																Optional:            false,
																Computed:            true,
															},

															"path": schema.StringAttribute{
																Description:         "Path is the path relative to the mount point of the file to project the token into.",
																MarkdownDescription: "Path is the path relative to the mount point of the file to project the token into.",
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

								"quobyte": schema.SingleNestedAttribute{
									Description:         "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
									MarkdownDescription: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
									Attributes: map[string]schema.Attribute{
										"group": schema.StringAttribute{
											Description:         "Group to map volume access to Default is no group",
											MarkdownDescription: "Group to map volume access to Default is no group",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
											MarkdownDescription: "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"registry": schema.StringAttribute{
											Description:         "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
											MarkdownDescription: "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"tenant": schema.StringAttribute{
											Description:         "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
											MarkdownDescription: "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"user": schema.StringAttribute{
											Description:         "User to map volume access to Defaults to serivceaccount user",
											MarkdownDescription: "User to map volume access to Defaults to serivceaccount user",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"volume": schema.StringAttribute{
											Description:         "Volume is a string that references an already created Quobyte volume by name.",
											MarkdownDescription: "Volume is a string that references an already created Quobyte volume by name.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"rbd": schema.SingleNestedAttribute{
									Description:         "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
									MarkdownDescription: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
											MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"image": schema.StringAttribute{
											Description:         "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"keyring": schema.StringAttribute{
											Description:         "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"monitors": schema.ListAttribute{
											Description:         "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											ElementType:         types.StringType,
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"pool": schema.StringAttribute{
											Description:         "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"user": schema.StringAttribute{
											Description:         "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											MarkdownDescription: "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"scale_io": schema.SingleNestedAttribute{
									Description:         "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
									MarkdownDescription: "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
											MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"gateway": schema.StringAttribute{
											Description:         "The host address of the ScaleIO API Gateway.",
											MarkdownDescription: "The host address of the ScaleIO API Gateway.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"protection_domain": schema.StringAttribute{
											Description:         "The name of the ScaleIO Protection Domain for the configured storage.",
											MarkdownDescription: "The name of the ScaleIO Protection Domain for the configured storage.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
											MarkdownDescription: "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"ssl_enabled": schema.BoolAttribute{
											Description:         "Flag to enable/disable SSL communication with Gateway, default false",
											MarkdownDescription: "Flag to enable/disable SSL communication with Gateway, default false",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"storage_mode": schema.StringAttribute{
											Description:         "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
											MarkdownDescription: "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"storage_pool": schema.StringAttribute{
											Description:         "The ScaleIO Storage Pool associated with the protection domain.",
											MarkdownDescription: "The ScaleIO Storage Pool associated with the protection domain.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"system": schema.StringAttribute{
											Description:         "The name of the storage system as configured in ScaleIO.",
											MarkdownDescription: "The name of the storage system as configured in ScaleIO.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"volume_name": schema.StringAttribute{
											Description:         "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
											MarkdownDescription: "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"secret": schema.SingleNestedAttribute{
									Description:         "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
									MarkdownDescription: "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
									Attributes: map[string]schema.Attribute{
										"default_mode": schema.Int64Attribute{
											Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"items": schema.ListNestedAttribute{
											Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
											MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"key": schema.StringAttribute{
														Description:         "The key to project.",
														MarkdownDescription: "The key to project.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"mode": schema.Int64Attribute{
														Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														Required:            false,
														Optional:            false,
														Computed:            true,
													},

													"path": schema.StringAttribute{
														Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
														MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

										"optional": schema.BoolAttribute{
											Description:         "Specify whether the Secret or its keys must be defined",
											MarkdownDescription: "Specify whether the Secret or its keys must be defined",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_name": schema.StringAttribute{
											Description:         "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
											MarkdownDescription: "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"storageos": schema.SingleNestedAttribute{
									Description:         "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
									MarkdownDescription: "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"secret_ref": schema.SingleNestedAttribute{
											Description:         "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
											MarkdownDescription: "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
													Required:            false,
													Optional:            false,
													Computed:            true,
												},
											},
											Required: false,
											Optional: false,
											Computed: true,
										},

										"volume_name": schema.StringAttribute{
											Description:         "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
											MarkdownDescription: "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"volume_namespace": schema.StringAttribute{
											Description:         "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
											MarkdownDescription: "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},
									},
									Required: false,
									Optional: false,
									Computed: true,
								},

								"vsphere_volume": schema.SingleNestedAttribute{
									Description:         "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
									MarkdownDescription: "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
									Attributes: map[string]schema.Attribute{
										"fs_type": schema.StringAttribute{
											Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"storage_policy_id": schema.StringAttribute{
											Description:         "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
											MarkdownDescription: "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"storage_policy_name": schema.StringAttribute{
											Description:         "Storage Policy Based Management (SPBM) profile name.",
											MarkdownDescription: "Storage Policy Based Management (SPBM) profile name.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"volume_path": schema.StringAttribute{
											Description:         "Path that identifies vSphere volume vmdk",
											MarkdownDescription: "Path that identifies vSphere volume vmdk",
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
						},
						Required: false,
						Optional: false,
						Computed: true,
					},

					"worker": schema.SingleNestedAttribute{
						Description:         "The component spec of Alluxio worker",
						MarkdownDescription: "The component spec of Alluxio worker",
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description:         "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								MarkdownDescription: "Enabled or Disabled for the components. For now, only  API Gateway is enabled or disabled.",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"env": schema.MapAttribute{
								Description:         "Environment variables that will be used by Alluxio component. <br>",
								MarkdownDescription: "Environment variables that will be used by Alluxio component. <br>",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"jvm_options": schema.ListAttribute{
								Description:         "Options for JVM",
								MarkdownDescription: "Options for JVM",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"network_mode": schema.StringAttribute{
								Description:         "Whether to use hostnetwork or not",
								MarkdownDescription: "Whether to use hostnetwork or not",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"node_selector": schema.MapAttribute{
								Description:         "NodeSelector is a selector which must be true for the master to fit on a node",
								MarkdownDescription: "NodeSelector is a selector which must be true for the master to fit on a node",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"pod_metadata": schema.SingleNestedAttribute{
								Description:         "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								MarkdownDescription: "PodMetadata defines labels and annotations that will be propagated to Alluxio's pods",
								Attributes: map[string]schema.Attribute{
									"annotations": schema.MapAttribute{
										Description:         "Annotations are annotations of pod specification",
										MarkdownDescription: "Annotations are annotations of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"labels": schema.MapAttribute{
										Description:         "Labels are labels of pod specification",
										MarkdownDescription: "Labels are labels of pod specification",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"ports": schema.MapAttribute{
								Description:         "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								MarkdownDescription: "Ports used by Alluxio(e.g. rpc: 19998 for master)",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"properties": schema.MapAttribute{
								Description:         "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								MarkdownDescription: "Configurable properties for the Alluxio component. <br> Refer to <a href='https://docs.alluxio.io/os/user/stable/en/reference/Properties-List.html'>Alluxio Configuration Properties</a> for more info",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								MarkdownDescription: "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								Required:            false,
								Optional:            false,
								Computed:            true,
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								MarkdownDescription: "Resources that will be requested by the Alluxio component. <br> <br> Resources are not allowed for ephemeral containers. Ephemeral containers use spare resources already allocated to the pod.",
								Attributes: map[string]schema.Attribute{
									"limits": schema.MapAttribute{
										Description:         "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Limits describes the maximum amount of compute resources allowed. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},

									"requests": schema.MapAttribute{
										Description:         "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										MarkdownDescription: "Requests describes the minimum amount of compute resources required. If Requests is omitted for a container, it defaults to Limits if that is explicitly specified, otherwise to an implementation-defined value. More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            false,
										Computed:            true,
									},
								},
								Required: false,
								Optional: false,
								Computed: true,
							},

							"volume_mounts": schema.ListNestedAttribute{
								Description:         "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								MarkdownDescription: "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the alluxio runtime component's filesystem.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted.  Must not contain ':'.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"name": schema.StringAttribute{
											Description:         "This must match the Name of a Volume.",
											MarkdownDescription: "This must match the Name of a Volume.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"read_only": schema.BoolAttribute{
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false.",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted. Defaults to '' (volume's root).",
											Required:            false,
											Optional:            false,
											Computed:            true,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive.",
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

func (r *DataFluidIoAlluxioRuntimeV1Alpha1DataSource) Configure(_ context.Context, request datasource.ConfigureRequest, response *datasource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if dataSourceData, ok := request.ProviderData.(*utilities.DataSourceData); ok {
		if dataSourceData.Offline {
			response.Diagnostics.Append(utilities.OfflineProviderError())
		} else {
			r.kubernetesClient = dataSourceData.Client
		}
	} else {
		response.Diagnostics.Append(utilities.UnexpectedDataSourceDataError(request.ProviderData))
	}
}

func (r *DataFluidIoAlluxioRuntimeV1Alpha1DataSource) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read data source k8s_data_fluid_io_alluxio_runtime_v1alpha1")

	var data DataFluidIoAlluxioRuntimeV1Alpha1DataSourceData
	response.Diagnostics.Append(request.Config.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "data.fluid.io", Version: "v1alpha1", Resource: "alluxioruntimes"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.Append(utilities.GetNamespacedResourceError(err, data.Metadata.Name, data.Metadata.Namespace))
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalJsonError(err))
		return
	}

	var readResponse DataFluidIoAlluxioRuntimeV1Alpha1DataSourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.Append(utilities.JsonUnmarshalError(err))
		return
	}

	data.ID = types.StringValue(fmt.Sprintf("%s/%s", data.Metadata.Namespace, data.Metadata.Name))
	data.ApiVersion = pointer.String("data.fluid.io/v1alpha1")
	data.Kind = pointer.String("AlluxioRuntime")
	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}
