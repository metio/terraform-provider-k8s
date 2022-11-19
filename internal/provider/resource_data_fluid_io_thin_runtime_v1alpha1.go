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

type DataFluidIoThinRuntimeV1Alpha1Resource struct{}

var (
	_ resource.Resource = (*DataFluidIoThinRuntimeV1Alpha1Resource)(nil)
)

type DataFluidIoThinRuntimeV1Alpha1TerraformModel struct {
	Id         types.Int64  `tfsdk:"id"`
	YAML       types.String `tfsdk:"yaml"`
	ApiVersion types.String `tfsdk:"api_version"`
	Kind       types.String `tfsdk:"kind"`
	Metadata   types.Object `tfsdk:"metadata"`
	Spec       types.Object `tfsdk:"spec"`
}

type DataFluidIoThinRuntimeV1Alpha1GoModel struct {
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
		DisablePrometheus *bool `tfsdk:"disable_prometheus" yaml:"disablePrometheus,omitempty"`

		Fuse *struct {
			Args *[]string `tfsdk:"args" yaml:"args,omitempty"`

			CleanPolicy *string `tfsdk:"clean_policy" yaml:"cleanPolicy,omitempty"`

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

			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

			ImageTag *string `tfsdk:"image_tag" yaml:"imageTag,omitempty"`

			LivenessProbe *struct {
				Exec *struct {
					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
				} `tfsdk:"exec" yaml:"exec,omitempty"`

				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

				TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

			NetworkMode *string `tfsdk:"network_mode" yaml:"networkMode,omitempty"`

			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

			Options *map[string]string `tfsdk:"options" yaml:"options,omitempty"`

			Ports *[]struct {
				ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

				HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

				HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
			} `tfsdk:"ports" yaml:"ports,omitempty"`

			ReadinessProbe *struct {
				Exec *struct {
					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
				} `tfsdk:"exec" yaml:"exec,omitempty"`

				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

				TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			VolumeMounts *[]struct {
				MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

				MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

				SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`
		} `tfsdk:"fuse" yaml:"fuse,omitempty"`

		ProfileName *string `tfsdk:"profile_name" yaml:"profileName,omitempty"`

		Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

		RunAs *struct {
			Gid *int64 `tfsdk:"gid" yaml:"gid,omitempty"`

			Group *string `tfsdk:"group" yaml:"group,omitempty"`

			Uid *int64 `tfsdk:"uid" yaml:"uid,omitempty"`

			User *string `tfsdk:"user" yaml:"user,omitempty"`
		} `tfsdk:"run_as" yaml:"runAs,omitempty"`

		Tieredstore *struct {
			Levels *[]struct {
				High *string `tfsdk:"high" yaml:"high,omitempty"`

				Low *string `tfsdk:"low" yaml:"low,omitempty"`

				Mediumtype *string `tfsdk:"mediumtype" yaml:"mediumtype,omitempty"`

				Path *string `tfsdk:"path" yaml:"path,omitempty"`

				Quota utilities.IntOrString `tfsdk:"quota" yaml:"quota,omitempty"`

				QuotaList *string `tfsdk:"quota_list" yaml:"quotaList,omitempty"`

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

				VolumeType *string `tfsdk:"volume_type" yaml:"volumeType,omitempty"`
			} `tfsdk:"levels" yaml:"levels,omitempty"`
		} `tfsdk:"tieredstore" yaml:"tieredstore,omitempty"`

		Version *struct {
			Image *string `tfsdk:"image" yaml:"image,omitempty"`

			ImagePullPolicy *string `tfsdk:"image_pull_policy" yaml:"imagePullPolicy,omitempty"`

			ImageTag *string `tfsdk:"image_tag" yaml:"imageTag,omitempty"`
		} `tfsdk:"version" yaml:"version,omitempty"`

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

		Worker *struct {
			Enabled *bool `tfsdk:"enabled" yaml:"enabled,omitempty"`

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

			LivenessProbe *struct {
				Exec *struct {
					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
				} `tfsdk:"exec" yaml:"exec,omitempty"`

				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

				TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"liveness_probe" yaml:"livenessProbe,omitempty"`

			NetworkMode *string `tfsdk:"network_mode" yaml:"networkMode,omitempty"`

			NodeSelector *map[string]string `tfsdk:"node_selector" yaml:"nodeSelector,omitempty"`

			Ports *[]struct {
				ContainerPort *int64 `tfsdk:"container_port" yaml:"containerPort,omitempty"`

				HostIP *string `tfsdk:"host_ip" yaml:"hostIP,omitempty"`

				HostPort *int64 `tfsdk:"host_port" yaml:"hostPort,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				Protocol *string `tfsdk:"protocol" yaml:"protocol,omitempty"`
			} `tfsdk:"ports" yaml:"ports,omitempty"`

			ReadinessProbe *struct {
				Exec *struct {
					Command *[]string `tfsdk:"command" yaml:"command,omitempty"`
				} `tfsdk:"exec" yaml:"exec,omitempty"`

				FailureThreshold *int64 `tfsdk:"failure_threshold" yaml:"failureThreshold,omitempty"`

				Grpc *struct {
					Port *int64 `tfsdk:"port" yaml:"port,omitempty"`

					Service *string `tfsdk:"service" yaml:"service,omitempty"`
				} `tfsdk:"grpc" yaml:"grpc,omitempty"`

				HttpGet *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					HttpHeaders *[]struct {
						Name *string `tfsdk:"name" yaml:"name,omitempty"`

						Value *string `tfsdk:"value" yaml:"value,omitempty"`
					} `tfsdk:"http_headers" yaml:"httpHeaders,omitempty"`

					Path *string `tfsdk:"path" yaml:"path,omitempty"`

					Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`

					Scheme *string `tfsdk:"scheme" yaml:"scheme,omitempty"`
				} `tfsdk:"http_get" yaml:"httpGet,omitempty"`

				InitialDelaySeconds *int64 `tfsdk:"initial_delay_seconds" yaml:"initialDelaySeconds,omitempty"`

				PeriodSeconds *int64 `tfsdk:"period_seconds" yaml:"periodSeconds,omitempty"`

				SuccessThreshold *int64 `tfsdk:"success_threshold" yaml:"successThreshold,omitempty"`

				TcpSocket *struct {
					Host *string `tfsdk:"host" yaml:"host,omitempty"`

					Port utilities.IntOrString `tfsdk:"port" yaml:"port,omitempty"`
				} `tfsdk:"tcp_socket" yaml:"tcpSocket,omitempty"`

				TerminationGracePeriodSeconds *int64 `tfsdk:"termination_grace_period_seconds" yaml:"terminationGracePeriodSeconds,omitempty"`

				TimeoutSeconds *int64 `tfsdk:"timeout_seconds" yaml:"timeoutSeconds,omitempty"`
			} `tfsdk:"readiness_probe" yaml:"readinessProbe,omitempty"`

			Replicas *int64 `tfsdk:"replicas" yaml:"replicas,omitempty"`

			Resources *struct {
				Limits *map[string]string `tfsdk:"limits" yaml:"limits,omitempty"`

				Requests *map[string]string `tfsdk:"requests" yaml:"requests,omitempty"`
			} `tfsdk:"resources" yaml:"resources,omitempty"`

			VolumeMounts *[]struct {
				MountPath *string `tfsdk:"mount_path" yaml:"mountPath,omitempty"`

				MountPropagation *string `tfsdk:"mount_propagation" yaml:"mountPropagation,omitempty"`

				Name *string `tfsdk:"name" yaml:"name,omitempty"`

				ReadOnly *bool `tfsdk:"read_only" yaml:"readOnly,omitempty"`

				SubPath *string `tfsdk:"sub_path" yaml:"subPath,omitempty"`

				SubPathExpr *string `tfsdk:"sub_path_expr" yaml:"subPathExpr,omitempty"`
			} `tfsdk:"volume_mounts" yaml:"volumeMounts,omitempty"`
		} `tfsdk:"worker" yaml:"worker,omitempty"`
	} `tfsdk:"spec" yaml:"spec,omitempty"`
}

func NewDataFluidIoThinRuntimeV1Alpha1Resource() resource.Resource {
	return &DataFluidIoThinRuntimeV1Alpha1Resource{}
}

func (r *DataFluidIoThinRuntimeV1Alpha1Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_data_fluid_io_thin_runtime_v1alpha1"
}

func (r *DataFluidIoThinRuntimeV1Alpha1Resource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description:         "ThinRuntime is the Schema for the thinruntimes API",
		MarkdownDescription: "ThinRuntime is the Schema for the thinruntimes API",
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
				Description:         "ThinRuntimeSpec defines the desired state of ThinRuntime",
				MarkdownDescription: "ThinRuntimeSpec defines the desired state of ThinRuntime",

				Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

					"disable_prometheus": {
						Description:         "Disable monitoring for Runtime Prometheus is enabled by default",
						MarkdownDescription: "Disable monitoring for Runtime Prometheus is enabled by default",

						Type: types.BoolType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"fuse": {
						Description:         "The component spec of thinRuntime",
						MarkdownDescription: "The component spec of thinRuntime",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"args": {
								Description:         "Arguments that will be passed to thinRuntime Fuse",
								MarkdownDescription: "Arguments that will be passed to thinRuntime Fuse",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"clean_policy": {
								Description:         "CleanPolicy decides when to clean thinRuntime Fuse pods. Currently Fluid supports two policies: OnDemand and OnRuntimeDeleted OnDemand cleans fuse pod once the fuse pod on some node is not needed OnRuntimeDeleted cleans fuse pod only when the cache runtime is deleted Defaults to OnDemand",
								MarkdownDescription: "CleanPolicy decides when to clean thinRuntime Fuse pods. Currently Fluid supports two policies: OnDemand and OnRuntimeDeleted OnDemand cleans fuse pod once the fuse pod on some node is not needed OnRuntimeDeleted cleans fuse pod only when the cache runtime is deleted Defaults to OnDemand",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"command": {
								Description:         "Command that will be passed to thinRuntime Fuse",
								MarkdownDescription: "Command that will be passed to thinRuntime Fuse",

								Type: types.ListType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"env": {
								Description:         "Environment variables that will be used by thinRuntime Fuse",
								MarkdownDescription: "Environment variables that will be used by thinRuntime Fuse",

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
										Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

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
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

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

							"image": {
								Description:         "Image for thinRuntime fuse",
								MarkdownDescription: "Image for thinRuntime fuse",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_policy": {
								Description:         "One of the three policies: 'Always', 'IfNotPresent', 'Never'",
								MarkdownDescription: "One of the three policies: 'Always', 'IfNotPresent', 'Never'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_tag": {
								Description:         "Image for thinRuntime fuse",
								MarkdownDescription: "Image for thinRuntime fuse",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"liveness_probe": {
								Description:         "livenessProbe of thin fuse pod",
								MarkdownDescription: "livenessProbe of thin fuse pod",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_headers": {
												Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
												MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The header field name",
														MarkdownDescription: "The header field name",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "The header field value",
														MarkdownDescription: "The header field value",

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

											"path": {
												Description:         "Path to access on the HTTP server.",
												MarkdownDescription: "Path to access on the HTTP server.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: utilities.IntOrStringType{},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"scheme": {
												Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
												MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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

									"initial_delay_seconds": {
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tcp_socket": {
										Description:         "TCPSocket specifies an action involving a TCP port.",
										MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Optional: Host name to connect to, defaults to the pod IP.",
												MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: utilities.IntOrStringType{},

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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

							"network_mode": {
								Description:         "Whether to use hostnetwork or not",
								MarkdownDescription: "Whether to use hostnetwork or not",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("HostNetwork", "", "ContainerNetwork"),
								},
							},

							"node_selector": {
								Description:         "NodeSelector is a selector which must be true for the fuse client to fit on a node, this option only effect when global is enabled",
								MarkdownDescription: "NodeSelector is a selector which must be true for the fuse client to fit on a node, this option only effect when global is enabled",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"options": {
								Description:         "Options configurable options of FUSE client, performance parameters usually. will be merged with Dataset.spec.mounts.options into fuse pod.",
								MarkdownDescription: "Options configurable options of FUSE client, performance parameters usually. will be merged with Dataset.spec.mounts.options into fuse pod.",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ports": {
								Description:         "Ports used thinRuntime",
								MarkdownDescription: "Ports used thinRuntime",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"container_port": {
										Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
										MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"host_ip": {
										Description:         "What host IP to bind the external port to.",
										MarkdownDescription: "What host IP to bind the external port to.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host_port": {
										Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
										MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
										MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol": {
										Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
										MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",

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

							"readiness_probe": {
								Description:         "readinessProbe of thin fuse pod",
								MarkdownDescription: "readinessProbe of thin fuse pod",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_headers": {
												Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
												MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The header field name",
														MarkdownDescription: "The header field name",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "The header field value",
														MarkdownDescription: "The header field value",

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

											"path": {
												Description:         "Path to access on the HTTP server.",
												MarkdownDescription: "Path to access on the HTTP server.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: utilities.IntOrStringType{},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"scheme": {
												Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
												MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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

									"initial_delay_seconds": {
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tcp_socket": {
										Description:         "TCPSocket specifies an action involving a TCP port.",
										MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Optional: Host name to connect to, defaults to the pod IP.",
												MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: utilities.IntOrStringType{},

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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

							"resources": {
								Description:         "Resources that will be requested by thinRuntime Fuse.",
								MarkdownDescription: "Resources that will be requested by thinRuntime Fuse.",

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

							"volume_mounts": {
								Description:         "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the thinruntime component's filesystem.",
								MarkdownDescription: "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into the thinruntime component's filesystem.",

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
						}),

						Required: false,
						Optional: true,
						Computed: false,
					},

					"profile_name": {
						Description:         "",
						MarkdownDescription: "",

						Type: types.StringType,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"replicas": {
						Description:         "The replicas of the worker, need to be specified",
						MarkdownDescription: "The replicas of the worker, need to be specified",

						Type: types.Int64Type,

						Required: false,
						Optional: true,
						Computed: false,
					},

					"run_as": {
						Description:         "Manage the user to run Runtime",
						MarkdownDescription: "Manage the user to run Runtime",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"gid": {
								Description:         "The gid to run the alluxio runtime",
								MarkdownDescription: "The gid to run the alluxio runtime",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"group": {
								Description:         "The group name to run the alluxio runtime",
								MarkdownDescription: "The group name to run the alluxio runtime",

								Type: types.StringType,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"uid": {
								Description:         "The uid to run the alluxio runtime",
								MarkdownDescription: "The uid to run the alluxio runtime",

								Type: types.Int64Type,

								Required: true,
								Optional: false,
								Computed: false,
							},

							"user": {
								Description:         "The user name to run the alluxio runtime",
								MarkdownDescription: "The user name to run the alluxio runtime",

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

					"tieredstore": {
						Description:         "Tiered storage",
						MarkdownDescription: "Tiered storage",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"levels": {
								Description:         "configurations for multiple tiers",
								MarkdownDescription: "configurations for multiple tiers",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"high": {
										Description:         "Ratio of high watermark of the tier (e.g. 0.9)",
										MarkdownDescription: "Ratio of high watermark of the tier (e.g. 0.9)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"low": {
										Description:         "Ratio of low watermark of the tier (e.g. 0.7)",
										MarkdownDescription: "Ratio of low watermark of the tier (e.g. 0.7)",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"mediumtype": {
										Description:         "Medium Type of the tier. One of the three types: 'MEM', 'SSD', 'HDD'",
										MarkdownDescription: "Medium Type of the tier. One of the three types: 'MEM', 'SSD', 'HDD'",

										Type: types.StringType,

										Required: true,
										Optional: false,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("MEM", "SSD", "HDD"),
										},
									},

									"path": {
										Description:         "File paths to be used for the tier. Multiple paths are supported. Multiple paths should be separated with comma. For example: '/mnt/cache1,/mnt/cache2'.",
										MarkdownDescription: "File paths to be used for the tier. Multiple paths are supported. Multiple paths should be separated with comma. For example: '/mnt/cache1,/mnt/cache2'.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.LengthAtLeast(1),
										},
									},

									"quota": {
										Description:         "Quota for the whole tier. (e.g. 100Gi) Please note that if there're multiple paths used for this tierstore, the quota will be equally divided into these paths. If you'd like to set quota for each, path, see QuotaList for more information.",
										MarkdownDescription: "Quota for the whole tier. (e.g. 100Gi) Please note that if there're multiple paths used for this tierstore, the quota will be equally divided into these paths. If you'd like to set quota for each, path, see QuotaList for more information.",

										Type: utilities.IntOrStringType{},

										Required: false,
										Optional: true,
										Computed: false,
									},

									"quota_list": {
										Description:         "QuotaList are quotas used to set quota on multiple paths. Quotas should be separated with comma. Quotas in this list will be set to paths with the same order in Path. For example, with Path defined with '/mnt/cache1,/mnt/cache2' and QuotaList set to '100Gi, 50Gi', then we get 100GiB cache storage under '/mnt/cache1' and 50GiB under '/mnt/cache2'. Also note that num of quotas must be consistent with the num of paths defined in Path.",
										MarkdownDescription: "QuotaList are quotas used to set quota on multiple paths. Quotas should be separated with comma. Quotas in this list will be set to paths with the same order in Path. For example, with Path defined with '/mnt/cache1,/mnt/cache2' and QuotaList set to '100Gi, 50Gi', then we get 100GiB cache storage under '/mnt/cache1' and 50GiB under '/mnt/cache2'. Also note that num of quotas must be consistent with the num of paths defined in Path.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.RegexMatches(regexp.MustCompile(`^((\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+)))),)+((\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))(([KMGTPE]i)|[numkMGTPE]|([eE](\+|-)?(([0-9]+(\.[0-9]*)?)|(\.[0-9]+))))?)$`), ""),
										},
									},

									"volume_source": {
										Description:         "VolumeSource is the volume source of the tier. It follows the form of corev1.VolumeSource. For now, users should only specify VolumeSource when VolumeType is set to emptyDir.",
										MarkdownDescription: "VolumeSource is the volume source of the tier. It follows the form of corev1.VolumeSource. For now, users should only specify VolumeSource when VolumeType is set to emptyDir.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

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
														Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
																Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
												Description:         "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
												MarkdownDescription: "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",

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
												Description:         "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
												MarkdownDescription: "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",

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
																		Description:         "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
																		MarkdownDescription: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

																		Type: types.ListType{ElemType: types.StringType},

																		Required: false,
																		Optional: true,
																		Computed: false,
																	},

																	"data_source": {
																		Description:         "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
																		MarkdownDescription: "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",

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
																		Description:         "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
																		MarkdownDescription: "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",

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
																		Description:         "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
																		MarkdownDescription: "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",

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
																		Description:         "A label query over volumes to consider for binding.",
																		MarkdownDescription: "A label query over volumes to consider for binding.",

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
																		Description:         "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
																		MarkdownDescription: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",

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
																		Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
																		MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",

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
														Description:         "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
																				Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
																				Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																				MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
														Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
														MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
																Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

									"volume_type": {
										Description:         "VolumeType is the volume type of the tier. Should be one of the three types: 'hostPath', 'emptyDir' and 'volumeTemplate'. If not set, defaults to hostPath.",
										MarkdownDescription: "VolumeType is the volume type of the tier. Should be one of the three types: 'hostPath', 'emptyDir' and 'volumeTemplate'. If not set, defaults to hostPath.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,

										Validators: []tfsdk.AttributeValidator{

											stringvalidator.OneOf("hostPath", "emptyDir"),
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

					"version": {
						Description:         "The version information that instructs fluid to orchestrate a particular version,",
						MarkdownDescription: "The version information that instructs fluid to orchestrate a particular version,",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"image": {
								Description:         "Image (e.g. alluxio/alluxio)",
								MarkdownDescription: "Image (e.g. alluxio/alluxio)",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_pull_policy": {
								Description:         "One of the three policies: 'Always', 'IfNotPresent', 'Never'",
								MarkdownDescription: "One of the three policies: 'Always', 'IfNotPresent', 'Never'",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"image_tag": {
								Description:         "Image tag (e.g. 2.3.0-SNAPSHOT)",
								MarkdownDescription: "Image tag (e.g. 2.3.0-SNAPSHOT)",

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
						Description:         "Volumes is the list of Kubernetes volumes that can be mounted by runtime components and/or fuses.",
						MarkdownDescription: "Volumes is the list of Kubernetes volumes that can be mounted by runtime components and/or fuses.",

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
										Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
										MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
												Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
												MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
								Description:         "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",
								MarkdownDescription: "CSI (Container Storage Interface) represents ephemeral storage that is handled by certain external CSI drivers (Beta feature).",

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
								Description:         "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",
								MarkdownDescription: "Ephemeral represents a volume that is handled by a cluster storage driver. The volume's lifecycle is tied to the pod that defines it - it will be created before the pod starts, and deleted when the pod is removed.  Use this if: a) the volume is only needed while the pod runs, b) features of normal volumes like restoring from snapshot or capacity tracking are needed, c) the storage driver is specified through a storage class, and d) the storage driver supports dynamic volume provisioning through a PersistentVolumeClaim (see EphemeralVolumeSource for more information on the connection between this volume type and PersistentVolumeClaim).  Use PersistentVolumeClaim or one of the vendor-specific APIs for volumes that persist for longer than the lifecycle of an individual pod.  Use CSI for light-weight local ephemeral volumes if the CSI driver is meant to be used that way - see the documentation of the driver for more information.  A pod can use both types of ephemeral volumes and persistent volumes at the same time.",

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
														Description:         "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",
														MarkdownDescription: "AccessModes contains the desired access modes the volume should have. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1",

														Type: types.ListType{ElemType: types.StringType},

														Required: false,
														Optional: true,
														Computed: false,
													},

													"data_source": {
														Description:         "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",
														MarkdownDescription: "This field can be used to specify either: * An existing VolumeSnapshot object (snapshot.storage.k8s.io/VolumeSnapshot) * An existing PVC (PersistentVolumeClaim) If the provisioner or an external controller can support the specified data source, it will create a new volume based on the contents of the specified data source. If the AnyVolumeDataSource feature gate is enabled, this field will always have the same contents as the DataSourceRef field.",

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
														Description:         "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",
														MarkdownDescription: "Specifies the object from which to populate the volume with data, if a non-empty volume is desired. This may be any local object from a non-empty API group (non core object) or a PersistentVolumeClaim object. When this field is specified, volume binding will only succeed if the type of the specified object matches some installed volume populator or dynamic provisioner. This field will replace the functionality of the DataSource field and as such if both fields are non-empty, they must have the same value. For backwards compatibility, both fields (DataSource and DataSourceRef) will be set to the same value automatically if one of them is empty and the other is non-empty. There are two important differences between DataSource and DataSourceRef: * While DataSource only allows two specific types of objects, DataSourceRef allows any non-core object, as well as PersistentVolumeClaim objects. * While DataSource ignores disallowed values (dropping them), DataSourceRef preserves all values, and generates an error if a disallowed value is specified. (Alpha) Using this field requires the AnyVolumeDataSource feature gate to be enabled.",

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
														Description:         "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",
														MarkdownDescription: "Resources represents the minimum resources the volume should have. If RecoverVolumeExpansionFailure feature is enabled users are allowed to specify resource requirements that are lower than previous value but must still be higher than capacity recorded in the status field of the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources",

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
														Description:         "A label query over volumes to consider for binding.",
														MarkdownDescription: "A label query over volumes to consider for binding.",

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
														Description:         "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",
														MarkdownDescription: "Name of the StorageClass required by the claim. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1",

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
														Description:         "VolumeName is the binding reference to the PersistentVolume backing this claim.",
														MarkdownDescription: "VolumeName is the binding reference to the PersistentVolume backing this claim.",

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
										Description:         "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
										MarkdownDescription: "Mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
																Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
																Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
																MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
										Description:         "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
										MarkdownDescription: "Optional: mode bits used to set permissions on created files by default. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. Defaults to 0644. Directories within the path are not affected by this setting. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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
												Description:         "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",
												MarkdownDescription: "Optional: mode bits used to set permissions on this file. Must be an octal value between 0000 and 0777 or a decimal value between 0 and 511. YAML accepts both octal and decimal values, JSON requires decimal values for mode bits. If not specified, the volume defaultMode will be used. This might be in conflict with other options that affect the file mode, like fsGroup, and the result can be other mode bits set.",

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

					"worker": {
						Description:         "The component spec of worker",
						MarkdownDescription: "The component spec of worker",

						Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

							"enabled": {
								Description:         "Enabled or Disabled for the components.",
								MarkdownDescription: "Enabled or Disabled for the components.",

								Type: types.BoolType,

								Required: false,
								Optional: true,
								Computed: false,
							},

							"env": {
								Description:         "Environment variables that will be used by thinRuntime component.",
								MarkdownDescription: "Environment variables that will be used by thinRuntime component.",

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
										Description:         "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",
										MarkdownDescription: "Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables. If a variable cannot be resolved, the reference in the input string will be unchanged. Double $$ are reduced to a single $, which allows for escaping the $(VAR_NAME) syntax: i.e. '$$(VAR_NAME)' will produce the string literal '$(VAR_NAME)'. Escaped references will never be expanded, regardless of whether the variable exists or not. Defaults to ''.",

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
												Description:         "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",
												MarkdownDescription: "Selects a field of the pod: supports metadata.name, metadata.namespace, 'metadata.labels['<KEY>']', 'metadata.annotations['<KEY>']', spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP, status.podIPs.",

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

							"liveness_probe": {
								Description:         "livenessProbe of thin fuse pod",
								MarkdownDescription: "livenessProbe of thin fuse pod",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_headers": {
												Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
												MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The header field name",
														MarkdownDescription: "The header field name",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "The header field value",
														MarkdownDescription: "The header field value",

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

											"path": {
												Description:         "Path to access on the HTTP server.",
												MarkdownDescription: "Path to access on the HTTP server.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: utilities.IntOrStringType{},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"scheme": {
												Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
												MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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

									"initial_delay_seconds": {
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tcp_socket": {
										Description:         "TCPSocket specifies an action involving a TCP port.",
										MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Optional: Host name to connect to, defaults to the pod IP.",
												MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: utilities.IntOrStringType{},

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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

							"network_mode": {
								Description:         "Whether to use hostnetwork or not",
								MarkdownDescription: "Whether to use hostnetwork or not",

								Type: types.StringType,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									stringvalidator.OneOf("HostNetwork", "", "ContainerNetwork"),
								},
							},

							"node_selector": {
								Description:         "NodeSelector is a selector",
								MarkdownDescription: "NodeSelector is a selector",

								Type: types.MapType{ElemType: types.StringType},

								Required: false,
								Optional: true,
								Computed: false,
							},

							"ports": {
								Description:         "Ports used thinRuntime",
								MarkdownDescription: "Ports used thinRuntime",

								Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

									"container_port": {
										Description:         "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",
										MarkdownDescription: "Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.",

										Type: types.Int64Type,

										Required: true,
										Optional: false,
										Computed: false,
									},

									"host_ip": {
										Description:         "What host IP to bind the external port to.",
										MarkdownDescription: "What host IP to bind the external port to.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"host_port": {
										Description:         "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",
										MarkdownDescription: "Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"name": {
										Description:         "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",
										MarkdownDescription: "If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.",

										Type: types.StringType,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"protocol": {
										Description:         "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",
										MarkdownDescription: "Protocol for port. Must be UDP, TCP, or SCTP. Defaults to 'TCP'.",

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

							"readiness_probe": {
								Description:         "readinessProbe of thin fuse pod",
								MarkdownDescription: "readinessProbe of thin fuse pod",

								Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

									"exec": {
										Description:         "Exec specifies the action to take.",
										MarkdownDescription: "Exec specifies the action to take.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"command": {
												Description:         "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",
												MarkdownDescription: "Command is the command line to execute inside the container, the working directory for the command  is root ('/') in the container's filesystem. The command is simply exec'd, it is not run inside a shell, so traditional shell instructions ('|', etc) won't work. To use a shell, you need to explicitly call out to that shell. Exit status of 0 is treated as live/healthy and non-zero is unhealthy.",

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

									"failure_threshold": {
										Description:         "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"grpc": {
										Description:         "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",
										MarkdownDescription: "GRPC specifies an action involving a GRPC port. This is an alpha field and requires enabling GRPCContainerProbe feature gate.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"port": {
												Description:         "Port number of the gRPC service. Number must be in the range 1 to 65535.",
												MarkdownDescription: "Port number of the gRPC service. Number must be in the range 1 to 65535.",

												Type: types.Int64Type,

												Required: true,
												Optional: false,
												Computed: false,
											},

											"service": {
												Description:         "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",
												MarkdownDescription: "Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).  If this is not specified, the default behavior is defined by gRPC.",

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

									"http_get": {
										Description:         "HTTPGet specifies the http request to perform.",
										MarkdownDescription: "HTTPGet specifies the http request to perform.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",
												MarkdownDescription: "Host name to connect to, defaults to the pod IP. You probably want to set 'Host' in httpHeaders instead.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"http_headers": {
												Description:         "Custom headers to set in the request. HTTP allows repeated headers.",
												MarkdownDescription: "Custom headers to set in the request. HTTP allows repeated headers.",

												Attributes: tfsdk.ListNestedAttributes(map[string]tfsdk.Attribute{

													"name": {
														Description:         "The header field name",
														MarkdownDescription: "The header field name",

														Type: types.StringType,

														Required: true,
														Optional: false,
														Computed: false,
													},

													"value": {
														Description:         "The header field value",
														MarkdownDescription: "The header field value",

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

											"path": {
												Description:         "Path to access on the HTTP server.",
												MarkdownDescription: "Path to access on the HTTP server.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Name or number of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: utilities.IntOrStringType{},

												Required: true,
												Optional: false,
												Computed: false,
											},

											"scheme": {
												Description:         "Scheme to use for connecting to the host. Defaults to HTTP.",
												MarkdownDescription: "Scheme to use for connecting to the host. Defaults to HTTP.",

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

									"initial_delay_seconds": {
										Description:         "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after the container has started before liveness probes are initiated. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"period_seconds": {
										Description:         "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",
										MarkdownDescription: "How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"success_threshold": {
										Description:         "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",
										MarkdownDescription: "Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness and startup. Minimum value is 1.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"tcp_socket": {
										Description:         "TCPSocket specifies an action involving a TCP port.",
										MarkdownDescription: "TCPSocket specifies an action involving a TCP port.",

										Attributes: tfsdk.SingleNestedAttributes(map[string]tfsdk.Attribute{

											"host": {
												Description:         "Optional: Host name to connect to, defaults to the pod IP.",
												MarkdownDescription: "Optional: Host name to connect to, defaults to the pod IP.",

												Type: types.StringType,

												Required: false,
												Optional: true,
												Computed: false,
											},

											"port": {
												Description:         "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",
												MarkdownDescription: "Number or name of the port to access on the container. Number must be in the range 1 to 65535. Name must be an IANA_SVC_NAME.",

												Type: utilities.IntOrStringType{},

												Required: true,
												Optional: false,
												Computed: false,
											},
										}),

										Required: false,
										Optional: true,
										Computed: false,
									},

									"termination_grace_period_seconds": {
										Description:         "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",
										MarkdownDescription: "Optional duration in seconds the pod needs to terminate gracefully upon probe failure. The grace period is the duration in seconds after the processes running in the pod are sent a termination signal and the time when the processes are forcibly halted with a kill signal. Set this value longer than the expected cleanup time for your process. If this value is nil, the pod's terminationGracePeriodSeconds will be used. Otherwise, this value overrides the value provided by the pod spec. Value must be non-negative integer. The value zero indicates stop immediately via the kill signal (no opportunity to shut down). This is a beta field and requires enabling ProbeTerminationGracePeriod feature gate. Minimum value is 1. spec.terminationGracePeriodSeconds is used if unset.",

										Type: types.Int64Type,

										Required: false,
										Optional: true,
										Computed: false,
									},

									"timeout_seconds": {
										Description:         "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",
										MarkdownDescription: "Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes",

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

							"replicas": {
								Description:         "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",
								MarkdownDescription: "Replicas is the desired number of replicas of the given template. If unspecified, defaults to 1. replicas is the min replicas of dataset in the cluster",

								Type: types.Int64Type,

								Required: false,
								Optional: true,
								Computed: false,

								Validators: []tfsdk.AttributeValidator{

									int64validator.AtLeast(1),
								},
							},

							"resources": {
								Description:         "Resources that will be requested by thinRuntime component.",
								MarkdownDescription: "Resources that will be requested by thinRuntime component.",

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

							"volume_mounts": {
								Description:         "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into runtime component's filesystem.",
								MarkdownDescription: "VolumeMounts specifies the volumes listed in '.spec.volumes' to mount into runtime component's filesystem.",

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

func (r *DataFluidIoThinRuntimeV1Alpha1Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Create resource k8s_data_fluid_io_thin_runtime_v1alpha1")

	var state DataFluidIoThinRuntimeV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DataFluidIoThinRuntimeV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("data.fluid.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ThinRuntime")

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

func (r *DataFluidIoThinRuntimeV1Alpha1Resource) Read(ctx context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	tflog.Debug(ctx, "Read resource k8s_data_fluid_io_thin_runtime_v1alpha1")
	// NO-OP: All data is already in Terraform state
}

func (r *DataFluidIoThinRuntimeV1Alpha1Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Update resource k8s_data_fluid_io_thin_runtime_v1alpha1")

	var state DataFluidIoThinRuntimeV1Alpha1TerraformModel
	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	var goModel DataFluidIoThinRuntimeV1Alpha1GoModel
	diags = req.Config.Get(ctx, &goModel)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	goModel.ApiVersion = utilities.Ptr("data.fluid.io/v1alpha1")
	goModel.Kind = utilities.Ptr("ThinRuntime")

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

func (r *DataFluidIoThinRuntimeV1Alpha1Resource) Delete(ctx context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	tflog.Debug(ctx, "Delete resource k8s_data_fluid_io_thin_runtime_v1alpha1")
	// NO-OP: Terraform removes the state automatically for us
}
