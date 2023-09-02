/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package litmuschaos_io_v1alpha1

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/metio/terraform-provider-k8s/internal/utilities"
	"github.com/metio/terraform-provider-k8s/internal/validators"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sSchema "k8s.io/apimachinery/pkg/runtime/schema"
	k8sTypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/utils/pointer"
	"regexp"
	"strings"
)

var (
	_ resource.Resource                = &LitmuschaosIoChaosEngineV1Alpha1Resource{}
	_ resource.ResourceWithConfigure   = &LitmuschaosIoChaosEngineV1Alpha1Resource{}
	_ resource.ResourceWithImportState = &LitmuschaosIoChaosEngineV1Alpha1Resource{}
)

func NewLitmuschaosIoChaosEngineV1Alpha1Resource() resource.Resource {
	return &LitmuschaosIoChaosEngineV1Alpha1Resource{}
}

type LitmuschaosIoChaosEngineV1Alpha1Resource struct {
	kubernetesClient dynamic.Interface
	fieldManager     string
	forceConflicts   bool
}

type LitmuschaosIoChaosEngineV1Alpha1ResourceData struct {
	ID             types.String `tfsdk:"id" json:"-"`
	ForceConflicts types.Bool   `tfsdk:"force_conflicts" json:"-"`
	FieldManager   types.String `tfsdk:"field_manager" json:"-"`
	WaitFor        types.List   `tfsdk:"wait_for" json:"-"`

	ApiVersion *string `tfsdk:"api_version" json:"apiVersion"`
	Kind       *string `tfsdk:"kind" json:"kind"`

	Metadata struct {
		Name        string            `tfsdk:"name" json:"name"`
		Namespace   string            `tfsdk:"namespace" json:"namespace"`
		Labels      map[string]string `tfsdk:"labels" json:"labels,omitempty"`
		Annotations map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
	} `tfsdk:"metadata" json:"metadata"`

	Spec *struct {
		Appinfo *struct {
			Appkind  *string `tfsdk:"appkind" json:"appkind,omitempty"`
			Applabel *string `tfsdk:"applabel" json:"applabel,omitempty"`
			Appns    *string `tfsdk:"appns" json:"appns,omitempty"`
		} `tfsdk:"appinfo" json:"appinfo,omitempty"`
		AuxiliaryAppInfo    *string `tfsdk:"auxiliary_app_info" json:"auxiliaryAppInfo,omitempty"`
		ChaosServiceAccount *string `tfsdk:"chaos_service_account" json:"chaosServiceAccount,omitempty"`
		Components          *struct {
			Runner *struct {
				Image             *string            `tfsdk:"image" json:"image,omitempty"`
				RunnerAnnotations *map[string]string `tfsdk:"runner_annotations" json:"runnerAnnotations,omitempty"`
				RunnerLabels      *map[string]string `tfsdk:"runner_labels" json:"runnerLabels,omitempty"`
				Tolerations       *[]struct {
					Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
					Key               *string `tfsdk:"key" json:"key,omitempty"`
					Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
					TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
					Value             *string `tfsdk:"value" json:"value,omitempty"`
				} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				Type *string `tfsdk:"type" json:"type,omitempty"`
			} `tfsdk:"runner" json:"runner,omitempty"`
			Sidecar *[]struct {
				Env *[]struct {
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
				Image           *string `tfsdk:"image" json:"image,omitempty"`
				ImagePullPolicy *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
				Secrets         *[]struct {
					MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
					Name      *string `tfsdk:"name" json:"name,omitempty"`
				} `tfsdk:"secrets" json:"secrets,omitempty"`
			} `tfsdk:"sidecar" json:"sidecar,omitempty"`
		} `tfsdk:"components" json:"components,omitempty"`
		DefaultHealthCheck *bool   `tfsdk:"default_health_check" json:"defaultHealthCheck,omitempty"`
		EngineState        *string `tfsdk:"engine_state" json:"engineState,omitempty"`
		Experiments        *[]struct {
			Name *string `tfsdk:"name" json:"name,omitempty"`
			Spec *struct {
				Components *struct {
					ConfigMaps *[]struct {
						MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"config_maps" json:"configMaps,omitempty"`
					Env *[]struct {
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
					ExperimentAnnotations *map[string]string `tfsdk:"experiment_annotations" json:"experimentAnnotations,omitempty"`
					ExperimentImage       *string            `tfsdk:"experiment_image" json:"experimentImage,omitempty"`
					NodeSelector          *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
					Secrets               *[]struct {
						MountPath *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
						Name      *string `tfsdk:"name" json:"name,omitempty"`
					} `tfsdk:"secrets" json:"secrets,omitempty"`
					StatusCheckTimeouts *struct {
						Delay   *int64 `tfsdk:"delay" json:"delay,omitempty"`
						Timeout *int64 `tfsdk:"timeout" json:"timeout,omitempty"`
					} `tfsdk:"status_check_timeouts" json:"statusCheckTimeouts,omitempty"`
					Tolerations *[]struct {
						Effect            *string `tfsdk:"effect" json:"effect,omitempty"`
						Key               *string `tfsdk:"key" json:"key,omitempty"`
						Operator          *string `tfsdk:"operator" json:"operator,omitempty"`
						TolerationSeconds *int64  `tfsdk:"toleration_seconds" json:"tolerationSeconds,omitempty"`
						Value             *string `tfsdk:"value" json:"value,omitempty"`
					} `tfsdk:"tolerations" json:"tolerations,omitempty"`
				} `tfsdk:"components" json:"components,omitempty"`
				Probe *[]struct {
					CmdProbe_inputs *struct {
						Command    *string `tfsdk:"command" json:"command,omitempty"`
						Comparator *struct {
							Criteria *string `tfsdk:"criteria" json:"criteria,omitempty"`
							Type     *string `tfsdk:"type" json:"type,omitempty"`
							Value    *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"comparator" json:"comparator,omitempty"`
						Source *struct {
							Annotations *map[string]string `tfsdk:"annotations" json:"annotations,omitempty"`
							Args        *[]string          `tfsdk:"args" json:"args,omitempty"`
							Command     *[]string          `tfsdk:"command" json:"command,omitempty"`
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
							HostNetwork      *bool   `tfsdk:"host_network" json:"hostNetwork,omitempty"`
							Image            *string `tfsdk:"image" json:"image,omitempty"`
							ImagePullPolicy  *string `tfsdk:"image_pull_policy" json:"imagePullPolicy,omitempty"`
							ImagePullSecrets *[]struct {
								Name *string `tfsdk:"name" json:"name,omitempty"`
							} `tfsdk:"image_pull_secrets" json:"imagePullSecrets,omitempty"`
							InheritInputs *bool              `tfsdk:"inherit_inputs" json:"inheritInputs,omitempty"`
							Labels        *map[string]string `tfsdk:"labels" json:"labels,omitempty"`
							NodeSelector  *map[string]string `tfsdk:"node_selector" json:"nodeSelector,omitempty"`
							Privileged    *bool              `tfsdk:"privileged" json:"privileged,omitempty"`
							VolumeMount   *[]struct {
								MountPath        *string `tfsdk:"mount_path" json:"mountPath,omitempty"`
								MountPropagation *string `tfsdk:"mount_propagation" json:"mountPropagation,omitempty"`
								Name             *string `tfsdk:"name" json:"name,omitempty"`
								ReadOnly         *bool   `tfsdk:"read_only" json:"readOnly,omitempty"`
								SubPath          *string `tfsdk:"sub_path" json:"subPath,omitempty"`
								SubPathExpr      *string `tfsdk:"sub_path_expr" json:"subPathExpr,omitempty"`
							} `tfsdk:"volume_mount" json:"volumeMount,omitempty"`
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
						} `tfsdk:"source" json:"source,omitempty"`
					} `tfsdk:"cmd_probe_inputs" json:"cmdProbe/inputs,omitempty"`
					Data             *string `tfsdk:"data" json:"data,omitempty"`
					HttpProbe_inputs *struct {
						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						Method             *struct {
							Get *struct {
								Criteria     *string `tfsdk:"criteria" json:"criteria,omitempty"`
								ResponseCode *string `tfsdk:"response_code" json:"responseCode,omitempty"`
							} `tfsdk:"get" json:"get,omitempty"`
							Post *struct {
								Body         *string `tfsdk:"body" json:"body,omitempty"`
								BodyPath     *string `tfsdk:"body_path" json:"bodyPath,omitempty"`
								ContentType  *string `tfsdk:"content_type" json:"contentType,omitempty"`
								Criteria     *string `tfsdk:"criteria" json:"criteria,omitempty"`
								ResponseCode *string `tfsdk:"response_code" json:"responseCode,omitempty"`
							} `tfsdk:"post" json:"post,omitempty"`
						} `tfsdk:"method" json:"method,omitempty"`
						Url *string `tfsdk:"url" json:"url,omitempty"`
					} `tfsdk:"http_probe_inputs" json:"httpProbe/inputs,omitempty"`
					K8sProbe_inputs *struct {
						FieldSelector *string `tfsdk:"field_selector" json:"fieldSelector,omitempty"`
						Group         *string `tfsdk:"group" json:"group,omitempty"`
						LabelSelector *string `tfsdk:"label_selector" json:"labelSelector,omitempty"`
						Namespace     *string `tfsdk:"namespace" json:"namespace,omitempty"`
						Operation     *string `tfsdk:"operation" json:"operation,omitempty"`
						Resource      *string `tfsdk:"resource" json:"resource,omitempty"`
						ResourceNames *string `tfsdk:"resource_names" json:"resourceNames,omitempty"`
						Version       *string `tfsdk:"version" json:"version,omitempty"`
					} `tfsdk:"k8s_probe_inputs" json:"k8sProbe/inputs,omitempty"`
					Mode             *string `tfsdk:"mode" json:"mode,omitempty"`
					Name             *string `tfsdk:"name" json:"name,omitempty"`
					PromProbe_inputs *struct {
						Comparator *struct {
							Criteria *string `tfsdk:"criteria" json:"criteria,omitempty"`
							Value    *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"comparator" json:"comparator,omitempty"`
						Endpoint  *string `tfsdk:"endpoint" json:"endpoint,omitempty"`
						Query     *string `tfsdk:"query" json:"query,omitempty"`
						QueryPath *string `tfsdk:"query_path" json:"queryPath,omitempty"`
					} `tfsdk:"prom_probe_inputs" json:"promProbe/inputs,omitempty"`
					RunProperties *struct {
						Attempt              *int64  `tfsdk:"attempt" json:"attempt,omitempty"`
						EvaluationTimeout    *string `tfsdk:"evaluation_timeout" json:"evaluationTimeout,omitempty"`
						InitialDelay         *string `tfsdk:"initial_delay" json:"initialDelay,omitempty"`
						InitialDelaySeconds  *int64  `tfsdk:"initial_delay_seconds" json:"initialDelaySeconds,omitempty"`
						Interval             *string `tfsdk:"interval" json:"interval,omitempty"`
						ProbePollingInterval *string `tfsdk:"probe_polling_interval" json:"probePollingInterval,omitempty"`
						ProbeTimeout         *string `tfsdk:"probe_timeout" json:"probeTimeout,omitempty"`
						Retry                *int64  `tfsdk:"retry" json:"retry,omitempty"`
						StopOnFailure        *bool   `tfsdk:"stop_on_failure" json:"stopOnFailure,omitempty"`
					} `tfsdk:"run_properties" json:"runProperties,omitempty"`
					SloProbe_inputs *struct {
						Comparator *struct {
							Criteria *string `tfsdk:"criteria" json:"criteria,omitempty"`
							Type     *string `tfsdk:"type" json:"type,omitempty"`
							Value    *string `tfsdk:"value" json:"value,omitempty"`
						} `tfsdk:"comparator" json:"comparator,omitempty"`
						EvaluationWindow *struct {
							EvaluationEndTime   *int64 `tfsdk:"evaluation_end_time" json:"evaluationEndTime,omitempty"`
							EvaluationStartTime *int64 `tfsdk:"evaluation_start_time" json:"evaluationStartTime,omitempty"`
						} `tfsdk:"evaluation_window" json:"evaluationWindow,omitempty"`
						InsecureSkipVerify *bool   `tfsdk:"insecure_skip_verify" json:"insecureSkipVerify,omitempty"`
						PlatformEndpoint   *string `tfsdk:"platform_endpoint" json:"platformEndpoint,omitempty"`
						SloIdentifier      *string `tfsdk:"slo_identifier" json:"sloIdentifier,omitempty"`
						SloSourceMetadata  *struct {
							ApiTokenSecret *string `tfsdk:"api_token_secret" json:"apiTokenSecret,omitempty"`
							Scope          *struct {
								AccountIdentifier *string `tfsdk:"account_identifier" json:"accountIdentifier,omitempty"`
								OrgIdentifier     *string `tfsdk:"org_identifier" json:"orgIdentifier,omitempty"`
								ProjectIdentifier *string `tfsdk:"project_identifier" json:"projectIdentifier,omitempty"`
							} `tfsdk:"scope" json:"scope,omitempty"`
						} `tfsdk:"slo_source_metadata" json:"sloSourceMetadata,omitempty"`
					} `tfsdk:"slo_probe_inputs" json:"sloProbe/inputs,omitempty"`
					Type *string `tfsdk:"type" json:"type,omitempty"`
				} `tfsdk:"probe" json:"probe,omitempty"`
			} `tfsdk:"spec" json:"spec,omitempty"`
		} `tfsdk:"experiments" json:"experiments,omitempty"`
		JobCleanUpPolicy *string `tfsdk:"job_clean_up_policy" json:"jobCleanUpPolicy,omitempty"`
		Selectors        *struct {
			Pods *[]struct {
				Names     *string `tfsdk:"names" json:"names,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"pods" json:"pods,omitempty"`
			Workloads *[]struct {
				Kind      *string `tfsdk:"kind" json:"kind,omitempty"`
				Labels    *string `tfsdk:"labels" json:"labels,omitempty"`
				Names     *string `tfsdk:"names" json:"names,omitempty"`
				Namespace *string `tfsdk:"namespace" json:"namespace,omitempty"`
			} `tfsdk:"workloads" json:"workloads,omitempty"`
		} `tfsdk:"selectors" json:"selectors,omitempty"`
		TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" json:"terminationGracePeriodSeconds,omitempty"`
	} `tfsdk:"spec" json:"spec,omitempty"`
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = request.ProviderTypeName + "_litmuschaos_io_chaos_engine_v1alpha1"
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
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

			"force_conflicts": schema.BoolAttribute{
				Description:         "If 'true', server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "If `true`, server-side apply will force the changes against conflicts. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"field_manager": schema.BoolAttribute{
				Description:         "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				MarkdownDescription: "The name of the manager used to track field ownership. If not specified uses the value from the provider configuration.",
				Required:            false,
				Optional:            true,
				Computed:            true,
			},

			"wait_for": schema.ListNestedAttribute{
				Description:         "Wait for specific conditions after create/update of resources.",
				MarkdownDescription: "Wait for specific conditions after create/update of resources.",
				Required:            false,
				Optional:            true,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"jsonpath": schema.StringAttribute{
							Description:         "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							MarkdownDescription: "Relaxed JSONPath expression to use. See https://pkg.go.dev/k8s.io/kubectl/pkg/cmd/get#RelaxedJSONPathExpression for details.",
							Required:            true,
							Optional:            false,
							Computed:            false,
						},
						"value": schema.StringAttribute{
							Description:         "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							MarkdownDescription: "The value to wait for. If not specified, waiting will complete as soon as JSONPath expression exists and has any non-empty value.",
							Required:            false,
							Optional:            true,
							Computed:            true,
						},
						"timeout": schema.StringAttribute{
							Description:         "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							MarkdownDescription: "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.",
							Required:            false,
							Optional:            true,
							Computed:            true,
							Default:             stringdefault.StaticString("30s"),
						},
					},
				},
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
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
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},

					"labels": schema.MapAttribute{
						Description:         "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						MarkdownDescription: "Keys and values that can be used to organize and categorize objects. See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/ for more details.",
						ElementType:         types.StringType,
						Required:            false,
						Optional:            true,
						Computed:            true,
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
						Computed:            true,
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
					"appinfo": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"appkind": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(^$|deployment|statefulset|daemonset|deploymentconfig|rollout)$`), ""),
								},
							},

							"applabel": schema.StringAttribute{
								Description:         "",
								MarkdownDescription: "",
								Required:            false,
								Optional:            true,
								Computed:            false,
							},

							"appns": schema.StringAttribute{
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

					"auxiliary_app_info": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"chaos_service_account": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"components": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"runner": schema.SingleNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								Attributes: map[string]schema.Attribute{
									"image": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"runner_annotations": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"runner_labels": schema.MapAttribute{
										Description:         "",
										MarkdownDescription: "",
										ElementType:         types.StringType,
										Required:            false,
										Optional:            true,
										Computed:            false,
									},

									"tolerations": schema.ListNestedAttribute{
										Description:         "Pod's tolerations.",
										MarkdownDescription: "Pod's tolerations.",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"effect": schema.StringAttribute{
													Description:         "Effect to match. Empty means all effects.",
													MarkdownDescription: "Effect to match. Empty means all effects.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"key": schema.StringAttribute{
													Description:         "Taint key the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists.",
													MarkdownDescription: "Taint key the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"operator": schema.StringAttribute{
													Description:         "Operators are Exists or Equal. Defaults to Equal.",
													MarkdownDescription: "Operators are Exists or Equal. Defaults to Equal.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"toleration_seconds": schema.Int64Attribute{
													Description:         "Period of time the toleration tolerates the taint.",
													MarkdownDescription: "Period of time the toleration tolerates the taint.",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"value": schema.StringAttribute{
													Description:         "If the operator is Exists, the value should be empty, otherwise just a regular string.",
													MarkdownDescription: "If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

									"type": schema.StringAttribute{
										Description:         "",
										MarkdownDescription: "",
										Required:            false,
										Optional:            true,
										Computed:            false,
										Validators: []validator.String{
											stringvalidator.RegexMatches(regexp.MustCompile(`^(go)$`), ""),
										},
									},
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"sidecar": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"env": schema.ListNestedAttribute{
											Description:         "ENV contains ENV passed to the sidecar container",
											MarkdownDescription: "ENV contains ENV passed to the sidecar container",
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
														Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
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

										"env_from": schema.ListNestedAttribute{
											Description:         "EnvFrom for the sidecar container",
											MarkdownDescription: "EnvFrom for the sidecar container",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"config_map_ref": schema.SingleNestedAttribute{
														Description:         "The ConfigMap to select from",
														MarkdownDescription: "The ConfigMap to select from",
														Attributes: map[string]schema.Attribute{
															"name": schema.StringAttribute{
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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
																Description:         "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
																MarkdownDescription: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names TODO: Add other useful fields. apiVersion, kind, uid?",
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

										"image": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"image_pull_policy": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"secrets": schema.ListNestedAttribute{
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

					"default_health_check": schema.BoolAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
					},

					"engine_state": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(active|stop)$`), ""),
						},
					},

					"experiments": schema.ListNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"name": schema.StringAttribute{
									Description:         "",
									MarkdownDescription: "",
									Required:            false,
									Optional:            true,
									Computed:            false,
								},

								"spec": schema.SingleNestedAttribute{
									Description:         "",
									MarkdownDescription: "",
									Attributes: map[string]schema.Attribute{
										"components": schema.SingleNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											Attributes: map[string]schema.Attribute{
												"config_maps": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"mount_path": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"env": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
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
																Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
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
																		Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.",
																		MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.",
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

												"experiment_annotations": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"experiment_image": schema.StringAttribute{
													Description:         "",
													MarkdownDescription: "",
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"node_selector": schema.MapAttribute{
													Description:         "",
													MarkdownDescription: "",
													ElementType:         types.StringType,
													Required:            false,
													Optional:            true,
													Computed:            false,
												},

												"secrets": schema.ListNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"mount_path": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
														},
													},
													Required: false,
													Optional: true,
													Computed: false,
												},

												"status_check_timeouts": schema.SingleNestedAttribute{
													Description:         "",
													MarkdownDescription: "",
													Attributes: map[string]schema.Attribute{
														"delay": schema.Int64Attribute{
															Description:         "",
															MarkdownDescription: "",
															Required:            false,
															Optional:            true,
															Computed:            false,
														},

														"timeout": schema.Int64Attribute{
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

												"tolerations": schema.ListNestedAttribute{
													Description:         "Pod's tolerations.",
													MarkdownDescription: "Pod's tolerations.",
													NestedObject: schema.NestedAttributeObject{
														Attributes: map[string]schema.Attribute{
															"effect": schema.StringAttribute{
																Description:         "Effect to match. Empty means all effects.",
																MarkdownDescription: "Effect to match. Empty means all effects.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"key": schema.StringAttribute{
																Description:         "Taint key the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists.",
																MarkdownDescription: "Taint key the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"operator": schema.StringAttribute{
																Description:         "Operators are Exists or Equal. Defaults to Equal.",
																MarkdownDescription: "Operators are Exists or Equal. Defaults to Equal.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"toleration_seconds": schema.Int64Attribute{
																Description:         "Period of time the toleration tolerates the taint.",
																MarkdownDescription: "Period of time the toleration tolerates the taint.",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"value": schema.StringAttribute{
																Description:         "If the operator is Exists, the value should be empty, otherwise just a regular string.",
																MarkdownDescription: "If the operator is Exists, the value should be empty, otherwise just a regular string.",
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

										"probe": schema.ListNestedAttribute{
											Description:         "",
											MarkdownDescription: "",
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{
													"cmd_probe_inputs": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"command": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																},
															},

															"comparator": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"criteria": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																		Validators: []validator.String{
																			stringvalidator.LengthAtLeast(1),
																			stringvalidator.RegexMatches(regexp.MustCompile(`^(int|float|string)$`), ""),
																		},
																	},

																	"value": schema.StringAttribute{
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

															"source": schema.SingleNestedAttribute{
																Description:         "The external pod where we have to run the probe commands. It will run the commands inside the experiment pod itself(inline mode) if source contains a nil value",
																MarkdownDescription: "The external pod where we have to run the probe commands. It will run the commands inside the experiment pod itself(inline mode) if source contains a nil value",
																Attributes: map[string]schema.Attribute{
																	"annotations": schema.MapAttribute{
																		Description:         "Annotations for the source pod",
																		MarkdownDescription: "Annotations for the source pod",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"args": schema.ListAttribute{
																		Description:         "Args for the source pod",
																		MarkdownDescription: "Args for the source pod",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"command": schema.ListAttribute{
																		Description:         "Command for the source pod",
																		MarkdownDescription: "Command for the source pod",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"env": schema.ListNestedAttribute{
																		Description:         "ENVList contains ENV passed to the source pod",
																		MarkdownDescription: "ENVList contains ENV passed to the source pod",
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
																					Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																					MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
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
																							Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.",
																							MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.",
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

																	"host_network": schema.BoolAttribute{
																		Description:         "HostNetwork define the hostNetwork of the external pod it supports boolean values and default value is false",
																		MarkdownDescription: "HostNetwork define the hostNetwork of the external pod it supports boolean values and default value is false",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"image": schema.StringAttribute{
																		Description:         "Image for the source pod",
																		MarkdownDescription: "Image for the source pod",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"image_pull_policy": schema.StringAttribute{
																		Description:         "ImagePullPolicy for the source pod",
																		MarkdownDescription: "ImagePullPolicy for the source pod",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"image_pull_secrets": schema.ListNestedAttribute{
																		Description:         "ImagePullSecrets for source pod",
																		MarkdownDescription: "ImagePullSecrets for source pod",
																		NestedObject: schema.NestedAttributeObject{
																			Attributes: map[string]schema.Attribute{
																				"name": schema.StringAttribute{
																					Description:         "Name of the referent",
																					MarkdownDescription: "Name of the referent",
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

																	"inherit_inputs": schema.BoolAttribute{
																		Description:         "InheritInputs define to inherit experiment details in probe pod it supports boolean values and default value is false.",
																		MarkdownDescription: "InheritInputs define to inherit experiment details in probe pod it supports boolean values and default value is false.",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"labels": schema.MapAttribute{
																		Description:         "Labels for the source pod",
																		MarkdownDescription: "Labels for the source pod",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"node_selector": schema.MapAttribute{
																		Description:         "NodeSelector for the source pod",
																		MarkdownDescription: "NodeSelector for the source pod",
																		ElementType:         types.StringType,
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"privileged": schema.BoolAttribute{
																		Description:         "Privileged for the source pod",
																		MarkdownDescription: "Privileged for the source pod",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"volume_mount": schema.ListNestedAttribute{
																		Description:         "VolumesMount for the source pod",
																		MarkdownDescription: "VolumesMount for the source pod",
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
																					Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive. This field is beta in 1.15.",
																					MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive. This field is beta in 1.15.",
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
																		Description:         "Volumes for the source pod",
																		MarkdownDescription: "Volumes for the source pod",
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
																							Optional:            true,
																							Computed:            false,
																						},

																						"partition": schema.Int64Attribute{
																							Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																							MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																							MarkdownDescription: "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_id": schema.StringAttribute{
																							Description:         "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																							MarkdownDescription: "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
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
																							Description:         "Host Caching mode: None, Read Only, Read Write.",
																							MarkdownDescription: "Host Caching mode: None, Read Only, Read Write.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"disk_name": schema.StringAttribute{
																							Description:         "The Name of the data disk in the blob storage",
																							MarkdownDescription: "The Name of the data disk in the blob storage",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"disk_uri": schema.StringAttribute{
																							Description:         "The URI the data disk in the blob storage",
																							MarkdownDescription: "The URI the data disk in the blob storage",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"kind": schema.StringAttribute{
																							Description:         "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																							MarkdownDescription: "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
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
																							Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_name": schema.StringAttribute{
																							Description:         "the name of secret that contains Azure Storage Account Name and Key",
																							MarkdownDescription: "the name of secret that contains Azure Storage Account Name and Key",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"share_name": schema.StringAttribute{
																							Description:         "Share Name",
																							MarkdownDescription: "Share Name",
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
																					Description:         "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
																					MarkdownDescription: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
																					Attributes: map[string]schema.Attribute{
																						"monitors": schema.ListAttribute{
																							Description:         "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																							MarkdownDescription: "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																							ElementType:         types.StringType,
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
																							Description:         "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																							MarkdownDescription: "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																							MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_file": schema.StringAttribute{
																							Description:         "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																							MarkdownDescription: "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.SingleNestedAttribute{
																							Description:         "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																							MarkdownDescription: "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
																							Description:         "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																							MarkdownDescription: "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
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
																					Description:         "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																					MarkdownDescription: "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																					Attributes: map[string]schema.Attribute{
																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																							MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																							MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.SingleNestedAttribute{
																							Description:         "Optional: points to a secret object containing parameters used to connect to OpenStack.",
																							MarkdownDescription: "Optional: points to a secret object containing parameters used to connect to OpenStack.",
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
																							Description:         "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																							MarkdownDescription: "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
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
																					Description:         "ConfigMap represents a configMap that should populate this volume",
																					MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",
																					Attributes: map[string]schema.Attribute{
																						"default_mode": schema.Int64Attribute{
																							Description:         "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"items": schema.ListNestedAttribute{
																							Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																							MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "The key to project.",
																										MarkdownDescription: "The key to project.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"mode": schema.Int64Attribute{
																										Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																										MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"path": schema.StringAttribute{
																										Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																										MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																							Description:         "Specify whether the ConfigMap or its keys must be defined",
																							MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",
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
																					Description:         "CSI (Container Storage Interface) represents storage that is handled by an external CSI driver (Alpha feature).",
																					MarkdownDescription: "CSI (Container Storage Interface) represents storage that is handled by an external CSI driver (Alpha feature).",
																					Attributes: map[string]schema.Attribute{
																						"driver": schema.StringAttribute{
																							Description:         "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																							MarkdownDescription: "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																							MarkdownDescription: "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"node_publish_secret_ref": schema.SingleNestedAttribute{
																							Description:         "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																							MarkdownDescription: "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
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
																							Description:         "Specifies a read-only configuration for the volume. Defaults to false (read/write).",
																							MarkdownDescription: "Specifies a read-only configuration for the volume. Defaults to false (read/write).",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_attributes": schema.MapAttribute{
																							Description:         "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
																							MarkdownDescription: "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
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
																					Description:         "DownwardAPI represents downward API about the pod that should populate this volume",
																					MarkdownDescription: "DownwardAPI represents downward API about the pod that should populate this volume",
																					Attributes: map[string]schema.Attribute{
																						"default_mode": schema.Int64Attribute{
																							Description:         "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
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
																										Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																										MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
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
																					Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																					MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																					Attributes: map[string]schema.Attribute{
																						"medium": schema.StringAttribute{
																							Description:         "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																							MarkdownDescription: "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"size_limit": schema.StringAttribute{
																							Description:         "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
																							MarkdownDescription: "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"fc": schema.SingleNestedAttribute{
																					Description:         "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
																					MarkdownDescription: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
																					Attributes: map[string]schema.Attribute{
																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																							MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"lun": schema.Int64Attribute{
																							Description:         "Optional: FC target lun number",
																							MarkdownDescription: "Optional: FC target lun number",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"target_ww_ns": schema.ListAttribute{
																							Description:         "Optional: FC target worldwide names (WWNs)",
																							MarkdownDescription: "Optional: FC target worldwide names (WWNs)",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"wwids": schema.ListAttribute{
																							Description:         "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
																							MarkdownDescription: "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
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
																					Description:         "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
																					MarkdownDescription: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
																					Attributes: map[string]schema.Attribute{
																						"driver": schema.StringAttribute{
																							Description:         "Driver is the name of the driver to use for this volume.",
																							MarkdownDescription: "Driver is the name of the driver to use for this volume.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																							MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"options": schema.MapAttribute{
																							Description:         "Optional: Extra command options if any.",
																							MarkdownDescription: "Optional: Extra command options if any.",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.SingleNestedAttribute{
																							Description:         "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																							MarkdownDescription: "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
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
																					Description:         "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
																					MarkdownDescription: "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
																					Attributes: map[string]schema.Attribute{
																						"dataset_name": schema.StringAttribute{
																							Description:         "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																							MarkdownDescription: "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"dataset_uuid": schema.StringAttribute{
																							Description:         "UUID of the dataset. This is unique identifier of a Flocker dataset",
																							MarkdownDescription: "UUID of the dataset. This is unique identifier of a Flocker dataset",
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
																					Description:         "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																					MarkdownDescription: "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																					Attributes: map[string]schema.Attribute{
																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																							MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"partition": schema.Int64Attribute{
																							Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																							MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"pd_name": schema.StringAttribute{
																							Description:         "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																							MarkdownDescription: "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																							MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
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
																					Description:         "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
																					MarkdownDescription: "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
																					Attributes: map[string]schema.Attribute{
																						"directory": schema.StringAttribute{
																							Description:         "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																							MarkdownDescription: "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"repository": schema.StringAttribute{
																							Description:         "Repository URL",
																							MarkdownDescription: "Repository URL",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"revision": schema.StringAttribute{
																							Description:         "Commit hash for the specified revision.",
																							MarkdownDescription: "Commit hash for the specified revision.",
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
																					Description:         "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
																					MarkdownDescription: "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
																					Attributes: map[string]schema.Attribute{
																						"endpoints": schema.StringAttribute{
																							Description:         "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																							MarkdownDescription: "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"path": schema.StringAttribute{
																							Description:         "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																							MarkdownDescription: "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																							MarkdownDescription: "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
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
																					Description:         "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
																					MarkdownDescription: "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
																					Attributes: map[string]schema.Attribute{
																						"path": schema.StringAttribute{
																							Description:         "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																							MarkdownDescription: "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"type": schema.StringAttribute{
																							Description:         "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																							MarkdownDescription: "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
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
																					Description:         "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
																					MarkdownDescription: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
																					Attributes: map[string]schema.Attribute{
																						"chap_auth_discovery": schema.BoolAttribute{
																							Description:         "whether support iSCSI Discovery CHAP authentication",
																							MarkdownDescription: "whether support iSCSI Discovery CHAP authentication",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"chap_auth_session": schema.BoolAttribute{
																							Description:         "whether support iSCSI Session CHAP authentication",
																							MarkdownDescription: "whether support iSCSI Session CHAP authentication",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																							MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"initiator_name": schema.StringAttribute{
																							Description:         "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																							MarkdownDescription: "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"iqn": schema.StringAttribute{
																							Description:         "Target iSCSI Qualified Name.",
																							MarkdownDescription: "Target iSCSI Qualified Name.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"iscsi_interface": schema.StringAttribute{
																							Description:         "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																							MarkdownDescription: "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"lun": schema.Int64Attribute{
																							Description:         "iSCSI Target Lun number.",
																							MarkdownDescription: "iSCSI Target Lun number.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"portals": schema.ListAttribute{
																							Description:         "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																							MarkdownDescription: "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																							ElementType:         types.StringType,
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																							MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.SingleNestedAttribute{
																							Description:         "CHAP Secret for iSCSI target and initiator authentication",
																							MarkdownDescription: "CHAP Secret for iSCSI target and initiator authentication",
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
																							Description:         "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																							MarkdownDescription: "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
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
																					Description:         "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																					MarkdownDescription: "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																					Required:            true,
																					Optional:            false,
																					Computed:            false,
																				},

																				"nfs": schema.SingleNestedAttribute{
																					Description:         "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																					MarkdownDescription: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																					Attributes: map[string]schema.Attribute{
																						"path": schema.StringAttribute{
																							Description:         "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																							MarkdownDescription: "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																							MarkdownDescription: "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"server": schema.StringAttribute{
																							Description:         "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																							MarkdownDescription: "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
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
																					Description:         "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																					MarkdownDescription: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																					Attributes: map[string]schema.Attribute{
																						"claim_name": schema.StringAttribute{
																							Description:         "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																							MarkdownDescription: "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Will force the ReadOnly setting in VolumeMounts. Default false.",
																							MarkdownDescription: "Will force the ReadOnly setting in VolumeMounts. Default false.",
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
																					Description:         "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
																					MarkdownDescription: "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
																					Attributes: map[string]schema.Attribute{
																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"pd_id": schema.StringAttribute{
																							Description:         "ID that identifies Photon Controller persistent disk",
																							MarkdownDescription: "ID that identifies Photon Controller persistent disk",
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
																					Description:         "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
																					MarkdownDescription: "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
																					Attributes: map[string]schema.Attribute{
																						"fs_type": schema.StringAttribute{
																							Description:         "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							MarkdownDescription: "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_id": schema.StringAttribute{
																							Description:         "VolumeID uniquely identifies a Portworx volume",
																							MarkdownDescription: "VolumeID uniquely identifies a Portworx volume",
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
																					Description:         "Items for all in one resources secrets, configmaps, and downward API",
																					MarkdownDescription: "Items for all in one resources secrets, configmaps, and downward API",
																					Attributes: map[string]schema.Attribute{
																						"default_mode": schema.Int64Attribute{
																							Description:         "Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
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
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"mode": schema.Int64Attribute{
																															Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																															MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																															Required:            false,
																															Optional:            true,
																															Computed:            false,
																														},

																														"path": schema.StringAttribute{
																															Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																															MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																												Description:         "Specify whether the ConfigMap or its keys must be defined",
																												MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",
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
																															Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																															MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
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
																															Required:            true,
																															Optional:            false,
																															Computed:            false,
																														},

																														"mode": schema.Int64Attribute{
																															Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																															MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																															Required:            false,
																															Optional:            true,
																															Computed:            false,
																														},

																														"path": schema.StringAttribute{
																															Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																															MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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

																									"service_account_token": schema.SingleNestedAttribute{
																										Description:         "information about the serviceAccountToken data to project",
																										MarkdownDescription: "information about the serviceAccountToken data to project",
																										Attributes: map[string]schema.Attribute{
																											"audience": schema.StringAttribute{
																												Description:         "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																												MarkdownDescription: "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"expiration_seconds": schema.Int64Attribute{
																												Description:         "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																												MarkdownDescription: "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																												Required:            false,
																												Optional:            true,
																												Computed:            false,
																											},

																											"path": schema.StringAttribute{
																												Description:         "Path is the path relative to the mount point of the file to project the token into.",
																												MarkdownDescription: "Path is the path relative to the mount point of the file to project the token into.",
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
																							Required: true,
																							Optional: false,
																							Computed: false,
																						},
																					},
																					Required: false,
																					Optional: true,
																					Computed: false,
																				},

																				"quobyte": schema.SingleNestedAttribute{
																					Description:         "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
																					MarkdownDescription: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
																					Attributes: map[string]schema.Attribute{
																						"group": schema.StringAttribute{
																							Description:         "Group to map volume access to Default is no group",
																							MarkdownDescription: "Group to map volume access to Default is no group",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																							MarkdownDescription: "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"registry": schema.StringAttribute{
																							Description:         "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																							MarkdownDescription: "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"tenant": schema.StringAttribute{
																							Description:         "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																							MarkdownDescription: "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"user": schema.StringAttribute{
																							Description:         "User to map volume access to Defaults to serivceaccount user",
																							MarkdownDescription: "User to map volume access to Defaults to serivceaccount user",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume": schema.StringAttribute{
																							Description:         "Volume is a string that references an already created Quobyte volume by name.",
																							MarkdownDescription: "Volume is a string that references an already created Quobyte volume by name.",
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
																					Description:         "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
																					MarkdownDescription: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
																					Attributes: map[string]schema.Attribute{
																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																							MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"image": schema.StringAttribute{
																							Description:         "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							MarkdownDescription: "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"keyring": schema.StringAttribute{
																							Description:         "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							MarkdownDescription: "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"monitors": schema.ListAttribute{
																							Description:         "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							MarkdownDescription: "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							ElementType:         types.StringType,
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"pool": schema.StringAttribute{
																							Description:         "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							MarkdownDescription: "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.SingleNestedAttribute{
																							Description:         "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							MarkdownDescription: "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
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
																							Description:         "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																							MarkdownDescription: "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
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
																					Description:         "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
																					MarkdownDescription: "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
																					Attributes: map[string]schema.Attribute{
																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																							MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"gateway": schema.StringAttribute{
																							Description:         "The host address of the ScaleIO API Gateway.",
																							MarkdownDescription: "The host address of the ScaleIO API Gateway.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"protection_domain": schema.StringAttribute{
																							Description:         "The name of the ScaleIO Protection Domain for the configured storage.",
																							MarkdownDescription: "The name of the ScaleIO Protection Domain for the configured storage.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.SingleNestedAttribute{
																							Description:         "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																							MarkdownDescription: "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
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
																							Description:         "Flag to enable/disable SSL communication with Gateway, default false",
																							MarkdownDescription: "Flag to enable/disable SSL communication with Gateway, default false",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"storage_mode": schema.StringAttribute{
																							Description:         "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																							MarkdownDescription: "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"storage_pool": schema.StringAttribute{
																							Description:         "The ScaleIO Storage Pool associated with the protection domain.",
																							MarkdownDescription: "The ScaleIO Storage Pool associated with the protection domain.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"system": schema.StringAttribute{
																							Description:         "The name of the storage system as configured in ScaleIO.",
																							MarkdownDescription: "The name of the storage system as configured in ScaleIO.",
																							Required:            true,
																							Optional:            false,
																							Computed:            false,
																						},

																						"volume_name": schema.StringAttribute{
																							Description:         "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
																							MarkdownDescription: "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
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
																					Description:         "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																					MarkdownDescription: "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																					Attributes: map[string]schema.Attribute{
																						"default_mode": schema.Int64Attribute{
																							Description:         "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"items": schema.ListNestedAttribute{
																							Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																							MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																							NestedObject: schema.NestedAttributeObject{
																								Attributes: map[string]schema.Attribute{
																									"key": schema.StringAttribute{
																										Description:         "The key to project.",
																										MarkdownDescription: "The key to project.",
																										Required:            true,
																										Optional:            false,
																										Computed:            false,
																									},

																									"mode": schema.Int64Attribute{
																										Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																										MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																										Required:            false,
																										Optional:            true,
																										Computed:            false,
																									},

																									"path": schema.StringAttribute{
																										Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																										MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
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
																							Description:         "Specify whether the Secret or its keys must be defined",
																							MarkdownDescription: "Specify whether the Secret or its keys must be defined",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_name": schema.StringAttribute{
																							Description:         "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																							MarkdownDescription: "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
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
																					Description:         "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
																					MarkdownDescription: "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
																					Attributes: map[string]schema.Attribute{
																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"read_only": schema.BoolAttribute{
																							Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"secret_ref": schema.SingleNestedAttribute{
																							Description:         "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
																							MarkdownDescription: "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
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
																							Description:         "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
																							MarkdownDescription: "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_namespace": schema.StringAttribute{
																							Description:         "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																							MarkdownDescription: "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
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
																					Description:         "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
																					MarkdownDescription: "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
																					Attributes: map[string]schema.Attribute{
																						"fs_type": schema.StringAttribute{
																							Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"storage_policy_id": schema.StringAttribute{
																							Description:         "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																							MarkdownDescription: "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"storage_policy_name": schema.StringAttribute{
																							Description:         "Storage Policy Based Management (SPBM) profile name.",
																							MarkdownDescription: "Storage Policy Based Management (SPBM) profile name.",
																							Required:            false,
																							Optional:            true,
																							Computed:            false,
																						},

																						"volume_path": schema.StringAttribute{
																							Description:         "Path that identifies vSphere volume vmdk",
																							MarkdownDescription: "Path that identifies vSphere volume vmdk",
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
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"data": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            false,
														Optional:            true,
														Computed:            false,
													},

													"http_probe_inputs": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"method": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"get": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"criteria": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},

																			"response_code": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																		Validators: []UNKNOWN{
																			schemavalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("post")),
																		},
																	},

																	"post": schema.SingleNestedAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Attributes: map[string]schema.Attribute{
																			"body": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"body_path": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																			},

																			"content_type": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            false,
																				Optional:            true,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},

																			"criteria": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																				Validators: []validator.String{
																					stringvalidator.LengthAtLeast(1),
																				},
																			},

																			"response_code": schema.StringAttribute{
																				Description:         "",
																				MarkdownDescription: "",
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
																		Validators: []UNKNOWN{
																			schemavalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("get")),
																		},
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"url": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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

													"k8s_probe_inputs": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"field_selector": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"group": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"label_selector": schema.StringAttribute{
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

															"operation": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
																Validators: []validator.String{
																	stringvalidator.LengthAtLeast(1),
																	stringvalidator.RegexMatches(regexp.MustCompile(`^(present|absent|create|delete)$`), ""),
																},
															},

															"resource": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"resource_names": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
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
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"mode": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.RegexMatches(regexp.MustCompile(`^(SOT|EOT|Edge|Continuous|OnChaos)$`), ""),
														},
													},

													"name": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
													},

													"prom_probe_inputs": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"comparator": schema.SingleNestedAttribute{
																Description:         "",
																MarkdownDescription: "",
																Attributes: map[string]schema.Attribute{
																	"criteria": schema.StringAttribute{
																		Description:         "",
																		MarkdownDescription: "",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
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

															"endpoint": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"query": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"query_path": schema.StringAttribute{
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

													"run_properties": schema.SingleNestedAttribute{
														Description:         "",
														MarkdownDescription: "",
														Attributes: map[string]schema.Attribute{
															"attempt": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"evaluation_timeout": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"initial_delay": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"initial_delay_seconds": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"interval": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"probe_polling_interval": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"probe_timeout": schema.StringAttribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"retry": schema.Int64Attribute{
																Description:         "",
																MarkdownDescription: "",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"stop_on_failure": schema.BoolAttribute{
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

													"slo_probe_inputs": schema.SingleNestedAttribute{
														Description:         "inputs needed for the SLO probe",
														MarkdownDescription: "inputs needed for the SLO probe",
														Attributes: map[string]schema.Attribute{
															"comparator": schema.SingleNestedAttribute{
																Description:         "Comparator check for the correctness of the probe output",
																MarkdownDescription: "Comparator check for the correctness of the probe output",
																Attributes: map[string]schema.Attribute{
																	"criteria": schema.StringAttribute{
																		Description:         "Criteria for matching data it supports >=, <=, ==, >, <, != for int and float it supports equal, notEqual, contains for string",
																		MarkdownDescription: "Criteria for matching data it supports >=, <=, ==, >, <, != for int and float it supports equal, notEqual, contains for string",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"type": schema.StringAttribute{
																		Description:         "Type of data it can be int, float, string",
																		MarkdownDescription: "Type of data it can be int, float, string",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"value": schema.StringAttribute{
																		Description:         "Value contains relative value for criteria",
																		MarkdownDescription: "Value contains relative value for criteria",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},
																},
																Required: true,
																Optional: false,
																Computed: false,
															},

															"evaluation_window": schema.SingleNestedAttribute{
																Description:         "EvaluationWindow is the time period for which the metrics will be evaluated",
																MarkdownDescription: "EvaluationWindow is the time period for which the metrics will be evaluated",
																Attributes: map[string]schema.Attribute{
																	"evaluation_end_time": schema.Int64Attribute{
																		Description:         "End time of evaluation",
																		MarkdownDescription: "End time of evaluation",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},

																	"evaluation_start_time": schema.Int64Attribute{
																		Description:         "Start time of evaluation",
																		MarkdownDescription: "Start time of evaluation",
																		Required:            false,
																		Optional:            true,
																		Computed:            false,
																	},
																},
																Required: false,
																Optional: true,
																Computed: false,
															},

															"insecure_skip_verify": schema.BoolAttribute{
																Description:         "InsecureSkipVerify flag to skip certificate checks",
																MarkdownDescription: "InsecureSkipVerify flag to skip certificate checks",
																Required:            false,
																Optional:            true,
																Computed:            false,
															},

															"platform_endpoint": schema.StringAttribute{
																Description:         "PlatformEndpoint for the monitoring service endpoint",
																MarkdownDescription: "PlatformEndpoint for the monitoring service endpoint",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"slo_identifier": schema.StringAttribute{
																Description:         "SLOIdentifier for fetching the details of the SLO",
																MarkdownDescription: "SLOIdentifier for fetching the details of the SLO",
																Required:            true,
																Optional:            false,
																Computed:            false,
															},

															"slo_source_metadata": schema.SingleNestedAttribute{
																Description:         "SLOSourceMetadata consists of required metadata details to fetch metric data",
																MarkdownDescription: "SLOSourceMetadata consists of required metadata details to fetch metric data",
																Attributes: map[string]schema.Attribute{
																	"api_token_secret": schema.StringAttribute{
																		Description:         "APITokenSecret for authenticating with the platform service",
																		MarkdownDescription: "APITokenSecret for authenticating with the platform service",
																		Required:            true,
																		Optional:            false,
																		Computed:            false,
																	},

																	"scope": schema.SingleNestedAttribute{
																		Description:         "Scope required for fetching details",
																		MarkdownDescription: "Scope required for fetching details",
																		Attributes: map[string]schema.Attribute{
																			"account_identifier": schema.StringAttribute{
																				Description:         "AccountIdentifier for account ID",
																				MarkdownDescription: "AccountIdentifier for account ID",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"org_identifier": schema.StringAttribute{
																				Description:         "OrgIdentifier for organization ID",
																				MarkdownDescription: "OrgIdentifier for organization ID",
																				Required:            true,
																				Optional:            false,
																				Computed:            false,
																			},

																			"project_identifier": schema.StringAttribute{
																				Description:         "ProjectIdentifier for project ID",
																				MarkdownDescription: "ProjectIdentifier for project ID",
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
																Required: true,
																Optional: false,
																Computed: false,
															},
														},
														Required: false,
														Optional: true,
														Computed: false,
													},

													"type": schema.StringAttribute{
														Description:         "",
														MarkdownDescription: "",
														Required:            true,
														Optional:            false,
														Computed:            false,
														Validators: []validator.String{
															stringvalidator.LengthAtLeast(1),
															stringvalidator.RegexMatches(regexp.MustCompile(`^(k8sProbe|httpProbe|cmdProbe|promProbe|sloProbe)$`), ""),
														},
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
						},
						Required: false,
						Optional: true,
						Computed: false,
					},

					"job_clean_up_policy": schema.StringAttribute{
						Description:         "",
						MarkdownDescription: "",
						Required:            false,
						Optional:            true,
						Computed:            false,
						Validators: []validator.String{
							stringvalidator.RegexMatches(regexp.MustCompile(`^(delete|retain)$`), ""),
						},
					},

					"selectors": schema.SingleNestedAttribute{
						Description:         "",
						MarkdownDescription: "",
						Attributes: map[string]schema.Attribute{
							"pods": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"names": schema.StringAttribute{
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
								},
								Required: false,
								Optional: true,
								Computed: false,
							},

							"workloads": schema.ListNestedAttribute{
								Description:         "",
								MarkdownDescription: "",
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"kind": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            true,
											Optional:            false,
											Computed:            false,
											Validators: []validator.String{
												stringvalidator.RegexMatches(regexp.MustCompile(`^(^$|deployment|statefulset|daemonset|deploymentconfig|rollout)$`), ""),
											},
										},

										"labels": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
											Computed:            false,
										},

										"names": schema.StringAttribute{
											Description:         "",
											MarkdownDescription: "",
											Required:            false,
											Optional:            true,
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

					"termination_grace_period_seconds": schema.Int64Attribute{
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
		},
	}
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Configure(_ context.Context, request resource.ConfigureRequest, response *resource.ConfigureResponse) {
	if request.ProviderData == nil {
		return
	}

	if resourceData, ok := request.ProviderData.(*utilities.ResourceData); ok {
		if resourceData.Offline {
			response.Diagnostics.AddError(
				"Provider in Offline Mode",
				"This provider has offline mode enabled and thus cannot connect to a Kubernetes cluster to create resources or read any data. "+
					"Disable offline mode to allow resource creation or remove the resource declaration from your configuration to get rid of this error.",
			)
		} else {
			r.kubernetesClient = resourceData.Client
			r.fieldManager = resourceData.FieldManager
			r.forceConflicts = resourceData.ForceConflicts
		}
	} else {
		response.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *dynamic.DynamicClient, got: %T. Please report this issue to the provider developers.", request.ProviderData),
		)
	}
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_litmuschaos_io_chaos_engine_v1alpha1")

	var model LitmuschaosIoChaosEngineV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ID = types.StringValue(fmt.Sprintf("%s/%s", model.Metadata.Name, model.Metadata.Namespace))
	model.ApiVersion = pointer.String("litmuschaos.io/v1alpha1")
	model.Kind = pointer.String("ChaosEngine")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "litmuschaos.io", Version: "v1alpha1", Resource: "ChaosEngine"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while creating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse LitmuschaosIoChaosEngineV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_litmuschaos_io_chaos_engine_v1alpha1")

	var data LitmuschaosIoChaosEngineV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	getResponse, err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "litmuschaos.io", Version: "v1alpha1", Resource: "ChaosEngine"}).
		Namespace(data.Metadata.Namespace).
		Get(ctx, data.Metadata.Name, meta.GetOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to GET resource",
			"An unexpected error occurred while reading the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"GET Error: "+err.Error(),
		)
		return
	}
	getBytes, err := getResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal GET response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse LitmuschaosIoChaosEngineV1Alpha1ResourceData
	err = json.Unmarshal(getBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal resource",
			"An unexpected error occurred while parsing the resource read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	data.Metadata = readResponse.Metadata
	data.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_litmuschaos_io_chaos_engine_v1alpha1")

	var model LitmuschaosIoChaosEngineV1Alpha1ResourceData
	response.Diagnostics.Append(request.Plan.Get(ctx, &model)...)
	if response.Diagnostics.HasError() {
		return
	}

	model.ApiVersion = pointer.String("litmuschaos.io/v1alpha1")
	model.Kind = pointer.String("ChaosEngine")

	bytes, err := json.Marshal(model)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal resource",
			"An unexpected error occurred while marshalling the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"JSON Error: "+err.Error(),
		)
		return
	}

	forceConflicts := r.forceConflicts
	if !model.ForceConflicts.IsNull() && !model.ForceConflicts.IsUnknown() {
		forceConflicts = model.ForceConflicts.ValueBool()
	}
	fieldManager := r.fieldManager
	if !model.FieldManager.IsNull() && !model.FieldManager.IsUnknown() {
		fieldManager = model.FieldManager.ValueString()
	}
	patchOptions := meta.PatchOptions{
		FieldManager:    fieldManager,
		Force:           pointer.Bool(forceConflicts),
		FieldValidation: "Strict",
	}

	patchResponse, err := r.kubernetesClient.Resource(k8sSchema.GroupVersionResource{Group: "litmuschaos.io", Version: "v1alpha1", Resource: "ChaosEngine"}).
		Namespace(model.Metadata.Namespace).
		Patch(ctx, model.Metadata.Name, k8sTypes.ApplyPatchType, bytes, patchOptions)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to PATCH resource",
			"An unexpected error occurred while updating the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"PATCH Error: "+err.Error(),
		)
		return
	}

	patchBytes, err := patchResponse.MarshalJSON()
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to marshal PATCH response",
			"Please report this issue to the provider developers.\n\n"+
				"Marshal Error: "+err.Error(),
		)
		return
	}

	var readResponse LitmuschaosIoChaosEngineV1Alpha1ResourceData
	err = json.Unmarshal(patchBytes, &readResponse)
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to unmarshal response",
			"An unexpected error occurred while unmarshalling read response. "+
				"Please report this issue to the provider developers.\n\n"+
				"Unmarshal Error: "+err.Error(),
		)
		return
	}

	model.Metadata = readResponse.Metadata
	model.Spec = readResponse.Spec

	response.Diagnostics.Append(response.State.Set(ctx, &model)...)
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_litmuschaos_io_chaos_engine_v1alpha1")

	var data LitmuschaosIoChaosEngineV1Alpha1ResourceData
	response.Diagnostics.Append(request.State.Get(ctx, &data)...)
	if response.Diagnostics.HasError() {
		return
	}

	err := r.kubernetesClient.
		Resource(k8sSchema.GroupVersionResource{Group: "litmuschaos.io", Version: "v1alpha1", Resource: "ChaosEngine"}).
		Namespace(data.Metadata.Namespace).
		Delete(ctx, data.Metadata.Name, meta.DeleteOptions{})
	if err != nil {
		response.Diagnostics.AddError(
			"Unable to DELETE resource",
			"An unexpected error occurred while deleting the resource. "+
				"Please report this issue to the provider developers.\n\n"+
				"DELETE Error: "+err.Error(),
		)
		return
	}
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) ImportState(ctx context.Context, request resource.ImportStateRequest, response *resource.ImportStateResponse) {
	idParts := strings.Split(request.ID, "/")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		response.Diagnostics.AddError(
			"Error importing resource",
			fmt.Sprintf("Expected import identifier with format: 'namespace/name' Got: '%q'", request.ID),
		)
		return
	}

	namespace := idParts[0]
	name := idParts[1]
	tflog.Trace(ctx, "parsed import ID", map[string]interface{}{
		"namespace": namespace,
		"name":      name,
	})
	resource.ImportStatePassthroughID(ctx, path.Root("id"), request, response)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("namespace"), namespace)...)
	response.Diagnostics.Append(response.State.SetAttribute(ctx, path.Root("metadata").AtName("name"), name)...)
}
