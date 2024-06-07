/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package projectcontour_io_v1alpha1

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
	_ datasource.DataSource = &ProjectcontourIoContourDeploymentV1Alpha1Manifest{}
)

func NewProjectcontourIoContourDeploymentV1Alpha1Manifest() datasource.DataSource {
	return &ProjectcontourIoContourDeploymentV1Alpha1Manifest{}
}

type ProjectcontourIoContourDeploymentV1Alpha1Manifest struct{}

type ProjectcontourIoContourDeploymentV1Alpha1ManifestData struct {
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
		Contour *struct {
			Deployment *struct {
				Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Strategy *struct {
					RollingUpdate *struct {
						MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
						MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"strategy" json:"strategy,omitempty"`
			} `tfsdk:"deployment" json:"deployment,omitempty"`
			DisabledFeatures   *[]string `tfsdk:"disabled_features" json:"disabledFeatures,omitempty"`
			KubernetesLogLevel *int64    `tfsdk:"kubernetes_log_level" json:"kubernetesLogLevel,omitempty"`
			LogLevel           *string   `tfsdk:"log_level" json:"logLevel,omitempty"`
			NodePlacement      *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"node_placement" json:"nodePlacement,omitempty"`
			PodAnnotations *map[string]string `tfsdk:"pod_annotations" json:"podAnnotations,omitempty"`
			Replicas       *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources      *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			WatchNamespaces *[]string `tfsdk:"watch_namespaces" json:"watchNamespaces,omitempty"`
		} `tfsdk:"contour" json:"contour,omitempty"`
		Envoy *struct {
			BaseID    *int64 `tfsdk:"base_id" json:"baseID,omitempty"`
			DaemonSet *struct {
				UpdateStrategy *struct {
					RollingUpdate *struct {
						MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
						MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"update_strategy" json:"updateStrategy,omitempty"`
			} `tfsdk:"daemon_set" json:"daemonSet,omitempty"`
			Deployment *struct {
				Replicas *int64 `tfsdk:"replicas" json:"replicas,omitempty"`
				Strategy *struct {
					RollingUpdate *struct {
						MaxSurge       *string `tfsdk:"max_surge" json:"maxSurge,omitempty"`
						MaxUnavailable *string `tfsdk:"max_unavailable" json:"maxUnavailable,omitempty"`
					} `tfsdk:"rolling_update" json:"rollingUpdate,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"strategy" json:"strategy,omitempty"`
			} `tfsdk:"deployment" json:"deployment,omitempty"`
			ExtraVolumeMounts *[]struct {
				MountPath         *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
				MountPropagation  *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
				Name              *string `tfsdk:"name" json:"name,omitempty"`
				ReadOnly          *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
				RecursiveReadOnly *string `tfsdk:"recursive_read_only" json:"recursiveReadOnly,omitempty"`
				SubPath           *string `tfsdk:"sub_path" json:"subPath,omitempty"`
				SubPathExpr       *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
			} `tfsdk:"extra_volume_mounts" json:"extraVolumeMounts,omitempty"`
			ExtraVolumes *[]struct {
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
			} `tfsdk:"extra_volumes" json:"extraVolumes,omitempty"`
			LogLevel          *string `tfsdk:"log_level" json:"logLevel,omitempty"`
			NetworkPublishing *struct {
				ExternalTrafficPolicy *string            `tfsdk:"external_traffic_policy" json:"externalTrafficPolicy,omitempty"`
				IpFamilyPolicy        *string            `tfsdk:"ip_family_policy" json:"ipFamilyPolicy,omitempty"`
				ServiceAnnotations    *map[string]string `tfsdk:"service_annotations" json:"serviceAnnotations,omitempty"`
				Type                  *string            `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"network_publishing" json:"networkPublishing,omitempty"`
			NodePlacement *struct {
				NodeSelector *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
				Tolerations  *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
			} `tfsdk:"node_placement" json:"nodePlacement,omitempty"`
			OverloadMaxHeapSize *int64             `tfsdk:"overload_max_heap_size" json:"overloadMaxHeapSize,omitempty"`
			PodAnnotations      *map[string]string `tfsdk:"pod_annotations" json:"podAnnotations,omitempty"`
			Replicas            *int64             `tfsdk:"replicas" json:"replicas,omitempty"`
			Resources           *struct {
				Claims *[]struct {
					Name *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"claims" json:"claims,omitempty"`
				Limits   *map[string]string `tfsdk:"limits" json:"limits,omitempty"`
				Requests *map[string]string `tfsdk:"requests" json:"requests,omitempty"`
			} `tfsdk:"resources" json:"resources,omitempty"`
			WorkloadType *string `tfsdk:"workload_type" json:"workloadType,omitempty"`
		} `tfsdk:"envoy" json:"envoy,omitempty"`
		ResourceLabels  *map[string]string `tfsdk:"resource_labels" json:"resourceLabels,omitempty"`
		RuntimeSettings *struct {
			Debug *struct {
				Address *string `tfsdk:"address" json:"address,omitempty"`
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"debug" json:"debug,omitempty"`
			EnableExternalNameService *bool `tfsdk:"enable_external_name_service" json:"enableExternalNameService,omitempty"`
			Envoy                     *struct {
				ClientCertificate *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"client_certificate" json:"clientCertificate,omitempty"`
				Cluster *struct {
					CircuitBreakers *struct {
						MaxConnections     *int64 `tfsdk:"max_connections" json:"maxConnections,omitempty"`
						MaxPendingRequests *int64 `tfsdk:"max_pending_requests" json:"maxPendingRequests,omitempty"`
						MaxRequests        *int64 `tfsdk:"max_requests" json:"maxRequests,omitempty"`
						MaxRetries         *int64 `tfsdk:"max_retries" json:"maxRetries,omitempty"`
					} `tfsdk:"circuit_breakers" json:"circuitBreakers,omitempty"`
					DnsLookupFamily                   *string `tfsdk:"dns_lookup_family" json:"dnsLookupFamily,omitempty"`
					MaxRequestsPerConnection          *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
					Per_connection_buffer_limit_bytes *int64  `tfsdk:"per_connection_buffer_limit_bytes" json:"per-connection-buffer-limit-bytes,omitempty"`
					UpstreamTLS                       *struct {
						CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
						MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
						MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
					} `tfsdk:"upstream_tls" json:"upstreamTLS,omitempty"`
				} `tfsdk:"cluster" json:"cluster,omitempty"`
				DefaultHTTPVersions *[]string `tfsdk:"default_http_versions" json:"defaultHTTPVersions,omitempty"`
				Health              *struct {
					Address *string `tfsdk:"address" json:"address,omitempty"`
					Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"health" json:"health,omitempty"`
				Http *struct {
					AccessLog *string `tfsdk:"access_log" json:"accessLog,omitempty"`
					Address   *string `tfsdk:"address" json:"address,omitempty"`
					Port      *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"http" json:"http,omitempty"`
				Https *struct {
					AccessLog *string `tfsdk:"access_log" json:"accessLog,omitempty"`
					Address   *string `tfsdk:"address" json:"address,omitempty"`
					Port      *int64  `tfsdk:"port" json:"port,omitempty"`
				} `tfsdk:"https" json:"https,omitempty"`
				Listener *struct {
					ConnectionBalancer                *string `tfsdk:"connection_balancer" json:"connectionBalancer,omitempty"`
					DisableAllowChunkedLength         *bool   `tfsdk:"disable_allow_chunked_length" json:"disableAllowChunkedLength,omitempty"`
					DisableMergeSlashes               *bool   `tfsdk:"disable_merge_slashes" json:"disableMergeSlashes,omitempty"`
					HttpMaxConcurrentStreams          *int64  `tfsdk:"http_max_concurrent_streams" json:"httpMaxConcurrentStreams,omitempty"`
					MaxConnectionsPerListener         *int64  `tfsdk:"max_connections_per_listener" json:"maxConnectionsPerListener,omitempty"`
					MaxRequestsPerConnection          *int64  `tfsdk:"max_requests_per_connection" json:"maxRequestsPerConnection,omitempty"`
					MaxRequestsPerIOCycle             *int64  `tfsdk:"max_requests_per_io_cycle" json:"maxRequestsPerIOCycle,omitempty"`
					Per_connection_buffer_limit_bytes *int64  `tfsdk:"per_connection_buffer_limit_bytes" json:"per-connection-buffer-limit-bytes,omitempty"`
					ServerHeaderTransformation        *string `tfsdk:"server_header_transformation" json:"serverHeaderTransformation,omitempty"`
					SocketOptions                     *struct {
						Tos          *int64 `tfsdk:"tos" json:"tos,omitempty"`
						TrafficClass *int64 `tfsdk:"traffic_class" json:"trafficClass,omitempty"`
					} `tfsdk:"socket_options" json:"socketOptions,omitempty"`
					Tls *struct {
						CipherSuites           *[]string `tfsdk:"cipher_suites" json:"cipherSuites,omitempty"`
						MaximumProtocolVersion *string   `tfsdk:"maximum_protocol_version" json:"maximumProtocolVersion,omitempty"`
						MinimumProtocolVersion *string   `tfsdk:"minimum_protocol_version" json:"minimumProtocolVersion,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
					UseProxyProtocol *bool `tfsdk:"use_proxy_protocol" json:"useProxyProtocol,omitempty"`
				} `tfsdk:"listener" json:"listener,omitempty"`
				Logging *struct {
					AccessLogFormat       *string   `tfsdk:"access_log_format" json:"accessLogFormat,omitempty"`
					AccessLogFormatString *string   `tfsdk:"access_log_format_string" json:"accessLogFormatString,omitempty"`
					AccessLogJSONFields   *[]string `tfsdk:"access_log_json_fields" json:"accessLogJSONFields,omitempty"`
					AccessLogLevel        *string   `tfsdk:"access_log_level" json:"accessLogLevel,omitempty"`
				} `tfsdk:"logging" json:"logging,omitempty"`
				Metrics *struct {
					Address *string `tfsdk:"address" json:"address,omitempty"`
					Port    *int64  `tfsdk:"port" json:"port,omitempty"`
					Tls     *struct {
						CaFile   *string `tfsdk:"ca_file" json:"caFile,omitempty"`
						CertFile *string `tfsdk:"cert_file" json:"certFile,omitempty"`
						KeyFile  *string `tfsdk:"key_file" json:"keyFile,omitempty"`
					} `tfsdk:"tls" json:"tls,omitempty"`
				} `tfsdk:"metrics" json:"metrics,omitempty"`
				Network *struct {
					AdminPort      *int64 `tfsdk:"admin_port" json:"adminPort,omitempty"`
					NumTrustedHops *int64 `tfsdk:"num_trusted_hops" json:"numTrustedHops,omitempty"`
				} `tfsdk:"network" json:"network,omitempty"`
				Service *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"service" json:"service,omitempty"`
				Timeouts *struct {
					ConnectTimeout                *string `tfsdk:"connect_timeout" json:"connectTimeout,omitempty"`
					ConnectionIdleTimeout         *string `tfsdk:"connection_idle_timeout" json:"connectionIdleTimeout,omitempty"`
					ConnectionShutdownGracePeriod *string `tfsdk:"connection_shutdown_grace_period" json:"connectionShutdownGracePeriod,omitempty"`
					DelayedCloseTimeout           *string `tfsdk:"delayed_close_timeout" json:"delayedCloseTimeout,omitempty"`
					MaxConnectionDuration         *string `tfsdk:"max_connection_duration" json:"maxConnectionDuration,omitempty"`
					RequestTimeout                *string `tfsdk:"request_timeout" json:"requestTimeout,omitempty"`
					StreamIdleTimeout             *string `tfsdk:"stream_idle_timeout" json:"streamIdleTimeout,omitempty"`
				} `tfsdk:"timeouts" json:"timeouts,omitempty"`
			} `tfsdk:"envoy" json:"envoy,omitempty"`
			FeatureFlags *[]string `tfsdk:"feature_flags" json:"featureFlags,omitempty"`
			Gateway      *struct {
				GatewayRef *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"gateway_ref" json:"gatewayRef,omitempty"`
			} `tfsdk:"gateway" json:"gateway,omitempty"`
			GlobalExtAuth *struct {
				AuthPolicy *struct {
					Context  *map[string]string `tfsdk:"context" json:"context,omitempty"`
					Disabled *bool              `tfsdk:"disabled" json:"disabled,omitempty"`
				} `tfsdk:"auth_policy" json:"authPolicy,omitempty"`
				ExtensionRef *struct {
					ApiVersion *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					Name       *string `tfsdk:"name" json:"name,omitempty"`
					Namespace  *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"extension_ref" json:"extensionRef,omitempty"`
				FailOpen        *bool   `tfsdk:"fail_open" json:"failOpen,omitempty"`
				ResponseTimeout *string `tfsdk:"response_timeout" json:"responseTimeout,omitempty"`
				WithRequestBody *struct {
					AllowPartialMessage *bool  `tfsdk:"allow_partial_message" json:"allowPartialMessage,omitempty"`
					MaxRequestBytes     *int64 `tfsdk:"max_request_bytes" json:"maxRequestBytes,omitempty"`
					PackAsBytes         *bool  `tfsdk:"pack_as_bytes" json:"packAsBytes,omitempty"`
				} `tfsdk:"with_request_body" json:"withRequestBody,omitempty"`
			} `tfsdk:"global_ext_auth" json:"globalExtAuth,omitempty"`
			Health *struct {
				Address *string `tfsdk:"address" json:"address,omitempty"`
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
			} `tfsdk:"health" json:"health,omitempty"`
			Httpproxy *struct {
				DisablePermitInsecure *bool `tfsdk:"disable_permit_insecure" json:"disablePermitInsecure,omitempty"`
				FallbackCertificate   *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"fallback_certificate" json:"fallbackCertificate,omitempty"`
				RootNamespaces *[]string `tfsdk:"root_namespaces" json:"rootNamespaces,omitempty"`
			} `tfsdk:"httpproxy" json:"httpproxy,omitempty"`
			Ingress *struct {
				ClassNames    *[]string `tfsdk:"class_names" json:"classNames,omitempty"`
				StatusAddress *string   `tfsdk:"status_address" json:"statusAddress,omitempty"`
			} `tfsdk:"ingress" json:"ingress,omitempty"`
			Metrics *struct {
				Address *string `tfsdk:"address" json:"address,omitempty"`
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				Tls     *struct {
					CaFile   *string `tfsdk:"ca_file" json:"caFile,omitempty"`
					CertFile *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					KeyFile  *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
			} `tfsdk:"metrics" json:"metrics,omitempty"`
			Policy *struct {
				ApplyToIngress *bool `tfsdk:"apply_to_ingress" json:"applyToIngress,omitempty"`
				RequestHeaders *struct {
					Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
					Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"request_headers" json:"requestHeaders,omitempty"`
				ResponseHeaders *struct {
					Remove *[]string          `tfsdk:"remove" json:"remove,omitempty"`
					Set    *map[string]string `tfsdk:"set" json:"set,omitempty"`
				} `tfsdk:"response_headers" json:"responseHeaders,omitempty"`
			} `tfsdk:"policy" json:"policy,omitempty"`
			RateLimitService *struct {
				DefaultGlobalRateLimitPolicy *struct {
					Descriptors *[]struct {
						Entries *[]struct {
							GenericKey *struct {
								Key   *string `tfsdk:"key" json:"key,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"generic_key" json:"genericKey,omitempty"`
							RemoteAddress *map[string]string `tfsdk:"remote_address" json:"remoteAddress,omitempty"`
							RequestHeader *struct {
								DescriptorKey *string `tfsdk:"descriptor_key" json:"descriptorKey,omitempty"`
								HeaderName    *string `tfsdk:"header_name" json:"headerName,omitempty"`
							} `tfsdk:"request_header" json:"requestHeader,omitempty"`
							RequestHeaderValueMatch *struct {
								ExpectMatch *bool `tfsdk:"expect_match" json:"expectMatch,omitempty"`
								Headers     *[]struct {
									Contains            *string `tfsdk:"contains" json:"contains,omitempty"`
									Exact               *string `tfsdk:"exact" json:"exact,omitempty"`
									IgnoreCase          *bool   `tfsdk:"ignore_case" json:"ignoreCase,omitempty"`
									Name                *string `tfsdk:"name" json:"name,omitempty"`
									Notcontains         *string `tfsdk:"notcontains" json:"notcontains,omitempty"`
									Notexact            *string `tfsdk:"notexact" json:"notexact,omitempty"`
									Notpresent          *bool   `tfsdk:"notpresent" json:"notpresent,omitempty"`
									Present             *bool   `tfsdk:"present" json:"present,omitempty"`
									Regex               *string `tfsdk:"regex" json:"regex,omitempty"`
									TreatMissingAsEmpty *bool   `tfsdk:"treat_missing_as_empty" json:"treatMissingAsEmpty,omitempty"`
								} `tfsdk:"headers" json:"headers,omitempty"`
								Value *string `tfsdk:"value" json:"value,omitempty"`
							} `tfsdk:"request_header_value_match" json:"requestHeaderValueMatch,omitempty"`
						} `tfsdk:"entries" json:"entries,omitempty"`
					} `tfsdk:"descriptors" json:"descriptors,omitempty"`
					Disabled *bool `tfsdk:"disabled" json:"disabled,omitempty"`
				} `tfsdk:"default_global_rate_limit_policy" json:"defaultGlobalRateLimitPolicy,omitempty"`
				Domain                      *string `tfsdk:"domain" json:"domain,omitempty"`
				EnableResourceExhaustedCode *bool   `tfsdk:"enable_resource_exhausted_code" json:"enableResourceExhaustedCode,omitempty"`
				EnableXRateLimitHeaders     *bool   `tfsdk:"enable_x_rate_limit_headers" json:"enableXRateLimitHeaders,omitempty"`
				ExtensionService            *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"extension_service" json:"extensionService,omitempty"`
				FailOpen *bool `tfsdk:"fail_open" json:"failOpen,omitempty"`
			} `tfsdk:"rate_limit_service" json:"rateLimitService,omitempty"`
			Tracing *struct {
				CustomTags *[]struct {
					Literal           *string `tfsdk:"literal" json:"literal,omitempty"`
					RequestHeaderName *string `tfsdk:"request_header_name" json:"requestHeaderName,omitempty"`
					TagName           *string `tfsdk:"tag_name" json:"tagName,omitempty"`
				} `tfsdk:"custom_tags" json:"customTags,omitempty"`
				ExtensionService *struct {
					Name      *string `tfsdk:"name" json:"name,omitempty"`
					Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
				} `tfsdk:"extension_service" json:"extensionService,omitempty"`
				IncludePodDetail *bool   `tfsdk:"include_pod_detail" json:"includePodDetail,omitempty"`
				MaxPathTagLength *int64  `tfsdk:"max_path_tag_length" json:"maxPathTagLength,omitempty"`
				OverallSampling  *string `tfsdk:"overall_sampling" json:"overallSampling,omitempty"`
				ServiceName      *string `tfsdk:"service_name" json:"serviceName,omitempty"`
			} `tfsdk:"tracing" json:"tracing,omitempty"`
			XdsServer *struct {
				Address *string `tfsdk:"address" json:"address,omitempty"`
				Port    *int64  `tfsdk:"port" json:"port,omitempty"`
				Tls     *struct {
					CaFile   *string `tfsdk:"ca_file" json:"caFile,omitempty"`
					CertFile *string `tfsdk:"cert_file" json:"certFile,omitempty"`
					Insecure *bool   `tfsdk:"insecure" json:"insecure,omitempty"`
					KeyFile  *string `tfsdk:"key_file" json:"keyFile,omitempty"`
				} `tfsdk:"tls" json:"tls,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"xds_server" json:"xdsServer,omitempty"`
		} `tfsdk:"runtime_settings" json:"runtimeSettings,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *ProjectcontourIoContourDeploymentV1Alpha1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_projectcontour_io_contour_deployment_v1alpha1_manifest"
}

func (r *ProjectcontourIoContourDeploymentV1Alpha1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "ContourDeployment is the schema for a Contour Deployment.",
		MarkdownDescription: "ContourDeployment is the schema for a Contour Deployment.",
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
				Description:         "ContourDeploymentSpec specifies options for how a Contourinstance should be provisioned.",
				MarkdownDescription: "ContourDeploymentSpec specifies options for how a Contourinstance should be provisioned.",
				Attributes: map[string]schema.Attribute{
					"contour": schema.SingleNestedAttribute{
						Description:         "Contour specifies deployment-time settings for the Contourpart of the installation, i.e. the xDS server/control planeand associated resources, including things like replica countfor the Deployment, and node placement constraints for the pods.",
						MarkdownDescription: "Contour specifies deployment-time settings for the Contourpart of the installation, i.e. the xDS server/control planeand associated resources, including things like replica countfor the Deployment, and node placement constraints for the pods.",
						Attributes: map[string]schema.Attribute{
							"deployment": schema.SingleNestedAttribute{
								Description:         "Deployment describes the settings for running contour as a 'Deployment'.",
								MarkdownDescription: "Deployment describes the settings for running contour as a 'Deployment'.",
								Attributes: map[string]schema.Attribute{
									"replicas": schema.Int64Attribute{
										Description:         "Replicas is the desired number of replicas.",
										MarkdownDescription: "Replicas is the desired number of replicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"strategy": schema.SingleNestedAttribute{
										Description:         "Strategy describes the deployment strategy to use to replace existing pods with new pods.",
										MarkdownDescription: "Strategy describes the deployment strategy to use to replace existing pods with new pods.",
										Attributes: map[string]schema.Attribute{
											"rolling_update": schema.SingleNestedAttribute{
												Description:         "Rolling update config params. Present only if DeploymentStrategyType =RollingUpdate.---TODO: Update this to follow our convention for oneOf, whatever we decide itto be.",
												MarkdownDescription: "Rolling update config params. Present only if DeploymentStrategyType =RollingUpdate.---TODO: Update this to follow our convention for oneOf, whatever we decide itto be.",
												Attributes: map[string]schema.Attribute{
													"max_surge": schema.StringAttribute{
														Description:         "The maximum number of pods that can be scheduled above the desired number ofpods.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 25%.Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately whenthe rolling update starts, such that the total number of old and new pods do not exceed130% of desired pods. Once old pods have been killed,new ReplicaSet can be scaled up further, ensuring that total number of pods runningat any time during the update is at most 130% of desired pods.",
														MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number ofpods.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 25%.Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately whenthe rolling update starts, such that the total number of old and new pods do not exceed130% of desired pods. Once old pods have been killed,new ReplicaSet can be scaled up further, ensuring that total number of pods runningat any time during the update is at most 130% of desired pods.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_unavailable": schema.StringAttribute{
														Description:         "The maximum number of pods that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 25%.Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired podsimmediately when the rolling update starts. Once new pods are ready, old ReplicaSetcan be scaled down further, followed by scaling up the new ReplicaSet, ensuringthat the total number of pods available at all times during the update is atleast 70% of desired pods.",
														MarkdownDescription: "The maximum number of pods that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 25%.Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired podsimmediately when the rolling update starts. Once new pods are ready, old ReplicaSetcan be scaled down further, followed by scaling up the new ReplicaSet, ensuringthat the total number of pods available at all times during the update is atleast 70% of desired pods.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.",
												MarkdownDescription: "Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.",
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

							"disabled_features": schema.ListAttribute{
								Description:         "DisabledFeatures defines an array of resources that will be ignored bycontour reconciler.",
								MarkdownDescription: "DisabledFeatures defines an array of resources that will be ignored bycontour reconciler.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"kubernetes_log_level": schema.Int64Attribute{
								Description:         "KubernetesLogLevel Enable Kubernetes client debug logging with log level. If unset,defaults to 0.",
								MarkdownDescription: "KubernetesLogLevel Enable Kubernetes client debug logging with log level. If unset,defaults to 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
									int64validator.AtMost(9),
								},
							},

							"log_level": schema.StringAttribute{
								Description:         "LogLevel sets the log level for ContourAllowed values are 'info', 'debug'.",
								MarkdownDescription: "LogLevel sets the log level for ContourAllowed values are 'info', 'debug'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"node_placement": schema.SingleNestedAttribute{
								Description:         "NodePlacement describes node scheduling configuration of Contour pods.",
								MarkdownDescription: "NodePlacement describes node scheduling configuration of Contour pods.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector is the simplest recommended form of node selection constraintand specifies a map of key-value pairs. For the pod to be eligibleto run on a node, the node must have each of the indicated key-value pairsas labels (it can have additional labels as well).If unset, the pod(s) will be scheduled to any available node.",
										MarkdownDescription: "NodeSelector is the simplest recommended form of node selection constraintand specifies a map of key-value pairs. For the pod to be eligibleto run on a node, the node must have each of the indicated key-value pairsas labels (it can have additional labels as well).If unset, the pod(s) will be scheduled to any available node.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations work with taints to ensure that pods are not scheduledonto inappropriate nodes. One or more taints are applied to a node; thismarks that the node should not accept any pods that do not tolerate thetaints.The default is an empty list.See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/for additional details.",
										MarkdownDescription: "Tolerations work with taints to ensure that pods are not scheduledonto inappropriate nodes. One or more taints are applied to a node; thismarks that the node should not accept any pods that do not tolerate thetaints.The default is an empty list.See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/for additional details.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"pod_annotations": schema.MapAttribute{
								Description:         "PodAnnotations defines annotations to add to the Contour pods.the annotations for Prometheus will be appended or overwritten with predefined value.",
								MarkdownDescription: "PodAnnotations defines annotations to add to the Contour pods.the annotations for Prometheus will be appended or overwritten with predefined value.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Deprecated: Use 'DeploymentSettings.Replicas' instead.Replicas is the desired number of Contour replicas. If if unset,defaults to 2.if both 'DeploymentSettings.Replicas' and this one is set, use 'DeploymentSettings.Replicas'.",
								MarkdownDescription: "Deprecated: Use 'DeploymentSettings.Replicas' instead.Replicas is the desired number of Contour replicas. If if unset,defaults to 2.if both 'DeploymentSettings.Replicas' and this one is set, use 'DeploymentSettings.Replicas'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Compute Resources required by contour container.Cannot be updated.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Compute Resources required by contour container.Cannot be updated.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"watch_namespaces": schema.ListAttribute{
								Description:         "WatchNamespaces is an array of namespaces. Setting it will instruct the contour instanceto only watch this subset of namespaces.",
								MarkdownDescription: "WatchNamespaces is an array of namespaces. Setting it will instruct the contour instanceto only watch this subset of namespaces.",
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

					"envoy": schema.SingleNestedAttribute{
						Description:         "Envoy specifies deployment-time settings for the Envoypart of the installation, i.e. the xDS client/data planeand associated resources, including things like the workloadtype to use (DaemonSet or Deployment), node placement constraintsfor the pods, and various options for the Envoy service.",
						MarkdownDescription: "Envoy specifies deployment-time settings for the Envoypart of the installation, i.e. the xDS client/data planeand associated resources, including things like the workloadtype to use (DaemonSet or Deployment), node placement constraintsfor the pods, and various options for the Envoy service.",
						Attributes: map[string]schema.Attribute{
							"base_id": schema.Int64Attribute{
								Description:         "The base ID to use when allocating shared memory regions.if Envoy needs to be run multiple times on the same machine, each running Envoy will need a unique base IDso that the shared memory regions do not conflict.defaults to 0.",
								MarkdownDescription: "The base ID to use when allocating shared memory regions.if Envoy needs to be run multiple times on the same machine, each running Envoy will need a unique base IDso that the shared memory regions do not conflict.defaults to 0.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"daemon_set": schema.SingleNestedAttribute{
								Description:         "DaemonSet describes the settings for running envoy as a 'DaemonSet'.if 'WorkloadType' is 'Deployment',it's must be nil",
								MarkdownDescription: "DaemonSet describes the settings for running envoy as a 'DaemonSet'.if 'WorkloadType' is 'Deployment',it's must be nil",
								Attributes: map[string]schema.Attribute{
									"update_strategy": schema.SingleNestedAttribute{
										Description:         "Strategy describes the deployment strategy to use to replace existing DaemonSet pods with new pods.",
										MarkdownDescription: "Strategy describes the deployment strategy to use to replace existing DaemonSet pods with new pods.",
										Attributes: map[string]schema.Attribute{
											"rolling_update": schema.SingleNestedAttribute{
												Description:         "Rolling update config params. Present only if type = 'RollingUpdate'.---TODO: Update this to follow our convention for oneOf, whatever we decide itto be. Same as Deployment 'strategy.rollingUpdate'.See https://github.com/kubernetes/kubernetes/issues/35345",
												MarkdownDescription: "Rolling update config params. Present only if type = 'RollingUpdate'.---TODO: Update this to follow our convention for oneOf, whatever we decide itto be. Same as Deployment 'strategy.rollingUpdate'.See https://github.com/kubernetes/kubernetes/issues/35345",
												Attributes: map[string]schema.Attribute{
													"max_surge": schema.StringAttribute{
														Description:         "The maximum number of nodes with an existing available DaemonSet pod thatcan have an updated DaemonSet pod during during an update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up to a minimum of 1.Default value is 0.Example: when this is set to 30%, at most 30% of the total number of nodesthat should be running the daemon pod (i.e. status.desiredNumberScheduled)can have their a new pod created before the old pod is marked as deleted.The update starts by launching new pods on 30% of nodes. Once an updatedpod is available (Ready for at least minReadySeconds) the old DaemonSet podon that node is marked deleted. If the old pod becomes unavailable for anyreason (Ready transitions to false, is evicted, or is drained) an updatedpod is immediatedly created on that node without considering surge limits.Allowing surge implies the possibility that the resources consumed by thedaemonset on any given node can double if the readiness check fails, andso resource intensive daemonsets should take into account that they maycause evictions during disruption.",
														MarkdownDescription: "The maximum number of nodes with an existing available DaemonSet pod thatcan have an updated DaemonSet pod during during an update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up to a minimum of 1.Default value is 0.Example: when this is set to 30%, at most 30% of the total number of nodesthat should be running the daemon pod (i.e. status.desiredNumberScheduled)can have their a new pod created before the old pod is marked as deleted.The update starts by launching new pods on 30% of nodes. Once an updatedpod is available (Ready for at least minReadySeconds) the old DaemonSet podon that node is marked deleted. If the old pod becomes unavailable for anyreason (Ready transitions to false, is evicted, or is drained) an updatedpod is immediatedly created on that node without considering surge limits.Allowing surge implies the possibility that the resources consumed by thedaemonset on any given node can double if the readiness check fails, andso resource intensive daemonsets should take into account that they maycause evictions during disruption.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_unavailable": schema.StringAttribute{
														Description:         "The maximum number of DaemonSet pods that can be unavailable during theupdate. Value can be an absolute number (ex: 5) or a percentage of totalnumber of DaemonSet pods at the start of the update (ex: 10%). Absolutenumber is calculated from percentage by rounding up.This cannot be 0 if MaxSurge is 0Default value is 1.Example: when this is set to 30%, at most 30% of the total number of nodesthat should be running the daemon pod (i.e. status.desiredNumberScheduled)can have their pods stopped for an update at any given time. The updatestarts by stopping at most 30% of those DaemonSet pods and then bringsup new DaemonSet pods in their place. Once the new pods are available,it then proceeds onto other DaemonSet pods, thus ensuring that at least70% of original number of DaemonSet pods are available at all times duringthe update.",
														MarkdownDescription: "The maximum number of DaemonSet pods that can be unavailable during theupdate. Value can be an absolute number (ex: 5) or a percentage of totalnumber of DaemonSet pods at the start of the update (ex: 10%). Absolutenumber is calculated from percentage by rounding up.This cannot be 0 if MaxSurge is 0Default value is 1.Example: when this is set to 30%, at most 30% of the total number of nodesthat should be running the daemon pod (i.e. status.desiredNumberScheduled)can have their pods stopped for an update at any given time. The updatestarts by stopping at most 30% of those DaemonSet pods and then bringsup new DaemonSet pods in their place. Once the new pods are available,it then proceeds onto other DaemonSet pods, thus ensuring that at least70% of original number of DaemonSet pods are available at all times duringthe update.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "Type of daemon set update. Can be 'RollingUpdate' or 'OnDelete'. Default is RollingUpdate.",
												MarkdownDescription: "Type of daemon set update. Can be 'RollingUpdate' or 'OnDelete'. Default is RollingUpdate.",
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

							"deployment": schema.SingleNestedAttribute{
								Description:         "Deployment describes the settings for running envoy as a 'Deployment'.if 'WorkloadType' is 'DaemonSet',it's must be nil",
								MarkdownDescription: "Deployment describes the settings for running envoy as a 'Deployment'.if 'WorkloadType' is 'DaemonSet',it's must be nil",
								Attributes: map[string]schema.Attribute{
									"replicas": schema.Int64Attribute{
										Description:         "Replicas is the desired number of replicas.",
										MarkdownDescription: "Replicas is the desired number of replicas.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.Int64{
											int64validator.AtLeast(0),
										},
									},

									"strategy": schema.SingleNestedAttribute{
										Description:         "Strategy describes the deployment strategy to use to replace existing pods with new pods.",
										MarkdownDescription: "Strategy describes the deployment strategy to use to replace existing pods with new pods.",
										Attributes: map[string]schema.Attribute{
											"rolling_update": schema.SingleNestedAttribute{
												Description:         "Rolling update config params. Present only if DeploymentStrategyType =RollingUpdate.---TODO: Update this to follow our convention for oneOf, whatever we decide itto be.",
												MarkdownDescription: "Rolling update config params. Present only if DeploymentStrategyType =RollingUpdate.---TODO: Update this to follow our convention for oneOf, whatever we decide itto be.",
												Attributes: map[string]schema.Attribute{
													"max_surge": schema.StringAttribute{
														Description:         "The maximum number of pods that can be scheduled above the desired number ofpods.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 25%.Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately whenthe rolling update starts, such that the total number of old and new pods do not exceed130% of desired pods. Once old pods have been killed,new ReplicaSet can be scaled up further, ensuring that total number of pods runningat any time during the update is at most 130% of desired pods.",
														MarkdownDescription: "The maximum number of pods that can be scheduled above the desired number ofpods.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).This can not be 0 if MaxUnavailable is 0.Absolute number is calculated from percentage by rounding up.Defaults to 25%.Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately whenthe rolling update starts, such that the total number of old and new pods do not exceed130% of desired pods. Once old pods have been killed,new ReplicaSet can be scaled up further, ensuring that total number of pods runningat any time during the update is at most 130% of desired pods.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_unavailable": schema.StringAttribute{
														Description:         "The maximum number of pods that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 25%.Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired podsimmediately when the rolling update starts. Once new pods are ready, old ReplicaSetcan be scaled down further, followed by scaling up the new ReplicaSet, ensuringthat the total number of pods available at all times during the update is atleast 70% of desired pods.",
														MarkdownDescription: "The maximum number of pods that can be unavailable during the update.Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).Absolute number is calculated from percentage by rounding down.This can not be 0 if MaxSurge is 0.Defaults to 25%.Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired podsimmediately when the rolling update starts. Once new pods are ready, old ReplicaSetcan be scaled down further, followed by scaling up the new ReplicaSet, ensuringthat the total number of pods available at all times during the update is atleast 70% of desired pods.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"type": schema.StringAttribute{
												Description:         "Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.",
												MarkdownDescription: "Type of deployment. Can be 'Recreate' or 'RollingUpdate'. Default is RollingUpdate.",
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

							"extra_volume_mounts": schema.ListNestedAttribute{
								Description:         "ExtraVolumeMounts holds the extra volume mounts to add (normally used with extraVolumes).",
								MarkdownDescription: "ExtraVolumeMounts holds the extra volume mounts to add (normally used with extraVolumes).",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"mount_path": schema.StringAttribute{
											Description:         "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
											MarkdownDescription: "Path within the container at which the volume should be mounted.  Mustnot contain ':'.",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"mount_propagation": schema.StringAttribute{
											Description:         "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified(which defaults to None).",
											MarkdownDescription: "mountPropagation determines how mounts are propagated from the hostto container and the other way around.When not set, MountPropagationNone is used.This field is beta in 1.10.When RecursiveReadOnly is set to IfPossible or to Enabled, MountPropagation must be None or unspecified(which defaults to None).",
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
											Description:         "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
											MarkdownDescription: "Mounted read-only if true, read-write otherwise (false or unspecified).Defaults to false.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"recursive_read_only": schema.StringAttribute{
											Description:         "RecursiveReadOnly specifies whether read-only mounts should be handledrecursively.If ReadOnly is false, this field has no meaning and must be unspecified.If ReadOnly is true, and this field is set to Disabled, the mount is not maderecursively read-only.  If this field is set to IfPossible, the mount is maderecursively read-only, if it is supported by the container runtime.  If thisfield is set to Enabled, the mount is made recursively read-only if it issupported by the container runtime, otherwise the pod will not be started andan error will be generated to indicate the reason.If this field is set to IfPossible or Enabled, MountPropagation must be set toNone (or be unspecified, which defaults to None).If this field is not specified, it is treated as an equivalent of Disabled.",
											MarkdownDescription: "RecursiveReadOnly specifies whether read-only mounts should be handledrecursively.If ReadOnly is false, this field has no meaning and must be unspecified.If ReadOnly is true, and this field is set to Disabled, the mount is not maderecursively read-only.  If this field is set to IfPossible, the mount is maderecursively read-only, if it is supported by the container runtime.  If thisfield is set to Enabled, the mount is made recursively read-only if it issupported by the container runtime, otherwise the pod will not be started andan error will be generated to indicate the reason.If this field is set to IfPossible or Enabled, MountPropagation must be set toNone (or be unspecified, which defaults to None).If this field is not specified, it is treated as an equivalent of Disabled.",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path": schema.StringAttribute{
											Description:         "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
											MarkdownDescription: "Path within the volume from which the container's volume should be mounted.Defaults to '' (volume's root).",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"sub_path_expr": schema.StringAttribute{
											Description:         "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
											MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted.Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment.Defaults to '' (volume's root).SubPathExpr and SubPath are mutually exclusive.",
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

							"extra_volumes": schema.ListNestedAttribute{
								Description:         "ExtraVolumes holds the extra volumes to add.",
								MarkdownDescription: "ExtraVolumes holds the extra volumes to add.",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"aws_elastic_block_store": schema.SingleNestedAttribute{
											Description:         "awsElasticBlockStore represents an AWS Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											MarkdownDescription: "awsElasticBlockStore represents an AWS Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstoreTODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstoreTODO: how do we prevent errors in the filesystem from compromising the machine",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"partition": schema.Int64Attribute{
													Description:         "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
													MarkdownDescription: "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly value true will force the readOnly setting in VolumeMounts.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
													MarkdownDescription: "readOnly value true will force the readOnly setting in VolumeMounts.More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_id": schema.StringAttribute{
													Description:         "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
													MarkdownDescription: "volumeID is unique ID of the persistent disk resource in AWS (Amazon EBS volume).More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
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
													Description:         "fsType is Filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fsType is Filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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
													Description:         "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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
													Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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
													Description:         "monitors is Required: Monitors is a collection of Ceph monitorsMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "monitors is Required: Monitors is a collection of Ceph monitorsMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
													Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_file": schema.StringAttribute{
													Description:         "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secretMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "secretFile is Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secretMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "secretRef is Optional: SecretRef is reference to the authentication secret for User, default is empty.More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
													Description:         "user is optional: User is the rados user name, default is adminMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
													MarkdownDescription: "user is optional: User is the rados user name, default is adminMore info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
											Description:         "cinder represents a cinder volume attached and mounted on kubelets host machine.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											MarkdownDescription: "cinder represents a cinder volume attached and mounted on kubelets host machine.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is optional: points to a secret object containing parameters used to connectto OpenStack.",
													MarkdownDescription: "secretRef is optional: points to a secret object containing parameters used to connectto OpenStack.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
													Description:         "volumeID used to identify the volume in cinder.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
													MarkdownDescription: "volumeID used to identify the volume in cinder.More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
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
													Description:         "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
													MarkdownDescription: "defaultMode is optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

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

										"csi": schema.SingleNestedAttribute{
											Description:         "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
											MarkdownDescription: "csi (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
											Attributes: map[string]schema.Attribute{
												"driver": schema.StringAttribute{
													Description:         "driver is the name of the CSI driver that handles this volume.Consult with your admin for the correct name as registered in the cluster.",
													MarkdownDescription: "driver is the name of the CSI driver that handles this volume.Consult with your admin for the correct name as registered in the cluster.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"fs_type": schema.StringAttribute{
													Description:         "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'.If not provided, the empty value is passed to the associated CSI driverwhich will determine the default filesystem to apply.",
													MarkdownDescription: "fsType to mount. Ex. 'ext4', 'xfs', 'ntfs'.If not provided, the empty value is passed to the associated CSI driverwhich will determine the default filesystem to apply.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_publish_secret_ref": schema.SingleNestedAttribute{
													Description:         "nodePublishSecretRef is a reference to the secret object containingsensitive information to pass to the CSI driver to complete the CSINodePublishVolume and NodeUnpublishVolume calls.This field is optional, and  may be empty if no secret is required. If thesecret object contains more than one secret, all secret references are passed.",
													MarkdownDescription: "nodePublishSecretRef is a reference to the secret object containingsensitive information to pass to the CSI driver to complete the CSINodePublishVolume and NodeUnpublishVolume calls.This field is optional, and  may be empty if no secret is required. If thesecret object contains more than one secret, all secret references are passed.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
													Description:         "readOnly specifies a read-only configuration for the volume.Defaults to false (read/write).",
													MarkdownDescription: "readOnly specifies a read-only configuration for the volume.Defaults to false (read/write).",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_attributes": schema.MapAttribute{
													Description:         "volumeAttributes stores driver-specific properties that are passed to the CSIdriver. Consult your driver's documentation for supported values.",
													MarkdownDescription: "volumeAttributes stores driver-specific properties that are passed to the CSIdriver. Consult your driver's documentation for supported values.",
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
													Description:         "Optional: mode bits to use on created files by default. Must be aOptional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
													MarkdownDescription: "Optional: mode bits to use on created files by default. Must be aOptional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal values for mode bits.Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
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

										"empty_dir": schema.SingleNestedAttribute{
											Description:         "emptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
											MarkdownDescription: "emptyDir represents a temporary directory that shares a pod's lifetime.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
											Attributes: map[string]schema.Attribute{
												"medium": schema.StringAttribute{
													Description:         "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
													MarkdownDescription: "medium represents what type of storage medium should back this directory.The default is '' which means to use the node's default medium.Must be an empty string (default) or Memory.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"size_limit": schema.StringAttribute{
													Description:         "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
													MarkdownDescription: "sizeLimit is the total amount of local storage required for this EmptyDir volume.The size limit is also applicable for memory medium.The maximum usage on memory medium EmptyDir would be the minimum value betweenthe SizeLimit specified here and the sum of memory limits of all containers in a pod.The default is nil which means that the limit is undefined.More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
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
											Description:         "ephemeral represents a volume that is handled by a cluster storage driver.The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,and deleted when the pod is removed.Use this if:a) the volume is only needed while the pod runs,b) features of normal volumes like restoring from snapshot or capacity   tracking are needed,c) the storage driver is specified through a storage class, andd) the storage driver supports dynamic volume provisioning through   a PersistentVolumeClaim (see EphemeralVolumeSource for more   information on the connection between this volume type   and PersistentVolumeClaim).Use PersistentVolumeClaim or one of the vendor-specificAPIs for volumes that persist for longer than the lifecycleof an individual pod.Use CSI for light-weight local ephemeral volumes if the CSI driver is meant tobe used that way - see the documentation of the driver formore information.A pod can use both types of ephemeral volumes andpersistent volumes at the same time.",
											MarkdownDescription: "ephemeral represents a volume that is handled by a cluster storage driver.The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts,and deleted when the pod is removed.Use this if:a) the volume is only needed while the pod runs,b) features of normal volumes like restoring from snapshot or capacity   tracking are needed,c) the storage driver is specified through a storage class, andd) the storage driver supports dynamic volume provisioning through   a PersistentVolumeClaim (see EphemeralVolumeSource for more   information on the connection between this volume type   and PersistentVolumeClaim).Use PersistentVolumeClaim or one of the vendor-specificAPIs for volumes that persist for longer than the lifecycleof an individual pod.Use CSI for light-weight local ephemeral volumes if the CSI driver is meant tobe used that way - see the documentation of the driver formore information.A pod can use both types of ephemeral volumes andpersistent volumes at the same time.",
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

										"fc": schema.SingleNestedAttribute{
											Description:         "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
											MarkdownDescription: "fc represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.TODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.TODO: how do we prevent errors in the filesystem from compromising the machine",
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
													Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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
													Description:         "wwids Optional: FC volume world wide identifiers (wwids)Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
													MarkdownDescription: "wwids Optional: FC volume world wide identifiers (wwids)Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
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
											Description:         "flexVolume represents a generic volume resource that isprovisioned/attached using an exec based plugin.",
											MarkdownDescription: "flexVolume represents a generic volume resource that isprovisioned/attached using an exec based plugin.",
											Attributes: map[string]schema.Attribute{
												"driver": schema.StringAttribute{
													Description:         "driver is the name of the driver to use for this volume.",
													MarkdownDescription: "driver is the name of the driver to use for this volume.",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
													MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
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
													Description:         "readOnly is Optional: defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly is Optional: defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is Optional: secretRef is reference to the secret object containingsensitive information to pass to the plugin scripts. This may beempty if no secret object is specified. If the secret objectcontains more than one secret, all secrets are passed to the pluginscripts.",
													MarkdownDescription: "secretRef is Optional: secretRef is reference to the secret object containingsensitive information to pass to the plugin scripts. This may beempty if no secret object is specified. If the secret objectcontains more than one secret, all secrets are passed to the pluginscripts.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
													Description:         "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flockershould be considered as deprecated",
													MarkdownDescription: "datasetName is Name of the dataset stored as metadata -> name on the dataset for Flockershould be considered as deprecated",
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
											Description:         "gcePersistentDisk represents a GCE Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											MarkdownDescription: "gcePersistentDisk represents a GCE Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdiskTODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdiskTODO: how do we prevent errors in the filesystem from compromising the machine",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"partition": schema.Int64Attribute{
													Description:         "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													MarkdownDescription: "partition is the partition in the volume that you want to mount.If omitted, the default is to mount by volume name.Examples: For volume /dev/sda1, you specify the partition as '1'.Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"pd_name": schema.StringAttribute{
													Description:         "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													MarkdownDescription: "pdName is unique name of the PD resource in GCE. Used to identify the disk in GCE.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
													MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
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
											Description:         "gitRepo represents a git repository at a particular revision.DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount anEmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDirinto the Pod's container.",
											MarkdownDescription: "gitRepo represents a git repository at a particular revision.DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount anEmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDirinto the Pod's container.",
											Attributes: map[string]schema.Attribute{
												"directory": schema.StringAttribute{
													Description:         "directory is the target directory name.Must not contain or start with '..'.  If '.' is supplied, the volume directory will be thegit repository.  Otherwise, if specified, the volume will contain the git repository inthe subdirectory with the given name.",
													MarkdownDescription: "directory is the target directory name.Must not contain or start with '..'.  If '.' is supplied, the volume directory will be thegit repository.  Otherwise, if specified, the volume will contain the git repository inthe subdirectory with the given name.",
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
											Description:         "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/glusterfs/README.md",
											MarkdownDescription: "glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/glusterfs/README.md",
											Attributes: map[string]schema.Attribute{
												"endpoints": schema.StringAttribute{
													Description:         "endpoints is the endpoint name that details Glusterfs topology.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"path": schema.StringAttribute{
													Description:         "path is the Glusterfs volume path.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													MarkdownDescription: "path is the Glusterfs volume path.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the Glusterfs volume to be mounted with read-only permissions.Defaults to false.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
													MarkdownDescription: "readOnly here will force the Glusterfs volume to be mounted with read-only permissions.Defaults to false.More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
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
											Description:         "hostPath represents a pre-existing file or directory on the hostmachine that is directly exposed to the container. This is generallyused for system agents or other privileged things that are allowedto see the host machine. Most containers will NOT need this.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath---TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can notmount host directories as read/write.",
											MarkdownDescription: "hostPath represents a pre-existing file or directory on the hostmachine that is directly exposed to the container. This is generallyused for system agents or other privileged things that are allowedto see the host machine. Most containers will NOT need this.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath---TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can notmount host directories as read/write.",
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description:         "path of the directory on the host.If the path is a symlink, it will follow the link to the real path.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
													MarkdownDescription: "path of the directory on the host.If the path is a symlink, it will follow the link to the real path.More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"type": schema.StringAttribute{
													Description:         "type for HostPath VolumeDefaults to ''More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
													MarkdownDescription: "type for HostPath VolumeDefaults to ''More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
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
											Description:         "iscsi represents an ISCSI Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://examples.k8s.io/volumes/iscsi/README.md",
											MarkdownDescription: "iscsi represents an ISCSI Disk resource that is attached to akubelet's host machine and then exposed to the pod.More info: https://examples.k8s.io/volumes/iscsi/README.md",
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
													Description:         "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsiTODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsiTODO: how do we prevent errors in the filesystem from compromising the machine",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"initiator_name": schema.StringAttribute{
													Description:         "initiatorName is the custom iSCSI Initiator Name.If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<target portal>:<volume name> will be created for the connection.",
													MarkdownDescription: "initiatorName is the custom iSCSI Initiator Name.If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface<target portal>:<volume name> will be created for the connection.",
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
													Description:         "iscsiInterface is the interface Name that uses an iSCSI transport.Defaults to 'default' (tcp).",
													MarkdownDescription: "iscsiInterface is the interface Name that uses an iSCSI transport.Defaults to 'default' (tcp).",
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
													Description:         "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
													MarkdownDescription: "portals is the iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.",
													MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
													MarkdownDescription: "secretRef is the CHAP Secret for iSCSI target and initiator authentication",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
													Description:         "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
													MarkdownDescription: "targetPortal is iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the portis other than default (typically TCP ports 860 and 3260).",
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
											Description:         "name of the volume.Must be a DNS_LABEL and unique within the pod.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											MarkdownDescription: "name of the volume.Must be a DNS_LABEL and unique within the pod.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
											Required:            true,
											Optional:            false,
											Computed:            false,
										},

										"nfs": schema.SingleNestedAttribute{
											Description:         "nfs represents an NFS mount on the host that shares a pod's lifetimeMore info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											MarkdownDescription: "nfs represents an NFS mount on the host that shares a pod's lifetimeMore info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
											Attributes: map[string]schema.Attribute{
												"path": schema.StringAttribute{
													Description:         "path that is exported by the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													MarkdownDescription: "path that is exported by the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the NFS export to be mounted with read-only permissions.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													MarkdownDescription: "readOnly here will force the NFS export to be mounted with read-only permissions.Defaults to false.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"server": schema.StringAttribute{
													Description:         "server is the hostname or IP address of the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
													MarkdownDescription: "server is the hostname or IP address of the NFS server.More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
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
											Description:         "persistentVolumeClaimVolumeSource represents a reference to aPersistentVolumeClaim in the same namespace.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											MarkdownDescription: "persistentVolumeClaimVolumeSource represents a reference to aPersistentVolumeClaim in the same namespace.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
											Attributes: map[string]schema.Attribute{
												"claim_name": schema.StringAttribute{
													Description:         "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
													MarkdownDescription: "claimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume.More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
													MarkdownDescription: "readOnly Will force the ReadOnly setting in VolumeMounts.Default false.",
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
													Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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
													Description:         "fSType represents the filesystem type to mountMust be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fSType represents the filesystem type to mountMust be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
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

										"quobyte": schema.SingleNestedAttribute{
											Description:         "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
											MarkdownDescription: "quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
											Attributes: map[string]schema.Attribute{
												"group": schema.StringAttribute{
													Description:         "group to map volume access toDefault is no group",
													MarkdownDescription: "group to map volume access toDefault is no group",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the Quobyte volume to be mounted with read-only permissions.Defaults to false.",
													MarkdownDescription: "readOnly here will force the Quobyte volume to be mounted with read-only permissions.Defaults to false.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"registry": schema.StringAttribute{
													Description:         "registry represents a single or multiple Quobyte Registry servicesspecified as a string as host:port pair (multiple entries are separated with commas)which acts as the central registry for volumes",
													MarkdownDescription: "registry represents a single or multiple Quobyte Registry servicesspecified as a string as host:port pair (multiple entries are separated with commas)which acts as the central registry for volumes",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"tenant": schema.StringAttribute{
													Description:         "tenant owning the given Quobyte volume in the BackendUsed with dynamically provisioned Quobyte volumes, value is set by the plugin",
													MarkdownDescription: "tenant owning the given Quobyte volume in the BackendUsed with dynamically provisioned Quobyte volumes, value is set by the plugin",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"user": schema.StringAttribute{
													Description:         "user to map volume access toDefaults to serivceaccount user",
													MarkdownDescription: "user to map volume access toDefaults to serivceaccount user",
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
											Description:         "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/rbd/README.md",
											MarkdownDescription: "rbd represents a Rados Block Device mount on the host that shares a pod's lifetime.More info: https://examples.k8s.io/volumes/rbd/README.md",
											Attributes: map[string]schema.Attribute{
												"fs_type": schema.StringAttribute{
													Description:         "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#rbdTODO: how do we prevent errors in the filesystem from compromising the machine",
													MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount.Tip: Ensure that the filesystem type is supported by the host operating system.Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.More info: https://kubernetes.io/docs/concepts/storage/volumes#rbdTODO: how do we prevent errors in the filesystem from compromising the machine",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"image": schema.StringAttribute{
													Description:         "image is the rados image name.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "image is the rados image name.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"keyring": schema.StringAttribute{
													Description:         "keyring is the path to key ring for RBDUser.Default is /etc/ceph/keyring.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "keyring is the path to key ring for RBDUser.Default is /etc/ceph/keyring.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"monitors": schema.ListAttribute{
													Description:         "monitors is a collection of Ceph monitors.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "monitors is a collection of Ceph monitors.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													ElementType:         types.StringType,
													Required:            true,
													Optional:            false,
													Computed:            false,
												},

												"pool": schema.StringAttribute{
													Description:         "pool is the rados pool name.Default is rbd.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "pool is the rados pool name.Default is rbd.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "readOnly here will force the ReadOnly setting in VolumeMounts.Defaults to false.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef is name of the authentication secret for RBDUser. If providedoverrides keyring.Default is nil.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "secretRef is name of the authentication secret for RBDUser. If providedoverrides keyring.Default is nil.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
													Description:         "user is the rados user name.Default is admin.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
													MarkdownDescription: "user is the rados user name.Default is admin.More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
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
													Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'.Default is 'xfs'.",
													MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'.Default is 'xfs'.",
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
													Description:         "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly Defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef references to the secret for ScaleIO user and othersensitive information. If this is not provided, Login operation will fail.",
													MarkdownDescription: "secretRef references to the secret for ScaleIO user and othersensitive information. If this is not provided, Login operation will fail.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
													Description:         "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.Default is ThinProvisioned.",
													MarkdownDescription: "storageMode indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned.Default is ThinProvisioned.",
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
													Description:         "volumeName is the name of a volume already created in the ScaleIO systemthat is associated with this volume source.",
													MarkdownDescription: "volumeName is the name of a volume already created in the ScaleIO systemthat is associated with this volume source.",
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
											Description:         "secret represents a secret that should populate this volume.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
											MarkdownDescription: "secret represents a secret that should populate this volume.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
											Attributes: map[string]schema.Attribute{
												"default_mode": schema.Int64Attribute{
													Description:         "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
													MarkdownDescription: "defaultMode is Optional: mode bits used to set permissions on created files by default.Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511.YAML accepts both octal and decimal values, JSON requires decimal valuesfor mode bits. Defaults to 0644.Directories within the path are not affected by this setting.This might be in conflict with other options that affect the filemode, like fsGroup, and the result can be other mode bits set.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"items": schema.ListNestedAttribute{
													Description:         "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
													MarkdownDescription: "items If unspecified, each key-value pair in the Data field of the referencedSecret will be projected into the volume as a file whose name is thekey and content is the value. If specified, the listed keys will beprojected into the specified paths, and unlisted keys will not bepresent. If a key is specified which is not present in the Secret,the volume setup will error unless it is marked optional. Paths must berelative and may not contain the '..' path or start with '..'.",
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

												"optional": schema.BoolAttribute{
													Description:         "optional field specify whether the Secret or its keys must be defined",
													MarkdownDescription: "optional field specify whether the Secret or its keys must be defined",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_name": schema.StringAttribute{
													Description:         "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
													MarkdownDescription: "secretName is the name of the secret in the pod's namespace to use.More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
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
													Description:         "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fsType is the filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"read_only": schema.BoolAttribute{
													Description:         "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will forcethe ReadOnly setting in VolumeMounts.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secret_ref": schema.SingleNestedAttribute{
													Description:         "secretRef specifies the secret to use for obtaining the StorageOS APIcredentials.  If not specified, default values will be attempted.",
													MarkdownDescription: "secretRef specifies the secret to use for obtaining the StorageOS APIcredentials.  If not specified, default values will be attempted.",
													Attributes: map[string]schema.Attribute{
														"name": schema.StringAttribute{
															Description:         "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
															MarkdownDescription: "Name of the referent.This field is effectively required, but due to backwards compatibility isallowed to be empty. Instances of this type with an empty value here arealmost certainly wrong.TODO: Add other useful fields. apiVersion, kind, uid?More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#namesTODO: Drop 'kubebuilder:default' when controller-gen doesn't need it https://github.com/kubernetes-sigs/kubebuilder/issues/3896.",
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
													Description:         "volumeName is the human-readable name of the StorageOS volume.  Volumenames are only unique within a namespace.",
													MarkdownDescription: "volumeName is the human-readable name of the StorageOS volume.  Volumenames are only unique within a namespace.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"volume_namespace": schema.StringAttribute{
													Description:         "volumeNamespace specifies the scope of the volume within StorageOS.  If nonamespace is specified then the Pod's namespace will be used.  This allows theKubernetes name scoping to be mirrored within StorageOS for tighter integration.Set VolumeName to any name to override the default behaviour.Set to 'default' if you are not using namespaces within StorageOS.Namespaces that do not pre-exist within StorageOS will be created.",
													MarkdownDescription: "volumeNamespace specifies the scope of the volume within StorageOS.  If nonamespace is specified then the Pod's namespace will be used.  This allows theKubernetes name scoping to be mirrored within StorageOS for tighter integration.Set VolumeName to any name to override the default behaviour.Set to 'default' if you are not using namespaces within StorageOS.Namespaces that do not pre-exist within StorageOS will be created.",
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
													Description:         "fsType is filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
													MarkdownDescription: "fsType is filesystem type to mount.Must be a filesystem type supported by the host operating system.Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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

							"log_level": schema.StringAttribute{
								Description:         "LogLevel sets the log level for Envoy.Allowed values are 'trace', 'debug', 'info', 'warn', 'error', 'critical', 'off'.",
								MarkdownDescription: "LogLevel sets the log level for Envoy.Allowed values are 'trace', 'debug', 'info', 'warn', 'error', 'critical', 'off'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"network_publishing": schema.SingleNestedAttribute{
								Description:         "NetworkPublishing defines how to expose Envoy to a network.",
								MarkdownDescription: "NetworkPublishing defines how to expose Envoy to a network.",
								Attributes: map[string]schema.Attribute{
									"external_traffic_policy": schema.StringAttribute{
										Description:         "ExternalTrafficPolicy describes how nodes distribute service traffic theyreceive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs,and LoadBalancer IPs).If unset, defaults to 'Local'.",
										MarkdownDescription: "ExternalTrafficPolicy describes how nodes distribute service traffic theyreceive on one of the Service's 'externally-facing' addresses (NodePorts, ExternalIPs,and LoadBalancer IPs).If unset, defaults to 'Local'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"ip_family_policy": schema.StringAttribute{
										Description:         "IPFamilyPolicy represents the dual-stack-ness requested or required bythis Service. If there is no value provided, then this field will be setto SingleStack. Services can be 'SingleStack' (a single IP family),'PreferDualStack' (two IP families on dual-stack configured clusters ora single IP family on single-stack clusters), or 'RequireDualStack'(two IP families on dual-stack configured clusters, otherwise fail).",
										MarkdownDescription: "IPFamilyPolicy represents the dual-stack-ness requested or required bythis Service. If there is no value provided, then this field will be setto SingleStack. Services can be 'SingleStack' (a single IP family),'PreferDualStack' (two IP families on dual-stack configured clusters ora single IP family on single-stack clusters), or 'RequireDualStack'(two IP families on dual-stack configured clusters, otherwise fail).",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_annotations": schema.MapAttribute{
										Description:         "ServiceAnnotations is the annotations to add tothe provisioned Envoy service.",
										MarkdownDescription: "ServiceAnnotations is the annotations to add tothe provisioned Envoy service.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"type": schema.StringAttribute{
										Description:         "NetworkPublishingType is the type of publishing strategy to use. Valid values are:* LoadBalancerServiceIn this configuration, network endpoints for Envoy use container networking.A Kubernetes LoadBalancer Service is created to publish Envoy networkendpoints.See: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer* NodePortServicePublishes Envoy network endpoints using a Kubernetes NodePort Service.In this configuration, Envoy network endpoints use container networking. A KubernetesNodePort Service is created to publish the network endpoints.See: https://kubernetes.io/docs/concepts/services-networking/service/#nodeportNOTE:When provisioning an Envoy 'NodePortService', use Gateway Listeners' port numbers to populatethe Service's node port values, there's no way to auto-allocate them.See: https://github.com/projectcontour/contour/issues/4499* ClusterIPServicePublishes Envoy network endpoints using a Kubernetes ClusterIP Service.In this configuration, Envoy network endpoints use container networking. A KubernetesClusterIP Service is created to publish the network endpoints.See: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-typesIf unset, defaults to LoadBalancerService.",
										MarkdownDescription: "NetworkPublishingType is the type of publishing strategy to use. Valid values are:* LoadBalancerServiceIn this configuration, network endpoints for Envoy use container networking.A Kubernetes LoadBalancer Service is created to publish Envoy networkendpoints.See: https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer* NodePortServicePublishes Envoy network endpoints using a Kubernetes NodePort Service.In this configuration, Envoy network endpoints use container networking. A KubernetesNodePort Service is created to publish the network endpoints.See: https://kubernetes.io/docs/concepts/services-networking/service/#nodeportNOTE:When provisioning an Envoy 'NodePortService', use Gateway Listeners' port numbers to populatethe Service's node port values, there's no way to auto-allocate them.See: https://github.com/projectcontour/contour/issues/4499* ClusterIPServicePublishes Envoy network endpoints using a Kubernetes ClusterIP Service.In this configuration, Envoy network endpoints use container networking. A KubernetesClusterIP Service is created to publish the network endpoints.See: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-typesIf unset, defaults to LoadBalancerService.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_placement": schema.SingleNestedAttribute{
								Description:         "NodePlacement describes node scheduling configuration of Envoy pods.",
								MarkdownDescription: "NodePlacement describes node scheduling configuration of Envoy pods.",
								Attributes: map[string]schema.Attribute{
									"node_selector": schema.MapAttribute{
										Description:         "NodeSelector is the simplest recommended form of node selection constraintand specifies a map of key-value pairs. For the pod to be eligibleto run on a node, the node must have each of the indicated key-value pairsas labels (it can have additional labels as well).If unset, the pod(s) will be scheduled to any available node.",
										MarkdownDescription: "NodeSelector is the simplest recommended form of node selection constraintand specifies a map of key-value pairs. For the pod to be eligibleto run on a node, the node must have each of the indicated key-value pairsas labels (it can have additional labels as well).If unset, the pod(s) will be scheduled to any available node.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Tolerations work with taints to ensure that pods are not scheduledonto inappropriate nodes. One or more taints are applied to a node; thismarks that the node should not accept any pods that do not tolerate thetaints.The default is an empty list.See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/for additional details.",
										MarkdownDescription: "Tolerations work with taints to ensure that pods are not scheduledonto inappropriate nodes. One or more taints are applied to a node; thismarks that the node should not accept any pods that do not tolerate thetaints.The default is an empty list.See https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/for additional details.",
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"overload_max_heap_size": schema.Int64Attribute{
								Description:         "OverloadMaxHeapSize defines the maximum heap memory of the envoy controlled by the overload manager.When the value is greater than 0, the overload manager is enabled,and when envoy reaches 95% of the maximum heap size, it performs a shrink heap operation,When it reaches 98% of the maximum heap size, Envoy Will stop accepting requests.More info: https://projectcontour.io/docs/main/config/overload-manager/",
								MarkdownDescription: "OverloadMaxHeapSize defines the maximum heap memory of the envoy controlled by the overload manager.When the value is greater than 0, the overload manager is enabled,and when envoy reaches 95% of the maximum heap size, it performs a shrink heap operation,When it reaches 98% of the maximum heap size, Envoy Will stop accepting requests.More info: https://projectcontour.io/docs/main/config/overload-manager/",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"pod_annotations": schema.MapAttribute{
								Description:         "PodAnnotations defines annotations to add to the Envoy pods.the annotations for Prometheus will be appended or overwritten with predefined value.",
								MarkdownDescription: "PodAnnotations defines annotations to add to the Envoy pods.the annotations for Prometheus will be appended or overwritten with predefined value.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"replicas": schema.Int64Attribute{
								Description:         "Deprecated: Use 'DeploymentSettings.Replicas' instead.Replicas is the desired number of Envoy replicas. If WorkloadTypeis not 'Deployment', this field is ignored. Otherwise, if unset,defaults to 2.if both 'DeploymentSettings.Replicas' and this one is set, use 'DeploymentSettings.Replicas'.",
								MarkdownDescription: "Deprecated: Use 'DeploymentSettings.Replicas' instead.Replicas is the desired number of Envoy replicas. If WorkloadTypeis not 'Deployment', this field is ignored. Otherwise, if unset,defaults to 2.if both 'DeploymentSettings.Replicas' and this one is set, use 'DeploymentSettings.Replicas'.",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.Int64{
									int64validator.AtLeast(0),
								},
							},

							"resources": schema.SingleNestedAttribute{
								Description:         "Compute Resources required by envoy container.Cannot be updated.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
								MarkdownDescription: "Compute Resources required by envoy container.Cannot be updated.More info: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/",
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

							"workload_type": schema.StringAttribute{
								Description:         "WorkloadType is the type of workload to install Envoyas. Choices are DaemonSet and Deployment. If unset, defaultsto DaemonSet.",
								MarkdownDescription: "WorkloadType is the type of workload to install Envoyas. Choices are DaemonSet and Deployment. If unset, defaultsto DaemonSet.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"resource_labels": schema.MapAttribute{
						Description:         "ResourceLabels is a set of labels to add to the provisioned Contour resources.Deprecated: use Gateway.Spec.Infrastructure.Labels instead. This field will beremoved in a future release.",
						MarkdownDescription: "ResourceLabels is a set of labels to add to the provisioned Contour resources.Deprecated: use Gateway.Spec.Infrastructure.Labels instead. This field will beremoved in a future release.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"runtime_settings": schema.SingleNestedAttribute{
						Description:         "RuntimeSettings is a ContourConfiguration spec to be used whenprovisioning a Contour instance that will influence aspects ofthe Contour instance's runtime behavior.",
						MarkdownDescription: "RuntimeSettings is a ContourConfiguration spec to be used whenprovisioning a Contour instance that will influence aspects ofthe Contour instance's runtime behavior.",
						Attributes: map[string]schema.Attribute{
							"debug": schema.SingleNestedAttribute{
								Description:         "Debug contains parameters to enable debug loggingand debug interfaces inside Contour.",
								MarkdownDescription: "Debug contains parameters to enable debug loggingand debug interfaces inside Contour.",
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description:         "Defines the Contour debug address interface.Contour's default is '127.0.0.1'.",
										MarkdownDescription: "Defines the Contour debug address interface.Contour's default is '127.0.0.1'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"port": schema.Int64Attribute{
										Description:         "Defines the Contour debug address port.Contour's default is 6060.",
										MarkdownDescription: "Defines the Contour debug address port.Contour's default is 6060.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"enable_external_name_service": schema.BoolAttribute{
								Description:         "EnableExternalNameService allows processing of ExternalNameServicesContour's default is false for security reasons.",
								MarkdownDescription: "EnableExternalNameService allows processing of ExternalNameServicesContour's default is false for security reasons.",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"envoy": schema.SingleNestedAttribute{
								Description:         "Envoy contains parameters for Envoy as wellas how to optionally configure a managed Envoy fleet.",
								MarkdownDescription: "Envoy contains parameters for Envoy as wellas how to optionally configure a managed Envoy fleet.",
								Attributes: map[string]schema.Attribute{
									"client_certificate": schema.SingleNestedAttribute{
										Description:         "ClientCertificate defines the namespace/name of the Kubernetessecret containing the client certificate and private keyto be used when establishing TLS connection to upstreamcluster.",
										MarkdownDescription: "ClientCertificate defines the namespace/name of the Kubernetessecret containing the client certificate and private keyto be used when establishing TLS connection to upstreamcluster.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
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

									"cluster": schema.SingleNestedAttribute{
										Description:         "Cluster holds various configurable Envoy cluster values that canbe set in the config file.",
										MarkdownDescription: "Cluster holds various configurable Envoy cluster values that canbe set in the config file.",
										Attributes: map[string]schema.Attribute{
											"circuit_breakers": schema.SingleNestedAttribute{
												Description:         "GlobalCircuitBreakerDefaults specifies default circuit breaker budget across all services.If defined, this will be used as the default for all services.",
												MarkdownDescription: "GlobalCircuitBreakerDefaults specifies default circuit breaker budget across all services.If defined, this will be used as the default for all services.",
												Attributes: map[string]schema.Attribute{
													"max_connections": schema.Int64Attribute{
														Description:         "The maximum number of connections that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
														MarkdownDescription: "The maximum number of connections that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_pending_requests": schema.Int64Attribute{
														Description:         "The maximum number of pending requests that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
														MarkdownDescription: "The maximum number of pending requests that a single Envoy instance allows to the Kubernetes Service; defaults to 1024.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_requests": schema.Int64Attribute{
														Description:         "The maximum parallel requests a single Envoy instance allows to the Kubernetes Service; defaults to 1024",
														MarkdownDescription: "The maximum parallel requests a single Envoy instance allows to the Kubernetes Service; defaults to 1024",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"max_retries": schema.Int64Attribute{
														Description:         "The maximum number of parallel retries a single Envoy instance allows to the Kubernetes Service; defaults to 3.",
														MarkdownDescription: "The maximum number of parallel retries a single Envoy instance allows to the Kubernetes Service; defaults to 3.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"dns_lookup_family": schema.StringAttribute{
												Description:         "DNSLookupFamily defines how external names are looked upWhen configured as V4, the DNS resolver will only perform a lookupfor addresses in the IPv4 family. If V6 is configured, the DNS resolverwill only perform a lookup for addresses in the IPv6 family.If AUTO is configured, the DNS resolver will first perform a lookupfor addresses in the IPv6 family and fallback to a lookup for addressesin the IPv4 family. If ALL is specified, the DNS resolver will perform a lookup forboth IPv4 and IPv6 families, and return all resolved addresses.When this is used, Happy Eyeballs will be enabled for upstream connections.Refer to Happy Eyeballs Support for more information.Note: This only applies to externalName clusters.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html#envoy-v3-api-enum-config-cluster-v3-cluster-dnslookupfamilyfor more information.Values: 'auto' (default), 'v4', 'v6', 'all'.Other values will produce an error.",
												MarkdownDescription: "DNSLookupFamily defines how external names are looked upWhen configured as V4, the DNS resolver will only perform a lookupfor addresses in the IPv4 family. If V6 is configured, the DNS resolverwill only perform a lookup for addresses in the IPv6 family.If AUTO is configured, the DNS resolver will first perform a lookupfor addresses in the IPv6 family and fallback to a lookup for addressesin the IPv4 family. If ALL is specified, the DNS resolver will perform a lookup forboth IPv4 and IPv6 families, and return all resolved addresses.When this is used, Happy Eyeballs will be enabled for upstream connections.Refer to Happy Eyeballs Support for more information.Note: This only applies to externalName clusters.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto.html#envoy-v3-api-enum-config-cluster-v3-cluster-dnslookupfamilyfor more information.Values: 'auto' (default), 'v4', 'v6', 'all'.Other values will produce an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_requests_per_connection": schema.Int64Attribute{
												Description:         "Defines the maximum requests for upstream connections. If not specified, there is no limit.see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-msg-config-core-v3-httpprotocoloptionsfor more information.",
												MarkdownDescription: "Defines the maximum requests for upstream connections. If not specified, there is no limit.see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-msg-config-core-v3-httpprotocoloptionsfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"per_connection_buffer_limit_bytes": schema.Int64Attribute{
												Description:         "Defines the soft limit on size of the cluster’s new connection read and write buffers in bytes.If unspecified, an implementation defined default is applied (1MiB).see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto#envoy-v3-api-field-config-cluster-v3-cluster-per-connection-buffer-limit-bytesfor more information.",
												MarkdownDescription: "Defines the soft limit on size of the cluster’s new connection read and write buffers in bytes.If unspecified, an implementation defined default is applied (1MiB).see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto#envoy-v3-api-field-config-cluster-v3-cluster-per-connection-buffer-limit-bytesfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"upstream_tls": schema.SingleNestedAttribute{
												Description:         "UpstreamTLS contains the TLS policy parameters for upstream connections",
												MarkdownDescription: "UpstreamTLS contains the TLS policy parameters for upstream connections",
												Attributes: map[string]schema.Attribute{
													"cipher_suites": schema.ListAttribute{
														Description:         "CipherSuites defines the TLS ciphers to be supported by Envoy TLSlisteners when negotiating TLS 1.2. Ciphers are validated against theset that Envoy supports by default. This parameter should only be usedby advanced users. Note that these will be ignored when TLS 1.3 is inuse.This field is optional; when it is undefined, a Contour-managed ciphersuite listwill be used, which may be updated to keep it secure.Contour's default list is:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'Ciphers provided are validated against the following list:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES128-GCM-SHA256'  - 'ECDHE-RSA-AES128-GCM-SHA256'  - 'ECDHE-ECDSA-AES128-SHA'  - 'ECDHE-RSA-AES128-SHA'  - 'AES128-GCM-SHA256'  - 'AES128-SHA'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'  - 'ECDHE-ECDSA-AES256-SHA'  - 'ECDHE-RSA-AES256-SHA'  - 'AES256-GCM-SHA384'  - 'AES256-SHA'Contour recommends leaving this undefined unless you are sure you must.See: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#extensions-transport-sockets-tls-v3-tlsparametersNote: This list is a superset of what is valid for stock Envoy builds and those using BoringSSL FIPS.",
														MarkdownDescription: "CipherSuites defines the TLS ciphers to be supported by Envoy TLSlisteners when negotiating TLS 1.2. Ciphers are validated against theset that Envoy supports by default. This parameter should only be usedby advanced users. Note that these will be ignored when TLS 1.3 is inuse.This field is optional; when it is undefined, a Contour-managed ciphersuite listwill be used, which may be updated to keep it secure.Contour's default list is:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'Ciphers provided are validated against the following list:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES128-GCM-SHA256'  - 'ECDHE-RSA-AES128-GCM-SHA256'  - 'ECDHE-ECDSA-AES128-SHA'  - 'ECDHE-RSA-AES128-SHA'  - 'AES128-GCM-SHA256'  - 'AES128-SHA'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'  - 'ECDHE-ECDSA-AES256-SHA'  - 'ECDHE-RSA-AES256-SHA'  - 'AES256-GCM-SHA384'  - 'AES256-SHA'Contour recommends leaving this undefined unless you are sure you must.See: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#extensions-transport-sockets-tls-v3-tlsparametersNote: This list is a superset of what is valid for stock Envoy builds and those using BoringSSL FIPS.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"maximum_protocol_version": schema.StringAttribute{
														Description:         "MaximumProtocolVersion is the maximum TLS version this vhost shouldnegotiate.Values: '1.2', '1.3'(default).Other values will produce an error.",
														MarkdownDescription: "MaximumProtocolVersion is the maximum TLS version this vhost shouldnegotiate.Values: '1.2', '1.3'(default).Other values will produce an error.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"minimum_protocol_version": schema.StringAttribute{
														Description:         "MinimumProtocolVersion is the minimum TLS version this vhost shouldnegotiate.Values: '1.2' (default), '1.3'.Other values will produce an error.",
														MarkdownDescription: "MinimumProtocolVersion is the minimum TLS version this vhost shouldnegotiate.Values: '1.2' (default), '1.3'.Other values will produce an error.",
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

									"default_http_versions": schema.ListAttribute{
										Description:         "DefaultHTTPVersions defines the default set of HTTPSversions the proxy should accept. HTTP versions arestrings of the form 'HTTP/xx'. Supported versions are'HTTP/1.1' and 'HTTP/2'.Values: 'HTTP/1.1', 'HTTP/2' (default: both).Other values will produce an error.",
										MarkdownDescription: "DefaultHTTPVersions defines the default set of HTTPSversions the proxy should accept. HTTP versions arestrings of the form 'HTTP/xx'. Supported versions are'HTTP/1.1' and 'HTTP/2'.Values: 'HTTP/1.1', 'HTTP/2' (default: both).Other values will produce an error.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"health": schema.SingleNestedAttribute{
										Description:         "Health defines the endpoint Envoy uses to serve health checks.Contour's default is { address: '0.0.0.0', port: 8002 }.",
										MarkdownDescription: "Health defines the endpoint Envoy uses to serve health checks.Contour's default is { address: '0.0.0.0', port: 8002 }.",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "Defines the health address interface.",
												MarkdownDescription: "Defines the health address interface.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "Defines the health port.",
												MarkdownDescription: "Defines the health port.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"http": schema.SingleNestedAttribute{
										Description:         "Defines the HTTP Listener for Envoy.Contour's default is { address: '0.0.0.0', port: 8080, accessLog: '/dev/stdout' }.",
										MarkdownDescription: "Defines the HTTP Listener for Envoy.Contour's default is { address: '0.0.0.0', port: 8080, accessLog: '/dev/stdout' }.",
										Attributes: map[string]schema.Attribute{
											"access_log": schema.StringAttribute{
												Description:         "AccessLog defines where Envoy logs are outputted for this listener.",
												MarkdownDescription: "AccessLog defines where Envoy logs are outputted for this listener.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"address": schema.StringAttribute{
												Description:         "Defines an Envoy Listener Address.",
												MarkdownDescription: "Defines an Envoy Listener Address.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "Defines an Envoy listener Port.",
												MarkdownDescription: "Defines an Envoy listener Port.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"https": schema.SingleNestedAttribute{
										Description:         "Defines the HTTPS Listener for Envoy.Contour's default is { address: '0.0.0.0', port: 8443, accessLog: '/dev/stdout' }.",
										MarkdownDescription: "Defines the HTTPS Listener for Envoy.Contour's default is { address: '0.0.0.0', port: 8443, accessLog: '/dev/stdout' }.",
										Attributes: map[string]schema.Attribute{
											"access_log": schema.StringAttribute{
												Description:         "AccessLog defines where Envoy logs are outputted for this listener.",
												MarkdownDescription: "AccessLog defines where Envoy logs are outputted for this listener.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"address": schema.StringAttribute{
												Description:         "Defines an Envoy Listener Address.",
												MarkdownDescription: "Defines an Envoy Listener Address.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "Defines an Envoy listener Port.",
												MarkdownDescription: "Defines an Envoy listener Port.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"listener": schema.SingleNestedAttribute{
										Description:         "Listener hold various configurable Envoy listener values.",
										MarkdownDescription: "Listener hold various configurable Envoy listener values.",
										Attributes: map[string]schema.Attribute{
											"connection_balancer": schema.StringAttribute{
												Description:         "ConnectionBalancer. If the value is exact, the listener will use the exact connection balancerSee https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/listener.proto#envoy-api-msg-listener-connectionbalanceconfigfor more information.Values: (empty string): use the default ConnectionBalancer, 'exact': use the Exact ConnectionBalancer.Other values will produce an error.",
												MarkdownDescription: "ConnectionBalancer. If the value is exact, the listener will use the exact connection balancerSee https://www.envoyproxy.io/docs/envoy/latest/api-v2/api/v2/listener.proto#envoy-api-msg-listener-connectionbalanceconfigfor more information.Values: (empty string): use the default ConnectionBalancer, 'exact': use the Exact ConnectionBalancer.Other values will produce an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"disable_allow_chunked_length": schema.BoolAttribute{
												Description:         "DisableAllowChunkedLength disables the RFC-compliant Envoy behavior tostrip the 'Content-Length' header if 'Transfer-Encoding: chunked' isalso set. This is an emergency off-switch to revert back to Envoy'sdefault behavior in case of failures. Please file an issue if failuresare encountered.See: https://github.com/projectcontour/contour/issues/3221Contour's default is false.",
												MarkdownDescription: "DisableAllowChunkedLength disables the RFC-compliant Envoy behavior tostrip the 'Content-Length' header if 'Transfer-Encoding: chunked' isalso set. This is an emergency off-switch to revert back to Envoy'sdefault behavior in case of failures. Please file an issue if failuresare encountered.See: https://github.com/projectcontour/contour/issues/3221Contour's default is false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"disable_merge_slashes": schema.BoolAttribute{
												Description:         "DisableMergeSlashes disables Envoy's non-standard merge_slashes path transformation optionwhich strips duplicate slashes from request URL paths.Contour's default is false.",
												MarkdownDescription: "DisableMergeSlashes disables Envoy's non-standard merge_slashes path transformation optionwhich strips duplicate slashes from request URL paths.Contour's default is false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"http_max_concurrent_streams": schema.Int64Attribute{
												Description:         "Defines the value for SETTINGS_MAX_CONCURRENT_STREAMS Envoy will advertise in theSETTINGS frame in HTTP/2 connections and the limit for concurrent streams allowedfor a peer on a single HTTP/2 connection. It is recommended to not set this lowerthan 100 but this field can be used to bound resource usage by HTTP/2 connectionsand mitigate attacks like CVE-2023-44487. The default value when this is not set isunlimited.",
												MarkdownDescription: "Defines the value for SETTINGS_MAX_CONCURRENT_STREAMS Envoy will advertise in theSETTINGS frame in HTTP/2 connections and the limit for concurrent streams allowedfor a peer on a single HTTP/2 connection. It is recommended to not set this lowerthan 100 but this field can be used to bound resource usage by HTTP/2 connectionsand mitigate attacks like CVE-2023-44487. The default value when this is not set isunlimited.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"max_connections_per_listener": schema.Int64Attribute{
												Description:         "Defines the limit on number of active connections to a listener. The limit is appliedper listener. The default value when this is not set is unlimited.",
												MarkdownDescription: "Defines the limit on number of active connections to a listener. The limit is appliedper listener. The default value when this is not set is unlimited.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"max_requests_per_connection": schema.Int64Attribute{
												Description:         "Defines the maximum requests for downstream connections. If not specified, there is no limit.see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-msg-config-core-v3-httpprotocoloptionsfor more information.",
												MarkdownDescription: "Defines the maximum requests for downstream connections. If not specified, there is no limit.see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-msg-config-core-v3-httpprotocoloptionsfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"max_requests_per_io_cycle": schema.Int64Attribute{
												Description:         "Defines the limit on number of HTTP requests that Envoy will process from a singleconnection in a single I/O cycle. Requests over this limit are processed in subsequentI/O cycles. Can be used as a mitigation for CVE-2023-44487 when abusive traffic isdetected. Configures the http.max_requests_per_io_cycle Envoy runtime setting. The defaultvalue when this is not set is no limit.",
												MarkdownDescription: "Defines the limit on number of HTTP requests that Envoy will process from a singleconnection in a single I/O cycle. Requests over this limit are processed in subsequentI/O cycles. Can be used as a mitigation for CVE-2023-44487 when abusive traffic isdetected. Configures the http.max_requests_per_io_cycle Envoy runtime setting. The defaultvalue when this is not set is no limit.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"per_connection_buffer_limit_bytes": schema.Int64Attribute{
												Description:         "Defines the soft limit on size of the listener’s new connection read and write buffers in bytes.If unspecified, an implementation defined default is applied (1MiB).see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener.proto#envoy-v3-api-field-config-listener-v3-listener-per-connection-buffer-limit-bytesfor more information.",
												MarkdownDescription: "Defines the soft limit on size of the listener’s new connection read and write buffers in bytes.If unspecified, an implementation defined default is applied (1MiB).see https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/listener/v3/listener.proto#envoy-v3-api-field-config-listener-v3-listener-per-connection-buffer-limit-bytesfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"server_header_transformation": schema.StringAttribute{
												Description:         "Defines the action to be applied to the Server header on the response path.When configured as overwrite, overwrites any Server header with 'envoy'.When configured as append_if_absent, if a Server header is present, pass it through, otherwise set it to 'envoy'.When configured as pass_through, pass through the value of the Server header, and do not append a header if none is present.Values: 'overwrite' (default), 'append_if_absent', 'pass_through'Other values will produce an error.Contour's default is overwrite.",
												MarkdownDescription: "Defines the action to be applied to the Server header on the response path.When configured as overwrite, overwrites any Server header with 'envoy'.When configured as append_if_absent, if a Server header is present, pass it through, otherwise set it to 'envoy'.When configured as pass_through, pass through the value of the Server header, and do not append a header if none is present.Values: 'overwrite' (default), 'append_if_absent', 'pass_through'Other values will produce an error.Contour's default is overwrite.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"socket_options": schema.SingleNestedAttribute{
												Description:         "SocketOptions defines configurable socket options for the listeners.Single set of options are applied to all listeners.",
												MarkdownDescription: "SocketOptions defines configurable socket options for the listeners.Single set of options are applied to all listeners.",
												Attributes: map[string]schema.Attribute{
													"tos": schema.Int64Attribute{
														Description:         "Defines the value for IPv4 TOS field (including 6 bit DSCP field) for IP packets originating from Envoy listeners.Single value is applied to all listeners.If listeners are bound to IPv6-only addresses, setting this option will cause an error.",
														MarkdownDescription: "Defines the value for IPv4 TOS field (including 6 bit DSCP field) for IP packets originating from Envoy listeners.Single value is applied to all listeners.If listeners are bound to IPv6-only addresses, setting this option will cause an error.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(255),
														},
													},

													"traffic_class": schema.Int64Attribute{
														Description:         "Defines the value for IPv6 Traffic Class field (including 6 bit DSCP field) for IP packets originating from the Envoy listeners.Single value is applied to all listeners.If listeners are bound to IPv4-only addresses, setting this option will cause an error.",
														MarkdownDescription: "Defines the value for IPv6 Traffic Class field (including 6 bit DSCP field) for IP packets originating from the Envoy listeners.Single value is applied to all listeners.If listeners are bound to IPv4-only addresses, setting this option will cause an error.",
														Required:            false,
														Optional:            true,
														Computed:            false,
														Validators: []validator.Int64{
															int64validator.AtLeast(0),
															int64validator.AtMost(255),
														},
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"tls": schema.SingleNestedAttribute{
												Description:         "TLS holds various configurable Envoy TLS listener values.",
												MarkdownDescription: "TLS holds various configurable Envoy TLS listener values.",
												Attributes: map[string]schema.Attribute{
													"cipher_suites": schema.ListAttribute{
														Description:         "CipherSuites defines the TLS ciphers to be supported by Envoy TLSlisteners when negotiating TLS 1.2. Ciphers are validated against theset that Envoy supports by default. This parameter should only be usedby advanced users. Note that these will be ignored when TLS 1.3 is inuse.This field is optional; when it is undefined, a Contour-managed ciphersuite listwill be used, which may be updated to keep it secure.Contour's default list is:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'Ciphers provided are validated against the following list:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES128-GCM-SHA256'  - 'ECDHE-RSA-AES128-GCM-SHA256'  - 'ECDHE-ECDSA-AES128-SHA'  - 'ECDHE-RSA-AES128-SHA'  - 'AES128-GCM-SHA256'  - 'AES128-SHA'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'  - 'ECDHE-ECDSA-AES256-SHA'  - 'ECDHE-RSA-AES256-SHA'  - 'AES256-GCM-SHA384'  - 'AES256-SHA'Contour recommends leaving this undefined unless you are sure you must.See: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#extensions-transport-sockets-tls-v3-tlsparametersNote: This list is a superset of what is valid for stock Envoy builds and those using BoringSSL FIPS.",
														MarkdownDescription: "CipherSuites defines the TLS ciphers to be supported by Envoy TLSlisteners when negotiating TLS 1.2. Ciphers are validated against theset that Envoy supports by default. This parameter should only be usedby advanced users. Note that these will be ignored when TLS 1.3 is inuse.This field is optional; when it is undefined, a Contour-managed ciphersuite listwill be used, which may be updated to keep it secure.Contour's default list is:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'Ciphers provided are validated against the following list:  - '[ECDHE-ECDSA-AES128-GCM-SHA256|ECDHE-ECDSA-CHACHA20-POLY1305]'  - '[ECDHE-RSA-AES128-GCM-SHA256|ECDHE-RSA-CHACHA20-POLY1305]'  - 'ECDHE-ECDSA-AES128-GCM-SHA256'  - 'ECDHE-RSA-AES128-GCM-SHA256'  - 'ECDHE-ECDSA-AES128-SHA'  - 'ECDHE-RSA-AES128-SHA'  - 'AES128-GCM-SHA256'  - 'AES128-SHA'  - 'ECDHE-ECDSA-AES256-GCM-SHA384'  - 'ECDHE-RSA-AES256-GCM-SHA384'  - 'ECDHE-ECDSA-AES256-SHA'  - 'ECDHE-RSA-AES256-SHA'  - 'AES256-GCM-SHA384'  - 'AES256-SHA'Contour recommends leaving this undefined unless you are sure you must.See: https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/transport_sockets/tls/v3/common.proto#extensions-transport-sockets-tls-v3-tlsparametersNote: This list is a superset of what is valid for stock Envoy builds and those using BoringSSL FIPS.",
														ElementType:         types.StringType,
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"maximum_protocol_version": schema.StringAttribute{
														Description:         "MaximumProtocolVersion is the maximum TLS version this vhost shouldnegotiate.Values: '1.2', '1.3'(default).Other values will produce an error.",
														MarkdownDescription: "MaximumProtocolVersion is the maximum TLS version this vhost shouldnegotiate.Values: '1.2', '1.3'(default).Other values will produce an error.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"minimum_protocol_version": schema.StringAttribute{
														Description:         "MinimumProtocolVersion is the minimum TLS version this vhost shouldnegotiate.Values: '1.2' (default), '1.3'.Other values will produce an error.",
														MarkdownDescription: "MinimumProtocolVersion is the minimum TLS version this vhost shouldnegotiate.Values: '1.2' (default), '1.3'.Other values will produce an error.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"use_proxy_protocol": schema.BoolAttribute{
												Description:         "Use PROXY protocol for all listeners.Contour's default is false.",
												MarkdownDescription: "Use PROXY protocol for all listeners.Contour's default is false.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"logging": schema.SingleNestedAttribute{
										Description:         "Logging defines how Envoy's logs can be configured.",
										MarkdownDescription: "Logging defines how Envoy's logs can be configured.",
										Attributes: map[string]schema.Attribute{
											"access_log_format": schema.StringAttribute{
												Description:         "AccessLogFormat sets the global access log format.Values: 'envoy' (default), 'json'.Other values will produce an error.",
												MarkdownDescription: "AccessLogFormat sets the global access log format.Values: 'envoy' (default), 'json'.Other values will produce an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"access_log_format_string": schema.StringAttribute{
												Description:         "AccessLogFormatString sets the access log format when format is set to 'envoy'.When empty, Envoy's default format is used.",
												MarkdownDescription: "AccessLogFormatString sets the access log format when format is set to 'envoy'.When empty, Envoy's default format is used.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"access_log_json_fields": schema.ListAttribute{
												Description:         "AccessLogJSONFields sets the fields that JSON logging willoutput when AccessLogFormat is json.",
												MarkdownDescription: "AccessLogJSONFields sets the fields that JSON logging willoutput when AccessLogFormat is json.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"access_log_level": schema.StringAttribute{
												Description:         "AccessLogLevel sets the verbosity level of the access log.Values: 'info' (default, all requests are logged), 'error' (all non-success requests, i.e. 300+ response code, are logged), 'critical' (all 5xx requests are logged) and 'disabled'.Other values will produce an error.",
												MarkdownDescription: "AccessLogLevel sets the verbosity level of the access log.Values: 'info' (default, all requests are logged), 'error' (all non-success requests, i.e. 300+ response code, are logged), 'critical' (all 5xx requests are logged) and 'disabled'.Other values will produce an error.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"metrics": schema.SingleNestedAttribute{
										Description:         "Metrics defines the endpoint Envoy uses to serve metrics.Contour's default is { address: '0.0.0.0', port: 8002 }.",
										MarkdownDescription: "Metrics defines the endpoint Envoy uses to serve metrics.Contour's default is { address: '0.0.0.0', port: 8002 }.",
										Attributes: map[string]schema.Attribute{
											"address": schema.StringAttribute{
												Description:         "Defines the metrics address interface.",
												MarkdownDescription: "Defines the metrics address interface.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
													stringvalidator.LengthAtMost(253),
												},
											},

											"port": schema.Int64Attribute{
												Description:         "Defines the metrics port.",
												MarkdownDescription: "Defines the metrics port.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"tls": schema.SingleNestedAttribute{
												Description:         "TLS holds TLS file config details.Metrics and health endpoints cannot have same port number when metrics is served over HTTPS.",
												MarkdownDescription: "TLS holds TLS file config details.Metrics and health endpoints cannot have same port number when metrics is served over HTTPS.",
												Attributes: map[string]schema.Attribute{
													"ca_file": schema.StringAttribute{
														Description:         "CA filename.",
														MarkdownDescription: "CA filename.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"cert_file": schema.StringAttribute{
														Description:         "Client certificate filename.",
														MarkdownDescription: "Client certificate filename.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"key_file": schema.StringAttribute{
														Description:         "Client key filename.",
														MarkdownDescription: "Client key filename.",
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

									"network": schema.SingleNestedAttribute{
										Description:         "Network holds various configurable Envoy network values.",
										MarkdownDescription: "Network holds various configurable Envoy network values.",
										Attributes: map[string]schema.Attribute{
											"admin_port": schema.Int64Attribute{
												Description:         "Configure the port used to access the Envoy Admin interface.If configured to port '0' then the admin interface is disabled.Contour's default is 9001.",
												MarkdownDescription: "Configure the port used to access the Envoy Admin interface.If configured to port '0' then the admin interface is disabled.Contour's default is 9001.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"num_trusted_hops": schema.Int64Attribute{
												Description:         "XffNumTrustedHops defines the number of additional ingress proxy hops from theright side of the x-forwarded-for HTTP header to trust when determining the originclient’s IP address.See https://www.envoyproxy.io/docs/envoy/v1.17.0/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto?highlight=xff_num_trusted_hopsfor more information.Contour's default is 0.",
												MarkdownDescription: "XffNumTrustedHops defines the number of additional ingress proxy hops from theright side of the x-forwarded-for HTTP header to trust when determining the originclient’s IP address.See https://www.envoyproxy.io/docs/envoy/v1.17.0/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto?highlight=xff_num_trusted_hopsfor more information.Contour's default is 0.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"service": schema.SingleNestedAttribute{
										Description:         "Service holds Envoy service parameters for setting Ingress status.Contour's default is { namespace: 'projectcontour', name: 'envoy' }.",
										MarkdownDescription: "Service holds Envoy service parameters for setting Ingress status.Contour's default is { namespace: 'projectcontour', name: 'envoy' }.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
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

									"timeouts": schema.SingleNestedAttribute{
										Description:         "Timeouts holds various configurable timeouts that canbe set in the config file.",
										MarkdownDescription: "Timeouts holds various configurable timeouts that canbe set in the config file.",
										Attributes: map[string]schema.Attribute{
											"connect_timeout": schema.StringAttribute{
												Description:         "ConnectTimeout defines how long the proxy should wait when establishing connection to upstream service.If not set, a default value of 2 seconds will be used.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto#envoy-v3-api-field-config-cluster-v3-cluster-connect-timeoutfor more information.",
												MarkdownDescription: "ConnectTimeout defines how long the proxy should wait when establishing connection to upstream service.If not set, a default value of 2 seconds will be used.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/cluster/v3/cluster.proto#envoy-v3-api-field-config-cluster-v3-cluster-connect-timeoutfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection_idle_timeout": schema.StringAttribute{
												Description:         "ConnectionIdleTimeout defines how long the proxy should wait while there areno active requests (for HTTP/1.1) or streams (for HTTP/2) before terminatingan HTTP connection. Set to 'infinity' to disable the timeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-idle-timeoutfor more information.",
												MarkdownDescription: "ConnectionIdleTimeout defines how long the proxy should wait while there areno active requests (for HTTP/1.1) or streams (for HTTP/2) before terminatingan HTTP connection. Set to 'infinity' to disable the timeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-idle-timeoutfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"connection_shutdown_grace_period": schema.StringAttribute{
												Description:         "ConnectionShutdownGracePeriod defines how long the proxy will wait between sending aninitial GOAWAY frame and a second, final GOAWAY frame when terminating an HTTP/2 connection.During this grace period, the proxy will continue to respond to new streams. After the finalGOAWAY frame has been sent, the proxy will refuse new streams.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-drain-timeoutfor more information.",
												MarkdownDescription: "ConnectionShutdownGracePeriod defines how long the proxy will wait between sending aninitial GOAWAY frame and a second, final GOAWAY frame when terminating an HTTP/2 connection.During this grace period, the proxy will continue to respond to new streams. After the finalGOAWAY frame has been sent, the proxy will refuse new streams.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-drain-timeoutfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"delayed_close_timeout": schema.StringAttribute{
												Description:         "DelayedCloseTimeout defines how long envoy will wait, once connectionclose processing has been initiated, for the downstream peer to closethe connection before Envoy closes the socket associated with the connection.Setting this timeout to 'infinity' will disable it, equivalent to setting it to '0'in Envoy. Leaving it unset will result in the Envoy default value being used.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-delayed-close-timeoutfor more information.",
												MarkdownDescription: "DelayedCloseTimeout defines how long envoy will wait, once connectionclose processing has been initiated, for the downstream peer to closethe connection before Envoy closes the socket associated with the connection.Setting this timeout to 'infinity' will disable it, equivalent to setting it to '0'in Envoy. Leaving it unset will result in the Envoy default value being used.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-delayed-close-timeoutfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_connection_duration": schema.StringAttribute{
												Description:         "MaxConnectionDuration defines the maximum period of time after an HTTP connectionhas been established from the client to the proxy before it is closed by the proxy,regardless of whether there has been activity or not. Omit or set to 'infinity' forno max duration.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-max-connection-durationfor more information.",
												MarkdownDescription: "MaxConnectionDuration defines the maximum period of time after an HTTP connectionhas been established from the client to the proxy before it is closed by the proxy,regardless of whether there has been activity or not. Omit or set to 'infinity' forno max duration.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/config/core/v3/protocol.proto#envoy-v3-api-field-config-core-v3-httpprotocoloptions-max-connection-durationfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"request_timeout": schema.StringAttribute{
												Description:         "RequestTimeout sets the client request timeout globally for Contour. Note thatthis is a timeout for the entire request, not an idle timeout. Omit or set to'infinity' to disable the timeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-request-timeoutfor more information.",
												MarkdownDescription: "RequestTimeout sets the client request timeout globally for Contour. Note thatthis is a timeout for the entire request, not an idle timeout. Omit or set to'infinity' to disable the timeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-request-timeoutfor more information.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"stream_idle_timeout": schema.StringAttribute{
												Description:         "StreamIdleTimeout defines how long the proxy should wait while there is norequest activity (for HTTP/1.1) or stream activity (for HTTP/2) beforeterminating the HTTP request or stream. Set to 'infinity' to disable thetimeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-stream-idle-timeoutfor more information.",
												MarkdownDescription: "StreamIdleTimeout defines how long the proxy should wait while there is norequest activity (for HTTP/1.1) or stream activity (for HTTP/2) beforeterminating the HTTP request or stream. Set to 'infinity' to disable thetimeout entirely.See https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/filters/network/http_connection_manager/v3/http_connection_manager.proto#envoy-v3-api-field-extensions-filters-network-http-connection-manager-v3-httpconnectionmanager-stream-idle-timeoutfor more information.",
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

							"feature_flags": schema.ListAttribute{
								Description:         "FeatureFlags defines toggle to enable new contour features.Available toggles are:useEndpointSlices - Configures contour to fetch endpoint datafrom k8s endpoint slices. defaults to true,If false then reads endpoint data from the k8s endpoints.",
								MarkdownDescription: "FeatureFlags defines toggle to enable new contour features.Available toggles are:useEndpointSlices - Configures contour to fetch endpoint datafrom k8s endpoint slices. defaults to true,If false then reads endpoint data from the k8s endpoints.",
								ElementType:         types.StringType,
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"gateway": schema.SingleNestedAttribute{
								Description:         "Gateway contains parameters for the gateway-api Gateway that Contouris configured to serve traffic.",
								MarkdownDescription: "Gateway contains parameters for the gateway-api Gateway that Contouris configured to serve traffic.",
								Attributes: map[string]schema.Attribute{
									"gateway_ref": schema.SingleNestedAttribute{
										Description:         "GatewayRef defines the specific Gateway that this Contourinstance corresponds to.",
										MarkdownDescription: "GatewayRef defines the specific Gateway that this Contourinstance corresponds to.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
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

							"global_ext_auth": schema.SingleNestedAttribute{
								Description:         "GlobalExternalAuthorization allows envoys external authorization filterto be enabled for all virtual hosts.",
								MarkdownDescription: "GlobalExternalAuthorization allows envoys external authorization filterto be enabled for all virtual hosts.",
								Attributes: map[string]schema.Attribute{
									"auth_policy": schema.SingleNestedAttribute{
										Description:         "AuthPolicy sets a default authorization policy for client requests.This policy will be used unless overridden by individual routes.",
										MarkdownDescription: "AuthPolicy sets a default authorization policy for client requests.This policy will be used unless overridden by individual routes.",
										Attributes: map[string]schema.Attribute{
											"context": schema.MapAttribute{
												Description:         "Context is a set of key/value pairs that are sent to theauthentication server in the check request. If a contextis provided at an enclosing scope, the entries are mergedsuch that the inner scope overrides matching keys from theouter scope.",
												MarkdownDescription: "Context is a set of key/value pairs that are sent to theauthentication server in the check request. If a contextis provided at an enclosing scope, the entries are mergedsuch that the inner scope overrides matching keys from theouter scope.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"disabled": schema.BoolAttribute{
												Description:         "When true, this field disables client request authenticationfor the scope of the policy.",
												MarkdownDescription: "When true, this field disables client request authenticationfor the scope of the policy.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"extension_ref": schema.SingleNestedAttribute{
										Description:         "ExtensionServiceRef specifies the extension resource that will authorize client requests.",
										MarkdownDescription: "ExtensionServiceRef specifies the extension resource that will authorize client requests.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "API version of the referent.If this field is not specified, the default 'projectcontour.io/v1alpha1' will be used",
												MarkdownDescription: "API version of the referent.If this field is not specified, the default 'projectcontour.io/v1alpha1' will be used",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.String{
													stringvalidator.LengthAtLeast(1),
												},
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the referent.If this field is not specifies, the namespace of the resource that targets the referent will be used.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												MarkdownDescription: "Namespace of the referent.If this field is not specifies, the namespace of the resource that targets the referent will be used.More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												Required:            false,
												Optional:            true,
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

									"fail_open": schema.BoolAttribute{
										Description:         "If FailOpen is true, the client request is forwarded to the upstream serviceeven if the authorization server fails to respond. This field should not beset in most cases. It is intended for use only while migrating applicationsfrom internal authorization to Contour external authorization.",
										MarkdownDescription: "If FailOpen is true, the client request is forwarded to the upstream serviceeven if the authorization server fails to respond. This field should not beset in most cases. It is intended for use only while migrating applicationsfrom internal authorization to Contour external authorization.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"response_timeout": schema.StringAttribute{
										Description:         "ResponseTimeout configures maximum time to wait for a check response from the authorization server.Timeout durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration).Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.The string 'infinity' is also a valid input and specifies no timeout.",
										MarkdownDescription: "ResponseTimeout configures maximum time to wait for a check response from the authorization server.Timeout durations are expressed in the Go [Duration format](https://godoc.org/time#ParseDuration).Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'.The string 'infinity' is also a valid input and specifies no timeout.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(((\d*(\.\d*)?h)|(\d*(\.\d*)?m)|(\d*(\.\d*)?s)|(\d*(\.\d*)?ms)|(\d*(\.\d*)?us)|(\d*(\.\d*)?µs)|(\d*(\.\d*)?ns))+|infinity|infinite)$`), ""),
										},
									},

									"with_request_body": schema.SingleNestedAttribute{
										Description:         "WithRequestBody specifies configuration for sending the client request's body to authorization server.",
										MarkdownDescription: "WithRequestBody specifies configuration for sending the client request's body to authorization server.",
										Attributes: map[string]schema.Attribute{
											"allow_partial_message": schema.BoolAttribute{
												Description:         "If AllowPartialMessage is true, then Envoy will buffer the body until MaxRequestBytes are reached.",
												MarkdownDescription: "If AllowPartialMessage is true, then Envoy will buffer the body until MaxRequestBytes are reached.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"max_request_bytes": schema.Int64Attribute{
												Description:         "MaxRequestBytes sets the maximum size of message body ExtAuthz filter will hold in-memory.",
												MarkdownDescription: "MaxRequestBytes sets the maximum size of message body ExtAuthz filter will hold in-memory.",
												Required:            false,
												Optional:            true,
												Computed:            false,
												Validators: []validator.Int64{
													int64validator.AtLeast(1),
												},
											},

											"pack_as_bytes": schema.BoolAttribute{
												Description:         "If PackAsBytes is true, the body sent to Authorization Server is in raw bytes.",
												MarkdownDescription: "If PackAsBytes is true, the body sent to Authorization Server is in raw bytes.",
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

							"health": schema.SingleNestedAttribute{
								Description:         "Health defines the endpoints Contour uses to serve health checks.Contour's default is { address: '0.0.0.0', port: 8000 }.",
								MarkdownDescription: "Health defines the endpoints Contour uses to serve health checks.Contour's default is { address: '0.0.0.0', port: 8000 }.",
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description:         "Defines the health address interface.",
										MarkdownDescription: "Defines the health address interface.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "Defines the health port.",
										MarkdownDescription: "Defines the health port.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"httpproxy": schema.SingleNestedAttribute{
								Description:         "HTTPProxy defines parameters on HTTPProxy.",
								MarkdownDescription: "HTTPProxy defines parameters on HTTPProxy.",
								Attributes: map[string]schema.Attribute{
									"disable_permit_insecure": schema.BoolAttribute{
										Description:         "DisablePermitInsecure disables the use of thepermitInsecure field in HTTPProxy.Contour's default is false.",
										MarkdownDescription: "DisablePermitInsecure disables the use of thepermitInsecure field in HTTPProxy.Contour's default is false.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"fallback_certificate": schema.SingleNestedAttribute{
										Description:         "FallbackCertificate defines the namespace/name of the Kubernetes secret touse as fallback when a non-SNI request is received.",
										MarkdownDescription: "FallbackCertificate defines the namespace/name of the Kubernetes secret touse as fallback when a non-SNI request is received.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
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

									"root_namespaces": schema.ListAttribute{
										Description:         "Restrict Contour to searching these namespaces for root ingress routes.",
										MarkdownDescription: "Restrict Contour to searching these namespaces for root ingress routes.",
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

							"ingress": schema.SingleNestedAttribute{
								Description:         "Ingress contains parameters for ingress options.",
								MarkdownDescription: "Ingress contains parameters for ingress options.",
								Attributes: map[string]schema.Attribute{
									"class_names": schema.ListAttribute{
										Description:         "Ingress Class Names Contour should use.",
										MarkdownDescription: "Ingress Class Names Contour should use.",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"status_address": schema.StringAttribute{
										Description:         "Address to set in Ingress object status.",
										MarkdownDescription: "Address to set in Ingress object status.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"metrics": schema.SingleNestedAttribute{
								Description:         "Metrics defines the endpoint Contour uses to serve metrics.Contour's default is { address: '0.0.0.0', port: 8000 }.",
								MarkdownDescription: "Metrics defines the endpoint Contour uses to serve metrics.Contour's default is { address: '0.0.0.0', port: 8000 }.",
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description:         "Defines the metrics address interface.",
										MarkdownDescription: "Defines the metrics address interface.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
											stringvalidator.LengthAtMost(253),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "Defines the metrics port.",
										MarkdownDescription: "Defines the metrics port.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS holds TLS file config details.Metrics and health endpoints cannot have same port number when metrics is served over HTTPS.",
										MarkdownDescription: "TLS holds TLS file config details.Metrics and health endpoints cannot have same port number when metrics is served over HTTPS.",
										Attributes: map[string]schema.Attribute{
											"ca_file": schema.StringAttribute{
												Description:         "CA filename.",
												MarkdownDescription: "CA filename.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_file": schema.StringAttribute{
												Description:         "Client certificate filename.",
												MarkdownDescription: "Client certificate filename.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key_file": schema.StringAttribute{
												Description:         "Client key filename.",
												MarkdownDescription: "Client key filename.",
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

							"policy": schema.SingleNestedAttribute{
								Description:         "Policy specifies default policy applied if not overridden by the user",
								MarkdownDescription: "Policy specifies default policy applied if not overridden by the user",
								Attributes: map[string]schema.Attribute{
									"apply_to_ingress": schema.BoolAttribute{
										Description:         "ApplyToIngress determines if the Policies will apply to ingress objectsContour's default is false.",
										MarkdownDescription: "ApplyToIngress determines if the Policies will apply to ingress objectsContour's default is false.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"request_headers": schema.SingleNestedAttribute{
										Description:         "RequestHeadersPolicy defines the request headers set/removed on all routes",
										MarkdownDescription: "RequestHeadersPolicy defines the request headers set/removed on all routes",
										Attributes: map[string]schema.Attribute{
											"remove": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"set": schema.MapAttribute{
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

									"response_headers": schema.SingleNestedAttribute{
										Description:         "ResponseHeadersPolicy defines the response headers set/removed on all routes",
										MarkdownDescription: "ResponseHeadersPolicy defines the response headers set/removed on all routes",
										Attributes: map[string]schema.Attribute{
											"remove": schema.ListAttribute{
												Description:         "",
												MarkdownDescription: "",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"set": schema.MapAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"rate_limit_service": schema.SingleNestedAttribute{
								Description:         "RateLimitService optionally holds properties of the Rate Limit Serviceto be used for global rate limiting.",
								MarkdownDescription: "RateLimitService optionally holds properties of the Rate Limit Serviceto be used for global rate limiting.",
								Attributes: map[string]schema.Attribute{
									"default_global_rate_limit_policy": schema.SingleNestedAttribute{
										Description:         "DefaultGlobalRateLimitPolicy allows setting a default global rate limit policy for every HTTPProxy.HTTPProxy can overwrite this configuration.",
										MarkdownDescription: "DefaultGlobalRateLimitPolicy allows setting a default global rate limit policy for every HTTPProxy.HTTPProxy can overwrite this configuration.",
										Attributes: map[string]schema.Attribute{
											"descriptors": schema.ListNestedAttribute{
												Description:         "Descriptors defines the list of descriptors that willbe generated and sent to the rate limit service. Eachdescriptor contains 1+ key-value pair entries.",
												MarkdownDescription: "Descriptors defines the list of descriptors that willbe generated and sent to the rate limit service. Eachdescriptor contains 1+ key-value pair entries.",
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{
														"entries": schema.ListNestedAttribute{
															Description:         "Entries is the list of key-value pair generators.",
															MarkdownDescription: "Entries is the list of key-value pair generators.",
															NestedObject: schema.NestedAttributeObject{
																Attributes: map[string]schema.Attribute{
																	"generic_key": schema.SingleNestedAttribute{
																		Description:         "GenericKey defines a descriptor entry with a static key and value.",
																		MarkdownDescription: "GenericKey defines a descriptor entry with a static key and value.",
																		Attributes: map[string]schema.Attribute{
																			"key": schema.StringAttribute{
																				Description:         "Key defines the key of the descriptor entry. If not set, thekey is set to 'generic_key'.",
																				MarkdownDescription: "Key defines the key of the descriptor entry. If not set, thekey is set to 'generic_key'.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"value": schema.StringAttribute{
																				Description:         "Value defines the value of the descriptor entry.",
																				MarkdownDescription: "Value defines the value of the descriptor entry.",
																				Required:            false,
																				Optional:            true,
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

																	"remote_address": schema.MapAttribute{
																		Description:         "RemoteAddress defines a descriptor entry with a key of 'remote_address'and a value equal to the client's IP address (from x-forwarded-for).",
																		MarkdownDescription: "RemoteAddress defines a descriptor entry with a key of 'remote_address'and a value equal to the client's IP address (from x-forwarded-for).",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"request_header": schema.SingleNestedAttribute{
																		Description:         "RequestHeader defines a descriptor entry that's populated only ifa given header is present on the request. The descriptor key is static,and the descriptor value is equal to the value of the header.",
																		MarkdownDescription: "RequestHeader defines a descriptor entry that's populated only ifa given header is present on the request. The descriptor key is static,and the descriptor value is equal to the value of the header.",
																		Attributes: map[string]schema.Attribute{
																			"descriptor_key": schema.StringAttribute{
																				Description:         "DescriptorKey defines the key to use on the descriptor entry.",
																				MarkdownDescription: "DescriptorKey defines the key to use on the descriptor entry.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},

																			"header_name": schema.StringAttribute{
																				Description:         "HeaderName defines the name of the header to look for on the request.",
																				MarkdownDescription: "HeaderName defines the name of the header to look for on the request.",
																				Required:            false,
																				Optional:            true,
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

																	"request_header_value_match": schema.SingleNestedAttribute{
																		Description:         "RequestHeaderValueMatch defines a descriptor entry that's populatedif the request's headers match a set of 1+ match criteria. Thedescriptor key is 'header_match', and the descriptor value is static.",
																		MarkdownDescription: "RequestHeaderValueMatch defines a descriptor entry that's populatedif the request's headers match a set of 1+ match criteria. Thedescriptor key is 'header_match', and the descriptor value is static.",
																		Attributes: map[string]schema.Attribute{
																			"expect_match": schema.BoolAttribute{
																				Description:         "ExpectMatch defines whether the request must positively match the matchcriteria in order to generate a descriptor entry (i.e. true), or notmatch the match criteria in order to generate a descriptor entry (i.e. false).The default is true.",
																				MarkdownDescription: "ExpectMatch defines whether the request must positively match the matchcriteria in order to generate a descriptor entry (i.e. true), or notmatch the match criteria in order to generate a descriptor entry (i.e. false).The default is true.",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"headers": schema.ListNestedAttribute{
																				Description:         "Headers is a list of 1+ match criteria to apply against the requestto determine whether to populate the descriptor entry or not.",
																				MarkdownDescription: "Headers is a list of 1+ match criteria to apply against the requestto determine whether to populate the descriptor entry or not.",
																				NestedObject: schema.NestedAttributeObject{
																					Attributes: map[string]schema.Attribute{
																						"contains": schema.StringAttribute{
																							Description:         "Contains specifies a substring that must be present inthe header value.",
																							MarkdownDescription: "Contains specifies a substring that must be present inthe header value.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"exact": schema.StringAttribute{
																							Description:         "Exact specifies a string that the header value must be equal to.",
																							MarkdownDescription: "Exact specifies a string that the header value must be equal to.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"ignore_case": schema.BoolAttribute{
																							Description:         "IgnoreCase specifies that string matching should be case insensitive.Note that this has no effect on the Regex parameter.",
																							MarkdownDescription: "IgnoreCase specifies that string matching should be case insensitive.Note that this has no effect on the Regex parameter.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"name": schema.StringAttribute{
																							Description:         "Name is the name of the header to match against. Name is required.Header names are case insensitive.",
																							MarkdownDescription: "Name is the name of the header to match against. Name is required.Header names are case insensitive.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"notcontains": schema.StringAttribute{
																							Description:         "NotContains specifies a substring that must not be presentin the header value.",
																							MarkdownDescription: "NotContains specifies a substring that must not be presentin the header value.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"notexact": schema.StringAttribute{
																							Description:         "NoExact specifies a string that the header value must not beequal to. The condition is true if the header has any other value.",
																							MarkdownDescription: "NoExact specifies a string that the header value must not beequal to. The condition is true if the header has any other value.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"notpresent": schema.BoolAttribute{
																							Description:         "NotPresent specifies that condition is true when the named headeris not present. Note that setting NotPresent to false does notmake the condition true if the named header is present.",
																							MarkdownDescription: "NotPresent specifies that condition is true when the named headeris not present. Note that setting NotPresent to false does notmake the condition true if the named header is present.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"present": schema.BoolAttribute{
																							Description:         "Present specifies that condition is true when the named headeris present, regardless of its value. Note that setting Presentto false does not make the condition true if the named headeris absent.",
																							MarkdownDescription: "Present specifies that condition is true when the named headeris present, regardless of its value. Note that setting Presentto false does not make the condition true if the named headeris absent.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"regex": schema.StringAttribute{
																							Description:         "Regex specifies a regular expression pattern that must match the headervalue.",
																							MarkdownDescription: "Regex specifies a regular expression pattern that must match the headervalue.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"treat_missing_as_empty": schema.BoolAttribute{
																							Description:         "TreatMissingAsEmpty specifies if the header match rule specified headerdoes not exist, this header value will be treated as empty. Defaults to false.Unlike the underlying Envoy implementation this is **only** supported fornegative matches (e.g. NotContains, NotExact).",
																							MarkdownDescription: "TreatMissingAsEmpty specifies if the header match rule specified headerdoes not exist, this header value will be treated as empty. Defaults to false.Unlike the underlying Envoy implementation this is **only** supported fornegative matches (e.g. NotContains, NotExact).",
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

																			"value": schema.StringAttribute{
																				Description:         "Value defines the value of the descriptor entry.",
																				MarkdownDescription: "Value defines the value of the descriptor entry.",
																				Required:            false,
																				Optional:            true,
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

											"disabled": schema.BoolAttribute{
												Description:         "Disabled configures the HTTPProxy to not usethe default global rate limit policy defined by the Contour configuration.",
												MarkdownDescription: "Disabled configures the HTTPProxy to not usethe default global rate limit policy defined by the Contour configuration.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"domain": schema.StringAttribute{
										Description:         "Domain is passed to the Rate Limit Service.",
										MarkdownDescription: "Domain is passed to the Rate Limit Service.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable_resource_exhausted_code": schema.BoolAttribute{
										Description:         "EnableResourceExhaustedCode enables translating error code 429 togrpc code RESOURCE_EXHAUSTED. When disabled it's translated to UNAVAILABLE",
										MarkdownDescription: "EnableResourceExhaustedCode enables translating error code 429 togrpc code RESOURCE_EXHAUSTED. When disabled it's translated to UNAVAILABLE",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"enable_x_rate_limit_headers": schema.BoolAttribute{
										Description:         "EnableXRateLimitHeaders defines whether to include the X-RateLimitheaders X-RateLimit-Limit, X-RateLimit-Remaining, and X-RateLimit-Reset(as defined by the IETF Internet-Draft linked below), on responsesto clients when the Rate Limit Service is consulted for a request.ref. https://tools.ietf.org/id/draft-polli-ratelimit-headers-03.html",
										MarkdownDescription: "EnableXRateLimitHeaders defines whether to include the X-RateLimitheaders X-RateLimit-Limit, X-RateLimit-Remaining, and X-RateLimit-Reset(as defined by the IETF Internet-Draft linked below), on responsesto clients when the Rate Limit Service is consulted for a request.ref. https://tools.ietf.org/id/draft-polli-ratelimit-headers-03.html",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"extension_service": schema.SingleNestedAttribute{
										Description:         "ExtensionService identifies the extension service defining the RLS.",
										MarkdownDescription: "ExtensionService identifies the extension service defining the RLS.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"fail_open": schema.BoolAttribute{
										Description:         "FailOpen defines whether to allow requests to proceed when theRate Limit Service fails to respond with a valid rate limitdecision within the timeout defined on the extension service.",
										MarkdownDescription: "FailOpen defines whether to allow requests to proceed when theRate Limit Service fails to respond with a valid rate limitdecision within the timeout defined on the extension service.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"tracing": schema.SingleNestedAttribute{
								Description:         "Tracing defines properties for exporting trace data to OpenTelemetry.",
								MarkdownDescription: "Tracing defines properties for exporting trace data to OpenTelemetry.",
								Attributes: map[string]schema.Attribute{
									"custom_tags": schema.ListNestedAttribute{
										Description:         "CustomTags defines a list of custom tags with unique tag name.",
										MarkdownDescription: "CustomTags defines a list of custom tags with unique tag name.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"literal": schema.StringAttribute{
													Description:         "Literal is a static custom tag value.Precisely one of Literal, RequestHeaderName must be set.",
													MarkdownDescription: "Literal is a static custom tag value.Precisely one of Literal, RequestHeaderName must be set.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"request_header_name": schema.StringAttribute{
													Description:         "RequestHeaderName indicates which request headerthe label value is obtained from.Precisely one of Literal, RequestHeaderName must be set.",
													MarkdownDescription: "RequestHeaderName indicates which request headerthe label value is obtained from.Precisely one of Literal, RequestHeaderName must be set.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"tag_name": schema.StringAttribute{
													Description:         "TagName is the unique name of the custom tag.",
													MarkdownDescription: "TagName is the unique name of the custom tag.",
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

									"extension_service": schema.SingleNestedAttribute{
										Description:         "ExtensionService identifies the extension service defining the otel-collector.",
										MarkdownDescription: "ExtensionService identifies the extension service defining the otel-collector.",
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "",
												MarkdownDescription: "",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: true,
										Optional: false,
										Computed: false,
									},

									"include_pod_detail": schema.BoolAttribute{
										Description:         "IncludePodDetail defines a flag.If it is true, contour will add the pod name and namespace to the span of the trace.the default is true.Note: The Envoy pods MUST have the HOSTNAME and CONTOUR_NAMESPACE environment variables set for this to work properly.",
										MarkdownDescription: "IncludePodDetail defines a flag.If it is true, contour will add the pod name and namespace to the span of the trace.the default is true.Note: The Envoy pods MUST have the HOSTNAME and CONTOUR_NAMESPACE environment variables set for this to work properly.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"max_path_tag_length": schema.Int64Attribute{
										Description:         "MaxPathTagLength defines maximum length of the request pathto extract and include in the HttpUrl tag.contour's default is 256.",
										MarkdownDescription: "MaxPathTagLength defines maximum length of the request pathto extract and include in the HttpUrl tag.contour's default is 256.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"overall_sampling": schema.StringAttribute{
										Description:         "OverallSampling defines the sampling rate of trace data.contour's default is 100.",
										MarkdownDescription: "OverallSampling defines the sampling rate of trace data.contour's default is 100.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"service_name": schema.StringAttribute{
										Description:         "ServiceName defines the name for the service.contour's default is contour.",
										MarkdownDescription: "ServiceName defines the name for the service.contour's default is contour.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"xds_server": schema.SingleNestedAttribute{
								Description:         "XDSServer contains parameters for the xDS server.",
								MarkdownDescription: "XDSServer contains parameters for the xDS server.",
								Attributes: map[string]schema.Attribute{
									"address": schema.StringAttribute{
										Description:         "Defines the xDS gRPC API address which Contour will serve.Contour's default is '0.0.0.0'.",
										MarkdownDescription: "Defines the xDS gRPC API address which Contour will serve.Contour's default is '0.0.0.0'.",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.LengthAtLeast(1),
										},
									},

									"port": schema.Int64Attribute{
										Description:         "Defines the xDS gRPC API port which Contour will serve.Contour's default is 8001.",
										MarkdownDescription: "Defines the xDS gRPC API port which Contour will serve.Contour's default is 8001.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tls": schema.SingleNestedAttribute{
										Description:         "TLS holds TLS file config details.Contour's default is { caFile: '/certs/ca.crt', certFile: '/certs/tls.cert', keyFile: '/certs/tls.key', insecure: false }.",
										MarkdownDescription: "TLS holds TLS file config details.Contour's default is { caFile: '/certs/ca.crt', certFile: '/certs/tls.cert', keyFile: '/certs/tls.key', insecure: false }.",
										Attributes: map[string]schema.Attribute{
											"ca_file": schema.StringAttribute{
												Description:         "CA filename.",
												MarkdownDescription: "CA filename.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"cert_file": schema.StringAttribute{
												Description:         "Client certificate filename.",
												MarkdownDescription: "Client certificate filename.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"insecure": schema.BoolAttribute{
												Description:         "Allow serving the xDS gRPC API without TLS.",
												MarkdownDescription: "Allow serving the xDS gRPC API without TLS.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"key_file": schema.StringAttribute{
												Description:         "Client key filename.",
												MarkdownDescription: "Client key filename.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"type": schema.StringAttribute{
										Description:         "Defines the XDSServer to use for 'contour serve'.Values: 'envoy' (default), 'contour (deprecated)'.Other values will produce an error.",
										MarkdownDescription: "Defines the XDSServer to use for 'contour serve'.Values: 'envoy' (default), 'contour (deprecated)'.Other values will produce an error.",
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
				Required: false,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func (r *ProjectcontourIoContourDeploymentV1Alpha1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_projectcontour_io_contour_deployment_v1alpha1_manifest")

	var model ProjectcontourIoContourDeploymentV1Alpha1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("projectcontour.io/v1alpha1")
	model.Kind = pointer.String("ContourDeployment")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.Append(utilities.MarshalYamlError(err))
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
