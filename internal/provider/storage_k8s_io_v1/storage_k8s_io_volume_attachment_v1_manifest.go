/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package storage_k8s_io_v1

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"
)

var (
	_ datasource.DataSource = &StorageK8SIoVolumeAttachmentV1Manifest{}
)

func NewStorageK8SIoVolumeAttachmentV1Manifest() datasource.DataSource {
	return &StorageK8SIoVolumeAttachmentV1Manifest{}
}

type StorageK8SIoVolumeAttachmentV1Manifest struct{}

type StorageK8SIoVolumeAttachmentV1ManifestData struct {
	ID   types.String `tfsdk:"id" json:"-"`
	YAML types.String `tfsdk:"yaml" json:"-"`

	ApiVersion *string `tfsdk:"-" json:"apiVersion"`
	Kind       *string `tfsdk:"-" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Attacher *string `tfsdk:"attacher" json:"attacher,omitempty"`
		NodeName *string `tfsdk:"node_name" json:"nodeName,omitempty"`
		Source   *struct {
			InlineVolumeSpec *struct {
				AccessModes          *[]string `tfsdk:"access_modes" json:"accessModes,omitempty"`
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
					ReadOnly        *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretName      *string `tfsdk:"secret_name" json:"secretName,omitempty"`
					SecretNamespace *string `tfsdk:"secret_namespace" json:"secretNamespace,omitempty"`
					ShareName       *string `tfsdk:"share_name" json:"shareName,omitempty"`
				} `tfsdk:"azure_file" json:"azureFile,omitempty"`
				Capacity *map[string]string `tfsdk:"capacity" json:"capacity,omitempty"`
				Cephfs   *struct {
					Monitors   *[]string `tfsdk:"monitors" json:"monitors,omitempty"`
					Path       *string   `tfsdk:"path" json:"path,omitempty"`
					ReadOnly   *bool     `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretFile *string   `tfsdk:"secret_file" json:"secretFile,omitempty"`
					SecretRef  *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					User *string `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"cephfs" json:"cephfs,omitempty"`
				Cinder *struct {
					FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
				} `tfsdk:"cinder" json:"cinder,omitempty"`
				ClaimRef *struct {
					ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
					FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
					Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
					Name            *string `tfsdk:"name" json:"name,omitempty"`
					Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
					ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
					Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
				} `tfsdk:"claim_ref" json:"claimRef,omitempty"`
				Csi *struct {
					ControllerExpandSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"controller_expand_secret_ref" json:"controllerExpandSecretRef,omitempty"`
					ControllerPublishSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"controller_publish_secret_ref" json:"controllerPublishSecretRef,omitempty"`
					Driver              *string `tfsdk:"driver" json:"driver,omitempty"`
					FsType              *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					NodeExpandSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"node_expand_secret_ref" json:"nodeExpandSecretRef,omitempty"`
					NodePublishSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"node_publish_secret_ref" json:"nodePublishSecretRef,omitempty"`
					NodeStageSecretRef *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"node_stage_secret_ref" json:"nodeStageSecretRef,omitempty"`
					ReadOnly         *bool              `tfsdk:"read_only" json:"readOnly,omitempty"`
					VolumeAttributes *map[string]string `tfsdk:"volume_attributes" json:"volumeAttributes,omitempty"`
					VolumeHandle     *string            `tfsdk:"volume_handle" json:"volumeHandle,omitempty"`
				} `tfsdk:"csi" json:"csi,omitempty"`
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
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
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
				Glusterfs *struct {
					Endpoints          *string `tfsdk:"endpoints" json:"endpoints,omitempty"`
					EndpointsNamespace *string `tfsdk:"endpoints_namespace" json:"endpointsNamespace,omitempty"`
					Path               *string `tfsdk:"path" json:"path,omitempty"`
					ReadOnly           *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
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
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					TargetPortal *string `tfsdk:"target_portal" json:"targetPortal,omitempty"`
				} `tfsdk:"iscsi" json:"iscsi,omitempty"`
				Local *struct {
					FsType *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					Path   *string `tfsdk:"path" json:"path,omitempty"`
				} `tfsdk:"local" json:"local,omitempty"`
				MountOptions *[]string `tfsdk:"mount_options" json:"mountOptions,omitempty"`
				Nfs          *struct {
					Path     *string `tfsdk:"path" json:"path,omitempty"`
					ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					Server   *string `tfsdk:"server" json:"server,omitempty"`
				} `tfsdk:"nfs" json:"nfs,omitempty"`
				NodeAffinity *struct {
					Required *struct {
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
					} `tfsdk:"required" json:"required,omitempty"`
				} `tfsdk:"node_affinity" json:"nodeAffinity,omitempty"`
				PersistentVolumeReclaimPolicy *string `tfsdk:"persistent_volume_reclaim_policy" json:"persistentVolumeReclaimPolicy,omitempty"`
				PhotonPersistentDisk          *struct {
					FsType *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					PdID   *string `tfsdk:"pd_id" json:"pdID,omitempty"`
				} `tfsdk:"photon_persistent_disk" json:"photonPersistentDisk,omitempty"`
				PortworxVolume *struct {
					FsType   *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					ReadOnly *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					VolumeID *string `tfsdk:"volume_id" json:"volumeID,omitempty"`
				} `tfsdk:"portworx_volume" json:"portworxVolume,omitempty"`
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
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					User *string `tfsdk:"user" json:"user,omitempty"`
				} `tfsdk:"rbd" json:"rbd,omitempty"`
				ScaleIO *struct {
					FsType           *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					Gateway          *string `tfsdk:"gateway" json:"gateway,omitempty"`
					ProtectionDomain *string `tfsdk:"protection_domain" json:"protectionDomain,omitempty"`
					ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretRef        *struct {
						Name      *string `tfsdk:"name" json:"name,omitempty"`
						Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					SslEnabled  *bool   `tfsdk:"ssl_enabled" json:"sslEnabled,omitempty"`
					StorageMode *string `tfsdk:"storage_mode" json:"storageMode,omitempty"`
					StoragePool *string `tfsdk:"storage_pool" json:"storagePool,omitempty"`
					System      *string `tfsdk:"system" json:"system,omitempty"`
					VolumeName  *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
				} `tfsdk:"scale_io" json:"scaleIO,omitempty"`
				StorageClassName *string `tfsdk:"storage_class_name" json:"storageClassName,omitempty"`
				Storageos        *struct {
					FsType    *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					ReadOnly  *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
					SecretRef *struct {
						ApiVersion      *string `tfsdk:"api_version" json:"apiVersion,omitempty"`
						FieldPath       *string `tfsdk:"field_path" json:"fieldPath,omitempty"`
						Kind            *string `tfsdk:"kind" json:"kind,omitempty"`
						Name            *string `tfsdk:"name" json:"name,omitempty"`
						Namespace       *string `tfsdk:"namespace" json:"namespace,omitempty"`
						ResourceVersion *string `tfsdk:"resource_version" json:"resourceVersion,omitempty"`
						Uid             *string `tfsdk:"uid" json:"uid,omitempty"`
					} `tfsdk:"secret_ref" json:"secretRef,omitempty"`
					VolumeName      *string `tfsdk:"volume_name" json:"volumeName,omitempty"`
					VolumeNamespace *string `tfsdk:"volume_namespace" json:"volumeNamespace,omitempty"`
				} `tfsdk:"storageos" json:"storageos,omitempty"`
				VolumeMode    *string `tfsdk:"volume_mode" json:"volumeMode,omitempty"`
				VsphereVolume *struct {
					FsType            *string `tfsdk:"fs_type" json:"fsType,omitempty"`
					StoragePolicyID   *string `tfsdk:"storage_policy_id" json:"storagePolicyID,omitempty"`
					StoragePolicyName *string `tfsdk:"storage_policy_name" json:"storagePolicyName,omitempty"`
					VolumePath        *string `tfsdk:"volume_path" json:"volumePath,omitempty"`
				} `tfsdk:"vsphere_volume" json:"vsphereVolume,omitempty"`
			} `tfsdk:"inline_volume_spec" json:"inlineVolumeSpec,omitempty"`
			PersistentVolumeName *string `tfsdk:"persistent_volume_name" json:"persistentVolumeName,omitempty"`
		} `tfsdk:"source" json:"source,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *StorageK8SIoVolumeAttachmentV1Manifest) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_storage_k8s_io_volume_attachment_v1_manifest"
}

func (r *StorageK8SIoVolumeAttachmentV1Manifest) Schema(_ context.Context, _ datasource.SchemaRequest, response *datasource.SchemaResponse) {
	response.Schema = schema.Schema{
		Description:         "VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.VolumeAttachment objects are non-namespaced.",
		MarkdownDescription: "VolumeAttachment captures the intent to attach or detach the specified volume to/from the specified node.VolumeAttachment objects are non-namespaced.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description:         "Contains the value 'metadata.name'.",
				MarkdownDescription: "Contains the value `metadata.name`.",
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
				Description:         "VolumeAttachmentSpec is the specification of a VolumeAttachment request.",
				MarkdownDescription: "VolumeAttachmentSpec is the specification of a VolumeAttachment request.",
				Attributes: map[string]schema.Attribute{
					"attacher": schema.StringAttribute{
						Description:         "attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().",
						MarkdownDescription: "attacher indicates the name of the volume driver that MUST handle this request. This is the name returned by GetPluginName().",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"node_name": schema.StringAttribute{
						Description:         "nodeName represents the node that the volume should be attached to.",
						MarkdownDescription: "nodeName represents the node that the volume should be attached to.",
						Required:            true,
						Optional:            false,
						Computed:            false,
					},

					"source": schema.SingleNestedAttribute{
						Description:         "VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.",
						MarkdownDescription: "VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.",
						Attributes: map[string]schema.Attribute{
							"inline_volume_spec": schema.SingleNestedAttribute{
								Description:         "PersistentVolumeSpec is the specification of a persistent volume.",
								MarkdownDescription: "PersistentVolumeSpec is the specification of a persistent volume.",
								Attributes: map[string]schema.Attribute{
									"access_modes": schema.ListAttribute{
										Description:         "accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes",
										MarkdownDescription: "accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"aws_elastic_block_store": schema.SingleNestedAttribute{
										Description:         "Represents a Persistent Disk resource in AWS.An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.",
										MarkdownDescription: "Represents a Persistent Disk resource in AWS.An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.",
										Attributes: map[string]schema.Attribute{
											"fs_type": schema.StringAttribute{
												Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
												MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
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
										Description:         "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
										MarkdownDescription: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
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
										Description:         "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
										MarkdownDescription: "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
										Attributes: map[string]schema.Attribute{
											"read_only": schema.BoolAttribute{
												Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret_name": schema.StringAttribute{
												Description:         "secretName is the name of secret that contains Azure Storage Account Name and Key",
												MarkdownDescription: "secretName is the name of secret that contains Azure Storage Account Name and Key",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"secret_namespace": schema.StringAttribute{
												Description:         "secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod",
												MarkdownDescription: "secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"share_name": schema.StringAttribute{
												Description:         "shareName is the azure Share Name",
												MarkdownDescription: "shareName is the azure Share Name",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"capacity": schema.MapAttribute{
										Description:         "capacity is the description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity",
										MarkdownDescription: "capacity is the description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"cephfs": schema.SingleNestedAttribute{
										Description:         "Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.",
										MarkdownDescription: "Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.",
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
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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
												Description:         "user is Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
												MarkdownDescription: "user is Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
										Description:         "Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.",
										MarkdownDescription: "Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.",
										Attributes: map[string]schema.Attribute{
											"fs_type": schema.StringAttribute{
												Description:         "fsType Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												MarkdownDescription: "fsType Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"read_only": schema.BoolAttribute{
												Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

									"claim_ref": schema.SingleNestedAttribute{
										Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
										MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",
										Attributes: map[string]schema.Attribute{
											"api_version": schema.StringAttribute{
												Description:         "API version of the referent.",
												MarkdownDescription: "API version of the referent.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"field_path": schema.StringAttribute{
												Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
												MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"kind": schema.StringAttribute{
												Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"name": schema.StringAttribute{
												Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"namespace": schema.StringAttribute{
												Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"resource_version": schema.StringAttribute{
												Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"uid": schema.StringAttribute{
												Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
												MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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
										Description:         "Represents storage that is managed by an external CSI volume driver (Beta feature)",
										MarkdownDescription: "Represents storage that is managed by an external CSI volume driver (Beta feature)",
										Attributes: map[string]schema.Attribute{
											"controller_expand_secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"controller_publish_secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"driver": schema.StringAttribute{
												Description:         "driver is the name of the driver to use for this volume. Required.",
												MarkdownDescription: "driver is the name of the driver to use for this volume. Required.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"fs_type": schema.StringAttribute{
												Description:         "fsType to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'.",
												MarkdownDescription: "fsType to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"node_expand_secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_publish_secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},
												},
												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_stage_secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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
												Description:         "readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).",
												MarkdownDescription: "readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"volume_attributes": schema.MapAttribute{
												Description:         "volumeAttributes of the volume to publish.",
												MarkdownDescription: "volumeAttributes of the volume to publish.",
												ElementType:         types.StringType,
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"volume_handle": schema.StringAttribute{
												Description:         "volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.",
												MarkdownDescription: "volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"fc": schema.SingleNestedAttribute{
										Description:         "Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.",
										MarkdownDescription: "Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.",
										Attributes: map[string]schema.Attribute{
											"fs_type": schema.StringAttribute{
												Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
												MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
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
										Description:         "FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.",
										MarkdownDescription: "FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.",
										Attributes: map[string]schema.Attribute{
											"driver": schema.StringAttribute{
												Description:         "driver is the name of the driver to use for this volume.",
												MarkdownDescription: "driver is the name of the driver to use for this volume.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"fs_type": schema.StringAttribute{
												Description:         "fsType is the Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
												MarkdownDescription: "fsType is the Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
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
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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
										Description:         "Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.",
										MarkdownDescription: "Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.",
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
										Description:         "Represents a Persistent Disk resource in Google Compute Engine.A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.",
										MarkdownDescription: "Represents a Persistent Disk resource in Google Compute Engine.A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.",
										Attributes: map[string]schema.Attribute{
											"fs_type": schema.StringAttribute{
												Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
												MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
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

									"glusterfs": schema.SingleNestedAttribute{
										Description:         "Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.",
										MarkdownDescription: "Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.",
										Attributes: map[string]schema.Attribute{
											"endpoints": schema.StringAttribute{
												Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
												MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"endpoints_namespace": schema.StringAttribute{
												Description:         "endpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
												MarkdownDescription: "endpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
												Required:            false,
												Optional:            true,
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
										Description:         "Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.",
										MarkdownDescription: "Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.",
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
										Description:         "ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.",
										MarkdownDescription: "ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.",
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
												Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi",
												MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi",
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
												Description:         "iqn is Target iSCSI Qualified Name.",
												MarkdownDescription: "iqn is Target iSCSI Qualified Name.",
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
												Description:         "lun is iSCSI Target Lun number.",
												MarkdownDescription: "lun is iSCSI Target Lun number.",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},

											"portals": schema.ListAttribute{
												Description:         "portals is the iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
												MarkdownDescription: "portals is the iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
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
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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

									"local": schema.SingleNestedAttribute{
										Description:         "Local represents directly-attached storage with node affinity (Beta feature)",
										MarkdownDescription: "Local represents directly-attached storage with node affinity (Beta feature)",
										Attributes: map[string]schema.Attribute{
											"fs_type": schema.StringAttribute{
												Description:         "fsType is the filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default value is to auto-select a filesystem if unspecified.",
												MarkdownDescription: "fsType is the filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default value is to auto-select a filesystem if unspecified.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"path": schema.StringAttribute{
												Description:         "path of the full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).",
												MarkdownDescription: "path of the full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).",
												Required:            true,
												Optional:            false,
												Computed:            false,
											},
										},
										Required: false,
										Optional: true,
										Computed: false,
									},

									"mount_options": schema.ListAttribute{
										Description:         "mountOptions is the list of mount options, e.g. ['ro', 'soft']. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options",
										MarkdownDescription: "mountOptions is the list of mount options, e.g. ['ro', 'soft']. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"nfs": schema.SingleNestedAttribute{
										Description:         "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",
										MarkdownDescription: "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",
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

									"node_affinity": schema.SingleNestedAttribute{
										Description:         "VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.",
										MarkdownDescription: "VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.",
										Attributes: map[string]schema.Attribute{
											"required": schema.SingleNestedAttribute{
												Description:         "A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.",
												MarkdownDescription: "A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.",
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

									"persistent_volume_reclaim_policy": schema.StringAttribute{
										Description:         "persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming",
										MarkdownDescription: "persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"photon_persistent_disk": schema.SingleNestedAttribute{
										Description:         "Represents a Photon Controller persistent disk resource.",
										MarkdownDescription: "Represents a Photon Controller persistent disk resource.",
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
										Description:         "PortworxVolumeSource represents a Portworx volume resource.",
										MarkdownDescription: "PortworxVolumeSource represents a Portworx volume resource.",
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

									"quobyte": schema.SingleNestedAttribute{
										Description:         "Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.",
										MarkdownDescription: "Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.",
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
										Description:         "Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.",
										MarkdownDescription: "Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.",
										Attributes: map[string]schema.Attribute{
											"fs_type": schema.StringAttribute{
												Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd",
												MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd",
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
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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
										Description:         "ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume",
										MarkdownDescription: "ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume",
										Attributes: map[string]schema.Attribute{
											"fs_type": schema.StringAttribute{
												Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'",
												MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'",
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
												Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
												Required:            false,
												Optional:            true,
												Computed:            false,
											},

											"secret_ref": schema.SingleNestedAttribute{
												Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
												Attributes: map[string]schema.Attribute{
													"name": schema.StringAttribute{
														Description:         "name is unique within a namespace to reference a secret resource.",
														MarkdownDescription: "name is unique within a namespace to reference a secret resource.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "namespace defines the space within which the secret name must be unique.",
														MarkdownDescription: "namespace defines the space within which the secret name must be unique.",
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
												Description:         "sslEnabled is the flag to enable/disable SSL communication with Gateway, default false",
												MarkdownDescription: "sslEnabled is the flag to enable/disable SSL communication with Gateway, default false",
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

									"storage_class_name": schema.StringAttribute{
										Description:         "storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.",
										MarkdownDescription: "storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"storageos": schema.SingleNestedAttribute{
										Description:         "Represents a StorageOS persistent volume resource.",
										MarkdownDescription: "Represents a StorageOS persistent volume resource.",
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
												Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
												MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",
												Attributes: map[string]schema.Attribute{
													"api_version": schema.StringAttribute{
														Description:         "API version of the referent.",
														MarkdownDescription: "API version of the referent.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"field_path": schema.StringAttribute{
														Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
														MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"kind": schema.StringAttribute{
														Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"name": schema.StringAttribute{
														Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"namespace": schema.StringAttribute{
														Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
														MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"resource_version": schema.StringAttribute{
														Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
														MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"uid": schema.StringAttribute{
														Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
														MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
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

									"volume_mode": schema.StringAttribute{
										Description:         "volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.",
										MarkdownDescription: "volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"vsphere_volume": schema.SingleNestedAttribute{
										Description:         "Represents a vSphere volume resource.",
										MarkdownDescription: "Represents a vSphere volume resource.",
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
								Required: false,
								Optional: true,
								Computed: false,
							},

							"persistent_volume_name": schema.StringAttribute{
								Description:         "persistentVolumeName represents the name of the persistent volume to attach.",
								MarkdownDescription: "persistentVolumeName represents the name of the persistent volume to attach.",
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

func (r *StorageK8SIoVolumeAttachmentV1Manifest) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_storage_k8s_io_volume_attachment_v1_manifest")

	var model StorageK8SIoVolumeAttachmentV1ManifestData
	response.Diagnostics.Append(request.Config.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(model.Metadata.Name)
	model.ApiVersion = pointer.String("storage.k8s.io/v1")
	model.Kind = pointer.String("VolumeAttachment")

	y, err := yaml.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"YAML Error: "+err.Error(),
		)
		return
	}
	model.YAML = types.StringValue(string(y))

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}
