/*
* SPDX-FileCopyrightText: The terraform-provider-k8s Authors
* SPDX-License-Identifier: 0BSD
 */

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"

	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/schemavalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"

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

type LitmuschaosIoChaosEngineV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*LitmuschaosIoChaosEngineV1Alpha1Resource)(nil)
)

type LitmuschaosIoChaosEngineV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type LitmuschaosIoChaosEngineV1Alpha1GoModel struct {
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
		AnnotationCheck *string `tfsdk:"annotation_check" yaml:"annotationCheck,omitempty"`

		Appinfo *struct {
			Appkind *string `tfsdk:"appkind" yaml:"appkind,omitempty"`

			Applabel *string `tfsdk:"applabel" yaml:"applabel,omitempty"`

			Appns *string `tfsdk:"appns" yaml:"appns,omitempty"`
		} `tfsdk:"appinfo" yaml:"appinfo,omitempty"`

		AuxiliaryAppInfo *string `tfsdk:"auxiliary_app_info" yaml:"auxiliaryAppInfo,omitempty"`

		ChaosServiceAccount *string `tfsdk:"chaos_service_account" yaml:"chaosServiceAccount,omitempty"`

		Components *struct {
			Runner *struct {
				Image *string `tfsdk:"image" yaml:"image,omitempty"`

				RunnerAnnotations *map[string]string `tfsdk:"runner_annotations" yaml:"runnerAnnotations,omitempty"`

				Tolerations *[]struct {
					Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

					Key *string `tfsdk:"key" yaml:"key,omitempty"`

					Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

					TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

					Value *string `tfsdk:"value" yaml:"value,omitempty"`
				} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`

				Type *string `tfsdk:"type" yaml:"type,omitempty"`
			} `tfsdk:"runner" yaml:"runner,omitempty"`
		} `tfsdk:"components" yaml:"components,omitempty"`

		DefaultAppHealthCheck *string `tfsdk:"default_app_health_check" yaml:"defaultAppHealthCheck,omitempty"`

		EngineState *string `tfsdk:"engine_state" yaml:"engineState,omitempty"`

		Experiments *[]struct {
			Name *string `tfsdk:"name" yaml:"name,omitempty"`

			Spec *struct {
				Components *struct {
					ConfigMaps *[]struct {
						MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"config_maps" yaml:"configMaps,omitempty"`

					Env *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`

						ValueFrom *struct {
							ConfigMapKeyRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

							FieldRef *struct {
								ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

								FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
							} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

							ResourceFieldRef *struct {
								ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

								Divisor utilities.IntOrString `tfsdk:"divisor" yaml:"divisor,omitempty"`

								Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
							} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

							SecretKeyRef *struct {
								Key *string `tfsdk:"key" yaml:"key,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
							} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
						} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
					} `tfsdk:"env" yaml:"env,omitempty"`

					ExperimentAnnotations *map[string]string `tfsdk:"experiment_annotations" yaml:"experimentAnnotations,omitempty"`

					ExperimentImage *string `tfsdk:"experiment_image" yaml:"experimentImage,omitempty"`

					NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

					Secrets *[]struct {
						MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

						Name *string `tfsdk:"name" yaml:"name,omitempty"`
					} `tfsdk:"secrets" yaml:"secrets,omitempty"`

					StatusCheckTimeouts *struct {
						Delay *int64 `tfsdk:"delay" yaml:"delay,omitempty"`

						Timeout *int64 `tfsdk:"timeout" yaml:"timeout,omitempty"`
					} `tfsdk:"status_check_timeouts" yaml:"statusCheckTimeouts,omitempty"`

					Tolerations *[]struct {
						Effect *string `tfsdk:"effect" yaml:"effect,omitempty"`

						Key *string `tfsdk:"key" yaml:"key,omitempty"`

						Operator *string `tfsdk:"operator" yaml:"operator,omitempty"`

						TolerationSeconds *int64 `tfsdk:"toleration_seconds" yaml:"tolerationSeconds,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"tolerations" yaml:"tolerations,omitempty"`
				} `tfsdk:"components" yaml:"components,omitempty"`

				Probe *[]struct {
					CmdProbe_inputs *struct {
						Command *string `tfsdk:"command" yaml:"command,omitempty"`

						Comparator *struct {
							Criteria *string `tfsdk:"criteria" yaml:"criteria,omitempty"`

							Type *string `tfsdk:"type" yaml:"type,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"comparator" yaml:"comparator,omitempty"`

						Source *struct {
							Annotations *map[string]string `tfsdk:"annotations" yaml:"annotations,omitempty"`

							Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

							Command *[]string `tfsdk:"command" yaml:"command,omitempty"`

							Env *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								Value *string `tfsdk:"value" yaml:"value,omitempty"`

								ValueFrom *struct {
									ConfigMapKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
									} `tfsdk:"config_map_key_ref" yaml:"configMapKeyRef,omitempty"`

									FieldRef *struct {
										ApiVersion *string `tfsdk:"api_version" yaml:"apiVersion,omitempty"`

										FieldPath *string `tfsdk:"field_path" yaml:"fieldPath,omitempty"`
									} `tfsdk:"field_ref" yaml:"fieldRef,omitempty"`

									ResourceFieldRef *struct {
										ContainerName *string `tfsdk:"container_name" yaml:"containerName,omitempty"`

										Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

										Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
									} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`

									SecretKeyRef *struct {
										Key *string `tfsdk:"key" yaml:"key,omitempty"`

										Name *string `tfsdk:"name" yaml:"name,omitempty"`

										Optional *bool `tfsdk:"optional" yaml:"optional,omitempty"`
									} `tfsdk:"secret_key_ref" yaml:"secretKeyRef,omitempty"`
								} `tfsdk:"value_from" yaml:"valueFrom,omitempty"`
							} `tfsdk:"env" yaml:"env,omitempty"`

							HostNetwork *bool `tfsdk:"host_network" yaml:"hostNetwork,omitempty"`

							Image *string `tfsdk:"image" yaml:"image,omitempty"`

							ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

							ImagePullSecrets *[]struct {
								Name *string `tfsdk:"name" yaml:"name,omitempty"`
							} `tfsdk:"image_pull_secrets" yaml:"imagePullSecrets,omitempty"`

							InheritInputs *bool `tfsdk:"inherit_inputs" yaml:"inheritInputs,omitempty"`

							Labels *map[string]string `tfsdk:"labels" yaml:"labels,omitempty"`

							NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

							Privileged *bool `tfsdk:"privileged" yaml:"privileged,omitempty"`

							VolumeMount *[]struct {
								MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

								MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

								Name *string `tfsdk:"name" yaml:"name,omitempty"`

								ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

								SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

								SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
							} `tfsdk:"volume_mount" yaml:"volumeMount,omitempty"`

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

											Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

											Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`
										} `tfsdk:"resource_field_ref" yaml:"resourceFieldRef,omitempty"`
									} `tfsdk:"items" yaml:"items,omitempty"`
								} `tfsdk:"downward_api" yaml:"downwardAPI,omitempty"`

								EmptyDir *struct {
									Medium *string `tfsdk:"medium" yaml:"medium,omitempty"`

									SizeLimit *string `tfsdk:"size_limit" yaml:"sizeLimit,omitempty"`
								} `tfsdk:"empty_dir" yaml:"emptyDir,omitempty"`

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

													Divisor *string `tfsdk:"divisor" yaml:"divisor,omitempty"`

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
						} `tfsdk:"source" yaml:"source,omitempty"`
					} `tfsdk:"cmd_probe_inputs" yaml:"cmdProbe/inputs,omitempty"`

					Data *string `tfsdk:"data" yaml:"data,omitempty"`

					HttpProbe_inputs *struct {
						InsecureSkipVerify *bool `tfsdk:"insecure_skip_verify" yaml:"insecureSkipVerify,omitempty"`

						Method *struct {
							Get *struct {
								Criteria *string `tfsdk:"criteria" yaml:"criteria,omitempty"`

								ResponseCode *string `tfsdk:"response_code" yaml:"responseCode,omitempty"`
							} `tfsdk:"get" yaml:"get,omitempty"`

							Post *struct {
								Body *string `tfsdk:"body" yaml:"body,omitempty"`

								BodyPath *string `tfsdk:"body_path" yaml:"bodyPath,omitempty"`

								ContentType *string `tfsdk:"content_type" yaml:"contentType,omitempty"`

								Criteria *string `tfsdk:"criteria" yaml:"criteria,omitempty"`

								ResponseCode *string `tfsdk:"response_code" yaml:"responseCode,omitempty"`
							} `tfsdk:"post" yaml:"post,omitempty"`
						} `tfsdk:"method" yaml:"method,omitempty"`

						Url *string `tfsdk:"url" yaml:"url,omitempty"`
					} `tfsdk:"http_probe_inputs" yaml:"httpProbe/inputs,omitempty"`

					K8sProbe_inputs *struct {
						FieldSelector *string `tfsdk:"field_selector" yaml:"fieldSelector,omitempty"`

						Group *string `tfsdk:"group" yaml:"group,omitempty"`

						LabelSelector *string `tfsdk:"label_selector" yaml:"labelSelector,omitempty"`

						Namespace *string `tfsdk:"namespace" yaml:"namespace,omitempty"`

						Operation *string `tfsdk:"operation" yaml:"operation,omitempty"`

						Resource *string `tfsdk:"resource" yaml:"resource,omitempty"`

						Version *string `tfsdk:"version" yaml:"version,omitempty"`
					} `tfsdk:"k8s_probe_inputs" yaml:"k8sProbe/inputs,omitempty"`

					Mode *string `tfsdk:"mode" yaml:"mode,omitempty"`

					Name *string `tfsdk:"name" yaml:"name,omitempty"`

					PromProbe_inputs *struct {
						Comparator *struct {
							Criteria *string `tfsdk:"criteria" yaml:"criteria,omitempty"`

							Value *string `tfsdk:"value" yaml:"value,omitempty"`
						} `tfsdk:"comparator" yaml:"comparator,omitempty"`

						Endpoint *string `tfsdk:"endpoint" yaml:"endpoint,omitempty"`

						Query *string `tfsdk:"query" yaml:"query,omitempty"`

						QueryPath *string `tfsdk:"query_path" yaml:"queryPath,omitempty"`
					} `tfsdk:"prom_probe_inputs" yaml:"promProbe/inputs,omitempty"`

					RunProperties *struct {
						InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

						Interval *int64 `tfsdk:"interval" yaml:"interval,omitempty"`

						ProbePollingInterval *int64 `tfsdk:"probe_polling_interval" yaml:"probePollingInterval,omitempty"`

						ProbeTimeout *int64 `tfsdk:"probe_timeout" yaml:"probeTimeout,omitempty"`

						Retry *int64 `tfsdk:"retry" yaml:"retry,omitempty"`

						StopOnFailure *bool `tfsdk:"stop_on_failure" yaml:"stopOnFailure,omitempty"`
					} `tfsdk:"run_properties" yaml:"runProperties,omitempty"`

					Type *string `tfsdk:"type" yaml:"type,omitempty"`
				} `tfsdk:"probe" yaml:"probe,omitempty"`
			} `tfsdk:"spec" yaml:"spec,omitempty"`
		} `tfsdk:"experiments" yaml:"experiments,omitempty"`

		JobCleanUpPolicy *string `tfsdk:"job_clean_up_policy" yaml:"jobCleanUpPolicy,omitempty"`

		TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewLitmuschaosIoChaosEngineV1Alpha1Resource() resource.Resource {
	return &LitmuschaosIoChaosEngineV1Alpha1Resource{}
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_litmuschaos_io_chaos_engine_v1alpha1"
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "",
		MarkdownDescription: "",
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
				Description:         "",
				MarkdownDescription: "",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"annotation_check": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^(true|false)$`), ""),
						},
					},

					"appinfo": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"appkind": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.RegexMatches(regexp.MustCompile(`^(^$|deployment|statefulset|daemonset|deploymentconfig|rollout)$`), ""),
								},
							},

							"applabel": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"appns": {
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

					"auxiliary_app_info": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"chaos_service_account": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"components": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"runner": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"image": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"runner_annotations": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.MapType{ElemType: types.StringType},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tolerations": {
										Description:         "Pod's tolerations.",
										MarkdownDescription: "Pod's tolerations.",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"effect": {
												Description:         "Effect to match. Empty means all effects.",
												MarkdownDescription: "Effect to match. Empty means all effects.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"key": {
												Description:         "Taint key the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists.",
												MarkdownDescription: "Taint key the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"operator": {
												Description:         "Operators are Exists or Equal. Defaults to Equal.",
												MarkdownDescription: "Operators are Exists or Equal. Defaults to Equal.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"toleration_seconds": {
												Description:         "Period of time the toleration tolerates the taint.",
												MarkdownDescription: "Period of time the toleration tolerates the taint.",

												Type: types.Int64Type,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"value": {
												Description:         "If the operator is Exists, the value should be empty, otherwise just a regular string.",
												MarkdownDescription: "If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

									"type": {
										Description:         "",
										MarkdownDescription: "",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^(go)$`), ""),
										},
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

					"default_app_health_check": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^(true|false)$`), ""),
						},
					},

					"engine_state": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^(active|stop)$`), ""),
						},
					},

					"experiments": {
						Description:         "",
						MarkdownDescription: "",

						Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

							"name": {
								Description:         "",
								MarkdownDescription: "",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"spec": {
								Description:         "",
								MarkdownDescription: "",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"components": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"config_maps": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"mount_path": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"env": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
														MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
														MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value_from": {
														Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
														MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"config_map_key_ref": {
																Description:         "Selects a key of a ConfigMap.",
																MarkdownDescription: "Selects a key of a ConfigMap.",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key to select.",
																		MarkdownDescription: "The key to select.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
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
																		Description:         "Specify whether the ConfigMap or its key must be defined",
																		MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

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

															"field_ref": {
																Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.",
																MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.",

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

															"resource_field_ref": {
																Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

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

															"secret_key_ref": {
																Description:         "Selects a key of a secret in the pod's namespace",
																MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"key": {
																		Description:         "The key of the secret to select from.  Must be a valid secret key.",
																		MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
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
																		Description:         "Specify whether the Secret or its key must be defined",
																		MarkdownDescription: "Specify whether the Secret or its key must be defined",

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

											"experiment_annotations": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"experiment_image": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"node_selector": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.MapType{ElemType: types.StringType},

												Required: false,
												Optional: true,
												Computed: false,
											},

											"secrets": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"mount_path": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

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
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"status_check_timeouts": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"delay": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"timeout": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

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
												Description:         "Pod's tolerations.",
												MarkdownDescription: "Pod's tolerations.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"effect": {
														Description:         "Effect to match. Empty means all effects.",
														MarkdownDescription: "Effect to match. Empty means all effects.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"key": {
														Description:         "Taint key the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists.",
														MarkdownDescription: "Taint key the toleration applies to. Empty means match all taint keys. If the key is empty, operator must be Exists.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"operator": {
														Description:         "Operators are Exists or Equal. Defaults to Equal.",
														MarkdownDescription: "Operators are Exists or Equal. Defaults to Equal.",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"toleration_seconds": {
														Description:         "Period of time the toleration tolerates the taint.",
														MarkdownDescription: "Period of time the toleration tolerates the taint.",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"value": {
														Description:         "If the operator is Exists, the value should be empty, otherwise just a regular string.",
														MarkdownDescription: "If the operator is Exists, the value should be empty, otherwise just a regular string.",

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

									"probe": {
										Description:         "",
										MarkdownDescription: "",

										Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

											"cmd_probe_inputs": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"command": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},

													"comparator": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"criteria": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"type": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	stringvalidator.LengthAtLeast(1),

																	stringvalidator.RegexMatches(regexp.MustCompile(`^(int|float|string)$`), ""),
																},
															},

															"value": {
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

													"source": {
														Description:         "The external pod where we have to run the probe commands. It will run the commands inside the experiment pod itself(inline mode) if source contains a nil value",
														MarkdownDescription: "The external pod where we have to run the probe commands. It will run the commands inside the experiment pod itself(inline mode) if source contains a nil value",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"annotations": {
																Description:         "Annotations for the source pod",
																MarkdownDescription: "Annotations for the source pod",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"args": {
																Description:         "Args for the source pod",
																MarkdownDescription: "Args for the source pod",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"command": {
																Description:         "Command for the source pod",
																MarkdownDescription: "Command for the source pod",

																Type: types.ListType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"env": {
																Description:         "ENVList contains ENV passed to the source pod",
																MarkdownDescription: "ENVList contains ENV passed to the source pod",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "Name of the environment variable. Must be a C_IDENTIFIER.",
																		MarkdownDescription: "Name of the environment variable. Must be a C_IDENTIFIER.",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"value": {
																		Description:         "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
																		MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"value_from": {
																		Description:         "Source for the environment variable's value. Cannot be used if value is not empty.",
																		MarkdownDescription: "Source for the environment variable's value. Cannot be used if value is not empty.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"config_map_key_ref": {
																				Description:         "Selects a key of a ConfigMap.",
																				MarkdownDescription: "Selects a key of a ConfigMap.",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "The key to select.",
																						MarkdownDescription: "The key to select.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
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
																						Description:         "Specify whether the ConfigMap or its key must be defined",
																						MarkdownDescription: "Specify whether the ConfigMap or its key must be defined",

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

																			"field_ref": {
																				Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.",
																				MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations, spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.",

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

																			"resource_field_ref": {
																				Description:         "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",
																				MarkdownDescription: "Selects a resource of the container: only resources limits and requests (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.",

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

																						Type: types.StringType,

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

																			"secret_key_ref": {
																				Description:         "Selects a key of a secret in the pod's namespace",
																				MarkdownDescription: "Selects a key of a secret in the pod's namespace",

																				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "The key of the secret to select from.  Must be a valid secret key.",
																						MarkdownDescription: "The key of the secret to select from.  Must be a valid secret key.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
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
																						Description:         "Specify whether the Secret or its key must be defined",
																						MarkdownDescription: "Specify whether the Secret or its key must be defined",

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

															"host_network": {
																Description:         "HostNetwork define the hostNetwork of the external pod it supports boolean values and default value is false",
																MarkdownDescription: "HostNetwork define the hostNetwork of the external pod it supports boolean values and default value is false",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"image": {
																Description:         "Image for the source pod",
																MarkdownDescription: "Image for the source pod",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"image_pull_policy": {
																Description:         "ImagePullPolicy for the source pod",
																MarkdownDescription: "ImagePullPolicy for the source pod",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"image_pull_secrets": {
																Description:         "ImagePullSecrets for source pod",
																MarkdownDescription: "ImagePullSecrets for source pod",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"name": {
																		Description:         "Name of the referent",
																		MarkdownDescription: "Name of the referent",

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

															"inherit_inputs": {
																Description:         "InheritInputs define to inherit experiment details in probe pod it supports boolean values and default value is false.",
																MarkdownDescription: "InheritInputs define to inherit experiment details in probe pod it supports boolean values and default value is false.",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"labels": {
																Description:         "Labels for the source pod",
																MarkdownDescription: "Labels for the source pod",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"node_selector": {
																Description:         "NodeSelector for the source pod",
																MarkdownDescription: "NodeSelector for the source pod",

																Type: types.MapType{ElemType: types.StringType},

																Required: false,
																Optional: true,
																Computed: false,
															},

															"privileged": {
																Description:         "Privileged for the source pod",
																MarkdownDescription: "Privileged for the source pod",

																Type: types.BoolType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"volume_mount": {
																Description:         "VolumesMount for the source pod",
																MarkdownDescription: "VolumesMount for the source pod",

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
																		Description:         "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive. This field is beta in 1.15.",
																		MarkdownDescription: "Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to '' (volume's root). SubPathExpr and SubPath are mutually exclusive. This field is beta in 1.15.",

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
																Description:         "Volumes for the source pod",
																MarkdownDescription: "Volumes for the source pod",

																Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																	"aws_elastic_block_store": {
																		Description:         "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																		MarkdownDescription: "AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",
																				MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore TODO: how do we prevent errors in the filesystem from compromising the machine",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"partition": {
																				Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",
																				MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty).",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																				MarkdownDescription: "Specify 'true' to force and set the ReadOnly property in VolumeMounts to 'true'. If omitted, the default is 'false'. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"volume_id": {
																				Description:         "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",
																				MarkdownDescription: "Unique ID of the persistent disk resource in AWS (Amazon EBS volume). More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore",

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
																				Description:         "Host Caching mode: None, Read Only, Read Write.",
																				MarkdownDescription: "Host Caching mode: None, Read Only, Read Write.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"disk_name": {
																				Description:         "The Name of the data disk in the blob storage",
																				MarkdownDescription: "The Name of the data disk in the blob storage",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"disk_uri": {
																				Description:         "The URI the data disk in the blob storage",
																				MarkdownDescription: "The URI the data disk in the blob storage",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"fs_type": {
																				Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																				MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"kind": {
																				Description:         "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",
																				MarkdownDescription: "Expected values Shared: multiple blob disks per storage account  Dedicated: single blob disk per storage account  Managed: azure managed data disk (only in managed availability set). defaults to shared",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																				MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

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
																				Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																				MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_name": {
																				Description:         "the name of secret that contains Azure Storage Account Name and Key",
																				MarkdownDescription: "the name of secret that contains Azure Storage Account Name and Key",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"share_name": {
																				Description:         "Share Name",
																				MarkdownDescription: "Share Name",

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
																		Description:         "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",
																		MarkdownDescription: "CephFS represents a Ceph FS mount on the host that shares a pod's lifetime",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"monitors": {
																				Description:         "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																				MarkdownDescription: "Required: Monitors is a collection of Ceph monitors More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"path": {
																				Description:         "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",
																				MarkdownDescription: "Optional: Used as the mounted root, rather than the full Ceph tree, default is /",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																				MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_file": {
																				Description:         "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																				MarkdownDescription: "Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_ref": {
																				Description:         "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																				MarkdownDescription: "Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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
																				Description:         "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",
																				MarkdownDescription: "Optional: User is the rados user name, default is admin More info: https://examples.k8s.io/volumes/cephfs/README.md#how-to-use-it",

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
																		Description:         "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																		MarkdownDescription: "Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																				MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																				MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_ref": {
																				Description:         "Optional: points to a secret object containing parameters used to connect to OpenStack.",
																				MarkdownDescription: "Optional: points to a secret object containing parameters used to connect to OpenStack.",

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
																				Description:         "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",
																				MarkdownDescription: "volume id used to identify the volume in cinder. More info: https://examples.k8s.io/mysql-cinder-pd/README.md",

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
																		Description:         "ConfigMap represents a configMap that should populate this volume",
																		MarkdownDescription: "ConfigMap represents a configMap that should populate this volume",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"default_mode": {
																				Description:         "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"items": {
																				Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																				MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "The key to project.",
																						MarkdownDescription: "The key to project.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"mode": {
																						Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																						Type: types.Int64Type,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"path": {
																						Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																						MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
																				Description:         "Specify whether the ConfigMap or its keys must be defined",
																				MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",

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
																		Description:         "CSI (Container Storage Interface) represents storage that is handled by an external CSI driver (Alpha feature).",
																		MarkdownDescription: "CSI (Container Storage Interface) represents storage that is handled by an external CSI driver (Alpha feature).",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"driver": {
																				Description:         "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",
																				MarkdownDescription: "Driver is the name of the CSI driver that handles this volume. Consult with your admin for the correct name as registered in the cluster.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"fs_type": {
																				Description:         "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",
																				MarkdownDescription: "Filesystem type to mount. Ex. 'ext4', 'xfs', 'ntfs'. If not provided, the empty value is passed to the associated CSI driver which will determine the default filesystem to apply.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"node_publish_secret_ref": {
																				Description:         "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",
																				MarkdownDescription: "NodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and  may be empty if no secret is required. If the secret object contains more than one secret, all secret references are passed.",

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
																				Description:         "Specifies a read-only configuration for the volume. Defaults to false (read/write).",
																				MarkdownDescription: "Specifies a read-only configuration for the volume. Defaults to false (read/write).",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"volume_attributes": {
																				Description:         "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",
																				MarkdownDescription: "VolumeAttributes stores driver-specific properties that are passed to the CSI driver. Consult your driver's documentation for supported values.",

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
																		Description:         "DownwardAPI represents downward API about the pod that should populate this volume",
																		MarkdownDescription: "DownwardAPI represents downward API about the pod that should populate this volume",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"default_mode": {
																				Description:         "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
																						Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

																								Type: types.StringType,

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
																		Description:         "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																		MarkdownDescription: "EmptyDir represents a temporary directory that shares a pod's lifetime. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"medium": {
																				Description:         "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",
																				MarkdownDescription: "What type of storage medium should back this directory. The default is '' which means to use the node's default medium. Must be an empty string (default) or Memory. More info: https://kubernetes.io/docs/concepts/storage/volumes#emptydir",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"size_limit": {
																				Description:         "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",
																				MarkdownDescription: "Total amount of local storage required for this EmptyDir volume. The size limit is also applicable for memory medium. The maximum usage on memory medium EmptyDir would be the minimum value between the SizeLimit specified here and the sum of memory limits of all containers in a pod. The default is nil which means that the limit is undefined. More info: http://kubernetes.io/docs/user-guide/volumes#emptydir",

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

																	"fc": {
																		Description:         "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",
																		MarkdownDescription: "FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",
																				MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. TODO: how do we prevent errors in the filesystem from compromising the machine",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"lun": {
																				Description:         "Optional: FC target lun number",
																				MarkdownDescription: "Optional: FC target lun number",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																				MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"target_ww_ns": {
																				Description:         "Optional: FC target worldwide names (WWNs)",
																				MarkdownDescription: "Optional: FC target worldwide names (WWNs)",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"wwids": {
																				Description:         "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",
																				MarkdownDescription: "Optional: FC volume world wide identifiers (wwids) Either wwids or combination of targetWWNs and lun must be set, but not both simultaneously.",

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
																		Description:         "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",
																		MarkdownDescription: "FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"driver": {
																				Description:         "Driver is the name of the driver to use for this volume.",
																				MarkdownDescription: "Driver is the name of the driver to use for this volume.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"fs_type": {
																				Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",
																				MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. The default filesystem depends on FlexVolume script.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"options": {
																				Description:         "Optional: Extra command options if any.",
																				MarkdownDescription: "Optional: Extra command options if any.",

																				Type: types.MapType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																				MarkdownDescription: "Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_ref": {
																				Description:         "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",
																				MarkdownDescription: "Optional: SecretRef is reference to the secret object containing sensitive information to pass to the plugin scripts. This may be empty if no secret object is specified. If the secret object contains more than one secret, all secrets are passed to the plugin scripts.",

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
																		Description:         "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",
																		MarkdownDescription: "Flocker represents a Flocker volume attached to a kubelet's host machine. This depends on the Flocker control service being running",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"dataset_name": {
																				Description:         "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",
																				MarkdownDescription: "Name of the dataset stored as metadata -> name on the dataset for Flocker should be considered as deprecated",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"dataset_uuid": {
																				Description:         "UUID of the dataset. This is unique identifier of a Flocker dataset",
																				MarkdownDescription: "UUID of the dataset. This is unique identifier of a Flocker dataset",

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
																		Description:         "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																		MarkdownDescription: "GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",
																				MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk TODO: how do we prevent errors in the filesystem from compromising the machine",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"partition": {
																				Description:         "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																				MarkdownDescription: "The partition in the volume that you want to mount. If omitted, the default is to mount by volume name. Examples: For volume /dev/sda1, you specify the partition as '1'. Similarly, the volume partition for /dev/sda is '0' (or you can leave the property empty). More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"pd_name": {
																				Description:         "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																				MarkdownDescription: "Unique name of the PD resource in GCE. Used to identify the disk in GCE. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",
																				MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk",

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
																		Description:         "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",
																		MarkdownDescription: "GitRepo represents a git repository at a particular revision. DEPRECATED: GitRepo is deprecated. To provision a container with a git repo, mount an EmptyDir into an InitContainer that clones the repo using git, then mount the EmptyDir into the Pod's container.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"directory": {
																				Description:         "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",
																				MarkdownDescription: "Target directory name. Must not contain or start with '..'.  If '.' is supplied, the volume directory will be the git repository.  Otherwise, if specified, the volume will contain the git repository in the subdirectory with the given name.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"repository": {
																				Description:         "Repository URL",
																				MarkdownDescription: "Repository URL",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"revision": {
																				Description:         "Commit hash for the specified revision.",
																				MarkdownDescription: "Commit hash for the specified revision.",

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
																		Description:         "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",
																		MarkdownDescription: "Glusterfs represents a Glusterfs mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/glusterfs/README.md",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"endpoints": {
																				Description:         "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																				MarkdownDescription: "EndpointsName is the endpoint name that details Glusterfs topology. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"path": {
																				Description:         "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																				MarkdownDescription: "Path is the Glusterfs volume path. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",
																				MarkdownDescription: "ReadOnly here will force the Glusterfs volume to be mounted with read-only permissions. Defaults to false. More info: https://examples.k8s.io/volumes/glusterfs/README.md#create-a-pod",

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
																		Description:         "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",
																		MarkdownDescription: "HostPath represents a pre-existing file or directory on the host machine that is directly exposed to the container. This is generally used for system agents or other privileged things that are allowed to see the host machine. Most containers will NOT need this. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath --- TODO(jonesdl) We need to restrict who can use host directory mounts and who can/can not mount host directories as read/write.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"path": {
																				Description:         "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																				MarkdownDescription: "Path of the directory on the host. If the path is a symlink, it will follow the link to the real path. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"type": {
																				Description:         "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",
																				MarkdownDescription: "Type for HostPath Volume Defaults to '' More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath",

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
																		Description:         "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",
																		MarkdownDescription: "ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://examples.k8s.io/volumes/iscsi/README.md",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"chap_auth_discovery": {
																				Description:         "whether support iSCSI Discovery CHAP authentication",
																				MarkdownDescription: "whether support iSCSI Discovery CHAP authentication",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"chap_auth_session": {
																				Description:         "whether support iSCSI Session CHAP authentication",
																				MarkdownDescription: "whether support iSCSI Session CHAP authentication",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"fs_type": {
																				Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",
																				MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#iscsi TODO: how do we prevent errors in the filesystem from compromising the machine",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"initiator_name": {
																				Description:         "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",
																				MarkdownDescription: "Custom iSCSI Initiator Name. If initiatorName is specified with iscsiInterface simultaneously, new iSCSI interface <target portal>:<volume name> will be created for the connection.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"iqn": {
																				Description:         "Target iSCSI Qualified Name.",
																				MarkdownDescription: "Target iSCSI Qualified Name.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"iscsi_interface": {
																				Description:         "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",
																				MarkdownDescription: "iSCSI Interface Name that uses an iSCSI transport. Defaults to 'default' (tcp).",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"lun": {
																				Description:         "iSCSI Target Lun number.",
																				MarkdownDescription: "iSCSI Target Lun number.",

																				Type: types.Int64Type,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"portals": {
																				Description:         "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																				MarkdownDescription: "iSCSI Target Portal List. The portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",
																				MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false.",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_ref": {
																				Description:         "CHAP Secret for iSCSI target and initiator authentication",
																				MarkdownDescription: "CHAP Secret for iSCSI target and initiator authentication",

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
																				Description:         "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",
																				MarkdownDescription: "iSCSI Target Portal. The Portal is either an IP or ip_addr:port if the port is other than default (typically TCP ports 860 and 3260).",

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
																		Description:         "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
																		MarkdownDescription: "Volume's name. Must be a DNS_LABEL and unique within the pod. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",

																		Type: types.StringType,

																		Required: true,
																		Optional: false,
																		Computed: false,
																	},

																	"nfs": {
																		Description:         "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																		MarkdownDescription: "NFS represents an NFS mount on the host that shares a pod's lifetime More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"path": {
																				Description:         "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																				MarkdownDescription: "Path that is exported by the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																				MarkdownDescription: "ReadOnly here will force the NFS export to be mounted with read-only permissions. Defaults to false. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"server": {
																				Description:         "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",
																				MarkdownDescription: "Server is the hostname or IP address of the NFS server. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs",

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
																		Description:         "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																		MarkdownDescription: "PersistentVolumeClaimVolumeSource represents a reference to a PersistentVolumeClaim in the same namespace. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"claim_name": {
																				Description:         "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",
																				MarkdownDescription: "ClaimName is the name of a PersistentVolumeClaim in the same namespace as the pod using this volume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#persistentvolumeclaims",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Will force the ReadOnly setting in VolumeMounts. Default false.",
																				MarkdownDescription: "Will force the ReadOnly setting in VolumeMounts. Default false.",

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
																		Description:         "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",
																		MarkdownDescription: "PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																				MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"pd_id": {
																				Description:         "ID that identifies Photon Controller persistent disk",
																				MarkdownDescription: "ID that identifies Photon Controller persistent disk",

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
																		Description:         "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",
																		MarkdownDescription: "PortworxVolume represents a portworx volume attached and mounted on kubelets host machine",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",
																				MarkdownDescription: "FSType represents the filesystem type to mount Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs'. Implicitly inferred to be 'ext4' if unspecified.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																				MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"volume_id": {
																				Description:         "VolumeID uniquely identifies a Portworx volume",
																				MarkdownDescription: "VolumeID uniquely identifies a Portworx volume",

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
																		Description:         "Items for all in one resources secrets, configmaps, and downward API",
																		MarkdownDescription: "Items for all in one resources secrets, configmaps, and downward API",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"default_mode": {
																				Description:         "Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Mode bits to use on created files by default. Must be a value between 0 and 0777. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"sources": {
																				Description:         "list of volume projections",
																				MarkdownDescription: "list of volume projections",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"config_map": {
																						Description:         "information about the configMap data to project",
																						MarkdownDescription: "information about the configMap data to project",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"items": {
																								Description:         "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																								MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the ConfigMap, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"key": {
																										Description:         "The key to project.",
																										MarkdownDescription: "The key to project.",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"mode": {
																										Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																										MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																										Type: types.Int64Type,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"path": {
																										Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																										MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
																								Description:         "Specify whether the ConfigMap or its keys must be defined",
																								MarkdownDescription: "Specify whether the ConfigMap or its keys must be defined",

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
																						Description:         "information about the downwardAPI data to project",
																						MarkdownDescription: "information about the downwardAPI data to project",

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
																										Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																										MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

																												Type: types.StringType,

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
																						Description:         "information about the secret data to project",
																						MarkdownDescription: "information about the secret data to project",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"items": {
																								Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																								MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																									"key": {
																										Description:         "The key to project.",
																										MarkdownDescription: "The key to project.",

																										Type: types.StringType,

																										Required: true,
																										Optional: false,
																										Computed: false,
																									},

																									"mode": {
																										Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																										MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																										Type: types.Int64Type,

																										Required: false,
																										Optional: true,
																										Computed: false,
																									},

																									"path": {
																										Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																										MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
																								Description:         "Specify whether the Secret or its key must be defined",
																								MarkdownDescription: "Specify whether the Secret or its key must be defined",

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
																						Description:         "information about the serviceAccountToken data to project",
																						MarkdownDescription: "information about the serviceAccountToken data to project",

																						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																							"audience": {
																								Description:         "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",
																								MarkdownDescription: "Audience is the intended audience of the token. A recipient of a token must identify itself with an identifier specified in the audience of the token, and otherwise should reject the token. The audience defaults to the identifier of the apiserver.",

																								Type: types.StringType,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"expiration_seconds": {
																								Description:         "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",
																								MarkdownDescription: "ExpirationSeconds is the requested duration of validity of the service account token. As the token approaches expiration, the kubelet volume plugin will proactively rotate the service account token. The kubelet will start trying to rotate the token if the token is older than 80 percent of its time to live or if the token is older than 24 hours.Defaults to 1 hour and must be at least 10 minutes.",

																								Type: types.Int64Type,

																								Required: false,
																								Optional: true,
																								Computed: false,
																							},

																							"path": {
																								Description:         "Path is the path relative to the mount point of the file to project the token into.",
																								MarkdownDescription: "Path is the path relative to the mount point of the file to project the token into.",

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

																	"quobyte": {
																		Description:         "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",
																		MarkdownDescription: "Quobyte represents a Quobyte mount on the host that shares a pod's lifetime",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"group": {
																				Description:         "Group to map volume access to Default is no group",
																				MarkdownDescription: "Group to map volume access to Default is no group",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",
																				MarkdownDescription: "ReadOnly here will force the Quobyte volume to be mounted with read-only permissions. Defaults to false.",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"registry": {
																				Description:         "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",
																				MarkdownDescription: "Registry represents a single or multiple Quobyte Registry services specified as a string as host:port pair (multiple entries are separated with commas) which acts as the central registry for volumes",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"tenant": {
																				Description:         "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",
																				MarkdownDescription: "Tenant owning the given Quobyte volume in the Backend Used with dynamically provisioned Quobyte volumes, value is set by the plugin",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"user": {
																				Description:         "User to map volume access to Defaults to serivceaccount user",
																				MarkdownDescription: "User to map volume access to Defaults to serivceaccount user",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"volume": {
																				Description:         "Volume is a string that references an already created Quobyte volume by name.",
																				MarkdownDescription: "Volume is a string that references an already created Quobyte volume by name.",

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
																		Description:         "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",
																		MarkdownDescription: "RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",
																				MarkdownDescription: "Filesystem type of the volume that you want to mount. Tip: Ensure that the filesystem type is supported by the host operating system. Examples: 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified. More info: https://kubernetes.io/docs/concepts/storage/volumes#rbd TODO: how do we prevent errors in the filesystem from compromising the machine",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"image": {
																				Description:         "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																				MarkdownDescription: "The rados image name. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"keyring": {
																				Description:         "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																				MarkdownDescription: "Keyring is the path to key ring for RBDUser. Default is /etc/ceph/keyring. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"monitors": {
																				Description:         "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																				MarkdownDescription: "A collection of Ceph monitors. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

																				Type: types.ListType{ElemType: types.StringType},

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"pool": {
																				Description:         "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																				MarkdownDescription: "The rados pool name. Default is rbd. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																				MarkdownDescription: "ReadOnly here will force the ReadOnly setting in VolumeMounts. Defaults to false. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_ref": {
																				Description:         "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																				MarkdownDescription: "SecretRef is name of the authentication secret for RBDUser. If provided overrides keyring. Default is nil. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

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
																				Description:         "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",
																				MarkdownDescription: "The rados user name. Default is admin. More info: https://examples.k8s.io/volumes/rbd/README.md#how-to-use-it",

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
																		Description:         "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",
																		MarkdownDescription: "ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",
																				MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Default is 'xfs'.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"gateway": {
																				Description:         "The host address of the ScaleIO API Gateway.",
																				MarkdownDescription: "The host address of the ScaleIO API Gateway.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"protection_domain": {
																				Description:         "The name of the ScaleIO Protection Domain for the configured storage.",
																				MarkdownDescription: "The name of the ScaleIO Protection Domain for the configured storage.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																				MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_ref": {
																				Description:         "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",
																				MarkdownDescription: "SecretRef references to the secret for ScaleIO user and other sensitive information. If this is not provided, Login operation will fail.",

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
																				Description:         "Flag to enable/disable SSL communication with Gateway, default false",
																				MarkdownDescription: "Flag to enable/disable SSL communication with Gateway, default false",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"storage_mode": {
																				Description:         "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",
																				MarkdownDescription: "Indicates whether the storage for a volume should be ThickProvisioned or ThinProvisioned. Default is ThinProvisioned.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"storage_pool": {
																				Description:         "The ScaleIO Storage Pool associated with the protection domain.",
																				MarkdownDescription: "The ScaleIO Storage Pool associated with the protection domain.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"system": {
																				Description:         "The name of the storage system as configured in ScaleIO.",
																				MarkdownDescription: "The name of the storage system as configured in ScaleIO.",

																				Type: types.StringType,

																				Required: true,
																				Optional: false,
																				Computed: false,
																			},

																			"volume_name": {
																				Description:         "The name of a volume already created in the ScaleIO system that is associated with this volume source.",
																				MarkdownDescription: "The name of a volume already created in the ScaleIO system that is associated with this volume source.",

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
																		Description:         "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																		MarkdownDescription: "Secret represents a secret that should populate this volume. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"default_mode": {
																				Description:         "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits to use on created files by default. Must be a value between 0 and 0777. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																				Type: types.Int64Type,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"items": {
																				Description:         "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",
																				MarkdownDescription: "If unspecified, each key-value pair in the Data field of the referenced Secret will be projected into the volume as a file whose name is the key and content is the value. If specified, the listed keys will be projected into the specified paths, and unlisted keys will not be present. If a key is specified which is not present in the Secret, the volume setup will error unless it is marked optional. Paths must be relative and may not contain the '..' path or start with '..'.",

																				Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

																					"key": {
																						Description:         "The key to project.",
																						MarkdownDescription: "The key to project.",

																						Type: types.StringType,

																						Required: true,
																						Optional: false,
																						Computed: false,
																					},

																					"mode": {
																						Description:         "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																						MarkdownDescription: "Optional: mode bits to use on this file, must be a value between 0 and 0777. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

																						Type: types.Int64Type,

																						Required: false,
																						Optional: true,
																						Computed: false,
																					},

																					"path": {
																						Description:         "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",
																						MarkdownDescription: "The relative path of the file to map the key to. May not be an absolute path. May not contain the path element '..'. May not start with the string '..'.",

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
																				Description:         "Specify whether the Secret or its keys must be defined",
																				MarkdownDescription: "Specify whether the Secret or its keys must be defined",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_name": {
																				Description:         "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",
																				MarkdownDescription: "Name of the secret in the pod's namespace to use. More info: https://kubernetes.io/docs/concepts/storage/volumes#secret",

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
																		Description:         "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",
																		MarkdownDescription: "StorageOS represents a StorageOS volume attached and mounted on Kubernetes nodes.",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																				MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"read_only": {
																				Description:         "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",
																				MarkdownDescription: "Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts.",

																				Type: types.BoolType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"secret_ref": {
																				Description:         "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",
																				MarkdownDescription: "SecretRef specifies the secret to use for obtaining the StorageOS API credentials.  If not specified, default values will be attempted.",

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
																				Description:         "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",
																				MarkdownDescription: "VolumeName is the human-readable name of the StorageOS volume.  Volume names are only unique within a namespace.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"volume_namespace": {
																				Description:         "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",
																				MarkdownDescription: "VolumeNamespace specifies the scope of the volume within StorageOS.  If no namespace is specified then the Pod's namespace will be used.  This allows the Kubernetes name scoping to be mirrored within StorageOS for tighter integration. Set VolumeName to any name to override the default behaviour. Set to 'default' if you are not using namespaces within StorageOS. Namespaces that do not pre-exist within StorageOS will be created.",

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
																		Description:         "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",
																		MarkdownDescription: "VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine",

																		Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																			"fs_type": {
																				Description:         "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",
																				MarkdownDescription: "Filesystem type to mount. Must be a filesystem type supported by the host operating system. Ex. 'ext4', 'xfs', 'ntfs'. Implicitly inferred to be 'ext4' if unspecified.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"storage_policy_id": {
																				Description:         "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",
																				MarkdownDescription: "Storage Policy Based Management (SPBM) profile ID associated with the StoragePolicyName.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"storage_policy_name": {
																				Description:         "Storage Policy Based Management (SPBM) profile name.",
																				MarkdownDescription: "Storage Policy Based Management (SPBM) profile name.",

																				Type: types.StringType,

																				Required: false,
																				Optional: true,
																				Computed: false,
																			},

																			"volume_path": {
																				Description:         "Path that identifies vSphere volume vmdk",
																				MarkdownDescription: "Path that identifies vSphere volume vmdk",

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

											"data": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_probe_inputs": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"insecure_skip_verify": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"method": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"get": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"criteria": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtLeast(1),
																		},
																	},

																	"response_code": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtLeast(1),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	schemavalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("post")),
																},
															},

															"post": {
																Description:         "",
																MarkdownDescription: "",

																Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

																	"body": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"body_path": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"content_type": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtLeast(1),
																		},
																	},

																	"criteria": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtLeast(1),
																		},
																	},

																	"response_code": {
																		Description:         "",
																		MarkdownDescription: "",

																		Type: types.StringType,

																		Required: false,
																		Optional: true,
																		Computed: false,

																		Validators: []tfsdk.AttributeValidator{

																			stringvalidator.LengthAtLeast(1),
																		},
																	},
																}),

																Required: false,
																Optional: true,
																Computed: false,

																Validators: []tfsdk.AttributeValidator{

																	schemavalidator.AtLeastOneOf(path.MatchRelative().AtParent().AtName("get")),
																},
															},
														}),

														Required: false,
														Optional: true,
														Computed: false,
													},

													"url": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),
														},
													},
												}),

												Required: false,
												Optional: true,
												Computed: false,
											},

											"k8s_probe_inputs": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"field_selector": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"group": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"label_selector": {
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

													"operation": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,

														Validators: []tfsdk.AttributeValidator{

															stringvalidator.LengthAtLeast(1),

															stringvalidator.RegexMatches(regexp.MustCompile(`^(present|absent|create|delete)$`), ""),
														},
													},

													"resource": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"version": {
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

											"mode": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.RegexMatches(regexp.MustCompile(`^(SOT|EOT|Edge|Continuous|OnChaos)$`), ""),
												},
											},

											"name": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"prom_probe_inputs": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"comparator": {
														Description:         "",
														MarkdownDescription: "",

														Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

															"criteria": {
																Description:         "",
																MarkdownDescription: "",

																Type: types.StringType,

																Required: false,
																Optional: true,
																Computed: false,
															},

															"value": {
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

													"endpoint": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"query": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.StringType,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"query_path": {
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

											"run_properties": {
												Description:         "",
												MarkdownDescription: "",

												Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

													"initial_delay_seconds": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"interval": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"probe_polling_interval": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: false,
														Optional: true,
														Computed: false,
													},

													"probe_timeout": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"retry": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.Int64Type,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"stop_on_failure": {
														Description:         "",
														MarkdownDescription: "",

														Type: types.BoolType,

														Required: false,
														Optional: true,
														Computed: false,
													},
												}),

												Required: true,
												Optional: false,
												Computed: false,
											},

											"type": {
												Description:         "",
												MarkdownDescription: "",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,

												Validators: []tfsdk.AttributeValidator{

													stringvalidator.LengthAtLeast(1),

													stringvalidator.RegexMatches(regexp.MustCompile(`^(k8sProbe|httpProbe|cmdProbe|promProbe)$`), ""),
												},
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

					"job_clean_up_policy": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,

						Validators: []tfsdk.AttributeValidator{

							stringvalidator.RegexMatches(regexp.MustCompile(`^(delete|retain)$`), ""),
						},
					},

					"termination_grace_period_seconds": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.Int64Type,

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

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_litmuschaos_io_chaos_engine_v1alpha1")

	var state LitmuschaosIoChaosEngineV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LitmuschaosIoChaosEngineV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("litmuschaos.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ChaosEngine")

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

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_litmuschaos_io_chaos_engine_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_litmuschaos_io_chaos_engine_v1alpha1")

	var state LitmuschaosIoChaosEngineV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel LitmuschaosIoChaosEngineV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("litmuschaos.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ChaosEngine")

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

func (r *LitmuschaosIoChaosEngineV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_litmuschaos_io_chaos_engine_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
