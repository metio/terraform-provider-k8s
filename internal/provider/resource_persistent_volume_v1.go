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

type PersistentVolumeV1Resource struct{}

var (
	_ resource.Resource = (*PersistentVolumeV1Resource)(nil)
)

type PersistentVolumeV1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type PersistentVolumeV1GoModel struct {
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
		AccessModes *[]string `tfsdk:"access_modes" yaml:"accessModes,omitempty"`

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

			SecretNamespace *string `tfsdk:"secret_namespace" yaml:"secretNamespace,omitempty"`

			ShareName *string `tfsdk:"share_name" yaml:"shareName,omitempty"`
		} `tfsdk:"azure_file" yaml:"azureFile,omitempty"`

		Capacity *map[string]string `tfsdk:"capacity" yaml:"capacity,omitempty"`

		Cephfs *struct {
			Monitors *[]string `tfsdk:"monitors" yaml:"monitors,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

			SecretFile *string `tfsdk:"secret_file" yaml:"secretFile,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

			User *string `tfsdk:"user" yaml:"user,omitempty"`
		} `tfsdk:"cephfs" yaml:"cephfs,omitempty"`

		Cinder *struct {
			FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

			ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

			SecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

			VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
		} `tfsdk:"cinder" yaml:"cinder,omitempty"`

		ClaimRef *struct {
			ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

			FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

			Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

			ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

			Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
		} `tfsdk:"claim_ref" yaml:"claimRef,omitempty"`

		Csi *struct {
			ControllerExpandSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"controller_expand_secret_ref" yaml:"controllerExpandSecretRef,omitempty"`

			ControllerPublishSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"controller_publish_secret_ref" yaml:"controllerPublishSecretRef,omitempty"`

			Driver *string `tfsdk:"driver" yaml:"driver,omitempty"`

			FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

			NodeExpandSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"node_expand_secret_ref" yaml:"nodeExpandSecretRef,omitempty"`

			NodePublishSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"node_publish_secret_ref" yaml:"nodePublishSecretRef,omitempty"`

			NodeStageSecretRef *struct {
				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"node_stage_secret_ref" yaml:"nodeStageSecretRef,omitempty"`

			ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

			VolumeAttributes *map[string]string `tfsdk:"volume_attributes" yaml:"volumeAttributes,omitempty"`

			VolumeHandle *string `tfsdk:"volume_handle" yaml:"volumeHandle,omitempty"`
		} `tfsdk:"csi" yaml:"csi,omitempty"`

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

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
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

		Glusterfs *struct {
			Endpoints *string `tfsdk:"endpoints" yaml:"endpoints,omitempty"`

			EndpointsNamespace *string `tfsdk:"endpoints_namespace" yaml:"endpointsNamespace,omitempty"`

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

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

			TargetPortal *string `tfsdk:"target_portal" yaml:"targetPortal,omitempty"`
		} `tfsdk:"iscsi" yaml:"iscsi,omitempty"`

		Local *struct {
			FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

			Path *string `tfsdk:"path" yaml:"path,omitempty"`
		} `tfsdk:"local" yaml:"local,omitempty"`

		MountOptions *[]string `tfsdk:"mount_options" yaml:"mountOptions,omitempty"`

		Nfs *struct {
			Path *string `tfsdk:"path" yaml:"path,omitempty"`

			ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

			Server *string `tfsdk:"server" yaml:"server,omitempty"`
		} `tfsdk:"nfs" yaml:"nfs,omitempty"`

		NodeAffinity *struct {
			Required *struct {
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
			} `tfsdk:"required" yaml:"required,omitempty"`
		} `tfsdk:"node_affinity" yaml:"nodeAffinity,omitempty"`

		PersistentVolumeReclaimPolicy *string `tfsdk:"persistent_volume_reclaim_policy" yaml:"persistentVolumeReclaimPolicy,omitempty"`

		PhotonPersistentDisk *struct {
			FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

			PdID *string `tfsdk:"pd_id" yaml:"pdID,omitempty"`
		} `tfsdk:"photon_persistent_disk" yaml:"photonPersistentDisk,omitempty"`

		PortworxVolume *struct {
			FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

			ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

			VolumeID *string `tfsdk:"volume_id" yaml:"volumeID,omitempty"`
		} `tfsdk:"portworx_volume" yaml:"portworxVolume,omitempty"`

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

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
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

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

			SslEnabled *bool `tfsdk:"ssl_enabled" yaml:"sslEnabled,omitempty"`

			StorageMode *string `tfsdk:"storage_mode" yaml:"storageMode,omitempty"`

			StoragePool *string `tfsdk:"storage_pool" yaml:"storagePool,omitempty"`

			System *string `tfsdk:"system" yaml:"system,omitempty"`

			VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`
		} `tfsdk:"scale_io" yaml:"scaleIO,omitempty"`

		StorageClassName *string `tfsdk:"storage_class_name" yaml:"storageClassName,omitempty"`

		Storageos *struct {
			FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

			ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

			SecretRef *struct {
				ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

				FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`

				Kind *string `tfsdk:"kind" yaml:"kind,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

				ResourceVersion *string `tfsdk:"resource_version" yaml:"resourceVersion,omitempty"`

				Uid *string `tfsdk:"uid" yaml:"uid,omitempty"`
			} `tfsdk:"secret_ref" yaml:"secretRef,omitempty"`

			VolumeName *string `tfsdk:"volume_name" yaml:"volumeName,omitempty"`

			VolumeNamespace *string `tfsdk:"volume_namespace" yaml:"volumeNamespace,omitempty"`
		} `tfsdk:"storageos" yaml:"storageos,omitempty"`

		VolumeMode *string `tfsdk:"volume_mode" yaml:"volumeMode,omitempty"`

		VsphereVolume *struct {
			FsType *string `tfsdk:"fs_type" yaml:"fsType,omitempty"`

			StoragePolicyID *string `tfsdk:"storage_policy_id" yaml:"storagePolicyID,omitempty"`

			StoragePolicyName *string `tfsdk:"storage_policy_name" yaml:"storagePolicyName,omitempty"`

			VolumePath *string `tfsdk:"volume_path" yaml:"volumePath,omitempty"`
		} `tfsdk:"vsphere_volume" yaml:"vsphereVolume,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewPersistentVolumeV1Resource() resource.Resource {
	return &PersistentVolumeV1Resource{}
}

func (r *PersistentVolumeV1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_persistent_volume_v1"
}

func (r *PersistentVolumeV1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes",
		MarkdownDescription: "PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes",
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
				Description:         "PersistentVolumeSpec is the specification of a persistent volume.",
				MarkdownDescription: "PersistentVolumeSpec is the specification of a persistent volume.",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"access_modes": {
						Description:         "accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes",
						MarkdownDescription: "accessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"aws_elastic_block_store": {
						Description:         "Represents a Persistent Disk resource in AWS.An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.",
						MarkdownDescription: "Represents a Persistent Disk resource in AWS.An AWS EBS disk must exist before mounting to a container. The disk must also be in the same AWS zone as the kubelet. An AWS EBS disk can only be mounted as read/write once. AWS EBS volumes support ownership management and SELinux relabeling.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"fs_type": {
								Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
								MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

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
						Description:         "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",
						MarkdownDescription: "AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.",

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
						Description:         "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",
						MarkdownDescription: "AzureFile represents an Azure File Service mount on the host and bind mount to the pod.",

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
								Description:         "secretName is the name of secret that contains Azure Storage Account Name and Key",
								MarkdownDescription: "secretName is the name of secret that contains Azure Storage Account Name and Key",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"secret_namespace": {
								Description:         "secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod",
								MarkdownDescription: "secretNamespace is the namespace of the secret that contains Azure Storage Account Name and Key default is the same as the Pod",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"share_name": {
								Description:         "shareName is the azure Share Name",
								MarkdownDescription: "shareName is the azure Share Name",

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

					"capacity": {
						Description:         "capacity is the description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity",
						MarkdownDescription: "capacity is the description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity",

						Type: types.MapType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"cephfs": {
						Description:         "Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.",
						MarkdownDescription: "Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.",

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
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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
								Description:         "user is Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
								MarkdownDescription: "user is Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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
						Description:         "Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.",
						MarkdownDescription: "Represents a cinder volume resource in Openstack. A Cinder volume must exist before mounting to a container. The volume must also be in the same region as the kubelet. Cinder volumes support ownership management and SELinux relabeling.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"fs_type": {
								Description:         "fsType Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
								MarkdownDescription: "fsType Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"read_only": {
								Description:         "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
								MarkdownDescription: "readOnly is Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_ref": {
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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

					"claim_ref": {
						Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
						MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"api_version": {
								Description:         "API version of the referent.",
								MarkdownDescription: "API version of the referent.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"field_path": {
								Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
								MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"kind": {
								Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
								MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"name": {
								Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
								MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"namespace": {
								Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
								MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"resource_version": {
								Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
								MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"uid": {
								Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
								MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",

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

					"csi": {
						Description:         "Represents storage that is managed by an external CSI volume driver (Beta feature)",
						MarkdownDescription: "Represents storage that is managed by an external CSI volume driver (Beta feature)",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"controller_expand_secret_ref": {
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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

							"controller_publish_secret_ref": {
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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

							"driver": {
								Description:         "driver is the name of the driver to use for this volume. Required.",
								MarkdownDescription: "driver is the name of the driver to use for this volume. Required.",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"fs_type": {
								Description:         "fsType to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'.",
								MarkdownDescription: "fsType to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"node_expand_secret_ref": {
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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

							"node_publish_secret_ref": {
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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

							"node_stage_secret_ref": {
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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
								Description:         "readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).",
								MarkdownDescription: "readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_attributes": {
								Description:         "volumeAttributes of the volume to publish.",
								MarkdownDescription: "volumeAttributes of the volume to publish.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"volume_handle": {
								Description:         "volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.",
								MarkdownDescription: "volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.",

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

					"fc": {
						Description:         "Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.",
						MarkdownDescription: "Represents a Fibre Channel volume. Fibre Channel volumes can only be mounted as read/write once. Fibre Channel volumes support ownership management and SELinux relabeling.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"fs_type": {
								Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
								MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

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
						Description:         "FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.",
						MarkdownDescription: "FlexPersistentVolumeSource represents a generic persistent volume resource that is provisioned/attached using an exec based plugin.",

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
								Description:         "fsType is the Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
								MarkdownDescription: "fsType is the Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",

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
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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
						Description:         "Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.",
						MarkdownDescription: "Represents a Flocker volume mounted by the Flocker agent. One and only one of datasetName and datasetUUID should be set. Flocker volumes do not support ownership management or SELinux relabeling.",

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
						Description:         "Represents a Persistent Disk resource in Google Compute Engine.A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.",
						MarkdownDescription: "Represents a Persistent Disk resource in Google Compute Engine.A GCE PD must exist before mounting to a container. The disk must also be in the same GCE project and zone as the kubelet. A GCE PD can only be mounted as read/write once or read-only many times. GCE PDs support ownership management and SELinux relabeling.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"fs_type": {
								Description:         "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
								MarkdownDescription: "fsType is filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

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

					"glusterfs": {
						Description:         "Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.",
						MarkdownDescription: "Represents a Glusterfs mount that lasts the lifetime of a pod. Glusterfs volumes do not support ownership management or SELinux relabeling.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"endpoints": {
								Description:         "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
								MarkdownDescription: "endpoints is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"endpoints_namespace": {
								Description:         "endpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
								MarkdownDescription: "endpointsNamespace is the namespace that contains Glusterfs endpoint. If this field is empty, the EndpointNamespace defaults to the same namespace as the bound PVC. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

								Type: types.StringType,

								Required: false,
								Optional: true,
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
						Description:         "Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.",
						MarkdownDescription: "Represents a host path mapped into a pod. Host path volumes do not support ownership management or SELinux relabeling.",

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
						Description:         "ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.",
						MarkdownDescription: "ISCSIPersistentVolumeSource represents an ISCSI disk. ISCSI volumes can only be mounted as read/write once. ISCSI volumes support ownership management and SELinux relabeling.",

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
								Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi",
								MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi",

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
								Description:         "iqn is Target iSCSI Qualified Name.",
								MarkdownDescription: "iqn is Target iSCSI Qualified Name.",

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
								Description:         "lun is iSCSI Target Lun number.",
								MarkdownDescription: "lun is iSCSI Target Lun number.",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"portals": {
								Description:         "portals is the iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
								MarkdownDescription: "portals is the iSCSI Target Portal List. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

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
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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

					"local": {
						Description:         "Local represents directly-attached storage with node affinity (Beta feature)",
						MarkdownDescription: "Local represents directly-attached storage with node affinity (Beta feature)",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"fs_type": {
								Description:         "fsType is the filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default value is to auto-select a filesystem if unspecified.",
								MarkdownDescription: "fsType is the filesystem type to mount. It applies only when the Path is a block device. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default value is to auto-select a filesystem if unspecified.",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"path": {
								Description:         "path of the full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).",
								MarkdownDescription: "path of the full path to the volume on the node. It can be either a directory or block device (disk, partition, ...).",

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

					"mount_options": {
						Description:         "mountOptions is the list of mount options, e.g. ['ro', 'soft']. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options",
						MarkdownDescription: "mountOptions is the list of mount options, e.g. ['ro', 'soft']. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options",

						Type: types.ListType{ElemType: types.StringType},

						Required: false,
						Optional: true,
						Computed: false,
					},

					"nfs": {
						Description:         "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",
						MarkdownDescription: "Represents an NFS mount that lasts the lifetime of a pod. NFS volumes do not support ownership management or SELinux relabeling.",

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

					"node_affinity": {
						Description:         "VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.",
						MarkdownDescription: "VolumeNodeAffinity defines constraints that limit what nodes this volume can be accessed from.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"required": {
								Description:         "A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.",
								MarkdownDescription: "A node selector represents the union of the results of one or more label queries over a set of nodes; that is, it represents the OR of the selectors represented by the node selector terms.",

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

					"persistent_volume_reclaim_policy": {
						Description:         "persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming",
						MarkdownDescription: "persistentVolumeReclaimPolicy defines what happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"photon_persistent_disk": {
						Description:         "Represents a Photon Controller persistent disk resource.",
						MarkdownDescription: "Represents a Photon Controller persistent disk resource.",

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
						Description:         "PortworxVolumeSource represents a Portworx volume resource.",
						MarkdownDescription: "PortworxVolumeSource represents a Portworx volume resource.",

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

					"quobyte": {
						Description:         "Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.",
						MarkdownDescription: "Represents a Quobyte mount that lasts the lifetime of a pod. Quobyte volumes do not support ownership management or SELinux relabeling.",

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
						Description:         "Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.",
						MarkdownDescription: "Represents a Rados Block Device mount that lasts the lifetime of a pod. RBD volumes support ownership management and SELinux relabeling.",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"fs_type": {
								Description:         "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd",
								MarkdownDescription: "fsType is the filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd",

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
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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
						Description:         "ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume",
						MarkdownDescription: "ScaleIOPersistentVolumeSource represents a persistent ScaleIO volume",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"fs_type": {
								Description:         "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'",
								MarkdownDescription: "fsType is the filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'",

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
								Description:         "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
								MarkdownDescription: "readOnly defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"secret_ref": {
								Description:         "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",
								MarkdownDescription: "SecretReference represents a Secret Reference. It has enough information to retrieve secret in any namespace",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"name": {
										Description:         "name is unique within a namespace to reference a secret resource.",
										MarkdownDescription: "name is unique within a namespace to reference a secret resource.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "namespace defines the space within which the secret name must be unique.",
										MarkdownDescription: "namespace defines the space within which the secret name must be unique.",

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
								Description:         "sslEnabled is the flag to enable/disable SSL communication with Gateway, default false",
								MarkdownDescription: "sslEnabled is the flag to enable/disable SSL communication with Gateway, default false",

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

					"storage_class_name": {
						Description:         "storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.",
						MarkdownDescription: "storageClassName is the name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"storageos": {
						Description:         "Represents a StorageOS persistent volume resource.",
						MarkdownDescription: "Represents a StorageOS persistent volume resource.",

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
								Description:         "ObjectReference contains enough information to let you inspect or modify the referred object.",
								MarkdownDescription: "ObjectReference contains enough information to let you inspect or modify the referred object.",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"api_version": {
										Description:         "API version of the referent.",
										MarkdownDescription: "API version of the referent.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"field_path": {
										Description:         "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",
										MarkdownDescription: "If referring to a piece of an object instead of an entire object, this string should contain a valid JSON/Go field access statement, such as desiredState.manifest.containers[2]. For example, if the object reference is to a container within a pod, this would take on a value like: 'spec.containers{name}' (where 'name' refers to the name of the container that triggered the event) or if no container name is specified 'spec.containers[2]' (container with index 2 in this pod). This syntax is chosen only to have some well-defined way of referencing a part of an object.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"kind": {
										Description:         "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
										MarkdownDescription: "Kind of the referent. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
										MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"namespace": {
										Description:         "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
										MarkdownDescription: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"resource_version": {
										Description:         "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",
										MarkdownDescription: "Specific resourceVersion to which this reference is made, if any. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#concurrency-control-and-consistency",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"uid": {
										Description:         "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",
										MarkdownDescription: "UID of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#uids",

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

					"volume_mode": {
						Description:         "volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.",
						MarkdownDescription: "volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec.",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"vsphere_volume": {
						Description:         "Represents a vSphere volume resource.",
						MarkdownDescription: "Represents a vSphere volume resource.",

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
		},
	}, nil
}

func (r *PersistentVolumeV1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_persistent_volume_v1")

	var state PersistentVolumeV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel PersistentVolumeV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("/v1")
	goModel.Kind = utilities.Ptr("PersistentVolume")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *PersistentVolumeV1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_persistent_volume_v1")
	// NO-OP: All data is already in Terraform state
}

func (r *PersistentVolumeV1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_persistent_volume_v1")

	var state PersistentVolumeV1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel PersistentVolumeV1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("/v1")
	goModel.Kind = utilities.Ptr("PersistentVolume")

	state.Id = types.Int64Value(time.Now().UnixNano())
	state.ApiVersion = types.StringValue(*goModel.ApiVersion)
	state.Kind = types.StringValue(*goModel.Kind)

	marshal, err := yaml.Marshal(goModel)
	if err != nil {
		resp.Diagnostics.AddError("Could not generate YAML", err.Error())
		return
	}
	state.YAML = types.StringValue(string(marshal))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *PersistentVolumeV1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_persistent_volume_v1")
	// NO-OP: Terraform removes the state automatically for us
}
