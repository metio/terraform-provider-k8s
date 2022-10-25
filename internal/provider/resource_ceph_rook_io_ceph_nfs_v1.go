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

type CephRookIoCephNFSV1Resource struct{}

var (
	_ resource.Resource = (*CephRookIoCephNFSV1Resource)(nil)
)

type CephRookIoCephNFSV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type CephRookIoCephNFSV1GoModel struct {
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
		Rados *struct {
			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			Pool *string `tfsdk:"pool" yaml:"pool,omitempty"`
		} `tfsdk:"rados" yaml:"rados,omitempty"`

		Security *struct {
			Kerberos *struct {
				ConfigFiles *struct {
					VolumeSource *struct {
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
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

									Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

									Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"metadata" yaml:"metadata,omitempty"`

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
					} `tfsdk:"volume_source" yaml:"volumeSource,omitempty"`
				} `tfsdk:"config_files" yaml:"configFiles,omitempty"`

				KeytabFile *struct {
					VolumeSource *struct {
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
								Metadata *struct {
									Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

									Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

									Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

									Name *string `tfsdk:"name" yaml:"name,omitempty"`

									Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
								} `tfsdk:"metadata" yaml:"metadata,omitempty"`

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
					} `tfsdk:"volume_source" yaml:"volumeSource,omitempty"`
				} `tfsdk:"keytab_file" yaml:"keytabFile,omitempty"`

				PrincipalName *string `tfsdk:"principal_name" yaml:"principalName,omitempty"`
			} `tfsdk:"kerberos" yaml:"kerberos,omitempty"`

			Sssd *struct {
				Sidecar *struct {
					AdditionalFiles *[]struct {
						SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

						VolumeSource *struct {
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
									Metadata *struct {
										Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

										Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

										Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
									} `tfsdk:"metadata" yaml:"metadata,omitempty"`

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
						} `tfsdk:"volume_source" yaml:"volumeSource,omitempty"`
					} `tfsdk:"additional_files" yaml:"additionalFiles,omitempty"`

					DebugLevel *int64 `tfsdk:"debug_level" yaml:"debugLevel,omitempty"`

					Image *string `tfsdk:"image" yaml:"image,omitempty"`

					Resources *struct {
						Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

						Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
					} `tfsdk:"resources" yaml:"resources,omitempty"`

					SssdConfigFile *struct {
						VolumeSource *struct {
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
									Metadata *struct {
										Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

										Finalizers *[]string `tfsdk:"finalizers" yaml:"finalizers,omitempty"`

										Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
									} `tfsdk:"metadata" yaml:"metadata,omitempty"`

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
						} `tfsdk:"volume_source" yaml:"volumeSource,omitempty"`
					} `tfsdk:"sssd_config_file" yaml:"sssdConfigFile,omitempty"`
				} `tfsdk:"sidecar" yaml:"sidecar,omitempty"`
			} `tfsdk:"sssd" yaml:"sssd,omitempty"`
		} `tfsdk:"security" yaml:"security,omitempty"`

		Server *struct {
			Active *int64 `tfsdk:"active" yaml:"active,omitempty"`

			Annotations utilities.Dynamic `tfsdk:"annotations" yaml:"annotations,omitempty"`

			Labels utilities.Dynamic `tfsdk:"labels" yaml:"labels,omitempty"`

			LogLevel *string `tfsdk:"log_level" yaml:"logLevel,omitempty"`

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

				TopologySpreadConstraints *[]struct {
					LabelSelector *struct {
						MatchExpressions *[]struct {
							Key *string `tfsdk:"key" yaml:"key,omitempty"`

							Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

							Values *[]string `tfsdk:"values" yaml:"values,omitempty"`
						} `tfsdk:"match_expressions" yaml:"matchExpressions,omitempty"`

						MatchLabels *map[string]string `tfsdk:"match_labels" yaml:"matchLabels,omitempty"`
					} `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

					MatchLabelKeys *[]string `tfsdk:"match_label_keys" yaml:"matchLabelKeys,omitempty"`

					MaxSkew *int64 `tfsdk:"max_skew" yaml:"maxSkew,omitempty"`

					MinDomains *int64 `tfsdk:"min_domains" yaml:"minDomains,omitempty"`

					NodeAffinityPolicy *string `tfsdk:"node_affinity_policy" yaml:"nodeAffinityPolicy,omitempty"`

					NodeTaintsPolicy *string `tfsdk:"node_taints_policy" yaml:"nodeTaintsPolicy,omitempty"`

					TopologyKey *string `tfsdk:"topology_key" yaml:"topologyKey,omitempty"`

					WhenUnsatisfiable *string `tfsdk:"when_unsatisfiable" yaml:"whenUnsatisfiable,omitempty"`
				} `tfsdk:"topology_spread_constraints" yaml:"topologySpreadConstraints,omitempty"`
			} `tfsdk:"placement" yaml:"placement,omitempty"`

			PriorityClassName *string `tfsdk:"priority_class_name" yaml:"priorityClassName,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`
		} `tfsdk:"server" yaml:"server,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewCephRookIoCephNFSV1Resource() resource.Resource {
	return &CephRookIoCephNFSV1Resource{}
}

func (r *CephRookIoCephNFSV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ceph_rook_io_ceph_nfs_v1"
}

func (r *CephRookIoCephNFSV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "CephNFS represents a Ceph NFS",
		MarkdownDescription: "CephNFS represents a Ceph NFS",
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
				Description:         "NFSGaneshaSpec represents the spec of an nfs ganesha server",
				MarkdownDescription: "NFSGaneshaSpec represents the spec of an nfs ganesha server",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"rados": {
						Description:         "RADOS is the Ganesha RADOS specification",
						MarkdownDescription: "RADOS is the Ganesha RADOS specification",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"namespace": {
								Description:         "The namespace inside the Ceph pool (set by 'pool') where shared NFS-Ganesha config is stored. This setting is required for Ceph v15 and ignored for Ceph v16. As of Ceph Pacific v16+, this is internally set to the name of the CephNFS.",
								MarkdownDescription: "The namespace inside the Ceph pool (set by 'pool') where shared NFS-Ganesha config is stored. This setting is required for Ceph v15 and ignored for Ceph v16. As of Ceph Pacific v16+, this is internally set to the name of the CephNFS.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"pool": {
								Description:         "The Ceph pool used store the shared configuration for NFS-Ganesha daemons. This setting is required for Ceph v15 and ignored for Ceph v16. As of Ceph Pacific 16.2.7+, this is internally hardcoded to '.nfs'.",
								MarkdownDescription: "The Ceph pool used store the shared configuration for NFS-Ganesha daemons. This setting is required for Ceph v15 and ignored for Ceph v16. As of Ceph Pacific 16.2.7+, this is internally hardcoded to '.nfs'.",

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

					"security": {
						Description:         "Security allows specifying security configurations for the NFS cluster",
						MarkdownDescription: "Security allows specifying security configurations for the NFS cluster",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"kerberos": {
								Description:         "Kerberos configures NFS-Ganesha to secure NFS client connections with Kerberos.",
								MarkdownDescription: "Kerberos configures NFS-Ganesha to secure NFS client connections with Kerberos.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"config_files": {
										Description:         "ConfigFiles defines where the Kerberos configuration should be sourced from. Config files will be placed into the '/etc/krb5.conf.rook/' directory.  If this is left empty, Rook will not add any files. This allows you to manage the files yourself however you wish. For example, you may build them into your custom Ceph container image or use the Vault agent injector to securely add the files via annotations on the CephNFS spec (passed to the NFS server pods).  Rook configures Kerberos to log to stderr. We suggest removing logging sections from config files to avoid consuming unnecessary disk space from logging to files.",
										MarkdownDescription: "ConfigFiles defines where the Kerberos configuration should be sourced from. Config files will be placed into the '/etc/krb5.conf.rook/' directory.  If this is left empty, Rook will not add any files. This allows you to manage the files yourself however you wish. For example, you may build them into your custom Ceph container image or use the Vault agent injector to securely add the files via annotations on the CephNFS spec (passed to the NFS server pods).  Rook configures Kerberos to log to stderr. We suggest removing logging sections from config files to avoid consuming unnecessary disk space from logging to files.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"volume_source": {
												Description:         "VolumeSource accepts a standard Kubernetes VolumeSource for Kerberos configuration files like what is normally used to configure Volumes for a Pod. For example, a ConfigMap, Secret, or HostPath. The volume may contain multiple files, all of which will be loaded.",
												MarkdownDescription: "VolumeSource accepts a standard Kubernetes VolumeSource for Kerberos configuration files like what is normally used to configure Volumes for a Pod. For example, a ConfigMap, Secret, or HostPath. The volume may contain multiple files, all of which will be loaded.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"annotations": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"finalizers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"labels": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespace": {
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

									"keytab_file": {
										Description:         "KeytabFile defines where the Kerberos keytab should be sourced from. The keytab file will be placed into '/etc/krb5.keytab'. If this is left empty, Rook will not add the file. This allows you to manage the 'krb5.keytab' file yourself however you wish. For example, you may build it into your custom Ceph container image or use the Vault agent injector to securely add the file via annotations on the CephNFS spec (passed to the NFS server pods).",
										MarkdownDescription: "KeytabFile defines where the Kerberos keytab should be sourced from. The keytab file will be placed into '/etc/krb5.keytab'. If this is left empty, Rook will not add the file. This allows you to manage the 'krb5.keytab' file yourself however you wish. For example, you may build it into your custom Ceph container image or use the Vault agent injector to securely add the file via annotations on the CephNFS spec (passed to the NFS server pods).",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"volume_source": {
												Description:         "VolumeSource accepts a standard Kubernetes VolumeSource for the Kerberos keytab file like what is normally used to configure Volumes for a Pod. For example, a Secret or HostPath. There are two requirements for the source's content:   1. The config file must be mountable via 'subPath: krb5.keytab'. For example, in a      Secret, the data item must be named 'krb5.keytab', or 'items' must be defined to      select the key and give it path 'krb5.keytab'. A HostPath directory must have the      'krb5.keytab' file.   2. The volume or config file must have mode 0600.",
												MarkdownDescription: "VolumeSource accepts a standard Kubernetes VolumeSource for the Kerberos keytab file like what is normally used to configure Volumes for a Pod. For example, a Secret or HostPath. There are two requirements for the source's content:   1. The config file must be mountable via 'subPath: krb5.keytab'. For example, in a      Secret, the data item must be named 'krb5.keytab', or 'items' must be defined to      select the key and give it path 'krb5.keytab'. A HostPath directory must have the      'krb5.keytab' file.   2. The volume or config file must have mode 0600.",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"annotations": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"finalizers": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"labels": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"name": {
																				Description:         "",
																				MarkdownDescription: "",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"namespace": {
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

									"principal_name": {
										Description:         "PrincipalName corresponds directly to NFS-Ganesha's NFS_KRB5:PrincipalName config. In practice, this is the service prefix of the principal name. The default is 'nfs'. This value is combined with (a) the namespace and name of the CephNFS (with a hyphen between) and (b) the Realm configured in the user-provided krb5.conf to determine the full principal name: <principalName>/<namespace>-<name>@<realm>. e.g., nfs/rook-ceph-my-nfs@example.net. See https://github.com/nfs-ganesha/nfs-ganesha/wiki/RPCSEC_GSS for more detail.",
										MarkdownDescription: "PrincipalName corresponds directly to NFS-Ganesha's NFS_KRB5:PrincipalName config. In practice, this is the service prefix of the principal name. The default is 'nfs'. This value is combined with (a) the namespace and name of the CephNFS (with a hyphen between) and (b) the Realm configured in the user-provided krb5.conf to determine the full principal name: <principalName>/<namespace>-<name>@<realm>. e.g., nfs/rook-ceph-my-nfs@example.net. See https://github.com/nfs-ganesha/nfs-ganesha/wiki/RPCSEC_GSS for more detail.",

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

							"sssd": {
								Description:         "SSSD enables integration with System Security Services Daemon (SSSD). SSSD can be used to provide user ID mapping from a number of sources. See https://sssd.io for more information about the SSSD project.",
								MarkdownDescription: "SSSD enables integration with System Security Services Daemon (SSSD). SSSD can be used to provide user ID mapping from a number of sources. See https://sssd.io for more information about the SSSD project.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"sidecar": {
										Description:         "Sidecar tells Rook to run SSSD in a sidecar alongside the NFS-Ganesha server in each NFS pod.",
										MarkdownDescription: "Sidecar tells Rook to run SSSD in a sidecar alongside the NFS-Ganesha server in each NFS pod.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"additional_files": {
												Description:         "AdditionalFiles defines any number of additional files that should be mounted into the SSSD sidecar. These files may be referenced by the sssd.conf config file.",
												MarkdownDescription: "AdditionalFiles defines any number of additional files that should be mounted into the SSSD sidecar. These files may be referenced by the sssd.conf config file.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"sub_path": {
														Description:         "SubPath defines the sub-path in '/etc/sssd/rook-additional/' where the additional file(s) will be placed. Each subPath definition must be unique and must not contain ':'.",
														MarkdownDescription: "SubPath defines the sub-path in '/etc/sssd/rook-additional/' where the additional file(s) will be placed. Each subPath definition must be unique and must not contain ':'.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),

															stringvalidator.RegexMatches(regexp.MustCompile(`^[^:]+$`), ""),
														},
													},

													"volume_source": {
														Description:         "VolumeSource accepts standard Kubernetes VolumeSource for the additional file(s) like what is normally used to configure Volumes for a Pod. Fore example, a ConfigMap, Secret, or HostPath. Each VolumeSource adds one or more additional files to the SSSD sidecar container in the '/etc/sssd/rook-additional/<subPath>' directory. Be aware that some files may need to have a specific file mode like 0600 due to requirements by SSSD for some files. For example, CA or TLS certificates.",
														MarkdownDescription: "VolumeSource accepts standard Kubernetes VolumeSource for the additional file(s) like what is normally used to configure Volumes for a Pod. Fore example, a ConfigMap, Secret, or HostPath. Each VolumeSource adds one or more additional files to the SSSD sidecar container in the '/etc/sssd/rook-additional/<subPath>' directory. Be aware that some files may need to have a specific file mode like 0600 due to requirements by SSSD for some files. For example, CA or TLS certificates.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"finalizers": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"namespace": {
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

														Required: true,
														Optional: false,
														Computed: false,
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"debug_level": {
												Description:         "DebugLevel sets the debug level for SSSD. If unset or set to 0, Rook does nothing. Otherwise, this may be a value between 1 and 10. See SSSD docs for more info: https://sssd.io/troubleshooting/basics.html#sssd-debug-logs",
												MarkdownDescription: "DebugLevel sets the debug level for SSSD. If unset or set to 0, Rook does nothing. Otherwise, this may be a value between 1 and 10. See SSSD docs for more info: https://sssd.io/troubleshooting/basics.html#sssd-debug-logs",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													int64validator.AtLeast(0),

													int64validator.AtMost(10),
												},
											},

											"image": {
												Description:         "Image defines the container image that should be used for the SSSD sidecar.",
												MarkdownDescription: "Image defines the container image that should be used for the SSSD sidecar.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),
												},
											},

											"resources": {
												Description:         "Resources allow specifying resource requests/limits on the SSSD sidecar container.",
												MarkdownDescription: "Resources allow specifying resource requests/limits on the SSSD sidecar container.",

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

											"sssd_config_file": {
												Description:         "SSSDConfigFile defines where the SSSD configuration should be sourced from. The config file will be placed into '/etc/sssd/sssd.conf'. If this is left empty, Rook will not add the file. This allows you to manage the 'sssd.conf' file yourself however you wish. For example, you may build it into your custom Ceph container image or use the Vault agent injector to securely add the file via annotations on the CephNFS spec (passed to the NFS server pods).",
												MarkdownDescription: "SSSDConfigFile defines where the SSSD configuration should be sourced from. The config file will be placed into '/etc/sssd/sssd.conf'. If this is left empty, Rook will not add the file. This allows you to manage the 'sssd.conf' file yourself however you wish. For example, you may build it into your custom Ceph container image or use the Vault agent injector to securely add the file via annotations on the CephNFS spec (passed to the NFS server pods).",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"volume_source": {
														Description:         "VolumeSource accepts a standard Kubernetes VolumeSource for the SSSD configuration file like what is normally used to configure Volumes for a Pod. For example, a ConfigMap, Secret, or HostPath. There are two requirements for the source's content:   1. The config file must be mountable via 'subPath: sssd.conf'. For example, in a ConfigMap,      the data item must be named 'sssd.conf', or 'items' must be defined to select the key      and give it path 'sssd.conf'. A HostPath directory must have the 'sssd.conf' file.   2. The volume or config file must have mode 0600.",
														MarkdownDescription: "VolumeSource accepts a standard Kubernetes VolumeSource for the SSSD configuration file like what is normally used to configure Volumes for a Pod. For example, a ConfigMap, Secret, or HostPath. There are two requirements for the source's content:   1. The config file must be mountable via 'subPath: sssd.conf'. For example, in a ConfigMap,      the data item must be named 'sssd.conf', or 'items' must be defined to select the key      and give it path 'sssd.conf'. A HostPath directory must have the 'sssd.conf' file.   2. The volume or config file must have mode 0600.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"annotations": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"finalizers": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.ListType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"labels": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.MapType{ElemType: types.StringType},

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"name": {
																						Description:         "",
																						MarkdownDescription: "",

																						Type: types.StringType,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"namespace": {
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
																Description:         "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
																MarkdownDescription: "hostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",

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

					"server": {
						Description:         "Server is the Ganesha Server specification",
						MarkdownDescription: "Server is the Ganesha Server specification",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"active": {
								Description:         "The number of active Ganesha servers",
								MarkdownDescription: "The number of active Ganesha servers",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"annotations": {
								Description:         "The annotations-related configuration to add/set on each Pod related object.",
								MarkdownDescription: "The annotations-related configuration to add/set on each Pod related object.",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"labels": {
								Description:         "The labels-related configuration to add/set on each Pod related object.",
								MarkdownDescription: "The labels-related configuration to add/set on each Pod related object.",

								Type: utilities.DynamicType{},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"log_level": {
								Description:         "LogLevel set logging level",
								MarkdownDescription: "LogLevel set logging level",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"placement": {
								Description:         "The affinity to place the ganesha pods",
								MarkdownDescription: "The affinity to place the ganesha pods",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"node_affinity": {
										Description:         "NodeAffinity is a group of node affinity scheduling rules",
										MarkdownDescription: "NodeAffinity is a group of node affinity scheduling rules",

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
										Description:         "PodAffinity is a group of inter pod affinity scheduling rules",
										MarkdownDescription: "PodAffinity is a group of inter pod affinity scheduling rules",

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
										Description:         "PodAntiAffinity is a group of inter pod anti affinity scheduling rules",
										MarkdownDescription: "PodAntiAffinity is a group of inter pod anti affinity scheduling rules",

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
										Description:         "The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>",
										MarkdownDescription: "The pod this Toleration is attached to tolerates any taint that matches the triple <key,value,effect> using the matching operator <operator>",

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

									"topology_spread_constraints": {
										Description:         "TopologySpreadConstraint specifies how to spread matching pods among the given topology",
										MarkdownDescription: "TopologySpreadConstraint specifies how to spread matching pods among the given topology",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"label_selector": {
												Description:         "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",
												MarkdownDescription: "LabelSelector is used to find matching pods. Pods that match this label selector are counted to determine the number of pods in their corresponding topology domain.",

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

											"match_label_keys": {
												Description:         "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.",
												MarkdownDescription: "MatchLabelKeys is a set of pod label keys to select the pods over which spreading will be calculated. The keys are used to lookup values from the incoming pod labels, those key-value labels are ANDed with labelSelector to select the group of existing pods over which spreading will be calculated for the incoming pod. Keys that don't exist in the incoming pod labels will be ignored. A null or empty list means only match against labelSelector.",

												Type: types.ListType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"max_skew": {
												Description:         "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",
												MarkdownDescription: "MaxSkew describes the degree to which pods may be unevenly distributed. When 'whenUnsatisfiable=DoNotSchedule', it is the maximum permitted difference between the number of matching pods in the target topology and the global minimum. The global minimum is the minimum number of matching pods in an eligible domain or zero if the number of eligible domains is less than MinDomains. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 2/2/1: In this case, the global minimum is 1. | zone1 | zone2 | zone3 | |  P P  |  P P  |   P   | - if MaxSkew is 1, incoming pod can only be scheduled to zone3 to become 2/2/2; scheduling it onto zone1(zone2) would make the ActualSkew(3-1) on zone1(zone2) violate MaxSkew(1). - if MaxSkew is 2, incoming pod can be scheduled onto any zone. When 'whenUnsatisfiable=ScheduleAnyway', it is used to give higher precedence to topologies that satisfy it. It's a required field. Default value is 1 and 0 is not allowed.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"min_domains": {
												Description:         "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",
												MarkdownDescription: "MinDomains indicates a minimum number of eligible domains. When the number of eligible domains with matching topology keys is less than minDomains, Pod Topology Spread treats 'global minimum' as 0, and then the calculation of Skew is performed. And when the number of eligible domains with matching topology keys equals or greater than minDomains, this value has no effect on scheduling. As a result, when the number of eligible domains is less than minDomains, scheduler won't schedule more than maxSkew Pods to those domains. If value is nil, the constraint behaves as if MinDomains is equal to 1. Valid values are integers greater than 0. When value is not nil, WhenUnsatisfiable must be DoNotSchedule.  For example, in a 3-zone cluster, MaxSkew is set to 2, MinDomains is set to 5 and pods with the same labelSelector spread as 2/2/2: | zone1 | zone2 | zone3 | |  P P  |  P P  |  P P  | The number of domains is less than 5(MinDomains), so 'global minimum' is treated as 0. In this situation, new pod with the same labelSelector cannot be scheduled, because computed skew will be 3(3 - 0) if new Pod is scheduled to any of the three zones, it will violate MaxSkew.  This is a beta field and requires the MinDomainsInPodTopologySpread feature gate to be enabled (enabled by default).",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_affinity_policy": {
												Description:         "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
												MarkdownDescription: "NodeAffinityPolicy indicates how we will treat Pod's nodeAffinity/nodeSelector when calculating pod topology spread skew. Options are: - Honor: only nodes matching nodeAffinity/nodeSelector are included in the calculations. - Ignore: nodeAffinity/nodeSelector are ignored. All nodes are included in the calculations.  If this value is nil, the behavior is equivalent to the Honor policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_taints_policy": {
												Description:         "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",
												MarkdownDescription: "NodeTaintsPolicy indicates how we will treat node taints when calculating pod topology spread skew. Options are: - Honor: nodes without taints, along with tainted nodes for which the incoming pod has a toleration, are included. - Ignore: node taints are ignored. All nodes are included.  If this value is nil, the behavior is equivalent to the Ignore policy. This is a alpha-level feature enabled by the NodeInclusionPolicyInPodTopologySpread feature flag.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"topology_key": {
												Description:         "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",
												MarkdownDescription: "TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each <key, value> as a 'bucket', and try to put balanced number of pods into each bucket. We define a domain as a particular instance of a topology. Also, we define an eligible domain as a domain whose nodes meet the requirements of nodeAffinityPolicy and nodeTaintsPolicy. e.g. If TopologyKey is 'kubernetes.io/hostname', each Node is a domain of that topology. And, if TopologyKey is 'topology.kubernetes.io/zone', each zone is a domain of that topology. It's a required field.",

												Type: types.StringType,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"when_unsatisfiable": {
												Description:         "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location,   but giving higher precedence to topologies that would help reduce the   skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",
												MarkdownDescription: "WhenUnsatisfiable indicates how to deal with a pod if it doesn't satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it. - ScheduleAnyway tells the scheduler to schedule the pod in any location,   but giving higher precedence to topologies that would help reduce the   skew. A constraint is considered 'Unsatisfiable' for an incoming pod if and only if every possible node assignment for that pod would violate 'MaxSkew' on some topology. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1: | zone1 | zone2 | zone3 | | P P P |   P   |   P   | If WhenUnsatisfiable is set to DoNotSchedule, incoming pod can only be scheduled to zone2(zone3) to become 3/2/1(3/1/2) as ActualSkew(2-1) on zone2(zone3) satisfies MaxSkew(1). In other words, the cluster can still be imbalanced, but scheduler won't make it *more* imbalanced. It's a required field.",

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

							"priority_class_name": {
								Description:         "PriorityClassName sets the priority class on the pods",
								MarkdownDescription: "PriorityClassName sets the priority class on the pods",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resources": {
								Description:         "Resources set resource requests and limits",
								MarkdownDescription: "Resources set resource requests and limits",

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
						}),

						Required: true,
						Optional: false,
						Computed: false,
					},
				}),

				Required: true,
				Optional: false,
				Computed: false,
			},
		},
	}, nil
}

func (r *CephRookIoCephNFSV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_ceph_rook_io_ceph_nfs_v1")

	var state CephRookIoCephNFSV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CephRookIoCephNFSV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ceph.rook.io/v1")
	goModel.Kind = utilities.Ptr("CephNFS")

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

func (r *CephRookIoCephNFSV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_ceph_rook_io_ceph_nfs_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *CephRookIoCephNFSV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_ceph_rook_io_ceph_nfs_v1")

	var state CephRookIoCephNFSV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel CephRookIoCephNFSV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("ceph.rook.io/v1")
	goModel.Kind = utilities.Ptr("CephNFS")

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

func (r *CephRookIoCephNFSV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_ceph_rook_io_ceph_nfs_v1")
	// NO-OP: Terraform removes the state automatically for us
}
