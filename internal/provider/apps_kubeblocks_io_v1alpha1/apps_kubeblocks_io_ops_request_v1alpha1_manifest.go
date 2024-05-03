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
	_ datasource.DataSource = &AppsKubeblocksIoOpsRequestV1Alpha1Manifest{}
)

func NewAppsKubeblocksIoOpsRequestV1Alpha1Manifest() datasource.DataSource {
	return &AppsKubeblocksIoOpsRequestV1Alpha1Manifest{}
}

type AppsKubeblocksIoOpsRequestV1Alpha1Manifest struct{}

type AppsKubeblocksIoOpsRequestV1Alpha1ManifestData struct {
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
		BackupSpec *struct {
			BackupMethod     *string `tfsdk:"backup_method" json:"backupMethod,omitempty"`
			BackupName       *string `tfsdk:"backup_name" json:"backupName,omitempty"`
			BackupPolicyName *string `tfsdk:"backup_policy_name" json:"backupPolicyName,omitempty"`
			DeletionPolicy   *string `tfsdk:"deletion_policy" json:"deletionPolicy,omitempty"`
			ParentBackupName *string `tfsdk:"parent_backup_name" json:"parentBackupName,omitempty"`
			RetentionPeriod  *string `tfsdk:"retention_period" json:"retentionPeriod,omitempty"`
		} `tfsdk:"backup_spec" json:"backupSpec,omitempty"`
		Cancel     *bool   `tfsdk:"cancel" json:"cancel,omitempty"`
		ClusterRef *string `tfsdk:"cluster_ref" json:"clusterRef,omitempty"`
		CustomSpec *struct {
			Components *[]struct {
				ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
				Parameters    *[]struct {
					Name  *string `tfsdk:"name" json:"name,omitempty"`
					Value *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"parameters" json:"parameters,omitempty"`
			} `tfsdk:"components" json:"components,omitempty"`
			OpsDefinitionRef   *string `tfsdk:"ops_definition_ref" json:"opsDefinitionRef,omitempty"`
			Parallelism        *string `tfsdk:"parallelism" json:"parallelism,omitempty"`
			ServiceAccountName *string `tfsdk:"service_account_name" json:"serviceAccountName,omitempty"`
		} `tfsdk:"custom_spec" json:"customSpec,omitempty"`
		Expose *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Services      *[]struct {
				Annotations    *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
				IpFamilies     *[]string          `tfsdk:"ip_families" json:"ipFamilies,omitempty"`
				IpFamilyPolicy *string            `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
				Name           *string            `tfsdk:"name" json:"name,omitempty"`
				Ports          *[]struct {
					AppProtocol *string `tfsdk:"app_protocol" json:"appProtocol,omitempty"`
					Name        *string `tfsdk:"name" json:"name,omitempty"`
					NodePort    *int64  `tfsdk:"node_port" json:"nodePort,omitempty"`
					Port        *int64  `tfsdk:"port" json:"port,omitempty"`
					Protocol    *string `tfsdk:"protocol" json:"protocol,omitempty"`
					TargetPort  *string `tfsdk:"target_port" json:"targetPort,omitempty"`
				} `tfsdk:"ports" json:"ports,omitempty"`
				RoleSelector *string            `tfsdk:"role_selector" json:"roleSelector,omitempty"`
				Selector     *map[string]string `tfsdk:"selector" json:"selector,omitempty"`
				ServiceType  *string            `tfsdk:"service_type" json:"serviceType,omitempty"`
			} `tfsdk:"services" json:"services,omitempty"`
			Switch *string `tfsdk:"switch" json:"switch,omitempty"`
		} `tfsdk:"expose" json:"expose,omitempty"`
		Force             *bool `tfsdk:"force" json:"force,omitempty"`
		HorizontalScaling *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Instances     *[]struct {
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
			OfflineInstances *[]string `tfsdk:"offline_instances" json:"offlineInstances,omitempty"`
			Replicas         *int64    `tfsdk:"replicas" json:"replicas,omitempty"`
		} `tfsdk:"horizontal_scaling" json:"horizontalScaling,omitempty"`
		RebuildFrom *[]struct {
			BackupName    *string            `tfsdk:"backup_name" json:"backupName,omitempty"`
			ComponentName *string            `tfsdk:"component_name" json:"componentName,omitempty"`
			EnvForRestore *map[string]string `tfsdk:"env_for_restore" json:"envForRestore,omitempty"`
			Instances     *[]struct {
				Name           *string `tfsdk:"name" json:"name,omitempty"`
				TargetNodeName *string `tfsdk:"target_node_name" json:"targetNodeName,omitempty"`
			} `tfsdk:"instances" json:"instances,omitempty"`
		} `tfsdk:"rebuild_from" json:"rebuildFrom,omitempty"`
		Reconfigure *struct {
			ComponentName  *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Configurations *[]struct {
				Keys *[]struct {
					FileContent *string `tfsdk:"file_content" json:"fileContent,omitempty"`
					Key         *string `tfsdk:"key" json:"key,omitempty"`
					Parameters  *[]struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
				} `tfsdk:"keys" json:"keys,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Policy *string `tfsdk:"policy" json:"policy,omitempty"`
			} `tfsdk:"configurations" json:"configurations,omitempty"`
		} `tfsdk:"reconfigure" json:"reconfigure,omitempty"`
		Reconfigures *[]struct {
			ComponentName  *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Configurations *[]struct {
				Keys *[]struct {
					FileContent *string `tfsdk:"file_content" json:"fileContent,omitempty"`
					Key         *string `tfsdk:"key" json:"key,omitempty"`
					Parameters  *[]struct {
						Key   *string `tfsdk:"key" json:"key,omitempty"`
						Value *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"parameters" json:"parameters,omitempty"`
				} `tfsdk:"keys" json:"keys,omitempty"`
				Name   *string `tfsdk:"name" json:"name,omitempty"`
				Policy *string `tfsdk:"policy" json:"policy,omitempty"`
			} `tfsdk:"configurations" json:"configurations,omitempty"`
		} `tfsdk:"reconfigures" json:"reconfigures,omitempty"`
		Restart *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
		} `tfsdk:"restart" json:"restart,omitempty"`
		RestoreFrom *struct {
			Backup *[]struct {
				Ref *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"ref" json:"ref,omitempty"`
			} `tfsdk:"backup" json:"backup,omitempty"`
			PointInTime *struct {
				Ref *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"ref" json:"ref,omitempty"`
				Time *string `tfsdk:"time" json:"time,omitempty"`
			} `tfsdk:"point_in_time" json:"pointInTime,omitempty"`
		} `tfsdk:"restore_from" json:"restoreFrom,omitempty"`
		RestoreSpec *struct {
			BackupName                        *string `tfsdk:"backup_name" json:"backupName,omitempty"`
			DoReadyRestoreAfterClusterRunning *bool   `tfsdk:"do_ready_restore_after_cluster_running" json:"doReadyRestoreAfterClusterRunning,omitempty"`
			RestoreTimeStr                    *string `tfsdk:"restore_time_str" json:"restoreTimeStr,omitempty"`
			VolumeRestorePolicy               *string `tfsdk:"volume_restore_policy" json:"volumeRestorePolicy,omitempty"`
		} `tfsdk:"restore_spec" json:"restoreSpec,omitempty"`
		ScriptSpec *struct {
			ComponentName *string   `tfsdk:"component_name" json:"componentName,omitempty"`
			Image         *string   `tfsdk:"image" json:"image,omitempty"`
			Script        *[]string `tfsdk:"script" json:"script,omitempty"`
			ScriptFrom    *struct {
				ConfigMapRef *[]struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"config_map_ref" json:"configMapRef,omitempty"`
				SecretRef *[]struct {
					Key      *string `tfsdk:"key" json:"key,omitempty"`
					Name     *string `tfsdk:"name" json:"name,omitempty"`
					Optional *bool   `tfsdk:"optional" json:"optional,omitempty"`
				} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
			} `tfsdk:"script_from" json:"scriptFrom,omitempty"`
			Secret *struct {
				Name        *string `tfsdk:"name" json:"name,omitempty"`
				PasswordKey *string `tfsdk:"password_key" json:"passwordKey,omitempty"`
				UsernameKey *string `tfsdk:"username_key" json:"usernameKey,omitempty"`
			} `tfsdk:"secret" json:"secret,omitempty"`
			Selector *struct {
				MatchExpressions *[]struct {
					Key      *string   `tfsdk:"key" json:"key,omitempty"`
					Operator *string   `tfsdk:"operator" json:"operator,omitempty"`
					Values   *[]string `tfsdk:"values" json:"values,omitempty"`
				} `tfsdk:"match_expressions" json:"matchExpressions,omitempty"`
				MatchLabels *map[string]string `tfsdk:"match_labels" json:"matchLabels,omitempty"`
			} `tfsdk:"selector" json:"selector,omitempty"`
		} `tfsdk:"script_spec" json:"scriptSpec,omitempty"`
		Switchover *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			InstanceName  *string `tfsdk:"instance_name" json:"instanceName,omitempty"`
		} `tfsdk:"switchover" json:"switchover,omitempty"`
		TtlSecondsAfterSucceed *int64  `tfsdk:"ttl_seconds_after_succeed" json:"ttlSecondsAfterSucceed,omitempty"`
		TtlSecondsBeforeAbort  *int64  `tfsdk:"ttl_seconds_before_abort" json:"ttlSecondsBeforeAbort,omitempty"`
		Type                   *string `tfsdk:"type" json:"type,omitempty"`
		Upgrade                *struct {
			ClusterVersionRef *string `tfsdk:"cluster_version_ref" json:"clusterVersionRef,omitempty"`
		} `tfsdk:"upgrade" json:"upgrade,omitempty"`
		VerticalScaling *[]map[string]string `tfsdk:"vertical_scaling" json:"verticalScaling,omitempty"`
		VolumeExpansion *[]struct {
			ComponentName *string `tfsdk:"component_name" json:"componentName,omitempty"`
			Instances     *[]struct {
				Name                 *string `tfsdk:"name" json:"name,omitempty"`
				VolumeClaimTemplates *[]struct {
					Name    *string `tfsdk:"name" json:"name,omitempty"`
					Storage *string `tfsdk:"storage" json:"storage,omitempty"`
				} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
			} `tfsdk:"instances" json:"instances,omitempty"`
			VolumeClaimTemplates *[]struct {
				Name    *string `tfsdk:"name" json:"name,omitempty"`
				Storage *string `tfsdk:"storage" json:"storage,omitempty"`
			} `tfsdk:"volume_claim_templates" json:"volumeClaimTemplates,omitempty"`
		} `tfsdk:"volume_expansion" json:"volumeExpansion,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *AppsKubeblocksIoOpsRequestV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_apps_kubeblocks_io_ops_request_v1alpha1_manifest"
}

func (r *AppsKubeblocksIoOpsRequestV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "OpsRequest is the Schema for the opsrequests API",
		MarkdownDescription: "OpsRequest is the Schema for the opsrequests API",
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
				Description:         "OpsRequestSpec defines the desired state of OpsRequest",
				MarkdownDescription: "OpsRequestSpec defines the desired state of OpsRequest",
				Attributes: map[string]schema.Attribute{
					"backup_spec": schema.SingleNestedAttribute{
						Description:         "Specifies the parameters to backup a Cluster.",
						MarkdownDescription: "Specifies the parameters to backup a Cluster.",
						Attributes: map[string]schema.Attribute{
							"backup_method": schema.StringAttribute{
								Description:         "Specifies the name of BackupMethod. The specified BackupMethod must be defined in the BackupPolicy.",
								MarkdownDescription: "Specifies the name of BackupMethod. The specified BackupMethod must be defined in the BackupPolicy.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"backup_name": schema.StringAttribute{
								Description:         "Specifies the name of the Backup custom resource.",
								MarkdownDescription: "Specifies the name of the Backup custom resource.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"backup_policy_name": schema.StringAttribute{
								Description:         "Indicates the name of the BackupPolicy applied to perform this Backup.",
								MarkdownDescription: "Indicates the name of the BackupPolicy applied to perform this Backup.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"deletion_policy": schema.StringAttribute{
								Description:         "Determines whether the backup contents stored in backup repository should be deleted when the Backup custom resource is deleted. Supported values are 'Retain' and 'Delete'. - 'Retain' means that the backup content and its physical snapshot on backup repository are kept. - 'Delete' means that the backup content and its physical snapshot on backup repository are deleted.",
								MarkdownDescription: "Determines whether the backup contents stored in backup repository should be deleted when the Backup custom resource is deleted. Supported values are 'Retain' and 'Delete'. - 'Retain' means that the backup content and its physical snapshot on backup repository are kept. - 'Delete' means that the backup content and its physical snapshot on backup repository are deleted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Delete", "Retain"),
								},
							},

							"parent_backup_name": schema.StringAttribute{
								Description:         "If the specified BackupMethod is incremental, 'parentBackupName' is required.",
								MarkdownDescription: "If the specified BackupMethod is incremental, 'parentBackupName' is required.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"retention_period": schema.StringAttribute{
								Description:         "Determines the duration for which the Backup custom resources should be retained.  The controller will automatically remove all Backup objects that are older than the specified RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the Backup objects of last 30 days. Sample duration format:  - years: 2y - months: 6mo - days: 30d - hours: 12h - minutes: 30m  You can also combine the above durations. For example: 30d12h30m. If not set, the Backup objects will be kept forever.  If the 'deletionPolicy' is set to 'Delete', then the associated backup data will also be deleted along with the Backup object. Otherwise, only the Backup custom resource will be deleted.",
								MarkdownDescription: "Determines the duration for which the Backup custom resources should be retained.  The controller will automatically remove all Backup objects that are older than the specified RetentionPeriod. For example, RetentionPeriod of '30d' will keep only the Backup objects of last 30 days. Sample duration format:  - years: 2y - months: 6mo - days: 30d - hours: 12h - minutes: 30m  You can also combine the above durations. For example: 30d12h30m. If not set, the Backup objects will be kept forever.  If the 'deletionPolicy' is set to 'Delete', then the associated backup data will also be deleted along with the Backup object. Otherwise, only the Backup custom resource will be deleted.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"cancel": schema.BoolAttribute{
						Description:         "Indicates whether the current operation should be canceled and terminated gracefully if it's in the 'Pending', 'Creating', or 'Running' state.  This field applies only to 'VerticalScaling' and 'HorizontalScaling' opsRequests.  Note: Setting 'cancel' to true is irreversible; further modifications to this field are ineffective.",
						MarkdownDescription: "Indicates whether the current operation should be canceled and terminated gracefully if it's in the 'Pending', 'Creating', or 'Running' state.  This field applies only to 'VerticalScaling' and 'HorizontalScaling' opsRequests.  Note: Setting 'cancel' to true is irreversible; further modifications to this field are ineffective.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"cluster_ref": schema.StringAttribute{
						Description:         "Specifies the name of the Cluster resource that this operation is targeting.",
						MarkdownDescription: "Specifies the name of the Cluster resource that this operation is targeting.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"custom_spec": schema.SingleNestedAttribute{
						Description:         "Specifies a custom operation defined by OpsDefinition.",
						MarkdownDescription: "Specifies a custom operation defined by OpsDefinition.",
						Attributes: map[string]schema.Attribute{
							"components": schema.ListNestedAttribute{
								Description:         "Specifies the components and their parameters for executing custom actions as defined in OpsDefinition. Requires at least one component.",
								MarkdownDescription: "Specifies the components and their parameters for executing custom actions as defined in OpsDefinition. Requires at least one component.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"component_name": schema.StringAttribute{
											Description:         "Specifies the name of the Component.",
											MarkdownDescription: "Specifies the name of the Component.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"parameters": schema.ListNestedAttribute{
											Description:         "Specifies the parameters that match the schema specified in the 'opsDefinition.spec.parametersSchema'.",
											MarkdownDescription: "Specifies the parameters that match the schema specified in the 'opsDefinition.spec.parametersSchema'.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "Specifies the identifier of the parameter as defined in the OpsDefinition.",
														MarkdownDescription: "Specifies the identifier of the parameter as defined in the OpsDefinition.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"value": schema.StringAttribute{
														Description:         "Holds the data associated with the parameter. If the parameter type is an array, the format should be 'v1,v2,v3'.",
														MarkdownDescription: "Holds the data associated with the parameter. If the parameter type is an array, the format should be 'v1,v2,v3'.",
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
								},
								Required: true,
								Optional: false,
								Computed: false,
							},

							"ops_definition_ref": schema.StringAttribute{
								Description:         "Specifies the name of the OpsDefinition.",
								MarkdownDescription: "Specifies the name of the OpsDefinition.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"parallelism": schema.StringAttribute{
								Description:         "Specifies the maximum number of components to be operated on concurrently to mitigate performance impact on clusters with multiple components.  It accepts an absolute number (e.g., 5) or a percentage of components to execute in parallel (e.g., '10%'). Percentages are rounded up to the nearest whole number of components. For example, if '10%' results in less than one, it rounds up to 1.  When unspecified, all components are processed simultaneously by default.  Note: This feature is not implemented yet.",
								MarkdownDescription: "Specifies the maximum number of components to be operated on concurrently to mitigate performance impact on clusters with multiple components.  It accepts an absolute number (e.g., 5) or a percentage of components to execute in parallel (e.g., '10%'). Percentages are rounded up to the nearest whole number of components. For example, if '10%' results in less than one, it rounds up to 1.  When unspecified, all components are processed simultaneously by default.  Note: This feature is not implemented yet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"service_account_name": schema.StringAttribute{
								Description:         "Specifies the name of the ServiceAccount to be used for executing the custom operation.",
								MarkdownDescription: "Specifies the name of the ServiceAccount to be used for executing the custom operation.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"expose": schema.ListNestedAttribute{
						Description:         "Lists Expose objects, each specifying a Component and its services to be exposed.",
						MarkdownDescription: "Lists Expose objects, each specifying a Component and its services to be exposed.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"services": schema.ListNestedAttribute{
									Description:         "Specifies a list of OpsService. When an OpsService is exposed, a corresponding ClusterService will be added to 'cluster.spec.services'. On the other hand, when an OpsService is unexposed, the corresponding ClusterService will be removed from 'cluster.spec.services'.  Note: If 'componentName' is not specified, the 'ports' and 'selector' fields must be provided in each OpsService definition.",
									MarkdownDescription: "Specifies a list of OpsService. When an OpsService is exposed, a corresponding ClusterService will be added to 'cluster.spec.services'. On the other hand, when an OpsService is unexposed, the corresponding ClusterService will be removed from 'cluster.spec.services'.  Note: If 'componentName' is not specified, the 'ports' and 'selector' fields must be provided in each OpsService definition.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"annotations": schema.MapAttribute{
												Description:         "Contains cloud provider related parameters if ServiceType is LoadBalancer.  More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												MarkdownDescription: "Contains cloud provider related parameters if ServiceType is LoadBalancer.  More info: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_families": schema.ListAttribute{
												Description:         "A list of IP families (e.g., IPv4, IPv6) assigned to this Service.  Usually assigned automatically based on the cluster configuration and the 'ipFamilyPolicy' field. If specified manually, the requested IP family must be available in the cluster and allowed by the 'ipFamilyPolicy'. If the requested IP family is not available or not allowed, the Service creation will fail.  Valid values:  - 'IPv4' - 'IPv6'  This field may hold a maximum of two entries (dual-stack families, in either order).  Common combinations of 'ipFamilies' and 'ipFamilyPolicy' are:  - ipFamilies=[] + ipFamilyPolicy='PreferDualStack' : The Service prefers dual-stack but can fall back to single-stack if the cluster does not support dual-stack. The IP family is automatically assigned based on the cluster configuration. - ipFamilies=['IPV4','IPV6'] + ipFamilyPolicy='RequiredDualStack' : The Service requires dual-stack and will only be created if the cluster supports both IPv4 and IPv6. The primary IP family is IPV4. - ipFamilies=['IPV6','IPV4'] + ipFamilyPolicy='RequiredDualStack' : The Service requires dual-stack and will only be created if the cluster supports both IPv4 and IPv6. The primary IP family is IPV6. - ipFamilies=['IPV4'] + ipFamilyPolicy='SingleStack' : The Service uses a single-stack with IPv4 only. - ipFamilies=['IPV6'] + ipFamilyPolicy='SingleStack' : The Service uses a single-stack with IPv6 only.",
												MarkdownDescription: "A list of IP families (e.g., IPv4, IPv6) assigned to this Service.  Usually assigned automatically based on the cluster configuration and the 'ipFamilyPolicy' field. If specified manually, the requested IP family must be available in the cluster and allowed by the 'ipFamilyPolicy'. If the requested IP family is not available or not allowed, the Service creation will fail.  Valid values:  - 'IPv4' - 'IPv6'  This field may hold a maximum of two entries (dual-stack families, in either order).  Common combinations of 'ipFamilies' and 'ipFamilyPolicy' are:  - ipFamilies=[] + ipFamilyPolicy='PreferDualStack' : The Service prefers dual-stack but can fall back to single-stack if the cluster does not support dual-stack. The IP family is automatically assigned based on the cluster configuration. - ipFamilies=['IPV4','IPV6'] + ipFamilyPolicy='RequiredDualStack' : The Service requires dual-stack and will only be created if the cluster supports both IPv4 and IPv6. The primary IP family is IPV4. - ipFamilies=['IPV6','IPV4'] + ipFamilyPolicy='RequiredDualStack' : The Service requires dual-stack and will only be created if the cluster supports both IPv4 and IPv6. The primary IP family is IPV6. - ipFamilies=['IPV4'] + ipFamilyPolicy='SingleStack' : The Service uses a single-stack with IPv4 only. - ipFamilies=['IPV6'] + ipFamilyPolicy='SingleStack' : The Service uses a single-stack with IPv6 only.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"ip_family_policy": schema.StringAttribute{
												Description:         "Specifies whether the Service should use a single IP family (SingleStack) or two IP families (DualStack).  Possible values:  - 'SingleStack' (default) : The Service uses a single IP family. If no value is provided, IPFamilyPolicy defaults to SingleStack. - 'PreferDualStack' : The Service prefers to use two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. - 'RequiredDualStack' : The Service requires two IP families on dual-stack configured clusters. If the cluster is not configured for dual-stack, the Service creation fails.",
												MarkdownDescription: "Specifies whether the Service should use a single IP family (SingleStack) or two IP families (DualStack).  Possible values:  - 'SingleStack' (default) : The Service uses a single IP family. If no value is provided, IPFamilyPolicy defaults to SingleStack. - 'PreferDualStack' : The Service prefers to use two IP families on dual-stack configured clusters or a single IP family on single-stack clusters. - 'RequiredDualStack' : The Service requires two IP families on dual-stack configured clusters. If the cluster is not configured for dual-stack, the Service creation fails.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Specifies the name of the Service. This name is used to set 'clusterService.name'.  Note: This field cannot be updated.",
												MarkdownDescription: "Specifies the name of the Service. This name is used to set 'clusterService.name'.  Note: This field cannot be updated.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"ports": schema.ListNestedAttribute{
												Description:         "Specifies Port definitions that are to be exposed by a ClusterService.  If not specified, the Port definitions from non-NodePort and non-LoadBalancer type ComponentService defined in the ComponentDefinition ('componentDefinition.spec.services') will be used. If no matching ComponentService is found, the expose operation will fail.  More info: https://kubernetes.io/docs/concepts/services-networking/service/#field-spec-ports",
												MarkdownDescription: "Specifies Port definitions that are to be exposed by a ClusterService.  If not specified, the Port definitions from non-NodePort and non-LoadBalancer type ComponentService defined in the ComponentDefinition ('componentDefinition.spec.services') will be used. If no matching ComponentService is found, the expose operation will fail.  More info: https://kubernetes.io/docs/concepts/services-networking/service/#field-spec-ports",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"app_protocol": schema.StringAttribute{
															Description:         "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:  * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).  * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540 * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455  * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
															MarkdownDescription: "The application protocol for this port. This is used as a hint for implementations to offer richer behavior for protocols that they understand. This field follows standard Kubernetes label syntax. Valid values are either:  * Un-prefixed protocol names - reserved for IANA standard service names (as per RFC-6335 and https://www.iana.org/assignments/service-names).  * Kubernetes-defined prefixed names: * 'kubernetes.io/h2c' - HTTP/2 over cleartext as described in https://www.rfc-editor.org/rfc/rfc7540 * 'kubernetes.io/ws'  - WebSocket over cleartext as described in https://www.rfc-editor.org/rfc/rfc6455 * 'kubernetes.io/wss' - WebSocket over TLS as described in https://www.rfc-editor.org/rfc/rfc6455  * Other protocols should use implementation-defined prefixed names such as mycompany.com/my-custom-protocol.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"name": schema.StringAttribute{
															Description:         "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
															MarkdownDescription: "The name of this port within the service. This must be a DNS_LABEL. All ports within a ServiceSpec must have unique names. When considering the endpoints for a Service, this must match the 'name' field in the EndpointPort. Optional if only one ServicePort is defined on this service.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"node_port": schema.Int64Attribute{
															Description:         "The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
															MarkdownDescription: "The port on each node on which this service is exposed when type is NodePort or LoadBalancer.  Usually assigned by the system. If a value is specified, in-range, and not in use it will be used, otherwise the operation will fail.  If not specified, a port will be allocated if this Service requires one.  If this field is specified when creating a Service which does not need it, creation will fail. This field will be wiped when updating a Service to no longer need it (e.g. changing type from NodePort to ClusterIP). More info: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport",
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
															Description:         "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
															MarkdownDescription: "The IP protocol for this port. Supports 'TCP', 'UDP', and 'SCTP'. Default is TCP.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"target_port": schema.StringAttribute{
															Description:         "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
															MarkdownDescription: "Number or name of the port to access on the pods targeted by the service. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME. If this is a string, it will be looked up as a named port in the target Pod's container ports. If this is not specified, the value of the 'port' field is used (an identity map). This field is ignored for services with clusterIP=None, and should be omitted or set equal to the 'port' field. More info: https://kubernetes.io/docs/concepts/services-networking/service/#defining-a-service",
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

											"role_selector": schema.StringAttribute{
												Description:         "Specifies a role to target with the service. If specified, the service will only be exposed to pods with the matching role.  Note: At least one of 'roleSelector' or 'selector' must be specified. If both are specified, a pod must match both conditions to be selected.",
												MarkdownDescription: "Specifies a role to target with the service. If specified, the service will only be exposed to pods with the matching role.  Note: At least one of 'roleSelector' or 'selector' must be specified. If both are specified, a pod must match both conditions to be selected.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"selector": schema.MapAttribute{
												Description:         "Routes service traffic to pods with matching label keys and values. If specified, the service will only be exposed to pods matching the selector.  Note: At least one of 'roleSelector' or 'selector' must be specified. If both are specified, a pod must match both conditions to be selected.",
												MarkdownDescription: "Routes service traffic to pods with matching label keys and values. If specified, the service will only be exposed to pods matching the selector.  Note: At least one of 'roleSelector' or 'selector' must be specified. If both are specified, a pod must match both conditions to be selected.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"service_type": schema.StringAttribute{
												Description:         "Determines how the Service is exposed. Defaults to 'ClusterIP'. Valid options are 'ClusterIP', 'NodePort', and 'LoadBalancer'.  - 'ClusterIP': allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. - 'NodePort': builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer': builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP.  Note: although K8s Service type allows the 'ExternalName' type, it is not a valid option for the expose operation.  For more info, see: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
												MarkdownDescription: "Determines how the Service is exposed. Defaults to 'ClusterIP'. Valid options are 'ClusterIP', 'NodePort', and 'LoadBalancer'.  - 'ClusterIP': allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, they are determined by manual construction of an Endpoints object or EndpointSlice objects. - 'NodePort': builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. - 'LoadBalancer': builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP.  Note: although K8s Service type allows the 'ExternalName' type, it is not a valid option for the expose operation.  For more info, see: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
									},
									Required: true,
									Optional: false,
									Computed: false,
								},

								"switch": schema.StringAttribute{
									Description:         "Indicates whether the services will be exposed. 'Enable' exposes the services. while 'Disable' removes the exposed Service.",
									MarkdownDescription: "Indicates whether the services will be exposed. 'Enable' exposes the services. while 'Disable' removes the exposed Service.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.String{
										stringvalidator.OneOf("Enable", "Disable"),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"force": schema.BoolAttribute{
						Description:         "Instructs the system to bypass pre-checks (including cluster state checks and customized pre-conditions hooks) and immediately execute the opsRequest, except for the opsRequest of 'Start' type, which will still undergo pre-checks even if 'force' is true.  This is useful for concurrent execution of 'VerticalScaling' and 'HorizontalScaling' opsRequests. By setting 'force' to true, you can bypass the default checks and demand these opsRequests to run simultaneously.  Note: Once set, the 'force' field is immutable and cannot be updated.",
						MarkdownDescription: "Instructs the system to bypass pre-checks (including cluster state checks and customized pre-conditions hooks) and immediately execute the opsRequest, except for the opsRequest of 'Start' type, which will still undergo pre-checks even if 'force' is true.  This is useful for concurrent execution of 'VerticalScaling' and 'HorizontalScaling' opsRequests. By setting 'force' to true, you can bypass the default checks and demand these opsRequests to run simultaneously.  Note: Once set, the 'force' field is immutable and cannot be updated.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"horizontal_scaling": schema.ListNestedAttribute{
						Description:         "Lists HorizontalScaling objects, each specifying scaling requirements for a Component, including desired total replica counts, configurations for new instances, modifications for existing instances, and instance downscaling options.",
						MarkdownDescription: "Lists HorizontalScaling objects, each specifying scaling requirements for a Component, including desired total replica counts, configurations for new instances, modifications for existing instances, and instance downscaling options.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"instances": schema.ListNestedAttribute{
									Description:         "Contains a list of InstanceTemplate objects. Each InstanceTemplate object allows for modifying replica counts or specifying configurations for new instances during scaling.  The field supports two main use cases:  - Modifying replica count: Specify the desired replica count for existing instances with a particular configuration using Name and Replicas fields. To modify the replica count, the Name and Replicas fields of the InstanceTemplate object should be provided. Only these fields are used for matching and adjusting replicas; other fields are ignored. The Replicas value overrides any existing count. - Configuring new instances: Define the configuration for new instances added during scaling, including resource requirements, labels, annotations, etc. New instances are created based on the provided InstanceTemplate.",
									MarkdownDescription: "Contains a list of InstanceTemplate objects. Each InstanceTemplate object allows for modifying replica counts or specifying configurations for new instances during scaling.  The field supports two main use cases:  - Modifying replica count: Specify the desired replica count for existing instances with a particular configuration using Name and Replicas fields. To modify the replica count, the Name and Replicas fields of the InstanceTemplate object should be provided. Only these fields are used for matching and adjusting replicas; other fields are ignored. The Replicas value overrides any existing count. - Configuring new instances: Define the configuration for new instances added during scaling, including resource requirements, labels, annotations, etc. New instances are created based on the provided InstanceTemplate.",
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

								"offline_instances": schema.ListAttribute{
									Description:         "Specifies the names of instances to be scaled down. This provides control over which specific instances are targeted for termination when reducing the replica count.",
									MarkdownDescription: "Specifies the names of instances to be scaled down. This provides control over which specific instances are targeted for termination when reducing the replica count.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"replicas": schema.Int64Attribute{
									Description:         "Specifies the number of total replicas.",
									MarkdownDescription: "Specifies the number of total replicas.",
									Required:            true,
									Optional:            false,
									Computed:            false,
									Validators: []validator.Int64{
										int64validator.AtLeast(0),
									},
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"rebuild_from": schema.ListNestedAttribute{
						Description:         "Specifies the parameters to rebuild some instances. Rebuilding an instance involves restoring its data from a backup or another database replica. The instances being rebuilt usually serve as standby in the cluster. Hence rebuilding instances is often also referred to as 'standby reconstruction'.",
						MarkdownDescription: "Specifies the parameters to rebuild some instances. Rebuilding an instance involves restoring its data from a backup or another database replica. The instances being rebuilt usually serve as standby in the cluster. Hence rebuilding instances is often also referred to as 'standby reconstruction'.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"backup_name": schema.StringAttribute{
									Description:         "Indicates the name of the Backup custom resource from which to recover the instance. Defaults to an empty PersistentVolume if unspecified.  Note: - Only full physical backups are supported for multi-replica Components (e.g., 'xtrabackup' for MySQL). - Logical backups (e.g., 'mysqldump' for MySQL) are unsupported in the current version.",
									MarkdownDescription: "Indicates the name of the Backup custom resource from which to recover the instance. Defaults to an empty PersistentVolume if unspecified.  Note: - Only full physical backups are supported for multi-replica Components (e.g., 'xtrabackup' for MySQL). - Logical backups (e.g., 'mysqldump' for MySQL) are unsupported in the current version.",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"env_for_restore": schema.MapAttribute{
									Description:         "Defines container environment variables for the restore process. merged with the ones specified in the Backup and ActionSet resources.  Merge priority: Restore env > Backup env > ActionSet env.  Purpose: Some databases require different configurations when being restored as a standby compared to being restored as a primary. For example, when restoring MySQL as a replica, you need to set 'skip_slave_start='ON'' for 5.7 or 'skip_replica_start='ON'' for 8.0. Allowing environment variables to be passed in makes it more convenient to control these behavioral differences during the restore process.",
									MarkdownDescription: "Defines container environment variables for the restore process. merged with the ones specified in the Backup and ActionSet resources.  Merge priority: Restore env > Backup env > ActionSet env.  Purpose: Some databases require different configurations when being restored as a standby compared to being restored as a primary. For example, when restoring MySQL as a replica, you need to set 'skip_slave_start='ON'' for 5.7 or 'skip_replica_start='ON'' for 8.0. Allowing environment variables to be passed in makes it more convenient to control these behavioral differences during the restore process.",
									ElementType:         types.StringType,
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"instances": schema.ListNestedAttribute{
									Description:         "Specifies the instances (Pods) that need to be rebuilt, typically operating as standbys.",
									MarkdownDescription: "Specifies the instances (Pods) that need to be rebuilt, typically operating as standbys.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Pod name of the instance.",
												MarkdownDescription: "Pod name of the instance.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"target_node_name": schema.StringAttribute{
												Description:         "The instance will rebuild on the specified node when the instance uses local PersistentVolume as the storage disk. If not set, it will rebuild on a random node.",
												MarkdownDescription: "The instance will rebuild on the specified node when the instance uses local PersistentVolume as the storage disk. If not set, it will rebuild on a random node.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
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

					"reconfigure": schema.SingleNestedAttribute{
						Description:         "Specifies a component and its configuration updates.  This field is deprecated and replaced by 'reconfigures'.",
						MarkdownDescription: "Specifies a component and its configuration updates.  This field is deprecated and replaced by 'reconfigures'.",
						Attributes: map[string]schema.Attribute{
							"component_name": schema.StringAttribute{
								Description:         "Specifies the name of the Component.",
								MarkdownDescription: "Specifies the name of the Component.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"configurations": schema.ListNestedAttribute{
								Description:         "Contains a list of ConfigurationItem objects, specifying the Component's configuration template name, upgrade policy, and parameter key-value pairs to be updated.",
								MarkdownDescription: "Contains a list of ConfigurationItem objects, specifying the Component's configuration template name, upgrade policy, and parameter key-value pairs to be updated.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"keys": schema.ListNestedAttribute{
											Description:         "Sets the configuration files and their associated parameters that need to be updated. It should contain at least one item.",
											MarkdownDescription: "Sets the configuration files and their associated parameters that need to be updated. It should contain at least one item.",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"file_content": schema.StringAttribute{
														Description:         "Specifies the content of the entire configuration file. This field is used to update the complete configuration file.  Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
														MarkdownDescription: "Specifies the content of the entire configuration file. This field is used to update the complete configuration file.  Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key": schema.StringAttribute{
														Description:         "Represents a key in the configuration template(as ConfigMap). Each key in the ConfigMap corresponds to a specific configuration file.",
														MarkdownDescription: "Represents a key in the configuration template(as ConfigMap). Each key in the ConfigMap corresponds to a specific configuration file.",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"parameters": schema.ListNestedAttribute{
														Description:         "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file.  Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
														MarkdownDescription: "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file.  Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
														NestedObject: schema.NestedAttributeObject{
															Attributes: map[string]schema.Attribute{
																"key": schema.StringAttribute{
																	Description:         "Represents the name of the parameter that is to be updated.",
																	MarkdownDescription: "Represents the name of the parameter that is to be updated.",
																	Required:            true,
																	Optional:            false,
																	Computed:            false,
																},

																"value": schema.StringAttribute{
																	Description:         "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
																	MarkdownDescription: "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
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

										"name": schema.StringAttribute{
											Description:         "Specifies the name of the configuration template.",
											MarkdownDescription: "Specifies the name of the configuration template.",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.LengthAtMost(63),
												stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
											},
										},

										"policy": schema.StringAttribute{
											Description:         "Defines the upgrade policy for the configuration.",
											MarkdownDescription: "Defines the upgrade policy for the configuration.",
											Required:            false,
											Optional:            true,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.OneOf("simple", "parallel", "rolling", "autoReload", "operatorSyncUpdate", "dynamicReloadBeginRestart"),
											},
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

					"reconfigures": schema.ListNestedAttribute{
						Description:         "Lists Reconfigure objects, each specifying a Component and its configuration updates.",
						MarkdownDescription: "Lists Reconfigure objects, each specifying a Component and its configuration updates.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"configurations": schema.ListNestedAttribute{
									Description:         "Contains a list of ConfigurationItem objects, specifying the Component's configuration template name, upgrade policy, and parameter key-value pairs to be updated.",
									MarkdownDescription: "Contains a list of ConfigurationItem objects, specifying the Component's configuration template name, upgrade policy, and parameter key-value pairs to be updated.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"keys": schema.ListNestedAttribute{
												Description:         "Sets the configuration files and their associated parameters that need to be updated. It should contain at least one item.",
												MarkdownDescription: "Sets the configuration files and their associated parameters that need to be updated. It should contain at least one item.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"file_content": schema.StringAttribute{
															Description:         "Specifies the content of the entire configuration file. This field is used to update the complete configuration file.  Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
															MarkdownDescription: "Specifies the content of the entire configuration file. This field is used to update the complete configuration file.  Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"key": schema.StringAttribute{
															Description:         "Represents a key in the configuration template(as ConfigMap). Each key in the ConfigMap corresponds to a specific configuration file.",
															MarkdownDescription: "Represents a key in the configuration template(as ConfigMap). Each key in the ConfigMap corresponds to a specific configuration file.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"parameters": schema.ListNestedAttribute{
															Description:         "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file.  Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
															MarkdownDescription: "Specifies a list of key-value pairs representing parameters and their corresponding values within a single configuration file. This field is used to override or set the values of parameters without modifying the entire configuration file.  Either the 'parameters' field or the 'fileContent' field must be set, but not both.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"key": schema.StringAttribute{
																		Description:         "Represents the name of the parameter that is to be updated.",
																		MarkdownDescription: "Represents the name of the parameter that is to be updated.",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
																		MarkdownDescription: "Represents the parameter values that are to be updated. If set to nil, the parameter defined by the Key field will be removed from the configuration file.",
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

											"name": schema.StringAttribute{
												Description:         "Specifies the name of the configuration template.",
												MarkdownDescription: "Specifies the name of the configuration template.",
												Required:            true,
												Optional:            false,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtMost(63),
													stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
												},
											},

											"policy": schema.StringAttribute{
												Description:         "Defines the upgrade policy for the configuration.",
												MarkdownDescription: "Defines the upgrade policy for the configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.OneOf("simple", "parallel", "rolling", "autoReload", "operatorSyncUpdate", "dynamicReloadBeginRestart"),
												},
											},
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

					"restart": schema.ListNestedAttribute{
						Description:         "Lists Components to be restarted.",
						MarkdownDescription: "Lists Components to be restarted.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
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

					"restore_from": schema.SingleNestedAttribute{
						Description:         "Cluster RestoreFrom backup or point in time.",
						MarkdownDescription: "Cluster RestoreFrom backup or point in time.",
						Attributes: map[string]schema.Attribute{
							"backup": schema.ListNestedAttribute{
								Description:         "Refers to the backup name and component name used for restoration. Supports recovery of multiple Components.",
								MarkdownDescription: "Refers to the backup name and component name used for restoration. Supports recovery of multiple Components.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"ref": schema.SingleNestedAttribute{
											Description:         "Refers to a reference backup that needs to be restored.",
											MarkdownDescription: "Refers to a reference backup that needs to be restored.",
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Description:         "Refers to the specific name of the resource.",
													MarkdownDescription: "Refers to the specific name of the resource.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"namespace": schema.StringAttribute{
													Description:         "Refers to the specific namespace of the resource.",
													MarkdownDescription: "Refers to the specific namespace of the resource.",
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

							"point_in_time": schema.SingleNestedAttribute{
								Description:         "Refers to the specific point in time for recovery.",
								MarkdownDescription: "Refers to the specific point in time for recovery.",
								Attributes: map[string]schema.Attribute{
									"ref": schema.SingleNestedAttribute{
										Description:         "Refers to a reference source cluster that needs to be restored.",
										MarkdownDescription: "Refers to a reference source cluster that needs to be restored.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Refers to the specific name of the resource.",
												MarkdownDescription: "Refers to the specific name of the resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Refers to the specific namespace of the resource.",
												MarkdownDescription: "Refers to the specific namespace of the resource.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"time": schema.StringAttribute{
										Description:         "Refers to the specific time point for restoration, with UTC as the time zone.",
										MarkdownDescription: "Refers to the specific time point for restoration, with UTC as the time zone.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											validators.DateTime64Validator(),
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

					"restore_spec": schema.SingleNestedAttribute{
						Description:         "Specifies the parameters to restore a Cluster. Note that this restore operation will roll back cluster services.",
						MarkdownDescription: "Specifies the parameters to restore a Cluster. Note that this restore operation will roll back cluster services.",
						Attributes: map[string]schema.Attribute{
							"backup_name": schema.StringAttribute{
								Description:         "Specifies the name of the Backup custom resource.",
								MarkdownDescription: "Specifies the name of the Backup custom resource.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"do_ready_restore_after_cluster_running": schema.BoolAttribute{
								Description:         "Controls the timing of PostReady actions during the recovery process.  If false (default), PostReady actions execute when the Component reaches the 'Running' state. If true, PostReady actions are delayed until the entire Cluster is 'Running,' ensuring the cluster's overall stability before proceeding.  This setting is useful for coordinating PostReady operations across the Cluster for optimal cluster conditions.",
								MarkdownDescription: "Controls the timing of PostReady actions during the recovery process.  If false (default), PostReady actions execute when the Component reaches the 'Running' state. If true, PostReady actions are delayed until the entire Cluster is 'Running,' ensuring the cluster's overall stability before proceeding.  This setting is useful for coordinating PostReady operations across the Cluster for optimal cluster conditions.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"restore_time_str": schema.StringAttribute{
								Description:         "Specifies the point in time to which the restore should be performed. Supported time formats:  - RFC3339 format, e.g. '2023-11-25T18:52:53Z' - A human-readable date-time format, e.g. 'Jul 25,2023 18:52:53 UTC+0800'",
								MarkdownDescription: "Specifies the point in time to which the restore should be performed. Supported time formats:  - RFC3339 format, e.g. '2023-11-25T18:52:53Z' - A human-readable date-time format, e.g. 'Jul 25,2023 18:52:53 UTC+0800'",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"volume_restore_policy": schema.StringAttribute{
								Description:         "Specifies the policy for restoring volume claims of a Component's Pods. It determines whether the volume claims should be restored sequentially (one by one) or in parallel (all at once). Support values:  - 'Serial' - 'Parallel'",
								MarkdownDescription: "Specifies the policy for restoring volume claims of a Component's Pods. It determines whether the volume claims should be restored sequentially (one by one) or in parallel (all at once). Support values:  - 'Serial' - 'Parallel'",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.OneOf("Serial", "Parallel"),
								},
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"script_spec": schema.SingleNestedAttribute{
						Description:         "Specifies the image and scripts for executing engine-specific operations such as creating databases or users. It supports limited engines including MySQL, PostgreSQL, Redis, MongoDB.  ScriptSpec has been replaced by the more versatile OpsDefinition. It is recommended to use OpsDefinition instead. ScriptSpec is deprecated and will be removed in a future version.",
						MarkdownDescription: "Specifies the image and scripts for executing engine-specific operations such as creating databases or users. It supports limited engines including MySQL, PostgreSQL, Redis, MongoDB.  ScriptSpec has been replaced by the more versatile OpsDefinition. It is recommended to use OpsDefinition instead. ScriptSpec is deprecated and will be removed in a future version.",
						Attributes: map[string]schema.Attribute{
							"component_name": schema.StringAttribute{
								Description:         "Specifies the name of the Component.",
								MarkdownDescription: "Specifies the name of the Component.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},

							"image": schema.StringAttribute{
								Description:         "Specifies the image to be used to execute scripts.  By default, the image 'apecloud/kubeblocks-datascript:latest' is used.",
								MarkdownDescription: "Specifies the image to be used to execute scripts.  By default, the image 'apecloud/kubeblocks-datascript:latest' is used.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"script": schema.ListAttribute{
								Description:         "Defines the content of scripts to be executed.  All scripts specified in this field will be executed in the order they are provided.  Note: this field cannot be modified once set.",
								MarkdownDescription: "Defines the content of scripts to be executed.  All scripts specified in this field will be executed in the order they are provided.  Note: this field cannot be modified once set.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"script_from": schema.SingleNestedAttribute{
								Description:         "Specifies the sources of the scripts to be executed. Each script can be imported either from a ConfigMap or a Secret.  All scripts obtained from the sources specified in this field will be executed after any scripts provided in the 'script' field.  Execution order: 1. Scripts provided in the 'script' field, in the order of the scripts listed. 2. Scripts imported from ConfigMaps, in the order of the sources listed. 3. Scripts imported from Secrets, in the order of the sources listed.  Note: this field cannot be modified once set.",
								MarkdownDescription: "Specifies the sources of the scripts to be executed. Each script can be imported either from a ConfigMap or a Secret.  All scripts obtained from the sources specified in this field will be executed after any scripts provided in the 'script' field.  Execution order: 1. Scripts provided in the 'script' field, in the order of the scripts listed. 2. Scripts imported from ConfigMaps, in the order of the sources listed. 3. Scripts imported from Secrets, in the order of the sources listed.  Note: this field cannot be modified once set.",
								Attributes: map[string]schema.Attribute{
									"config_map_ref": schema.ListNestedAttribute{
										Description:         "A list of ConfigMapKeySelector objects, each specifies a ConfigMap and a key containing the script.  Note: This field cannot be modified once set.",
										MarkdownDescription: "A list of ConfigMapKeySelector objects, each specifies a ConfigMap and a key containing the script.  Note: This field cannot be modified once set.",
										NestedObject: schema.NestedAttributeObject{
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
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"secret_ref": schema.ListNestedAttribute{
										Description:         "A list of SecretKeySelector objects, each specifies a Secret and a key containing the script.  Note: This field cannot be modified once set.",
										MarkdownDescription: "A list of SecretKeySelector objects, each specifies a Secret and a key containing the script.  Note: This field cannot be modified once set.",
										NestedObject: schema.NestedAttributeObject{
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
								Description:         "Defines the secret to be used to execute the script. If not specified, the default cluster root credential secret is used.",
								MarkdownDescription: "Defines the secret to be used to execute the script. If not specified, the default cluster root credential secret is used.",
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Description:         "Specifies the name of the secret.",
										MarkdownDescription: "Specifies the name of the secret.",
										Required:            true,
										Optional:            false,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtMost(63),
											stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9]([a-z0-9\.\-]*[a-z0-9])?$`), ""),
										},
									},

									"password_key": schema.StringAttribute{
										Description:         "Used to specify the password part of the secret.",
										MarkdownDescription: "Used to specify the password part of the secret.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"username_key": schema.StringAttribute{
										Description:         "Used to specify the username part of the secret.",
										MarkdownDescription: "Used to specify the username part of the secret.",
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
								Description:         "Specifies the labels used to select the Pods on which the script should be executed.  By default, the script is executed on the Pod associated with the service named '{clusterName}-{componentName}', which typically routes to the Pod with the primary/leader role.  However, some Components, such as Redis, do not synchronize account information between primary and secondary Pods. In these cases, the script must be executed on all replica Pods matching the selector.  Note: this field cannot be modified once set.",
								MarkdownDescription: "Specifies the labels used to select the Pods on which the script should be executed.  By default, the script is executed on the Pod associated with the service named '{clusterName}-{componentName}', which typically routes to the Pod with the primary/leader role.  However, some Components, such as Redis, do not synchronize account information between primary and secondary Pods. In these cases, the script must be executed on all replica Pods matching the selector.  Note: this field cannot be modified once set.",
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"switchover": schema.ListNestedAttribute{
						Description:         "Lists Switchover objects, each specifying a Component to perform the switchover operation.",
						MarkdownDescription: "Lists Switchover objects, each specifying a Component to perform the switchover operation.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"instance_name": schema.StringAttribute{
									Description:         "Specifies the instance to become the primary or leader during a switchover operation.  The value of 'instanceName' can be either:  1. '*' (wildcard value): - Indicates no specific instance is designated as the primary or leader. - Executes the switchover action from 'clusterDefinition.componentDefs[*].switchoverSpec.withoutCandidate'. - 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' must be defined when using '*'.  2. A valid instance name (pod name): - Designates a specific instance (pod) as the primary or leader. - The name must match one of the pods in the component. Any non-valid pod name is considered invalid. - Executes the switchover action from 'clusterDefinition.componentDefs[*].switchoverSpec.withCandidate'. - 'clusterDefinition.componentDefs[*].switchoverSpec.withCandidate' must be defined when specifying a valid instance name.",
									MarkdownDescription: "Specifies the instance to become the primary or leader during a switchover operation.  The value of 'instanceName' can be either:  1. '*' (wildcard value): - Indicates no specific instance is designated as the primary or leader. - Executes the switchover action from 'clusterDefinition.componentDefs[*].switchoverSpec.withoutCandidate'. - 'clusterDefinition.componentDefs[x].switchoverSpec.withoutCandidate' must be defined when using '*'.  2. A valid instance name (pod name): - Designates a specific instance (pod) as the primary or leader. - The name must match one of the pods in the component. Any non-valid pod name is considered invalid. - Executes the switchover action from 'clusterDefinition.componentDefs[*].switchoverSpec.withCandidate'. - 'clusterDefinition.componentDefs[*].switchoverSpec.withCandidate' must be defined when specifying a valid instance name.",
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

					"ttl_seconds_after_succeed": schema.Int64Attribute{
						Description:         "Specifies the duration in seconds that an OpsRequest will remain in the system after successfully completing (when 'opsRequest.status.phase' is 'Succeed') before automatic deletion.",
						MarkdownDescription: "Specifies the duration in seconds that an OpsRequest will remain in the system after successfully completing (when 'opsRequest.status.phase' is 'Succeed') before automatic deletion.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"ttl_seconds_before_abort": schema.Int64Attribute{
						Description:         "Specifies the maximum time in seconds that the OpsRequest will wait for its pre-conditions to be met before it aborts the operation. If set to 0 (default), pre-conditions must be satisfied immediately for the OpsRequest to proceed.",
						MarkdownDescription: "Specifies the maximum time in seconds that the OpsRequest will wait for its pre-conditions to be met before it aborts the operation. If set to 0 (default), pre-conditions must be satisfied immediately for the OpsRequest to proceed.",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"type": schema.StringAttribute{
						Description:         "Specifies the type of this operation. Supported types include 'Start', 'Stop', 'Restart', 'Switchover', 'VerticalScaling', 'HorizontalScaling', 'VolumeExpansion', 'Reconfiguring', 'Upgrade', 'Backup', 'Restore', 'Expose', 'DataScript', 'RebuildInstance', 'Custom'.  Note: This field is immutable once set.",
						MarkdownDescription: "Specifies the type of this operation. Supported types include 'Start', 'Stop', 'Restart', 'Switchover', 'VerticalScaling', 'HorizontalScaling', 'VolumeExpansion', 'Reconfiguring', 'Upgrade', 'Backup', 'Restore', 'Expose', 'DataScript', 'RebuildInstance', 'Custom'.  Note: This field is immutable once set.",
						Required:            true,
						Optional:            false,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.OneOf("Upgrade", "VerticalScaling", "VolumeExpansion", "HorizontalScaling", "Restart", "Reconfiguring", "Start", "Stop", "Expose", "Switchover", "DataScript", "Backup", "Restore", "RebuildInstance", "Custom"),
						},
					},

					"upgrade": schema.SingleNestedAttribute{
						Description:         "Specifies the desired new version of the Cluster.  Note: This field is immutable once set.",
						MarkdownDescription: "Specifies the desired new version of the Cluster.  Note: This field is immutable once set.",
						Attributes: map[string]schema.Attribute{
							"cluster_version_ref": schema.StringAttribute{
								Description:         "Specifies the name of the target ClusterVersion for the upgrade.  This field is deprecated since v0.9 because ClusterVersion is deprecated.",
								MarkdownDescription: "Specifies the name of the target ClusterVersion for the upgrade.  This field is deprecated since v0.9 because ClusterVersion is deprecated.",
								Required:            true,
								Optional:            false,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"vertical_scaling": schema.ListAttribute{
						Description:         "Lists VerticalScaling objects, each specifying a component and its desired compute resources for vertical scaling.",
						MarkdownDescription: "Lists VerticalScaling objects, each specifying a component and its desired compute resources for vertical scaling.",
						ElementType:         types.MapType{ElemType: types.StringType},
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"volume_expansion": schema.ListNestedAttribute{
						Description:         "Lists VolumeExpansion objects, each specifying a component and its corresponding volumeClaimTemplates that requires storage expansion.",
						MarkdownDescription: "Lists VolumeExpansion objects, each specifying a component and its corresponding volumeClaimTemplates that requires storage expansion.",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"component_name": schema.StringAttribute{
									Description:         "Specifies the name of the Component.",
									MarkdownDescription: "Specifies the name of the Component.",
									Required:            true,
									Optional:            false,
									Computed:            false,
								},

								"instances": schema.ListNestedAttribute{
									Description:         "Specifies the instance template that need to volume expand.",
									MarkdownDescription: "Specifies the instance template that need to volume expand.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Refer to the instance template name of the component or sharding.",
												MarkdownDescription: "Refer to the instance template name of the component or sharding.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"volume_claim_templates": schema.ListNestedAttribute{
												Description:         "volumeClaimTemplates specifies the storage size and volumeClaimTemplate name.",
												MarkdownDescription: "volumeClaimTemplates specifies the storage size and volumeClaimTemplate name.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Specify the name of the volumeClaimTemplate in the Component. The specified name must match one of the volumeClaimTemplates defined in the 'clusterComponentSpec.volumeClaimTemplates' field.",
															MarkdownDescription: "Specify the name of the volumeClaimTemplate in the Component. The specified name must match one of the volumeClaimTemplates defined in the 'clusterComponentSpec.volumeClaimTemplates' field.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},

														"storage": schema.StringAttribute{
															Description:         "Specifies the desired storage size for the volume.",
															MarkdownDescription: "Specifies the desired storage size for the volume.",
															Required:            true,
															Optional:            false,
															Computed:            false,
														},
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

								"volume_claim_templates": schema.ListNestedAttribute{
									Description:         "Specifies a list of OpsRequestVolumeClaimTemplate objects, defining the volumeClaimTemplates that are used to expand the storage and the desired storage size for each one.",
									MarkdownDescription: "Specifies a list of OpsRequestVolumeClaimTemplate objects, defining the volumeClaimTemplates that are used to expand the storage and the desired storage size for each one.",
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "Specify the name of the volumeClaimTemplate in the Component. The specified name must match one of the volumeClaimTemplates defined in the 'clusterComponentSpec.volumeClaimTemplates' field.",
												MarkdownDescription: "Specify the name of the volumeClaimTemplate in the Component. The specified name must match one of the volumeClaimTemplates defined in the 'clusterComponentSpec.volumeClaimTemplates' field.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"storage": schema.StringAttribute{
												Description:         "Specifies the desired storage size for the volume.",
												MarkdownDescription: "Specifies the desired storage size for the volume.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
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
				},
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *AppsKubeblocksIoOpsRequestV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_apps_kubeblocks_io_ops_request_v1alpha1_manifest")

	var model AppsKubeblocksIoOpsRequestV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("apps.kubeblocks.io/v1alpha1")
	model.Kind = pointer.String("OpsRequest")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
